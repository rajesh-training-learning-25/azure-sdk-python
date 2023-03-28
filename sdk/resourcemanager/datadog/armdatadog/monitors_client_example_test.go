//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatadog_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datadog/resource-manager/Microsoft.Datadog/stable/2021-03-01/examples/ApiKeys_List.json
func ExampleMonitorsClient_NewListAPIKeysPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatadog.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMonitorsClient().NewListAPIKeysPager("myResourceGroup", "myMonitor", nil)
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
		// page.APIKeyListResponse = armdatadog.APIKeyListResponse{
		// 	Value: []*armdatadog.APIKey{
		// 		{
		// 			Name: to.Ptr("<API_KEY_NAME>"),
		// 			Created: to.Ptr("2019-04-05 09:20:30"),
		// 			CreatedBy: to.Ptr("john@example.com"),
		// 			Key: to.Ptr("1111111111111111aaaaaaaaaaaaaaaa"),
		// 		},
		// 		{
		// 			Name: to.Ptr("<API_KEY_NAME_2>"),
		// 			Created: to.Ptr("2019-04-05 09:19:53"),
		// 			CreatedBy: to.Ptr("jane@example.com"),
		// 			Key: to.Ptr("2111111111111111aaaaaaaaaaaaaaaa"),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datadog/resource-manager/Microsoft.Datadog/stable/2021-03-01/examples/ApiKeys_GetDefaultKey.json
