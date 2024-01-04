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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/639ecfad68419328658bd4cfe7094af4ce472be2/specification/network/resource-manager/Microsoft.Network/stable/2023-06-01/examples/IpAllocationDelete.json
func ExampleIPAllocationsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIPAllocationsClient().BeginDelete(ctx, "rg1", "test-ipallocation", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/639ecfad68419328658bd4cfe7094af4ce472be2/specification/network/resource-manager/Microsoft.Network/stable/2023-06-01/examples/IpAllocationGet.json
func ExampleIPAllocationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIPAllocationsClient().Get(ctx, "rg1", "test-ipallocation", &armnetwork.IPAllocationsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IPAllocation = armnetwork.IPAllocation{
	// 	Name: to.Ptr("test-ipallocation"),
	// 	Type: to.Ptr("Microsoft.Network/IpAllocations"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/IpAllocations/test-ipallocation"),
	// 	Location: to.Ptr("centraluseuap"),
	// 	Properties: &armnetwork.IPAllocationPropertiesFormat{
	// 		Type: to.Ptr(armnetwork.IPAllocationTypeHypernet),
	// 		AllocationTags: map[string]*string{
	// 			"VNetID": to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/HypernetVnet1"),
	// 		},
	// 		IpamAllocationID: to.Ptr("916d3b28-663f-448b-9abc-1bea9d5fed8f"),
	// 		Prefix: to.Ptr("3.2.5.0/24"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/639ecfad68419328658bd4cfe7094af4ce472be2/specification/network/resource-manager/Microsoft.Network/stable/2023-06-01/examples/IpAllocationCreate.json
func ExampleIPAllocationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIPAllocationsClient().BeginCreateOrUpdate(ctx, "rg1", "test-ipallocation", armnetwork.IPAllocation{
		Location: to.Ptr("centraluseuap"),
		Properties: &armnetwork.IPAllocationPropertiesFormat{
			Type: to.Ptr(armnetwork.IPAllocationTypeHypernet),
			AllocationTags: map[string]*string{
				"VNetID": to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/HypernetVnet1"),
			},
			Prefix: to.Ptr("3.2.5.0/24"),
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
	// res.IPAllocation = armnetwork.IPAllocation{
	// 	Name: to.Ptr("test-ipallocation"),
	// 	Type: to.Ptr("Microsoft.Network/IpAllocations"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/IpAllocations/test-ipallocation"),
	// 	Location: to.Ptr("centraluseuap"),
	// 	Properties: &armnetwork.IPAllocationPropertiesFormat{
	// 		Type: to.Ptr(armnetwork.IPAllocationTypeHypernet),
	// 		AllocationTags: map[string]*string{
	// 			"VNetID": to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/HypernetVnet1"),
	// 		},
	// 		IpamAllocationID: to.Ptr("916d3b28-663f-448b-9abc-1bea9d5fed8f"),
	// 		Prefix: to.Ptr("3.2.5.0/24"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/639ecfad68419328658bd4cfe7094af4ce472be2/specification/network/resource-manager/Microsoft.Network/stable/2023-06-01/examples/IpAllocationUpdateTags.json
func ExampleIPAllocationsClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIPAllocationsClient().UpdateTags(ctx, "rg1", "test-ipallocation", armnetwork.TagsObject{
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
	// res.IPAllocation = armnetwork.IPAllocation{
	// 	Name: to.Ptr("test-ipallocation"),
	// 	Type: to.Ptr("Microsoft.Network/IpAllocations"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/IpAllocations/test-ipallocation"),
	// 	Location: to.Ptr("centraluseuap"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armnetwork.IPAllocationPropertiesFormat{
	// 		Type: to.Ptr(armnetwork.IPAllocationTypeHypernet),
	// 		AllocationTags: map[string]*string{
	// 			"VNetID": to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/HypernetVnet1"),
	// 		},
	// 		IpamAllocationID: to.Ptr("916d3b28-663f-448b-9abc-1bea9d5fed8f"),
	// 		Prefix: to.Ptr("3.2.5.0/24"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/639ecfad68419328658bd4cfe7094af4ce472be2/specification/network/resource-manager/Microsoft.Network/stable/2023-06-01/examples/IpAllocationList.json
func ExampleIPAllocationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIPAllocationsClient().NewListPager(nil)
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
		// page.IPAllocationListResult = armnetwork.IPAllocationListResult{
		// 	Value: []*armnetwork.IPAllocation{
		// 		{
		// 			Name: to.Ptr("test-ipallocation1"),
		// 			Type: to.Ptr("Microsoft.Network/IpAllocations"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/IpAllocations/test-ipallocation1"),
		// 			Location: to.Ptr("centraluseuap"),
		// 			Properties: &armnetwork.IPAllocationPropertiesFormat{
		// 				Type: to.Ptr(armnetwork.IPAllocationTypeHypernet),
		// 				AllocationTags: map[string]*string{
		// 					"VNetID": to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/HypernetVnet1"),
		// 				},
		// 				IpamAllocationID: to.Ptr("916d3b28-663f-448b-9abc-1bea9d5fed8f"),
		// 				Prefix: to.Ptr("3.2.5.0/24"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-ipallocation2"),
		// 			Type: to.Ptr("Microsoft.Network/IpAllocations"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/IpAllocations/test-ipallocation2"),
		// 			Location: to.Ptr("centraluseuap"),
		// 			Properties: &armnetwork.IPAllocationPropertiesFormat{
		// 				Type: to.Ptr(armnetwork.IPAllocationTypeHypernet),
		// 				AllocationTags: map[string]*string{
		// 					"VNetID": to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/HypernetVnet2"),
		// 				},
		// 				IpamAllocationID: to.Ptr("57dc7256-2ff7-43f2-b9c8-85a70b5c6408"),
		// 				Prefix: to.Ptr("3.2.6.0/24"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/639ecfad68419328658bd4cfe7094af4ce472be2/specification/network/resource-manager/Microsoft.Network/stable/2023-06-01/examples/IpAllocationListByResourceGroup.json
func ExampleIPAllocationsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIPAllocationsClient().NewListByResourceGroupPager("rg1", nil)
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
		// page.IPAllocationListResult = armnetwork.IPAllocationListResult{
		// 	Value: []*armnetwork.IPAllocation{
		// 		{
		// 			Name: to.Ptr("test-ipallocation1"),
		// 			Type: to.Ptr("Microsoft.Network/IpAllocations"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/IpAllocations/test-ipallocation1"),
		// 			Location: to.Ptr("centraluseuap"),
		// 			Properties: &armnetwork.IPAllocationPropertiesFormat{
		// 				Type: to.Ptr(armnetwork.IPAllocationTypeHypernet),
		// 				AllocationTags: map[string]*string{
		// 					"VNetID": to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/HypernetVnet1"),
		// 				},
		// 				IpamAllocationID: to.Ptr("916d3b28-663f-448b-9abc-1bea9d5fed8f"),
		// 				Prefix: to.Ptr("3.2.5.0/24"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-ipallocation2"),
		// 			Type: to.Ptr("Microsoft.Network/IpAllocations"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/IpAllocations/test-ipallocation2"),
		// 			Location: to.Ptr("centraluseuap"),
		// 			Properties: &armnetwork.IPAllocationPropertiesFormat{
		// 				Type: to.Ptr(armnetwork.IPAllocationTypeHypernet),
		// 				AllocationTags: map[string]*string{
		// 					"VNetID": to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/HypernetVnet2"),
		// 				},
		// 				IpamAllocationID: to.Ptr("57dc7256-2ff7-43f2-b9c8-85a70b5c6408"),
		// 				Prefix: to.Ptr("3.2.6.0/24"),
		// 			},
		// 	}},
		// }
	}
}
