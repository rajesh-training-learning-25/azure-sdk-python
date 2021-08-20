// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package containerservice

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2021-07-01/containerservice"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AgentPoolMode = original.AgentPoolMode

const (
	AgentPoolModeSystem AgentPoolMode = original.AgentPoolModeSystem
	AgentPoolModeUser   AgentPoolMode = original.AgentPoolModeUser
)

type AgentPoolType = original.AgentPoolType

const (
	AgentPoolTypeAvailabilitySet         AgentPoolType = original.AgentPoolTypeAvailabilitySet
	AgentPoolTypeVirtualMachineScaleSets AgentPoolType = original.AgentPoolTypeVirtualMachineScaleSets
)

type Code = original.Code

const (
	CodeRunning Code = original.CodeRunning
	CodeStopped Code = original.CodeStopped
)

type ConnectionStatus = original.ConnectionStatus

const (
	ConnectionStatusApproved     ConnectionStatus = original.ConnectionStatusApproved
	ConnectionStatusDisconnected ConnectionStatus = original.ConnectionStatusDisconnected
	ConnectionStatusPending      ConnectionStatus = original.ConnectionStatusPending
	ConnectionStatusRejected     ConnectionStatus = original.ConnectionStatusRejected
)

type CreatedByType = original.CreatedByType

const (
	CreatedByTypeApplication     CreatedByType = original.CreatedByTypeApplication
	CreatedByTypeKey             CreatedByType = original.CreatedByTypeKey
	CreatedByTypeManagedIdentity CreatedByType = original.CreatedByTypeManagedIdentity
	CreatedByTypeUser            CreatedByType = original.CreatedByTypeUser
)

type Expander = original.Expander

const (
	ExpanderLeastWaste Expander = original.ExpanderLeastWaste
	ExpanderMostPods   Expander = original.ExpanderMostPods
	ExpanderPriority   Expander = original.ExpanderPriority
	ExpanderRandom     Expander = original.ExpanderRandom
)

type ExtendedLocationTypes = original.ExtendedLocationTypes

const (
	ExtendedLocationTypesEdgeZone ExtendedLocationTypes = original.ExtendedLocationTypesEdgeZone
)

type GPUInstanceProfile = original.GPUInstanceProfile

const (
	GPUInstanceProfileMIG1g GPUInstanceProfile = original.GPUInstanceProfileMIG1g
	GPUInstanceProfileMIG2g GPUInstanceProfile = original.GPUInstanceProfileMIG2g
	GPUInstanceProfileMIG3g GPUInstanceProfile = original.GPUInstanceProfileMIG3g
	GPUInstanceProfileMIG4g GPUInstanceProfile = original.GPUInstanceProfileMIG4g
	GPUInstanceProfileMIG7g GPUInstanceProfile = original.GPUInstanceProfileMIG7g
)

type KubeletDiskType = original.KubeletDiskType

const (
	KubeletDiskTypeOS        KubeletDiskType = original.KubeletDiskTypeOS
	KubeletDiskTypeTemporary KubeletDiskType = original.KubeletDiskTypeTemporary
)

type LicenseType = original.LicenseType

const (
	LicenseTypeNone          LicenseType = original.LicenseTypeNone
	LicenseTypeWindowsServer LicenseType = original.LicenseTypeWindowsServer
)

type LoadBalancerSku = original.LoadBalancerSku

const (
	LoadBalancerSkuBasic    LoadBalancerSku = original.LoadBalancerSkuBasic
	LoadBalancerSkuStandard LoadBalancerSku = original.LoadBalancerSkuStandard
)

type ManagedClusterPodIdentityProvisioningState = original.ManagedClusterPodIdentityProvisioningState

const (
	ManagedClusterPodIdentityProvisioningStateAssigned ManagedClusterPodIdentityProvisioningState = original.ManagedClusterPodIdentityProvisioningStateAssigned
	ManagedClusterPodIdentityProvisioningStateDeleting ManagedClusterPodIdentityProvisioningState = original.ManagedClusterPodIdentityProvisioningStateDeleting
	ManagedClusterPodIdentityProvisioningStateFailed   ManagedClusterPodIdentityProvisioningState = original.ManagedClusterPodIdentityProvisioningStateFailed
	ManagedClusterPodIdentityProvisioningStateUpdating ManagedClusterPodIdentityProvisioningState = original.ManagedClusterPodIdentityProvisioningStateUpdating
)

type ManagedClusterSKUName = original.ManagedClusterSKUName

const (
	ManagedClusterSKUNameBasic ManagedClusterSKUName = original.ManagedClusterSKUNameBasic
)

type ManagedClusterSKUTier = original.ManagedClusterSKUTier

const (
	ManagedClusterSKUTierFree ManagedClusterSKUTier = original.ManagedClusterSKUTierFree
	ManagedClusterSKUTierPaid ManagedClusterSKUTier = original.ManagedClusterSKUTierPaid
)

type NetworkMode = original.NetworkMode

const (
	NetworkModeBridge      NetworkMode = original.NetworkModeBridge
	NetworkModeTransparent NetworkMode = original.NetworkModeTransparent
)

type NetworkPlugin = original.NetworkPlugin

const (
	NetworkPluginAzure   NetworkPlugin = original.NetworkPluginAzure
	NetworkPluginKubenet NetworkPlugin = original.NetworkPluginKubenet
)

type NetworkPolicy = original.NetworkPolicy

const (
	NetworkPolicyAzure  NetworkPolicy = original.NetworkPolicyAzure
	NetworkPolicyCalico NetworkPolicy = original.NetworkPolicyCalico
)

type OSDiskType = original.OSDiskType

