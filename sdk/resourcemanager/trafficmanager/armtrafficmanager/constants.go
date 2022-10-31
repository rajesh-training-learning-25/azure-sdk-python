//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armtrafficmanager

const (
	moduleName    = "armtrafficmanager"
	moduleVersion = "1.1.0-beta.1"
)

// AllowedEndpointRecordType - The allowed type DNS record types for this profile.
type AllowedEndpointRecordType string

const (
	AllowedEndpointRecordTypeAny         AllowedEndpointRecordType = "Any"
	AllowedEndpointRecordTypeDomainName  AllowedEndpointRecordType = "DomainName"
	AllowedEndpointRecordTypeIPv4Address AllowedEndpointRecordType = "IPv4Address"
	AllowedEndpointRecordTypeIPv6Address AllowedEndpointRecordType = "IPv6Address"
)

// PossibleAllowedEndpointRecordTypeValues returns the possible values for the AllowedEndpointRecordType const type.
func PossibleAllowedEndpointRecordTypeValues() []AllowedEndpointRecordType {
	return []AllowedEndpointRecordType{
		AllowedEndpointRecordTypeAny,
		AllowedEndpointRecordTypeDomainName,
		AllowedEndpointRecordTypeIPv4Address,
		AllowedEndpointRecordTypeIPv6Address,
	}
}

// AlwaysServe - If Always Serve is enabled, probing for endpoint health will be disabled and endpoints will be included in
// the traffic routing method.
type AlwaysServe string

const (
	AlwaysServeDisabled AlwaysServe = "Disabled"
	AlwaysServeEnabled  AlwaysServe = "Enabled"
)

// PossibleAlwaysServeValues returns the possible values for the AlwaysServe const type.
func PossibleAlwaysServeValues() []AlwaysServe {
	return []AlwaysServe{
		AlwaysServeDisabled,
		AlwaysServeEnabled,
	}
}

// EndpointMonitorStatus - The monitoring status of the endpoint.
type EndpointMonitorStatus string

const (
	EndpointMonitorStatusCheckingEndpoint EndpointMonitorStatus = "CheckingEndpoint"
	EndpointMonitorStatusDegraded         EndpointMonitorStatus = "Degraded"
	EndpointMonitorStatusDisabled         EndpointMonitorStatus = "Disabled"
	EndpointMonitorStatusInactive         EndpointMonitorStatus = "Inactive"
	EndpointMonitorStatusOnline           EndpointMonitorStatus = "Online"
	EndpointMonitorStatusStopped          EndpointMonitorStatus = "Stopped"
)

// PossibleEndpointMonitorStatusValues returns the possible values for the EndpointMonitorStatus const type.
func PossibleEndpointMonitorStatusValues() []EndpointMonitorStatus {
	return []EndpointMonitorStatus{
		EndpointMonitorStatusCheckingEndpoint,
		EndpointMonitorStatusDegraded,
		EndpointMonitorStatusDisabled,
		EndpointMonitorStatusInactive,
		EndpointMonitorStatusOnline,
		EndpointMonitorStatusStopped,
	}
}

// EndpointStatus - The status of the endpoint. If the endpoint is Enabled, it is probed for endpoint health and is included
// in the traffic routing method.
type EndpointStatus string

const (
	EndpointStatusDisabled EndpointStatus = "Disabled"
	EndpointStatusEnabled  EndpointStatus = "Enabled"
)

// PossibleEndpointStatusValues returns the possible values for the EndpointStatus const type.
func PossibleEndpointStatusValues() []EndpointStatus {
	return []EndpointStatus{
		EndpointStatusDisabled,
		EndpointStatusEnabled,
	}
}

type EndpointType string

const (
	EndpointTypeAzureEndpoints    EndpointType = "AzureEndpoints"
	EndpointTypeExternalEndpoints EndpointType = "ExternalEndpoints"
	EndpointTypeNestedEndpoints   EndpointType = "NestedEndpoints"
)

// PossibleEndpointTypeValues returns the possible values for the EndpointType const type.
func PossibleEndpointTypeValues() []EndpointType {
	return []EndpointType{
		EndpointTypeAzureEndpoints,
		EndpointTypeExternalEndpoints,
		EndpointTypeNestedEndpoints,
	}
}

