// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package servicemap

import original "github.com/Azure/azure-sdk-for-go/services/operationalinsights/mgmt/2015-11-01-preview/servicemap"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}

type ClientGroupsClient = original.ClientGroupsClient

func NewClientGroupsClient(subscriptionID string) ClientGroupsClient {
	return original.NewClientGroupsClient(subscriptionID)
}
func NewClientGroupsClientWithBaseURI(baseURI string, subscriptionID string) ClientGroupsClient {
	return original.NewClientGroupsClientWithBaseURI(baseURI, subscriptionID)
}

type MachineGroupsClient = original.MachineGroupsClient

func NewMachineGroupsClient(subscriptionID string) MachineGroupsClient {
	return original.NewMachineGroupsClient(subscriptionID)
}
func NewMachineGroupsClientWithBaseURI(baseURI string, subscriptionID string) MachineGroupsClient {
	return original.NewMachineGroupsClientWithBaseURI(baseURI, subscriptionID)
}

type MachinesClient = original.MachinesClient

func NewMachinesClient(subscriptionID string) MachinesClient {
	return original.NewMachinesClient(subscriptionID)
}
func NewMachinesClientWithBaseURI(baseURI string, subscriptionID string) MachinesClient {
	return original.NewMachinesClientWithBaseURI(baseURI, subscriptionID)
}

type MapsClient = original.MapsClient

func NewMapsClient(subscriptionID string) MapsClient {
	return original.NewMapsClient(subscriptionID)
}
func NewMapsClientWithBaseURI(baseURI string, subscriptionID string) MapsClient {
	return original.NewMapsClientWithBaseURI(baseURI, subscriptionID)
}

type Accuracy = original.Accuracy

const (
	Actual    Accuracy = original.Actual
	Estimated Accuracy = original.Estimated
)

func PossibleAccuracyValues() [2]Accuracy {
	return original.PossibleAccuracyValues()
}

type AzureCloudServiceRoleType = original.AzureCloudServiceRoleType

const (
	Unknown AzureCloudServiceRoleType = original.Unknown
	Web     AzureCloudServiceRoleType = original.Web
	Worker  AzureCloudServiceRoleType = original.Worker
)

func PossibleAzureCloudServiceRoleTypeValues() [3]AzureCloudServiceRoleType {
	return original.PossibleAzureCloudServiceRoleTypeValues()
}

type Bitness = original.Bitness

const (
	SixFourbit  Bitness = original.SixFourbit
	ThreeTwobit Bitness = original.ThreeTwobit
)

func PossibleBitnessValues() [2]Bitness {
	return original.PossibleBitnessValues()
}

type ConnectionFailureState = original.ConnectionFailureState

const (
	Failed ConnectionFailureState = original.Failed
	Mixed  ConnectionFailureState = original.Mixed
	Ok     ConnectionFailureState = original.Ok
)

func PossibleConnectionFailureStateValues() [3]ConnectionFailureState {
	return original.PossibleConnectionFailureStateValues()
}

type HypervisorType = original.HypervisorType

const (
	HypervisorTypeHyperv  HypervisorType = original.HypervisorTypeHyperv
	HypervisorTypeUnknown HypervisorType = original.HypervisorTypeUnknown
)

func PossibleHypervisorTypeValues() [2]HypervisorType {
	return original.PossibleHypervisorTypeValues()
}

type Kind = original.Kind

const (
	KindRefclientgroup      Kind = original.KindRefclientgroup
	KindRefmachine          Kind = original.KindRefmachine
	KindRefmachinewithhints Kind = original.KindRefmachinewithhints
	KindRefport             Kind = original.KindRefport
	KindRefprocess          Kind = original.KindRefprocess
	KindResourceReference   Kind = original.KindResourceReference
)

func PossibleKindValues() [6]Kind {
	return original.PossibleKindValues()
}

type KindBasicCoreResource = original.KindBasicCoreResource

const (
	KindClientGroup  KindBasicCoreResource = original.KindClientGroup
	KindCoreResource KindBasicCoreResource = original.KindCoreResource
	KindMachine      KindBasicCoreResource = original.KindMachine
	KindMachineGroup KindBasicCoreResource = original.KindMachineGroup
	KindPort         KindBasicCoreResource = original.KindPort
	KindProcess      KindBasicCoreResource = original.KindProcess
)

