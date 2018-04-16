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

package network

import original "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2015-06-15/network"

type VirtualNetworkGatewayConnectionsClient = original.VirtualNetworkGatewayConnectionsClient
type ExpressRouteCircuitAuthorizationsClient = original.ExpressRouteCircuitAuthorizationsClient
type InterfacesClient = original.InterfacesClient
type VirtualNetworkGatewaysClient = original.VirtualNetworkGatewaysClient
type SecurityRulesClient = original.SecurityRulesClient
type ExpressRouteServiceProvidersClient = original.ExpressRouteServiceProvidersClient
type SubnetsClient = original.SubnetsClient
type LoadBalancersClient = original.LoadBalancersClient
type ExpressRouteCircuitPeeringsClient = original.ExpressRouteCircuitPeeringsClient
type ExpressRouteCircuitsClient = original.ExpressRouteCircuitsClient
type ApplicationGatewaysClient = original.ApplicationGatewaysClient
type LocalNetworkGatewaysClient = original.LocalNetworkGatewaysClient
type RouteTablesClient = original.RouteTablesClient
type PublicIPAddressesClient = original.PublicIPAddressesClient
type RoutesClient = original.RoutesClient
type UsagesClient = original.UsagesClient
type ApplicationGatewayCookieBasedAffinity = original.ApplicationGatewayCookieBasedAffinity

const (
	Disabled ApplicationGatewayCookieBasedAffinity = original.Disabled
	Enabled  ApplicationGatewayCookieBasedAffinity = original.Enabled
)

type ApplicationGatewayOperationalState = original.ApplicationGatewayOperationalState

const (
	Running  ApplicationGatewayOperationalState = original.Running
	Starting ApplicationGatewayOperationalState = original.Starting
	Stopped  ApplicationGatewayOperationalState = original.Stopped
	Stopping ApplicationGatewayOperationalState = original.Stopping
)

type ApplicationGatewayProtocol = original.ApplicationGatewayProtocol

const (
	HTTP  ApplicationGatewayProtocol = original.HTTP
	HTTPS ApplicationGatewayProtocol = original.HTTPS
)

type ApplicationGatewayRequestRoutingRuleType = original.ApplicationGatewayRequestRoutingRuleType

const (
	Basic            ApplicationGatewayRequestRoutingRuleType = original.Basic
	PathBasedRouting ApplicationGatewayRequestRoutingRuleType = original.PathBasedRouting
)

type ApplicationGatewaySkuName = original.ApplicationGatewaySkuName

const (
	StandardLarge  ApplicationGatewaySkuName = original.StandardLarge
	StandardMedium ApplicationGatewaySkuName = original.StandardMedium
	StandardSmall  ApplicationGatewaySkuName = original.StandardSmall
)

type ApplicationGatewayTier = original.ApplicationGatewayTier

const (
	Standard ApplicationGatewayTier = original.Standard
)

type AuthorizationUseStatus = original.AuthorizationUseStatus

const (
	Available AuthorizationUseStatus = original.Available
	InUse     AuthorizationUseStatus = original.InUse
)

type ExpressRouteCircuitPeeringAdvertisedPublicPrefixState = original.ExpressRouteCircuitPeeringAdvertisedPublicPrefixState

const (
	Configured       ExpressRouteCircuitPeeringAdvertisedPublicPrefixState = original.Configured
	Configuring      ExpressRouteCircuitPeeringAdvertisedPublicPrefixState = original.Configuring
	NotConfigured    ExpressRouteCircuitPeeringAdvertisedPublicPrefixState = original.NotConfigured
	ValidationNeeded ExpressRouteCircuitPeeringAdvertisedPublicPrefixState = original.ValidationNeeded
)

type ExpressRouteCircuitPeeringState = original.ExpressRouteCircuitPeeringState

const (
	ExpressRouteCircuitPeeringStateDisabled ExpressRouteCircuitPeeringState = original.ExpressRouteCircuitPeeringStateDisabled
	ExpressRouteCircuitPeeringStateEnabled  ExpressRouteCircuitPeeringState = original.ExpressRouteCircuitPeeringStateEnabled
)

type ExpressRouteCircuitPeeringType = original.ExpressRouteCircuitPeeringType

const (
	AzurePrivatePeering ExpressRouteCircuitPeeringType = original.AzurePrivatePeering
	AzurePublicPeering  ExpressRouteCircuitPeeringType = original.AzurePublicPeering
	MicrosoftPeering    ExpressRouteCircuitPeeringType = original.MicrosoftPeering
)

type ExpressRouteCircuitSkuFamily = original.ExpressRouteCircuitSkuFamily

const (
	MeteredData   ExpressRouteCircuitSkuFamily = original.MeteredData
	UnlimitedData ExpressRouteCircuitSkuFamily = original.UnlimitedData
)

type ExpressRouteCircuitSkuTier = original.ExpressRouteCircuitSkuTier

const (
	ExpressRouteCircuitSkuTierPremium  ExpressRouteCircuitSkuTier = original.ExpressRouteCircuitSkuTierPremium
	ExpressRouteCircuitSkuTierStandard ExpressRouteCircuitSkuTier = original.ExpressRouteCircuitSkuTierStandard
)

