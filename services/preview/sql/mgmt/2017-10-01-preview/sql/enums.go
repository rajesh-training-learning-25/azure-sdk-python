package sql

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

// CapabilityGroup enumerates the values for capability group.
type CapabilityGroup string

const (
	// SupportedEditions ...
	SupportedEditions CapabilityGroup = "supportedEditions"
	// SupportedElasticPoolEditions ...
	SupportedElasticPoolEditions CapabilityGroup = "supportedElasticPoolEditions"
	// SupportedManagedInstanceVersions ...
	SupportedManagedInstanceVersions CapabilityGroup = "supportedManagedInstanceVersions"
)

// PossibleCapabilityGroupValues returns an array of possible values for the CapabilityGroup const type.
func PossibleCapabilityGroupValues() []CapabilityGroup {
	return []CapabilityGroup{SupportedEditions, SupportedElasticPoolEditions, SupportedManagedInstanceVersions}
}

// CapabilityStatus enumerates the values for capability status.
type CapabilityStatus string

const (
	// Available ...
	Available CapabilityStatus = "Available"
	// Default ...
	Default CapabilityStatus = "Default"
	// Disabled ...
	Disabled CapabilityStatus = "Disabled"
	// Visible ...
	Visible CapabilityStatus = "Visible"
)

// PossibleCapabilityStatusValues returns an array of possible values for the CapabilityStatus const type.
func PossibleCapabilityStatusValues() []CapabilityStatus {
	return []CapabilityStatus{Available, Default, Disabled, Visible}
}

// CatalogCollationType enumerates the values for catalog collation type.
type CatalogCollationType string

const (
	// DATABASEDEFAULT ...
	DATABASEDEFAULT CatalogCollationType = "DATABASE_DEFAULT"
	// SQLLatin1GeneralCP1CIAS ...
	SQLLatin1GeneralCP1CIAS CatalogCollationType = "SQL_Latin1_General_CP1_CI_AS"
)

// PossibleCatalogCollationTypeValues returns an array of possible values for the CatalogCollationType const type.
func PossibleCatalogCollationTypeValues() []CatalogCollationType {
	return []CatalogCollationType{DATABASEDEFAULT, SQLLatin1GeneralCP1CIAS}
}

// CreateMode enumerates the values for create mode.
type CreateMode string

const (
	// CreateModeCopy ...
	CreateModeCopy CreateMode = "Copy"
	// CreateModeDefault ...
	CreateModeDefault CreateMode = "Default"
	// CreateModeOnlineSecondary ...
	CreateModeOnlineSecondary CreateMode = "OnlineSecondary"
	// CreateModePointInTimeRestore ...
	CreateModePointInTimeRestore CreateMode = "PointInTimeRestore"
	// CreateModeRecovery ...
	CreateModeRecovery CreateMode = "Recovery"
	// CreateModeRestore ...
	CreateModeRestore CreateMode = "Restore"
	// CreateModeRestoreExternalBackup ...
	CreateModeRestoreExternalBackup CreateMode = "RestoreExternalBackup"
	// CreateModeRestoreExternalBackupSecondary ...
	CreateModeRestoreExternalBackupSecondary CreateMode = "RestoreExternalBackupSecondary"
	// CreateModeRestoreLongTermRetentionBackup ...
	CreateModeRestoreLongTermRetentionBackup CreateMode = "RestoreLongTermRetentionBackup"
	// CreateModeSecondary ...
	CreateModeSecondary CreateMode = "Secondary"
)

// PossibleCreateModeValues returns an array of possible values for the CreateMode const type.
func PossibleCreateModeValues() []CreateMode {
	return []CreateMode{CreateModeCopy, CreateModeDefault, CreateModeOnlineSecondary, CreateModePointInTimeRestore, CreateModeRecovery, CreateModeRestore, CreateModeRestoreExternalBackup, CreateModeRestoreExternalBackupSecondary, CreateModeRestoreLongTermRetentionBackup, CreateModeSecondary}
}

// DatabaseLicenseType enumerates the values for database license type.
type DatabaseLicenseType string

