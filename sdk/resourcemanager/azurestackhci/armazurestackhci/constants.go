//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurestackhci

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci"
	moduleVersion = "v1.2.0"
)

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// ArcSettingAggregateState - Aggregate state of Arc agent across the nodes in this HCI cluster.
type ArcSettingAggregateState string

const (
	ArcSettingAggregateStateCanceled           ArcSettingAggregateState = "Canceled"
	ArcSettingAggregateStateConnected          ArcSettingAggregateState = "Connected"
	ArcSettingAggregateStateCreating           ArcSettingAggregateState = "Creating"
	ArcSettingAggregateStateDeleted            ArcSettingAggregateState = "Deleted"
	ArcSettingAggregateStateDeleting           ArcSettingAggregateState = "Deleting"
	ArcSettingAggregateStateDisconnected       ArcSettingAggregateState = "Disconnected"
	ArcSettingAggregateStateError              ArcSettingAggregateState = "Error"
	ArcSettingAggregateStateFailed             ArcSettingAggregateState = "Failed"
	ArcSettingAggregateStateInProgress         ArcSettingAggregateState = "InProgress"
	ArcSettingAggregateStateMoving             ArcSettingAggregateState = "Moving"
	ArcSettingAggregateStateNotSpecified       ArcSettingAggregateState = "NotSpecified"
	ArcSettingAggregateStatePartiallyConnected ArcSettingAggregateState = "PartiallyConnected"
	ArcSettingAggregateStatePartiallySucceeded ArcSettingAggregateState = "PartiallySucceeded"
	ArcSettingAggregateStateSucceeded          ArcSettingAggregateState = "Succeeded"
	ArcSettingAggregateStateUpdating           ArcSettingAggregateState = "Updating"
)

