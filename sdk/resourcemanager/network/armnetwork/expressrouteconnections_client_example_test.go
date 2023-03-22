//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/network/resource-manager/Microsoft.Network/stable/2022-09-01/examples/ExpressRouteConnectionCreate.json
func ExampleExpressRouteConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewExpressRouteConnectionsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "resourceGroupName", "gateway-2", "connectionName", armnetwork.ExpressRouteConnection{
		ID:   to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteGateways/gateway-2/expressRouteConnections/connectionName"),
		Name: to.Ptr("connectionName"),
		Properties: &armnetwork.ExpressRouteConnectionProperties{
			AuthorizationKey: to.Ptr("authorizationKey"),
			ExpressRouteCircuitPeering: &armnetwork.ExpressRouteCircuitPeeringID{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteCircuits/circuitName/peerings/AzurePrivatePeering"),
			},
			RoutingConfiguration: &armnetwork.RoutingConfiguration{
				AssociatedRouteTable: &armnetwork.SubResource{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable1"),
				},
				InboundRouteMap: &armnetwork.SubResource{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap1"),
				},
				OutboundRouteMap: &armnetwork.SubResource{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap2"),
				},
				PropagatedRouteTables: &armnetwork.PropagatedRouteTable{
					IDs: []*armnetwork.SubResource{
						{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable1"),
						},
						{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable2"),
						},
						{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable3"),
						}},
					Labels: []*string{
						to.Ptr("label1"),
						to.Ptr("label2")},
				},
			},
			RoutingWeight: to.Ptr[int32](2),
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
	// res.ExpressRouteConnection = armnetwork.ExpressRouteConnection{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteGateways/gateway-2/expressRouteConnections/connectionName"),
	// 	Name: to.Ptr("connectionName"),
	// 	Properties: &armnetwork.ExpressRouteConnectionProperties{
	// 		AuthorizationKey: to.Ptr("authorizationKey"),
	// 		EnableInternetSecurity: to.Ptr(false),
	// 		EnablePrivateLinkFastPath: to.Ptr(false),
	// 		ExpressRouteCircuitPeering: &armnetwork.ExpressRouteCircuitPeeringID{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteCircuits/circuitName/peerings/AzurePrivatePeering"),
	// 		},
	// 		ExpressRouteGatewayBypass: to.Ptr(false),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		RoutingConfiguration: &armnetwork.RoutingConfiguration{
	// 			AssociatedRouteTable: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable1"),
	// 			},
	// 			InboundRouteMap: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap1"),
	// 			},
	// 			OutboundRouteMap: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap2"),
	// 			},
	// 			PropagatedRouteTables: &armnetwork.PropagatedRouteTable{
	// 				IDs: []*armnetwork.SubResource{
	// 					{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable1"),
	// 					},
	// 					{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable2"),
	// 					},
	// 					{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable3"),
	// 				}},
	// 				Labels: []*string{
	// 					to.Ptr("label1"),
	// 					to.Ptr("label2")},
	// 				},
	// 			},
	// 			RoutingWeight: to.Ptr[int32](2),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/network/resource-manager/Microsoft.Network/stable/2022-09-01/examples/ExpressRouteConnectionGet.json
func ExampleExpressRouteConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewExpressRouteConnectionsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "resourceGroupName", "expressRouteGatewayName", "connectionName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExpressRouteConnection = armnetwork.ExpressRouteConnection{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteGateways/expressRouteGatewayName/expressRouteConnections/connectionName"),
	// 	Name: to.Ptr("connectionName"),
	// 	Properties: &armnetwork.ExpressRouteConnectionProperties{
	// 		AuthorizationKey: to.Ptr("authorizationKey"),
	// 		EnableInternetSecurity: to.Ptr(false),
	// 		EnablePrivateLinkFastPath: to.Ptr(false),
	// 		ExpressRouteCircuitPeering: &armnetwork.ExpressRouteCircuitPeeringID{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteCircuits/circuitName/peerings/AzurePrivatePeering"),
	// 		},
	// 		ExpressRouteGatewayBypass: to.Ptr(false),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		RoutingConfiguration: &armnetwork.RoutingConfiguration{
	// 			AssociatedRouteTable: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable1"),
	// 			},
	// 			InboundRouteMap: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap1"),
	// 			},
	// 			OutboundRouteMap: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap2"),
	// 			},
	// 			PropagatedRouteTables: &armnetwork.PropagatedRouteTable{
	// 				IDs: []*armnetwork.SubResource{
	// 					{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable1"),
	// 					},
	// 					{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable2"),
	// 					},
	// 					{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable3"),
	// 				}},
	// 				Labels: []*string{
	// 					to.Ptr("label1"),
	// 					to.Ptr("label2")},
	// 				},
	// 			},
	// 			RoutingWeight: to.Ptr[int32](1),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/network/resource-manager/Microsoft.Network/stable/2022-09-01/examples/ExpressRouteConnectionDelete.json
func ExampleExpressRouteConnectionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewExpressRouteConnectionsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx, "resourceGroupName", "expressRouteGatewayName", "connectionName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/network/resource-manager/Microsoft.Network/stable/2022-09-01/examples/ExpressRouteConnectionList.json
func ExampleExpressRouteConnectionsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewExpressRouteConnectionsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.List(ctx, "resourceGroupName", "expressRouteGatewayName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExpressRouteConnectionList = armnetwork.ExpressRouteConnectionList{
	// 	Value: []*armnetwork.ExpressRouteConnection{
	// 		{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteGateways/expressRouteGatewayName/expressRouteConnections/connectionName"),
	// 			Name: to.Ptr("connectionName"),
	// 			Properties: &armnetwork.ExpressRouteConnectionProperties{
	// 				AuthorizationKey: to.Ptr("authorizationKey"),
	// 				EnableInternetSecurity: to.Ptr(false),
	// 				EnablePrivateLinkFastPath: to.Ptr(false),
	// 				ExpressRouteCircuitPeering: &armnetwork.ExpressRouteCircuitPeeringID{
	// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteCircuits/circuitName/peerings/AzurePrivatePeering"),
	// 				},
	// 				ExpressRouteGatewayBypass: to.Ptr(false),
	// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 				RoutingConfiguration: &armnetwork.RoutingConfiguration{
	// 					AssociatedRouteTable: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable1"),
	// 					},
	// 					InboundRouteMap: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap1"),
	// 					},
	// 					OutboundRouteMap: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap2"),
	// 					},
	// 					PropagatedRouteTables: &armnetwork.PropagatedRouteTable{
	// 						IDs: []*armnetwork.SubResource{
	// 							{
	// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable1"),
	// 							},
	// 							{
	// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable2"),
	// 							},
	// 							{
	// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable3"),
	// 						}},
	// 						Labels: []*string{
	// 							to.Ptr("label1"),
	// 							to.Ptr("label2")},
	// 						},
	// 					},
	// 					RoutingWeight: to.Ptr[int32](1),
	// 				},
	// 		}},
	// 	}
}
