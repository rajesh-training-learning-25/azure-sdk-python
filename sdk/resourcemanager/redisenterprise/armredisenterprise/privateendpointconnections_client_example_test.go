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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a8698bb86e66e2d29ce5e8987b6aaa8fc7f7f04b/specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2024-03-01-preview/examples/RedisEnterpriseListPrivateEndpointConnections.json
func ExamplePrivateEndpointConnectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionsClient().NewListPager("rg1", "cache1", nil)
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
		// page.PrivateEndpointConnectionListResult = armredisenterprise.PrivateEndpointConnectionListResult{
		// 	Value: []*armredisenterprise.PrivateEndpointConnection{
		// 		{
		// 			Name: to.Ptr("pectest01"),
		// 			Type: to.Ptr("Microsoft.Cache/redisEnterprise/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1/privateEndpointConnections/pectest01"),
		// 			Properties: &armredisenterprise.PrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armredisenterprise.PrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/petest01"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armredisenterprise.PrivateLinkServiceConnectionState{
		// 					Description: to.Ptr("Auto-Approved"),
		// 					ActionsRequired: to.Ptr("None"),
		// 					Status: to.Ptr(armredisenterprise.PrivateEndpointServiceConnectionStatusApproved),
		// 				},
		// 				ProvisioningState: to.Ptr(armredisenterprise.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("pectest01"),
		// 			Type: to.Ptr("Microsoft.Cache/redisEnterprise/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1/privateEndpointConnections/pectest01"),
		// 			Properties: &armredisenterprise.PrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armredisenterprise.PrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/petest01"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armredisenterprise.PrivateLinkServiceConnectionState{
		// 					Description: to.Ptr("Auto-Approved"),
		// 					ActionsRequired: to.Ptr("None"),
		// 					Status: to.Ptr(armredisenterprise.PrivateEndpointServiceConnectionStatusApproved),
		// 				},
		// 				ProvisioningState: to.Ptr(armredisenterprise.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a8698bb86e66e2d29ce5e8987b6aaa8fc7f7f04b/specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2024-03-01-preview/examples/RedisEnterpriseGetPrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().Get(ctx, "rg1", "cache1", "pectest01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnection = armredisenterprise.PrivateEndpointConnection{
	// 	Name: to.Ptr("pectest01"),
	// 	Type: to.Ptr("Microsoft.Cache/redisEnterprise/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1/privateEndpointConnections/pectest01"),
	// 	Properties: &armredisenterprise.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armredisenterprise.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/petest01"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armredisenterprise.PrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("Auto-Approved"),
	// 			ActionsRequired: to.Ptr("None"),
	// 			Status: to.Ptr(armredisenterprise.PrivateEndpointServiceConnectionStatusApproved),
	// 		},
	// 		ProvisioningState: to.Ptr(armredisenterprise.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a8698bb86e66e2d29ce5e8987b6aaa8fc7f7f04b/specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2024-03-01-preview/examples/RedisEnterprisePutPrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionsClient_BeginPut() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionsClient().BeginPut(ctx, "rg1", "cache1", "pectest01", armredisenterprise.PrivateEndpointConnection{
		Properties: &armredisenterprise.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armredisenterprise.PrivateLinkServiceConnectionState{
				Description: to.Ptr("Auto-Approved"),
				Status:      to.Ptr(armredisenterprise.PrivateEndpointServiceConnectionStatusApproved),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a8698bb86e66e2d29ce5e8987b6aaa8fc7f7f04b/specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2024-03-01-preview/examples/RedisEnterpriseDeletePrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionsClient().BeginDelete(ctx, "rg1", "cache1", "pectest01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