type IPAllocationMethod = original.IPAllocationMethod

const (
	Dynamic IPAllocationMethod = original.Dynamic
	Static  IPAllocationMethod = original.Static
)

type LoadDistribution = original.LoadDistribution

const (
	Default          LoadDistribution = original.Default
	SourceIP         LoadDistribution = original.SourceIP
	SourceIPProtocol LoadDistribution = original.SourceIPProtocol
)

type OperationStatus = original.OperationStatus

const (
	Failed     OperationStatus = original.Failed
	InProgress OperationStatus = original.InProgress
	Succeeded  OperationStatus = original.Succeeded
)

type ProbeProtocol = original.ProbeProtocol

const (
	ProbeProtocolHTTP ProbeProtocol = original.ProbeProtocolHTTP
	ProbeProtocolTCP  ProbeProtocol = original.ProbeProtocolTCP
)

type ProcessorArchitecture = original.ProcessorArchitecture

const (
	Amd64 ProcessorArchitecture = original.Amd64
	X86   ProcessorArchitecture = original.X86
)

type RouteNextHopType = original.RouteNextHopType

const (
	RouteNextHopTypeInternet              RouteNextHopType = original.RouteNextHopTypeInternet
	RouteNextHopTypeNone                  RouteNextHopType = original.RouteNextHopTypeNone
	RouteNextHopTypeVirtualAppliance      RouteNextHopType = original.RouteNextHopTypeVirtualAppliance
	RouteNextHopTypeVirtualNetworkGateway RouteNextHopType = original.RouteNextHopTypeVirtualNetworkGateway
	RouteNextHopTypeVnetLocal             RouteNextHopType = original.RouteNextHopTypeVnetLocal
)

type SecurityRuleAccess = original.SecurityRuleAccess

const (
	Allow SecurityRuleAccess = original.Allow
	Deny  SecurityRuleAccess = original.Deny
)

type SecurityRuleDirection = original.SecurityRuleDirection

const (
	Inbound  SecurityRuleDirection = original.Inbound
	Outbound SecurityRuleDirection = original.Outbound
)

type SecurityRuleProtocol = original.SecurityRuleProtocol

const (
	Asterisk SecurityRuleProtocol = original.Asterisk
	TCP      SecurityRuleProtocol = original.TCP
	UDP      SecurityRuleProtocol = original.UDP
)

type ServiceProviderProvisioningState = original.ServiceProviderProvisioningState

const (
	Deprovisioning ServiceProviderProvisioningState = original.Deprovisioning
	NotProvisioned ServiceProviderProvisioningState = original.NotProvisioned
	Provisioned    ServiceProviderProvisioningState = original.Provisioned
	Provisioning   ServiceProviderProvisioningState = original.Provisioning
)

type TransportProtocol = original.TransportProtocol

const (
	TransportProtocolTCP TransportProtocol = original.TransportProtocolTCP
	TransportProtocolUDP TransportProtocol = original.TransportProtocolUDP
)

type VirtualNetworkGatewayConnectionStatus = original.VirtualNetworkGatewayConnectionStatus

const (
	Connected    VirtualNetworkGatewayConnectionStatus = original.Connected
	Connecting   VirtualNetworkGatewayConnectionStatus = original.Connecting
	NotConnected VirtualNetworkGatewayConnectionStatus = original.NotConnected
	Unknown      VirtualNetworkGatewayConnectionStatus = original.Unknown
)

type VirtualNetworkGatewayConnectionType = original.VirtualNetworkGatewayConnectionType

const (
	ExpressRoute VirtualNetworkGatewayConnectionType = original.ExpressRoute
	IPsec        VirtualNetworkGatewayConnectionType = original.IPsec
	Vnet2Vnet    VirtualNetworkGatewayConnectionType = original.Vnet2Vnet
	VPNClient    VirtualNetworkGatewayConnectionType = original.VPNClient
)

type VirtualNetworkGatewaySkuName = original.VirtualNetworkGatewaySkuName

const (
	VirtualNetworkGatewaySkuNameBasic           VirtualNetworkGatewaySkuName = original.VirtualNetworkGatewaySkuNameBasic
	VirtualNetworkGatewaySkuNameHighPerformance VirtualNetworkGatewaySkuName = original.VirtualNetworkGatewaySkuNameHighPerformance
	VirtualNetworkGatewaySkuNameStandard        VirtualNetworkGatewaySkuName = original.VirtualNetworkGatewaySkuNameStandard
)

type VirtualNetworkGatewaySkuTier = original.VirtualNetworkGatewaySkuTier

const (
	VirtualNetworkGatewaySkuTierBasic           VirtualNetworkGatewaySkuTier = original.VirtualNetworkGatewaySkuTierBasic
	VirtualNetworkGatewaySkuTierHighPerformance VirtualNetworkGatewaySkuTier = original.VirtualNetworkGatewaySkuTierHighPerformance
	VirtualNetworkGatewaySkuTierStandard        VirtualNetworkGatewaySkuTier = original.VirtualNetworkGatewaySkuTierStandard
)

