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

package adhybridhealthservice

import original "github.com/Azure/azure-sdk-for-go/services/adhybridhealthservice/mgmt/2014-01-01/adhybridhealthservice"

type AddomainservicemembersClient = original.AddomainservicemembersClient
type AddsservicemembersClient = original.AddsservicemembersClient
type AddsservicesClient = original.AddsservicesClient
type AlertsClient = original.AlertsClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type ConfigurationClient = original.ConfigurationClient
type DimensionsClient = original.DimensionsClient
type AlgorithmStepType = original.AlgorithmStepType

const (
	ConnectorFilter         AlgorithmStepType = original.ConnectorFilter
	Deprovisioning          AlgorithmStepType = original.Deprovisioning
	ExportFlow              AlgorithmStepType = original.ExportFlow
	ImportFlow              AlgorithmStepType = original.ImportFlow
	Join                    AlgorithmStepType = original.Join
	MvDeletion              AlgorithmStepType = original.MvDeletion
	MvObjectTypeChange      AlgorithmStepType = original.MvObjectTypeChange
	Projection              AlgorithmStepType = original.Projection
	Provisioning            AlgorithmStepType = original.Provisioning
	Recall                  AlgorithmStepType = original.Recall
	Staging                 AlgorithmStepType = original.Staging
	Undefined               AlgorithmStepType = original.Undefined
	ValidateConnectorFilter AlgorithmStepType = original.ValidateConnectorFilter
)

type AttributeDeltaOperationType = original.AttributeDeltaOperationType

const (
	AttributeDeltaOperationTypeAdd       AttributeDeltaOperationType = original.AttributeDeltaOperationTypeAdd
	AttributeDeltaOperationTypeDelete    AttributeDeltaOperationType = original.AttributeDeltaOperationTypeDelete
	AttributeDeltaOperationTypeReplace   AttributeDeltaOperationType = original.AttributeDeltaOperationTypeReplace
	AttributeDeltaOperationTypeUndefined AttributeDeltaOperationType = original.AttributeDeltaOperationTypeUndefined
	AttributeDeltaOperationTypeUpdate    AttributeDeltaOperationType = original.AttributeDeltaOperationTypeUpdate
)

type AttributeMappingType = original.AttributeMappingType

const (
	Constant AttributeMappingType = original.Constant
	Direct   AttributeMappingType = original.Direct
	DnPart   AttributeMappingType = original.DnPart
	Script   AttributeMappingType = original.Script
)

type DeltaOperationType = original.DeltaOperationType

const (
	DeltaOperationTypeAdd       DeltaOperationType = original.DeltaOperationTypeAdd
	DeltaOperationTypeDeletAdd  DeltaOperationType = original.DeltaOperationTypeDeletAdd
	DeltaOperationTypeDelete    DeltaOperationType = original.DeltaOperationTypeDelete
	DeltaOperationTypeNone      DeltaOperationType = original.DeltaOperationTypeNone
	DeltaOperationTypeObsolete  DeltaOperationType = original.DeltaOperationTypeObsolete
	DeltaOperationTypeReplace   DeltaOperationType = original.DeltaOperationTypeReplace
	DeltaOperationTypeUndefined DeltaOperationType = original.DeltaOperationTypeUndefined
	DeltaOperationTypeUpdate    DeltaOperationType = original.DeltaOperationTypeUpdate
)

type HealthStatus = original.HealthStatus

const (
	Error        HealthStatus = original.Error
	Healthy      HealthStatus = original.Healthy
	Missing      HealthStatus = original.Missing
	NotMonitored HealthStatus = original.NotMonitored
	Warning      HealthStatus = original.Warning
)

type Level = original.Level

const (
	LevelError      Level = original.LevelError
	LevelPreWarning Level = original.LevelPreWarning
	LevelWarning    Level = original.LevelWarning
)

type MonitoringLevel = original.MonitoringLevel

const (
	Full    MonitoringLevel = original.Full
	Off     MonitoringLevel = original.Off
	Partial MonitoringLevel = original.Partial
)

type PasswordOperationTypes = original.PasswordOperationTypes

const (
	PasswordOperationTypesChange    PasswordOperationTypes = original.PasswordOperationTypesChange
	PasswordOperationTypesSet       PasswordOperationTypes = original.PasswordOperationTypesSet
	PasswordOperationTypesUndefined PasswordOperationTypes = original.PasswordOperationTypesUndefined
)

