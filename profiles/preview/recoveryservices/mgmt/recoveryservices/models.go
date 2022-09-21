//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package recoveryservices

import (
	"context"

	original "github.com/Azure/dev/azure-sdk-for-go/services/recoveryservices/mgmt/2021-08-01/recoveryservices"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AuthType = original.AuthType

const (
	AuthTypeAAD                  AuthType = original.AuthTypeAAD
	AuthTypeAccessControlService AuthType = original.AuthTypeAccessControlService
	AuthTypeACS                  AuthType = original.AuthTypeACS
	AuthTypeAzureActiveDirectory AuthType = original.AuthTypeAzureActiveDirectory
	AuthTypeInvalid              AuthType = original.AuthTypeInvalid
)

type AuthTypeBasicResourceCertificateDetails = original.AuthTypeBasicResourceCertificateDetails

const (
	AuthTypeBasicResourceCertificateDetailsAuthTypeAccessControlService       AuthTypeBasicResourceCertificateDetails = original.AuthTypeBasicResourceCertificateDetailsAuthTypeAccessControlService
	AuthTypeBasicResourceCertificateDetailsAuthTypeAzureActiveDirectory       AuthTypeBasicResourceCertificateDetails = original.AuthTypeBasicResourceCertificateDetailsAuthTypeAzureActiveDirectory
	AuthTypeBasicResourceCertificateDetailsAuthTypeResourceCertificateDetails AuthTypeBasicResourceCertificateDetails = original.AuthTypeBasicResourceCertificateDetailsAuthTypeResourceCertificateDetails
)

type CreatedByType = original.CreatedByType

const (
	CreatedByTypeApplication     CreatedByType = original.CreatedByTypeApplication
	CreatedByTypeKey             CreatedByType = original.CreatedByTypeKey
	CreatedByTypeManagedIdentity CreatedByType = original.CreatedByTypeManagedIdentity
	CreatedByTypeUser            CreatedByType = original.CreatedByTypeUser
)

type InfrastructureEncryptionState = original.InfrastructureEncryptionState

const (
	InfrastructureEncryptionStateDisabled InfrastructureEncryptionState = original.InfrastructureEncryptionStateDisabled
	InfrastructureEncryptionStateEnabled  InfrastructureEncryptionState = original.InfrastructureEncryptionStateEnabled
)

type PrivateEndpointConnectionStatus = original.PrivateEndpointConnectionStatus

const (
	PrivateEndpointConnectionStatusApproved     PrivateEndpointConnectionStatus = original.PrivateEndpointConnectionStatusApproved
	PrivateEndpointConnectionStatusDisconnected PrivateEndpointConnectionStatus = original.PrivateEndpointConnectionStatusDisconnected
	PrivateEndpointConnectionStatusPending      PrivateEndpointConnectionStatus = original.PrivateEndpointConnectionStatusPending
	PrivateEndpointConnectionStatusRejected     PrivateEndpointConnectionStatus = original.PrivateEndpointConnectionStatusRejected
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateDeleting  ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateFailed    ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStatePending   ProvisioningState = original.ProvisioningStatePending
	ProvisioningStateSucceeded ProvisioningState = original.ProvisioningStateSucceeded
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	ResourceIdentityTypeNone                       ResourceIdentityType = original.ResourceIdentityTypeNone
	ResourceIdentityTypeSystemAssigned             ResourceIdentityType = original.ResourceIdentityTypeSystemAssigned
	ResourceIdentityTypeSystemAssignedUserAssigned ResourceIdentityType = original.ResourceIdentityTypeSystemAssignedUserAssigned
	ResourceIdentityTypeUserAssigned               ResourceIdentityType = original.ResourceIdentityTypeUserAssigned
)

type ResourceMoveState = original.ResourceMoveState

const (
	ResourceMoveStateCommitFailed    ResourceMoveState = original.ResourceMoveStateCommitFailed
	ResourceMoveStateCommitTimedout  ResourceMoveState = original.ResourceMoveStateCommitTimedout
	ResourceMoveStateCriticalFailure ResourceMoveState = original.ResourceMoveStateCriticalFailure
	ResourceMoveStateFailure         ResourceMoveState = original.ResourceMoveStateFailure
	ResourceMoveStateInProgress      ResourceMoveState = original.ResourceMoveStateInProgress
	ResourceMoveStateMoveSucceeded   ResourceMoveState = original.ResourceMoveStateMoveSucceeded
	ResourceMoveStatePartialSuccess  ResourceMoveState = original.ResourceMoveStatePartialSuccess
	ResourceMoveStatePrepareFailed   ResourceMoveState = original.ResourceMoveStatePrepareFailed
	ResourceMoveStatePrepareTimedout ResourceMoveState = original.ResourceMoveStatePrepareTimedout
	ResourceMoveStateUnknown         ResourceMoveState = original.ResourceMoveStateUnknown
)

type SkuName = original.SkuName

const (
	SkuNameRS0      SkuName = original.SkuNameRS0
	SkuNameStandard SkuName = original.SkuNameStandard
)

type TriggerType = original.TriggerType

const (
	TriggerTypeForcedUpgrade TriggerType = original.TriggerTypeForcedUpgrade
	TriggerTypeUserTriggered TriggerType = original.TriggerTypeUserTriggered
)

type UsagesUnit = original.UsagesUnit

const (
	UsagesUnitBytes          UsagesUnit = original.UsagesUnitBytes
	UsagesUnitBytesPerSecond UsagesUnit = original.UsagesUnitBytesPerSecond
	UsagesUnitCount          UsagesUnit = original.UsagesUnitCount
	UsagesUnitCountPerSecond UsagesUnit = original.UsagesUnitCountPerSecond
	UsagesUnitPercent        UsagesUnit = original.UsagesUnitPercent
	UsagesUnitSeconds        UsagesUnit = original.UsagesUnitSeconds
)

type VaultPrivateEndpointState = original.VaultPrivateEndpointState

const (
	VaultPrivateEndpointStateEnabled VaultPrivateEndpointState = original.VaultPrivateEndpointStateEnabled
	VaultPrivateEndpointStateNone    VaultPrivateEndpointState = original.VaultPrivateEndpointStateNone
)

type VaultUpgradeState = original.VaultUpgradeState

const (
	VaultUpgradeStateFailed     VaultUpgradeState = original.VaultUpgradeStateFailed
	VaultUpgradeStateInProgress VaultUpgradeState = original.VaultUpgradeStateInProgress
	VaultUpgradeStateUnknown    VaultUpgradeState = original.VaultUpgradeStateUnknown
	VaultUpgradeStateUpgraded   VaultUpgradeState = original.VaultUpgradeStateUpgraded
)

type BaseClient = original.BaseClient
type BasicResourceCertificateDetails = original.BasicResourceCertificateDetails
type CertificateRequest = original.CertificateRequest
type CheckNameAvailabilityParameters = original.CheckNameAvailabilityParameters
type CheckNameAvailabilityResult = original.CheckNameAvailabilityResult
type Client = original.Client
type ClientDiscoveryDisplay = original.ClientDiscoveryDisplay
type ClientDiscoveryForLogSpecification = original.ClientDiscoveryForLogSpecification
type ClientDiscoveryForProperties = original.ClientDiscoveryForProperties
type ClientDiscoveryForServiceSpecification = original.ClientDiscoveryForServiceSpecification
type ClientDiscoveryResponse = original.ClientDiscoveryResponse
type ClientDiscoveryResponseIterator = original.ClientDiscoveryResponseIterator
type ClientDiscoveryResponsePage = original.ClientDiscoveryResponsePage
type ClientDiscoveryValueForSingleAPI = original.ClientDiscoveryValueForSingleAPI
type CloudError = original.CloudError
type CmkKekIdentity = original.CmkKekIdentity
type CmkKeyVaultProperties = original.CmkKeyVaultProperties
type Error = original.Error
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type IdentityData = original.IdentityData
type JobsSummary = original.JobsSummary
type MonitoringSummary = original.MonitoringSummary
type NameInfo = original.NameInfo
type OperationResource = original.OperationResource
type OperationsClient = original.OperationsClient
type PatchTrackedResource = original.PatchTrackedResource
type PatchVault = original.PatchVault
type PrivateEndpoint = original.PrivateEndpoint
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionVaultProperties = original.PrivateEndpointConnectionVaultProperties
type PrivateLinkResource = original.PrivateLinkResource
type PrivateLinkResourceProperties = original.PrivateLinkResourceProperties
type PrivateLinkResources = original.PrivateLinkResources
type PrivateLinkResourcesClient = original.PrivateLinkResourcesClient
type PrivateLinkResourcesIterator = original.PrivateLinkResourcesIterator
type PrivateLinkResourcesPage = original.PrivateLinkResourcesPage
type PrivateLinkServiceConnectionState = original.PrivateLinkServiceConnectionState
type RawCertificateData = original.RawCertificateData
type RegisteredIdentitiesClient = original.RegisteredIdentitiesClient
type ReplicationUsage = original.ReplicationUsage
type ReplicationUsageList = original.ReplicationUsageList
type ReplicationUsagesClient = original.ReplicationUsagesClient
type Resource = original.Resource
type ResourceCertificateAndAadDetails = original.ResourceCertificateAndAadDetails
type ResourceCertificateAndAcsDetails = original.ResourceCertificateAndAcsDetails
type ResourceCertificateDetails = original.ResourceCertificateDetails
type Sku = original.Sku
type SystemData = original.SystemData
type TrackedResource = original.TrackedResource
type UpgradeDetails = original.UpgradeDetails
type UsagesClient = original.UsagesClient
type UserIdentity = original.UserIdentity
type Vault = original.Vault
type VaultCertificateResponse = original.VaultCertificateResponse
type VaultCertificatesClient = original.VaultCertificatesClient
type VaultExtendedInfo = original.VaultExtendedInfo
type VaultExtendedInfoClient = original.VaultExtendedInfoClient
type VaultExtendedInfoResource = original.VaultExtendedInfoResource
type VaultList = original.VaultList
type VaultListIterator = original.VaultListIterator
type VaultListPage = original.VaultListPage
type VaultProperties = original.VaultProperties
type VaultPropertiesEncryption = original.VaultPropertiesEncryption
type VaultPropertiesMoveDetails = original.VaultPropertiesMoveDetails
type VaultUsage = original.VaultUsage
type VaultUsageList = original.VaultUsageList
type VaultsClient = original.VaultsClient
type VaultsCreateOrUpdateFuture = original.VaultsCreateOrUpdateFuture
type VaultsUpdateFuture = original.VaultsUpdateFuture

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewClient(subscriptionID string) Client {
	return original.NewClient(subscriptionID)
}
func NewClientDiscoveryResponseIterator(page ClientDiscoveryResponsePage) ClientDiscoveryResponseIterator {
	return original.NewClientDiscoveryResponseIterator(page)
}
func NewClientDiscoveryResponsePage(cur ClientDiscoveryResponse, getNextPage func(context.Context, ClientDiscoveryResponse) (ClientDiscoveryResponse, error)) ClientDiscoveryResponsePage {
	return original.NewClientDiscoveryResponsePage(cur, getNextPage)
}
func NewClientWithBaseURI(baseURI string, subscriptionID string) Client {
	return original.NewClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateLinkResourcesClient(subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClient(subscriptionID)
}
func NewPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateLinkResourcesIterator(page PrivateLinkResourcesPage) PrivateLinkResourcesIterator {
	return original.NewPrivateLinkResourcesIterator(page)
}
func NewPrivateLinkResourcesPage(cur PrivateLinkResources, getNextPage func(context.Context, PrivateLinkResources) (PrivateLinkResources, error)) PrivateLinkResourcesPage {
	return original.NewPrivateLinkResourcesPage(cur, getNextPage)
}
func NewRegisteredIdentitiesClient(subscriptionID string) RegisteredIdentitiesClient {
	return original.NewRegisteredIdentitiesClient(subscriptionID)
}
func NewRegisteredIdentitiesClientWithBaseURI(baseURI string, subscriptionID string) RegisteredIdentitiesClient {
	return original.NewRegisteredIdentitiesClientWithBaseURI(baseURI, subscriptionID)
}
func NewReplicationUsagesClient(subscriptionID string) ReplicationUsagesClient {
	return original.NewReplicationUsagesClient(subscriptionID)
}
func NewReplicationUsagesClientWithBaseURI(baseURI string, subscriptionID string) ReplicationUsagesClient {
	return original.NewReplicationUsagesClientWithBaseURI(baseURI, subscriptionID)
}
func NewUsagesClient(subscriptionID string) UsagesClient {
	return original.NewUsagesClient(subscriptionID)
}
func NewUsagesClientWithBaseURI(baseURI string, subscriptionID string) UsagesClient {
	return original.NewUsagesClientWithBaseURI(baseURI, subscriptionID)
}
func NewVaultCertificatesClient(subscriptionID string) VaultCertificatesClient {
	return original.NewVaultCertificatesClient(subscriptionID)
}
func NewVaultCertificatesClientWithBaseURI(baseURI string, subscriptionID string) VaultCertificatesClient {
	return original.NewVaultCertificatesClientWithBaseURI(baseURI, subscriptionID)
}
func NewVaultExtendedInfoClient(subscriptionID string) VaultExtendedInfoClient {
	return original.NewVaultExtendedInfoClient(subscriptionID)
}
func NewVaultExtendedInfoClientWithBaseURI(baseURI string, subscriptionID string) VaultExtendedInfoClient {
	return original.NewVaultExtendedInfoClientWithBaseURI(baseURI, subscriptionID)
}
func NewVaultListIterator(page VaultListPage) VaultListIterator {
	return original.NewVaultListIterator(page)
}
func NewVaultListPage(cur VaultList, getNextPage func(context.Context, VaultList) (VaultList, error)) VaultListPage {
	return original.NewVaultListPage(cur, getNextPage)
}
func NewVaultsClient(subscriptionID string) VaultsClient {
	return original.NewVaultsClient(subscriptionID)
}
func NewVaultsClientWithBaseURI(baseURI string, subscriptionID string) VaultsClient {
	return original.NewVaultsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAuthTypeBasicResourceCertificateDetailsValues() []AuthTypeBasicResourceCertificateDetails {
	return original.PossibleAuthTypeBasicResourceCertificateDetailsValues()
}
func PossibleAuthTypeValues() []AuthType {
	return original.PossibleAuthTypeValues()
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleInfrastructureEncryptionStateValues() []InfrastructureEncryptionState {
	return original.PossibleInfrastructureEncryptionStateValues()
}
func PossiblePrivateEndpointConnectionStatusValues() []PrivateEndpointConnectionStatus {
	return original.PossiblePrivateEndpointConnectionStatusValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func PossibleResourceMoveStateValues() []ResourceMoveState {
	return original.PossibleResourceMoveStateValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleTriggerTypeValues() []TriggerType {
	return original.PossibleTriggerTypeValues()
}
func PossibleUsagesUnitValues() []UsagesUnit {
	return original.PossibleUsagesUnitValues()
}
func PossibleVaultPrivateEndpointStateValues() []VaultPrivateEndpointState {
	return original.PossibleVaultPrivateEndpointStateValues()
}
func PossibleVaultUpgradeStateValues() []VaultUpgradeState {
	return original.PossibleVaultUpgradeStateValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
