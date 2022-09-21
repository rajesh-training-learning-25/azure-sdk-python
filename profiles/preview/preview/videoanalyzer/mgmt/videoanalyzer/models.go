//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package videoanalyzer

import (
	"context"

	original "github.com/Azure/dev/azure-sdk-for-go/services/preview/videoanalyzer/mgmt/2021-11-01-preview/videoanalyzer"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AccessPolicyEccAlgo = original.AccessPolicyEccAlgo

const (
	AccessPolicyEccAlgoES256 AccessPolicyEccAlgo = original.AccessPolicyEccAlgoES256
	AccessPolicyEccAlgoES384 AccessPolicyEccAlgo = original.AccessPolicyEccAlgoES384
	AccessPolicyEccAlgoES512 AccessPolicyEccAlgo = original.AccessPolicyEccAlgoES512
)

type AccessPolicyRole = original.AccessPolicyRole

const (
	AccessPolicyRoleReader AccessPolicyRole = original.AccessPolicyRoleReader
)

type AccessPolicyRsaAlgo = original.AccessPolicyRsaAlgo

const (
	AccessPolicyRsaAlgoRS256 AccessPolicyRsaAlgo = original.AccessPolicyRsaAlgoRS256
	AccessPolicyRsaAlgoRS384 AccessPolicyRsaAlgo = original.AccessPolicyRsaAlgoRS384
	AccessPolicyRsaAlgoRS512 AccessPolicyRsaAlgo = original.AccessPolicyRsaAlgoRS512
)

type AccountEncryptionKeyType = original.AccountEncryptionKeyType

const (
	AccountEncryptionKeyTypeCustomerKey AccountEncryptionKeyType = original.AccountEncryptionKeyTypeCustomerKey
	AccountEncryptionKeyTypeSystemKey   AccountEncryptionKeyType = original.AccountEncryptionKeyTypeSystemKey
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

type EncoderSystemPresetType = original.EncoderSystemPresetType

const (
	EncoderSystemPresetTypeSingleLayer1080pH264AAC EncoderSystemPresetType = original.EncoderSystemPresetTypeSingleLayer1080pH264AAC
	EncoderSystemPresetTypeSingleLayer2160pH264AAC EncoderSystemPresetType = original.EncoderSystemPresetTypeSingleLayer2160pH264AAC
	EncoderSystemPresetTypeSingleLayer540pH264AAC  EncoderSystemPresetType = original.EncoderSystemPresetTypeSingleLayer540pH264AAC
	EncoderSystemPresetTypeSingleLayer720pH264AAC  EncoderSystemPresetType = original.EncoderSystemPresetTypeSingleLayer720pH264AAC
)

type Kind = original.Kind

const (
	KindBatch Kind = original.KindBatch
	KindLive  Kind = original.KindLive
)

type LivePipelineState = original.LivePipelineState

const (
	LivePipelineStateActivating   LivePipelineState = original.LivePipelineStateActivating
	LivePipelineStateActive       LivePipelineState = original.LivePipelineStateActive
	LivePipelineStateDeactivating LivePipelineState = original.LivePipelineStateDeactivating
	LivePipelineStateInactive     LivePipelineState = original.LivePipelineStateInactive
)

type MetricAggregationType = original.MetricAggregationType

const (
	MetricAggregationTypeAverage MetricAggregationType = original.MetricAggregationTypeAverage
	MetricAggregationTypeCount   MetricAggregationType = original.MetricAggregationTypeCount
	MetricAggregationTypeTotal   MetricAggregationType = original.MetricAggregationTypeTotal
)

type MetricUnit = original.MetricUnit

const (
	MetricUnitBytes        MetricUnit = original.MetricUnitBytes
	MetricUnitCount        MetricUnit = original.MetricUnitCount
	MetricUnitMilliseconds MetricUnit = original.MetricUnitMilliseconds
)

type ParameterType = original.ParameterType

const (
	ParameterTypeBool         ParameterType = original.ParameterTypeBool
	ParameterTypeDouble       ParameterType = original.ParameterTypeDouble
	ParameterTypeInt          ParameterType = original.ParameterTypeInt
	ParameterTypeSecretString ParameterType = original.ParameterTypeSecretString
	ParameterTypeString       ParameterType = original.ParameterTypeString
)

type PipelineJobState = original.PipelineJobState

const (
	PipelineJobStateCanceled   PipelineJobState = original.PipelineJobStateCanceled
	PipelineJobStateCompleted  PipelineJobState = original.PipelineJobStateCompleted
	PipelineJobStateFailed     PipelineJobState = original.PipelineJobStateFailed
	PipelineJobStateProcessing PipelineJobState = original.PipelineJobStateProcessing
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
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatusApproved
	PrivateEndpointServiceConnectionStatusPending  PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatusPending
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatusRejected
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateFailed     ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateInProgress ProvisioningState = original.ProvisioningStateInProgress
	ProvisioningStateSucceeded  ProvisioningState = original.ProvisioningStateSucceeded
)

type PublicNetworkAccess = original.PublicNetworkAccess

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = original.PublicNetworkAccessDisabled
	PublicNetworkAccessEnabled  PublicNetworkAccess = original.PublicNetworkAccessEnabled
)

type RtspTransport = original.RtspTransport

const (
	RtspTransportHTTP RtspTransport = original.RtspTransportHTTP
	RtspTransportTCP  RtspTransport = original.RtspTransportTCP
)

type SkuName = original.SkuName

const (
	SkuNameBatchS1 SkuName = original.SkuNameBatchS1
	SkuNameLiveS1  SkuName = original.SkuNameLiveS1
)

type SkuTier = original.SkuTier

const (
	SkuTierStandard SkuTier = original.SkuTierStandard
)

type Type = original.Type

const (
	TypeMicrosoftVideoAnalyzerEncoderProcessor  Type = original.TypeMicrosoftVideoAnalyzerEncoderProcessor
	TypeMicrosoftVideoAnalyzerProcessorNodeBase Type = original.TypeMicrosoftVideoAnalyzerProcessorNodeBase
	TypeMicrosoftVideoAnalyzerRtspSource        Type = original.TypeMicrosoftVideoAnalyzerRtspSource
	TypeMicrosoftVideoAnalyzerSinkNodeBase      Type = original.TypeMicrosoftVideoAnalyzerSinkNodeBase
	TypeMicrosoftVideoAnalyzerSourceNodeBase    Type = original.TypeMicrosoftVideoAnalyzerSourceNodeBase
	TypeMicrosoftVideoAnalyzerVideoSink         Type = original.TypeMicrosoftVideoAnalyzerVideoSink
	TypeMicrosoftVideoAnalyzerVideoSource       Type = original.TypeMicrosoftVideoAnalyzerVideoSource
	TypeNodeBase                                Type = original.TypeNodeBase
)

type TypeBasicAudioEncoderBase = original.TypeBasicAudioEncoderBase

const (
	TypeBasicAudioEncoderBaseTypeAudioEncoderBase                      TypeBasicAudioEncoderBase = original.TypeBasicAudioEncoderBaseTypeAudioEncoderBase
	TypeBasicAudioEncoderBaseTypeMicrosoftVideoAnalyzerAudioEncoderAac TypeBasicAudioEncoderBase = original.TypeBasicAudioEncoderBaseTypeMicrosoftVideoAnalyzerAudioEncoderAac
)

type TypeBasicAuthenticationBase = original.TypeBasicAuthenticationBase

const (
	TypeBasicAuthenticationBaseTypeAuthenticationBase                      TypeBasicAuthenticationBase = original.TypeBasicAuthenticationBaseTypeAuthenticationBase
	TypeBasicAuthenticationBaseTypeMicrosoftVideoAnalyzerJwtAuthentication TypeBasicAuthenticationBase = original.TypeBasicAuthenticationBaseTypeMicrosoftVideoAnalyzerJwtAuthentication
)

type TypeBasicCertificateSource = original.TypeBasicCertificateSource

const (
	TypeBasicCertificateSourceTypeCertificateSource                        TypeBasicCertificateSource = original.TypeBasicCertificateSourceTypeCertificateSource
	TypeBasicCertificateSourceTypeMicrosoftVideoAnalyzerPemCertificateList TypeBasicCertificateSource = original.TypeBasicCertificateSourceTypeMicrosoftVideoAnalyzerPemCertificateList
)

type TypeBasicCredentialsBase = original.TypeBasicCredentialsBase

const (
	TypeBasicCredentialsBaseTypeCredentialsBase                                   TypeBasicCredentialsBase = original.TypeBasicCredentialsBaseTypeCredentialsBase
	TypeBasicCredentialsBaseTypeMicrosoftVideoAnalyzerUsernamePasswordCredentials TypeBasicCredentialsBase = original.TypeBasicCredentialsBaseTypeMicrosoftVideoAnalyzerUsernamePasswordCredentials
)

type TypeBasicEncoderPresetBase = original.TypeBasicEncoderPresetBase

const (
	TypeBasicEncoderPresetBaseTypeEncoderPresetBase                         TypeBasicEncoderPresetBase = original.TypeBasicEncoderPresetBaseTypeEncoderPresetBase
	TypeBasicEncoderPresetBaseTypeMicrosoftVideoAnalyzerEncoderCustomPreset TypeBasicEncoderPresetBase = original.TypeBasicEncoderPresetBaseTypeMicrosoftVideoAnalyzerEncoderCustomPreset
	TypeBasicEncoderPresetBaseTypeMicrosoftVideoAnalyzerEncoderSystemPreset TypeBasicEncoderPresetBase = original.TypeBasicEncoderPresetBaseTypeMicrosoftVideoAnalyzerEncoderSystemPreset
)

type TypeBasicEndpointBase = original.TypeBasicEndpointBase

const (
	TypeBasicEndpointBaseTypeEndpointBase                            TypeBasicEndpointBase = original.TypeBasicEndpointBaseTypeEndpointBase
	TypeBasicEndpointBaseTypeMicrosoftVideoAnalyzerTLSEndpoint       TypeBasicEndpointBase = original.TypeBasicEndpointBaseTypeMicrosoftVideoAnalyzerTLSEndpoint
	TypeBasicEndpointBaseTypeMicrosoftVideoAnalyzerUnsecuredEndpoint TypeBasicEndpointBase = original.TypeBasicEndpointBaseTypeMicrosoftVideoAnalyzerUnsecuredEndpoint
)

type TypeBasicTimeSequenceBase = original.TypeBasicTimeSequenceBase

const (
	TypeBasicTimeSequenceBaseTypeMicrosoftVideoAnalyzerVideoSequenceAbsoluteTimeMarkers TypeBasicTimeSequenceBase = original.TypeBasicTimeSequenceBaseTypeMicrosoftVideoAnalyzerVideoSequenceAbsoluteTimeMarkers
	TypeBasicTimeSequenceBaseTypeTimeSequenceBase                                       TypeBasicTimeSequenceBase = original.TypeBasicTimeSequenceBaseTypeTimeSequenceBase
)

type TypeBasicTokenKey = original.TypeBasicTokenKey

const (
	TypeBasicTokenKeyTypeMicrosoftVideoAnalyzerEccTokenKey TypeBasicTokenKey = original.TypeBasicTokenKeyTypeMicrosoftVideoAnalyzerEccTokenKey
	TypeBasicTokenKeyTypeMicrosoftVideoAnalyzerRsaTokenKey TypeBasicTokenKey = original.TypeBasicTokenKeyTypeMicrosoftVideoAnalyzerRsaTokenKey
	TypeBasicTokenKeyTypeTokenKey                          TypeBasicTokenKey = original.TypeBasicTokenKeyTypeTokenKey
)

type TypeBasicTunnelBase = original.TypeBasicTunnelBase

const (
	TypeBasicTunnelBaseTypeMicrosoftVideoAnalyzerSecureIotDeviceRemoteTunnel TypeBasicTunnelBase = original.TypeBasicTunnelBaseTypeMicrosoftVideoAnalyzerSecureIotDeviceRemoteTunnel
	TypeBasicTunnelBaseTypeTunnelBase                                        TypeBasicTunnelBase = original.TypeBasicTunnelBaseTypeTunnelBase
)

type TypeBasicVideoEncoderBase = original.TypeBasicVideoEncoderBase

const (
	TypeBasicVideoEncoderBaseTypeMicrosoftVideoAnalyzerVideoEncoderH264 TypeBasicVideoEncoderBase = original.TypeBasicVideoEncoderBaseTypeMicrosoftVideoAnalyzerVideoEncoderH264
	TypeBasicVideoEncoderBaseTypeVideoEncoderBase                       TypeBasicVideoEncoderBase = original.TypeBasicVideoEncoderBaseTypeVideoEncoderBase
)

type VideoScaleMode = original.VideoScaleMode

const (
	VideoScaleModePad                 VideoScaleMode = original.VideoScaleModePad
	VideoScaleModePreserveAspectRatio VideoScaleMode = original.VideoScaleModePreserveAspectRatio
	VideoScaleModeStretch             VideoScaleMode = original.VideoScaleModeStretch
)

type VideoType = original.VideoType

const (
	VideoTypeArchive VideoType = original.VideoTypeArchive
	VideoTypeFile    VideoType = original.VideoTypeFile
)

type AccessPoliciesClient = original.AccessPoliciesClient
type AccessPolicyEntity = original.AccessPolicyEntity
type AccessPolicyEntityCollection = original.AccessPolicyEntityCollection
type AccessPolicyEntityCollectionIterator = original.AccessPolicyEntityCollectionIterator
type AccessPolicyEntityCollectionPage = original.AccessPolicyEntityCollectionPage
type AccessPolicyProperties = original.AccessPolicyProperties
type AccountEncryption = original.AccountEncryption
type AudioEncoderAac = original.AudioEncoderAac
type AudioEncoderBase = original.AudioEncoderBase
type AuthenticationBase = original.AuthenticationBase
type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type BasicAudioEncoderBase = original.BasicAudioEncoderBase
type BasicAuthenticationBase = original.BasicAuthenticationBase
type BasicCertificateSource = original.BasicCertificateSource
type BasicCredentialsBase = original.BasicCredentialsBase
type BasicEncoderPresetBase = original.BasicEncoderPresetBase
type BasicEndpointBase = original.BasicEndpointBase
type BasicNodeBase = original.BasicNodeBase
type BasicProcessorNodeBase = original.BasicProcessorNodeBase
type BasicSinkNodeBase = original.BasicSinkNodeBase
type BasicSourceNodeBase = original.BasicSourceNodeBase
type BasicTimeSequenceBase = original.BasicTimeSequenceBase
type BasicTokenKey = original.BasicTokenKey
type BasicTunnelBase = original.BasicTunnelBase
type BasicVideoEncoderBase = original.BasicVideoEncoderBase
type CertificateSource = original.CertificateSource
type CheckNameAvailabilityRequest = original.CheckNameAvailabilityRequest
type CheckNameAvailabilityResponse = original.CheckNameAvailabilityResponse
type Collection = original.Collection
type CredentialsBase = original.CredentialsBase
type EccTokenKey = original.EccTokenKey
type EdgeModuleEntity = original.EdgeModuleEntity
type EdgeModuleEntityCollection = original.EdgeModuleEntityCollection
type EdgeModuleEntityCollectionIterator = original.EdgeModuleEntityCollectionIterator
type EdgeModuleEntityCollectionPage = original.EdgeModuleEntityCollectionPage
type EdgeModuleProperties = original.EdgeModuleProperties
type EdgeModuleProvisioningToken = original.EdgeModuleProvisioningToken
type EdgeModulesClient = original.EdgeModulesClient
type EncoderCustomPreset = original.EncoderCustomPreset
type EncoderPresetBase = original.EncoderPresetBase
type EncoderProcessor = original.EncoderProcessor
type EncoderSystemPreset = original.EncoderSystemPreset
type Endpoint = original.Endpoint
type EndpointBase = original.EndpointBase
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type GroupLevelAccessControl = original.GroupLevelAccessControl
type Identity = original.Identity
type IotHub = original.IotHub
type JwtAuthentication = original.JwtAuthentication
type KeyVaultProperties = original.KeyVaultProperties
type ListProvisioningTokenInput = original.ListProvisioningTokenInput
type LivePipeline = original.LivePipeline
type LivePipelineCollection = original.LivePipelineCollection
type LivePipelineCollectionIterator = original.LivePipelineCollectionIterator
type LivePipelineCollectionPage = original.LivePipelineCollectionPage
type LivePipelineOperationStatus = original.LivePipelineOperationStatus
type LivePipelineOperationStatusesClient = original.LivePipelineOperationStatusesClient
type LivePipelineProperties = original.LivePipelineProperties
type LivePipelinePropertiesUpdate = original.LivePipelinePropertiesUpdate
type LivePipelineUpdate = original.LivePipelineUpdate
type LivePipelinesActivateFuture = original.LivePipelinesActivateFuture
type LivePipelinesClient = original.LivePipelinesClient
type LivePipelinesDeactivateFuture = original.LivePipelinesDeactivateFuture
type LocationsClient = original.LocationsClient
type LogSpecification = original.LogSpecification
type MetricDimension = original.MetricDimension
type MetricSpecification = original.MetricSpecification
type Model = original.Model
type NetworkAccessControl = original.NetworkAccessControl
type NodeBase = original.NodeBase
type NodeInput = original.NodeInput
type Operation = original.Operation
type OperationCollection = original.OperationCollection
type OperationDisplay = original.OperationDisplay
type OperationResultsClient = original.OperationResultsClient
type OperationResultsGroupClient = original.OperationResultsGroupClient
type OperationStatus = original.OperationStatus
type OperationStatusesClient = original.OperationStatusesClient
type OperationStatusesGroupClient = original.OperationStatusesGroupClient
type OperationsClient = original.OperationsClient
type ParameterDeclaration = original.ParameterDeclaration
type ParameterDefinition = original.ParameterDefinition
type PemCertificateList = original.PemCertificateList
type PipelineJob = original.PipelineJob
type PipelineJobCollection = original.PipelineJobCollection
type PipelineJobCollectionIterator = original.PipelineJobCollectionIterator
type PipelineJobCollectionPage = original.PipelineJobCollectionPage
type PipelineJobError = original.PipelineJobError
type PipelineJobOperationStatus = original.PipelineJobOperationStatus
type PipelineJobOperationStatusesClient = original.PipelineJobOperationStatusesClient
type PipelineJobProperties = original.PipelineJobProperties
type PipelineJobPropertiesUpdate = original.PipelineJobPropertiesUpdate
type PipelineJobUpdate = original.PipelineJobUpdate
type PipelineJobsCancelFuture = original.PipelineJobsCancelFuture
type PipelineJobsClient = original.PipelineJobsClient
type PipelineTopologiesClient = original.PipelineTopologiesClient
type PipelineTopology = original.PipelineTopology
type PipelineTopologyCollection = original.PipelineTopologyCollection
type PipelineTopologyCollectionIterator = original.PipelineTopologyCollectionIterator
type PipelineTopologyCollectionPage = original.PipelineTopologyCollectionPage
type PipelineTopologyProperties = original.PipelineTopologyProperties
type PipelineTopologyPropertiesUpdate = original.PipelineTopologyPropertiesUpdate
type PipelineTopologyUpdate = original.PipelineTopologyUpdate
type PrivateEndpoint = original.PrivateEndpoint
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionListResult = original.PrivateEndpointConnectionListResult
type PrivateEndpointConnectionOperationStatus = original.PrivateEndpointConnectionOperationStatus
type PrivateEndpointConnectionProperties = original.PrivateEndpointConnectionProperties
type PrivateEndpointConnectionsClient = original.PrivateEndpointConnectionsClient
type PrivateLinkResource = original.PrivateLinkResource
type PrivateLinkResourceListResult = original.PrivateLinkResourceListResult
type PrivateLinkResourceProperties = original.PrivateLinkResourceProperties
type PrivateLinkResourcesClient = original.PrivateLinkResourcesClient
type PrivateLinkServiceConnectionState = original.PrivateLinkServiceConnectionState
type ProcessorNodeBase = original.ProcessorNodeBase
type Properties = original.Properties
type PropertiesType = original.PropertiesType
type PropertiesUpdate = original.PropertiesUpdate
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type ResourceIdentity = original.ResourceIdentity
type RsaTokenKey = original.RsaTokenKey
type RtspSource = original.RtspSource
type SecureIotDeviceRemoteTunnel = original.SecureIotDeviceRemoteTunnel
type ServiceSpecification = original.ServiceSpecification
type SinkNodeBase = original.SinkNodeBase
type Sku = original.Sku
type SourceNodeBase = original.SourceNodeBase
type StorageAccount = original.StorageAccount
type SystemData = original.SystemData
type TLSEndpoint = original.TLSEndpoint
type TLSValidationOptions = original.TLSValidationOptions
type TimeSequenceBase = original.TimeSequenceBase
type TokenClaim = original.TokenClaim
type TokenKey = original.TokenKey
type TrackedResource = original.TrackedResource
type TunnelBase = original.TunnelBase
type UnsecuredEndpoint = original.UnsecuredEndpoint
type Update = original.Update
type UserAssignedManagedIdentity = original.UserAssignedManagedIdentity
type UsernamePasswordCredentials = original.UsernamePasswordCredentials
type VideoAnalyzersClient = original.VideoAnalyzersClient
type VideoAnalyzersCreateOrUpdateFutureType = original.VideoAnalyzersCreateOrUpdateFutureType
type VideoAnalyzersUpdateFutureType = original.VideoAnalyzersUpdateFutureType
type VideoArchival = original.VideoArchival
type VideoContentToken = original.VideoContentToken
type VideoContentUrls = original.VideoContentUrls
type VideoCreationProperties = original.VideoCreationProperties
type VideoEncoderBase = original.VideoEncoderBase
type VideoEncoderH264 = original.VideoEncoderH264
type VideoEntity = original.VideoEntity
type VideoEntityCollection = original.VideoEntityCollection
type VideoEntityCollectionIterator = original.VideoEntityCollectionIterator
type VideoEntityCollectionPage = original.VideoEntityCollectionPage
type VideoFlags = original.VideoFlags
type VideoMediaInfo = original.VideoMediaInfo
type VideoPreviewImageUrls = original.VideoPreviewImageUrls
type VideoProperties = original.VideoProperties
type VideoPublishingOptions = original.VideoPublishingOptions
type VideoScale = original.VideoScale
type VideoSequenceAbsoluteTimeMarkers = original.VideoSequenceAbsoluteTimeMarkers
type VideoSink = original.VideoSink
type VideoSource = original.VideoSource
type VideosClient = original.VideosClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAccessPoliciesClient(subscriptionID string) AccessPoliciesClient {
	return original.NewAccessPoliciesClient(subscriptionID)
}
func NewAccessPoliciesClientWithBaseURI(baseURI string, subscriptionID string) AccessPoliciesClient {
	return original.NewAccessPoliciesClientWithBaseURI(baseURI, subscriptionID)
}
func NewAccessPolicyEntityCollectionIterator(page AccessPolicyEntityCollectionPage) AccessPolicyEntityCollectionIterator {
	return original.NewAccessPolicyEntityCollectionIterator(page)
}
func NewAccessPolicyEntityCollectionPage(cur AccessPolicyEntityCollection, getNextPage func(context.Context, AccessPolicyEntityCollection) (AccessPolicyEntityCollection, error)) AccessPolicyEntityCollectionPage {
	return original.NewAccessPolicyEntityCollectionPage(cur, getNextPage)
}
func NewEdgeModuleEntityCollectionIterator(page EdgeModuleEntityCollectionPage) EdgeModuleEntityCollectionIterator {
	return original.NewEdgeModuleEntityCollectionIterator(page)
}
func NewEdgeModuleEntityCollectionPage(cur EdgeModuleEntityCollection, getNextPage func(context.Context, EdgeModuleEntityCollection) (EdgeModuleEntityCollection, error)) EdgeModuleEntityCollectionPage {
	return original.NewEdgeModuleEntityCollectionPage(cur, getNextPage)
}
func NewEdgeModulesClient(subscriptionID string) EdgeModulesClient {
	return original.NewEdgeModulesClient(subscriptionID)
}
func NewEdgeModulesClientWithBaseURI(baseURI string, subscriptionID string) EdgeModulesClient {
	return original.NewEdgeModulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewLivePipelineCollectionIterator(page LivePipelineCollectionPage) LivePipelineCollectionIterator {
	return original.NewLivePipelineCollectionIterator(page)
}
func NewLivePipelineCollectionPage(cur LivePipelineCollection, getNextPage func(context.Context, LivePipelineCollection) (LivePipelineCollection, error)) LivePipelineCollectionPage {
	return original.NewLivePipelineCollectionPage(cur, getNextPage)
}
func NewLivePipelineOperationStatusesClient(subscriptionID string) LivePipelineOperationStatusesClient {
	return original.NewLivePipelineOperationStatusesClient(subscriptionID)
}
func NewLivePipelineOperationStatusesClientWithBaseURI(baseURI string, subscriptionID string) LivePipelineOperationStatusesClient {
	return original.NewLivePipelineOperationStatusesClientWithBaseURI(baseURI, subscriptionID)
}
func NewLivePipelinesClient(subscriptionID string) LivePipelinesClient {
	return original.NewLivePipelinesClient(subscriptionID)
}
func NewLivePipelinesClientWithBaseURI(baseURI string, subscriptionID string) LivePipelinesClient {
	return original.NewLivePipelinesClientWithBaseURI(baseURI, subscriptionID)
}
func NewLocationsClient(subscriptionID string) LocationsClient {
	return original.NewLocationsClient(subscriptionID)
}
func NewLocationsClientWithBaseURI(baseURI string, subscriptionID string) LocationsClient {
	return original.NewLocationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationResultsClient(subscriptionID string) OperationResultsClient {
	return original.NewOperationResultsClient(subscriptionID)
}
func NewOperationResultsClientWithBaseURI(baseURI string, subscriptionID string) OperationResultsClient {
	return original.NewOperationResultsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationResultsGroupClient(subscriptionID string) OperationResultsGroupClient {
	return original.NewOperationResultsGroupClient(subscriptionID)
}
func NewOperationResultsGroupClientWithBaseURI(baseURI string, subscriptionID string) OperationResultsGroupClient {
	return original.NewOperationResultsGroupClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationStatusesClient(subscriptionID string) OperationStatusesClient {
	return original.NewOperationStatusesClient(subscriptionID)
}
func NewOperationStatusesClientWithBaseURI(baseURI string, subscriptionID string) OperationStatusesClient {
	return original.NewOperationStatusesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationStatusesGroupClient(subscriptionID string) OperationStatusesGroupClient {
	return original.NewOperationStatusesGroupClient(subscriptionID)
}
func NewOperationStatusesGroupClientWithBaseURI(baseURI string, subscriptionID string) OperationStatusesGroupClient {
	return original.NewOperationStatusesGroupClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPipelineJobCollectionIterator(page PipelineJobCollectionPage) PipelineJobCollectionIterator {
	return original.NewPipelineJobCollectionIterator(page)
}
func NewPipelineJobCollectionPage(cur PipelineJobCollection, getNextPage func(context.Context, PipelineJobCollection) (PipelineJobCollection, error)) PipelineJobCollectionPage {
	return original.NewPipelineJobCollectionPage(cur, getNextPage)
}
func NewPipelineJobOperationStatusesClient(subscriptionID string) PipelineJobOperationStatusesClient {
	return original.NewPipelineJobOperationStatusesClient(subscriptionID)
}
func NewPipelineJobOperationStatusesClientWithBaseURI(baseURI string, subscriptionID string) PipelineJobOperationStatusesClient {
	return original.NewPipelineJobOperationStatusesClientWithBaseURI(baseURI, subscriptionID)
}
func NewPipelineJobsClient(subscriptionID string) PipelineJobsClient {
	return original.NewPipelineJobsClient(subscriptionID)
}
func NewPipelineJobsClientWithBaseURI(baseURI string, subscriptionID string) PipelineJobsClient {
	return original.NewPipelineJobsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPipelineTopologiesClient(subscriptionID string) PipelineTopologiesClient {
	return original.NewPipelineTopologiesClient(subscriptionID)
}
func NewPipelineTopologiesClientWithBaseURI(baseURI string, subscriptionID string) PipelineTopologiesClient {
	return original.NewPipelineTopologiesClientWithBaseURI(baseURI, subscriptionID)
}
func NewPipelineTopologyCollectionIterator(page PipelineTopologyCollectionPage) PipelineTopologyCollectionIterator {
	return original.NewPipelineTopologyCollectionIterator(page)
}
func NewPipelineTopologyCollectionPage(cur PipelineTopologyCollection, getNextPage func(context.Context, PipelineTopologyCollection) (PipelineTopologyCollection, error)) PipelineTopologyCollectionPage {
	return original.NewPipelineTopologyCollectionPage(cur, getNextPage)
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
func NewVideoAnalyzersClient(subscriptionID string) VideoAnalyzersClient {
	return original.NewVideoAnalyzersClient(subscriptionID)
}
func NewVideoAnalyzersClientWithBaseURI(baseURI string, subscriptionID string) VideoAnalyzersClient {
	return original.NewVideoAnalyzersClientWithBaseURI(baseURI, subscriptionID)
}
func NewVideoEntityCollectionIterator(page VideoEntityCollectionPage) VideoEntityCollectionIterator {
	return original.NewVideoEntityCollectionIterator(page)
}
func NewVideoEntityCollectionPage(cur VideoEntityCollection, getNextPage func(context.Context, VideoEntityCollection) (VideoEntityCollection, error)) VideoEntityCollectionPage {
	return original.NewVideoEntityCollectionPage(cur, getNextPage)
}
func NewVideosClient(subscriptionID string) VideosClient {
	return original.NewVideosClient(subscriptionID)
}
func NewVideosClientWithBaseURI(baseURI string, subscriptionID string) VideosClient {
	return original.NewVideosClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAccessPolicyEccAlgoValues() []AccessPolicyEccAlgo {
	return original.PossibleAccessPolicyEccAlgoValues()
}
func PossibleAccessPolicyRoleValues() []AccessPolicyRole {
	return original.PossibleAccessPolicyRoleValues()
}
func PossibleAccessPolicyRsaAlgoValues() []AccessPolicyRsaAlgo {
	return original.PossibleAccessPolicyRsaAlgoValues()
}
func PossibleAccountEncryptionKeyTypeValues() []AccountEncryptionKeyType {
	return original.PossibleAccountEncryptionKeyTypeValues()
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
func PossibleEncoderSystemPresetTypeValues() []EncoderSystemPresetType {
	return original.PossibleEncoderSystemPresetTypeValues()
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleLivePipelineStateValues() []LivePipelineState {
	return original.PossibleLivePipelineStateValues()
}
func PossibleMetricAggregationTypeValues() []MetricAggregationType {
	return original.PossibleMetricAggregationTypeValues()
}
func PossibleMetricUnitValues() []MetricUnit {
	return original.PossibleMetricUnitValues()
}
func PossibleParameterTypeValues() []ParameterType {
	return original.PossibleParameterTypeValues()
}
func PossiblePipelineJobStateValues() []PipelineJobState {
	return original.PossiblePipelineJobStateValues()
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
func PossibleRtspTransportValues() []RtspTransport {
	return original.PossibleRtspTransportValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func PossibleTypeBasicAudioEncoderBaseValues() []TypeBasicAudioEncoderBase {
	return original.PossibleTypeBasicAudioEncoderBaseValues()
}
func PossibleTypeBasicAuthenticationBaseValues() []TypeBasicAuthenticationBase {
	return original.PossibleTypeBasicAuthenticationBaseValues()
}
func PossibleTypeBasicCertificateSourceValues() []TypeBasicCertificateSource {
	return original.PossibleTypeBasicCertificateSourceValues()
}
func PossibleTypeBasicCredentialsBaseValues() []TypeBasicCredentialsBase {
	return original.PossibleTypeBasicCredentialsBaseValues()
}
func PossibleTypeBasicEncoderPresetBaseValues() []TypeBasicEncoderPresetBase {
	return original.PossibleTypeBasicEncoderPresetBaseValues()
}
func PossibleTypeBasicEndpointBaseValues() []TypeBasicEndpointBase {
	return original.PossibleTypeBasicEndpointBaseValues()
}
func PossibleTypeBasicTimeSequenceBaseValues() []TypeBasicTimeSequenceBase {
	return original.PossibleTypeBasicTimeSequenceBaseValues()
}
func PossibleTypeBasicTokenKeyValues() []TypeBasicTokenKey {
	return original.PossibleTypeBasicTokenKeyValues()
}
func PossibleTypeBasicTunnelBaseValues() []TypeBasicTunnelBase {
	return original.PossibleTypeBasicTunnelBaseValues()
}
func PossibleTypeBasicVideoEncoderBaseValues() []TypeBasicVideoEncoderBase {
	return original.PossibleTypeBasicVideoEncoderBaseValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func PossibleVideoScaleModeValues() []VideoScaleMode {
	return original.PossibleVideoScaleModeValues()
}
func PossibleVideoTypeValues() []VideoType {
	return original.PossibleVideoTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
