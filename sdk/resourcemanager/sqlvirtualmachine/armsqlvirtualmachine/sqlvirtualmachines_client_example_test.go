//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsqlvirtualmachine_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-07-01-preview/examples/ListBySqlVirtualMachineGroupSqlVirtualMachine.json
func ExampleSQLVirtualMachinesClient_NewListBySQLVMGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsqlvirtualmachine.NewSQLVirtualMachinesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListBySQLVMGroupPager("testrg", "testvm", nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-07-01-preview/examples/ListSubscriptionSqlVirtualMachine.json
func ExampleSQLVirtualMachinesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsqlvirtualmachine.NewSQLVirtualMachinesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-07-01-preview/examples/StartAssessmentOnSqlVirtualMachine.json
func ExampleSQLVirtualMachinesClient_BeginStartAssessment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsqlvirtualmachine.NewSQLVirtualMachinesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginStartAssessment(ctx, "testrg", "testvm", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-07-01-preview/examples/RedeploySqlVirtualMachine.json
func ExampleSQLVirtualMachinesClient_BeginRedeploy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsqlvirtualmachine.NewSQLVirtualMachinesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginRedeploy(ctx, "testrg", "testvm", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-07-01-preview/examples/GetSqlVirtualMachine.json
func ExampleSQLVirtualMachinesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsqlvirtualmachine.NewSQLVirtualMachinesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "testrg", "testvm", &armsqlvirtualmachine.SQLVirtualMachinesClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-07-01-preview/examples/CreateOrUpdateVirtualMachineWithVMGroup.json
func ExampleSQLVirtualMachinesClient_BeginCreateOrUpdate_createsOrUpdatesASqlVirtualMachineAndJoinsItToASqlVirtualMachineGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsqlvirtualmachine.NewSQLVirtualMachinesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "testrg", "testvm", armsqlvirtualmachine.SQLVirtualMachine{
		Location: to.Ptr("northeurope"),
		Properties: &armsqlvirtualmachine.Properties{
			SQLVirtualMachineGroupResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/testvmgroup"),
			VirtualMachineResourceID:         to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachines/testvm2"),
			WsfcDomainCredentials: &armsqlvirtualmachine.WsfcDomainCredentials{
				ClusterBootstrapAccountPassword: to.Ptr("<Password>"),
				ClusterOperatorAccountPassword:  to.Ptr("<Password>"),
				SQLServiceAccountPassword:       to.Ptr("<Password>"),
			},
			WsfcStaticIP: to.Ptr("10.0.0.7"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-07-01-preview/examples/CreateOrUpdateSqlVirtualMachineAutomatedBackupWeekly.json
func ExampleSQLVirtualMachinesClient_BeginCreateOrUpdate_createsOrUpdatesASqlVirtualMachineForAutomatedBackUpSettingsWithWeeklyAndDaysOfTheWeekToRunTheBackUp() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsqlvirtualmachine.NewSQLVirtualMachinesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "testrg", "testvm", armsqlvirtualmachine.SQLVirtualMachine{
		Location: to.Ptr("northeurope"),
		Properties: &armsqlvirtualmachine.Properties{
			AutoBackupSettings: &armsqlvirtualmachine.AutoBackupSettings{
				BackupScheduleType: to.Ptr(armsqlvirtualmachine.BackupScheduleTypeManual),
				BackupSystemDbs:    to.Ptr(true),
				DaysOfWeek: []*armsqlvirtualmachine.AutoBackupDaysOfWeek{
					to.Ptr(armsqlvirtualmachine.AutoBackupDaysOfWeekMonday),
					to.Ptr(armsqlvirtualmachine.AutoBackupDaysOfWeekFriday)},
				Enable:                to.Ptr(true),
				EnableEncryption:      to.Ptr(true),
				FullBackupFrequency:   to.Ptr(armsqlvirtualmachine.FullBackupFrequencyTypeWeekly),
				FullBackupStartTime:   to.Ptr[int32](6),
				FullBackupWindowHours: to.Ptr[int32](11),
				LogBackupFrequency:    to.Ptr[int32](10),
				Password:              to.Ptr("<Password>"),
				RetentionPeriod:       to.Ptr[int32](17),
				StorageAccessKey:      to.Ptr("<primary storage access key>"),
				StorageAccountURL:     to.Ptr("https://teststorage.blob.core.windows.net/"),
				StorageContainerName:  to.Ptr("testcontainer"),
			},
			AutoPatchingSettings: &armsqlvirtualmachine.AutoPatchingSettings{
				DayOfWeek:                     to.Ptr(armsqlvirtualmachine.DayOfWeekSunday),
				Enable:                        to.Ptr(true),
				MaintenanceWindowDuration:     to.Ptr[int32](60),
				MaintenanceWindowStartingHour: to.Ptr[int32](2),
			},
			KeyVaultCredentialSettings: &armsqlvirtualmachine.KeyVaultCredentialSettings{
				Enable: to.Ptr(false),
			},
			ServerConfigurationsManagementSettings: &armsqlvirtualmachine.ServerConfigurationsManagementSettings{
				AdditionalFeaturesServerConfigurations: &armsqlvirtualmachine.AdditionalFeaturesServerConfigurations{
					IsRServicesEnabled: to.Ptr(false),
				},
				SQLConnectivityUpdateSettings: &armsqlvirtualmachine.SQLConnectivityUpdateSettings{
					ConnectivityType:      to.Ptr(armsqlvirtualmachine.ConnectivityTypePRIVATE),
					Port:                  to.Ptr[int32](1433),
					SQLAuthUpdatePassword: to.Ptr("<password>"),
					SQLAuthUpdateUserName: to.Ptr("sqllogin"),
				},
				SQLStorageUpdateSettings: &armsqlvirtualmachine.SQLStorageUpdateSettings{
					DiskConfigurationType: to.Ptr(armsqlvirtualmachine.DiskConfigurationTypeNEW),
					DiskCount:             to.Ptr[int32](1),
					StartingDeviceID:      to.Ptr[int32](2),
				},
				SQLWorkloadTypeUpdateSettings: &armsqlvirtualmachine.SQLWorkloadTypeUpdateSettings{
					SQLWorkloadType: to.Ptr(armsqlvirtualmachine.SQLWorkloadTypeOLTP),
				},
			},
			SQLImageSKU:              to.Ptr(armsqlvirtualmachine.SQLImageSKUEnterprise),
			SQLManagement:            to.Ptr(armsqlvirtualmachine.SQLManagementModeFull),
			SQLServerLicenseType:     to.Ptr(armsqlvirtualmachine.SQLServerLicenseTypePAYG),
			VirtualMachineResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachines/testvm"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-07-01-preview/examples/CreateOrUpdateSqlVirtualMachineStorageConfigurationEXTEND.json
func ExampleSQLVirtualMachinesClient_BeginCreateOrUpdate_createsOrUpdatesASqlVirtualMachineForStorageConfigurationSettingsToExtendDataLogOrTempDbStoragePool() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsqlvirtualmachine.NewSQLVirtualMachinesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "testrg", "testvm", armsqlvirtualmachine.SQLVirtualMachine{
		Location: to.Ptr("northeurope"),
		Properties: &armsqlvirtualmachine.Properties{
			StorageConfigurationSettings: &armsqlvirtualmachine.StorageConfigurationSettings{
				DiskConfigurationType: to.Ptr(armsqlvirtualmachine.DiskConfigurationTypeEXTEND),
				SQLDataSettings: &armsqlvirtualmachine.SQLStorageSettings{
					Luns: []*int32{
						to.Ptr[int32](2)},
				},
			},
			VirtualMachineResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachines/testvm"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-07-01-preview/examples/CreateOrUpdateSqlVirtualMachineStorageConfigurationNEW.json
func ExampleSQLVirtualMachinesClient_BeginCreateOrUpdate_createsOrUpdatesASqlVirtualMachineForStorageConfigurationSettingsToNewDataLogAndTempDbStoragePool() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsqlvirtualmachine.NewSQLVirtualMachinesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "testrg", "testvm", armsqlvirtualmachine.SQLVirtualMachine{
		Location: to.Ptr("northeurope"),
		Properties: &armsqlvirtualmachine.Properties{
			StorageConfigurationSettings: &armsqlvirtualmachine.StorageConfigurationSettings{
				DiskConfigurationType: to.Ptr(armsqlvirtualmachine.DiskConfigurationTypeNEW),
				SQLDataSettings: &armsqlvirtualmachine.SQLStorageSettings{
					DefaultFilePath: to.Ptr("F:\\folderpath\\"),
					Luns: []*int32{
						to.Ptr[int32](0)},
				},
				SQLLogSettings: &armsqlvirtualmachine.SQLStorageSettings{
					DefaultFilePath: to.Ptr("G:\\folderpath\\"),
					Luns: []*int32{
						to.Ptr[int32](1)},
				},
				SQLSystemDbOnDataDisk: to.Ptr(true),
				SQLTempDbSettings: &armsqlvirtualmachine.SQLTempDbSettings{
					DataFileCount:   to.Ptr[int32](8),
					DataFileSize:    to.Ptr[int32](256),
					DataGrowth:      to.Ptr[int32](512),
					DefaultFilePath: to.Ptr("D:\\TEMP"),
					LogFileSize:     to.Ptr[int32](256),
					LogGrowth:       to.Ptr[int32](512),
				},
				StorageWorkloadType: to.Ptr(armsqlvirtualmachine.StorageWorkloadTypeOLTP),
			},
			VirtualMachineResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachines/testvm"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-07-01-preview/examples/CreateOrUpdateSqlVirtualMachineMAX.json
func ExampleSQLVirtualMachinesClient_BeginCreateOrUpdate_createsOrUpdatesASqlVirtualMachineWithMaxParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsqlvirtualmachine.NewSQLVirtualMachinesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "testrg", "testvm", armsqlvirtualmachine.SQLVirtualMachine{
		Location: to.Ptr("northeurope"),
		Properties: &armsqlvirtualmachine.Properties{
			AssessmentSettings: &armsqlvirtualmachine.AssessmentSettings{
				Enable:         to.Ptr(true),
				RunImmediately: to.Ptr(true),
				Schedule: &armsqlvirtualmachine.Schedule{
					DayOfWeek:      to.Ptr(armsqlvirtualmachine.AssessmentDayOfWeekSunday),
					Enable:         to.Ptr(true),
					StartTime:      to.Ptr("23:17"),
					WeeklyInterval: to.Ptr[int32](1),
				},
			},
			AutoBackupSettings: &armsqlvirtualmachine.AutoBackupSettings{
				BackupScheduleType:    to.Ptr(armsqlvirtualmachine.BackupScheduleTypeManual),
				BackupSystemDbs:       to.Ptr(true),
				Enable:                to.Ptr(true),
				EnableEncryption:      to.Ptr(true),
				FullBackupFrequency:   to.Ptr(armsqlvirtualmachine.FullBackupFrequencyTypeDaily),
				FullBackupStartTime:   to.Ptr[int32](6),
				FullBackupWindowHours: to.Ptr[int32](11),
				LogBackupFrequency:    to.Ptr[int32](10),
				Password:              to.Ptr("<Password>"),
				RetentionPeriod:       to.Ptr[int32](17),
				StorageAccessKey:      to.Ptr("<primary storage access key>"),
				StorageAccountURL:     to.Ptr("https://teststorage.blob.core.windows.net/"),
				StorageContainerName:  to.Ptr("testcontainer"),
			},
			AutoPatchingSettings: &armsqlvirtualmachine.AutoPatchingSettings{
				DayOfWeek:                     to.Ptr(armsqlvirtualmachine.DayOfWeekSunday),
				Enable:                        to.Ptr(true),
				MaintenanceWindowDuration:     to.Ptr[int32](60),
				MaintenanceWindowStartingHour: to.Ptr[int32](2),
			},
			EnableAutomaticUpgrade: to.Ptr(true),
			KeyVaultCredentialSettings: &armsqlvirtualmachine.KeyVaultCredentialSettings{
				Enable: to.Ptr(false),
			},
			LeastPrivilegeMode: to.Ptr(armsqlvirtualmachine.LeastPrivilegeModeEnabled),
			ServerConfigurationsManagementSettings: &armsqlvirtualmachine.ServerConfigurationsManagementSettings{
				AdditionalFeaturesServerConfigurations: &armsqlvirtualmachine.AdditionalFeaturesServerConfigurations{
					IsRServicesEnabled: to.Ptr(false),
				},
				SQLConnectivityUpdateSettings: &armsqlvirtualmachine.SQLConnectivityUpdateSettings{
					ConnectivityType:      to.Ptr(armsqlvirtualmachine.ConnectivityTypePRIVATE),
					Port:                  to.Ptr[int32](1433),
					SQLAuthUpdatePassword: to.Ptr("<password>"),
					SQLAuthUpdateUserName: to.Ptr("sqllogin"),
				},
				SQLInstanceSettings: &armsqlvirtualmachine.SQLInstanceSettings{
					Collation:                          to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
					IsIfiEnabled:                       to.Ptr(true),
					IsLpimEnabled:                      to.Ptr(true),
					IsOptimizeForAdHocWorkloadsEnabled: to.Ptr(true),
					MaxDop:                             to.Ptr[int32](8),
					MaxServerMemoryMB:                  to.Ptr[int32](128),
					MinServerMemoryMB:                  to.Ptr[int32](0),
				},
				SQLStorageUpdateSettings: &armsqlvirtualmachine.SQLStorageUpdateSettings{
					DiskConfigurationType: to.Ptr(armsqlvirtualmachine.DiskConfigurationTypeNEW),
					DiskCount:             to.Ptr[int32](1),
					StartingDeviceID:      to.Ptr[int32](2),
				},
				SQLWorkloadTypeUpdateSettings: &armsqlvirtualmachine.SQLWorkloadTypeUpdateSettings{
					SQLWorkloadType: to.Ptr(armsqlvirtualmachine.SQLWorkloadTypeOLTP),
				},
			},
			SQLImageSKU:              to.Ptr(armsqlvirtualmachine.SQLImageSKUEnterprise),
			SQLManagement:            to.Ptr(armsqlvirtualmachine.SQLManagementModeFull),
			SQLServerLicenseType:     to.Ptr(armsqlvirtualmachine.SQLServerLicenseTypePAYG),
			VirtualMachineResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachines/testvm"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-07-01-preview/examples/CreateOrUpdateSqlVirtualMachineMIN.json
func ExampleSQLVirtualMachinesClient_BeginCreateOrUpdate_createsOrUpdatesASqlVirtualMachineWithMinParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsqlvirtualmachine.NewSQLVirtualMachinesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "testrg", "testvm", armsqlvirtualmachine.SQLVirtualMachine{
		Location: to.Ptr("northeurope"),
		Properties: &armsqlvirtualmachine.Properties{
			VirtualMachineResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachines/testvm"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-07-01-preview/examples/DeleteSqlVirtualMachine.json
func ExampleSQLVirtualMachinesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsqlvirtualmachine.NewSQLVirtualMachinesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx, "testrg", "testvm1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-07-01-preview/examples/UpdateSqlVirtualMachine.json
func ExampleSQLVirtualMachinesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsqlvirtualmachine.NewSQLVirtualMachinesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx, "testrg", "testvm", armsqlvirtualmachine.Update{
		Tags: map[string]*string{
			"mytag": to.Ptr("myval"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-07-01-preview/examples/ListByResourceGroupSqlVirtualMachine.json
func ExampleSQLVirtualMachinesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsqlvirtualmachine.NewSQLVirtualMachinesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByResourceGroupPager("testrg", nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}