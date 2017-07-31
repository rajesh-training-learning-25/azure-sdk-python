package network

import (
	 original "github.com/Azure/azure-sdk-for-go/service/network/management/2016-03-30/network"
)

type (
	 SecurityRulesClient = original.SecurityRulesClient
	 SubnetsClient = original.SubnetsClient
	 UsagesClient = original.UsagesClient
	 ExpressRouteCircuitAuthorizationsClient = original.ExpressRouteCircuitAuthorizationsClient
	 ExpressRouteCircuitPeeringsClient = original.ExpressRouteCircuitPeeringsClient
	 ExpressRouteServiceProvidersClient = original.ExpressRouteServiceProvidersClient
	 LocalNetworkGatewaysClient = original.LocalNetworkGatewaysClient
	 PublicIPAddressesClient = original.PublicIPAddressesClient
	 VirtualNetworksClient = original.VirtualNetworksClient
	 ManagementClient = original.ManagementClient
	 ExpressRouteCircuitsClient = original.ExpressRouteCircuitsClient
	 LoadBalancersClient = original.LoadBalancersClient
	 RouteTablesClient = original.RouteTablesClient
	 SecurityGroupsClient = original.SecurityGroupsClient
	 VirtualNetworkGatewayConnectionsClient = original.VirtualNetworkGatewayConnectionsClient
	 VirtualNetworkGatewaysClient = original.VirtualNetworkGatewaysClient
	 ApplicationGatewaysClient = original.ApplicationGatewaysClient
	 InterfacesClient = original.InterfacesClient
	 ApplicationGatewayCookieBasedAffinity = original.ApplicationGatewayCookieBasedAffinity
	 ApplicationGatewayOperationalState = original.ApplicationGatewayOperationalState
	 ApplicationGatewayProtocol = original.ApplicationGatewayProtocol
	 ApplicationGatewayRequestRoutingRuleType = original.ApplicationGatewayRequestRoutingRuleType
	 ApplicationGatewaySkuName = original.ApplicationGatewaySkuName
	 ApplicationGatewayTier = original.ApplicationGatewayTier
	 AuthorizationUseStatus = original.AuthorizationUseStatus
	 ExpressRouteCircuitPeeringAdvertisedPublicPrefixState = original.ExpressRouteCircuitPeeringAdvertisedPublicPrefixState
	 ExpressRouteCircuitPeeringState = original.ExpressRouteCircuitPeeringState
	 ExpressRouteCircuitPeeringType = original.ExpressRouteCircuitPeeringType
	 ExpressRouteCircuitSkuFamily = original.ExpressRouteCircuitSkuFamily
	 ExpressRouteCircuitSkuTier = original.ExpressRouteCircuitSkuTier
	 IPAllocationMethod = original.IPAllocationMethod
	 IPVersion = original.IPVersion
	 LoadDistribution = original.LoadDistribution
	 OperationStatus = original.OperationStatus
	 ProbeProtocol = original.ProbeProtocol
	 ProcessorArchitecture = original.ProcessorArchitecture
	 RouteNextHopType = original.RouteNextHopType
	 SecurityRuleAccess = original.SecurityRuleAccess
	 SecurityRuleDirection = original.SecurityRuleDirection
	 SecurityRuleProtocol = original.SecurityRuleProtocol
	 ServiceProviderProvisioningState = original.ServiceProviderProvisioningState
	 TransportProtocol = original.TransportProtocol
	 VirtualNetworkGatewayConnectionStatus = original.VirtualNetworkGatewayConnectionStatus
	 VirtualNetworkGatewayConnectionType = original.VirtualNetworkGatewayConnectionType
	 VirtualNetworkGatewaySkuName = original.VirtualNetworkGatewaySkuName
	 VirtualNetworkGatewaySkuTier = original.VirtualNetworkGatewaySkuTier
	 VirtualNetworkGatewayType = original.VirtualNetworkGatewayType
	 VpnType = original.VpnType
	 AddressSpace = original.AddressSpace
	 ApplicationGateway = original.ApplicationGateway
	 ApplicationGatewayBackendAddress = original.ApplicationGatewayBackendAddress
	 ApplicationGatewayBackendAddressPool = original.ApplicationGatewayBackendAddressPool
	 ApplicationGatewayBackendAddressPoolPropertiesFormat = original.ApplicationGatewayBackendAddressPoolPropertiesFormat
	 ApplicationGatewayBackendHTTPSettings = original.ApplicationGatewayBackendHTTPSettings
	 ApplicationGatewayBackendHTTPSettingsPropertiesFormat = original.ApplicationGatewayBackendHTTPSettingsPropertiesFormat
	 ApplicationGatewayFrontendIPConfiguration = original.ApplicationGatewayFrontendIPConfiguration
	 ApplicationGatewayFrontendIPConfigurationPropertiesFormat = original.ApplicationGatewayFrontendIPConfigurationPropertiesFormat
	 ApplicationGatewayFrontendPort = original.ApplicationGatewayFrontendPort
	 ApplicationGatewayFrontendPortPropertiesFormat = original.ApplicationGatewayFrontendPortPropertiesFormat
	 ApplicationGatewayHTTPListener = original.ApplicationGatewayHTTPListener
	 ApplicationGatewayHTTPListenerPropertiesFormat = original.ApplicationGatewayHTTPListenerPropertiesFormat
	 ApplicationGatewayIPConfiguration = original.ApplicationGatewayIPConfiguration
	 ApplicationGatewayIPConfigurationPropertiesFormat = original.ApplicationGatewayIPConfigurationPropertiesFormat
	 ApplicationGatewayListResult = original.ApplicationGatewayListResult
	 ApplicationGatewayPathRule = original.ApplicationGatewayPathRule
	 ApplicationGatewayPathRulePropertiesFormat = original.ApplicationGatewayPathRulePropertiesFormat
	 ApplicationGatewayProbe = original.ApplicationGatewayProbe
	 ApplicationGatewayProbePropertiesFormat = original.ApplicationGatewayProbePropertiesFormat
	 ApplicationGatewayPropertiesFormat = original.ApplicationGatewayPropertiesFormat
	 ApplicationGatewayRequestRoutingRule = original.ApplicationGatewayRequestRoutingRule
	 ApplicationGatewayRequestRoutingRulePropertiesFormat = original.ApplicationGatewayRequestRoutingRulePropertiesFormat
	 ApplicationGatewaySku = original.ApplicationGatewaySku
	 ApplicationGatewaySslCertificate = original.ApplicationGatewaySslCertificate
	 ApplicationGatewaySslCertificatePropertiesFormat = original.ApplicationGatewaySslCertificatePropertiesFormat
	 ApplicationGatewayURLPathMap = original.ApplicationGatewayURLPathMap
	 ApplicationGatewayURLPathMapPropertiesFormat = original.ApplicationGatewayURLPathMapPropertiesFormat
	 AuthorizationListResult = original.AuthorizationListResult
	 AuthorizationPropertiesFormat = original.AuthorizationPropertiesFormat
	 AzureAsyncOperationResult = original.AzureAsyncOperationResult
	 BackendAddressPool = original.BackendAddressPool
	 BackendAddressPoolPropertiesFormat = original.BackendAddressPoolPropertiesFormat
	 BgpSettings = original.BgpSettings
	 ConnectionResetSharedKey = original.ConnectionResetSharedKey
	 ConnectionSharedKey = original.ConnectionSharedKey
	 ConnectionSharedKeyResult = original.ConnectionSharedKeyResult
	 DhcpOptions = original.DhcpOptions
	 DNSNameAvailabilityResult = original.DNSNameAvailabilityResult
	 Error = original.Error
	 ErrorDetails = original.ErrorDetails
	 ExpressRouteCircuit = original.ExpressRouteCircuit
	 ExpressRouteCircuitArpTable = original.ExpressRouteCircuitArpTable
	 ExpressRouteCircuitAuthorization = original.ExpressRouteCircuitAuthorization
	 ExpressRouteCircuitListResult = original.ExpressRouteCircuitListResult
	 ExpressRouteCircuitPeering = original.ExpressRouteCircuitPeering
	 ExpressRouteCircuitPeeringConfig = original.ExpressRouteCircuitPeeringConfig
	 ExpressRouteCircuitPeeringListResult = original.ExpressRouteCircuitPeeringListResult
	 ExpressRouteCircuitPeeringPropertiesFormat = original.ExpressRouteCircuitPeeringPropertiesFormat
	 ExpressRouteCircuitPropertiesFormat = original.ExpressRouteCircuitPropertiesFormat
	 ExpressRouteCircuitRoutesTable = original.ExpressRouteCircuitRoutesTable
	 ExpressRouteCircuitRoutesTableSummary = original.ExpressRouteCircuitRoutesTableSummary
	 ExpressRouteCircuitsArpTableListResult = original.ExpressRouteCircuitsArpTableListResult
	 ExpressRouteCircuitServiceProviderProperties = original.ExpressRouteCircuitServiceProviderProperties
	 ExpressRouteCircuitSku = original.ExpressRouteCircuitSku
	 ExpressRouteCircuitsRoutesTableListResult = original.ExpressRouteCircuitsRoutesTableListResult
	 ExpressRouteCircuitsRoutesTableSummaryListResult = original.ExpressRouteCircuitsRoutesTableSummaryListResult
	 ExpressRouteCircuitStats = original.ExpressRouteCircuitStats
	 ExpressRouteServiceProvider = original.ExpressRouteServiceProvider
	 ExpressRouteServiceProviderBandwidthsOffered = original.ExpressRouteServiceProviderBandwidthsOffered
	 ExpressRouteServiceProviderListResult = original.ExpressRouteServiceProviderListResult
	 ExpressRouteServiceProviderPropertiesFormat = original.ExpressRouteServiceProviderPropertiesFormat
	 FrontendIPConfiguration = original.FrontendIPConfiguration
	 FrontendIPConfigurationPropertiesFormat = original.FrontendIPConfigurationPropertiesFormat
	 InboundNatPool = original.InboundNatPool
	 InboundNatPoolPropertiesFormat = original.InboundNatPoolPropertiesFormat
	 InboundNatRule = original.InboundNatRule
	 InboundNatRulePropertiesFormat = original.InboundNatRulePropertiesFormat
	 Interface = original.Interface
	 InterfaceDNSSettings = original.InterfaceDNSSettings
	 InterfaceIPConfiguration = original.InterfaceIPConfiguration
	 InterfaceIPConfigurationPropertiesFormat = original.InterfaceIPConfigurationPropertiesFormat
	 InterfaceListResult = original.InterfaceListResult
	 InterfacePropertiesFormat = original.InterfacePropertiesFormat
	 IPConfiguration = original.IPConfiguration
	 IPConfigurationPropertiesFormat = original.IPConfigurationPropertiesFormat
	 LoadBalancer = original.LoadBalancer
	 LoadBalancerListResult = original.LoadBalancerListResult
	 LoadBalancerPropertiesFormat = original.LoadBalancerPropertiesFormat
	 LoadBalancingRule = original.LoadBalancingRule
	 LoadBalancingRulePropertiesFormat = original.LoadBalancingRulePropertiesFormat
	 LocalNetworkGateway = original.LocalNetworkGateway
	 LocalNetworkGatewayListResult = original.LocalNetworkGatewayListResult
	 LocalNetworkGatewayPropertiesFormat = original.LocalNetworkGatewayPropertiesFormat
	 OutboundNatRule = original.OutboundNatRule
	 OutboundNatRulePropertiesFormat = original.OutboundNatRulePropertiesFormat
	 Probe = original.Probe
	 ProbePropertiesFormat = original.ProbePropertiesFormat
	 PublicIPAddress = original.PublicIPAddress
	 PublicIPAddressDNSSettings = original.PublicIPAddressDNSSettings
	 PublicIPAddressListResult = original.PublicIPAddressListResult
	 PublicIPAddressPropertiesFormat = original.PublicIPAddressPropertiesFormat
	 Resource = original.Resource
	 Route = original.Route
	 RouteListResult = original.RouteListResult
	 RoutePropertiesFormat = original.RoutePropertiesFormat
	 RouteTable = original.RouteTable
	 RouteTableListResult = original.RouteTableListResult
	 RouteTablePropertiesFormat = original.RouteTablePropertiesFormat
	 SecurityGroup = original.SecurityGroup
	 SecurityGroupListResult = original.SecurityGroupListResult
	 SecurityGroupPropertiesFormat = original.SecurityGroupPropertiesFormat
	 SecurityRule = original.SecurityRule
	 SecurityRuleListResult = original.SecurityRuleListResult
	 SecurityRulePropertiesFormat = original.SecurityRulePropertiesFormat
	 String = original.String
	 Subnet = original.Subnet
	 SubnetListResult = original.SubnetListResult
	 SubnetPropertiesFormat = original.SubnetPropertiesFormat
	 SubResource = original.SubResource
	 Usage = original.Usage
	 UsageName = original.UsageName
	 UsagesListResult = original.UsagesListResult
	 VirtualNetwork = original.VirtualNetwork
	 VirtualNetworkGateway = original.VirtualNetworkGateway
	 VirtualNetworkGatewayConnection = original.VirtualNetworkGatewayConnection
	 VirtualNetworkGatewayConnectionListResult = original.VirtualNetworkGatewayConnectionListResult
	 VirtualNetworkGatewayConnectionPropertiesFormat = original.VirtualNetworkGatewayConnectionPropertiesFormat
	 VirtualNetworkGatewayIPConfiguration = original.VirtualNetworkGatewayIPConfiguration
	 VirtualNetworkGatewayIPConfigurationPropertiesFormat = original.VirtualNetworkGatewayIPConfigurationPropertiesFormat
	 VirtualNetworkGatewayListResult = original.VirtualNetworkGatewayListResult
	 VirtualNetworkGatewayPropertiesFormat = original.VirtualNetworkGatewayPropertiesFormat
	 VirtualNetworkGatewaySku = original.VirtualNetworkGatewaySku
	 VirtualNetworkListResult = original.VirtualNetworkListResult
	 VirtualNetworkPropertiesFormat = original.VirtualNetworkPropertiesFormat
	 VpnClientConfiguration = original.VpnClientConfiguration
	 VpnClientParameters = original.VpnClientParameters
	 VpnClientRevokedCertificate = original.VpnClientRevokedCertificate
	 VpnClientRevokedCertificatePropertiesFormat = original.VpnClientRevokedCertificatePropertiesFormat
	 VpnClientRootCertificate = original.VpnClientRootCertificate
	 VpnClientRootCertificatePropertiesFormat = original.VpnClientRootCertificatePropertiesFormat
	 RoutesClient = original.RoutesClient
)

