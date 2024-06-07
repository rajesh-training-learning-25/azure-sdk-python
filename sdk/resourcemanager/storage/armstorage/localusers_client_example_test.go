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

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b22c642b361e6d6e7d72a2347a09b0bcf6075d70/specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/LocalUsersList.json
func ExampleLocalUsersClient_NewListPager_listLocalUsers() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLocalUsersClient().NewListPager("res6977", "sto2527", &armstorage.LocalUsersClientListOptions{Maxpagesize: nil,
		Filter:  nil,
		Include: nil,
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
		// page.LocalUsers = armstorage.LocalUsers{
		// 	Value: []*armstorage.LocalUser{
		// 		{
		// 			Name: to.Ptr("user1"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/localUsers"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res6977/providers/Microsoft.Storage/storageAccounts/sto2527/loalUsers/user1"),
		// 			Properties: &armstorage.LocalUserProperties{
		// 				AllowACLAuthorization: to.Ptr(true),
		// 				GroupID: to.Ptr[int32](2000),
		// 				HasSharedKey: to.Ptr(true),
		// 				HasSSHKey: to.Ptr(true),
		// 				HasSSHPassword: to.Ptr(true),
		// 				HomeDirectory: to.Ptr("homedirectory"),
		// 				PermissionScopes: []*armstorage.PermissionScope{
		// 					{
		// 						Permissions: to.Ptr("rwd"),
		// 						ResourceName: to.Ptr("share1"),
		// 						Service: to.Ptr("file"),
		// 					},
		// 					{
		// 						Permissions: to.Ptr("rw"),
		// 						ResourceName: to.Ptr("share2"),
		// 						Service: to.Ptr("file"),
		// 				}},
		// 				Sid: to.Ptr("S-1-2-0-125132-153423-36235-1000"),
		// 				UserID: to.Ptr[int32](1000),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("user2"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/localUsers"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9290/providers/Microsoft.Storage/storageAccounts/sto1590/loalUsers/user2"),
		// 			Properties: &armstorage.LocalUserProperties{
		// 				AllowACLAuthorization: to.Ptr(true),
		// 				GroupID: to.Ptr[int32](2000),
		// 				HasSharedKey: to.Ptr(true),
		// 				HasSSHKey: to.Ptr(false),
		// 				HasSSHPassword: to.Ptr(true),
		// 				PermissionScopes: []*armstorage.PermissionScope{
		// 					{
		// 						Permissions: to.Ptr("rw"),
		// 						ResourceName: to.Ptr("resourcename"),
		// 						Service: to.Ptr("blob"),
		// 				}},
		// 				Sid: to.Ptr("S-1-2-0-533672-235636-66334-1001"),
		// 				UserID: to.Ptr[int32](1001),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b22c642b361e6d6e7d72a2347a09b0bcf6075d70/specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/LocalUsersListNFSv3Enabled.json
func ExampleLocalUsersClient_NewListPager_listNfSv3EnabledLocalUsers() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLocalUsersClient().NewListPager("res6977", "sto2527", &armstorage.LocalUsersClientListOptions{Maxpagesize: nil,
		Filter:  nil,
		Include: to.Ptr(armstorage.ListLocalUserIncludeParamNfsv3),
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
		// page.LocalUsers = armstorage.LocalUsers{
		// 	Value: []*armstorage.LocalUser{
		// 		{
		// 			Name: to.Ptr("user1"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/localUsers"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res6977/providers/Microsoft.Storage/storageAccounts/sto2527/loalUsers/user1"),
		// 			Properties: &armstorage.LocalUserProperties{
		// 				AllowACLAuthorization: to.Ptr(true),
		// 				ExtendedGroups: []*int32{
		// 					to.Ptr[int32](1001),
		// 					to.Ptr[int32](1005),
		// 					to.Ptr[int32](2005)},
		// 					GroupID: to.Ptr[int32](2000),
		// 					HasSharedKey: to.Ptr(true),
		// 					HasSSHKey: to.Ptr(true),
		// 					HasSSHPassword: to.Ptr(true),
		// 					HomeDirectory: to.Ptr("homedirectory"),
		// 					IsNFSv3Enabled: to.Ptr(true),
		// 					PermissionScopes: []*armstorage.PermissionScope{
		// 						{
		// 							Permissions: to.Ptr("rwd"),
		// 							ResourceName: to.Ptr("share1"),
		// 							Service: to.Ptr("file"),
		// 						},
		// 						{
		// 							Permissions: to.Ptr("rw"),
		// 							ResourceName: to.Ptr("share2"),
		// 							Service: to.Ptr("file"),
		// 					}},
		// 					Sid: to.Ptr("S-1-2-0-125132-153423-36235-1000"),
		// 					UserID: to.Ptr[int32](1000),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("user2"),
		// 				Type: to.Ptr("Microsoft.Storage/storageAccounts/localUsers"),
		// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9290/providers/Microsoft.Storage/storageAccounts/sto1590/loalUsers/user2"),
		// 				Properties: &armstorage.LocalUserProperties{
		// 					AllowACLAuthorization: to.Ptr(true),
		// 					ExtendedGroups: []*int32{
		// 						to.Ptr[int32](1001),
		// 						to.Ptr[int32](1005),
		// 						to.Ptr[int32](2005)},
		// 						GroupID: to.Ptr[int32](2000),
		// 						HasSharedKey: to.Ptr(true),
		// 						HasSSHKey: to.Ptr(false),
		// 						HasSSHPassword: to.Ptr(true),
		// 						IsNFSv3Enabled: to.Ptr(true),
		// 						PermissionScopes: []*armstorage.PermissionScope{
		// 							{
		// 								Permissions: to.Ptr("rw"),
		// 								ResourceName: to.Ptr("resourcename"),
		// 								Service: to.Ptr("blob"),
		// 						}},
		// 						Sid: to.Ptr("S-1-2-0-533672-235636-66334-1001"),
		// 						UserID: to.Ptr[int32](1001),
		// 					},
		// 			}},
		// 		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b22c642b361e6d6e7d72a2347a09b0bcf6075d70/specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/LocalUserGet.json
func ExampleLocalUsersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLocalUsersClient().Get(ctx, "res6977", "sto2527", "user1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LocalUser = armstorage.LocalUser{
	// 	Name: to.Ptr("user1"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/localUsers"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res6977/providers/Microsoft.Storage/storageAccounts/sto2527/loalUsers/user1"),
	// 	Properties: &armstorage.LocalUserProperties{
	// 		AllowACLAuthorization: to.Ptr(true),
	// 		ExtendedGroups: []*int32{
	// 			to.Ptr[int32](1001),
	// 			to.Ptr[int32](1005),
	// 			to.Ptr[int32](2005)},
	// 			GroupID: to.Ptr[int32](2000),
	// 			HasSharedKey: to.Ptr(true),
	// 			HasSSHKey: to.Ptr(true),
	// 			HasSSHPassword: to.Ptr(true),
	// 			HomeDirectory: to.Ptr("homedirectory"),
	// 			IsNFSv3Enabled: to.Ptr(true),
	// 			PermissionScopes: []*armstorage.PermissionScope{
	// 				{
	// 					Permissions: to.Ptr("rwd"),
	// 					ResourceName: to.Ptr("share1"),
	// 					Service: to.Ptr("file"),
	// 				},
	// 				{
	// 					Permissions: to.Ptr("rw"),
	// 					ResourceName: to.Ptr("share2"),
	// 					Service: to.Ptr("file"),
	// 			}},
	// 			Sid: to.Ptr("S-1-2-0-125132-153423-36235-1000"),
	// 			UserID: to.Ptr[int32](1000),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b22c642b361e6d6e7d72a2347a09b0bcf6075d70/specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/LocalUserCreate.json
func ExampleLocalUsersClient_CreateOrUpdate_createLocalUser() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLocalUsersClient().CreateOrUpdate(ctx, "res6977", "sto2527", "user1", armstorage.LocalUser{
		Properties: &armstorage.LocalUserProperties{
			AllowACLAuthorization: to.Ptr(true),
			GroupID:               to.Ptr[int32](2000),
			HasSSHPassword:        to.Ptr(true),
			HomeDirectory:         to.Ptr("homedirectory"),
			PermissionScopes: []*armstorage.PermissionScope{
				{
					Permissions:  to.Ptr("rwd"),
					ResourceName: to.Ptr("share1"),
					Service:      to.Ptr("file"),
				},
				{
					Permissions:  to.Ptr("rw"),
					ResourceName: to.Ptr("share2"),
					Service:      to.Ptr("file"),
				}},
			SSHAuthorizedKeys: []*armstorage.SSHPublicKey{
				{
					Description: to.Ptr("key name"),
					Key:         to.Ptr("ssh-rsa keykeykeykeykey="),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LocalUser = armstorage.LocalUser{
	// 	Name: to.Ptr("user1"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/localUsers"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res6977/providers/Microsoft.Storage/storageAccounts/sto2527/loalUsers/user1"),
	// 	Properties: &armstorage.LocalUserProperties{
	// 		AllowACLAuthorization: to.Ptr(true),
	// 		GroupID: to.Ptr[int32](2000),
	// 		HomeDirectory: to.Ptr("homedirectory"),
	// 		PermissionScopes: []*armstorage.PermissionScope{
	// 			{
	// 				Permissions: to.Ptr("rwd"),
	// 				ResourceName: to.Ptr("share1"),
	// 				Service: to.Ptr("file"),
	// 			},
	// 			{
	// 				Permissions: to.Ptr("rw"),
	// 				ResourceName: to.Ptr("share2"),
	// 				Service: to.Ptr("file"),
	// 		}},
	// 		Sid: to.Ptr("S-1-2-0-125132-153423-36235-1000"),
	// 		SSHAuthorizedKeys: []*armstorage.SSHPublicKey{
	// 			{
	// 				Description: to.Ptr("key name"),
	// 				Key: to.Ptr("ssh-rsa keykeykeykeykey="),
	// 		}},
	// 		UserID: to.Ptr[int32](1000),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b22c642b361e6d6e7d72a2347a09b0bcf6075d70/specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/LocalUserCreateNFSv3Enabled.json
func ExampleLocalUsersClient_CreateOrUpdate_createNfSv3EnabledLocalUser() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLocalUsersClient().CreateOrUpdate(ctx, "res6977", "sto2527", "user1", armstorage.LocalUser{
		Properties: &armstorage.LocalUserProperties{
			ExtendedGroups: []*int32{
				to.Ptr[int32](1001),
				to.Ptr[int32](1005),
				to.Ptr[int32](2005)},
			IsNFSv3Enabled: to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LocalUser = armstorage.LocalUser{
	// 	Name: to.Ptr("user1"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/localUsers"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res6977/providers/Microsoft.Storage/storageAccounts/sto2527/loalUsers/user1"),
	// 	Properties: &armstorage.LocalUserProperties{
	// 		AllowACLAuthorization: to.Ptr(true),
	// 		ExtendedGroups: []*int32{
	// 			to.Ptr[int32](1001),
	// 			to.Ptr[int32](1005),
	// 			to.Ptr[int32](2005)},
	// 			GroupID: to.Ptr[int32](2000),
	// 			HomeDirectory: to.Ptr("homedirectory"),
	// 			IsNFSv3Enabled: to.Ptr(true),
	// 			PermissionScopes: []*armstorage.PermissionScope{
	// 				{
	// 					Permissions: to.Ptr("rwd"),
	// 					ResourceName: to.Ptr("share1"),
	// 					Service: to.Ptr("file"),
	// 				},
	// 				{
	// 					Permissions: to.Ptr("rw"),
	// 					ResourceName: to.Ptr("share2"),
	// 					Service: to.Ptr("file"),
	// 			}},
	// 			Sid: to.Ptr("S-1-2-0-125132-153423-36235-1000"),
	// 			SSHAuthorizedKeys: []*armstorage.SSHPublicKey{
	// 				{
	// 					Description: to.Ptr("key name"),
	// 					Key: to.Ptr("ssh-rsa keykeykeykeykey="),
	// 			}},
	// 			UserID: to.Ptr[int32](1000),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b22c642b361e6d6e7d72a2347a09b0bcf6075d70/specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/LocalUserUpdate.json
func ExampleLocalUsersClient_CreateOrUpdate_updateLocalUser() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLocalUsersClient().CreateOrUpdate(ctx, "res6977", "sto2527", "user1", armstorage.LocalUser{
		Properties: &armstorage.LocalUserProperties{
			AllowACLAuthorization: to.Ptr(false),
			ExtendedGroups: []*int32{
				to.Ptr[int32](1001),
				to.Ptr[int32](1005),
				to.Ptr[int32](2005)},
			GroupID:        to.Ptr[int32](3000),
			HasSharedKey:   to.Ptr(false),
			HasSSHKey:      to.Ptr(false),
			HasSSHPassword: to.Ptr(false),
			HomeDirectory:  to.Ptr("homedirectory2"),
			IsNFSv3Enabled: to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LocalUser = armstorage.LocalUser{
	// 	Name: to.Ptr("user1"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/localUsers"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res6977/providers/Microsoft.Storage/storageAccounts/sto2527/loalUsers/user1"),
	// 	Properties: &armstorage.LocalUserProperties{
	// 		AllowACLAuthorization: to.Ptr(false),
	// 		ExtendedGroups: []*int32{
	// 			to.Ptr[int32](1001),
	// 			to.Ptr[int32](1005),
	// 			to.Ptr[int32](2005)},
	// 			GroupID: to.Ptr[int32](3000),
	// 			HasSharedKey: to.Ptr(false),
	// 			HasSSHKey: to.Ptr(false),
	// 			HasSSHPassword: to.Ptr(false),
	// 			HomeDirectory: to.Ptr("homedirectory2"),
	// 			IsNFSv3Enabled: to.Ptr(true),
	// 			Sid: to.Ptr("S-1-2-0-3528686663-1788730862-2791910117-1000"),
	// 			UserID: to.Ptr[int32](1000),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b22c642b361e6d6e7d72a2347a09b0bcf6075d70/specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/LocalUserDelete.json
func ExampleLocalUsersClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewLocalUsersClient().Delete(ctx, "res6977", "sto2527", "user1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b22c642b361e6d6e7d72a2347a09b0bcf6075d70/specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/LocalUserListKeys.json
func ExampleLocalUsersClient_ListKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLocalUsersClient().ListKeys(ctx, "res6977", "sto2527", "user1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LocalUserKeys = armstorage.LocalUserKeys{
	// 	SharedKey: to.Ptr("<REDACTED>"),
	// 	SSHAuthorizedKeys: []*armstorage.SSHPublicKey{
	// 		{
	// 			Description: to.Ptr("key name"),
	// 			Key: to.Ptr("ssh-rsa keykeykeykeykew="),
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b22c642b361e6d6e7d72a2347a09b0bcf6075d70/specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/LocalUserRegeneratePassword.json
func ExampleLocalUsersClient_RegeneratePassword() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLocalUsersClient().RegeneratePassword(ctx, "res6977", "sto2527", "user1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LocalUserRegeneratePasswordResult = armstorage.LocalUserRegeneratePasswordResult{
	// 	SSHPassword: to.Ptr("<REDACTED>"),
	// }
}
