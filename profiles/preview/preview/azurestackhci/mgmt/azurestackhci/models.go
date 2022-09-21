//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package azurestackhci

import (
	"context"

	original "github.com/Azure/dev/azure-sdk-for-go/services/preview/azurestackhci/mgmt/2021-01-01-preview/azurestackhci"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ActionType = original.ActionType

const (
	ActionTypeInternal ActionType = original.ActionTypeInternal
)

type ArcSettingAggregateState = original.ArcSettingAggregateState

const (
	ArcSettingAggregateStateCanceled           ArcSettingAggregateState = original.ArcSettingAggregateStateCanceled
	ArcSettingAggregateStateConnected          ArcSettingAggregateState = original.ArcSettingAggregateStateConnected
	ArcSettingAggregateStateCreating           ArcSettingAggregateState = original.ArcSettingAggregateStateCreating
	ArcSettingAggregateStateDeleted            ArcSettingAggregateState = original.ArcSettingAggregateStateDeleted
	ArcSettingAggregateStateDeleting           ArcSettingAggregateState = original.ArcSettingAggregateStateDeleting
	ArcSettingAggregateStateDisconnected       ArcSettingAggregateState = original.ArcSettingAggregateStateDisconnected
	ArcSettingAggregateStateError              ArcSettingAggregateState = original.ArcSettingAggregateStateError
	ArcSettingAggregateStateFailed             ArcSettingAggregateState = original.ArcSettingAggregateStateFailed
	ArcSettingAggregateStateInProgress         ArcSettingAggregateState = original.ArcSettingAggregateStateInProgress
	ArcSettingAggregateStateMoving             ArcSettingAggregateState = original.ArcSettingAggregateStateMoving
	ArcSettingAggregateStateNotSpecified       ArcSettingAggregateState = original.ArcSettingAggregateStateNotSpecified
	ArcSettingAggregateStatePartiallyConnected ArcSettingAggregateState = original.ArcSettingAggregateStatePartiallyConnected
	ArcSettingAggregateStatePartiallySucceeded ArcSettingAggregateState = original.ArcSettingAggregateStatePartiallySucceeded
	ArcSettingAggregateStateSucceeded          ArcSettingAggregateState = original.ArcSettingAggregateStateSucceeded
	ArcSettingAggregateStateUpdating           ArcSettingAggregateState = original.ArcSettingAggregateStateUpdating
)

type CreatedByType = original.CreatedByType

const (
	CreatedByTypeApplication     CreatedByType = original.CreatedByTypeApplication
	CreatedByTypeKey             CreatedByType = original.CreatedByTypeKey
	CreatedByTypeManagedIdentity CreatedByType = original.CreatedByTypeManagedIdentity
	CreatedByTypeUser            CreatedByType = original.CreatedByTypeUser
)

type ExtensionAggregateState = original.ExtensionAggregateState

const (
	ExtensionAggregateStateCanceled           ExtensionAggregateState = original.ExtensionAggregateStateCanceled
	ExtensionAggregateStateConnected          ExtensionAggregateState = original.ExtensionAggregateStateConnected
	ExtensionAggregateStateCreating           ExtensionAggregateState = original.ExtensionAggregateStateCreating
	ExtensionAggregateStateDeleted            ExtensionAggregateState = original.ExtensionAggregateStateDeleted
	ExtensionAggregateStateDeleting           ExtensionAggregateState = original.ExtensionAggregateStateDeleting
	ExtensionAggregateStateDisconnected       ExtensionAggregateState = original.ExtensionAggregateStateDisconnected
	ExtensionAggregateStateError              ExtensionAggregateState = original.ExtensionAggregateStateError
	ExtensionAggregateStateFailed             ExtensionAggregateState = original.ExtensionAggregateStateFailed
	ExtensionAggregateStateInProgress         ExtensionAggregateState = original.ExtensionAggregateStateInProgress
	ExtensionAggregateStateMoving             ExtensionAggregateState = original.ExtensionAggregateStateMoving
	ExtensionAggregateStateNotSpecified       ExtensionAggregateState = original.ExtensionAggregateStateNotSpecified
	ExtensionAggregateStatePartiallyConnected ExtensionAggregateState = original.ExtensionAggregateStatePartiallyConnected
	ExtensionAggregateStatePartiallySucceeded ExtensionAggregateState = original.ExtensionAggregateStatePartiallySucceeded
	ExtensionAggregateStateSucceeded          ExtensionAggregateState = original.ExtensionAggregateStateSucceeded
	ExtensionAggregateStateUpdating           ExtensionAggregateState = original.ExtensionAggregateStateUpdating
)

type NodeArcState = original.NodeArcState

const (
	NodeArcStateCanceled     NodeArcState = original.NodeArcStateCanceled
	NodeArcStateConnected    NodeArcState = original.NodeArcStateConnected
	NodeArcStateCreating     NodeArcState = original.NodeArcStateCreating
	NodeArcStateDeleted      NodeArcState = original.NodeArcStateDeleted
	NodeArcStateDeleting     NodeArcState = original.NodeArcStateDeleting
	NodeArcStateDisconnected NodeArcState = original.NodeArcStateDisconnected
	NodeArcStateError        NodeArcState = original.NodeArcStateError
	NodeArcStateFailed       NodeArcState = original.NodeArcStateFailed
	NodeArcStateMoving       NodeArcState = original.NodeArcStateMoving
	NodeArcStateNotSpecified NodeArcState = original.NodeArcStateNotSpecified
	NodeArcStateSucceeded    NodeArcState = original.NodeArcStateSucceeded
	NodeArcStateUpdating     NodeArcState = original.NodeArcStateUpdating
)

type NodeExtensionState = original.NodeExtensionState

const (
	NodeExtensionStateCanceled     NodeExtensionState = original.NodeExtensionStateCanceled
	NodeExtensionStateConnected    NodeExtensionState = original.NodeExtensionStateConnected
	NodeExtensionStateCreating     NodeExtensionState = original.NodeExtensionStateCreating
	NodeExtensionStateDeleted      NodeExtensionState = original.NodeExtensionStateDeleted
	NodeExtensionStateDeleting     NodeExtensionState = original.NodeExtensionStateDeleting
	NodeExtensionStateDisconnected NodeExtensionState = original.NodeExtensionStateDisconnected
	NodeExtensionStateError        NodeExtensionState = original.NodeExtensionStateError
	NodeExtensionStateFailed       NodeExtensionState = original.NodeExtensionStateFailed
	NodeExtensionStateMoving       NodeExtensionState = original.NodeExtensionStateMoving
	NodeExtensionStateNotSpecified NodeExtensionState = original.NodeExtensionStateNotSpecified
	NodeExtensionStateSucceeded    NodeExtensionState = original.NodeExtensionStateSucceeded
	NodeExtensionStateUpdating     NodeExtensionState = original.NodeExtensionStateUpdating
)

type Origin = original.Origin

const (
	OriginSystem     Origin = original.OriginSystem
	OriginUser       Origin = original.OriginUser
	OriginUsersystem Origin = original.OriginUsersystem
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateAccepted     ProvisioningState = original.ProvisioningStateAccepted
	ProvisioningStateCanceled     ProvisioningState = original.ProvisioningStateCanceled
	ProvisioningStateFailed       ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateProvisioning ProvisioningState = original.ProvisioningStateProvisioning
	ProvisioningStateSucceeded    ProvisioningState = original.ProvisioningStateSucceeded
)

type Status = original.Status

const (
	StatusConnectedRecently    Status = original.StatusConnectedRecently
	StatusDisconnected         Status = original.StatusDisconnected
	StatusError                Status = original.StatusError
	StatusNotConnectedRecently Status = original.StatusNotConnectedRecently
	StatusNotYetRegistered     Status = original.StatusNotYetRegistered
)

type ArcSetting = original.ArcSetting
type ArcSettingList = original.ArcSettingList
type ArcSettingListIterator = original.ArcSettingListIterator
type ArcSettingListPage = original.ArcSettingListPage
type ArcSettingProperties = original.ArcSettingProperties
type ArcSettingsClient = original.ArcSettingsClient
type ArcSettingsDeleteFuture = original.ArcSettingsDeleteFuture
type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type Cluster = original.Cluster
type ClusterList = original.ClusterList
type ClusterListIterator = original.ClusterListIterator
type ClusterListPage = original.ClusterListPage
type ClusterNode = original.ClusterNode
type ClusterPatch = original.ClusterPatch
type ClusterPatchProperties = original.ClusterPatchProperties
type ClusterProperties = original.ClusterProperties
type ClusterReportedProperties = original.ClusterReportedProperties
type ClustersClient = original.ClustersClient
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type Extension = original.Extension
type ExtensionList = original.ExtensionList
type ExtensionListIterator = original.ExtensionListIterator
type ExtensionListPage = original.ExtensionListPage
type ExtensionParameters = original.ExtensionParameters
type ExtensionProperties = original.ExtensionProperties
type ExtensionsClient = original.ExtensionsClient
type ExtensionsCreateFuture = original.ExtensionsCreateFuture
type ExtensionsDeleteFuture = original.ExtensionsDeleteFuture
type ExtensionsUpdateFuture = original.ExtensionsUpdateFuture
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationsClient = original.OperationsClient
type PerNodeExtensionState = original.PerNodeExtensionState
type PerNodeState = original.PerNodeState
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type SystemData = original.SystemData
type TrackedResource = original.TrackedResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewArcSettingListIterator(page ArcSettingListPage) ArcSettingListIterator {
	return original.NewArcSettingListIterator(page)
}
func NewArcSettingListPage(cur ArcSettingList, getNextPage func(context.Context, ArcSettingList) (ArcSettingList, error)) ArcSettingListPage {
	return original.NewArcSettingListPage(cur, getNextPage)
}
func NewArcSettingsClient(subscriptionID string) ArcSettingsClient {
	return original.NewArcSettingsClient(subscriptionID)
}
func NewArcSettingsClientWithBaseURI(baseURI string, subscriptionID string) ArcSettingsClient {
	return original.NewArcSettingsClientWithBaseURI(baseURI, subscriptionID)
}
func NewClusterListIterator(page ClusterListPage) ClusterListIterator {
	return original.NewClusterListIterator(page)
}
func NewClusterListPage(cur ClusterList, getNextPage func(context.Context, ClusterList) (ClusterList, error)) ClusterListPage {
	return original.NewClusterListPage(cur, getNextPage)
}
func NewClustersClient(subscriptionID string) ClustersClient {
	return original.NewClustersClient(subscriptionID)
}
func NewClustersClientWithBaseURI(baseURI string, subscriptionID string) ClustersClient {
	return original.NewClustersClientWithBaseURI(baseURI, subscriptionID)
}
func NewExtensionListIterator(page ExtensionListPage) ExtensionListIterator {
	return original.NewExtensionListIterator(page)
}
func NewExtensionListPage(cur ExtensionList, getNextPage func(context.Context, ExtensionList) (ExtensionList, error)) ExtensionListPage {
	return original.NewExtensionListPage(cur, getNextPage)
}
func NewExtensionsClient(subscriptionID string) ExtensionsClient {
	return original.NewExtensionsClient(subscriptionID)
}
func NewExtensionsClientWithBaseURI(baseURI string, subscriptionID string) ExtensionsClient {
	return original.NewExtensionsClientWithBaseURI(baseURI, subscriptionID)
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
func PossibleActionTypeValues() []ActionType {
	return original.PossibleActionTypeValues()
}
func PossibleArcSettingAggregateStateValues() []ArcSettingAggregateState {
	return original.PossibleArcSettingAggregateStateValues()
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleExtensionAggregateStateValues() []ExtensionAggregateState {
	return original.PossibleExtensionAggregateStateValues()
}
func PossibleNodeArcStateValues() []NodeArcState {
	return original.PossibleNodeArcStateValues()
}
func PossibleNodeExtensionStateValues() []NodeExtensionState {
	return original.PossibleNodeExtensionStateValues()
}
func PossibleOriginValues() []Origin {
	return original.PossibleOriginValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleStatusValues() []Status {
	return original.PossibleStatusValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
