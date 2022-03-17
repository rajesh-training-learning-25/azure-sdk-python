//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package maps

import (
	"context"

	original "github.com/Azure/temp/github.com/Azure/azure-sdk-for-go/services/maps/mgmt/2021-02-01/maps"
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

type KeyType = original.KeyType

const (
	KeyTypePrimary   KeyType = original.KeyTypePrimary
	KeyTypeSecondary KeyType = original.KeyTypeSecondary
)

type Kind = original.Kind

const (
	KindGen1 Kind = original.KindGen1
	KindGen2 Kind = original.KindGen2
)

type Name = original.Name

const (
	NameG2 Name = original.NameG2
	NameS0 Name = original.NameS0
	NameS1 Name = original.NameS1
)

type Account = original.Account
type AccountKeys = original.AccountKeys
type AccountProperties = original.AccountProperties
type AccountUpdateParameters = original.AccountUpdateParameters
type Accounts = original.Accounts
type AccountsClient = original.AccountsClient
type AccountsIterator = original.AccountsIterator
type AccountsPage = original.AccountsPage
type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type Client = original.Client
type Creator = original.Creator
type CreatorList = original.CreatorList
type CreatorListIterator = original.CreatorListIterator
type CreatorListPage = original.CreatorListPage
type CreatorProperties = original.CreatorProperties
type CreatorUpdateParameters = original.CreatorUpdateParameters
type CreatorsClient = original.CreatorsClient
type Dimension = original.Dimension
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type KeySpecification = original.KeySpecification
type MetricSpecification = original.MetricSpecification
type OperationDetail = original.OperationDetail
type OperationDisplay = original.OperationDisplay
type OperationProperties = original.OperationProperties
type Operations = original.Operations
type OperationsIterator = original.OperationsIterator
type OperationsPage = original.OperationsPage
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type ServiceSpecification = original.ServiceSpecification
type Sku = original.Sku
type SystemData = original.SystemData
type TrackedResource = original.TrackedResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAccountsClient(subscriptionID string) AccountsClient {
	return original.NewAccountsClient(subscriptionID)
}
func NewAccountsClientWithBaseURI(baseURI string, subscriptionID string) AccountsClient {
	return original.NewAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewAccountsIterator(page AccountsPage) AccountsIterator {
	return original.NewAccountsIterator(page)
}
func NewAccountsPage(cur Accounts, getNextPage func(context.Context, Accounts) (Accounts, error)) AccountsPage {
	return original.NewAccountsPage(cur, getNextPage)
}
func NewClient(subscriptionID string) Client {
	return original.NewClient(subscriptionID)
}
func NewClientWithBaseURI(baseURI string, subscriptionID string) Client {
	return original.NewClientWithBaseURI(baseURI, subscriptionID)
}
func NewCreatorListIterator(page CreatorListPage) CreatorListIterator {
	return original.NewCreatorListIterator(page)
}
func NewCreatorListPage(cur CreatorList, getNextPage func(context.Context, CreatorList) (CreatorList, error)) CreatorListPage {
	return original.NewCreatorListPage(cur, getNextPage)
}
func NewCreatorsClient(subscriptionID string) CreatorsClient {
	return original.NewCreatorsClient(subscriptionID)
}
func NewCreatorsClientWithBaseURI(baseURI string, subscriptionID string) CreatorsClient {
	return original.NewCreatorsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsIterator(page OperationsPage) OperationsIterator {
	return original.NewOperationsIterator(page)
}
func NewOperationsPage(cur Operations, getNextPage func(context.Context, Operations) (Operations, error)) OperationsPage {
	return original.NewOperationsPage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleKeyTypeValues() []KeyType {
	return original.PossibleKeyTypeValues()
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleNameValues() []Name {
	return original.PossibleNameValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
