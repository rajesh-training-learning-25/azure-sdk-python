//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package machinelearningservices

import (
	"context"

	original "github.com/Azure/dev/azure-sdk-for-go/services/machinelearningservices/mgmt/2021-07-01/machinelearningservices"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AllocationState = original.AllocationState

const (
	AllocationStateResizing AllocationState = original.AllocationStateResizing
	AllocationStateSteady   AllocationState = original.AllocationStateSteady
)

type ApplicationSharingPolicy = original.ApplicationSharingPolicy

const (
	ApplicationSharingPolicyPersonal ApplicationSharingPolicy = original.ApplicationSharingPolicyPersonal
	ApplicationSharingPolicyShared   ApplicationSharingPolicy = original.ApplicationSharingPolicyShared
)

type ClusterPurpose = original.ClusterPurpose

const (
	ClusterPurposeDenseProd ClusterPurpose = original.ClusterPurposeDenseProd
	ClusterPurposeDevTest   ClusterPurpose = original.ClusterPurposeDevTest
	ClusterPurposeFastProd  ClusterPurpose = original.ClusterPurposeFastProd
)

type ComputeInstanceAuthorizationType = original.ComputeInstanceAuthorizationType

const (
	ComputeInstanceAuthorizationTypePersonal ComputeInstanceAuthorizationType = original.ComputeInstanceAuthorizationTypePersonal
)

type ComputeInstanceState = original.ComputeInstanceState

const (
	ComputeInstanceStateCreateFailed    ComputeInstanceState = original.ComputeInstanceStateCreateFailed
	ComputeInstanceStateCreating        ComputeInstanceState = original.ComputeInstanceStateCreating
	ComputeInstanceStateDeleting        ComputeInstanceState = original.ComputeInstanceStateDeleting
	ComputeInstanceStateJobRunning      ComputeInstanceState = original.ComputeInstanceStateJobRunning
	ComputeInstanceStateRestarting      ComputeInstanceState = original.ComputeInstanceStateRestarting
	ComputeInstanceStateRunning         ComputeInstanceState = original.ComputeInstanceStateRunning
	ComputeInstanceStateSettingUp       ComputeInstanceState = original.ComputeInstanceStateSettingUp
	ComputeInstanceStateSetupFailed     ComputeInstanceState = original.ComputeInstanceStateSetupFailed
	ComputeInstanceStateStarting        ComputeInstanceState = original.ComputeInstanceStateStarting
	ComputeInstanceStateStopped         ComputeInstanceState = original.ComputeInstanceStateStopped
	ComputeInstanceStateStopping        ComputeInstanceState = original.ComputeInstanceStateStopping
	ComputeInstanceStateUnknown         ComputeInstanceState = original.ComputeInstanceStateUnknown
	ComputeInstanceStateUnusable        ComputeInstanceState = original.ComputeInstanceStateUnusable
	ComputeInstanceStateUserSettingUp   ComputeInstanceState = original.ComputeInstanceStateUserSettingUp
	ComputeInstanceStateUserSetupFailed ComputeInstanceState = original.ComputeInstanceStateUserSetupFailed
)

type ComputeType = original.ComputeType

const (
	ComputeTypeAKS               ComputeType = original.ComputeTypeAKS
	ComputeTypeAmlCompute        ComputeType = original.ComputeTypeAmlCompute
	ComputeTypeComputeInstance   ComputeType = original.ComputeTypeComputeInstance
	ComputeTypeDatabricks        ComputeType = original.ComputeTypeDatabricks
	ComputeTypeDataFactory       ComputeType = original.ComputeTypeDataFactory
	ComputeTypeDataLakeAnalytics ComputeType = original.ComputeTypeDataLakeAnalytics
	ComputeTypeHDInsight         ComputeType = original.ComputeTypeHDInsight
	ComputeTypeKubernetes        ComputeType = original.ComputeTypeKubernetes
	ComputeTypeSynapseSpark      ComputeType = original.ComputeTypeSynapseSpark
	ComputeTypeVirtualMachine    ComputeType = original.ComputeTypeVirtualMachine
)

type ComputeTypeBasicCompute = original.ComputeTypeBasicCompute

const (
	ComputeTypeBasicComputeComputeTypeAKS               ComputeTypeBasicCompute = original.ComputeTypeBasicComputeComputeTypeAKS
	ComputeTypeBasicComputeComputeTypeAmlCompute        ComputeTypeBasicCompute = original.ComputeTypeBasicComputeComputeTypeAmlCompute
	ComputeTypeBasicComputeComputeTypeCompute           ComputeTypeBasicCompute = original.ComputeTypeBasicComputeComputeTypeCompute
	ComputeTypeBasicComputeComputeTypeComputeInstance   ComputeTypeBasicCompute = original.ComputeTypeBasicComputeComputeTypeComputeInstance
	ComputeTypeBasicComputeComputeTypeDatabricks        ComputeTypeBasicCompute = original.ComputeTypeBasicComputeComputeTypeDatabricks
	ComputeTypeBasicComputeComputeTypeDataFactory       ComputeTypeBasicCompute = original.ComputeTypeBasicComputeComputeTypeDataFactory
	ComputeTypeBasicComputeComputeTypeDataLakeAnalytics ComputeTypeBasicCompute = original.ComputeTypeBasicComputeComputeTypeDataLakeAnalytics
	ComputeTypeBasicComputeComputeTypeHDInsight         ComputeTypeBasicCompute = original.ComputeTypeBasicComputeComputeTypeHDInsight
	ComputeTypeBasicComputeComputeTypeSynapseSpark      ComputeTypeBasicCompute = original.ComputeTypeBasicComputeComputeTypeSynapseSpark
	ComputeTypeBasicComputeComputeTypeVirtualMachine    ComputeTypeBasicCompute = original.ComputeTypeBasicComputeComputeTypeVirtualMachine
)

type ComputeTypeBasicComputeSecrets = original.ComputeTypeBasicComputeSecrets

const (
	ComputeTypeBasicComputeSecretsComputeTypeComputeSecrets ComputeTypeBasicComputeSecrets = original.ComputeTypeBasicComputeSecretsComputeTypeComputeSecrets
	ComputeTypeBasicComputeSecretsComputeTypeVirtualMachine ComputeTypeBasicComputeSecrets = original.ComputeTypeBasicComputeSecretsComputeTypeVirtualMachine
)

type CreatedByType = original.CreatedByType

const (
	CreatedByTypeApplication     CreatedByType = original.CreatedByTypeApplication
	CreatedByTypeKey             CreatedByType = original.CreatedByTypeKey
	CreatedByTypeManagedIdentity CreatedByType = original.CreatedByTypeManagedIdentity
	CreatedByTypeUser            CreatedByType = original.CreatedByTypeUser
)

type DiagnoseResultLevel = original.DiagnoseResultLevel

const (
	DiagnoseResultLevelError       DiagnoseResultLevel = original.DiagnoseResultLevelError
	DiagnoseResultLevelInformation DiagnoseResultLevel = original.DiagnoseResultLevelInformation
	DiagnoseResultLevelWarning     DiagnoseResultLevel = original.DiagnoseResultLevelWarning
)

type EncryptionStatus = original.EncryptionStatus

const (
	EncryptionStatusDisabled EncryptionStatus = original.EncryptionStatusDisabled
	EncryptionStatusEnabled  EncryptionStatus = original.EncryptionStatusEnabled
)

type LoadBalancerType = original.LoadBalancerType

const (
	LoadBalancerTypeInternalLoadBalancer LoadBalancerType = original.LoadBalancerTypeInternalLoadBalancer
	LoadBalancerTypePublicIP             LoadBalancerType = original.LoadBalancerTypePublicIP
)

type NodeState = original.NodeState

const (
	NodeStateIdle      NodeState = original.NodeStateIdle
	NodeStateLeaving   NodeState = original.NodeStateLeaving
	NodeStatePreempted NodeState = original.NodeStatePreempted
	NodeStatePreparing NodeState = original.NodeStatePreparing
	NodeStateRunning   NodeState = original.NodeStateRunning
	NodeStateUnusable  NodeState = original.NodeStateUnusable
)

type OperationName = original.OperationName

const (
	OperationNameCreate  OperationName = original.OperationNameCreate
	OperationNameDelete  OperationName = original.OperationNameDelete
	OperationNameReimage OperationName = original.OperationNameReimage
	OperationNameRestart OperationName = original.OperationNameRestart
	OperationNameStart   OperationName = original.OperationNameStart
	OperationNameStop    OperationName = original.OperationNameStop
)

type OperationStatus = original.OperationStatus

const (
	OperationStatusCreateFailed  OperationStatus = original.OperationStatusCreateFailed
	OperationStatusDeleteFailed  OperationStatus = original.OperationStatusDeleteFailed
	OperationStatusInProgress    OperationStatus = original.OperationStatusInProgress
	OperationStatusReimageFailed OperationStatus = original.OperationStatusReimageFailed
	OperationStatusRestartFailed OperationStatus = original.OperationStatusRestartFailed
	OperationStatusStartFailed   OperationStatus = original.OperationStatusStartFailed
	OperationStatusStopFailed    OperationStatus = original.OperationStatusStopFailed
	OperationStatusSucceeded     OperationStatus = original.OperationStatusSucceeded
)

type OsType = original.OsType

const (
	OsTypeLinux   OsType = original.OsTypeLinux
	OsTypeWindows OsType = original.OsTypeWindows
)

type PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningState

const (
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateCreating
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateDeleting
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateFailed
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateSucceeded
)

type PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatus

const (
	PrivateEndpointServiceConnectionStatusApproved     PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatusApproved
	PrivateEndpointServiceConnectionStatusDisconnected PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatusDisconnected
	PrivateEndpointServiceConnectionStatusPending      PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatusPending
	PrivateEndpointServiceConnectionStatusRejected     PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatusRejected
	PrivateEndpointServiceConnectionStatusTimeout      PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatusTimeout
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateCanceled  ProvisioningState = original.ProvisioningStateCanceled
	ProvisioningStateCreating  ProvisioningState = original.ProvisioningStateCreating
	ProvisioningStateDeleting  ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateFailed    ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateSucceeded ProvisioningState = original.ProvisioningStateSucceeded
	ProvisioningStateUnknown   ProvisioningState = original.ProvisioningStateUnknown
	ProvisioningStateUpdating  ProvisioningState = original.ProvisioningStateUpdating
)

type PublicNetworkAccess = original.PublicNetworkAccess

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = original.PublicNetworkAccessDisabled
	PublicNetworkAccessEnabled  PublicNetworkAccess = original.PublicNetworkAccessEnabled
)

