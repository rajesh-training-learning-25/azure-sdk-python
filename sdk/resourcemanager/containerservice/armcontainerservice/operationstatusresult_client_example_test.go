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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/2a3294803a8df71659e27711e5fbfe4e97826828/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2024-06-02-preview/examples/OperationStatusResultList.json
func ExampleOperationStatusResultClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationStatusResultClient().NewListPager("rg1", "clustername1", nil)
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
		// page.OperationStatusResultList = armcontainerservice.OperationStatusResultList{
		// 	Value: []*armcontainerservice.OperationStatusResult{
		// 		{
		// 			Name: to.Ptr("d11edb09-6e27-429f-9fe5-17baf773bc4a"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/operations/d11edb09-6e27-429f-9fe5-17baf773bc4a"),
		// 			PercentComplete: to.Ptr[float32](40),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-26T12:14:26.317Z"); return t}()),
		// 			Status: to.Ptr("InProgress"),
		// 		},
		// 		{
		// 			Name: to.Ptr("d11edb09-6e27-429f-9fe5-17baf773bc4b"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-26T12:14:50.317Z"); return t}()),
		// 			Error: &armcontainerservice.ErrorDetail{
		// 				Code: to.Ptr("ReconcileAgentPoolIdentityError"),
		// 				Message: to.Ptr("Reconcile agent pool nodepool1 identity failed"),
		// 			},
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/operations/d11edb09-6e27-429f-9fe5-17baf773bc4b"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-26T12:14:26.317Z"); return t}()),
		// 			Status: to.Ptr("Failed"),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/2a3294803a8df71659e27711e5fbfe4e97826828/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2024-06-02-preview/examples/OperationStatusResultGet.json
func ExampleOperationStatusResultClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationStatusResultClient().Get(ctx, "rg1", "clustername1", "00000000-0000-0000-0000-000000000001", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationStatusResult = armcontainerservice.OperationStatusResult{
	// 	Name: to.Ptr("00000000-0000-0000-0000-000000000001"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/operations/00000000-0000-0000-0000-000000000001"),
	// 	PercentComplete: to.Ptr[float32](40),
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-26T12:14:26.317Z"); return t}()),
	// 	Status: to.Ptr("InProgress"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/2a3294803a8df71659e27711e5fbfe4e97826828/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2024-06-02-preview/examples/OperationStatusResultGetByAgentPool.json
func ExampleOperationStatusResultClient_GetByAgentPool() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationStatusResultClient().GetByAgentPool(ctx, "rg1", "clustername1", "agentpool1", "00000000-0000-0000-0000-000000000001", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationStatusResult = armcontainerservice.OperationStatusResult{
	// 	Name: to.Ptr("00000000-0000-0000-0000-000000000001"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/agentPools/agentpool1/operations/00000000-0000-0000-0000-000000000001"),
	// 	PercentComplete: to.Ptr[float32](40),
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-26T12:14:26.317Z"); return t}()),
	// 	Status: to.Ptr("InProgress"),
	// }
}
