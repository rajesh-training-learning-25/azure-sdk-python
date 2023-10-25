//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7b551033155a63739b6d28f79b9c07569f6179b8/specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/Routes_ListByEndpoint.json
func ExampleRoutesClient_NewListByEndpointPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRoutesClient().NewListByEndpointPager("RG", "profile1", "endpoint1", nil)
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
		// page.RouteListResult = armcdn.RouteListResult{
		// 	Value: []*armcdn.Route{
		// 		{
		// 			Name: to.Ptr("route1"),
		// 			Type: to.Ptr("Microsoft.Cdn/profiles/afdendpoints/routes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/afdendpoints/endpoint1/routes/route1"),
		// 			Properties: &armcdn.RouteProperties{
		// 				DeploymentStatus: to.Ptr(armcdn.DeploymentStatusNotStarted),
		// 				ProvisioningState: to.Ptr(armcdn.AfdProvisioningStateSucceeded),
		// 				CacheConfiguration: &armcdn.AfdRouteCacheConfiguration{
		// 					CompressionSettings: &armcdn.CompressionSettings{
		// 						ContentTypesToCompress: []*string{
		// 							to.Ptr("text/html"),
		// 							to.Ptr("application/octet-stream")},
		// 							IsCompressionEnabled: to.Ptr(true),
		// 						},
		// 						QueryStringCachingBehavior: to.Ptr(armcdn.AfdQueryStringCachingBehaviorIgnoreQueryString),
		// 					},
		// 					CustomDomains: []*armcdn.ActivatedResourceReference{
		// 						{
		// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/customDomains/domain1"),
		// 					}},
		// 					EnabledState: to.Ptr(armcdn.EnabledStateEnabled),
		// 					ForwardingProtocol: to.Ptr(armcdn.ForwardingProtocolMatchRequest),
		// 					HTTPSRedirect: to.Ptr(armcdn.HTTPSRedirectEnabled),
		// 					LinkToDefaultDomain: to.Ptr(armcdn.LinkToDefaultDomainEnabled),
		// 					OriginGroup: &armcdn.ResourceReference{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/originGroups/originGroup1"),
		// 					},
		// 					PatternsToMatch: []*string{
		// 						to.Ptr("/*")},
		// 						RuleSets: []*armcdn.ResourceReference{
		// 							{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/ruleSets/ruleSet1"),
		// 						}},
		// 						SupportedProtocols: []*armcdn.AFDEndpointProtocols{
		// 							to.Ptr(armcdn.AFDEndpointProtocolsHTTPS),
		// 							to.Ptr(armcdn.AFDEndpointProtocolsHTTP)},
		// 						},
		// 				}},
		// 			}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7b551033155a63739b6d28f79b9c07569f6179b8/specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/Routes_Get.json
func ExampleRoutesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoutesClient().Get(ctx, "RG", "profile1", "endpoint1", "route1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Route = armcdn.Route{
	// 	Name: to.Ptr("route1"),
	// 	Type: to.Ptr("Microsoft.Cdn/profiles/afdendpoints/routes"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/afdendpoints/endpoint1/routes/route1"),
	// 	Properties: &armcdn.RouteProperties{
	// 		DeploymentStatus: to.Ptr(armcdn.DeploymentStatusNotStarted),
	// 		ProvisioningState: to.Ptr(armcdn.AfdProvisioningStateSucceeded),
	// 		CacheConfiguration: &armcdn.AfdRouteCacheConfiguration{
	// 			CompressionSettings: &armcdn.CompressionSettings{
	// 				ContentTypesToCompress: []*string{
	// 					to.Ptr("text/html"),
	// 					to.Ptr("application/octet-stream")},
	// 					IsCompressionEnabled: to.Ptr(true),
	// 				},
	// 				QueryStringCachingBehavior: to.Ptr(armcdn.AfdQueryStringCachingBehaviorIgnoreQueryString),
	// 			},
	// 			CustomDomains: []*armcdn.ActivatedResourceReference{
	// 				{
	// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/customDomains/domain1"),
	// 			}},
	// 			EnabledState: to.Ptr(armcdn.EnabledStateEnabled),
	// 			ForwardingProtocol: to.Ptr(armcdn.ForwardingProtocolMatchRequest),
	// 			HTTPSRedirect: to.Ptr(armcdn.HTTPSRedirectEnabled),
	// 			LinkToDefaultDomain: to.Ptr(armcdn.LinkToDefaultDomainEnabled),
	// 			OriginGroup: &armcdn.ResourceReference{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/originGroups/originGroup1"),
	// 			},
	// 			PatternsToMatch: []*string{
	// 				to.Ptr("/*")},
	// 				RuleSets: []*armcdn.ResourceReference{
	// 					{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/ruleSets/ruleSet1"),
	// 				}},
	// 				SupportedProtocols: []*armcdn.AFDEndpointProtocols{
	// 					to.Ptr(armcdn.AFDEndpointProtocolsHTTPS),
	// 					to.Ptr(armcdn.AFDEndpointProtocolsHTTP)},
	// 				},
	// 			}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7b551033155a63739b6d28f79b9c07569f6179b8/specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/Routes_Create.json
func ExampleRoutesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRoutesClient().BeginCreate(ctx, "RG", "profile1", "endpoint1", "route1", armcdn.Route{
		Properties: &armcdn.RouteProperties{
			CacheConfiguration: &armcdn.AfdRouteCacheConfiguration{
				CompressionSettings: &armcdn.CompressionSettings{
					ContentTypesToCompress: []*string{
						to.Ptr("text/html"),
						to.Ptr("application/octet-stream")},
					IsCompressionEnabled: to.Ptr(true),
				},
				QueryParameters:            to.Ptr("querystring=test"),
				QueryStringCachingBehavior: to.Ptr(armcdn.AfdQueryStringCachingBehaviorIgnoreSpecifiedQueryStrings),
			},
			CustomDomains: []*armcdn.ActivatedResourceReference{
				{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/customDomains/domain1"),
				}},
			EnabledState:        to.Ptr(armcdn.EnabledStateEnabled),
			ForwardingProtocol:  to.Ptr(armcdn.ForwardingProtocolMatchRequest),
			HTTPSRedirect:       to.Ptr(armcdn.HTTPSRedirectEnabled),
			LinkToDefaultDomain: to.Ptr(armcdn.LinkToDefaultDomainEnabled),
			OriginGroup: &armcdn.ResourceReference{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/originGroups/originGroup1"),
			},
			PatternsToMatch: []*string{
				to.Ptr("/*")},
			RuleSets: []*armcdn.ResourceReference{
				{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/ruleSets/ruleSet1"),
				}},
			SupportedProtocols: []*armcdn.AFDEndpointProtocols{
				to.Ptr(armcdn.AFDEndpointProtocolsHTTPS),
				to.Ptr(armcdn.AFDEndpointProtocolsHTTP)},
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
	// res.Route = armcdn.Route{
	// 	Name: to.Ptr("route1"),
	// 	Type: to.Ptr("Microsoft.Cdn/profiles/afdendpoints/routes"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/afdendpoints/endpoint1/routes/route1"),
	// 	Properties: &armcdn.RouteProperties{
	// 		DeploymentStatus: to.Ptr(armcdn.DeploymentStatusNotStarted),
	// 		ProvisioningState: to.Ptr(armcdn.AfdProvisioningStateSucceeded),
	// 		CacheConfiguration: &armcdn.AfdRouteCacheConfiguration{
	// 			CompressionSettings: &armcdn.CompressionSettings{
	// 				ContentTypesToCompress: []*string{
	// 					to.Ptr("text/html"),
	// 					to.Ptr("application/octet-stream")},
	// 					IsCompressionEnabled: to.Ptr(true),
	// 				},
	// 				QueryParameters: to.Ptr("querystring=test"),
	// 				QueryStringCachingBehavior: to.Ptr(armcdn.AfdQueryStringCachingBehaviorIgnoreSpecifiedQueryStrings),
	// 			},
	// 			CustomDomains: []*armcdn.ActivatedResourceReference{
	// 				{
	// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/customDomains/domain1"),
	// 			}},
	// 			EnabledState: to.Ptr(armcdn.EnabledStateEnabled),
	// 			ForwardingProtocol: to.Ptr(armcdn.ForwardingProtocolMatchRequest),
	// 			HTTPSRedirect: to.Ptr(armcdn.HTTPSRedirectEnabled),
	// 			LinkToDefaultDomain: to.Ptr(armcdn.LinkToDefaultDomainEnabled),
	// 			OriginGroup: &armcdn.ResourceReference{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/originGroups/originGroup1"),
	// 			},
	// 			PatternsToMatch: []*string{
	// 				to.Ptr("/*")},
	// 				RuleSets: []*armcdn.ResourceReference{
	// 					{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/ruleSets/ruleSet1"),
	// 				}},
	// 				SupportedProtocols: []*armcdn.AFDEndpointProtocols{
	// 					to.Ptr(armcdn.AFDEndpointProtocolsHTTPS),
	// 					to.Ptr(armcdn.AFDEndpointProtocolsHTTP)},
	// 				},
	// 			}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7b551033155a63739b6d28f79b9c07569f6179b8/specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/Routes_Update.json
func ExampleRoutesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRoutesClient().BeginUpdate(ctx, "RG", "profile1", "endpoint1", "route1", armcdn.RouteUpdateParameters{
		Properties: &armcdn.RouteUpdatePropertiesParameters{
			CacheConfiguration: &armcdn.AfdRouteCacheConfiguration{
				CompressionSettings: &armcdn.CompressionSettings{
					ContentTypesToCompress: []*string{
						to.Ptr("text/html"),
						to.Ptr("application/octet-stream")},
					IsCompressionEnabled: to.Ptr(true),
				},
				QueryStringCachingBehavior: to.Ptr(armcdn.AfdQueryStringCachingBehaviorIgnoreQueryString),
			},
			CustomDomains: []*armcdn.ActivatedResourceReference{
				{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/customDomains/domain1"),
				}},
			EnabledState:        to.Ptr(armcdn.EnabledStateEnabled),
			ForwardingProtocol:  to.Ptr(armcdn.ForwardingProtocolMatchRequest),
			HTTPSRedirect:       to.Ptr(armcdn.HTTPSRedirectEnabled),
			LinkToDefaultDomain: to.Ptr(armcdn.LinkToDefaultDomainEnabled),
			OriginGroup: &armcdn.ResourceReference{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/originGroups/originGroup1"),
			},
			PatternsToMatch: []*string{
				to.Ptr("/*")},
			RuleSets: []*armcdn.ResourceReference{
				{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/ruleSets/ruleSet1"),
				}},
			SupportedProtocols: []*armcdn.AFDEndpointProtocols{
				to.Ptr(armcdn.AFDEndpointProtocolsHTTPS),
				to.Ptr(armcdn.AFDEndpointProtocolsHTTP)},
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
	// res.Route = armcdn.Route{
	// 	Name: to.Ptr("route1"),
	// 	Type: to.Ptr("Microsoft.Cdn/profiles/afdendpoints/routes"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/afdendpoints/endpoint1/routes/route1"),
	// 	Properties: &armcdn.RouteProperties{
	// 		DeploymentStatus: to.Ptr(armcdn.DeploymentStatusNotStarted),
	// 		ProvisioningState: to.Ptr(armcdn.AfdProvisioningStateSucceeded),
	// 		CacheConfiguration: &armcdn.AfdRouteCacheConfiguration{
	// 			CompressionSettings: &armcdn.CompressionSettings{
	// 				ContentTypesToCompress: []*string{
	// 					to.Ptr("text/html"),
	// 					to.Ptr("application/octet-stream")},
	// 					IsCompressionEnabled: to.Ptr(true),
	// 				},
	// 				QueryStringCachingBehavior: to.Ptr(armcdn.AfdQueryStringCachingBehaviorIgnoreQueryString),
	// 			},
	// 			CustomDomains: []*armcdn.ActivatedResourceReference{
	// 				{
	// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/customDomains/domain1"),
	// 			}},
	// 			EnabledState: to.Ptr(armcdn.EnabledStateEnabled),
	// 			ForwardingProtocol: to.Ptr(armcdn.ForwardingProtocolMatchRequest),
	// 			HTTPSRedirect: to.Ptr(armcdn.HTTPSRedirectEnabled),
	// 			LinkToDefaultDomain: to.Ptr(armcdn.LinkToDefaultDomainEnabled),
	// 			OriginGroup: &armcdn.ResourceReference{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/originGroups/originGroup1"),
	// 			},
	// 			PatternsToMatch: []*string{
	// 				to.Ptr("/*")},
	// 				RuleSets: []*armcdn.ResourceReference{
	// 					{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/ruleSets/ruleSet1"),
	// 				}},
	// 				SupportedProtocols: []*armcdn.AFDEndpointProtocols{
	// 					to.Ptr(armcdn.AFDEndpointProtocolsHTTPS),
	// 					to.Ptr(armcdn.AFDEndpointProtocolsHTTP)},
	// 				},
	// 			}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7b551033155a63739b6d28f79b9c07569f6179b8/specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/Routes_Delete.json
func ExampleRoutesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRoutesClient().BeginDelete(ctx, "RG", "profile1", "endpoint1", "route1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
