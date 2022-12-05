//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armwindowsesu

const (
	moduleName    = "armwindowsesu"
	moduleVersion = "v0.5.0"
)

// OsType - Type of OS for which the key is requested.
type OsType string

const (
	OsTypeWindows7            OsType = "Windows7"
	OsTypeWindowsServer2008   OsType = "WindowsServer2008"
	OsTypeWindowsServer2008R2 OsType = "WindowsServer2008R2"
)

// PossibleOsTypeValues returns the possible values for the OsType const type.
func PossibleOsTypeValues() []OsType {
	return []OsType{
		OsTypeWindows7,
		OsTypeWindowsServer2008,
		OsTypeWindowsServer2008R2,
	}
}

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

// SupportType - Type of support
type SupportType string

const (
	SupportTypePremiumAssurance      SupportType = "PremiumAssurance"
	SupportTypeSupplementalServicing SupportType = "SupplementalServicing"
)

// PossibleSupportTypeValues returns the possible values for the SupportType const type.
func PossibleSupportTypeValues() []SupportType {
	return []SupportType{
		SupportTypePremiumAssurance,
		SupportTypeSupplementalServicing,
	}
}