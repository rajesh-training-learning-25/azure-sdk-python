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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/ExpressRoutePortDelete.json
func ExampleExpressRoutePortsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExpressRoutePortsClient().BeginDelete(ctx, "rg1", "portName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/ExpressRoutePortGet.json
func ExampleExpressRoutePortsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExpressRoutePortsClient().Get(ctx, "rg1", "portName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExpressRoutePort = armnetwork.ExpressRoutePort{
	// 	Name: to.Ptr("portName"),
	// 	Type: to.Ptr("Microsoft.Network/expressRoutePorts"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armnetwork.ExpressRoutePortPropertiesFormat{
	// 		AllocationDate: to.Ptr("Friday, July 1, 2018"),
	// 		BandwidthInGbps: to.Ptr[int32](100),
	// 		BillingType: to.Ptr(armnetwork.ExpressRoutePortsBillingTypeUnlimitedData),
	// 		Circuits: []*armnetwork.SubResource{
	// 		},
	// 		Encapsulation: to.Ptr(armnetwork.ExpressRoutePortsEncapsulationQinQ),
	// 		EtherType: to.Ptr("0x8100"),
	// 		Links: []*armnetwork.ExpressRouteLink{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName/links/link1"),
	// 				Name: to.Ptr("link1"),
	// 				Properties: &armnetwork.ExpressRouteLinkPropertiesFormat{
	// 					AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateDisabled),
	// 					ColoLocation: to.Ptr("coloLocation1"),
	// 					ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
	// 					InterfaceName: to.Ptr("Ethernet 0/0"),
	// 					PatchPanelID: to.Ptr("patchPanelId1"),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					RackID: to.Ptr("rackId1"),
	// 					RouterName: to.Ptr("router1"),
	// 				},
	// 			},
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName/links/link2"),
	// 				Name: to.Ptr("link2"),
	// 				Properties: &armnetwork.ExpressRouteLinkPropertiesFormat{
	// 					AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateDisabled),
	// 					ColoLocation: to.Ptr("coloLocation2"),
	// 					ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
	// 					InterfaceName: to.Ptr("Ethernet 0/0"),
	// 					PatchPanelID: to.Ptr("patchPanelId2"),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					RackID: to.Ptr("rackId2"),
	// 					RouterName: to.Ptr("router2"),
	// 				},
	// 		}},
	// 		Mtu: to.Ptr("1500"),
	// 		PeeringLocation: to.Ptr("peeringLocationName"),
	// 		ProvisionedBandwidthInGbps: to.Ptr[float32](0),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/ExpressRoutePortCreate.json
func ExampleExpressRoutePortsClient_BeginCreateOrUpdate_expressRoutePortCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExpressRoutePortsClient().BeginCreateOrUpdate(ctx, "rg1", "portName", armnetwork.ExpressRoutePort{
		Location: to.Ptr("westus"),
		Properties: &armnetwork.ExpressRoutePortPropertiesFormat{
			BandwidthInGbps: to.Ptr[int32](100),
			BillingType:     to.Ptr(armnetwork.ExpressRoutePortsBillingTypeUnlimitedData),
			Encapsulation:   to.Ptr(armnetwork.ExpressRoutePortsEncapsulationQinQ),
			PeeringLocation: to.Ptr("peeringLocationName"),
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
	// res.ExpressRoutePort = armnetwork.ExpressRoutePort{
	// 	Name: to.Ptr("portName"),
	// 	Type: to.Ptr("Microsoft.Network/expressRoutePorts"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armnetwork.ExpressRoutePortPropertiesFormat{
	// 		AllocationDate: to.Ptr("Friday, July 1, 2018"),
	// 		BandwidthInGbps: to.Ptr[int32](100),
	// 		BillingType: to.Ptr(armnetwork.ExpressRoutePortsBillingTypeUnlimitedData),
	// 		Circuits: []*armnetwork.SubResource{
	// 		},
	// 		Encapsulation: to.Ptr(armnetwork.ExpressRoutePortsEncapsulationQinQ),
	// 		EtherType: to.Ptr("0x8100"),
	// 		Links: []*armnetwork.ExpressRouteLink{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName/links/link1"),
	// 				Name: to.Ptr("link1"),
	// 				Properties: &armnetwork.ExpressRouteLinkPropertiesFormat{
	// 					AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateDisabled),
	// 					ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
	// 					InterfaceName: to.Ptr("Ethernet 0/0"),
	// 					PatchPanelID: to.Ptr("patchPanelId1"),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					RackID: to.Ptr("rackId1"),
	// 					RouterName: to.Ptr("router1"),
	// 				},
	// 			},
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName/links/link2"),
	// 				Name: to.Ptr("link2"),
	// 				Properties: &armnetwork.ExpressRouteLinkPropertiesFormat{
	// 					AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateDisabled),
	// 					ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
	// 					InterfaceName: to.Ptr("Ethernet 0/0"),
	// 					PatchPanelID: to.Ptr("patchPanelId2"),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					RackID: to.Ptr("rackId2"),
	// 					RouterName: to.Ptr("router2"),
	// 				},
	// 		}},
	// 		Mtu: to.Ptr("1500"),
	// 		PeeringLocation: to.Ptr("peeringLocationName"),
	// 		ProvisionedBandwidthInGbps: to.Ptr[float32](0),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/ExpressRoutePortUpdateLink.json
func ExampleExpressRoutePortsClient_BeginCreateOrUpdate_expressRoutePortUpdateLink() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExpressRoutePortsClient().BeginCreateOrUpdate(ctx, "rg1", "portName", armnetwork.ExpressRoutePort{
		Location: to.Ptr("westus"),
		Properties: &armnetwork.ExpressRoutePortPropertiesFormat{
			BandwidthInGbps: to.Ptr[int32](100),
			BillingType:     to.Ptr(armnetwork.ExpressRoutePortsBillingTypeUnlimitedData),
			Encapsulation:   to.Ptr(armnetwork.ExpressRoutePortsEncapsulationQinQ),
			Links: []*armnetwork.ExpressRouteLink{
				{
					Name: to.Ptr("link1"),
					Properties: &armnetwork.ExpressRouteLinkPropertiesFormat{
						AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateEnabled),
					},
				}},
			PeeringLocation: to.Ptr("peeringLocationName"),
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
	// res.ExpressRoutePort = armnetwork.ExpressRoutePort{
	// 	Name: to.Ptr("portName"),
	// 	Type: to.Ptr("Microsoft.Network/expressRoutePorts"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armnetwork.ExpressRoutePortPropertiesFormat{
	// 		AllocationDate: to.Ptr("Friday, July 1, 2018"),
	// 		BandwidthInGbps: to.Ptr[int32](100),
	// 		BillingType: to.Ptr(armnetwork.ExpressRoutePortsBillingTypeUnlimitedData),
	// 		Circuits: []*armnetwork.SubResource{
	// 		},
	// 		Encapsulation: to.Ptr(armnetwork.ExpressRoutePortsEncapsulationQinQ),
	// 		EtherType: to.Ptr("0x8100"),
	// 		Links: []*armnetwork.ExpressRouteLink{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName/links/link1"),
	// 				Name: to.Ptr("link1"),
	// 				Properties: &armnetwork.ExpressRouteLinkPropertiesFormat{
	// 					AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateEnabled),
	// 					ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
	// 					InterfaceName: to.Ptr("Ethernet 0/0"),
	// 					PatchPanelID: to.Ptr("patchPanelId1"),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					RackID: to.Ptr("rackId1"),
	// 					RouterName: to.Ptr("router1"),
	// 				},
	// 			},
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName/links/link2"),
	// 				Name: to.Ptr("link2"),
	// 				Properties: &armnetwork.ExpressRouteLinkPropertiesFormat{
	// 					AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateDisabled),
	// 					ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
	// 					InterfaceName: to.Ptr("Ethernet 0/0"),
	// 					PatchPanelID: to.Ptr("patchPanelId2"),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					RackID: to.Ptr("rackId2"),
	// 					RouterName: to.Ptr("router2"),
	// 				},
	// 		}},
	// 		Mtu: to.Ptr("1500"),
	// 		PeeringLocation: to.Ptr("peeringLocationName"),
	// 		ProvisionedBandwidthInGbps: to.Ptr[float32](0),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/ExpressRoutePortUpdateTags.json
func ExampleExpressRoutePortsClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExpressRoutePortsClient().UpdateTags(ctx, "rg1", "portName", armnetwork.TagsObject{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExpressRoutePort = armnetwork.ExpressRoutePort{
	// 	Name: to.Ptr("portName"),
	// 	Type: to.Ptr("Microsoft.Network/expressRoutePorts"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armnetwork.ExpressRoutePortPropertiesFormat{
	// 		AllocationDate: to.Ptr("Friday, July 1, 2018"),
	// 		BandwidthInGbps: to.Ptr[int32](100),
	// 		BillingType: to.Ptr(armnetwork.ExpressRoutePortsBillingTypeUnlimitedData),
	// 		Circuits: []*armnetwork.SubResource{
	// 		},
	// 		Encapsulation: to.Ptr(armnetwork.ExpressRoutePortsEncapsulationQinQ),
	// 		EtherType: to.Ptr("0x8100"),
	// 		Links: []*armnetwork.ExpressRouteLink{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName/links/link1"),
	// 				Name: to.Ptr("link1"),
	// 				Properties: &armnetwork.ExpressRouteLinkPropertiesFormat{
	// 					AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateDisabled),
	// 					ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
	// 					InterfaceName: to.Ptr("Ethernet 0/0"),
	// 					PatchPanelID: to.Ptr("patchPanelId1"),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					RackID: to.Ptr("rackId1"),
	// 					RouterName: to.Ptr("router1"),
	// 				},
	// 			},
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName/links/link2"),
	// 				Name: to.Ptr("link2"),
	// 				Properties: &armnetwork.ExpressRouteLinkPropertiesFormat{
	// 					AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateDisabled),
	// 					ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
	// 					InterfaceName: to.Ptr("Ethernet 0/0"),
	// 					PatchPanelID: to.Ptr("patchPanelId2"),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					RackID: to.Ptr("rackId2"),
	// 					RouterName: to.Ptr("router2"),
	// 				},
	// 		}},
	// 		Mtu: to.Ptr("1500"),
	// 		PeeringLocation: to.Ptr("peeringLocationName"),
	// 		ProvisionedBandwidthInGbps: to.Ptr[float32](0),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/ExpressRoutePortListByResourceGroup.json
func ExampleExpressRoutePortsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExpressRoutePortsClient().NewListByResourceGroupPager("rg1", nil)
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
		// page.ExpressRoutePortListResult = armnetwork.ExpressRoutePortListResult{
		// 	Value: []*armnetwork.ExpressRoutePort{
		// 		{
		// 			Name: to.Ptr("portName"),
		// 			Type: to.Ptr("Microsoft.Network/expressRoutePorts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armnetwork.ExpressRoutePortPropertiesFormat{
		// 				AllocationDate: to.Ptr("Friday, July 1, 2018"),
		// 				BandwidthInGbps: to.Ptr[int32](100),
		// 				BillingType: to.Ptr(armnetwork.ExpressRoutePortsBillingTypeUnlimitedData),
		// 				Circuits: []*armnetwork.SubResource{
		// 				},
		// 				Encapsulation: to.Ptr(armnetwork.ExpressRoutePortsEncapsulationQinQ),
		// 				EtherType: to.Ptr("0x8100"),
		// 				Links: []*armnetwork.ExpressRouteLink{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName/links/link1"),
		// 						Name: to.Ptr("link1"),
		// 						Properties: &armnetwork.ExpressRouteLinkPropertiesFormat{
		// 							AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateDisabled),
		// 							ColoLocation: to.Ptr("coloLocation1"),
		// 							ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
		// 							InterfaceName: to.Ptr("Ethernet 0/0"),
		// 							PatchPanelID: to.Ptr("patchPanelId1"),
		// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							RackID: to.Ptr("rackId1"),
		// 							RouterName: to.Ptr("router1"),
		// 						},
		// 					},
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName/links/link2"),
		// 						Name: to.Ptr("link2"),
		// 						Properties: &armnetwork.ExpressRouteLinkPropertiesFormat{
		// 							AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateDisabled),
		// 							ColoLocation: to.Ptr("coloLocation2"),
		// 							ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
		// 							InterfaceName: to.Ptr("Ethernet 0/0"),
		// 							PatchPanelID: to.Ptr("patchPanelId2"),
		// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							RackID: to.Ptr("rackId2"),
		// 							RouterName: to.Ptr("router2"),
		// 						},
		// 				}},
		// 				Mtu: to.Ptr("1500"),
		// 				PeeringLocation: to.Ptr("peeringLocationName"),
		// 				ProvisionedBandwidthInGbps: to.Ptr[float32](0),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/ExpressRoutePortList.json
func ExampleExpressRoutePortsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExpressRoutePortsClient().NewListPager(nil)
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
		// page.ExpressRoutePortListResult = armnetwork.ExpressRoutePortListResult{
		// 	Value: []*armnetwork.ExpressRoutePort{
		// 		{
		// 			Name: to.Ptr("portName"),
		// 			Type: to.Ptr("Microsoft.Network/expressRoutePorts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armnetwork.ExpressRoutePortPropertiesFormat{
		// 				AllocationDate: to.Ptr("Friday, July 1, 2018"),
		// 				BandwidthInGbps: to.Ptr[int32](100),
		// 				BillingType: to.Ptr(armnetwork.ExpressRoutePortsBillingTypeUnlimitedData),
		// 				Circuits: []*armnetwork.SubResource{
		// 				},
		// 				Encapsulation: to.Ptr(armnetwork.ExpressRoutePortsEncapsulationQinQ),
		// 				EtherType: to.Ptr("0x8100"),
		// 				Links: []*armnetwork.ExpressRouteLink{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName/links/link1"),
		// 						Name: to.Ptr("link1"),
		// 						Properties: &armnetwork.ExpressRouteLinkPropertiesFormat{
		// 							AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateDisabled),
		// 							ColoLocation: to.Ptr("coloLocation1"),
		// 							ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
		// 							InterfaceName: to.Ptr("Ethernet 0/0"),
		// 							PatchPanelID: to.Ptr("patchPanelId1"),
		// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							RackID: to.Ptr("rackId1"),
		// 							RouterName: to.Ptr("router1"),
		// 						},
		// 					},
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName/links/link2"),
		// 						Name: to.Ptr("link2"),
		// 						Properties: &armnetwork.ExpressRouteLinkPropertiesFormat{
		// 							AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateDisabled),
		// 							ColoLocation: to.Ptr("coloLocation2"),
		// 							ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
		// 							InterfaceName: to.Ptr("Ethernet 0/0"),
		// 							PatchPanelID: to.Ptr("patchPanelId2"),
		// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							RackID: to.Ptr("rackId2"),
		// 							RouterName: to.Ptr("router2"),
		// 						},
		// 				}},
		// 				Mtu: to.Ptr("1500"),
		// 				PeeringLocation: to.Ptr("peeringLocationName"),
		// 				ProvisionedBandwidthInGbps: to.Ptr[float32](0),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/GenerateExpressRoutePortsLOA.json
func ExampleExpressRoutePortsClient_GenerateLOA() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExpressRoutePortsClient().GenerateLOA(ctx, "rg1", "portName", armnetwork.GenerateExpressRoutePortsLOARequest{
		CustomerName: to.Ptr("customerName"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GenerateExpressRoutePortsLOAResult = armnetwork.GenerateExpressRoutePortsLOAResult{
	// 	EncodedContent: to.Ptr("TWFuIGlzIGRpc3"),
	// }
}