// MonitorProtocol - The protocol (HTTP, HTTPS or TCP) used to probe for endpoint health.
type MonitorProtocol string

const (
	MonitorProtocolHTTP  MonitorProtocol = "HTTP"
	MonitorProtocolHTTPS MonitorProtocol = "HTTPS"
	MonitorProtocolTCP   MonitorProtocol = "TCP"
)

// PossibleMonitorProtocolValues returns the possible values for the MonitorProtocol const type.
func PossibleMonitorProtocolValues() []MonitorProtocol {
	return []MonitorProtocol{
		MonitorProtocolHTTP,
		MonitorProtocolHTTPS,
		MonitorProtocolTCP,
	}
}

// ProfileMonitorStatus - The profile-level monitoring status of the Traffic Manager profile.
type ProfileMonitorStatus string

const (
	ProfileMonitorStatusCheckingEndpoints ProfileMonitorStatus = "CheckingEndpoints"
	ProfileMonitorStatusDegraded          ProfileMonitorStatus = "Degraded"
	ProfileMonitorStatusDisabled          ProfileMonitorStatus = "Disabled"
	ProfileMonitorStatusInactive          ProfileMonitorStatus = "Inactive"
	ProfileMonitorStatusOnline            ProfileMonitorStatus = "Online"
)

// PossibleProfileMonitorStatusValues returns the possible values for the ProfileMonitorStatus const type.
func PossibleProfileMonitorStatusValues() []ProfileMonitorStatus {
	return []ProfileMonitorStatus{
		ProfileMonitorStatusCheckingEndpoints,
		ProfileMonitorStatusDegraded,
		ProfileMonitorStatusDisabled,
		ProfileMonitorStatusInactive,
		ProfileMonitorStatusOnline,
	}
}

// ProfileStatus - The status of the Traffic Manager profile.
type ProfileStatus string

const (
	ProfileStatusDisabled ProfileStatus = "Disabled"
	ProfileStatusEnabled  ProfileStatus = "Enabled"
)

// PossibleProfileStatusValues returns the possible values for the ProfileStatus const type.
func PossibleProfileStatusValues() []ProfileStatus {
	return []ProfileStatus{
		ProfileStatusDisabled,
		ProfileStatusEnabled,
	}
}

// TrafficRoutingMethod - The traffic routing method of the Traffic Manager profile.
type TrafficRoutingMethod string

const (
	TrafficRoutingMethodGeographic  TrafficRoutingMethod = "Geographic"
	TrafficRoutingMethodMultiValue  TrafficRoutingMethod = "MultiValue"
	TrafficRoutingMethodPerformance TrafficRoutingMethod = "Performance"
	TrafficRoutingMethodPriority    TrafficRoutingMethod = "Priority"
	TrafficRoutingMethodSubnet      TrafficRoutingMethod = "Subnet"
	TrafficRoutingMethodWeighted    TrafficRoutingMethod = "Weighted"
)

// PossibleTrafficRoutingMethodValues returns the possible values for the TrafficRoutingMethod const type.
func PossibleTrafficRoutingMethodValues() []TrafficRoutingMethod {
	return []TrafficRoutingMethod{
		TrafficRoutingMethodGeographic,
		TrafficRoutingMethodMultiValue,
		TrafficRoutingMethodPerformance,
		TrafficRoutingMethodPriority,
		TrafficRoutingMethodSubnet,
		TrafficRoutingMethodWeighted,
	}
}

// TrafficViewEnrollmentStatus - Indicates whether Traffic View is 'Enabled' or 'Disabled' for the Traffic Manager profile.
// Null, indicates 'Disabled'. Enabling this feature will increase the cost of the Traffic Manage profile.
type TrafficViewEnrollmentStatus string

const (
	TrafficViewEnrollmentStatusDisabled TrafficViewEnrollmentStatus = "Disabled"
	TrafficViewEnrollmentStatusEnabled  TrafficViewEnrollmentStatus = "Enabled"
)

// PossibleTrafficViewEnrollmentStatusValues returns the possible values for the TrafficViewEnrollmentStatus const type.
func PossibleTrafficViewEnrollmentStatusValues() []TrafficViewEnrollmentStatus {
	return []TrafficViewEnrollmentStatus{
		TrafficViewEnrollmentStatusDisabled,
		TrafficViewEnrollmentStatusEnabled,
	}
}