type QuotaUnit = original.QuotaUnit

const (
	QuotaUnitCount QuotaUnit = original.QuotaUnitCount
)

type ReasonCode = original.ReasonCode

const (
	ReasonCodeNotAvailableForRegion       ReasonCode = original.ReasonCodeNotAvailableForRegion
	ReasonCodeNotAvailableForSubscription ReasonCode = original.ReasonCodeNotAvailableForSubscription
	ReasonCodeNotSpecified                ReasonCode = original.ReasonCodeNotSpecified
)

type RemoteLoginPortPublicAccess = original.RemoteLoginPortPublicAccess

const (
	RemoteLoginPortPublicAccessDisabled     RemoteLoginPortPublicAccess = original.RemoteLoginPortPublicAccessDisabled
	RemoteLoginPortPublicAccessEnabled      RemoteLoginPortPublicAccess = original.RemoteLoginPortPublicAccessEnabled
	RemoteLoginPortPublicAccessNotSpecified RemoteLoginPortPublicAccess = original.RemoteLoginPortPublicAccessNotSpecified
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	ResourceIdentityTypeNone                       ResourceIdentityType = original.ResourceIdentityTypeNone
	ResourceIdentityTypeSystemAssigned             ResourceIdentityType = original.ResourceIdentityTypeSystemAssigned
	ResourceIdentityTypeSystemAssignedUserAssigned ResourceIdentityType = original.ResourceIdentityTypeSystemAssignedUserAssigned
	ResourceIdentityTypeUserAssigned               ResourceIdentityType = original.ResourceIdentityTypeUserAssigned
)

