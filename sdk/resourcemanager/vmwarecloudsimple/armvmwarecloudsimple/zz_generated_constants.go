//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvmwarecloudsimple

const (
	moduleName    = "armvmwarecloudsimple"
	moduleVersion = "v1.0.0"
)

// AggregationType - Metric's aggregation type for e.g. (Average, Total)
type AggregationType string

const (
	AggregationTypeAverage AggregationType = "Average"
	AggregationTypeTotal   AggregationType = "Total"
)

// PossibleAggregationTypeValues returns the possible values for the AggregationType const type.
func PossibleAggregationTypeValues() []AggregationType {
	return []AggregationType{
		AggregationTypeAverage,
		AggregationTypeTotal,
	}
}

// CustomizationHostNameType - Type of host name
type CustomizationHostNameType string

const (
	CustomizationHostNameTypeCUSTOMNAME         CustomizationHostNameType = "CUSTOM_NAME"
	CustomizationHostNameTypeFIXED              CustomizationHostNameType = "FIXED"
	CustomizationHostNameTypePREFIXBASED        CustomizationHostNameType = "PREFIX_BASED"
	CustomizationHostNameTypeUSERDEFINED        CustomizationHostNameType = "USER_DEFINED"
	CustomizationHostNameTypeVIRTUALMACHINENAME CustomizationHostNameType = "VIRTUAL_MACHINE_NAME"
)

// PossibleCustomizationHostNameTypeValues returns the possible values for the CustomizationHostNameType const type.
func PossibleCustomizationHostNameTypeValues() []CustomizationHostNameType {
	return []CustomizationHostNameType{
		CustomizationHostNameTypeCUSTOMNAME,
		CustomizationHostNameTypeFIXED,
		CustomizationHostNameTypePREFIXBASED,
		CustomizationHostNameTypeUSERDEFINED,
		CustomizationHostNameTypeVIRTUALMACHINENAME,
	}
}

// CustomizationIPAddressType - Customization Specification ip type
type CustomizationIPAddressType string

const (
	CustomizationIPAddressTypeCUSTOM      CustomizationIPAddressType = "CUSTOM"
	CustomizationIPAddressTypeDHCPIP      CustomizationIPAddressType = "DHCP_IP"
	CustomizationIPAddressTypeFIXEDIP     CustomizationIPAddressType = "FIXED_IP"
	CustomizationIPAddressTypeUSERDEFINED CustomizationIPAddressType = "USER_DEFINED"
)

// PossibleCustomizationIPAddressTypeValues returns the possible values for the CustomizationIPAddressType const type.
func PossibleCustomizationIPAddressTypeValues() []CustomizationIPAddressType {
	return []CustomizationIPAddressType{
		CustomizationIPAddressTypeCUSTOM,
		CustomizationIPAddressTypeDHCPIP,
		CustomizationIPAddressTypeFIXEDIP,
		CustomizationIPAddressTypeUSERDEFINED,
	}
}

// CustomizationIdentityType - Identity type
type CustomizationIdentityType string

const (
	CustomizationIdentityTypeLINUX       CustomizationIdentityType = "LINUX"
	CustomizationIdentityTypeWINDOWS     CustomizationIdentityType = "WINDOWS"
	CustomizationIdentityTypeWINDOWSTEXT CustomizationIdentityType = "WINDOWS_TEXT"
)

// PossibleCustomizationIdentityTypeValues returns the possible values for the CustomizationIdentityType const type.
func PossibleCustomizationIdentityTypeValues() []CustomizationIdentityType {
	return []CustomizationIdentityType{
		CustomizationIdentityTypeLINUX,
		CustomizationIdentityTypeWINDOWS,
		CustomizationIdentityTypeWINDOWSTEXT,
	}
}

// CustomizationPolicyPropertiesType - The type of customization (Linux or Windows)
type CustomizationPolicyPropertiesType string

