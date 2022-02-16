//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapp

const (
	moduleName    = "armapp"
	moduleVersion = "v0.1.0"
)

// AccessMode - Access mode for storage
type AccessMode string

const (
	AccessModeReadOnly  AccessMode = "ReadOnly"
	AccessModeReadWrite AccessMode = "ReadWrite"
)

// PossibleAccessModeValues returns the possible values for the AccessMode const type.
func PossibleAccessModeValues() []AccessMode {
	return []AccessMode{
		AccessModeReadOnly,
		AccessModeReadWrite,
	}
}

// ToPtr returns a *AccessMode pointing to the current value.
func (c AccessMode) ToPtr() *AccessMode {
	return &c
}

// ActiveRevisionsMode - ActiveRevisionsMode controls how active revisions are handled for the Container app:Multiple: multiple
// revisions can be active. If no value if provided, this is the defaultSingle: Only one revision
// can be active at a time. Revision weights can not be used in this mode
type ActiveRevisionsMode string

const (
	ActiveRevisionsModeMultiple ActiveRevisionsMode = "multiple"
	ActiveRevisionsModeSingle   ActiveRevisionsMode = "single"
)

// PossibleActiveRevisionsModeValues returns the possible values for the ActiveRevisionsMode const type.
func PossibleActiveRevisionsModeValues() []ActiveRevisionsMode {
	return []ActiveRevisionsMode{
		ActiveRevisionsModeMultiple,
		ActiveRevisionsModeSingle,
	}
}

// ToPtr returns a *ActiveRevisionsMode pointing to the current value.
func (c ActiveRevisionsMode) ToPtr() *ActiveRevisionsMode {
	return &c
}

// AppProtocol - Tells Dapr which protocol your application is using. Valid options are http and grpc. Default is http
type AppProtocol string

const (
	AppProtocolGrpc AppProtocol = "grpc"
	AppProtocolHTTP AppProtocol = "http"
)

// PossibleAppProtocolValues returns the possible values for the AppProtocol const type.
func PossibleAppProtocolValues() []AppProtocol {
	return []AppProtocol{
		AppProtocolGrpc,
		AppProtocolHTTP,
	}
}

// ToPtr returns a *AppProtocol pointing to the current value.
func (c AppProtocol) ToPtr() *AppProtocol {
	return &c
}

// BindingType - Custom Domain binding type.
type BindingType string

const (
	BindingTypeDisabled   BindingType = "Disabled"
	BindingTypeSniEnabled BindingType = "SniEnabled"
)

// PossibleBindingTypeValues returns the possible values for the BindingType const type.
func PossibleBindingTypeValues() []BindingType {
	return []BindingType{
		BindingTypeDisabled,
		BindingTypeSniEnabled,
	}
}

// ToPtr returns a *BindingType pointing to the current value.
func (c BindingType) ToPtr() *BindingType {
	return &c
}

// ContainerAppProvisioningState - Provisioning state of the Container App.
type ContainerAppProvisioningState string

const (
	ContainerAppProvisioningStateCanceled   ContainerAppProvisioningState = "Canceled"
	ContainerAppProvisioningStateFailed     ContainerAppProvisioningState = "Failed"
	ContainerAppProvisioningStateInProgress ContainerAppProvisioningState = "InProgress"
	ContainerAppProvisioningStateSucceeded  ContainerAppProvisioningState = "Succeeded"
)

// PossibleContainerAppProvisioningStateValues returns the possible values for the ContainerAppProvisioningState const type.
func PossibleContainerAppProvisioningStateValues() []ContainerAppProvisioningState {
	return []ContainerAppProvisioningState{
		ContainerAppProvisioningStateCanceled,
		ContainerAppProvisioningStateFailed,
		ContainerAppProvisioningStateInProgress,
		ContainerAppProvisioningStateSucceeded,
	}
}