const (
	// BasePrice ...
	BasePrice DatabaseLicenseType = "BasePrice"
	// LicenseIncluded ...
	LicenseIncluded DatabaseLicenseType = "LicenseIncluded"
)

// PossibleDatabaseLicenseTypeValues returns an array of possible values for the DatabaseLicenseType const type.
func PossibleDatabaseLicenseTypeValues() []DatabaseLicenseType {
	return []DatabaseLicenseType{BasePrice, LicenseIncluded}
}

// DatabaseReadScale enumerates the values for database read scale.
type DatabaseReadScale string

const (
	// DatabaseReadScaleDisabled ...
	DatabaseReadScaleDisabled DatabaseReadScale = "Disabled"
	// DatabaseReadScaleEnabled ...
	DatabaseReadScaleEnabled DatabaseReadScale = "Enabled"
)

// PossibleDatabaseReadScaleValues returns an array of possible values for the DatabaseReadScale const type.
func PossibleDatabaseReadScaleValues() []DatabaseReadScale {
	return []DatabaseReadScale{DatabaseReadScaleDisabled, DatabaseReadScaleEnabled}
}

// DatabaseStatus enumerates the values for database status.
type DatabaseStatus string

const (
	// DatabaseStatusAutoClosed ...
	DatabaseStatusAutoClosed DatabaseStatus = "AutoClosed"
	// DatabaseStatusCopying ...
	DatabaseStatusCopying DatabaseStatus = "Copying"
	// DatabaseStatusCreating ...
	DatabaseStatusCreating DatabaseStatus = "Creating"
	// DatabaseStatusDisabled ...
	DatabaseStatusDisabled DatabaseStatus = "Disabled"
	// DatabaseStatusEmergencyMode ...
	DatabaseStatusEmergencyMode DatabaseStatus = "EmergencyMode"
	// DatabaseStatusInaccessible ...
	DatabaseStatusInaccessible DatabaseStatus = "Inaccessible"
	// DatabaseStatusOffline ...
	DatabaseStatusOffline DatabaseStatus = "Offline"
	// DatabaseStatusOfflineChangingDwPerformanceTiers ...
	DatabaseStatusOfflineChangingDwPerformanceTiers DatabaseStatus = "OfflineChangingDwPerformanceTiers"
	// DatabaseStatusOfflineSecondary ...
	DatabaseStatusOfflineSecondary DatabaseStatus = "OfflineSecondary"
	// DatabaseStatusOnline ...
	DatabaseStatusOnline DatabaseStatus = "Online"
	// DatabaseStatusOnlineChangingDwPerformanceTiers ...
	DatabaseStatusOnlineChangingDwPerformanceTiers DatabaseStatus = "OnlineChangingDwPerformanceTiers"
	// DatabaseStatusPaused ...
	DatabaseStatusPaused DatabaseStatus = "Paused"
	// DatabaseStatusPausing ...
	DatabaseStatusPausing DatabaseStatus = "Pausing"
	// DatabaseStatusRecovering ...
	DatabaseStatusRecovering DatabaseStatus = "Recovering"
	// DatabaseStatusRecoveryPending ...
	DatabaseStatusRecoveryPending DatabaseStatus = "RecoveryPending"
	// DatabaseStatusRestoring ...
	DatabaseStatusRestoring DatabaseStatus = "Restoring"
	// DatabaseStatusResuming ...
	DatabaseStatusResuming DatabaseStatus = "Resuming"
	// DatabaseStatusScaling ...
	DatabaseStatusScaling DatabaseStatus = "Scaling"
	// DatabaseStatusShutdown ...
	DatabaseStatusShutdown DatabaseStatus = "Shutdown"
	// DatabaseStatusStandby ...
	DatabaseStatusStandby DatabaseStatus = "Standby"
	// DatabaseStatusSuspect ...
	DatabaseStatusSuspect DatabaseStatus = "Suspect"
)

