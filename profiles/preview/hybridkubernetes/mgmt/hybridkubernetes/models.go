//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package hybridkubernetes

import (
	"context"

	original "github.com/Azure/dev/azure-sdk-for-go/services/hybridkubernetes/mgmt/2021-10-01/hybridkubernetes"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AuthenticationMethod = original.AuthenticationMethod

const (
	AuthenticationMethodAAD   AuthenticationMethod = original.AuthenticationMethodAAD
	AuthenticationMethodToken AuthenticationMethod = original.AuthenticationMethodToken
)

type ConnectivityStatus = original.ConnectivityStatus

const (
	ConnectivityStatusConnected  ConnectivityStatus = original.ConnectivityStatusConnected
	ConnectivityStatusConnecting ConnectivityStatus = original.ConnectivityStatusConnecting
	ConnectivityStatusExpired    ConnectivityStatus = original.ConnectivityStatusExpired
	ConnectivityStatusOffline    ConnectivityStatus = original.ConnectivityStatusOffline
)

type CreatedByType = original.CreatedByType

const (
	CreatedByTypeApplication     CreatedByType = original.CreatedByTypeApplication
	CreatedByTypeKey             CreatedByType = original.CreatedByTypeKey
	CreatedByTypeManagedIdentity CreatedByType = original.CreatedByTypeManagedIdentity
	CreatedByTypeUser            CreatedByType = original.CreatedByTypeUser
)

type LastModifiedByType = original.LastModifiedByType

const (
	LastModifiedByTypeApplication     LastModifiedByType = original.LastModifiedByTypeApplication
	LastModifiedByTypeKey             LastModifiedByType = original.LastModifiedByTypeKey
	LastModifiedByTypeManagedIdentity LastModifiedByType = original.LastModifiedByTypeManagedIdentity
	LastModifiedByTypeUser            LastModifiedByType = original.LastModifiedByTypeUser
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateAccepted     ProvisioningState = original.ProvisioningStateAccepted
	ProvisioningStateCanceled     ProvisioningState = original.ProvisioningStateCanceled
	ProvisioningStateDeleting     ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateFailed       ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateProvisioning ProvisioningState = original.ProvisioningStateProvisioning
	ProvisioningStateSucceeded    ProvisioningState = original.ProvisioningStateSucceeded
	ProvisioningStateUpdating     ProvisioningState = original.ProvisioningStateUpdating
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	ResourceIdentityTypeNone           ResourceIdentityType = original.ResourceIdentityTypeNone
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = original.ResourceIdentityTypeSystemAssigned
)

type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type ConnectedCluster = original.ConnectedCluster
type ConnectedClusterClient = original.ConnectedClusterClient
type ConnectedClusterCreateFuture = original.ConnectedClusterCreateFuture
type ConnectedClusterDeleteFuture = original.ConnectedClusterDeleteFuture
type ConnectedClusterIdentity = original.ConnectedClusterIdentity
type ConnectedClusterList = original.ConnectedClusterList
type ConnectedClusterListIterator = original.ConnectedClusterListIterator
type ConnectedClusterListPage = original.ConnectedClusterListPage
type ConnectedClusterPatch = original.ConnectedClusterPatch
type ConnectedClusterProperties = original.ConnectedClusterProperties
type CredentialResult = original.CredentialResult
type CredentialResults = original.CredentialResults
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type HybridConnectionConfig = original.HybridConnectionConfig
type ListClusterUserCredentialProperties = original.ListClusterUserCredentialProperties
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationList = original.OperationList
type OperationListIterator = original.OperationListIterator
type OperationListPage = original.OperationListPage
type OperationsClient = original.OperationsClient
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type SystemData = original.SystemData
type TrackedResource = original.TrackedResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewConnectedClusterClient(subscriptionID string) ConnectedClusterClient {
	return original.NewConnectedClusterClient(subscriptionID)
}
func NewConnectedClusterClientWithBaseURI(baseURI string, subscriptionID string) ConnectedClusterClient {
	return original.NewConnectedClusterClientWithBaseURI(baseURI, subscriptionID)
}
func NewConnectedClusterListIterator(page ConnectedClusterListPage) ConnectedClusterListIterator {
	return original.NewConnectedClusterListIterator(page)
}
func NewConnectedClusterListPage(cur ConnectedClusterList, getNextPage func(context.Context, ConnectedClusterList) (ConnectedClusterList, error)) ConnectedClusterListPage {
	return original.NewConnectedClusterListPage(cur, getNextPage)
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
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAuthenticationMethodValues() []AuthenticationMethod {
	return original.PossibleAuthenticationMethodValues()
}
func PossibleConnectivityStatusValues() []ConnectivityStatus {
	return original.PossibleConnectivityStatusValues()
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleLastModifiedByTypeValues() []LastModifiedByType {
	return original.PossibleLastModifiedByTypeValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
