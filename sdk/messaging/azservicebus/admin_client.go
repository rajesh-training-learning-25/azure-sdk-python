// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azservicebus

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/atom"
	"github.com/Azure/go-autorest/autorest/date"
)

// AdminClient allows you to administer resources in a Service Bus Namespace.
// For example, you can create queues, enabling capabilities like partitioning, duplicate detection, etc..
// NOTE: For sending and receiving messages you'll need to use the `Client` type instead.
type AdminClient struct {
	em *atom.EntityManager
}

type AdminClientOptions struct {
	// for future expansion
}

// NewAdminClient creates an AdminClient authenticating using a connection string.
func NewAdminClientWithConnectionString(connectionString string, options *AdminClientOptions) (*AdminClient, error) {
	em, err := atom.NewEntityManagerWithConnectionString(connectionString, internal.Version)

	if err != nil {
		return nil, err
	}

	return &AdminClient{em: em}, nil
}

// NewAdminClient creates an AdminClient authenticating using a TokenCredential.
func NewAdminClient(fullyQualifiedNamespace string, tokenCredential azcore.TokenCredential, options *AdminClientOptions) (*AdminClient, error) {
	em, err := atom.NewEntityManager(fullyQualifiedNamespace, tokenCredential, internal.Version)

	if err != nil {
		return nil, err
	}

	return &AdminClient{em: em}, nil
}

// AddQueue creates a queue using defaults for all options.
func (ac *AdminClient) AddQueue(ctx context.Context, queueName string) (*QueueProperties, error) {
	return ac.AddQueueWithProperties(ctx, &QueueProperties{
		Name: queueName,
	})
}

// CreateQueue creates a queue with configurable properties.
func (ac *AdminClient) AddQueueWithProperties(ctx context.Context, properties *QueueProperties) (*QueueProperties, error) {
	return ac.createOrUpdateQueueImpl(ctx, properties, true)
}

func (ac *AdminClient) GetQueue(ctx context.Context, queueName string) (*QueueProperties, error) {
	name, desc, err := ac.getQueueImpl(ctx, queueName)

	if err != nil {
		return nil, err
	}

	return newQueueProperties(name, desc)
}

// GetQueueRuntimeProperties gets runtime properties of a queue, like the SizeInBytes, or ActiveMessageCount.
func (ac *AdminClient) GetQueueRuntimeProperties(ctx context.Context, queueName string) (*QueueRuntimeProperties, error) {
	name, desc, err := ac.getQueueImpl(ctx, queueName)

	if err != nil {
		return nil, err
	}

	return &QueueRuntimeProperties{
		Name:                           name,
		SizeInBytes:                    Int64OrZero(desc.SizeInBytes),
		CreatedAt:                      DateTimeToTime(desc.CreatedAt),
		UpdatedAt:                      DateTimeToTime(desc.UpdatedAt),
		AccessedAt:                     DateTimeToTime(desc.AccessedAt),
		TotalMessageCount:              Int64OrZero(desc.MessageCount),
		ActiveMessageCount:             Int32OrZero(desc.CountDetails.ActiveMessageCount),
		DeadLetterMessageCount:         Int32OrZero(desc.CountDetails.DeadLetterMessageCount),
		ScheduledMessageCount:          Int32OrZero(desc.CountDetails.ScheduledMessageCount),
		TransferDeadLetterMessageCount: Int32OrZero(desc.CountDetails.TransferDeadLetterMessageCount),
		TransferMessageCount:           Int32OrZero(desc.CountDetails.TransferMessageCount),
	}, nil
}

func DateTimeToTime(t *date.Time) time.Time {
	if t == nil {
		return time.Time{}
	}

	return t.Time
}

func Int32OrZero(i *int32) int32 {
	if i == nil {
		return 0
	}

	return *i
}

func Int64OrZero(i *int64) int64 {
	if i == nil {
		return 0
	}

	return *i
}

// QueueExists checks if a queue exists.
// Returns true if the queue is found
// (false, nil) if the queue is not found
// (false, err) if an error occurred while trying to check if the queue exists.
func (ac *AdminClient) QueueExists(ctx context.Context, queueName string) (bool, error) {
	_, err := ac.GetQueue(ctx, queueName)

	if err == nil {
		return true, nil
	}

	var httpResponse azcore.HTTPResponse

	if errors.As(err, &httpResponse) && httpResponse.RawResponse().StatusCode == 404 {
		return false, nil
	}

	return false, err
}

// UpdateQueue updates an existing queue.
func (ac *AdminClient) UpdateQueue(ctx context.Context, properties *QueueProperties) (*QueueProperties, error) {
	return ac.createOrUpdateQueueImpl(ctx, properties, false)
}

func (ac *AdminClient) DeleteQueue(ctx context.Context, queueName string) (*http.Response, error) {
	resp, err := ac.em.Delete(ctx, "/"+queueName)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (ac *AdminClient) ListQueues() {

}

func (ac *AdminClient) ListQueuesRuntimeProperties() {

}

// func (ac *AdminClient) GetNamespaceProperties() {}

// func (ac *AdminClient) CreateTopic()                        {}
// func (ac *AdminClient) CreateSubscription()                 {}
// func (ac *AdminClient) CreateRule()                         {}
// func (ac *AdminClient) DeleteTopic()                        {}
// func (ac *AdminClient) DeleteSubscription()                 {}
// func (ac *AdminClient) DeleteRule()                         {}
// func (ac *AdminClient) GetRule()                            {}
// func (ac *AdminClient) GetSubscription()                    {}
// func (ac *AdminClient) GetSubscriptionRuntimeProperties()   {}
// func (ac *AdminClient) GetTopic()                           {}
// func (ac *AdminClient) GetTopicRuntimeProperties()          {}
// func (ac *AdminClient) ListRules()                          {}
// func (ac *AdminClient) ListTopics()                         {}
// func (ac *AdminClient) ListTopicsRuntimeProperties()        {}
// func (ac *AdminClient) ListSubscriptions()                  {}
// func (ac *AdminClient) ListSubscriptionsRuntimeProperties() {}

// func (ac *AdminClient) TopicExists()               {}
// func (ac *AdminClient) SubscriptionExists()        {}
// func (ac *AdminClient) RuleExists()                {}

// func (ac *AdminClient) UpdateTopic()        {}
// func (ac *AdminClient) UpdateSubscription() {}
// func (ac *AdminClient) UpdateRule()         {}
