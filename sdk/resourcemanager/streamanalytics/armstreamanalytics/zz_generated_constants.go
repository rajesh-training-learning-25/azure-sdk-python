//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstreamanalytics

const (
	moduleName    = "armstreamanalytics"
	moduleVersion = "v0.4.0"
)

// AuthenticationMode - Authentication Mode. Valid modes are ConnectionString, Msi and 'UserToken'.
type AuthenticationMode string

const (
	AuthenticationModeConnectionString AuthenticationMode = "ConnectionString"
	AuthenticationModeMsi              AuthenticationMode = "Msi"
	AuthenticationModeUserToken        AuthenticationMode = "UserToken"
)

// PossibleAuthenticationModeValues returns the possible values for the AuthenticationMode const type.
func PossibleAuthenticationModeValues() []AuthenticationMode {
	return []AuthenticationMode{
		AuthenticationModeConnectionString,
		AuthenticationModeMsi,
		AuthenticationModeUserToken,
	}
}

// ClusterProvisioningState - The status of the cluster provisioning. The three terminal states are: Succeeded, Failed and
// Canceled
type ClusterProvisioningState string

const (
	// ClusterProvisioningStateCanceled - The cluster provisioning was canceled.
	ClusterProvisioningStateCanceled ClusterProvisioningState = "Canceled"
	// ClusterProvisioningStateFailed - The cluster provisioning failed.
	ClusterProvisioningStateFailed ClusterProvisioningState = "Failed"
	// ClusterProvisioningStateInProgress - The cluster provisioning was inprogress.
	ClusterProvisioningStateInProgress ClusterProvisioningState = "InProgress"
	// ClusterProvisioningStateSucceeded - The cluster provisioning succeeded.
	ClusterProvisioningStateSucceeded ClusterProvisioningState = "Succeeded"
)

// PossibleClusterProvisioningStateValues returns the possible values for the ClusterProvisioningState const type.
func PossibleClusterProvisioningStateValues() []ClusterProvisioningState {
	return []ClusterProvisioningState{
		ClusterProvisioningStateCanceled,
		ClusterProvisioningStateFailed,
		ClusterProvisioningStateInProgress,
		ClusterProvisioningStateSucceeded,
	}
}

// ClusterSKUName - Specifies the SKU name of the cluster. Required on PUT (CreateOrUpdate) requests.
type ClusterSKUName string

const (
	// ClusterSKUNameDefault - The default SKU.
	ClusterSKUNameDefault ClusterSKUName = "Default"
)

// PossibleClusterSKUNameValues returns the possible values for the ClusterSKUName const type.
func PossibleClusterSKUNameValues() []ClusterSKUName {
	return []ClusterSKUName{
		ClusterSKUNameDefault,
	}
}

// CompatibilityLevel - Controls certain runtime behaviors of the streaming job.
type CompatibilityLevel string

const (
	CompatibilityLevelOne0 CompatibilityLevel = "1.0"
	CompatibilityLevelOne2 CompatibilityLevel = "1.2"
)

// PossibleCompatibilityLevelValues returns the possible values for the CompatibilityLevel const type.
func PossibleCompatibilityLevelValues() []CompatibilityLevel {
	return []CompatibilityLevel{
		CompatibilityLevelOne0,
		CompatibilityLevelOne2,
	}
}

// CompressionType - Indicates the type of compression that the input uses. Required on PUT (CreateOrReplace) requests.
type CompressionType string

const (
	CompressionTypeDeflate CompressionType = "Deflate"
	CompressionTypeGZip    CompressionType = "GZip"
	CompressionTypeNone    CompressionType = "None"
)

// PossibleCompressionTypeValues returns the possible values for the CompressionType const type.
func PossibleCompressionTypeValues() []CompressionType {
	return []CompressionType{
		CompressionTypeDeflate,
		CompressionTypeGZip,
		CompressionTypeNone,
	}
}

// ContentStoragePolicy - Valid values are JobStorageAccount and SystemAccount. If set to JobStorageAccount, this requires
// the user to also specify jobStorageAccount property. .
type ContentStoragePolicy string

