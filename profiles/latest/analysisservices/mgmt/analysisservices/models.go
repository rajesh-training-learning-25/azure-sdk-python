//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package analysisservices

import (
	"context"

	original "github.com/Azure/temp/github.com/Azure/azure-sdk-for-go/services/analysisservices/mgmt/2017-08-01/analysisservices"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ConnectionMode = original.ConnectionMode

const (
	All      ConnectionMode = original.All
	ReadOnly ConnectionMode = original.ReadOnly
)

type ProvisioningState = original.ProvisioningState

const (
	Deleting     ProvisioningState = original.Deleting
	Failed       ProvisioningState = original.Failed
	Paused       ProvisioningState = original.Paused
	Pausing      ProvisioningState = original.Pausing
	Preparing    ProvisioningState = original.Preparing
	Provisioning ProvisioningState = original.Provisioning
	Resuming     ProvisioningState = original.Resuming
	Scaling      ProvisioningState = original.Scaling
	Succeeded    ProvisioningState = original.Succeeded
	Suspended    ProvisioningState = original.Suspended
	Suspending   ProvisioningState = original.Suspending
	Updating     ProvisioningState = original.Updating
)

type SkuTier = original.SkuTier

const (
	Basic       SkuTier = original.Basic
	Development SkuTier = original.Development
	Standard    SkuTier = original.Standard
)

type State = original.State

const (
	StateDeleting     State = original.StateDeleting
	StateFailed       State = original.StateFailed
	StatePaused       State = original.StatePaused
	StatePausing      State = original.StatePausing
	StatePreparing    State = original.StatePreparing
	StateProvisioning State = original.StateProvisioning
	StateResuming     State = original.StateResuming
	StateScaling      State = original.StateScaling
	StateSucceeded    State = original.StateSucceeded
	StateSuspended    State = original.StateSuspended
	StateSuspending   State = original.StateSuspending
	StateUpdating     State = original.StateUpdating
)

type Status = original.Status

const (
	Live Status = original.Live
)

type BaseClient = original.BaseClient
type CheckServerNameAvailabilityParameters = original.CheckServerNameAvailabilityParameters
type CheckServerNameAvailabilityResult = original.CheckServerNameAvailabilityResult
type ErrorResponse = original.ErrorResponse
type GatewayDetails = original.GatewayDetails
type GatewayError = original.GatewayError
type GatewayListStatusError = original.GatewayListStatusError
type GatewayListStatusLive = original.GatewayListStatusLive
type IPv4FirewallRule = original.IPv4FirewallRule
type IPv4FirewallSettings = original.IPv4FirewallSettings
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationStatus = original.OperationStatus
type OperationsClient = original.OperationsClient
type Resource = original.Resource
type ResourceSku = original.ResourceSku
type Server = original.Server
type ServerAdministrators = original.ServerAdministrators
type ServerMutableProperties = original.ServerMutableProperties
type ServerProperties = original.ServerProperties
type ServerUpdateParameters = original.ServerUpdateParameters
type Servers = original.Servers
type ServersClient = original.ServersClient
type ServersCreateFuture = original.ServersCreateFuture
type ServersDeleteFuture = original.ServersDeleteFuture
type ServersResumeFuture = original.ServersResumeFuture
type ServersSuspendFuture = original.ServersSuspendFuture
type ServersUpdateFuture = original.ServersUpdateFuture
type SkuDetailsForExistingResource = original.SkuDetailsForExistingResource
type SkuEnumerationForExistingResourceResult = original.SkuEnumerationForExistingResourceResult
type SkuEnumerationForNewResourceResult = original.SkuEnumerationForNewResourceResult

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
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
func NewServersClient(subscriptionID string) ServersClient {
	return original.NewServersClient(subscriptionID)
}
func NewServersClientWithBaseURI(baseURI string, subscriptionID string) ServersClient {
	return original.NewServersClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleConnectionModeValues() []ConnectionMode {
	return original.PossibleConnectionModeValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func PossibleStateValues() []State {
	return original.PossibleStateValues()
}
func PossibleStatusValues() []Status {
	return original.PossibleStatusValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
