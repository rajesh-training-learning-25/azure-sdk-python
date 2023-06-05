//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdeviceprovisioningservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0d41e635294dce73dfa99b07f3da4b68a9c9e29c/specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/preview/2023-03-01-preview/examples/DPSOperations.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceprovisioningservices.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armdeviceprovisioningservices.OperationListResult{
		// 	Value: []*armdeviceprovisioningservices.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/register/action"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Register Resource Provider"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("IotHubs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/IotHubs/diagnosticSettings/read"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Get Diagnostic Setting"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("IotHubs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/IotHubs/diagnosticSettings/write"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Set Diagnostic Setting"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("IotHubs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/IotHubs/metricDefinitions/read"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Read IotHub service metric definitions"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("IotHubs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/IotHubs/logDefinitions/read"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Read IotHub service log definitions"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("IotHubs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/operations/Read"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Get All ResourceProvider Operations"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("IotHubs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/checkNameAvailability/Action"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Check If IotHub name is available"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("IotHubs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/usages/Read"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Get Subscription Usages"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("IotHubs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/iotHubs/Read"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Get IotHub(s)"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("IotHubs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/iotHubs/Write"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Create or update IotHub Resource"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("IotHubs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/iotHubs/Delete"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Delete IotHub Resource"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("IotHubs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/iotHubs/iotHubStats/Read"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Get IotHub Statistics"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("IotHubs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/iotHubs/skus/Read"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Get valid IotHub Skus"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("IotHubs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/iotHubs/listkeys/Action"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Get all IotHub Keys"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("IotHubs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/iotHubs/iotHubKeys/listkeys/Action"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Get IotHub Key for the given name"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("IotHubs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/iotHubs/eventHubEndpoints/consumerGroups/Write"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Create EventHub Consumer Group"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("IotHubs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/iotHubs/eventHubEndpoints/consumerGroups/Read"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Get EventHub Consumer Group(s)"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("IotHubs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/iotHubs/eventHubEndpoints/consumerGroups/Delete"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Delete EventHub Consumer Group"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("IotHubs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/iotHubs/exportDevices/Action"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Export Devices"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("IotHubs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/iotHubs/importDevices/Action"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Import Devices"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("IotHubs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/iotHubs/jobs/Read"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Get the Job(s) on IotHub"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("IotHubs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/iotHubs/quotaMetrics/Read"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Get Quota Metrics"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("IotHubs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/iotHubs/routing/routes/$testall/Action"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Routing Rule Test All"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("IotHubs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/iotHubs/routing/routes/$testnew/Action"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Routing Rule Test Route"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("IotHubs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/iotHubs/routingEndpointsHealth/Read"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Get Endpoint Health"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("IotHubs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/ProvisioningServices/diagnosticSettings/read"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Get Diagnostic Setting"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("IotHubs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/ProvisioningServices/diagnosticSettings/write"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Set Diagnostic Setting"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("IotHubs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/ProvisioningServices/metricDefinitions/read"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Read DPS service metric definitions"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("IotHubs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/ProvisioningServices/logDefinitions/read"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Read DPS service log definitions"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("IotHubs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/checkProvisioningServiceNameAvailability/Action"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Check If Provisioning Service name is available"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("ProvisioningServives"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/provisioningServices/Read"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Get Provisioning Service resource"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("ProvisioningServices"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/provisioningServices/Write"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Create Provisioning Service resource"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("ProvisioningServices"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/provisioningServices/Delete"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Delete Provisioning Service resource"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("ProvisioningServices"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/provisioningServices/skus/Read"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("Delete Provisioning Service resource"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("ProvisioningServices"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Devices/provisioningServices/listkeys/Action"),
		// 			Display: &armdeviceprovisioningservices.OperationDisplay{
		// 				Operation: to.Ptr("get security related metadata"),
		// 				Provider: to.Ptr("Microsoft Devices"),
		// 				Resource: to.Ptr("ProvisioningServices"),
		// 			},
		// 	}},
		// }
	}
}
