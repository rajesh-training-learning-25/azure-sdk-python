//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcehealth

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	internal       *arm.Client
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	internal, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID,
		internal:       internal,
	}, nil
}

// NewAvailabilityStatusesClient creates a new instance of AvailabilityStatusesClient.
func (c *ClientFactory) NewAvailabilityStatusesClient() *AvailabilityStatusesClient {
	return &AvailabilityStatusesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewChildAvailabilityStatusesClient creates a new instance of ChildAvailabilityStatusesClient.
func (c *ClientFactory) NewChildAvailabilityStatusesClient() *ChildAvailabilityStatusesClient {
	return &ChildAvailabilityStatusesClient{
		internal: c.internal,
	}
}

// NewChildResourcesClient creates a new instance of ChildResourcesClient.
func (c *ClientFactory) NewChildResourcesClient() *ChildResourcesClient {
	return &ChildResourcesClient{
		internal: c.internal,
	}
}

// NewEmergingIssuesClient creates a new instance of EmergingIssuesClient.
func (c *ClientFactory) NewEmergingIssuesClient() *EmergingIssuesClient {
	return &EmergingIssuesClient{
		internal: c.internal,
	}
}

// NewEventClient creates a new instance of EventClient.
func (c *ClientFactory) NewEventClient() *EventClient {
	return &EventClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewEventsClient creates a new instance of EventsClient.
func (c *ClientFactory) NewEventsClient() *EventsClient {
	return &EventsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewImpactedResourcesClient creates a new instance of ImpactedResourcesClient.
func (c *ClientFactory) NewImpactedResourcesClient() *ImpactedResourcesClient {
	return &ImpactedResourcesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewMetadataClient creates a new instance of MetadataClient.
func (c *ClientFactory) NewMetadataClient() *MetadataClient {
	return &MetadataClient{
		internal: c.internal,
	}
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		internal: c.internal,
	}
}

// NewSecurityAdvisoryImpactedResourcesClient creates a new instance of SecurityAdvisoryImpactedResourcesClient.
func (c *ClientFactory) NewSecurityAdvisoryImpactedResourcesClient() *SecurityAdvisoryImpactedResourcesClient {
	return &SecurityAdvisoryImpactedResourcesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