const (
	CustomizationPolicyPropertiesTypeLINUX   CustomizationPolicyPropertiesType = "LINUX"
	CustomizationPolicyPropertiesTypeWINDOWS CustomizationPolicyPropertiesType = "WINDOWS"
)

// PossibleCustomizationPolicyPropertiesTypeValues returns the possible values for the CustomizationPolicyPropertiesType const type.
func PossibleCustomizationPolicyPropertiesTypeValues() []CustomizationPolicyPropertiesType {
	return []CustomizationPolicyPropertiesType{
		CustomizationPolicyPropertiesTypeLINUX,
		CustomizationPolicyPropertiesTypeWINDOWS,
	}
}

// DiskIndependenceMode - Disk's independence mode type
type DiskIndependenceMode string

const (
	DiskIndependenceModePersistent               DiskIndependenceMode = "persistent"
	DiskIndependenceModeIndependentPersistent    DiskIndependenceMode = "independent_persistent"
	DiskIndependenceModeIndependentNonpersistent DiskIndependenceMode = "independent_nonpersistent"
)

// PossibleDiskIndependenceModeValues returns the possible values for the DiskIndependenceMode const type.
func PossibleDiskIndependenceModeValues() []DiskIndependenceMode {
	return []DiskIndependenceMode{
		DiskIndependenceModePersistent,
		DiskIndependenceModeIndependentPersistent,
		DiskIndependenceModeIndependentNonpersistent,
	}
}

// GuestOSNICCustomizationAllocation - IP address allocation method
type GuestOSNICCustomizationAllocation string

const (
	GuestOSNICCustomizationAllocationDynamic GuestOSNICCustomizationAllocation = "dynamic"
	GuestOSNICCustomizationAllocationStatic  GuestOSNICCustomizationAllocation = "static"
)

// PossibleGuestOSNICCustomizationAllocationValues returns the possible values for the GuestOSNICCustomizationAllocation const type.
func PossibleGuestOSNICCustomizationAllocationValues() []GuestOSNICCustomizationAllocation {
	return []GuestOSNICCustomizationAllocation{
		GuestOSNICCustomizationAllocationDynamic,
		GuestOSNICCustomizationAllocationStatic,
	}
}

// GuestOSType - The Guest OS type
type GuestOSType string

const (
	GuestOSTypeLinux   GuestOSType = "linux"
	GuestOSTypeWindows GuestOSType = "windows"
	GuestOSTypeOther   GuestOSType = "other"
)

// PossibleGuestOSTypeValues returns the possible values for the GuestOSType const type.
func PossibleGuestOSTypeValues() []GuestOSType {
	return []GuestOSType{
		GuestOSTypeLinux,
		GuestOSTypeWindows,
		GuestOSTypeOther,
	}
}

// NICType - NIC type
type NICType string

const (
	NICTypeE1000   NICType = "E1000"
	NICTypeE1000E  NICType = "E1000E"
	NICTypePCNET32 NICType = "PCNET32"
	NICTypeVMXNET  NICType = "VMXNET"
	NICTypeVMXNET2 NICType = "VMXNET2"
	NICTypeVMXNET3 NICType = "VMXNET3"
)

// PossibleNICTypeValues returns the possible values for the NICType const type.
func PossibleNICTypeValues() []NICType {
	return []NICType{
		NICTypeE1000,
		NICTypeE1000E,
		NICTypePCNET32,
		NICTypeVMXNET,
		NICTypeVMXNET2,
		NICTypeVMXNET3,
	}
}

// NodeStatus - Node status, indicates is private cloud set up on this node or not
type NodeStatus string

const (
	NodeStatusUnused NodeStatus = "unused"
	NodeStatusUsed   NodeStatus = "used"
)

// PossibleNodeStatusValues returns the possible values for the NodeStatus const type.
func PossibleNodeStatusValues() []NodeStatus {
	return []NodeStatus{
		NodeStatusUnused,
		NodeStatusUsed,
	}
}

