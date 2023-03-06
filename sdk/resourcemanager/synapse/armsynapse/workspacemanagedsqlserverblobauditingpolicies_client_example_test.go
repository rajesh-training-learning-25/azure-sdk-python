//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetWorkspaceManagedSqlServerBlobAuditingSettings.json
func ExampleWorkspaceManagedSQLServerBlobAuditingPoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewWorkspaceManagedSQLServerBlobAuditingPoliciesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "wsg-7398", "testWorkspace", armsynapse.BlobAuditingPolicyNameDefault, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServerBlobAuditingPolicy = armsynapse.ServerBlobAuditingPolicy{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces/auditingSettings"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/wsg-7398/providers/Microsoft.Synapse/workspaces/testWorkspace/auditingSettings/default"),
	// 	Properties: &armsynapse.ServerBlobAuditingPolicyProperties{
	// 		AuditActionsAndGroups: []*string{
	// 		},
	// 		IsAzureMonitorTargetEnabled: to.Ptr(false),
	// 		IsDevopsAuditEnabled: to.Ptr(false),
	// 		IsStorageSecondaryKeyInUse: to.Ptr(false),
	// 		RetentionDays: to.Ptr[int32](0),
	// 		State: to.Ptr(armsynapse.BlobAuditingPolicyStateDisabled),
	// 		StorageAccountSubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		StorageEndpoint: to.Ptr(""),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateWorkspaceManagedSqlServerBlobAuditingSettingsWithAllParameters.json
func ExampleWorkspaceManagedSQLServerBlobAuditingPoliciesClient_BeginCreateOrUpdate_createOrUpdateBlobAuditingPolicyOfWorkspaceSqlServerWithAllParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewWorkspaceManagedSQLServerBlobAuditingPoliciesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "wsg-7398", "testWorkspace", armsynapse.BlobAuditingPolicyNameDefault, armsynapse.ServerBlobAuditingPolicy{
		Properties: &armsynapse.ServerBlobAuditingPolicyProperties{
			AuditActionsAndGroups: []*string{
				to.Ptr("SUCCESSFUL_DATABASE_AUTHENTICATION_GROUP"),
				to.Ptr("FAILED_DATABASE_AUTHENTICATION_GROUP"),
				to.Ptr("BATCH_COMPLETED_GROUP")},
			IsAzureMonitorTargetEnabled:  to.Ptr(true),
			IsStorageSecondaryKeyInUse:   to.Ptr(false),
			QueueDelayMs:                 to.Ptr[int32](4000),
			RetentionDays:                to.Ptr[int32](6),
			State:                        to.Ptr(armsynapse.BlobAuditingPolicyStateEnabled),
			StorageAccountAccessKey:      to.Ptr("sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD=="),
			StorageAccountSubscriptionID: to.Ptr("00000000-1234-0000-5678-000000000000"),
			StorageEndpoint:              to.Ptr("https://mystorage.blob.core.windows.net"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServerBlobAuditingPolicy = armsynapse.ServerBlobAuditingPolicy{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces/auditingSettings"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/wsg-7398/providers/Microsoft.Synapse/workspaces/testWorkspace/auditingSettings/default"),
	// 	Properties: &armsynapse.ServerBlobAuditingPolicyProperties{
	// 		AuditActionsAndGroups: []*string{
	// 			to.Ptr("SUCCESSFUL_DATABASE_AUTHENTICATION_GROUP"),
	// 			to.Ptr("FAILED_DATABASE_AUTHENTICATION_GROUP"),
	// 			to.Ptr("BATCH_COMPLETED_GROUP")},
	// 			IsAzureMonitorTargetEnabled: to.Ptr(true),
	// 			IsStorageSecondaryKeyInUse: to.Ptr(false),
	// 			QueueDelayMs: to.Ptr[int32](4000),
	// 			RetentionDays: to.Ptr[int32](6),
	// 			State: to.Ptr(armsynapse.BlobAuditingPolicyStateEnabled),
	// 			StorageAccountSubscriptionID: to.Ptr("00000000-1234-0000-5678-000000000000"),
	// 			StorageEndpoint: to.Ptr("https://mystorage.blob.core.windows.net"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateWorkspaceManagedSqlServerBlobAuditingSettingsWithMinParameters.json
func ExampleWorkspaceManagedSQLServerBlobAuditingPoliciesClient_BeginCreateOrUpdate_createOrUpdateBlobAuditingPolicyOfWorkspaceManagedSqlServerWithMinimalParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewWorkspaceManagedSQLServerBlobAuditingPoliciesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "wsg-7398", "testWorkspace", armsynapse.BlobAuditingPolicyNameDefault, armsynapse.ServerBlobAuditingPolicy{
		Properties: &armsynapse.ServerBlobAuditingPolicyProperties{
			State:                   to.Ptr(armsynapse.BlobAuditingPolicyStateEnabled),
			StorageAccountAccessKey: to.Ptr("sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD=="),
			StorageEndpoint:         to.Ptr("https://mystorage.blob.core.windows.net"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServerBlobAuditingPolicy = armsynapse.ServerBlobAuditingPolicy{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces/auditingSettings"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/wsg-7398/providers/Microsoft.Synapse/workspaces/testWorkspace/auditingSettings/default"),
	// 	Properties: &armsynapse.ServerBlobAuditingPolicyProperties{
	// 		AuditActionsAndGroups: []*string{
	// 			to.Ptr("SUCCESSFUL_DATABASE_AUTHENTICATION_GROUP"),
	// 			to.Ptr("FAILED_DATABASE_AUTHENTICATION_GROUP"),
	// 			to.Ptr("BATCH_COMPLETED_GROUP")},
	// 			IsStorageSecondaryKeyInUse: to.Ptr(false),
	// 			RetentionDays: to.Ptr[int32](6),
	// 			State: to.Ptr(armsynapse.BlobAuditingPolicyStateEnabled),
	// 			StorageAccountSubscriptionID: to.Ptr("00000000-1234-0000-5678-000000000000"),
	// 			StorageEndpoint: to.Ptr("https://mystorage.blob.core.windows.net"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListWorkspaceManagedSqlServerBlobAuditingSettings.json
func ExampleWorkspaceManagedSQLServerBlobAuditingPoliciesClient_NewListByWorkspacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewWorkspaceManagedSQLServerBlobAuditingPoliciesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByWorkspacePager("wsg-7398", "testWorkspace", nil)
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
		// page.ServerBlobAuditingPolicyListResult = armsynapse.ServerBlobAuditingPolicyListResult{
		// 	Value: []*armsynapse.ServerBlobAuditingPolicy{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.Synapse/workspaces/auditingSettings"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/wsg-7398/providers/Microsoft.Synapse/workspaces/testWorkspace/auditingSettings/default"),
		// 			Properties: &armsynapse.ServerBlobAuditingPolicyProperties{
		// 				AuditActionsAndGroups: []*string{
		// 				},
		// 				IsAzureMonitorTargetEnabled: to.Ptr(false),
		// 				IsStorageSecondaryKeyInUse: to.Ptr(false),
		// 				RetentionDays: to.Ptr[int32](0),
		// 				State: to.Ptr(armsynapse.BlobAuditingPolicyStateDisabled),
		// 				StorageAccountSubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				StorageEndpoint: to.Ptr(""),
		// 			},
		// 	}},
		// }
	}
}
