// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package advisor

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/advisor/mgmt/2020-01-01/advisor"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type CPUThreshold = original.CPUThreshold

const (
	Five    CPUThreshold = original.Five
	OneFive CPUThreshold = original.OneFive
	OneZero CPUThreshold = original.OneZero
	TwoZero CPUThreshold = original.TwoZero
)

type Category = original.Category

const (
	Cost                  Category = original.Cost
	HighAvailability      Category = original.HighAvailability
	OperationalExcellence Category = original.OperationalExcellence
	Performance           Category = original.Performance
	Security              Category = original.Security
)

type DigestConfigState = original.DigestConfigState

const (
	Active   DigestConfigState = original.Active
	Disabled DigestConfigState = original.Disabled
)

type Impact = original.Impact

const (
	High   Impact = original.High
	Low    Impact = original.Low
	Medium Impact = original.Medium
)

type Risk = original.Risk

const (
	Error   Risk = original.Error
	None    Risk = original.None
	Warning Risk = original.Warning
)

type Scenario = original.Scenario

const (
	Alerts Scenario = original.Alerts
)

type ARMErrorResponseBody = original.ARMErrorResponseBody
type ArmErrorResponse = original.ArmErrorResponse
type BaseClient = original.BaseClient
type ConfigData = original.ConfigData
type ConfigDataProperties = original.ConfigDataProperties
type ConfigurationListResult = original.ConfigurationListResult
type ConfigurationListResultIterator = original.ConfigurationListResultIterator
type ConfigurationListResultPage = original.ConfigurationListResultPage
type ConfigurationsClient = original.ConfigurationsClient
type DigestConfig = original.DigestConfig
type MetadataEntity = original.MetadataEntity
type MetadataEntityListResult = original.MetadataEntityListResult
type MetadataEntityListResultIterator = original.MetadataEntityListResultIterator
type MetadataEntityListResultPage = original.MetadataEntityListResultPage
type MetadataEntityProperties = original.MetadataEntityProperties
type MetadataSupportedValueDetail = original.MetadataSupportedValueDetail
type OperationDisplayInfo = original.OperationDisplayInfo
type OperationEntity = original.OperationEntity
type OperationEntityListResult = original.OperationEntityListResult
type OperationEntityListResultIterator = original.OperationEntityListResultIterator
type OperationEntityListResultPage = original.OperationEntityListResultPage
type OperationsClient = original.OperationsClient
type RecommendationMetadataClient = original.RecommendationMetadataClient
type RecommendationProperties = original.RecommendationProperties
type RecommendationsClient = original.RecommendationsClient
type Resource = original.Resource
type ResourceMetadata = original.ResourceMetadata
type ResourceRecommendationBase = original.ResourceRecommendationBase
type ResourceRecommendationBaseListResult = original.ResourceRecommendationBaseListResult
type ResourceRecommendationBaseListResultIterator = original.ResourceRecommendationBaseListResultIterator
type ResourceRecommendationBaseListResultPage = original.ResourceRecommendationBaseListResultPage
type SetObject = original.SetObject
type ShortDescription = original.ShortDescription
type SuppressionContract = original.SuppressionContract
type SuppressionContractListResult = original.SuppressionContractListResult
type SuppressionContractListResultIterator = original.SuppressionContractListResultIterator
type SuppressionContractListResultPage = original.SuppressionContractListResultPage
type SuppressionProperties = original.SuppressionProperties
type SuppressionsClient = original.SuppressionsClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewConfigurationListResultIterator(page ConfigurationListResultPage) ConfigurationListResultIterator {
	return original.NewConfigurationListResultIterator(page)
}
func NewConfigurationListResultPage(cur ConfigurationListResult, getNextPage func(context.Context, ConfigurationListResult) (ConfigurationListResult, error)) ConfigurationListResultPage {
	return original.NewConfigurationListResultPage(cur, getNextPage)
}
func NewConfigurationsClient(subscriptionID string) ConfigurationsClient {
	return original.NewConfigurationsClient(subscriptionID)
}
func NewConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationsClient {
	return original.NewConfigurationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewMetadataEntityListResultIterator(page MetadataEntityListResultPage) MetadataEntityListResultIterator {
	return original.NewMetadataEntityListResultIterator(page)
}
func NewMetadataEntityListResultPage(cur MetadataEntityListResult, getNextPage func(context.Context, MetadataEntityListResult) (MetadataEntityListResult, error)) MetadataEntityListResultPage {
	return original.NewMetadataEntityListResultPage(cur, getNextPage)
}
func NewOperationEntityListResultIterator(page OperationEntityListResultPage) OperationEntityListResultIterator {
	return original.NewOperationEntityListResultIterator(page)
}
func NewOperationEntityListResultPage(cur OperationEntityListResult, getNextPage func(context.Context, OperationEntityListResult) (OperationEntityListResult, error)) OperationEntityListResultPage {
	return original.NewOperationEntityListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewRecommendationMetadataClient(subscriptionID string) RecommendationMetadataClient {
	return original.NewRecommendationMetadataClient(subscriptionID)
}
func NewRecommendationMetadataClientWithBaseURI(baseURI string, subscriptionID string) RecommendationMetadataClient {
	return original.NewRecommendationMetadataClientWithBaseURI(baseURI, subscriptionID)
}
func NewRecommendationsClient(subscriptionID string) RecommendationsClient {
	return original.NewRecommendationsClient(subscriptionID)
}
func NewRecommendationsClientWithBaseURI(baseURI string, subscriptionID string) RecommendationsClient {
	return original.NewRecommendationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceRecommendationBaseListResultIterator(page ResourceRecommendationBaseListResultPage) ResourceRecommendationBaseListResultIterator {
	return original.NewResourceRecommendationBaseListResultIterator(page)
}
func NewResourceRecommendationBaseListResultPage(cur ResourceRecommendationBaseListResult, getNextPage func(context.Context, ResourceRecommendationBaseListResult) (ResourceRecommendationBaseListResult, error)) ResourceRecommendationBaseListResultPage {
	return original.NewResourceRecommendationBaseListResultPage(cur, getNextPage)
}
func NewSuppressionContractListResultIterator(page SuppressionContractListResultPage) SuppressionContractListResultIterator {
	return original.NewSuppressionContractListResultIterator(page)
}
func NewSuppressionContractListResultPage(cur SuppressionContractListResult, getNextPage func(context.Context, SuppressionContractListResult) (SuppressionContractListResult, error)) SuppressionContractListResultPage {
	return original.NewSuppressionContractListResultPage(cur, getNextPage)
}
func NewSuppressionsClient(subscriptionID string) SuppressionsClient {
	return original.NewSuppressionsClient(subscriptionID)
}
func NewSuppressionsClientWithBaseURI(baseURI string, subscriptionID string) SuppressionsClient {
	return original.NewSuppressionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleCPUThresholdValues() []CPUThreshold {
	return original.PossibleCPUThresholdValues()
}
func PossibleCategoryValues() []Category {
	return original.PossibleCategoryValues()
}
func PossibleDigestConfigStateValues() []DigestConfigState {
	return original.PossibleDigestConfigStateValues()
}
func PossibleImpactValues() []Impact {
	return original.PossibleImpactValues()
}
func PossibleRiskValues() []Risk {
	return original.PossibleRiskValues()
}
func PossibleScenarioValues() []Scenario {
	return original.PossibleScenarioValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
