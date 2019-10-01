package backupapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/recoveryservices/mgmt/2016-12-01/backup"
	"github.com/Azure/go-autorest/autorest"
)

// ResourceVaultConfigsClientAPI contains the set of methods on the ResourceVaultConfigsClient type.
type ResourceVaultConfigsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string) (result backup.ResourceVaultConfigResource, err error)
	Update(ctx context.Context, vaultName string, resourceGroupName string, parameters backup.ResourceVaultConfigResource) (result backup.ResourceVaultConfigResource, err error)
}

var _ ResourceVaultConfigsClientAPI = (*backup.ResourceVaultConfigsClient)(nil)

// EnginesClientAPI contains the set of methods on the EnginesClient type.
type EnginesClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, backupEngineName string, filter string, skipToken string) (result backup.EngineBaseResource, err error)
	List(ctx context.Context, vaultName string, resourceGroupName string, filter string, skipToken string) (result backup.EngineBaseResourceListPage, err error)
}

var _ EnginesClientAPI = (*backup.EnginesClient)(nil)

// ProtectionContainerRefreshOperationResultsClientAPI contains the set of methods on the ProtectionContainerRefreshOperationResultsClient type.
type ProtectionContainerRefreshOperationResultsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, operationID string) (result autorest.Response, err error)
}

var _ ProtectionContainerRefreshOperationResultsClientAPI = (*backup.ProtectionContainerRefreshOperationResultsClient)(nil)

// ProtectableContainersClientAPI contains the set of methods on the ProtectableContainersClient type.
type ProtectableContainersClientAPI interface {
	List(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, filter string) (result backup.ProtectableContainerResourceListPage, err error)
}

var _ ProtectableContainersClientAPI = (*backup.ProtectableContainersClient)(nil)

// ProtectionContainersClientAPI contains the set of methods on the ProtectionContainersClient type.
type ProtectionContainersClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string) (result backup.ProtectionContainerResource, err error)
	Inquire(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, filter string) (result autorest.Response, err error)
	Refresh(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, filter string) (result autorest.Response, err error)
	Register(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, parameters backup.ProtectionContainerResource) (result backup.ProtectionContainerResource, err error)
	Unregister(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string) (result autorest.Response, err error)
}

var _ ProtectionContainersClientAPI = (*backup.ProtectionContainersClient)(nil)

// WorkloadItemsClientAPI contains the set of methods on the WorkloadItemsClient type.
type WorkloadItemsClientAPI interface {
	List(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, filter string, skipToken string) (result backup.WorkloadItemResourceListPage, err error)
}

var _ WorkloadItemsClientAPI = (*backup.WorkloadItemsClient)(nil)

// ProtectionContainerOperationResultsClientAPI contains the set of methods on the ProtectionContainerOperationResultsClient type.
type ProtectionContainerOperationResultsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, operationID string) (result backup.ProtectionContainerResource, err error)
}

var _ ProtectionContainerOperationResultsClientAPI = (*backup.ProtectionContainerOperationResultsClient)(nil)

// ProtectedItemsClientAPI contains the set of methods on the ProtectedItemsClient type.
type ProtectedItemsClientAPI interface {
	Delete(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string) (result autorest.Response, err error)
}

var _ ProtectedItemsClientAPI = (*backup.ProtectedItemsClient)(nil)

// BackupsClientAPI contains the set of methods on the BackupsClient type.
type BackupsClientAPI interface {
	Trigger(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, parameters backup.RequestResource) (result autorest.Response, err error)
}

var _ BackupsClientAPI = (*backup.BackupsClient)(nil)

// ProtectedItemOperationResultsClientAPI contains the set of methods on the ProtectedItemOperationResultsClient type.
type ProtectedItemOperationResultsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, operationID string) (result backup.ProtectedItemResource, err error)
}

var _ ProtectedItemOperationResultsClientAPI = (*backup.ProtectedItemOperationResultsClient)(nil)

// ProtectedItemOperationStatusesClientAPI contains the set of methods on the ProtectedItemOperationStatusesClient type.
type ProtectedItemOperationStatusesClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, operationID string) (result backup.OperationStatus, err error)
}

var _ ProtectedItemOperationStatusesClientAPI = (*backup.ProtectedItemOperationStatusesClient)(nil)

