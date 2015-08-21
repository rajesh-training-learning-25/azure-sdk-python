package compute

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.11.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/azure/go-autorest/autorest"
	"github.com/azure/go-autorest/autorest/date"
)

type CachingTypes string

const (
	None      CachingTypes = "None"
	ReadOnly  CachingTypes = "ReadOnly"
	ReadWrite CachingTypes = "ReadWrite"
)

type ComponentNames string

const (
	MicrosoftWindowsShellSetup ComponentNames = "Microsoft-Windows-Shell-Setup"
)

type ComputeOperationStatus string

const (
	ComputeOperationStatusFailed     ComputeOperationStatus = "Failed"
	ComputeOperationStatusInProgress ComputeOperationStatus = "InProgress"
	ComputeOperationStatusPreempted  ComputeOperationStatus = "Preempted"
	ComputeOperationStatusSucceeded  ComputeOperationStatus = "Succeeded"
)

type DiskCreateOptionTypes string

const (
	Attach    DiskCreateOptionTypes = "attach"
	Empty     DiskCreateOptionTypes = "empty"
	FromImage DiskCreateOptionTypes = "fromImage"
)

type OperatingSystemTypes string

const (
	Linux   OperatingSystemTypes = "Linux"
	Windows OperatingSystemTypes = "Windows"
)

type OperationStatus string

const (
	OperationStatusFailed     OperationStatus = "Failed"
	OperationStatusInProgress OperationStatus = "InProgress"
	OperationStatusSucceeded  OperationStatus = "Succeeded"
)

type PassNames string

const (
	OobeSystem PassNames = "oobeSystem"
)

type ProtocolTypes string

const (
	Http  ProtocolTypes = "Http"
	Https ProtocolTypes = "Https"
)

type SettingNames string

const (
	AutoLogon          SettingNames = "AutoLogon"
	FirstLogonCommands SettingNames = "FirstLogonCommands"
)

type StatusLevelTypes string

const (
	Error   StatusLevelTypes = "Error"
	Info    StatusLevelTypes = "Info"
	Warning StatusLevelTypes = "Warning"
)

type UsageUnit string

const (
	Count UsageUnit = "Count"
)

type VirtualMachineSizeTypes string

const (
	BasicA0    VirtualMachineSizeTypes = "Basic_A0"
	BasicA1    VirtualMachineSizeTypes = "Basic_A1"
	BasicA2    VirtualMachineSizeTypes = "Basic_A2"
	BasicA3    VirtualMachineSizeTypes = "Basic_A3"
	BasicA4    VirtualMachineSizeTypes = "Basic_A4"
	StandardA0 VirtualMachineSizeTypes = "Standard_A0"
	StandardA1 VirtualMachineSizeTypes = "Standard_A1"
	StandardA2 VirtualMachineSizeTypes = "Standard_A2"
	StandardA3 VirtualMachineSizeTypes = "Standard_A3"
	StandardA4 VirtualMachineSizeTypes = "Standard_A4"
	StandardA5 VirtualMachineSizeTypes = "Standard_A5"
	StandardA6 VirtualMachineSizeTypes = "Standard_A6"
	StandardA7 VirtualMachineSizeTypes = "Standard_A7"
	StandardA8 VirtualMachineSizeTypes = "Standard_A8"
	StandardA9 VirtualMachineSizeTypes = "Standard_A9"
	StandardG1 VirtualMachineSizeTypes = "Standard_G1"
	StandardG2 VirtualMachineSizeTypes = "Standard_G2"
	StandardG3 VirtualMachineSizeTypes = "Standard_G3"
	StandardG4 VirtualMachineSizeTypes = "Standard_G4"
	StandardG5 VirtualMachineSizeTypes = "Standard_G5"
)

