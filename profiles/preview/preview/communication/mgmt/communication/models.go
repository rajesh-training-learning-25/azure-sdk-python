//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package communication

import (
	"context"

	original "github.com/Azure/dev/azure-sdk-for-go/services/preview/communication/mgmt/2020-08-20-preview/communication"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AggregationType = original.AggregationType

const (
	Average AggregationType = original.Average
	Count   AggregationType = original.Count
	Maximum AggregationType = original.Maximum
	Minimum AggregationType = original.Minimum
	Total   AggregationType = original.Total
)

type KeyType = original.KeyType

const (
	Primary   KeyType = original.Primary
	Secondary KeyType = original.Secondary
)

type ProvisioningState = original.ProvisioningState

const (
	Canceled  ProvisioningState = original.Canceled
	Creating  ProvisioningState = original.Creating
	Deleting  ProvisioningState = original.Deleting
	Failed    ProvisioningState = original.Failed
	Moving    ProvisioningState = original.Moving
	Running   ProvisioningState = original.Running
	Succeeded ProvisioningState = original.Succeeded
	Unknown   ProvisioningState = original.Unknown
	Updating  ProvisioningState = original.Updating
)

type Status = original.Status

const (
	StatusCanceled  Status = original.StatusCanceled
	StatusCreating  Status = original.StatusCreating
	StatusDeleting  Status = original.StatusDeleting
	StatusFailed    Status = original.StatusFailed
	StatusMoving    Status = original.StatusMoving
	StatusSucceeded Status = original.StatusSucceeded
)

type BaseClient = original.BaseClient
type Dimension = original.Dimension
type ErrorResponse = original.ErrorResponse
type ErrorResponseError = original.ErrorResponseError
type LinkNotificationHubParameters = original.LinkNotificationHubParameters
type LinkedNotificationHub = original.LinkedNotificationHub
type LocationResource = original.LocationResource
type MetricSpecification = original.MetricSpecification
type NameAvailability = original.NameAvailability
type NameAvailabilityParameters = original.NameAvailabilityParameters
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationList = original.OperationList
type OperationListIterator = original.OperationListIterator
type OperationListPage = original.OperationListPage
type OperationProperties = original.OperationProperties
type OperationStatus = original.OperationStatus
type OperationStatusesClient = original.OperationStatusesClient
type OperationsClient = original.OperationsClient
type RegenerateKeyParameters = original.RegenerateKeyParameters
type Resource = original.Resource
type ServiceClient = original.ServiceClient
type ServiceCreateOrUpdateFuture = original.ServiceCreateOrUpdateFuture
type ServiceDeleteFuture = original.ServiceDeleteFuture
type ServiceKeys = original.ServiceKeys
type ServiceProperties = original.ServiceProperties
type ServiceResource = original.ServiceResource
type ServiceResourceList = original.ServiceResourceList
type ServiceResourceListIterator = original.ServiceResourceListIterator
type ServiceResourceListPage = original.ServiceResourceListPage
type ServiceSpecification = original.ServiceSpecification
type TaggedResource = original.TaggedResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewOperationListIterator(page OperationListPage) OperationListIterator {
	return original.NewOperationListIterator(page)
}
func NewOperationListPage(cur OperationList, getNextPage func(context.Context, OperationList) (OperationList, error)) OperationListPage {
	return original.NewOperationListPage(cur, getNextPage)
}
func NewOperationStatusesClient(subscriptionID string) OperationStatusesClient {
	return original.NewOperationStatusesClient(subscriptionID)
}
func NewOperationStatusesClientWithBaseURI(baseURI string, subscriptionID string) OperationStatusesClient {
	return original.NewOperationStatusesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewServiceClient(subscriptionID string) ServiceClient {
	return original.NewServiceClient(subscriptionID)
}
func NewServiceClientWithBaseURI(baseURI string, subscriptionID string) ServiceClient {
	return original.NewServiceClientWithBaseURI(baseURI, subscriptionID)
}
func NewServiceResourceListIterator(page ServiceResourceListPage) ServiceResourceListIterator {
	return original.NewServiceResourceListIterator(page)
}
func NewServiceResourceListPage(cur ServiceResourceList, getNextPage func(context.Context, ServiceResourceList) (ServiceResourceList, error)) ServiceResourceListPage {
	return original.NewServiceResourceListPage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAggregationTypeValues() []AggregationType {
	return original.PossibleAggregationTypeValues()
}
func PossibleKeyTypeValues() []KeyType {
	return original.PossibleKeyTypeValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleStatusValues() []Status {
	return original.PossibleStatusValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
