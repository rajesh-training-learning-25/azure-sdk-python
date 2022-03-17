//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package confidentialledger

import (
	"context"

	original "github.com/Azure/temp/github.com/Azure/azure-sdk-for-go/services/preview/confidentialledger/mgmt/2020-12-01-preview/confidentialledger"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type CreatedByType = original.CreatedByType

const (
	CreatedByTypeApplication     CreatedByType = original.CreatedByTypeApplication
	CreatedByTypeKey             CreatedByType = original.CreatedByTypeKey
	CreatedByTypeManagedIdentity CreatedByType = original.CreatedByTypeManagedIdentity
	CreatedByTypeUser            CreatedByType = original.CreatedByTypeUser
)

type LedgerRoleName = original.LedgerRoleName

const (
	LedgerRoleNameAdministrator LedgerRoleName = original.LedgerRoleNameAdministrator
	LedgerRoleNameContributor   LedgerRoleName = original.LedgerRoleNameContributor
	LedgerRoleNameReader        LedgerRoleName = original.LedgerRoleNameReader
)

type LedgerType = original.LedgerType

const (
	LedgerTypePrivate LedgerType = original.LedgerTypePrivate
	LedgerTypePublic  LedgerType = original.LedgerTypePublic
	LedgerTypeUnknown LedgerType = original.LedgerTypeUnknown
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

type AADBasedSecurityPrincipal = original.AADBasedSecurityPrincipal
type BaseClient = original.BaseClient
type CertBasedSecurityPrincipal = original.CertBasedSecurityPrincipal
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type LedgerClient = original.LedgerClient
type LedgerCreateFuture = original.LedgerCreateFuture
type LedgerDeleteFuture = original.LedgerDeleteFuture
type LedgerProperties = original.LedgerProperties
type LedgerUpdateFuture = original.LedgerUpdateFuture
type List = original.List
type ListIterator = original.ListIterator
type ListPage = original.ListPage
type Location = original.Location
type Model = original.Model
type OperationsClient = original.OperationsClient
type Resource = original.Resource
type ResourceProviderOperationDefinition = original.ResourceProviderOperationDefinition
type ResourceProviderOperationDisplay = original.ResourceProviderOperationDisplay
type ResourceProviderOperationList = original.ResourceProviderOperationList
type ResourceProviderOperationListIterator = original.ResourceProviderOperationListIterator
type ResourceProviderOperationListPage = original.ResourceProviderOperationListPage
type SystemData = original.SystemData
type Tags = original.Tags

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewLedgerClient(subscriptionID string) LedgerClient {
	return original.NewLedgerClient(subscriptionID)
}
func NewLedgerClientWithBaseURI(baseURI string, subscriptionID string) LedgerClient {
	return original.NewLedgerClientWithBaseURI(baseURI, subscriptionID)
}
func NewListIterator(page ListPage) ListIterator {
	return original.NewListIterator(page)
}
func NewListPage(cur List, getNextPage func(context.Context, List) (List, error)) ListPage {
	return original.NewListPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceProviderOperationListIterator(page ResourceProviderOperationListPage) ResourceProviderOperationListIterator {
	return original.NewResourceProviderOperationListIterator(page)
}
func NewResourceProviderOperationListPage(cur ResourceProviderOperationList, getNextPage func(context.Context, ResourceProviderOperationList) (ResourceProviderOperationList, error)) ResourceProviderOperationListPage {
	return original.NewResourceProviderOperationListPage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleLedgerRoleNameValues() []LedgerRoleName {
	return original.PossibleLedgerRoleNameValues()
}
func PossibleLedgerTypeValues() []LedgerType {
	return original.PossibleLedgerTypeValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