// Gets or sets additional XML formatted information that can be included in
// the Unattend.xml file, which is used by Windows Setup. Contents are
// defined by setting name, component name, and the pass in which the content
// is a applied.
type AdditionalUnattendContent struct {
	PassName      PassNames      `json:"passName,omitempty"`
	ComponentName ComponentNames `json:"componentName,omitempty"`
	SettingName   SettingNames   `json:"settingName,omitempty"`
	Content       string         `json:"content,omitempty"`
}

// Api error base.
type ApiErrorBase struct {
	Code    string `json:"code,omitempty"`
	Target  string `json:"target,omitempty"`
	Message string `json:"message,omitempty"`
}

// Create or update Availability Set parameters.
type AvailabilitySet struct {
	autorest.Response `json:"-"`
	Id                string            `json:"id,omitempty"`
	Name              string            `json:"name,omitempty"`
	Type              string            `json:"type,omitempty"`
	Location          string            `json:"location,omitempty"`
	Tags              map[string]string `json:"tags,omitempty"`
	Properties        struct {
		PlatformUpdateDomainCount int                  `json:"platformUpdateDomainCount,omitempty"`
		PlatformFaultDomainCount  int                  `json:"platformFaultDomainCount,omitempty"`
		VirtualMachines           []SubResource        `json:"virtualMachines,omitempty"`
		Statuses                  []InstanceViewStatus `json:"statuses,omitempty"`
	} `json:"properties,omitempty"`
}

// The List Availability Set operation response.
type AvailabilitySetListResult struct {
	autorest.Response `json:"-"`
	Value             []AvailabilitySet `json:"value,omitempty"`
}

// The Compute service response for long-running operations.
type ComputeLongRunningOperationResult struct {
	autorest.Response `json:"-"`
	OperationId       string                 `json:"operationId,omitempty"`
	Status            ComputeOperationStatus `json:"status,omitempty"`
	StartTime         date.Time              `json:"startTime,omitempty"`
	EndTime           date.Time              `json:"endTime,omitempty"`
	Properties        struct {
		Output map[string]string `json:"output,omitempty"`
	} `json:"properties,omitempty"`
	Error struct {
		Details    []ApiErrorBase `json:"details,omitempty"`
		Innererror struct {
			Exceptiontype string `json:"exceptiontype,omitempty"`
			Errordetail   string `json:"errordetail,omitempty"`
		} `json:"innererror,omitempty"`
		Code    string `json:"code,omitempty"`
		Target  string `json:"target,omitempty"`
		Message string `json:"message,omitempty"`
	} `json:"error,omitempty"`
}

// Describes a data disk.
type DataDisk struct {
	Lun        int    `json:"lun,omitempty"`
	DiskSizeGB int    `json:"diskSizeGB,omitempty"`
	Name       string `json:"name,omitempty"`
	Vhd        struct {
		Uri string `json:"uri,omitempty"`
	} `json:"vhd,omitempty"`
	Image struct {
		Uri string `json:"uri,omitempty"`
	} `json:"image,omitempty"`
	Caching      CachingTypes          `json:"caching,omitempty"`
	CreateOption DiskCreateOptionTypes `json:"createOption,omitempty"`
}

// Contains the data disk images information.
type DataDiskImage struct {
	Lun int `json:"lun,omitempty"`
}

// The instance view of the disk.
type DiskInstanceView struct {
	Name     string               `json:"name,omitempty"`
	Statuses []InstanceViewStatus `json:"statuses,omitempty"`
}

// Instance view status.
type InstanceViewStatus struct {
	Code          string           `json:"code,omitempty"`
	Level         StatusLevelTypes `json:"level,omitempty"`
	DisplayStatus string           `json:"displayStatus,omitempty"`
	Message       string           `json:"message,omitempty"`
	Time          date.Time        `json:"time,omitempty"`
}

// The List Usages operation response.
type ListUsagesResult struct {
	autorest.Response `json:"-"`
	Value             []Usage `json:"value,omitempty"`
}

// Describes a network interface reference.
type NetworkInterfaceReference struct {
	Id         string `json:"id,omitempty"`
	Properties struct {
		Primary bool `json:"primary,omitempty"`
	} `json:"properties,omitempty"`
}

