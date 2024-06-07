//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b22c642b361e6d6e7d72a2347a09b0bcf6075d70/specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/OperationsList.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armstorage.OperationListResult{
		// 	Value: []*armstorage.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage/storageAccounts/write"),
		// 			Display: &armstorage.OperationDisplay{
		// 				Description: to.Ptr("Creates a storage account with the specified parameters or update the properties or tags or adds custom domain for the specified storage account."),
		// 				Operation: to.Ptr("Create/Update Storage Account"),
		// 				Provider: to.Ptr("Microsoft Storage"),
		// 				Resource: to.Ptr("Storage Accounts"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage/storageAccounts/delete"),
		// 			Display: &armstorage.OperationDisplay{
		// 				Description: to.Ptr("Deletes an existing storage account."),
		// 				Operation: to.Ptr("Delete Storage Account"),
		// 				Provider: to.Ptr("Microsoft Storage"),
		// 				Resource: to.Ptr("Storage Accounts"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage/storageAccounts/listkeys/action"),
		// 			Display: &armstorage.OperationDisplay{
		// 				Description: to.Ptr("Returns the access keys for the specified storage account."),
		// 				Operation: to.Ptr("List Storage Account Keys"),
		// 				Provider: to.Ptr("Microsoft Storage"),
		// 				Resource: to.Ptr("Storage Accounts"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage/storageAccounts/regeneratekey/action"),
		// 			Display: &armstorage.OperationDisplay{
		// 				Description: to.Ptr("Regenerates the access keys for the specified storage account."),
		// 				Operation: to.Ptr("Regenerate Storage Account Keys"),
		// 				Provider: to.Ptr("Microsoft Storage"),
		// 				Resource: to.Ptr("Storage Accounts"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage/checknameavailability/read"),
		// 			Display: &armstorage.OperationDisplay{
		// 				Description: to.Ptr("Checks that account name is valid and is not in use."),
		// 				Operation: to.Ptr("Check Name Availability"),
		// 				Provider: to.Ptr("Microsoft Storage"),
		// 				Resource: to.Ptr("Name Availability"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage/storageAccounts/read"),
		// 			Display: &armstorage.OperationDisplay{
		// 				Description: to.Ptr("Returns the list of storage accounts or gets the properties for the specified storage account."),
		// 				Operation: to.Ptr("List/Get Storage Account(s)"),
		// 				Provider: to.Ptr("Microsoft Storage"),
		// 				Resource: to.Ptr("Storage Accounts"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage/usages/read"),
		// 			Display: &armstorage.OperationDisplay{
		// 				Description: to.Ptr("Returns the limit and the current usage count for resources in the specified subscription"),
		// 				Operation: to.Ptr("Get Subscription Usages"),
		// 				Provider: to.Ptr("Microsoft Storage"),
		// 				Resource: to.Ptr("Usage Metrics"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage/storageAccounts/listAccountSas/action"),
		// 			Display: &armstorage.OperationDisplay{
		// 				Description: to.Ptr("Returns the Account SAS token for the specified storage account."),
		// 				Operation: to.Ptr("Returns Storage Account SAS Token"),
		// 				Provider: to.Ptr("Microsoft Storage"),
		// 				Resource: to.Ptr("Storage Account SAS Token"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage/storageAccounts/listServiceSas/action"),
		// 			Display: &armstorage.OperationDisplay{
		// 				Description: to.Ptr("Storage Service SAS Token"),
		// 				Operation: to.Ptr("Returns Storage Service SAS Token"),
		// 				Provider: to.Ptr("Microsoft Storage"),
		// 				Resource: to.Ptr("Returns the Service SAS token for the specified storage account."),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage/locations/deleteVirtualNetworkOrSubnets/action"),
		// 			Display: &armstorage.OperationDisplay{
		// 				Description: to.Ptr("Notifies Microsoft.Storage that virtual network or subnet is being deleted"),
		// 				Operation: to.Ptr("Delete virtual network or subnets notifications"),
		// 				Provider: to.Ptr("Microsoft Storage"),
		// 				Resource: to.Ptr("Location"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage/operations/read"),
		// 			Display: &armstorage.OperationDisplay{
		// 				Description: to.Ptr("Polls the status of an asynchronous operation."),
		// 				Operation: to.Ptr("Poll Asynchronous Operation"),
		// 				Provider: to.Ptr("Microsoft Storage"),
		// 				Resource: to.Ptr("Operations"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage/register/action"),
		// 			Display: &armstorage.OperationDisplay{
		// 				Description: to.Ptr("Registers the subscription for the storage resource provider and enables the creation of storage accounts."),
		// 				Operation: to.Ptr("Registers the Storage Resource Provider"),
		// 				Provider: to.Ptr("Microsoft Storage"),
		// 				Resource: to.Ptr("Storage Resource Provider"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage/skus/read"),
		// 			Display: &armstorage.OperationDisplay{
		// 				Description: to.Ptr("Lists the Skus supported by Microsoft.Storage."),
		// 				Operation: to.Ptr("List Skus"),
		// 				Provider: to.Ptr("Microsoft Storage"),
		// 				Resource: to.Ptr("Skus"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage/storageAccounts/services/diagnosticSettings/write"),
		// 			Display: &armstorage.OperationDisplay{
		// 				Description: to.Ptr("Create/Update storage account diagnostic settings."),
		// 				Operation: to.Ptr("Create/Update Diagnostic Settings"),
		// 				Provider: to.Ptr("Microsoft Storage"),
		// 				Resource: to.Ptr("Storage Accounts"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage/storageAccounts/providers/Microsoft.Insights/metricDefinitions/read"),
		// 			Display: &armstorage.OperationDisplay{
		// 				Description: to.Ptr("Get list of Microsoft Storage Metrics definitions."),
		// 				Operation: to.Ptr("Get list of Microsoft Storage Metrics definitions"),
		// 				Provider: to.Ptr("Microsoft Storage"),
		// 				Resource: to.Ptr("Storage Accounts"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 			OperationProperties: &armstorage.OperationProperties{
		// 				ServiceSpecification: &armstorage.ServiceSpecification{
		// 					MetricSpecifications: []*armstorage.MetricSpecification{
		// 						{
		// 							Name: to.Ptr("UsedCapacity"),
		// 							AggregationType: to.Ptr("Average"),
		// 							Category: to.Ptr("Capacity"),
		// 							DisplayDescription: to.Ptr("Account used capacity"),
		// 							DisplayName: to.Ptr("Used capacity"),
		// 							FillGapWithZero: to.Ptr(false),
		// 							ResourceIDDimensionNameOverride: to.Ptr("AccountResourceId"),
		// 							Unit: to.Ptr("Bytes"),
		// 					}},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage/storageAccounts/providers/Microsoft.Insights/diagnosticSettings/read"),
		// 			Display: &armstorage.OperationDisplay{
		// 				Description: to.Ptr("Gets the diagnostic setting for the resource."),
		// 				Operation: to.Ptr("Read diagnostic setting"),
		// 				Provider: to.Ptr("Microsoft Storage"),
		// 				Resource: to.Ptr("Storage Accounts"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage/storageAccounts/providers/Microsoft.Insights/diagnosticSettings/write"),
		// 			Display: &armstorage.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the diagnostic setting for the resource."),
		// 				Operation: to.Ptr("Write diagnostic setting"),
		// 				Provider: to.Ptr("Microsoft Storage"),
		// 				Resource: to.Ptr("Storage Accounts"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage/storageAccounts/blobServices/providers/Microsoft.Insights/metricDefinitions/read"),
		// 			Display: &armstorage.OperationDisplay{
		// 				Description: to.Ptr("Get list of Microsoft Storage Metrics definitions."),
		// 				Operation: to.Ptr("Get list of Microsoft Storage Metrics definitions"),
		// 				Provider: to.Ptr("Microsoft Storage"),
		// 				Resource: to.Ptr("Blob service"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 			OperationProperties: &armstorage.OperationProperties{
		// 				ServiceSpecification: &armstorage.ServiceSpecification{
		// 					MetricSpecifications: []*armstorage.MetricSpecification{
		// 						{
		// 							Name: to.Ptr("BlobCapacity"),
		// 							AggregationType: to.Ptr("Average"),
		// 							Category: to.Ptr("Capacity"),
		// 							Dimensions: []*armstorage.Dimension{
		// 								{
		// 									Name: to.Ptr("BlobType"),
		// 									DisplayName: to.Ptr("Blob type"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The amount of storage used by the storage account’s Blob service in bytes."),
		// 							DisplayName: to.Ptr("Blob Capacity"),
		// 							FillGapWithZero: to.Ptr(false),
		// 							Unit: to.Ptr("Bytes"),
		// 						},
		// 						{
		// 							Name: to.Ptr("BlobCount"),
		// 							AggregationType: to.Ptr("Average"),
		// 							Category: to.Ptr("Capacity"),
		// 							Dimensions: []*armstorage.Dimension{
		// 								{
		// 									Name: to.Ptr("BlobType"),
		// 									DisplayName: to.Ptr("Blob type"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The number of Blob in the storage account’s Blob service."),
		// 							DisplayName: to.Ptr("Blob Count"),
		// 							FillGapWithZero: to.Ptr(false),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("ContainerCount"),
		// 							AggregationType: to.Ptr("Average"),
		// 							Category: to.Ptr("Capacity"),
		// 							DisplayDescription: to.Ptr("The number of containers in the storage account’s Blob service."),
		// 							DisplayName: to.Ptr("Blob Container Count"),
		// 							FillGapWithZero: to.Ptr(false),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("BlobProvisionedSize"),
		// 							AggregationType: to.Ptr("Average"),
		// 							Category: to.Ptr("Capacity"),
		// 							Dimensions: []*armstorage.Dimension{
		// 								{
		// 									Name: to.Ptr("BlobType"),
		// 									DisplayName: to.Ptr("Blob type"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The amount of storage provisioned in the storage account’s Blob service in bytes."),
		// 							DisplayName: to.Ptr("Blob Provisioned Size"),
		// 							FillGapWithZero: to.Ptr(false),
		// 							Unit: to.Ptr("Bytes"),
		// 					}},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage/storageAccounts/blobServices/providers/Microsoft.Insights/diagnosticSettings/read"),
		// 			Display: &armstorage.OperationDisplay{
		// 				Description: to.Ptr("Gets the diagnostic setting for the resource."),
		// 				Operation: to.Ptr("Read diagnostic setting"),
		// 				Provider: to.Ptr("Microsoft Storage"),
		// 				Resource: to.Ptr("Blob service"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage/storageAccounts/blobServices/providers/Microsoft.Insights/diagnosticSettings/write"),
		// 			Display: &armstorage.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the diagnostic setting for the resource."),
		// 				Operation: to.Ptr("Write diagnostic setting"),
		// 				Provider: to.Ptr("Microsoft Storage"),
		// 				Resource: to.Ptr("Blob service"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage/storageAccounts/tableServices/providers/Microsoft.Insights/metricDefinitions/read"),
		// 			Display: &armstorage.OperationDisplay{
		// 				Description: to.Ptr("Get list of Microsoft Storage Metrics definitions."),
		// 				Operation: to.Ptr("Get list of Microsoft Storage Metrics definitions"),
		// 				Provider: to.Ptr("Microsoft Storage"),
		// 				Resource: to.Ptr("Table service"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 			OperationProperties: &armstorage.OperationProperties{
		// 				ServiceSpecification: &armstorage.ServiceSpecification{
		// 					MetricSpecifications: []*armstorage.MetricSpecification{
		// 						{
		// 							Name: to.Ptr("TableCapacity"),
		// 							AggregationType: to.Ptr("Average"),
		// 							Category: to.Ptr("Capacity"),
		// 							DisplayDescription: to.Ptr("The amount of storage used by the storage account’s Table service in bytes."),
		// 							DisplayName: to.Ptr("Table Capacity"),
		// 							FillGapWithZero: to.Ptr(false),
		// 							Unit: to.Ptr("Bytes"),
		// 						},
		// 						{
		// 							Name: to.Ptr("TableCount"),
		// 							AggregationType: to.Ptr("Average"),
		// 							Category: to.Ptr("Capacity"),
		// 							DisplayDescription: to.Ptr("The number of table in the storage account’s Table service."),
		// 							DisplayName: to.Ptr("Table Count"),
		// 							FillGapWithZero: to.Ptr(false),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("TableEntityCount"),
		// 							AggregationType: to.Ptr("Average"),
		// 							Category: to.Ptr("Capacity"),
		// 							DisplayDescription: to.Ptr("The number of table entities in the storage account’s Table service."),
		// 							DisplayName: to.Ptr("Table Entity Count"),
		// 							FillGapWithZero: to.Ptr(false),
		// 							Unit: to.Ptr("Count"),
		// 					}},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage/storageAccounts/tableServices/providers/Microsoft.Insights/diagnosticSettings/read"),
		// 			Display: &armstorage.OperationDisplay{
		// 				Description: to.Ptr("Gets the diagnostic setting for the resource."),
		// 				Operation: to.Ptr("Read diagnostic setting"),
		// 				Provider: to.Ptr("Microsoft Storage"),
		// 				Resource: to.Ptr("Table service"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage/storageAccounts/tableServices/providers/Microsoft.Insights/diagnosticSettings/write"),
		// 			Display: &armstorage.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the diagnostic setting for the resource."),
		// 				Operation: to.Ptr("Write diagnostic setting"),
		// 				Provider: to.Ptr("Microsoft Storage"),
		// 				Resource: to.Ptr("Table service"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage/storageAccounts/queueServices/providers/Microsoft.Insights/metricDefinitions/read"),
		// 			Display: &armstorage.OperationDisplay{
		// 				Description: to.Ptr("Get list of Microsoft Storage Metrics definitions."),
		// 				Operation: to.Ptr("Get list of Microsoft Storage Metrics definitions"),
		// 				Provider: to.Ptr("Microsoft Storage"),
		// 				Resource: to.Ptr("Queue service"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 			OperationProperties: &armstorage.OperationProperties{
		// 				ServiceSpecification: &armstorage.ServiceSpecification{
		// 					MetricSpecifications: []*armstorage.MetricSpecification{
		// 						{
		// 							Name: to.Ptr("QueueCapacity"),
		// 							AggregationType: to.Ptr("Average"),
		// 							Category: to.Ptr("Capacity"),
		// 							DisplayDescription: to.Ptr("The amount of storage used by the storage account’s Queue service in bytes."),
		// 							DisplayName: to.Ptr("Queue Capacity"),
		// 							FillGapWithZero: to.Ptr(false),
		// 							Unit: to.Ptr("Bytes"),
		// 						},
		// 						{
		// 							Name: to.Ptr("QueueCount"),
		// 							AggregationType: to.Ptr("Average"),
		// 							Category: to.Ptr("Capacity"),
		// 							DisplayDescription: to.Ptr("The number of queue in the storage account’s Queue service."),
		// 							DisplayName: to.Ptr("Queue Count"),
		// 							FillGapWithZero: to.Ptr(false),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("QueueMessageCount"),
		// 							AggregationType: to.Ptr("Average"),
		// 							Category: to.Ptr("Capacity"),
		// 							DisplayDescription: to.Ptr("The approximate number of queue messages in the storage account’s Queue service."),
		// 							DisplayName: to.Ptr("Queue Message Count"),
		// 							FillGapWithZero: to.Ptr(false),
		// 							Unit: to.Ptr("Count"),
		// 					}},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage/storageAccounts/queueServices/providers/Microsoft.Insights/diagnosticSettings/read"),
		// 			Display: &armstorage.OperationDisplay{
		// 				Description: to.Ptr("Gets the diagnostic setting for the resource."),
		// 				Operation: to.Ptr("Read diagnostic setting"),
		// 				Provider: to.Ptr("Microsoft Storage"),
		// 				Resource: to.Ptr("Queue service"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage/storageAccounts/queueServices/providers/Microsoft.Insights/diagnosticSettings/write"),
		// 			Display: &armstorage.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the diagnostic setting for the resource."),
		// 				Operation: to.Ptr("Write diagnostic setting"),
		// 				Provider: to.Ptr("Microsoft Storage"),
		// 				Resource: to.Ptr("Queue service"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage/storageAccounts/fileServices/providers/Microsoft.Insights/metricDefinitions/read"),
		// 			Display: &armstorage.OperationDisplay{
		// 				Description: to.Ptr("Get list of Microsoft Storage Metrics definitions."),
		// 				Operation: to.Ptr("Get list of Microsoft Storage Metrics definitions"),
		// 				Provider: to.Ptr("Microsoft Storage"),
		// 				Resource: to.Ptr("File service"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 			OperationProperties: &armstorage.OperationProperties{
		// 				ServiceSpecification: &armstorage.ServiceSpecification{
		// 					MetricSpecifications: []*armstorage.MetricSpecification{
		// 						{
		// 							Name: to.Ptr("FileCapacity"),
		// 							AggregationType: to.Ptr("Average"),
		// 							Category: to.Ptr("Capacity"),
		// 							DisplayDescription: to.Ptr("The amount of storage used by the storage account’s File service in bytes."),
		// 							DisplayName: to.Ptr("File Capacity"),
		// 							FillGapWithZero: to.Ptr(false),
		// 							Unit: to.Ptr("Bytes"),
		// 						},
		// 						{
		// 							Name: to.Ptr("FileProvisionedSize"),
		// 							AggregationType: to.Ptr("Average"),
		// 							Category: to.Ptr("Capacity"),
		// 							DisplayDescription: to.Ptr("The amount of storage provisioned in the storage account’s File service in bytes."),
		// 							DisplayName: to.Ptr("File Provisioned Size"),
		// 							FillGapWithZero: to.Ptr(false),
		// 							Unit: to.Ptr("Bytes"),
		// 						},
		// 						{
		// 							Name: to.Ptr("FileCount"),
		// 							AggregationType: to.Ptr("Average"),
		// 							Category: to.Ptr("Capacity"),
		// 							DisplayDescription: to.Ptr("The number of file in the storage account’s File service."),
		// 							DisplayName: to.Ptr("File Count"),
		// 							FillGapWithZero: to.Ptr(false),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("FileShareCount"),
		// 							AggregationType: to.Ptr("Average"),
		// 							Category: to.Ptr("Capacity"),
		// 							DisplayDescription: to.Ptr("The number of file shares in the storage account’s File service."),
		// 							DisplayName: to.Ptr("File Share Count"),
		// 							FillGapWithZero: to.Ptr(false),
		// 							Unit: to.Ptr("Count"),
		// 					}},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage/storageAccounts/fileServices/providers/Microsoft.Insights/diagnosticSettings/read"),
		// 			Display: &armstorage.OperationDisplay{
		// 				Description: to.Ptr("Gets the diagnostic setting for the resource."),
		// 				Operation: to.Ptr("Read diagnostic setting"),
		// 				Provider: to.Ptr("Microsoft Storage"),
		// 				Resource: to.Ptr("File service"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage/storageAccounts/fileServices/providers/Microsoft.Insights/diagnosticSettings/write"),
		// 			Display: &armstorage.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the diagnostic setting for the resource."),
		// 				Operation: to.Ptr("Write diagnostic setting"),
		// 				Provider: to.Ptr("Microsoft Storage"),
		// 				Resource: to.Ptr("File service"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 	}},
		// }
	}
}
