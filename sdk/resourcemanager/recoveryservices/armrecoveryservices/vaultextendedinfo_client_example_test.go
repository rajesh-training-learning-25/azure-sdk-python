//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrecoveryservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2023-01-01/examples/GETVaultExtendedInfo.json
func ExampleVaultExtendedInfoClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservices.NewVaultExtendedInfoClient("77777777-b0c6-47a2-b37c-d8e65a629c18", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "Default-RecoveryServices-ResourceGroup", "swaggerExample", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VaultExtendedInfoResource = armrecoveryservices.VaultExtendedInfoResource{
	// 	Name: to.Ptr("vaultExtendedInfo"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/extendedInformation"),
	// 	Etag: to.Ptr("f0d0260b-b92d-4458-ba0a-32c6cdabacb7"),
	// 	ID: to.Ptr("/subscriptions/77777777-b0c6-47a2-b37c-d8e65a629c18/resourceGroups/Default-RecoveryServices-ResourceGroup/providers/Microsoft.RecoveryServices/vaults/swaggerExample/extendedInformation/vaultExtendedInfo"),
	// 	Properties: &armrecoveryservices.VaultExtendedInfo{
	// 		Algorithm: to.Ptr("None"),
	// 		IntegrityKey: to.Ptr("J09wzS27fnJ+Wjot7xO5wA=="),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2023-01-01/examples/UpdateVaultExtendedInfo.json
func ExampleVaultExtendedInfoClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservices.NewVaultExtendedInfoClient("77777777-b0c6-47a2-b37c-d8e65a629c18", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "Default-RecoveryServices-ResourceGroup", "swaggerExample", armrecoveryservices.VaultExtendedInfoResource{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VaultExtendedInfoResource = armrecoveryservices.VaultExtendedInfoResource{
	// 	Name: to.Ptr("vaultExtendedInfo"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/extendedInformation"),
	// 	Etag: to.Ptr("f0d0260b-b92d-4458-ba0a-32c6cdabacb7"),
	// 	ID: to.Ptr("/subscriptions/77777777-b0c6-47a2-b37c-d8e65a629c18/resourceGroups/Default-RecoveryServices-ResourceGroup/providers/Microsoft.RecoveryServices/vaults/swaggerExample/extendedInformation/vaultExtendedInfo"),
	// 	Properties: &armrecoveryservices.VaultExtendedInfo{
	// 		Algorithm: to.Ptr("None"),
	// 		IntegrityKey: to.Ptr("J99wzS27fmJ+Wjot7xO5wA=="),
	// 	},
	// }
}