func PossibleKindBasicCoreResourceValues() [6]KindBasicCoreResource {
	return original.PossibleKindBasicCoreResourceValues()
}

type KindBasicHostingConfiguration = original.KindBasicHostingConfiguration

const (
	KindHostingConfiguration KindBasicHostingConfiguration = original.KindHostingConfiguration
	KindProviderazure        KindBasicHostingConfiguration = original.KindProviderazure
)

func PossibleKindBasicHostingConfigurationValues() [2]KindBasicHostingConfiguration {
	return original.PossibleKindBasicHostingConfigurationValues()
}

type KindBasicMapRequest = original.KindBasicMapRequest

const (
	KindMapmachineGroupDependency  KindBasicMapRequest = original.KindMapmachineGroupDependency
	KindMapmachineListDependency   KindBasicMapRequest = original.KindMapmachineListDependency
	KindMapRequest                 KindBasicMapRequest = original.KindMapRequest
	KindMapsingleMachineDependency KindBasicMapRequest = original.KindMapsingleMachineDependency
	KindMultipleMachinesMapRequest KindBasicMapRequest = original.KindMultipleMachinesMapRequest
)

func PossibleKindBasicMapRequestValues() [5]KindBasicMapRequest {
	return original.PossibleKindBasicMapRequestValues()
}

type KindBasicProcessHostingConfiguration = original.KindBasicProcessHostingConfiguration

const (
	KindBasicProcessHostingConfigurationKindProcessHostingConfiguration KindBasicProcessHostingConfiguration = original.KindBasicProcessHostingConfigurationKindProcessHostingConfiguration
	KindBasicProcessHostingConfigurationKindProviderazure               KindBasicProcessHostingConfiguration = original.KindBasicProcessHostingConfigurationKindProviderazure
)

func PossibleKindBasicProcessHostingConfigurationValues() [2]KindBasicProcessHostingConfiguration {
	return original.PossibleKindBasicProcessHostingConfigurationValues()
}

type KindBasicRelationship = original.KindBasicRelationship

const (
	KindRelacceptor   KindBasicRelationship = original.KindRelacceptor
	KindRelationship  KindBasicRelationship = original.KindRelationship
	KindRelconnection KindBasicRelationship = original.KindRelconnection
)

func PossibleKindBasicRelationshipValues() [3]KindBasicRelationship {
	return original.PossibleKindBasicRelationshipValues()
}

type MachineGroupType = original.MachineGroupType

const (
	MachineGroupTypeAzureCs    MachineGroupType = original.MachineGroupTypeAzureCs
	MachineGroupTypeAzureSf    MachineGroupType = original.MachineGroupTypeAzureSf
	MachineGroupTypeAzureVmss  MachineGroupType = original.MachineGroupTypeAzureVmss
	MachineGroupTypeUnknown    MachineGroupType = original.MachineGroupTypeUnknown
	MachineGroupTypeUserStatic MachineGroupType = original.MachineGroupTypeUserStatic
)

func PossibleMachineGroupTypeValues() [5]MachineGroupType {
	return original.PossibleMachineGroupTypeValues()
}

type MachineRebootStatus = original.MachineRebootStatus

const (
	MachineRebootStatusNotRebooted MachineRebootStatus = original.MachineRebootStatusNotRebooted
	MachineRebootStatusRebooted    MachineRebootStatus = original.MachineRebootStatusRebooted
	MachineRebootStatusUnknown     MachineRebootStatus = original.MachineRebootStatusUnknown
)

func PossibleMachineRebootStatusValues() [3]MachineRebootStatus {
	return original.PossibleMachineRebootStatusValues()
}

type MonitoringState = original.MonitoringState

const (
	Discovered MonitoringState = original.Discovered
	Monitored  MonitoringState = original.Monitored
)

func PossibleMonitoringStateValues() [2]MonitoringState {
	return original.PossibleMonitoringStateValues()
}

type OperatingSystemFamily = original.OperatingSystemFamily

const (
	OperatingSystemFamilyAix     OperatingSystemFamily = original.OperatingSystemFamilyAix
	OperatingSystemFamilyLinux   OperatingSystemFamily = original.OperatingSystemFamilyLinux
	OperatingSystemFamilySolaris OperatingSystemFamily = original.OperatingSystemFamilySolaris
	OperatingSystemFamilyUnknown OperatingSystemFamily = original.OperatingSystemFamilyUnknown
	OperatingSystemFamilyWindows OperatingSystemFamily = original.OperatingSystemFamilyWindows
)

