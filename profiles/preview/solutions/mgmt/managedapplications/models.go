//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package managedapplications

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/solutions/mgmt/2018-02-01/managedapplications"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ApplicationArtifactName = original.ApplicationArtifactName

const (
	Authorizations       ApplicationArtifactName = original.Authorizations
	CustomRoleDefinition ApplicationArtifactName = original.CustomRoleDefinition
	NotSpecified         ApplicationArtifactName = original.NotSpecified
	ViewDefinition       ApplicationArtifactName = original.ViewDefinition
)

type ApplicationArtifactType = original.ApplicationArtifactType

const (
	ApplicationArtifactTypeCustom       ApplicationArtifactType = original.ApplicationArtifactTypeCustom
	ApplicationArtifactTypeNotSpecified ApplicationArtifactType = original.ApplicationArtifactTypeNotSpecified
	ApplicationArtifactTypeTemplate     ApplicationArtifactType = original.ApplicationArtifactTypeTemplate
)

type ApplicationDefinitionArtifactName = original.ApplicationDefinitionArtifactName

const (
	ApplicationDefinitionArtifactNameApplicationResourceTemplate ApplicationDefinitionArtifactName = original.ApplicationDefinitionArtifactNameApplicationResourceTemplate
	ApplicationDefinitionArtifactNameCreateUIDefinition          ApplicationDefinitionArtifactName = original.ApplicationDefinitionArtifactNameCreateUIDefinition
	ApplicationDefinitionArtifactNameMainTemplateParameters      ApplicationDefinitionArtifactName = original.ApplicationDefinitionArtifactNameMainTemplateParameters
	ApplicationDefinitionArtifactNameNotSpecified                ApplicationDefinitionArtifactName = original.ApplicationDefinitionArtifactNameNotSpecified
)

type ApplicationLockLevel = original.ApplicationLockLevel

const (
	CanNotDelete ApplicationLockLevel = original.CanNotDelete
	None         ApplicationLockLevel = original.None
	ReadOnly     ApplicationLockLevel = original.ReadOnly
)

type ApplicationManagementMode = original.ApplicationManagementMode

const (
	ApplicationManagementModeManaged      ApplicationManagementMode = original.ApplicationManagementModeManaged
	ApplicationManagementModeNotSpecified ApplicationManagementMode = original.ApplicationManagementModeNotSpecified
	ApplicationManagementModeUnmanaged    ApplicationManagementMode = original.ApplicationManagementModeUnmanaged
)

type DeploymentMode = original.DeploymentMode

const (
	DeploymentModeComplete     DeploymentMode = original.DeploymentModeComplete
	DeploymentModeIncremental  DeploymentMode = original.DeploymentModeIncremental
	DeploymentModeNotSpecified DeploymentMode = original.DeploymentModeNotSpecified
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateAccepted     ProvisioningState = original.ProvisioningStateAccepted
	ProvisioningStateCanceled     ProvisioningState = original.ProvisioningStateCanceled
	ProvisioningStateCreated      ProvisioningState = original.ProvisioningStateCreated
	ProvisioningStateCreating     ProvisioningState = original.ProvisioningStateCreating
	ProvisioningStateDeleted      ProvisioningState = original.ProvisioningStateDeleted
	ProvisioningStateDeleting     ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateFailed       ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateNotSpecified ProvisioningState = original.ProvisioningStateNotSpecified
	ProvisioningStateReady        ProvisioningState = original.ProvisioningStateReady
	ProvisioningStateRunning      ProvisioningState = original.ProvisioningStateRunning
	ProvisioningStateSucceeded    ProvisioningState = original.ProvisioningStateSucceeded
	ProvisioningStateUpdating     ProvisioningState = original.ProvisioningStateUpdating
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	ResourceIdentityTypeNone                       ResourceIdentityType = original.ResourceIdentityTypeNone
	ResourceIdentityTypeSystemAssigned             ResourceIdentityType = original.ResourceIdentityTypeSystemAssigned
	ResourceIdentityTypeSystemAssignedUserAssigned ResourceIdentityType = original.ResourceIdentityTypeSystemAssignedUserAssigned
	ResourceIdentityTypeUserAssigned               ResourceIdentityType = original.ResourceIdentityTypeUserAssigned
)

