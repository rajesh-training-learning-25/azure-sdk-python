//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package servicefabric

import (
	"context"

	original "github.com/Azure/temp/github.com/Azure/azure-sdk-for-go/services/servicefabric/mgmt/2021-06-01/servicefabric"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ArmServicePackageActivationMode = original.ArmServicePackageActivationMode

const (
	ArmServicePackageActivationModeExclusiveProcess ArmServicePackageActivationMode = original.ArmServicePackageActivationModeExclusiveProcess
	ArmServicePackageActivationModeSharedProcess    ArmServicePackageActivationMode = original.ArmServicePackageActivationModeSharedProcess
)

type ArmUpgradeFailureAction = original.ArmUpgradeFailureAction

const (
	ArmUpgradeFailureActionManual   ArmUpgradeFailureAction = original.ArmUpgradeFailureActionManual
	ArmUpgradeFailureActionRollback ArmUpgradeFailureAction = original.ArmUpgradeFailureActionRollback
)

type ClusterState = original.ClusterState

const (
	ClusterStateAutoScale                 ClusterState = original.ClusterStateAutoScale
	ClusterStateBaselineUpgrade           ClusterState = original.ClusterStateBaselineUpgrade
	ClusterStateDeploying                 ClusterState = original.ClusterStateDeploying
	ClusterStateEnforcingClusterVersion   ClusterState = original.ClusterStateEnforcingClusterVersion
	ClusterStateReady                     ClusterState = original.ClusterStateReady
	ClusterStateUpdatingInfrastructure    ClusterState = original.ClusterStateUpdatingInfrastructure
	ClusterStateUpdatingUserCertificate   ClusterState = original.ClusterStateUpdatingUserCertificate
	ClusterStateUpdatingUserConfiguration ClusterState = original.ClusterStateUpdatingUserConfiguration
	ClusterStateUpgradeServiceUnreachable ClusterState = original.ClusterStateUpgradeServiceUnreachable
	ClusterStateWaitingForNodes           ClusterState = original.ClusterStateWaitingForNodes
)

type ClusterUpgradeCadence = original.ClusterUpgradeCadence

const (
	ClusterUpgradeCadenceWave0 ClusterUpgradeCadence = original.ClusterUpgradeCadenceWave0
	ClusterUpgradeCadenceWave1 ClusterUpgradeCadence = original.ClusterUpgradeCadenceWave1
	ClusterUpgradeCadenceWave2 ClusterUpgradeCadence = original.ClusterUpgradeCadenceWave2
)

type DurabilityLevel = original.DurabilityLevel

const (
	DurabilityLevelBronze DurabilityLevel = original.DurabilityLevelBronze
	DurabilityLevelGold   DurabilityLevel = original.DurabilityLevelGold
	DurabilityLevelSilver DurabilityLevel = original.DurabilityLevelSilver
)

type Environment = original.Environment

const (
	EnvironmentLinux   Environment = original.EnvironmentLinux
	EnvironmentWindows Environment = original.EnvironmentWindows
)

type ManagedIdentityType = original.ManagedIdentityType

const (
	ManagedIdentityTypeNone                       ManagedIdentityType = original.ManagedIdentityTypeNone
	ManagedIdentityTypeSystemAssigned             ManagedIdentityType = original.ManagedIdentityTypeSystemAssigned
	ManagedIdentityTypeSystemAssignedUserAssigned ManagedIdentityType = original.ManagedIdentityTypeSystemAssignedUserAssigned
	ManagedIdentityTypeUserAssigned               ManagedIdentityType = original.ManagedIdentityTypeUserAssigned
)

type MoveCost = original.MoveCost

const (
	MoveCostHigh   MoveCost = original.MoveCostHigh
	MoveCostLow    MoveCost = original.MoveCostLow
	MoveCostMedium MoveCost = original.MoveCostMedium
	MoveCostZero   MoveCost = original.MoveCostZero
)

type NotificationChannel = original.NotificationChannel

const (
	NotificationChannelEmailSubscription NotificationChannel = original.NotificationChannelEmailSubscription
	NotificationChannelEmailUser         NotificationChannel = original.NotificationChannelEmailUser
)

type NotificationLevel = original.NotificationLevel

const (
	NotificationLevelAll      NotificationLevel = original.NotificationLevelAll
	NotificationLevelCritical NotificationLevel = original.NotificationLevelCritical
)

type PartitionScheme = original.PartitionScheme

const (
	PartitionSchemeInvalid           PartitionScheme = original.PartitionSchemeInvalid
	PartitionSchemeNamed             PartitionScheme = original.PartitionSchemeNamed
	PartitionSchemeSingleton         PartitionScheme = original.PartitionSchemeSingleton
	PartitionSchemeUniformInt64Range PartitionScheme = original.PartitionSchemeUniformInt64Range
)

type PartitionSchemeBasicPartitionSchemeDescription = original.PartitionSchemeBasicPartitionSchemeDescription

const (
	PartitionSchemeBasicPartitionSchemeDescriptionPartitionSchemeNamed                      PartitionSchemeBasicPartitionSchemeDescription = original.PartitionSchemeBasicPartitionSchemeDescriptionPartitionSchemeNamed
	PartitionSchemeBasicPartitionSchemeDescriptionPartitionSchemePartitionSchemeDescription PartitionSchemeBasicPartitionSchemeDescription = original.PartitionSchemeBasicPartitionSchemeDescriptionPartitionSchemePartitionSchemeDescription
	PartitionSchemeBasicPartitionSchemeDescriptionPartitionSchemeSingleton                  PartitionSchemeBasicPartitionSchemeDescription = original.PartitionSchemeBasicPartitionSchemeDescriptionPartitionSchemeSingleton
	PartitionSchemeBasicPartitionSchemeDescriptionPartitionSchemeUniformInt64Range          PartitionSchemeBasicPartitionSchemeDescription = original.PartitionSchemeBasicPartitionSchemeDescriptionPartitionSchemeUniformInt64Range
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateCanceled  ProvisioningState = original.ProvisioningStateCanceled
	ProvisioningStateFailed    ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateSucceeded ProvisioningState = original.ProvisioningStateSucceeded
	ProvisioningStateUpdating  ProvisioningState = original.ProvisioningStateUpdating
)

type ReliabilityLevel = original.ReliabilityLevel

const (
	ReliabilityLevelBronze   ReliabilityLevel = original.ReliabilityLevelBronze
	ReliabilityLevelGold     ReliabilityLevel = original.ReliabilityLevelGold
	ReliabilityLevelNone     ReliabilityLevel = original.ReliabilityLevelNone
	ReliabilityLevelPlatinum ReliabilityLevel = original.ReliabilityLevelPlatinum
	ReliabilityLevelSilver   ReliabilityLevel = original.ReliabilityLevelSilver
)

type ReliabilityLevel1 = original.ReliabilityLevel1

const (
	ReliabilityLevel1Bronze   ReliabilityLevel1 = original.ReliabilityLevel1Bronze
	ReliabilityLevel1Gold     ReliabilityLevel1 = original.ReliabilityLevel1Gold
	ReliabilityLevel1None     ReliabilityLevel1 = original.ReliabilityLevel1None
	ReliabilityLevel1Platinum ReliabilityLevel1 = original.ReliabilityLevel1Platinum
	ReliabilityLevel1Silver   ReliabilityLevel1 = original.ReliabilityLevel1Silver
)

type RollingUpgradeMode = original.RollingUpgradeMode

const (
	RollingUpgradeModeInvalid           RollingUpgradeMode = original.RollingUpgradeModeInvalid
	RollingUpgradeModeMonitored         RollingUpgradeMode = original.RollingUpgradeModeMonitored
	RollingUpgradeModeUnmonitoredAuto   RollingUpgradeMode = original.RollingUpgradeModeUnmonitoredAuto
	RollingUpgradeModeUnmonitoredManual RollingUpgradeMode = original.RollingUpgradeModeUnmonitoredManual
)

type ServiceCorrelationScheme = original.ServiceCorrelationScheme

const (
	ServiceCorrelationSchemeAffinity           ServiceCorrelationScheme = original.ServiceCorrelationSchemeAffinity
	ServiceCorrelationSchemeAlignedAffinity    ServiceCorrelationScheme = original.ServiceCorrelationSchemeAlignedAffinity
	ServiceCorrelationSchemeInvalid            ServiceCorrelationScheme = original.ServiceCorrelationSchemeInvalid
	ServiceCorrelationSchemeNonAlignedAffinity ServiceCorrelationScheme = original.ServiceCorrelationSchemeNonAlignedAffinity
)

type ServiceKind = original.ServiceKind

const (
	ServiceKindInvalid   ServiceKind = original.ServiceKindInvalid
	ServiceKindStateful  ServiceKind = original.ServiceKindStateful
	ServiceKindStateless ServiceKind = original.ServiceKindStateless
)

type ServiceKindBasicServiceResourceProperties = original.ServiceKindBasicServiceResourceProperties

const (
	ServiceKindBasicServiceResourcePropertiesServiceKindServiceResourceProperties ServiceKindBasicServiceResourceProperties = original.ServiceKindBasicServiceResourcePropertiesServiceKindServiceResourceProperties
	ServiceKindBasicServiceResourcePropertiesServiceKindStateful                  ServiceKindBasicServiceResourceProperties = original.ServiceKindBasicServiceResourcePropertiesServiceKindStateful
	ServiceKindBasicServiceResourcePropertiesServiceKindStateless                 ServiceKindBasicServiceResourceProperties = original.ServiceKindBasicServiceResourcePropertiesServiceKindStateless
)

type ServiceKindBasicServiceResourceUpdateProperties = original.ServiceKindBasicServiceResourceUpdateProperties

const (
	ServiceKindBasicServiceResourceUpdatePropertiesServiceKindServiceResourceUpdateProperties ServiceKindBasicServiceResourceUpdateProperties = original.ServiceKindBasicServiceResourceUpdatePropertiesServiceKindServiceResourceUpdateProperties
	ServiceKindBasicServiceResourceUpdatePropertiesServiceKindStateful                        ServiceKindBasicServiceResourceUpdateProperties = original.ServiceKindBasicServiceResourceUpdatePropertiesServiceKindStateful
	ServiceKindBasicServiceResourceUpdatePropertiesServiceKindStateless                       ServiceKindBasicServiceResourceUpdateProperties = original.ServiceKindBasicServiceResourceUpdatePropertiesServiceKindStateless
)

type ServiceLoadMetricWeight = original.ServiceLoadMetricWeight

const (
	ServiceLoadMetricWeightHigh   ServiceLoadMetricWeight = original.ServiceLoadMetricWeightHigh
	ServiceLoadMetricWeightLow    ServiceLoadMetricWeight = original.ServiceLoadMetricWeightLow
	ServiceLoadMetricWeightMedium ServiceLoadMetricWeight = original.ServiceLoadMetricWeightMedium
	ServiceLoadMetricWeightZero   ServiceLoadMetricWeight = original.ServiceLoadMetricWeightZero
)

type ServicePlacementPolicyType = original.ServicePlacementPolicyType

const (
	ServicePlacementPolicyTypeInvalid                    ServicePlacementPolicyType = original.ServicePlacementPolicyTypeInvalid
	ServicePlacementPolicyTypeInvalidDomain              ServicePlacementPolicyType = original.ServicePlacementPolicyTypeInvalidDomain
	ServicePlacementPolicyTypeNonPartiallyPlaceService   ServicePlacementPolicyType = original.ServicePlacementPolicyTypeNonPartiallyPlaceService
	ServicePlacementPolicyTypePreferredPrimaryDomain     ServicePlacementPolicyType = original.ServicePlacementPolicyTypePreferredPrimaryDomain
	ServicePlacementPolicyTypeRequiredDomain             ServicePlacementPolicyType = original.ServicePlacementPolicyTypeRequiredDomain
	ServicePlacementPolicyTypeRequiredDomainDistribution ServicePlacementPolicyType = original.ServicePlacementPolicyTypeRequiredDomainDistribution
)

type SfZonalUpgradeMode = original.SfZonalUpgradeMode

const (
	SfZonalUpgradeModeHierarchical SfZonalUpgradeMode = original.SfZonalUpgradeModeHierarchical
	SfZonalUpgradeModeParallel     SfZonalUpgradeMode = original.SfZonalUpgradeModeParallel
)

type Type = original.Type

const (
	TypeServicePlacementPolicyDescription Type = original.TypeServicePlacementPolicyDescription
)

type UpgradeMode = original.UpgradeMode

const (
	UpgradeModeAutomatic UpgradeMode = original.UpgradeModeAutomatic
	UpgradeModeManual    UpgradeMode = original.UpgradeModeManual
)

type VmssZonalUpgradeMode = original.VmssZonalUpgradeMode

const (
	VmssZonalUpgradeModeHierarchical VmssZonalUpgradeMode = original.VmssZonalUpgradeModeHierarchical
	VmssZonalUpgradeModeParallel     VmssZonalUpgradeMode = original.VmssZonalUpgradeModeParallel
)

type X509StoreName = original.X509StoreName

const (
	X509StoreNameAddressBook          X509StoreName = original.X509StoreNameAddressBook
	X509StoreNameAuthRoot             X509StoreName = original.X509StoreNameAuthRoot
	X509StoreNameCertificateAuthority X509StoreName = original.X509StoreNameCertificateAuthority
	X509StoreNameDisallowed           X509StoreName = original.X509StoreNameDisallowed
	X509StoreNameMy                   X509StoreName = original.X509StoreNameMy
	X509StoreNameRoot                 X509StoreName = original.X509StoreNameRoot
	X509StoreNameTrustedPeople        X509StoreName = original.X509StoreNameTrustedPeople
	X509StoreNameTrustedPublisher     X509StoreName = original.X509StoreNameTrustedPublisher
)

type X509StoreName1 = original.X509StoreName1

const (
	X509StoreName1AddressBook          X509StoreName1 = original.X509StoreName1AddressBook
	X509StoreName1AuthRoot             X509StoreName1 = original.X509StoreName1AuthRoot
	X509StoreName1CertificateAuthority X509StoreName1 = original.X509StoreName1CertificateAuthority
	X509StoreName1Disallowed           X509StoreName1 = original.X509StoreName1Disallowed
	X509StoreName1My                   X509StoreName1 = original.X509StoreName1My
	X509StoreName1Root                 X509StoreName1 = original.X509StoreName1Root
	X509StoreName1TrustedPeople        X509StoreName1 = original.X509StoreName1TrustedPeople
	X509StoreName1TrustedPublisher     X509StoreName1 = original.X509StoreName1TrustedPublisher
)

type ApplicationDeltaHealthPolicy = original.ApplicationDeltaHealthPolicy
type ApplicationHealthPolicy = original.ApplicationHealthPolicy
type ApplicationMetricDescription = original.ApplicationMetricDescription
type ApplicationResource = original.ApplicationResource
type ApplicationResourceList = original.ApplicationResourceList
type ApplicationResourceProperties = original.ApplicationResourceProperties
type ApplicationResourceUpdate = original.ApplicationResourceUpdate
type ApplicationResourceUpdateProperties = original.ApplicationResourceUpdateProperties
type ApplicationTypeResource = original.ApplicationTypeResource
type ApplicationTypeResourceList = original.ApplicationTypeResourceList
type ApplicationTypeResourceProperties = original.ApplicationTypeResourceProperties
type ApplicationTypeVersionResource = original.ApplicationTypeVersionResource
type ApplicationTypeVersionResourceList = original.ApplicationTypeVersionResourceList
type ApplicationTypeVersionResourceProperties = original.ApplicationTypeVersionResourceProperties
type ApplicationTypeVersionsCleanupPolicy = original.ApplicationTypeVersionsCleanupPolicy
type ApplicationTypeVersionsClient = original.ApplicationTypeVersionsClient
type ApplicationTypeVersionsCreateOrUpdateFuture = original.ApplicationTypeVersionsCreateOrUpdateFuture
type ApplicationTypeVersionsDeleteFuture = original.ApplicationTypeVersionsDeleteFuture
type ApplicationTypesClient = original.ApplicationTypesClient
type ApplicationTypesDeleteFuture = original.ApplicationTypesDeleteFuture
type ApplicationUpgradePolicy = original.ApplicationUpgradePolicy
type ApplicationUserAssignedIdentity = original.ApplicationUserAssignedIdentity
type ApplicationsClient = original.ApplicationsClient
type ApplicationsCreateOrUpdateFuture = original.ApplicationsCreateOrUpdateFuture
type ApplicationsDeleteFuture = original.ApplicationsDeleteFuture
type ApplicationsUpdateFuture = original.ApplicationsUpdateFuture
type ArmApplicationHealthPolicy = original.ArmApplicationHealthPolicy
type ArmRollingUpgradeMonitoringPolicy = original.ArmRollingUpgradeMonitoringPolicy
type ArmServiceTypeHealthPolicy = original.ArmServiceTypeHealthPolicy
type AvailableOperationDisplay = original.AvailableOperationDisplay
type AzureActiveDirectory = original.AzureActiveDirectory
type BaseClient = original.BaseClient
type BasicPartitionSchemeDescription = original.BasicPartitionSchemeDescription
type BasicServicePlacementPolicyDescription = original.BasicServicePlacementPolicyDescription
type BasicServiceResourceProperties = original.BasicServiceResourceProperties
type BasicServiceResourceUpdateProperties = original.BasicServiceResourceUpdateProperties
type CertificateDescription = original.CertificateDescription
type ClientCertificateCommonName = original.ClientCertificateCommonName
type ClientCertificateThumbprint = original.ClientCertificateThumbprint
type Cluster = original.Cluster
type ClusterCodeVersionsListResult = original.ClusterCodeVersionsListResult
type ClusterCodeVersionsResult = original.ClusterCodeVersionsResult
type ClusterHealthPolicy = original.ClusterHealthPolicy
type ClusterListResult = original.ClusterListResult
type ClusterProperties = original.ClusterProperties
type ClusterPropertiesUpdateParameters = original.ClusterPropertiesUpdateParameters
type ClusterUpdateParameters = original.ClusterUpdateParameters
type ClusterUpgradeDeltaHealthPolicy = original.ClusterUpgradeDeltaHealthPolicy
type ClusterUpgradePolicy = original.ClusterUpgradePolicy
type ClusterVersionDetails = original.ClusterVersionDetails
type ClusterVersionsClient = original.ClusterVersionsClient
type ClustersClient = original.ClustersClient
type ClustersCreateOrUpdateFuture = original.ClustersCreateOrUpdateFuture
type ClustersUpdateFuture = original.ClustersUpdateFuture
type DiagnosticsStorageAccountConfig = original.DiagnosticsStorageAccountConfig
type EndpointRangeDescription = original.EndpointRangeDescription
type ErrorModel = original.ErrorModel
type ErrorModelError = original.ErrorModelError
type ManagedIdentity = original.ManagedIdentity
type NamedPartitionSchemeDescription = original.NamedPartitionSchemeDescription
type NodeTypeDescription = original.NodeTypeDescription
type Notification = original.Notification
type NotificationTarget = original.NotificationTarget
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationResult = original.OperationResult
type OperationsClient = original.OperationsClient
type PartitionSchemeDescription = original.PartitionSchemeDescription
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type ServerCertificateCommonName = original.ServerCertificateCommonName
type ServerCertificateCommonNames = original.ServerCertificateCommonNames
type ServiceCorrelationDescription = original.ServiceCorrelationDescription
type ServiceLoadMetricDescription = original.ServiceLoadMetricDescription
type ServicePlacementPolicyDescription = original.ServicePlacementPolicyDescription
type ServiceResource = original.ServiceResource
type ServiceResourceList = original.ServiceResourceList
type ServiceResourceProperties = original.ServiceResourceProperties
type ServiceResourcePropertiesBase = original.ServiceResourcePropertiesBase
type ServiceResourceUpdate = original.ServiceResourceUpdate
type ServiceResourceUpdateProperties = original.ServiceResourceUpdateProperties
type ServiceTypeDeltaHealthPolicy = original.ServiceTypeDeltaHealthPolicy
type ServiceTypeHealthPolicy = original.ServiceTypeHealthPolicy
type ServicesClient = original.ServicesClient
type ServicesCreateOrUpdateFuture = original.ServicesCreateOrUpdateFuture
type ServicesDeleteFuture = original.ServicesDeleteFuture
type ServicesUpdateFuture = original.ServicesUpdateFuture
type SettingsParameterDescription = original.SettingsParameterDescription
type SettingsSectionDescription = original.SettingsSectionDescription
type SingletonPartitionSchemeDescription = original.SingletonPartitionSchemeDescription
type StatefulServiceProperties = original.StatefulServiceProperties
type StatefulServiceUpdateProperties = original.StatefulServiceUpdateProperties
type StatelessServiceProperties = original.StatelessServiceProperties
type StatelessServiceUpdateProperties = original.StatelessServiceUpdateProperties
type SystemData = original.SystemData
type UniformInt64RangePartitionSchemeDescription = original.UniformInt64RangePartitionSchemeDescription
type UpgradableVersionPathResult = original.UpgradableVersionPathResult
type UpgradableVersionsDescription = original.UpgradableVersionsDescription
type UserAssignedIdentity = original.UserAssignedIdentity

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewApplicationTypeVersionsClient(subscriptionID string) ApplicationTypeVersionsClient {
	return original.NewApplicationTypeVersionsClient(subscriptionID)
}
func NewApplicationTypeVersionsClientWithBaseURI(baseURI string, subscriptionID string) ApplicationTypeVersionsClient {
	return original.NewApplicationTypeVersionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewApplicationTypesClient(subscriptionID string) ApplicationTypesClient {
	return original.NewApplicationTypesClient(subscriptionID)
}
func NewApplicationTypesClientWithBaseURI(baseURI string, subscriptionID string) ApplicationTypesClient {
	return original.NewApplicationTypesClientWithBaseURI(baseURI, subscriptionID)
}
func NewApplicationsClient(subscriptionID string) ApplicationsClient {
	return original.NewApplicationsClient(subscriptionID)
}
func NewApplicationsClientWithBaseURI(baseURI string, subscriptionID string) ApplicationsClient {
	return original.NewApplicationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewClusterVersionsClient(subscriptionID string) ClusterVersionsClient {
	return original.NewClusterVersionsClient(subscriptionID)
}
func NewClusterVersionsClientWithBaseURI(baseURI string, subscriptionID string) ClusterVersionsClient {
	return original.NewClusterVersionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewClustersClient(subscriptionID string) ClustersClient {
	return original.NewClustersClient(subscriptionID)
}
func NewClustersClientWithBaseURI(baseURI string, subscriptionID string) ClustersClient {
	return original.NewClustersClientWithBaseURI(baseURI, subscriptionID)
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
func NewServicesClient(subscriptionID string) ServicesClient {
	return original.NewServicesClient(subscriptionID)
}
func NewServicesClientWithBaseURI(baseURI string, subscriptionID string) ServicesClient {
	return original.NewServicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleArmServicePackageActivationModeValues() []ArmServicePackageActivationMode {
	return original.PossibleArmServicePackageActivationModeValues()
}
func PossibleArmUpgradeFailureActionValues() []ArmUpgradeFailureAction {
	return original.PossibleArmUpgradeFailureActionValues()
}
func PossibleClusterStateValues() []ClusterState {
	return original.PossibleClusterStateValues()
}
func PossibleClusterUpgradeCadenceValues() []ClusterUpgradeCadence {
	return original.PossibleClusterUpgradeCadenceValues()
}
func PossibleDurabilityLevelValues() []DurabilityLevel {
	return original.PossibleDurabilityLevelValues()
}
func PossibleEnvironmentValues() []Environment {
	return original.PossibleEnvironmentValues()
}
func PossibleManagedIdentityTypeValues() []ManagedIdentityType {
	return original.PossibleManagedIdentityTypeValues()
}
func PossibleMoveCostValues() []MoveCost {
	return original.PossibleMoveCostValues()
}
func PossibleNotificationChannelValues() []NotificationChannel {
	return original.PossibleNotificationChannelValues()
}
func PossibleNotificationLevelValues() []NotificationLevel {
	return original.PossibleNotificationLevelValues()
}
func PossiblePartitionSchemeBasicPartitionSchemeDescriptionValues() []PartitionSchemeBasicPartitionSchemeDescription {
	return original.PossiblePartitionSchemeBasicPartitionSchemeDescriptionValues()
}
func PossiblePartitionSchemeValues() []PartitionScheme {
	return original.PossiblePartitionSchemeValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleReliabilityLevel1Values() []ReliabilityLevel1 {
	return original.PossibleReliabilityLevel1Values()
}
func PossibleReliabilityLevelValues() []ReliabilityLevel {
	return original.PossibleReliabilityLevelValues()
}
func PossibleRollingUpgradeModeValues() []RollingUpgradeMode {
	return original.PossibleRollingUpgradeModeValues()
}
func PossibleServiceCorrelationSchemeValues() []ServiceCorrelationScheme {
	return original.PossibleServiceCorrelationSchemeValues()
}
func PossibleServiceKindBasicServiceResourcePropertiesValues() []ServiceKindBasicServiceResourceProperties {
	return original.PossibleServiceKindBasicServiceResourcePropertiesValues()
}
func PossibleServiceKindBasicServiceResourceUpdatePropertiesValues() []ServiceKindBasicServiceResourceUpdateProperties {
	return original.PossibleServiceKindBasicServiceResourceUpdatePropertiesValues()
}
func PossibleServiceKindValues() []ServiceKind {
	return original.PossibleServiceKindValues()
}
func PossibleServiceLoadMetricWeightValues() []ServiceLoadMetricWeight {
	return original.PossibleServiceLoadMetricWeightValues()
}
func PossibleServicePlacementPolicyTypeValues() []ServicePlacementPolicyType {
	return original.PossibleServicePlacementPolicyTypeValues()
}
func PossibleSfZonalUpgradeModeValues() []SfZonalUpgradeMode {
	return original.PossibleSfZonalUpgradeModeValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func PossibleUpgradeModeValues() []UpgradeMode {
	return original.PossibleUpgradeModeValues()
}
func PossibleVmssZonalUpgradeModeValues() []VmssZonalUpgradeMode {
	return original.PossibleVmssZonalUpgradeModeValues()
}
func PossibleX509StoreName1Values() []X509StoreName1 {
	return original.PossibleX509StoreName1Values()
}
func PossibleX509StoreNameValues() []X509StoreName {
	return original.PossibleX509StoreNameValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
