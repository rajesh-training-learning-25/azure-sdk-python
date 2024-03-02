//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e4009d2f8d3bf0271757e522c7d1c1997e193d44/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/availabilitySetExamples/AvailabilitySet_Create.json
func ExampleAvailabilitySetsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAvailabilitySetsClient().CreateOrUpdate(ctx, "myResourceGroup", "myAvailabilitySet", armcompute.AvailabilitySet{
		Location: to.Ptr("westus"),
		Properties: &armcompute.AvailabilitySetProperties{
			PlatformFaultDomainCount:  to.Ptr[int32](2),
			PlatformUpdateDomainCount: to.Ptr[int32](20),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AvailabilitySet = armcompute.AvailabilitySet{
	// 	Name: to.Ptr("myAvailabilitySet"),
	// 	Type: to.Ptr("Microsoft.Compute/availabilitySets"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/availabilitySets/myAvailabilitySet"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armcompute.AvailabilitySetProperties{
	// 		PlatformFaultDomainCount: to.Ptr[int32](2),
	// 		PlatformUpdateDomainCount: to.Ptr[int32](20),
	// 	},
	// 	SKU: &armcompute.SKU{
	// 		Name: to.Ptr("Classic"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e4009d2f8d3bf0271757e522c7d1c1997e193d44/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/availabilitySetExamples/AvailabilitySet_Update_MaximumSet_Gen.json
func ExampleAvailabilitySetsClient_Update_availabilitySetUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAvailabilitySetsClient().Update(ctx, "rgcompute", "aaaaaaaaaaaaaaaaaaa", armcompute.AvailabilitySetUpdate{
		Tags: map[string]*string{
			"key2574": to.Ptr("aaaaaaaa"),
		},
		Properties: &armcompute.AvailabilitySetProperties{
			PlatformFaultDomainCount:  to.Ptr[int32](2),
			PlatformUpdateDomainCount: to.Ptr[int32](20),
			ProximityPlacementGroup: &armcompute.SubResource{
				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
			},
			VirtualMachines: []*armcompute.SubResource{
				{
					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
				}},
		},
		SKU: &armcompute.SKU{
			Name:     to.Ptr("DSv3-Type1"),
			Capacity: to.Ptr[int64](7),
			Tier:     to.Ptr("aaa"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AvailabilitySet = armcompute.AvailabilitySet{
	// 	Name: to.Ptr("myAvailabilitySet"),
	// 	Type: to.Ptr("Microsoft.Compute/availabilitySets"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/availabilitySets/myAvailabilitySet"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"key2505": to.Ptr("aa"),
	// 		"key9626": to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 	},
	// 	Properties: &armcompute.AvailabilitySetProperties{
	// 		PlatformFaultDomainCount: to.Ptr[int32](2),
	// 		PlatformUpdateDomainCount: to.Ptr[int32](20),
	// 		ProximityPlacementGroup: &armcompute.SubResource{
	// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
	// 		},
	// 		Statuses: []*armcompute.InstanceViewStatus{
	// 			{
	// 				Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
	// 				DisplayStatus: to.Ptr("aaaaaa"),
	// 				Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 				Message: to.Ptr("a"),
	// 				Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t}()),
	// 		}},
	// 		VirtualMachines: []*armcompute.SubResource{
	// 			{
	// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
	// 		}},
	// 	},
	// 	SKU: &armcompute.SKU{
	// 		Name: to.Ptr("Classic"),
	// 		Capacity: to.Ptr[int64](29),
	// 		Tier: to.Ptr("aaaaaaaaaaaaaa"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e4009d2f8d3bf0271757e522c7d1c1997e193d44/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/availabilitySetExamples/AvailabilitySet_Update_MinimumSet_Gen.json
func ExampleAvailabilitySetsClient_Update_availabilitySetUpdateMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAvailabilitySetsClient().Update(ctx, "rgcompute", "aaaaaaaaaaaaaaaaaaaa", armcompute.AvailabilitySetUpdate{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AvailabilitySet = armcompute.AvailabilitySet{
	// 	Location: to.Ptr("westus"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e4009d2f8d3bf0271757e522c7d1c1997e193d44/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/availabilitySetExamples/AvailabilitySet_Delete_MaximumSet_Gen.json
func ExampleAvailabilitySetsClient_Delete_availabilitySetDeleteMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAvailabilitySetsClient().Delete(ctx, "rgcompute", "aaaaaaaaaaaaaaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e4009d2f8d3bf0271757e522c7d1c1997e193d44/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/availabilitySetExamples/AvailabilitySet_Delete_MinimumSet_Gen.json
func ExampleAvailabilitySetsClient_Delete_availabilitySetDeleteMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAvailabilitySetsClient().Delete(ctx, "rgcompute", "aaaaaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e4009d2f8d3bf0271757e522c7d1c1997e193d44/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/availabilitySetExamples/AvailabilitySet_Get_MaximumSet_Gen.json
func ExampleAvailabilitySetsClient_Get_availabilitySetGetMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAvailabilitySetsClient().Get(ctx, "rgcompute", "aaaaaaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AvailabilitySet = armcompute.AvailabilitySet{
	// 	Name: to.Ptr("myAvailabilitySet"),
	// 	Type: to.Ptr("Microsoft.Compute/availabilitySets"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/availabilitySets/myAvailabilitySet"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"key2505": to.Ptr("aa"),
	// 		"key9626": to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 	},
	// 	Properties: &armcompute.AvailabilitySetProperties{
	// 		PlatformFaultDomainCount: to.Ptr[int32](2),
	// 		PlatformUpdateDomainCount: to.Ptr[int32](20),
	// 		ProximityPlacementGroup: &armcompute.SubResource{
	// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
	// 		},
	// 		Statuses: []*armcompute.InstanceViewStatus{
	// 			{
	// 				Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
	// 				DisplayStatus: to.Ptr("aaaaaa"),
	// 				Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 				Message: to.Ptr("a"),
	// 				Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t}()),
	// 		}},
	// 		VirtualMachines: []*armcompute.SubResource{
	// 			{
	// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
	// 		}},
	// 	},
	// 	SKU: &armcompute.SKU{
	// 		Name: to.Ptr("Classic"),
	// 		Capacity: to.Ptr[int64](29),
	// 		Tier: to.Ptr("aaaaaaaaaaaaaa"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e4009d2f8d3bf0271757e522c7d1c1997e193d44/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/availabilitySetExamples/AvailabilitySet_Get_MinimumSet_Gen.json
func ExampleAvailabilitySetsClient_Get_availabilitySetGetMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAvailabilitySetsClient().Get(ctx, "rgcompute", "aaaaaaaaaaaaaaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AvailabilitySet = armcompute.AvailabilitySet{
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/rgcompute/providers/Microsoft.Compute/availabilitySets/myAvailabilitySet"),
	// 	Location: to.Ptr("westus"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e4009d2f8d3bf0271757e522c7d1c1997e193d44/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/availabilitySetExamples/AvailabilitySet_ListBySubscription.json
func ExampleAvailabilitySetsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAvailabilitySetsClient().NewListBySubscriptionPager(&armcompute.AvailabilitySetsClientListBySubscriptionOptions{Expand: to.Ptr("virtualMachines\\$ref")})
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
		// page.AvailabilitySetListResult = armcompute.AvailabilitySetListResult{
		// 	Value: []*armcompute.AvailabilitySet{
		// 		{
		// 			Name: to.Ptr("{availabilitySetName}"),
		// 			Type: to.Ptr("Microsoft.Compute/availabilitySets"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 			Location: to.Ptr("australiasoutheast"),
		// 			Properties: &armcompute.AvailabilitySetProperties{
		// 				PlatformFaultDomainCount: to.Ptr[int32](3),
		// 				PlatformUpdateDomainCount: to.Ptr[int32](5),
		// 				VirtualMachines: []*armcompute.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}"),
		// 				}},
		// 			},
		// 			SKU: &armcompute.SKU{
		// 				Name: to.Ptr("Classic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("{availabilitySetName}"),
		// 			Type: to.Ptr("Microsoft.Compute/availabilitySets"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 			Location: to.Ptr("australiasoutheast"),
		// 			Properties: &armcompute.AvailabilitySetProperties{
		// 				PlatformFaultDomainCount: to.Ptr[int32](3),
		// 				PlatformUpdateDomainCount: to.Ptr[int32](5),
		// 				VirtualMachines: []*armcompute.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}"),
		// 				}},
		// 			},
		// 			SKU: &armcompute.SKU{
		// 				Name: to.Ptr("Classic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("{availabilitySetName}"),
		// 			Type: to.Ptr("Microsoft.Compute/availabilitySets"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 			Location: to.Ptr("westcentralus"),
		// 			Tags: map[string]*string{
		// 				"{tagName}": to.Ptr("{tagValue}"),
		// 			},
		// 			Properties: &armcompute.AvailabilitySetProperties{
		// 				PlatformFaultDomainCount: to.Ptr[int32](3),
		// 				PlatformUpdateDomainCount: to.Ptr[int32](5),
		// 				VirtualMachines: []*armcompute.SubResource{
		// 				},
		// 			},
		// 			SKU: &armcompute.SKU{
		// 				Name: to.Ptr("Classic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("{availabilitySetName}"),
		// 			Type: to.Ptr("Microsoft.Compute/availabilitySets"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 			Location: to.Ptr("westcentralus"),
		// 			Tags: map[string]*string{
		// 				"{tagName}": to.Ptr("{tagValue}"),
		// 			},
		// 			Properties: &armcompute.AvailabilitySetProperties{
		// 				PlatformFaultDomainCount: to.Ptr[int32](3),
		// 				PlatformUpdateDomainCount: to.Ptr[int32](5),
		// 				VirtualMachines: []*armcompute.SubResource{
		// 				},
		// 			},
		// 			SKU: &armcompute.SKU{
		// 				Name: to.Ptr("Classic"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e4009d2f8d3bf0271757e522c7d1c1997e193d44/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/availabilitySetExamples/AvailabilitySet_List_MaximumSet_Gen.json
func ExampleAvailabilitySetsClient_NewListPager_availabilitySetListMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAvailabilitySetsClient().NewListPager("rgcompute", nil)
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
		// page.AvailabilitySetListResult = armcompute.AvailabilitySetListResult{
		// 	Value: []*armcompute.AvailabilitySet{
		// 		{
		// 			Name: to.Ptr("{availabilitySetName}"),
		// 			Type: to.Ptr("Microsoft.Compute/availabilitySets"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 			Location: to.Ptr("australiasoutheast"),
		// 			Tags: map[string]*string{
		// 				"key2039": to.Ptr("aaaaaaaaaaaaa"),
		// 			},
		// 			Properties: &armcompute.AvailabilitySetProperties{
		// 				PlatformFaultDomainCount: to.Ptr[int32](3),
		// 				PlatformUpdateDomainCount: to.Ptr[int32](5),
		// 				ProximityPlacementGroup: &armcompute.SubResource{
		// 					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 				},
		// 				Statuses: []*armcompute.InstanceViewStatus{
		// 					{
		// 						Code: to.Ptr("aaaaaaaaaaaaaaa"),
		// 						DisplayStatus: to.Ptr("aaaaaaaaaaa"),
		// 						Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 						Message: to.Ptr("aaaaaa"),
		// 						Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T13:39:56.427Z"); return t}()),
		// 				}},
		// 				VirtualMachines: []*armcompute.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}"),
		// 				}},
		// 			},
		// 			SKU: &armcompute.SKU{
		// 				Name: to.Ptr("Classic"),
		// 				Capacity: to.Ptr[int64](22),
		// 				Tier: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("{availabilitySetName}"),
		// 			Type: to.Ptr("Microsoft.Compute/availabilitySets"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 			Location: to.Ptr("australiasoutheast"),
		// 			Tags: map[string]*string{
		// 				"key5849": to.Ptr("aaaaaaaaaaaaaaa"),
		// 			},
		// 			Properties: &armcompute.AvailabilitySetProperties{
		// 				PlatformFaultDomainCount: to.Ptr[int32](3),
		// 				PlatformUpdateDomainCount: to.Ptr[int32](5),
		// 				ProximityPlacementGroup: &armcompute.SubResource{
		// 					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 				},
		// 				Statuses: []*armcompute.InstanceViewStatus{
		// 					{
		// 						Code: to.Ptr("aaaaaaaaaaaaaaa"),
		// 						DisplayStatus: to.Ptr("aaaaaaaaaaa"),
		// 						Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 						Message: to.Ptr("aaaaaa"),
		// 						Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T13:39:56.427Z"); return t}()),
		// 				}},
		// 				VirtualMachines: []*armcompute.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}"),
		// 				}},
		// 			},
		// 			SKU: &armcompute.SKU{
		// 				Name: to.Ptr("Classic"),
		// 				Capacity: to.Ptr[int64](23),
		// 				Tier: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("{availabilitySetName}"),
		// 			Type: to.Ptr("Microsoft.Compute/availabilitySets"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 			Location: to.Ptr("westcentralus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armcompute.AvailabilitySetProperties{
		// 				PlatformFaultDomainCount: to.Ptr[int32](3),
		// 				PlatformUpdateDomainCount: to.Ptr[int32](5),
		// 				ProximityPlacementGroup: &armcompute.SubResource{
		// 					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 				},
		// 				Statuses: []*armcompute.InstanceViewStatus{
		// 					{
		// 						Code: to.Ptr("aaaaaaaaaaaaaaa"),
		// 						DisplayStatus: to.Ptr("aaaaaaaaaaa"),
		// 						Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 						Message: to.Ptr("aaaaaa"),
		// 						Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T13:39:56.427Z"); return t}()),
		// 				}},
		// 				VirtualMachines: []*armcompute.SubResource{
		// 				},
		// 			},
		// 			SKU: &armcompute.SKU{
		// 				Name: to.Ptr("Classic"),
		// 				Capacity: to.Ptr[int64](26),
		// 				Tier: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("{availabilitySetName}"),
		// 			Type: to.Ptr("Microsoft.Compute/availabilitySets"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 			Location: to.Ptr("westcentralus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armcompute.AvailabilitySetProperties{
		// 				PlatformFaultDomainCount: to.Ptr[int32](3),
		// 				PlatformUpdateDomainCount: to.Ptr[int32](5),
		// 				ProximityPlacementGroup: &armcompute.SubResource{
		// 					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 				},
		// 				Statuses: []*armcompute.InstanceViewStatus{
		// 					{
		// 						Code: to.Ptr("aaaaaaaaaaaaaaa"),
		// 						DisplayStatus: to.Ptr("aaaaaaaaaaa"),
		// 						Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 						Message: to.Ptr("aaaaaa"),
		// 						Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T13:39:56.427Z"); return t}()),
		// 				}},
		// 				VirtualMachines: []*armcompute.SubResource{
		// 				},
		// 			},
		// 			SKU: &armcompute.SKU{
		// 				Name: to.Ptr("Classic"),
		// 				Capacity: to.Ptr[int64](6),
		// 				Tier: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e4009d2f8d3bf0271757e522c7d1c1997e193d44/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/availabilitySetExamples/AvailabilitySet_List_MinimumSet_Gen.json
func ExampleAvailabilitySetsClient_NewListPager_availabilitySetListMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAvailabilitySetsClient().NewListPager("rgcompute", nil)
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
		// page.AvailabilitySetListResult = armcompute.AvailabilitySetListResult{
		// 	Value: []*armcompute.AvailabilitySet{
		// 		{
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/rgcompute/providers/Microsoft.Compute/availabilitySets/myAvailabilitySet1"),
		// 			Location: to.Ptr("australiasoutheast"),
		// 		},
		// 		{
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/rgcompute/providers/Microsoft.Compute/availabilitySets/myAvailabilitySet2"),
		// 			Location: to.Ptr("australiasoutheast"),
		// 		},
		// 		{
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/rgcompute/providers/Microsoft.Compute/availabilitySets/myAvailabilitySet3"),
		// 			Location: to.Ptr("westcentralus"),
		// 		},
		// 		{
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/rgcompute/providers/Microsoft.Compute/availabilitySets/myAvailabilitySet4"),
		// 			Location: to.Ptr("westcentralus"),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e4009d2f8d3bf0271757e522c7d1c1997e193d44/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/availabilitySetExamples/AvailabilitySet_ListAvailableSizes_MaximumSet_Gen.json
func ExampleAvailabilitySetsClient_NewListAvailableSizesPager_availabilitySetListAvailableSizesMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAvailabilitySetsClient().NewListAvailableSizesPager("rgcompute", "aaaaaaaaaaaaaaaaaaaa", nil)
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
		// page.VirtualMachineSizeListResult = armcompute.VirtualMachineSizeListResult{
		// 	Value: []*armcompute.VirtualMachineSize{
		// 		{
		// 			Name: to.Ptr("Standard_A1_V2"),
		// 			MaxDataDiskCount: to.Ptr[int32](2),
		// 			MemoryInMB: to.Ptr[int32](2048),
		// 			NumberOfCores: to.Ptr[int32](1),
		// 			OSDiskSizeInMB: to.Ptr[int32](1047552),
		// 			ResourceDiskSizeInMB: to.Ptr[int32](10240),
		// 		},
		// 		{
		// 			Name: to.Ptr("Standard_A2_V2"),
		// 			MaxDataDiskCount: to.Ptr[int32](4),
		// 			MemoryInMB: to.Ptr[int32](4096),
		// 			NumberOfCores: to.Ptr[int32](2),
		// 			OSDiskSizeInMB: to.Ptr[int32](1047552),
		// 			ResourceDiskSizeInMB: to.Ptr[int32](20480),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e4009d2f8d3bf0271757e522c7d1c1997e193d44/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/availabilitySetExamples/AvailabilitySet_ListAvailableSizes_MinimumSet_Gen.json
func ExampleAvailabilitySetsClient_NewListAvailableSizesPager_availabilitySetListAvailableSizesMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAvailabilitySetsClient().NewListAvailableSizesPager("rgcompute", "aa", nil)
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
		// page.VirtualMachineSizeListResult = armcompute.VirtualMachineSizeListResult{
		// }
	}
}
