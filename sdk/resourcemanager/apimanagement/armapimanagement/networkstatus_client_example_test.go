//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementServiceGetNetworkStatus.json
func ExampleNetworkStatusClient_ListByService() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNetworkStatusClient().ListByService(ctx, "rg1", "apimService1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkStatusContractByLocationArray = []*armapimanagement.NetworkStatusContractByLocation{
	// 	{
	// 		Location: to.Ptr("West US"),
	// 		NetworkStatus: &armapimanagement.NetworkStatusContract{
	// 			ConnectivityStatus: []*armapimanagement.ConnectivityStatusContract{
	// 				{
	// 					Name: to.Ptr("apimgmtst6xxxxxxxxxxx.blob.core.windows.net"),
	// 					Error: to.Ptr(""),
	// 					IsOptional: to.Ptr(false),
	// 					LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:55.9365931Z"); return t}()),
	// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:55:14.7035899Z"); return t}()),
	// 					ResourceType: to.Ptr("BlobStorage"),
	// 					Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 				},
	// 				{
	// 					Name: to.Ptr("apimgmtst6xxxxxxxxxxx.file.core.windows.net"),
	// 					Error: to.Ptr(""),
	// 					IsOptional: to.Ptr(true),
	// 					LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:55.9265938Z"); return t}()),
	// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:55:41.5322463Z"); return t}()),
	// 					ResourceType: to.Ptr("FileStorage"),
	// 					Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 				},
	// 				{
	// 					Name: to.Ptr("apimgmtst6xxxxxxxxxxx.queue.core.windows.net"),
	// 					Error: to.Ptr(""),
	// 					IsOptional: to.Ptr(true),
	// 					LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:55.8410477Z"); return t}()),
	// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:55:30.645994Z"); return t}()),
	// 					ResourceType: to.Ptr("Queue"),
	// 					Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 				},
	// 				{
	// 					Name: to.Ptr("apimgmtst6xxxxxxxxxxx.table.core.windows.net"),
	// 					Error: to.Ptr(""),
	// 					IsOptional: to.Ptr(false),
	// 					LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:55.9365931Z"); return t}()),
	// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:55:23.8789171Z"); return t}()),
	// 					ResourceType: to.Ptr("TableStorage"),
	// 					Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 				},
	// 				{
	// 					Name: to.Ptr("gcs.prod.monitoring.core.windows.net"),
	// 					Error: to.Ptr(""),
	// 					IsOptional: to.Ptr(true),
	// 					LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T08:07:37.5486932Z"); return t}()),
	// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:57:34.8666833Z"); return t}()),
	// 					ResourceType: to.Ptr("Monitoring"),
	// 					Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 				},
	// 				{
	// 					Name: to.Ptr("https://gcs.ppe.warm.ingestion.monitoring.azure.com"),
	// 					Error: to.Ptr(""),
	// 					IsOptional: to.Ptr(true),
	// 					LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:56.1060523Z"); return t}()),
	// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:56:26.1870188Z"); return t}()),
	// 					ResourceType: to.Ptr("Monitoring"),
	// 					Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 				},
	// 				{
	// 					Name: to.Ptr("https://global.metrics.nsatc.net/"),
	// 					Error: to.Ptr(""),
	// 					IsOptional: to.Ptr(true),
	// 					LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:56.0510519Z"); return t}()),
	// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:56:35.9620612Z"); return t}()),
	// 					ResourceType: to.Ptr("Monitoring"),
	// 					Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 				},
	// 				{
	// 					Name: to.Ptr("https://login.windows.net"),
	// 					Error: to.Ptr(""),
	// 					IsOptional: to.Ptr(true),
	// 					LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:56.1060523Z"); return t}()),
	// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:56:30.8047708Z"); return t}()),
	// 					ResourceType: to.Ptr("AzureActiveDirectory"),
	// 					Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 				},
	// 				{
	// 					Name: to.Ptr("https://prod2.metrics.nsatc.net:1886/RecoveryService"),
	// 					Error: to.Ptr(""),
	// 					IsOptional: to.Ptr(true),
	// 					LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:56.2796235Z"); return t}()),
	// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:56:45.2095302Z"); return t}()),
	// 					ResourceType: to.Ptr("Metrics"),
	// 					Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 				},
	// 				{
	// 					Name: to.Ptr("LocalGatewayRedis"),
	// 					Error: to.Ptr(""),
	// 					IsOptional: to.Ptr(true),
	// 					LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:55.9365931Z"); return t}()),
	// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:55:15.1345836Z"); return t}()),
	// 					ResourceType: to.Ptr("InternalCache"),
	// 					Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 				},
	// 				{
	// 					Name: to.Ptr("prod.warmpath.msftcloudes.com"),
	// 					Error: to.Ptr(""),
	// 					IsOptional: to.Ptr(false),
	// 					LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:55.8410477Z"); return t}()),
	// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:55:57.8992141Z"); return t}()),
	// 					ResourceType: to.Ptr("Monitoring"),
	// 					Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 				},
	// 				{
	// 					Name: to.Ptr("Scm"),
	// 					Error: to.Ptr(""),
	// 					IsOptional: to.Ptr(true),
	// 					LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-20T02:25:48.7066996Z"); return t}()),
	// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T23:01:24.0553684Z"); return t}()),
	// 					ResourceType: to.Ptr("SourceControl"),
	// 					Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 				},
	// 				{
	// 					Name: to.Ptr("smtpi-ch1.msn.com:25028"),
	// 					Error: to.Ptr(""),
	// 					IsOptional: to.Ptr(true),
	// 					LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:56.3510577Z"); return t}()),
	// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:58:22.2430074Z"); return t}()),
	// 					ResourceType: to.Ptr("Email"),
	// 					Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 				},
	// 				{
	// 					Name: to.Ptr("zwcvuxxxx.database.windows.net"),
	// 					Error: to.Ptr(""),
	// 					IsOptional: to.Ptr(false),
	// 					LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:56.0410467Z"); return t}()),
	// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:55:44.3582171Z"); return t}()),
	// 					ResourceType: to.Ptr("SQLDatabase"),
	// 					Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 			}},
	// 			DNSServers: []*string{
	// 				to.Ptr("10.82.98.10")},
	// 			},
	// 	}}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementServiceGetNetworkStatusByLocation.json
func ExampleNetworkStatusClient_ListByLocation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNetworkStatusClient().ListByLocation(ctx, "rg1", "apimService1", "North Central US", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkStatusContract = armapimanagement.NetworkStatusContract{
	// 	ConnectivityStatus: []*armapimanagement.ConnectivityStatusContract{
	// 		{
	// 			Name: to.Ptr("apimgmtst6tnxxxxxxxxxxx.blob.core.windows.net"),
	// 			Error: to.Ptr(""),
	// 			IsOptional: to.Ptr(false),
	// 			LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:55.9365931Z"); return t}()),
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:55:14.7035899Z"); return t}()),
	// 			ResourceType: to.Ptr("BlobStorage"),
	// 			Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 		},
	// 		{
	// 			Name: to.Ptr("apimgmtst6tnxxxxxxxxxxx.file.core.windows.net"),
	// 			Error: to.Ptr(""),
	// 			IsOptional: to.Ptr(true),
	// 			LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:55.9265938Z"); return t}()),
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:55:41.5322463Z"); return t}()),
	// 			ResourceType: to.Ptr("FileStorage"),
	// 			Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 		},
	// 		{
	// 			Name: to.Ptr("apimgmtst6tnxxxxxxxxxxx.queue.core.windows.net"),
	// 			Error: to.Ptr(""),
	// 			IsOptional: to.Ptr(true),
	// 			LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:55.8410477Z"); return t}()),
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:55:30.645994Z"); return t}()),
	// 			ResourceType: to.Ptr("Queue"),
	// 			Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 		},
	// 		{
	// 			Name: to.Ptr("apimgmtst6tnxxxxxxxxxxx.table.core.windows.net"),
	// 			Error: to.Ptr(""),
	// 			IsOptional: to.Ptr(false),
	// 			LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:55.9365931Z"); return t}()),
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:55:23.8789171Z"); return t}()),
	// 			ResourceType: to.Ptr("TableStorage"),
	// 			Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 		},
	// 		{
	// 			Name: to.Ptr("gcs.prod.monitoring.core.windows.net"),
	// 			Error: to.Ptr(""),
	// 			IsOptional: to.Ptr(true),
	// 			LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T08:07:37.5486932Z"); return t}()),
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:57:34.8666833Z"); return t}()),
	// 			ResourceType: to.Ptr("Monitoring"),
	// 			Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 		},
	// 		{
	// 			Name: to.Ptr("https://gcs.ppe.warm.ingestion.monitoring.azure.com"),
	// 			Error: to.Ptr(""),
	// 			IsOptional: to.Ptr(true),
	// 			LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:56.1060523Z"); return t}()),
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:56:26.1870188Z"); return t}()),
	// 			ResourceType: to.Ptr("Monitoring"),
	// 			Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 		},
	// 		{
	// 			Name: to.Ptr("https://global.metrics.nsatc.net/"),
	// 			Error: to.Ptr(""),
	// 			IsOptional: to.Ptr(true),
	// 			LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:56.0510519Z"); return t}()),
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:56:35.9620612Z"); return t}()),
	// 			ResourceType: to.Ptr("Monitoring"),
	// 			Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 		},
	// 		{
	// 			Name: to.Ptr("https://login.windows.net"),
	// 			Error: to.Ptr(""),
	// 			IsOptional: to.Ptr(true),
	// 			LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:56.1060523Z"); return t}()),
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:56:30.8047708Z"); return t}()),
	// 			ResourceType: to.Ptr("AzureActiveDirectory"),
	// 			Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 		},
	// 		{
	// 			Name: to.Ptr("https://prod2.metrics.nsatc.net:1886/RecoveryService"),
	// 			Error: to.Ptr(""),
	// 			IsOptional: to.Ptr(true),
	// 			LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:56.2796235Z"); return t}()),
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:56:45.2095302Z"); return t}()),
	// 			ResourceType: to.Ptr("Metrics"),
	// 			Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 		},
	// 		{
	// 			Name: to.Ptr("LocalGatewayRedis"),
	// 			Error: to.Ptr(""),
	// 			IsOptional: to.Ptr(true),
	// 			LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:55.9365931Z"); return t}()),
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:55:15.1345836Z"); return t}()),
	// 			ResourceType: to.Ptr("InternalCache"),
	// 			Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 		},
	// 		{
	// 			Name: to.Ptr("prod.warmpath.msftcloudes.com"),
	// 			Error: to.Ptr(""),
	// 			IsOptional: to.Ptr(false),
	// 			LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:55.8410477Z"); return t}()),
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:55:57.8992141Z"); return t}()),
	// 			ResourceType: to.Ptr("Monitoring"),
	// 			Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 		},
	// 		{
	// 			Name: to.Ptr("Scm"),
	// 			Error: to.Ptr(""),
	// 			IsOptional: to.Ptr(true),
	// 			LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:57.325384Z"); return t}()),
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T23:03:57.6187917Z"); return t}()),
	// 			ResourceType: to.Ptr("SourceControl"),
	// 			Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 		},
	// 		{
	// 			Name: to.Ptr("smtpi-xxx.msn.com:25028"),
	// 			Error: to.Ptr(""),
	// 			IsOptional: to.Ptr(true),
	// 			LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:56.3510577Z"); return t}()),
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:58:22.2430074Z"); return t}()),
	// 			ResourceType: to.Ptr("Email"),
	// 			Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 		},
	// 		{
	// 			Name: to.Ptr("zwcvuxxxx.database.windows.net"),
	// 			Error: to.Ptr(""),
	// 			IsOptional: to.Ptr(false),
	// 			LastStatusChange: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-20T07:54:56.0410467Z"); return t}()),
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-24T22:55:44.3582171Z"); return t}()),
	// 			ResourceType: to.Ptr("SQLDatabase"),
	// 			Status: to.Ptr(armapimanagement.ConnectivityStatusTypeSuccess),
	// 	}},
	// 	DNSServers: []*string{
	// 		to.Ptr("10.82.98.10")},
	// 	}
}
