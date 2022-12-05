//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armiothub

const (
	moduleName    = "armiothub"
	moduleVersion = "v1.0.0"
)

// AccessRights - The permissions assigned to the shared access policy.
type AccessRights string

const (
	AccessRightsRegistryRead                                         AccessRights = "RegistryRead"
	AccessRightsRegistryWrite                                        AccessRights = "RegistryWrite"
	AccessRightsServiceConnect                                       AccessRights = "ServiceConnect"
	AccessRightsDeviceConnect                                        AccessRights = "DeviceConnect"
	AccessRightsRegistryReadRegistryWrite                            AccessRights = "RegistryRead, RegistryWrite"
	AccessRightsRegistryReadServiceConnect                           AccessRights = "RegistryRead, ServiceConnect"
	AccessRightsRegistryReadDeviceConnect                            AccessRights = "RegistryRead, DeviceConnect"
	AccessRightsRegistryWriteServiceConnect                          AccessRights = "RegistryWrite, ServiceConnect"
	AccessRightsRegistryWriteDeviceConnect                           AccessRights = "RegistryWrite, DeviceConnect"
	AccessRightsServiceConnectDeviceConnect                          AccessRights = "ServiceConnect, DeviceConnect"
	AccessRightsRegistryReadRegistryWriteServiceConnect              AccessRights = "RegistryRead, RegistryWrite, ServiceConnect"
	AccessRightsRegistryReadRegistryWriteDeviceConnect               AccessRights = "RegistryRead, RegistryWrite, DeviceConnect"
	AccessRightsRegistryReadServiceConnectDeviceConnect              AccessRights = "RegistryRead, ServiceConnect, DeviceConnect"
	AccessRightsRegistryWriteServiceConnectDeviceConnect             AccessRights = "RegistryWrite, ServiceConnect, DeviceConnect"
	AccessRightsRegistryReadRegistryWriteServiceConnectDeviceConnect AccessRights = "RegistryRead, RegistryWrite, ServiceConnect, DeviceConnect"
)

// PossibleAccessRightsValues returns the possible values for the AccessRights const type.
func PossibleAccessRightsValues() []AccessRights {
	return []AccessRights{
		AccessRightsRegistryRead,
		AccessRightsRegistryWrite,
		AccessRightsServiceConnect,
		AccessRightsDeviceConnect,
		AccessRightsRegistryReadRegistryWrite,
		AccessRightsRegistryReadServiceConnect,
		AccessRightsRegistryReadDeviceConnect,
		AccessRightsRegistryWriteServiceConnect,
		AccessRightsRegistryWriteDeviceConnect,
		AccessRightsServiceConnectDeviceConnect,
		AccessRightsRegistryReadRegistryWriteServiceConnect,
		AccessRightsRegistryReadRegistryWriteDeviceConnect,
		AccessRightsRegistryReadServiceConnectDeviceConnect,
		AccessRightsRegistryWriteServiceConnectDeviceConnect,
		AccessRightsRegistryReadRegistryWriteServiceConnectDeviceConnect,
	}
}

// AuthenticationType - Specifies authentication type being used for connecting to the storage account.
type AuthenticationType string

const (
	AuthenticationTypeIdentityBased AuthenticationType = "identityBased"
	AuthenticationTypeKeyBased      AuthenticationType = "keyBased"
)

// PossibleAuthenticationTypeValues returns the possible values for the AuthenticationType const type.
func PossibleAuthenticationTypeValues() []AuthenticationType {
	return []AuthenticationType{
		AuthenticationTypeIdentityBased,
		AuthenticationTypeKeyBased,
	}
}

// Capabilities - The capabilities and features enabled for the IoT hub.
type Capabilities string

const (
	CapabilitiesDeviceManagement Capabilities = "DeviceManagement"
	CapabilitiesNone             Capabilities = "None"
)

