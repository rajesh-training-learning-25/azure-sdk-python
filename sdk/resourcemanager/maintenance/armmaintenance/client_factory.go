//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmaintenance

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
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
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

// NewApplyUpdateForResourceGroupClient creates a new instance of ApplyUpdateForResourceGroupClient.
func (c *ClientFactory) NewApplyUpdateForResourceGroupClient() *ApplyUpdateForResourceGroupClient {
	return &ApplyUpdateForResourceGroupClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewApplyUpdatesClient creates a new instance of ApplyUpdatesClient.
func (c *ClientFactory) NewApplyUpdatesClient() *ApplyUpdatesClient {
	return &ApplyUpdatesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewConfigurationAssignmentsClient creates a new instance of ConfigurationAssignmentsClient.
func (c *ClientFactory) NewConfigurationAssignmentsClient() *ConfigurationAssignmentsClient {
	return &ConfigurationAssignmentsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewConfigurationAssignmentsForResourceGroupClient creates a new instance of ConfigurationAssignmentsForResourceGroupClient.
func (c *ClientFactory) NewConfigurationAssignmentsForResourceGroupClient() *ConfigurationAssignmentsForResourceGroupClient {
	return &ConfigurationAssignmentsForResourceGroupClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewConfigurationAssignmentsForSubscriptionsClient creates a new instance of ConfigurationAssignmentsForSubscriptionsClient.
func (c *ClientFactory) NewConfigurationAssignmentsForSubscriptionsClient() *ConfigurationAssignmentsForSubscriptionsClient {
	return &ConfigurationAssignmentsForSubscriptionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewConfigurationAssignmentsWithinSubscriptionClient creates a new instance of ConfigurationAssignmentsWithinSubscriptionClient.
func (c *ClientFactory) NewConfigurationAssignmentsWithinSubscriptionClient() *ConfigurationAssignmentsWithinSubscriptionClient {
	return &ConfigurationAssignmentsWithinSubscriptionClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewConfigurationsClient creates a new instance of ConfigurationsClient.
func (c *ClientFactory) NewConfigurationsClient() *ConfigurationsClient {
	return &ConfigurationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewConfigurationsForResourceGroupClient creates a new instance of ConfigurationsForResourceGroupClient.
func (c *ClientFactory) NewConfigurationsForResourceGroupClient() *ConfigurationsForResourceGroupClient {
	return &ConfigurationsForResourceGroupClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		internal: c.internal,
	}
}

// NewPublicMaintenanceConfigurationsClient creates a new instance of PublicMaintenanceConfigurationsClient.
func (c *ClientFactory) NewPublicMaintenanceConfigurationsClient() *PublicMaintenanceConfigurationsClient {
	return &PublicMaintenanceConfigurationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewScheduledEventClient creates a new instance of ScheduledEventClient.
func (c *ClientFactory) NewScheduledEventClient() *ScheduledEventClient {
	return &ScheduledEventClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewUpdatesClient creates a new instance of UpdatesClient.
func (c *ClientFactory) NewUpdatesClient() *UpdatesClient {
	return &UpdatesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