const (
	OSDiskTypeEphemeral OSDiskType = original.OSDiskTypeEphemeral
	OSDiskTypeManaged   OSDiskType = original.OSDiskTypeManaged
)

type OSSKU = original.OSSKU

const (
	OSSKUCBLMariner OSSKU = original.OSSKUCBLMariner
	OSSKUUbuntu     OSSKU = original.OSSKUUbuntu
)

type OSType = original.OSType

const (
	OSTypeLinux   OSType = original.OSTypeLinux
	OSTypeWindows OSType = original.OSTypeWindows
)

type OutboundType = original.OutboundType

const (
	OutboundTypeLoadBalancer           OutboundType = original.OutboundTypeLoadBalancer
	OutboundTypeManagedNATGateway      OutboundType = original.OutboundTypeManagedNATGateway
	OutboundTypeUserAssignedNATGateway OutboundType = original.OutboundTypeUserAssignedNATGateway
	OutboundTypeUserDefinedRouting     OutboundType = original.OutboundTypeUserDefinedRouting
)

type PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningState

const (
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateCreating
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateDeleting
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateFailed
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateSucceeded
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	ResourceIdentityTypeNone           ResourceIdentityType = original.ResourceIdentityTypeNone
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = original.ResourceIdentityTypeSystemAssigned
	ResourceIdentityTypeUserAssigned   ResourceIdentityType = original.ResourceIdentityTypeUserAssigned
)

type ScaleDownMode = original.ScaleDownMode

const (
	ScaleDownModeDeallocate ScaleDownMode = original.ScaleDownModeDeallocate
	ScaleDownModeDelete     ScaleDownMode = original.ScaleDownModeDelete
)

type ScaleSetEvictionPolicy = original.ScaleSetEvictionPolicy

const (
	ScaleSetEvictionPolicyDeallocate ScaleSetEvictionPolicy = original.ScaleSetEvictionPolicyDeallocate
	ScaleSetEvictionPolicyDelete     ScaleSetEvictionPolicy = original.ScaleSetEvictionPolicyDelete
)

type ScaleSetPriority = original.ScaleSetPriority

const (
	ScaleSetPriorityRegular ScaleSetPriority = original.ScaleSetPriorityRegular
	ScaleSetPrioritySpot    ScaleSetPriority = original.ScaleSetPrioritySpot
)

type StorageProfileTypes = original.StorageProfileTypes

const (
	StorageProfileTypesManagedDisks   StorageProfileTypes = original.StorageProfileTypesManagedDisks
	StorageProfileTypesStorageAccount StorageProfileTypes = original.StorageProfileTypesStorageAccount
)

type UpgradeChannel = original.UpgradeChannel

const (
	UpgradeChannelNodeImage UpgradeChannel = original.UpgradeChannelNodeImage
	UpgradeChannelNone      UpgradeChannel = original.UpgradeChannelNone
	UpgradeChannelPatch     UpgradeChannel = original.UpgradeChannelPatch
	UpgradeChannelRapid     UpgradeChannel = original.UpgradeChannelRapid
	UpgradeChannelStable    UpgradeChannel = original.UpgradeChannelStable
)

type VMSizeTypes = original.VMSizeTypes

