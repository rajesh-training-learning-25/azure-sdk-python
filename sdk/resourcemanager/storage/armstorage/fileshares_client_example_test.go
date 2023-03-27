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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b32e1896f30e6ea155449cb49719a6286e32b961/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/DeletedFileSharesList.json
func ExampleFileSharesClient_NewListPager_listDeletedShares() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFileSharesClient().NewListPager("res9290", "sto1590", &armstorage.FileSharesClientListOptions{Maxpagesize: nil,
		Filter: nil,
		Expand: to.Ptr("deleted"),
	})
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
		// page.FileShareItems = armstorage.FileShareItems{
		// 	Value: []*armstorage.FileShareItem{
		// 		{
		// 			Name: to.Ptr("share1644"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/fileServices/shares"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9290/providers/Microsoft.Storage/storageAccounts/sto1590/fileServices/default/shares/share1644_1234567890"),
		// 			Etag: to.Ptr("\"0x8D589847D51C7DE\""),
		// 			Properties: &armstorage.FileShareProperties{
		// 				Deleted: to.Ptr(true),
		// 				DeletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-14T08:20:47Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-14T08:20:47Z"); return t}()),
		// 				RemainingRetentionDays: to.Ptr[int32](30),
		// 				ShareQuota: to.Ptr[int32](1024),
		// 				Version: to.Ptr("1234567890"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("share4052"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/fileServices/shares"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9290/providers/Microsoft.Storage/storageAccounts/sto1590/fileServices/default/shares/share4052"),
		// 			Etag: to.Ptr("\"0x8D589847DAB5AF9\""),
		// 			Properties: &armstorage.FileShareProperties{
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-14T08:20:47Z"); return t}()),
		// 				ShareQuota: to.Ptr[int32](1024),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b32e1896f30e6ea155449cb49719a6286e32b961/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/FileShareSnapshotsList.json
func ExampleFileSharesClient_NewListPager_listShareSnapshots() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFileSharesClient().NewListPager("res9290", "sto1590", &armstorage.FileSharesClientListOptions{Maxpagesize: nil,
		Filter: nil,
		Expand: to.Ptr("snapshots"),
	})
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
		// page.FileShareItems = armstorage.FileShareItems{
		// 	Value: []*armstorage.FileShareItem{
		// 		{
		// 			Name: to.Ptr("share4052"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/fileServices/shares"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9290/providers/Microsoft.Storage/storageAccounts/sto1590/fileServices/default/shares/share4052"),
		// 			Etag: to.Ptr("\"0x8D589847DAB5AF9\""),
		// 			Properties: &armstorage.FileShareProperties{
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-26T05:47:05.0000000Z"); return t}()),
		// 				ShareQuota: to.Ptr[int32](1024),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("share4052"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/fileServices/shares"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9290/providers/Microsoft.Storage/storageAccounts/sto1590/fileServices/default/shares/share4052"),
		// 			Etag: to.Ptr("\"0x8D589847DAB5AF9\""),
		// 			Properties: &armstorage.FileShareProperties{
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-26T05:47:05.0000000Z"); return t}()),
		// 				ShareQuota: to.Ptr[int32](1024),
		// 				SnapshotTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-26T05:48:07.0000000Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b32e1896f30e6ea155449cb49719a6286e32b961/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/FileSharesList.json
func ExampleFileSharesClient_NewListPager_listShares() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFileSharesClient().NewListPager("res9290", "sto1590", &armstorage.FileSharesClientListOptions{Maxpagesize: nil,
		Filter: nil,
		Expand: nil,
	})
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
		// page.FileShareItems = armstorage.FileShareItems{
		// 	Value: []*armstorage.FileShareItem{
		// 		{
		// 			Name: to.Ptr("share1644"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/fileServices/shares"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9290/providers/Microsoft.Storage/storageAccounts/sto1590/fileServices/default/shares/share1644"),
		// 			Etag: to.Ptr("\"0x8D589847D51C7DE\""),
		// 			Properties: &armstorage.FileShareProperties{
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-14T08:20:47Z"); return t}()),
		// 				ShareQuota: to.Ptr[int32](1024),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("share4052"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/fileServices/shares"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9290/providers/Microsoft.Storage/storageAccounts/sto1590/fileServices/default/shares/share4052"),
		// 			Etag: to.Ptr("\"0x8D589847DAB5AF9\""),
		// 			Properties: &armstorage.FileShareProperties{
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-14T08:20:47Z"); return t}()),
		// 				ShareQuota: to.Ptr[int32](1024),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b32e1896f30e6ea155449cb49719a6286e32b961/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/FileSharesPut_NFS.json