// ToPtr returns a *ContainerAppProvisioningState pointing to the current value.
func (c ContainerAppProvisioningState) ToPtr() *ContainerAppProvisioningState {
	return &c
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

// ToPtr returns a *CreatedByType pointing to the current value.
func (c CreatedByType) ToPtr() *CreatedByType {
	return &c
}

// DNSVerificationTestResult - DNS verification test result.
type DNSVerificationTestResult string

const (
	DNSVerificationTestResultPassed  DNSVerificationTestResult = "Passed"
	DNSVerificationTestResultFailed  DNSVerificationTestResult = "Failed"
	DNSVerificationTestResultSkipped DNSVerificationTestResult = "Skipped"
)

// PossibleDNSVerificationTestResultValues returns the possible values for the DNSVerificationTestResult const type.
func PossibleDNSVerificationTestResultValues() []DNSVerificationTestResult {
	return []DNSVerificationTestResult{
		DNSVerificationTestResultPassed,
		DNSVerificationTestResultFailed,
		DNSVerificationTestResultSkipped,
	}
}

// ToPtr returns a *DNSVerificationTestResult pointing to the current value.
func (c DNSVerificationTestResult) ToPtr() *DNSVerificationTestResult {
	return &c
}

// DisableWwwAuthenticateMode - true if the www-authenticate provider should be omitted from the request; otherwise, false.
type DisableWwwAuthenticateMode string

const (
	DisableWwwAuthenticateModeFalse DisableWwwAuthenticateMode = "False"
	DisableWwwAuthenticateModeTrue  DisableWwwAuthenticateMode = "True"
)

// PossibleDisableWwwAuthenticateModeValues returns the possible values for the DisableWwwAuthenticateMode const type.
func PossibleDisableWwwAuthenticateModeValues() []DisableWwwAuthenticateMode {
	return []DisableWwwAuthenticateMode{
		DisableWwwAuthenticateModeFalse,
		DisableWwwAuthenticateModeTrue,
	}
}

// ToPtr returns a *DisableWwwAuthenticateMode pointing to the current value.
func (c DisableWwwAuthenticateMode) ToPtr() *DisableWwwAuthenticateMode {
	return &c
}

// EasyAuthState - Enabled if the Authentication / Authorization feature is enabled for the current app; otherwise, Disabled.
type EasyAuthState string

const (
	EasyAuthStateDisabled EasyAuthState = "Disabled"
	EasyAuthStateEnabled  EasyAuthState = "Enabled"
)

// PossibleEasyAuthStateValues returns the possible values for the EasyAuthState const type.
func PossibleEasyAuthStateValues() []EasyAuthState {
	return []EasyAuthState{
		EasyAuthStateDisabled,
		EasyAuthStateEnabled,
	}
}

// ToPtr returns a *EasyAuthState pointing to the current value.
func (c EasyAuthState) ToPtr() *EasyAuthState {
	return &c
}

// EnvironmentProvisioningState - Provisioning state of the Environment.
type EnvironmentProvisioningState string

const (
	EnvironmentProvisioningStateCanceled                      EnvironmentProvisioningState = "Canceled"
	EnvironmentProvisioningStateFailed                        EnvironmentProvisioningState = "Failed"
	EnvironmentProvisioningStateInfrastructureSetupComplete   EnvironmentProvisioningState = "InfrastructureSetupComplete"
	EnvironmentProvisioningStateInfrastructureSetupInProgress EnvironmentProvisioningState = "InfrastructureSetupInProgress"
	EnvironmentProvisioningStateInitializationInProgress      EnvironmentProvisioningState = "InitializationInProgress"
	EnvironmentProvisioningStateScheduledForDelete            EnvironmentProvisioningState = "ScheduledForDelete"
	EnvironmentProvisioningStateSucceeded                     EnvironmentProvisioningState = "Succeeded"
	EnvironmentProvisioningStateUpgradeFailed                 EnvironmentProvisioningState = "UpgradeFailed"
	EnvironmentProvisioningStateUpgradeRequested              EnvironmentProvisioningState = "UpgradeRequested"
	EnvironmentProvisioningStateWaiting                       EnvironmentProvisioningState = "Waiting"
)

// PossibleEnvironmentProvisioningStateValues returns the possible values for the EnvironmentProvisioningState const type.
func PossibleEnvironmentProvisioningStateValues() []EnvironmentProvisioningState {
	return []EnvironmentProvisioningState{
		EnvironmentProvisioningStateCanceled,
		EnvironmentProvisioningStateFailed,
		EnvironmentProvisioningStateInfrastructureSetupComplete,
		EnvironmentProvisioningStateInfrastructureSetupInProgress,
		EnvironmentProvisioningStateInitializationInProgress,
		EnvironmentProvisioningStateScheduledForDelete,
		EnvironmentProvisioningStateSucceeded,
		EnvironmentProvisioningStateUpgradeFailed,
		EnvironmentProvisioningStateUpgradeRequested,
		EnvironmentProvisioningStateWaiting,
	}
}

// ToPtr returns a *EnvironmentProvisioningState pointing to the current value.
func (c EnvironmentProvisioningState) ToPtr() *EnvironmentProvisioningState {
	return &c
}

// IdentityProviderState - Indicate whether identity provider is enabled or disabled.
type IdentityProviderState string

const (
	IdentityProviderStateDisabled IdentityProviderState = "Disabled"
	IdentityProviderStateEnabled  IdentityProviderState = "Enabled"
)

// PossibleIdentityProviderStateValues returns the possible values for the IdentityProviderState const type.
func PossibleIdentityProviderStateValues() []IdentityProviderState {
	return []IdentityProviderState{
		IdentityProviderStateDisabled,
		IdentityProviderStateEnabled,
	}
}

// ToPtr returns a *IdentityProviderState pointing to the current value.
func (c IdentityProviderState) ToPtr() *IdentityProviderState {
	return &c
}

// IngressTransportMethod - Ingress transport protocol
type IngressTransportMethod string

const (
	IngressTransportMethodAuto  IngressTransportMethod = "auto"
	IngressTransportMethodHTTP  IngressTransportMethod = "http"
	IngressTransportMethodHTTP2 IngressTransportMethod = "http2"
)

// PossibleIngressTransportMethodValues returns the possible values for the IngressTransportMethod const type.
func PossibleIngressTransportMethodValues() []IngressTransportMethod {
	return []IngressTransportMethod{
		IngressTransportMethodAuto,
		IngressTransportMethodHTTP,
		IngressTransportMethodHTTP2,
	}
}

// ToPtr returns a *IngressTransportMethod pointing to the current value.
func (c IngressTransportMethod) ToPtr() *IngressTransportMethod {
	return &c
}

// ManagedServiceIdentityType - Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                       ManagedServiceIdentityType = "None"
	ManagedServiceIdentityTypeSystemAssigned             ManagedServiceIdentityType = "SystemAssigned"
	ManagedServiceIdentityTypeSystemAssignedUserAssigned ManagedServiceIdentityType = "SystemAssigned,UserAssigned"
	ManagedServiceIdentityTypeUserAssigned               ManagedServiceIdentityType = "UserAssigned"
)

