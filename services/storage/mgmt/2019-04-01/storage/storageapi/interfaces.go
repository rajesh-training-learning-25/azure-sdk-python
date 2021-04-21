package storageapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result storage.OperationListResult, err error)
}

var _ OperationsClientAPI = (*storage.OperationsClient)(nil)

// SkusClientAPI contains the set of methods on the SkusClient type.
type SkusClientAPI interface {
	List(ctx context.Context) (result storage.SkuListResult, err error)
}

var _ SkusClientAPI = (*storage.SkusClient)(nil)

// AccountsClientAPI contains the set of methods on the AccountsClient type.
type AccountsClientAPI interface {
	CheckNameAvailability(ctx context.Context, accountName storage.AccountCheckNameAvailabilityParameters) (result storage.CheckNameAvailabilityResult, err error)
	Create(ctx context.Context, resourceGroupName string, accountName string, parameters storage.AccountCreateParameters) (result storage.AccountsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string) (result autorest.Response, err error)
	Failover(ctx context.Context, resourceGroupName string, accountName string) (result storage.AccountsFailoverFuture, err error)
	GetProperties(ctx context.Context, resourceGroupName string, accountName string, expand storage.AccountExpand) (result storage.Account, err error)
	List(ctx context.Context) (result storage.AccountListResultPage, err error)
	ListComplete(ctx context.Context) (result storage.AccountListResultIterator, err error)
	ListAccountSAS(ctx context.Context, resourceGroupName string, accountName string, parameters storage.AccountSasParameters) (result storage.ListAccountSasResponse, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result storage.AccountListResult, err error)
	ListKeys(ctx context.Context, resourceGroupName string, accountName string, expand storage.ListKeyExpand) (result storage.AccountListKeysResult, err error)
	ListServiceSAS(ctx context.Context, resourceGroupName string, accountName string, parameters storage.ServiceSasParameters) (result storage.ListServiceSasResponse, err error)
	RegenerateKey(ctx context.Context, resourceGroupName string, accountName string, regenerateKey storage.AccountRegenerateKeyParameters) (result storage.AccountListKeysResult, err error)
	RevokeUserDelegationKeys(ctx context.Context, resourceGroupName string, accountName string) (result autorest.Response, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, parameters storage.AccountUpdateParameters) (result storage.Account, err error)
}

var _ AccountsClientAPI = (*storage.AccountsClient)(nil)

// UsagesClientAPI contains the set of methods on the UsagesClient type.
type UsagesClientAPI interface {
	ListByLocation(ctx context.Context, location string) (result storage.UsageListResult, err error)
}

var _ UsagesClientAPI = (*storage.UsagesClient)(nil)

// ManagementPoliciesClientAPI contains the set of methods on the ManagementPoliciesClient type.
type ManagementPoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, properties storage.ManagementPolicy) (result storage.ManagementPolicy, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string) (result storage.ManagementPolicy, err error)
}

var _ ManagementPoliciesClientAPI = (*storage.ManagementPoliciesClient)(nil)

// BlobServicesClientAPI contains the set of methods on the BlobServicesClient type.
type BlobServicesClientAPI interface {
	GetServiceProperties(ctx context.Context, resourceGroupName string, accountName string) (result storage.BlobServiceProperties, err error)
	List(ctx context.Context, resourceGroupName string, accountName string) (result storage.BlobServiceItems, err error)
	SetServiceProperties(ctx context.Context, resourceGroupName string, accountName string, parameters storage.BlobServiceProperties) (result storage.BlobServiceProperties, err error)
}

var _ BlobServicesClientAPI = (*storage.BlobServicesClient)(nil)

// BlobContainersClientAPI contains the set of methods on the BlobContainersClient type.
type BlobContainersClientAPI interface {
	ClearLegalHold(ctx context.Context, resourceGroupName string, accountName string, containerName string, legalHold storage.LegalHold) (result storage.LegalHold, err error)
	Create(ctx context.Context, resourceGroupName string, accountName string, containerName string, blobContainer storage.BlobContainer) (result storage.BlobContainer, err error)
	CreateOrUpdateImmutabilityPolicy(ctx context.Context, resourceGroupName string, accountName string, containerName string, parameters *storage.ImmutabilityPolicy, ifMatch string) (result storage.ImmutabilityPolicy, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, containerName string) (result autorest.Response, err error)
	DeleteImmutabilityPolicy(ctx context.Context, resourceGroupName string, accountName string, containerName string, ifMatch string) (result storage.ImmutabilityPolicy, err error)
	ExtendImmutabilityPolicy(ctx context.Context, resourceGroupName string, accountName string, containerName string, ifMatch string, parameters *storage.ImmutabilityPolicy) (result storage.ImmutabilityPolicy, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, containerName string) (result storage.BlobContainer, err error)
	GetImmutabilityPolicy(ctx context.Context, resourceGroupName string, accountName string, containerName string, ifMatch string) (result storage.ImmutabilityPolicy, err error)
	Lease(ctx context.Context, resourceGroupName string, accountName string, containerName string, parameters *storage.LeaseContainerRequest) (result storage.LeaseContainerResponse, err error)
	List(ctx context.Context, resourceGroupName string, accountName string, skipToken string, maxpagesize string, filter string) (result storage.ListContainerItemsPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, accountName string, skipToken string, maxpagesize string, filter string) (result storage.ListContainerItemsIterator, err error)
	LockImmutabilityPolicy(ctx context.Context, resourceGroupName string, accountName string, containerName string, ifMatch string) (result storage.ImmutabilityPolicy, err error)
	SetLegalHold(ctx context.Context, resourceGroupName string, accountName string, containerName string, legalHold storage.LegalHold) (result storage.LegalHold, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, containerName string, blobContainer storage.BlobContainer) (result storage.BlobContainer, err error)
}

var _ BlobContainersClientAPI = (*storage.BlobContainersClient)(nil)

// FileServicesClientAPI contains the set of methods on the FileServicesClient type.
type FileServicesClientAPI interface {
	GetServiceProperties(ctx context.Context, resourceGroupName string, accountName string) (result storage.FileServiceProperties, err error)
	List(ctx context.Context, resourceGroupName string, accountName string) (result storage.FileServiceItems, err error)
	SetServiceProperties(ctx context.Context, resourceGroupName string, accountName string, parameters storage.FileServiceProperties) (result storage.FileServiceProperties, err error)
}

var _ FileServicesClientAPI = (*storage.FileServicesClient)(nil)

// FileSharesClientAPI contains the set of methods on the FileSharesClient type.
type FileSharesClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, accountName string, shareName string, fileShare storage.FileShare) (result storage.FileShare, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, shareName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, shareName string) (result storage.FileShare, err error)
	List(ctx context.Context, resourceGroupName string, accountName string, skipToken string, maxpagesize string, filter string) (result storage.FileShareItemsPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, accountName string, skipToken string, maxpagesize string, filter string) (result storage.FileShareItemsIterator, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, shareName string, fileShare storage.FileShare) (result storage.FileShare, err error)
}

var _ FileSharesClientAPI = (*storage.FileSharesClient)(nil)