// ItemLevelRecoveryConnectionsClientAPI contains the set of methods on the ItemLevelRecoveryConnectionsClient type.
type ItemLevelRecoveryConnectionsClientAPI interface {
	Provision(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string, parameters backup.ILRRequestResource) (result autorest.Response, err error)
	Revoke(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string) (result autorest.Response, err error)
}

var _ ItemLevelRecoveryConnectionsClientAPI = (*backup.ItemLevelRecoveryConnectionsClient)(nil)

// JobCancellationsClientAPI contains the set of methods on the JobCancellationsClient type.
type JobCancellationsClientAPI interface {
	Trigger(ctx context.Context, vaultName string, resourceGroupName string, jobName string) (result autorest.Response, err error)
}

var _ JobCancellationsClientAPI = (*backup.JobCancellationsClient)(nil)

// JobOperationResultsClientAPI contains the set of methods on the JobOperationResultsClient type.
type JobOperationResultsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, jobName string, operationID string) (result autorest.Response, err error)
}

var _ JobOperationResultsClientAPI = (*backup.JobOperationResultsClient)(nil)

// OperationResultsClientAPI contains the set of methods on the OperationResultsClient type.
type OperationResultsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, operationID string) (result autorest.Response, err error)
}

var _ OperationResultsClientAPI = (*backup.OperationResultsClient)(nil)

// OperationStatusesClientAPI contains the set of methods on the OperationStatusesClient type.
type OperationStatusesClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, operationID string) (result backup.OperationStatus, err error)
}

var _ OperationStatusesClientAPI = (*backup.OperationStatusesClient)(nil)

// ProtectionPoliciesClientAPI contains the set of methods on the ProtectionPoliciesClient type.
type ProtectionPoliciesClientAPI interface {
	Delete(ctx context.Context, vaultName string, resourceGroupName string, policyName string) (result autorest.Response, err error)
	Get(ctx context.Context, vaultName string, resourceGroupName string, policyName string) (result backup.ProtectionPolicyResource, err error)
}

var _ ProtectionPoliciesClientAPI = (*backup.ProtectionPoliciesClient)(nil)

// ProtectionPolicyOperationResultsClientAPI contains the set of methods on the ProtectionPolicyOperationResultsClient type.
type ProtectionPolicyOperationResultsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, policyName string, operationID string) (result backup.ProtectionPolicyResource, err error)
}

var _ ProtectionPolicyOperationResultsClientAPI = (*backup.ProtectionPolicyOperationResultsClient)(nil)

// ProtectionPolicyOperationStatusesClientAPI contains the set of methods on the ProtectionPolicyOperationStatusesClient type.
type ProtectionPolicyOperationStatusesClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, policyName string, operationID string) (result backup.OperationStatus, err error)
}

var _ ProtectionPolicyOperationStatusesClientAPI = (*backup.ProtectionPolicyOperationStatusesClient)(nil)

// ProtectableItemsClientAPI contains the set of methods on the ProtectableItemsClient type.
type ProtectableItemsClientAPI interface {
	List(ctx context.Context, vaultName string, resourceGroupName string, filter string, skipToken string) (result backup.WorkloadProtectableItemResourceListPage, err error)
}

var _ ProtectableItemsClientAPI = (*backup.ProtectableItemsClient)(nil)

// ProtectionContainersGroupClientAPI contains the set of methods on the ProtectionContainersGroupClient type.
type ProtectionContainersGroupClientAPI interface {
	List(ctx context.Context, vaultName string, resourceGroupName string, filter string) (result backup.ProtectionContainerResourceListPage, err error)
}

var _ ProtectionContainersGroupClientAPI = (*backup.ProtectionContainersGroupClient)(nil)

// SecurityPINsClientAPI contains the set of methods on the SecurityPINsClient type.
type SecurityPINsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string) (result backup.TokenInformation, err error)
}

var _ SecurityPINsClientAPI = (*backup.SecurityPINsClient)(nil)

// ResourceStorageConfigsClientAPI contains the set of methods on the ResourceStorageConfigsClient type.
type ResourceStorageConfigsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string) (result backup.ResourceConfigResource, err error)
	Patch(ctx context.Context, vaultName string, resourceGroupName string, parameters backup.ResourceConfigResource) (result autorest.Response, err error)
	Update(ctx context.Context, vaultName string, resourceGroupName string, parameters backup.ResourceConfigResource) (result backup.ResourceConfigResource, err error)
}

var _ ResourceStorageConfigsClientAPI = (*backup.ResourceStorageConfigsClient)(nil)
