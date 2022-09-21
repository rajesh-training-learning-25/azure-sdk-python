//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package deviceupdate

import (
	"context"

	original "github.com/Azure/dev/azure-sdk-for-go/services/preview/deviceupdate/mgmt/2022-04-01-preview/deviceupdate"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ActionType = original.ActionType

const (
	ActionTypeInternal ActionType = original.ActionTypeInternal
)

type CheckNameAvailabilityReason = original.CheckNameAvailabilityReason

const (
	CheckNameAvailabilityReasonAlreadyExists CheckNameAvailabilityReason = original.CheckNameAvailabilityReasonAlreadyExists
	CheckNameAvailabilityReasonInvalid       CheckNameAvailabilityReason = original.CheckNameAvailabilityReasonInvalid
)

type CreatedByType = original.CreatedByType

const (
	CreatedByTypeApplication     CreatedByType = original.CreatedByTypeApplication
	CreatedByTypeKey             CreatedByType = original.CreatedByTypeKey
	CreatedByTypeManagedIdentity CreatedByType = original.CreatedByTypeManagedIdentity
	CreatedByTypeUser            CreatedByType = original.CreatedByTypeUser
)

type GroupIDProvisioningState = original.GroupIDProvisioningState

const (
	GroupIDProvisioningStateCanceled  GroupIDProvisioningState = original.GroupIDProvisioningStateCanceled
	GroupIDProvisioningStateFailed    GroupIDProvisioningState = original.GroupIDProvisioningStateFailed
	GroupIDProvisioningStateSucceeded GroupIDProvisioningState = original.GroupIDProvisioningStateSucceeded
)

type ManagedServiceIdentityType = original.ManagedServiceIdentityType

const (
	ManagedServiceIdentityTypeNone                       ManagedServiceIdentityType = original.ManagedServiceIdentityTypeNone
	ManagedServiceIdentityTypeSystemAssigned             ManagedServiceIdentityType = original.ManagedServiceIdentityTypeSystemAssigned
	ManagedServiceIdentityTypeSystemAssignedUserAssigned ManagedServiceIdentityType = original.ManagedServiceIdentityTypeSystemAssignedUserAssigned
	ManagedServiceIdentityTypeUserAssigned               ManagedServiceIdentityType = original.ManagedServiceIdentityTypeUserAssigned
)

type Origin = original.Origin

const (
	OriginSystem     Origin = original.OriginSystem
	OriginUser       Origin = original.OriginUser
	OriginUsersystem Origin = original.OriginUsersystem
)

type PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningState

const (
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateCreating
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateDeleting
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateFailed
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateSucceeded
)

type PrivateEndpointConnectionProxyProvisioningState = original.PrivateEndpointConnectionProxyProvisioningState

const (
	PrivateEndpointConnectionProxyProvisioningStateCreating  PrivateEndpointConnectionProxyProvisioningState = original.PrivateEndpointConnectionProxyProvisioningStateCreating
	PrivateEndpointConnectionProxyProvisioningStateDeleting  PrivateEndpointConnectionProxyProvisioningState = original.PrivateEndpointConnectionProxyProvisioningStateDeleting
	PrivateEndpointConnectionProxyProvisioningStateFailed    PrivateEndpointConnectionProxyProvisioningState = original.PrivateEndpointConnectionProxyProvisioningStateFailed
	PrivateEndpointConnectionProxyProvisioningStateSucceeded PrivateEndpointConnectionProxyProvisioningState = original.PrivateEndpointConnectionProxyProvisioningStateSucceeded
)

type PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatus

const (
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatusApproved
	PrivateEndpointServiceConnectionStatusPending  PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatusPending
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatusRejected
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateAccepted  ProvisioningState = original.ProvisioningStateAccepted
	ProvisioningStateCanceled  ProvisioningState = original.ProvisioningStateCanceled
	ProvisioningStateCreating  ProvisioningState = original.ProvisioningStateCreating
	ProvisioningStateDeleted   ProvisioningState = original.ProvisioningStateDeleted
	ProvisioningStateFailed    ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateSucceeded ProvisioningState = original.ProvisioningStateSucceeded
)

type PublicNetworkAccess = original.PublicNetworkAccess

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = original.PublicNetworkAccessDisabled
	PublicNetworkAccessEnabled  PublicNetworkAccess = original.PublicNetworkAccessEnabled
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = original.ResourceIdentityTypeSystemAssigned
)