// PossibleCapabilitiesValues returns the possible values for the Capabilities const type.
func PossibleCapabilitiesValues() []Capabilities {
	return []Capabilities{
		CapabilitiesDeviceManagement,
		CapabilitiesNone,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// DefaultAction - Default Action for Network Rule Set
type DefaultAction string

const (
	DefaultActionAllow DefaultAction = "Allow"
	DefaultActionDeny  DefaultAction = "Deny"
)

// PossibleDefaultActionValues returns the possible values for the DefaultAction const type.
func PossibleDefaultActionValues() []DefaultAction {
	return []DefaultAction{
		DefaultActionAllow,
		DefaultActionDeny,
	}
}

// EndpointHealthStatus - Health statuses have following meanings. The 'healthy' status shows that the endpoint is accepting
// messages as expected. The 'unhealthy' status shows that the endpoint is not accepting messages as
// expected and IoT Hub is retrying to send data to this endpoint. The status of an unhealthy endpoint will be updated to
// healthy when IoT Hub has established an eventually consistent state of health.
// The 'dead' status shows that the endpoint is not accepting messages, after IoT Hub retried sending messages for the retrial
// period. See IoT Hub metrics to identify errors and monitor issues with
// endpoints. The 'unknown' status shows that the IoT Hub has not established a connection with the endpoint. No messages
// have been delivered to or rejected from this endpoint
type EndpointHealthStatus string

const (
	EndpointHealthStatusDead      EndpointHealthStatus = "dead"
	EndpointHealthStatusDegraded  EndpointHealthStatus = "degraded"
	EndpointHealthStatusHealthy   EndpointHealthStatus = "healthy"
	EndpointHealthStatusUnhealthy EndpointHealthStatus = "unhealthy"
	EndpointHealthStatusUnknown   EndpointHealthStatus = "unknown"
)

// PossibleEndpointHealthStatusValues returns the possible values for the EndpointHealthStatus const type.
func PossibleEndpointHealthStatusValues() []EndpointHealthStatus {
	return []EndpointHealthStatus{
		EndpointHealthStatusDead,
		EndpointHealthStatusDegraded,
		EndpointHealthStatusHealthy,
		EndpointHealthStatusUnhealthy,
		EndpointHealthStatusUnknown,
	}
}

// IPFilterActionType - The desired action for requests captured by this rule.
type IPFilterActionType string

const (
	IPFilterActionTypeAccept IPFilterActionType = "Accept"
	IPFilterActionTypeReject IPFilterActionType = "Reject"
)

// PossibleIPFilterActionTypeValues returns the possible values for the IPFilterActionType const type.
func PossibleIPFilterActionTypeValues() []IPFilterActionType {
	return []IPFilterActionType{
		IPFilterActionTypeAccept,
		IPFilterActionTypeReject,
	}
}

// IotHubNameUnavailabilityReason - The reason for unavailability.
type IotHubNameUnavailabilityReason string

const (
	IotHubNameUnavailabilityReasonInvalid       IotHubNameUnavailabilityReason = "Invalid"
	IotHubNameUnavailabilityReasonAlreadyExists IotHubNameUnavailabilityReason = "AlreadyExists"
)

// PossibleIotHubNameUnavailabilityReasonValues returns the possible values for the IotHubNameUnavailabilityReason const type.
func PossibleIotHubNameUnavailabilityReasonValues() []IotHubNameUnavailabilityReason {
	return []IotHubNameUnavailabilityReason{
		IotHubNameUnavailabilityReasonInvalid,
		IotHubNameUnavailabilityReasonAlreadyExists,
	}
}

// IotHubReplicaRoleType - The role of the region, can be either primary or secondary. The primary region is where the IoT
// hub is currently provisioned. The secondary region is the Azure disaster recovery (DR) paired region and
// also the region where the IoT hub can failover to.
type IotHubReplicaRoleType string

const (
	IotHubReplicaRoleTypePrimary   IotHubReplicaRoleType = "primary"
	IotHubReplicaRoleTypeSecondary IotHubReplicaRoleType = "secondary"
)

// PossibleIotHubReplicaRoleTypeValues returns the possible values for the IotHubReplicaRoleType const type.
func PossibleIotHubReplicaRoleTypeValues() []IotHubReplicaRoleType {
	return []IotHubReplicaRoleType{
		IotHubReplicaRoleTypePrimary,
		IotHubReplicaRoleTypeSecondary,
	}
}

// IotHubSKU - The name of the SKU.
type IotHubSKU string

const (
	IotHubSKUB1 IotHubSKU = "B1"
	IotHubSKUB2 IotHubSKU = "B2"
	IotHubSKUB3 IotHubSKU = "B3"
	IotHubSKUF1 IotHubSKU = "F1"
	IotHubSKUS1 IotHubSKU = "S1"
	IotHubSKUS2 IotHubSKU = "S2"
	IotHubSKUS3 IotHubSKU = "S3"
)

// PossibleIotHubSKUValues returns the possible values for the IotHubSKU const type.
func PossibleIotHubSKUValues() []IotHubSKU {
	return []IotHubSKU{
		IotHubSKUB1,
		IotHubSKUB2,
		IotHubSKUB3,
		IotHubSKUF1,
		IotHubSKUS1,
		IotHubSKUS2,
		IotHubSKUS3,
	}
}

// IotHubSKUTier - The billing tier for the IoT hub.
type IotHubSKUTier string

const (
	IotHubSKUTierFree     IotHubSKUTier = "Free"
	IotHubSKUTierStandard IotHubSKUTier = "Standard"
	IotHubSKUTierBasic    IotHubSKUTier = "Basic"
)

// PossibleIotHubSKUTierValues returns the possible values for the IotHubSKUTier const type.
func PossibleIotHubSKUTierValues() []IotHubSKUTier {
	return []IotHubSKUTier{
		IotHubSKUTierFree,
		IotHubSKUTierStandard,
		IotHubSKUTierBasic,
	}
}

// IotHubScaleType - The type of the scaling enabled.
type IotHubScaleType string

const (
	IotHubScaleTypeAutomatic IotHubScaleType = "Automatic"
	IotHubScaleTypeManual    IotHubScaleType = "Manual"
	IotHubScaleTypeNone      IotHubScaleType = "None"
)

// PossibleIotHubScaleTypeValues returns the possible values for the IotHubScaleType const type.
func PossibleIotHubScaleTypeValues() []IotHubScaleType {
	return []IotHubScaleType{
		IotHubScaleTypeAutomatic,
		IotHubScaleTypeManual,
		IotHubScaleTypeNone,
	}
}

// JobStatus - The status of the job.
type JobStatus string

const (
	JobStatusUnknown   JobStatus = "unknown"
	JobStatusEnqueued  JobStatus = "enqueued"
	JobStatusRunning   JobStatus = "running"
	JobStatusCompleted JobStatus = "completed"
	JobStatusFailed    JobStatus = "failed"
	JobStatusCancelled JobStatus = "cancelled"
)

// PossibleJobStatusValues returns the possible values for the JobStatus const type.
func PossibleJobStatusValues() []JobStatus {
	return []JobStatus{
		JobStatusUnknown,
		JobStatusEnqueued,
		JobStatusRunning,
		JobStatusCompleted,
		JobStatusFailed,
		JobStatusCancelled,
	}
}

// JobType - The type of the job.
type JobType string

const (
	JobTypeBackup                    JobType = "backup"
	JobTypeExport                    JobType = "export"
	JobTypeFactoryResetDevice        JobType = "factoryResetDevice"
	JobTypeFirmwareUpdate            JobType = "firmwareUpdate"
	JobTypeImport                    JobType = "import"
	JobTypeReadDeviceProperties      JobType = "readDeviceProperties"
	JobTypeRebootDevice              JobType = "rebootDevice"
	JobTypeUnknown                   JobType = "unknown"
	JobTypeUpdateDeviceConfiguration JobType = "updateDeviceConfiguration"
	JobTypeWriteDeviceProperties     JobType = "writeDeviceProperties"
)

// PossibleJobTypeValues returns the possible values for the JobType const type.
func PossibleJobTypeValues() []JobType {
	return []JobType{
		JobTypeBackup,
		JobTypeExport,
		JobTypeFactoryResetDevice,
		JobTypeFirmwareUpdate,
		JobTypeImport,
		JobTypeReadDeviceProperties,
		JobTypeRebootDevice,
		JobTypeUnknown,
		JobTypeUpdateDeviceConfiguration,
		JobTypeWriteDeviceProperties,
	}
}

// NetworkRuleIPAction - IP Filter Action
type NetworkRuleIPAction string

const (
	NetworkRuleIPActionAllow NetworkRuleIPAction = "Allow"
)

// PossibleNetworkRuleIPActionValues returns the possible values for the NetworkRuleIPAction const type.
func PossibleNetworkRuleIPActionValues() []NetworkRuleIPAction {
	return []NetworkRuleIPAction{
		NetworkRuleIPActionAllow,
	}
}

// PrivateLinkServiceConnectionStatus - The status of a private endpoint connection
type PrivateLinkServiceConnectionStatus string

const (
	PrivateLinkServiceConnectionStatusApproved     PrivateLinkServiceConnectionStatus = "Approved"
	PrivateLinkServiceConnectionStatusDisconnected PrivateLinkServiceConnectionStatus = "Disconnected"
	PrivateLinkServiceConnectionStatusPending      PrivateLinkServiceConnectionStatus = "Pending"
	PrivateLinkServiceConnectionStatusRejected     PrivateLinkServiceConnectionStatus = "Rejected"
)

// PossiblePrivateLinkServiceConnectionStatusValues returns the possible values for the PrivateLinkServiceConnectionStatus const type.
func PossiblePrivateLinkServiceConnectionStatusValues() []PrivateLinkServiceConnectionStatus {
	return []PrivateLinkServiceConnectionStatus{
		PrivateLinkServiceConnectionStatusApproved,
		PrivateLinkServiceConnectionStatusDisconnected,
		PrivateLinkServiceConnectionStatusPending,
		PrivateLinkServiceConnectionStatusRejected,
	}
}

// PublicNetworkAccess - Whether requests from Public Network are allowed
type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

// PossiblePublicNetworkAccessValues returns the possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{
		PublicNetworkAccessDisabled,
		PublicNetworkAccessEnabled,
	}
}

