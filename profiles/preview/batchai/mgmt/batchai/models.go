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

package batchai

import original "github.com/Azure/azure-sdk-for-go/services/batchai/mgmt/2018-03-01/batchai"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type OperationsClient = original.OperationsClient
type UsageClient = original.UsageClient
type JobsClient = original.JobsClient
type FileServersClient = original.FileServersClient
type AllocationState = original.AllocationState

const (
	Resizing AllocationState = original.Resizing
	Steady   AllocationState = original.Steady
)

type CachingType = original.CachingType

const (
	None      CachingType = original.None
	Readonly  CachingType = original.Readonly
	Readwrite CachingType = original.Readwrite
)

type DeallocationOption = original.DeallocationOption

const (
	Requeue              DeallocationOption = original.Requeue
	Terminate            DeallocationOption = original.Terminate
	Unknown              DeallocationOption = original.Unknown
	Waitforjobcompletion DeallocationOption = original.Waitforjobcompletion
)

type ExecutionState = original.ExecutionState

const (
	Failed      ExecutionState = original.Failed
	Queued      ExecutionState = original.Queued
	Running     ExecutionState = original.Running
	Succeeded   ExecutionState = original.Succeeded
	Terminating ExecutionState = original.Terminating
)

type FileServerProvisioningState = original.FileServerProvisioningState

const (
	FileServerProvisioningStateCreating  FileServerProvisioningState = original.FileServerProvisioningStateCreating
	FileServerProvisioningStateDeleting  FileServerProvisioningState = original.FileServerProvisioningStateDeleting
	FileServerProvisioningStateFailed    FileServerProvisioningState = original.FileServerProvisioningStateFailed
	FileServerProvisioningStateSucceeded FileServerProvisioningState = original.FileServerProvisioningStateSucceeded
	FileServerProvisioningStateUpdating  FileServerProvisioningState = original.FileServerProvisioningStateUpdating
)

type FileServerType = original.FileServerType

const (
	Glusterfs FileServerType = original.Glusterfs
	Nfs       FileServerType = original.Nfs
)

type OutputType = original.OutputType

const (
	Custom  OutputType = original.Custom
	Logs    OutputType = original.Logs
	Model   OutputType = original.Model
	Summary OutputType = original.Summary
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateCreating  ProvisioningState = original.ProvisioningStateCreating
	ProvisioningStateDeleting  ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateFailed    ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateSucceeded ProvisioningState = original.ProvisioningStateSucceeded
)

type StorageAccountType = original.StorageAccountType

const (
	PremiumLRS  StorageAccountType = original.PremiumLRS
	StandardLRS StorageAccountType = original.StandardLRS
)

type ToolType = original.ToolType

const (
	ToolTypeCaffe      ToolType = original.ToolTypeCaffe
	ToolTypeCaffe2     ToolType = original.ToolTypeCaffe2
	ToolTypeChainer    ToolType = original.ToolTypeChainer
	ToolTypeCntk       ToolType = original.ToolTypeCntk
	ToolTypeCustom     ToolType = original.ToolTypeCustom
	ToolTypeTensorflow ToolType = original.ToolTypeTensorflow
)

type VMPriority = original.VMPriority

const (
	Dedicated   VMPriority = original.Dedicated
	Lowpriority VMPriority = original.Lowpriority
)