func PossibleOperatingSystemFamilyValues() [5]OperatingSystemFamily {
	return original.PossibleOperatingSystemFamilyValues()
}

type ProcessRole = original.ProcessRole

const (
	AppServer      ProcessRole = original.AppServer
	DatabaseServer ProcessRole = original.DatabaseServer
	LdapServer     ProcessRole = original.LdapServer
	SmbServer      ProcessRole = original.SmbServer
	WebServer      ProcessRole = original.WebServer
)

func PossibleProcessRoleValues() [5]ProcessRole {
	return original.PossibleProcessRoleValues()
}

type Provider = original.Provider

const (
	Azure Provider = original.Azure
)

func PossibleProviderValues() [1]Provider {
	return original.PossibleProviderValues()
}

type Provider1 = original.Provider1

const (
	Provider1Azure Provider1 = original.Provider1Azure
)

func PossibleProvider1Values() [1]Provider1 {
	return original.PossibleProvider1Values()
}

type VirtualizationState = original.VirtualizationState

const (
	VirtualizationStateHypervisor VirtualizationState = original.VirtualizationStateHypervisor
	VirtualizationStatePhysical   VirtualizationState = original.VirtualizationStatePhysical
	VirtualizationStateUnknown    VirtualizationState = original.VirtualizationStateUnknown
	VirtualizationStateVirtual    VirtualizationState = original.VirtualizationStateVirtual
)

func PossibleVirtualizationStateValues() [4]VirtualizationState {
	return original.PossibleVirtualizationStateValues()
}

type VirtualMachineType = original.VirtualMachineType

const (
	VirtualMachineTypeHyperv    VirtualMachineType = original.VirtualMachineTypeHyperv
	VirtualMachineTypeLdom      VirtualMachineType = original.VirtualMachineTypeLdom
	VirtualMachineTypeLpar      VirtualMachineType = original.VirtualMachineTypeLpar
	VirtualMachineTypeUnknown   VirtualMachineType = original.VirtualMachineTypeUnknown
	VirtualMachineTypeVirtualPc VirtualMachineType = original.VirtualMachineTypeVirtualPc
	VirtualMachineTypeVmware    VirtualMachineType = original.VirtualMachineTypeVmware
	VirtualMachineTypeXen       VirtualMachineType = original.VirtualMachineTypeXen
)

func PossibleVirtualMachineTypeValues() [7]VirtualMachineType {
	return original.PossibleVirtualMachineTypeValues()
}