// PossibleManagedServiceIdentityTypeValues returns the possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{
		ManagedServiceIdentityTypeNone,
		ManagedServiceIdentityTypeSystemAssigned,
		ManagedServiceIdentityTypeSystemAssignedUserAssigned,
		ManagedServiceIdentityTypeUserAssigned,
	}
}

// ToPtr returns a *ManagedServiceIdentityType pointing to the current value.
func (c ManagedServiceIdentityType) ToPtr() *ManagedServiceIdentityType {
	return &c
}

// PreserveURLFragmentsForLoginsMode - True if the fragments from the request are preserved after the login request is made;
// otherwise, False.
type PreserveURLFragmentsForLoginsMode string

const (
	PreserveURLFragmentsForLoginsModeFalse PreserveURLFragmentsForLoginsMode = "False"
	PreserveURLFragmentsForLoginsModeTrue  PreserveURLFragmentsForLoginsMode = "True"
)

// PossiblePreserveURLFragmentsForLoginsModeValues returns the possible values for the PreserveURLFragmentsForLoginsMode const type.
func PossiblePreserveURLFragmentsForLoginsModeValues() []PreserveURLFragmentsForLoginsMode {
	return []PreserveURLFragmentsForLoginsMode{
		PreserveURLFragmentsForLoginsModeFalse,
		PreserveURLFragmentsForLoginsModeTrue,
	}
}

// ToPtr returns a *PreserveURLFragmentsForLoginsMode pointing to the current value.
func (c PreserveURLFragmentsForLoginsMode) ToPtr() *PreserveURLFragmentsForLoginsMode {
	return &c
}

// RequireHTTPSMode - false if the authentication/authorization responses not having the HTTPS scheme are permissible; otherwise,
// true.
type RequireHTTPSMode string

const (
	RequireHTTPSModeFalse RequireHTTPSMode = "False"
	RequireHTTPSModeTrue  RequireHTTPSMode = "True"
)

// PossibleRequireHTTPSModeValues returns the possible values for the RequireHTTPSMode const type.
func PossibleRequireHTTPSModeValues() []RequireHTTPSMode {
	return []RequireHTTPSMode{
		RequireHTTPSModeFalse,
		RequireHTTPSModeTrue,
	}
}

// ToPtr returns a *RequireHTTPSMode pointing to the current value.
func (c RequireHTTPSMode) ToPtr() *RequireHTTPSMode {
	return &c
}

// RevisionHealthState - Current health State of the revision
type RevisionHealthState string

const (
	RevisionHealthStateHealthy   RevisionHealthState = "Healthy"
	RevisionHealthStateNone      RevisionHealthState = "None"
	RevisionHealthStateUnhealthy RevisionHealthState = "Unhealthy"
)

// PossibleRevisionHealthStateValues returns the possible values for the RevisionHealthState const type.
func PossibleRevisionHealthStateValues() []RevisionHealthState {
	return []RevisionHealthState{
		RevisionHealthStateHealthy,
		RevisionHealthStateNone,
		RevisionHealthStateUnhealthy,
	}
}