type AppInsightsReference = original.AppInsightsReference
type AutoScaleSettings = original.AutoScaleSettings
type AzureBlobFileSystemReference = original.AzureBlobFileSystemReference
type AzureFileShareReference = original.AzureFileShareReference
type AzureStorageCredentialsInfo = original.AzureStorageCredentialsInfo
type Caffe2Settings = original.Caffe2Settings
type CaffeSettings = original.CaffeSettings
type ChainerSettings = original.ChainerSettings
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type Cluster = original.Cluster
type ClusterBaseProperties = original.ClusterBaseProperties
type ClusterCreateParameters = original.ClusterCreateParameters
type ClusterListResult = original.ClusterListResult
type ClusterListResultIterator = original.ClusterListResultIterator
type ClusterListResultPage = original.ClusterListResultPage
type ClusterProperties = original.ClusterProperties
type ClustersCreateFuture = original.ClustersCreateFuture
type ClustersDeleteFuture = original.ClustersDeleteFuture
type ClusterUpdateParameters = original.ClusterUpdateParameters
type ClusterUpdateProperties = original.ClusterUpdateProperties
type CNTKsettings = original.CNTKsettings
type ContainerSettings = original.ContainerSettings
type CustomToolkitSettings = original.CustomToolkitSettings
type DataDisks = original.DataDisks
type EnvironmentVariable = original.EnvironmentVariable
type EnvironmentVariableWithSecretValue = original.EnvironmentVariableWithSecretValue
type Error = original.Error
type File = original.File
type FileListResult = original.FileListResult
type FileListResultIterator = original.FileListResultIterator
type FileListResultPage = original.FileListResultPage
type FileProperties = original.FileProperties
type FileServer = original.FileServer
type FileServerBaseProperties = original.FileServerBaseProperties
type FileServerCreateParameters = original.FileServerCreateParameters
type FileServerListResult = original.FileServerListResult
type FileServerListResultIterator = original.FileServerListResultIterator
type FileServerListResultPage = original.FileServerListResultPage
type FileServerProperties = original.FileServerProperties
type FileServerReference = original.FileServerReference
type FileServersCreateFuture = original.FileServersCreateFuture
type FileServersDeleteFuture = original.FileServersDeleteFuture
type ImageReference = original.ImageReference
type ImageSourceRegistry = original.ImageSourceRegistry
type InputDirectory = original.InputDirectory
type Job = original.Job
type JobBaseProperties = original.JobBaseProperties
type JobBasePropertiesConstraints = original.JobBasePropertiesConstraints
type JobCreateParameters = original.JobCreateParameters
type JobListResult = original.JobListResult
type JobListResultIterator = original.JobListResultIterator
type JobListResultPage = original.JobListResultPage
type JobPreparation = original.JobPreparation
type JobProperties = original.JobProperties
type JobPropertiesConstraints = original.JobPropertiesConstraints
type JobPropertiesExecutionInfo = original.JobPropertiesExecutionInfo
type JobsCreateFuture = original.JobsCreateFuture
type JobsDeleteFuture = original.JobsDeleteFuture
type JobsTerminateFuture = original.JobsTerminateFuture
type KeyVaultKeyReference = original.KeyVaultKeyReference
type KeyVaultSecretReference = original.KeyVaultSecretReference
type ListUsagesResult = original.ListUsagesResult
type ListUsagesResultIterator = original.ListUsagesResultIterator
type ListUsagesResultPage = original.ListUsagesResultPage
type LocalDataVolume = original.LocalDataVolume
type ManualScaleSettings = original.ManualScaleSettings
type MountSettings = original.MountSettings
type MountVolumes = original.MountVolumes
type NameValuePair = original.NameValuePair
type NodeSetup = original.NodeSetup
type NodeStateCounts = original.NodeStateCounts
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OutputDirectory = original.OutputDirectory
type PerformanceCountersSettings = original.PerformanceCountersSettings
type PrivateRegistryCredentials = original.PrivateRegistryCredentials
type PyTorchSettings = original.PyTorchSettings
type RemoteLoginInformation = original.RemoteLoginInformation
type RemoteLoginInformationListResult = original.RemoteLoginInformationListResult
type RemoteLoginInformationListResultIterator = original.RemoteLoginInformationListResultIterator
type RemoteLoginInformationListResultPage = original.RemoteLoginInformationListResultPage
type Resource = original.Resource
type ResourceID = original.ResourceID
type ScaleSettings = original.ScaleSettings
type SetupTask = original.SetupTask
type SSHConfiguration = original.SSHConfiguration
type TensorFlowSettings = original.TensorFlowSettings
type UnmanagedFileSystemReference = original.UnmanagedFileSystemReference
type Usage = original.Usage
type UsageName = original.UsageName
type UserAccountSettings = original.UserAccountSettings
type VirtualMachineConfiguration = original.VirtualMachineConfiguration
type ClustersClient = original.ClustersClient

func NewFileServersClient(subscriptionID string) FileServersClient {
	return original.NewFileServersClient(subscriptionID)
}
func NewFileServersClientWithBaseURI(baseURI string, subscriptionID string) FileServersClient {
	return original.NewFileServersClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleAllocationStateValues() []AllocationState {
	return original.PossibleAllocationStateValues()
}
func PossibleCachingTypeValues() []CachingType {
	return original.PossibleCachingTypeValues()
}
func PossibleDeallocationOptionValues() []DeallocationOption {
	return original.PossibleDeallocationOptionValues()
}
func PossibleExecutionStateValues() []ExecutionState {
	return original.PossibleExecutionStateValues()
}
func PossibleFileServerProvisioningStateValues() []FileServerProvisioningState {
	return original.PossibleFileServerProvisioningStateValues()
}
func PossibleFileServerTypeValues() []FileServerType {
	return original.PossibleFileServerTypeValues()
}
func PossibleOutputTypeValues() []OutputType {
	return original.PossibleOutputTypeValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleStorageAccountTypeValues() []StorageAccountType {
	return original.PossibleStorageAccountTypeValues()
}
func PossibleToolTypeValues() []ToolType {
	return original.PossibleToolTypeValues()
}
func PossibleVMPriorityValues() []VMPriority {
	return original.PossibleVMPriorityValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
func NewClustersClient(subscriptionID string) ClustersClient {
	return original.NewClustersClient(subscriptionID)
}
func NewClustersClientWithBaseURI(baseURI string, subscriptionID string) ClustersClient {
	return original.NewClustersClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewUsageClient(subscriptionID string) UsageClient {
	return original.NewUsageClient(subscriptionID)
}
func NewUsageClientWithBaseURI(baseURI string, subscriptionID string) UsageClient {
	return original.NewUsageClientWithBaseURI(baseURI, subscriptionID)
}
func NewJobsClient(subscriptionID string) JobsClient {
	return original.NewJobsClient(subscriptionID)
}
func NewJobsClientWithBaseURI(baseURI string, subscriptionID string) JobsClient {
	return original.NewJobsClientWithBaseURI(baseURI, subscriptionID)
}
