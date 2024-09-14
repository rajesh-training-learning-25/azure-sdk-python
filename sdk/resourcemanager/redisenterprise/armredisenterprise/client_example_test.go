//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armredisenterprise_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a287afb3721dee0d88f11502ec123470bc52a28/specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2024-09-01-preview/examples/RedisEnterpriseCreate.json
func ExampleClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginCreate(ctx, "rg1", "cache1", armredisenterprise.Cluster{
		Location: to.Ptr("West US"),
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
		},
		Identity: &armredisenterprise.ManagedServiceIdentity{
			Type: to.Ptr(armredisenterprise.ManagedServiceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armredisenterprise.UserAssignedIdentity{
				"/subscriptions/your-subscription/resourceGroups/your-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/your-identity": {},
			},
		},
		Properties: &armredisenterprise.ClusterProperties{
			Encryption: &armredisenterprise.ClusterPropertiesEncryption{
				CustomerManagedKeyEncryption: &armredisenterprise.ClusterPropertiesEncryptionCustomerManagedKeyEncryption{
					KeyEncryptionKeyIdentity: &armredisenterprise.ClusterPropertiesEncryptionCustomerManagedKeyEncryptionKeyIdentity{
						IdentityType:                   to.Ptr(armredisenterprise.CmkIdentityTypeUserAssignedIdentity),
						UserAssignedIdentityResourceID: to.Ptr("/subscriptions/your-subscription/resourceGroups/your-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/your-identity"),
					},
					KeyEncryptionKeyURL: to.Ptr("https://your-kv.vault.azure.net/keys/your-key/your-key-version"),
				},
			},
			MinimumTLSVersion: to.Ptr(armredisenterprise.TLSVersionOne2),
		},
		SKU: &armredisenterprise.SKU{
			Name:     to.Ptr(armredisenterprise.SKUNameEnterpriseFlashF300),
			Capacity: to.Ptr[int32](3),
		},
		Zones: []*string{
			to.Ptr("1"),
			to.Ptr("2"),
			to.Ptr("3")},
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
	// res.Cluster = armredisenterprise.Cluster{
	// 	Name: to.Ptr("cache1"),
	// 	Type: to.Ptr("Microsoft.Cache/redisEnterprise"),
	// 	ID: to.Ptr("/subscriptions/e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 	},
	// 	Identity: &armredisenterprise.ManagedServiceIdentity{
	// 		Type: to.Ptr(armredisenterprise.ManagedServiceIdentityTypeUserAssigned),
	// 		UserAssignedIdentities: map[string]*armredisenterprise.UserAssignedIdentity{
	// 			"/subscriptions/your-subscription/resourceGroups/your-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/your-identity": &armredisenterprise.UserAssignedIdentity{
	// 				ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			},
	// 		},
	// 	},
	// 	Properties: &armredisenterprise.ClusterProperties{
	// 		HostName: to.Ptr("cache1.westus.something.azure.net"),
	// 		MinimumTLSVersion: to.Ptr(armredisenterprise.TLSVersionOne2),
	// 		ProvisioningState: to.Ptr(armredisenterprise.ProvisioningStateSucceeded),
	// 		RedisVersion: to.Ptr("5"),
	// 		ResourceState: to.Ptr(armredisenterprise.ResourceStateRunning),
	// 	},
	// 	SKU: &armredisenterprise.SKU{
	// 		Name: to.Ptr(armredisenterprise.SKUNameEnterpriseFlashF300),
	// 		Capacity: to.Ptr[int32](3),
	// 	},
	// 	Zones: []*string{
	// 		to.Ptr("1"),
	// 		to.Ptr("2"),
	// 		to.Ptr("3")},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a287afb3721dee0d88f11502ec123470bc52a28/specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2024-09-01-preview/examples/RedisEnterpriseUpdate.json
func ExampleClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginUpdate(ctx, "rg1", "cache1", armredisenterprise.ClusterUpdate{
		Properties: &armredisenterprise.ClusterProperties{
			MinimumTLSVersion: to.Ptr(armredisenterprise.TLSVersionOne2),
		},
		SKU: &armredisenterprise.SKU{
			Name:     to.Ptr(armredisenterprise.SKUNameEnterpriseFlashF300),
			Capacity: to.Ptr[int32](9),
		},
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
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
	// res.Cluster = armredisenterprise.Cluster{
	// 	Name: to.Ptr("cache1"),
	// 	Type: to.Ptr("Microsoft.Cache/redisEnterprise"),
	// 	ID: to.Ptr("/subscriptions/e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 	},
	// 	Identity: &armredisenterprise.ManagedServiceIdentity{
	// 		Type: to.Ptr(armredisenterprise.ManagedServiceIdentityTypeNone),
	// 	},
	// 	Properties: &armredisenterprise.ClusterProperties{
	// 		Encryption: &armredisenterprise.ClusterPropertiesEncryption{
	// 		},
	// 		HostName: to.Ptr("cache1.westus.something.azure.com"),
	// 		MinimumTLSVersion: to.Ptr(armredisenterprise.TLSVersionOne2),
	// 		ProvisioningState: to.Ptr(armredisenterprise.ProvisioningStateSucceeded),
	// 		RedisVersion: to.Ptr("5"),
	// 		ResourceState: to.Ptr(armredisenterprise.ResourceStateUpdating),
	// 	},
	// 	SKU: &armredisenterprise.SKU{
	// 		Name: to.Ptr(armredisenterprise.SKUNameEnterpriseFlashF300),
	// 		Capacity: to.Ptr[int32](9),
	// 	},
	// 	Zones: []*string{
	// 		to.Ptr("1"),
	// 		to.Ptr("2"),
	// 		to.Ptr("3")},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a287afb3721dee0d88f11502ec123470bc52a28/specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2024-09-01-preview/examples/RedisEnterpriseDelete.json
func ExampleClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginDelete(ctx, "rg1", "cache1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a287afb3721dee0d88f11502ec123470bc52a28/specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2024-09-01-preview/examples/RedisEnterpriseGet.json
func ExampleClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().Get(ctx, "rg1", "cache1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Cluster = armredisenterprise.Cluster{
	// 	Name: to.Ptr("cache1"),
	// 	Type: to.Ptr("Microsoft.Cache/redisEnterprise"),
	// 	ID: to.Ptr("/subscriptions/e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armredisenterprise.ClusterProperties{
	// 		HostName: to.Ptr("cache1.westus.something.azure.com"),
	// 		MinimumTLSVersion: to.Ptr(armredisenterprise.TLSVersionOne2),
	// 		PrivateEndpointConnections: []*armredisenterprise.PrivateEndpointConnection{
	// 			{
	// 				ID: to.Ptr("/subscriptions/e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1/privateEndpointConnections/cachePec"),
	// 				Properties: &armredisenterprise.PrivateEndpointConnectionProperties{
	// 					PrivateEndpoint: &armredisenterprise.PrivateEndpoint{
	// 						ID: to.Ptr("/subscriptions/e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/cachePe"),
	// 					},
	// 					PrivateLinkServiceConnectionState: &armredisenterprise.PrivateLinkServiceConnectionState{
	// 						Description: to.Ptr("Please approve my connection"),
	// 						ActionsRequired: to.Ptr("None"),
	// 						Status: to.Ptr(armredisenterprise.PrivateEndpointServiceConnectionStatusApproved),
	// 					},
	// 				},
	// 		}},
	// 		ProvisioningState: to.Ptr(armredisenterprise.ProvisioningStateSucceeded),
	// 		RedisVersion: to.Ptr("6"),
	// 		ResourceState: to.Ptr(armredisenterprise.ResourceStateRunning),
	// 	},
	// 	SKU: &armredisenterprise.SKU{
	// 		Name: to.Ptr(armredisenterprise.SKUNameEnterpriseFlashF300),
	// 		Capacity: to.Ptr[int32](3),
	// 	},
	// 	Zones: []*string{
	// 		to.Ptr("1"),
	// 		to.Ptr("2"),
	// 		to.Ptr("3")},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a287afb3721dee0d88f11502ec123470bc52a28/specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2024-09-01-preview/examples/RedisEnterpriseListByResourceGroup.json
func ExampleClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListByResourceGroupPager("rg1", nil)
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
		// page.ClusterList = armredisenterprise.ClusterList{
		// 	Value: []*armredisenterprise.Cluster{
		// 		{
		// 			Name: to.Ptr("cache1"),
		// 			Type: to.Ptr("Microsoft.Cache/redisEnterprise"),
		// 			ID: to.Ptr("/subscriptions/e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armredisenterprise.ClusterProperties{
		// 				HostName: to.Ptr("cache1.westus.something.azure.com"),
		// 				MinimumTLSVersion: to.Ptr(armredisenterprise.TLSVersionOne2),
		// 				ProvisioningState: to.Ptr(armredisenterprise.ProvisioningStateSucceeded),
		// 				RedisVersion: to.Ptr("5"),
		// 				ResourceState: to.Ptr(armredisenterprise.ResourceStateRunning),
		// 			},
		// 			SKU: &armredisenterprise.SKU{
		// 				Name: to.Ptr(armredisenterprise.SKUNameEnterpriseFlashF300),
		// 				Capacity: to.Ptr[int32](3),
		// 			},
		// 			Zones: []*string{
		// 				to.Ptr("1"),
		// 				to.Ptr("2"),
		// 				to.Ptr("3")},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a287afb3721dee0d88f11502ec123470bc52a28/specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2024-09-01-preview/examples/RedisEnterpriseList.json
func ExampleClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListPager(nil)
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
		// page.ClusterList = armredisenterprise.ClusterList{
		// 	Value: []*armredisenterprise.Cluster{
		// 		{
		// 			Name: to.Ptr("cache1"),
		// 			Type: to.Ptr("Microsoft.Cache/redisEnterprise"),
		// 			ID: to.Ptr("/subscriptions/e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armredisenterprise.ClusterProperties{
		// 				HostName: to.Ptr("cache1.westus.something.azure.com"),
		// 				MinimumTLSVersion: to.Ptr(armredisenterprise.TLSVersionOne2),
		// 				ProvisioningState: to.Ptr(armredisenterprise.ProvisioningStateSucceeded),
		// 				RedisVersion: to.Ptr("6"),
		// 				ResourceState: to.Ptr(armredisenterprise.ResourceStateRunning),
		// 			},
		// 			SKU: &armredisenterprise.SKU{
		// 				Name: to.Ptr(armredisenterprise.SKUNameEnterpriseFlashF300),
		// 				Capacity: to.Ptr[int32](3),
		// 			},
		// 	}},
		// }
	}
}
