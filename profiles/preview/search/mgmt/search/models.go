// +build go1.9

// Copyright 2020 Microsoft Corporation
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

package search

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/search/mgmt/2020-03-13/search"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AdminKeyKind = original.AdminKeyKind

const (
	Primary   AdminKeyKind = original.Primary
	Secondary AdminKeyKind = original.Secondary
)

type HostingMode = original.HostingMode

const (
	Default     HostingMode = original.Default
	HighDensity HostingMode = original.HighDensity
)

type IdentityType = original.IdentityType

const (
	None           IdentityType = original.None
	SystemAssigned IdentityType = original.SystemAssigned
)

type PrivateLinkServiceConnectionStatus = original.PrivateLinkServiceConnectionStatus

const (
	Approved     PrivateLinkServiceConnectionStatus = original.Approved
	Disconnected PrivateLinkServiceConnectionStatus = original.Disconnected
	Pending      PrivateLinkServiceConnectionStatus = original.Pending
	Rejected     PrivateLinkServiceConnectionStatus = original.Rejected
)

type ProvisioningState = original.ProvisioningState

const (
	Failed       ProvisioningState = original.Failed
	Provisioning ProvisioningState = original.Provisioning
	Succeeded    ProvisioningState = original.Succeeded
)

type PublicNetworkAccess = original.PublicNetworkAccess

const (
	Disabled PublicNetworkAccess = original.Disabled
	Enabled  PublicNetworkAccess = original.Enabled
)

type ServiceStatus = original.ServiceStatus

const (
	ServiceStatusDegraded     ServiceStatus = original.ServiceStatusDegraded
	ServiceStatusDeleting     ServiceStatus = original.ServiceStatusDeleting
	ServiceStatusDisabled     ServiceStatus = original.ServiceStatusDisabled
	ServiceStatusError        ServiceStatus = original.ServiceStatusError
	ServiceStatusProvisioning ServiceStatus = original.ServiceStatusProvisioning
	ServiceStatusRunning      ServiceStatus = original.ServiceStatusRunning
)

type SharedPrivateLinkResourceStatus = original.SharedPrivateLinkResourceStatus

const (
	SharedPrivateLinkResourceStatusApproved     SharedPrivateLinkResourceStatus = original.SharedPrivateLinkResourceStatusApproved
	SharedPrivateLinkResourceStatusDisconnected SharedPrivateLinkResourceStatus = original.SharedPrivateLinkResourceStatusDisconnected
	SharedPrivateLinkResourceStatusPending      SharedPrivateLinkResourceStatus = original.SharedPrivateLinkResourceStatusPending
	SharedPrivateLinkResourceStatusRejected     SharedPrivateLinkResourceStatus = original.SharedPrivateLinkResourceStatusRejected
	SharedPrivateLinkResourceStatusTimeout      SharedPrivateLinkResourceStatus = original.SharedPrivateLinkResourceStatusTimeout
)

type SkuName = original.SkuName

const (
	Basic              SkuName = original.Basic
	Free               SkuName = original.Free
	Standard           SkuName = original.Standard
	Standard2          SkuName = original.Standard2
	Standard3          SkuName = original.Standard3
	StorageOptimizedL1 SkuName = original.StorageOptimizedL1
	StorageOptimizedL2 SkuName = original.StorageOptimizedL2
)

type UnavailableNameReason = original.UnavailableNameReason

const (
	AlreadyExists UnavailableNameReason = original.AlreadyExists
	Invalid       UnavailableNameReason = original.Invalid
)

