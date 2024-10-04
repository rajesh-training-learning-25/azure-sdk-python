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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4883fa5dbf6f2c9093fac8ce334547e9dfac68fa/specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/NetworkManagerSecurityUserRuleList.json
func ExampleSecurityUserRulesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSecurityUserRulesClient().NewListPager("rg1", "testNetworkManager", "myTestSecurityConfig", "testRuleCollection", &armnetwork.SecurityUserRulesClientListOptions{Top: nil,
		SkipToken: nil,
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
		// page.SecurityUserRuleListResult = armnetwork.SecurityUserRuleListResult{
		// 	Value: []*armnetwork.SecurityUserRule{
		// 		{
		// 			Name: to.Ptr("SampleUserRule"),
		// 			Type: to.Ptr("Microsoft.Network/networkManagers/securityConfigurations/ruleCollections/rules"),
		// 			ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/networkmanagers/testNetworkManager/securityUserConfigurations/myTestSecurityConfig/ruleCollections/testRuleCollection/rules/SampleUserRule"),
		// 			Properties: &armnetwork.SecurityUserRulePropertiesFormat{
		// 				Description: to.Ptr("Sample User Rule"),
		// 				DestinationPortRanges: []*string{
		// 					to.Ptr("22")},
		// 					Destinations: []*armnetwork.AddressPrefixItem{
		// 						{
		// 							AddressPrefix: to.Ptr("*"),
		// 							AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeIPPrefix),
		// 					}},
		// 					Direction: to.Ptr(armnetwork.SecurityConfigurationRuleDirectionInbound),
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					ResourceGUID: to.Ptr("f930553b-f78d-48c5-9445-6cf86b85e615"),
		// 					SourcePortRanges: []*string{
		// 						to.Ptr("0-65535")},
		// 						Sources: []*armnetwork.AddressPrefixItem{
		// 							{
		// 								AddressPrefix: to.Ptr("*"),
		// 								AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeIPPrefix),
		// 						}},
		// 						Protocol: to.Ptr(armnetwork.SecurityConfigurationRuleProtocolTCP),
		// 					},
		// 					SystemData: &armnetwork.SystemData{
		// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27.000Z"); return t}()),
		// 						CreatedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
		// 						CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 						LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27.000Z"); return t}()),
		// 						LastModifiedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
		// 						LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 					},
		// 			}},
		// 		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4883fa5dbf6f2c9093fac8ce334547e9dfac68fa/specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/NetworkManagerSecurityUserRuleGet.json
func ExampleSecurityUserRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSecurityUserRulesClient().Get(ctx, "rg1", "testNetworkManager", "myTestSecurityConfig", "testRuleCollection", "SampleUserRule", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SecurityUserRule = armnetwork.SecurityUserRule{
	// 	Name: to.Ptr("SampleUserRule"),
	// 	Type: to.Ptr("Microsoft.Network/networkManagers/SecurityUserConfigurations/ruleCollections/rules"),
	// 	ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager/SecurityUserConfigurations/myTestSecurityConfig/ruleCollections/testRuleCollection/rules/SampleUserRule"),
	// 	Properties: &armnetwork.SecurityUserRulePropertiesFormat{
	// 		Description: to.Ptr("Sample User Rule"),
	// 		DestinationPortRanges: []*string{
	// 			to.Ptr("22")},
	// 			Destinations: []*armnetwork.AddressPrefixItem{
	// 				{
	// 					AddressPrefix: to.Ptr("*"),
	// 					AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeIPPrefix),
	// 			}},
	// 			Direction: to.Ptr(armnetwork.SecurityConfigurationRuleDirectionInbound),
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			ResourceGUID: to.Ptr("f930553b-f78d-48c5-9445-6cf86b85e615"),
	// 			SourcePortRanges: []*string{
	// 				to.Ptr("0-65535")},
	// 				Sources: []*armnetwork.AddressPrefixItem{
	// 					{
	// 						AddressPrefix: to.Ptr("*"),
	// 						AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeIPPrefix),
	// 				}},
	// 				Protocol: to.Ptr(armnetwork.SecurityConfigurationRuleProtocolTCP),
	// 			},
	// 			SystemData: &armnetwork.SystemData{
	// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27.000Z"); return t}()),
	// 				CreatedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
	// 				CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27.000Z"); return t}()),
	// 				LastModifiedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
	// 				LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4883fa5dbf6f2c9093fac8ce334547e9dfac68fa/specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/NetworkManagerSecurityUserRulePut.json
func ExampleSecurityUserRulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSecurityUserRulesClient().CreateOrUpdate(ctx, "rg1", "testNetworkManager", "myTestSecurityConfig", "testRuleCollection", "SampleUserRule", armnetwork.SecurityUserRule{
		Properties: &armnetwork.SecurityUserRulePropertiesFormat{
			Description: to.Ptr("Sample User Rule"),
			DestinationPortRanges: []*string{
				to.Ptr("22")},
			Destinations: []*armnetwork.AddressPrefixItem{
				{
					AddressPrefix:     to.Ptr("*"),
					AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeIPPrefix),
				}},
			Direction: to.Ptr(armnetwork.SecurityConfigurationRuleDirectionInbound),
			SourcePortRanges: []*string{
				to.Ptr("0-65535")},
			Sources: []*armnetwork.AddressPrefixItem{
				{
					AddressPrefix:     to.Ptr("*"),
					AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeIPPrefix),
				}},
			Protocol: to.Ptr(armnetwork.SecurityConfigurationRuleProtocolTCP),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SecurityUserRule = armnetwork.SecurityUserRule{
	// 	Name: to.Ptr("SampleUserRule"),
	// 	Type: to.Ptr("Microsoft.Network/networkManagers/SecurityUserConfigurations/ruleCollections/rules"),
	// 	ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager/SecurityUserConfigurations/Policy1/ruleCollections/testRuleCollection/rules/SampleUserRule"),
	// 	Properties: &armnetwork.SecurityUserRulePropertiesFormat{
	// 		Description: to.Ptr("Sample User Rule"),
	// 		DestinationPortRanges: []*string{
	// 			to.Ptr("22")},
	// 			Destinations: []*armnetwork.AddressPrefixItem{
	// 				{
	// 					AddressPrefix: to.Ptr("*"),
	// 					AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeIPPrefix),
	// 			}},
	// 			Direction: to.Ptr(armnetwork.SecurityConfigurationRuleDirectionInbound),
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			ResourceGUID: to.Ptr("f930553b-f78d-48c5-9445-6cf86b85e615"),
	// 			SourcePortRanges: []*string{
	// 				to.Ptr("0-65535")},
	// 				Sources: []*armnetwork.AddressPrefixItem{
	// 					{
	// 						AddressPrefix: to.Ptr("*"),
	// 						AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeIPPrefix),
	// 				}},
	// 				Protocol: to.Ptr(armnetwork.SecurityConfigurationRuleProtocolTCP),
	// 			},
	// 			SystemData: &armnetwork.SystemData{
	// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27.000Z"); return t}()),
	// 				CreatedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
	// 				CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27.000Z"); return t}()),
	// 				LastModifiedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
	// 				LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4883fa5dbf6f2c9093fac8ce334547e9dfac68fa/specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/NetworkManagerSecurityUserRuleDelete.json
func ExampleSecurityUserRulesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSecurityUserRulesClient().BeginDelete(ctx, "rg1", "testNetworkManager", "myTestSecurityConfig", "testRuleCollection", "SampleUserRule", &armnetwork.SecurityUserRulesClientBeginDeleteOptions{Force: to.Ptr(false)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
