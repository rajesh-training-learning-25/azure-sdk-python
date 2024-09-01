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

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e79d9ef3e065f2dcb6bd1db51e29c62a99dff5cb/specification/batch/resource-manager/Microsoft.Batch/stable/2024-07-01/examples/ApplicationCreate.json
func ExampleApplicationClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationClient().Create(ctx, "default-azurebatch-japaneast", "sampleacct", "app1", &armbatch.ApplicationClientCreateOptions{Parameters: &armbatch.Application{
		Properties: &armbatch.ApplicationProperties{
			AllowUpdates: to.Ptr(false),
			DisplayName:  to.Ptr("myAppName"),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Application = armbatch.Application{
	// 	Name: to.Ptr("app1"),
	// 	Type: to.Ptr("Microsoft.Batch/batchAccounts/applications"),
	// 	Etag: to.Ptr("W/\"0x8D64F8EBB3DC411\""),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/applications/app1"),
	// 	Properties: &armbatch.ApplicationProperties{
	// 		AllowUpdates: to.Ptr(false),
	// 		DisplayName: to.Ptr("myAppName"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e79d9ef3e065f2dcb6bd1db51e29c62a99dff5cb/specification/batch/resource-manager/Microsoft.Batch/stable/2024-07-01/examples/ApplicationDelete.json
func ExampleApplicationClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewApplicationClient().Delete(ctx, "default-azurebatch-japaneast", "sampleacct", "app1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e79d9ef3e065f2dcb6bd1db51e29c62a99dff5cb/specification/batch/resource-manager/Microsoft.Batch/stable/2024-07-01/examples/ApplicationGet.json
func ExampleApplicationClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationClient().Get(ctx, "default-azurebatch-japaneast", "sampleacct", "app1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Application = armbatch.Application{
	// 	Name: to.Ptr("app1"),
	// 	Type: to.Ptr("Microsoft.Batch/batchAccounts/applications"),
	// 	Etag: to.Ptr("W/\"0x8D64F915BDF7F00\""),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/applications/app1"),
	// 	Properties: &armbatch.ApplicationProperties{
	// 		AllowUpdates: to.Ptr(true),
	// 		DisplayName: to.Ptr("Sample Application"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e79d9ef3e065f2dcb6bd1db51e29c62a99dff5cb/specification/batch/resource-manager/Microsoft.Batch/stable/2024-07-01/examples/ApplicationUpdate.json
func ExampleApplicationClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationClient().Update(ctx, "default-azurebatch-japaneast", "sampleacct", "app1", armbatch.Application{
		Properties: &armbatch.ApplicationProperties{
			AllowUpdates:   to.Ptr(true),
			DefaultVersion: to.Ptr("2"),
			DisplayName:    to.Ptr("myAppName"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Application = armbatch.Application{
	// 	Name: to.Ptr("app1"),
	// 	Type: to.Ptr("Microsoft.Batch/batchAccounts/applications"),
	// 	Etag: to.Ptr("W/\"0x8D64F915BDF7F00\""),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/applications/app1"),
	// 	Properties: &armbatch.ApplicationProperties{
	// 		AllowUpdates: to.Ptr(true),
	// 		DefaultVersion: to.Ptr("2"),
	// 		DisplayName: to.Ptr("myAppName"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e79d9ef3e065f2dcb6bd1db51e29c62a99dff5cb/specification/batch/resource-manager/Microsoft.Batch/stable/2024-07-01/examples/ApplicationList.json
func ExampleApplicationClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewApplicationClient().NewListPager("default-azurebatch-japaneast", "sampleacct", &armbatch.ApplicationClientListOptions{Maxresults: nil})
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
		// page.ListApplicationsResult = armbatch.ListApplicationsResult{
		// 	Value: []*armbatch.Application{
		// 		{
		// 			Name: to.Ptr("app1"),
		// 			Type: to.Ptr("Microsoft.Batch/batchAccounts/applications"),
		// 			Etag: to.Ptr("W/\"0x8D64F91A9089879\""),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/applications/app1"),
		// 			Properties: &armbatch.ApplicationProperties{
		// 				AllowUpdates: to.Ptr(false),
		// 				DefaultVersion: to.Ptr("1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("app1"),
		// 			Type: to.Ptr("Microsoft.Batch/batchAccounts/applications"),
		// 			Etag: to.Ptr("W/\"0x8D64F91A9089879\""),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/applications/app2"),
		// 			Properties: &armbatch.ApplicationProperties{
		// 				AllowUpdates: to.Ptr(false),
		// 				DefaultVersion: to.Ptr("2.0"),
		// 				DisplayName: to.Ptr("myAppName"),
		// 			},
		// 	}},
		// }
	}
}
