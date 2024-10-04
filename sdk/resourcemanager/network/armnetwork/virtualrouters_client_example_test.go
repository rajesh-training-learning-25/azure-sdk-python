//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4883fa5dbf6f2c9093fac8ce334547e9dfac68fa/specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/VirtualRouterDelete.json
func ExampleVirtualRoutersClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualRoutersClient().BeginDelete(ctx, "rg1", "virtualRouter", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4883fa5dbf6f2c9093fac8ce334547e9dfac68fa/specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/VirtualRouterGet.json
func ExampleVirtualRoutersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualRoutersClient().Get(ctx, "rg1", "virtualRouter", &armnetwork.VirtualRoutersClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualRouter = armnetwork.VirtualRouter{
	// 	Name: to.Ptr("virtualRouter"),
	// 	Type: to.Ptr("Microsoft.Network/virtualRouters"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualRouters/virtualRouter"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 	},
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.VirtualRouterPropertiesFormat{
	// 		HostedGateway: &armnetwork.SubResource{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vnetGateway"),
	// 		},
	// 		Peerings: []*armnetwork.SubResource{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualRouters/virtualRouter/peerings/peering1"),
	// 		}},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		VirtualRouterAsn: to.Ptr[int64](10000),
	// 		VirtualRouterIPs: []*string{
	// 			to.Ptr("192.168.1.1"),
	// 			to.Ptr("192.168.1.2")},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4883fa5dbf6f2c9093fac8ce334547e9dfac68fa/specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/VirtualRouterPut.json
func ExampleVirtualRoutersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualRoutersClient().BeginCreateOrUpdate(ctx, "rg1", "virtualRouter", armnetwork.VirtualRouter{
		Location: to.Ptr("West US"),
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
		},
		Properties: &armnetwork.VirtualRouterPropertiesFormat{
			HostedGateway: &armnetwork.SubResource{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vnetGateway"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualRouter = armnetwork.VirtualRouter{
	// 	Name: to.Ptr("virtualRouter"),
	// 	Type: to.Ptr("Microsoft.Network/virtualRouters"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualRouters/virtualRouter"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 	},
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.VirtualRouterPropertiesFormat{
	// 		HostedGateway: &armnetwork.SubResource{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vnetGateway"),
	// 		},
	// 		Peerings: []*armnetwork.SubResource{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualRouters/virtualRouter/peerings/peering1"),
	// 		}},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		VirtualRouterAsn: to.Ptr[int64](10000),
	// 		VirtualRouterIPs: []*string{
	// 			to.Ptr("192.168.1.1"),
	// 			to.Ptr("192.168.1.2")},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4883fa5dbf6f2c9093fac8ce334547e9dfac68fa/specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/VirtualRouterListByResourceGroup.json
func ExampleVirtualRoutersClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualRoutersClient().NewListByResourceGroupPager("rg1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.VirtualRouterListResult = armnetwork.VirtualRouterListResult{
		// 	Value: []*armnetwork.VirtualRouter{
		// 		{
		// 			Name: to.Ptr("virtualRouter"),
		// 			Type: to.Ptr("Microsoft.Network/virtualRouters"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualRouters/virtualRouter"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 			},
		// 			Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 			Properties: &armnetwork.VirtualRouterPropertiesFormat{
		// 				HostedGateway: &armnetwork.SubResource{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vnetGateway"),
		// 				},
		// 				Peerings: []*armnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualRouters/virtualRouter/peerings/peering1"),
		// 				}},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				VirtualRouterAsn: to.Ptr[int64](10000),
		// 				VirtualRouterIPs: []*string{
		// 					to.Ptr("192.168.1.1"),
		// 					to.Ptr("192.168.1.2")},
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4883fa5dbf6f2c9093fac8ce334547e9dfac68fa/specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/VirtualRouterListBySubscription.json
func ExampleVirtualRoutersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualRoutersClient().NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.VirtualRouterListResult = armnetwork.VirtualRouterListResult{
		// 	Value: []*armnetwork.VirtualRouter{
		// 		{
		// 			Name: to.Ptr("virtualRouter"),
		// 			Type: to.Ptr("Microsoft.Network/virtualRouters"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualRouters/virtualRouter"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 			},
		// 			Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 			Properties: &armnetwork.VirtualRouterPropertiesFormat{
		// 				HostedGateway: &armnetwork.SubResource{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vnetGateway"),
		// 				},
		// 				Peerings: []*armnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualRouters/virtualRouter/peerings/peering1"),
		// 				}},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				VirtualRouterAsn: to.Ptr[int64](10000),
		// 				VirtualRouterIPs: []*string{
		// 					to.Ptr("192.168.1.1"),
		// 					to.Ptr("192.168.1.2")},
		// 				},
		// 		}},
		// 	}
	}
}