func ExampleFileSharesClient_Create_createNfsShares() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFileSharesClient().Create(ctx, "res346", "sto666", "share1235", armstorage.FileShare{
		FileShareProperties: &armstorage.FileShareProperties{
			EnabledProtocols: to.Ptr(armstorage.EnabledProtocolsNFS),
		},
	}, &armstorage.FileSharesClientCreateOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FileShare = armstorage.FileShare{
	// 	Name: to.Ptr("share1235"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/fileServices/shares"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res346/providers/Microsoft.Storage/storageAccounts/sto666/fileServices/default/shares/share1235"),
	// 	FileShareProperties: &armstorage.FileShareProperties{
	// 		EnabledProtocols: to.Ptr(armstorage.EnabledProtocolsNFS),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b32e1896f30e6ea155449cb49719a6286e32b961/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/FileSharesPut.json
func ExampleFileSharesClient_Create_putShares() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFileSharesClient().Create(ctx, "res3376", "sto328", "share6185", armstorage.FileShare{}, &armstorage.FileSharesClientCreateOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FileShare = armstorage.FileShare{
	// 	Name: to.Ptr("share6185"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/fileServices/shares"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res3376/providers/Microsoft.Storage/storageAccounts/sto328/fileServices/default/shares/share6185"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b32e1896f30e6ea155449cb49719a6286e32b961/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/FileSharesPut_AccessTier.json
func ExampleFileSharesClient_Create_putSharesWithAccessTier() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFileSharesClient().Create(ctx, "res346", "sto666", "share1235", armstorage.FileShare{
		FileShareProperties: &armstorage.FileShareProperties{
			AccessTier: to.Ptr(armstorage.ShareAccessTierHot),
		},
	}, &armstorage.FileSharesClientCreateOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FileShare = armstorage.FileShare{
	// 	Name: to.Ptr("share1235"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/fileServices/shares"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res346/providers/Microsoft.Storage/storageAccounts/sto666/fileServices/default/shares/share1235"),
	// 	FileShareProperties: &armstorage.FileShareProperties{
	// 		AccessTier: to.Ptr(armstorage.ShareAccessTierHot),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b32e1896f30e6ea155449cb49719a6286e32b961/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/FileShareAclsPatch.json
func ExampleFileSharesClient_Update_updateShareAcls() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFileSharesClient().Update(ctx, "res3376", "sto328", "share6185", armstorage.FileShare{
		FileShareProperties: &armstorage.FileShareProperties{
			SignedIdentifiers: []*armstorage.SignedIdentifier{
				{
					AccessPolicy: &armstorage.AccessPolicy{
						ExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-01T08:49:37.0000000Z"); return t }()),
						Permission: to.Ptr("rwd"),
						StartTime:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T08:49:37.0000000Z"); return t }()),
					},
					ID: to.Ptr("MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI"),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FileShare = armstorage.FileShare{
	// 	Name: to.Ptr("share6185"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/fileServices/shares"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res3376/providers/Microsoft.Storage/storageAccounts/sto328/fileServices/default/shares/share6185"),
	// 	FileShareProperties: &armstorage.FileShareProperties{
	// 		SignedIdentifiers: []*armstorage.SignedIdentifier{
	// 			{
	// 				AccessPolicy: &armstorage.AccessPolicy{
	// 					ExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-01T08:49:37.0000000Z"); return t}()),
	// 					Permission: to.Ptr("rwd"),
	// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T08:49:37.0000000Z"); return t}()),
	// 				},
	// 				ID: to.Ptr("MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI"),
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b32e1896f30e6ea155449cb49719a6286e32b961/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/FileSharesPatch.json
func ExampleFileSharesClient_Update_updateShares() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFileSharesClient().Update(ctx, "res3376", "sto328", "share6185", armstorage.FileShare{
		FileShareProperties: &armstorage.FileShareProperties{
			Metadata: map[string]*string{
				"type": to.Ptr("image"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FileShare = armstorage.FileShare{
	// 	Name: to.Ptr("share6185"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/fileServices/shares"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res3376/providers/Microsoft.Storage/storageAccounts/sto328/fileServices/default/shares/share6185"),
	// 	FileShareProperties: &armstorage.FileShareProperties{
	// 		Metadata: map[string]*string{
	// 			"type": to.Ptr("image"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b32e1896f30e6ea155449cb49719a6286e32b961/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/FileSharesGet_Stats.json
func ExampleFileSharesClient_Get_getShareStats() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFileSharesClient().Get(ctx, "res9871", "sto6217", "share1634", &armstorage.FileSharesClientGetOptions{Expand: to.Ptr("stats"),
		XMSSnapshot: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FileShare = armstorage.FileShare{
	// 	Name: to.Ptr("share1634"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/fileServices/shares"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9871/providers/Microsoft.Storage/storageAccounts/sto6217/fileServices/default/shares/share1634"),
	// 	Etag: to.Ptr("\"0x8D592D74CC20EBA\""),
	// 	FileShareProperties: &armstorage.FileShareProperties{
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-26T05:06:14Z"); return t}()),
	// 		ShareQuota: to.Ptr[int32](1024),
	// 		ShareUsageBytes: to.Ptr[int64](652945),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b32e1896f30e6ea155449cb49719a6286e32b961/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/FileSharesGet.json
func ExampleFileSharesClient_Get_getShares() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFileSharesClient().Get(ctx, "res9871", "sto6217", "share1634", &armstorage.FileSharesClientGetOptions{Expand: nil,
		XMSSnapshot: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FileShare = armstorage.FileShare{
	// 	Name: to.Ptr("share1634"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/fileServices/shares"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9871/providers/Microsoft.Storage/storageAccounts/sto6217/fileServices/default/shares/share1634"),
	// 	Etag: to.Ptr("\"0x8D592D74CC20EBA\""),
	// 	FileShareProperties: &armstorage.FileShareProperties{
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-26T05:06:14Z"); return t}()),
	// 		ShareQuota: to.Ptr[int32](1024),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b32e1896f30e6ea155449cb49719a6286e32b961/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/FileSharesDelete.json
func ExampleFileSharesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewFileSharesClient().Delete(ctx, "res4079", "sto4506", "share9689", &armstorage.FileSharesClientDeleteOptions{XMSSnapshot: nil,
		Include: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b32e1896f30e6ea155449cb49719a6286e32b961/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/FileSharesRestore.json
func ExampleFileSharesClient_Restore() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewFileSharesClient().Restore(ctx, "res3376", "sto328", "share1249", armstorage.DeletedShare{
		DeletedShareName:    to.Ptr("share1249"),
		DeletedShareVersion: to.Ptr("1234567890"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b32e1896f30e6ea155449cb49719a6286e32b961/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/FileSharesLease_Acquire.json
func ExampleFileSharesClient_Lease_acquireALeaseOnAShare() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFileSharesClient().Lease(ctx, "res3376", "sto328", "share124", &armstorage.FileSharesClientLeaseOptions{XMSSnapshot: nil,
		Parameters: &armstorage.LeaseShareRequest{
			Action:        to.Ptr(armstorage.LeaseShareActionAcquire),
			LeaseDuration: to.Ptr[int32](-1),
		},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LeaseShareResponse = armstorage.LeaseShareResponse{
	// 	LeaseID: to.Ptr("8698f513-fa75-44a1-b8eb-30ba336af27d"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b32e1896f30e6ea155449cb49719a6286e32b961/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/FileSharesLease_Break.json
func ExampleFileSharesClient_Lease_breakALeaseOnAShare() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFileSharesClient().Lease(ctx, "res3376", "sto328", "share12", &armstorage.FileSharesClientLeaseOptions{XMSSnapshot: nil,
		Parameters: &armstorage.LeaseShareRequest{
			Action:  to.Ptr(armstorage.LeaseShareActionBreak),
			LeaseID: to.Ptr("8698f513-fa75-44a1-b8eb-30ba336af27d"),
		},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LeaseShareResponse = armstorage.LeaseShareResponse{
	// 	LeaseTimeSeconds: to.Ptr("0"),
	// }
}