// PossibleDatabaseStatusValues returns an array of possible values for the DatabaseStatus const type.
func PossibleDatabaseStatusValues() []DatabaseStatus {
	return []DatabaseStatus{DatabaseStatusAutoClosed, DatabaseStatusCopying, DatabaseStatusCreating, DatabaseStatusDisabled, DatabaseStatusEmergencyMode, DatabaseStatusInaccessible, DatabaseStatusOffline, DatabaseStatusOfflineChangingDwPerformanceTiers, DatabaseStatusOfflineSecondary, DatabaseStatusOnline, DatabaseStatusOnlineChangingDwPerformanceTiers, DatabaseStatusPaused, DatabaseStatusPausing, DatabaseStatusRecovering, DatabaseStatusRecoveryPending, DatabaseStatusRestoring, DatabaseStatusResuming, DatabaseStatusScaling, DatabaseStatusShutdown, DatabaseStatusStandby, DatabaseStatusSuspect}
}

// ElasticPoolLicenseType enumerates the values for elastic pool license type.
type ElasticPoolLicenseType string

const (
	// ElasticPoolLicenseTypeBasePrice ...
	ElasticPoolLicenseTypeBasePrice ElasticPoolLicenseType = "BasePrice"
	// ElasticPoolLicenseTypeLicenseIncluded ...
	ElasticPoolLicenseTypeLicenseIncluded ElasticPoolLicenseType = "LicenseIncluded"
)

// PossibleElasticPoolLicenseTypeValues returns an array of possible values for the ElasticPoolLicenseType const type.
func PossibleElasticPoolLicenseTypeValues() []ElasticPoolLicenseType {
	return []ElasticPoolLicenseType{ElasticPoolLicenseTypeBasePrice, ElasticPoolLicenseTypeLicenseIncluded}
}

// ElasticPoolState enumerates the values for elastic pool state.
type ElasticPoolState string

const (
	// ElasticPoolStateCreating ...
	ElasticPoolStateCreating ElasticPoolState = "Creating"
	// ElasticPoolStateDisabled ...
	ElasticPoolStateDisabled ElasticPoolState = "Disabled"
	// ElasticPoolStateReady ...
	ElasticPoolStateReady ElasticPoolState = "Ready"
)

// PossibleElasticPoolStateValues returns an array of possible values for the ElasticPoolState const type.
func PossibleElasticPoolStateValues() []ElasticPoolState {
	return []ElasticPoolState{ElasticPoolStateCreating, ElasticPoolStateDisabled, ElasticPoolStateReady}
}

// InstanceFailoverGroupReplicationRole enumerates the values for instance failover group replication role.
type InstanceFailoverGroupReplicationRole string

const (
	// Primary ...
	Primary InstanceFailoverGroupReplicationRole = "Primary"
	// Secondary ...
	Secondary InstanceFailoverGroupReplicationRole = "Secondary"
)

// PossibleInstanceFailoverGroupReplicationRoleValues returns an array of possible values for the InstanceFailoverGroupReplicationRole const type.
func PossibleInstanceFailoverGroupReplicationRoleValues() []InstanceFailoverGroupReplicationRole {
	return []InstanceFailoverGroupReplicationRole{Primary, Secondary}
}

// LogSizeUnit enumerates the values for log size unit.
type LogSizeUnit string

const (
	// Gigabytes ...
	Gigabytes LogSizeUnit = "Gigabytes"
	// Megabytes ...
	Megabytes LogSizeUnit = "Megabytes"
	// Percent ...
	Percent LogSizeUnit = "Percent"
	// Petabytes ...
	Petabytes LogSizeUnit = "Petabytes"
	// Terabytes ...
	Terabytes LogSizeUnit = "Terabytes"
)

// PossibleLogSizeUnitValues returns an array of possible values for the LogSizeUnit const type.
func PossibleLogSizeUnitValues() []LogSizeUnit {
	return []LogSizeUnit{Gigabytes, Megabytes, Percent, Petabytes, Terabytes}
}

// ManagementOperationState enumerates the values for management operation state.
type ManagementOperationState string

