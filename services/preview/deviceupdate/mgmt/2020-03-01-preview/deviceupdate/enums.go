package deviceupdate

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ActionType enumerates the values for action type.
type ActionType string

const (
	// Internal ...
	Internal ActionType = "Internal"
)

// PossibleActionTypeValues returns an array of possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{Internal}
}

// CheckNameAvailabilityReason enumerates the values for check name availability reason.
type CheckNameAvailabilityReason string

const (
	// AlreadyExists ...
	AlreadyExists CheckNameAvailabilityReason = "AlreadyExists"
	// Invalid ...
	Invalid CheckNameAvailabilityReason = "Invalid"
)

// PossibleCheckNameAvailabilityReasonValues returns an array of possible values for the CheckNameAvailabilityReason const type.
func PossibleCheckNameAvailabilityReasonValues() []CheckNameAvailabilityReason {
	return []CheckNameAvailabilityReason{AlreadyExists, Invalid}
}

// CreatedByType enumerates the values for created by type.
type CreatedByType string

const (
	// Application ...
	Application CreatedByType = "Application"
	// Key ...
	Key CreatedByType = "Key"
	// ManagedIdentity ...
	ManagedIdentity CreatedByType = "ManagedIdentity"
	// User ...
	User CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns an array of possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{Application, Key, ManagedIdentity, User}
}

// GroupIDProvisioningState enumerates the values for group id provisioning state.
type GroupIDProvisioningState string

const (
	// Canceled ...
	Canceled GroupIDProvisioningState = "Canceled"
	// Failed ...
	Failed GroupIDProvisioningState = "Failed"
	// Succeeded ...
	Succeeded GroupIDProvisioningState = "Succeeded"
)

// PossibleGroupIDProvisioningStateValues returns an array of possible values for the GroupIDProvisioningState const type.
func PossibleGroupIDProvisioningStateValues() []GroupIDProvisioningState {
	return []GroupIDProvisioningState{Canceled, Failed, Succeeded}
}

// ManagedServiceIdentityType enumerates the values for managed service identity type.
type ManagedServiceIdentityType string

const (
	// None ...
	None ManagedServiceIdentityType = "None"
	// SystemAssigned ...
	SystemAssigned ManagedServiceIdentityType = "SystemAssigned"
	// SystemAssignedUserAssigned ...
	SystemAssignedUserAssigned ManagedServiceIdentityType = "SystemAssigned,UserAssigned"
	// UserAssigned ...
	UserAssigned ManagedServiceIdentityType = "UserAssigned"
)

// PossibleManagedServiceIdentityTypeValues returns an array of possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{None, SystemAssigned, SystemAssignedUserAssigned, UserAssigned}
}

// Origin enumerates the values for origin.
type Origin string

const (
	// OriginSystem ...
	OriginSystem Origin = "system"
	// OriginUser ...
	OriginUser Origin = "user"
	// OriginUsersystem ...
	OriginUsersystem Origin = "user,system"
)

// PossibleOriginValues returns an array of possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{OriginSystem, OriginUser, OriginUsersystem}
}

// PrivateEndpointConnectionProvisioningState enumerates the values for private endpoint connection
// provisioning state.
type PrivateEndpointConnectionProvisioningState string

const (
	// PrivateEndpointConnectionProvisioningStateCreating ...
	PrivateEndpointConnectionProvisioningStateCreating PrivateEndpointConnectionProvisioningState = "Creating"
	// PrivateEndpointConnectionProvisioningStateDeleting ...
	PrivateEndpointConnectionProvisioningStateDeleting PrivateEndpointConnectionProvisioningState = "Deleting"
	// PrivateEndpointConnectionProvisioningStateFailed ...
	PrivateEndpointConnectionProvisioningStateFailed PrivateEndpointConnectionProvisioningState = "Failed"
	// PrivateEndpointConnectionProvisioningStateSucceeded ...
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = "Succeeded"
)

// PossiblePrivateEndpointConnectionProvisioningStateValues returns an array of possible values for the PrivateEndpointConnectionProvisioningState const type.
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{PrivateEndpointConnectionProvisioningStateCreating, PrivateEndpointConnectionProvisioningStateDeleting, PrivateEndpointConnectionProvisioningStateFailed, PrivateEndpointConnectionProvisioningStateSucceeded}
}

// PrivateEndpointServiceConnectionStatus enumerates the values for private endpoint service connection status.
type PrivateEndpointServiceConnectionStatus string

const (
	// Approved ...
	Approved PrivateEndpointServiceConnectionStatus = "Approved"
	// Pending ...
	Pending PrivateEndpointServiceConnectionStatus = "Pending"
	// Rejected ...
	Rejected PrivateEndpointServiceConnectionStatus = "Rejected"
)

// PossiblePrivateEndpointServiceConnectionStatusValues returns an array of possible values for the PrivateEndpointServiceConnectionStatus const type.
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return []PrivateEndpointServiceConnectionStatus{Approved, Pending, Rejected}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// ProvisioningStateAccepted ...
	ProvisioningStateAccepted ProvisioningState = "Accepted"
	// ProvisioningStateCanceled ...
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateCreating ...
	ProvisioningStateCreating ProvisioningState = "Creating"
	// ProvisioningStateDeleted ...
	ProvisioningStateDeleted ProvisioningState = "Deleted"
	// ProvisioningStateFailed ...
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateSucceeded ...
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{ProvisioningStateAccepted, ProvisioningStateCanceled, ProvisioningStateCreating, ProvisioningStateDeleted, ProvisioningStateFailed, ProvisioningStateSucceeded}
}

// PublicNetworkAccess enumerates the values for public network access.
type PublicNetworkAccess string

const (
	// Disabled ...
	Disabled PublicNetworkAccess = "Disabled"
	// Enabled ...
	Enabled PublicNetworkAccess = "Enabled"
)

// PossiblePublicNetworkAccessValues returns an array of possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{Disabled, Enabled}
}
