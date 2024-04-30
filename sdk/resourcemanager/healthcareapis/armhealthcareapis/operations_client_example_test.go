//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhealthcareapis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/OperationsList.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthcareapis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page.ListOperations = armhealthcareapis.ListOperations{
		// 	Value: []*armhealthcareapis.OperationDetail{
		// 		{
		// 			Name: to.Ptr("Microsoft.HealthcareApis/services/read"),
		// 			Display: &armhealthcareapis.OperationDisplay{
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HealthcareApis/services/write"),
		// 			Display: &armhealthcareapis.OperationDisplay{
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HealthcareApis/services/delete"),
		// 			Display: &armhealthcareapis.OperationDisplay{
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HealthcareApis/locations/operationresults/read"),
		// 			Display: &armhealthcareapis.OperationDisplay{
		// 				Description: to.Ptr("Get the status of an asynchronous operation"),
		// 				Operation: to.Ptr("read"),
		// 				Provider: to.Ptr("Microsoft.HealthcareApis"),
		// 				Resource: to.Ptr("operationresults"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HealthcareApis/workspaces/read"),
		// 			Display: &armhealthcareapis.OperationDisplay{
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HealthcareApis/workspaces/write"),
		// 			Display: &armhealthcareapis.OperationDisplay{
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HealthcareApis/workspaces/delete"),
		// 			Display: &armhealthcareapis.OperationDisplay{
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HealthcareApis/workspaces/dicomservices/read"),
		// 			Display: &armhealthcareapis.OperationDisplay{
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HealthcareApis/workspaces/dicomservices/write"),
		// 			Display: &armhealthcareapis.OperationDisplay{
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HealthcareApis/workspaces/dicomservices/delete"),
		// 			Display: &armhealthcareapis.OperationDisplay{
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HealthcareApis/workspaces/iotconnectors/read"),
		// 			Display: &armhealthcareapis.OperationDisplay{
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HealthcareApis/workspaces/iotconnectors/write"),
		// 			Display: &armhealthcareapis.OperationDisplay{
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HealthcareApis/workspaces/iotconnectors/delete"),
		// 			Display: &armhealthcareapis.OperationDisplay{
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HealthcareApis/workspaces/fhirservices/read"),
		// 			Display: &armhealthcareapis.OperationDisplay{
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HealthcareApis/workspaces/fhirservices/write"),
		// 			Display: &armhealthcareapis.OperationDisplay{
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HealthcareApis/workspaces/fhirservices/delete"),
		// 			Display: &armhealthcareapis.OperationDisplay{
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HealthcareApis/checkNameAvailability/post"),
		// 			Display: &armhealthcareapis.OperationDisplay{
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HealthcareApis/Operations/read"),
		// 			Display: &armhealthcareapis.OperationDisplay{
		// 				Description: to.Ptr("Get the list of operations supported by this Resource Provider."),
		// 				Operation: to.Ptr("read"),
		// 				Provider: to.Ptr("Microsoft.HealthcareApis"),
		// 				Resource: to.Ptr("operations"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 	}},
		// }
	}
}
