//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdomainservices

const (
	module  = "armdomainservices"
	version = "v0.1.0"
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

// ToPtr returns a *CreatedByType pointing to the current value.
func (c CreatedByType) ToPtr() *CreatedByType {
	return &c
}

// ExternalAccess - A flag to determine whether or not Secure LDAP access over the internet is enabled or disabled.
type ExternalAccess string

const (
	ExternalAccessDisabled ExternalAccess = "Disabled"
	ExternalAccessEnabled  ExternalAccess = "Enabled"
)

// PossibleExternalAccessValues returns the possible values for the ExternalAccess const type.
func PossibleExternalAccessValues() []ExternalAccess {
	return []ExternalAccess{
		ExternalAccessDisabled,
		ExternalAccessEnabled,
	}
}

// ToPtr returns a *ExternalAccess pointing to the current value.
func (c ExternalAccess) ToPtr() *ExternalAccess {
	return &c
}

// FilteredSync - Enabled or Disabled flag to turn on Group-based filtered sync
type FilteredSync string

const (
	FilteredSyncDisabled FilteredSync = "Disabled"
	FilteredSyncEnabled  FilteredSync = "Enabled"
)

// PossibleFilteredSyncValues returns the possible values for the FilteredSync const type.
func PossibleFilteredSyncValues() []FilteredSync {
	return []FilteredSync{
		FilteredSyncDisabled,
		FilteredSyncEnabled,
	}
}

// ToPtr returns a *FilteredSync pointing to the current value.
func (c FilteredSync) ToPtr() *FilteredSync {
	return &c
}

// KerberosArmoring - A flag to determine whether or not KerberosArmoring is enabled or disabled.
type KerberosArmoring string

const (
	KerberosArmoringDisabled KerberosArmoring = "Disabled"
	KerberosArmoringEnabled  KerberosArmoring = "Enabled"
)

// PossibleKerberosArmoringValues returns the possible values for the KerberosArmoring const type.
func PossibleKerberosArmoringValues() []KerberosArmoring {
	return []KerberosArmoring{
		KerberosArmoringDisabled,
		KerberosArmoringEnabled,
	}
}

// ToPtr returns a *KerberosArmoring pointing to the current value.
func (c KerberosArmoring) ToPtr() *KerberosArmoring {
	return &c
}

// KerberosRc4Encryption - A flag to determine whether or not KerberosRc4Encryption is enabled or disabled.
type KerberosRc4Encryption string

const (
	KerberosRc4EncryptionDisabled KerberosRc4Encryption = "Disabled"
	KerberosRc4EncryptionEnabled  KerberosRc4Encryption = "Enabled"
)

// PossibleKerberosRc4EncryptionValues returns the possible values for the KerberosRc4Encryption const type.
func PossibleKerberosRc4EncryptionValues() []KerberosRc4Encryption {
	return []KerberosRc4Encryption{
		KerberosRc4EncryptionDisabled,
		KerberosRc4EncryptionEnabled,
	}
}

// ToPtr returns a *KerberosRc4Encryption pointing to the current value.
func (c KerberosRc4Encryption) ToPtr() *KerberosRc4Encryption {
	return &c
}

// Ldaps - A flag to determine whether or not Secure LDAP is enabled or disabled.
type Ldaps string

const (
	LdapsDisabled Ldaps = "Disabled"
	LdapsEnabled  Ldaps = "Enabled"
)

// PossibleLdapsValues returns the possible values for the Ldaps const type.
func PossibleLdapsValues() []Ldaps {
	return []Ldaps{
		LdapsDisabled,
		LdapsEnabled,
	}
}

// ToPtr returns a *Ldaps pointing to the current value.
func (c Ldaps) ToPtr() *Ldaps {
	return &c
}

// NotifyDcAdmins - Should domain controller admins be notified
type NotifyDcAdmins string

const (
	NotifyDcAdminsDisabled NotifyDcAdmins = "Disabled"
	NotifyDcAdminsEnabled  NotifyDcAdmins = "Enabled"
)

// PossibleNotifyDcAdminsValues returns the possible values for the NotifyDcAdmins const type.
func PossibleNotifyDcAdminsValues() []NotifyDcAdmins {
	return []NotifyDcAdmins{
		NotifyDcAdminsDisabled,
		NotifyDcAdminsEnabled,
	}
}

// ToPtr returns a *NotifyDcAdmins pointing to the current value.
func (c NotifyDcAdmins) ToPtr() *NotifyDcAdmins {
	return &c
}

// NotifyGlobalAdmins - Should global admins be notified
type NotifyGlobalAdmins string

const (
	NotifyGlobalAdminsDisabled NotifyGlobalAdmins = "Disabled"
	NotifyGlobalAdminsEnabled  NotifyGlobalAdmins = "Enabled"
)

// PossibleNotifyGlobalAdminsValues returns the possible values for the NotifyGlobalAdmins const type.
func PossibleNotifyGlobalAdminsValues() []NotifyGlobalAdmins {
	return []NotifyGlobalAdmins{
		NotifyGlobalAdminsDisabled,
		NotifyGlobalAdminsEnabled,
	}
}

// ToPtr returns a *NotifyGlobalAdmins pointing to the current value.
func (c NotifyGlobalAdmins) ToPtr() *NotifyGlobalAdmins {
	return &c
}

// NtlmV1 - A flag to determine whether or not NtlmV1 is enabled or disabled.
type NtlmV1 string

const (
	NtlmV1Disabled NtlmV1 = "Disabled"
	NtlmV1Enabled  NtlmV1 = "Enabled"
)

// PossibleNtlmV1Values returns the possible values for the NtlmV1 const type.
func PossibleNtlmV1Values() []NtlmV1 {
	return []NtlmV1{
		NtlmV1Disabled,
		NtlmV1Enabled,
	}
}

// ToPtr returns a *NtlmV1 pointing to the current value.
func (c NtlmV1) ToPtr() *NtlmV1 {
	return &c
}

// Status - Status for individual validator after running diagnostics.
type Status string

const (
	StatusFailure Status = "Failure"
	StatusNone    Status = "None"
	StatusOK      Status = "OK"
	StatusRunning Status = "Running"
	StatusSkipped Status = "Skipped"
	StatusWarning Status = "Warning"
)

// PossibleStatusValues returns the possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{
		StatusFailure,
		StatusNone,
		StatusOK,
		StatusRunning,
		StatusSkipped,
		StatusWarning,
	}
}

