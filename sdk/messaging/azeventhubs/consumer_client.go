// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
package azeventhubs

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/log"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/internal/amqpwrap"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/internal/conn"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/internal/go-amqp"
)

// DefaultConsumerGroup is the name of the default consumer group in the Event Hubs service.
const DefaultConsumerGroup = "$Default"

// ConsumerClientOptions contains options for the `NewConsumerClient` and `NewConsumerClientFromConnectionString`
// functions.
type ConsumerClientOptions struct {
	// TLSConfig configures a client with a custom *tls.Config.
	TLSConfig *tls.Config

	// Application ID that will be passed to the namespace.
	ApplicationID string

	// NewWebSocketConn is a function that can create a net.Conn for use with websockets.
	// For an example, see ExampleNewClient_usingWebsockets() function in example_client_test.go.
	NewWebSocketConn func(ctx context.Context, args NewWebSocketConnArgs) (net.Conn, error)

	// RetryOptions controls how often operations are retried from this client and any
	// Receivers and Senders created from this client.
	RetryOptions RetryOptions

	// StartPosition is the position we will start receiving events from,
	// either an offset (inclusive) with Offset, or receiving events received
	// after a specific time using EnqueuedTime.
	StartPosition StartPosition

	// OwnerLevel is the priority for this consumer, also known as the 'epoch' level.
	// When used, a consumer with a higher OwnerLevel will take ownership of a partition
	// from consumers with a lower OwnerLevel.
	// Default is off.
	OwnerLevel *uint64
}

type StartPosition struct {
	// Offset will start the consumer after the specified offset. Can be exclusive
	// or inclusive, based on the Inclusive property.
	// NOTE: offsets are not stable values, and might refer to different events over time
	// as the Event Hub events reach their age limit and are discarded.
	Offset *uint64

	// SequenceNumber will start the consumer after the specified sequence number. Can be exclusive
	// or inclusive, based on the Inclusive property.
	SequenceNumber *int64

	// EnqueuedTime will start the consumer before events that were enqueued on or after EnqueuedTime.
	// Can be exclusive or inclusive, based on the Inclusive property.
	EnqueuedTime *time.Time

	// Inclusive configures whether the events directly at Offset, SequenceNumber or EnqueuedTime will be included (true)
	// or excluded (false).
	Inclusive bool

	// Earliest will start the consumer at the earliest event.
	Earliest *bool

	// Latest will start the consumer after the last event.
	Latest *bool
}

// ConsumerClient is used to receive events from an Event Hub partition.
type ConsumerClient struct {
	retryOptions  RetryOptions
	namespace     *internal.Namespace
	eventHub      string
	consumerGroup string
	partitionID   string
	ownerLevel    *uint64

	offsetExpression string

	links *internal.Links[amqpwrap.AMQPReceiverCloser]
}

// NewConsumerClient creates a ConsumerClient which uses an azcore.TokenCredential for authentication.
// The consumerGroup is the consumer group for this consumer.
// The fullyQualifiedNamespace is the Event Hubs namespace name (ex: myeventhub.servicebus.windows.net)
// The credential is one of the credentials in the `github.com/Azure/azure-sdk-for-go/sdk/azidentity` package.
func NewConsumerClient(consumerGroup string, fullyQualifiedNamespace string, eventHub string, partitionID string, credential azcore.TokenCredential, options *ConsumerClientOptions) (*ConsumerClient, error) {
	return newConsumerClientImpl(consumerClientArgs{
		fullyQualifiedNamespace: fullyQualifiedNamespace,
		credential:              credential,
		eventHub:                eventHub,
		partitionID:             partitionID,
		consumerGroup:           consumerGroup,
	}, options)
}

// NewConsumerClientFromConnectionString creates a ConsumerClient with a connection string that contains an EntityPath value.
// The consumerGroup is the consumer group for this consumer.
// ex: Endpoint=sb://<your-namespace>.servicebus.windows.net/;SharedAccessKeyName=<key-name>;SharedAccessKey=<key>;EntityPath=<entity path>
func NewConsumerClientFromConnectionString(consumerGroup string, connectionString string, partitionID string, options *ConsumerClientOptions) (*ConsumerClient, error) {
	parsedConn, err := conn.ParsedConnectionFromStr(connectionString)

	if err != nil {
		return nil, err
	}

	if parsedConn.HubName == "" {
		return nil, errors.New("Connection string needs to contain an EntityPath. Use NewConsumerClientForHubFromConnectionString to specify eventHub as a parameter.")
	}

	return newConsumerClientImpl(consumerClientArgs{
		connectionString: connectionString,
		// eventHub will come from the connection string itself.
		partitionID:   partitionID,
		consumerGroup: consumerGroup,
	}, options)
}

