//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-02-01/examples/ManagedHsm_ListPrivateEndpointConnectionsByResource.json
func ExampleMHSMPrivateEndpointConnectionsClient_NewListByResourcePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMHSMPrivateEndpointConnectionsClient().NewListByResourcePager("sample-group", "sample-mhsm", nil)
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
		// page.MHSMPrivateEndpointConnectionsListResult = armkeyvault.MHSMPrivateEndpointConnectionsListResult{
		// 	Value: []*armkeyvault.MHSMPrivateEndpointConnection{
		// 		{
		// 			Name: to.Ptr("sample-pec1"),
		// 			Type: to.Ptr("Microsoft.KeyVault/managedhsms/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/managedhsms/sample-mhsm/privateEndpointConnections/sample-pec1"),
		// 			Etag: to.Ptr(""),
		// 			Properties: &armkeyvault.MHSMPrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armkeyvault.MHSMPrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-1234-000000000000/resourceGroups/sample-group/providers/Microsoft.Network/privateEndpoints/sample-pe1"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armkeyvault.MHSMPrivateLinkServiceConnectionState{
		// 					Description: to.Ptr("This was automatically approved by user1234@contoso.com"),
		// 					ActionsRequired: to.Ptr(armkeyvault.ActionsRequiredNone),
		// 					Status: to.Ptr(armkeyvault.PrivateEndpointServiceConnectionStatusApproved),
		// 				},
		// 				ProvisioningState: to.Ptr(armkeyvault.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sample-pec2"),
		// 			Type: to.Ptr("Microsoft.KeyVault/managedhsms/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/managedhsms/sample-mhsm/privateEndpointConnections/sample-pec2"),
		// 			Etag: to.Ptr(""),
		// 			Properties: &armkeyvault.MHSMPrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armkeyvault.MHSMPrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-1234-000000000000/resourceGroups/sample-group/providers/Microsoft.Network/privateEndpoints/sample-pe2"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armkeyvault.MHSMPrivateLinkServiceConnectionState{
		// 					Description: to.Ptr("This was automatically approved by user1234@contoso.com"),
		// 					ActionsRequired: to.Ptr(armkeyvault.ActionsRequiredNone),
		// 					Status: to.Ptr(armkeyvault.PrivateEndpointServiceConnectionStatusApproved),
		// 				},
		// 				ProvisioningState: to.Ptr(armkeyvault.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-02-01/examples/ManagedHsm_getPrivateEndpointConnection.json
func ExampleMHSMPrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMHSMPrivateEndpointConnectionsClient().Get(ctx, "sample-group", "sample-mhsm", "sample-pec", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MHSMPrivateEndpointConnection = armkeyvault.MHSMPrivateEndpointConnection{
	// 	Name: to.Ptr("sample-pec"),
	// 	Type: to.Ptr("Microsoft.KeyVault/managedhsms/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/managedhsms/sample-mhsm/privateEndpointConnections/sample-pec"),
	// 	Etag: to.Ptr(""),
	// 	Properties: &armkeyvault.MHSMPrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armkeyvault.MHSMPrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-1234-000000000000/resourceGroups/sample-group/providers/Microsoft.Network/privateEndpoints/sample-pe"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armkeyvault.MHSMPrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("This was automatically approved by user1234@contoso.com"),
	// 			ActionsRequired: to.Ptr(armkeyvault.ActionsRequiredNone),
	// 			Status: to.Ptr(armkeyvault.PrivateEndpointServiceConnectionStatusApproved),
	// 		},
	// 		ProvisioningState: to.Ptr(armkeyvault.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-02-01/examples/ManagedHsm_putPrivateEndpointConnection.json
func ExampleMHSMPrivateEndpointConnectionsClient_Put() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMHSMPrivateEndpointConnectionsClient().Put(ctx, "sample-group", "sample-mhsm", "sample-pec", armkeyvault.MHSMPrivateEndpointConnection{
		Properties: &armkeyvault.MHSMPrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armkeyvault.MHSMPrivateLinkServiceConnectionState{
				Description: to.Ptr("My name is Joe and I'm approving this."),
				Status:      to.Ptr(armkeyvault.PrivateEndpointServiceConnectionStatusApproved),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MHSMPrivateEndpointConnection = armkeyvault.MHSMPrivateEndpointConnection{
	// 	Name: to.Ptr("sample-pec"),
	// 	Type: to.Ptr("Microsoft.KeyVault/managedhsms/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/managedhsms/sample-mhsm/privateEndpointConnections/sample-pec"),
	// 	Properties: &armkeyvault.MHSMPrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armkeyvault.MHSMPrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-1234-000000000000/resourceGroups/sample-group/providers/Microsoft.Network/privateEndpoints/sample-pe"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armkeyvault.MHSMPrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("My name is Joe and I'm approving this."),
	// 			ActionsRequired: to.Ptr(armkeyvault.ActionsRequiredNone),
	// 			Status: to.Ptr(armkeyvault.PrivateEndpointServiceConnectionStatusApproved),
	// 		},
	// 		ProvisioningState: to.Ptr(armkeyvault.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-02-01/examples/ManagedHsm_deletePrivateEndpointConnection.json
func ExampleMHSMPrivateEndpointConnectionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMHSMPrivateEndpointConnectionsClient().BeginDelete(ctx, "sample-group", "sample-mhsm", "sample-pec", nil)
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
	// res.MHSMPrivateEndpointConnection = armkeyvault.MHSMPrivateEndpointConnection{
	// 	Name: to.Ptr("sample-pec"),
	// 	Type: to.Ptr("Microsoft.KeyVault/managedhsms/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/managedhsms/sample-vault/privateEndpointConnections/sample-pec"),
	// 	Properties: &armkeyvault.MHSMPrivateEndpointConnectionProperties{
	// 		ProvisioningState: to.Ptr(armkeyvault.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 	},
	// }
}
