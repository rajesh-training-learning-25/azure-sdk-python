//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/batch/resource-manager/Microsoft.Batch/stable/2023-05-01/examples/PrivateLinkResourcesList.json
func ExamplePrivateLinkResourceClient_NewListByBatchAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkResourceClient().NewListByBatchAccountPager("default-azurebatch-japaneast", "sampleacct", &armbatch.PrivateLinkResourceClientListByBatchAccountOptions{Maxresults: nil})
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
		// page.ListPrivateLinkResourcesResult = armbatch.ListPrivateLinkResourcesResult{
		// 	Value: []*armbatch.PrivateLinkResource{
		// 		{
		// 			Name: to.Ptr("batchAccount"),
		// 			Type: to.Ptr("Microsoft.Batch/batchAccounts/privateLinkResources"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/privateLinkResources/batchAccount"),
		// 			Properties: &armbatch.PrivateLinkResourceProperties{
		// 				GroupID: to.Ptr("batchAccount"),
		// 				RequiredMembers: []*string{
		// 					to.Ptr("batchAccount")},
		// 					RequiredZoneNames: []*string{
		// 						to.Ptr("privatelink.japaneast.batch.azure.com")},
		// 					},
		// 			}},
		// 		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/batch/resource-manager/Microsoft.Batch/stable/2023-05-01/examples/PrivateLinkResourceGet.json
func ExamplePrivateLinkResourceClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkResourceClient().Get(ctx, "default-azurebatch-japaneast", "sampleacct", "batchAccount", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkResource = armbatch.PrivateLinkResource{
	// 	Name: to.Ptr("sampleacct"),
	// 	Type: to.Ptr("Microsoft.Batch/batchAccounts/privateLinkResources"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/privateLinkResources/batchAccount"),
	// 	Properties: &armbatch.PrivateLinkResourceProperties{
	// 		GroupID: to.Ptr("batchAccount"),
	// 		RequiredMembers: []*string{
	// 			to.Ptr("batchAccount")},
	// 			RequiredZoneNames: []*string{
	// 				to.Ptr("privatelink.japaneast.batch.azure.com")},
	// 			},
	// 		}
}
