//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/BackupInstanceOperations/ListBackupInstances.json
func ExampleBackupInstancesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("04cf684a-d41f-4550-9f70-7708a3a2283b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("000pikumar", "PratikPrivatePreviewVault1", nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/BackupInstanceOperations/GetBackupInstance.json
func ExampleBackupInstancesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("04cf684a-d41f-4550-9f70-7708a3a2283b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "000pikumar", "PratikPrivatePreviewVault1", "testInstance1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/BackupInstanceOperations/PutBackupInstance.json
func ExampleBackupInstancesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("04cf684a-d41f-4550-9f70-7708a3a2283b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "000pikumar", "PratikPrivatePreviewVault1", "testInstance1", armdataprotection.BackupInstanceResource{
		Tags: map[string]*string{
			"key1": to.Ptr("val1"),
		},
		Properties: &armdataprotection.BackupInstance{
			DataSourceInfo: &armdataprotection.Datasource{
				DatasourceType:   to.Ptr("Microsoft.DBforPostgreSQL/servers/databases"),
				ObjectType:       to.Ptr("Datasource"),
				ResourceID:       to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest/databases/testdb"),
				ResourceLocation: to.Ptr(""),
				ResourceName:     to.Ptr("testdb"),
				ResourceType:     to.Ptr("Microsoft.DBforPostgreSQL/servers/databases"),
				ResourceURI:      to.Ptr(""),
			},
			DataSourceSetInfo: &armdataprotection.DatasourceSet{
				DatasourceType:   to.Ptr("Microsoft.DBforPostgreSQL/servers/databases"),
				ObjectType:       to.Ptr("DatasourceSet"),
				ResourceID:       to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest"),
				ResourceLocation: to.Ptr(""),
				ResourceName:     to.Ptr("viveksipgtest"),
				ResourceType:     to.Ptr("Microsoft.DBforPostgreSQL/servers"),
				ResourceURI:      to.Ptr(""),
			},
			DatasourceAuthCredentials: &armdataprotection.SecretStoreBasedAuthCredentials{
				ObjectType: to.Ptr("SecretStoreBasedAuthCredentials"),
				SecretStoreResource: &armdataprotection.SecretStoreResource{
					SecretStoreType: to.Ptr(armdataprotection.SecretStoreTypeAzureKeyVault),
					URI:             to.Ptr("https://samplevault.vault.azure.net/secrets/credentials"),
				},
			},
			FriendlyName: to.Ptr("harshitbi2"),
			ObjectType:   to.Ptr("BackupInstance"),
			PolicyInfo: &armdataprotection.PolicyInfo{
				PolicyID: to.Ptr("/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/000pikumar/providers/Microsoft.DataProtection/Backupvaults/PratikPrivatePreviewVault1/backupPolicies/PratikPolicy1"),
				PolicyParameters: &armdataprotection.PolicyParameters{
					DataStoreParametersList: []armdataprotection.DataStoreParametersClassification{
						&armdataprotection.AzureOperationalStoreParameters{
							DataStoreType:   to.Ptr(armdataprotection.DataStoreTypesOperationalStore),
							ObjectType:      to.Ptr("AzureOperationalStoreParameters"),
							ResourceGroupID: to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest"),
						}},
				},
			},
			ValidationType: to.Ptr(armdataprotection.ValidationTypeShallowValidation),
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/BackupInstanceOperations/DeleteBackupInstance.json
func ExampleBackupInstancesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("04cf684a-d41f-4550-9f70-7708a3a2283b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx, "000pikumar", "PratikPrivatePreviewVault1", "testInstance1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/BackupInstanceOperations/TriggerBackup.json
func ExampleBackupInstancesClient_BeginAdhocBackup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("04cf684a-d41f-4550-9f70-7708a3a2283b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginAdhocBackup(ctx, "000pikumar", "PratikPrivatePreviewVault1", "testInstance1", armdataprotection.TriggerBackupRequest{
		BackupRuleOptions: &armdataprotection.AdHocBackupRuleOptions{
			RuleName: to.Ptr("BackupWeekly"),
			TriggerOption: &armdataprotection.AdhocBackupTriggerOption{
				RetentionTagOverride: to.Ptr("yearly"),
			},
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/BackupInstanceOperations/ValidateForBackup.json
func ExampleBackupInstancesClient_BeginValidateForBackup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("04cf684a-d41f-4550-9f70-7708a3a2283b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginValidateForBackup(ctx, "000pikumar", "PratikPrivatePreviewVault1", armdataprotection.ValidateForBackupRequest{
		BackupInstance: &armdataprotection.BackupInstance{
			DataSourceInfo: &armdataprotection.Datasource{
				DatasourceType:   to.Ptr("OssDB"),
				ObjectType:       to.Ptr("Datasource"),
				ResourceID:       to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest/databases/testdb"),
				ResourceLocation: to.Ptr(""),
				ResourceName:     to.Ptr("testdb"),
				ResourceType:     to.Ptr("Microsoft.DBforPostgreSQL/servers/databases"),
				ResourceURI:      to.Ptr(""),
			},
			DataSourceSetInfo: &armdataprotection.DatasourceSet{
				DatasourceType:   to.Ptr("OssDB"),
				ObjectType:       to.Ptr("DatasourceSet"),
				ResourceID:       to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest"),
				ResourceLocation: to.Ptr(""),
				ResourceName:     to.Ptr("viveksipgtest"),
				ResourceType:     to.Ptr("Microsoft.DBforPostgreSQL/servers"),
				ResourceURI:      to.Ptr(""),
			},
			DatasourceAuthCredentials: &armdataprotection.SecretStoreBasedAuthCredentials{
				ObjectType: to.Ptr("SecretStoreBasedAuthCredentials"),
				SecretStoreResource: &armdataprotection.SecretStoreResource{
					SecretStoreType: to.Ptr(armdataprotection.SecretStoreTypeAzureKeyVault),
					URI:             to.Ptr("https://samplevault.vault.azure.net/secrets/credentials"),
				},
			},
			FriendlyName: to.Ptr("harshitbi2"),
			ObjectType:   to.Ptr("BackupInstance"),
			PolicyInfo: &armdataprotection.PolicyInfo{
				PolicyID: to.Ptr("/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/000pikumar/providers/Microsoft.DataProtection/Backupvaults/PratikPrivatePreviewVault1/backupPolicies/PratikPolicy1"),
			},
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/BackupInstanceOperations/GetBackupInstanceOperationResult.json
func ExampleBackupInstancesClient_GetBackupInstanceOperationResult() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("04cf684a-d41f-4550-9f70-7708a3a2283b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetBackupInstanceOperationResult(ctx, "SampleResourceGroup", "swaggerExample", "testInstance1", "YWUzNDFkMzQtZmM5OS00MmUyLWEzNDMtZGJkMDIxZjlmZjgzOzdmYzBiMzhmLTc2NmItNDM5NS05OWQ1LTVmOGEzNzg4MWQzNA==", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/BackupInstanceOperations/TriggerRehydrate.json
func ExampleBackupInstancesClient_BeginTriggerRehydrate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("04cf684a-d41f-4550-9f70-7708a3a2283b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginTriggerRehydrate(ctx, "000pikumar", "PratikPrivatePreviewVault1", "testInstance1", armdataprotection.AzureBackupRehydrationRequest{
		RecoveryPointID:              to.Ptr("hardcodedRP"),
		RehydrationPriority:          to.Ptr(armdataprotection.RehydrationPriorityHigh),
		RehydrationRetentionDuration: to.Ptr("7D"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/BackupInstanceOperations/TriggerRestore.json
func ExampleBackupInstancesClient_BeginTriggerRestore_triggerRestore() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("04cf684a-d41f-4550-9f70-7708a3a2283b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginTriggerRestore(ctx, "000pikumar", "PratikPrivatePreviewVault1", "testInstance1", &armdataprotection.AzureBackupRecoveryPointBasedRestoreRequest{
		ObjectType: to.Ptr("AzureBackupRecoveryPointBasedRestoreRequest"),
		RestoreTargetInfo: &armdataprotection.RestoreTargetInfo{
			ObjectType:      to.Ptr("RestoreTargetInfo"),
			RecoveryOption:  to.Ptr(armdataprotection.RecoveryOptionFailIfExists),
			RestoreLocation: to.Ptr("southeastasia"),
			DatasourceAuthCredentials: &armdataprotection.SecretStoreBasedAuthCredentials{
				ObjectType: to.Ptr("SecretStoreBasedAuthCredentials"),
				SecretStoreResource: &armdataprotection.SecretStoreResource{
					SecretStoreType: to.Ptr(armdataprotection.SecretStoreTypeAzureKeyVault),
					URI:             to.Ptr("https://samplevault.vault.azure.net/secrets/credentials"),
				},
			},
			DatasourceInfo: &armdataprotection.Datasource{
				DatasourceType:   to.Ptr("Microsoft.DBforPostgreSQL/servers/databases"),
				ObjectType:       to.Ptr("Datasource"),
				ResourceID:       to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest/databases/targetdb"),
				ResourceLocation: to.Ptr(""),
				ResourceName:     to.Ptr("targetdb"),
				ResourceType:     to.Ptr("Microsoft.DBforPostgreSQL/servers/databases"),
				ResourceURI:      to.Ptr(""),
			},
			DatasourceSetInfo: &armdataprotection.DatasourceSet{
				DatasourceType:   to.Ptr("Microsoft.DBforPostgreSQL/servers/databases"),
				ObjectType:       to.Ptr("DatasourceSet"),
				ResourceID:       to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest"),
				ResourceLocation: to.Ptr(""),
				ResourceName:     to.Ptr("viveksipgtest"),
				ResourceType:     to.Ptr("Microsoft.DBforPostgreSQL/servers"),
				ResourceURI:      to.Ptr(""),
			},
		},
		SourceDataStoreType: to.Ptr(armdataprotection.SourceDataStoreTypeVaultStore),
		SourceResourceID:    to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest/databases/testdb"),
		RecoveryPointID:     to.Ptr("hardcodedRP"),
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/BackupInstanceOperations/TriggerRestoreAsFiles.json
func ExampleBackupInstancesClient_BeginTriggerRestore_triggerRestoreAsFiles() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("04cf684a-d41f-4550-9f70-7708a3a2283b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginTriggerRestore(ctx, "000pikumar", "PrivatePreviewVault1", "testInstance1", &armdataprotection.AzureBackupRecoveryPointBasedRestoreRequest{
		ObjectType: to.Ptr("AzureBackupRecoveryPointBasedRestoreRequest"),
		RestoreTargetInfo: &armdataprotection.RestoreFilesTargetInfo{
			ObjectType:      to.Ptr("RestoreFilesTargetInfo"),
			RecoveryOption:  to.Ptr(armdataprotection.RecoveryOptionFailIfExists),
			RestoreLocation: to.Ptr("southeastasia"),
			TargetDetails: &armdataprotection.TargetDetails{
				FilePrefix:                to.Ptr("restoredblob"),
				RestoreTargetLocationType: to.Ptr(armdataprotection.RestoreTargetLocationTypeAzureBlobs),
				URL:                       to.Ptr("https://teststorage.blob.core.windows.net/restoretest"),
			},
		},
		SourceDataStoreType: to.Ptr(armdataprotection.SourceDataStoreTypeVaultStore),
		SourceResourceID:    to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest/databases/testdb"),
		RecoveryPointID:     to.Ptr("hardcodedRP"),
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/BackupInstanceOperations/TriggerRestoreWithRehydration.json
func ExampleBackupInstancesClient_BeginTriggerRestore_triggerRestoreWithRehydration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("04cf684a-d41f-4550-9f70-7708a3a2283b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginTriggerRestore(ctx, "000pikumar", "PratikPrivatePreviewVault1", "testInstance1", &armdataprotection.AzureBackupRestoreWithRehydrationRequest{
		ObjectType: to.Ptr("AzureBackupRestoreWithRehydrationRequest"),
		RestoreTargetInfo: &armdataprotection.RestoreTargetInfo{
			ObjectType:      to.Ptr("RestoreTargetInfo"),
			RecoveryOption:  to.Ptr(armdataprotection.RecoveryOptionFailIfExists),
			RestoreLocation: to.Ptr("southeastasia"),
			DatasourceInfo: &armdataprotection.Datasource{
				DatasourceType:   to.Ptr("OssDB"),
				ObjectType:       to.Ptr("Datasource"),
				ResourceID:       to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest/databases/testdb"),
				ResourceLocation: to.Ptr(""),
				ResourceName:     to.Ptr("testdb"),
				ResourceType:     to.Ptr("Microsoft.DBforPostgreSQL/servers/databases"),
				ResourceURI:      to.Ptr(""),
			},
			DatasourceSetInfo: &armdataprotection.DatasourceSet{
				DatasourceType:   to.Ptr("OssDB"),
				ObjectType:       to.Ptr("DatasourceSet"),
				ResourceID:       to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest"),
				ResourceLocation: to.Ptr(""),
				ResourceName:     to.Ptr("viveksipgtest"),
				ResourceType:     to.Ptr("Microsoft.DBforPostgreSQL/servers"),
				ResourceURI:      to.Ptr(""),
			},
		},
		SourceDataStoreType:          to.Ptr(armdataprotection.SourceDataStoreTypeVaultStore),
		SourceResourceID:             to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest/databases/testdb"),
		RecoveryPointID:              to.Ptr("hardcodedRP"),
		RehydrationPriority:          to.Ptr(armdataprotection.RehydrationPriorityHigh),
		RehydrationRetentionDuration: to.Ptr("7D"),
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/BackupInstanceOperations/ResumeBackups.json
func ExampleBackupInstancesClient_BeginResumeBackups() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("04cf684a-d41f-4550-9f70-7708a3a2283b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginResumeBackups(ctx, "testrg", "testvault", "testbi", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/BackupInstanceOperations/ResumeProtection.json
func ExampleBackupInstancesClient_BeginResumeProtection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("04cf684a-d41f-4550-9f70-7708a3a2283b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginResumeProtection(ctx, "testrg", "testvault", "testbi", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/BackupInstanceOperations/StopProtection.json
func ExampleBackupInstancesClient_BeginStopProtection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("04cf684a-d41f-4550-9f70-7708a3a2283b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginStopProtection(ctx, "testrg", "testvault", "testbi", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/BackupInstanceOperations/SuspendBackups.json
func ExampleBackupInstancesClient_BeginSuspendBackups() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("04cf684a-d41f-4550-9f70-7708a3a2283b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginSuspendBackups(ctx, "testrg", "testvault", "testbi", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/BackupInstanceOperations/SyncBackupInstance.json
func ExampleBackupInstancesClient_BeginSyncBackupInstance() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("04cf684a-d41f-4550-9f70-7708a3a2283b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginSyncBackupInstance(ctx, "testrg", "testvault", "testbi", armdataprotection.SyncBackupInstanceRequest{
		SyncType: to.Ptr(armdataprotection.SyncTypeDefault),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/BackupInstanceOperations/ValidateRestore.json
func ExampleBackupInstancesClient_BeginValidateForRestore() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("04cf684a-d41f-4550-9f70-7708a3a2283b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginValidateForRestore(ctx, "000pikumar", "PratikPrivatePreviewVault1", "testInstance1", armdataprotection.ValidateRestoreRequestObject{
		RestoreRequestObject: &armdataprotection.AzureBackupRecoveryPointBasedRestoreRequest{
			ObjectType: to.Ptr("AzureBackupRecoveryPointBasedRestoreRequest"),
			RestoreTargetInfo: &armdataprotection.RestoreTargetInfo{
				ObjectType:      to.Ptr("RestoreTargetInfo"),
				RecoveryOption:  to.Ptr(armdataprotection.RecoveryOptionFailIfExists),
				RestoreLocation: to.Ptr("southeastasia"),
				DatasourceAuthCredentials: &armdataprotection.SecretStoreBasedAuthCredentials{
					ObjectType: to.Ptr("SecretStoreBasedAuthCredentials"),
					SecretStoreResource: &armdataprotection.SecretStoreResource{
						SecretStoreType: to.Ptr(armdataprotection.SecretStoreTypeAzureKeyVault),
						URI:             to.Ptr("https://samplevault.vault.azure.net/secrets/credentials"),
					},
				},
				DatasourceInfo: &armdataprotection.Datasource{
					DatasourceType:   to.Ptr("Microsoft.DBforPostgreSQL/servers/databases"),
					ObjectType:       to.Ptr("Datasource"),
					ResourceID:       to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest/databases/targetdb"),
					ResourceLocation: to.Ptr(""),
					ResourceName:     to.Ptr("targetdb"),
					ResourceType:     to.Ptr("Microsoft.DBforPostgreSQL/servers/databases"),
					ResourceURI:      to.Ptr(""),
				},
				DatasourceSetInfo: &armdataprotection.DatasourceSet{
					DatasourceType:   to.Ptr("Microsoft.DBforPostgreSQL/servers/databases"),
					ObjectType:       to.Ptr("DatasourceSet"),
					ResourceID:       to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest"),
					ResourceLocation: to.Ptr(""),
					ResourceName:     to.Ptr("viveksipgtest"),
					ResourceType:     to.Ptr("Microsoft.DBforPostgreSQL/servers"),
					ResourceURI:      to.Ptr(""),
				},
			},
			SourceDataStoreType: to.Ptr(armdataprotection.SourceDataStoreTypeVaultStore),
			SourceResourceID:    to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest/databases/testdb"),
			RecoveryPointID:     to.Ptr("hardcodedRP"),
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
