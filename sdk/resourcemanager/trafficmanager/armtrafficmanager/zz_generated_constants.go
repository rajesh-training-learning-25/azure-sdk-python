//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtrafficmanager

const (
	module  = "armtrafficmanager"
	version = "v0.1.0"
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

// ToPtr returns a *AllowedEndpointRecordType pointing to the current value.
func (c AllowedEndpointRecordType) ToPtr() *AllowedEndpointRecordType {
	return &c
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

// ToPtr returns a *EndpointMonitorStatus pointing to the current value.
func (c EndpointMonitorStatus) ToPtr() *EndpointMonitorStatus {
	return &c
}

// EndpointStatus - The status of the endpoint. If the endpoint is Enabled, it is probed for endpoint health and is included in the traffic routing method.
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

// ToPtr returns a *EndpointStatus pointing to the current value.
func (c EndpointStatus) ToPtr() *EndpointStatus {
	return &c
}

type Enum8 string

const (
	Enum8Default Enum8 = "default"
)

// PossibleEnum8Values returns the possible values for the Enum8 const type.
func PossibleEnum8Values() []Enum8 {
	return []Enum8{
		Enum8Default,
	}
}

// ToPtr returns a *Enum8 pointing to the current value.
func (c Enum8) ToPtr() *Enum8 {
	return &c
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

// ToPtr returns a *MonitorProtocol pointing to the current value.
func (c MonitorProtocol) ToPtr() *MonitorProtocol {
	return &c
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

// ToPtr returns a *ProfileMonitorStatus pointing to the current value.
func (c ProfileMonitorStatus) ToPtr() *ProfileMonitorStatus {
	return &c
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

// ToPtr returns a *ProfileStatus pointing to the current value.
func (c ProfileStatus) ToPtr() *ProfileStatus {
	return &c
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

// ToPtr returns a *TrafficRoutingMethod pointing to the current value.
func (c TrafficRoutingMethod) ToPtr() *TrafficRoutingMethod {
	return &c
}

// TrafficViewEnrollmentStatus - Indicates whether Traffic View is 'Enabled' or 'Disabled' for the Traffic Manager profile. Null, indicates 'Disabled'.
// Enabling this feature will increase the cost of the Traffic Manage profile.
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

// ToPtr returns a *TrafficViewEnrollmentStatus pointing to the current value.
func (c TrafficViewEnrollmentStatus) ToPtr() *TrafficViewEnrollmentStatus {
	return &c
}