type Acceptor = original.Acceptor
type AcceptorProperties = original.AcceptorProperties
type AgentConfiguration = original.AgentConfiguration
type AzureCloudServiceConfiguration = original.AzureCloudServiceConfiguration
type AzureHostingConfiguration = original.AzureHostingConfiguration
type AzureProcessHostingConfiguration = original.AzureProcessHostingConfiguration
type AzureServiceFabricClusterConfiguration = original.AzureServiceFabricClusterConfiguration
type AzureVMScaleSetConfiguration = original.AzureVMScaleSetConfiguration
type ClientGroup = original.ClientGroup
type ClientGroupMember = original.ClientGroupMember
type ClientGroupMemberProperties = original.ClientGroupMemberProperties
type ClientGroupMembersCollection = original.ClientGroupMembersCollection
type ClientGroupMembersCollectionIterator = original.ClientGroupMembersCollectionIterator
type ClientGroupMembersCollectionPage = original.ClientGroupMembersCollectionPage
type ClientGroupMembersCount = original.ClientGroupMembersCount
type ClientGroupProperties = original.ClientGroupProperties
type ClientGroupReference = original.ClientGroupReference
type Connection = original.Connection
type ConnectionCollection = original.ConnectionCollection
type ConnectionCollectionIterator = original.ConnectionCollectionIterator
type ConnectionCollectionPage = original.ConnectionCollectionPage
type ConnectionProperties = original.ConnectionProperties
type BasicCoreResource = original.BasicCoreResource
type CoreResource = original.CoreResource
type Error = original.Error
type ErrorResponse = original.ErrorResponse
type BasicHostingConfiguration = original.BasicHostingConfiguration
type HostingConfiguration = original.HostingConfiguration
type HypervisorConfiguration = original.HypervisorConfiguration
type ImageConfiguration = original.ImageConfiguration
type Ipv4NetworkInterface = original.Ipv4NetworkInterface
type Ipv6NetworkInterface = original.Ipv6NetworkInterface
type Liveness = original.Liveness
type Machine = original.Machine
type MachineCollection = original.MachineCollection
type MachineCollectionIterator = original.MachineCollectionIterator
type MachineCollectionPage = original.MachineCollectionPage
type MachineCountsByOperatingSystem = original.MachineCountsByOperatingSystem
type MachineGroup = original.MachineGroup
type MachineGroupCollection = original.MachineGroupCollection
type MachineGroupCollectionIterator = original.MachineGroupCollectionIterator
type MachineGroupCollectionPage = original.MachineGroupCollectionPage
type MachineGroupMapRequest = original.MachineGroupMapRequest
type MachineGroupProperties = original.MachineGroupProperties
type MachineListMapRequest = original.MachineListMapRequest
type MachineProperties = original.MachineProperties
type MachineReference = original.MachineReference
type MachineReferenceWithHints = original.MachineReferenceWithHints
type MachineReferenceWithHintsProperties = original.MachineReferenceWithHintsProperties
type MachineResourcesConfiguration = original.MachineResourcesConfiguration
type MachinesSummary = original.MachinesSummary
type MachinesSummaryProperties = original.MachinesSummaryProperties
type Map = original.Map
type MapEdges = original.MapEdges
type MapNodes = original.MapNodes
type BasicMapRequest = original.BasicMapRequest
type MapRequest = original.MapRequest
type MapResponse = original.MapResponse
type BasicMultipleMachinesMapRequest = original.BasicMultipleMachinesMapRequest
type MultipleMachinesMapRequest = original.MultipleMachinesMapRequest
type NetworkConfiguration = original.NetworkConfiguration
type OperatingSystemConfiguration = original.OperatingSystemConfiguration
type Port = original.Port
type PortCollection = original.PortCollection
type PortCollectionIterator = original.PortCollectionIterator
type PortCollectionPage = original.PortCollectionPage
type PortProperties = original.PortProperties
type PortReference = original.PortReference
type PortReferenceProperties = original.PortReferenceProperties
type Process = original.Process
type ProcessCollection = original.ProcessCollection
type ProcessCollectionIterator = original.ProcessCollectionIterator
type ProcessCollectionPage = original.ProcessCollectionPage
type ProcessDetails = original.ProcessDetails
type ProcessHostedService = original.ProcessHostedService
type BasicProcessHostingConfiguration = original.BasicProcessHostingConfiguration
type ProcessHostingConfiguration = original.ProcessHostingConfiguration
type ProcessProperties = original.ProcessProperties
type ProcessReference = original.ProcessReference
type ProcessReferenceProperties = original.ProcessReferenceProperties
type ProcessUser = original.ProcessUser
type BasicRelationship = original.BasicRelationship
type Relationship = original.Relationship
type RelationshipProperties = original.RelationshipProperties
type Resource = original.Resource
type BasicResourceReference = original.BasicResourceReference
type ResourceReference = original.ResourceReference
type SingleMachineDependencyMapRequest = original.SingleMachineDependencyMapRequest
type Summary = original.Summary
type SummaryProperties = original.SummaryProperties
type Timezone = original.Timezone
type VirtualMachineConfiguration = original.VirtualMachineConfiguration
type PortsClient = original.PortsClient

func NewPortsClient(subscriptionID string) PortsClient {
	return original.NewPortsClient(subscriptionID)
}
func NewPortsClientWithBaseURI(baseURI string, subscriptionID string) PortsClient {
	return original.NewPortsClientWithBaseURI(baseURI, subscriptionID)
}

type ProcessesClient = original.ProcessesClient

func NewProcessesClient(subscriptionID string) ProcessesClient {
	return original.NewProcessesClient(subscriptionID)
}
func NewProcessesClientWithBaseURI(baseURI string, subscriptionID string) ProcessesClient {
	return original.NewProcessesClientWithBaseURI(baseURI, subscriptionID)
}

type SummariesClient = original.SummariesClient

func NewSummariesClient(subscriptionID string) SummariesClient {
	return original.NewSummariesClient(subscriptionID)
}
func NewSummariesClientWithBaseURI(baseURI string, subscriptionID string) SummariesClient {
	return original.NewSummariesClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
