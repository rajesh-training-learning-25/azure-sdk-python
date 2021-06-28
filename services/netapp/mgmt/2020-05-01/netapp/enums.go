package netapp

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// BackupType enumerates the values for backup type.
type BackupType string

const (
	// Manual Manual backup
	Manual BackupType = "Manual"
	// Scheduled Scheduled backup
	Scheduled BackupType = "Scheduled"
)

// PossibleBackupTypeValues returns an array of possible values for the BackupType const type.
func PossibleBackupTypeValues() []BackupType {
	return []BackupType{Manual, Scheduled}
}

// CheckNameResourceTypes enumerates the values for check name resource types.
type CheckNameResourceTypes string

const (
	// MicrosoftNetAppnetAppAccounts ...
	MicrosoftNetAppnetAppAccounts CheckNameResourceTypes = "Microsoft.NetApp/netAppAccounts"
	// MicrosoftNetAppnetAppAccountscapacityPools ...
	MicrosoftNetAppnetAppAccountscapacityPools CheckNameResourceTypes = "Microsoft.NetApp/netAppAccounts/capacityPools"
	// MicrosoftNetAppnetAppAccountscapacityPoolsvolumes ...
	MicrosoftNetAppnetAppAccountscapacityPoolsvolumes CheckNameResourceTypes = "Microsoft.NetApp/netAppAccounts/capacityPools/volumes"
	// MicrosoftNetAppnetAppAccountscapacityPoolsvolumessnapshots ...
	MicrosoftNetAppnetAppAccountscapacityPoolsvolumessnapshots CheckNameResourceTypes = "Microsoft.NetApp/netAppAccounts/capacityPools/volumes/snapshots"
)

// PossibleCheckNameResourceTypesValues returns an array of possible values for the CheckNameResourceTypes const type.
func PossibleCheckNameResourceTypesValues() []CheckNameResourceTypes {
	return []CheckNameResourceTypes{MicrosoftNetAppnetAppAccounts, MicrosoftNetAppnetAppAccountscapacityPools, MicrosoftNetAppnetAppAccountscapacityPoolsvolumes, MicrosoftNetAppnetAppAccountscapacityPoolsvolumessnapshots}
}

// CheckQuotaNameResourceTypes enumerates the values for check quota name resource types.
type CheckQuotaNameResourceTypes string

const (
	// CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccounts ...
	CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccounts CheckQuotaNameResourceTypes = "Microsoft.NetApp/netAppAccounts"
	// CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPools ...
	CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPools CheckQuotaNameResourceTypes = "Microsoft.NetApp/netAppAccounts/capacityPools"
	// CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPoolsvolumes ...
	CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPoolsvolumes CheckQuotaNameResourceTypes = "Microsoft.NetApp/netAppAccounts/capacityPools/volumes"
	// CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPoolsvolumessnapshots ...
	CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPoolsvolumessnapshots CheckQuotaNameResourceTypes = "Microsoft.NetApp/netAppAccounts/capacityPools/volumes/snapshots"
)

// PossibleCheckQuotaNameResourceTypesValues returns an array of possible values for the CheckQuotaNameResourceTypes const type.
func PossibleCheckQuotaNameResourceTypesValues() []CheckQuotaNameResourceTypes {
	return []CheckQuotaNameResourceTypes{CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccounts, CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPools, CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPoolsvolumes, CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPoolsvolumessnapshots}
}

// EndpointType enumerates the values for endpoint type.
type EndpointType string

const (
	// Dst ...
	Dst EndpointType = "dst"
	// Src ...
	Src EndpointType = "src"
)

// PossibleEndpointTypeValues returns an array of possible values for the EndpointType const type.
func PossibleEndpointTypeValues() []EndpointType {
	return []EndpointType{Dst, Src}
}

// InAvailabilityReasonType enumerates the values for in availability reason type.
type InAvailabilityReasonType string

const (
	// AlreadyExists ...
	AlreadyExists InAvailabilityReasonType = "AlreadyExists"
	// Invalid ...
	Invalid InAvailabilityReasonType = "Invalid"
)

// PossibleInAvailabilityReasonTypeValues returns an array of possible values for the InAvailabilityReasonType const type.
func PossibleInAvailabilityReasonTypeValues() []InAvailabilityReasonType {
	return []InAvailabilityReasonType{AlreadyExists, Invalid}
}

// MirrorState enumerates the values for mirror state.
type MirrorState string

const (
	// Broken ...
	Broken MirrorState = "Broken"
	// Mirrored ...
	Mirrored MirrorState = "Mirrored"
	// Uninitialized ...
	Uninitialized MirrorState = "Uninitialized"
)

// PossibleMirrorStateValues returns an array of possible values for the MirrorState const type.
func PossibleMirrorStateValues() []MirrorState {
	return []MirrorState{Broken, Mirrored, Uninitialized}
}

// PatchServiceLevel enumerates the values for patch service level.
type PatchServiceLevel string

const (
	// Premium Premium service level
	Premium PatchServiceLevel = "Premium"
	// Standard Standard service level
	Standard PatchServiceLevel = "Standard"
	// Ultra Ultra service level
	Ultra PatchServiceLevel = "Ultra"
)

// PossiblePatchServiceLevelValues returns an array of possible values for the PatchServiceLevel const type.
func PossiblePatchServiceLevelValues() []PatchServiceLevel {
	return []PatchServiceLevel{Premium, Standard, Ultra}
}

// RelationshipStatus enumerates the values for relationship status.
type RelationshipStatus string

const (
	// Idle ...
	Idle RelationshipStatus = "Idle"
	// Transferring ...
	Transferring RelationshipStatus = "Transferring"
)

// PossibleRelationshipStatusValues returns an array of possible values for the RelationshipStatus const type.
func PossibleRelationshipStatusValues() []RelationshipStatus {
	return []RelationshipStatus{Idle, Transferring}
}

// ReplicationSchedule enumerates the values for replication schedule.
type ReplicationSchedule string

const (
	// OneZerominutely ...
	OneZerominutely ReplicationSchedule = "_10minutely"
	// Daily ...
	Daily ReplicationSchedule = "daily"
	// Hourly ...
	Hourly ReplicationSchedule = "hourly"
)

// PossibleReplicationScheduleValues returns an array of possible values for the ReplicationSchedule const type.
func PossibleReplicationScheduleValues() []ReplicationSchedule {
	return []ReplicationSchedule{OneZerominutely, Daily, Hourly}
}

// ServiceLevel enumerates the values for service level.
type ServiceLevel string

const (
	// ServiceLevelPremium Premium service level
	ServiceLevelPremium ServiceLevel = "Premium"
	// ServiceLevelStandard Standard service level
	ServiceLevelStandard ServiceLevel = "Standard"
	// ServiceLevelUltra Ultra service level
	ServiceLevelUltra ServiceLevel = "Ultra"
)

// PossibleServiceLevelValues returns an array of possible values for the ServiceLevel const type.
func PossibleServiceLevelValues() []ServiceLevel {
	return []ServiceLevel{ServiceLevelPremium, ServiceLevelStandard, ServiceLevelUltra}
}
