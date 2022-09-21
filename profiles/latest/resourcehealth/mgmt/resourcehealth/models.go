//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package resourcehealth

import (
	"context"

	original "github.com/Azure/dev/azure-sdk-for-go/services/resourcehealth/mgmt/2020-05-01/resourcehealth"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AvailabilityStateValues = original.AvailabilityStateValues

const (
	AvailabilityStateValuesAvailable   AvailabilityStateValues = original.AvailabilityStateValuesAvailable
	AvailabilityStateValuesDegraded    AvailabilityStateValues = original.AvailabilityStateValuesDegraded
	AvailabilityStateValuesUnavailable AvailabilityStateValues = original.AvailabilityStateValuesUnavailable
	AvailabilityStateValuesUnknown     AvailabilityStateValues = original.AvailabilityStateValuesUnknown
)

type ReasonChronicityTypes = original.ReasonChronicityTypes

const (
	ReasonChronicityTypesPersistent ReasonChronicityTypes = original.ReasonChronicityTypesPersistent
	ReasonChronicityTypesTransient  ReasonChronicityTypes = original.ReasonChronicityTypesTransient
)

type ReasonTypeValues = original.ReasonTypeValues

const (
	ReasonTypeValuesPlanned       ReasonTypeValues = original.ReasonTypeValuesPlanned
	ReasonTypeValuesUnplanned     ReasonTypeValues = original.ReasonTypeValuesUnplanned
	ReasonTypeValuesUserInitiated ReasonTypeValues = original.ReasonTypeValuesUserInitiated
)

type AvailabilityStatus = original.AvailabilityStatus
type AvailabilityStatusListResult = original.AvailabilityStatusListResult
type AvailabilityStatusListResultIterator = original.AvailabilityStatusListResultIterator
type AvailabilityStatusListResultPage = original.AvailabilityStatusListResultPage
type AvailabilityStatusProperties = original.AvailabilityStatusProperties
type AvailabilityStatusPropertiesRecentlyResolved = original.AvailabilityStatusPropertiesRecentlyResolved
type AvailabilityStatusesClient = original.AvailabilityStatusesClient
type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type ErrorResponse = original.ErrorResponse
type ErrorResponseError = original.ErrorResponseError
type ImpactedRegion = original.ImpactedRegion
type ImpactedResourceStatus = original.ImpactedResourceStatus
type ImpactedResourceStatusProperties = original.ImpactedResourceStatusProperties
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationsClient = original.OperationsClient
type ProxyResource = original.ProxyResource
type RecommendedAction = original.RecommendedAction
type Resource = original.Resource
type ServiceImpactingEvent = original.ServiceImpactingEvent
type ServiceImpactingEventIncidentProperties = original.ServiceImpactingEventIncidentProperties
type ServiceImpactingEventStatus = original.ServiceImpactingEventStatus
type StatusBanner = original.StatusBanner
type TrackedResource = original.TrackedResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAvailabilityStatusListResultIterator(page AvailabilityStatusListResultPage) AvailabilityStatusListResultIterator {
	return original.NewAvailabilityStatusListResultIterator(page)
}
func NewAvailabilityStatusListResultPage(cur AvailabilityStatusListResult, getNextPage func(context.Context, AvailabilityStatusListResult) (AvailabilityStatusListResult, error)) AvailabilityStatusListResultPage {
	return original.NewAvailabilityStatusListResultPage(cur, getNextPage)
}
func NewAvailabilityStatusesClient(subscriptionID string) AvailabilityStatusesClient {
	return original.NewAvailabilityStatusesClient(subscriptionID)
}
func NewAvailabilityStatusesClientWithBaseURI(baseURI string, subscriptionID string) AvailabilityStatusesClient {
	return original.NewAvailabilityStatusesClientWithBaseURI(baseURI, subscriptionID)
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
func PossibleAvailabilityStateValuesValues() []AvailabilityStateValues {
	return original.PossibleAvailabilityStateValuesValues()
}
func PossibleReasonChronicityTypesValues() []ReasonChronicityTypes {
	return original.PossibleReasonChronicityTypesValues()
}
func PossibleReasonTypeValuesValues() []ReasonTypeValues {
	return original.PossibleReasonTypeValuesValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