const (
	// CancelInProgress ...
	CancelInProgress ManagementOperationState = "CancelInProgress"
	// Cancelled ...
	Cancelled ManagementOperationState = "Cancelled"
	// Failed ...
	Failed ManagementOperationState = "Failed"
	// InProgress ...
	InProgress ManagementOperationState = "InProgress"
	// Pending ...
	Pending ManagementOperationState = "Pending"
	// Succeeded ...
	Succeeded ManagementOperationState = "Succeeded"
)

// PossibleManagementOperationStateValues returns an array of possible values for the ManagementOperationState const type.
func PossibleManagementOperationStateValues() []ManagementOperationState {
	return []ManagementOperationState{CancelInProgress, Cancelled, Failed, InProgress, Pending, Succeeded}
}

// MaxSizeUnit enumerates the values for max size unit.
type MaxSizeUnit string

const (
	// MaxSizeUnitGigabytes ...
	MaxSizeUnitGigabytes MaxSizeUnit = "Gigabytes"
	// MaxSizeUnitMegabytes ...
	MaxSizeUnitMegabytes MaxSizeUnit = "Megabytes"
	// MaxSizeUnitPetabytes ...
	MaxSizeUnitPetabytes MaxSizeUnit = "Petabytes"
	// MaxSizeUnitTerabytes ...
	MaxSizeUnitTerabytes MaxSizeUnit = "Terabytes"
)

// PossibleMaxSizeUnitValues returns an array of possible values for the MaxSizeUnit const type.
func PossibleMaxSizeUnitValues() []MaxSizeUnit {
	return []MaxSizeUnit{MaxSizeUnitGigabytes, MaxSizeUnitMegabytes, MaxSizeUnitPetabytes, MaxSizeUnitTerabytes}
}

// PerformanceLevelUnit enumerates the values for performance level unit.
type PerformanceLevelUnit string

const (
	// DTU ...
	DTU PerformanceLevelUnit = "DTU"
	// VCores ...
	VCores PerformanceLevelUnit = "VCores"
)

// PossiblePerformanceLevelUnitValues returns an array of possible values for the PerformanceLevelUnit const type.
func PossiblePerformanceLevelUnitValues() []PerformanceLevelUnit {
	return []PerformanceLevelUnit{DTU, VCores}
}

// ReadOnlyEndpointFailoverPolicy enumerates the values for read only endpoint failover policy.
type ReadOnlyEndpointFailoverPolicy string

const (
	// ReadOnlyEndpointFailoverPolicyDisabled ...
	ReadOnlyEndpointFailoverPolicyDisabled ReadOnlyEndpointFailoverPolicy = "Disabled"
	// ReadOnlyEndpointFailoverPolicyEnabled ...
	ReadOnlyEndpointFailoverPolicyEnabled ReadOnlyEndpointFailoverPolicy = "Enabled"
)

// PossibleReadOnlyEndpointFailoverPolicyValues returns an array of possible values for the ReadOnlyEndpointFailoverPolicy const type.
func PossibleReadOnlyEndpointFailoverPolicyValues() []ReadOnlyEndpointFailoverPolicy {
	return []ReadOnlyEndpointFailoverPolicy{ReadOnlyEndpointFailoverPolicyDisabled, ReadOnlyEndpointFailoverPolicyEnabled}
}

// ReadWriteEndpointFailoverPolicy enumerates the values for read write endpoint failover policy.
type ReadWriteEndpointFailoverPolicy string

const (
	// Automatic ...
	Automatic ReadWriteEndpointFailoverPolicy = "Automatic"
	// Manual ...
	Manual ReadWriteEndpointFailoverPolicy = "Manual"
)

// PossibleReadWriteEndpointFailoverPolicyValues returns an array of possible values for the ReadWriteEndpointFailoverPolicy const type.
func PossibleReadWriteEndpointFailoverPolicyValues() []ReadWriteEndpointFailoverPolicy {
	return []ReadWriteEndpointFailoverPolicy{Automatic, Manual}
}

// SampleName enumerates the values for sample name.
type SampleName string

const (
	// AdventureWorksLT ...
	AdventureWorksLT SampleName = "AdventureWorksLT"
	// WideWorldImportersFull ...
	WideWorldImportersFull SampleName = "WideWorldImportersFull"
	// WideWorldImportersStd ...
	WideWorldImportersStd SampleName = "WideWorldImportersStd"
)

