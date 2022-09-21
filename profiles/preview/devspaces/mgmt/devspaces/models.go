//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package devspaces

import (
	"context"

	original "github.com/Azure/dev/azure-sdk-for-go/services/devspaces/mgmt/2019-04-01/devspaces"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type InstanceType = original.InstanceType

const (
	InstanceTypeKubernetes                            InstanceType = original.InstanceTypeKubernetes
	InstanceTypeOrchestratorSpecificConnectionDetails InstanceType = original.InstanceTypeOrchestratorSpecificConnectionDetails
)

type ProvisioningState = original.ProvisioningState

const (
	Canceled  ProvisioningState = original.Canceled
	Creating  ProvisioningState = original.Creating
	Deleted   ProvisioningState = original.Deleted
	Deleting  ProvisioningState = original.Deleting
	Failed    ProvisioningState = original.Failed
	Succeeded ProvisioningState = original.Succeeded
	Updating  ProvisioningState = original.Updating
)

type SkuTier = original.SkuTier

const (
	Standard SkuTier = original.Standard
)

type BaseClient = original.BaseClient
type BasicOrchestratorSpecificConnectionDetails = original.BasicOrchestratorSpecificConnectionDetails
type ContainerHostMapping = original.ContainerHostMapping
type ContainerHostMappingsClient = original.ContainerHostMappingsClient
type Controller = original.Controller
type ControllerConnectionDetails = original.ControllerConnectionDetails
type ControllerConnectionDetailsList = original.ControllerConnectionDetailsList
type ControllerList = original.ControllerList
type ControllerListIterator = original.ControllerListIterator
type ControllerListPage = original.ControllerListPage
type ControllerProperties = original.ControllerProperties
type ControllerUpdateParameters = original.ControllerUpdateParameters
type ControllerUpdateParametersProperties = original.ControllerUpdateParametersProperties
type ControllersClient = original.ControllersClient
type ControllersCreateFuture = original.ControllersCreateFuture
type ControllersDeleteFuture = original.ControllersDeleteFuture
type ErrorDetails = original.ErrorDetails
type ErrorResponse = original.ErrorResponse
type KubernetesConnectionDetails = original.KubernetesConnectionDetails
type ListConnectionDetailsParameters = original.ListConnectionDetailsParameters
type OperationsClient = original.OperationsClient
type OrchestratorSpecificConnectionDetails = original.OrchestratorSpecificConnectionDetails
type Resource = original.Resource
type ResourceProviderOperationDefinition = original.ResourceProviderOperationDefinition
type ResourceProviderOperationDisplay = original.ResourceProviderOperationDisplay
type ResourceProviderOperationList = original.ResourceProviderOperationList
type ResourceProviderOperationListIterator = original.ResourceProviderOperationListIterator
type ResourceProviderOperationListPage = original.ResourceProviderOperationListPage
type Sku = original.Sku
type TrackedResource = original.TrackedResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewContainerHostMappingsClient(subscriptionID string) ContainerHostMappingsClient {
	return original.NewContainerHostMappingsClient(subscriptionID)
}
func NewContainerHostMappingsClientWithBaseURI(baseURI string, subscriptionID string) ContainerHostMappingsClient {
	return original.NewContainerHostMappingsClientWithBaseURI(baseURI, subscriptionID)
}
func NewControllerListIterator(page ControllerListPage) ControllerListIterator {
	return original.NewControllerListIterator(page)
}
func NewControllerListPage(cur ControllerList, getNextPage func(context.Context, ControllerList) (ControllerList, error)) ControllerListPage {
	return original.NewControllerListPage(cur, getNextPage)
}
func NewControllersClient(subscriptionID string) ControllersClient {
	return original.NewControllersClient(subscriptionID)
}
func NewControllersClientWithBaseURI(baseURI string, subscriptionID string) ControllersClient {
	return original.NewControllersClientWithBaseURI(baseURI, subscriptionID)
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
func PossibleInstanceTypeValues() []InstanceType {
	return original.PossibleInstanceTypeValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
