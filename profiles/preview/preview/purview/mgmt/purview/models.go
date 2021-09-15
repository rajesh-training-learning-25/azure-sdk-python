//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package purview

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/purview/mgmt/2020-12-01-preview/purview"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type Name = original.Name

const (
	Standard Name = original.Standard
)

type ProvisioningState = original.ProvisioningState

const (
	Creating     ProvisioningState = original.Creating
	Deleting     ProvisioningState = original.Deleting
	Failed       ProvisioningState = original.Failed
	Moving       ProvisioningState = original.Moving
	SoftDeleted  ProvisioningState = original.SoftDeleted
	SoftDeleting ProvisioningState = original.SoftDeleting
	Succeeded    ProvisioningState = original.Succeeded
	Unknown      ProvisioningState = original.Unknown
)

type PublicNetworkAccess = original.PublicNetworkAccess

const (
	Disabled     PublicNetworkAccess = original.Disabled
	Enabled      PublicNetworkAccess = original.Enabled
	NotSpecified PublicNetworkAccess = original.NotSpecified
)

type Reason = original.Reason

const (
	AlreadyExists Reason = original.AlreadyExists
	Invalid       Reason = original.Invalid
)

type ScopeType = original.ScopeType

const (
	Subscription ScopeType = original.Subscription
	Tenant       ScopeType = original.Tenant
)

type Status = original.Status

const (
	StatusApproved     Status = original.StatusApproved
	StatusDisconnected Status = original.StatusDisconnected
	StatusPending      Status = original.StatusPending
	StatusRejected     Status = original.StatusRejected
	StatusUnknown      Status = original.StatusUnknown
)

type Type = original.Type

const (
	SystemAssigned Type = original.SystemAssigned
)