type RunStepOperationType = original.RunStepOperationType

const (
	RunStepOperationTypeApplyRules                RunStepOperationType = original.RunStepOperationTypeApplyRules
	RunStepOperationTypeDeltaImport               RunStepOperationType = original.RunStepOperationTypeDeltaImport
	RunStepOperationTypeDeltaSynchronization      RunStepOperationType = original.RunStepOperationTypeDeltaSynchronization
	RunStepOperationTypeExport                    RunStepOperationType = original.RunStepOperationTypeExport
	RunStepOperationTypeFullExport                RunStepOperationType = original.RunStepOperationTypeFullExport
	RunStepOperationTypeFullImport                RunStepOperationType = original.RunStepOperationTypeFullImport
	RunStepOperationTypeFullImportReevaluateRules RunStepOperationType = original.RunStepOperationTypeFullImportReevaluateRules
	RunStepOperationTypeFullSynchornization       RunStepOperationType = original.RunStepOperationTypeFullSynchornization
	RunStepOperationTypeUndefined                 RunStepOperationType = original.RunStepOperationTypeUndefined
)

type ServerDisabledReason = original.ServerDisabledReason

const (
	DeletedFromPortal       ServerDisabledReason = original.DeletedFromPortal
	DisabledDueToInactivity ServerDisabledReason = original.DisabledDueToInactivity
	GdprStopCollection      ServerDisabledReason = original.GdprStopCollection
	None                    ServerDisabledReason = original.None
)

type ServiceType = original.ServiceType

const (
	ServiceTypeAadConnectSync ServiceType = original.ServiceTypeAadConnectSync
	ServiceTypeDirSync        ServiceType = original.ServiceTypeDirSync
	ServiceTypeUndefined      ServiceType = original.ServiceTypeUndefined
)

type State = original.State

const (
	Active                   State = original.Active
	ResolvedByPositiveResult State = original.ResolvedByPositiveResult
	ResolvedByStateChange    State = original.ResolvedByStateChange
	ResolvedByTimer          State = original.ResolvedByTimer
	ResolvedManually         State = original.ResolvedManually
)

type ValueDeltaOperationType = original.ValueDeltaOperationType

const (
	ValueDeltaOperationTypeAdd       ValueDeltaOperationType = original.ValueDeltaOperationTypeAdd
	ValueDeltaOperationTypeDelete    ValueDeltaOperationType = original.ValueDeltaOperationTypeDelete
	ValueDeltaOperationTypeUndefined ValueDeltaOperationType = original.ValueDeltaOperationTypeUndefined
	ValueDeltaOperationTypeUpdate    ValueDeltaOperationType = original.ValueDeltaOperationTypeUpdate
)

type ValueType = original.ValueType

const (
	ValueTypeBinary    ValueType = original.ValueTypeBinary
	ValueTypeBoolean   ValueType = original.ValueTypeBoolean
	ValueTypeDn        ValueType = original.ValueTypeDn
	ValueTypeInteger   ValueType = original.ValueTypeInteger
	ValueTypeString    ValueType = original.ValueTypeString
	ValueTypeUndefined ValueType = original.ValueTypeUndefined
)

