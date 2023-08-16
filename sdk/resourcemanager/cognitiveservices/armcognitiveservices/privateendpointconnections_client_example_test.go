//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2023-05-01/examples/ListPrivateEndpointConnections.json
func ExamplePrivateEndpointConnectionsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().List(ctx, "res6977", "sto2527", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnectionListResult = armcognitiveservices.PrivateEndpointConnectionListResult{
	// 	Value: []*armcognitiveservices.PrivateEndpointConnection{
	// 		{
	// 			Name: to.Ptr("{privateEndpointConnectionName}"),
	// 			Type: to.Ptr("Microsoft.CognitiveServices/accounts/privateEndpointConnections"),
	// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res7231/providers/Microsoft.CognitiveServices/accounts/sto288/privateEndpointConnections/{privateEndpointConnectionName}"),
	// 			Properties: &armcognitiveservices.PrivateEndpointConnectionProperties{
	// 				PrivateEndpoint: &armcognitiveservices.PrivateEndpoint{
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res7231/providers/Microsoft.Network/privateEndpoints/petest01"),
	// 				},
	// 				PrivateLinkServiceConnectionState: &armcognitiveservices.PrivateLinkServiceConnectionState{
	// 					Description: to.Ptr("Auto-Approved"),
	// 					ActionsRequired: to.Ptr("None"),
	// 					Status: to.Ptr(armcognitiveservices.PrivateEndpointServiceConnectionStatusApproved),
	// 				},
	// 			},
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2023-05-01/examples/GetPrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().Get(ctx, "res6977", "sto2527", "{privateEndpointConnectionName}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnection = armcognitiveservices.PrivateEndpointConnection{
	// 	Name: to.Ptr("{privateEndpointConnectionName}"),
	// 	Type: to.Ptr("Microsoft.CognitiveServices/accounts/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res7231/providers/Microsoft.CognitiveServices/accounts/sto288/privateEndpointConnections/{privateEndpointConnectionName}"),
	// 	Properties: &armcognitiveservices.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armcognitiveservices.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res7231/providers/Microsoft.Network/privateEndpoints/petest01"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armcognitiveservices.PrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("Auto-Approved"),
	// 			ActionsRequired: to.Ptr("None"),
	// 			Status: to.Ptr(armcognitiveservices.PrivateEndpointServiceConnectionStatusApproved),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2023-05-01/examples/PutPrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionsClient().BeginCreateOrUpdate(ctx, "res7687", "sto9699", "{privateEndpointConnectionName}", armcognitiveservices.PrivateEndpointConnection{
		Properties: &armcognitiveservices.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armcognitiveservices.PrivateLinkServiceConnectionState{
				Description: to.Ptr("Auto-Approved"),
				Status:      to.Ptr(armcognitiveservices.PrivateEndpointServiceConnectionStatusApproved),
			},
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
	// res.PrivateEndpointConnection = armcognitiveservices.PrivateEndpointConnection{
	// 	Name: to.Ptr("{privateEndpointConnectionName}"),
	// 	Type: to.Ptr("Microsoft.CognitiveServices/accounts/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res7231/providers/Microsoft.CognitiveServices/accounts/sto288/privateEndpointConnections/{privateEndpointConnectionName}"),
	// 	Properties: &armcognitiveservices.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armcognitiveservices.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res7231/providers/Microsoft.Network/privateEndpoints/petest01"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armcognitiveservices.PrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("Auto-Approved"),
	// 			ActionsRequired: to.Ptr("None"),
	// 			Status: to.Ptr(armcognitiveservices.PrivateEndpointServiceConnectionStatusApproved),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2023-05-01/examples/DeletePrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionsClient().BeginDelete(ctx, "res6977", "sto2527", "{privateEndpointConnectionName}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