// ToPtr returns a *Status pointing to the current value.
func (c Status) ToPtr() *Status {
	return &c
}

// SyncKerberosPasswords - A flag to determine whether or not SyncKerberosPasswords is enabled or disabled.
type SyncKerberosPasswords string

const (
	SyncKerberosPasswordsDisabled SyncKerberosPasswords = "Disabled"
	SyncKerberosPasswordsEnabled  SyncKerberosPasswords = "Enabled"
)

// PossibleSyncKerberosPasswordsValues returns the possible values for the SyncKerberosPasswords const type.
func PossibleSyncKerberosPasswordsValues() []SyncKerberosPasswords {
	return []SyncKerberosPasswords{
		SyncKerberosPasswordsDisabled,
		SyncKerberosPasswordsEnabled,
	}
}

// ToPtr returns a *SyncKerberosPasswords pointing to the current value.
func (c SyncKerberosPasswords) ToPtr() *SyncKerberosPasswords {
	return &c
}

// SyncNtlmPasswords - A flag to determine whether or not SyncNtlmPasswords is enabled or disabled.
type SyncNtlmPasswords string

const (
	SyncNtlmPasswordsDisabled SyncNtlmPasswords = "Disabled"
	SyncNtlmPasswordsEnabled  SyncNtlmPasswords = "Enabled"
)

// PossibleSyncNtlmPasswordsValues returns the possible values for the SyncNtlmPasswords const type.
func PossibleSyncNtlmPasswordsValues() []SyncNtlmPasswords {
	return []SyncNtlmPasswords{
		SyncNtlmPasswordsDisabled,
		SyncNtlmPasswordsEnabled,
	}
}

// ToPtr returns a *SyncNtlmPasswords pointing to the current value.
func (c SyncNtlmPasswords) ToPtr() *SyncNtlmPasswords {
	return &c
}

// SyncOnPremPasswords - A flag to determine whether or not SyncOnPremPasswords is enabled or disabled.
type SyncOnPremPasswords string

const (
	SyncOnPremPasswordsDisabled SyncOnPremPasswords = "Disabled"
	SyncOnPremPasswordsEnabled  SyncOnPremPasswords = "Enabled"
)

// PossibleSyncOnPremPasswordsValues returns the possible values for the SyncOnPremPasswords const type.
func PossibleSyncOnPremPasswordsValues() []SyncOnPremPasswords {
	return []SyncOnPremPasswords{
		SyncOnPremPasswordsDisabled,
		SyncOnPremPasswordsEnabled,
	}
}

// ToPtr returns a *SyncOnPremPasswords pointing to the current value.
func (c SyncOnPremPasswords) ToPtr() *SyncOnPremPasswords {
	return &c
}

// TLSV1 - A flag to determine whether or not TlsV1 is enabled or disabled.
type TLSV1 string

const (
	TLSV1Disabled TLSV1 = "Disabled"
	TLSV1Enabled  TLSV1 = "Enabled"
)

// PossibleTLSV1Values returns the possible values for the TLSV1 const type.
func PossibleTLSV1Values() []TLSV1 {
	return []TLSV1{
		TLSV1Disabled,
		TLSV1Enabled,
	}
}

// ToPtr returns a *TLSV1 pointing to the current value.
func (c TLSV1) ToPtr() *TLSV1 {
	return &c
}
