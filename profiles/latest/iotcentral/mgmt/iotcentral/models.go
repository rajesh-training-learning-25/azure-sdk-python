//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package iotcentral

import (
	"context"

	original "github.com/Azure/temp/github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2021-06-01/iotcentral"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AppSku = original.AppSku

const (
	AppSkuST0 AppSku = original.AppSkuST0
	AppSkuST1 AppSku = original.AppSkuST1
	AppSkuST2 AppSku = original.AppSkuST2
)

type AppState = original.AppState

const (
	AppStateCreated   AppState = original.AppStateCreated
	AppStateSuspended AppState = original.AppStateSuspended
)

type SystemAssignedServiceIdentityType = original.SystemAssignedServiceIdentityType

const (
	SystemAssignedServiceIdentityTypeNone           SystemAssignedServiceIdentityType = original.SystemAssignedServiceIdentityTypeNone
	SystemAssignedServiceIdentityTypeSystemAssigned SystemAssignedServiceIdentityType = original.SystemAssignedServiceIdentityTypeSystemAssigned
)

type App = original.App
type AppAvailabilityInfo = original.AppAvailabilityInfo
type AppListResult = original.AppListResult
type AppListResultIterator = original.AppListResultIterator
type AppListResultPage = original.AppListResultPage
type AppPatch = original.AppPatch
type AppProperties = original.AppProperties
type AppSkuInfo = original.AppSkuInfo
type AppTemplate = original.AppTemplate
type AppTemplateLocations = original.AppTemplateLocations
type AppTemplatesResult = original.AppTemplatesResult
type AppTemplatesResultIterator = original.AppTemplatesResultIterator
type AppTemplatesResultPage = original.AppTemplatesResultPage
type AppsClient = original.AppsClient
type AppsCreateOrUpdateFuture = original.AppsCreateOrUpdateFuture
type AppsDeleteFuture = original.AppsDeleteFuture
type AppsUpdateFuture = original.AppsUpdateFuture
type BaseClient = original.BaseClient
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationInputs = original.OperationInputs
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type Resource = original.Resource
type SystemAssignedServiceIdentity = original.SystemAssignedServiceIdentity

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAppListResultIterator(page AppListResultPage) AppListResultIterator {
	return original.NewAppListResultIterator(page)
}
func NewAppListResultPage(cur AppListResult, getNextPage func(context.Context, AppListResult) (AppListResult, error)) AppListResultPage {
	return original.NewAppListResultPage(cur, getNextPage)
}
func NewAppTemplatesResultIterator(page AppTemplatesResultPage) AppTemplatesResultIterator {
	return original.NewAppTemplatesResultIterator(page)
}
func NewAppTemplatesResultPage(cur AppTemplatesResult, getNextPage func(context.Context, AppTemplatesResult) (AppTemplatesResult, error)) AppTemplatesResultPage {
	return original.NewAppTemplatesResultPage(cur, getNextPage)
}
func NewAppsClient(subscriptionID string) AppsClient {
	return original.NewAppsClient(subscriptionID)
}
func NewAppsClientWithBaseURI(baseURI string, subscriptionID string) AppsClient {
	return original.NewAppsClientWithBaseURI(baseURI, subscriptionID)
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
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAppSkuValues() []AppSku {
	return original.PossibleAppSkuValues()
}
func PossibleAppStateValues() []AppState {
	return original.PossibleAppStateValues()
}
func PossibleSystemAssignedServiceIdentityTypeValues() []SystemAssignedServiceIdentityType {
	return original.PossibleSystemAssignedServiceIdentityTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
