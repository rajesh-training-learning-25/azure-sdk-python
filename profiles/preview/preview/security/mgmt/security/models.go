// +build go1.9

// Copyright 2019 Microsoft Corporation
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

package security

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v2.0/security"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AadConnectivityState = original.AadConnectivityState

const (
	Connected   AadConnectivityState = original.Connected
	Discovered  AadConnectivityState = original.Discovered
	NotLicensed AadConnectivityState = original.NotLicensed
)

type AlertNotifications = original.AlertNotifications

const (
	Off AlertNotifications = original.Off
	On  AlertNotifications = original.On
)

type AlertsToAdmins = original.AlertsToAdmins

const (
	AlertsToAdminsOff AlertsToAdmins = original.AlertsToAdminsOff
	AlertsToAdminsOn  AlertsToAdmins = original.AlertsToAdminsOn
)

type AutoProvision = original.AutoProvision

const (
	AutoProvisionOff AutoProvision = original.AutoProvisionOff
	AutoProvisionOn  AutoProvision = original.AutoProvisionOn
)

type ConnectionType = original.ConnectionType

const (
	External ConnectionType = original.External
	Internal ConnectionType = original.Internal
)

type Direction = original.Direction

const (
	Inbound  Direction = original.Inbound
	Outbound Direction = original.Outbound
)

type ExternalSecuritySolutionKind = original.ExternalSecuritySolutionKind

const (
	AAD ExternalSecuritySolutionKind = original.AAD
	ATA ExternalSecuritySolutionKind = original.ATA
	CEF ExternalSecuritySolutionKind = original.CEF
)

type Family = original.Family

const (
	Ngfw    Family = original.Ngfw
	SaasWaf Family = original.SaasWaf
	Va      Family = original.Va
	Waf     Family = original.Waf
)

type KindEnum = original.KindEnum

const (
	KindAAD                      KindEnum = original.KindAAD
	KindATA                      KindEnum = original.KindATA
	KindCEF                      KindEnum = original.KindCEF
	KindExternalSecuritySolution KindEnum = original.KindExternalSecuritySolution
)

type Operator = original.Operator

const (
	Contains             Operator = original.Contains
	EndsWith             Operator = original.EndsWith
	Equals               Operator = original.Equals
	EqualsIgnoreCase     Operator = original.EqualsIgnoreCase
	GreaterThan          Operator = original.GreaterThan
	GreaterThanOrEqualTo Operator = original.GreaterThanOrEqualTo
	LesserThan           Operator = original.LesserThan
	LesserThanOrEqualTo  Operator = original.LesserThanOrEqualTo
	NotEquals            Operator = original.NotEquals
	NotEqualsIgnoreCase  Operator = original.NotEqualsIgnoreCase
	StartsWith           Operator = original.StartsWith
)

type PricingTier = original.PricingTier

const (
	Free     PricingTier = original.Free
	Standard PricingTier = original.Standard
)

type PropertyType = original.PropertyType

const (
	Boolean PropertyType = original.Boolean
	Integer PropertyType = original.Integer
	Number  PropertyType = original.Number
	String  PropertyType = original.String
)

type Protocol = original.Protocol

const (
	All Protocol = original.All
	TCP Protocol = original.TCP
	UDP Protocol = original.UDP
)

type ReportedSeverity = original.ReportedSeverity

const (
	High          ReportedSeverity = original.High
	Informational ReportedSeverity = original.Informational
	Low           ReportedSeverity = original.Low
	Medium        ReportedSeverity = original.Medium
)

type SettingKind = original.SettingKind

const (
	SettingKindAlertSuppressionSetting SettingKind = original.SettingKindAlertSuppressionSetting
	SettingKindDataExportSetting       SettingKind = original.SettingKindDataExportSetting
)

type State = original.State

const (
	Failed      State = original.Failed
	Passed      State = original.Passed
	Skipped     State = original.Skipped
	Unsupported State = original.Unsupported
)

type Status = original.Status

const (
	Initiated Status = original.Initiated
	Revoked   Status = original.Revoked
)

type StatusReason = original.StatusReason

const (
	Expired               StatusReason = original.Expired
	NewerRequestInitiated StatusReason = original.NewerRequestInitiated
	UserRequested         StatusReason = original.UserRequested
)

type TransportProtocol = original.TransportProtocol

const (
	TransportProtocolTCP TransportProtocol = original.TransportProtocolTCP
	TransportProtocolUDP TransportProtocol = original.TransportProtocolUDP
)

type ValueType = original.ValueType

const (
	ValueTypeIPCidr ValueType = original.ValueTypeIPCidr
	ValueTypeString ValueType = original.ValueTypeString
)

