//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnginx

const (
	moduleName    = "armnginx"
	moduleVersion = "v2.0.0"
)

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

type IdentityType string

const (
	IdentityTypeNone                       IdentityType = "None"
	IdentityTypeSystemAssigned             IdentityType = "SystemAssigned"
	IdentityTypeSystemAssignedUserAssigned IdentityType = "SystemAssigned, UserAssigned"
	IdentityTypeUserAssigned               IdentityType = "UserAssigned"
)

// PossibleIdentityTypeValues returns the possible values for the IdentityType const type.
func PossibleIdentityTypeValues() []IdentityType {
	return []IdentityType{
		IdentityTypeNone,
		IdentityTypeSystemAssigned,
		IdentityTypeSystemAssignedUserAssigned,
		IdentityTypeUserAssigned,
	}
}

type NginxPrivateIPAllocationMethod string

const (
	NginxPrivateIPAllocationMethodDynamic NginxPrivateIPAllocationMethod = "Dynamic"
	NginxPrivateIPAllocationMethodStatic  NginxPrivateIPAllocationMethod = "Static"
)

// PossibleNginxPrivateIPAllocationMethodValues returns the possible values for the NginxPrivateIPAllocationMethod const type.
func PossibleNginxPrivateIPAllocationMethodValues() []NginxPrivateIPAllocationMethod {
	return []NginxPrivateIPAllocationMethod{
		NginxPrivateIPAllocationMethodDynamic,
		NginxPrivateIPAllocationMethodStatic,
	}
}

type ProvisioningState string

const (
	ProvisioningStateAccepted     ProvisioningState = "Accepted"
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateCreating     ProvisioningState = "Creating"
	ProvisioningStateDeleted      ProvisioningState = "Deleted"
	ProvisioningStateDeleting     ProvisioningState = "Deleting"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateNotSpecified ProvisioningState = "NotSpecified"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleted,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateNotSpecified,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}