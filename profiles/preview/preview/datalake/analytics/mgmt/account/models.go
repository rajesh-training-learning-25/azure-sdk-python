// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package account

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/datalake/analytics/mgmt/2015-10-01-preview/account"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type DataLakeAnalyticsAccountState = original.DataLakeAnalyticsAccountState

const (
	Active    DataLakeAnalyticsAccountState = original.Active
	Suspended DataLakeAnalyticsAccountState = original.Suspended
)

type DataLakeAnalyticsAccountStatus = original.DataLakeAnalyticsAccountStatus

const (
	Creating   DataLakeAnalyticsAccountStatus = original.Creating
	Deleted    DataLakeAnalyticsAccountStatus = original.Deleted
	Deleting   DataLakeAnalyticsAccountStatus = original.Deleting
	Failed     DataLakeAnalyticsAccountStatus = original.Failed
	Patching   DataLakeAnalyticsAccountStatus = original.Patching
	Resuming   DataLakeAnalyticsAccountStatus = original.Resuming
	Running    DataLakeAnalyticsAccountStatus = original.Running
	Succeeded  DataLakeAnalyticsAccountStatus = original.Succeeded
	Suspending DataLakeAnalyticsAccountStatus = original.Suspending
)

type OperationStatus = original.OperationStatus

const (
	OperationStatusFailed     OperationStatus = original.OperationStatusFailed
	OperationStatusInProgress OperationStatus = original.OperationStatusInProgress
	OperationStatusSucceeded  OperationStatus = original.OperationStatusSucceeded
)