func ExampleMonitorsClient_GetDefaultKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatadog.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMonitorsClient().GetDefaultKey(ctx, "myResourceGroup", "myMonitor", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.APIKey = armdatadog.APIKey{
	// 	Name: to.Ptr("<API_KEY_NAME>"),
	// 	Created: to.Ptr("2019-04-05 09:20:30"),
	// 	CreatedBy: to.Ptr("john@example.com"),
	// 	Key: to.Ptr("1111111111111111aaaaaaaaaaaaaaaa"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datadog/resource-manager/Microsoft.Datadog/stable/2021-03-01/examples/ApiKeys_SetDefaultKey.json
func ExampleMonitorsClient_SetDefaultKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatadog.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewMonitorsClient().SetDefaultKey(ctx, "myResourceGroup", "myMonitor", &armdatadog.MonitorsClientSetDefaultKeyOptions{Body: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datadog/resource-manager/Microsoft.Datadog/stable/2021-03-01/examples/Hosts_List.json
func ExampleMonitorsClient_NewListHostsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatadog.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMonitorsClient().NewListHostsPager("myResourceGroup", "myMonitor", nil)
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
		// page.HostListResponse = armdatadog.HostListResponse{
		// 	Value: []*armdatadog.Host{
		// 		{
		// 			Name: to.Ptr("vm1"),
		// 			Aliases: []*string{
		// 				to.Ptr("vm1"),
		// 				to.Ptr("65f2dd83-95ae-4f56-b6aa-a5dafc05f4cd")},
		// 				Apps: []*string{
		// 					to.Ptr("ntp"),
		// 					to.Ptr("agent")},
		// 					Meta: &armdatadog.HostMetadata{
		// 						AgentVersion: to.Ptr("7.19.2"),
		// 						InstallMethod: &armdatadog.InstallMethod{
		// 							InstallerVersion: to.Ptr("install_script-1.0.0"),
		// 							Tool: to.Ptr("install_script"),
		// 							ToolVersion: to.Ptr("install_script"),
		// 						},
		// 						LogsAgent: &armdatadog.LogsAgent{
		// 							Transport: to.Ptr(""),
		// 						},
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("vm2"),
		// 					Aliases: []*string{
		// 						to.Ptr("vm2"),
		// 						to.Ptr("df631d9a-8178-4580-bf60-c697a5e8df4d")},
		// 						Apps: []*string{
		// 							to.Ptr("infra"),
		// 							to.Ptr("agent")},
		// 							Meta: &armdatadog.HostMetadata{
		// 								AgentVersion: to.Ptr("7.18.1"),
		// 								InstallMethod: &armdatadog.InstallMethod{
		// 									InstallerVersion: to.Ptr("install_script-1.0.0"),
		// 									Tool: to.Ptr("install_script"),
		// 									ToolVersion: to.Ptr("install_script"),
		// 								},
		// 								LogsAgent: &armdatadog.LogsAgent{
		// 									Transport: to.Ptr("HTTP"),
		// 								},
		// 							},
		// 					}},
		// 				}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datadog/resource-manager/Microsoft.Datadog/stable/2021-03-01/examples/LinkedResources_List.json
func ExampleMonitorsClient_NewListLinkedResourcesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatadog.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMonitorsClient().NewListLinkedResourcesPager("myResourceGroup", "myMonitor", nil)
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
		// page.LinkedResourceListResponse = armdatadog.LinkedResourceListResponse{
		// 	Value: []*armdatadog.LinkedResource{
		// 		{
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Datadog/monitors/myMonitor"),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datadog/resource-manager/Microsoft.Datadog/stable/2021-03-01/examples/MonitoredResources_List.json
func ExampleMonitorsClient_NewListMonitoredResourcesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatadog.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMonitorsClient().NewListMonitoredResourcesPager("myResourceGroup", "myMonitor", nil)
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
		// page.MonitoredResourceListResponse = armdatadog.MonitoredResourceListResponse{
		// 	Value: []*armdatadog.MonitoredResource{
		// 		{
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVault"),
		// 			ReasonForLogsStatus: to.Ptr("CapturedByRules"),
		// 			ReasonForMetricsStatus: to.Ptr("CapturedByRules"),
		// 			SendingLogs: to.Ptr(true),
		// 			SendingMetrics: to.Ptr(true),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datadog/resource-manager/Microsoft.Datadog/stable/2021-03-01/examples/Monitors_List.json
func ExampleMonitorsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatadog.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMonitorsClient().NewListPager(nil)
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
		// page.MonitorResourceListResponse = armdatadog.MonitorResourceListResponse{
		// 	Value: []*armdatadog.MonitorResource{
		// 		{
		// 			Name: to.Ptr("myMonitor"),
		// 			Type: to.Ptr("Microsoft.Datadog/monitors"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/monitors/myMonitor"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armdatadog.MonitorProperties{
		// 				DatadogOrganizationProperties: &armdatadog.OrganizationProperties{
		// 					Name: to.Ptr("myOrg"),
		// 					ID: to.Ptr("myOrg123"),
		// 				},
		// 				LiftrResourceCategory: to.Ptr(armdatadog.LiftrResourceCategoriesMonitorLogs),
		// 				LiftrResourcePreference: to.Ptr[int32](1),
		// 				MonitoringStatus: to.Ptr(armdatadog.MonitoringStatusEnabled),
		// 				ProvisioningState: to.Ptr(armdatadog.ProvisioningStateSucceeded),
		// 			},
		// 			Tags: map[string]*string{
		// 				"Environment": to.Ptr("Dev"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datadog/resource-manager/Microsoft.Datadog/stable/2021-03-01/examples/Monitors_ListByResourceGroup.json
func ExampleMonitorsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatadog.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMonitorsClient().NewListByResourceGroupPager("myResourceGroup", nil)
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
		// page.MonitorResourceListResponse = armdatadog.MonitorResourceListResponse{
		// 	Value: []*armdatadog.MonitorResource{
		// 		{
		// 			Name: to.Ptr("myMonitor"),
		// 			Type: to.Ptr("Microsoft.Datadog/monitors"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/monitors/myMonitor"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armdatadog.MonitorProperties{
		// 				DatadogOrganizationProperties: &armdatadog.OrganizationProperties{
		// 					Name: to.Ptr("myOrg"),
		// 					ID: to.Ptr("myOrg123"),
		// 				},
		// 				LiftrResourceCategory: to.Ptr(armdatadog.LiftrResourceCategoriesMonitorLogs),
		// 				LiftrResourcePreference: to.Ptr[int32](1),
		// 				MonitoringStatus: to.Ptr(armdatadog.MonitoringStatusEnabled),
		// 				ProvisioningState: to.Ptr(armdatadog.ProvisioningStateSucceeded),
		// 			},
		// 			Tags: map[string]*string{
		// 				"Environment": to.Ptr("Dev"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datadog/resource-manager/Microsoft.Datadog/stable/2021-03-01/examples/Monitors_Get.json
func ExampleMonitorsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatadog.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMonitorsClient().Get(ctx, "myResourceGroup", "myMonitor", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MonitorResource = armdatadog.MonitorResource{
	// 	Name: to.Ptr("myMonitor"),
	// 	Type: to.Ptr("Microsoft.Datadog/monitors"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/monitors/myMonitor"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armdatadog.MonitorProperties{
	// 		DatadogOrganizationProperties: &armdatadog.OrganizationProperties{
	// 			Name: to.Ptr("myOrg"),
	// 			ID: to.Ptr("myOrg123"),
	// 		},
	// 		LiftrResourceCategory: to.Ptr(armdatadog.LiftrResourceCategoriesMonitorLogs),
	// 		LiftrResourcePreference: to.Ptr[int32](1),
	// 		MonitoringStatus: to.Ptr(armdatadog.MonitoringStatusEnabled),
	// 		ProvisioningState: to.Ptr(armdatadog.ProvisioningStateSucceeded),
	// 	},
	// 	Tags: map[string]*string{
	// 		"Environment": to.Ptr("Dev"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datadog/resource-manager/Microsoft.Datadog/stable/2021-03-01/examples/Monitors_Create.json
func ExampleMonitorsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatadog.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMonitorsClient().BeginCreate(ctx, "myResourceGroup", "myMonitor", &armdatadog.MonitorsClientBeginCreateOptions{Body: nil})
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
	// res.MonitorResource = armdatadog.MonitorResource{
	// 	Name: to.Ptr("myMonitor"),
	// 	Type: to.Ptr("Microsoft.Datadog/monitors"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/monitors/myMonitor"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armdatadog.MonitorProperties{
	// 		DatadogOrganizationProperties: &armdatadog.OrganizationProperties{
	// 			Name: to.Ptr("myOrg"),
	// 			ID: to.Ptr("myOrg123"),
	// 		},
	// 		LiftrResourceCategory: to.Ptr(armdatadog.LiftrResourceCategoriesMonitorLogs),
	// 		LiftrResourcePreference: to.Ptr[int32](1),
	// 		MonitoringStatus: to.Ptr(armdatadog.MonitoringStatusEnabled),
	// 		ProvisioningState: to.Ptr(armdatadog.ProvisioningStateSucceeded),
	// 	},
	// 	SKU: &armdatadog.ResourceSKU{
	// 		Name: to.Ptr("free_Monthly"),
	// 	},
	// 	Tags: map[string]*string{
	// 		"Environment": to.Ptr("Dev"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datadog/resource-manager/Microsoft.Datadog/stable/2021-03-01/examples/Monitors_Update.json
func ExampleMonitorsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatadog.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMonitorsClient().BeginUpdate(ctx, "myResourceGroup", "myMonitor", &armdatadog.MonitorsClientBeginUpdateOptions{Body: nil})
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
	// res.MonitorResource = armdatadog.MonitorResource{
	// 	Name: to.Ptr("myMonitor"),
	// 	Type: to.Ptr("Microsoft.Datadog/monitors"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/monitors/myMonitor"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armdatadog.MonitorProperties{
	// 		DatadogOrganizationProperties: &armdatadog.OrganizationProperties{
	// 			Name: to.Ptr("myOrg"),
	// 			ID: to.Ptr("myOrg123"),
	// 		},
	// 		LiftrResourceCategory: to.Ptr(armdatadog.LiftrResourceCategoriesMonitorLogs),
	// 		LiftrResourcePreference: to.Ptr[int32](1),
	// 		MonitoringStatus: to.Ptr(armdatadog.MonitoringStatusEnabled),
	// 		ProvisioningState: to.Ptr(armdatadog.ProvisioningStateSucceeded),
	// 	},
	// 	SKU: &armdatadog.ResourceSKU{
	// 		Name: to.Ptr("free_Monthly"),
	// 	},
	// 	Tags: map[string]*string{
	// 		"Environment": to.Ptr("Dev"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datadog/resource-manager/Microsoft.Datadog/stable/2021-03-01/examples/Monitors_Delete.json
func ExampleMonitorsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatadog.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMonitorsClient().BeginDelete(ctx, "myResourceGroup", "myMonitor", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datadog/resource-manager/Microsoft.Datadog/stable/2021-03-01/examples/RefreshSetPassword_Get.json
func ExampleMonitorsClient_RefreshSetPasswordLink() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatadog.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMonitorsClient().RefreshSetPasswordLink(ctx, "myResourceGroup", "myMonitor", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SetPasswordLink = armdatadog.SetPasswordLink{
	// 	SetPasswordLink: to.Ptr("https://datadoghq.com/reset_password/tokenvalue123"),
	// }
}
