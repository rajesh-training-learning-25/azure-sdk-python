//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a498cae6d1a93f4c33073f0747b93b22815c09b7/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-02-01/examples/Common/BackupProtectionContainers_UsageSummary_Get.json
func ExampleBackupUsageSummariesClient_NewListPager_getProtectedContainersUsagesSummary() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicesbackup.NewBackupUsageSummariesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("testVault", "testRG", &armrecoveryservicesbackup.BackupUsageSummariesClientListOptions{Filter: to.Ptr("type eq 'BackupProtectionContainerCountSummary'"),
		SkipToken: nil,
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
		// page.BackupManagementUsageList = armrecoveryservicesbackup.BackupManagementUsageList{
		// 	Value: []*armrecoveryservicesbackup.BackupManagementUsage{
		// 		{
		// 			Name: &armrecoveryservicesbackup.NameInfo{
		// 				LocalizedValue: to.Ptr("Azure Backup Server"),
		// 				Value: to.Ptr("AzureBackupServer"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](2),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armrecoveryservicesbackup.UsagesUnitCount),
		// 		},
		// 		{
		// 			Name: &armrecoveryservicesbackup.NameInfo{
		// 				LocalizedValue: to.Ptr("Azure Backup Agent"),
		// 				Value: to.Ptr("MAB"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](3),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armrecoveryservicesbackup.UsagesUnitCount),
		// 		},
		// 		{
		// 			Name: &armrecoveryservicesbackup.NameInfo{
		// 				LocalizedValue: to.Ptr("SQL in Azure VM"),
		// 				Value: to.Ptr("AzureWorkload"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](1),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armrecoveryservicesbackup.UsagesUnitCount),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a498cae6d1a93f4c33073f0747b93b22815c09b7/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-02-01/examples/Common/BackupProtectedItem_UsageSummary_Get.json
func ExampleBackupUsageSummariesClient_NewListPager_getProtectedItemsUsagesSummary() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicesbackup.NewBackupUsageSummariesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("testVault", "testRG", &armrecoveryservicesbackup.BackupUsageSummariesClientListOptions{Filter: to.Ptr("type eq 'BackupProtectedItemCountSummary'"),
		SkipToken: nil,
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
		// page.BackupManagementUsageList = armrecoveryservicesbackup.BackupManagementUsageList{
		// 	Value: []*armrecoveryservicesbackup.BackupManagementUsage{
		// 		{
		// 			Name: &armrecoveryservicesbackup.NameInfo{
		// 				LocalizedValue: to.Ptr("Azure Virtual Machine"),
		// 				Value: to.Ptr("AzureIaasVM"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](7),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armrecoveryservicesbackup.UsagesUnitCount),
		// 		},
		// 		{
		// 			Name: &armrecoveryservicesbackup.NameInfo{
		// 				LocalizedValue: to.Ptr("Azure Backup Agent"),
		// 				Value: to.Ptr("MAB"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](3),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armrecoveryservicesbackup.UsagesUnitCount),
		// 		},
		// 		{
		// 			Name: &armrecoveryservicesbackup.NameInfo{
		// 				LocalizedValue: to.Ptr("Azure Backup Server"),
		// 				Value: to.Ptr("AzureBackupServer"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](1),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armrecoveryservicesbackup.UsagesUnitCount),
		// 		},
		// 		{
		// 			Name: &armrecoveryservicesbackup.NameInfo{
		// 				LocalizedValue: to.Ptr("Azure Storage (Azure Files)"),
		// 				Value: to.Ptr("AzureStorage"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](2),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armrecoveryservicesbackup.UsagesUnitCount),
		// 		},
		// 		{
		// 			Name: &armrecoveryservicesbackup.NameInfo{
		// 				LocalizedValue: to.Ptr("SQL in Azure VM"),
		// 				Value: to.Ptr("AzureWorkload"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](2),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armrecoveryservicesbackup.UsagesUnitCount),
		// 	}},
		// }
	}
}
