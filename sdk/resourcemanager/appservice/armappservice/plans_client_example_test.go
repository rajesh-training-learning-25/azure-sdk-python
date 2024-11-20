//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b5d78da207e9c5d8f82e95224039867271f47cdf/specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/ListAppServicePlans.json
func ExamplePlansClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPlansClient().NewListPager(&armappservice.PlansClientListOptions{Detailed: nil})
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
		// page.PlanCollection = armappservice.PlanCollection{
		// 	Value: []*armappservice.Plan{
		// 		{
		// 			Name: to.Ptr("testsf6141"),
		// 			Type: to.Ptr("Microsoft.Web/serverfarms"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/serverfarms/testsf6141"),
		// 			Kind: to.Ptr("app"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappservice.PlanProperties{
		// 				GeoRegion: to.Ptr("East US"),
		// 				IsSpot: to.Ptr(false),
		// 				MaximumNumberOfWorkers: to.Ptr[int32](20),
		// 				NumberOfSites: to.Ptr[int32](4),
		// 				NumberOfWorkers: to.Ptr[int32](19),
		// 				ProvisioningState: to.Ptr(armappservice.ProvisioningStateSucceeded),
		// 				Reserved: to.Ptr(false),
		// 				Status: to.Ptr(armappservice.StatusOptionsReady),
		// 				TargetWorkerCount: to.Ptr[int32](0),
		// 				TargetWorkerSizeID: to.Ptr[int32](0),
		// 			},
		// 			SKU: &armappservice.SKUDescription{
		// 				Name: to.Ptr("P1"),
		// 				Capacity: to.Ptr[int32](1),
		// 				Family: to.Ptr("P"),
		// 				Size: to.Ptr("P1"),
		// 				Tier: to.Ptr("Premium"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testsf7252"),
		// 			Type: to.Ptr("Microsoft.Web/serverfarms"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/serverfarms/testsf7252"),
		// 			Kind: to.Ptr("app"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappservice.PlanProperties{
		// 				GeoRegion: to.Ptr("East US"),
		// 				IsSpot: to.Ptr(false),
		// 				MaximumNumberOfWorkers: to.Ptr[int32](20),
		// 				NumberOfSites: to.Ptr[int32](4),
		// 				NumberOfWorkers: to.Ptr[int32](19),
		// 				ProvisioningState: to.Ptr(armappservice.ProvisioningStateSucceeded),
		// 				Reserved: to.Ptr(false),
		// 				Status: to.Ptr(armappservice.StatusOptionsReady),
		// 				TargetWorkerCount: to.Ptr[int32](0),
		// 				TargetWorkerSizeID: to.Ptr[int32](0),
		// 			},
		// 			SKU: &armappservice.SKUDescription{
		// 				Name: to.Ptr("P1"),
		// 				Capacity: to.Ptr[int32](1),
		// 				Family: to.Ptr("P"),
		// 				Size: to.Ptr("P1"),
		// 				Tier: to.Ptr("Premium"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b5d78da207e9c5d8f82e95224039867271f47cdf/specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/ListAppServicePlansByResourceGroup.json
func ExamplePlansClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPlansClient().NewListByResourceGroupPager("testrg123", nil)
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
		// page.PlanCollection = armappservice.PlanCollection{
		// 	Value: []*armappservice.Plan{
		// 		{
		// 			Name: to.Ptr("testsf6141"),
		// 			Type: to.Ptr("Microsoft.Web/serverfarms"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/serverfarms/testsf6141"),
		// 			Kind: to.Ptr("app"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappservice.PlanProperties{
		// 				GeoRegion: to.Ptr("East US"),
		// 				IsSpot: to.Ptr(false),
		// 				MaximumNumberOfWorkers: to.Ptr[int32](20),
		// 				NumberOfSites: to.Ptr[int32](4),
		// 				NumberOfWorkers: to.Ptr[int32](19),
		// 				ProvisioningState: to.Ptr(armappservice.ProvisioningStateSucceeded),
		// 				Reserved: to.Ptr(false),
		// 				Status: to.Ptr(armappservice.StatusOptionsReady),
		// 				TargetWorkerCount: to.Ptr[int32](0),
		// 				TargetWorkerSizeID: to.Ptr[int32](0),
		// 			},
		// 			SKU: &armappservice.SKUDescription{
		// 				Name: to.Ptr("P1"),
		// 				Capacity: to.Ptr[int32](1),
		// 				Family: to.Ptr("P"),
		// 				Size: to.Ptr("P1"),
		// 				Tier: to.Ptr("Premium"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testsf7252"),
		// 			Type: to.Ptr("Microsoft.Web/serverfarms"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/serverfarms/testsf7252"),
		// 			Kind: to.Ptr("app"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappservice.PlanProperties{
		// 				GeoRegion: to.Ptr("East US"),
		// 				IsSpot: to.Ptr(false),
		// 				MaximumNumberOfWorkers: to.Ptr[int32](20),
		// 				NumberOfSites: to.Ptr[int32](4),
		// 				NumberOfWorkers: to.Ptr[int32](19),
		// 				ProvisioningState: to.Ptr(armappservice.ProvisioningStateSucceeded),
		// 				Reserved: to.Ptr(false),
		// 				Status: to.Ptr(armappservice.StatusOptionsReady),
		// 				TargetWorkerCount: to.Ptr[int32](0),
		// 				TargetWorkerSizeID: to.Ptr[int32](0),
		// 			},
		// 			SKU: &armappservice.SKUDescription{
		// 				Name: to.Ptr("P1"),
		// 				Capacity: to.Ptr[int32](1),
		// 				Family: to.Ptr("P"),
		// 				Size: to.Ptr("P1"),
		// 				Tier: to.Ptr("Premium"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b5d78da207e9c5d8f82e95224039867271f47cdf/specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/GetAppServicePlan.json
func ExamplePlansClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPlansClient().Get(ctx, "testrg123", "testsf6141", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Plan = armappservice.Plan{
	// 	Name: to.Ptr("testsf6141"),
	// 	Type: to.Ptr("Microsoft.Web/serverfarms"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/serverfarms/testsf6141"),
	// 	Kind: to.Ptr("app"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armappservice.PlanProperties{
	// 		GeoRegion: to.Ptr("East US"),
	// 		IsSpot: to.Ptr(false),
	// 		MaximumNumberOfWorkers: to.Ptr[int32](20),
	// 		NumberOfSites: to.Ptr[int32](4),
	// 		NumberOfWorkers: to.Ptr[int32](19),
	// 		ProvisioningState: to.Ptr(armappservice.ProvisioningStateSucceeded),
	// 		Reserved: to.Ptr(false),
	// 		Status: to.Ptr(armappservice.StatusOptionsReady),
	// 		TargetWorkerCount: to.Ptr[int32](0),
	// 		TargetWorkerSizeID: to.Ptr[int32](0),
	// 	},
	// 	SKU: &armappservice.SKUDescription{
	// 		Name: to.Ptr("P1"),
	// 		Capacity: to.Ptr[int32](1),
	// 		Family: to.Ptr("P"),
	// 		Size: to.Ptr("P1"),
	// 		Tier: to.Ptr("Premium"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b5d78da207e9c5d8f82e95224039867271f47cdf/specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/CreateOrUpdateAppServicePlan.json
func ExamplePlansClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPlansClient().BeginCreateOrUpdate(ctx, "testrg123", "testsf6141", armappservice.Plan{
		Kind:       to.Ptr("app"),
		Location:   to.Ptr("East US"),
		Properties: &armappservice.PlanProperties{},
		SKU: &armappservice.SKUDescription{
			Name:     to.Ptr("P1"),
			Capacity: to.Ptr[int32](1),
			Family:   to.Ptr("P"),
			Size:     to.Ptr("P1"),
			Tier:     to.Ptr("Premium"),
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
	// res.Plan = armappservice.Plan{
	// 	Name: to.Ptr("testsf6141"),
	// 	Type: to.Ptr("Microsoft.Web/serverfarms"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/serverfarms/testsf6141"),
	// 	Kind: to.Ptr("app"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armappservice.PlanProperties{
	// 		GeoRegion: to.Ptr("East US"),
	// 		MaximumNumberOfWorkers: to.Ptr[int32](20),
	// 		NumberOfSites: to.Ptr[int32](4),
	// 		ProvisioningState: to.Ptr(armappservice.ProvisioningStateSucceeded),
	// 		Reserved: to.Ptr(false),
	// 		Status: to.Ptr(armappservice.StatusOptionsReady),
	// 		TargetWorkerCount: to.Ptr[int32](0),
	// 		TargetWorkerSizeID: to.Ptr[int32](0),
	// 	},
	// 	SKU: &armappservice.SKUDescription{
	// 		Name: to.Ptr("P1"),
	// 		Capacity: to.Ptr[int32](1),
	// 		Family: to.Ptr("P"),
	// 		Size: to.Ptr("P1"),
	// 		Tier: to.Ptr("Premium"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b5d78da207e9c5d8f82e95224039867271f47cdf/specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/DeleteAppServicePlan.json
func ExamplePlansClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewPlansClient().Delete(ctx, "testrg123", "testsf6141", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b5d78da207e9c5d8f82e95224039867271f47cdf/specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/PatchAppServicePlan.json
func ExamplePlansClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPlansClient().Update(ctx, "testrg123", "testsf6141", armappservice.PlanPatchResource{
		Kind:       to.Ptr("app"),
		Properties: &armappservice.PlanPatchResourceProperties{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Plan = armappservice.Plan{
	// 	Name: to.Ptr("testsf6141"),
	// 	Type: to.Ptr("Microsoft.Web/serverfarms"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/serverfarms/testsf6141"),
	// 	Kind: to.Ptr("app"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armappservice.PlanProperties{
	// 		GeoRegion: to.Ptr("East US"),
	// 		MaximumNumberOfWorkers: to.Ptr[int32](20),
	// 		NumberOfSites: to.Ptr[int32](4),
	// 		ProvisioningState: to.Ptr(armappservice.ProvisioningStateSucceeded),
	// 		Reserved: to.Ptr(false),
	// 		Status: to.Ptr(armappservice.StatusOptionsReady),
	// 		TargetWorkerCount: to.Ptr[int32](0),
	// 		TargetWorkerSizeID: to.Ptr[int32](0),
	// 	},
	// 	SKU: &armappservice.SKUDescription{
	// 		Name: to.Ptr("P1"),
	// 		Capacity: to.Ptr[int32](1),
	// 		Family: to.Ptr("P"),
	// 		Size: to.Ptr("P1"),
	// 		Tier: to.Ptr("Premium"),
	// 	},
	// }
}
