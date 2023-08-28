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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/baac183ffa684d94f697f0fc6f480e02cfb00f3d/specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/VirtualNetworkTapDelete.json
func ExampleVirtualNetworkTapsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualNetworkTapsClient().BeginDelete(ctx, "rg1", "test-vtap", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/baac183ffa684d94f697f0fc6f480e02cfb00f3d/specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/VirtualNetworkTapGet.json
func ExampleVirtualNetworkTapsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualNetworkTapsClient().Get(ctx, "rg1", "testvtap", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualNetworkTap = armnetwork.VirtualNetworkTap{
	// 	Name: to.Ptr("testvtap"),
	// 	Type: to.Ptr("Microsoft.Network/virtualNetworkTaps"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkTaps/testvtap"),
	// 	Location: to.Ptr("centraluseuap"),
	// 	Etag: to.Ptr("etag"),
	// 	Properties: &armnetwork.VirtualNetworkTapPropertiesFormat{
	// 		DestinationNetworkInterfaceIPConfiguration: &armnetwork.InterfaceIPConfiguration{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/testNetworkInterface/ipConfigurations/testIPConfig1"),
	// 		},
	// 		DestinationPort: to.Ptr[int32](4789),
	// 		NetworkInterfaceTapConfigurations: []*armnetwork.InterfaceTapConfiguration{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/testNetworkInterface2/tapConfigurations/testtapConfiguration"),
	// 		}},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		ResourceGUID: to.Ptr("6A7C139D-8B8D-499B-B7CB-4F3F02A8A44F"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/baac183ffa684d94f697f0fc6f480e02cfb00f3d/specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/VirtualNetworkTapCreate.json
func ExampleVirtualNetworkTapsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualNetworkTapsClient().BeginCreateOrUpdate(ctx, "rg1", "test-vtap", armnetwork.VirtualNetworkTap{
		Location: to.Ptr("centraluseuap"),
		Properties: &armnetwork.VirtualNetworkTapPropertiesFormat{
			DestinationNetworkInterfaceIPConfiguration: &armnetwork.InterfaceIPConfiguration{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/testNetworkInterface/ipConfigurations/ipconfig1"),
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
	// res.VirtualNetworkTap = armnetwork.VirtualNetworkTap{
	// 	Name: to.Ptr("testvtap"),
	// 	Type: to.Ptr("Microsoft.Network/virtualNetworkTaps"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkTaps/testvtap"),
	// 	Location: to.Ptr("centraluseuap"),
	// 	Etag: to.Ptr("etag"),
	// 	Properties: &armnetwork.VirtualNetworkTapPropertiesFormat{
	// 		DestinationNetworkInterfaceIPConfiguration: &armnetwork.InterfaceIPConfiguration{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/testNetworkInterface/ipConfigurations/testIPConfig1"),
	// 		},
	// 		DestinationPort: to.Ptr[int32](4789),
	// 		NetworkInterfaceTapConfigurations: []*armnetwork.InterfaceTapConfiguration{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/testNetworkInterface2/tapConfigurations/testtapConfiguration"),
	// 		}},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		ResourceGUID: to.Ptr("6A7C139D-8B8D-499B-B7CB-4F3F02A8A44F"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/baac183ffa684d94f697f0fc6f480e02cfb00f3d/specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/VirtualNetworkTapUpdateTags.json
func ExampleVirtualNetworkTapsClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualNetworkTapsClient().UpdateTags(ctx, "rg1", "test-vtap", armnetwork.TagsObject{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualNetworkTap = armnetwork.VirtualNetworkTap{
	// 	Name: to.Ptr("test-vtap"),
	// 	Type: to.Ptr("Microsoft.Network/virtualNetworkTaps"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkTaps/test-vtap"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armnetwork.VirtualNetworkTapPropertiesFormat{
	// 		DestinationNetworkInterfaceIPConfiguration: &armnetwork.InterfaceIPConfiguration{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/testNetworkInterface/ipConfigurations/testIPConfig1"),
	// 		},
	// 		DestinationPort: to.Ptr[int32](4789),
	// 		NetworkInterfaceTapConfigurations: []*armnetwork.InterfaceTapConfiguration{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/testNetworkInterface2/tapConfigurations/testtapConfiguration"),
	// 		}},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/baac183ffa684d94f697f0fc6f480e02cfb00f3d/specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/VirtualNetworkTapListAll.json
func ExampleVirtualNetworkTapsClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualNetworkTapsClient().NewListAllPager(nil)
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
		// page.VirtualNetworkTapListResult = armnetwork.VirtualNetworkTapListResult{
		// 	Value: []*armnetwork.VirtualNetworkTap{
		// 		{
		// 			Name: to.Ptr("testvtap"),
		// 			Type: to.Ptr("Microsoft.Network/virtualNetworkTaps"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkTaps/testvtap"),
		// 			Location: to.Ptr("centraluseuap"),
		// 			Etag: to.Ptr("etag"),
		// 			Properties: &armnetwork.VirtualNetworkTapPropertiesFormat{
		// 				DestinationNetworkInterfaceIPConfiguration: &armnetwork.InterfaceIPConfiguration{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/testNetworkInterface/ipConfigurations/testIPConfig1"),
		// 				},
		// 				DestinationPort: to.Ptr[int32](4789),
		// 				NetworkInterfaceTapConfigurations: []*armnetwork.InterfaceTapConfiguration{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/testNetworkInterface2/tapConfigurations/testtapConfiguration"),
		// 				}},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				ResourceGUID: to.Ptr("6A7C139D-8B8D-499B-B7CB-4F3F02A8A44F"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testvtap2"),
		// 			Type: to.Ptr("Microsoft.Network/virtualNetworkTaps"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkTaps/testvtap2"),
		// 			Location: to.Ptr("centraluseuap"),
		// 			Etag: to.Ptr("etag"),
		// 			Properties: &armnetwork.VirtualNetworkTapPropertiesFormat{
		// 				DestinationNetworkInterfaceIPConfiguration: &armnetwork.InterfaceIPConfiguration{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/testNetworkInterface/ipConfigurations/testIPConfig1"),
		// 				},
		// 				DestinationPort: to.Ptr[int32](4789),
		// 				NetworkInterfaceTapConfigurations: []*armnetwork.InterfaceTapConfiguration{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/testNetworkInterface3/tapConfigurations/testtapConfiguration"),
		// 				}},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				ResourceGUID: to.Ptr("6A7C139D-8B8D-499B-B7CB-4F3F02A8A44F"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/baac183ffa684d94f697f0fc6f480e02cfb00f3d/specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/VirtualNetworkTapList.json
func ExampleVirtualNetworkTapsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualNetworkTapsClient().NewListByResourceGroupPager("rg1", nil)
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
		// page.VirtualNetworkTapListResult = armnetwork.VirtualNetworkTapListResult{
		// 	Value: []*armnetwork.VirtualNetworkTap{
		// 		{
		// 			Name: to.Ptr("testvtap"),
		// 			Type: to.Ptr("Microsoft.Network/virtualNetworkTaps"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkTaps/testvtap"),
		// 			Location: to.Ptr("centraluseuap"),
		// 			Etag: to.Ptr("etag"),
		// 			Properties: &armnetwork.VirtualNetworkTapPropertiesFormat{
		// 				DestinationNetworkInterfaceIPConfiguration: &armnetwork.InterfaceIPConfiguration{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/testNetworkInterface/ipConfigurations/testIPConfig1"),
		// 				},
		// 				DestinationPort: to.Ptr[int32](4789),
		// 				NetworkInterfaceTapConfigurations: []*armnetwork.InterfaceTapConfiguration{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/testNetworkInterface2/tapConfigurations/testtapConfiguration"),
		// 				}},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				ResourceGUID: to.Ptr("6A7C139D-8B8D-499B-B7CB-4F3F02A8A44F"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testvtap2"),
		// 			Type: to.Ptr("Microsoft.Network/virtualNetworkTaps"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkTaps/testvtap2"),
		// 			Location: to.Ptr("centraluseuap"),
		// 			Etag: to.Ptr("etag"),
		// 			Properties: &armnetwork.VirtualNetworkTapPropertiesFormat{
		// 				DestinationNetworkInterfaceIPConfiguration: &armnetwork.InterfaceIPConfiguration{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/testNetworkInterface/ipConfigurations/testIPConfig1"),
		// 				},
		// 				DestinationPort: to.Ptr[int32](4789),
		// 				NetworkInterfaceTapConfigurations: []*armnetwork.InterfaceTapConfiguration{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/testNetworkInterface3/tapConfigurations/testtapConfiguration"),
		// 				}},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				ResourceGUID: to.Ptr("6A7C139D-8B8D-499B-B7CB-4F3F02A8A44F"),
		// 			},
		// 	}},
		// }
	}
}