const (
	ContentStoragePolicyJobStorageAccount ContentStoragePolicy = "JobStorageAccount"
	ContentStoragePolicySystemAccount     ContentStoragePolicy = "SystemAccount"
)

// PossibleContentStoragePolicyValues returns the possible values for the ContentStoragePolicy const type.
func PossibleContentStoragePolicyValues() []ContentStoragePolicy {
	return []ContentStoragePolicy{
		ContentStoragePolicyJobStorageAccount,
		ContentStoragePolicySystemAccount,
	}
}

// Encoding - Specifies the encoding of the incoming data in the case of input and the encoding of outgoing data in the case
// of output.
type Encoding string

const (
	EncodingUTF8 Encoding = "UTF8"
)

// PossibleEncodingValues returns the possible values for the Encoding const type.
func PossibleEncodingValues() []Encoding {
	return []Encoding{
		EncodingUTF8,
	}
}

// EventSerializationType - Indicates the type of serialization that the input or output uses. Required on PUT (CreateOrReplace)
// requests.
type EventSerializationType string

const (
	EventSerializationTypeAvro    EventSerializationType = "Avro"
	EventSerializationTypeCSV     EventSerializationType = "Csv"
	EventSerializationTypeJSON    EventSerializationType = "Json"
	EventSerializationTypeParquet EventSerializationType = "Parquet"
)

// PossibleEventSerializationTypeValues returns the possible values for the EventSerializationType const type.
func PossibleEventSerializationTypeValues() []EventSerializationType {
	return []EventSerializationType{
		EventSerializationTypeAvro,
		EventSerializationTypeCSV,
		EventSerializationTypeJSON,
		EventSerializationTypeParquet,
	}
}

// EventsOutOfOrderPolicy - Indicates the policy to apply to events that arrive out of order in the input event stream.
type EventsOutOfOrderPolicy string

const (
	EventsOutOfOrderPolicyAdjust EventsOutOfOrderPolicy = "Adjust"
	EventsOutOfOrderPolicyDrop   EventsOutOfOrderPolicy = "Drop"
)

// PossibleEventsOutOfOrderPolicyValues returns the possible values for the EventsOutOfOrderPolicy const type.
func PossibleEventsOutOfOrderPolicyValues() []EventsOutOfOrderPolicy {
	return []EventsOutOfOrderPolicy{
		EventsOutOfOrderPolicyAdjust,
		EventsOutOfOrderPolicyDrop,
	}
}

// JSONOutputSerializationFormat - Specifies the format of the JSON the output will be written in. The currently supported
// values are 'lineSeparated' indicating the output will be formatted by having each JSON object separated by a new
// line and 'array' indicating the output will be formatted as an array of JSON objects.
type JSONOutputSerializationFormat string

const (
	JSONOutputSerializationFormatArray         JSONOutputSerializationFormat = "Array"
	JSONOutputSerializationFormatLineSeparated JSONOutputSerializationFormat = "LineSeparated"
)

// PossibleJSONOutputSerializationFormatValues returns the possible values for the JSONOutputSerializationFormat const type.
func PossibleJSONOutputSerializationFormatValues() []JSONOutputSerializationFormat {
	return []JSONOutputSerializationFormat{
		JSONOutputSerializationFormatArray,
		JSONOutputSerializationFormatLineSeparated,
	}
}

// JobState - The current execution state of the streaming job.
type JobState string