// OnboardingStatus - indicates whether account onboarded or not in a given region
type OnboardingStatus string

const (
	OnboardingStatusNotOnBoarded     OnboardingStatus = "notOnBoarded"
	OnboardingStatusOnBoarded        OnboardingStatus = "onBoarded"
	OnboardingStatusOnBoardingFailed OnboardingStatus = "onBoardingFailed"
	OnboardingStatusOnBoarding       OnboardingStatus = "onBoarding"
)

// PossibleOnboardingStatusValues returns the possible values for the OnboardingStatus const type.
func PossibleOnboardingStatusValues() []OnboardingStatus {
	return []OnboardingStatus{
		OnboardingStatusNotOnBoarded,
		OnboardingStatusOnBoarded,
		OnboardingStatusOnBoardingFailed,
		OnboardingStatusOnBoarding,
	}
}

// OperationOrigin - The origin of operation
type OperationOrigin string

const (
	OperationOriginUser       OperationOrigin = "user"
	OperationOriginSystem     OperationOrigin = "system"
	OperationOriginUserSystem OperationOrigin = "user,system"
)

// PossibleOperationOriginValues returns the possible values for the OperationOrigin const type.
func PossibleOperationOriginValues() []OperationOrigin {
	return []OperationOrigin{
		OperationOriginUser,
		OperationOriginSystem,
		OperationOriginUserSystem,
	}
}

// StopMode - mode indicates a type of stop operation - reboot, suspend, shutdown or power-off
type StopMode string

const (
	StopModeReboot   StopMode = "reboot"
	StopModeSuspend  StopMode = "suspend"
	StopModeShutdown StopMode = "shutdown"
	StopModePoweroff StopMode = "poweroff"
)

// PossibleStopModeValues returns the possible values for the StopMode const type.
func PossibleStopModeValues() []StopMode {
	return []StopMode{
		StopModeReboot,
		StopModeSuspend,
		StopModeShutdown,
		StopModePoweroff,
	}
}

// UsageCount - The usages' unit
type UsageCount string

const (
	UsageCountCount          UsageCount = "Count"
	UsageCountBytes          UsageCount = "Bytes"
	UsageCountSeconds        UsageCount = "Seconds"
	UsageCountPercent        UsageCount = "Percent"
	UsageCountCountPerSecond UsageCount = "CountPerSecond"
	UsageCountBytesPerSecond UsageCount = "BytesPerSecond"
)

// PossibleUsageCountValues returns the possible values for the UsageCount const type.
func PossibleUsageCountValues() []UsageCount {
	return []UsageCount{
		UsageCountCount,
		UsageCountBytes,
		UsageCountSeconds,
		UsageCountPercent,
		UsageCountCountPerSecond,
		UsageCountBytesPerSecond,
	}
}

// VirtualMachineStatus - The status of Virtual machine
type VirtualMachineStatus string

const (
	VirtualMachineStatusRunning      VirtualMachineStatus = "running"
	VirtualMachineStatusSuspended    VirtualMachineStatus = "suspended"
	VirtualMachineStatusPoweredoff   VirtualMachineStatus = "poweredoff"
	VirtualMachineStatusUpdating     VirtualMachineStatus = "updating"
	VirtualMachineStatusDeallocating VirtualMachineStatus = "deallocating"
	VirtualMachineStatusDeleting     VirtualMachineStatus = "deleting"
)

// PossibleVirtualMachineStatusValues returns the possible values for the VirtualMachineStatus const type.
func PossibleVirtualMachineStatusValues() []VirtualMachineStatus {
	return []VirtualMachineStatus{
		VirtualMachineStatusRunning,
		VirtualMachineStatusSuspended,
		VirtualMachineStatusPoweredoff,
		VirtualMachineStatusUpdating,
		VirtualMachineStatusDeallocating,
		VirtualMachineStatusDeleting,
	}
}