// PossibleArcSettingAggregateStateValues returns the possible values for the ArcSettingAggregateState const type.
func PossibleArcSettingAggregateStateValues() []ArcSettingAggregateState {
	return []ArcSettingAggregateState{
		ArcSettingAggregateStateCanceled,
		ArcSettingAggregateStateConnected,
		ArcSettingAggregateStateCreating,
		ArcSettingAggregateStateDeleted,
		ArcSettingAggregateStateDeleting,
		ArcSettingAggregateStateDisconnected,
		ArcSettingAggregateStateError,
		ArcSettingAggregateStateFailed,
		ArcSettingAggregateStateInProgress,
		ArcSettingAggregateStateMoving,
		ArcSettingAggregateStateNotSpecified,
		ArcSettingAggregateStatePartiallyConnected,
		ArcSettingAggregateStatePartiallySucceeded,
		ArcSettingAggregateStateSucceeded,
		ArcSettingAggregateStateUpdating,
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

// DiagnosticLevel - Desired level of diagnostic data emitted by the cluster.
type DiagnosticLevel string

const (
	DiagnosticLevelBasic    DiagnosticLevel = "Basic"
	DiagnosticLevelEnhanced DiagnosticLevel = "Enhanced"
	DiagnosticLevelOff      DiagnosticLevel = "Off"
)

// PossibleDiagnosticLevelValues returns the possible values for the DiagnosticLevel const type.
func PossibleDiagnosticLevelValues() []DiagnosticLevel {
	return []DiagnosticLevel{
		DiagnosticLevelBasic,
		DiagnosticLevelEnhanced,
		DiagnosticLevelOff,
	}
}

// ExtensionAggregateState - Aggregate state of Arc Extensions across the nodes in this HCI cluster.
type ExtensionAggregateState string

const (
	ExtensionAggregateStateCanceled           ExtensionAggregateState = "Canceled"
	ExtensionAggregateStateConnected          ExtensionAggregateState = "Connected"
	ExtensionAggregateStateCreating           ExtensionAggregateState = "Creating"
	ExtensionAggregateStateDeleted            ExtensionAggregateState = "Deleted"
	ExtensionAggregateStateDeleting           ExtensionAggregateState = "Deleting"
	ExtensionAggregateStateDisconnected       ExtensionAggregateState = "Disconnected"
	ExtensionAggregateStateError              ExtensionAggregateState = "Error"
	ExtensionAggregateStateFailed             ExtensionAggregateState = "Failed"
	ExtensionAggregateStateInProgress         ExtensionAggregateState = "InProgress"
	ExtensionAggregateStateMoving             ExtensionAggregateState = "Moving"
	ExtensionAggregateStateNotSpecified       ExtensionAggregateState = "NotSpecified"
	ExtensionAggregateStatePartiallyConnected ExtensionAggregateState = "PartiallyConnected"
	ExtensionAggregateStatePartiallySucceeded ExtensionAggregateState = "PartiallySucceeded"
	ExtensionAggregateStateSucceeded          ExtensionAggregateState = "Succeeded"
	ExtensionAggregateStateUpdating           ExtensionAggregateState = "Updating"
)

// PossibleExtensionAggregateStateValues returns the possible values for the ExtensionAggregateState const type.
func PossibleExtensionAggregateStateValues() []ExtensionAggregateState {
	return []ExtensionAggregateState{
		ExtensionAggregateStateCanceled,
		ExtensionAggregateStateConnected,
		ExtensionAggregateStateCreating,
		ExtensionAggregateStateDeleted,
		ExtensionAggregateStateDeleting,
		ExtensionAggregateStateDisconnected,
		ExtensionAggregateStateError,
		ExtensionAggregateStateFailed,
		ExtensionAggregateStateInProgress,
		ExtensionAggregateStateMoving,
		ExtensionAggregateStateNotSpecified,
		ExtensionAggregateStatePartiallyConnected,
		ExtensionAggregateStatePartiallySucceeded,
		ExtensionAggregateStateSucceeded,
		ExtensionAggregateStateUpdating,
	}
}

// ImdsAttestation - IMDS attestation status of the cluster.
type ImdsAttestation string

const (
	ImdsAttestationDisabled ImdsAttestation = "Disabled"
	ImdsAttestationEnabled  ImdsAttestation = "Enabled"
)

// PossibleImdsAttestationValues returns the possible values for the ImdsAttestation const type.
func PossibleImdsAttestationValues() []ImdsAttestation {
	return []ImdsAttestation{
		ImdsAttestationDisabled,
		ImdsAttestationEnabled,
	}
}

// NodeArcState - State of Arc agent in this node.
type NodeArcState string

const (
	NodeArcStateCanceled     NodeArcState = "Canceled"
	NodeArcStateConnected    NodeArcState = "Connected"
	NodeArcStateCreating     NodeArcState = "Creating"
	NodeArcStateDeleted      NodeArcState = "Deleted"
	NodeArcStateDeleting     NodeArcState = "Deleting"
	NodeArcStateDisconnected NodeArcState = "Disconnected"
	NodeArcStateError        NodeArcState = "Error"
	NodeArcStateFailed       NodeArcState = "Failed"
	NodeArcStateMoving       NodeArcState = "Moving"
	NodeArcStateNotSpecified NodeArcState = "NotSpecified"
	NodeArcStateSucceeded    NodeArcState = "Succeeded"
	NodeArcStateUpdating     NodeArcState = "Updating"
)

// PossibleNodeArcStateValues returns the possible values for the NodeArcState const type.
func PossibleNodeArcStateValues() []NodeArcState {
	return []NodeArcState{
		NodeArcStateCanceled,
		NodeArcStateConnected,
		NodeArcStateCreating,
		NodeArcStateDeleted,
		NodeArcStateDeleting,
		NodeArcStateDisconnected,
		NodeArcStateError,
		NodeArcStateFailed,
		NodeArcStateMoving,
		NodeArcStateNotSpecified,
		NodeArcStateSucceeded,
		NodeArcStateUpdating,
	}
}

// NodeExtensionState - State of Arc Extension in this node.
type NodeExtensionState string

const (
	NodeExtensionStateCanceled     NodeExtensionState = "Canceled"
	NodeExtensionStateConnected    NodeExtensionState = "Connected"
	NodeExtensionStateCreating     NodeExtensionState = "Creating"
	NodeExtensionStateDeleted      NodeExtensionState = "Deleted"
	NodeExtensionStateDeleting     NodeExtensionState = "Deleting"
	NodeExtensionStateDisconnected NodeExtensionState = "Disconnected"
	NodeExtensionStateError        NodeExtensionState = "Error"
	NodeExtensionStateFailed       NodeExtensionState = "Failed"
	NodeExtensionStateMoving       NodeExtensionState = "Moving"
	NodeExtensionStateNotSpecified NodeExtensionState = "NotSpecified"
	NodeExtensionStateSucceeded    NodeExtensionState = "Succeeded"
	NodeExtensionStateUpdating     NodeExtensionState = "Updating"
)

// PossibleNodeExtensionStateValues returns the possible values for the NodeExtensionState const type.
func PossibleNodeExtensionStateValues() []NodeExtensionState {
	return []NodeExtensionState{
		NodeExtensionStateCanceled,
		NodeExtensionStateConnected,
		NodeExtensionStateCreating,
		NodeExtensionStateDeleted,
		NodeExtensionStateDeleting,
		NodeExtensionStateDisconnected,
		NodeExtensionStateError,
		NodeExtensionStateFailed,
		NodeExtensionStateMoving,
		NodeExtensionStateNotSpecified,
		NodeExtensionStateSucceeded,
		NodeExtensionStateUpdating,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	OriginSystem     Origin = "system"
	OriginUser       Origin = "user"
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// ProvisioningState - Provisioning state of the ArcSetting proxy resource.
type ProvisioningState string

const (
	ProvisioningStateAccepted     ProvisioningState = "Accepted"
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateFailed,
		ProvisioningStateProvisioning,
		ProvisioningStateSucceeded,
	}
}

// Status - Status of the cluster agent.
type Status string

const (
	StatusConnectedRecently    Status = "ConnectedRecently"
	StatusDisconnected         Status = "Disconnected"
	StatusError                Status = "Error"
	StatusNotConnectedRecently Status = "NotConnectedRecently"
	StatusNotYetRegistered     Status = "NotYetRegistered"
)

// PossibleStatusValues returns the possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{
		StatusConnectedRecently,
		StatusDisconnected,
		StatusError,
		StatusNotConnectedRecently,
		StatusNotYetRegistered,
	}
}

// WindowsServerSubscription - Desired state of Windows Server Subscription.
type WindowsServerSubscription string

const (
	WindowsServerSubscriptionDisabled WindowsServerSubscription = "Disabled"
	WindowsServerSubscriptionEnabled  WindowsServerSubscription = "Enabled"
)

// PossibleWindowsServerSubscriptionValues returns the possible values for the WindowsServerSubscription const type.
func PossibleWindowsServerSubscriptionValues() []WindowsServerSubscription {
	return []WindowsServerSubscription{
		WindowsServerSubscriptionDisabled,
		WindowsServerSubscriptionEnabled,
	}
}