const (
	// JobStateCreated - The job is currently in the Created state.
	JobStateCreated JobState = "Created"
	// JobStateDegraded - The job is currently in the Degraded state.
	JobStateDegraded JobState = "Degraded"
	// JobStateDeleting - The job is currently in the Deleting state.
	JobStateDeleting JobState = "Deleting"
	// JobStateFailed - The job is currently in the Failed state.
	JobStateFailed JobState = "Failed"
	// JobStateRestarting - The job is currently in the Restarting state.
	JobStateRestarting JobState = "Restarting"
	// JobStateRunning - The job is currently in the Running state.
	JobStateRunning JobState = "Running"
	// JobStateScaling - The job is currently in the Scaling state.
	JobStateScaling JobState = "Scaling"
	// JobStateStarting - The job is currently in the Starting state.
	JobStateStarting JobState = "Starting"
	// JobStateStopped - The job is currently in the Stopped state.
	JobStateStopped JobState = "Stopped"
	// JobStateStopping - The job is currently in the Stopping state.
	JobStateStopping JobState = "Stopping"
)

// PossibleJobStateValues returns the possible values for the JobState const type.
func PossibleJobStateValues() []JobState {
	return []JobState{
		JobStateCreated,
		JobStateDegraded,
		JobStateDeleting,
		JobStateFailed,
		JobStateRestarting,
		JobStateRunning,
		JobStateScaling,
		JobStateStarting,
		JobStateStopped,
		JobStateStopping,
	}
}

// JobType - Describes the type of the job. Valid modes are Cloud and 'Edge'.
type JobType string

const (
	JobTypeCloud JobType = "Cloud"
	JobTypeEdge  JobType = "Edge"
)

// PossibleJobTypeValues returns the possible values for the JobType const type.
func PossibleJobTypeValues() []JobType {
	return []JobType{
		JobTypeCloud,
		JobTypeEdge,
	}
}

// OutputErrorPolicy - Indicates the policy to apply to events that arrive at the output and cannot be written to the external
// storage due to being malformed (missing column values, column values of wrong type or size).
type OutputErrorPolicy string

const (
	OutputErrorPolicyDrop OutputErrorPolicy = "Drop"
	OutputErrorPolicyStop OutputErrorPolicy = "Stop"
)

// PossibleOutputErrorPolicyValues returns the possible values for the OutputErrorPolicy const type.
func PossibleOutputErrorPolicyValues() []OutputErrorPolicy {
	return []OutputErrorPolicy{
		OutputErrorPolicyDrop,
		OutputErrorPolicyStop,
	}
}

// OutputStartMode - Value may be JobStartTime, CustomTime, or LastOutputEventTime to indicate whether the starting point
// of the output event stream should start whenever the job is started, start at a custom user time
// stamp specified via the outputStartTime property, or start from the last event output time.
type OutputStartMode string

const (
	OutputStartModeCustomTime          OutputStartMode = "CustomTime"
	OutputStartModeJobStartTime        OutputStartMode = "JobStartTime"
	OutputStartModeLastOutputEventTime OutputStartMode = "LastOutputEventTime"
)

// PossibleOutputStartModeValues returns the possible values for the OutputStartMode const type.
func PossibleOutputStartModeValues() []OutputStartMode {
	return []OutputStartMode{
		OutputStartModeCustomTime,
		OutputStartModeJobStartTime,
		OutputStartModeLastOutputEventTime,
	}
}

// RefreshType - Indicates the type of data refresh option.
type RefreshType string

const (
	RefreshTypeRefreshPeriodicallyWithDelta RefreshType = "RefreshPeriodicallyWithDelta"
	RefreshTypeRefreshPeriodicallyWithFull  RefreshType = "RefreshPeriodicallyWithFull"
	RefreshTypeStatic                       RefreshType = "Static"
)

// PossibleRefreshTypeValues returns the possible values for the RefreshType const type.
func PossibleRefreshTypeValues() []RefreshType {
	return []RefreshType{
		RefreshTypeRefreshPeriodicallyWithDelta,
		RefreshTypeRefreshPeriodicallyWithFull,
		RefreshTypeStatic,
	}
}

// SKUName - The name of the SKU. Required on PUT (CreateOrReplace) requests.
type SKUName string

const (
	SKUNameStandard SKUName = "Standard"
)

// PossibleSKUNameValues returns the possible values for the SKUName const type.
func PossibleSKUNameValues() []SKUName {
	return []SKUName{
		SKUNameStandard,
	}
}