type AdditionalInformation = original.AdditionalInformation
type AddsConfiguration = original.AddsConfiguration
type AddsConfigurationIterator = original.AddsConfigurationIterator
type AddsConfigurationPage = original.AddsConfigurationPage
type AddsServiceMember = original.AddsServiceMember
type AddsServiceMembers = original.AddsServiceMembers
type AddsServiceMembersIterator = original.AddsServiceMembersIterator
type AddsServiceMembersPage = original.AddsServiceMembersPage
type Agent = original.Agent
type Alert = original.Alert
type AlertFeedback = original.AlertFeedback
type AlertFeedbacks = original.AlertFeedbacks
type Alerts = original.Alerts
type AlertsIterator = original.AlertsIterator
type AlertsPage = original.AlertsPage
type AssociatedObject = original.AssociatedObject
type AttributeDelta = original.AttributeDelta
type AttributeMapping = original.AttributeMapping
type AttributeMppingSource = original.AttributeMppingSource
type Bool = original.Bool
type ChangeNotReimported = original.ChangeNotReimported
type ChangeNotReimportedDelta = original.ChangeNotReimportedDelta
type ChangeNotReimportedEntry = original.ChangeNotReimportedEntry
type Connector = original.Connector
type ConnectorConnectionError = original.ConnectorConnectionError
type ConnectorConnectionErrors = original.ConnectorConnectionErrors
type ConnectorObjectError = original.ConnectorObjectError
type ConnectorObjectErrors = original.ConnectorObjectErrors
type Connectors = original.Connectors
type Credential = original.Credential
type Credentials = original.Credentials
type DataFreshnessDetails = original.DataFreshnessDetails
type Dimension = original.Dimension
type Dimensions = original.Dimensions
type DimensionsIterator = original.DimensionsIterator
type DimensionsPage = original.DimensionsPage
type Display = original.Display
type ErrorCount = original.ErrorCount
type ErrorCounts = original.ErrorCounts
type ErrorDetail = original.ErrorDetail
type ErrorReportUsersEntries = original.ErrorReportUsersEntries
type ErrorReportUsersEntry = original.ErrorReportUsersEntry
type ExportError = original.ExportError
type ExportErrors = original.ExportErrors
type ExportStatus = original.ExportStatus
type ExportStatuses = original.ExportStatuses
type ExportStatusesIterator = original.ExportStatusesIterator
type ExportStatusesPage = original.ExportStatusesPage
type ExtensionErrorInfo = original.ExtensionErrorInfo
type ForestSummary = original.ForestSummary
type GlobalConfiguration = original.GlobalConfiguration
type GlobalConfigurations = original.GlobalConfigurations
type HelpLink = original.HelpLink
type Hotfix = original.Hotfix
type Hotfixes = original.Hotfixes
type ImportError = original.ImportError
type ImportErrors = original.ImportErrors
type InboundReplicationNeighbor = original.InboundReplicationNeighbor
type InboundReplicationNeighbors = original.InboundReplicationNeighbors
type Item = original.Item
type Items = original.Items
type MergedExportError = original.MergedExportError
type MergedExportErrors = original.MergedExportErrors
type MetricGroup = original.MetricGroup
type MetricMetadata = original.MetricMetadata
type MetricMetadataList = original.MetricMetadataList
type MetricMetadataListIterator = original.MetricMetadataListIterator
type MetricMetadataListPage = original.MetricMetadataListPage
type Metrics = original.Metrics
type MetricSet = original.MetricSet
type MetricSets = original.MetricSets
type MetricsIterator = original.MetricsIterator
type MetricsPage = original.MetricsPage
type ModuleConfiguration = original.ModuleConfiguration
type ModuleConfigurations = original.ModuleConfigurations
type ObjectWithSyncError = original.ObjectWithSyncError
type Operation = original.Operation
type OperationListResponse = original.OperationListResponse
type OperationListResponseIterator = original.OperationListResponseIterator
type OperationListResponsePage = original.OperationListResponsePage
type Partition = original.Partition
type PartitionScope = original.PartitionScope
type PasswordHashSyncConfiguration = original.PasswordHashSyncConfiguration
type PasswordManagementSettings = original.PasswordManagementSettings
type ReplicationStatus = original.ReplicationStatus
type ReplicationSummary = original.ReplicationSummary
type ReplicationSummaryList = original.ReplicationSummaryList
type RuleErrorInfo = original.RuleErrorInfo
type RunProfile = original.RunProfile
type RunProfiles = original.RunProfiles
type RunStep = original.RunStep
type Service = original.Service
type ServiceConfiguration = original.ServiceConfiguration
type ServiceMember = original.ServiceMember
type ServiceMemberProperties = original.ServiceMemberProperties
type ServiceMembers = original.ServiceMembers
type ServiceMembersIterator = original.ServiceMembersIterator
type ServiceMembersPage = original.ServiceMembersPage
type ServiceProperties = original.ServiceProperties
type Services = original.Services
type ServicesIterator = original.ServicesIterator
type ServicesPage = original.ServicesPage
type TabularExportError = original.TabularExportError
type Tenant = original.Tenant
type TenantOnboardingDetails = original.TenantOnboardingDetails
type ValueDelta = original.ValueDelta
type OperationsClient = original.OperationsClient
type ServicemembersClient = original.ServicemembersClient
type ServicesClient = original.ServicesClient
type ServimembersClient = original.ServimembersClient

