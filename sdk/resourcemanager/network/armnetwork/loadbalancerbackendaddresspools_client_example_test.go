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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/LBBackendAddressPoolListWithBackendAddressesPoolType.json
func ExampleLoadBalancerBackendAddressPoolsClient_NewListPager_loadBalancerWithBackendAddressPoolContainingBackendAddresses() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLoadBalancerBackendAddressPoolsClient().NewListPager("testrg", "lb", nil)
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
		// page.LoadBalancerBackendAddressPoolListResult = armnetwork.LoadBalancerBackendAddressPoolListResult{
		// 	Value: []*armnetwork.BackendAddressPool{
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/backend"),
		// 			Name: to.Ptr("backend"),
		// 			Type: to.Ptr("Microsoft.Network/loadBalancers/backendAddressPools"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
		// 			Properties: &armnetwork.BackendAddressPoolPropertiesFormat{
		// 				LoadBalancerBackendAddresses: []*armnetwork.LoadBalancerBackendAddress{
		// 					{
		// 						Name: to.Ptr("address1"),
		// 						Properties: &armnetwork.LoadBalancerBackendAddressPropertiesFormat{
		// 							IPAddress: to.Ptr("10.0.0.4"),
		// 							VirtualNetwork: &armnetwork.SubResource{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb"),
		// 							},
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("address2"),
		// 						Properties: &armnetwork.LoadBalancerBackendAddressPropertiesFormat{
		// 							IPAddress: to.Ptr("10.0.0.5"),
		// 							VirtualNetwork: &armnetwork.SubResource{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb"),
		// 							},
		// 						},
		// 				}},
		// 				LoadBalancingRules: []*armnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb/loadBalancingRules/rulelb"),
		// 				}},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/LoadBalancerBackendAddressPoolList.json
func ExampleLoadBalancerBackendAddressPoolsClient_NewListPager_loadBalancerBackendAddressPoolList() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLoadBalancerBackendAddressPoolsClient().NewListPager("testrg", "lb", nil)
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
		// page.LoadBalancerBackendAddressPoolListResult = armnetwork.LoadBalancerBackendAddressPoolListResult{
		// 	Value: []*armnetwork.BackendAddressPool{
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/backend"),
		// 			Name: to.Ptr("backend"),
		// 			Type: to.Ptr("Microsoft.Network/loadBalancers/backendAddressPools"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
		// 			Properties: &armnetwork.BackendAddressPoolPropertiesFormat{
		// 				BackendIPConfigurations: []*armnetwork.InterfaceIPConfiguration{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/networkInterfaces/nic/ipConfigurations/default-ip-config"),
		// 				}},
		// 				LoadBalancingRules: []*armnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb/loadBalancingRules/rulelb"),
		// 				}},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/LBBackendAddressPoolWithBackendAddressesGet.json
func ExampleLoadBalancerBackendAddressPoolsClient_Get_loadBalancerWithBackendAddressPoolWithBackendAddresses() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLoadBalancerBackendAddressPoolsClient().Get(ctx, "testrg", "lb", "backend", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BackendAddressPool = armnetwork.BackendAddressPool{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/backend"),
	// 	Name: to.Ptr("backend"),
	// 	Type: to.Ptr("Microsoft.Network/loadBalancers/backendAddressPools"),
	// 	Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
	// 	Properties: &armnetwork.BackendAddressPoolPropertiesFormat{
	// 		LoadBalancerBackendAddresses: []*armnetwork.LoadBalancerBackendAddress{
	// 			{
	// 				Name: to.Ptr("address1"),
	// 				Properties: &armnetwork.LoadBalancerBackendAddressPropertiesFormat{
	// 					IPAddress: to.Ptr("10.0.0.4"),
	// 					VirtualNetwork: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb"),
	// 					},
	// 				},
	// 			},
	// 			{
	// 				Name: to.Ptr("address2"),
	// 				Properties: &armnetwork.LoadBalancerBackendAddressPropertiesFormat{
	// 					IPAddress: to.Ptr("10.0.0.5"),
	// 					VirtualNetwork: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb"),
	// 					},
	// 				},
	// 		}},
	// 		LoadBalancingRules: []*armnetwork.SubResource{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb/loadBalancingRules/rulelb"),
	// 		}},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/LoadBalancerBackendAddressPoolGet.json
func ExampleLoadBalancerBackendAddressPoolsClient_Get_loadBalancerBackendAddressPoolGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLoadBalancerBackendAddressPoolsClient().Get(ctx, "testrg", "lb", "backend", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BackendAddressPool = armnetwork.BackendAddressPool{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/backend"),
	// 	Name: to.Ptr("backend"),
	// 	Type: to.Ptr("Microsoft.Network/loadBalancers/backendAddressPools"),
	// 	Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
	// 	Properties: &armnetwork.BackendAddressPoolPropertiesFormat{
	// 		BackendIPConfigurations: []*armnetwork.InterfaceIPConfiguration{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/networkInterfaces/nic/ipConfigurations/default-ip-config"),
	// 		}},
	// 		LoadBalancingRules: []*armnetwork.SubResource{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb/loadBalancingRules/rulelb"),
	// 		}},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/LBBackendAddressPoolWithBackendAddressesPut.json
func ExampleLoadBalancerBackendAddressPoolsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLoadBalancerBackendAddressPoolsClient().BeginCreateOrUpdate(ctx, "testrg", "lb", "backend", armnetwork.BackendAddressPool{
		Properties: &armnetwork.BackendAddressPoolPropertiesFormat{
			LoadBalancerBackendAddresses: []*armnetwork.LoadBalancerBackendAddress{
				{
					Name: to.Ptr("address1"),
					Properties: &armnetwork.LoadBalancerBackendAddressPropertiesFormat{
						IPAddress: to.Ptr("10.0.0.4"),
						VirtualNetwork: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb"),
						},
					},
				},
				{
					Name: to.Ptr("address2"),
					Properties: &armnetwork.LoadBalancerBackendAddressPropertiesFormat{
						IPAddress: to.Ptr("10.0.0.5"),
						VirtualNetwork: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb"),
						},
					},
				}},
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
	// res.BackendAddressPool = armnetwork.BackendAddressPool{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/backend"),
	// 	Name: to.Ptr("backend"),
	// 	Type: to.Ptr("Microsoft.Network/loadBalancers/backendAddressPools"),
	// 	Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
	// 	Properties: &armnetwork.BackendAddressPoolPropertiesFormat{
	// 		LoadBalancerBackendAddresses: []*armnetwork.LoadBalancerBackendAddress{
	// 			{
	// 				Name: to.Ptr("address1"),
	// 				Properties: &armnetwork.LoadBalancerBackendAddressPropertiesFormat{
	// 					IPAddress: to.Ptr("10.0.0.4"),
	// 					VirtualNetwork: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb"),
	// 					},
	// 				},
	// 			},
	// 			{
	// 				Name: to.Ptr("address2"),
	// 				Properties: &armnetwork.LoadBalancerBackendAddressPropertiesFormat{
	// 					IPAddress: to.Ptr("10.0.0.5"),
	// 					VirtualNetwork: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb"),
	// 					},
	// 				},
	// 		}},
	// 		LoadBalancingRules: []*armnetwork.SubResource{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb/loadBalancingRules/rulelb"),
	// 		}},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/LoadBalancerBackendAddressPoolDelete.json
func ExampleLoadBalancerBackendAddressPoolsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLoadBalancerBackendAddressPoolsClient().BeginDelete(ctx, "testrg", "lb", "backend", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
