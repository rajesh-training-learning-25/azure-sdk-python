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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b22c642b361e6d6e7d72a2347a09b0bcf6075d70/specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/storageTaskAssignmentsList/ListStorageTaskAssignmentInstancesReportSummary.json
func ExampleTaskAssignmentInstancesReportClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTaskAssignmentInstancesReportClient().NewListPager("res4228", "sto4445", "myassignment1", &armstorage.TaskAssignmentInstancesReportClientListOptions{Maxpagesize: nil,
		Filter: nil,
	})
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
		// page.TaskReportSummary = armstorage.TaskReportSummary{
		// 	Value: []*armstorage.TaskReportInstance{
		// 		{
		// 			Name: to.Ptr("instance1"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/reports"),
		// 			ID: to.Ptr("/subscriptions/1f31ba14-ce16-4281-b9b4-3e78da6e1616/resourceGroups/res4228/providers/Microsoft.Storage/storageAccounts/sto4445/storageTaskAssignments/instance1"),
		// 			Properties: &armstorage.TaskReportProperties{
		// 				FinishTime: to.Ptr("2023-06-23T00:40:10.2931264Z"),
		// 				ObjectFailedCount: to.Ptr("0"),
		// 				ObjectsOperatedOnCount: to.Ptr("150"),
		// 				ObjectsSucceededCount: to.Ptr("150"),
		// 				ObjectsTargetedCount: to.Ptr("150"),
		// 				RunResult: to.Ptr(armstorage.RunResultSucceeded),
		// 				RunStatusEnum: to.Ptr(armstorage.RunStatusEnumFinished),
		// 				RunStatusError: to.Ptr("0"),
		// 				StartTime: to.Ptr("2023-06-23T00:30:43.226744Z"),
		// 				StorageAccountID: to.Ptr("/subscriptions/1f31ba14-ce16-4281-b9b4-3e78da6e1616/resourcegroups/res4228/providers/Microsoft.Storage/storageAccounts/sto4445"),
		// 				SummaryReportPath: to.Ptr("https://acc123.blob.core.windows.net/result-container/{folderpath}/SummaryReport.json"),
		// 				TaskAssignmentID: to.Ptr("/subscriptions/1f31ba14-ce16-4281-b9b4-3e78da6e1616/resourcegroups/res4228/providers/Microsoft.Storage/storageAccounts/sto4445/storageTaskAssignments/myassignment1"),
		// 				TaskID: to.Ptr("/subscriptions/1f31ba14-ce16-4281-b9b4-3e78da6e1616/resourcegroups/res4228/providers/Microsoft.StorageActions/storageTasks/mytask1"),
		// 				TaskVersion: to.Ptr("1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("instance2"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/reports"),
		// 			ID: to.Ptr("/subscriptions/1f31ba14-ce16-4281-b9b4-3e78da6e1616/resourceGroups/res4228/providers/Microsoft.Storage/storageAccounts/sto4445/storageTaskAssignments/instance2"),
		// 			Properties: &armstorage.TaskReportProperties{
		// 				FinishTime: to.Ptr("2023-06-23T00:40:10.2931264Z"),
		// 				ObjectFailedCount: to.Ptr("0"),
		// 				ObjectsOperatedOnCount: to.Ptr("150"),
		// 				ObjectsSucceededCount: to.Ptr("150"),
		// 				ObjectsTargetedCount: to.Ptr("150"),
		// 				RunResult: to.Ptr(armstorage.RunResultSucceeded),
		// 				RunStatusEnum: to.Ptr(armstorage.RunStatusEnumFinished),
		// 				RunStatusError: to.Ptr("0"),
		// 				StartTime: to.Ptr("2023-06-23T00:30:43.226744Z"),
		// 				StorageAccountID: to.Ptr("/subscriptions/1f31ba14-ce16-4281-b9b4-3e78da6e1616/resourcegroups/res4228/providers/Microsoft.Storage/storageAccounts/sto4445"),
		// 				SummaryReportPath: to.Ptr("https://acc123.blob.core.windows.net/result-container/{folderpath}/SummaryReport.json"),
		// 				TaskAssignmentID: to.Ptr("/subscriptions/1f31ba14-ce16-4281-b9b4-3e78da6e1616/resourcegroups/res4228/providers/Microsoft.Storage/storageAccounts/sto4445/storageTaskAssignments/myassignment1"),
		// 				TaskID: to.Ptr("/subscriptions/1f31ba14-ce16-4281-b9b4-3e78da6e1616/resourcegroups/res4228/providers/Microsoft.StorageActions/storageTasks/mytask1"),
		// 				TaskVersion: to.Ptr("1"),
		// 			},
		// 	}},
		// }
	}
}