func NewAddomainservicemembersClient() AddomainservicemembersClient {
	return original.NewAddomainservicemembersClient()
}
func NewAddomainservicemembersClientWithBaseURI(baseURI string) AddomainservicemembersClient {
	return original.NewAddomainservicemembersClientWithBaseURI(baseURI)
}
func NewAddsservicemembersClient() AddsservicemembersClient {
	return original.NewAddsservicemembersClient()
}
func NewAddsservicemembersClientWithBaseURI(baseURI string) AddsservicemembersClient {
	return original.NewAddsservicemembersClientWithBaseURI(baseURI)
}
func NewAddsservicesClient() AddsservicesClient {
	return original.NewAddsservicesClient()
}
func NewAddsservicesClientWithBaseURI(baseURI string) AddsservicesClient {
	return original.NewAddsservicesClientWithBaseURI(baseURI)
}
func NewAlertsClient() AlertsClient {
	return original.NewAlertsClient()
}
func NewAlertsClientWithBaseURI(baseURI string) AlertsClient {
	return original.NewAlertsClientWithBaseURI(baseURI)
}
func New() BaseClient {
	return original.New()
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func NewConfigurationClient() ConfigurationClient {
	return original.NewConfigurationClient()
}
func NewConfigurationClientWithBaseURI(baseURI string) ConfigurationClient {
	return original.NewConfigurationClientWithBaseURI(baseURI)
}
func NewDimensionsClient() DimensionsClient {
	return original.NewDimensionsClient()
}
func NewDimensionsClientWithBaseURI(baseURI string) DimensionsClient {
	return original.NewDimensionsClientWithBaseURI(baseURI)
}
func PossibleAlgorithmStepTypeValues() []AlgorithmStepType {
	return original.PossibleAlgorithmStepTypeValues()
}
func PossibleAttributeDeltaOperationTypeValues() []AttributeDeltaOperationType {
	return original.PossibleAttributeDeltaOperationTypeValues()
}
func PossibleAttributeMappingTypeValues() []AttributeMappingType {
	return original.PossibleAttributeMappingTypeValues()
}
func PossibleDeltaOperationTypeValues() []DeltaOperationType {
	return original.PossibleDeltaOperationTypeValues()
}
func PossibleHealthStatusValues() []HealthStatus {
	return original.PossibleHealthStatusValues()
}
func PossibleLevelValues() []Level {
	return original.PossibleLevelValues()
}
func PossibleMonitoringLevelValues() []MonitoringLevel {
	return original.PossibleMonitoringLevelValues()
}
func PossiblePasswordOperationTypesValues() []PasswordOperationTypes {
	return original.PossiblePasswordOperationTypesValues()
}
func PossibleRunStepOperationTypeValues() []RunStepOperationType {
	return original.PossibleRunStepOperationTypeValues()
}
func PossibleServerDisabledReasonValues() []ServerDisabledReason {
	return original.PossibleServerDisabledReasonValues()
}
func PossibleServiceTypeValues() []ServiceType {
	return original.PossibleServiceTypeValues()
}
func PossibleStateValues() []State {
	return original.PossibleStateValues()
}
func PossibleValueDeltaOperationTypeValues() []ValueDeltaOperationType {
	return original.PossibleValueDeltaOperationTypeValues()
}
func PossibleValueTypeValues() []ValueType {
	return original.PossibleValueTypeValues()
}
func NewOperationsClient() OperationsClient {
	return original.NewOperationsClient()
}
func NewOperationsClientWithBaseURI(baseURI string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI)
}
func NewServicemembersClient() ServicemembersClient {
	return original.NewServicemembersClient()
}
func NewServicemembersClientWithBaseURI(baseURI string) ServicemembersClient {
	return original.NewServicemembersClientWithBaseURI(baseURI)
}
func NewServicesClient() ServicesClient {
	return original.NewServicesClient()
}
func NewServicesClientWithBaseURI(baseURI string) ServicesClient {
	return original.NewServicesClientWithBaseURI(baseURI)
}
func NewServimembersClient() ServimembersClient {
	return original.NewServimembersClient()
}
func NewServimembersClientWithBaseURI(baseURI string) ServimembersClient {
	return original.NewServimembersClientWithBaseURI(baseURI)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
