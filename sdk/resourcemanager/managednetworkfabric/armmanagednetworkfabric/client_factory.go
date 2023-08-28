//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmanagednetworkfabric

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	credential     azcore.TokenCredential
	options        *arm.ClientOptions
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName+".ClientFactory", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, credential: credential,
		options: options.Clone(),
	}, nil
}

func (c *ClientFactory) NewAccessControlListsClient() *AccessControlListsClient {
	subClient, _ := NewAccessControlListsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewInternetGatewaysClient() *InternetGatewaysClient {
	subClient, _ := NewInternetGatewaysClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewInternetGatewayRulesClient() *InternetGatewayRulesClient {
	subClient, _ := NewInternetGatewayRulesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewIPCommunitiesClient() *IPCommunitiesClient {
	subClient, _ := NewIPCommunitiesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewIPExtendedCommunitiesClient() *IPExtendedCommunitiesClient {
	subClient, _ := NewIPExtendedCommunitiesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewIPPrefixesClient() *IPPrefixesClient {
	subClient, _ := NewIPPrefixesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewL2IsolationDomainsClient() *L2IsolationDomainsClient {
	subClient, _ := NewL2IsolationDomainsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewL3IsolationDomainsClient() *L3IsolationDomainsClient {
	subClient, _ := NewL3IsolationDomainsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewInternalNetworksClient() *InternalNetworksClient {
	subClient, _ := NewInternalNetworksClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewExternalNetworksClient() *ExternalNetworksClient {
	subClient, _ := NewExternalNetworksClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewNeighborGroupsClient() *NeighborGroupsClient {
	subClient, _ := NewNeighborGroupsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewNetworkDeviceSKUsClient() *NetworkDeviceSKUsClient {
	subClient, _ := NewNetworkDeviceSKUsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewNetworkDevicesClient() *NetworkDevicesClient {
	subClient, _ := NewNetworkDevicesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewNetworkInterfacesClient() *NetworkInterfacesClient {
	subClient, _ := NewNetworkInterfacesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewNetworkFabricControllersClient() *NetworkFabricControllersClient {
	subClient, _ := NewNetworkFabricControllersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewNetworkFabricSKUsClient() *NetworkFabricSKUsClient {
	subClient, _ := NewNetworkFabricSKUsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewNetworkFabricsClient() *NetworkFabricsClient {
	subClient, _ := NewNetworkFabricsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewNetworkToNetworkInterconnectsClient() *NetworkToNetworkInterconnectsClient {
	subClient, _ := NewNetworkToNetworkInterconnectsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewNetworkPacketBrokersClient() *NetworkPacketBrokersClient {
	subClient, _ := NewNetworkPacketBrokersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewNetworkRacksClient() *NetworkRacksClient {
	subClient, _ := NewNetworkRacksClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewNetworkTapRulesClient() *NetworkTapRulesClient {
	subClient, _ := NewNetworkTapRulesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewNetworkTapsClient() *NetworkTapsClient {
	subClient, _ := NewNetworkTapsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRoutePoliciesClient() *RoutePoliciesClient {
	subClient, _ := NewRoutePoliciesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
