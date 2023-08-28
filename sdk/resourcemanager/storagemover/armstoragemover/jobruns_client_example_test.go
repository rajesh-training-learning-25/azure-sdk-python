//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armstoragemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagemover/armstoragemover/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/storagemover/resource-manager/Microsoft.StorageMover/preview/2023-07-01-preview/examples/JobRuns_List.json
func ExampleJobRunsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragemover.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJobRunsClient().NewListPager("examples-rg", "examples-storageMoverName", "examples-projectName", "examples-jobDefinitionName", nil)
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
		// page.JobRunList = armstoragemover.JobRunList{
		// 	Value: []*armstoragemover.JobRun{
		// 		{
		// 			Name: to.Ptr("examples-jobRunName1"),
		// 			Type: to.Ptr("Microsoft.StorageMover/storageMovers/projects/jobDefinitions/jobRuns"),
		// 			ID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/projectName/examples-projectName/jobDefinitions/examples-jobDefinitionName1/jobRuns/examples-jobRunName1"),
		// 			Properties: &armstoragemover.JobRunProperties{
		// 				AgentName: to.Ptr("migration-agent"),
		// 				AgentResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/agents/migration-agent"),
		// 				BytesExcluded: to.Ptr[int64](995116277760),
		// 				BytesFailed: to.Ptr[int64](5116277760),
		// 				BytesNoTransferNeeded: to.Ptr[int64](2995116277760),
		// 				BytesScanned: to.Ptr[int64](49951162777600),
		// 				BytesTransferred: to.Ptr[int64](1995116277760),
		// 				BytesUnsupported: to.Ptr[int64](495116277760),
		// 				ExecutionStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T02:11:01.1075056Z"); return t}()),
		// 				ItemsExcluded: to.Ptr[int64](50),
		// 				ItemsFailed: to.Ptr[int64](3),
		// 				ItemsNoTransferNeeded: to.Ptr[int64](150),
		// 				ItemsScanned: to.Ptr[int64](351),
		// 				ItemsTransferred: to.Ptr[int64](100),
		// 				ItemsUnsupported: to.Ptr[int64](27),
		// 				JobDefinitionProperties: map[string]any{
		// 				},
		// 				LastStatusUpdate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T02:21:01.1075056Z"); return t}()),
		// 				ScanStatus: to.Ptr(armstoragemover.JobRunScanStatusScanning),
		// 				SourceName: to.Ptr("sourceEndpoint"),
		// 				SourceProperties: map[string]any{
		// 				},
		// 				SourceResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/sourceEndpoint"),
		// 				Status: to.Ptr(armstoragemover.JobRunStatusRunning),
		// 				TargetName: to.Ptr("targetEndpoint"),
		// 				TargetProperties: map[string]any{
		// 				},
		// 				TargetResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/targetEndpoint"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("examples-jobRunName2"),
		// 			Type: to.Ptr("Microsoft.StorageMover/storageMovers/projects/jobDefinitions/jobRuns"),
		// 			ID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/projectName/examples-projectName/jobDefinitions/examples-jobDefinitionName1/jobRuns/examples-jobRunName2"),
		// 			Properties: &armstoragemover.JobRunProperties{
		// 				AgentName: to.Ptr("migration-agent"),
		// 				AgentResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/agents/migration-agent"),
		// 				BytesExcluded: to.Ptr[int64](995116277760),
		// 				BytesFailed: to.Ptr[int64](5116277760),
		// 				BytesNoTransferNeeded: to.Ptr[int64](2995116277760),
		// 				BytesScanned: to.Ptr[int64](49951162777600),
		// 				BytesTransferred: to.Ptr[int64](1995116277760),
		// 				BytesUnsupported: to.Ptr[int64](495116277760),
		// 				ExecutionStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T02:11:01.1075056Z"); return t}()),
		// 				ItemsExcluded: to.Ptr[int64](50),
		// 				ItemsFailed: to.Ptr[int64](3),
		// 				ItemsNoTransferNeeded: to.Ptr[int64](150),
		// 				ItemsScanned: to.Ptr[int64](351),
		// 				ItemsTransferred: to.Ptr[int64](100),
		// 				ItemsUnsupported: to.Ptr[int64](27),
		// 				JobDefinitionProperties: map[string]any{
		// 				},
		// 				LastStatusUpdate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T02:21:01.1075056Z"); return t}()),
		// 				ScanStatus: to.Ptr(armstoragemover.JobRunScanStatusScanning),
		// 				SourceName: to.Ptr("sourceEndpoint"),
		// 				SourceProperties: map[string]any{
		// 				},
		// 				SourceResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/sourceEndpoint"),
		// 				Status: to.Ptr(armstoragemover.JobRunStatusFailed),
		// 				TargetName: to.Ptr("targetEndpoint"),
		// 				TargetProperties: map[string]any{
		// 				},
		// 				TargetResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/targetEndpoint"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("examples-jobRunName3"),
		// 			Type: to.Ptr("Microsoft.StorageMover/storageMovers/projects/jobDefinitions/jobRuns"),
		// 			ID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/projectName/examples-projectName/jobDefinitions/examples-jobDefinitionName1/jobRuns/examples-jobRunName3"),
		// 			Properties: &armstoragemover.JobRunProperties{
		// 				AgentName: to.Ptr("migration-agent"),
		// 				AgentResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/agents/migration-agent"),
		// 				BytesExcluded: to.Ptr[int64](995116277760),
		// 				BytesFailed: to.Ptr[int64](5116277760),
		// 				BytesNoTransferNeeded: to.Ptr[int64](2995116277760),
		// 				BytesScanned: to.Ptr[int64](49951162777600),
		// 				BytesTransferred: to.Ptr[int64](1995116277760),
		// 				BytesUnsupported: to.Ptr[int64](495116277760),
		// 				ExecutionStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T02:11:01.1075056Z"); return t}()),
		// 				ItemsExcluded: to.Ptr[int64](50),
		// 				ItemsFailed: to.Ptr[int64](3),
		// 				ItemsNoTransferNeeded: to.Ptr[int64](150),
		// 				ItemsScanned: to.Ptr[int64](351),
		// 				ItemsTransferred: to.Ptr[int64](100),
		// 				ItemsUnsupported: to.Ptr[int64](27),
		// 				JobDefinitionProperties: map[string]any{
		// 				},
		// 				LastStatusUpdate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T02:21:01.1075056Z"); return t}()),
		// 				ScanStatus: to.Ptr(armstoragemover.JobRunScanStatusScanning),
		// 				SourceName: to.Ptr("sourceEndpoint"),
		// 				SourceProperties: map[string]any{
		// 				},
		// 				SourceResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/sourceEndpoint"),
		// 				Status: to.Ptr(armstoragemover.JobRunStatusFailed),
		// 				TargetName: to.Ptr("targetEndpoint"),
		// 				TargetProperties: map[string]any{
		// 				},
		// 				TargetResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/targetEndpoint"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/storagemover/resource-manager/Microsoft.StorageMover/preview/2023-07-01-preview/examples/JobRuns_Get.json
func ExampleJobRunsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragemover.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobRunsClient().Get(ctx, "examples-rg", "examples-storageMoverName", "examples-projectName", "examples-jobDefinitionName", "examples-jobRunName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.JobRun = armstoragemover.JobRun{
	// 	Name: to.Ptr("examples-jobRunName"),
	// 	Type: to.Ptr("Microsoft.StorageMover/storageMovers/projects/jobDefinitions/jobRuns"),
	// 	ID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/projects/examples-projectName/jobDefinitions/examples-jobDefinitionName/jobRuns/examples-jobRunName"),
	// 	Properties: &armstoragemover.JobRunProperties{
	// 		AgentName: to.Ptr("migration-agent"),
	// 		AgentResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/agents/migration-agent"),
	// 		BytesExcluded: to.Ptr[int64](995116277760),
	// 		BytesFailed: to.Ptr[int64](5116277760),
	// 		BytesNoTransferNeeded: to.Ptr[int64](2995116277760),
	// 		BytesScanned: to.Ptr[int64](49951162777600),
	// 		BytesTransferred: to.Ptr[int64](1995116277760),
	// 		BytesUnsupported: to.Ptr[int64](495116277760),
	// 		ExecutionStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T02:11:01.1075056Z"); return t}()),
	// 		ItemsExcluded: to.Ptr[int64](50),
	// 		ItemsFailed: to.Ptr[int64](3),
	// 		ItemsNoTransferNeeded: to.Ptr[int64](150),
	// 		ItemsScanned: to.Ptr[int64](351),
	// 		ItemsTransferred: to.Ptr[int64](100),
	// 		ItemsUnsupported: to.Ptr[int64](27),
	// 		JobDefinitionProperties: map[string]any{
	// 		},
	// 		LastStatusUpdate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T02:21:01.1075056Z"); return t}()),
	// 		ScanStatus: to.Ptr(armstoragemover.JobRunScanStatusScanning),
	// 		SourceName: to.Ptr("sourceEndpoint"),
	// 		SourceProperties: map[string]any{
	// 		},
	// 		SourceResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/sourceEndpoint"),
	// 		Status: to.Ptr(armstoragemover.JobRunStatusRunning),
	// 		TargetName: to.Ptr("targetEndpoint"),
	// 		TargetProperties: map[string]any{
	// 		},
	// 		TargetResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/targetEndpoint"),
	// 	},
	// }
}