// ResourceIdentityType - The type of identity used for the resource. The type 'SystemAssigned, UserAssigned' includes both
// an implicitly created identity and a set of user assigned identities. The type 'None' will remove any
// identities from the service.
type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned             ResourceIdentityType = "SystemAssigned"
	ResourceIdentityTypeUserAssigned               ResourceIdentityType = "UserAssigned"
	ResourceIdentityTypeSystemAssignedUserAssigned ResourceIdentityType = "SystemAssigned, UserAssigned"
	ResourceIdentityTypeNone                       ResourceIdentityType = "None"
)

// PossibleResourceIdentityTypeValues returns the possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{
		ResourceIdentityTypeSystemAssigned,
		ResourceIdentityTypeUserAssigned,
		ResourceIdentityTypeSystemAssignedUserAssigned,
		ResourceIdentityTypeNone,
	}
}

// RouteErrorSeverity - Severity of the route error
type RouteErrorSeverity string

const (
	RouteErrorSeverityError   RouteErrorSeverity = "error"
	RouteErrorSeverityWarning RouteErrorSeverity = "warning"
)

// PossibleRouteErrorSeverityValues returns the possible values for the RouteErrorSeverity const type.
func PossibleRouteErrorSeverityValues() []RouteErrorSeverity {
	return []RouteErrorSeverity{
		RouteErrorSeverityError,
		RouteErrorSeverityWarning,
	}
}

