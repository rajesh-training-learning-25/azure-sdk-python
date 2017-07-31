package armservicemap

import (
	 original "github.com/Azure/azure-sdk-for-go/service/service-map/management/2015-11-01-preview/arm-service-map"
)

type (
	 MachinesClient = original.MachinesClient
	 Accuracy = original.Accuracy
	 Bitness = original.Bitness
	 ConnectionFailureState = original.ConnectionFailureState
	 HypervisorType = original.HypervisorType
	 MachineRebootStatus = original.MachineRebootStatus
	 MonitoringState = original.MonitoringState
	 OperatingSystemFamily = original.OperatingSystemFamily
	 ProcessRole = original.ProcessRole
	 VirtualizationState = original.VirtualizationState
	 VirtualMachineType = original.VirtualMachineType
	 Acceptor = original.Acceptor
	 AcceptorProperties = original.AcceptorProperties
	 AgentConfiguration = original.AgentConfiguration
	 ClientGroup = original.ClientGroup
	 ClientGroupProperties = original.ClientGroupProperties
	 ClientGroupMember = original.ClientGroupMember
	 ClientGroupMemberProperties = original.ClientGroupMemberProperties
	 ClientGroupMembersCollection = original.ClientGroupMembersCollection
	 ClientGroupMembersCount = original.ClientGroupMembersCount
	 Connection = original.Connection
	 ConnectionCollection = original.ConnectionCollection
	 ConnectionProperties = original.ConnectionProperties
	 CoreResource = original.CoreResource
	 Error = original.Error
	 ErrorResponse = original.ErrorResponse
	 HypervisorConfiguration = original.HypervisorConfiguration
	 Ipv4NetworkInterface = original.Ipv4NetworkInterface
	 Ipv6NetworkInterface = original.Ipv6NetworkInterface
	 Liveness = original.Liveness
	 Machine = original.Machine
	 MachineProperties = original.MachineProperties
	 MachineCollection = original.MachineCollection
	 MachineCountsByOperatingSystem = original.MachineCountsByOperatingSystem
	 MachineGroup = original.MachineGroup
	 MachineGroupProperties = original.MachineGroupProperties
	 MachineGroupCollection = original.MachineGroupCollection
	 MachineGroupMapRequest = original.MachineGroupMapRequest
	 MachineReference = original.MachineReference
	 MachineReferenceWithHints = original.MachineReferenceWithHints
	 MachineReferenceWithHintsProperties = original.MachineReferenceWithHintsProperties
	 MachineResourcesConfiguration = original.MachineResourcesConfiguration
	 MachinesSummary = original.MachinesSummary
	 MachinesSummaryProperties = original.MachinesSummaryProperties
	 Map = original.Map
	 MapEdges = original.MapEdges
	 MapNodes = original.MapNodes
	 MapRequest = original.MapRequest
	 MapResponse = original.MapResponse
	 NetworkConfiguration = original.NetworkConfiguration
	 OperatingSystemConfiguration = original.OperatingSystemConfiguration
	 Port = original.Port
	 PortProperties = original.PortProperties
	 PortCollection = original.PortCollection
	 PortReference = original.PortReference
	 PortReferenceProperties = original.PortReferenceProperties
	 Process = original.Process
	 ProcessProperties = original.ProcessProperties
	 ProcessCollection = original.ProcessCollection
	 ProcessDetails = original.ProcessDetails
	 ProcessReference = original.ProcessReference
	 ProcessReferenceProperties = original.ProcessReferenceProperties
	 ProcessUser = original.ProcessUser
	 Relationship = original.Relationship
	 RelationshipProperties = original.RelationshipProperties
	 Resource = original.Resource
	 ResourceReference = original.ResourceReference
	 SingleMachineDependencyMapRequest = original.SingleMachineDependencyMapRequest
	 Summary = original.Summary
	 SummaryProperties = original.SummaryProperties
	 Timezone = original.Timezone
	 VirtualMachineConfiguration = original.VirtualMachineConfiguration
	 PortsClient = original.PortsClient
	 ProcessesClient = original.ProcessesClient
	 ManagementClient = original.ManagementClient
	 ClientGroupsClient = original.ClientGroupsClient
	 MachineGroupsClient = original.MachineGroupsClient
	 MapsClient = original.MapsClient
	 SummariesClient = original.SummariesClient
)

const (
	 Actual = original.Actual
	 Estimated = original.Estimated
	 SixFourbit = original.SixFourbit
	 ThreeTwobit = original.ThreeTwobit
	 Failed = original.Failed
	 Mixed = original.Mixed
	 Ok = original.Ok
	 Hyperv = original.Hyperv
	 Unknown = original.Unknown
	 MachineRebootStatusNotRebooted = original.MachineRebootStatusNotRebooted
	 MachineRebootStatusRebooted = original.MachineRebootStatusRebooted
	 MachineRebootStatusUnknown = original.MachineRebootStatusUnknown
	 Discovered = original.Discovered
	 Monitored = original.Monitored
	 OperatingSystemFamilyAix = original.OperatingSystemFamilyAix
	 OperatingSystemFamilyLinux = original.OperatingSystemFamilyLinux
	 OperatingSystemFamilySolaris = original.OperatingSystemFamilySolaris
	 OperatingSystemFamilyUnknown = original.OperatingSystemFamilyUnknown
	 OperatingSystemFamilyWindows = original.OperatingSystemFamilyWindows
	 AppServer = original.AppServer
	 DatabaseServer = original.DatabaseServer
	 LdapServer = original.LdapServer
	 SmbServer = original.SmbServer
	 WebServer = original.WebServer
	 VirtualizationStateHypervisor = original.VirtualizationStateHypervisor
	 VirtualizationStatePhysical = original.VirtualizationStatePhysical
	 VirtualizationStateUnknown = original.VirtualizationStateUnknown
	 VirtualizationStateVirtual = original.VirtualizationStateVirtual
	 VirtualMachineTypeHyperv = original.VirtualMachineTypeHyperv
	 VirtualMachineTypeLdom = original.VirtualMachineTypeLdom
	 VirtualMachineTypeLpar = original.VirtualMachineTypeLpar
	 VirtualMachineTypeUnknown = original.VirtualMachineTypeUnknown
	 VirtualMachineTypeVirtualPc = original.VirtualMachineTypeVirtualPc
	 VirtualMachineTypeVmware = original.VirtualMachineTypeVmware
	 VirtualMachineTypeXen = original.VirtualMachineTypeXen
	 DefaultBaseURI = original.DefaultBaseURI
)