type VirtualNetworkGatewayType = original.VirtualNetworkGatewayType

const (
	VirtualNetworkGatewayTypeExpressRoute VirtualNetworkGatewayType = original.VirtualNetworkGatewayTypeExpressRoute
	VirtualNetworkGatewayTypeVpn          VirtualNetworkGatewayType = original.VirtualNetworkGatewayTypeVpn
)

type VpnType = original.VpnType

const (
	PolicyBased VpnType = original.PolicyBased
	RouteBased  VpnType = original.RouteBased
)

type AddressSpace = original.AddressSpace
type ApplicationGateway = original.ApplicationGateway
type ApplicationGatewayBackendAddress = original.ApplicationGatewayBackendAddress
type ApplicationGatewayBackendAddressPool = original.ApplicationGatewayBackendAddressPool
type ApplicationGatewayBackendAddressPoolPropertiesFormat = original.ApplicationGatewayBackendAddressPoolPropertiesFormat
type ApplicationGatewayBackendHTTPSettings = original.ApplicationGatewayBackendHTTPSettings
type ApplicationGatewayBackendHTTPSettingsPropertiesFormat = original.ApplicationGatewayBackendHTTPSettingsPropertiesFormat
type ApplicationGatewayFrontendIPConfiguration = original.ApplicationGatewayFrontendIPConfiguration
type ApplicationGatewayFrontendIPConfigurationPropertiesFormat = original.ApplicationGatewayFrontendIPConfigurationPropertiesFormat
type ApplicationGatewayFrontendPort = original.ApplicationGatewayFrontendPort
type ApplicationGatewayFrontendPortPropertiesFormat = original.ApplicationGatewayFrontendPortPropertiesFormat
type ApplicationGatewayHTTPListener = original.ApplicationGatewayHTTPListener
type ApplicationGatewayHTTPListenerPropertiesFormat = original.ApplicationGatewayHTTPListenerPropertiesFormat
type ApplicationGatewayIPConfiguration = original.ApplicationGatewayIPConfiguration
type ApplicationGatewayIPConfigurationPropertiesFormat = original.ApplicationGatewayIPConfigurationPropertiesFormat
type ApplicationGatewayListResult = original.ApplicationGatewayListResult
type ApplicationGatewayListResultIterator = original.ApplicationGatewayListResultIterator
type ApplicationGatewayListResultPage = original.ApplicationGatewayListResultPage
type ApplicationGatewayPathRule = original.ApplicationGatewayPathRule
type ApplicationGatewayPathRulePropertiesFormat = original.ApplicationGatewayPathRulePropertiesFormat
type ApplicationGatewayProbe = original.ApplicationGatewayProbe
type ApplicationGatewayProbePropertiesFormat = original.ApplicationGatewayProbePropertiesFormat
type ApplicationGatewayPropertiesFormat = original.ApplicationGatewayPropertiesFormat
type ApplicationGatewayRequestRoutingRule = original.ApplicationGatewayRequestRoutingRule
type ApplicationGatewayRequestRoutingRulePropertiesFormat = original.ApplicationGatewayRequestRoutingRulePropertiesFormat
type ApplicationGatewaysCreateOrUpdateFuture = original.ApplicationGatewaysCreateOrUpdateFuture
type ApplicationGatewaysDeleteFuture = original.ApplicationGatewaysDeleteFuture
type ApplicationGatewaySku = original.ApplicationGatewaySku
type ApplicationGatewaySslCertificate = original.ApplicationGatewaySslCertificate
type ApplicationGatewaySslCertificatePropertiesFormat = original.ApplicationGatewaySslCertificatePropertiesFormat
type ApplicationGatewaysStartFuture = original.ApplicationGatewaysStartFuture
type ApplicationGatewaysStopFuture = original.ApplicationGatewaysStopFuture
type ApplicationGatewayURLPathMap = original.ApplicationGatewayURLPathMap
type ApplicationGatewayURLPathMapPropertiesFormat = original.ApplicationGatewayURLPathMapPropertiesFormat
type AuthorizationListResult = original.AuthorizationListResult
type AuthorizationListResultIterator = original.AuthorizationListResultIterator
type AuthorizationListResultPage = original.AuthorizationListResultPage
type AuthorizationPropertiesFormat = original.AuthorizationPropertiesFormat
type AzureAsyncOperationResult = original.AzureAsyncOperationResult
type BackendAddressPool = original.BackendAddressPool
type BackendAddressPoolPropertiesFormat = original.BackendAddressPoolPropertiesFormat
type BgpSettings = original.BgpSettings
type ConnectionResetSharedKey = original.ConnectionResetSharedKey
type ConnectionSharedKey = original.ConnectionSharedKey
type ConnectionSharedKeyResult = original.ConnectionSharedKeyResult
type DhcpOptions = original.DhcpOptions
type DNSNameAvailabilityResult = original.DNSNameAvailabilityResult
type Error = original.Error
type ErrorDetails = original.ErrorDetails
type ExpressRouteCircuit = original.ExpressRouteCircuit
type ExpressRouteCircuitArpTable = original.ExpressRouteCircuitArpTable
type ExpressRouteCircuitAuthorization = original.ExpressRouteCircuitAuthorization
type ExpressRouteCircuitAuthorizationsCreateOrUpdateFuture = original.ExpressRouteCircuitAuthorizationsCreateOrUpdateFuture
type ExpressRouteCircuitAuthorizationsDeleteFuture = original.ExpressRouteCircuitAuthorizationsDeleteFuture
type ExpressRouteCircuitListResult = original.ExpressRouteCircuitListResult
type ExpressRouteCircuitListResultIterator = original.ExpressRouteCircuitListResultIterator
type ExpressRouteCircuitListResultPage = original.ExpressRouteCircuitListResultPage
type ExpressRouteCircuitPeering = original.ExpressRouteCircuitPeering
type ExpressRouteCircuitPeeringConfig = original.ExpressRouteCircuitPeeringConfig
type ExpressRouteCircuitPeeringListResult = original.ExpressRouteCircuitPeeringListResult
type ExpressRouteCircuitPeeringListResultIterator = original.ExpressRouteCircuitPeeringListResultIterator
type ExpressRouteCircuitPeeringListResultPage = original.ExpressRouteCircuitPeeringListResultPage
type ExpressRouteCircuitPeeringPropertiesFormat = original.ExpressRouteCircuitPeeringPropertiesFormat
type ExpressRouteCircuitPeeringsCreateOrUpdateFuture = original.ExpressRouteCircuitPeeringsCreateOrUpdateFuture
type ExpressRouteCircuitPeeringsDeleteFuture = original.ExpressRouteCircuitPeeringsDeleteFuture
type ExpressRouteCircuitPropertiesFormat = original.ExpressRouteCircuitPropertiesFormat
type ExpressRouteCircuitRoutesTable = original.ExpressRouteCircuitRoutesTable
type ExpressRouteCircuitsArpTableListResult = original.ExpressRouteCircuitsArpTableListResult
type ExpressRouteCircuitsArpTableListResultIterator = original.ExpressRouteCircuitsArpTableListResultIterator
type ExpressRouteCircuitsArpTableListResultPage = original.ExpressRouteCircuitsArpTableListResultPage
type ExpressRouteCircuitsCreateOrUpdateFuture = original.ExpressRouteCircuitsCreateOrUpdateFuture
type ExpressRouteCircuitsDeleteFuture = original.ExpressRouteCircuitsDeleteFuture
type ExpressRouteCircuitServiceProviderProperties = original.ExpressRouteCircuitServiceProviderProperties
type ExpressRouteCircuitSku = original.ExpressRouteCircuitSku
type ExpressRouteCircuitsRoutesTableListResult = original.ExpressRouteCircuitsRoutesTableListResult
type ExpressRouteCircuitsRoutesTableListResultIterator = original.ExpressRouteCircuitsRoutesTableListResultIterator
type ExpressRouteCircuitsRoutesTableListResultPage = original.ExpressRouteCircuitsRoutesTableListResultPage
type ExpressRouteCircuitsStatsListResult = original.ExpressRouteCircuitsStatsListResult
type ExpressRouteCircuitsStatsListResultIterator = original.ExpressRouteCircuitsStatsListResultIterator
type ExpressRouteCircuitsStatsListResultPage = original.ExpressRouteCircuitsStatsListResultPage
type ExpressRouteCircuitStats = original.ExpressRouteCircuitStats
type ExpressRouteServiceProvider = original.ExpressRouteServiceProvider
type ExpressRouteServiceProviderBandwidthsOffered = original.ExpressRouteServiceProviderBandwidthsOffered
type ExpressRouteServiceProviderListResult = original.ExpressRouteServiceProviderListResult
type ExpressRouteServiceProviderListResultIterator = original.ExpressRouteServiceProviderListResultIterator
type ExpressRouteServiceProviderListResultPage = original.ExpressRouteServiceProviderListResultPage
type ExpressRouteServiceProviderPropertiesFormat = original.ExpressRouteServiceProviderPropertiesFormat
type FrontendIPConfiguration = original.FrontendIPConfiguration
type FrontendIPConfigurationPropertiesFormat = original.FrontendIPConfigurationPropertiesFormat
type InboundNatPool = original.InboundNatPool
type InboundNatPoolPropertiesFormat = original.InboundNatPoolPropertiesFormat
type InboundNatRule = original.InboundNatRule
type InboundNatRulePropertiesFormat = original.InboundNatRulePropertiesFormat
type Interface = original.Interface
type InterfaceDNSSettings = original.InterfaceDNSSettings
type InterfaceIPConfiguration = original.InterfaceIPConfiguration
type InterfaceIPConfigurationPropertiesFormat = original.InterfaceIPConfigurationPropertiesFormat
type InterfaceListResult = original.InterfaceListResult
type InterfaceListResultIterator = original.InterfaceListResultIterator
type InterfaceListResultPage = original.InterfaceListResultPage
type InterfacePropertiesFormat = original.InterfacePropertiesFormat
type InterfacesCreateOrUpdateFuture = original.InterfacesCreateOrUpdateFuture
type InterfacesDeleteFuture = original.InterfacesDeleteFuture
type IPConfiguration = original.IPConfiguration
type IPConfigurationPropertiesFormat = original.IPConfigurationPropertiesFormat
type LoadBalancer = original.LoadBalancer
type LoadBalancerListResult = original.LoadBalancerListResult
type LoadBalancerListResultIterator = original.LoadBalancerListResultIterator
type LoadBalancerListResultPage = original.LoadBalancerListResultPage
type LoadBalancerPropertiesFormat = original.LoadBalancerPropertiesFormat
type LoadBalancersCreateOrUpdateFuture = original.LoadBalancersCreateOrUpdateFuture
type LoadBalancersDeleteFuture = original.LoadBalancersDeleteFuture
type LoadBalancingRule = original.LoadBalancingRule
type LoadBalancingRulePropertiesFormat = original.LoadBalancingRulePropertiesFormat
type LocalNetworkGateway = original.LocalNetworkGateway
type LocalNetworkGatewayListResult = original.LocalNetworkGatewayListResult
type LocalNetworkGatewayListResultIterator = original.LocalNetworkGatewayListResultIterator
type LocalNetworkGatewayListResultPage = original.LocalNetworkGatewayListResultPage
type LocalNetworkGatewayPropertiesFormat = original.LocalNetworkGatewayPropertiesFormat
type LocalNetworkGatewaysCreateOrUpdateFuture = original.LocalNetworkGatewaysCreateOrUpdateFuture
type LocalNetworkGatewaysDeleteFuture = original.LocalNetworkGatewaysDeleteFuture
type OutboundNatRule = original.OutboundNatRule
type OutboundNatRulePropertiesFormat = original.OutboundNatRulePropertiesFormat
type Probe = original.Probe
type ProbePropertiesFormat = original.ProbePropertiesFormat
type PublicIPAddress = original.PublicIPAddress
type PublicIPAddressDNSSettings = original.PublicIPAddressDNSSettings
type PublicIPAddressesCreateOrUpdateFuture = original.PublicIPAddressesCreateOrUpdateFuture
type PublicIPAddressesDeleteFuture = original.PublicIPAddressesDeleteFuture
type PublicIPAddressListResult = original.PublicIPAddressListResult
type PublicIPAddressListResultIterator = original.PublicIPAddressListResultIterator
type PublicIPAddressListResultPage = original.PublicIPAddressListResultPage
type PublicIPAddressPropertiesFormat = original.PublicIPAddressPropertiesFormat
type Resource = original.Resource
type Route = original.Route
type RouteListResult = original.RouteListResult
type RouteListResultIterator = original.RouteListResultIterator
type RouteListResultPage = original.RouteListResultPage
type RoutePropertiesFormat = original.RoutePropertiesFormat
type RoutesCreateOrUpdateFuture = original.RoutesCreateOrUpdateFuture
type RoutesDeleteFuture = original.RoutesDeleteFuture
type RouteTable = original.RouteTable
type RouteTableListResult = original.RouteTableListResult
type RouteTableListResultIterator = original.RouteTableListResultIterator
type RouteTableListResultPage = original.RouteTableListResultPage
type RouteTablePropertiesFormat = original.RouteTablePropertiesFormat
type RouteTablesCreateOrUpdateFuture = original.RouteTablesCreateOrUpdateFuture
type RouteTablesDeleteFuture = original.RouteTablesDeleteFuture
type SecurityGroup = original.SecurityGroup
type SecurityGroupListResult = original.SecurityGroupListResult
type SecurityGroupListResultIterator = original.SecurityGroupListResultIterator
type SecurityGroupListResultPage = original.SecurityGroupListResultPage
type SecurityGroupPropertiesFormat = original.SecurityGroupPropertiesFormat
type SecurityGroupsCreateOrUpdateFuture = original.SecurityGroupsCreateOrUpdateFuture
type SecurityGroupsDeleteFuture = original.SecurityGroupsDeleteFuture
type SecurityRule = original.SecurityRule
type SecurityRuleListResult = original.SecurityRuleListResult
type SecurityRuleListResultIterator = original.SecurityRuleListResultIterator
type SecurityRuleListResultPage = original.SecurityRuleListResultPage
type SecurityRulePropertiesFormat = original.SecurityRulePropertiesFormat
type SecurityRulesCreateOrUpdateFuture = original.SecurityRulesCreateOrUpdateFuture
type SecurityRulesDeleteFuture = original.SecurityRulesDeleteFuture
type String = original.String
type Subnet = original.Subnet
type SubnetListResult = original.SubnetListResult
type SubnetListResultIterator = original.SubnetListResultIterator
type SubnetListResultPage = original.SubnetListResultPage
type SubnetPropertiesFormat = original.SubnetPropertiesFormat
type SubnetsCreateOrUpdateFuture = original.SubnetsCreateOrUpdateFuture
type SubnetsDeleteFuture = original.SubnetsDeleteFuture
type SubResource = original.SubResource
type Usage = original.Usage
type UsageName = original.UsageName
type UsagesListResult = original.UsagesListResult
type UsagesListResultIterator = original.UsagesListResultIterator
type UsagesListResultPage = original.UsagesListResultPage
type VirtualNetwork = original.VirtualNetwork
type VirtualNetworkGateway = original.VirtualNetworkGateway
type VirtualNetworkGatewayConnection = original.VirtualNetworkGatewayConnection
type VirtualNetworkGatewayConnectionListResult = original.VirtualNetworkGatewayConnectionListResult
type VirtualNetworkGatewayConnectionListResultIterator = original.VirtualNetworkGatewayConnectionListResultIterator
type VirtualNetworkGatewayConnectionListResultPage = original.VirtualNetworkGatewayConnectionListResultPage
type VirtualNetworkGatewayConnectionPropertiesFormat = original.VirtualNetworkGatewayConnectionPropertiesFormat
type VirtualNetworkGatewayConnectionsCreateOrUpdateFuture = original.VirtualNetworkGatewayConnectionsCreateOrUpdateFuture
type VirtualNetworkGatewayConnectionsDeleteFuture = original.VirtualNetworkGatewayConnectionsDeleteFuture
type VirtualNetworkGatewayConnectionsResetSharedKeyFuture = original.VirtualNetworkGatewayConnectionsResetSharedKeyFuture
type VirtualNetworkGatewayConnectionsSetSharedKeyFuture = original.VirtualNetworkGatewayConnectionsSetSharedKeyFuture
type VirtualNetworkGatewayIPConfiguration = original.VirtualNetworkGatewayIPConfiguration
type VirtualNetworkGatewayIPConfigurationPropertiesFormat = original.VirtualNetworkGatewayIPConfigurationPropertiesFormat
type VirtualNetworkGatewayListResult = original.VirtualNetworkGatewayListResult
type VirtualNetworkGatewayListResultIterator = original.VirtualNetworkGatewayListResultIterator
type VirtualNetworkGatewayListResultPage = original.VirtualNetworkGatewayListResultPage
type VirtualNetworkGatewayPropertiesFormat = original.VirtualNetworkGatewayPropertiesFormat
type VirtualNetworkGatewaysCreateOrUpdateFuture = original.VirtualNetworkGatewaysCreateOrUpdateFuture
type VirtualNetworkGatewaysDeleteFuture = original.VirtualNetworkGatewaysDeleteFuture
type VirtualNetworkGatewaySku = original.VirtualNetworkGatewaySku
type VirtualNetworkGatewaysResetFuture = original.VirtualNetworkGatewaysResetFuture
type VirtualNetworkListResult = original.VirtualNetworkListResult
type VirtualNetworkListResultIterator = original.VirtualNetworkListResultIterator
type VirtualNetworkListResultPage = original.VirtualNetworkListResultPage
type VirtualNetworkPropertiesFormat = original.VirtualNetworkPropertiesFormat
type VirtualNetworksCreateOrUpdateFuture = original.VirtualNetworksCreateOrUpdateFuture
type VirtualNetworksDeleteFuture = original.VirtualNetworksDeleteFuture
type VpnClientConfiguration = original.VpnClientConfiguration
type VpnClientParameters = original.VpnClientParameters
type VpnClientRevokedCertificate = original.VpnClientRevokedCertificate
type VpnClientRevokedCertificatePropertiesFormat = original.VpnClientRevokedCertificatePropertiesFormat
type VpnClientRootCertificate = original.VpnClientRootCertificate
type VpnClientRootCertificatePropertiesFormat = original.VpnClientRootCertificatePropertiesFormat
type VirtualNetworksClient = original.VirtualNetworksClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type SecurityGroupsClient = original.SecurityGroupsClient

