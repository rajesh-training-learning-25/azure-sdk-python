//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package features

import (
	"context"

	original "github.com/Azure/temp/github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2021-07-01/features"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type SubscriptionFeatureRegistrationApprovalType = original.SubscriptionFeatureRegistrationApprovalType

const (
	SubscriptionFeatureRegistrationApprovalTypeApprovalRequired SubscriptionFeatureRegistrationApprovalType = original.SubscriptionFeatureRegistrationApprovalTypeApprovalRequired
	SubscriptionFeatureRegistrationApprovalTypeAutoApproval     SubscriptionFeatureRegistrationApprovalType = original.SubscriptionFeatureRegistrationApprovalTypeAutoApproval
	SubscriptionFeatureRegistrationApprovalTypeNotSpecified     SubscriptionFeatureRegistrationApprovalType = original.SubscriptionFeatureRegistrationApprovalTypeNotSpecified
)

type SubscriptionFeatureRegistrationState = original.SubscriptionFeatureRegistrationState

const (
	SubscriptionFeatureRegistrationStateNotRegistered SubscriptionFeatureRegistrationState = original.SubscriptionFeatureRegistrationStateNotRegistered
	SubscriptionFeatureRegistrationStateNotSpecified  SubscriptionFeatureRegistrationState = original.SubscriptionFeatureRegistrationStateNotSpecified
	SubscriptionFeatureRegistrationStatePending       SubscriptionFeatureRegistrationState = original.SubscriptionFeatureRegistrationStatePending
	SubscriptionFeatureRegistrationStateRegistered    SubscriptionFeatureRegistrationState = original.SubscriptionFeatureRegistrationStateRegistered
	SubscriptionFeatureRegistrationStateRegistering   SubscriptionFeatureRegistrationState = original.SubscriptionFeatureRegistrationStateRegistering
	SubscriptionFeatureRegistrationStateUnregistered  SubscriptionFeatureRegistrationState = original.SubscriptionFeatureRegistrationStateUnregistered
	SubscriptionFeatureRegistrationStateUnregistering SubscriptionFeatureRegistrationState = original.SubscriptionFeatureRegistrationStateUnregistering
)

type AuthorizationProfile = original.AuthorizationProfile
type BaseClient = original.BaseClient
type Client = original.Client
type ErrorDefinition = original.ErrorDefinition
type ErrorResponse = original.ErrorResponse
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsListResult = original.OperationsListResult
type OperationsListResultIterator = original.OperationsListResultIterator
type OperationsListResultPage = original.OperationsListResultPage
type Properties = original.Properties
type ProxyResource = original.ProxyResource
type Result = original.Result
type SubscriptionFeatureRegistration = original.SubscriptionFeatureRegistration
type SubscriptionFeatureRegistrationList = original.SubscriptionFeatureRegistrationList
type SubscriptionFeatureRegistrationListIterator = original.SubscriptionFeatureRegistrationListIterator
type SubscriptionFeatureRegistrationListPage = original.SubscriptionFeatureRegistrationListPage
type SubscriptionFeatureRegistrationProperties = original.SubscriptionFeatureRegistrationProperties
type SubscriptionFeatureRegistrationsClient = original.SubscriptionFeatureRegistrationsClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewClient(subscriptionID string) Client {
	return original.NewClient(subscriptionID)
}
func NewClientWithBaseURI(baseURI string, subscriptionID string) Client {
	return original.NewClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsListResultIterator(page OperationsListResultPage) OperationsListResultIterator {
	return original.NewOperationsListResultIterator(page)
}
func NewOperationsListResultPage(cur OperationsListResult, getNextPage func(context.Context, OperationsListResult) (OperationsListResult, error)) OperationsListResultPage {
	return original.NewOperationsListResultPage(cur, getNextPage)
}
func NewSubscriptionFeatureRegistrationListIterator(page SubscriptionFeatureRegistrationListPage) SubscriptionFeatureRegistrationListIterator {
	return original.NewSubscriptionFeatureRegistrationListIterator(page)
}
func NewSubscriptionFeatureRegistrationListPage(cur SubscriptionFeatureRegistrationList, getNextPage func(context.Context, SubscriptionFeatureRegistrationList) (SubscriptionFeatureRegistrationList, error)) SubscriptionFeatureRegistrationListPage {
	return original.NewSubscriptionFeatureRegistrationListPage(cur, getNextPage)
}
func NewSubscriptionFeatureRegistrationsClient(subscriptionID string) SubscriptionFeatureRegistrationsClient {
	return original.NewSubscriptionFeatureRegistrationsClient(subscriptionID)
}
func NewSubscriptionFeatureRegistrationsClientWithBaseURI(baseURI string, subscriptionID string) SubscriptionFeatureRegistrationsClient {
	return original.NewSubscriptionFeatureRegistrationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleSubscriptionFeatureRegistrationApprovalTypeValues() []SubscriptionFeatureRegistrationApprovalType {
	return original.PossibleSubscriptionFeatureRegistrationApprovalTypeValues()
}
func PossibleSubscriptionFeatureRegistrationStateValues() []SubscriptionFeatureRegistrationState {
	return original.PossibleSubscriptionFeatureRegistrationStateValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