type AadConnectivityState1 = original.AadConnectivityState1
type AadExternalSecuritySolution = original.AadExternalSecuritySolution
type AadSolutionProperties = original.AadSolutionProperties
type AdaptiveNetworkHardening = original.AdaptiveNetworkHardening
type AdaptiveNetworkHardeningEnforceRequest = original.AdaptiveNetworkHardeningEnforceRequest
type AdaptiveNetworkHardeningProperties = original.AdaptiveNetworkHardeningProperties
type AdaptiveNetworkHardeningsClient = original.AdaptiveNetworkHardeningsClient
type AdaptiveNetworkHardeningsEnforceFuture = original.AdaptiveNetworkHardeningsEnforceFuture
type AdaptiveNetworkHardeningsList = original.AdaptiveNetworkHardeningsList
type AdaptiveNetworkHardeningsListIterator = original.AdaptiveNetworkHardeningsListIterator
type AdaptiveNetworkHardeningsListPage = original.AdaptiveNetworkHardeningsListPage
type AdvancedThreatProtectionClient = original.AdvancedThreatProtectionClient
type AdvancedThreatProtectionProperties = original.AdvancedThreatProtectionProperties
type AdvancedThreatProtectionSetting = original.AdvancedThreatProtectionSetting
type Alert = original.Alert
type AlertConfidenceReason = original.AlertConfidenceReason
type AlertEntity = original.AlertEntity
type AlertList = original.AlertList
type AlertListIterator = original.AlertListIterator
type AlertListPage = original.AlertListPage
type AlertProperties = original.AlertProperties
type AlertsClient = original.AlertsClient
type AllowedConnectionsClient = original.AllowedConnectionsClient
type AllowedConnectionsList = original.AllowedConnectionsList
type AllowedConnectionsListIterator = original.AllowedConnectionsListIterator
type AllowedConnectionsListPage = original.AllowedConnectionsListPage
type AllowedConnectionsResource = original.AllowedConnectionsResource
type AllowedConnectionsResourceProperties = original.AllowedConnectionsResourceProperties
type AllowlistCustomAlertRule = original.AllowlistCustomAlertRule
type AscLocation = original.AscLocation
type AscLocationList = original.AscLocationList
type AscLocationListIterator = original.AscLocationListIterator
type AscLocationListPage = original.AscLocationListPage
type AtaExternalSecuritySolution = original.AtaExternalSecuritySolution
type AtaSolutionProperties = original.AtaSolutionProperties
type AutoProvisioningSetting = original.AutoProvisioningSetting
type AutoProvisioningSettingList = original.AutoProvisioningSettingList
type AutoProvisioningSettingListIterator = original.AutoProvisioningSettingListIterator
type AutoProvisioningSettingListPage = original.AutoProvisioningSettingListPage
type AutoProvisioningSettingProperties = original.AutoProvisioningSettingProperties
type AutoProvisioningSettingsClient = original.AutoProvisioningSettingsClient
type BaseClient = original.BaseClient
type BasicExternalSecuritySolution = original.BasicExternalSecuritySolution
type CefExternalSecuritySolution = original.CefExternalSecuritySolution
type CefSolutionProperties = original.CefSolutionProperties
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type Compliance = original.Compliance
type ComplianceList = original.ComplianceList
type ComplianceListIterator = original.ComplianceListIterator
type ComplianceListPage = original.ComplianceListPage
type ComplianceProperties = original.ComplianceProperties
type ComplianceSegment = original.ComplianceSegment
type CompliancesClient = original.CompliancesClient
type ConnectableResource = original.ConnectableResource
type ConnectedResource = original.ConnectedResource
type ConnectedWorkspace = original.ConnectedWorkspace
type Contact = original.Contact
type ContactList = original.ContactList
type ContactListIterator = original.ContactListIterator
type ContactListPage = original.ContactListPage
type ContactProperties = original.ContactProperties
type ContactsClient = original.ContactsClient
type CustomAlertRule = original.CustomAlertRule
type DataExportSetting = original.DataExportSetting
type DataExportSettingProperties = original.DataExportSettingProperties
type DenylistCustomAlertRule = original.DenylistCustomAlertRule
type DeviceSecurityGroup = original.DeviceSecurityGroup
type DeviceSecurityGroupList = original.DeviceSecurityGroupList
type DeviceSecurityGroupListIterator = original.DeviceSecurityGroupListIterator
type DeviceSecurityGroupListPage = original.DeviceSecurityGroupListPage
type DeviceSecurityGroupProperties = original.DeviceSecurityGroupProperties
type DeviceSecurityGroupsClient = original.DeviceSecurityGroupsClient
type DiscoveredSecuritySolution = original.DiscoveredSecuritySolution
type DiscoveredSecuritySolutionList = original.DiscoveredSecuritySolutionList
type DiscoveredSecuritySolutionListIterator = original.DiscoveredSecuritySolutionListIterator
type DiscoveredSecuritySolutionListPage = original.DiscoveredSecuritySolutionListPage
type DiscoveredSecuritySolutionProperties = original.DiscoveredSecuritySolutionProperties
type DiscoveredSecuritySolutionsClient = original.DiscoveredSecuritySolutionsClient
type ETag = original.ETag
type EffectiveNetworkSecurityGroups = original.EffectiveNetworkSecurityGroups
type ExternalSecuritySolution = original.ExternalSecuritySolution
type ExternalSecuritySolutionKind1 = original.ExternalSecuritySolutionKind1
type ExternalSecuritySolutionList = original.ExternalSecuritySolutionList
type ExternalSecuritySolutionListIterator = original.ExternalSecuritySolutionListIterator
type ExternalSecuritySolutionListPage = original.ExternalSecuritySolutionListPage
type ExternalSecuritySolutionModel = original.ExternalSecuritySolutionModel
type ExternalSecuritySolutionProperties = original.ExternalSecuritySolutionProperties
type ExternalSecuritySolutionsClient = original.ExternalSecuritySolutionsClient
type InformationProtectionKeyword = original.InformationProtectionKeyword
type InformationProtectionPoliciesClient = original.InformationProtectionPoliciesClient
type InformationProtectionPolicy = original.InformationProtectionPolicy
type InformationProtectionPolicyList = original.InformationProtectionPolicyList
type InformationProtectionPolicyListIterator = original.InformationProtectionPolicyListIterator
type InformationProtectionPolicyListPage = original.InformationProtectionPolicyListPage
type InformationProtectionPolicyProperties = original.InformationProtectionPolicyProperties
type InformationType = original.InformationType
type JitNetworkAccessPoliciesClient = original.JitNetworkAccessPoliciesClient
type JitNetworkAccessPoliciesList = original.JitNetworkAccessPoliciesList
type JitNetworkAccessPoliciesListIterator = original.JitNetworkAccessPoliciesListIterator
type JitNetworkAccessPoliciesListPage = original.JitNetworkAccessPoliciesListPage
type JitNetworkAccessPolicy = original.JitNetworkAccessPolicy
type JitNetworkAccessPolicyInitiatePort = original.JitNetworkAccessPolicyInitiatePort
type JitNetworkAccessPolicyInitiateRequest = original.JitNetworkAccessPolicyInitiateRequest
type JitNetworkAccessPolicyInitiateVirtualMachine = original.JitNetworkAccessPolicyInitiateVirtualMachine
type JitNetworkAccessPolicyProperties = original.JitNetworkAccessPolicyProperties
type JitNetworkAccessPolicyVirtualMachine = original.JitNetworkAccessPolicyVirtualMachine
type JitNetworkAccessPortRule = original.JitNetworkAccessPortRule
type JitNetworkAccessRequest = original.JitNetworkAccessRequest
type JitNetworkAccessRequestPort = original.JitNetworkAccessRequestPort
type JitNetworkAccessRequestVirtualMachine = original.JitNetworkAccessRequestVirtualMachine
type Kind = original.Kind
type ListCustomAlertRule = original.ListCustomAlertRule
type Location = original.Location
type LocationsClient = original.LocationsClient
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationList = original.OperationList
type OperationListIterator = original.OperationListIterator
type OperationListPage = original.OperationListPage
type OperationsClient = original.OperationsClient
type PlaybookConfiguration = original.PlaybookConfiguration
type PlaybookConfigurationAction = original.PlaybookConfigurationAction
type PlaybookConfigurationList = original.PlaybookConfigurationList
type PlaybookConfigurationListIterator = original.PlaybookConfigurationListIterator
type PlaybookConfigurationListPage = original.PlaybookConfigurationListPage
type PlaybookConfigurationMetadata = original.PlaybookConfigurationMetadata
type PlaybookConfigurationProperties = original.PlaybookConfigurationProperties
type PlaybookConfigurationRuleSet = original.PlaybookConfigurationRuleSet
type PlaybookConfigurationScope = original.PlaybookConfigurationScope
type PlaybookConfigurationSource = original.PlaybookConfigurationSource
type PlaybookConfigurationTriggeringRule = original.PlaybookConfigurationTriggeringRule
type PlaybookConfigurationValidationStatus = original.PlaybookConfigurationValidationStatus
type PlaybookConfigurationsClient = original.PlaybookConfigurationsClient
type Pricing = original.Pricing
type PricingList = original.PricingList
type PricingProperties = original.PricingProperties
type PricingsClient = original.PricingsClient
type RegulatoryComplianceAssessment = original.RegulatoryComplianceAssessment
type RegulatoryComplianceAssessmentList = original.RegulatoryComplianceAssessmentList
type RegulatoryComplianceAssessmentListIterator = original.RegulatoryComplianceAssessmentListIterator
type RegulatoryComplianceAssessmentListPage = original.RegulatoryComplianceAssessmentListPage
type RegulatoryComplianceAssessmentProperties = original.RegulatoryComplianceAssessmentProperties
type RegulatoryComplianceAssessmentsClient = original.RegulatoryComplianceAssessmentsClient
type RegulatoryComplianceControl = original.RegulatoryComplianceControl
type RegulatoryComplianceControlList = original.RegulatoryComplianceControlList
type RegulatoryComplianceControlListIterator = original.RegulatoryComplianceControlListIterator
type RegulatoryComplianceControlListPage = original.RegulatoryComplianceControlListPage
type RegulatoryComplianceControlProperties = original.RegulatoryComplianceControlProperties
type RegulatoryComplianceControlsClient = original.RegulatoryComplianceControlsClient
type RegulatoryComplianceStandard = original.RegulatoryComplianceStandard
type RegulatoryComplianceStandardList = original.RegulatoryComplianceStandardList
type RegulatoryComplianceStandardListIterator = original.RegulatoryComplianceStandardListIterator
type RegulatoryComplianceStandardListPage = original.RegulatoryComplianceStandardListPage
type RegulatoryComplianceStandardProperties = original.RegulatoryComplianceStandardProperties
type RegulatoryComplianceStandardsClient = original.RegulatoryComplianceStandardsClient
type Resource = original.Resource
type Rule = original.Rule
type SensitivityLabel = original.SensitivityLabel
type Setting = original.Setting
type SettingResource = original.SettingResource
type SettingsClient = original.SettingsClient
type SettingsList = original.SettingsList
type SettingsListIterator = original.SettingsListIterator
type SettingsListPage = original.SettingsListPage
type Tags = original.Tags
type Task = original.Task
type TaskList = original.TaskList
type TaskListIterator = original.TaskListIterator
type TaskListPage = original.TaskListPage
type TaskParameters = original.TaskParameters
type TaskProperties = original.TaskProperties
type TasksClient = original.TasksClient
type ThresholdCustomAlertRule = original.ThresholdCustomAlertRule
type TimeWindowCustomAlertRule = original.TimeWindowCustomAlertRule
type TopologyClient = original.TopologyClient
type TopologyList = original.TopologyList
type TopologyListIterator = original.TopologyListIterator
type TopologyListPage = original.TopologyListPage
type TopologyResource = original.TopologyResource
type TopologyResourceProperties = original.TopologyResourceProperties
type TopologySingleResource = original.TopologySingleResource
type TopologySingleResourceChild = original.TopologySingleResourceChild
type TopologySingleResourceParent = original.TopologySingleResourceParent
type TrackedResource = original.TrackedResource
type WorkspaceSetting = original.WorkspaceSetting
type WorkspaceSettingList = original.WorkspaceSettingList
type WorkspaceSettingListIterator = original.WorkspaceSettingListIterator
type WorkspaceSettingListPage = original.WorkspaceSettingListPage
type WorkspaceSettingProperties = original.WorkspaceSettingProperties
type WorkspaceSettingsClient = original.WorkspaceSettingsClient