func NewRoutesClient(subscriptionID string) RoutesClient {
	return original.NewRoutesClient(subscriptionID)
}
func NewRoutesClientWithBaseURI(baseURI string, subscriptionID string) RoutesClient {
	return original.NewRoutesClientWithBaseURI(baseURI, subscriptionID)
}
func NewPublicIPAddressesClient(subscriptionID string) PublicIPAddressesClient {
	return original.NewPublicIPAddressesClient(subscriptionID)
}
func NewPublicIPAddressesClientWithBaseURI(baseURI string, subscriptionID string) PublicIPAddressesClient {
	return original.NewPublicIPAddressesClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleApplicationGatewayCookieBasedAffinityValues() []ApplicationGatewayCookieBasedAffinity {
	return original.PossibleApplicationGatewayCookieBasedAffinityValues()
}
func PossibleApplicationGatewayOperationalStateValues() []ApplicationGatewayOperationalState {
	return original.PossibleApplicationGatewayOperationalStateValues()
}
func PossibleApplicationGatewayProtocolValues() []ApplicationGatewayProtocol {
	return original.PossibleApplicationGatewayProtocolValues()
}
func PossibleApplicationGatewayRequestRoutingRuleTypeValues() []ApplicationGatewayRequestRoutingRuleType {
	return original.PossibleApplicationGatewayRequestRoutingRuleTypeValues()
}
func PossibleApplicationGatewaySkuNameValues() []ApplicationGatewaySkuName {
	return original.PossibleApplicationGatewaySkuNameValues()
}
func PossibleApplicationGatewayTierValues() []ApplicationGatewayTier {
	return original.PossibleApplicationGatewayTierValues()
}
func PossibleAuthorizationUseStatusValues() []AuthorizationUseStatus {
	return original.PossibleAuthorizationUseStatusValues()
}
func PossibleExpressRouteCircuitPeeringAdvertisedPublicPrefixStateValues() []ExpressRouteCircuitPeeringAdvertisedPublicPrefixState {
	return original.PossibleExpressRouteCircuitPeeringAdvertisedPublicPrefixStateValues()
}
func PossibleExpressRouteCircuitPeeringStateValues() []ExpressRouteCircuitPeeringState {
	return original.PossibleExpressRouteCircuitPeeringStateValues()
}
func PossibleExpressRouteCircuitPeeringTypeValues() []ExpressRouteCircuitPeeringType {
	return original.PossibleExpressRouteCircuitPeeringTypeValues()
}
func PossibleExpressRouteCircuitSkuFamilyValues() []ExpressRouteCircuitSkuFamily {
	return original.PossibleExpressRouteCircuitSkuFamilyValues()
}
func PossibleExpressRouteCircuitSkuTierValues() []ExpressRouteCircuitSkuTier {
	return original.PossibleExpressRouteCircuitSkuTierValues()
}
func PossibleIPAllocationMethodValues() []IPAllocationMethod {
	return original.PossibleIPAllocationMethodValues()
}
func PossibleLoadDistributionValues() []LoadDistribution {
	return original.PossibleLoadDistributionValues()
}
func PossibleOperationStatusValues() []OperationStatus {
	return original.PossibleOperationStatusValues()
}
func PossibleProbeProtocolValues() []ProbeProtocol {
	return original.PossibleProbeProtocolValues()
}
func PossibleProcessorArchitectureValues() []ProcessorArchitecture {
	return original.PossibleProcessorArchitectureValues()
}
func PossibleRouteNextHopTypeValues() []RouteNextHopType {
	return original.PossibleRouteNextHopTypeValues()
}
func PossibleSecurityRuleAccessValues() []SecurityRuleAccess {
	return original.PossibleSecurityRuleAccessValues()
}
func PossibleSecurityRuleDirectionValues() []SecurityRuleDirection {
	return original.PossibleSecurityRuleDirectionValues()
}
func PossibleSecurityRuleProtocolValues() []SecurityRuleProtocol {
	return original.PossibleSecurityRuleProtocolValues()
}
func PossibleServiceProviderProvisioningStateValues() []ServiceProviderProvisioningState {
	return original.PossibleServiceProviderProvisioningStateValues()
}
func PossibleTransportProtocolValues() []TransportProtocol {
	return original.PossibleTransportProtocolValues()
}
func PossibleVirtualNetworkGatewayConnectionStatusValues() []VirtualNetworkGatewayConnectionStatus {
	return original.PossibleVirtualNetworkGatewayConnectionStatusValues()
}
func PossibleVirtualNetworkGatewayConnectionTypeValues() []VirtualNetworkGatewayConnectionType {
	return original.PossibleVirtualNetworkGatewayConnectionTypeValues()
}
func PossibleVirtualNetworkGatewaySkuNameValues() []VirtualNetworkGatewaySkuName {
	return original.PossibleVirtualNetworkGatewaySkuNameValues()
}
func PossibleVirtualNetworkGatewaySkuTierValues() []VirtualNetworkGatewaySkuTier {
	return original.PossibleVirtualNetworkGatewaySkuTierValues()
}
func PossibleVirtualNetworkGatewayTypeValues() []VirtualNetworkGatewayType {
	return original.PossibleVirtualNetworkGatewayTypeValues()
}
func PossibleVpnTypeValues() []VpnType {
	return original.PossibleVpnTypeValues()
}
func NewVirtualNetworksClient(subscriptionID string) VirtualNetworksClient {
	return original.NewVirtualNetworksClient(subscriptionID)
}
func NewVirtualNetworksClientWithBaseURI(baseURI string, subscriptionID string) VirtualNetworksClient {
	return original.NewVirtualNetworksClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/2017-03-09"
}
func Version() string {
	return original.Version()
}
func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewSecurityGroupsClient(subscriptionID string) SecurityGroupsClient {
	return original.NewSecurityGroupsClient(subscriptionID)
}
func NewSecurityGroupsClientWithBaseURI(baseURI string, subscriptionID string) SecurityGroupsClient {
	return original.NewSecurityGroupsClientWithBaseURI(baseURI, subscriptionID)
}
func NewUsagesClient(subscriptionID string) UsagesClient {
	return original.NewUsagesClient(subscriptionID)
}
func NewUsagesClientWithBaseURI(baseURI string, subscriptionID string) UsagesClient {
	return original.NewUsagesClientWithBaseURI(baseURI, subscriptionID)
}
func NewExpressRouteCircuitAuthorizationsClient(subscriptionID string) ExpressRouteCircuitAuthorizationsClient {
	return original.NewExpressRouteCircuitAuthorizationsClient(subscriptionID)
}
func NewExpressRouteCircuitAuthorizationsClientWithBaseURI(baseURI string, subscriptionID string) ExpressRouteCircuitAuthorizationsClient {
	return original.NewExpressRouteCircuitAuthorizationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewInterfacesClient(subscriptionID string) InterfacesClient {
	return original.NewInterfacesClient(subscriptionID)
}
func NewInterfacesClientWithBaseURI(baseURI string, subscriptionID string) InterfacesClient {
	return original.NewInterfacesClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualNetworkGatewaysClient(subscriptionID string) VirtualNetworkGatewaysClient {
	return original.NewVirtualNetworkGatewaysClient(subscriptionID)
}
func NewVirtualNetworkGatewaysClientWithBaseURI(baseURI string, subscriptionID string) VirtualNetworkGatewaysClient {
	return original.NewVirtualNetworkGatewaysClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualNetworkGatewayConnectionsClient(subscriptionID string) VirtualNetworkGatewayConnectionsClient {
	return original.NewVirtualNetworkGatewayConnectionsClient(subscriptionID)
}
func NewVirtualNetworkGatewayConnectionsClientWithBaseURI(baseURI string, subscriptionID string) VirtualNetworkGatewayConnectionsClient {
	return original.NewVirtualNetworkGatewayConnectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSubnetsClient(subscriptionID string) SubnetsClient {
	return original.NewSubnetsClient(subscriptionID)
}
func NewSubnetsClientWithBaseURI(baseURI string, subscriptionID string) SubnetsClient {
	return original.NewSubnetsClientWithBaseURI(baseURI, subscriptionID)
}
func NewLoadBalancersClient(subscriptionID string) LoadBalancersClient {
	return original.NewLoadBalancersClient(subscriptionID)
}
func NewLoadBalancersClientWithBaseURI(baseURI string, subscriptionID string) LoadBalancersClient {
	return original.NewLoadBalancersClientWithBaseURI(baseURI, subscriptionID)
}
func NewExpressRouteCircuitPeeringsClient(subscriptionID string) ExpressRouteCircuitPeeringsClient {
	return original.NewExpressRouteCircuitPeeringsClient(subscriptionID)
}
func NewExpressRouteCircuitPeeringsClientWithBaseURI(baseURI string, subscriptionID string) ExpressRouteCircuitPeeringsClient {
	return original.NewExpressRouteCircuitPeeringsClientWithBaseURI(baseURI, subscriptionID)
}
func NewExpressRouteCircuitsClient(subscriptionID string) ExpressRouteCircuitsClient {
	return original.NewExpressRouteCircuitsClient(subscriptionID)
}
func NewExpressRouteCircuitsClientWithBaseURI(baseURI string, subscriptionID string) ExpressRouteCircuitsClient {
	return original.NewExpressRouteCircuitsClientWithBaseURI(baseURI, subscriptionID)
}
func NewApplicationGatewaysClient(subscriptionID string) ApplicationGatewaysClient {
	return original.NewApplicationGatewaysClient(subscriptionID)
}
func NewApplicationGatewaysClientWithBaseURI(baseURI string, subscriptionID string) ApplicationGatewaysClient {
	return original.NewApplicationGatewaysClientWithBaseURI(baseURI, subscriptionID)
}
func NewLocalNetworkGatewaysClient(subscriptionID string) LocalNetworkGatewaysClient {
	return original.NewLocalNetworkGatewaysClient(subscriptionID)
}
func NewLocalNetworkGatewaysClientWithBaseURI(baseURI string, subscriptionID string) LocalNetworkGatewaysClient {
	return original.NewLocalNetworkGatewaysClientWithBaseURI(baseURI, subscriptionID)
}
func NewSecurityRulesClient(subscriptionID string) SecurityRulesClient {
	return original.NewSecurityRulesClient(subscriptionID)
}
func NewSecurityRulesClientWithBaseURI(baseURI string, subscriptionID string) SecurityRulesClient {
	return original.NewSecurityRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewExpressRouteServiceProvidersClient(subscriptionID string) ExpressRouteServiceProvidersClient {
	return original.NewExpressRouteServiceProvidersClient(subscriptionID)
}
func NewExpressRouteServiceProvidersClientWithBaseURI(baseURI string, subscriptionID string) ExpressRouteServiceProvidersClient {
	return original.NewExpressRouteServiceProvidersClientWithBaseURI(baseURI, subscriptionID)
}
func NewRouteTablesClient(subscriptionID string) RouteTablesClient {
	return original.NewRouteTablesClient(subscriptionID)
}
func NewRouteTablesClientWithBaseURI(baseURI string, subscriptionID string) RouteTablesClient {
	return original.NewRouteTablesClientWithBaseURI(baseURI, subscriptionID)
}
