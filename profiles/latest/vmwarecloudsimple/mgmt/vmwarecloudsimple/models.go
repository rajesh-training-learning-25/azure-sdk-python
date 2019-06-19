// +build go1.9

// Copyright 2019 Microsoft Corporation
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

package vmwarecloudsimple

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/vmwarecloudsimple/mgmt/2019-04-01/vmwarecloudsimple"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AggregationType = original.AggregationType

const (
	Average AggregationType = original.Average
	Total   AggregationType = original.Total
)

type DiskIndependenceMode = original.DiskIndependenceMode

const (
	IndependentNonpersistent DiskIndependenceMode = original.IndependentNonpersistent
	IndependentPersistent    DiskIndependenceMode = original.IndependentPersistent
	Persistent               DiskIndependenceMode = original.Persistent
)

type GuestOSType = original.GuestOSType

const (
	Linux   GuestOSType = original.Linux
	Other   GuestOSType = original.Other
	Windows GuestOSType = original.Windows
)

type NICType = original.NICType

const (
	E1000   NICType = original.E1000
	E1000E  NICType = original.E1000E
	PCNET32 NICType = original.PCNET32
	VMXNET  NICType = original.VMXNET
	VMXNET2 NICType = original.VMXNET2
	VMXNET3 NICType = original.VMXNET3
)

type NodeStatus = original.NodeStatus

const (
	Unused NodeStatus = original.Unused
	Used   NodeStatus = original.Used
)

type OnboardingStatus = original.OnboardingStatus

const (
	NotOnBoarded     OnboardingStatus = original.NotOnBoarded
	OnBoarded        OnboardingStatus = original.OnBoarded
	OnBoarding       OnboardingStatus = original.OnBoarding
	OnBoardingFailed OnboardingStatus = original.OnBoardingFailed
)

type OperationOrigin = original.OperationOrigin

const (
	System     OperationOrigin = original.System
	User       OperationOrigin = original.User
	Usersystem OperationOrigin = original.Usersystem
)

type PrivateCloudResourceType = original.PrivateCloudResourceType

const (
	MicrosoftVMwareCloudSimpleprivateClouds PrivateCloudResourceType = original.MicrosoftVMwareCloudSimpleprivateClouds
)

type StopMode = original.StopMode

const (
	Poweroff StopMode = original.Poweroff
	Reboot   StopMode = original.Reboot
	Shutdown StopMode = original.Shutdown
	Suspend  StopMode = original.Suspend
)

type UsageCount = original.UsageCount

const (
	Bytes          UsageCount = original.Bytes
	BytesPerSecond UsageCount = original.BytesPerSecond
	Count          UsageCount = original.Count
	CountPerSecond UsageCount = original.CountPerSecond
	Percent        UsageCount = original.Percent
	Seconds        UsageCount = original.Seconds
)

type VirtualMachineStatus = original.VirtualMachineStatus

const (
	Deallocating VirtualMachineStatus = original.Deallocating
	Deleting     VirtualMachineStatus = original.Deleting
	Poweredoff   VirtualMachineStatus = original.Poweredoff
	Running      VirtualMachineStatus = original.Running
	Suspended    VirtualMachineStatus = original.Suspended
	Updating     VirtualMachineStatus = original.Updating
)

