// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
package azeventhubs

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	azlog "github.com/Azure/azure-sdk-for-go/sdk/internal/log"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/internal/amqpwrap"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/internal/conn"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/internal/exported"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/internal/go-amqp"
)

// NewWebSocketConnArgs are passed to your web socket creation function (ClientOptions.NewWebSocketConn)
type NewWebSocketConnArgs = exported.NewWebSocketConnArgs

// RetryOptions represent the options for retries.
type RetryOptions = exported.RetryOptions

// ProducerClientOptions contains options for the `NewProducerClient` and `NewProducerClientFromConnectionString`
// functions.
type ProducerClientOptions struct {
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
}

type ProducerClient struct {
	retryOptions RetryOptions
	namespace    internal.NamespaceForProducerOrConsumer
	eventHub     string

	links *internal.Links[amqpwrap.AMQPSenderCloser]
}

// anyPartitionID is what we target if we want to send a message and let Event Hubs pick a partition
// or if we're doing an operation that isn't partition specific, such as querying the management link
// to get event hub properties or partition properties.
const anyPartitionID = ""

// NewProducerClient creates a ProducerClient which uses an azcore.TokenCredential for authentication.
// The fullyQualifiedNamespace is the Event Hubs namespace name (ex: myeventhub.servicebus.windows.net)
// The credential is one of the credentials in the `github.com/Azure/azure-sdk-for-go/sdk/azidentity` package.
func NewProducerClient(fullyQualifiedNamespace string, eventHub string, credential azcore.TokenCredential, options *ProducerClientOptions) (*ProducerClient, error) {
	return newProducerClientImpl(producerClientCreds{
		fullyQualifiedNamespace: fullyQualifiedNamespace,
		credential:              credential,
		eventHub:                eventHub,
	}, options)
}

// NewProducerClientFromConnectionString creates a ProducerClient with a connection string that contains an EntityPath value.
// ex: Endpoint=sb://<your-namespace>.servicebus.windows.net/;SharedAccessKeyName=<key-name>;SharedAccessKey=<key>;EntityPath=<entity path>
func NewProducerClientFromConnectionString(connectionString string, options *ProducerClientOptions) (*ProducerClient, error) {
	parsedConn, err := conn.ParsedConnectionFromStr(connectionString)

	if err != nil {
		return nil, err
	}

	if parsedConn.HubName == "" {
		return nil, errors.New("Connection string needs to contain an EntityPath. Use NewProducerClientForHubFromConnectionString to specify eventHub as a parameter.")
	}

	return newProducerClientImpl(producerClientCreds{
		connectionString: connectionString,
		// eventHub will come from the connection string itself.
	}, options)
}

// NewProducerClientForHubFromConnectionString creates a ProducerClient with a connection string that does not have an entity path value.
// ex: Endpoint=sb://<your-namespace>.servicebus.windows.net/;SharedAccessKeyName=<key-name>;SharedAccessKey=<key>
func NewProducerClientForHubFromConnectionString(connectionString string, eventHub string, options *ProducerClientOptions) (*ProducerClient, error) {
	return newProducerClientImpl(producerClientCreds{
		connectionString: connectionString,
		eventHub:         eventHub,
	}, options)
}

// NewEventDataBatchOptions contains optional parameters for the NewEventDataBatch function
type NewEventDataBatchOptions struct {
	// MaxBytes overrides the max size (in bytes) for a batch.
	// By default NewMessageBatch will use the max message size provided by the service.
	MaxBytes uint64

	// PartitionKey is hashed to calculate the partition assignment. Messages and message
	// batches with the same PartitionKey are guaranteed to end up in the same partition.
	// Note that if you use this option then PartitionID cannot be set.
	PartitionKey *string

	// PartitionID is the ID of the partition to send these messages to.
	// Note that if you use this option then PartitionKey cannot be set.
	PartitionID *string
}

func (pc *ProducerClient) NewEventDataBatch(ctx context.Context, options *NewEventDataBatchOptions) (*EventDataBatch, error) {
	if options == nil {
		options = &NewEventDataBatchOptions{}
	}

	if options.PartitionID != nil && options.PartitionKey != nil {
		return nil, errors.New("either PartitionID or PartitionKey can be set, but not both")
	}

	var batch EventDataBatch

	if options.PartitionID != nil {
		// they want to send to a particular partition. The batch size should be the same for any
		// link but we might as well use the one they're going to send to.
		batch.partitionID = options.PartitionID
	} else if options.PartitionKey != nil {
		batch.partitionKey = options.PartitionKey
	}

	partitionID := anyPartitionID

	if batch.partitionID != nil {
		partitionID = *batch.partitionID
	}

	err := pc.links.Retry(ctx, exported.EventProducer, "NewEventDataBatch", partitionID, pc.retryOptions, func(ctx context.Context, lwid internal.LinkWithID[amqpwrap.AMQPSenderCloser]) error {
		if options.MaxBytes == 0 {
			batch.maxBytes = lwid.Link.MaxMessageSize()
			return nil
		}

		if options.MaxBytes > lwid.Link.MaxMessageSize() {
			return internal.NewErrNonRetriable(fmt.Sprintf("maximum message size for batch was set to %d bytes, which is larger than the maximum size allowed by link (%d)", options.MaxBytes, lwid.Link.MaxMessageSize()))
		}

		batch.maxBytes = options.MaxBytes
		return nil
	})

	if err != nil {
		return nil, internal.TransformError(err)
	}

	return &batch, nil
}