// PossibleSampleNameValues returns an array of possible values for the SampleName const type.
func PossibleSampleNameValues() []SampleName {
	return []SampleName{AdventureWorksLT, WideWorldImportersFull, WideWorldImportersStd}
}

// ServerKeyType enumerates the values for server key type.
type ServerKeyType string

const (
	// AzureKeyVault ...
	AzureKeyVault ServerKeyType = "AzureKeyVault"
	// ServiceManaged ...
	ServiceManaged ServerKeyType = "ServiceManaged"
)

// PossibleServerKeyTypeValues returns an array of possible values for the ServerKeyType const type.
func PossibleServerKeyTypeValues() []ServerKeyType {
	return []ServerKeyType{AzureKeyVault, ServiceManaged}
}

// VulnerabilityAssessmentPolicyBaselineName enumerates the values for vulnerability assessment policy baseline
// name.
type VulnerabilityAssessmentPolicyBaselineName string

const (
	// VulnerabilityAssessmentPolicyBaselineNameDefault ...
	VulnerabilityAssessmentPolicyBaselineNameDefault VulnerabilityAssessmentPolicyBaselineName = "default"
	// VulnerabilityAssessmentPolicyBaselineNameMaster ...
	VulnerabilityAssessmentPolicyBaselineNameMaster VulnerabilityAssessmentPolicyBaselineName = "master"
)

// PossibleVulnerabilityAssessmentPolicyBaselineNameValues returns an array of possible values for the VulnerabilityAssessmentPolicyBaselineName const type.
func PossibleVulnerabilityAssessmentPolicyBaselineNameValues() []VulnerabilityAssessmentPolicyBaselineName {
	return []VulnerabilityAssessmentPolicyBaselineName{VulnerabilityAssessmentPolicyBaselineNameDefault, VulnerabilityAssessmentPolicyBaselineNameMaster}
}

// VulnerabilityAssessmentScanState enumerates the values for vulnerability assessment scan state.
type VulnerabilityAssessmentScanState string

const (
	// VulnerabilityAssessmentScanStateFailed ...
	VulnerabilityAssessmentScanStateFailed VulnerabilityAssessmentScanState = "Failed"
	// VulnerabilityAssessmentScanStateFailedToRun ...
	VulnerabilityAssessmentScanStateFailedToRun VulnerabilityAssessmentScanState = "FailedToRun"
	// VulnerabilityAssessmentScanStateInProgress ...
	VulnerabilityAssessmentScanStateInProgress VulnerabilityAssessmentScanState = "InProgress"
	// VulnerabilityAssessmentScanStatePassed ...
	VulnerabilityAssessmentScanStatePassed VulnerabilityAssessmentScanState = "Passed"
)

// PossibleVulnerabilityAssessmentScanStateValues returns an array of possible values for the VulnerabilityAssessmentScanState const type.
func PossibleVulnerabilityAssessmentScanStateValues() []VulnerabilityAssessmentScanState {
	return []VulnerabilityAssessmentScanState{VulnerabilityAssessmentScanStateFailed, VulnerabilityAssessmentScanStateFailedToRun, VulnerabilityAssessmentScanStateInProgress, VulnerabilityAssessmentScanStatePassed}
}

// VulnerabilityAssessmentScanTriggerType enumerates the values for vulnerability assessment scan trigger type.
type VulnerabilityAssessmentScanTriggerType string

const (
	// OnDemand ...
	OnDemand VulnerabilityAssessmentScanTriggerType = "OnDemand"
	// Recurring ...
	Recurring VulnerabilityAssessmentScanTriggerType = "Recurring"
)

// PossibleVulnerabilityAssessmentScanTriggerTypeValues returns an array of possible values for the VulnerabilityAssessmentScanTriggerType const type.
func PossibleVulnerabilityAssessmentScanTriggerTypeValues() []VulnerabilityAssessmentScanTriggerType {
	return []VulnerabilityAssessmentScanTriggerType{OnDemand, Recurring}
}