type AvailableOperation = original.AvailableOperation
type AvailableOperationDisplay = original.AvailableOperationDisplay
type AvailableOperationDisplayPropertyServiceSpecification = original.AvailableOperationDisplayPropertyServiceSpecification
type AvailableOperationDisplayPropertyServiceSpecificationMetricsItem = original.AvailableOperationDisplayPropertyServiceSpecificationMetricsItem
type AvailableOperationDisplayPropertyServiceSpecificationMetricsList = original.AvailableOperationDisplayPropertyServiceSpecificationMetricsList
type AvailableOperationsClient = original.AvailableOperationsClient
type AvailableOperationsListResponse = original.AvailableOperationsListResponse
type AvailableOperationsListResponseIterator = original.AvailableOperationsListResponseIterator
type AvailableOperationsListResponsePage = original.AvailableOperationsListResponsePage
type BaseClient = original.BaseClient
type CSRPError = original.CSRPError
type CSRPErrorBody = original.CSRPErrorBody
type DedicatedCloudNode = original.DedicatedCloudNode
type DedicatedCloudNodeClient = original.DedicatedCloudNodeClient
type DedicatedCloudNodeCreateOrUpdateFuture = original.DedicatedCloudNodeCreateOrUpdateFuture
type DedicatedCloudNodeListResponse = original.DedicatedCloudNodeListResponse
type DedicatedCloudNodeListResponseIterator = original.DedicatedCloudNodeListResponseIterator
type DedicatedCloudNodeListResponsePage = original.DedicatedCloudNodeListResponsePage
type DedicatedCloudNodeProperties = original.DedicatedCloudNodeProperties
type DedicatedCloudService = original.DedicatedCloudService
type DedicatedCloudServiceClient = original.DedicatedCloudServiceClient
type DedicatedCloudServiceDeleteFuture = original.DedicatedCloudServiceDeleteFuture
type DedicatedCloudServiceListResponse = original.DedicatedCloudServiceListResponse
type DedicatedCloudServiceListResponseIterator = original.DedicatedCloudServiceListResponseIterator
type DedicatedCloudServiceListResponsePage = original.DedicatedCloudServiceListResponsePage
type DedicatedCloudServiceProperties = original.DedicatedCloudServiceProperties
type GetOperationResultByRegionFuture = original.GetOperationResultByRegionFuture
type OperationError = original.OperationError
type OperationResource = original.OperationResource
type PatchPayload = original.PatchPayload
type PrivateCloud = original.PrivateCloud
type PrivateCloudByRegionClient = original.PrivateCloudByRegionClient
type PrivateCloudList = original.PrivateCloudList
type PrivateCloudListIterator = original.PrivateCloudListIterator
type PrivateCloudListPage = original.PrivateCloudListPage
type PrivateCloudProperties = original.PrivateCloudProperties
type ResourcePool = original.ResourcePool
type ResourcePoolProperties = original.ResourcePoolProperties
type ResourcePoolsListResponse = original.ResourcePoolsListResponse
type ResourcePoolsListResponseIterator = original.ResourcePoolsListResponseIterator
type ResourcePoolsListResponsePage = original.ResourcePoolsListResponsePage
type ResourcepoolByPCClient = original.ResourcepoolByPCClient
type ResourcepoolsByPCClient = original.ResourcepoolsByPCClient
type Sku = original.Sku
type SkuAvailability = original.SkuAvailability
type SkuAvailabilityListResponse = original.SkuAvailabilityListResponse
type SkuAvailabilityListResponseIterator = original.SkuAvailabilityListResponseIterator
type SkuAvailabilityListResponsePage = original.SkuAvailabilityListResponsePage
type SkuDescription = original.SkuDescription
type SkusAvailabilityWithinRegionClient = original.SkusAvailabilityWithinRegionClient
type Usage = original.Usage
type UsageListResponse = original.UsageListResponse
type UsageListResponseIterator = original.UsageListResponseIterator
type UsageListResponsePage = original.UsageListResponsePage
type UsageName = original.UsageName
type UsagesWithinRegionClient = original.UsagesWithinRegionClient
type VirtualDisk = original.VirtualDisk
type VirtualDiskController = original.VirtualDiskController
type VirtualMachine = original.VirtualMachine
type VirtualMachineClient = original.VirtualMachineClient
type VirtualMachineCreateOrUpdateFuture = original.VirtualMachineCreateOrUpdateFuture
type VirtualMachineDeleteFuture = original.VirtualMachineDeleteFuture
type VirtualMachineListResponse = original.VirtualMachineListResponse
type VirtualMachineListResponseIterator = original.VirtualMachineListResponseIterator
type VirtualMachineListResponsePage = original.VirtualMachineListResponsePage
type VirtualMachineProperties = original.VirtualMachineProperties
type VirtualMachineStartFuture = original.VirtualMachineStartFuture
type VirtualMachineStopFuture = original.VirtualMachineStopFuture
type VirtualMachineStopMode = original.VirtualMachineStopMode
type VirtualMachineTemplate = original.VirtualMachineTemplate
type VirtualMachineTemplateByPCClient = original.VirtualMachineTemplateByPCClient
type VirtualMachineTemplateListResponse = original.VirtualMachineTemplateListResponse
type VirtualMachineTemplateListResponseIterator = original.VirtualMachineTemplateListResponseIterator
type VirtualMachineTemplateListResponsePage = original.VirtualMachineTemplateListResponsePage
type VirtualMachineTemplateProperties = original.VirtualMachineTemplateProperties
type VirtualMachineTemplatesByPCClient = original.VirtualMachineTemplatesByPCClient
type VirtualMachineUpdateFuture = original.VirtualMachineUpdateFuture
type VirtualNetwork = original.VirtualNetwork
type VirtualNetworkByPCClient = original.VirtualNetworkByPCClient
type VirtualNetworkListResponse = original.VirtualNetworkListResponse
type VirtualNetworkListResponseIterator = original.VirtualNetworkListResponseIterator
type VirtualNetworkListResponsePage = original.VirtualNetworkListResponsePage
type VirtualNetworkProperties = original.VirtualNetworkProperties
type VirtualNetworksByPCClient = original.VirtualNetworksByPCClient
type VirtualNic = original.VirtualNic