type Role = original.Role

const (
	RoleFailover Role = original.RoleFailover
	RolePrimary  Role = original.RolePrimary
)

type SKU = original.SKU

const (
	SKUFree     SKU = original.SKUFree
	SKUStandard SKU = original.SKUStandard
)

type SkuTier = original.SkuTier

const (
	SkuTierBasic    SkuTier = original.SkuTierBasic
	SkuTierFree     SkuTier = original.SkuTierFree
	SkuTierPremium  SkuTier = original.SkuTierPremium
	SkuTierStandard SkuTier = original.SkuTierStandard
)

type Account = original.Account
type AccountList = original.AccountList
type AccountListIterator = original.AccountListIterator
type AccountListPage = original.AccountListPage
type AccountProperties = original.AccountProperties
type AccountUpdate = original.AccountUpdate
type AccountsClient = original.AccountsClient
type AccountsCreateFuture = original.AccountsCreateFuture
type AccountsDeleteFuture = original.AccountsDeleteFuture
type AccountsUpdateFuture = original.AccountsUpdateFuture
type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type CheckNameAvailabilityRequest = original.CheckNameAvailabilityRequest
type CheckNameAvailabilityResponse = original.CheckNameAvailabilityResponse
type ConnectionDetails = original.ConnectionDetails
type DiagnosticStorageProperties = original.DiagnosticStorageProperties
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type GroupConnectivityInformation = original.GroupConnectivityInformation
type GroupInformation = original.GroupInformation
type GroupInformationProperties = original.GroupInformationProperties
type Identity = original.Identity
type Instance = original.Instance
type InstanceList = original.InstanceList
type InstanceListIterator = original.InstanceListIterator
type InstanceListPage = original.InstanceListPage
type InstanceProperties = original.InstanceProperties
type InstancesClient = original.InstancesClient
type InstancesCreateFuture = original.InstancesCreateFuture
type InstancesDeleteFuture = original.InstancesDeleteFuture
type IotHubSettings = original.IotHubSettings
type Location = original.Location
type ManagedServiceIdentity = original.ManagedServiceIdentity
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type Plan = original.Plan
type PrivateEndpoint = original.PrivateEndpoint
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionListResult = original.PrivateEndpointConnectionListResult
type PrivateEndpointConnectionProperties = original.PrivateEndpointConnectionProperties
type PrivateEndpointConnectionProxiesClient = original.PrivateEndpointConnectionProxiesClient
type PrivateEndpointConnectionProxiesCreateOrUpdateFuture = original.PrivateEndpointConnectionProxiesCreateOrUpdateFuture
type PrivateEndpointConnectionProxiesDeleteFuture = original.PrivateEndpointConnectionProxiesDeleteFuture
type PrivateEndpointConnectionProxy = original.PrivateEndpointConnectionProxy
type PrivateEndpointConnectionProxyListResult = original.PrivateEndpointConnectionProxyListResult
type PrivateEndpointConnectionProxyProperties = original.PrivateEndpointConnectionProxyProperties
type PrivateEndpointConnectionProxyPropertiesModel = original.PrivateEndpointConnectionProxyPropertiesModel
type PrivateEndpointConnectionsClient = original.PrivateEndpointConnectionsClient
type PrivateEndpointConnectionsCreateOrUpdateFuture = original.PrivateEndpointConnectionsCreateOrUpdateFuture
type PrivateEndpointConnectionsDeleteFuture = original.PrivateEndpointConnectionsDeleteFuture
type PrivateEndpointUpdate = original.PrivateEndpointUpdate
type PrivateLinkResourceListResult = original.PrivateLinkResourceListResult
type PrivateLinkResourceProperties = original.PrivateLinkResourceProperties
type PrivateLinkResourcesClient = original.PrivateLinkResourcesClient
type PrivateLinkServiceConnection = original.PrivateLinkServiceConnection
type PrivateLinkServiceConnectionState = original.PrivateLinkServiceConnectionState
type PrivateLinkServiceProxy = original.PrivateLinkServiceProxy
type PrivateLinkServiceProxyRemotePrivateEndpointConnection = original.PrivateLinkServiceProxyRemotePrivateEndpointConnection
type ProxyResource = original.ProxyResource
type RemotePrivateEndpoint = original.RemotePrivateEndpoint
type RemotePrivateEndpointConnection = original.RemotePrivateEndpointConnection
type Resource = original.Resource
type ResourceModelWithAllowedPropertySet = original.ResourceModelWithAllowedPropertySet
type ResourceModelWithAllowedPropertySetIdentity = original.ResourceModelWithAllowedPropertySetIdentity
type ResourceModelWithAllowedPropertySetPlan = original.ResourceModelWithAllowedPropertySetPlan
type ResourceModelWithAllowedPropertySetSku = original.ResourceModelWithAllowedPropertySetSku
type Sku = original.Sku
type SystemData = original.SystemData
type TagUpdate = original.TagUpdate
type TrackedResource = original.TrackedResource
type UserAssignedIdentity = original.UserAssignedIdentity

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
func NewInstanceListIterator(page InstanceListPage) InstanceListIterator {
	return original.NewInstanceListIterator(page)
}
func NewInstanceListPage(cur InstanceList, getNextPage func(context.Context, InstanceList) (InstanceList, error)) InstanceListPage {
	return original.NewInstanceListPage(cur, getNextPage)
}
func NewInstancesClient(subscriptionID string) InstancesClient {
	return original.NewInstancesClient(subscriptionID)
}
func NewInstancesClientWithBaseURI(baseURI string, subscriptionID string) InstancesClient {
	return original.NewInstancesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateEndpointConnectionProxiesClient(subscriptionID string) PrivateEndpointConnectionProxiesClient {
	return original.NewPrivateEndpointConnectionProxiesClient(subscriptionID)
}
func NewPrivateEndpointConnectionProxiesClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionProxiesClient {
	return original.NewPrivateEndpointConnectionProxiesClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateEndpointConnectionsClient(subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClient(subscriptionID)
}
func NewPrivateEndpointConnectionsClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClientWithBaseURI(baseURI, subscriptionID)
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
func PossibleActionTypeValues() []ActionType {
	return original.PossibleActionTypeValues()
}
func PossibleCheckNameAvailabilityReasonValues() []CheckNameAvailabilityReason {
	return original.PossibleCheckNameAvailabilityReasonValues()
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleGroupIDProvisioningStateValues() []GroupIDProvisioningState {
	return original.PossibleGroupIDProvisioningStateValues()
}
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return original.PossibleManagedServiceIdentityTypeValues()
}
func PossibleOriginValues() []Origin {
	return original.PossibleOriginValues()
}
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return original.PossiblePrivateEndpointConnectionProvisioningStateValues()
}
func PossiblePrivateEndpointConnectionProxyProvisioningStateValues() []PrivateEndpointConnectionProxyProvisioningState {
	return original.PossiblePrivateEndpointConnectionProxyProvisioningStateValues()
}
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return original.PossiblePrivateEndpointServiceConnectionStatusValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return original.PossiblePublicNetworkAccessValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func PossibleRoleValues() []Role {
	return original.PossibleRoleValues()
}
func PossibleSKUValues() []SKU {
	return original.PossibleSKUValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
