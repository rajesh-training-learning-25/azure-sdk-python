package cognitiveservicesapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/mgmt/2021-04-30/cognitiveservices"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	CheckDomainAvailability(ctx context.Context, parameters cognitiveservices.CheckDomainAvailabilityParameter) (result cognitiveservices.DomainAvailability, err error)
	CheckSkuAvailability(ctx context.Context, location string, parameters cognitiveservices.CheckSkuAvailabilityParameter) (result cognitiveservices.SkuAvailabilityListResult, err error)
}

var _ BaseClientAPI = (*cognitiveservices.BaseClient)(nil)

// AccountsClientAPI contains the set of methods on the AccountsClient type.
type AccountsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, accountName string, account cognitiveservices.Account) (result cognitiveservices.AccountsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string) (result cognitiveservices.AccountsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string) (result cognitiveservices.Account, err error)
	List(ctx context.Context) (result cognitiveservices.AccountListResultPage, err error)
	ListComplete(ctx context.Context) (result cognitiveservices.AccountListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result cognitiveservices.AccountListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result cognitiveservices.AccountListResultIterator, err error)
	ListKeys(ctx context.Context, resourceGroupName string, accountName string) (result cognitiveservices.APIKeys, err error)
	ListSkus(ctx context.Context, resourceGroupName string, accountName string) (result cognitiveservices.AccountSkuListResult, err error)
	ListUsages(ctx context.Context, resourceGroupName string, accountName string, filter string) (result cognitiveservices.UsageListResult, err error)
	RegenerateKey(ctx context.Context, resourceGroupName string, accountName string, parameters cognitiveservices.RegenerateKeyParameters) (result cognitiveservices.APIKeys, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, account cognitiveservices.Account) (result cognitiveservices.AccountsUpdateFuture, err error)
}

var _ AccountsClientAPI = (*cognitiveservices.AccountsClient)(nil)

// DeletedAccountsClientAPI contains the set of methods on the DeletedAccountsClient type.
type DeletedAccountsClientAPI interface {
	Get(ctx context.Context, location string, resourceGroupName string, accountName string) (result cognitiveservices.Account, err error)
	List(ctx context.Context) (result cognitiveservices.AccountListResultPage, err error)
	ListComplete(ctx context.Context) (result cognitiveservices.AccountListResultIterator, err error)
	Purge(ctx context.Context, location string, resourceGroupName string, accountName string) (result cognitiveservices.DeletedAccountsPurgeFuture, err error)
}

var _ DeletedAccountsClientAPI = (*cognitiveservices.DeletedAccountsClient)(nil)

// ResourceSkusClientAPI contains the set of methods on the ResourceSkusClient type.
type ResourceSkusClientAPI interface {
	List(ctx context.Context) (result cognitiveservices.ResourceSkuListResultPage, err error)
	ListComplete(ctx context.Context) (result cognitiveservices.ResourceSkuListResultIterator, err error)
}

var _ ResourceSkusClientAPI = (*cognitiveservices.ResourceSkusClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result cognitiveservices.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result cognitiveservices.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*cognitiveservices.OperationsClient)(nil)

// PrivateEndpointConnectionsClientAPI contains the set of methods on the PrivateEndpointConnectionsClient type.
type PrivateEndpointConnectionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionName string, properties cognitiveservices.PrivateEndpointConnection) (result cognitiveservices.PrivateEndpointConnectionsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionName string) (result cognitiveservices.PrivateEndpointConnectionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionName string) (result cognitiveservices.PrivateEndpointConnection, err error)
	List(ctx context.Context, resourceGroupName string, accountName string) (result cognitiveservices.PrivateEndpointConnectionListResult, err error)
}

var _ PrivateEndpointConnectionsClientAPI = (*cognitiveservices.PrivateEndpointConnectionsClient)(nil)

// PrivateLinkResourcesClientAPI contains the set of methods on the PrivateLinkResourcesClient type.
type PrivateLinkResourcesClientAPI interface {
	List(ctx context.Context, resourceGroupName string, accountName string) (result cognitiveservices.PrivateLinkResourceListResult, err error)
}

var _ PrivateLinkResourcesClientAPI = (*cognitiveservices.PrivateLinkResourcesClient)(nil)