func New(referer string, regionID string, subscriptionID string) BaseClient {
	return original.New(referer, regionID, subscriptionID)
}
func NewAvailableOperationsClient(referer string, regionID string, subscriptionID string) AvailableOperationsClient {
	return original.NewAvailableOperationsClient(referer, regionID, subscriptionID)
}
func NewAvailableOperationsClientWithBaseURI(baseURI string, referer string, regionID string, subscriptionID string) AvailableOperationsClient {
	return original.NewAvailableOperationsClientWithBaseURI(baseURI, referer, regionID, subscriptionID)
}
func NewAvailableOperationsListResponseIterator(page AvailableOperationsListResponsePage) AvailableOperationsListResponseIterator {
	return original.NewAvailableOperationsListResponseIterator(page)
}
func NewAvailableOperationsListResponsePage(getNextPage func(context.Context, AvailableOperationsListResponse) (AvailableOperationsListResponse, error)) AvailableOperationsListResponsePage {
	return original.NewAvailableOperationsListResponsePage(getNextPage)
}
func NewDedicatedCloudNodeClient(referer string, regionID string, subscriptionID string) DedicatedCloudNodeClient {
	return original.NewDedicatedCloudNodeClient(referer, regionID, subscriptionID)
}
func NewDedicatedCloudNodeClientWithBaseURI(baseURI string, referer string, regionID string, subscriptionID string) DedicatedCloudNodeClient {
	return original.NewDedicatedCloudNodeClientWithBaseURI(baseURI, referer, regionID, subscriptionID)
}
func NewDedicatedCloudNodeListResponseIterator(page DedicatedCloudNodeListResponsePage) DedicatedCloudNodeListResponseIterator {
	return original.NewDedicatedCloudNodeListResponseIterator(page)
}
func NewDedicatedCloudNodeListResponsePage(getNextPage func(context.Context, DedicatedCloudNodeListResponse) (DedicatedCloudNodeListResponse, error)) DedicatedCloudNodeListResponsePage {
	return original.NewDedicatedCloudNodeListResponsePage(getNextPage)
}
func NewDedicatedCloudServiceClient(referer string, regionID string, subscriptionID string) DedicatedCloudServiceClient {
	return original.NewDedicatedCloudServiceClient(referer, regionID, subscriptionID)
}
func NewDedicatedCloudServiceClientWithBaseURI(baseURI string, referer string, regionID string, subscriptionID string) DedicatedCloudServiceClient {
	return original.NewDedicatedCloudServiceClientWithBaseURI(baseURI, referer, regionID, subscriptionID)
}
func NewDedicatedCloudServiceListResponseIterator(page DedicatedCloudServiceListResponsePage) DedicatedCloudServiceListResponseIterator {
	return original.NewDedicatedCloudServiceListResponseIterator(page)
}
func NewDedicatedCloudServiceListResponsePage(getNextPage func(context.Context, DedicatedCloudServiceListResponse) (DedicatedCloudServiceListResponse, error)) DedicatedCloudServiceListResponsePage {
	return original.NewDedicatedCloudServiceListResponsePage(getNextPage)
}
func NewPrivateCloudByRegionClient(referer string, regionID string, subscriptionID string) PrivateCloudByRegionClient {
	return original.NewPrivateCloudByRegionClient(referer, regionID, subscriptionID)
}
func NewPrivateCloudByRegionClientWithBaseURI(baseURI string, referer string, regionID string, subscriptionID string) PrivateCloudByRegionClient {
	return original.NewPrivateCloudByRegionClientWithBaseURI(baseURI, referer, regionID, subscriptionID)
}
func NewPrivateCloudListIterator(page PrivateCloudListPage) PrivateCloudListIterator {
	return original.NewPrivateCloudListIterator(page)
}
func NewPrivateCloudListPage(getNextPage func(context.Context, PrivateCloudList) (PrivateCloudList, error)) PrivateCloudListPage {
	return original.NewPrivateCloudListPage(getNextPage)
}
func NewResourcePoolsListResponseIterator(page ResourcePoolsListResponsePage) ResourcePoolsListResponseIterator {
	return original.NewResourcePoolsListResponseIterator(page)
}
func NewResourcePoolsListResponsePage(getNextPage func(context.Context, ResourcePoolsListResponse) (ResourcePoolsListResponse, error)) ResourcePoolsListResponsePage {
	return original.NewResourcePoolsListResponsePage(getNextPage)
}
func NewResourcepoolByPCClient(referer string, regionID string, subscriptionID string) ResourcepoolByPCClient {
	return original.NewResourcepoolByPCClient(referer, regionID, subscriptionID)
}
func NewResourcepoolByPCClientWithBaseURI(baseURI string, referer string, regionID string, subscriptionID string) ResourcepoolByPCClient {
	return original.NewResourcepoolByPCClientWithBaseURI(baseURI, referer, regionID, subscriptionID)
}
func NewResourcepoolsByPCClient(referer string, regionID string, subscriptionID string) ResourcepoolsByPCClient {
	return original.NewResourcepoolsByPCClient(referer, regionID, subscriptionID)
}
func NewResourcepoolsByPCClientWithBaseURI(baseURI string, referer string, regionID string, subscriptionID string) ResourcepoolsByPCClient {
	return original.NewResourcepoolsByPCClientWithBaseURI(baseURI, referer, regionID, subscriptionID)
}
func NewSkuAvailabilityListResponseIterator(page SkuAvailabilityListResponsePage) SkuAvailabilityListResponseIterator {
	return original.NewSkuAvailabilityListResponseIterator(page)
}
func NewSkuAvailabilityListResponsePage(getNextPage func(context.Context, SkuAvailabilityListResponse) (SkuAvailabilityListResponse, error)) SkuAvailabilityListResponsePage {
	return original.NewSkuAvailabilityListResponsePage(getNextPage)
}
func NewSkusAvailabilityWithinRegionClient(referer string, regionID string, subscriptionID string) SkusAvailabilityWithinRegionClient {
	return original.NewSkusAvailabilityWithinRegionClient(referer, regionID, subscriptionID)
}
func NewSkusAvailabilityWithinRegionClientWithBaseURI(baseURI string, referer string, regionID string, subscriptionID string) SkusAvailabilityWithinRegionClient {
	return original.NewSkusAvailabilityWithinRegionClientWithBaseURI(baseURI, referer, regionID, subscriptionID)
}
func NewUsageListResponseIterator(page UsageListResponsePage) UsageListResponseIterator {
	return original.NewUsageListResponseIterator(page)
}
func NewUsageListResponsePage(getNextPage func(context.Context, UsageListResponse) (UsageListResponse, error)) UsageListResponsePage {
	return original.NewUsageListResponsePage(getNextPage)
}
func NewUsagesWithinRegionClient(referer string, regionID string, subscriptionID string) UsagesWithinRegionClient {
	return original.NewUsagesWithinRegionClient(referer, regionID, subscriptionID)
}
func NewUsagesWithinRegionClientWithBaseURI(baseURI string, referer string, regionID string, subscriptionID string) UsagesWithinRegionClient {
	return original.NewUsagesWithinRegionClientWithBaseURI(baseURI, referer, regionID, subscriptionID)
}
func NewVirtualMachineClient(referer string, regionID string, subscriptionID string) VirtualMachineClient {
	return original.NewVirtualMachineClient(referer, regionID, subscriptionID)
}
func NewVirtualMachineClientWithBaseURI(baseURI string, referer string, regionID string, subscriptionID string) VirtualMachineClient {
	return original.NewVirtualMachineClientWithBaseURI(baseURI, referer, regionID, subscriptionID)
}
func NewVirtualMachineListResponseIterator(page VirtualMachineListResponsePage) VirtualMachineListResponseIterator {
	return original.NewVirtualMachineListResponseIterator(page)
}
func NewVirtualMachineListResponsePage(getNextPage func(context.Context, VirtualMachineListResponse) (VirtualMachineListResponse, error)) VirtualMachineListResponsePage {
	return original.NewVirtualMachineListResponsePage(getNextPage)
}
func NewVirtualMachineTemplateByPCClient(referer string, regionID string, subscriptionID string) VirtualMachineTemplateByPCClient {
	return original.NewVirtualMachineTemplateByPCClient(referer, regionID, subscriptionID)
}
func NewVirtualMachineTemplateByPCClientWithBaseURI(baseURI string, referer string, regionID string, subscriptionID string) VirtualMachineTemplateByPCClient {
	return original.NewVirtualMachineTemplateByPCClientWithBaseURI(baseURI, referer, regionID, subscriptionID)
}
func NewVirtualMachineTemplateListResponseIterator(page VirtualMachineTemplateListResponsePage) VirtualMachineTemplateListResponseIterator {
	return original.NewVirtualMachineTemplateListResponseIterator(page)
}
func NewVirtualMachineTemplateListResponsePage(getNextPage func(context.Context, VirtualMachineTemplateListResponse) (VirtualMachineTemplateListResponse, error)) VirtualMachineTemplateListResponsePage {
	return original.NewVirtualMachineTemplateListResponsePage(getNextPage)
}
func NewVirtualMachineTemplatesByPCClient(referer string, regionID string, subscriptionID string) VirtualMachineTemplatesByPCClient {
	return original.NewVirtualMachineTemplatesByPCClient(referer, regionID, subscriptionID)
}
func NewVirtualMachineTemplatesByPCClientWithBaseURI(baseURI string, referer string, regionID string, subscriptionID string) VirtualMachineTemplatesByPCClient {
	return original.NewVirtualMachineTemplatesByPCClientWithBaseURI(baseURI, referer, regionID, subscriptionID)
}
func NewVirtualNetworkByPCClient(referer string, regionID string, subscriptionID string) VirtualNetworkByPCClient {
	return original.NewVirtualNetworkByPCClient(referer, regionID, subscriptionID)
}
func NewVirtualNetworkByPCClientWithBaseURI(baseURI string, referer string, regionID string, subscriptionID string) VirtualNetworkByPCClient {
	return original.NewVirtualNetworkByPCClientWithBaseURI(baseURI, referer, regionID, subscriptionID)
}
func NewVirtualNetworkListResponseIterator(page VirtualNetworkListResponsePage) VirtualNetworkListResponseIterator {
	return original.NewVirtualNetworkListResponseIterator(page)
}
func NewVirtualNetworkListResponsePage(getNextPage func(context.Context, VirtualNetworkListResponse) (VirtualNetworkListResponse, error)) VirtualNetworkListResponsePage {
	return original.NewVirtualNetworkListResponsePage(getNextPage)
}
func NewVirtualNetworksByPCClient(referer string, regionID string, subscriptionID string) VirtualNetworksByPCClient {
	return original.NewVirtualNetworksByPCClient(referer, regionID, subscriptionID)
}
func NewVirtualNetworksByPCClientWithBaseURI(baseURI string, referer string, regionID string, subscriptionID string) VirtualNetworksByPCClient {
	return original.NewVirtualNetworksByPCClientWithBaseURI(baseURI, referer, regionID, subscriptionID)
}
func NewWithBaseURI(baseURI string, referer string, regionID string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, referer, regionID, subscriptionID)
}
func PossibleAggregationTypeValues() []AggregationType {
	return original.PossibleAggregationTypeValues()
}
func PossibleDiskIndependenceModeValues() []DiskIndependenceMode {
	return original.PossibleDiskIndependenceModeValues()
}
func PossibleGuestOSTypeValues() []GuestOSType {
	return original.PossibleGuestOSTypeValues()
}
func PossibleNICTypeValues() []NICType {
	return original.PossibleNICTypeValues()
}
func PossibleNodeStatusValues() []NodeStatus {
	return original.PossibleNodeStatusValues()
}
func PossibleOnboardingStatusValues() []OnboardingStatus {
	return original.PossibleOnboardingStatusValues()
}
func PossibleOperationOriginValues() []OperationOrigin {
	return original.PossibleOperationOriginValues()
}
func PossiblePrivateCloudResourceTypeValues() []PrivateCloudResourceType {
	return original.PossiblePrivateCloudResourceTypeValues()
}
func PossibleStopModeValues() []StopMode {
	return original.PossibleStopModeValues()
}
func PossibleUsageCountValues() []UsageCount {
	return original.PossibleUsageCountValues()
}
func PossibleVirtualMachineStatusValues() []VirtualMachineStatus {
	return original.PossibleVirtualMachineStatusValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