// Contains information about SSH certificate public key and the path on the
// Linux VM where the public key is placed.
type SshPublicKey struct {
	Path    string `json:"path,omitempty"`
	KeyData string `json:"keyData,omitempty"`
}

type SubResource struct {
	Id string `json:"id,omitempty"`
}

// Describes Compute Resource Usage.
type Usage struct {
	Unit         UsageUnit `json:"unit,omitempty"`
	CurrentValue int       `json:"currentValue,omitempty"`
	Limit        int32     `json:"limit,omitempty"`
	Name         struct {
		Value          string `json:"value,omitempty"`
		LocalizedValue string `json:"localizedValue,omitempty"`
	} `json:"name,omitempty"`
}

// Describes a single certificate reference in a Key Vault, and where the
// certificate should reside on the VM.
type VaultCertificate struct {
	CertificateUrl   string `json:"certificateUrl,omitempty"`
	CertificateStore string `json:"certificateStore,omitempty"`
}

// Describes a set of certificates which are all in the same Key Vault.
type VaultSecretGroup struct {
	SourceVault struct {
		Id string `json:"id,omitempty"`
	} `json:"sourceVault,omitempty"`
	VaultCertificates []VaultCertificate `json:"vaultCertificates,omitempty"`
}

// Describes a Virtual Machine.
type VirtualMachine struct {
	autorest.Response `json:"-"`
	Id                string            `json:"id,omitempty"`
	Name              string            `json:"name,omitempty"`
	Type              string            `json:"type,omitempty"`
	Location          string            `json:"location,omitempty"`
	Tags              map[string]string `json:"tags,omitempty"`
	Plan              struct {
		Name          string `json:"name,omitempty"`
		Publisher     string `json:"publisher,omitempty"`
		Product       string `json:"product,omitempty"`
		PromotionCode string `json:"promotionCode,omitempty"`
	} `json:"plan,omitempty"`
	Properties struct {
		HardwareProfile struct {
			VmSize VirtualMachineSizeTypes `json:"vmSize,omitempty"`
		} `json:"hardwareProfile,omitempty"`
		StorageProfile struct {
			ImageReference struct {
				Publisher string `json:"publisher,omitempty"`
				Offer     string `json:"offer,omitempty"`
				Sku       string `json:"sku,omitempty"`
				Version   string `json:"version,omitempty"`
			} `json:"imageReference,omitempty"`
			OsDisk struct {
				OsType OperatingSystemTypes `json:"osType,omitempty"`
				Name   string               `json:"name,omitempty"`
				Vhd    struct {
					Uri string `json:"uri,omitempty"`
				} `json:"vhd,omitempty"`
				Image struct {
					Uri string `json:"uri,omitempty"`
				} `json:"image,omitempty"`
				Caching      CachingTypes          `json:"caching,omitempty"`
				CreateOption DiskCreateOptionTypes `json:"createOption,omitempty"`
			} `json:"osDisk,omitempty"`
			DataDisks []DataDisk `json:"dataDisks,omitempty"`
		} `json:"storageProfile,omitempty"`
		OsProfile struct {
			ComputerName         string `json:"computerName,omitempty"`
			AdminUsername        string `json:"adminUsername,omitempty"`
			AdminPassword        string `json:"adminPassword,omitempty"`
			CustomData           string `json:"customData,omitempty"`
			WindowsConfiguration struct {
				ProvisionVMAgent          bool                        `json:"provisionVMAgent,omitempty"`
				EnableAutomaticUpdates    bool                        `json:"enableAutomaticUpdates,omitempty"`
				TimeZone                  string                      `json:"timeZone,omitempty"`
				AdditionalUnattendContent []AdditionalUnattendContent `json:"additionalUnattendContent,omitempty"`
				WinRM                     struct {
					Listeners []WinRMListener `json:"listeners,omitempty"`
				} `json:"winRM,omitempty"`
			} `json:"windowsConfiguration,omitempty"`
			LinuxConfiguration struct {
				DisablePasswordAuthentication bool `json:"disablePasswordAuthentication,omitempty"`
				Ssh                           struct {
					PublicKeys []SshPublicKey `json:"publicKeys,omitempty"`
				} `json:"ssh,omitempty"`
			} `json:"linuxConfiguration,omitempty"`
			Secrets []VaultSecretGroup `json:"secrets,omitempty"`
		} `json:"osProfile,omitempty"`
		NetworkProfile struct {
			NetworkInterfaces []NetworkInterfaceReference `json:"networkInterfaces,omitempty"`
		} `json:"networkProfile,omitempty"`
		AvailabilitySet struct {
			Id string `json:"id,omitempty"`
		} `json:"availabilitySet,omitempty"`
		ProvisioningState string `json:"provisioningState,omitempty"`
		InstanceView      struct {
			PlatformUpdateDomain int    `json:"platformUpdateDomain,omitempty"`
			PlatformFaultDomain  int    `json:"platformFaultDomain,omitempty"`
			RdpThumbPrint        string `json:"rdpThumbPrint,omitempty"`
			VmAgent              struct {
				VmAgentVersion    string                                       `json:"vmAgentVersion,omitempty"`
				ExtensionHandlers []VirtualMachineExtensionHandlerInstanceView `json:"extensionHandlers,omitempty"`
				Statuses          []InstanceViewStatus                         `json:"statuses,omitempty"`
			} `json:"vmAgent,omitempty"`
			Disks      []DiskInstanceView                    `json:"disks,omitempty"`
			Extensions []VirtualMachineExtensionInstanceView `json:"extensions,omitempty"`
			Statuses   []InstanceViewStatus                  `json:"statuses,omitempty"`
		} `json:"instanceView,omitempty"`
	} `json:"properties,omitempty"`
	Resources []VirtualMachineExtension `json:"resources,omitempty"`
}