// ToPtr returns a *RevisionHealthState pointing to the current value.
func (c RevisionHealthState) ToPtr() *RevisionHealthState {
	return &c
}

// RevisionProvisioningState - Current provisioning State of the revision
type RevisionProvisioningState string

const (
	RevisionProvisioningStateDeprovisioned  RevisionProvisioningState = "Deprovisioned"
	RevisionProvisioningStateDeprovisioning RevisionProvisioningState = "Deprovisioning"
	RevisionProvisioningStateFailed         RevisionProvisioningState = "Failed"
	RevisionProvisioningStateProvisioned    RevisionProvisioningState = "Provisioned"
	RevisionProvisioningStateProvisioning   RevisionProvisioningState = "Provisioning"
)

// PossibleRevisionProvisioningStateValues returns the possible values for the RevisionProvisioningState const type.
func PossibleRevisionProvisioningStateValues() []RevisionProvisioningState {
	return []RevisionProvisioningState{
		RevisionProvisioningStateDeprovisioned,
		RevisionProvisioningStateDeprovisioning,
		RevisionProvisioningStateFailed,
		RevisionProvisioningStateProvisioned,
		RevisionProvisioningStateProvisioning,
	}
}

// ToPtr returns a *RevisionProvisioningState pointing to the current value.
func (c RevisionProvisioningState) ToPtr() *RevisionProvisioningState {
	return &c
}

// SourceControlOperationState - Current provisioning State of the operation
type SourceControlOperationState string

const (
	SourceControlOperationStateCanceled   SourceControlOperationState = "Canceled"
	SourceControlOperationStateFailed     SourceControlOperationState = "Failed"
	SourceControlOperationStateInProgress SourceControlOperationState = "InProgress"
	SourceControlOperationStateSucceeded  SourceControlOperationState = "Succeeded"
)

// PossibleSourceControlOperationStateValues returns the possible values for the SourceControlOperationState const type.
func PossibleSourceControlOperationStateValues() []SourceControlOperationState {
	return []SourceControlOperationState{
		SourceControlOperationStateCanceled,
		SourceControlOperationStateFailed,
		SourceControlOperationStateInProgress,
		SourceControlOperationStateSucceeded,
	}
}

// ToPtr returns a *SourceControlOperationState pointing to the current value.
func (c SourceControlOperationState) ToPtr() *SourceControlOperationState {
	return &c
}

// StorageType - Storage type for the volume. If not provided, use EmptyDir.
type StorageType string

const (
	StorageTypeAzureFile StorageType = "AzureFile"
	StorageTypeEmptyDir  StorageType = "EmptyDir"
)

// PossibleStorageTypeValues returns the possible values for the StorageType const type.
func PossibleStorageTypeValues() []StorageType {
	return []StorageType{
		StorageTypeAzureFile,
		StorageTypeEmptyDir,
	}
}

// ToPtr returns a *StorageType pointing to the current value.
func (c StorageType) ToPtr() *StorageType {
	return &c
}

// Type - The type of probe.
type Type string

const (
	TypeLiveness  Type = "liveness"
	TypeReadiness Type = "readiness"
	TypeStartup   Type = "startup"
)

// PossibleTypeValues returns the possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{
		TypeLiveness,
		TypeReadiness,
		TypeStartup,
	}
}

// ToPtr returns a *Type pointing to the current value.
func (c Type) ToPtr() *Type {
	return &c
}

// UnauthenticatedClientAction - The action to take when an unauthenticated client attempts to access the app.
type UnauthenticatedClientAction string

const (
	UnauthenticatedClientActionAllowAnonymous      UnauthenticatedClientAction = "AllowAnonymous"
	UnauthenticatedClientActionRedirectToLoginPage UnauthenticatedClientAction = "RedirectToLoginPage"
	UnauthenticatedClientActionReturn401           UnauthenticatedClientAction = "Return401"
	UnauthenticatedClientActionReturn403           UnauthenticatedClientAction = "Return403"
)

// PossibleUnauthenticatedClientActionValues returns the possible values for the UnauthenticatedClientAction const type.
func PossibleUnauthenticatedClientActionValues() []UnauthenticatedClientAction {
	return []UnauthenticatedClientAction{
		UnauthenticatedClientActionAllowAnonymous,
		UnauthenticatedClientActionRedirectToLoginPage,
		UnauthenticatedClientActionReturn401,
		UnauthenticatedClientActionReturn403,
	}
}

// ToPtr returns a *UnauthenticatedClientAction pointing to the current value.
func (c UnauthenticatedClientAction) ToPtr() *UnauthenticatedClientAction {
	return &c
}