type SSHPublicAccess = original.SSHPublicAccess

const (
	SSHPublicAccessDisabled SSHPublicAccess = original.SSHPublicAccessDisabled
	SSHPublicAccessEnabled  SSHPublicAccess = original.SSHPublicAccessEnabled
)

type Status = original.Status

const (
	StatusFailure                              Status = original.StatusFailure
	StatusInvalidQuotaBelowClusterMinimum      Status = original.StatusInvalidQuotaBelowClusterMinimum
	StatusInvalidQuotaExceedsSubscriptionLimit Status = original.StatusInvalidQuotaExceedsSubscriptionLimit
	StatusInvalidVMFamilyName                  Status = original.StatusInvalidVMFamilyName
	StatusOperationNotEnabledForRegion         Status = original.StatusOperationNotEnabledForRegion
	StatusOperationNotSupportedForSku          Status = original.StatusOperationNotSupportedForSku
	StatusSuccess                              Status = original.StatusSuccess
	StatusUndefined                            Status = original.StatusUndefined
)

type Status1 = original.Status1

const (
	Status1Auto     Status1 = original.Status1Auto
	Status1Disabled Status1 = original.Status1Disabled
	Status1Enabled  Status1 = original.Status1Enabled
)

type UnderlyingResourceAction = original.UnderlyingResourceAction