func New(subscriptionID string, ascLocation string) BaseClient {
	return original.New(subscriptionID, ascLocation)
}
func NewAdaptiveNetworkHardeningsClient(subscriptionID string, ascLocation string) AdaptiveNetworkHardeningsClient {
	return original.NewAdaptiveNetworkHardeningsClient(subscriptionID, ascLocation)
}
func NewAdaptiveNetworkHardeningsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) AdaptiveNetworkHardeningsClient {
	return original.NewAdaptiveNetworkHardeningsClientWithBaseURI(baseURI, subscriptionID, ascLocation)
}
func NewAdaptiveNetworkHardeningsListIterator(page AdaptiveNetworkHardeningsListPage) AdaptiveNetworkHardeningsListIterator {
	return original.NewAdaptiveNetworkHardeningsListIterator(page)
}
func NewAdaptiveNetworkHardeningsListPage(getNextPage func(context.Context, AdaptiveNetworkHardeningsList) (AdaptiveNetworkHardeningsList, error)) AdaptiveNetworkHardeningsListPage {
	return original.NewAdaptiveNetworkHardeningsListPage(getNextPage)
}
func NewAdvancedThreatProtectionClient(subscriptionID string, ascLocation string) AdvancedThreatProtectionClient {
	return original.NewAdvancedThreatProtectionClient(subscriptionID, ascLocation)
}
func NewAdvancedThreatProtectionClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) AdvancedThreatProtectionClient {
	return original.NewAdvancedThreatProtectionClientWithBaseURI(baseURI, subscriptionID, ascLocation)
}
func NewAlertListIterator(page AlertListPage) AlertListIterator {
	return original.NewAlertListIterator(page)
}
func NewAlertListPage(getNextPage func(context.Context, AlertList) (AlertList, error)) AlertListPage {
	return original.NewAlertListPage(getNextPage)
}
func NewAlertsClient(subscriptionID string, ascLocation string) AlertsClient {
	return original.NewAlertsClient(subscriptionID, ascLocation)
}
func NewAlertsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) AlertsClient {
	return original.NewAlertsClientWithBaseURI(baseURI, subscriptionID, ascLocation)
}
func NewAllowedConnectionsClient(subscriptionID string, ascLocation string) AllowedConnectionsClient {
	return original.NewAllowedConnectionsClient(subscriptionID, ascLocation)
}
func NewAllowedConnectionsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) AllowedConnectionsClient {
	return original.NewAllowedConnectionsClientWithBaseURI(baseURI, subscriptionID, ascLocation)
}
func NewAllowedConnectionsListIterator(page AllowedConnectionsListPage) AllowedConnectionsListIterator {
	return original.NewAllowedConnectionsListIterator(page)
}
func NewAllowedConnectionsListPage(getNextPage func(context.Context, AllowedConnectionsList) (AllowedConnectionsList, error)) AllowedConnectionsListPage {
	return original.NewAllowedConnectionsListPage(getNextPage)
}
func NewAscLocationListIterator(page AscLocationListPage) AscLocationListIterator {
	return original.NewAscLocationListIterator(page)
}
func NewAscLocationListPage(getNextPage func(context.Context, AscLocationList) (AscLocationList, error)) AscLocationListPage {
	return original.NewAscLocationListPage(getNextPage)
}
func NewAutoProvisioningSettingListIterator(page AutoProvisioningSettingListPage) AutoProvisioningSettingListIterator {
	return original.NewAutoProvisioningSettingListIterator(page)
}
func NewAutoProvisioningSettingListPage(getNextPage func(context.Context, AutoProvisioningSettingList) (AutoProvisioningSettingList, error)) AutoProvisioningSettingListPage {
	return original.NewAutoProvisioningSettingListPage(getNextPage)
}
func NewAutoProvisioningSettingsClient(subscriptionID string, ascLocation string) AutoProvisioningSettingsClient {
	return original.NewAutoProvisioningSettingsClient(subscriptionID, ascLocation)
}
func NewAutoProvisioningSettingsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) AutoProvisioningSettingsClient {
	return original.NewAutoProvisioningSettingsClientWithBaseURI(baseURI, subscriptionID, ascLocation)
}
func NewComplianceListIterator(page ComplianceListPage) ComplianceListIterator {
	return original.NewComplianceListIterator(page)
}
func NewComplianceListPage(getNextPage func(context.Context, ComplianceList) (ComplianceList, error)) ComplianceListPage {
	return original.NewComplianceListPage(getNextPage)
}
func NewCompliancesClient(subscriptionID string, ascLocation string) CompliancesClient {
	return original.NewCompliancesClient(subscriptionID, ascLocation)
}
func NewCompliancesClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) CompliancesClient {
	return original.NewCompliancesClientWithBaseURI(baseURI, subscriptionID, ascLocation)
}
func NewContactListIterator(page ContactListPage) ContactListIterator {
	return original.NewContactListIterator(page)
}
func NewContactListPage(getNextPage func(context.Context, ContactList) (ContactList, error)) ContactListPage {
	return original.NewContactListPage(getNextPage)
}
func NewContactsClient(subscriptionID string, ascLocation string) ContactsClient {
	return original.NewContactsClient(subscriptionID, ascLocation)
}
func NewContactsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) ContactsClient {
	return original.NewContactsClientWithBaseURI(baseURI, subscriptionID, ascLocation)
}
func NewDeviceSecurityGroupListIterator(page DeviceSecurityGroupListPage) DeviceSecurityGroupListIterator {
	return original.NewDeviceSecurityGroupListIterator(page)
}
func NewDeviceSecurityGroupListPage(getNextPage func(context.Context, DeviceSecurityGroupList) (DeviceSecurityGroupList, error)) DeviceSecurityGroupListPage {
	return original.NewDeviceSecurityGroupListPage(getNextPage)
}
func NewDeviceSecurityGroupsClient(subscriptionID string, ascLocation string) DeviceSecurityGroupsClient {
	return original.NewDeviceSecurityGroupsClient(subscriptionID, ascLocation)
}
func NewDeviceSecurityGroupsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) DeviceSecurityGroupsClient {
	return original.NewDeviceSecurityGroupsClientWithBaseURI(baseURI, subscriptionID, ascLocation)
}
func NewDiscoveredSecuritySolutionListIterator(page DiscoveredSecuritySolutionListPage) DiscoveredSecuritySolutionListIterator {
	return original.NewDiscoveredSecuritySolutionListIterator(page)
}
func NewDiscoveredSecuritySolutionListPage(getNextPage func(context.Context, DiscoveredSecuritySolutionList) (DiscoveredSecuritySolutionList, error)) DiscoveredSecuritySolutionListPage {
	return original.NewDiscoveredSecuritySolutionListPage(getNextPage)
}
func NewDiscoveredSecuritySolutionsClient(subscriptionID string, ascLocation string) DiscoveredSecuritySolutionsClient {
	return original.NewDiscoveredSecuritySolutionsClient(subscriptionID, ascLocation)
}
func NewDiscoveredSecuritySolutionsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) DiscoveredSecuritySolutionsClient {
	return original.NewDiscoveredSecuritySolutionsClientWithBaseURI(baseURI, subscriptionID, ascLocation)
}
func NewExternalSecuritySolutionListIterator(page ExternalSecuritySolutionListPage) ExternalSecuritySolutionListIterator {
	return original.NewExternalSecuritySolutionListIterator(page)
}
func NewExternalSecuritySolutionListPage(getNextPage func(context.Context, ExternalSecuritySolutionList) (ExternalSecuritySolutionList, error)) ExternalSecuritySolutionListPage {
	return original.NewExternalSecuritySolutionListPage(getNextPage)
}
func NewExternalSecuritySolutionsClient(subscriptionID string, ascLocation string) ExternalSecuritySolutionsClient {
	return original.NewExternalSecuritySolutionsClient(subscriptionID, ascLocation)
}
func NewExternalSecuritySolutionsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) ExternalSecuritySolutionsClient {
	return original.NewExternalSecuritySolutionsClientWithBaseURI(baseURI, subscriptionID, ascLocation)
}
func NewInformationProtectionPoliciesClient(subscriptionID string, ascLocation string) InformationProtectionPoliciesClient {
	return original.NewInformationProtectionPoliciesClient(subscriptionID, ascLocation)
}
func NewInformationProtectionPoliciesClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) InformationProtectionPoliciesClient {
	return original.NewInformationProtectionPoliciesClientWithBaseURI(baseURI, subscriptionID, ascLocation)
}
func NewInformationProtectionPolicyListIterator(page InformationProtectionPolicyListPage) InformationProtectionPolicyListIterator {
	return original.NewInformationProtectionPolicyListIterator(page)
}
func NewInformationProtectionPolicyListPage(getNextPage func(context.Context, InformationProtectionPolicyList) (InformationProtectionPolicyList, error)) InformationProtectionPolicyListPage {
	return original.NewInformationProtectionPolicyListPage(getNextPage)
}
func NewJitNetworkAccessPoliciesClient(subscriptionID string, ascLocation string) JitNetworkAccessPoliciesClient {
	return original.NewJitNetworkAccessPoliciesClient(subscriptionID, ascLocation)
}
func NewJitNetworkAccessPoliciesClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) JitNetworkAccessPoliciesClient {
	return original.NewJitNetworkAccessPoliciesClientWithBaseURI(baseURI, subscriptionID, ascLocation)
}
func NewJitNetworkAccessPoliciesListIterator(page JitNetworkAccessPoliciesListPage) JitNetworkAccessPoliciesListIterator {
	return original.NewJitNetworkAccessPoliciesListIterator(page)
}
func NewJitNetworkAccessPoliciesListPage(getNextPage func(context.Context, JitNetworkAccessPoliciesList) (JitNetworkAccessPoliciesList, error)) JitNetworkAccessPoliciesListPage {
	return original.NewJitNetworkAccessPoliciesListPage(getNextPage)
}
func NewLocationsClient(subscriptionID string, ascLocation string) LocationsClient {
	return original.NewLocationsClient(subscriptionID, ascLocation)
}
func NewLocationsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) LocationsClient {
	return original.NewLocationsClientWithBaseURI(baseURI, subscriptionID, ascLocation)
}
func NewOperationListIterator(page OperationListPage) OperationListIterator {
	return original.NewOperationListIterator(page)
}
func NewOperationListPage(getNextPage func(context.Context, OperationList) (OperationList, error)) OperationListPage {
	return original.NewOperationListPage(getNextPage)
}
func NewOperationsClient(subscriptionID string, ascLocation string) OperationsClient {
	return original.NewOperationsClient(subscriptionID, ascLocation)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID, ascLocation)
}
func NewPlaybookConfigurationListIterator(page PlaybookConfigurationListPage) PlaybookConfigurationListIterator {
	return original.NewPlaybookConfigurationListIterator(page)
}
func NewPlaybookConfigurationListPage(getNextPage func(context.Context, PlaybookConfigurationList) (PlaybookConfigurationList, error)) PlaybookConfigurationListPage {
	return original.NewPlaybookConfigurationListPage(getNextPage)
}
func NewPlaybookConfigurationsClient(subscriptionID string, ascLocation string) PlaybookConfigurationsClient {
	return original.NewPlaybookConfigurationsClient(subscriptionID, ascLocation)
}
func NewPlaybookConfigurationsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) PlaybookConfigurationsClient {
	return original.NewPlaybookConfigurationsClientWithBaseURI(baseURI, subscriptionID, ascLocation)
}
func NewPricingsClient(subscriptionID string, ascLocation string) PricingsClient {
	return original.NewPricingsClient(subscriptionID, ascLocation)
}
func NewPricingsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) PricingsClient {
	return original.NewPricingsClientWithBaseURI(baseURI, subscriptionID, ascLocation)
}
func NewRegulatoryComplianceAssessmentListIterator(page RegulatoryComplianceAssessmentListPage) RegulatoryComplianceAssessmentListIterator {
	return original.NewRegulatoryComplianceAssessmentListIterator(page)
}
func NewRegulatoryComplianceAssessmentListPage(getNextPage func(context.Context, RegulatoryComplianceAssessmentList) (RegulatoryComplianceAssessmentList, error)) RegulatoryComplianceAssessmentListPage {
	return original.NewRegulatoryComplianceAssessmentListPage(getNextPage)
}
func NewRegulatoryComplianceAssessmentsClient(subscriptionID string, ascLocation string) RegulatoryComplianceAssessmentsClient {
	return original.NewRegulatoryComplianceAssessmentsClient(subscriptionID, ascLocation)
}
func NewRegulatoryComplianceAssessmentsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) RegulatoryComplianceAssessmentsClient {
	return original.NewRegulatoryComplianceAssessmentsClientWithBaseURI(baseURI, subscriptionID, ascLocation)
}
func NewRegulatoryComplianceControlListIterator(page RegulatoryComplianceControlListPage) RegulatoryComplianceControlListIterator {
	return original.NewRegulatoryComplianceControlListIterator(page)
}
func NewRegulatoryComplianceControlListPage(getNextPage func(context.Context, RegulatoryComplianceControlList) (RegulatoryComplianceControlList, error)) RegulatoryComplianceControlListPage {
	return original.NewRegulatoryComplianceControlListPage(getNextPage)
}
func NewRegulatoryComplianceControlsClient(subscriptionID string, ascLocation string) RegulatoryComplianceControlsClient {
	return original.NewRegulatoryComplianceControlsClient(subscriptionID, ascLocation)
}
func NewRegulatoryComplianceControlsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) RegulatoryComplianceControlsClient {
	return original.NewRegulatoryComplianceControlsClientWithBaseURI(baseURI, subscriptionID, ascLocation)
}
func NewRegulatoryComplianceStandardListIterator(page RegulatoryComplianceStandardListPage) RegulatoryComplianceStandardListIterator {
	return original.NewRegulatoryComplianceStandardListIterator(page)
}
func NewRegulatoryComplianceStandardListPage(getNextPage func(context.Context, RegulatoryComplianceStandardList) (RegulatoryComplianceStandardList, error)) RegulatoryComplianceStandardListPage {
	return original.NewRegulatoryComplianceStandardListPage(getNextPage)
}
func NewRegulatoryComplianceStandardsClient(subscriptionID string, ascLocation string) RegulatoryComplianceStandardsClient {
	return original.NewRegulatoryComplianceStandardsClient(subscriptionID, ascLocation)
}
func NewRegulatoryComplianceStandardsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) RegulatoryComplianceStandardsClient {
	return original.NewRegulatoryComplianceStandardsClientWithBaseURI(baseURI, subscriptionID, ascLocation)
}
func NewSettingsClient(subscriptionID string, ascLocation string) SettingsClient {
	return original.NewSettingsClient(subscriptionID, ascLocation)
}
func NewSettingsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) SettingsClient {
	return original.NewSettingsClientWithBaseURI(baseURI, subscriptionID, ascLocation)
}
func NewSettingsListIterator(page SettingsListPage) SettingsListIterator {
	return original.NewSettingsListIterator(page)
}
func NewSettingsListPage(getNextPage func(context.Context, SettingsList) (SettingsList, error)) SettingsListPage {
	return original.NewSettingsListPage(getNextPage)
}
func NewTaskListIterator(page TaskListPage) TaskListIterator {
	return original.NewTaskListIterator(page)
}
func NewTaskListPage(getNextPage func(context.Context, TaskList) (TaskList, error)) TaskListPage {
	return original.NewTaskListPage(getNextPage)
}
func NewTasksClient(subscriptionID string, ascLocation string) TasksClient {
	return original.NewTasksClient(subscriptionID, ascLocation)
}
func NewTasksClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) TasksClient {
	return original.NewTasksClientWithBaseURI(baseURI, subscriptionID, ascLocation)
}
func NewTopologyClient(subscriptionID string, ascLocation string) TopologyClient {
	return original.NewTopologyClient(subscriptionID, ascLocation)
}
func NewTopologyClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) TopologyClient {
	return original.NewTopologyClientWithBaseURI(baseURI, subscriptionID, ascLocation)
}
func NewTopologyListIterator(page TopologyListPage) TopologyListIterator {
	return original.NewTopologyListIterator(page)
}
func NewTopologyListPage(getNextPage func(context.Context, TopologyList) (TopologyList, error)) TopologyListPage {
	return original.NewTopologyListPage(getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string, ascLocation string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID, ascLocation)
}
func NewWorkspaceSettingListIterator(page WorkspaceSettingListPage) WorkspaceSettingListIterator {
	return original.NewWorkspaceSettingListIterator(page)
}
func NewWorkspaceSettingListPage(getNextPage func(context.Context, WorkspaceSettingList) (WorkspaceSettingList, error)) WorkspaceSettingListPage {
	return original.NewWorkspaceSettingListPage(getNextPage)
}
func NewWorkspaceSettingsClient(subscriptionID string, ascLocation string) WorkspaceSettingsClient {
	return original.NewWorkspaceSettingsClient(subscriptionID, ascLocation)
}
func NewWorkspaceSettingsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) WorkspaceSettingsClient {
	return original.NewWorkspaceSettingsClientWithBaseURI(baseURI, subscriptionID, ascLocation)
}
func PossibleAadConnectivityStateValues() []AadConnectivityState {
	return original.PossibleAadConnectivityStateValues()
}
func PossibleAlertNotificationsValues() []AlertNotifications {
	return original.PossibleAlertNotificationsValues()
}
func PossibleAlertsToAdminsValues() []AlertsToAdmins {
	return original.PossibleAlertsToAdminsValues()
}
func PossibleAutoProvisionValues() []AutoProvision {
	return original.PossibleAutoProvisionValues()
}
func PossibleConnectionTypeValues() []ConnectionType {
	return original.PossibleConnectionTypeValues()
}
func PossibleDirectionValues() []Direction {
	return original.PossibleDirectionValues()
}
func PossibleExternalSecuritySolutionKindValues() []ExternalSecuritySolutionKind {
	return original.PossibleExternalSecuritySolutionKindValues()
}
func PossibleFamilyValues() []Family {
	return original.PossibleFamilyValues()
}
func PossibleKindEnumValues() []KindEnum {
	return original.PossibleKindEnumValues()
}
func PossibleOperatorValues() []Operator {
	return original.PossibleOperatorValues()
}
func PossiblePricingTierValues() []PricingTier {
	return original.PossiblePricingTierValues()
}
func PossiblePropertyTypeValues() []PropertyType {
	return original.PossiblePropertyTypeValues()
}
func PossibleProtocolValues() []Protocol {
	return original.PossibleProtocolValues()
}
func PossibleReportedSeverityValues() []ReportedSeverity {
	return original.PossibleReportedSeverityValues()
}
func PossibleSettingKindValues() []SettingKind {
	return original.PossibleSettingKindValues()
}
func PossibleStateValues() []State {
	return original.PossibleStateValues()
}
func PossibleStatusReasonValues() []StatusReason {
	return original.PossibleStatusReasonValues()
}
func PossibleStatusValues() []Status {
	return original.PossibleStatusValues()
}
func PossibleTransportProtocolValues() []TransportProtocol {
	return original.PossibleTransportProtocolValues()
}
func PossibleValueTypeValues() []ValueType {
	return original.PossibleValueTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