type AddDataLakeStoreParameters = original.AddDataLakeStoreParameters
type AddStorageAccountParameters = original.AddStorageAccountParameters
type AzureAsyncOperationResult = original.AzureAsyncOperationResult
type BaseClient = original.BaseClient
type BlobContainer = original.BlobContainer
type BlobContainerProperties = original.BlobContainerProperties
type Client = original.Client
type CreateFuture = original.CreateFuture
type DataLakeAnalyticsAccount = original.DataLakeAnalyticsAccount
type DataLakeAnalyticsAccountListDataLakeStoreResult = original.DataLakeAnalyticsAccountListDataLakeStoreResult
type DataLakeAnalyticsAccountListDataLakeStoreResultIterator = original.DataLakeAnalyticsAccountListDataLakeStoreResultIterator
type DataLakeAnalyticsAccountListDataLakeStoreResultPage = original.DataLakeAnalyticsAccountListDataLakeStoreResultPage
type DataLakeAnalyticsAccountListResult = original.DataLakeAnalyticsAccountListResult
type DataLakeAnalyticsAccountListResultIterator = original.DataLakeAnalyticsAccountListResultIterator
type DataLakeAnalyticsAccountListResultPage = original.DataLakeAnalyticsAccountListResultPage
type DataLakeAnalyticsAccountListStorageAccountsResult = original.DataLakeAnalyticsAccountListStorageAccountsResult
type DataLakeAnalyticsAccountListStorageAccountsResultIterator = original.DataLakeAnalyticsAccountListStorageAccountsResultIterator
type DataLakeAnalyticsAccountListStorageAccountsResultPage = original.DataLakeAnalyticsAccountListStorageAccountsResultPage
type DataLakeAnalyticsAccountProperties = original.DataLakeAnalyticsAccountProperties
type DataLakeStoreAccountInfo = original.DataLakeStoreAccountInfo
type DataLakeStoreAccountInfoProperties = original.DataLakeStoreAccountInfoProperties
type DeleteFuture = original.DeleteFuture
type Error = original.Error
type ErrorDetails = original.ErrorDetails
type InnerError = original.InnerError
type ListBlobContainersResult = original.ListBlobContainersResult
type ListBlobContainersResultIterator = original.ListBlobContainersResultIterator
type ListBlobContainersResultPage = original.ListBlobContainersResultPage
type ListSasTokensResult = original.ListSasTokensResult
type ListSasTokensResultIterator = original.ListSasTokensResultIterator
type ListSasTokensResultPage = original.ListSasTokensResultPage
type SasTokenInfo = original.SasTokenInfo
type StorageAccountInfo = original.StorageAccountInfo
type StorageAccountProperties = original.StorageAccountProperties
type UpdateFuture = original.UpdateFuture

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewClient(subscriptionID string) Client {
	return original.NewClient(subscriptionID)
}
func NewClientWithBaseURI(baseURI string, subscriptionID string) Client {
	return original.NewClientWithBaseURI(baseURI, subscriptionID)
}
func NewDataLakeAnalyticsAccountListDataLakeStoreResultIterator(page DataLakeAnalyticsAccountListDataLakeStoreResultPage) DataLakeAnalyticsAccountListDataLakeStoreResultIterator {
	return original.NewDataLakeAnalyticsAccountListDataLakeStoreResultIterator(page)
}
func NewDataLakeAnalyticsAccountListDataLakeStoreResultPage(getNextPage func(context.Context, DataLakeAnalyticsAccountListDataLakeStoreResult) (DataLakeAnalyticsAccountListDataLakeStoreResult, error)) DataLakeAnalyticsAccountListDataLakeStoreResultPage {
	return original.NewDataLakeAnalyticsAccountListDataLakeStoreResultPage(getNextPage)
}
func NewDataLakeAnalyticsAccountListResultIterator(page DataLakeAnalyticsAccountListResultPage) DataLakeAnalyticsAccountListResultIterator {
	return original.NewDataLakeAnalyticsAccountListResultIterator(page)
}
func NewDataLakeAnalyticsAccountListResultPage(getNextPage func(context.Context, DataLakeAnalyticsAccountListResult) (DataLakeAnalyticsAccountListResult, error)) DataLakeAnalyticsAccountListResultPage {
	return original.NewDataLakeAnalyticsAccountListResultPage(getNextPage)
}
func NewDataLakeAnalyticsAccountListStorageAccountsResultIterator(page DataLakeAnalyticsAccountListStorageAccountsResultPage) DataLakeAnalyticsAccountListStorageAccountsResultIterator {
	return original.NewDataLakeAnalyticsAccountListStorageAccountsResultIterator(page)
}
func NewDataLakeAnalyticsAccountListStorageAccountsResultPage(getNextPage func(context.Context, DataLakeAnalyticsAccountListStorageAccountsResult) (DataLakeAnalyticsAccountListStorageAccountsResult, error)) DataLakeAnalyticsAccountListStorageAccountsResultPage {
	return original.NewDataLakeAnalyticsAccountListStorageAccountsResultPage(getNextPage)
}
func NewListBlobContainersResultIterator(page ListBlobContainersResultPage) ListBlobContainersResultIterator {
	return original.NewListBlobContainersResultIterator(page)
}
func NewListBlobContainersResultPage(getNextPage func(context.Context, ListBlobContainersResult) (ListBlobContainersResult, error)) ListBlobContainersResultPage {
	return original.NewListBlobContainersResultPage(getNextPage)
}
func NewListSasTokensResultIterator(page ListSasTokensResultPage) ListSasTokensResultIterator {
	return original.NewListSasTokensResultIterator(page)
}
func NewListSasTokensResultPage(getNextPage func(context.Context, ListSasTokensResult) (ListSasTokensResult, error)) ListSasTokensResultPage {
	return original.NewListSasTokensResultPage(getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleDataLakeAnalyticsAccountStateValues() []DataLakeAnalyticsAccountState {
	return original.PossibleDataLakeAnalyticsAccountStateValues()
}
func PossibleDataLakeAnalyticsAccountStatusValues() []DataLakeAnalyticsAccountStatus {
	return original.PossibleDataLakeAnalyticsAccountStatusValues()
}
func PossibleOperationStatusValues() []OperationStatus {
	return original.PossibleOperationStatusValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