const (
	VMSizeTypesStandardA1          VMSizeTypes = original.VMSizeTypesStandardA1
	VMSizeTypesStandardA10         VMSizeTypes = original.VMSizeTypesStandardA10
	VMSizeTypesStandardA11         VMSizeTypes = original.VMSizeTypesStandardA11
	VMSizeTypesStandardA1V2        VMSizeTypes = original.VMSizeTypesStandardA1V2
	VMSizeTypesStandardA2          VMSizeTypes = original.VMSizeTypesStandardA2
	VMSizeTypesStandardA2mV2       VMSizeTypes = original.VMSizeTypesStandardA2mV2
	VMSizeTypesStandardA2V2        VMSizeTypes = original.VMSizeTypesStandardA2V2
	VMSizeTypesStandardA3          VMSizeTypes = original.VMSizeTypesStandardA3
	VMSizeTypesStandardA4          VMSizeTypes = original.VMSizeTypesStandardA4
	VMSizeTypesStandardA4mV2       VMSizeTypes = original.VMSizeTypesStandardA4mV2
	VMSizeTypesStandardA4V2        VMSizeTypes = original.VMSizeTypesStandardA4V2
	VMSizeTypesStandardA5          VMSizeTypes = original.VMSizeTypesStandardA5
	VMSizeTypesStandardA6          VMSizeTypes = original.VMSizeTypesStandardA6
	VMSizeTypesStandardA7          VMSizeTypes = original.VMSizeTypesStandardA7
	VMSizeTypesStandardA8          VMSizeTypes = original.VMSizeTypesStandardA8
	VMSizeTypesStandardA8mV2       VMSizeTypes = original.VMSizeTypesStandardA8mV2
	VMSizeTypesStandardA8V2        VMSizeTypes = original.VMSizeTypesStandardA8V2
	VMSizeTypesStandardA9          VMSizeTypes = original.VMSizeTypesStandardA9
	VMSizeTypesStandardB2ms        VMSizeTypes = original.VMSizeTypesStandardB2ms
	VMSizeTypesStandardB2s         VMSizeTypes = original.VMSizeTypesStandardB2s
	VMSizeTypesStandardB4ms        VMSizeTypes = original.VMSizeTypesStandardB4ms
	VMSizeTypesStandardB8ms        VMSizeTypes = original.VMSizeTypesStandardB8ms
	VMSizeTypesStandardD1          VMSizeTypes = original.VMSizeTypesStandardD1
	VMSizeTypesStandardD11         VMSizeTypes = original.VMSizeTypesStandardD11
	VMSizeTypesStandardD11V2       VMSizeTypes = original.VMSizeTypesStandardD11V2
	VMSizeTypesStandardD11V2Promo  VMSizeTypes = original.VMSizeTypesStandardD11V2Promo
	VMSizeTypesStandardD12         VMSizeTypes = original.VMSizeTypesStandardD12
	VMSizeTypesStandardD12V2       VMSizeTypes = original.VMSizeTypesStandardD12V2
	VMSizeTypesStandardD12V2Promo  VMSizeTypes = original.VMSizeTypesStandardD12V2Promo
	VMSizeTypesStandardD13         VMSizeTypes = original.VMSizeTypesStandardD13
	VMSizeTypesStandardD13V2       VMSizeTypes = original.VMSizeTypesStandardD13V2
	VMSizeTypesStandardD13V2Promo  VMSizeTypes = original.VMSizeTypesStandardD13V2Promo
	VMSizeTypesStandardD14         VMSizeTypes = original.VMSizeTypesStandardD14
	VMSizeTypesStandardD14V2       VMSizeTypes = original.VMSizeTypesStandardD14V2
	VMSizeTypesStandardD14V2Promo  VMSizeTypes = original.VMSizeTypesStandardD14V2Promo
	VMSizeTypesStandardD15V2       VMSizeTypes = original.VMSizeTypesStandardD15V2
	VMSizeTypesStandardD16sV3      VMSizeTypes = original.VMSizeTypesStandardD16sV3
	VMSizeTypesStandardD16V3       VMSizeTypes = original.VMSizeTypesStandardD16V3
	VMSizeTypesStandardD1V2        VMSizeTypes = original.VMSizeTypesStandardD1V2
	VMSizeTypesStandardD2          VMSizeTypes = original.VMSizeTypesStandardD2
	VMSizeTypesStandardD2sV3       VMSizeTypes = original.VMSizeTypesStandardD2sV3
	VMSizeTypesStandardD2V2        VMSizeTypes = original.VMSizeTypesStandardD2V2
	VMSizeTypesStandardD2V2Promo   VMSizeTypes = original.VMSizeTypesStandardD2V2Promo
	VMSizeTypesStandardD2V3        VMSizeTypes = original.VMSizeTypesStandardD2V3
	VMSizeTypesStandardD3          VMSizeTypes = original.VMSizeTypesStandardD3
	VMSizeTypesStandardD32sV3      VMSizeTypes = original.VMSizeTypesStandardD32sV3
	VMSizeTypesStandardD32V3       VMSizeTypes = original.VMSizeTypesStandardD32V3
	VMSizeTypesStandardD3V2        VMSizeTypes = original.VMSizeTypesStandardD3V2
	VMSizeTypesStandardD3V2Promo   VMSizeTypes = original.VMSizeTypesStandardD3V2Promo
	VMSizeTypesStandardD4          VMSizeTypes = original.VMSizeTypesStandardD4
	VMSizeTypesStandardD4sV3       VMSizeTypes = original.VMSizeTypesStandardD4sV3
	VMSizeTypesStandardD4V2        VMSizeTypes = original.VMSizeTypesStandardD4V2
	VMSizeTypesStandardD4V2Promo   VMSizeTypes = original.VMSizeTypesStandardD4V2Promo
	VMSizeTypesStandardD4V3        VMSizeTypes = original.VMSizeTypesStandardD4V3
	VMSizeTypesStandardD5V2        VMSizeTypes = original.VMSizeTypesStandardD5V2
	VMSizeTypesStandardD5V2Promo   VMSizeTypes = original.VMSizeTypesStandardD5V2Promo
	VMSizeTypesStandardD64sV3      VMSizeTypes = original.VMSizeTypesStandardD64sV3
	VMSizeTypesStandardD64V3       VMSizeTypes = original.VMSizeTypesStandardD64V3
	VMSizeTypesStandardD8sV3       VMSizeTypes = original.VMSizeTypesStandardD8sV3
	VMSizeTypesStandardD8V3        VMSizeTypes = original.VMSizeTypesStandardD8V3
	VMSizeTypesStandardDS1         VMSizeTypes = original.VMSizeTypesStandardDS1
	VMSizeTypesStandardDS11        VMSizeTypes = original.VMSizeTypesStandardDS11
	VMSizeTypesStandardDS11V2      VMSizeTypes = original.VMSizeTypesStandardDS11V2
	VMSizeTypesStandardDS11V2Promo VMSizeTypes = original.VMSizeTypesStandardDS11V2Promo
	VMSizeTypesStandardDS12        VMSizeTypes = original.VMSizeTypesStandardDS12
	VMSizeTypesStandardDS12V2      VMSizeTypes = original.VMSizeTypesStandardDS12V2
	VMSizeTypesStandardDS12V2Promo VMSizeTypes = original.VMSizeTypesStandardDS12V2Promo
	VMSizeTypesStandardDS13        VMSizeTypes = original.VMSizeTypesStandardDS13
	VMSizeTypesStandardDS132V2     VMSizeTypes = original.VMSizeTypesStandardDS132V2
	VMSizeTypesStandardDS134V2     VMSizeTypes = original.VMSizeTypesStandardDS134V2
	VMSizeTypesStandardDS13V2      VMSizeTypes = original.VMSizeTypesStandardDS13V2
	VMSizeTypesStandardDS13V2Promo VMSizeTypes = original.VMSizeTypesStandardDS13V2Promo
	VMSizeTypesStandardDS14        VMSizeTypes = original.VMSizeTypesStandardDS14
	VMSizeTypesStandardDS144V2     VMSizeTypes = original.VMSizeTypesStandardDS144V2
	VMSizeTypesStandardDS148V2     VMSizeTypes = original.VMSizeTypesStandardDS148V2
	VMSizeTypesStandardDS14V2      VMSizeTypes = original.VMSizeTypesStandardDS14V2
	VMSizeTypesStandardDS14V2Promo VMSizeTypes = original.VMSizeTypesStandardDS14V2Promo
	VMSizeTypesStandardDS15V2      VMSizeTypes = original.VMSizeTypesStandardDS15V2
	VMSizeTypesStandardDS1V2       VMSizeTypes = original.VMSizeTypesStandardDS1V2
	VMSizeTypesStandardDS2         VMSizeTypes = original.VMSizeTypesStandardDS2
	VMSizeTypesStandardDS2V2       VMSizeTypes = original.VMSizeTypesStandardDS2V2
	VMSizeTypesStandardDS2V2Promo  VMSizeTypes = original.VMSizeTypesStandardDS2V2Promo
	VMSizeTypesStandardDS3         VMSizeTypes = original.VMSizeTypesStandardDS3
	VMSizeTypesStandardDS3V2       VMSizeTypes = original.VMSizeTypesStandardDS3V2
	VMSizeTypesStandardDS3V2Promo  VMSizeTypes = original.VMSizeTypesStandardDS3V2Promo
	VMSizeTypesStandardDS4         VMSizeTypes = original.VMSizeTypesStandardDS4
	VMSizeTypesStandardDS4V2       VMSizeTypes = original.VMSizeTypesStandardDS4V2
	VMSizeTypesStandardDS4V2Promo  VMSizeTypes = original.VMSizeTypesStandardDS4V2Promo
	VMSizeTypesStandardDS5V2       VMSizeTypes = original.VMSizeTypesStandardDS5V2
	VMSizeTypesStandardDS5V2Promo  VMSizeTypes = original.VMSizeTypesStandardDS5V2Promo
	VMSizeTypesStandardE16sV3      VMSizeTypes = original.VMSizeTypesStandardE16sV3
	VMSizeTypesStandardE16V3       VMSizeTypes = original.VMSizeTypesStandardE16V3
	VMSizeTypesStandardE2sV3       VMSizeTypes = original.VMSizeTypesStandardE2sV3
	VMSizeTypesStandardE2V3        VMSizeTypes = original.VMSizeTypesStandardE2V3
	VMSizeTypesStandardE3216sV3    VMSizeTypes = original.VMSizeTypesStandardE3216sV3
	VMSizeTypesStandardE328sV3     VMSizeTypes = original.VMSizeTypesStandardE328sV3
	VMSizeTypesStandardE32sV3      VMSizeTypes = original.VMSizeTypesStandardE32sV3
	VMSizeTypesStandardE32V3       VMSizeTypes = original.VMSizeTypesStandardE32V3
	VMSizeTypesStandardE4sV3       VMSizeTypes = original.VMSizeTypesStandardE4sV3
	VMSizeTypesStandardE4V3        VMSizeTypes = original.VMSizeTypesStandardE4V3
	VMSizeTypesStandardE6416sV3    VMSizeTypes = original.VMSizeTypesStandardE6416sV3
	VMSizeTypesStandardE6432sV3    VMSizeTypes = original.VMSizeTypesStandardE6432sV3
	VMSizeTypesStandardE64sV3      VMSizeTypes = original.VMSizeTypesStandardE64sV3
	VMSizeTypesStandardE64V3       VMSizeTypes = original.VMSizeTypesStandardE64V3
	VMSizeTypesStandardE8sV3       VMSizeTypes = original.VMSizeTypesStandardE8sV3
	VMSizeTypesStandardE8V3        VMSizeTypes = original.VMSizeTypesStandardE8V3
	VMSizeTypesStandardF1          VMSizeTypes = original.VMSizeTypesStandardF1
	VMSizeTypesStandardF16         VMSizeTypes = original.VMSizeTypesStandardF16
	VMSizeTypesStandardF16s        VMSizeTypes = original.VMSizeTypesStandardF16s
	VMSizeTypesStandardF16sV2      VMSizeTypes = original.VMSizeTypesStandardF16sV2
	VMSizeTypesStandardF1s         VMSizeTypes = original.VMSizeTypesStandardF1s
	VMSizeTypesStandardF2          VMSizeTypes = original.VMSizeTypesStandardF2
	VMSizeTypesStandardF2s         VMSizeTypes = original.VMSizeTypesStandardF2s
	VMSizeTypesStandardF2sV2       VMSizeTypes = original.VMSizeTypesStandardF2sV2
	VMSizeTypesStandardF32sV2      VMSizeTypes = original.VMSizeTypesStandardF32sV2
	VMSizeTypesStandardF4          VMSizeTypes = original.VMSizeTypesStandardF4
	VMSizeTypesStandardF4s         VMSizeTypes = original.VMSizeTypesStandardF4s
	VMSizeTypesStandardF4sV2       VMSizeTypes = original.VMSizeTypesStandardF4sV2
	VMSizeTypesStandardF64sV2      VMSizeTypes = original.VMSizeTypesStandardF64sV2
	VMSizeTypesStandardF72sV2      VMSizeTypes = original.VMSizeTypesStandardF72sV2
	VMSizeTypesStandardF8          VMSizeTypes = original.VMSizeTypesStandardF8
	VMSizeTypesStandardF8s         VMSizeTypes = original.VMSizeTypesStandardF8s
	VMSizeTypesStandardF8sV2       VMSizeTypes = original.VMSizeTypesStandardF8sV2
	VMSizeTypesStandardG1          VMSizeTypes = original.VMSizeTypesStandardG1
	VMSizeTypesStandardG2          VMSizeTypes = original.VMSizeTypesStandardG2
	VMSizeTypesStandardG3          VMSizeTypes = original.VMSizeTypesStandardG3
	VMSizeTypesStandardG4          VMSizeTypes = original.VMSizeTypesStandardG4
	VMSizeTypesStandardG5          VMSizeTypes = original.VMSizeTypesStandardG5
	VMSizeTypesStandardGS1         VMSizeTypes = original.VMSizeTypesStandardGS1
	VMSizeTypesStandardGS2         VMSizeTypes = original.VMSizeTypesStandardGS2
	VMSizeTypesStandardGS3         VMSizeTypes = original.VMSizeTypesStandardGS3
	VMSizeTypesStandardGS4         VMSizeTypes = original.VMSizeTypesStandardGS4
	VMSizeTypesStandardGS44        VMSizeTypes = original.VMSizeTypesStandardGS44
	VMSizeTypesStandardGS48        VMSizeTypes = original.VMSizeTypesStandardGS48
	VMSizeTypesStandardGS5         VMSizeTypes = original.VMSizeTypesStandardGS5
	VMSizeTypesStandardGS516       VMSizeTypes = original.VMSizeTypesStandardGS516
	VMSizeTypesStandardGS58        VMSizeTypes = original.VMSizeTypesStandardGS58
	VMSizeTypesStandardH16         VMSizeTypes = original.VMSizeTypesStandardH16
	VMSizeTypesStandardH16m        VMSizeTypes = original.VMSizeTypesStandardH16m
	VMSizeTypesStandardH16mr       VMSizeTypes = original.VMSizeTypesStandardH16mr
	VMSizeTypesStandardH16r        VMSizeTypes = original.VMSizeTypesStandardH16r
	VMSizeTypesStandardH8          VMSizeTypes = original.VMSizeTypesStandardH8
	VMSizeTypesStandardH8m         VMSizeTypes = original.VMSizeTypesStandardH8m
	VMSizeTypesStandardL16s        VMSizeTypes = original.VMSizeTypesStandardL16s
	VMSizeTypesStandardL32s        VMSizeTypes = original.VMSizeTypesStandardL32s
	VMSizeTypesStandardL4s         VMSizeTypes = original.VMSizeTypesStandardL4s
	VMSizeTypesStandardL8s         VMSizeTypes = original.VMSizeTypesStandardL8s
	VMSizeTypesStandardM12832ms    VMSizeTypes = original.VMSizeTypesStandardM12832ms
	VMSizeTypesStandardM12864ms    VMSizeTypes = original.VMSizeTypesStandardM12864ms
	VMSizeTypesStandardM128ms      VMSizeTypes = original.VMSizeTypesStandardM128ms
	VMSizeTypesStandardM128s       VMSizeTypes = original.VMSizeTypesStandardM128s
	VMSizeTypesStandardM6416ms     VMSizeTypes = original.VMSizeTypesStandardM6416ms
	VMSizeTypesStandardM6432ms     VMSizeTypes = original.VMSizeTypesStandardM6432ms
	VMSizeTypesStandardM64ms       VMSizeTypes = original.VMSizeTypesStandardM64ms
	VMSizeTypesStandardM64s        VMSizeTypes = original.VMSizeTypesStandardM64s
	VMSizeTypesStandardNC12        VMSizeTypes = original.VMSizeTypesStandardNC12
	VMSizeTypesStandardNC12sV2     VMSizeTypes = original.VMSizeTypesStandardNC12sV2
	VMSizeTypesStandardNC12sV3     VMSizeTypes = original.VMSizeTypesStandardNC12sV3
	VMSizeTypesStandardNC24        VMSizeTypes = original.VMSizeTypesStandardNC24
	VMSizeTypesStandardNC24r       VMSizeTypes = original.VMSizeTypesStandardNC24r
	VMSizeTypesStandardNC24rsV2    VMSizeTypes = original.VMSizeTypesStandardNC24rsV2
	VMSizeTypesStandardNC24rsV3    VMSizeTypes = original.VMSizeTypesStandardNC24rsV3
	VMSizeTypesStandardNC24sV2     VMSizeTypes = original.VMSizeTypesStandardNC24sV2
	VMSizeTypesStandardNC24sV3     VMSizeTypes = original.VMSizeTypesStandardNC24sV3
	VMSizeTypesStandardNC6         VMSizeTypes = original.VMSizeTypesStandardNC6
	VMSizeTypesStandardNC6sV2      VMSizeTypes = original.VMSizeTypesStandardNC6sV2
	VMSizeTypesStandardNC6sV3      VMSizeTypes = original.VMSizeTypesStandardNC6sV3
	VMSizeTypesStandardND12s       VMSizeTypes = original.VMSizeTypesStandardND12s
	VMSizeTypesStandardND24rs      VMSizeTypes = original.VMSizeTypesStandardND24rs
	VMSizeTypesStandardND24s       VMSizeTypes = original.VMSizeTypesStandardND24s
	VMSizeTypesStandardND6s        VMSizeTypes = original.VMSizeTypesStandardND6s
	VMSizeTypesStandardNV12        VMSizeTypes = original.VMSizeTypesStandardNV12
	VMSizeTypesStandardNV24        VMSizeTypes = original.VMSizeTypesStandardNV24
	VMSizeTypesStandardNV6         VMSizeTypes = original.VMSizeTypesStandardNV6
)

