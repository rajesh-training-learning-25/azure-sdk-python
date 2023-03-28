//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmaintenance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0b1c0be5c97dcd5dadea4e7f975a556a78f58d60/specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2022-07-01-preview/examples/ConfigurationAssignments_GetParent.json
func ExampleConfigurationAssignmentsClient_GetParent() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmaintenance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConfigurationAssignmentsClient().GetParent(ctx, "examplerg", "Microsoft.Compute", "virtualMachineScaleSets", "smdtest1", "virtualMachines", "smdvm1", "workervmPolicy", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConfigurationAssignment = armmaintenance.ConfigurationAssignment{
	// 	Name: to.Ptr("workervmPolicy"),
	// 	Type: to.Ptr("Microsoft.Maintenance/configurationAssignments"),
	// 	ID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Compute/virtualMachineScaleSets/smdtest1/virtualMachines/smdvm1/providers/Microsoft.Maintenance/configurationAssignments/workervmPolicy"),
	// 	Properties: &armmaintenance.ConfigurationAssignmentProperties{
	// 		MaintenanceConfigurationID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Maintenance/maintenanceConfigurations/policy1"),
	// 		ResourceID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Compute/virtualMachineScaleSets/smdtest1/virtualMachines/smdvm1"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0b1c0be5c97dcd5dadea4e7f975a556a78f58d60/specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2022-07-01-preview/examples/ConfigurationAssignments_CreateOrUpdateParent.json
func ExampleConfigurationAssignmentsClient_CreateOrUpdateParent() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmaintenance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConfigurationAssignmentsClient().CreateOrUpdateParent(ctx, "examplerg", "Microsoft.Compute", "virtualMachineScaleSets", "smdtest1", "virtualMachines", "smdvm1", "workervmPolicy", armmaintenance.ConfigurationAssignment{
		Properties: &armmaintenance.ConfigurationAssignmentProperties{
			MaintenanceConfigurationID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Maintenance/maintenanceConfigurations/policy1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConfigurationAssignment = armmaintenance.ConfigurationAssignment{
	// 	Name: to.Ptr("workervmPolicy"),
	// 	Type: to.Ptr("Microsoft.Maintenance/configurationAssignments"),
	// 	ID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Compute/virtualMachineScaleSets/smdtest1/virtualMachines/smdvm1/providers/Microsoft.Maintenance/configurationAssignments/workervmPolicy"),
	// 	Properties: &armmaintenance.ConfigurationAssignmentProperties{
	// 		MaintenanceConfigurationID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Maintenance/maintenanceConfigurations/policy1"),
	// 		ResourceID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Compute/virtualMachineScaleSets/smdtest1/virtualMachines/smdvm1"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0b1c0be5c97dcd5dadea4e7f975a556a78f58d60/specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2022-07-01-preview/examples/ConfigurationAssignments_DeleteParent.json
func ExampleConfigurationAssignmentsClient_DeleteParent() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmaintenance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConfigurationAssignmentsClient().DeleteParent(ctx, "examplerg", "Microsoft.Compute", "virtualMachineScaleSets", "smdtest1", "virtualMachines", "smdvm1", "workervmConfiguration", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConfigurationAssignment = armmaintenance.ConfigurationAssignment{
	// 	Name: to.Ptr("workervmConfiguration"),
	// 	Type: to.Ptr("Microsoft.Maintenance/configurationAssignments"),
	// 	ID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Compute/virtualMachineScaleSets/smdtest1/providers/Microsoft.Maintenance/configurationAssignments/workervmConfiguration"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0b1c0be5c97dcd5dadea4e7f975a556a78f58d60/specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2022-07-01-preview/examples/ConfigurationAssignments_Get.json
func ExampleConfigurationAssignmentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmaintenance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConfigurationAssignmentsClient().Get(ctx, "examplerg", "Microsoft.Compute", "virtualMachineScaleSets", "smdtest1", "workervmConfiguration", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConfigurationAssignment = armmaintenance.ConfigurationAssignment{
	// 	Name: to.Ptr("workervmConfiguration"),
	// 	Type: to.Ptr("Microsoft.Maintenance/configurationAssignments"),
	// 	ID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Compute/virtualMachineScaleSets/smdtest1/providers/Microsoft.Maintenance/configurationAssignments/workervmConfiguration"),
	// 	Properties: &armmaintenance.ConfigurationAssignmentProperties{
	// 		MaintenanceConfigurationID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Maintenance/maintenanceConfigurations/configuration1"),
	// 		ResourceID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Compute/virtualMachineScaleSets/smdtest1"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0b1c0be5c97dcd5dadea4e7f975a556a78f58d60/specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2022-07-01-preview/examples/ConfigurationAssignments_CreateOrUpdate.json
func ExampleConfigurationAssignmentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmaintenance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConfigurationAssignmentsClient().CreateOrUpdate(ctx, "examplerg", "Microsoft.Compute", "virtualMachineScaleSets", "smdtest1", "workervmConfiguration", armmaintenance.ConfigurationAssignment{
		Properties: &armmaintenance.ConfigurationAssignmentProperties{
			MaintenanceConfigurationID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Maintenance/maintenanceConfigurations/configuration1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConfigurationAssignment = armmaintenance.ConfigurationAssignment{
	// 	Name: to.Ptr("workervmConfiguration"),
	// 	Type: to.Ptr("Microsoft.Maintenance/configurationAssignments"),
	// 	ID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Compute/virtualMachineScaleSets/smdtest1/providers/Microsoft.Maintenance/configurationAssignments/workervmConfiguration"),
	// 	Properties: &armmaintenance.ConfigurationAssignmentProperties{
	// 		MaintenanceConfigurationID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Maintenance/maintenanceConfigurations/configuration1"),
	// 		ResourceID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Compute/virtualMachineScaleSets/smdtest1"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0b1c0be5c97dcd5dadea4e7f975a556a78f58d60/specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2022-07-01-preview/examples/ConfigurationAssignments_Delete.json
func ExampleConfigurationAssignmentsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmaintenance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConfigurationAssignmentsClient().Delete(ctx, "examplerg", "Microsoft.Compute", "virtualMachineScaleSets", "smdtest1", "workervmConfiguration", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConfigurationAssignment = armmaintenance.ConfigurationAssignment{
	// 	Name: to.Ptr("workervmConfiguration"),
	// 	Type: to.Ptr("Microsoft.Maintenance/configurationAssignments"),
	// 	ID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Compute/virtualMachineScaleSets/smdtest1/providers/Microsoft.Maintenance/configurationAssignments/workervmConfiguration"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0b1c0be5c97dcd5dadea4e7f975a556a78f58d60/specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2022-07-01-preview/examples/ConfigurationAssignments_ListParent.json
func ExampleConfigurationAssignmentsClient_NewListParentPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmaintenance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConfigurationAssignmentsClient().NewListParentPager("examplerg", "Microsoft.Compute", "virtualMachineScaleSets", "smdtest1", "virtualMachines", "smdtestvm1", nil)
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
		// page.ListConfigurationAssignmentsResult = armmaintenance.ListConfigurationAssignmentsResult{
		// 	Value: []*armmaintenance.ConfigurationAssignment{
		// 		{
		// 			Name: to.Ptr("workervmPolicy"),
		// 			Type: to.Ptr("Microsoft.Maintenance/configurationAssignments"),
		// 			ID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Compute/virtualMachineScaleSets/smdtest1/virtualMachines/smdtestvm1/providers/Microsoft.Maintenance/configurationAssignments/workervmPolicy"),
		// 			Properties: &armmaintenance.ConfigurationAssignmentProperties{
		// 				MaintenanceConfigurationID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Maintenance/maintenanceConfigurations/policy1"),
		// 				ResourceID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Compute/virtualMachineScaleSets/smdtest1/virtualMachines/smdtestvm1"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0b1c0be5c97dcd5dadea4e7f975a556a78f58d60/specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2022-07-01-preview/examples/ConfigurationAssignments_List.json
func ExampleConfigurationAssignmentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmaintenance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConfigurationAssignmentsClient().NewListPager("examplerg", "Microsoft.Compute", "virtualMachineScaleSets", "smdtest1", nil)
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
		// page.ListConfigurationAssignmentsResult = armmaintenance.ListConfigurationAssignmentsResult{
		// 	Value: []*armmaintenance.ConfigurationAssignment{
		// 		{
		// 			Name: to.Ptr("workervmConfiguration"),
		// 			Type: to.Ptr("Microsoft.Maintenance/configurationAssignments"),
		// 			ID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Compute/virtualMachineScaleSets/smdtest1/providers/Microsoft.Maintenance/configurationAssignments/workervmConfiguration"),
		// 			Properties: &armmaintenance.ConfigurationAssignmentProperties{
		// 				MaintenanceConfigurationID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Maintenance/maintenanceConfigurations/configuration1"),
		// 				ResourceID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Compute/virtualMachineScaleSets/smdtest1"),
		// 			},
		// 	}},
		// }
	}
}