// NewConsumerClientFromConnectionString creates a ConsumerClient with a connection string that does not have an entity path value.
// The consumerGroup is the consumer group for this consumer.
// ex: Endpoint=sb://<your-namespace>.servicebus.windows.net/;SharedAccessKeyName=<key-name>;SharedAccessKey=<key>
func NewConsumerClientForHubFromConnectionString(consumerGroup string, connectionString string, eventHub string, partitionID string, options *ConsumerClientOptions) (*ConsumerClient, error) {
	return newConsumerClientImpl(consumerClientArgs{
		connectionString: connectionString,
		eventHub:         eventHub,
		partitionID:      partitionID,
		consumerGroup:    consumerGroup,
	}, options)
}

// ReceiveEventsOptions contains optional parameters for the ReceiveEvents function
type ReceiveEventsOptions struct {
	// For future expansion
}

// ReceiveEvents receives events until the context has expired or been cancelled.
func (cc *ConsumerClient) ReceiveEvents(ctx context.Context, count int, options *ReceiveEventsOptions) ([]*ReceivedEventData, error) {
	var events []*ReceivedEventData

	err := cc.links.Retry(ctx, EventConsumer, "ReceiveEvents", cc.partitionID, cc.retryOptions, func(ctx context.Context, lwid internal.LinkWithID[amqpwrap.AMQPReceiverCloser]) error {
		events = nil

		outstandingCredits := lwid.Link.Credits()

		if count > int(outstandingCredits) {
			newCredits := uint32(count) - outstandingCredits

			log.Writef(EventConsumer, "Have %d outstanding credit, only issuing %d credits", outstandingCredits, newCredits)

			if err := lwid.Link.IssueCredit(newCredits); err != nil {
				return err
			}
		}

		for {
			amqpMessage, err := lwid.Link.Receive(ctx)

			if err != nil {
				prefetched := getAllPrefetched(lwid.Link, count-len(events))

				for _, amqpMsg := range prefetched {
					events = append(events, newReceivedEventData(amqpMsg))
				}

				// this lets cancel errors just return
				return err
			}

			receivedEvent := newReceivedEventData(amqpMessage)
			events = append(events, receivedEvent)

			if len(events) == count {
				return nil
			}
		}
	})

	if err != nil && len(events) == 0 {
		// TODO: if we get a "partition ownership lost" we need to think about whether that's retryable.
		return nil, internal.TransformError(err)
	}

	cc.offsetExpression = getOffsetExpression(StartPosition{
		SequenceNumber: to.Ptr(events[len(events)-1].SequenceNumber),
		Inclusive:      false,
	})

	return events, nil
}

// GetEventHubProperties gets event hub properties, like the available partition IDs and when the Event Hub was created.
func (cc *ConsumerClient) GetEventHubProperties(ctx context.Context, options *GetEventHubPropertiesOptions) (EventHubProperties, error) {
	rpcLink, err := cc.links.GetManagementLink(ctx)

	if err != nil {
		return EventHubProperties{}, err
	}

	return getEventHubProperties(ctx, cc.namespace, rpcLink.Link, cc.eventHub, options)
}

// GetPartitionProperties gets properties for a specific partition. This includes data like the last enqueued sequence number, the first sequence
// number and when an event was last enqueued to the partition.
func (cc *ConsumerClient) GetPartitionProperties(ctx context.Context, partitionID string, options *GetPartitionPropertiesOptions) (PartitionProperties, error) {
	rpcLink, err := cc.links.GetManagementLink(ctx)

	if err != nil {
		return PartitionProperties{}, err
	}

	return getPartitionProperties(ctx, cc.namespace, rpcLink.Link, cc.eventHub, partitionID, options)
}

// Close closes the consumer's link and the underlying AMQP connection.
func (cc *ConsumerClient) Close(ctx context.Context) error {
	if err := cc.links.Close(ctx); err != nil {
		log.Writef(EventConsumer, "Failed to close link (error might be cached): %s", err.Error())
	}
	return cc.namespace.Close(ctx, true)
}