// SendEventsOptions are options for the SendEvents function.
type SendEventsOptions struct {
	// PartitionID is the intended partition for the messages
	PartitionID *string
}

// SendEventBatchOptions contains optional parameters for the SendEventBatch function
type SendEventBatchOptions struct {
	// For future expansion
}

func (pc *ProducerClient) SendEventBatch(ctx context.Context, batch *EventDataBatch, options *SendEventBatchOptions) error {
	err := pc.links.Retry(ctx, exported.EventProducer, "SendEventBatch", *batch.partitionID, pc.retryOptions, func(ctx context.Context, lwid internal.LinkWithID[amqpwrap.AMQPSenderCloser]) error {
		return lwid.Link.Send(ctx, batch.toAMQPMessage())
	})
	return internal.TransformError(err)
}

// GetPartitionProperties gets properties for a specific partition. This includes data like the last enqueued sequence number, the first sequence
// number and when an event was last enqueued to the partition.
func (pc *ProducerClient) GetPartitionProperties(ctx context.Context, partitionID string, options *GetPartitionPropertiesOptions) (PartitionProperties, error) {
	rpcLink, err := pc.links.GetManagementLink(ctx)

	if err != nil {
		return PartitionProperties{}, err
	}

	return getPartitionProperties(ctx, pc.namespace, rpcLink.Link, pc.eventHub, partitionID, options)
}

// GetEventHubProperties gets event hub properties, like the available partition IDs and when the Event Hub was created.
func (pc *ProducerClient) GetEventHubProperties(ctx context.Context, options *GetEventHubPropertiesOptions) (EventHubProperties, error) {
	rpcLink, err := pc.links.GetManagementLink(ctx)

	if err != nil {
		return EventHubProperties{}, err
	}

	return getEventHubProperties(ctx, pc.namespace, rpcLink.Link, pc.eventHub, options)
}

// Close closes the producer's AMQP links and the underlying AMQP connection.
func (pc *ProducerClient) Close(ctx context.Context) error {
	if err := pc.links.Close(ctx); err != nil {
		azlog.Writef(EventProducer, "Failed when closing links while shutting down producer client: %s", err.Error())
	}
	return pc.namespace.Close(ctx, true)
}

func (pc *ProducerClient) getEntityPath(partitionID string) string {
	if partitionID != anyPartitionID {
		return fmt.Sprintf("%s/Partitions/%s", pc.eventHub, partitionID)
	} else {
		// this is the "let Event Hubs" decide link - any sends that occur here will
		// end up getting distributed to different partitions on the service side, rather
		// then being specified in the client.
		return pc.eventHub
	}
}

func (pc *ProducerClient) newEventHubProducerLink(ctx context.Context, session amqpwrap.AMQPSession, entityPath string) (amqpwrap.AMQPSenderCloser, error) {
	sender, err := session.NewSender(ctx, entityPath, &amqp.SenderOptions{
		SettlementMode:              to.Ptr(amqp.ModeMixed),
		RequestedReceiverSettleMode: to.Ptr(amqp.ModeFirst),
		IgnoreDispositionErrors:     true,
	})

	if err != nil {
		return nil, err
	}

	return sender, nil
}

type producerClientCreds struct {
	connectionString string

	// the Event Hubs namespace name (ex: myservicebus.servicebus.windows.net)
	fullyQualifiedNamespace string
	credential              azcore.TokenCredential

	eventHub string
}

func newProducerClientImpl(creds producerClientCreds, options *ProducerClientOptions) (*ProducerClient, error) {
	client := &ProducerClient{
		eventHub: creds.eventHub,
	}

	var err error
	var nsOptions []internal.NamespaceOption

	if creds.connectionString != "" {
		nsOptions = append(nsOptions, internal.NamespaceWithConnectionString(creds.connectionString))
	} else if creds.credential != nil {
		option := internal.NamespaceWithTokenCredential(
			creds.fullyQualifiedNamespace,
			creds.credential)

		nsOptions = append(nsOptions, option)
	}

	if options != nil {
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
	}

	tmpNS, err := internal.NewNamespace(nsOptions...)

	if err != nil {
		return nil, err
	}

	client.namespace = tmpNS

	client.links = internal.NewLinks(tmpNS, fmt.Sprintf("%s/$management", client.eventHub), client.getEntityPath, client.newEventHubProducerLink)

	return client, err
}
