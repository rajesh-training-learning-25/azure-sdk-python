package msiapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/msi/mgmt/2018-11-30/msi"
	"github.com/Azure/go-autorest/autorest"
)

// SystemAssignedIdentitiesClientAPI contains the set of methods on the SystemAssignedIdentitiesClient type.
type SystemAssignedIdentitiesClientAPI interface {
	GetByScope(ctx context.Context, scope string) (result msi.SystemAssignedIdentity, err error)
}

var _ SystemAssignedIdentitiesClientAPI = (*msi.SystemAssignedIdentitiesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result msi.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result msi.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*msi.OperationsClient)(nil)

// UserAssignedIdentitiesClientAPI contains the set of methods on the UserAssignedIdentitiesClient type.
type UserAssignedIdentitiesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters msi.Identity) (result msi.Identity, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result msi.Identity, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result msi.UserAssignedIdentitiesListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result msi.UserAssignedIdentitiesListResultIterator, err error)
	ListBySubscription(ctx context.Context) (result msi.UserAssignedIdentitiesListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result msi.UserAssignedIdentitiesListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, resourceName string, parameters msi.IdentityUpdate) (result msi.Identity, err error)
}

var _ UserAssignedIdentitiesClientAPI = (*msi.UserAssignedIdentitiesClient)(nil)