const (
	UnderlyingResourceActionDelete UnderlyingResourceAction = original.UnderlyingResourceActionDelete
	UnderlyingResourceActionDetach UnderlyingResourceAction = original.UnderlyingResourceActionDetach
)

type UsageUnit = original.UsageUnit

const (
	UsageUnitCount UsageUnit = original.UsageUnitCount
)

type VMPriceOSType = original.VMPriceOSType

const (
	VMPriceOSTypeLinux   VMPriceOSType = original.VMPriceOSTypeLinux
	VMPriceOSTypeWindows VMPriceOSType = original.VMPriceOSTypeWindows
)

type VMPriority = original.VMPriority

const (
	VMPriorityDedicated   VMPriority = original.VMPriorityDedicated
	VMPriorityLowPriority VMPriority = original.VMPriorityLowPriority
)

type VMTier = original.VMTier

const (
	VMTierLowPriority VMTier = original.VMTierLowPriority
	VMTierSpot        VMTier = original.VMTierSpot
	VMTierStandard    VMTier = original.VMTierStandard
)

type ValueFormat = original.ValueFormat

const (
	ValueFormatJSON ValueFormat = original.ValueFormatJSON
)

type AKS = original.AKS
type AKSProperties = original.AKSProperties
type AksComputeSecrets = original.AksComputeSecrets
type AksComputeSecretsProperties = original.AksComputeSecretsProperties
type AksNetworkingConfiguration = original.AksNetworkingConfiguration
type AmlCompute = original.AmlCompute
type AmlComputeNodeInformation = original.AmlComputeNodeInformation
type AmlComputeNodesInformation = original.AmlComputeNodesInformation
type AmlComputeNodesInformationIterator = original.AmlComputeNodesInformationIterator
type AmlComputeNodesInformationPage = original.AmlComputeNodesInformationPage
type AmlComputeProperties = original.AmlComputeProperties
type AmlUserFeature = original.AmlUserFeature
type AssignedUser = original.AssignedUser
type AutoPauseProperties = original.AutoPauseProperties
type AutoScaleProperties = original.AutoScaleProperties
type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type BasicCompute = original.BasicCompute
type BasicComputeSecrets = original.BasicComputeSecrets
type ClusterUpdateParameters = original.ClusterUpdateParameters
type ClusterUpdateProperties = original.ClusterUpdateProperties
type Compute = original.Compute
type ComputeClient = original.ComputeClient
type ComputeCreateOrUpdateFuture = original.ComputeCreateOrUpdateFuture
type ComputeDeleteFuture = original.ComputeDeleteFuture
type ComputeInstance = original.ComputeInstance
type ComputeInstanceApplication = original.ComputeInstanceApplication
type ComputeInstanceConnectivityEndpoints = original.ComputeInstanceConnectivityEndpoints
type ComputeInstanceCreatedBy = original.ComputeInstanceCreatedBy
type ComputeInstanceLastOperation = original.ComputeInstanceLastOperation
type ComputeInstanceProperties = original.ComputeInstanceProperties
type ComputeInstanceSSHSettings = original.ComputeInstanceSSHSettings
type ComputeResource = original.ComputeResource
type ComputeRestartFuture = original.ComputeRestartFuture
type ComputeSecrets = original.ComputeSecrets
type ComputeSecretsModel = original.ComputeSecretsModel
type ComputeStartFuture = original.ComputeStartFuture
type ComputeStopFuture = original.ComputeStopFuture
type ComputeUpdateFuture = original.ComputeUpdateFuture
type ContainerResourceRequirements = original.ContainerResourceRequirements
type CosmosDbSettings = original.CosmosDbSettings
type DataFactory = original.DataFactory
type DataLakeAnalytics = original.DataLakeAnalytics
type DataLakeAnalyticsProperties = original.DataLakeAnalyticsProperties
type Databricks = original.Databricks
type DatabricksComputeSecrets = original.DatabricksComputeSecrets
type DatabricksComputeSecretsProperties = original.DatabricksComputeSecretsProperties
type DatabricksProperties = original.DatabricksProperties
type DiagnoseRequestProperties = original.DiagnoseRequestProperties
type DiagnoseResponseResult = original.DiagnoseResponseResult
type DiagnoseResponseResultValue = original.DiagnoseResponseResultValue
type DiagnoseResult = original.DiagnoseResult
type DiagnoseWorkspaceParameters = original.DiagnoseWorkspaceParameters
type EncryptionProperty = original.EncryptionProperty
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type EstimatedVMPrice = original.EstimatedVMPrice
type EstimatedVMPrices = original.EstimatedVMPrices
type ExternalFQDNResponse = original.ExternalFQDNResponse
type FQDNEndpoint = original.FQDNEndpoint
type FQDNEndpointDetail = original.FQDNEndpointDetail
type FQDNEndpoints = original.FQDNEndpoints
type FQDNEndpointsProperties = original.FQDNEndpointsProperties
type HDInsight = original.HDInsight
type HDInsightProperties = original.HDInsightProperties
type Identity = original.Identity
type IdentityForCmk = original.IdentityForCmk
type InstanceTypeSchema = original.InstanceTypeSchema
type InstanceTypeSchemaResources = original.InstanceTypeSchemaResources
type KeyVaultProperties = original.KeyVaultProperties
type Kubernetes = original.Kubernetes
type KubernetesProperties = original.KubernetesProperties
type KubernetesSchema = original.KubernetesSchema
type ListAmlUserFeatureResult = original.ListAmlUserFeatureResult
type ListAmlUserFeatureResultIterator = original.ListAmlUserFeatureResultIterator
type ListAmlUserFeatureResultPage = original.ListAmlUserFeatureResultPage
type ListNotebookKeysResult = original.ListNotebookKeysResult
type ListStorageAccountKeysResult = original.ListStorageAccountKeysResult
type ListUsagesResult = original.ListUsagesResult
type ListUsagesResultIterator = original.ListUsagesResultIterator
type ListUsagesResultPage = original.ListUsagesResultPage
type ListWorkspaceKeysResult = original.ListWorkspaceKeysResult
type ListWorkspaceQuotas = original.ListWorkspaceQuotas
type ListWorkspaceQuotasIterator = original.ListWorkspaceQuotasIterator
type ListWorkspaceQuotasPage = original.ListWorkspaceQuotasPage
type NodeStateCounts = original.NodeStateCounts
type NotebookAccessTokenResult = original.NotebookAccessTokenResult
type NotebookPreparationError = original.NotebookPreparationError
type NotebookResourceInfo = original.NotebookResourceInfo
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationsClient = original.OperationsClient
type PaginatedComputeResourcesList = original.PaginatedComputeResourcesList
type PaginatedComputeResourcesListIterator = original.PaginatedComputeResourcesListIterator
type PaginatedComputeResourcesListPage = original.PaginatedComputeResourcesListPage
type PaginatedWorkspaceConnectionsList = original.PaginatedWorkspaceConnectionsList
type Password = original.Password
type PersonalComputeInstanceSettings = original.PersonalComputeInstanceSettings
type PrivateEndpoint = original.PrivateEndpoint
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionListResult = original.PrivateEndpointConnectionListResult
type PrivateEndpointConnectionProperties = original.PrivateEndpointConnectionProperties
type PrivateEndpointConnectionsClient = original.PrivateEndpointConnectionsClient
type PrivateLinkResource = original.PrivateLinkResource
type PrivateLinkResourceListResult = original.PrivateLinkResourceListResult
type PrivateLinkResourceProperties = original.PrivateLinkResourceProperties
type PrivateLinkResourcesClient = original.PrivateLinkResourcesClient
type PrivateLinkServiceConnectionState = original.PrivateLinkServiceConnectionState
type ProxyResource = original.ProxyResource
type QuotaBaseProperties = original.QuotaBaseProperties
type QuotaUpdateParameters = original.QuotaUpdateParameters
type QuotasClient = original.QuotasClient
type RegistryListCredentialsResult = original.RegistryListCredentialsResult
type Resource = original.Resource
type ResourceID = original.ResourceID
type ResourceName = original.ResourceName
type ResourceQuota = original.ResourceQuota
type ResourceSkuLocationInfo = original.ResourceSkuLocationInfo
type ResourceSkuZoneDetails = original.ResourceSkuZoneDetails
type Restriction = original.Restriction
type SKUCapability = original.SKUCapability
type ScaleSettings = original.ScaleSettings
type ScaleSettingsInformation = original.ScaleSettingsInformation
type ScriptReference = original.ScriptReference
type ScriptsToExecute = original.ScriptsToExecute
type ServiceManagedResourcesSettings = original.ServiceManagedResourcesSettings
type ServicePrincipalCredentials = original.ServicePrincipalCredentials
type SetupScripts = original.SetupScripts
type SharedPrivateLinkResource = original.SharedPrivateLinkResource
type SharedPrivateLinkResourceProperty = original.SharedPrivateLinkResourceProperty
type Sku = original.Sku
type SkuListResult = original.SkuListResult
type SkuListResultIterator = original.SkuListResultIterator
type SkuListResultPage = original.SkuListResultPage
type SslConfiguration = original.SslConfiguration
type SynapseSpark = original.SynapseSpark
type SynapseSparkProperties = original.SynapseSparkProperties
type SystemData = original.SystemData
type SystemService = original.SystemService
type TrackedResource = original.TrackedResource
type UpdateWorkspaceQuotas = original.UpdateWorkspaceQuotas
type UpdateWorkspaceQuotasResult = original.UpdateWorkspaceQuotasResult
type Usage = original.Usage
type UsageName = original.UsageName
type UsagesClient = original.UsagesClient
type UserAccountCredentials = original.UserAccountCredentials
type UserAssignedIdentity = original.UserAssignedIdentity
type VirtualMachine = original.VirtualMachine
type VirtualMachineImage = original.VirtualMachineImage
type VirtualMachineProperties = original.VirtualMachineProperties
type VirtualMachineSSHCredentials = original.VirtualMachineSSHCredentials
type VirtualMachineSecrets = original.VirtualMachineSecrets
type VirtualMachineSize = original.VirtualMachineSize
type VirtualMachineSizeListResult = original.VirtualMachineSizeListResult
type VirtualMachineSizesClient = original.VirtualMachineSizesClient
type Workspace = original.Workspace
type WorkspaceConnection = original.WorkspaceConnection
type WorkspaceConnectionProps = original.WorkspaceConnectionProps
type WorkspaceConnectionsClient = original.WorkspaceConnectionsClient
type WorkspaceFeaturesClient = original.WorkspaceFeaturesClient
type WorkspaceListResult = original.WorkspaceListResult
type WorkspaceListResultIterator = original.WorkspaceListResultIterator
type WorkspaceListResultPage = original.WorkspaceListResultPage
type WorkspaceProperties = original.WorkspaceProperties
type WorkspacePropertiesUpdateParameters = original.WorkspacePropertiesUpdateParameters
type WorkspaceSku = original.WorkspaceSku
type WorkspaceSkusClient = original.WorkspaceSkusClient
type WorkspaceUpdateParameters = original.WorkspaceUpdateParameters
type WorkspacesClient = original.WorkspacesClient
type WorkspacesCreateOrUpdateFuture = original.WorkspacesCreateOrUpdateFuture
type WorkspacesDeleteFuture = original.WorkspacesDeleteFuture
type WorkspacesDiagnoseFuture = original.WorkspacesDiagnoseFuture
type WorkspacesPrepareNotebookFuture = original.WorkspacesPrepareNotebookFuture
type WorkspacesResyncKeysFuture = original.WorkspacesResyncKeysFuture

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAmlComputeNodesInformationIterator(page AmlComputeNodesInformationPage) AmlComputeNodesInformationIterator {
	return original.NewAmlComputeNodesInformationIterator(page)
}
func NewAmlComputeNodesInformationPage(cur AmlComputeNodesInformation, getNextPage func(context.Context, AmlComputeNodesInformation) (AmlComputeNodesInformation, error)) AmlComputeNodesInformationPage {
	return original.NewAmlComputeNodesInformationPage(cur, getNextPage)
}
func NewComputeClient(subscriptionID string) ComputeClient {
	return original.NewComputeClient(subscriptionID)
}
func NewComputeClientWithBaseURI(baseURI string, subscriptionID string) ComputeClient {
	return original.NewComputeClientWithBaseURI(baseURI, subscriptionID)
}
func NewListAmlUserFeatureResultIterator(page ListAmlUserFeatureResultPage) ListAmlUserFeatureResultIterator {
	return original.NewListAmlUserFeatureResultIterator(page)
}
func NewListAmlUserFeatureResultPage(cur ListAmlUserFeatureResult, getNextPage func(context.Context, ListAmlUserFeatureResult) (ListAmlUserFeatureResult, error)) ListAmlUserFeatureResultPage {
	return original.NewListAmlUserFeatureResultPage(cur, getNextPage)
}
func NewListUsagesResultIterator(page ListUsagesResultPage) ListUsagesResultIterator {
	return original.NewListUsagesResultIterator(page)
}
func NewListUsagesResultPage(cur ListUsagesResult, getNextPage func(context.Context, ListUsagesResult) (ListUsagesResult, error)) ListUsagesResultPage {
	return original.NewListUsagesResultPage(cur, getNextPage)
}
func NewListWorkspaceQuotasIterator(page ListWorkspaceQuotasPage) ListWorkspaceQuotasIterator {
	return original.NewListWorkspaceQuotasIterator(page)
}
func NewListWorkspaceQuotasPage(cur ListWorkspaceQuotas, getNextPage func(context.Context, ListWorkspaceQuotas) (ListWorkspaceQuotas, error)) ListWorkspaceQuotasPage {
	return original.NewListWorkspaceQuotasPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPaginatedComputeResourcesListIterator(page PaginatedComputeResourcesListPage) PaginatedComputeResourcesListIterator {
	return original.NewPaginatedComputeResourcesListIterator(page)
}
func NewPaginatedComputeResourcesListPage(cur PaginatedComputeResourcesList, getNextPage func(context.Context, PaginatedComputeResourcesList) (PaginatedComputeResourcesList, error)) PaginatedComputeResourcesListPage {
	return original.NewPaginatedComputeResourcesListPage(cur, getNextPage)
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
func NewQuotasClient(subscriptionID string) QuotasClient {
	return original.NewQuotasClient(subscriptionID)
}
func NewQuotasClientWithBaseURI(baseURI string, subscriptionID string) QuotasClient {
	return original.NewQuotasClientWithBaseURI(baseURI, subscriptionID)
}
func NewSkuListResultIterator(page SkuListResultPage) SkuListResultIterator {
	return original.NewSkuListResultIterator(page)
}
func NewSkuListResultPage(cur SkuListResult, getNextPage func(context.Context, SkuListResult) (SkuListResult, error)) SkuListResultPage {
	return original.NewSkuListResultPage(cur, getNextPage)
}
func NewUsagesClient(subscriptionID string) UsagesClient {
	return original.NewUsagesClient(subscriptionID)
}
func NewUsagesClientWithBaseURI(baseURI string, subscriptionID string) UsagesClient {
	return original.NewUsagesClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualMachineSizesClient(subscriptionID string) VirtualMachineSizesClient {
	return original.NewVirtualMachineSizesClient(subscriptionID)
}
func NewVirtualMachineSizesClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineSizesClient {
	return original.NewVirtualMachineSizesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewWorkspaceConnectionsClient(subscriptionID string) WorkspaceConnectionsClient {
	return original.NewWorkspaceConnectionsClient(subscriptionID)
}
func NewWorkspaceConnectionsClientWithBaseURI(baseURI string, subscriptionID string) WorkspaceConnectionsClient {
	return original.NewWorkspaceConnectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWorkspaceFeaturesClient(subscriptionID string) WorkspaceFeaturesClient {
	return original.NewWorkspaceFeaturesClient(subscriptionID)
}
func NewWorkspaceFeaturesClientWithBaseURI(baseURI string, subscriptionID string) WorkspaceFeaturesClient {
	return original.NewWorkspaceFeaturesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWorkspaceListResultIterator(page WorkspaceListResultPage) WorkspaceListResultIterator {
	return original.NewWorkspaceListResultIterator(page)
}
func NewWorkspaceListResultPage(cur WorkspaceListResult, getNextPage func(context.Context, WorkspaceListResult) (WorkspaceListResult, error)) WorkspaceListResultPage {
	return original.NewWorkspaceListResultPage(cur, getNextPage)
}
func NewWorkspaceSkusClient(subscriptionID string) WorkspaceSkusClient {
	return original.NewWorkspaceSkusClient(subscriptionID)
}
func NewWorkspaceSkusClientWithBaseURI(baseURI string, subscriptionID string) WorkspaceSkusClient {
	return original.NewWorkspaceSkusClientWithBaseURI(baseURI, subscriptionID)
}
func NewWorkspacesClient(subscriptionID string) WorkspacesClient {
	return original.NewWorkspacesClient(subscriptionID)
}
func NewWorkspacesClientWithBaseURI(baseURI string, subscriptionID string) WorkspacesClient {
	return original.NewWorkspacesClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleAllocationStateValues() []AllocationState {
	return original.PossibleAllocationStateValues()
}
func PossibleApplicationSharingPolicyValues() []ApplicationSharingPolicy {
	return original.PossibleApplicationSharingPolicyValues()
}
func PossibleClusterPurposeValues() []ClusterPurpose {
	return original.PossibleClusterPurposeValues()
}
func PossibleComputeInstanceAuthorizationTypeValues() []ComputeInstanceAuthorizationType {
	return original.PossibleComputeInstanceAuthorizationTypeValues()
}
func PossibleComputeInstanceStateValues() []ComputeInstanceState {
	return original.PossibleComputeInstanceStateValues()
}
func PossibleComputeTypeBasicComputeSecretsValues() []ComputeTypeBasicComputeSecrets {
	return original.PossibleComputeTypeBasicComputeSecretsValues()
}
func PossibleComputeTypeBasicComputeValues() []ComputeTypeBasicCompute {
	return original.PossibleComputeTypeBasicComputeValues()
}
func PossibleComputeTypeValues() []ComputeType {
	return original.PossibleComputeTypeValues()
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleDiagnoseResultLevelValues() []DiagnoseResultLevel {
	return original.PossibleDiagnoseResultLevelValues()
}
func PossibleEncryptionStatusValues() []EncryptionStatus {
	return original.PossibleEncryptionStatusValues()
}
func PossibleLoadBalancerTypeValues() []LoadBalancerType {
	return original.PossibleLoadBalancerTypeValues()
}
func PossibleNodeStateValues() []NodeState {
	return original.PossibleNodeStateValues()
}
func PossibleOperationNameValues() []OperationName {
	return original.PossibleOperationNameValues()
}
func PossibleOperationStatusValues() []OperationStatus {
	return original.PossibleOperationStatusValues()
}
func PossibleOsTypeValues() []OsType {
	return original.PossibleOsTypeValues()
}
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return original.PossiblePrivateEndpointConnectionProvisioningStateValues()
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
func PossibleQuotaUnitValues() []QuotaUnit {
	return original.PossibleQuotaUnitValues()
}
func PossibleReasonCodeValues() []ReasonCode {
	return original.PossibleReasonCodeValues()
}
func PossibleRemoteLoginPortPublicAccessValues() []RemoteLoginPortPublicAccess {
	return original.PossibleRemoteLoginPortPublicAccessValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func PossibleSSHPublicAccessValues() []SSHPublicAccess {
	return original.PossibleSSHPublicAccessValues()
}
func PossibleStatus1Values() []Status1 {
	return original.PossibleStatus1Values()
}
func PossibleStatusValues() []Status {
	return original.PossibleStatusValues()
}
func PossibleUnderlyingResourceActionValues() []UnderlyingResourceAction {
	return original.PossibleUnderlyingResourceActionValues()
}
func PossibleUsageUnitValues() []UsageUnit {
	return original.PossibleUsageUnitValues()
}
func PossibleVMPriceOSTypeValues() []VMPriceOSType {
	return original.PossibleVMPriceOSTypeValues()
}
func PossibleVMPriorityValues() []VMPriority {
	return original.PossibleVMPriorityValues()
}
func PossibleVMTierValues() []VMTier {
	return original.PossibleVMTierValues()
}
func PossibleValueFormatValues() []ValueFormat {
	return original.PossibleValueFormatValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