type Application = original.Application
type ApplicationArtifact = original.ApplicationArtifact
type ApplicationAuthorization = original.ApplicationAuthorization
type ApplicationBillingDetailsDefinition = original.ApplicationBillingDetailsDefinition
type ApplicationClientDetails = original.ApplicationClientDetails
type ApplicationDefinition = original.ApplicationDefinition
type ApplicationDefinitionArtifact = original.ApplicationDefinitionArtifact
type ApplicationDefinitionListResult = original.ApplicationDefinitionListResult
type ApplicationDefinitionListResultIterator = original.ApplicationDefinitionListResultIterator
type ApplicationDefinitionListResultPage = original.ApplicationDefinitionListResultPage
type ApplicationDefinitionPatchable = original.ApplicationDefinitionPatchable
type ApplicationDefinitionProperties = original.ApplicationDefinitionProperties
type ApplicationDefinitionsClient = original.ApplicationDefinitionsClient
type ApplicationListResult = original.ApplicationListResult
type ApplicationListResultIterator = original.ApplicationListResultIterator
type ApplicationListResultPage = original.ApplicationListResultPage
type ApplicationPackageContact = original.ApplicationPackageContact
type ApplicationPackageSupportUrls = original.ApplicationPackageSupportUrls
type ApplicationPatchable = original.ApplicationPatchable
type ApplicationPolicy = original.ApplicationPolicy
type ApplicationProperties = original.ApplicationProperties
type ApplicationPropertiesPatchable = original.ApplicationPropertiesPatchable
type ApplicationsClient = original.ApplicationsClient
type ApplicationsCreateOrUpdateByIDFuture = original.ApplicationsCreateOrUpdateByIDFuture
type ApplicationsCreateOrUpdateFuture = original.ApplicationsCreateOrUpdateFuture
type ApplicationsDeleteByIDFuture = original.ApplicationsDeleteByIDFuture
type ApplicationsDeleteFuture = original.ApplicationsDeleteFuture
type ApplicationsRefreshPermissionsFuture = original.ApplicationsRefreshPermissionsFuture
type ApplicationsUpdateByIDFuture = original.ApplicationsUpdateByIDFuture
type ApplicationsUpdateFuture = original.ApplicationsUpdateFuture
type BaseClient = original.BaseClient
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type GenericResource = original.GenericResource
type Identity = original.Identity
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type Plan = original.Plan
type PlanPatchable = original.PlanPatchable
type Resource = original.Resource
type Sku = original.Sku
type UserAssignedResourceIdentity = original.UserAssignedResourceIdentity

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewApplicationDefinitionListResultIterator(page ApplicationDefinitionListResultPage) ApplicationDefinitionListResultIterator {
	return original.NewApplicationDefinitionListResultIterator(page)
}
func NewApplicationDefinitionListResultPage(cur ApplicationDefinitionListResult, getNextPage func(context.Context, ApplicationDefinitionListResult) (ApplicationDefinitionListResult, error)) ApplicationDefinitionListResultPage {
	return original.NewApplicationDefinitionListResultPage(cur, getNextPage)
}
func NewApplicationDefinitionsClient(subscriptionID string) ApplicationDefinitionsClient {
	return original.NewApplicationDefinitionsClient(subscriptionID)
}
func NewApplicationDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) ApplicationDefinitionsClient {
	return original.NewApplicationDefinitionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewApplicationListResultIterator(page ApplicationListResultPage) ApplicationListResultIterator {
	return original.NewApplicationListResultIterator(page)
}
func NewApplicationListResultPage(cur ApplicationListResult, getNextPage func(context.Context, ApplicationListResult) (ApplicationListResult, error)) ApplicationListResultPage {
	return original.NewApplicationListResultPage(cur, getNextPage)
}
func NewApplicationsClient(subscriptionID string) ApplicationsClient {
	return original.NewApplicationsClient(subscriptionID)
}
func NewApplicationsClientWithBaseURI(baseURI string, subscriptionID string) ApplicationsClient {
	return original.NewApplicationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleApplicationArtifactNameValues() []ApplicationArtifactName {
	return original.PossibleApplicationArtifactNameValues()
}
func PossibleApplicationArtifactTypeValues() []ApplicationArtifactType {
	return original.PossibleApplicationArtifactTypeValues()
}
func PossibleApplicationDefinitionArtifactNameValues() []ApplicationDefinitionArtifactName {
	return original.PossibleApplicationDefinitionArtifactNameValues()
}
func PossibleApplicationLockLevelValues() []ApplicationLockLevel {
	return original.PossibleApplicationLockLevelValues()
}
func PossibleApplicationManagementModeValues() []ApplicationManagementMode {
	return original.PossibleApplicationManagementModeValues()
}
func PossibleDeploymentModeValues() []DeploymentMode {
	return original.PossibleDeploymentModeValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
