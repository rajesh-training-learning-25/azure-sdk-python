// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package netapp

import original "github.com/Azure/azure-sdk-for-go/services/preview/netapp/mgmt/2017-08-15/netapp"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ServiceLevel = original.ServiceLevel

const (
	Premium  ServiceLevel = original.Premium
	Standard ServiceLevel = original.Standard
	Ultra    ServiceLevel = original.Ultra
)

type Account = original.Account
type AccountList = original.AccountList
type AccountPatch = original.AccountPatch
type AccountProperties = original.AccountProperties
type AccountsClient = original.AccountsClient
type AccountsCreateOrUpdateFuture = original.AccountsCreateOrUpdateFuture
type AccountsDeleteFuture = original.AccountsDeleteFuture
type ActiveDirectory = original.ActiveDirectory
type BaseClient = original.BaseClient
type CapacityPool = original.CapacityPool
type CapacityPoolList = original.CapacityPoolList
type CapacityPoolPatch = original.CapacityPoolPatch
type Dimension = original.Dimension
type Error = original.Error
type ExportPolicyRule = original.ExportPolicyRule
type MetricSpecification = original.MetricSpecification
type MountTarget = original.MountTarget
type MountTargetList = original.MountTargetList
type MountTargetProperties = original.MountTargetProperties
type MountTargetsClient = original.MountTargetsClient
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationProperties = original.OperationProperties
type OperationsClient = original.OperationsClient
type PoolPatchProperties = original.PoolPatchProperties
type PoolProperties = original.PoolProperties
type PoolsClient = original.PoolsClient
type PoolsCreateOrUpdateFuture = original.PoolsCreateOrUpdateFuture
type PoolsDeleteFuture = original.PoolsDeleteFuture
type ServiceSpecification = original.ServiceSpecification
type Snapshot = original.Snapshot
type SnapshotPatch = original.SnapshotPatch
type SnapshotProperties = original.SnapshotProperties
type SnapshotsClient = original.SnapshotsClient
type SnapshotsCreateFuture = original.SnapshotsCreateFuture
type SnapshotsDeleteFuture = original.SnapshotsDeleteFuture
type SnapshotsList = original.SnapshotsList
type Volume = original.Volume
type VolumeList = original.VolumeList
type VolumePatch = original.VolumePatch
type VolumePatchProperties = original.VolumePatchProperties
type VolumePatchPropertiesExportPolicy = original.VolumePatchPropertiesExportPolicy
type VolumeProperties = original.VolumeProperties
type VolumePropertiesExportPolicy = original.VolumePropertiesExportPolicy
type VolumesClient = original.VolumesClient
type VolumesCreateOrUpdateFuture = original.VolumesCreateOrUpdateFuture
type VolumesDeleteFuture = original.VolumesDeleteFuture

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAccountsClient(subscriptionID string) AccountsClient {
	return original.NewAccountsClient(subscriptionID)
}
func NewAccountsClientWithBaseURI(baseURI string, subscriptionID string) AccountsClient {
	return original.NewAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewMountTargetsClient(subscriptionID string) MountTargetsClient {
	return original.NewMountTargetsClient(subscriptionID)
}
func NewMountTargetsClientWithBaseURI(baseURI string, subscriptionID string) MountTargetsClient {
	return original.NewMountTargetsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPoolsClient(subscriptionID string) PoolsClient {
	return original.NewPoolsClient(subscriptionID)
}
func NewPoolsClientWithBaseURI(baseURI string, subscriptionID string) PoolsClient {
	return original.NewPoolsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSnapshotsClient(subscriptionID string) SnapshotsClient {
	return original.NewSnapshotsClient(subscriptionID)
}
func NewSnapshotsClientWithBaseURI(baseURI string, subscriptionID string) SnapshotsClient {
	return original.NewSnapshotsClientWithBaseURI(baseURI, subscriptionID)
}
func NewVolumesClient(subscriptionID string) VolumesClient {
	return original.NewVolumesClient(subscriptionID)
}
func NewVolumesClientWithBaseURI(baseURI string, subscriptionID string) VolumesClient {
	return original.NewVolumesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleServiceLevelValues() []ServiceLevel {
	return original.PossibleServiceLevelValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
