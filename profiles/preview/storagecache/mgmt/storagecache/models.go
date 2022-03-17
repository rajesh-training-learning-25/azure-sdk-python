//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package storagecache

import (
	"context"

	original "github.com/Azure/temp/github.com/Azure/azure-sdk-for-go/services/storagecache/mgmt/2021-09-01/storagecache"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type CacheIdentityType = original.CacheIdentityType

const (
	CacheIdentityTypeNone                       CacheIdentityType = original.CacheIdentityTypeNone
	CacheIdentityTypeSystemAssigned             CacheIdentityType = original.CacheIdentityTypeSystemAssigned
	CacheIdentityTypeSystemAssignedUserAssigned CacheIdentityType = original.CacheIdentityTypeSystemAssignedUserAssigned
	CacheIdentityTypeUserAssigned               CacheIdentityType = original.CacheIdentityTypeUserAssigned
)

type CreatedByType = original.CreatedByType

const (
	CreatedByTypeApplication     CreatedByType = original.CreatedByTypeApplication
	CreatedByTypeKey             CreatedByType = original.CreatedByTypeKey
	CreatedByTypeManagedIdentity CreatedByType = original.CreatedByTypeManagedIdentity
	CreatedByTypeUser            CreatedByType = original.CreatedByTypeUser
)

type DomainJoinedType = original.DomainJoinedType

const (
	DomainJoinedTypeError DomainJoinedType = original.DomainJoinedTypeError
	DomainJoinedTypeNo    DomainJoinedType = original.DomainJoinedTypeNo
	DomainJoinedTypeYes   DomainJoinedType = original.DomainJoinedTypeYes
)

type FirmwareStatusType = original.FirmwareStatusType

const (
	FirmwareStatusTypeAvailable   FirmwareStatusType = original.FirmwareStatusTypeAvailable
	FirmwareStatusTypeUnavailable FirmwareStatusType = original.FirmwareStatusTypeUnavailable
)

type HealthStateType = original.HealthStateType

const (
	HealthStateTypeDegraded      HealthStateType = original.HealthStateTypeDegraded
	HealthStateTypeDown          HealthStateType = original.HealthStateTypeDown
	HealthStateTypeFlushing      HealthStateType = original.HealthStateTypeFlushing
	HealthStateTypeHealthy       HealthStateType = original.HealthStateTypeHealthy
	HealthStateTypeStopped       HealthStateType = original.HealthStateTypeStopped
	HealthStateTypeStopping      HealthStateType = original.HealthStateTypeStopping
	HealthStateTypeTransitioning HealthStateType = original.HealthStateTypeTransitioning
	HealthStateTypeUnknown       HealthStateType = original.HealthStateTypeUnknown
	HealthStateTypeUpgrading     HealthStateType = original.HealthStateTypeUpgrading
)

type MetricAggregationType = original.MetricAggregationType

const (
	MetricAggregationTypeAverage      MetricAggregationType = original.MetricAggregationTypeAverage
	MetricAggregationTypeCount        MetricAggregationType = original.MetricAggregationTypeCount
	MetricAggregationTypeMaximum      MetricAggregationType = original.MetricAggregationTypeMaximum
	MetricAggregationTypeMinimum      MetricAggregationType = original.MetricAggregationTypeMinimum
	MetricAggregationTypeNone         MetricAggregationType = original.MetricAggregationTypeNone
	MetricAggregationTypeNotSpecified MetricAggregationType = original.MetricAggregationTypeNotSpecified
	MetricAggregationTypeTotal        MetricAggregationType = original.MetricAggregationTypeTotal
)

type NfsAccessRuleAccess = original.NfsAccessRuleAccess

const (
	NfsAccessRuleAccessNo NfsAccessRuleAccess = original.NfsAccessRuleAccessNo
	NfsAccessRuleAccessRo NfsAccessRuleAccess = original.NfsAccessRuleAccessRo
	NfsAccessRuleAccessRw NfsAccessRuleAccess = original.NfsAccessRuleAccessRw
)

type NfsAccessRuleScope = original.NfsAccessRuleScope

const (
	NfsAccessRuleScopeDefault NfsAccessRuleScope = original.NfsAccessRuleScopeDefault
	NfsAccessRuleScopeHost    NfsAccessRuleScope = original.NfsAccessRuleScopeHost
	NfsAccessRuleScopeNetwork NfsAccessRuleScope = original.NfsAccessRuleScopeNetwork
)

type OperationalStateType = original.OperationalStateType

const (
	OperationalStateTypeBusy      OperationalStateType = original.OperationalStateTypeBusy
	OperationalStateTypeFlushing  OperationalStateType = original.OperationalStateTypeFlushing
	OperationalStateTypeReady     OperationalStateType = original.OperationalStateTypeReady
	OperationalStateTypeSuspended OperationalStateType = original.OperationalStateTypeSuspended
)

type ProvisioningStateType = original.ProvisioningStateType

const (
	ProvisioningStateTypeCancelled ProvisioningStateType = original.ProvisioningStateTypeCancelled
	ProvisioningStateTypeCreating  ProvisioningStateType = original.ProvisioningStateTypeCreating
	ProvisioningStateTypeDeleting  ProvisioningStateType = original.ProvisioningStateTypeDeleting
	ProvisioningStateTypeFailed    ProvisioningStateType = original.ProvisioningStateTypeFailed
	ProvisioningStateTypeSucceeded ProvisioningStateType = original.ProvisioningStateTypeSucceeded
	ProvisioningStateTypeUpdating  ProvisioningStateType = original.ProvisioningStateTypeUpdating
)

type ReasonCode = original.ReasonCode

const (
	ReasonCodeNotAvailableForSubscription ReasonCode = original.ReasonCodeNotAvailableForSubscription
	ReasonCodeQuotaID                     ReasonCode = original.ReasonCodeQuotaID
)

type StorageTargetType = original.StorageTargetType

const (
	StorageTargetTypeBlobNfs StorageTargetType = original.StorageTargetTypeBlobNfs
	StorageTargetTypeClfs    StorageTargetType = original.StorageTargetTypeClfs
	StorageTargetTypeNfs3    StorageTargetType = original.StorageTargetTypeNfs3
	StorageTargetTypeUnknown StorageTargetType = original.StorageTargetTypeUnknown
)

type UsernameDownloadedType = original.UsernameDownloadedType

const (
	UsernameDownloadedTypeError UsernameDownloadedType = original.UsernameDownloadedTypeError
	UsernameDownloadedTypeNo    UsernameDownloadedType = original.UsernameDownloadedTypeNo
	UsernameDownloadedTypeYes   UsernameDownloadedType = original.UsernameDownloadedTypeYes
)

type UsernameSource = original.UsernameSource

const (
	UsernameSourceAD   UsernameSource = original.UsernameSourceAD
	UsernameSourceFile UsernameSource = original.UsernameSourceFile
	UsernameSourceLDAP UsernameSource = original.UsernameSourceLDAP
	UsernameSourceNone UsernameSource = original.UsernameSourceNone
)

type APIOperation = original.APIOperation
type APIOperationDisplay = original.APIOperationDisplay
type APIOperationListResult = original.APIOperationListResult
type APIOperationListResultIterator = original.APIOperationListResultIterator
type APIOperationListResultPage = original.APIOperationListResultPage
type APIOperationProperties = original.APIOperationProperties
type APIOperationPropertiesServiceSpecification = original.APIOperationPropertiesServiceSpecification
type AscOperation = original.AscOperation
type AscOperationProperties = original.AscOperationProperties
type AscOperationsClient = original.AscOperationsClient
type BaseClient = original.BaseClient
type BlobNfsTarget = original.BlobNfsTarget
type Cache = original.Cache
type CacheActiveDirectorySettings = original.CacheActiveDirectorySettings
type CacheActiveDirectorySettingsCredentials = original.CacheActiveDirectorySettingsCredentials
type CacheDirectorySettings = original.CacheDirectorySettings
type CacheEncryptionSettings = original.CacheEncryptionSettings
type CacheHealth = original.CacheHealth
type CacheIdentity = original.CacheIdentity
type CacheIdentityUserAssignedIdentitiesValue = original.CacheIdentityUserAssignedIdentitiesValue
type CacheNetworkSettings = original.CacheNetworkSettings
type CacheProperties = original.CacheProperties
type CacheSecuritySettings = original.CacheSecuritySettings
type CacheSku = original.CacheSku
type CacheUpgradeStatus = original.CacheUpgradeStatus
type CacheUsernameDownloadSettings = original.CacheUsernameDownloadSettings
type CacheUsernameDownloadSettingsCredentials = original.CacheUsernameDownloadSettingsCredentials
type CachesClient = original.CachesClient
type CachesCreateOrUpdateFuture = original.CachesCreateOrUpdateFuture
type CachesDebugInfoFuture = original.CachesDebugInfoFuture
type CachesDeleteFuture = original.CachesDeleteFuture
type CachesFlushFuture = original.CachesFlushFuture
type CachesListResult = original.CachesListResult
type CachesListResultIterator = original.CachesListResultIterator
type CachesListResultPage = original.CachesListResultPage
type CachesStartFuture = original.CachesStartFuture
type CachesStopFuture = original.CachesStopFuture
type CachesUpgradeFirmwareFuture = original.CachesUpgradeFirmwareFuture
type ClfsTarget = original.ClfsTarget
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type Condition = original.Condition
type ErrorResponse = original.ErrorResponse
type KeyVaultKeyReference = original.KeyVaultKeyReference
type KeyVaultKeyReferenceSourceVault = original.KeyVaultKeyReferenceSourceVault
type MetricDimension = original.MetricDimension
type MetricSpecification = original.MetricSpecification
type NamespaceJunction = original.NamespaceJunction
type Nfs3Target = original.Nfs3Target
type NfsAccessPolicy = original.NfsAccessPolicy
type NfsAccessRule = original.NfsAccessRule
type OperationsClient = original.OperationsClient
type ResourceSku = original.ResourceSku
type ResourceSkuCapabilities = original.ResourceSkuCapabilities
type ResourceSkuLocationInfo = original.ResourceSkuLocationInfo
type ResourceSkusResult = original.ResourceSkusResult
type ResourceSkusResultIterator = original.ResourceSkusResultIterator
type ResourceSkusResultPage = original.ResourceSkusResultPage
type Restriction = original.Restriction
type SkusClient = original.SkusClient
type StorageTarget = original.StorageTarget
type StorageTargetClient = original.StorageTargetClient
type StorageTargetFlushFuture = original.StorageTargetFlushFuture
type StorageTargetProperties = original.StorageTargetProperties
type StorageTargetResource = original.StorageTargetResource
type StorageTargetResumeFuture = original.StorageTargetResumeFuture
type StorageTargetSuspendFuture = original.StorageTargetSuspendFuture
type StorageTargetsClient = original.StorageTargetsClient
type StorageTargetsCreateOrUpdateFuture = original.StorageTargetsCreateOrUpdateFuture
type StorageTargetsDNSRefreshFuture = original.StorageTargetsDNSRefreshFuture
type StorageTargetsDeleteFuture = original.StorageTargetsDeleteFuture
type StorageTargetsResult = original.StorageTargetsResult
type StorageTargetsResultIterator = original.StorageTargetsResultIterator
type StorageTargetsResultPage = original.StorageTargetsResultPage
type SystemData = original.SystemData
type UnknownTarget = original.UnknownTarget
type UsageModel = original.UsageModel
type UsageModelDisplay = original.UsageModelDisplay
type UsageModelsClient = original.UsageModelsClient
type UsageModelsResult = original.UsageModelsResult
type UsageModelsResultIterator = original.UsageModelsResultIterator
type UsageModelsResultPage = original.UsageModelsResultPage

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAPIOperationListResultIterator(page APIOperationListResultPage) APIOperationListResultIterator {
	return original.NewAPIOperationListResultIterator(page)
}
func NewAPIOperationListResultPage(cur APIOperationListResult, getNextPage func(context.Context, APIOperationListResult) (APIOperationListResult, error)) APIOperationListResultPage {
	return original.NewAPIOperationListResultPage(cur, getNextPage)
}
func NewAscOperationsClient(subscriptionID string) AscOperationsClient {
	return original.NewAscOperationsClient(subscriptionID)
}
func NewAscOperationsClientWithBaseURI(baseURI string, subscriptionID string) AscOperationsClient {
	return original.NewAscOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewCachesClient(subscriptionID string) CachesClient {
	return original.NewCachesClient(subscriptionID)
}
func NewCachesClientWithBaseURI(baseURI string, subscriptionID string) CachesClient {
	return original.NewCachesClientWithBaseURI(baseURI, subscriptionID)
}
func NewCachesListResultIterator(page CachesListResultPage) CachesListResultIterator {
	return original.NewCachesListResultIterator(page)
}
func NewCachesListResultPage(cur CachesListResult, getNextPage func(context.Context, CachesListResult) (CachesListResult, error)) CachesListResultPage {
	return original.NewCachesListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceSkusResultIterator(page ResourceSkusResultPage) ResourceSkusResultIterator {
	return original.NewResourceSkusResultIterator(page)
}
func NewResourceSkusResultPage(cur ResourceSkusResult, getNextPage func(context.Context, ResourceSkusResult) (ResourceSkusResult, error)) ResourceSkusResultPage {
	return original.NewResourceSkusResultPage(cur, getNextPage)
}
func NewSkusClient(subscriptionID string) SkusClient {
	return original.NewSkusClient(subscriptionID)
}
func NewSkusClientWithBaseURI(baseURI string, subscriptionID string) SkusClient {
	return original.NewSkusClientWithBaseURI(baseURI, subscriptionID)
}
func NewStorageTargetClient(subscriptionID string) StorageTargetClient {
	return original.NewStorageTargetClient(subscriptionID)
}
func NewStorageTargetClientWithBaseURI(baseURI string, subscriptionID string) StorageTargetClient {
	return original.NewStorageTargetClientWithBaseURI(baseURI, subscriptionID)
}
func NewStorageTargetsClient(subscriptionID string) StorageTargetsClient {
	return original.NewStorageTargetsClient(subscriptionID)
}
func NewStorageTargetsClientWithBaseURI(baseURI string, subscriptionID string) StorageTargetsClient {
	return original.NewStorageTargetsClientWithBaseURI(baseURI, subscriptionID)
}
func NewStorageTargetsResultIterator(page StorageTargetsResultPage) StorageTargetsResultIterator {
	return original.NewStorageTargetsResultIterator(page)
}
func NewStorageTargetsResultPage(cur StorageTargetsResult, getNextPage func(context.Context, StorageTargetsResult) (StorageTargetsResult, error)) StorageTargetsResultPage {
	return original.NewStorageTargetsResultPage(cur, getNextPage)
}
func NewUsageModelsClient(subscriptionID string) UsageModelsClient {
	return original.NewUsageModelsClient(subscriptionID)
}
func NewUsageModelsClientWithBaseURI(baseURI string, subscriptionID string) UsageModelsClient {
	return original.NewUsageModelsClientWithBaseURI(baseURI, subscriptionID)
}
func NewUsageModelsResultIterator(page UsageModelsResultPage) UsageModelsResultIterator {
	return original.NewUsageModelsResultIterator(page)
}
func NewUsageModelsResultPage(cur UsageModelsResult, getNextPage func(context.Context, UsageModelsResult) (UsageModelsResult, error)) UsageModelsResultPage {
	return original.NewUsageModelsResultPage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleCacheIdentityTypeValues() []CacheIdentityType {
	return original.PossibleCacheIdentityTypeValues()
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleDomainJoinedTypeValues() []DomainJoinedType {
	return original.PossibleDomainJoinedTypeValues()
}
func PossibleFirmwareStatusTypeValues() []FirmwareStatusType {
	return original.PossibleFirmwareStatusTypeValues()
}
func PossibleHealthStateTypeValues() []HealthStateType {
	return original.PossibleHealthStateTypeValues()
}
func PossibleMetricAggregationTypeValues() []MetricAggregationType {
	return original.PossibleMetricAggregationTypeValues()
}
func PossibleNfsAccessRuleAccessValues() []NfsAccessRuleAccess {
	return original.PossibleNfsAccessRuleAccessValues()
}
func PossibleNfsAccessRuleScopeValues() []NfsAccessRuleScope {
	return original.PossibleNfsAccessRuleScopeValues()
}
func PossibleOperationalStateTypeValues() []OperationalStateType {
	return original.PossibleOperationalStateTypeValues()
}
func PossibleProvisioningStateTypeValues() []ProvisioningStateType {
	return original.PossibleProvisioningStateTypeValues()
}
func PossibleReasonCodeValues() []ReasonCode {
	return original.PossibleReasonCodeValues()
}
func PossibleStorageTargetTypeValues() []StorageTargetType {
	return original.PossibleStorageTargetTypeValues()
}
func PossibleUsernameDownloadedTypeValues() []UsernameDownloadedType {
	return original.PossibleUsernameDownloadedTypeValues()
}
func PossibleUsernameSourceValues() []UsernameSource {
	return original.PossibleUsernameSourceValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
