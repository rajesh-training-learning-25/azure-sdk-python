//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package authorization

import (
	"context"

	original "github.com/Azure/dev/azure-sdk-for-go/services/authorization/mgmt/2020-10-01/authorization"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ApprovalMode = original.ApprovalMode

const (
	ApprovalModeNoApproval  ApprovalMode = original.ApprovalModeNoApproval
	ApprovalModeParallel    ApprovalMode = original.ApprovalModeParallel
	ApprovalModeSerial      ApprovalMode = original.ApprovalModeSerial
	ApprovalModeSingleStage ApprovalMode = original.ApprovalModeSingleStage
)

type AssignmentType = original.AssignmentType

const (
	AssignmentTypeActivated AssignmentType = original.AssignmentTypeActivated
	AssignmentTypeAssigned  AssignmentType = original.AssignmentTypeAssigned
)

type EnablementRules = original.EnablementRules

const (
	EnablementRulesJustification             EnablementRules = original.EnablementRulesJustification
	EnablementRulesMultiFactorAuthentication EnablementRules = original.EnablementRulesMultiFactorAuthentication
	EnablementRulesTicketing                 EnablementRules = original.EnablementRulesTicketing
)

type MemberType = original.MemberType

const (
	MemberTypeDirect    MemberType = original.MemberTypeDirect
	MemberTypeGroup     MemberType = original.MemberTypeGroup
	MemberTypeInherited MemberType = original.MemberTypeInherited
)

type NotificationDeliveryMechanism = original.NotificationDeliveryMechanism

const (
	NotificationDeliveryMechanismEmail NotificationDeliveryMechanism = original.NotificationDeliveryMechanismEmail
)

type NotificationLevel = original.NotificationLevel

const (
	NotificationLevelAll      NotificationLevel = original.NotificationLevelAll
	NotificationLevelCritical NotificationLevel = original.NotificationLevelCritical
	NotificationLevelNone     NotificationLevel = original.NotificationLevelNone
)

type PrincipalType = original.PrincipalType

const (
	PrincipalTypeDevice           PrincipalType = original.PrincipalTypeDevice
	PrincipalTypeForeignGroup     PrincipalType = original.PrincipalTypeForeignGroup
	PrincipalTypeGroup            PrincipalType = original.PrincipalTypeGroup
	PrincipalTypeServicePrincipal PrincipalType = original.PrincipalTypeServicePrincipal
	PrincipalTypeUser             PrincipalType = original.PrincipalTypeUser
)

type RecipientType = original.RecipientType

const (
	RecipientTypeAdmin     RecipientType = original.RecipientTypeAdmin
	RecipientTypeApprover  RecipientType = original.RecipientTypeApprover
	RecipientTypeRequestor RecipientType = original.RecipientTypeRequestor
)

type RequestType = original.RequestType

const (
	RequestTypeAdminAssign    RequestType = original.RequestTypeAdminAssign
	RequestTypeAdminExtend    RequestType = original.RequestTypeAdminExtend
	RequestTypeAdminRemove    RequestType = original.RequestTypeAdminRemove
	RequestTypeAdminRenew     RequestType = original.RequestTypeAdminRenew
	RequestTypeAdminUpdate    RequestType = original.RequestTypeAdminUpdate
	RequestTypeSelfActivate   RequestType = original.RequestTypeSelfActivate
	RequestTypeSelfDeactivate RequestType = original.RequestTypeSelfDeactivate
	RequestTypeSelfExtend     RequestType = original.RequestTypeSelfExtend
	RequestTypeSelfRenew      RequestType = original.RequestTypeSelfRenew
)

type RoleManagementPolicyRuleType = original.RoleManagementPolicyRuleType

const (
	RoleManagementPolicyRuleTypeRoleManagementPolicyApprovalRule              RoleManagementPolicyRuleType = original.RoleManagementPolicyRuleTypeRoleManagementPolicyApprovalRule
	RoleManagementPolicyRuleTypeRoleManagementPolicyAuthenticationContextRule RoleManagementPolicyRuleType = original.RoleManagementPolicyRuleTypeRoleManagementPolicyAuthenticationContextRule
	RoleManagementPolicyRuleTypeRoleManagementPolicyEnablementRule            RoleManagementPolicyRuleType = original.RoleManagementPolicyRuleTypeRoleManagementPolicyEnablementRule
	RoleManagementPolicyRuleTypeRoleManagementPolicyExpirationRule            RoleManagementPolicyRuleType = original.RoleManagementPolicyRuleTypeRoleManagementPolicyExpirationRule
	RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule          RoleManagementPolicyRuleType = original.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule
)

type RuleType = original.RuleType

const (
	RuleTypeRoleManagementPolicyApprovalRule              RuleType = original.RuleTypeRoleManagementPolicyApprovalRule
	RuleTypeRoleManagementPolicyAuthenticationContextRule RuleType = original.RuleTypeRoleManagementPolicyAuthenticationContextRule
	RuleTypeRoleManagementPolicyEnablementRule            RuleType = original.RuleTypeRoleManagementPolicyEnablementRule
	RuleTypeRoleManagementPolicyExpirationRule            RuleType = original.RuleTypeRoleManagementPolicyExpirationRule
	RuleTypeRoleManagementPolicyNotificationRule          RuleType = original.RuleTypeRoleManagementPolicyNotificationRule
	RuleTypeRoleManagementPolicyRule                      RuleType = original.RuleTypeRoleManagementPolicyRule
)

type Status = original.Status

const (
	StatusAccepted                    Status = original.StatusAccepted
	StatusAdminApproved               Status = original.StatusAdminApproved
	StatusAdminDenied                 Status = original.StatusAdminDenied
	StatusCanceled                    Status = original.StatusCanceled
	StatusDenied                      Status = original.StatusDenied
	StatusFailed                      Status = original.StatusFailed
	StatusFailedAsResourceIsLocked    Status = original.StatusFailedAsResourceIsLocked
	StatusGranted                     Status = original.StatusGranted
	StatusInvalid                     Status = original.StatusInvalid
	StatusPendingAdminDecision        Status = original.StatusPendingAdminDecision
	StatusPendingApproval             Status = original.StatusPendingApproval
	StatusPendingApprovalProvisioning Status = original.StatusPendingApprovalProvisioning
	StatusPendingEvaluation           Status = original.StatusPendingEvaluation
	StatusPendingExternalProvisioning Status = original.StatusPendingExternalProvisioning
	StatusPendingProvisioning         Status = original.StatusPendingProvisioning
	StatusPendingRevocation           Status = original.StatusPendingRevocation
	StatusPendingScheduleCreation     Status = original.StatusPendingScheduleCreation
	StatusProvisioned                 Status = original.StatusProvisioned
	StatusProvisioningStarted         Status = original.StatusProvisioningStarted
	StatusRevoked                     Status = original.StatusRevoked
	StatusScheduleCreated             Status = original.StatusScheduleCreated
	StatusTimedOut                    Status = original.StatusTimedOut
)

type Type = original.Type

const (
	TypeAfterDateTime Type = original.TypeAfterDateTime
	TypeAfterDuration Type = original.TypeAfterDuration
	TypeNoExpiration  Type = original.TypeNoExpiration
)

type UserType = original.UserType

const (
	UserTypeGroup UserType = original.UserTypeGroup
	UserTypeUser  UserType = original.UserTypeUser
)

type ApprovalSettings = original.ApprovalSettings
type ApprovalStage = original.ApprovalStage
type BaseClient = original.BaseClient
type BasicRoleManagementPolicyRule = original.BasicRoleManagementPolicyRule
type ClassicAdministrator = original.ClassicAdministrator
type ClassicAdministratorListResult = original.ClassicAdministratorListResult
type ClassicAdministratorListResultIterator = original.ClassicAdministratorListResultIterator
type ClassicAdministratorListResultPage = original.ClassicAdministratorListResultPage
type ClassicAdministratorProperties = original.ClassicAdministratorProperties
type ClassicAdministratorsClient = original.ClassicAdministratorsClient
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type EligibleChildResource = original.EligibleChildResource
type EligibleChildResourcesClient = original.EligibleChildResourcesClient
type EligibleChildResourcesListResult = original.EligibleChildResourcesListResult
type EligibleChildResourcesListResultIterator = original.EligibleChildResourcesListResultIterator
type EligibleChildResourcesListResultPage = original.EligibleChildResourcesListResultPage
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type ExpandedProperties = original.ExpandedProperties
type ExpandedPropertiesPrincipal = original.ExpandedPropertiesPrincipal
type ExpandedPropertiesRoleDefinition = original.ExpandedPropertiesRoleDefinition
type ExpandedPropertiesScope = original.ExpandedPropertiesScope
type GlobalAdministratorClient = original.GlobalAdministratorClient
type Permission = original.Permission
type PermissionGetResult = original.PermissionGetResult
type PermissionGetResultIterator = original.PermissionGetResultIterator
type PermissionGetResultPage = original.PermissionGetResultPage
type PermissionsClient = original.PermissionsClient
type PolicyAssignmentProperties = original.PolicyAssignmentProperties
type PolicyAssignmentPropertiesPolicy = original.PolicyAssignmentPropertiesPolicy
type PolicyAssignmentPropertiesRoleDefinition = original.PolicyAssignmentPropertiesRoleDefinition
type PolicyAssignmentPropertiesScope = original.PolicyAssignmentPropertiesScope
type PolicyProperties = original.PolicyProperties
type PolicyPropertiesScope = original.PolicyPropertiesScope
type Principal = original.Principal
type ProviderOperation = original.ProviderOperation
type ProviderOperationsMetadata = original.ProviderOperationsMetadata
type ProviderOperationsMetadataClient = original.ProviderOperationsMetadataClient
type ProviderOperationsMetadataListResult = original.ProviderOperationsMetadataListResult
type ProviderOperationsMetadataListResultIterator = original.ProviderOperationsMetadataListResultIterator
type ProviderOperationsMetadataListResultPage = original.ProviderOperationsMetadataListResultPage
type ResourceType = original.ResourceType
type RoleAssignment = original.RoleAssignment
type RoleAssignmentCreateParameters = original.RoleAssignmentCreateParameters
type RoleAssignmentFilter = original.RoleAssignmentFilter
type RoleAssignmentListResult = original.RoleAssignmentListResult
type RoleAssignmentListResultIterator = original.RoleAssignmentListResultIterator
type RoleAssignmentListResultPage = original.RoleAssignmentListResultPage
type RoleAssignmentProperties = original.RoleAssignmentProperties
type RoleAssignmentPropertiesWithScope = original.RoleAssignmentPropertiesWithScope
type RoleAssignmentSchedule = original.RoleAssignmentSchedule
type RoleAssignmentScheduleFilter = original.RoleAssignmentScheduleFilter
type RoleAssignmentScheduleInstance = original.RoleAssignmentScheduleInstance
type RoleAssignmentScheduleInstanceFilter = original.RoleAssignmentScheduleInstanceFilter
type RoleAssignmentScheduleInstanceListResult = original.RoleAssignmentScheduleInstanceListResult
type RoleAssignmentScheduleInstanceListResultIterator = original.RoleAssignmentScheduleInstanceListResultIterator
type RoleAssignmentScheduleInstanceListResultPage = original.RoleAssignmentScheduleInstanceListResultPage
type RoleAssignmentScheduleInstanceProperties = original.RoleAssignmentScheduleInstanceProperties
type RoleAssignmentScheduleInstancesClient = original.RoleAssignmentScheduleInstancesClient
type RoleAssignmentScheduleListResult = original.RoleAssignmentScheduleListResult
type RoleAssignmentScheduleListResultIterator = original.RoleAssignmentScheduleListResultIterator
type RoleAssignmentScheduleListResultPage = original.RoleAssignmentScheduleListResultPage
type RoleAssignmentScheduleProperties = original.RoleAssignmentScheduleProperties
type RoleAssignmentScheduleRequest = original.RoleAssignmentScheduleRequest
type RoleAssignmentScheduleRequestFilter = original.RoleAssignmentScheduleRequestFilter
type RoleAssignmentScheduleRequestListResult = original.RoleAssignmentScheduleRequestListResult
type RoleAssignmentScheduleRequestListResultIterator = original.RoleAssignmentScheduleRequestListResultIterator
type RoleAssignmentScheduleRequestListResultPage = original.RoleAssignmentScheduleRequestListResultPage
type RoleAssignmentScheduleRequestProperties = original.RoleAssignmentScheduleRequestProperties
type RoleAssignmentScheduleRequestPropertiesScheduleInfo = original.RoleAssignmentScheduleRequestPropertiesScheduleInfo
type RoleAssignmentScheduleRequestPropertiesScheduleInfoExpiration = original.RoleAssignmentScheduleRequestPropertiesScheduleInfoExpiration
type RoleAssignmentScheduleRequestPropertiesTicketInfo = original.RoleAssignmentScheduleRequestPropertiesTicketInfo
type RoleAssignmentScheduleRequestsClient = original.RoleAssignmentScheduleRequestsClient
type RoleAssignmentSchedulesClient = original.RoleAssignmentSchedulesClient
type RoleAssignmentsClient = original.RoleAssignmentsClient
type RoleDefinition = original.RoleDefinition
type RoleDefinitionFilter = original.RoleDefinitionFilter
type RoleDefinitionListResult = original.RoleDefinitionListResult
type RoleDefinitionListResultIterator = original.RoleDefinitionListResultIterator
type RoleDefinitionListResultPage = original.RoleDefinitionListResultPage
type RoleDefinitionProperties = original.RoleDefinitionProperties
type RoleDefinitionsClient = original.RoleDefinitionsClient
type RoleEligibilitySchedule = original.RoleEligibilitySchedule
type RoleEligibilityScheduleFilter = original.RoleEligibilityScheduleFilter
type RoleEligibilityScheduleInstance = original.RoleEligibilityScheduleInstance
type RoleEligibilityScheduleInstanceFilter = original.RoleEligibilityScheduleInstanceFilter
type RoleEligibilityScheduleInstanceListResult = original.RoleEligibilityScheduleInstanceListResult
type RoleEligibilityScheduleInstanceListResultIterator = original.RoleEligibilityScheduleInstanceListResultIterator
type RoleEligibilityScheduleInstanceListResultPage = original.RoleEligibilityScheduleInstanceListResultPage
type RoleEligibilityScheduleInstanceProperties = original.RoleEligibilityScheduleInstanceProperties
type RoleEligibilityScheduleInstancesClient = original.RoleEligibilityScheduleInstancesClient
type RoleEligibilityScheduleListResult = original.RoleEligibilityScheduleListResult
type RoleEligibilityScheduleListResultIterator = original.RoleEligibilityScheduleListResultIterator
type RoleEligibilityScheduleListResultPage = original.RoleEligibilityScheduleListResultPage
type RoleEligibilityScheduleProperties = original.RoleEligibilityScheduleProperties
type RoleEligibilityScheduleRequest = original.RoleEligibilityScheduleRequest
type RoleEligibilityScheduleRequestFilter = original.RoleEligibilityScheduleRequestFilter
type RoleEligibilityScheduleRequestListResult = original.RoleEligibilityScheduleRequestListResult
type RoleEligibilityScheduleRequestListResultIterator = original.RoleEligibilityScheduleRequestListResultIterator
type RoleEligibilityScheduleRequestListResultPage = original.RoleEligibilityScheduleRequestListResultPage
type RoleEligibilityScheduleRequestProperties = original.RoleEligibilityScheduleRequestProperties
type RoleEligibilityScheduleRequestPropertiesScheduleInfo = original.RoleEligibilityScheduleRequestPropertiesScheduleInfo
type RoleEligibilityScheduleRequestPropertiesScheduleInfoExpiration = original.RoleEligibilityScheduleRequestPropertiesScheduleInfoExpiration
type RoleEligibilityScheduleRequestPropertiesTicketInfo = original.RoleEligibilityScheduleRequestPropertiesTicketInfo
type RoleEligibilityScheduleRequestsClient = original.RoleEligibilityScheduleRequestsClient
type RoleEligibilitySchedulesClient = original.RoleEligibilitySchedulesClient
type RoleManagementPoliciesClient = original.RoleManagementPoliciesClient
type RoleManagementPolicy = original.RoleManagementPolicy
type RoleManagementPolicyApprovalRule = original.RoleManagementPolicyApprovalRule
type RoleManagementPolicyAssignment = original.RoleManagementPolicyAssignment
type RoleManagementPolicyAssignmentListResult = original.RoleManagementPolicyAssignmentListResult
type RoleManagementPolicyAssignmentListResultIterator = original.RoleManagementPolicyAssignmentListResultIterator
type RoleManagementPolicyAssignmentListResultPage = original.RoleManagementPolicyAssignmentListResultPage
type RoleManagementPolicyAssignmentProperties = original.RoleManagementPolicyAssignmentProperties
type RoleManagementPolicyAssignmentsClient = original.RoleManagementPolicyAssignmentsClient
type RoleManagementPolicyAuthenticationContextRule = original.RoleManagementPolicyAuthenticationContextRule
type RoleManagementPolicyEnablementRule = original.RoleManagementPolicyEnablementRule
type RoleManagementPolicyExpirationRule = original.RoleManagementPolicyExpirationRule
type RoleManagementPolicyListResult = original.RoleManagementPolicyListResult
type RoleManagementPolicyListResultIterator = original.RoleManagementPolicyListResultIterator
type RoleManagementPolicyListResultPage = original.RoleManagementPolicyListResultPage
type RoleManagementPolicyNotificationRule = original.RoleManagementPolicyNotificationRule
type RoleManagementPolicyProperties = original.RoleManagementPolicyProperties
type RoleManagementPolicyRule = original.RoleManagementPolicyRule
type RoleManagementPolicyRuleTarget = original.RoleManagementPolicyRuleTarget
type UserSet = original.UserSet

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewClassicAdministratorListResultIterator(page ClassicAdministratorListResultPage) ClassicAdministratorListResultIterator {
	return original.NewClassicAdministratorListResultIterator(page)
}
func NewClassicAdministratorListResultPage(cur ClassicAdministratorListResult, getNextPage func(context.Context, ClassicAdministratorListResult) (ClassicAdministratorListResult, error)) ClassicAdministratorListResultPage {
	return original.NewClassicAdministratorListResultPage(cur, getNextPage)
}
func NewClassicAdministratorsClient(subscriptionID string) ClassicAdministratorsClient {
	return original.NewClassicAdministratorsClient(subscriptionID)
}
func NewClassicAdministratorsClientWithBaseURI(baseURI string, subscriptionID string) ClassicAdministratorsClient {
	return original.NewClassicAdministratorsClientWithBaseURI(baseURI, subscriptionID)
}
func NewEligibleChildResourcesClient(subscriptionID string) EligibleChildResourcesClient {
	return original.NewEligibleChildResourcesClient(subscriptionID)
}
func NewEligibleChildResourcesClientWithBaseURI(baseURI string, subscriptionID string) EligibleChildResourcesClient {
	return original.NewEligibleChildResourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewEligibleChildResourcesListResultIterator(page EligibleChildResourcesListResultPage) EligibleChildResourcesListResultIterator {
	return original.NewEligibleChildResourcesListResultIterator(page)
}
func NewEligibleChildResourcesListResultPage(cur EligibleChildResourcesListResult, getNextPage func(context.Context, EligibleChildResourcesListResult) (EligibleChildResourcesListResult, error)) EligibleChildResourcesListResultPage {
	return original.NewEligibleChildResourcesListResultPage(cur, getNextPage)
}
func NewGlobalAdministratorClient(subscriptionID string) GlobalAdministratorClient {
	return original.NewGlobalAdministratorClient(subscriptionID)
}
func NewGlobalAdministratorClientWithBaseURI(baseURI string, subscriptionID string) GlobalAdministratorClient {
	return original.NewGlobalAdministratorClientWithBaseURI(baseURI, subscriptionID)
}
func NewPermissionGetResultIterator(page PermissionGetResultPage) PermissionGetResultIterator {
	return original.NewPermissionGetResultIterator(page)
}
func NewPermissionGetResultPage(cur PermissionGetResult, getNextPage func(context.Context, PermissionGetResult) (PermissionGetResult, error)) PermissionGetResultPage {
	return original.NewPermissionGetResultPage(cur, getNextPage)
}
func NewPermissionsClient(subscriptionID string) PermissionsClient {
	return original.NewPermissionsClient(subscriptionID)
}
func NewPermissionsClientWithBaseURI(baseURI string, subscriptionID string) PermissionsClient {
	return original.NewPermissionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewProviderOperationsMetadataClient(subscriptionID string) ProviderOperationsMetadataClient {
	return original.NewProviderOperationsMetadataClient(subscriptionID)
}
func NewProviderOperationsMetadataClientWithBaseURI(baseURI string, subscriptionID string) ProviderOperationsMetadataClient {
	return original.NewProviderOperationsMetadataClientWithBaseURI(baseURI, subscriptionID)
}
func NewProviderOperationsMetadataListResultIterator(page ProviderOperationsMetadataListResultPage) ProviderOperationsMetadataListResultIterator {
	return original.NewProviderOperationsMetadataListResultIterator(page)
}
func NewProviderOperationsMetadataListResultPage(cur ProviderOperationsMetadataListResult, getNextPage func(context.Context, ProviderOperationsMetadataListResult) (ProviderOperationsMetadataListResult, error)) ProviderOperationsMetadataListResultPage {
	return original.NewProviderOperationsMetadataListResultPage(cur, getNextPage)
}
func NewRoleAssignmentListResultIterator(page RoleAssignmentListResultPage) RoleAssignmentListResultIterator {
	return original.NewRoleAssignmentListResultIterator(page)
}
func NewRoleAssignmentListResultPage(cur RoleAssignmentListResult, getNextPage func(context.Context, RoleAssignmentListResult) (RoleAssignmentListResult, error)) RoleAssignmentListResultPage {
	return original.NewRoleAssignmentListResultPage(cur, getNextPage)
}
func NewRoleAssignmentScheduleInstanceListResultIterator(page RoleAssignmentScheduleInstanceListResultPage) RoleAssignmentScheduleInstanceListResultIterator {
	return original.NewRoleAssignmentScheduleInstanceListResultIterator(page)
}
func NewRoleAssignmentScheduleInstanceListResultPage(cur RoleAssignmentScheduleInstanceListResult, getNextPage func(context.Context, RoleAssignmentScheduleInstanceListResult) (RoleAssignmentScheduleInstanceListResult, error)) RoleAssignmentScheduleInstanceListResultPage {
	return original.NewRoleAssignmentScheduleInstanceListResultPage(cur, getNextPage)
}
func NewRoleAssignmentScheduleInstancesClient(subscriptionID string) RoleAssignmentScheduleInstancesClient {
	return original.NewRoleAssignmentScheduleInstancesClient(subscriptionID)
}
func NewRoleAssignmentScheduleInstancesClientWithBaseURI(baseURI string, subscriptionID string) RoleAssignmentScheduleInstancesClient {
	return original.NewRoleAssignmentScheduleInstancesClientWithBaseURI(baseURI, subscriptionID)
}
func NewRoleAssignmentScheduleListResultIterator(page RoleAssignmentScheduleListResultPage) RoleAssignmentScheduleListResultIterator {
	return original.NewRoleAssignmentScheduleListResultIterator(page)
}
func NewRoleAssignmentScheduleListResultPage(cur RoleAssignmentScheduleListResult, getNextPage func(context.Context, RoleAssignmentScheduleListResult) (RoleAssignmentScheduleListResult, error)) RoleAssignmentScheduleListResultPage {
	return original.NewRoleAssignmentScheduleListResultPage(cur, getNextPage)
}
func NewRoleAssignmentScheduleRequestListResultIterator(page RoleAssignmentScheduleRequestListResultPage) RoleAssignmentScheduleRequestListResultIterator {
	return original.NewRoleAssignmentScheduleRequestListResultIterator(page)
}
func NewRoleAssignmentScheduleRequestListResultPage(cur RoleAssignmentScheduleRequestListResult, getNextPage func(context.Context, RoleAssignmentScheduleRequestListResult) (RoleAssignmentScheduleRequestListResult, error)) RoleAssignmentScheduleRequestListResultPage {
	return original.NewRoleAssignmentScheduleRequestListResultPage(cur, getNextPage)
}
func NewRoleAssignmentScheduleRequestsClient(subscriptionID string) RoleAssignmentScheduleRequestsClient {
	return original.NewRoleAssignmentScheduleRequestsClient(subscriptionID)
}
func NewRoleAssignmentScheduleRequestsClientWithBaseURI(baseURI string, subscriptionID string) RoleAssignmentScheduleRequestsClient {
	return original.NewRoleAssignmentScheduleRequestsClientWithBaseURI(baseURI, subscriptionID)
}
func NewRoleAssignmentSchedulesClient(subscriptionID string) RoleAssignmentSchedulesClient {
	return original.NewRoleAssignmentSchedulesClient(subscriptionID)
}
func NewRoleAssignmentSchedulesClientWithBaseURI(baseURI string, subscriptionID string) RoleAssignmentSchedulesClient {
	return original.NewRoleAssignmentSchedulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewRoleAssignmentsClient(subscriptionID string) RoleAssignmentsClient {
	return original.NewRoleAssignmentsClient(subscriptionID)
}
func NewRoleAssignmentsClientWithBaseURI(baseURI string, subscriptionID string) RoleAssignmentsClient {
	return original.NewRoleAssignmentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewRoleDefinitionListResultIterator(page RoleDefinitionListResultPage) RoleDefinitionListResultIterator {
	return original.NewRoleDefinitionListResultIterator(page)
}
func NewRoleDefinitionListResultPage(cur RoleDefinitionListResult, getNextPage func(context.Context, RoleDefinitionListResult) (RoleDefinitionListResult, error)) RoleDefinitionListResultPage {
	return original.NewRoleDefinitionListResultPage(cur, getNextPage)
}
func NewRoleDefinitionsClient(subscriptionID string) RoleDefinitionsClient {
	return original.NewRoleDefinitionsClient(subscriptionID)
}
func NewRoleDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) RoleDefinitionsClient {
	return original.NewRoleDefinitionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewRoleEligibilityScheduleInstanceListResultIterator(page RoleEligibilityScheduleInstanceListResultPage) RoleEligibilityScheduleInstanceListResultIterator {
	return original.NewRoleEligibilityScheduleInstanceListResultIterator(page)
}
func NewRoleEligibilityScheduleInstanceListResultPage(cur RoleEligibilityScheduleInstanceListResult, getNextPage func(context.Context, RoleEligibilityScheduleInstanceListResult) (RoleEligibilityScheduleInstanceListResult, error)) RoleEligibilityScheduleInstanceListResultPage {
	return original.NewRoleEligibilityScheduleInstanceListResultPage(cur, getNextPage)
}
func NewRoleEligibilityScheduleInstancesClient(subscriptionID string) RoleEligibilityScheduleInstancesClient {
	return original.NewRoleEligibilityScheduleInstancesClient(subscriptionID)
}
func NewRoleEligibilityScheduleInstancesClientWithBaseURI(baseURI string, subscriptionID string) RoleEligibilityScheduleInstancesClient {
	return original.NewRoleEligibilityScheduleInstancesClientWithBaseURI(baseURI, subscriptionID)
}
func NewRoleEligibilityScheduleListResultIterator(page RoleEligibilityScheduleListResultPage) RoleEligibilityScheduleListResultIterator {
	return original.NewRoleEligibilityScheduleListResultIterator(page)
}
func NewRoleEligibilityScheduleListResultPage(cur RoleEligibilityScheduleListResult, getNextPage func(context.Context, RoleEligibilityScheduleListResult) (RoleEligibilityScheduleListResult, error)) RoleEligibilityScheduleListResultPage {
	return original.NewRoleEligibilityScheduleListResultPage(cur, getNextPage)
}
func NewRoleEligibilityScheduleRequestListResultIterator(page RoleEligibilityScheduleRequestListResultPage) RoleEligibilityScheduleRequestListResultIterator {
	return original.NewRoleEligibilityScheduleRequestListResultIterator(page)
}
func NewRoleEligibilityScheduleRequestListResultPage(cur RoleEligibilityScheduleRequestListResult, getNextPage func(context.Context, RoleEligibilityScheduleRequestListResult) (RoleEligibilityScheduleRequestListResult, error)) RoleEligibilityScheduleRequestListResultPage {
	return original.NewRoleEligibilityScheduleRequestListResultPage(cur, getNextPage)
}
func NewRoleEligibilityScheduleRequestsClient(subscriptionID string) RoleEligibilityScheduleRequestsClient {
	return original.NewRoleEligibilityScheduleRequestsClient(subscriptionID)
}
func NewRoleEligibilityScheduleRequestsClientWithBaseURI(baseURI string, subscriptionID string) RoleEligibilityScheduleRequestsClient {
	return original.NewRoleEligibilityScheduleRequestsClientWithBaseURI(baseURI, subscriptionID)
}
func NewRoleEligibilitySchedulesClient(subscriptionID string) RoleEligibilitySchedulesClient {
	return original.NewRoleEligibilitySchedulesClient(subscriptionID)
}
func NewRoleEligibilitySchedulesClientWithBaseURI(baseURI string, subscriptionID string) RoleEligibilitySchedulesClient {
	return original.NewRoleEligibilitySchedulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewRoleManagementPoliciesClient(subscriptionID string) RoleManagementPoliciesClient {
	return original.NewRoleManagementPoliciesClient(subscriptionID)
}
func NewRoleManagementPoliciesClientWithBaseURI(baseURI string, subscriptionID string) RoleManagementPoliciesClient {
	return original.NewRoleManagementPoliciesClientWithBaseURI(baseURI, subscriptionID)
}
func NewRoleManagementPolicyAssignmentListResultIterator(page RoleManagementPolicyAssignmentListResultPage) RoleManagementPolicyAssignmentListResultIterator {
	return original.NewRoleManagementPolicyAssignmentListResultIterator(page)
}
func NewRoleManagementPolicyAssignmentListResultPage(cur RoleManagementPolicyAssignmentListResult, getNextPage func(context.Context, RoleManagementPolicyAssignmentListResult) (RoleManagementPolicyAssignmentListResult, error)) RoleManagementPolicyAssignmentListResultPage {
	return original.NewRoleManagementPolicyAssignmentListResultPage(cur, getNextPage)
}
func NewRoleManagementPolicyAssignmentsClient(subscriptionID string) RoleManagementPolicyAssignmentsClient {
	return original.NewRoleManagementPolicyAssignmentsClient(subscriptionID)
}
func NewRoleManagementPolicyAssignmentsClientWithBaseURI(baseURI string, subscriptionID string) RoleManagementPolicyAssignmentsClient {
	return original.NewRoleManagementPolicyAssignmentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewRoleManagementPolicyListResultIterator(page RoleManagementPolicyListResultPage) RoleManagementPolicyListResultIterator {
	return original.NewRoleManagementPolicyListResultIterator(page)
}
func NewRoleManagementPolicyListResultPage(cur RoleManagementPolicyListResult, getNextPage func(context.Context, RoleManagementPolicyListResult) (RoleManagementPolicyListResult, error)) RoleManagementPolicyListResultPage {
	return original.NewRoleManagementPolicyListResultPage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleApprovalModeValues() []ApprovalMode {
	return original.PossibleApprovalModeValues()
}
func PossibleAssignmentTypeValues() []AssignmentType {
	return original.PossibleAssignmentTypeValues()
}
func PossibleEnablementRulesValues() []EnablementRules {
	return original.PossibleEnablementRulesValues()
}
func PossibleMemberTypeValues() []MemberType {
	return original.PossibleMemberTypeValues()
}
func PossibleNotificationDeliveryMechanismValues() []NotificationDeliveryMechanism {
	return original.PossibleNotificationDeliveryMechanismValues()
}
func PossibleNotificationLevelValues() []NotificationLevel {
	return original.PossibleNotificationLevelValues()
}
func PossiblePrincipalTypeValues() []PrincipalType {
	return original.PossiblePrincipalTypeValues()
}
func PossibleRecipientTypeValues() []RecipientType {
	return original.PossibleRecipientTypeValues()
}
func PossibleRequestTypeValues() []RequestType {
	return original.PossibleRequestTypeValues()
}
func PossibleRoleManagementPolicyRuleTypeValues() []RoleManagementPolicyRuleType {
	return original.PossibleRoleManagementPolicyRuleTypeValues()
}
func PossibleRuleTypeValues() []RuleType {
	return original.PossibleRuleTypeValues()
}
func PossibleStatusValues() []Status {
	return original.PossibleStatusValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func PossibleUserTypeValues() []UserType {
	return original.PossibleUserTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
