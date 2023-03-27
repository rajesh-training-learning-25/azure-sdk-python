//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armstorage_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b32e1896f30e6ea155449cb49719a6286e32b961/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/TableOperationPut.json
func ExampleTableClient_Create_tableOperationPut() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTableClient().Create(ctx, "res3376", "sto328", "table6185", &armstorage.TableClientCreateOptions{Parameters: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Table = armstorage.Table{
	// 	Name: to.Ptr("table6185"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/tableServices/tables"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res3376/providers/Microsoft.Storage/storageAccounts/sto328/tableServices/default/tables/table6185"),
	// 	TableProperties: &armstorage.TableProperties{
	// 		TableName: to.Ptr("table6185"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b32e1896f30e6ea155449cb49719a6286e32b961/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/TableOperationPutOrPatchAcls.json
func ExampleTableClient_Create_tableOperationPutOrPatchAcls() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTableClient().Create(ctx, "res3376", "sto328", "table6185", &armstorage.TableClientCreateOptions{Parameters: &armstorage.Table{
		TableProperties: &armstorage.TableProperties{
			SignedIdentifiers: []*armstorage.TableSignedIdentifier{
				{
					AccessPolicy: &armstorage.TableAccessPolicy{
						ExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-20T08:49:37.0000000Z"); return t }()),
						Permission: to.Ptr("raud"),
						StartTime:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-17T08:49:37.0000000Z"); return t }()),
					},
					ID: to.Ptr("MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI"),
				},
				{
					AccessPolicy: &armstorage.TableAccessPolicy{
						ExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-20T08:49:37.0000000Z"); return t }()),
						Permission: to.Ptr("rad"),
						StartTime:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-17T08:49:37.0000000Z"); return t }()),
					},
					ID: to.Ptr("PTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODklMTI"),
				}},
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Table = armstorage.Table{
	// 	Name: to.Ptr("table6185"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/tableServices/tables"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res3376/providers/Microsoft.Storage/storageAccounts/sto328/tableServices/default/tables/table6185"),
	// 	TableProperties: &armstorage.TableProperties{
	// 		SignedIdentifiers: []*armstorage.TableSignedIdentifier{
	// 			{
	// 				AccessPolicy: &armstorage.TableAccessPolicy{
	// 					ExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-20T08:49:37.0000000Z"); return t}()),
	// 					Permission: to.Ptr("raud"),
	// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-17T08:49:37.0000000Z"); return t}()),
	// 				},
	// 				ID: to.Ptr("MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI"),
	// 			},
	// 			{
	// 				AccessPolicy: &armstorage.TableAccessPolicy{
	// 					ExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-20T08:49:37.0000000Z"); return t}()),
	// 					Permission: to.Ptr("rad"),
	// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-17T08:49:37.0000000Z"); return t}()),
	// 				},
	// 				ID: to.Ptr("PTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODklMTI"),
	// 		}},
	// 		TableName: to.Ptr("table6185"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b32e1896f30e6ea155449cb49719a6286e32b961/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/TableOperationPatch.json
func ExampleTableClient_Update_tableOperationPatch() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTableClient().Update(ctx, "res3376", "sto328", "table6185", &armstorage.TableClientUpdateOptions{Parameters: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Table = armstorage.Table{
	// 	Name: to.Ptr("table6185"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/tableServices/tables"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res3376/providers/Microsoft.Storage/storageAccounts/sto328/tableServices/default/tables/table6185"),
	// 	TableProperties: &armstorage.TableProperties{
	// 		TableName: to.Ptr("table6185"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b32e1896f30e6ea155449cb49719a6286e32b961/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/TableOperationGet.json
func ExampleTableClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTableClient().Get(ctx, "res3376", "sto328", "table6185", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Table = armstorage.Table{
	// 	Name: to.Ptr("table6185"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/tableServices/tables"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res3376/providers/Microsoft.Storage/storageAccounts/sto328/tableServices/default/tables/table6185"),
	// 	TableProperties: &armstorage.TableProperties{
	// 		TableName: to.Ptr("table6185"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b32e1896f30e6ea155449cb49719a6286e32b961/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/TableOperationDelete.json
func ExampleTableClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewTableClient().Delete(ctx, "res3376", "sto328", "table6185", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b32e1896f30e6ea155449cb49719a6286e32b961/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/TableOperationList.json
func ExampleTableClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTableClient().NewListPager("res9290", "sto328", nil)
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
		// page.ListTableResource = armstorage.ListTableResource{
		// 	Value: []*armstorage.Table{
		// 		{
		// 			Name: to.Ptr("table6185"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/tableServices/tables"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res3376/providers/Microsoft.Storage/storageAccounts/sto328/tableServices/default/tables/table6185"),
		// 			TableProperties: &armstorage.TableProperties{
		// 				TableName: to.Ptr("table6185"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("table6186"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/tableServices/tables"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res3376/providers/Microsoft.Storage/storageAccounts/sto328/tableServices/default/tables/table6186"),
		// 			TableProperties: &armstorage.TableProperties{
		// 				TableName: to.Ptr("table6186"),
		// 			},
		// 	}},
		// }
	}
}
