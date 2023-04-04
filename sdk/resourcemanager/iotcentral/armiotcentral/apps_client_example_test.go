//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armiotcentral_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotcentral/armiotcentral/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/Apps_Get.json
func ExampleAppsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotcentral.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAppsClient().Get(ctx, "resRg", "myIoTCentralApp", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.App = armiotcentral.App{
	// 	Name: to.Ptr("myIoTCentralApp"),
	// 	Type: to.Ptr("Microsoft.IoTCentral/iotApps"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.IoTCentral/iotApps/myIoTCentralApp"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"key": to.Ptr("value"),
	// 	},
	// 	Identity: &armiotcentral.SystemAssignedServiceIdentity{
	// 		Type: to.Ptr(armiotcentral.SystemAssignedServiceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("8988ab15-3e7a-4586-8a1c-ed07a73a53e9"),
	// 		TenantID: to.Ptr("f686d426-8d16-42db-81b7-ab578e110ccd"),
	// 	},
	// 	Properties: &armiotcentral.AppProperties{
	// 		ApplicationID: to.Ptr("6ebd8fd3-6e34-419e-908f-9be61ec6f6d6"),
	// 		DisplayName: to.Ptr("My IoT Central App"),
	// 		NetworkRuleSets: &armiotcentral.NetworkRuleSets{
	// 			ApplyToDevices: to.Ptr(true),
	// 			ApplyToIoTCentral: to.Ptr(false),
	// 			DefaultAction: to.Ptr(armiotcentral.NetworkActionDeny),
	// 			IPRules: []*armiotcentral.NetworkRuleSetIPRule{
	// 				{
	// 					FilterName: to.Ptr("My Computer"),
	// 					IPMask: to.Ptr("21.19.35.40/24"),
	// 			}},
	// 		},
	// 		PrivateEndpointConnections: []*armiotcentral.PrivateEndpointConnection{
	// 			{
	// 				Name: to.Ptr("myIoTCentralAppEndpoint.a791c6b5-874d-4f03-9092-718490d33770"),
	// 				Type: to.Ptr("Microsoft.IoTCentral/iotApps/privateEndpointConnections"),
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.IoTCentral/iotApps/myIoTCentralApp/PrivateEndpointConnections/myIoTCentralAppEndpoint.a791c6b5-874d-4f03-9092-718490d33770"),
	// 				Properties: &armiotcentral.PrivateEndpointConnectionProperties{
	// 					GroupIDs: []*string{
	// 						to.Ptr("iotApp")},
	// 						PrivateEndpoint: &armiotcentral.PrivateEndpoint{
	// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/privateEndpoints/myIoTCentralAppEndpoint"),
	// 						},
	// 						PrivateLinkServiceConnectionState: &armiotcentral.PrivateLinkServiceConnectionState{
	// 							Description: to.Ptr("Auto-approved"),
	// 							ActionsRequired: to.Ptr("None"),
	// 							Status: to.Ptr(armiotcentral.PrivateEndpointServiceConnectionStatusApproved),
	// 						},
	// 					},
	// 			}},
	// 			PublicNetworkAccess: to.Ptr(armiotcentral.PublicNetworkAccessEnabled),
	// 			State: to.Ptr(armiotcentral.AppStateCreated),
	// 			Subdomain: to.Ptr("my-iot-central-app"),
	// 			Template: to.Ptr("iotc-pnp-preview@1.0.0"),
	// 		},
	// 		SKU: &armiotcentral.AppSKUInfo{
	// 			Name: to.Ptr(armiotcentral.AppSKUST2),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/Apps_CreateOrUpdate.json
func ExampleAppsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotcentral.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAppsClient().BeginCreateOrUpdate(ctx, "resRg", "myIoTCentralApp", armiotcentral.App{
		Location: to.Ptr("westus"),
		Identity: &armiotcentral.SystemAssignedServiceIdentity{
			Type: to.Ptr(armiotcentral.SystemAssignedServiceIdentityTypeSystemAssigned),
		},
		Properties: &armiotcentral.AppProperties{
			DisplayName: to.Ptr("My IoT Central App"),
			Subdomain:   to.Ptr("my-iot-central-app"),
			Template:    to.Ptr("iotc-pnp-preview@1.0.0"),
		},
		SKU: &armiotcentral.AppSKUInfo{
			Name: to.Ptr(armiotcentral.AppSKUST2),
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
	// res.App = armiotcentral.App{
	// 	Name: to.Ptr("myIoTCentralApp"),
	// 	Type: to.Ptr("Microsoft.IoTCentral/iotApps"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.IoTCentral/iotApps/myIoTCentralApp"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"key": to.Ptr("value"),
	// 	},
	// 	Identity: &armiotcentral.SystemAssignedServiceIdentity{
	// 		Type: to.Ptr(armiotcentral.SystemAssignedServiceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("8988ab15-3e7a-4586-8a1c-ed07a73a53e9"),
	// 		TenantID: to.Ptr("f686d426-8d16-42db-81b7-ab578e110ccd"),
	// 	},
	// 	Properties: &armiotcentral.AppProperties{
	// 		ApplicationID: to.Ptr("6ebd8fd3-6e34-419e-908f-9be61ec6f6d6"),
	// 		DisplayName: to.Ptr("My IoT Central App"),
	// 		NetworkRuleSets: &armiotcentral.NetworkRuleSets{
	// 			ApplyToDevices: to.Ptr(false),
	// 			ApplyToIoTCentral: to.Ptr(false),
	// 			DefaultAction: to.Ptr(armiotcentral.NetworkActionDeny),
	// 			IPRules: []*armiotcentral.NetworkRuleSetIPRule{
	// 			},
	// 		},
	// 		PrivateEndpointConnections: []*armiotcentral.PrivateEndpointConnection{
	// 		},
	// 		ProvisioningState: to.Ptr(armiotcentral.ProvisioningStateSucceeded),
	// 		PublicNetworkAccess: to.Ptr(armiotcentral.PublicNetworkAccessEnabled),
	// 		State: to.Ptr(armiotcentral.AppStateCreated),
	// 		Subdomain: to.Ptr("my-iot-central-app"),
	// 		Template: to.Ptr("iotc-pnp-preview@1.0.0"),
	// 	},
	// 	SKU: &armiotcentral.AppSKUInfo{
	// 		Name: to.Ptr(armiotcentral.AppSKUST2),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/Apps_Update.json
func ExampleAppsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotcentral.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAppsClient().BeginUpdate(ctx, "resRg", "myIoTCentralApp", armiotcentral.AppPatch{
		Identity: &armiotcentral.SystemAssignedServiceIdentity{
			Type: to.Ptr(armiotcentral.SystemAssignedServiceIdentityTypeSystemAssigned),
		},
		Properties: &armiotcentral.AppProperties{
			DisplayName: to.Ptr("My IoT Central App 2"),
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/Apps_Delete.json
func ExampleAppsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotcentral.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAppsClient().BeginDelete(ctx, "resRg", "myIoTCentralApp", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/Apps_ListBySubscription.json
func ExampleAppsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotcentral.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAppsClient().NewListBySubscriptionPager(nil)
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
		// page.AppListResult = armiotcentral.AppListResult{
		// 	Value: []*armiotcentral.App{
		// 		{
		// 			Name: to.Ptr("myIoTCentralApp"),
		// 			Type: to.Ptr("Microsoft.IoTCentral/iotApps"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.IoTCentral/iotApps/myIoTCentralApp"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"key": to.Ptr("value"),
		// 			},
		// 			Identity: &armiotcentral.SystemAssignedServiceIdentity{
		// 				Type: to.Ptr(armiotcentral.SystemAssignedServiceIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("8988ab15-3e7a-4586-8a1c-ed07a73a53e9"),
		// 				TenantID: to.Ptr("f686d426-8d16-42db-81b7-ab578e110ccd"),
		// 			},
		// 			Properties: &armiotcentral.AppProperties{
		// 				ApplicationID: to.Ptr("6ebd8fd3-6e34-419e-908f-9be61ec6f6d6"),
		// 				DisplayName: to.Ptr("My IoT Central App"),
		// 				NetworkRuleSets: &armiotcentral.NetworkRuleSets{
		// 					ApplyToDevices: to.Ptr(true),
		// 					ApplyToIoTCentral: to.Ptr(false),
		// 					DefaultAction: to.Ptr(armiotcentral.NetworkActionDeny),
		// 					IPRules: []*armiotcentral.NetworkRuleSetIPRule{
		// 						{
		// 							FilterName: to.Ptr("My Computer"),
		// 							IPMask: to.Ptr("21.19.35.40/24"),
		// 					}},
		// 				},
		// 				PrivateEndpointConnections: []*armiotcentral.PrivateEndpointConnection{
		// 					{
		// 						Name: to.Ptr("myIoTCentralAppEndpoint.a791c6b5-874d-4f03-9092-718490d33770"),
		// 						Type: to.Ptr("Microsoft.IoTCentral/iotApps/privateEndpointConnections"),
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.IoTCentral/iotApps/myIoTCentralApp/PrivateEndpointConnections/myIoTCentralAppEndpoint.a791c6b5-874d-4f03-9092-718490d33770"),
		// 						Properties: &armiotcentral.PrivateEndpointConnectionProperties{
		// 							GroupIDs: []*string{
		// 								to.Ptr("iotApp")},
		// 								PrivateEndpoint: &armiotcentral.PrivateEndpoint{
		// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/privateEndpoints/myIoTCentralAppEndpoint"),
		// 								},
		// 								PrivateLinkServiceConnectionState: &armiotcentral.PrivateLinkServiceConnectionState{
		// 									Description: to.Ptr("Auto-approved"),
		// 									ActionsRequired: to.Ptr("None"),
		// 									Status: to.Ptr(armiotcentral.PrivateEndpointServiceConnectionStatusApproved),
		// 								},
		// 							},
		// 					}},
		// 					ProvisioningState: to.Ptr(armiotcentral.ProvisioningStateSucceeded),
		// 					PublicNetworkAccess: to.Ptr(armiotcentral.PublicNetworkAccessEnabled),
		// 					State: to.Ptr(armiotcentral.AppStateCreated),
		// 					Subdomain: to.Ptr("my-iot-central-app"),
		// 					Template: to.Ptr("iotc-pnp-preview@1.0.0"),
		// 				},
		// 				SKU: &armiotcentral.AppSKUInfo{
		// 					Name: to.Ptr(armiotcentral.AppSKU("F1")),
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/Apps_ListByResourceGroup.json
func ExampleAppsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotcentral.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAppsClient().NewListByResourceGroupPager("resRg", nil)
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
		// page.AppListResult = armiotcentral.AppListResult{
		// 	Value: []*armiotcentral.App{
		// 		{
		// 			Name: to.Ptr("myIoTCentralApp"),
		// 			Type: to.Ptr("Microsoft.IoTCentral/iotApps"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.IoTCentral/iotApps/myIoTCentralApp"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"key": to.Ptr("value"),
		// 			},
		// 			Identity: &armiotcentral.SystemAssignedServiceIdentity{
		// 				Type: to.Ptr(armiotcentral.SystemAssignedServiceIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("8988ab15-3e7a-4586-8a1c-ed07a73a53e9"),
		// 				TenantID: to.Ptr("f686d426-8d16-42db-81b7-ab578e110ccd"),
		// 			},
		// 			Properties: &armiotcentral.AppProperties{
		// 				ApplicationID: to.Ptr("6ebd8fd3-6e34-419e-908f-9be61ec6f6d6"),
		// 				DisplayName: to.Ptr("My IoT Central App"),
		// 				NetworkRuleSets: &armiotcentral.NetworkRuleSets{
		// 					ApplyToDevices: to.Ptr(true),
		// 					ApplyToIoTCentral: to.Ptr(false),
		// 					DefaultAction: to.Ptr(armiotcentral.NetworkActionDeny),
		// 					IPRules: []*armiotcentral.NetworkRuleSetIPRule{
		// 						{
		// 							FilterName: to.Ptr("My Computer"),
		// 							IPMask: to.Ptr("21.19.35.40/24"),
		// 					}},
		// 				},
		// 				PrivateEndpointConnections: []*armiotcentral.PrivateEndpointConnection{
		// 					{
		// 						Name: to.Ptr("myIoTCentralAppEndpoint.a791c6b5-874d-4f03-9092-718490d33770"),
		// 						Type: to.Ptr("Microsoft.IoTCentral/iotApps/privateEndpointConnections"),
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.IoTCentral/iotApps/myIoTCentralApp/PrivateEndpointConnections/myIoTCentralAppEndpoint.a791c6b5-874d-4f03-9092-718490d33770"),
		// 						Properties: &armiotcentral.PrivateEndpointConnectionProperties{
		// 							GroupIDs: []*string{
		// 								to.Ptr("iotApp")},
		// 								PrivateEndpoint: &armiotcentral.PrivateEndpoint{
		// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/privateEndpoints/myIoTCentralAppEndpoint"),
		// 								},
		// 								PrivateLinkServiceConnectionState: &armiotcentral.PrivateLinkServiceConnectionState{
		// 									Description: to.Ptr("Auto-approved"),
		// 									ActionsRequired: to.Ptr("None"),
		// 									Status: to.Ptr(armiotcentral.PrivateEndpointServiceConnectionStatusApproved),
		// 								},
		// 							},
		// 					}},
		// 					PublicNetworkAccess: to.Ptr(armiotcentral.PublicNetworkAccessEnabled),
		// 					State: to.Ptr(armiotcentral.AppStateCreated),
		// 					Subdomain: to.Ptr("my-iot-central-app"),
		// 					Template: to.Ptr("iotc-pnp-preview@1.0.0"),
		// 				},
		// 				SKU: &armiotcentral.AppSKUInfo{
		// 					Name: to.Ptr(armiotcentral.AppSKU("F1")),
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/Apps_CheckNameAvailability.json
func ExampleAppsClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotcentral.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAppsClient().CheckNameAvailability(ctx, armiotcentral.OperationInputs{
		Name: to.Ptr("myiotcentralapp"),
		Type: to.Ptr("IoTApps"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AppAvailabilityInfo = armiotcentral.AppAvailabilityInfo{
	// 	NameAvailable: to.Ptr(true),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/Apps_CheckSubdomainAvailability.json
func ExampleAppsClient_CheckSubdomainAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotcentral.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAppsClient().CheckSubdomainAvailability(ctx, armiotcentral.OperationInputs{
		Name: to.Ptr("myiotcentralapp"),
		Type: to.Ptr("IoTApps"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AppAvailabilityInfo = armiotcentral.AppAvailabilityInfo{
	// 	NameAvailable: to.Ptr(true),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/Apps_Templates.json
func ExampleAppsClient_NewListTemplatesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotcentral.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAppsClient().NewListTemplatesPager(nil)
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
		// page.AppTemplatesResult = armiotcentral.AppTemplatesResult{
		// 	Value: []*armiotcentral.AppTemplate{
		// 		{
		// 			Name: to.Ptr("Store Analytics – Condition Monitoring"),
		// 			Description: to.Ptr("Digitally connect and monitor your store environment to reduce operating costs and create experiences that customers love."),
		// 			Industry: to.Ptr("Retail"),
		// 			Locations: []*armiotcentral.AppTemplateLocations{
		// 				{
		// 					DisplayName: to.Ptr("United States"),
		// 					ID: to.Ptr("unitedstates"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Europe"),
		// 					ID: to.Ptr("europe"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Asia Pacific"),
		// 					ID: to.Ptr("asiapacific"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Australia"),
		// 					ID: to.Ptr("australia"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("United Kingdom"),
		// 					ID: to.Ptr("uk"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Japan"),
		// 					ID: to.Ptr("japan"),
		// 			}},
		// 			ManifestID: to.Ptr("iotc-condition"),
		// 			ManifestVersion: to.Ptr("1.0.0"),
		// 			Order: to.Ptr[float32](99),
		// 			Title: to.Ptr("In-store Analytics – Condition Monitoring"),
		// 		},
		// 		{
		// 			Name: to.Ptr("IoT Central Water Consumption application template"),
		// 			Description: to.Ptr("Enable remote tracking of water consumption to reduce field operations, detect leaks in time, while empowering cities to conserve water."),
		// 			Industry: to.Ptr("Government"),
		// 			Locations: []*armiotcentral.AppTemplateLocations{
		// 				{
		// 					DisplayName: to.Ptr("United States"),
		// 					ID: to.Ptr("unitedstates"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Europe"),
		// 					ID: to.Ptr("europe"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Asia Pacific"),
		// 					ID: to.Ptr("asiapacific"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Australia"),
		// 					ID: to.Ptr("australia"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("United Kingdom"),
		// 					ID: to.Ptr("uk"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Japan"),
		// 					ID: to.Ptr("japan"),
		// 			}},
		// 			ManifestID: to.Ptr("iotc-consumption"),
		// 			ManifestVersion: to.Ptr("1.0.0"),
		// 			Order: to.Ptr[float32](99),
		// 			Title: to.Ptr("Water Consumption Monitoring"),
		// 		},
		// 		{
		// 			Name: to.Ptr("IoT Central Digital Distribution Center application template"),
		// 			Description: to.Ptr("Digitally manage warehouse conveyor belt system efficiency using object detection and tracking."),
		// 			Industry: to.Ptr("Retail"),
		// 			Locations: []*armiotcentral.AppTemplateLocations{
		// 				{
		// 					DisplayName: to.Ptr("United States"),
		// 					ID: to.Ptr("unitedstates"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Europe"),
		// 					ID: to.Ptr("europe"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Asia Pacific"),
		// 					ID: to.Ptr("asiapacific"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Australia"),
		// 					ID: to.Ptr("australia"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("United Kingdom"),
		// 					ID: to.Ptr("uk"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Japan"),
		// 					ID: to.Ptr("japan"),
		// 			}},
		// 			ManifestID: to.Ptr("iotc-distribution"),
		// 			ManifestVersion: to.Ptr("1.0.0"),
		// 			Order: to.Ptr[float32](99),
		// 			Title: to.Ptr("Digital Distribution Center"),
		// 		},
		// 		{
		// 			Name: to.Ptr("IoT Central Smart Inventory Management application template"),
		// 			Description: to.Ptr("Enable accurate inventory tracking and ensure shelves are always stocked."),
		// 			Industry: to.Ptr("Retail"),
		// 			Locations: []*armiotcentral.AppTemplateLocations{
		// 				{
		// 					DisplayName: to.Ptr("United States"),
		// 					ID: to.Ptr("unitedstates"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Europe"),
		// 					ID: to.Ptr("europe"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Asia Pacific"),
		// 					ID: to.Ptr("asiapacific"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Australia"),
		// 					ID: to.Ptr("australia"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("United Kingdom"),
		// 					ID: to.Ptr("uk"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Japan"),
		// 					ID: to.Ptr("japan"),
		// 			}},
		// 			ManifestID: to.Ptr("iotc-inventory"),
		// 			ManifestVersion: to.Ptr("1.0.0"),
		// 			Order: to.Ptr[float32](99),
		// 			Title: to.Ptr("Smart Inventory Management"),
		// 		},
		// 		{
		// 			Name: to.Ptr("IoT Central Connected Logistics application template"),
		// 			Description: to.Ptr("Track your shipment in real-time across air, water and land with location and condition monitoring."),
		// 			Industry: to.Ptr("Retail"),
		// 			Locations: []*armiotcentral.AppTemplateLocations{
		// 				{
		// 					DisplayName: to.Ptr("United States"),
		// 					ID: to.Ptr("unitedstates"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Europe"),
		// 					ID: to.Ptr("europe"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Asia Pacific"),
		// 					ID: to.Ptr("asiapacific"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Australia"),
		// 					ID: to.Ptr("australia"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("United Kingdom"),
		// 					ID: to.Ptr("uk"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Japan"),
		// 					ID: to.Ptr("japan"),
		// 			}},
		// 			ManifestID: to.Ptr("iotc-logistics"),
		// 			ManifestVersion: to.Ptr("1.0.0"),
		// 			Order: to.Ptr[float32](99),
		// 			Title: to.Ptr("Connected Logistics"),
		// 		},
		// 		{
		// 			Name: to.Ptr("IoT Central Smart Meter Analytics application template"),
		// 			Description: to.Ptr("Connect utility meters to gain insights into billing, forecast consumption, and proactively detect outages."),
		// 			Industry: to.Ptr("Energy"),
		// 			Locations: []*armiotcentral.AppTemplateLocations{
		// 				{
		// 					DisplayName: to.Ptr("United States"),
		// 					ID: to.Ptr("unitedstates"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Europe"),
		// 					ID: to.Ptr("europe"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Asia Pacific"),
		// 					ID: to.Ptr("asiapacific"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Australia"),
		// 					ID: to.Ptr("australia"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("United Kingdom"),
		// 					ID: to.Ptr("uk"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Japan"),
		// 					ID: to.Ptr("japan"),
		// 			}},
		// 			ManifestID: to.Ptr("iotc-meter"),
		// 			ManifestVersion: to.Ptr("1.0.0"),
		// 			Order: to.Ptr[float32](99),
		// 			Title: to.Ptr("Smart Meter Analytics"),
		// 		},
		// 		{
		// 			Name: to.Ptr("IoT Central Micro-fulfillment Center"),
		// 			Description: to.Ptr("Digitally connect, monitor and manage all aspects of a fully automated fulfillment center to reduce costs by eliminating downtime while increasing security and overall efficiency."),
		// 			Industry: to.Ptr("Retail"),
		// 			Locations: []*armiotcentral.AppTemplateLocations{
		// 				{
		// 					DisplayName: to.Ptr("United States"),
		// 					ID: to.Ptr("unitedstates"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Europe"),
		// 					ID: to.Ptr("europe"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Asia Pacific"),
		// 					ID: to.Ptr("asiapacific"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Australia"),
		// 					ID: to.Ptr("australia"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("United Kingdom"),
		// 					ID: to.Ptr("uk"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Japan"),
		// 					ID: to.Ptr("japan"),
		// 			}},
		// 			ManifestID: to.Ptr("iotc-mfc"),
		// 			ManifestVersion: to.Ptr("1.0.0"),
		// 			Order: to.Ptr[float32](99),
		// 			Title: to.Ptr("Micro-fulfillment Center"),
		// 		},
		// 		{
		// 			Name: to.Ptr("IoT Central Phone-as-a-device application template"),
		// 			Description: to.Ptr("Create application with [\"paad\"] capabilities."),
		// 			Industry: to.Ptr("Utility"),
		// 			Locations: []*armiotcentral.AppTemplateLocations{
		// 				{
		// 					DisplayName: to.Ptr("United States"),
		// 					ID: to.Ptr("unitedstates"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Europe"),
		// 					ID: to.Ptr("europe"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Asia Pacific"),
		// 					ID: to.Ptr("asiapacific"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Australia"),
		// 					ID: to.Ptr("australia"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("United Kingdom"),
		// 					ID: to.Ptr("uk"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Japan"),
		// 					ID: to.Ptr("japan"),
		// 			}},
		// 			ManifestID: to.Ptr("iotc-paad"),
		// 			ManifestVersion: to.Ptr("1.0.0"),
		// 			Order: to.Ptr[float32](99),
		// 			Title: to.Ptr("Paad"),
		// 		},
		// 		{
		// 			Name: to.Ptr("IoT Central Continuous Patient Monitoring application template"),
		// 			Description: to.Ptr("Connect and manage devices for in-patient and remote monitoring to improve patient outcomes, reduce re-admissions, and manage chronic diseases."),
		// 			Industry: to.Ptr("Health"),
		// 			Locations: []*armiotcentral.AppTemplateLocations{
		// 				{
		// 					DisplayName: to.Ptr("United States"),
		// 					ID: to.Ptr("unitedstates"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Europe"),
		// 					ID: to.Ptr("europe"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Asia Pacific"),
		// 					ID: to.Ptr("asiapacific"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Australia"),
		// 					ID: to.Ptr("australia"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("United Kingdom"),
		// 					ID: to.Ptr("uk"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Japan"),
		// 					ID: to.Ptr("japan"),
		// 			}},
		// 			ManifestID: to.Ptr("iotc-patient"),
		// 			ManifestVersion: to.Ptr("1.0.0"),
		// 			Order: to.Ptr[float32](99),
		// 			Title: to.Ptr("Continuous Patient Monitoring"),
		// 		},
		// 		{
		// 			Name: to.Ptr("IoT Central PnP template (preview)"),
		// 			Description: to.Ptr("Create an application with Azure IoT Plug and Play."),
		// 			Industry: to.Ptr(""),
		// 			Locations: []*armiotcentral.AppTemplateLocations{
		// 				{
		// 					DisplayName: to.Ptr("United States"),
		// 					ID: to.Ptr("unitedstates"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Europe"),
		// 					ID: to.Ptr("europe"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Asia Pacific"),
		// 					ID: to.Ptr("asiapacific"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Australia"),
		// 					ID: to.Ptr("australia"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("United Kingdom"),
		// 					ID: to.Ptr("uk"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Japan"),
		// 					ID: to.Ptr("japan"),
		// 			}},
		// 			ManifestID: to.Ptr("iotc-pnp-preview"),
		// 			ManifestVersion: to.Ptr("1.0.0"),
		// 			Order: to.Ptr[float32](1),
		// 			Title: to.Ptr("Custom application"),
		// 		},
		// 		{
		// 			Name: to.Ptr("IoT Central Solar Power Monitoring application template"),
		// 			Description: to.Ptr("Connect, monitor, and manage your solar panels and energy generation."),
		// 			Industry: to.Ptr("Energy"),
		// 			Locations: []*armiotcentral.AppTemplateLocations{
		// 				{
		// 					DisplayName: to.Ptr("United States"),
		// 					ID: to.Ptr("unitedstates"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Europe"),
		// 					ID: to.Ptr("europe"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Asia Pacific"),
		// 					ID: to.Ptr("asiapacific"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Australia"),
		// 					ID: to.Ptr("australia"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("United Kingdom"),
		// 					ID: to.Ptr("uk"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Japan"),
		// 					ID: to.Ptr("japan"),
		// 			}},
		// 			ManifestID: to.Ptr("iotc-power"),
		// 			ManifestVersion: to.Ptr("1.0.0"),
		// 			Order: to.Ptr[float32](99),
		// 			Title: to.Ptr("Solar Power Monitoring"),
		// 		},
		// 		{
		// 			Name: to.Ptr("IoT Central Water Quality Monitoring application template"),
		// 			Description: to.Ptr("Improve water quality and detect issues earlier by analyzing real-time measurements across your environment."),
		// 			Industry: to.Ptr("Government"),
		// 			Locations: []*armiotcentral.AppTemplateLocations{
		// 				{
		// 					DisplayName: to.Ptr("United States"),
		// 					ID: to.Ptr("unitedstates"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Europe"),
		// 					ID: to.Ptr("europe"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Asia Pacific"),
		// 					ID: to.Ptr("asiapacific"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Australia"),
		// 					ID: to.Ptr("australia"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("United Kingdom"),
		// 					ID: to.Ptr("uk"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Japan"),
		// 					ID: to.Ptr("japan"),
		// 			}},
		// 			ManifestID: to.Ptr("iotc-quality"),
		// 			ManifestVersion: to.Ptr("1.0.0"),
		// 			Order: to.Ptr[float32](99),
		// 			Title: to.Ptr("Water Quality Monitoring"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Store Analytics – Checkout"),
		// 			Description: to.Ptr("Monitor and manage the checkout flow inside your store to improve efficiency and reduce wait times."),
		// 			Industry: to.Ptr("Retail"),
		// 			Locations: []*armiotcentral.AppTemplateLocations{
		// 				{
		// 					DisplayName: to.Ptr("United States"),
		// 					ID: to.Ptr("unitedstates"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Europe"),
		// 					ID: to.Ptr("europe"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Asia Pacific"),
		// 					ID: to.Ptr("asiapacific"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Australia"),
		// 					ID: to.Ptr("australia"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("United Kingdom"),
		// 					ID: to.Ptr("uk"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Japan"),
		// 					ID: to.Ptr("japan"),
		// 			}},
		// 			ManifestID: to.Ptr("iotc-store"),
		// 			ManifestVersion: to.Ptr("1.0.0"),
		// 			Order: to.Ptr[float32](99),
		// 			Title: to.Ptr("In-store Analytics – Checkout"),
		// 		},
		// 		{
		// 			Name: to.Ptr("IoT Central Video analytics - object and motion detection application template"),
		// 			Description: to.Ptr("Use cameras as a sensor in intelligent video analytics solutions powered by Azure IoT Edge, AI, and Azure Media Services."),
		// 			Industry: to.Ptr("Retail"),
		// 			Locations: []*armiotcentral.AppTemplateLocations{
		// 				{
		// 					DisplayName: to.Ptr("United States"),
		// 					ID: to.Ptr("unitedstates"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Europe"),
		// 					ID: to.Ptr("europe"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Asia Pacific"),
		// 					ID: to.Ptr("asiapacific"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Australia"),
		// 					ID: to.Ptr("australia"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("United Kingdom"),
		// 					ID: to.Ptr("uk"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Japan"),
		// 					ID: to.Ptr("japan"),
		// 			}},
		// 			ManifestID: to.Ptr("iotc-video-analytics-om"),
		// 			ManifestVersion: to.Ptr("1.0.0"),
		// 			Order: to.Ptr[float32](99),
		// 			Title: to.Ptr("Video analytics - object and motion detection"),
		// 		},
		// 		{
		// 			Name: to.Ptr("IoT Central Connected Waste Management application template"),
		// 			Description: to.Ptr("Maximize efficiency in the collection of solid wastes by dispatching field operators at the right time along an optimized collection route."),
		// 			Industry: to.Ptr("Government"),
		// 			Locations: []*armiotcentral.AppTemplateLocations{
		// 				{
		// 					DisplayName: to.Ptr("United States"),
		// 					ID: to.Ptr("unitedstates"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Europe"),
		// 					ID: to.Ptr("europe"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Asia Pacific"),
		// 					ID: to.Ptr("asiapacific"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Australia"),
		// 					ID: to.Ptr("australia"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("United Kingdom"),
		// 					ID: to.Ptr("uk"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Japan"),
		// 					ID: to.Ptr("japan"),
		// 			}},
		// 			ManifestID: to.Ptr("iotc-waste"),
		// 			ManifestVersion: to.Ptr("1.0.0"),
		// 			Order: to.Ptr[float32](99),
		// 			Title: to.Ptr("Connected Waste Management"),
		// 	}},
		// }
	}
}
