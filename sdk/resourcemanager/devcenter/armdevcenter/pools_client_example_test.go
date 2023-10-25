//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/Pools_List.json
func ExamplePoolsClient_NewListByProjectPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPoolsClient().NewListByProjectPager("rg1", "DevProject", &armdevcenter.PoolsClientListByProjectOptions{Top: nil})
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
		// page.PoolListResult = armdevcenter.PoolListResult{
		// 	Value: []*armdevcenter.Pool{
		// 		{
		// 			Name: to.Ptr("DevPool"),
		// 			Type: to.Ptr("Microsoft.DevCenter/pools"),
		// 			ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/projects/DevProject/pools/DevPool"),
		// 			SystemData: &armdevcenter.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:00:36.993Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:30:36.993Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user1"),
		// 				LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("centralus"),
		// 			Properties: &armdevcenter.PoolProperties{
		// 				DevBoxDefinitionName: to.Ptr("WebDevBox"),
		// 				DisplayName: to.Ptr("Developer Pool"),
		// 				LicenseType: to.Ptr(armdevcenter.LicenseTypeWindowsClient),
		// 				LocalAdministrator: to.Ptr(armdevcenter.LocalAdminStatusEnabled),
		// 				ManagedVirtualNetworkRegions: []*string{
		// 					to.Ptr("centralus")},
		// 					NetworkConnectionName: to.Ptr("Network1-westus2"),
		// 					SingleSignOnStatus: to.Ptr(armdevcenter.SingleSignOnStatusDisabled),
		// 					StopOnDisconnect: &armdevcenter.StopOnDisconnectConfiguration{
		// 						GracePeriodMinutes: to.Ptr[int32](60),
		// 						Status: to.Ptr(armdevcenter.StopOnDisconnectEnableStatusEnabled),
		// 					},
		// 					VirtualNetworkType: to.Ptr(armdevcenter.VirtualNetworkTypeManaged),
		// 					DevBoxCount: to.Ptr[int32](1),
		// 					HealthStatus: to.Ptr(armdevcenter.HealthStatusHealthy),
		// 					ProvisioningState: to.Ptr(armdevcenter.ProvisioningStateSucceeded),
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/Pools_Get.json
func ExamplePoolsClient_Get_poolsGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPoolsClient().Get(ctx, "rg1", "DevProject", "DevPool", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Pool = armdevcenter.Pool{
	// 	Name: to.Ptr("DevPool"),
	// 	Type: to.Ptr("Microsoft.DevCenter/pools"),
	// 	ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/projects/DevProject/pools/DevPool"),
	// 	SystemData: &armdevcenter.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:00:36.993Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:30:36.993Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user1"),
	// 		LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("centralus"),
	// 	Properties: &armdevcenter.PoolProperties{
	// 		DevBoxDefinitionName: to.Ptr("WebDevBox"),
	// 		DisplayName: to.Ptr("Developer Pool"),
	// 		LicenseType: to.Ptr(armdevcenter.LicenseTypeWindowsClient),
	// 		LocalAdministrator: to.Ptr(armdevcenter.LocalAdminStatusEnabled),
	// 		ManagedVirtualNetworkRegions: []*string{
	// 			to.Ptr("centralus")},
	// 			NetworkConnectionName: to.Ptr("Network1-westus2"),
	// 			SingleSignOnStatus: to.Ptr(armdevcenter.SingleSignOnStatusDisabled),
	// 			StopOnDisconnect: &armdevcenter.StopOnDisconnectConfiguration{
	// 				GracePeriodMinutes: to.Ptr[int32](60),
	// 				Status: to.Ptr(armdevcenter.StopOnDisconnectEnableStatusEnabled),
	// 			},
	// 			VirtualNetworkType: to.Ptr(armdevcenter.VirtualNetworkTypeManaged),
	// 			DevBoxCount: to.Ptr[int32](1),
	// 			HealthStatus: to.Ptr(armdevcenter.HealthStatusHealthy),
	// 			ProvisioningState: to.Ptr(armdevcenter.ProvisioningStateSucceeded),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/Pools_GetUnhealthyStatus.json
func ExamplePoolsClient_Get_poolsGetUnhealthyStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPoolsClient().Get(ctx, "rg1", "DevProject", "DevPool", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Pool = armdevcenter.Pool{
	// 	Name: to.Ptr("DevPool"),
	// 	Type: to.Ptr("Microsoft.DevCenter/pools"),
	// 	ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/projects/DevProject/pools/DevPool"),
	// 	SystemData: &armdevcenter.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:00:36.993Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:30:36.993Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user1"),
	// 		LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("centralus"),
	// 	Properties: &armdevcenter.PoolProperties{
	// 		DevBoxDefinitionName: to.Ptr("WebDevBox"),
	// 		DisplayName: to.Ptr("Developer Pool"),
	// 		LicenseType: to.Ptr(armdevcenter.LicenseTypeWindowsClient),
	// 		LocalAdministrator: to.Ptr(armdevcenter.LocalAdminStatusEnabled),
	// 		ManagedVirtualNetworkRegions: []*string{
	// 			to.Ptr("centralus")},
	// 			NetworkConnectionName: to.Ptr("Network1-westus2"),
	// 			SingleSignOnStatus: to.Ptr(armdevcenter.SingleSignOnStatusDisabled),
	// 			StopOnDisconnect: &armdevcenter.StopOnDisconnectConfiguration{
	// 				GracePeriodMinutes: to.Ptr[int32](60),
	// 				Status: to.Ptr(armdevcenter.StopOnDisconnectEnableStatusEnabled),
	// 			},
	// 			VirtualNetworkType: to.Ptr(armdevcenter.VirtualNetworkTypeManaged),
	// 			DevBoxCount: to.Ptr[int32](1),
	// 			HealthStatus: to.Ptr(armdevcenter.HealthStatusUnhealthy),
	// 			HealthStatusDetails: []*armdevcenter.HealthStatusDetail{
	// 				{
	// 					Code: to.Ptr("NetworkConnectionUnhealthy"),
	// 					Message: to.Ptr("The Pool's Network Connection is in an unhealthy state. Check the Network Connection's health status for more details."),
	// 				},
	// 				{
	// 					Code: to.Ptr("ImageValidationFailed"),
	// 					Message: to.Ptr("Image validation has failed. Check the Dev Box Definition's image validation status for more details."),
	// 				},
	// 				{
	// 					Code: to.Ptr("IntuneValidationFailed"),
	// 					Message: to.Ptr("Intune license validation has failed. The tenant does not have a valid Intune license."),
	// 			}},
	// 			ProvisioningState: to.Ptr(armdevcenter.ProvisioningStateSucceeded),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/Pools_Put.json
func ExamplePoolsClient_BeginCreateOrUpdate_poolsCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPoolsClient().BeginCreateOrUpdate(ctx, "rg1", "DevProject", "DevPool", armdevcenter.Pool{
		Location: to.Ptr("centralus"),
		Properties: &armdevcenter.PoolProperties{
			DevBoxDefinitionName:  to.Ptr("WebDevBox"),
			DisplayName:           to.Ptr("Developer Pool"),
			LicenseType:           to.Ptr(armdevcenter.LicenseTypeWindowsClient),
			LocalAdministrator:    to.Ptr(armdevcenter.LocalAdminStatusEnabled),
			NetworkConnectionName: to.Ptr("Network1-westus2"),
			SingleSignOnStatus:    to.Ptr(armdevcenter.SingleSignOnStatusDisabled),
			StopOnDisconnect: &armdevcenter.StopOnDisconnectConfiguration{
				GracePeriodMinutes: to.Ptr[int32](60),
				Status:             to.Ptr(armdevcenter.StopOnDisconnectEnableStatusEnabled),
			},
			VirtualNetworkType: to.Ptr(armdevcenter.VirtualNetworkTypeUnmanaged),
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
	// res.Pool = armdevcenter.Pool{
	// 	Name: to.Ptr("DevPool"),
	// 	Type: to.Ptr("Microsoft.DevCenter/pools"),
	// 	ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/projects/DevProject/pools/DevPool"),
	// 	SystemData: &armdevcenter.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:00:36.993Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:30:36.993Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user1"),
	// 		LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("centralus"),
	// 	Properties: &armdevcenter.PoolProperties{
	// 		DevBoxDefinitionName: to.Ptr("WebDevBox"),
	// 		DisplayName: to.Ptr("Developer Pool"),
	// 		LicenseType: to.Ptr(armdevcenter.LicenseTypeWindowsClient),
	// 		LocalAdministrator: to.Ptr(armdevcenter.LocalAdminStatusEnabled),
	// 		ManagedVirtualNetworkRegions: []*string{
	// 		},
	// 		NetworkConnectionName: to.Ptr("Network1-westus2"),
	// 		SingleSignOnStatus: to.Ptr(armdevcenter.SingleSignOnStatusDisabled),
	// 		StopOnDisconnect: &armdevcenter.StopOnDisconnectConfiguration{
	// 			GracePeriodMinutes: to.Ptr[int32](60),
	// 			Status: to.Ptr(armdevcenter.StopOnDisconnectEnableStatusEnabled),
	// 		},
	// 		VirtualNetworkType: to.Ptr(armdevcenter.VirtualNetworkTypeUnmanaged),
	// 		DevBoxCount: to.Ptr[int32](1),
	// 		HealthStatus: to.Ptr(armdevcenter.HealthStatusHealthy),
	// 		ProvisioningState: to.Ptr(armdevcenter.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/Pools_PutWithManagedNetwork.json
func ExamplePoolsClient_BeginCreateOrUpdate_poolsCreateOrUpdateWithManagedNetwork() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPoolsClient().BeginCreateOrUpdate(ctx, "rg1", "DevProject", "DevPool", armdevcenter.Pool{
		Location: to.Ptr("centralus"),
		Properties: &armdevcenter.PoolProperties{
			DevBoxDefinitionName: to.Ptr("WebDevBox"),
			DisplayName:          to.Ptr("Developer Pool"),
			LicenseType:          to.Ptr(armdevcenter.LicenseTypeWindowsClient),
			LocalAdministrator:   to.Ptr(armdevcenter.LocalAdminStatusEnabled),
			ManagedVirtualNetworkRegions: []*string{
				to.Ptr("centralus")},
			NetworkConnectionName: to.Ptr("managedNetwork"),
			SingleSignOnStatus:    to.Ptr(armdevcenter.SingleSignOnStatusDisabled),
			StopOnDisconnect: &armdevcenter.StopOnDisconnectConfiguration{
				GracePeriodMinutes: to.Ptr[int32](60),
				Status:             to.Ptr(armdevcenter.StopOnDisconnectEnableStatusEnabled),
			},
			VirtualNetworkType: to.Ptr(armdevcenter.VirtualNetworkTypeManaged),
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
	// res.Pool = armdevcenter.Pool{
	// 	Name: to.Ptr("DevPool"),
	// 	Type: to.Ptr("Microsoft.DevCenter/pools"),
	// 	ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/projects/DevProject/pools/DevPool"),
	// 	SystemData: &armdevcenter.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:00:36.993Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:30:36.993Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user1"),
	// 		LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("centralus"),
	// 	Properties: &armdevcenter.PoolProperties{
	// 		DevBoxDefinitionName: to.Ptr("WebDevBox"),
	// 		DisplayName: to.Ptr("Developer Pool"),
	// 		LicenseType: to.Ptr(armdevcenter.LicenseTypeWindowsClient),
	// 		LocalAdministrator: to.Ptr(armdevcenter.LocalAdminStatusEnabled),
	// 		ManagedVirtualNetworkRegions: []*string{
	// 			to.Ptr("centralus")},
	// 			NetworkConnectionName: to.Ptr("managedNetwork"),
	// 			SingleSignOnStatus: to.Ptr(armdevcenter.SingleSignOnStatusDisabled),
	// 			StopOnDisconnect: &armdevcenter.StopOnDisconnectConfiguration{
	// 				GracePeriodMinutes: to.Ptr[int32](60),
	// 				Status: to.Ptr(armdevcenter.StopOnDisconnectEnableStatusEnabled),
	// 			},
	// 			VirtualNetworkType: to.Ptr(armdevcenter.VirtualNetworkTypeManaged),
	// 			DevBoxCount: to.Ptr[int32](1),
	// 			HealthStatus: to.Ptr(armdevcenter.HealthStatusHealthy),
	// 			ProvisioningState: to.Ptr(armdevcenter.ProvisioningStateSucceeded),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/Pools_Patch.json
func ExamplePoolsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPoolsClient().BeginUpdate(ctx, "rg1", "DevProject", "DevPool", armdevcenter.PoolUpdate{
		Properties: &armdevcenter.PoolUpdateProperties{
			DevBoxDefinitionName: to.Ptr("WebDevBox2"),
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
	// res.Pool = armdevcenter.Pool{
	// 	Name: to.Ptr("DevPool"),
	// 	Type: to.Ptr("Microsoft.DevCenter/pools"),
	// 	ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/projects/DevProject/pools/DevPool"),
	// 	SystemData: &armdevcenter.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:00:36.993Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:30:36.993Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user1"),
	// 		LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("centralus"),
	// 	Properties: &armdevcenter.PoolProperties{
	// 		DevBoxDefinitionName: to.Ptr("WebDevBox2"),
	// 		DisplayName: to.Ptr("Developer Pool"),
	// 		LicenseType: to.Ptr(armdevcenter.LicenseTypeWindowsClient),
	// 		LocalAdministrator: to.Ptr(armdevcenter.LocalAdminStatusEnabled),
	// 		ManagedVirtualNetworkRegions: []*string{
	// 			to.Ptr("centralus")},
	// 			NetworkConnectionName: to.Ptr("Network1-westus2"),
	// 			SingleSignOnStatus: to.Ptr(armdevcenter.SingleSignOnStatusDisabled),
	// 			StopOnDisconnect: &armdevcenter.StopOnDisconnectConfiguration{
	// 				GracePeriodMinutes: to.Ptr[int32](60),
	// 				Status: to.Ptr(armdevcenter.StopOnDisconnectEnableStatusEnabled),
	// 			},
	// 			VirtualNetworkType: to.Ptr(armdevcenter.VirtualNetworkTypeManaged),
	// 			DevBoxCount: to.Ptr[int32](1),
	// 			HealthStatus: to.Ptr(armdevcenter.HealthStatusHealthy),
	// 			ProvisioningState: to.Ptr(armdevcenter.ProvisioningStateSucceeded),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/Pools_Delete.json
func ExamplePoolsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPoolsClient().BeginDelete(ctx, "rg1", "DevProject", "poolName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/Pools_RunHealthChecks.json
func ExamplePoolsClient_BeginRunHealthChecks() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPoolsClient().BeginRunHealthChecks(ctx, "rg1", "DevProject", "DevPool", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