// Capture Virtual Machine parameters.
type VirtualMachineCaptureParameters struct {
	VhdPrefix                string `json:"vhdPrefix,omitempty"`
	DestinationContainerName string `json:"destinationContainerName,omitempty"`
	OverwriteVhds            bool   `json:"overwriteVhds,omitempty"`
}

// Describes a Virtual Machine Extension.
type VirtualMachineExtension struct {
	autorest.Response `json:"-"`
	Id                string            `json:"id,omitempty"`
	Name              string            `json:"name,omitempty"`
	Type              string            `json:"type,omitempty"`
	Location          string            `json:"location,omitempty"`
	Tags              map[string]string `json:"tags,omitempty"`
	Properties        struct {
		Publisher               string            `json:"publisher,omitempty"`
		Type                    string            `json:"type,omitempty"`
		TypeHandlerVersion      string            `json:"typeHandlerVersion,omitempty"`
		AutoUpgradeMinorVersion bool              `json:"autoUpgradeMinorVersion,omitempty"`
		Settings                map[string]string `json:"settings,omitempty"`
		ProtectedSettings       map[string]string `json:"protectedSettings,omitempty"`
		ProvisioningState       string            `json:"provisioningState,omitempty"`
		InstanceView            struct {
			Name               string               `json:"name,omitempty"`
			Type               string               `json:"type,omitempty"`
			TypeHandlerVersion string               `json:"typeHandlerVersion,omitempty"`
			Substatuses        []InstanceViewStatus `json:"substatuses,omitempty"`
			Statuses           []InstanceViewStatus `json:"statuses,omitempty"`
		} `json:"instanceView,omitempty"`
	} `json:"properties,omitempty"`
}

// The instance view of a virtual machine extension handler.
type VirtualMachineExtensionHandlerInstanceView struct {
	Type               string `json:"type,omitempty"`
	TypeHandlerVersion string `json:"typeHandlerVersion,omitempty"`
	Status             struct {
		Code          string           `json:"code,omitempty"`
		Level         StatusLevelTypes `json:"level,omitempty"`
		DisplayStatus string           `json:"displayStatus,omitempty"`
		Message       string           `json:"message,omitempty"`
		Time          date.Time        `json:"time,omitempty"`
	} `json:"status,omitempty"`
}