func getOffsetExpression(startPosition StartPosition) string {
	lt := ">"

	if startPosition.Inclusive {
		lt = ">="
	}

	if startPosition.EnqueuedTime != nil {
		// time-based, non-inclusive
		return fmt.Sprintf("amqp.annotation.x-opt-enqueued-time %s '%v'", lt, startPosition.EnqueuedTime.UnixNano()/int64(time.Millisecond))
	}

	if startPosition.Offset != nil {
		// offset-based, non-inclusive
		// ex: amqp.annotation.x-opt-enqueued-time %s '165805323000'
		return fmt.Sprintf("amqp.annotation.x-opt-offset %s '%v'", lt, startPosition.Offset)
	}

	if startPosition.Latest != nil && *startPosition.Latest {
		return "amqp.annotation.x-opt-offset > '@latest'"
	}

	if startPosition.SequenceNumber != nil {
		return fmt.Sprintf("amqp.annotation.x-opt-sequence-number %s '%d", lt, *startPosition.SequenceNumber)
	}

	// default to the start
	return "amqp.annotation.x-opt-offset > '-1'"
}

func (cc *ConsumerClient) getEntityPath(partitionID string) string {
	if cc.ownerLevel == nil {
		return fmt.Sprintf("%s/ConsumerGroups/%s/Partitions/%s", cc.eventHub, cc.consumerGroup, partitionID)
	} else {
		return fmt.Sprintf("%s/ConsumerGroups/%s/Partitions/%s/epoch/%d", cc.eventHub, cc.consumerGroup, partitionID, *cc.ownerLevel)
	}
}

const defaultLinkRxBuffer = 2048

func (cc *ConsumerClient) newEventHubConsumerLink(ctx context.Context, session amqpwrap.AMQPSession, entityPath string) (internal.AMQPReceiverCloser, error) {
	receiver, err := session.NewReceiver(ctx, entityPath, &amqp.ReceiverOptions{
		SettlementMode: to.Ptr(amqp.ModeFirst),
		ManualCredits:  true,
		Credit:         defaultLinkRxBuffer,
		Filters: []amqp.LinkFilter{
			amqp.LinkFilterSelector(cc.offsetExpression),
		},
	})

	if err != nil {
		return nil, err
	}

	return receiver, nil
}

type consumerClientArgs struct {
	connectionString string

	// the Event Hubs namespace name (ex: myservicebus.servicebus.windows.net)
	fullyQualifiedNamespace string
	credential              azcore.TokenCredential

	eventHub    string
	partitionID string

	consumerGroup string
}

func newConsumerClientImpl(args consumerClientArgs, options *ConsumerClientOptions) (*ConsumerClient, error) {
	if options == nil {
		options = &ConsumerClientOptions{}
	}

	client := &ConsumerClient{
		eventHub:         args.eventHub,
		partitionID:      args.partitionID,
		ownerLevel:       options.OwnerLevel,
		consumerGroup:    args.consumerGroup,
		offsetExpression: getOffsetExpression(options.StartPosition),
	}

	var err error
	var nsOptions []internal.NamespaceOption

	if args.connectionString != "" {
		nsOptions = append(nsOptions, internal.NamespaceWithConnectionString(args.connectionString))
	} else if args.credential != nil {
		option := internal.NamespaceWithTokenCredential(
			args.fullyQualifiedNamespace,
			args.credential)

		nsOptions = append(nsOptions, option)
	}

	client.retryOptions = options.RetryOptions

	if options.TLSConfig != nil {
		nsOptions = append(nsOptions, internal.NamespaceWithTLSConfig(options.TLSConfig))
	}

	if options.NewWebSocketConn != nil {
		nsOptions = append(nsOptions, internal.NamespaceWithWebSocket(options.NewWebSocketConn))
	}

	if options.ApplicationID != "" {
		nsOptions = append(nsOptions, internal.NamespaceWithUserAgent(options.ApplicationID))
	}

	nsOptions = append(nsOptions, internal.NamespaceWithRetryOptions(options.RetryOptions))

	tempNS, err := internal.NewNamespace(nsOptions...)

	if err != nil {
		return nil, err
	}

	client.namespace = tempNS
	client.links = internal.NewLinks(tempNS, fmt.Sprintf("%s/$management", client.eventHub), client.getEntityPath, client.newEventHubConsumerLink)

	return client, nil
}

func getAllPrefetched(receiver amqpwrap.AMQPReceiver, max int) []*amqp.Message {
	var messages []*amqp.Message

	for i := 0; i < max; i++ {
		msg := receiver.Prefetched()

		if msg == nil {
			break
		}

		messages = append(messages, msg)
	}

	return messages
}
