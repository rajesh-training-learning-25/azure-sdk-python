//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package iothub

import (
	"context"

	original "github.com/Azure/dev/azure-sdk-for-go/services/provisioningservices/mgmt/2022-02-05/iothub"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AccessRightsDescription = original.AccessRightsDescription

const (
	AccessRightsDescriptionDeviceConnect           AccessRightsDescription = original.AccessRightsDescriptionDeviceConnect
	AccessRightsDescriptionEnrollmentRead          AccessRightsDescription = original.AccessRightsDescriptionEnrollmentRead
	AccessRightsDescriptionEnrollmentWrite         AccessRightsDescription = original.AccessRightsDescriptionEnrollmentWrite
	AccessRightsDescriptionRegistrationStatusRead  AccessRightsDescription = original.AccessRightsDescriptionRegistrationStatusRead
	AccessRightsDescriptionRegistrationStatusWrite AccessRightsDescription = original.AccessRightsDescriptionRegistrationStatusWrite
	AccessRightsDescriptionServiceConfig           AccessRightsDescription = original.AccessRightsDescriptionServiceConfig
)

type AllocationPolicy = original.AllocationPolicy

const (
	AllocationPolicyGeoLatency AllocationPolicy = original.AllocationPolicyGeoLatency
	AllocationPolicyHashed     AllocationPolicy = original.AllocationPolicyHashed
	AllocationPolicyStatic     AllocationPolicy = original.AllocationPolicyStatic
)

type CertificatePurpose = original.CertificatePurpose

const (
	CertificatePurposeClientAuthentication CertificatePurpose = original.CertificatePurposeClientAuthentication
	CertificatePurposeServerAuthentication CertificatePurpose = original.CertificatePurposeServerAuthentication
)

type CreatedByType = original.CreatedByType

const (
	CreatedByTypeApplication     CreatedByType = original.CreatedByTypeApplication
	CreatedByTypeKey             CreatedByType = original.CreatedByTypeKey
	CreatedByTypeManagedIdentity CreatedByType = original.CreatedByTypeManagedIdentity
	CreatedByTypeUser            CreatedByType = original.CreatedByTypeUser
)

type IPFilterActionType = original.IPFilterActionType

const (
	IPFilterActionTypeAccept IPFilterActionType = original.IPFilterActionTypeAccept
	IPFilterActionTypeReject IPFilterActionType = original.IPFilterActionTypeReject
)

type IPFilterTargetType = original.IPFilterTargetType

const (
	IPFilterTargetTypeAll        IPFilterTargetType = original.IPFilterTargetTypeAll
	IPFilterTargetTypeDeviceAPI  IPFilterTargetType = original.IPFilterTargetTypeDeviceAPI
	IPFilterTargetTypeServiceAPI IPFilterTargetType = original.IPFilterTargetTypeServiceAPI
)

type IotDpsSku = original.IotDpsSku

const (
	IotDpsSkuS1 IotDpsSku = original.IotDpsSkuS1
)

type NameUnavailabilityReason = original.NameUnavailabilityReason

const (
	NameUnavailabilityReasonAlreadyExists NameUnavailabilityReason = original.NameUnavailabilityReasonAlreadyExists
	NameUnavailabilityReasonInvalid       NameUnavailabilityReason = original.NameUnavailabilityReasonInvalid
)

type PrivateLinkServiceConnectionStatus = original.PrivateLinkServiceConnectionStatus

const (
	PrivateLinkServiceConnectionStatusApproved     PrivateLinkServiceConnectionStatus = original.PrivateLinkServiceConnectionStatusApproved
	PrivateLinkServiceConnectionStatusDisconnected PrivateLinkServiceConnectionStatus = original.PrivateLinkServiceConnectionStatusDisconnected
	PrivateLinkServiceConnectionStatusPending      PrivateLinkServiceConnectionStatus = original.PrivateLinkServiceConnectionStatusPending
	PrivateLinkServiceConnectionStatusRejected     PrivateLinkServiceConnectionStatus = original.PrivateLinkServiceConnectionStatusRejected
)

type PublicNetworkAccess = original.PublicNetworkAccess

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = original.PublicNetworkAccessDisabled
	PublicNetworkAccessEnabled  PublicNetworkAccess = original.PublicNetworkAccessEnabled
)