// Describes a Virtual Machine Extension Image.
type VirtualMachineExtensionImage struct {
	autorest.Response `json:"-"`
	Id                string `json:"id,omitempty"`
	Properties        struct {
		OperatingSystem            string `json:"operatingSystem,omitempty"`
		ComputeRole                string `json:"computeRole,omitempty"`
		HandlerSchema              string `json:"handlerSchema,omitempty"`
		VmScaleSetEnabled          bool   `json:"vmScaleSetEnabled,omitempty"`
		SupportsMultipleExtensions bool   `json:"supportsMultipleExtensions,omitempty"`
	} `json:"properties,omitempty"`
	Name     string            `json:"name,omitempty"`
	Location string            `json:"location,omitempty"`
	Tags     map[string]string `json:"tags,omitempty"`
}

// The instance view of a virtual machine extension.
type VirtualMachineExtensionInstanceView struct {
	Name               string               `json:"name,omitempty"`
	Type               string               `json:"type,omitempty"`
	TypeHandlerVersion string               `json:"typeHandlerVersion,omitempty"`
	Substatuses        []InstanceViewStatus `json:"substatuses,omitempty"`
	Statuses           []InstanceViewStatus `json:"statuses,omitempty"`
}

// Describes a Virtual Machine Image.
type VirtualMachineImage struct {
	autorest.Response `json:"-"`
	Id                string `json:"id,omitempty"`
	Properties        struct {
		Plan struct {
			Publisher string `json:"publisher,omitempty"`
			Name      string `json:"name,omitempty"`
			Product   string `json:"product,omitempty"`
		} `json:"plan,omitempty"`
		OsDiskImage struct {
			OperatingSystem OperatingSystemTypes `json:"operatingSystem,omitempty"`
		} `json:"osDiskImage,omitempty"`
		DataDiskImages []DataDiskImage `json:"dataDiskImages,omitempty"`
	} `json:"properties,omitempty"`
	Name     string            `json:"name,omitempty"`
	Location string            `json:"location,omitempty"`
	Tags     map[string]string `json:"tags,omitempty"`
}

// Virtual machine image resource information.
type VirtualMachineImageResource struct {
	Id       string            `json:"id,omitempty"`
	Name     string            `json:"name,omitempty"`
	Location string            `json:"location,omitempty"`
	Tags     map[string]string `json:"tags,omitempty"`
}

type VirtualMachineImageResourceList struct {
	autorest.Response `json:"-"`
	Value             []VirtualMachineImageResource `json:"value,omitempty"`
}

// The List Virtual Machine operation response.
type VirtualMachineListResult struct {
	autorest.Response `json:"-"`
	Value             []VirtualMachine `json:"value,omitempty"`
	NextLink          string           `json:"nextLink,omitempty"`
}

// Describes the properties of a VM size.
type VirtualMachineSize struct {
	Name                 string `json:"name,omitempty"`
	NumberOfCores        int    `json:"numberOfCores,omitempty"`
	OsDiskSizeInMB       int    `json:"osDiskSizeInMB,omitempty"`
	ResourceDiskSizeInMB int    `json:"resourceDiskSizeInMB,omitempty"`
	MemoryInMB           int    `json:"memoryInMB,omitempty"`
	MaxDataDiskCount     int    `json:"maxDataDiskCount,omitempty"`
}

// The List Virtual Machine operation response.
type VirtualMachineSizeListResult struct {
	autorest.Response `json:"-"`
	Value             []VirtualMachineSize `json:"value,omitempty"`
}

// Describes Protocol and thumbprint of Windows Remote Management listener
type WinRMListener struct {
	Protocol       ProtocolTypes `json:"protocol,omitempty"`
	CertificateUrl string        `json:"certificateUrl,omitempty"`
}