const (
	 DefaultBaseURI = original.DefaultBaseURI
	 Disabled = original.Disabled
	 Enabled = original.Enabled
	 Running = original.Running
	 Starting = original.Starting
	 Stopped = original.Stopped
	 Stopping = original.Stopping
	 HTTP = original.HTTP
	 HTTPS = original.HTTPS
	 Basic = original.Basic
	 PathBasedRouting = original.PathBasedRouting
	 StandardLarge = original.StandardLarge
	 StandardMedium = original.StandardMedium
	 StandardSmall = original.StandardSmall
	 Standard = original.Standard
	 Available = original.Available
	 InUse = original.InUse
	 Configured = original.Configured
	 Configuring = original.Configuring
	 NotConfigured = original.NotConfigured
	 ValidationNeeded = original.ValidationNeeded
	 ExpressRouteCircuitPeeringStateDisabled = original.ExpressRouteCircuitPeeringStateDisabled
	 ExpressRouteCircuitPeeringStateEnabled = original.ExpressRouteCircuitPeeringStateEnabled
	 AzurePrivatePeering = original.AzurePrivatePeering
	 AzurePublicPeering = original.AzurePublicPeering
	 MicrosoftPeering = original.MicrosoftPeering
	 MeteredData = original.MeteredData
	 UnlimitedData = original.UnlimitedData
	 ExpressRouteCircuitSkuTierPremium = original.ExpressRouteCircuitSkuTierPremium
	 ExpressRouteCircuitSkuTierStandard = original.ExpressRouteCircuitSkuTierStandard
	 Dynamic = original.Dynamic
	 Static = original.Static
	 IPv4 = original.IPv4
	 IPv6 = original.IPv6
	 Default = original.Default
	 SourceIP = original.SourceIP
	 SourceIPProtocol = original.SourceIPProtocol
	 Failed = original.Failed
	 InProgress = original.InProgress
	 Succeeded = original.Succeeded
	 ProbeProtocolHTTP = original.ProbeProtocolHTTP
	 ProbeProtocolTCP = original.ProbeProtocolTCP
	 Amd64 = original.Amd64
	 X86 = original.X86
	 RouteNextHopTypeInternet = original.RouteNextHopTypeInternet
	 RouteNextHopTypeNone = original.RouteNextHopTypeNone
	 RouteNextHopTypeVirtualAppliance = original.RouteNextHopTypeVirtualAppliance
	 RouteNextHopTypeVirtualNetworkGateway = original.RouteNextHopTypeVirtualNetworkGateway
	 RouteNextHopTypeVnetLocal = original.RouteNextHopTypeVnetLocal
	 Allow = original.Allow
	 Deny = original.Deny
	 Inbound = original.Inbound
	 Outbound = original.Outbound
	 Asterisk = original.Asterisk
	 TCP = original.TCP
	 UDP = original.UDP
	 Deprovisioning = original.Deprovisioning
	 NotProvisioned = original.NotProvisioned
	 Provisioned = original.Provisioned
	 Provisioning = original.Provisioning
	 TransportProtocolTCP = original.TransportProtocolTCP
	 TransportProtocolUDP = original.TransportProtocolUDP
	 Connected = original.Connected
	 Connecting = original.Connecting
	 NotConnected = original.NotConnected
	 Unknown = original.Unknown
	 ExpressRoute = original.ExpressRoute
	 IPsec = original.IPsec
	 Vnet2Vnet = original.Vnet2Vnet
	 VPNClient = original.VPNClient
	 VirtualNetworkGatewaySkuNameBasic = original.VirtualNetworkGatewaySkuNameBasic
	 VirtualNetworkGatewaySkuNameHighPerformance = original.VirtualNetworkGatewaySkuNameHighPerformance
	 VirtualNetworkGatewaySkuNameStandard = original.VirtualNetworkGatewaySkuNameStandard
	 VirtualNetworkGatewaySkuTierBasic = original.VirtualNetworkGatewaySkuTierBasic
	 VirtualNetworkGatewaySkuTierHighPerformance = original.VirtualNetworkGatewaySkuTierHighPerformance
	 VirtualNetworkGatewaySkuTierStandard = original.VirtualNetworkGatewaySkuTierStandard
	 VirtualNetworkGatewayTypeExpressRoute = original.VirtualNetworkGatewayTypeExpressRoute
	 VirtualNetworkGatewayTypeVpn = original.VirtualNetworkGatewayTypeVpn
	 PolicyBased = original.PolicyBased
	 RouteBased = original.RouteBased
)