type AdminKeyResult = original.AdminKeyResult
type AdminKeysClient = original.AdminKeysClient
type BaseClient = original.BaseClient
type CheckNameAvailabilityInput = original.CheckNameAvailabilityInput
type CheckNameAvailabilityOutput = original.CheckNameAvailabilityOutput
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type IPRule = original.IPRule
type Identity = original.Identity
type ListQueryKeysResult = original.ListQueryKeysResult
type ListQueryKeysResultIterator = original.ListQueryKeysResultIterator
type ListQueryKeysResultPage = original.ListQueryKeysResultPage
type NetworkRuleSet = original.NetworkRuleSet
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationsClient = original.OperationsClient
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionListResult = original.PrivateEndpointConnectionListResult
type PrivateEndpointConnectionListResultIterator = original.PrivateEndpointConnectionListResultIterator
type PrivateEndpointConnectionListResultPage = original.PrivateEndpointConnectionListResultPage
type PrivateEndpointConnectionProperties = original.PrivateEndpointConnectionProperties
type PrivateEndpointConnectionPropertiesPrivateEndpoint = original.PrivateEndpointConnectionPropertiesPrivateEndpoint
type PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState = original.PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState
type PrivateEndpointConnectionsClient = original.PrivateEndpointConnectionsClient
type PrivateLinkResource = original.PrivateLinkResource
type PrivateLinkResourceProperties = original.PrivateLinkResourceProperties
type PrivateLinkResourcesClient = original.PrivateLinkResourcesClient
type PrivateLinkResourcesResult = original.PrivateLinkResourcesResult
type QueryKey = original.QueryKey
type QueryKeysClient = original.QueryKeysClient
type Resource = original.Resource
type Service = original.Service
type ServiceListResult = original.ServiceListResult
type ServiceListResultIterator = original.ServiceListResultIterator
type ServiceListResultPage = original.ServiceListResultPage
type ServiceProperties = original.ServiceProperties
type ServicesClient = original.ServicesClient
type ServicesCreateOrUpdateFuture = original.ServicesCreateOrUpdateFuture
type ShareablePrivateLinkResourceProperties = original.ShareablePrivateLinkResourceProperties
type ShareablePrivateLinkResourceType = original.ShareablePrivateLinkResourceType
type SharedPrivateLinkResource = original.SharedPrivateLinkResource
type SharedPrivateLinkResourceListResult = original.SharedPrivateLinkResourceListResult
type SharedPrivateLinkResourceListResultIterator = original.SharedPrivateLinkResourceListResultIterator
type SharedPrivateLinkResourceListResultPage = original.SharedPrivateLinkResourceListResultPage
type SharedPrivateLinkResourceProperties = original.SharedPrivateLinkResourceProperties
type SharedPrivateLinkResourcesClient = original.SharedPrivateLinkResourcesClient
type Sku = original.Sku

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAdminKeysClient(subscriptionID string) AdminKeysClient {
	return original.NewAdminKeysClient(subscriptionID)
}
func NewAdminKeysClientWithBaseURI(baseURI string, subscriptionID string) AdminKeysClient {
	return original.NewAdminKeysClientWithBaseURI(baseURI, subscriptionID)
}
func NewListQueryKeysResultIterator(page ListQueryKeysResultPage) ListQueryKeysResultIterator {
	return original.NewListQueryKeysResultIterator(page)
}
func NewListQueryKeysResultPage(getNextPage func(context.Context, ListQueryKeysResult) (ListQueryKeysResult, error)) ListQueryKeysResultPage {
	return original.NewListQueryKeysResultPage(getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateEndpointConnectionListResultIterator(page PrivateEndpointConnectionListResultPage) PrivateEndpointConnectionListResultIterator {
	return original.NewPrivateEndpointConnectionListResultIterator(page)
}
func NewPrivateEndpointConnectionListResultPage(getNextPage func(context.Context, PrivateEndpointConnectionListResult) (PrivateEndpointConnectionListResult, error)) PrivateEndpointConnectionListResultPage {
	return original.NewPrivateEndpointConnectionListResultPage(getNextPage)
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
func NewQueryKeysClient(subscriptionID string) QueryKeysClient {
	return original.NewQueryKeysClient(subscriptionID)
}
func NewQueryKeysClientWithBaseURI(baseURI string, subscriptionID string) QueryKeysClient {
	return original.NewQueryKeysClientWithBaseURI(baseURI, subscriptionID)
}
func NewServiceListResultIterator(page ServiceListResultPage) ServiceListResultIterator {
	return original.NewServiceListResultIterator(page)
}
func NewServiceListResultPage(getNextPage func(context.Context, ServiceListResult) (ServiceListResult, error)) ServiceListResultPage {
	return original.NewServiceListResultPage(getNextPage)
}
func NewServicesClient(subscriptionID string) ServicesClient {
	return original.NewServicesClient(subscriptionID)
}
func NewServicesClientWithBaseURI(baseURI string, subscriptionID string) ServicesClient {
	return original.NewServicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewSharedPrivateLinkResourceListResultIterator(page SharedPrivateLinkResourceListResultPage) SharedPrivateLinkResourceListResultIterator {
	return original.NewSharedPrivateLinkResourceListResultIterator(page)
}
func NewSharedPrivateLinkResourceListResultPage(getNextPage func(context.Context, SharedPrivateLinkResourceListResult) (SharedPrivateLinkResourceListResult, error)) SharedPrivateLinkResourceListResultPage {
	return original.NewSharedPrivateLinkResourceListResultPage(getNextPage)
}
func NewSharedPrivateLinkResourcesClient(subscriptionID string) SharedPrivateLinkResourcesClient {
	return original.NewSharedPrivateLinkResourcesClient(subscriptionID)
}
func NewSharedPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) SharedPrivateLinkResourcesClient {
	return original.NewSharedPrivateLinkResourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAdminKeyKindValues() []AdminKeyKind {
	return original.PossibleAdminKeyKindValues()
}
func PossibleHostingModeValues() []HostingMode {
	return original.PossibleHostingModeValues()
}
func PossibleIdentityTypeValues() []IdentityType {
	return original.PossibleIdentityTypeValues()
}
func PossiblePrivateLinkServiceConnectionStatusValues() []PrivateLinkServiceConnectionStatus {
	return original.PossiblePrivateLinkServiceConnectionStatusValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return original.PossiblePublicNetworkAccessValues()
}
func PossibleServiceStatusValues() []ServiceStatus {
	return original.PossibleServiceStatusValues()
}
func PossibleSharedPrivateLinkResourceStatusValues() []SharedPrivateLinkResourceStatus {
	return original.PossibleSharedPrivateLinkResourceStatusValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleUnavailableNameReasonValues() []UnavailableNameReason {
	return original.PossibleUnavailableNameReasonValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
