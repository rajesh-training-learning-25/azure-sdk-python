//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armkubernetesruntime_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesruntime/armkubernetesruntime"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ad60d7f8eba124edc6999677c55aba2184e303b0/specification/kubernetesruntime/resource-manager/Microsoft.KubernetesRuntime/stable/2024-03-01/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkubernetesruntime.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page.OperationListResult = armkubernetesruntime.OperationListResult{
		// 	Value: []*armkubernetesruntime.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesRuntime/register/action"),
		// 			Display: &armkubernetesruntime.OperationDisplay{
		// 				Description: to.Ptr("Register the subscription for Microsoft.KubernetesRuntime"),
		// 				Operation: to.Ptr("Register the Microsoft.KubernetesRuntime"),
		// 				Provider: to.Ptr("Microsoft.KubernetesRuntime"),
		// 				Resource: to.Ptr("Microsoft.KubernetesRuntime"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesRuntime/unregister/action"),
		// 			Display: &armkubernetesruntime.OperationDisplay{
		// 				Description: to.Ptr("Unregister the subscription for Microsoft.KubernetesRuntime"),
		// 				Operation: to.Ptr("Unregister the Microsoft.KubernetesRuntime"),
		// 				Provider: to.Ptr("Microsoft.KubernetesRuntime"),
		// 				Resource: to.Ptr("Microsoft.KubernetesRuntime"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesRuntime/operations/read"),
		// 			Display: &armkubernetesruntime.OperationDisplay{
		// 				Description: to.Ptr("read operations"),
		// 				Operation: to.Ptr("read_operations"),
		// 				Provider: to.Ptr("Microsoft.KubernetesRuntime"),
		// 				Resource: to.Ptr("operations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesRuntime/locations/operationStatuses/read"),
		// 			Display: &armkubernetesruntime.OperationDisplay{
		// 				Description: to.Ptr("read operationStatuses"),
		// 				Operation: to.Ptr("read_operationStatuses"),
		// 				Provider: to.Ptr("Microsoft.KubernetesRuntime"),
		// 				Resource: to.Ptr("locations/operationStatuses"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesRuntime/locations/operationStatuses/write"),
		// 			Display: &armkubernetesruntime.OperationDisplay{
		// 				Description: to.Ptr("write operationStatuses"),
		// 				Operation: to.Ptr("write_operationStatuses"),
		// 				Provider: to.Ptr("Microsoft.KubernetesRuntime"),
		// 				Resource: to.Ptr("locations/operationStatuses"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesRuntime/storageClasses/read"),
		// 			Display: &armkubernetesruntime.OperationDisplay{
		// 				Description: to.Ptr("List StorageClassResource resources by parent"),
		// 				Operation: to.Ptr("StorageClass_List"),
		// 				Provider: to.Ptr("Microsoft.KubernetesRuntime"),
		// 				Resource: to.Ptr("storageClasses"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesRuntime/storageClasses/read"),
		// 			Display: &armkubernetesruntime.OperationDisplay{
		// 				Description: to.Ptr("Get a StorageClassResource"),
		// 				Operation: to.Ptr("StorageClass_Get"),
		// 				Provider: to.Ptr("Microsoft.KubernetesRuntime"),
		// 				Resource: to.Ptr("storageClasses"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesRuntime/storageClasses/write"),
		// 			Display: &armkubernetesruntime.OperationDisplay{
		// 				Description: to.Ptr("Create a StorageClassResource"),
		// 				Operation: to.Ptr("StorageClass_CreateOrUpdate"),
		// 				Provider: to.Ptr("Microsoft.KubernetesRuntime"),
		// 				Resource: to.Ptr("storageClasses"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesRuntime/storageClasses/delete"),
		// 			Display: &armkubernetesruntime.OperationDisplay{
		// 				Description: to.Ptr("Delete a StorageClassResource"),
		// 				Operation: to.Ptr("StorageClass_Delete"),
		// 				Provider: to.Ptr("Microsoft.KubernetesRuntime"),
		// 				Resource: to.Ptr("storageClasses"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesRuntime/storageClasses/write"),
		// 			Display: &armkubernetesruntime.OperationDisplay{
		// 				Description: to.Ptr("Update a StorageClassResource"),
		// 				Operation: to.Ptr("StorageClass_Update"),
		// 				Provider: to.Ptr("Microsoft.KubernetesRuntime"),
		// 				Resource: to.Ptr("storageClasses"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesRuntime/services/read"),
		// 			Display: &armkubernetesruntime.OperationDisplay{
		// 				Description: to.Ptr("List ServiceResource resources by parent"),
		// 				Operation: to.Ptr("Services_List"),
		// 				Provider: to.Ptr("Microsoft.KubernetesRuntime"),
		// 				Resource: to.Ptr("services"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesRuntime/services/read"),
		// 			Display: &armkubernetesruntime.OperationDisplay{
		// 				Description: to.Ptr("Get a ServiceResource"),
		// 				Operation: to.Ptr("Services_Get"),
		// 				Provider: to.Ptr("Microsoft.KubernetesRuntime"),
		// 				Resource: to.Ptr("services"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesRuntime/services/write"),
		// 			Display: &armkubernetesruntime.OperationDisplay{
		// 				Description: to.Ptr("Create a ServiceResource"),
		// 				Operation: to.Ptr("Services_CreateOrUpdate"),
		// 				Provider: to.Ptr("Microsoft.KubernetesRuntime"),
		// 				Resource: to.Ptr("services"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesRuntime/services/delete"),
		// 			Display: &armkubernetesruntime.OperationDisplay{
		// 				Description: to.Ptr("Delete a ServiceResource"),
		// 				Operation: to.Ptr("Services_Delete"),
		// 				Provider: to.Ptr("Microsoft.KubernetesRuntime"),
		// 				Resource: to.Ptr("services"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesRuntime/loadBalancers/read"),
		// 			Display: &armkubernetesruntime.OperationDisplay{
		// 				Description: to.Ptr("List LoadBalancer resources by parent"),
		// 				Operation: to.Ptr("LoadBalancers_List"),
		// 				Provider: to.Ptr("Microsoft.KubernetesRuntime"),
		// 				Resource: to.Ptr("loadBalancers"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesRuntime/loadBalancers/read"),
		// 			Display: &armkubernetesruntime.OperationDisplay{
		// 				Description: to.Ptr("Get a LoadBalancer"),
		// 				Operation: to.Ptr("LoadBalancers_Get"),
		// 				Provider: to.Ptr("Microsoft.KubernetesRuntime"),
		// 				Resource: to.Ptr("loadBalancers"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesRuntime/loadBalancers/write"),
		// 			Display: &armkubernetesruntime.OperationDisplay{
		// 				Description: to.Ptr("Create a LoadBalancer"),
		// 				Operation: to.Ptr("LoadBalancers_CreateOrUpdate"),
		// 				Provider: to.Ptr("Microsoft.KubernetesRuntime"),
		// 				Resource: to.Ptr("loadBalancers"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesRuntime/loadBalancers/delete"),
		// 			Display: &armkubernetesruntime.OperationDisplay{
		// 				Description: to.Ptr("Delete a LoadBalancer"),
		// 				Operation: to.Ptr("LoadBalancers_Delete"),
		// 				Provider: to.Ptr("Microsoft.KubernetesRuntime"),
		// 				Resource: to.Ptr("loadBalancers"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesRuntime/bgpPeers/read"),
		// 			Display: &armkubernetesruntime.OperationDisplay{
		// 				Description: to.Ptr("List BgpPeer resources by parent"),
		// 				Operation: to.Ptr("BgpPeers_List"),
		// 				Provider: to.Ptr("Microsoft.KubernetesRuntime"),
		// 				Resource: to.Ptr("bgpPeers"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesRuntime/bgpPeers/read"),
		// 			Display: &armkubernetesruntime.OperationDisplay{
		// 				Description: to.Ptr("Get a BgpPeer"),
		// 				Operation: to.Ptr("BgpPeers_Get"),
		// 				Provider: to.Ptr("Microsoft.KubernetesRuntime"),
		// 				Resource: to.Ptr("bgpPeers"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesRuntime/bgpPeers/write"),
		// 			Display: &armkubernetesruntime.OperationDisplay{
		// 				Description: to.Ptr("Create a BgpPeer"),
		// 				Operation: to.Ptr("BgpPeers_CreateOrUpdate"),
		// 				Provider: to.Ptr("Microsoft.KubernetesRuntime"),
		// 				Resource: to.Ptr("bgpPeers"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesRuntime/bgpPeers/delete"),
		// 			Display: &armkubernetesruntime.OperationDisplay{
		// 				Description: to.Ptr("Delete a BgpPeer"),
		// 				Operation: to.Ptr("BgpPeers_Delete"),
		// 				Provider: to.Ptr("Microsoft.KubernetesRuntime"),
		// 				Resource: to.Ptr("bgpPeers"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 	}},
		// }
	}
}