type State = original.State

const (
	StateActivating       State = original.StateActivating
	StateActivationFailed State = original.StateActivationFailed
	StateActive           State = original.StateActive
	StateDeleted          State = original.StateDeleted
	StateDeleting         State = original.StateDeleting
	StateDeletionFailed   State = original.StateDeletionFailed
	StateFailingOver      State = original.StateFailingOver
	StateFailoverFailed   State = original.StateFailoverFailed
	StateResuming         State = original.StateResuming
	StateSuspended        State = original.StateSuspended
	StateSuspending       State = original.StateSuspending
	StateTransitioning    State = original.StateTransitioning
)

type AsyncOperationResult = original.AsyncOperationResult
type BaseClient = original.BaseClient
type CertificateBodyDescription = original.CertificateBodyDescription
type CertificateListDescription = original.CertificateListDescription
type CertificateProperties = original.CertificateProperties
type CertificateResponse = original.CertificateResponse
type DefinitionDescription = original.DefinitionDescription
type DpsCertificateClient = original.DpsCertificateClient
type ErrorDetails = original.ErrorDetails
type ErrorMessage = original.ErrorMessage
type GroupIDInformation = original.GroupIDInformation
type GroupIDInformationProperties = original.GroupIDInformationProperties
type IPFilterRule = original.IPFilterRule
type IotDpsPropertiesDescription = original.IotDpsPropertiesDescription
type IotDpsResourceClient = original.IotDpsResourceClient
type IotDpsResourceCreateOrUpdateFuture = original.IotDpsResourceCreateOrUpdateFuture
type IotDpsResourceCreateOrUpdatePrivateEndpointConnectionFuture = original.IotDpsResourceCreateOrUpdatePrivateEndpointConnectionFuture
type IotDpsResourceDeleteFuture = original.IotDpsResourceDeleteFuture
type IotDpsResourceDeletePrivateEndpointConnectionFuture = original.IotDpsResourceDeletePrivateEndpointConnectionFuture
type IotDpsResourceUpdateFuture = original.IotDpsResourceUpdateFuture
type IotDpsSkuDefinition = original.IotDpsSkuDefinition
type IotDpsSkuDefinitionListResult = original.IotDpsSkuDefinitionListResult
type IotDpsSkuDefinitionListResultIterator = original.IotDpsSkuDefinitionListResultIterator
type IotDpsSkuDefinitionListResultPage = original.IotDpsSkuDefinitionListResultPage
type IotDpsSkuInfo = original.IotDpsSkuInfo
type ListPrivateEndpointConnection = original.ListPrivateEndpointConnection
type NameAvailabilityInfo = original.NameAvailabilityInfo
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationInputs = original.OperationInputs
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type PrivateEndpoint = original.PrivateEndpoint
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionProperties = original.PrivateEndpointConnectionProperties
type PrivateLinkResources = original.PrivateLinkResources
type PrivateLinkServiceConnectionState = original.PrivateLinkServiceConnectionState
type ProvisioningServiceDescription = original.ProvisioningServiceDescription
type ProvisioningServiceDescriptionListResult = original.ProvisioningServiceDescriptionListResult
type ProvisioningServiceDescriptionListResultIterator = original.ProvisioningServiceDescriptionListResultIterator
type ProvisioningServiceDescriptionListResultPage = original.ProvisioningServiceDescriptionListResultPage
type Resource = original.Resource
type SharedAccessSignatureAuthorizationRuleAccessRightsDescription = original.SharedAccessSignatureAuthorizationRuleAccessRightsDescription
type SharedAccessSignatureAuthorizationRuleListResult = original.SharedAccessSignatureAuthorizationRuleListResult
type SharedAccessSignatureAuthorizationRuleListResultIterator = original.SharedAccessSignatureAuthorizationRuleListResultIterator
type SharedAccessSignatureAuthorizationRuleListResultPage = original.SharedAccessSignatureAuthorizationRuleListResultPage
type SystemData = original.SystemData
type TagsResource = original.TagsResource
type VerificationCodeRequest = original.VerificationCodeRequest
type VerificationCodeResponse = original.VerificationCodeResponse
type VerificationCodeResponseProperties = original.VerificationCodeResponseProperties

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewDpsCertificateClient(subscriptionID string) DpsCertificateClient {
	return original.NewDpsCertificateClient(subscriptionID)
}
func NewDpsCertificateClientWithBaseURI(baseURI string, subscriptionID string) DpsCertificateClient {
	return original.NewDpsCertificateClientWithBaseURI(baseURI, subscriptionID)
}
func NewIotDpsResourceClient(subscriptionID string) IotDpsResourceClient {
	return original.NewIotDpsResourceClient(subscriptionID)
}
func NewIotDpsResourceClientWithBaseURI(baseURI string, subscriptionID string) IotDpsResourceClient {
	return original.NewIotDpsResourceClientWithBaseURI(baseURI, subscriptionID)
}
func NewIotDpsSkuDefinitionListResultIterator(page IotDpsSkuDefinitionListResultPage) IotDpsSkuDefinitionListResultIterator {
	return original.NewIotDpsSkuDefinitionListResultIterator(page)
}
func NewIotDpsSkuDefinitionListResultPage(cur IotDpsSkuDefinitionListResult, getNextPage func(context.Context, IotDpsSkuDefinitionListResult) (IotDpsSkuDefinitionListResult, error)) IotDpsSkuDefinitionListResultPage {
	return original.NewIotDpsSkuDefinitionListResultPage(cur, getNextPage)
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
func NewProvisioningServiceDescriptionListResultIterator(page ProvisioningServiceDescriptionListResultPage) ProvisioningServiceDescriptionListResultIterator {
	return original.NewProvisioningServiceDescriptionListResultIterator(page)
}
func NewProvisioningServiceDescriptionListResultPage(cur ProvisioningServiceDescriptionListResult, getNextPage func(context.Context, ProvisioningServiceDescriptionListResult) (ProvisioningServiceDescriptionListResult, error)) ProvisioningServiceDescriptionListResultPage {
	return original.NewProvisioningServiceDescriptionListResultPage(cur, getNextPage)
}
func NewSharedAccessSignatureAuthorizationRuleListResultIterator(page SharedAccessSignatureAuthorizationRuleListResultPage) SharedAccessSignatureAuthorizationRuleListResultIterator {
	return original.NewSharedAccessSignatureAuthorizationRuleListResultIterator(page)
}
func NewSharedAccessSignatureAuthorizationRuleListResultPage(cur SharedAccessSignatureAuthorizationRuleListResult, getNextPage func(context.Context, SharedAccessSignatureAuthorizationRuleListResult) (SharedAccessSignatureAuthorizationRuleListResult, error)) SharedAccessSignatureAuthorizationRuleListResultPage {
	return original.NewSharedAccessSignatureAuthorizationRuleListResultPage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAccessRightsDescriptionValues() []AccessRightsDescription {
	return original.PossibleAccessRightsDescriptionValues()
}
func PossibleAllocationPolicyValues() []AllocationPolicy {
	return original.PossibleAllocationPolicyValues()
}
func PossibleCertificatePurposeValues() []CertificatePurpose {
	return original.PossibleCertificatePurposeValues()
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleIPFilterActionTypeValues() []IPFilterActionType {
	return original.PossibleIPFilterActionTypeValues()
}
func PossibleIPFilterTargetTypeValues() []IPFilterTargetType {
	return original.PossibleIPFilterTargetTypeValues()
}
func PossibleIotDpsSkuValues() []IotDpsSku {
	return original.PossibleIotDpsSkuValues()
}
func PossibleNameUnavailabilityReasonValues() []NameUnavailabilityReason {
	return original.PossibleNameUnavailabilityReasonValues()
}
func PossiblePrivateLinkServiceConnectionStatusValues() []PrivateLinkServiceConnectionStatus {
	return original.PossiblePrivateLinkServiceConnectionStatusValues()
}
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return original.PossiblePublicNetworkAccessValues()
}
func PossibleStateValues() []State {
	return original.PossibleStateValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