type WeekDay = original.WeekDay

const (
	WeekDayFriday    WeekDay = original.WeekDayFriday
	WeekDayMonday    WeekDay = original.WeekDayMonday
	WeekDaySaturday  WeekDay = original.WeekDaySaturday
	WeekDaySunday    WeekDay = original.WeekDaySunday
	WeekDayThursday  WeekDay = original.WeekDayThursday
	WeekDayTuesday   WeekDay = original.WeekDayTuesday
	WeekDayWednesday WeekDay = original.WeekDayWednesday
)

type AccessProfile = original.AccessProfile
type AgentPool = original.AgentPool
type AgentPoolAvailableVersions = original.AgentPoolAvailableVersions
type AgentPoolAvailableVersionsProperties = original.AgentPoolAvailableVersionsProperties
type AgentPoolAvailableVersionsPropertiesAgentPoolVersionsItem = original.AgentPoolAvailableVersionsPropertiesAgentPoolVersionsItem
type AgentPoolListResult = original.AgentPoolListResult
type AgentPoolListResultIterator = original.AgentPoolListResultIterator
type AgentPoolListResultPage = original.AgentPoolListResultPage
type AgentPoolUpgradeProfile = original.AgentPoolUpgradeProfile
type AgentPoolUpgradeProfileProperties = original.AgentPoolUpgradeProfileProperties
type AgentPoolUpgradeProfilePropertiesUpgradesItem = original.AgentPoolUpgradeProfilePropertiesUpgradesItem
type AgentPoolUpgradeSettings = original.AgentPoolUpgradeSettings
type AgentPoolsClient = original.AgentPoolsClient
type AgentPoolsCreateOrUpdateFuture = original.AgentPoolsCreateOrUpdateFuture
type AgentPoolsDeleteFuture = original.AgentPoolsDeleteFuture
type AgentPoolsUpgradeNodeImageVersionFuture = original.AgentPoolsUpgradeNodeImageVersionFuture
type BaseClient = original.BaseClient
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type CommandResultProperties = original.CommandResultProperties
type CredentialResult = original.CredentialResult
type CredentialResults = original.CredentialResults
type DiagnosticsProfile = original.DiagnosticsProfile
type EndpointDependency = original.EndpointDependency
type EndpointDetail = original.EndpointDetail
type ExtendedLocation = original.ExtendedLocation
type KubeletConfig = original.KubeletConfig
type LinuxOSConfig = original.LinuxOSConfig
type LinuxProfile = original.LinuxProfile
type MaintenanceConfiguration = original.MaintenanceConfiguration
type MaintenanceConfigurationListResult = original.MaintenanceConfigurationListResult
type MaintenanceConfigurationListResultIterator = original.MaintenanceConfigurationListResultIterator
type MaintenanceConfigurationListResultPage = original.MaintenanceConfigurationListResultPage
type MaintenanceConfigurationProperties = original.MaintenanceConfigurationProperties
type MaintenanceConfigurationsClient = original.MaintenanceConfigurationsClient
type ManagedCluster = original.ManagedCluster
type ManagedClusterAADProfile = original.ManagedClusterAADProfile
type ManagedClusterAPIServerAccessProfile = original.ManagedClusterAPIServerAccessProfile
type ManagedClusterAccessProfile = original.ManagedClusterAccessProfile
type ManagedClusterAddonProfile = original.ManagedClusterAddonProfile
type ManagedClusterAddonProfileIdentity = original.ManagedClusterAddonProfileIdentity
type ManagedClusterAgentPoolProfile = original.ManagedClusterAgentPoolProfile
type ManagedClusterAgentPoolProfileProperties = original.ManagedClusterAgentPoolProfileProperties
type ManagedClusterAutoUpgradeProfile = original.ManagedClusterAutoUpgradeProfile
type ManagedClusterHTTPProxyConfig = original.ManagedClusterHTTPProxyConfig
type ManagedClusterIdentity = original.ManagedClusterIdentity
type ManagedClusterIdentityUserAssignedIdentitiesValue = original.ManagedClusterIdentityUserAssignedIdentitiesValue
type ManagedClusterListResult = original.ManagedClusterListResult
type ManagedClusterListResultIterator = original.ManagedClusterListResultIterator
type ManagedClusterListResultPage = original.ManagedClusterListResultPage
type ManagedClusterLoadBalancerProfile = original.ManagedClusterLoadBalancerProfile
type ManagedClusterLoadBalancerProfileManagedOutboundIPs = original.ManagedClusterLoadBalancerProfileManagedOutboundIPs
type ManagedClusterLoadBalancerProfileOutboundIPPrefixes = original.ManagedClusterLoadBalancerProfileOutboundIPPrefixes
type ManagedClusterLoadBalancerProfileOutboundIPs = original.ManagedClusterLoadBalancerProfileOutboundIPs
type ManagedClusterManagedOutboundIPProfile = original.ManagedClusterManagedOutboundIPProfile
type ManagedClusterNATGatewayProfile = original.ManagedClusterNATGatewayProfile
type ManagedClusterPodIdentity = original.ManagedClusterPodIdentity
type ManagedClusterPodIdentityException = original.ManagedClusterPodIdentityException
type ManagedClusterPodIdentityProfile = original.ManagedClusterPodIdentityProfile
type ManagedClusterPodIdentityProvisioningError = original.ManagedClusterPodIdentityProvisioningError
type ManagedClusterPodIdentityProvisioningErrorBody = original.ManagedClusterPodIdentityProvisioningErrorBody
type ManagedClusterPodIdentityProvisioningInfo = original.ManagedClusterPodIdentityProvisioningInfo
type ManagedClusterPoolUpgradeProfile = original.ManagedClusterPoolUpgradeProfile
type ManagedClusterPoolUpgradeProfileUpgradesItem = original.ManagedClusterPoolUpgradeProfileUpgradesItem
type ManagedClusterProperties = original.ManagedClusterProperties
type ManagedClusterPropertiesAutoScalerProfile = original.ManagedClusterPropertiesAutoScalerProfile
type ManagedClusterSKU = original.ManagedClusterSKU
type ManagedClusterSecurityProfile = original.ManagedClusterSecurityProfile
type ManagedClusterSecurityProfileAzureDefender = original.ManagedClusterSecurityProfileAzureDefender
type ManagedClusterServicePrincipalProfile = original.ManagedClusterServicePrincipalProfile
type ManagedClusterUpgradeProfile = original.ManagedClusterUpgradeProfile
type ManagedClusterUpgradeProfileProperties = original.ManagedClusterUpgradeProfileProperties
type ManagedClusterWindowsProfile = original.ManagedClusterWindowsProfile
type ManagedClustersClient = original.ManagedClustersClient
type ManagedClustersCreateOrUpdateFuture = original.ManagedClustersCreateOrUpdateFuture
type ManagedClustersDeleteFuture = original.ManagedClustersDeleteFuture
type ManagedClustersResetAADProfileFuture = original.ManagedClustersResetAADProfileFuture
type ManagedClustersResetServicePrincipalProfileFuture = original.ManagedClustersResetServicePrincipalProfileFuture
type ManagedClustersRotateClusterCertificatesFuture = original.ManagedClustersRotateClusterCertificatesFuture
type ManagedClustersRunCommandFuture = original.ManagedClustersRunCommandFuture
type ManagedClustersStartFuture = original.ManagedClustersStartFuture
type ManagedClustersStopFuture = original.ManagedClustersStopFuture
type ManagedClustersUpdateTagsFuture = original.ManagedClustersUpdateTagsFuture
type MasterProfile = original.MasterProfile
type NetworkProfile = original.NetworkProfile
type OSOptionProfile = original.OSOptionProfile
type OSOptionProperty = original.OSOptionProperty
type OSOptionPropertyList = original.OSOptionPropertyList
type OperationListResult = original.OperationListResult
type OperationValue = original.OperationValue
type OperationValueDisplay = original.OperationValueDisplay
type OperationsClient = original.OperationsClient
type OutboundEnvironmentEndpoint = original.OutboundEnvironmentEndpoint
type OutboundEnvironmentEndpointCollection = original.OutboundEnvironmentEndpointCollection
type OutboundEnvironmentEndpointCollectionIterator = original.OutboundEnvironmentEndpointCollectionIterator
type OutboundEnvironmentEndpointCollectionPage = original.OutboundEnvironmentEndpointCollectionPage
type PowerState = original.PowerState
type PrivateEndpoint = original.PrivateEndpoint
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionListResult = original.PrivateEndpointConnectionListResult
type PrivateEndpointConnectionProperties = original.PrivateEndpointConnectionProperties
type PrivateEndpointConnectionsClient = original.PrivateEndpointConnectionsClient
type PrivateEndpointConnectionsDeleteFuture = original.PrivateEndpointConnectionsDeleteFuture
type PrivateLinkResource = original.PrivateLinkResource
type PrivateLinkResourcesClient = original.PrivateLinkResourcesClient
type PrivateLinkResourcesListResult = original.PrivateLinkResourcesListResult
type PrivateLinkServiceConnectionState = original.PrivateLinkServiceConnectionState
type ResolvePrivateLinkServiceIDClient = original.ResolvePrivateLinkServiceIDClient
type Resource = original.Resource
type ResourceReference = original.ResourceReference
type RunCommandRequest = original.RunCommandRequest
type RunCommandResult = original.RunCommandResult
type SSHConfiguration = original.SSHConfiguration
type SSHPublicKey = original.SSHPublicKey
type SubResource = original.SubResource
type SysctlConfig = original.SysctlConfig
type SystemData = original.SystemData
type TagsObject = original.TagsObject
type TimeInWeek = original.TimeInWeek
type TimeSpan = original.TimeSpan
type UserAssignedIdentity = original.UserAssignedIdentity
type VMDiagnostics = original.VMDiagnostics

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAgentPoolListResultIterator(page AgentPoolListResultPage) AgentPoolListResultIterator {
	return original.NewAgentPoolListResultIterator(page)
}
func NewAgentPoolListResultPage(cur AgentPoolListResult, getNextPage func(context.Context, AgentPoolListResult) (AgentPoolListResult, error)) AgentPoolListResultPage {
	return original.NewAgentPoolListResultPage(cur, getNextPage)
}
func NewAgentPoolsClient(subscriptionID string) AgentPoolsClient {
	return original.NewAgentPoolsClient(subscriptionID)
}
func NewAgentPoolsClientWithBaseURI(baseURI string, subscriptionID string) AgentPoolsClient {
	return original.NewAgentPoolsClientWithBaseURI(baseURI, subscriptionID)
}
func NewMaintenanceConfigurationListResultIterator(page MaintenanceConfigurationListResultPage) MaintenanceConfigurationListResultIterator {
	return original.NewMaintenanceConfigurationListResultIterator(page)
}
func NewMaintenanceConfigurationListResultPage(cur MaintenanceConfigurationListResult, getNextPage func(context.Context, MaintenanceConfigurationListResult) (MaintenanceConfigurationListResult, error)) MaintenanceConfigurationListResultPage {
	return original.NewMaintenanceConfigurationListResultPage(cur, getNextPage)
}
func NewMaintenanceConfigurationsClient(subscriptionID string) MaintenanceConfigurationsClient {
	return original.NewMaintenanceConfigurationsClient(subscriptionID)
}
func NewMaintenanceConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) MaintenanceConfigurationsClient {
	return original.NewMaintenanceConfigurationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewManagedClusterListResultIterator(page ManagedClusterListResultPage) ManagedClusterListResultIterator {
	return original.NewManagedClusterListResultIterator(page)
}
func NewManagedClusterListResultPage(cur ManagedClusterListResult, getNextPage func(context.Context, ManagedClusterListResult) (ManagedClusterListResult, error)) ManagedClusterListResultPage {
	return original.NewManagedClusterListResultPage(cur, getNextPage)
}
func NewManagedClustersClient(subscriptionID string) ManagedClustersClient {
	return original.NewManagedClustersClient(subscriptionID)
}
func NewManagedClustersClientWithBaseURI(baseURI string, subscriptionID string) ManagedClustersClient {
	return original.NewManagedClustersClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOutboundEnvironmentEndpointCollectionIterator(page OutboundEnvironmentEndpointCollectionPage) OutboundEnvironmentEndpointCollectionIterator {
	return original.NewOutboundEnvironmentEndpointCollectionIterator(page)
}
func NewOutboundEnvironmentEndpointCollectionPage(cur OutboundEnvironmentEndpointCollection, getNextPage func(context.Context, OutboundEnvironmentEndpointCollection) (OutboundEnvironmentEndpointCollection, error)) OutboundEnvironmentEndpointCollectionPage {
	return original.NewOutboundEnvironmentEndpointCollectionPage(cur, getNextPage)
}
func NewPrivateEndpointConnectionsClient(subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClient(subscriptionID)
}
func NewPrivateEndpointConnectionsClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateLinkResourcesClient(subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClient(subscriptionID)
}
func NewPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewResolvePrivateLinkServiceIDClient(subscriptionID string) ResolvePrivateLinkServiceIDClient {
	return original.NewResolvePrivateLinkServiceIDClient(subscriptionID)
}
func NewResolvePrivateLinkServiceIDClientWithBaseURI(baseURI string, subscriptionID string) ResolvePrivateLinkServiceIDClient {
	return original.NewResolvePrivateLinkServiceIDClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAgentPoolModeValues() []AgentPoolMode {
	return original.PossibleAgentPoolModeValues()
}
func PossibleAgentPoolTypeValues() []AgentPoolType {
	return original.PossibleAgentPoolTypeValues()
}
func PossibleCodeValues() []Code {
	return original.PossibleCodeValues()
}
func PossibleConnectionStatusValues() []ConnectionStatus {
	return original.PossibleConnectionStatusValues()
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleExpanderValues() []Expander {
	return original.PossibleExpanderValues()
}
func PossibleExtendedLocationTypesValues() []ExtendedLocationTypes {
	return original.PossibleExtendedLocationTypesValues()
}
func PossibleGPUInstanceProfileValues() []GPUInstanceProfile {
	return original.PossibleGPUInstanceProfileValues()
}
func PossibleKubeletDiskTypeValues() []KubeletDiskType {
	return original.PossibleKubeletDiskTypeValues()
}
func PossibleLicenseTypeValues() []LicenseType {
	return original.PossibleLicenseTypeValues()
}
func PossibleLoadBalancerSkuValues() []LoadBalancerSku {
	return original.PossibleLoadBalancerSkuValues()
}
func PossibleManagedClusterPodIdentityProvisioningStateValues() []ManagedClusterPodIdentityProvisioningState {
	return original.PossibleManagedClusterPodIdentityProvisioningStateValues()
}
func PossibleManagedClusterSKUNameValues() []ManagedClusterSKUName {
	return original.PossibleManagedClusterSKUNameValues()
}
func PossibleManagedClusterSKUTierValues() []ManagedClusterSKUTier {
	return original.PossibleManagedClusterSKUTierValues()
}
func PossibleNetworkModeValues() []NetworkMode {
	return original.PossibleNetworkModeValues()
}
func PossibleNetworkPluginValues() []NetworkPlugin {
	return original.PossibleNetworkPluginValues()
}
func PossibleNetworkPolicyValues() []NetworkPolicy {
	return original.PossibleNetworkPolicyValues()
}
func PossibleOSDiskTypeValues() []OSDiskType {
	return original.PossibleOSDiskTypeValues()
}
func PossibleOSSKUValues() []OSSKU {
	return original.PossibleOSSKUValues()
}
func PossibleOSTypeValues() []OSType {
	return original.PossibleOSTypeValues()
}
func PossibleOutboundTypeValues() []OutboundType {
	return original.PossibleOutboundTypeValues()
}
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return original.PossiblePrivateEndpointConnectionProvisioningStateValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func PossibleScaleDownModeValues() []ScaleDownMode {
	return original.PossibleScaleDownModeValues()
}
func PossibleScaleSetEvictionPolicyValues() []ScaleSetEvictionPolicy {
	return original.PossibleScaleSetEvictionPolicyValues()
}
func PossibleScaleSetPriorityValues() []ScaleSetPriority {
	return original.PossibleScaleSetPriorityValues()
}
func PossibleStorageProfileTypesValues() []StorageProfileTypes {
	return original.PossibleStorageProfileTypesValues()
}
func PossibleUpgradeChannelValues() []UpgradeChannel {
	return original.PossibleUpgradeChannelValues()
}
func PossibleVMSizeTypesValues() []VMSizeTypes {
	return original.PossibleVMSizeTypesValues()
}
func PossibleWeekDayValues() []WeekDay {
	return original.PossibleWeekDayValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