// RoutingSource - The source that the routing rule is to be applied to, such as DeviceMessages.
type RoutingSource string

const (
	RoutingSourceDeviceConnectionStateEvents RoutingSource = "DeviceConnectionStateEvents"
	RoutingSourceDeviceJobLifecycleEvents    RoutingSource = "DeviceJobLifecycleEvents"
	RoutingSourceDeviceLifecycleEvents       RoutingSource = "DeviceLifecycleEvents"
	RoutingSourceDeviceMessages              RoutingSource = "DeviceMessages"
	RoutingSourceInvalid                     RoutingSource = "Invalid"
	RoutingSourceTwinChangeEvents            RoutingSource = "TwinChangeEvents"
)

// PossibleRoutingSourceValues returns the possible values for the RoutingSource const type.
func PossibleRoutingSourceValues() []RoutingSource {
	return []RoutingSource{
		RoutingSourceDeviceConnectionStateEvents,
		RoutingSourceDeviceJobLifecycleEvents,
		RoutingSourceDeviceLifecycleEvents,
		RoutingSourceDeviceMessages,
		RoutingSourceInvalid,
		RoutingSourceTwinChangeEvents,
	}
}

// RoutingStorageContainerPropertiesEncoding - Encoding that is used to serialize messages to blobs. Supported values are
// 'avro', 'avrodeflate', and 'JSON'. Default value is 'avro'.
type RoutingStorageContainerPropertiesEncoding string

const (
	RoutingStorageContainerPropertiesEncodingAvro        RoutingStorageContainerPropertiesEncoding = "Avro"
	RoutingStorageContainerPropertiesEncodingAvroDeflate RoutingStorageContainerPropertiesEncoding = "AvroDeflate"
	RoutingStorageContainerPropertiesEncodingJSON        RoutingStorageContainerPropertiesEncoding = "JSON"
)

// PossibleRoutingStorageContainerPropertiesEncodingValues returns the possible values for the RoutingStorageContainerPropertiesEncoding const type.
func PossibleRoutingStorageContainerPropertiesEncodingValues() []RoutingStorageContainerPropertiesEncoding {
	return []RoutingStorageContainerPropertiesEncoding{
		RoutingStorageContainerPropertiesEncodingAvro,
		RoutingStorageContainerPropertiesEncodingAvroDeflate,
		RoutingStorageContainerPropertiesEncodingJSON,
	}
}

// TestResultStatus - Result of testing route
type TestResultStatus string

const (
	TestResultStatusFalse     TestResultStatus = "false"
	TestResultStatusTrue      TestResultStatus = "true"
	TestResultStatusUndefined TestResultStatus = "undefined"
)

// PossibleTestResultStatusValues returns the possible values for the TestResultStatus const type.
func PossibleTestResultStatusValues() []TestResultStatus {
	return []TestResultStatus{
		TestResultStatusFalse,
		TestResultStatusTrue,
		TestResultStatusUndefined,
	}
}