// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package sql

import original "github.com/Azure/azure-sdk-for-go/services/sql/mgmt/2017-10-01-preview/sql"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type DatabaseOperationsClient = original.DatabaseOperationsClient
type DatabaseVulnerabilityAssessmentScansClient = original.DatabaseVulnerabilityAssessmentScansClient
type ElasticPoolOperationsClient = original.ElasticPoolOperationsClient
type ManagementOperationState = original.ManagementOperationState

const (
	CancelInProgress ManagementOperationState = original.CancelInProgress
	Cancelled        ManagementOperationState = original.Cancelled
	Failed           ManagementOperationState = original.Failed
	InProgress       ManagementOperationState = original.InProgress
	Pending          ManagementOperationState = original.Pending
	Succeeded        ManagementOperationState = original.Succeeded
)

type VulnerabilityAssessmentScanState = original.VulnerabilityAssessmentScanState

const (
	VulnerabilityAssessmentScanStateFailed      VulnerabilityAssessmentScanState = original.VulnerabilityAssessmentScanStateFailed
	VulnerabilityAssessmentScanStateFailedToRun VulnerabilityAssessmentScanState = original.VulnerabilityAssessmentScanStateFailedToRun
	VulnerabilityAssessmentScanStateInProgress  VulnerabilityAssessmentScanState = original.VulnerabilityAssessmentScanStateInProgress
	VulnerabilityAssessmentScanStatePassed      VulnerabilityAssessmentScanState = original.VulnerabilityAssessmentScanStatePassed
)

type VulnerabilityAssessmentScanTriggerType = original.VulnerabilityAssessmentScanTriggerType

const (
	OnDemand  VulnerabilityAssessmentScanTriggerType = original.OnDemand
	Recurring VulnerabilityAssessmentScanTriggerType = original.Recurring
)

type DatabaseOperation = original.DatabaseOperation
type DatabaseOperationListResult = original.DatabaseOperationListResult
type DatabaseOperationListResultIterator = original.DatabaseOperationListResultIterator
type DatabaseOperationListResultPage = original.DatabaseOperationListResultPage
type DatabaseOperationProperties = original.DatabaseOperationProperties
type DatabaseVulnerabilityAssessment = original.DatabaseVulnerabilityAssessment
type DatabaseVulnerabilityAssessmentProperties = original.DatabaseVulnerabilityAssessmentProperties
type DatabaseVulnerabilityAssessmentScanExportProperties = original.DatabaseVulnerabilityAssessmentScanExportProperties
type DatabaseVulnerabilityAssessmentScansExecuteFuture = original.DatabaseVulnerabilityAssessmentScansExecuteFuture
type DatabaseVulnerabilityAssessmentScansExport = original.DatabaseVulnerabilityAssessmentScansExport
type ElasticPoolOperation = original.ElasticPoolOperation
type ElasticPoolOperationListResult = original.ElasticPoolOperationListResult
type ElasticPoolOperationListResultIterator = original.ElasticPoolOperationListResultIterator
type ElasticPoolOperationListResultPage = original.ElasticPoolOperationListResultPage
type ElasticPoolOperationProperties = original.ElasticPoolOperationProperties
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type Sku = original.Sku
type VulnerabilityAssessmentRecurringScansProperties = original.VulnerabilityAssessmentRecurringScansProperties
type VulnerabilityAssessmentScanError = original.VulnerabilityAssessmentScanError
type VulnerabilityAssessmentScanRecord = original.VulnerabilityAssessmentScanRecord
type VulnerabilityAssessmentScanRecordListResult = original.VulnerabilityAssessmentScanRecordListResult
type VulnerabilityAssessmentScanRecordListResultIterator = original.VulnerabilityAssessmentScanRecordListResultIterator
type VulnerabilityAssessmentScanRecordListResultPage = original.VulnerabilityAssessmentScanRecordListResultPage
type VulnerabilityAssessmentScanRecordProperties = original.VulnerabilityAssessmentScanRecordProperties

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewDatabaseOperationsClient(subscriptionID string) DatabaseOperationsClient {
	return original.NewDatabaseOperationsClient(subscriptionID)
}
func NewDatabaseOperationsClientWithBaseURI(baseURI string, subscriptionID string) DatabaseOperationsClient {
	return original.NewDatabaseOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDatabaseVulnerabilityAssessmentScansClient(subscriptionID string) DatabaseVulnerabilityAssessmentScansClient {
	return original.NewDatabaseVulnerabilityAssessmentScansClient(subscriptionID)
}
func NewDatabaseVulnerabilityAssessmentScansClientWithBaseURI(baseURI string, subscriptionID string) DatabaseVulnerabilityAssessmentScansClient {
	return original.NewDatabaseVulnerabilityAssessmentScansClientWithBaseURI(baseURI, subscriptionID)
}
func NewElasticPoolOperationsClient(subscriptionID string) ElasticPoolOperationsClient {
	return original.NewElasticPoolOperationsClient(subscriptionID)
}
func NewElasticPoolOperationsClientWithBaseURI(baseURI string, subscriptionID string) ElasticPoolOperationsClient {
	return original.NewElasticPoolOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleManagementOperationStateValues() []ManagementOperationState {
	return original.PossibleManagementOperationStateValues()
}
func PossibleVulnerabilityAssessmentScanStateValues() []VulnerabilityAssessmentScanState {
	return original.PossibleVulnerabilityAssessmentScanStateValues()
}
func PossibleVulnerabilityAssessmentScanTriggerTypeValues() []VulnerabilityAssessmentScanTriggerType {
	return original.PossibleVulnerabilityAssessmentScanTriggerTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
