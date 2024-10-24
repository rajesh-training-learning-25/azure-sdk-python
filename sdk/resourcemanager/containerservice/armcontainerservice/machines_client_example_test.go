//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ad60d7f8eba124edc6999677c55aba2184e303b0/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2024-08-01/examples/MachineList.json
func ExampleMachinesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMachinesClient().NewListPager("rg1", "clustername1", "agentpool1", nil)
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
		// page.MachineListResult = armcontainerservice.MachineListResult{
		// 	Value: []*armcontainerservice.Machine{
		// 		{
		// 			Name: to.Ptr("aks-nodepool1-25481572-vmss000000"),
		// 			Type: to.Ptr("Microsoft.ContainerService/managedClusters/agentPools/machines"),
		// 			ID: to.Ptr("/subscriptions/26fe00f8-9173-4872-9134-bb1d2e00343a/resourceGroups/dummyRG/providers/Microsoft.ContainerService/managedClusters/round/agentPools/nodepool1/machines/aks-nodepool1-25481572-vmss000000"),
		// 			Properties: &armcontainerservice.MachineProperties{
		// 				Network: &armcontainerservice.MachineNetworkProperties{
		// 					IPAddresses: []*armcontainerservice.MachineIPAddress{
		// 						{
		// 							Family: to.Ptr(armcontainerservice.IPFamilyIPv4),
		// 							IP: to.Ptr("172.20.2.4"),
		// 						},
		// 						{
		// 							Family: to.Ptr(armcontainerservice.IPFamilyIPv4),
		// 							IP: to.Ptr("10.0.0.1"),
		// 					}},
		// 				},
		// 				ResourceID: to.Ptr("/subscriptions/26fe00f8-9173-4872-9134-bb1d2e00343a/resourceGroups/dummyRG/providers/Microsoft.Compute/virtualMachineScaleSets/aks-nodepool1-25481572-vmss/virtualMachines/0"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ad60d7f8eba124edc6999677c55aba2184e303b0/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2024-08-01/examples/MachineGet.json
func ExampleMachinesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMachinesClient().Get(ctx, "rg1", "clustername1", "agentpool1", "aks-nodepool1-42263519-vmss00000t", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Machine = armcontainerservice.Machine{
	// 	Name: to.Ptr("aks-nodepool1-25481572-vmss000000"),
	// 	Type: to.Ptr("Microsoft.ContainerService/managedClusters/agentPools/machines"),
	// 	ID: to.Ptr("/subscriptions/26fe00f8-9173-4872-9134-bb1d2e00343a/resourceGroups/dummyRG/providers/Microsoft.ContainerService/managedClusters/round/agentPools/nodepool1/machines/aks-nodepool1-25481572-vmss000000"),
	// 	Properties: &armcontainerservice.MachineProperties{
	// 		Network: &armcontainerservice.MachineNetworkProperties{
	// 			IPAddresses: []*armcontainerservice.MachineIPAddress{
	// 				{
	// 					Family: to.Ptr(armcontainerservice.IPFamilyIPv4),
	// 					IP: to.Ptr("172.20.2.4"),
	// 				},
	// 				{
	// 					Family: to.Ptr(armcontainerservice.IPFamilyIPv4),
	// 					IP: to.Ptr("10.0.0.1"),
	// 			}},
	// 		},
	// 		ResourceID: to.Ptr("/subscriptions/26fe00f8-9173-4872-9134-bb1d2e00343a/resourceGroups/dummyRG/providers/Microsoft.Compute/virtualMachineScaleSets/aks-nodepool1-25481572-vmss/virtualMachines/0"),
	// 	},
	// }
}
