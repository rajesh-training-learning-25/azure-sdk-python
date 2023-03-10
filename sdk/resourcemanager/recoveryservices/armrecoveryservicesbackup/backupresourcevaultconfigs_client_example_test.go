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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a498cae6d1a93f4c33073f0747b93b22815c09b7/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-02-01/examples/Common/BackupResourceVaultConfigs_Get.json
func ExampleBackupResourceVaultConfigsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicesbackup.NewBackupResourceVaultConfigsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "SwaggerTest", "SwaggerTestRg", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BackupResourceVaultConfigResource = armrecoveryservicesbackup.BackupResourceVaultConfigResource{
	// 	Name: to.Ptr("vaultconfig"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupconfig"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/SwaggerTestRg/providers/Microsoft.RecoveryServices/vaults/SwaggerTest/backupconfig/vaultconfig"),
	// 	Properties: &armrecoveryservicesbackup.BackupResourceVaultConfig{
	// 		EnhancedSecurityState: to.Ptr(armrecoveryservicesbackup.EnhancedSecurityStateEnabled),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a498cae6d1a93f4c33073f0747b93b22815c09b7/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-02-01/examples/Common/BackupResourceVaultConfigs_Patch.json
func ExampleBackupResourceVaultConfigsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicesbackup.NewBackupResourceVaultConfigsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "SwaggerTest", "SwaggerTestRg", armrecoveryservicesbackup.BackupResourceVaultConfigResource{
		Properties: &armrecoveryservicesbackup.BackupResourceVaultConfig{
			EnhancedSecurityState: to.Ptr(armrecoveryservicesbackup.EnhancedSecurityStateEnabled),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BackupResourceVaultConfigResource = armrecoveryservicesbackup.BackupResourceVaultConfigResource{
	// 	Name: to.Ptr("vaultconfig"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupconfig"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/SwaggerTestRg/providers/Microsoft.RecoveryServices/vaults/SwaggerTest/backupconfig/vaultconfig"),
	// 	Properties: &armrecoveryservicesbackup.BackupResourceVaultConfig{
	// 		EnhancedSecurityState: to.Ptr(armrecoveryservicesbackup.EnhancedSecurityStateEnabled),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a498cae6d1a93f4c33073f0747b93b22815c09b7/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-02-01/examples/Common/BackupResourceVaultConfigs_Put.json
func ExampleBackupResourceVaultConfigsClient_Put() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicesbackup.NewBackupResourceVaultConfigsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Put(ctx, "SwaggerTest", "SwaggerTestRg", armrecoveryservicesbackup.BackupResourceVaultConfigResource{
		Properties: &armrecoveryservicesbackup.BackupResourceVaultConfig{
			EnhancedSecurityState:  to.Ptr(armrecoveryservicesbackup.EnhancedSecurityStateEnabled),
			SoftDeleteFeatureState: to.Ptr(armrecoveryservicesbackup.SoftDeleteFeatureStateEnabled),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BackupResourceVaultConfigResource = armrecoveryservicesbackup.BackupResourceVaultConfigResource{
	// 	Name: to.Ptr("vaultconfig"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupconfig"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/SwaggerTestRg/providers/Microsoft.RecoveryServices/vaults/SwaggerTest/backupconfig/vaultconfig"),
	// 	Properties: &armrecoveryservicesbackup.BackupResourceVaultConfig{
	// 		EnhancedSecurityState: to.Ptr(armrecoveryservicesbackup.EnhancedSecurityStateEnabled),
	// 		SoftDeleteFeatureState: to.Ptr(armrecoveryservicesbackup.SoftDeleteFeatureStateEnabled),
	// 	},
	// }
}
