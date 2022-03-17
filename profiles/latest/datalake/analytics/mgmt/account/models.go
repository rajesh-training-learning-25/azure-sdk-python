//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package account

import (
	"context"

	original "github.com/Azure/temp/github.com/Azure/azure-sdk-for-go/services/datalake/analytics/mgmt/2016-11-01/account"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AADObjectType = original.AADObjectType

const (
	Group            AADObjectType = original.Group
	ServicePrincipal AADObjectType = original.ServicePrincipal
	User             AADObjectType = original.User
)

type DataLakeAnalyticsAccountState = original.DataLakeAnalyticsAccountState

const (
	Active    DataLakeAnalyticsAccountState = original.Active
	Suspended DataLakeAnalyticsAccountState = original.Suspended
)

type DataLakeAnalyticsAccountStatus = original.DataLakeAnalyticsAccountStatus

const (
	Canceled   DataLakeAnalyticsAccountStatus = original.Canceled
	Creating   DataLakeAnalyticsAccountStatus = original.Creating
	Deleted    DataLakeAnalyticsAccountStatus = original.Deleted
	Deleting   DataLakeAnalyticsAccountStatus = original.Deleting
	Failed     DataLakeAnalyticsAccountStatus = original.Failed
	Patching   DataLakeAnalyticsAccountStatus = original.Patching
	Resuming   DataLakeAnalyticsAccountStatus = original.Resuming
	Running    DataLakeAnalyticsAccountStatus = original.Running
	Succeeded  DataLakeAnalyticsAccountStatus = original.Succeeded
	Suspending DataLakeAnalyticsAccountStatus = original.Suspending
	Undeleting DataLakeAnalyticsAccountStatus = original.Undeleting
)

type FirewallAllowAzureIpsState = original.FirewallAllowAzureIpsState

const (
	Disabled FirewallAllowAzureIpsState = original.Disabled
	Enabled  FirewallAllowAzureIpsState = original.Enabled
)

type FirewallState = original.FirewallState

const (
	FirewallStateDisabled FirewallState = original.FirewallStateDisabled
	FirewallStateEnabled  FirewallState = original.FirewallStateEnabled
)

type OperationOrigin = original.OperationOrigin

const (
	OperationOriginSystem     OperationOrigin = original.OperationOriginSystem
	OperationOriginUser       OperationOrigin = original.OperationOriginUser
	OperationOriginUsersystem OperationOrigin = original.OperationOriginUsersystem
)

type SubscriptionState = original.SubscriptionState

const (
	SubscriptionStateDeleted      SubscriptionState = original.SubscriptionStateDeleted
	SubscriptionStateRegistered   SubscriptionState = original.SubscriptionStateRegistered
	SubscriptionStateSuspended    SubscriptionState = original.SubscriptionStateSuspended
	SubscriptionStateUnregistered SubscriptionState = original.SubscriptionStateUnregistered
	SubscriptionStateWarned       SubscriptionState = original.SubscriptionStateWarned
)

type TierType = original.TierType

const (
	Commitment100000AUHours TierType = original.Commitment100000AUHours
	Commitment10000AUHours  TierType = original.Commitment10000AUHours
	Commitment1000AUHours   TierType = original.Commitment1000AUHours
	Commitment100AUHours    TierType = original.Commitment100AUHours
	Commitment500000AUHours TierType = original.Commitment500000AUHours
	Commitment50000AUHours  TierType = original.Commitment50000AUHours
	Commitment5000AUHours   TierType = original.Commitment5000AUHours
	Commitment500AUHours    TierType = original.Commitment500AUHours
	Consumption             TierType = original.Consumption
)

type AccountsClient = original.AccountsClient
type AccountsCreateFutureType = original.AccountsCreateFutureType
type AccountsDeleteFutureType = original.AccountsDeleteFutureType
type AccountsUpdateFutureType = original.AccountsUpdateFutureType
type AddDataLakeStoreParameters = original.AddDataLakeStoreParameters
type AddDataLakeStoreProperties = original.AddDataLakeStoreProperties
type AddDataLakeStoreWithAccountParameters = original.AddDataLakeStoreWithAccountParameters
type AddStorageAccountParameters = original.AddStorageAccountParameters
type AddStorageAccountProperties = original.AddStorageAccountProperties
type AddStorageAccountWithAccountParameters = original.AddStorageAccountWithAccountParameters
type BaseClient = original.BaseClient
type CapabilityInformation = original.CapabilityInformation
type CheckNameAvailabilityParameters = original.CheckNameAvailabilityParameters
type ComputePoliciesClient = original.ComputePoliciesClient
type ComputePolicy = original.ComputePolicy
type ComputePolicyListResult = original.ComputePolicyListResult
type ComputePolicyListResultIterator = original.ComputePolicyListResultIterator
type ComputePolicyListResultPage = original.ComputePolicyListResultPage
type ComputePolicyProperties = original.ComputePolicyProperties
type CreateComputePolicyWithAccountParameters = original.CreateComputePolicyWithAccountParameters
type CreateDataLakeAnalyticsAccountParameters = original.CreateDataLakeAnalyticsAccountParameters
type CreateDataLakeAnalyticsAccountProperties = original.CreateDataLakeAnalyticsAccountProperties
type CreateFirewallRuleWithAccountParameters = original.CreateFirewallRuleWithAccountParameters
type CreateOrUpdateComputePolicyParameters = original.CreateOrUpdateComputePolicyParameters
type CreateOrUpdateComputePolicyProperties = original.CreateOrUpdateComputePolicyProperties
type CreateOrUpdateFirewallRuleParameters = original.CreateOrUpdateFirewallRuleParameters
type CreateOrUpdateFirewallRuleProperties = original.CreateOrUpdateFirewallRuleProperties
type DataLakeAnalyticsAccount = original.DataLakeAnalyticsAccount
type DataLakeAnalyticsAccountBasic = original.DataLakeAnalyticsAccountBasic
type DataLakeAnalyticsAccountListResult = original.DataLakeAnalyticsAccountListResult
type DataLakeAnalyticsAccountListResultIterator = original.DataLakeAnalyticsAccountListResultIterator
type DataLakeAnalyticsAccountListResultPage = original.DataLakeAnalyticsAccountListResultPage
type DataLakeAnalyticsAccountProperties = original.DataLakeAnalyticsAccountProperties
type DataLakeAnalyticsAccountPropertiesBasic = original.DataLakeAnalyticsAccountPropertiesBasic
type DataLakeStoreAccountInformation = original.DataLakeStoreAccountInformation
type DataLakeStoreAccountInformationListResult = original.DataLakeStoreAccountInformationListResult
type DataLakeStoreAccountInformationListResultIterator = original.DataLakeStoreAccountInformationListResultIterator
type DataLakeStoreAccountInformationListResultPage = original.DataLakeStoreAccountInformationListResultPage
type DataLakeStoreAccountInformationProperties = original.DataLakeStoreAccountInformationProperties
type DataLakeStoreAccountsClient = original.DataLakeStoreAccountsClient
type FirewallRule = original.FirewallRule
type FirewallRuleListResult = original.FirewallRuleListResult
type FirewallRuleListResultIterator = original.FirewallRuleListResultIterator
type FirewallRuleListResultPage = original.FirewallRuleListResultPage
type FirewallRuleProperties = original.FirewallRuleProperties
type FirewallRulesClient = original.FirewallRulesClient
type LocationsClient = original.LocationsClient
type NameAvailabilityInformation = original.NameAvailabilityInformation
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationsClient = original.OperationsClient
type Resource = original.Resource
type SasTokenInformation = original.SasTokenInformation
type SasTokenInformationListResult = original.SasTokenInformationListResult
type SasTokenInformationListResultIterator = original.SasTokenInformationListResultIterator
type SasTokenInformationListResultPage = original.SasTokenInformationListResultPage
type StorageAccountInformation = original.StorageAccountInformation
type StorageAccountInformationListResult = original.StorageAccountInformationListResult
type StorageAccountInformationListResultIterator = original.StorageAccountInformationListResultIterator
type StorageAccountInformationListResultPage = original.StorageAccountInformationListResultPage
type StorageAccountInformationProperties = original.StorageAccountInformationProperties
type StorageAccountsClient = original.StorageAccountsClient
type StorageContainer = original.StorageContainer
type StorageContainerListResult = original.StorageContainerListResult
type StorageContainerListResultIterator = original.StorageContainerListResultIterator
type StorageContainerListResultPage = original.StorageContainerListResultPage
type StorageContainerProperties = original.StorageContainerProperties
type SubResource = original.SubResource
type UpdateComputePolicyParameters = original.UpdateComputePolicyParameters
type UpdateComputePolicyProperties = original.UpdateComputePolicyProperties
type UpdateComputePolicyWithAccountParameters = original.UpdateComputePolicyWithAccountParameters
type UpdateDataLakeAnalyticsAccountParameters = original.UpdateDataLakeAnalyticsAccountParameters
type UpdateDataLakeAnalyticsAccountProperties = original.UpdateDataLakeAnalyticsAccountProperties
type UpdateDataLakeStoreProperties = original.UpdateDataLakeStoreProperties
type UpdateDataLakeStoreWithAccountParameters = original.UpdateDataLakeStoreWithAccountParameters
type UpdateFirewallRuleParameters = original.UpdateFirewallRuleParameters
type UpdateFirewallRuleProperties = original.UpdateFirewallRuleProperties
type UpdateFirewallRuleWithAccountParameters = original.UpdateFirewallRuleWithAccountParameters
type UpdateStorageAccountParameters = original.UpdateStorageAccountParameters
type UpdateStorageAccountProperties = original.UpdateStorageAccountProperties
type UpdateStorageAccountWithAccountParameters = original.UpdateStorageAccountWithAccountParameters

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAccountsClient(subscriptionID string) AccountsClient {
	return original.NewAccountsClient(subscriptionID)
}
func NewAccountsClientWithBaseURI(baseURI string, subscriptionID string) AccountsClient {
	return original.NewAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewComputePoliciesClient(subscriptionID string) ComputePoliciesClient {
	return original.NewComputePoliciesClient(subscriptionID)
}
func NewComputePoliciesClientWithBaseURI(baseURI string, subscriptionID string) ComputePoliciesClient {
	return original.NewComputePoliciesClientWithBaseURI(baseURI, subscriptionID)
}
func NewComputePolicyListResultIterator(page ComputePolicyListResultPage) ComputePolicyListResultIterator {
	return original.NewComputePolicyListResultIterator(page)
}
func NewComputePolicyListResultPage(cur ComputePolicyListResult, getNextPage func(context.Context, ComputePolicyListResult) (ComputePolicyListResult, error)) ComputePolicyListResultPage {
	return original.NewComputePolicyListResultPage(cur, getNextPage)
}
func NewDataLakeAnalyticsAccountListResultIterator(page DataLakeAnalyticsAccountListResultPage) DataLakeAnalyticsAccountListResultIterator {
	return original.NewDataLakeAnalyticsAccountListResultIterator(page)
}
func NewDataLakeAnalyticsAccountListResultPage(cur DataLakeAnalyticsAccountListResult, getNextPage func(context.Context, DataLakeAnalyticsAccountListResult) (DataLakeAnalyticsAccountListResult, error)) DataLakeAnalyticsAccountListResultPage {
	return original.NewDataLakeAnalyticsAccountListResultPage(cur, getNextPage)
}
func NewDataLakeStoreAccountInformationListResultIterator(page DataLakeStoreAccountInformationListResultPage) DataLakeStoreAccountInformationListResultIterator {
	return original.NewDataLakeStoreAccountInformationListResultIterator(page)
}
func NewDataLakeStoreAccountInformationListResultPage(cur DataLakeStoreAccountInformationListResult, getNextPage func(context.Context, DataLakeStoreAccountInformationListResult) (DataLakeStoreAccountInformationListResult, error)) DataLakeStoreAccountInformationListResultPage {
	return original.NewDataLakeStoreAccountInformationListResultPage(cur, getNextPage)
}
func NewDataLakeStoreAccountsClient(subscriptionID string) DataLakeStoreAccountsClient {
	return original.NewDataLakeStoreAccountsClient(subscriptionID)
}
func NewDataLakeStoreAccountsClientWithBaseURI(baseURI string, subscriptionID string) DataLakeStoreAccountsClient {
	return original.NewDataLakeStoreAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewFirewallRuleListResultIterator(page FirewallRuleListResultPage) FirewallRuleListResultIterator {
	return original.NewFirewallRuleListResultIterator(page)
}
func NewFirewallRuleListResultPage(cur FirewallRuleListResult, getNextPage func(context.Context, FirewallRuleListResult) (FirewallRuleListResult, error)) FirewallRuleListResultPage {
	return original.NewFirewallRuleListResultPage(cur, getNextPage)
}
func NewFirewallRulesClient(subscriptionID string) FirewallRulesClient {
	return original.NewFirewallRulesClient(subscriptionID)
}
func NewFirewallRulesClientWithBaseURI(baseURI string, subscriptionID string) FirewallRulesClient {
	return original.NewFirewallRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewLocationsClient(subscriptionID string) LocationsClient {
	return original.NewLocationsClient(subscriptionID)
}
func NewLocationsClientWithBaseURI(baseURI string, subscriptionID string) LocationsClient {
	return original.NewLocationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSasTokenInformationListResultIterator(page SasTokenInformationListResultPage) SasTokenInformationListResultIterator {
	return original.NewSasTokenInformationListResultIterator(page)
}
func NewSasTokenInformationListResultPage(cur SasTokenInformationListResult, getNextPage func(context.Context, SasTokenInformationListResult) (SasTokenInformationListResult, error)) SasTokenInformationListResultPage {
	return original.NewSasTokenInformationListResultPage(cur, getNextPage)
}
func NewStorageAccountInformationListResultIterator(page StorageAccountInformationListResultPage) StorageAccountInformationListResultIterator {
	return original.NewStorageAccountInformationListResultIterator(page)
}
func NewStorageAccountInformationListResultPage(cur StorageAccountInformationListResult, getNextPage func(context.Context, StorageAccountInformationListResult) (StorageAccountInformationListResult, error)) StorageAccountInformationListResultPage {
	return original.NewStorageAccountInformationListResultPage(cur, getNextPage)
}
func NewStorageAccountsClient(subscriptionID string) StorageAccountsClient {
	return original.NewStorageAccountsClient(subscriptionID)
}
func NewStorageAccountsClientWithBaseURI(baseURI string, subscriptionID string) StorageAccountsClient {
	return original.NewStorageAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewStorageContainerListResultIterator(page StorageContainerListResultPage) StorageContainerListResultIterator {
	return original.NewStorageContainerListResultIterator(page)
}
func NewStorageContainerListResultPage(cur StorageContainerListResult, getNextPage func(context.Context, StorageContainerListResult) (StorageContainerListResult, error)) StorageContainerListResultPage {
	return original.NewStorageContainerListResultPage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAADObjectTypeValues() []AADObjectType {
	return original.PossibleAADObjectTypeValues()
}
func PossibleDataLakeAnalyticsAccountStateValues() []DataLakeAnalyticsAccountState {
	return original.PossibleDataLakeAnalyticsAccountStateValues()
}
func PossibleDataLakeAnalyticsAccountStatusValues() []DataLakeAnalyticsAccountStatus {
	return original.PossibleDataLakeAnalyticsAccountStatusValues()
}
func PossibleFirewallAllowAzureIpsStateValues() []FirewallAllowAzureIpsState {
	return original.PossibleFirewallAllowAzureIpsStateValues()
}
func PossibleFirewallStateValues() []FirewallState {
	return original.PossibleFirewallStateValues()
}
func PossibleOperationOriginValues() []OperationOrigin {
	return original.PossibleOperationOriginValues()
}
func PossibleSubscriptionStateValues() []SubscriptionState {
	return original.PossibleSubscriptionStateValues()
}
func PossibleTierTypeValues() []TierType {
	return original.PossibleTierTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