type AccessKeys = original.AccessKeys
type Account = original.Account
type AccountEndpoints = original.AccountEndpoints
type AccountList = original.AccountList
type AccountListIterator = original.AccountListIterator
type AccountListPage = original.AccountListPage
type AccountProperties = original.AccountProperties
type AccountPropertiesEndpoints = original.AccountPropertiesEndpoints
type AccountPropertiesManagedResources = original.AccountPropertiesManagedResources
type AccountSku = original.AccountSku
type AccountUpdateParameters = original.AccountUpdateParameters
type AccountsClient = original.AccountsClient
type AccountsCreateOrUpdateFuture = original.AccountsCreateOrUpdateFuture
type AccountsDeleteFuture = original.AccountsDeleteFuture
type AccountsUpdateFuture = original.AccountsUpdateFuture
type BaseClient = original.BaseClient
type CheckNameAvailabilityRequest = original.CheckNameAvailabilityRequest
type CheckNameAvailabilityResult = original.CheckNameAvailabilityResult
type CloudConnectors = original.CloudConnectors
type DefaultAccountPayload = original.DefaultAccountPayload
type DefaultAccountsClient = original.DefaultAccountsClient
type DeletedAccount = original.DeletedAccount
type DeletedAccountList = original.DeletedAccountList
type DeletedAccountProperties = original.DeletedAccountProperties
type DeletedAccountPropertiesModel = original.DeletedAccountPropertiesModel
type DimensionProperties = original.DimensionProperties
type ErrorModel = original.ErrorModel
type ErrorResponseModel = original.ErrorResponseModel
type ErrorResponseModelError = original.ErrorResponseModelError
type Identity = original.Identity
type ManagedResources = original.ManagedResources
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationList = original.OperationList
type OperationListIterator = original.OperationListIterator
type OperationListPage = original.OperationListPage
type OperationMetaLogSpecification = original.OperationMetaLogSpecification
type OperationMetaMetricSpecification = original.OperationMetaMetricSpecification
type OperationMetaServiceSpecification = original.OperationMetaServiceSpecification
type OperationProperties = original.OperationProperties
type OperationsClient = original.OperationsClient
type PrivateEndpoint = original.PrivateEndpoint
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionList = original.PrivateEndpointConnectionList
type PrivateEndpointConnectionListIterator = original.PrivateEndpointConnectionListIterator
type PrivateEndpointConnectionListPage = original.PrivateEndpointConnectionListPage
type PrivateEndpointConnectionProperties = original.PrivateEndpointConnectionProperties
type PrivateEndpointConnectionsClient = original.PrivateEndpointConnectionsClient
type PrivateEndpointConnectionsDeleteFuture = original.PrivateEndpointConnectionsDeleteFuture
type PrivateLinkResource = original.PrivateLinkResource
type PrivateLinkResourceList = original.PrivateLinkResourceList
type PrivateLinkResourceListIterator = original.PrivateLinkResourceListIterator
type PrivateLinkResourceListPage = original.PrivateLinkResourceListPage
type PrivateLinkResourceProperties = original.PrivateLinkResourceProperties
type PrivateLinkResourcesClient = original.PrivateLinkResourcesClient
type PrivateLinkServiceConnectionState = original.PrivateLinkServiceConnectionState
type ProxyResource = original.ProxyResource
type TrackedResource = original.TrackedResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAccountListIterator(page AccountListPage) AccountListIterator {
	return original.NewAccountListIterator(page)
}
func NewAccountListPage(cur AccountList, getNextPage func(context.Context, AccountList) (AccountList, error)) AccountListPage {
	return original.NewAccountListPage(cur, getNextPage)
}
func NewAccountsClient(subscriptionID string) AccountsClient {
	return original.NewAccountsClient(subscriptionID)
}
func NewAccountsClientWithBaseURI(baseURI string, subscriptionID string) AccountsClient {
	return original.NewAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDefaultAccountsClient(subscriptionID string) DefaultAccountsClient {
	return original.NewDefaultAccountsClient(subscriptionID)
}
func NewDefaultAccountsClientWithBaseURI(baseURI string, subscriptionID string) DefaultAccountsClient {
	return original.NewDefaultAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListIterator(page OperationListPage) OperationListIterator {
	return original.NewOperationListIterator(page)
}
func NewOperationListPage(cur OperationList, getNextPage func(context.Context, OperationList) (OperationList, error)) OperationListPage {
	return original.NewOperationListPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateEndpointConnectionListIterator(page PrivateEndpointConnectionListPage) PrivateEndpointConnectionListIterator {
	return original.NewPrivateEndpointConnectionListIterator(page)
}
func NewPrivateEndpointConnectionListPage(cur PrivateEndpointConnectionList, getNextPage func(context.Context, PrivateEndpointConnectionList) (PrivateEndpointConnectionList, error)) PrivateEndpointConnectionListPage {
	return original.NewPrivateEndpointConnectionListPage(cur, getNextPage)
}
func NewPrivateEndpointConnectionsClient(subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClient(subscriptionID)
}
func NewPrivateEndpointConnectionsClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateLinkResourceListIterator(page PrivateLinkResourceListPage) PrivateLinkResourceListIterator {
	return original.NewPrivateLinkResourceListIterator(page)
}
func NewPrivateLinkResourceListPage(cur PrivateLinkResourceList, getNextPage func(context.Context, PrivateLinkResourceList) (PrivateLinkResourceList, error)) PrivateLinkResourceListPage {
	return original.NewPrivateLinkResourceListPage(cur, getNextPage)
}
func NewPrivateLinkResourcesClient(subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClient(subscriptionID)
}
func NewPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleNameValues() []Name {
	return original.PossibleNameValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return original.PossiblePublicNetworkAccessValues()
}
func PossibleReasonValues() []Reason {
	return original.PossibleReasonValues()
}
func PossibleScopeTypeValues() []ScopeType {
	return original.PossibleScopeTypeValues()
}
func PossibleStatusValues() []Status {
	return original.PossibleStatusValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
