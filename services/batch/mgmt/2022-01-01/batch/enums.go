package batch

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AccountKeyType enumerates the values for account key type.
type AccountKeyType string

const (
	// AccountKeyTypePrimary The primary account key.
	AccountKeyTypePrimary AccountKeyType = "Primary"
	// AccountKeyTypeSecondary The secondary account key.
	AccountKeyTypeSecondary AccountKeyType = "Secondary"
)

// PossibleAccountKeyTypeValues returns an array of possible values for the AccountKeyType const type.
func PossibleAccountKeyTypeValues() []AccountKeyType {
	return []AccountKeyType{AccountKeyTypePrimary, AccountKeyTypeSecondary}
}

// AllocationState enumerates the values for allocation state.
type AllocationState string

const (
	// AllocationStateResizing The pool is resizing; that is, compute nodes are being added to or removed from
	// the pool.
	AllocationStateResizing AllocationState = "Resizing"
	// AllocationStateSteady The pool is not resizing. There are no changes to the number of nodes in the pool
	// in progress. A pool enters this state when it is created and when no operations are being performed on
	// the pool to change the number of nodes.
	AllocationStateSteady AllocationState = "Steady"
	// AllocationStateStopping The pool was resizing, but the user has requested that the resize be stopped,
	// but the stop request has not yet been completed.
	AllocationStateStopping AllocationState = "Stopping"
)

// PossibleAllocationStateValues returns an array of possible values for the AllocationState const type.
func PossibleAllocationStateValues() []AllocationState {
	return []AllocationState{AllocationStateResizing, AllocationStateSteady, AllocationStateStopping}
}

// AuthenticationMode enumerates the values for authentication mode.
type AuthenticationMode string

const (
	// AuthenticationModeAAD The authentication mode using Azure Active Directory.
	AuthenticationModeAAD AuthenticationMode = "AAD"
	// AuthenticationModeSharedKey The authentication mode using shared keys.
	AuthenticationModeSharedKey AuthenticationMode = "SharedKey"
	// AuthenticationModeTaskAuthenticationToken The authentication mode using task authentication tokens.
	AuthenticationModeTaskAuthenticationToken AuthenticationMode = "TaskAuthenticationToken"
)

// PossibleAuthenticationModeValues returns an array of possible values for the AuthenticationMode const type.
func PossibleAuthenticationModeValues() []AuthenticationMode {
	return []AuthenticationMode{AuthenticationModeAAD, AuthenticationModeSharedKey, AuthenticationModeTaskAuthenticationToken}
}

// AutoStorageAuthenticationMode enumerates the values for auto storage authentication mode.
type AutoStorageAuthenticationMode string

const (
	// AutoStorageAuthenticationModeBatchAccountManagedIdentity The Batch service will authenticate requests to
	// auto-storage using the managed identity assigned to the Batch account.
	AutoStorageAuthenticationModeBatchAccountManagedIdentity AutoStorageAuthenticationMode = "BatchAccountManagedIdentity"
	// AutoStorageAuthenticationModeStorageKeys The Batch service will authenticate requests to auto-storage
	// using storage account keys.
	AutoStorageAuthenticationModeStorageKeys AutoStorageAuthenticationMode = "StorageKeys"
)

// PossibleAutoStorageAuthenticationModeValues returns an array of possible values for the AutoStorageAuthenticationMode const type.
func PossibleAutoStorageAuthenticationModeValues() []AutoStorageAuthenticationMode {
	return []AutoStorageAuthenticationMode{AutoStorageAuthenticationModeBatchAccountManagedIdentity, AutoStorageAuthenticationModeStorageKeys}
}

// AutoUserScope enumerates the values for auto user scope.
type AutoUserScope string

const (
	// AutoUserScopePool Specifies that the task runs as the common auto user account which is created on every
	// node in a pool.
	AutoUserScopePool AutoUserScope = "Pool"
	// AutoUserScopeTask Specifies that the service should create a new user for the task.
	AutoUserScopeTask AutoUserScope = "Task"
)

// PossibleAutoUserScopeValues returns an array of possible values for the AutoUserScope const type.
func PossibleAutoUserScopeValues() []AutoUserScope {
	return []AutoUserScope{AutoUserScopePool, AutoUserScopeTask}
}

// CachingType enumerates the values for caching type.
type CachingType string

const (
	// CachingTypeNone The caching mode for the disk is not enabled.
	CachingTypeNone CachingType = "None"
	// CachingTypeReadOnly The caching mode for the disk is read only.
	CachingTypeReadOnly CachingType = "ReadOnly"
	// CachingTypeReadWrite The caching mode for the disk is read and write.
	CachingTypeReadWrite CachingType = "ReadWrite"
)

// PossibleCachingTypeValues returns an array of possible values for the CachingType const type.
func PossibleCachingTypeValues() []CachingType {
	return []CachingType{CachingTypeNone, CachingTypeReadOnly, CachingTypeReadWrite}
}

// CertificateFormat enumerates the values for certificate format.
type CertificateFormat string

const (
	// CertificateFormatCer The certificate is a base64-encoded X.509 certificate.
	CertificateFormatCer CertificateFormat = "Cer"
	// CertificateFormatPfx The certificate is a PFX (PKCS#12) formatted certificate or certificate chain.
	CertificateFormatPfx CertificateFormat = "Pfx"
)

// PossibleCertificateFormatValues returns an array of possible values for the CertificateFormat const type.
func PossibleCertificateFormatValues() []CertificateFormat {
	return []CertificateFormat{CertificateFormatCer, CertificateFormatPfx}
}

// CertificateProvisioningState enumerates the values for certificate provisioning state.
type CertificateProvisioningState string

const (
	// CertificateProvisioningStateDeleting The user has requested that the certificate be deleted, but the
	// delete operation has not yet completed. You may not reference the certificate when creating or updating
	// pools.
	CertificateProvisioningStateDeleting CertificateProvisioningState = "Deleting"
	// CertificateProvisioningStateFailed The user requested that the certificate be deleted, but there are
	// pools that still have references to the certificate, or it is still installed on one or more compute
	// nodes. (The latter can occur if the certificate has been removed from the pool, but the node has not yet
	// restarted. Nodes refresh their certificates only when they restart.) You may use the cancel certificate
	// delete operation to cancel the delete, or the delete certificate operation to retry the delete.
	CertificateProvisioningStateFailed CertificateProvisioningState = "Failed"
	// CertificateProvisioningStateSucceeded The certificate is available for use in pools.
	CertificateProvisioningStateSucceeded CertificateProvisioningState = "Succeeded"
)

// PossibleCertificateProvisioningStateValues returns an array of possible values for the CertificateProvisioningState const type.
func PossibleCertificateProvisioningStateValues() []CertificateProvisioningState {
	return []CertificateProvisioningState{CertificateProvisioningStateDeleting, CertificateProvisioningStateFailed, CertificateProvisioningStateSucceeded}
}

// CertificateStoreLocation enumerates the values for certificate store location.
type CertificateStoreLocation string

const (
	// CertificateStoreLocationCurrentUser Certificates should be installed to the CurrentUser certificate
	// store.
	CertificateStoreLocationCurrentUser CertificateStoreLocation = "CurrentUser"
	// CertificateStoreLocationLocalMachine Certificates should be installed to the LocalMachine certificate
	// store.
	CertificateStoreLocationLocalMachine CertificateStoreLocation = "LocalMachine"
)

// PossibleCertificateStoreLocationValues returns an array of possible values for the CertificateStoreLocation const type.
func PossibleCertificateStoreLocationValues() []CertificateStoreLocation {
	return []CertificateStoreLocation{CertificateStoreLocationCurrentUser, CertificateStoreLocationLocalMachine}
}

// CertificateVisibility enumerates the values for certificate visibility.
type CertificateVisibility string

const (
	// CertificateVisibilityRemoteUser The certificate should be visible to the user accounts under which users
	// remotely access the node.
	CertificateVisibilityRemoteUser CertificateVisibility = "RemoteUser"
	// CertificateVisibilityStartTask The certificate should be visible to the user account under which the
	// start task is run. Note that if AutoUser Scope is Pool for both the StartTask and a Task, this
	// certificate will be visible to the Task as well.
	CertificateVisibilityStartTask CertificateVisibility = "StartTask"
	// CertificateVisibilityTask The certificate should be visible to the user accounts under which job tasks
	// are run.
	CertificateVisibilityTask CertificateVisibility = "Task"
)

// PossibleCertificateVisibilityValues returns an array of possible values for the CertificateVisibility const type.
func PossibleCertificateVisibilityValues() []CertificateVisibility {
	return []CertificateVisibility{CertificateVisibilityRemoteUser, CertificateVisibilityStartTask, CertificateVisibilityTask}
}

// ComputeNodeDeallocationOption enumerates the values for compute node deallocation option.
type ComputeNodeDeallocationOption string

const (
	// ComputeNodeDeallocationOptionRequeue Terminate running task processes and requeue the tasks. The tasks
	// will run again when a node is available. Remove nodes as soon as tasks have been terminated.
	ComputeNodeDeallocationOptionRequeue ComputeNodeDeallocationOption = "Requeue"
	// ComputeNodeDeallocationOptionRetainedData Allow currently running tasks to complete, then wait for all
	// task data retention periods to expire. Schedule no new tasks while waiting. Remove nodes when all task
	// retention periods have expired.
	ComputeNodeDeallocationOptionRetainedData ComputeNodeDeallocationOption = "RetainedData"
	// ComputeNodeDeallocationOptionTaskCompletion Allow currently running tasks to complete. Schedule no new
	// tasks while waiting. Remove nodes when all tasks have completed.
	ComputeNodeDeallocationOptionTaskCompletion ComputeNodeDeallocationOption = "TaskCompletion"
	// ComputeNodeDeallocationOptionTerminate Terminate running tasks. The tasks will be completed with
	// failureInfo indicating that they were terminated, and will not run again. Remove nodes as soon as tasks
	// have been terminated.
	ComputeNodeDeallocationOptionTerminate ComputeNodeDeallocationOption = "Terminate"
)

// PossibleComputeNodeDeallocationOptionValues returns an array of possible values for the ComputeNodeDeallocationOption const type.
func PossibleComputeNodeDeallocationOptionValues() []ComputeNodeDeallocationOption {
	return []ComputeNodeDeallocationOption{ComputeNodeDeallocationOptionRequeue, ComputeNodeDeallocationOptionRetainedData, ComputeNodeDeallocationOptionTaskCompletion, ComputeNodeDeallocationOptionTerminate}
}

// ComputeNodeFillType enumerates the values for compute node fill type.
type ComputeNodeFillType string

const (
	// ComputeNodeFillTypePack As many tasks as possible (taskSlotsPerNode) should be assigned to each node in
	// the pool before any tasks are assigned to the next node in the pool.
	ComputeNodeFillTypePack ComputeNodeFillType = "Pack"
	// ComputeNodeFillTypeSpread Tasks should be assigned evenly across all nodes in the pool.
	ComputeNodeFillTypeSpread ComputeNodeFillType = "Spread"
)

// PossibleComputeNodeFillTypeValues returns an array of possible values for the ComputeNodeFillType const type.
func PossibleComputeNodeFillTypeValues() []ComputeNodeFillType {
	return []ComputeNodeFillType{ComputeNodeFillTypePack, ComputeNodeFillTypeSpread}
}

// ContainerWorkingDirectory enumerates the values for container working directory.
type ContainerWorkingDirectory string

const (
	// ContainerWorkingDirectoryContainerImageDefault Using container image defined working directory. Beware
	// that this directory will not contain the resource files downloaded by Batch.
	ContainerWorkingDirectoryContainerImageDefault ContainerWorkingDirectory = "ContainerImageDefault"
	// ContainerWorkingDirectoryTaskWorkingDirectory Use the standard Batch service task working directory,
	// which will contain the Task resource files populated by Batch.
	ContainerWorkingDirectoryTaskWorkingDirectory ContainerWorkingDirectory = "TaskWorkingDirectory"
)

// PossibleContainerWorkingDirectoryValues returns an array of possible values for the ContainerWorkingDirectory const type.
func PossibleContainerWorkingDirectoryValues() []ContainerWorkingDirectory {
	return []ContainerWorkingDirectory{ContainerWorkingDirectoryContainerImageDefault, ContainerWorkingDirectoryTaskWorkingDirectory}
}

// DiffDiskPlacement enumerates the values for diff disk placement.
type DiffDiskPlacement string

const (
	// DiffDiskPlacementCacheDisk The Ephemeral OS Disk is stored on the VM cache.
	DiffDiskPlacementCacheDisk DiffDiskPlacement = "CacheDisk"
)

// PossibleDiffDiskPlacementValues returns an array of possible values for the DiffDiskPlacement const type.
func PossibleDiffDiskPlacementValues() []DiffDiskPlacement {
	return []DiffDiskPlacement{DiffDiskPlacementCacheDisk}
}

// DiskEncryptionTarget enumerates the values for disk encryption target.
type DiskEncryptionTarget string

const (
	// DiskEncryptionTargetOsDisk The OS Disk on the compute node is encrypted.
	DiskEncryptionTargetOsDisk DiskEncryptionTarget = "OsDisk"
	// DiskEncryptionTargetTemporaryDisk The temporary disk on the compute node is encrypted. On Linux this
	// encryption applies to other partitions (such as those on mounted data disks) when encryption occurs at
	// boot time.
	DiskEncryptionTargetTemporaryDisk DiskEncryptionTarget = "TemporaryDisk"
)

// PossibleDiskEncryptionTargetValues returns an array of possible values for the DiskEncryptionTarget const type.
func PossibleDiskEncryptionTargetValues() []DiskEncryptionTarget {
	return []DiskEncryptionTarget{DiskEncryptionTargetOsDisk, DiskEncryptionTargetTemporaryDisk}
}

// DynamicVNetAssignmentScope enumerates the values for dynamic v net assignment scope.
type DynamicVNetAssignmentScope string

const (
	// DynamicVNetAssignmentScopeJob Dynamic VNet assignment is done per-job. Don't use this option unless your
	// batch account has been approved to use this feature.
	DynamicVNetAssignmentScopeJob DynamicVNetAssignmentScope = "job"
	// DynamicVNetAssignmentScopeNone No dynamic VNet assignment is enabled.
	DynamicVNetAssignmentScopeNone DynamicVNetAssignmentScope = "none"
)

// PossibleDynamicVNetAssignmentScopeValues returns an array of possible values for the DynamicVNetAssignmentScope const type.
func PossibleDynamicVNetAssignmentScopeValues() []DynamicVNetAssignmentScope {
	return []DynamicVNetAssignmentScope{DynamicVNetAssignmentScopeJob, DynamicVNetAssignmentScopeNone}
}

// ElevationLevel enumerates the values for elevation level.
type ElevationLevel string

const (
	// ElevationLevelAdmin The user is a user with elevated access and operates with full Administrator
	// permissions.
	ElevationLevelAdmin ElevationLevel = "Admin"
	// ElevationLevelNonAdmin The user is a standard user without elevated access.
	ElevationLevelNonAdmin ElevationLevel = "NonAdmin"
)

// PossibleElevationLevelValues returns an array of possible values for the ElevationLevel const type.
func PossibleElevationLevelValues() []ElevationLevel {
	return []ElevationLevel{ElevationLevelAdmin, ElevationLevelNonAdmin}
}

// InboundEndpointProtocol enumerates the values for inbound endpoint protocol.
type InboundEndpointProtocol string

const (
	// InboundEndpointProtocolTCP Use TCP for the endpoint.
	InboundEndpointProtocolTCP InboundEndpointProtocol = "TCP"
	// InboundEndpointProtocolUDP Use UDP for the endpoint.
	InboundEndpointProtocolUDP InboundEndpointProtocol = "UDP"
)

// PossibleInboundEndpointProtocolValues returns an array of possible values for the InboundEndpointProtocol const type.
func PossibleInboundEndpointProtocolValues() []InboundEndpointProtocol {
	return []InboundEndpointProtocol{InboundEndpointProtocolTCP, InboundEndpointProtocolUDP}
}

// InterNodeCommunicationState enumerates the values for inter node communication state.
type InterNodeCommunicationState string

const (
	// InterNodeCommunicationStateDisabled Disable network communication between virtual machines.
	InterNodeCommunicationStateDisabled InterNodeCommunicationState = "Disabled"
	// InterNodeCommunicationStateEnabled Enable network communication between virtual machines.
	InterNodeCommunicationStateEnabled InterNodeCommunicationState = "Enabled"
)

// PossibleInterNodeCommunicationStateValues returns an array of possible values for the InterNodeCommunicationState const type.
func PossibleInterNodeCommunicationStateValues() []InterNodeCommunicationState {
	return []InterNodeCommunicationState{InterNodeCommunicationStateDisabled, InterNodeCommunicationStateEnabled}
}

// IPAddressProvisioningType enumerates the values for ip address provisioning type.
type IPAddressProvisioningType string

const (
	// IPAddressProvisioningTypeBatchManaged A public IP will be created and managed by Batch. There may be
	// multiple public IPs depending on the size of the Pool.
	IPAddressProvisioningTypeBatchManaged IPAddressProvisioningType = "BatchManaged"
	// IPAddressProvisioningTypeNoPublicIPAddresses No public IP Address will be created for the Compute Nodes
	// in the Pool.
	IPAddressProvisioningTypeNoPublicIPAddresses IPAddressProvisioningType = "NoPublicIPAddresses"
	// IPAddressProvisioningTypeUserManaged Public IPs are provided by the user and will be used to provision
	// the Compute Nodes.
	IPAddressProvisioningTypeUserManaged IPAddressProvisioningType = "UserManaged"
)

// PossibleIPAddressProvisioningTypeValues returns an array of possible values for the IPAddressProvisioningType const type.
func PossibleIPAddressProvisioningTypeValues() []IPAddressProvisioningType {
	return []IPAddressProvisioningType{IPAddressProvisioningTypeBatchManaged, IPAddressProvisioningTypeNoPublicIPAddresses, IPAddressProvisioningTypeUserManaged}
}

// KeySource enumerates the values for key source.
type KeySource string

const (
	// KeySourceMicrosoftBatch Batch creates and manages the encryption keys used to protect the account data.
	KeySourceMicrosoftBatch KeySource = "Microsoft.Batch"
	// KeySourceMicrosoftKeyVault The encryption keys used to protect the account data are stored in an
	// external key vault. If this is set then the Batch Account identity must be set to `SystemAssigned` and a
	// valid Key Identifier must also be supplied under the keyVaultProperties.
	KeySourceMicrosoftKeyVault KeySource = "Microsoft.KeyVault"
)

// PossibleKeySourceValues returns an array of possible values for the KeySource const type.
func PossibleKeySourceValues() []KeySource {
	return []KeySource{KeySourceMicrosoftBatch, KeySourceMicrosoftKeyVault}
}

// LoginMode enumerates the values for login mode.
type LoginMode string

const (
	// LoginModeBatch The LOGON32_LOGON_BATCH Win32 login mode. The batch login mode is recommended for long
	// running parallel processes.
	LoginModeBatch LoginMode = "Batch"
	// LoginModeInteractive The LOGON32_LOGON_INTERACTIVE Win32 login mode. Some applications require having
	// permissions associated with the interactive login mode. If this is the case for an application used in
	// your task, then this option is recommended.
	LoginModeInteractive LoginMode = "Interactive"
)

// PossibleLoginModeValues returns an array of possible values for the LoginMode const type.
func PossibleLoginModeValues() []LoginMode {
	return []LoginMode{LoginModeBatch, LoginModeInteractive}
}

// NameAvailabilityReason enumerates the values for name availability reason.
type NameAvailabilityReason string

const (
	// NameAvailabilityReasonAlreadyExists The requested name is already in use.
	NameAvailabilityReasonAlreadyExists NameAvailabilityReason = "AlreadyExists"
	// NameAvailabilityReasonInvalid The requested name is invalid.
	NameAvailabilityReasonInvalid NameAvailabilityReason = "Invalid"
)

// PossibleNameAvailabilityReasonValues returns an array of possible values for the NameAvailabilityReason const type.
func PossibleNameAvailabilityReasonValues() []NameAvailabilityReason {
	return []NameAvailabilityReason{NameAvailabilityReasonAlreadyExists, NameAvailabilityReasonInvalid}
}

// NetworkSecurityGroupRuleAccess enumerates the values for network security group rule access.
type NetworkSecurityGroupRuleAccess string

const (
	// NetworkSecurityGroupRuleAccessAllow Allow access.
	NetworkSecurityGroupRuleAccessAllow NetworkSecurityGroupRuleAccess = "Allow"
	// NetworkSecurityGroupRuleAccessDeny Deny access.
	NetworkSecurityGroupRuleAccessDeny NetworkSecurityGroupRuleAccess = "Deny"
)

// PossibleNetworkSecurityGroupRuleAccessValues returns an array of possible values for the NetworkSecurityGroupRuleAccess const type.
func PossibleNetworkSecurityGroupRuleAccessValues() []NetworkSecurityGroupRuleAccess {
	return []NetworkSecurityGroupRuleAccess{NetworkSecurityGroupRuleAccessAllow, NetworkSecurityGroupRuleAccessDeny}
}

// NodePlacementPolicyType enumerates the values for node placement policy type.
type NodePlacementPolicyType string

const (
	// NodePlacementPolicyTypeRegional All nodes in the pool will be allocated in the same region.
	NodePlacementPolicyTypeRegional NodePlacementPolicyType = "Regional"
	// NodePlacementPolicyTypeZonal Nodes in the pool will be spread across different zones with best effort
	// balancing.
	NodePlacementPolicyTypeZonal NodePlacementPolicyType = "Zonal"
)

// PossibleNodePlacementPolicyTypeValues returns an array of possible values for the NodePlacementPolicyType const type.
func PossibleNodePlacementPolicyTypeValues() []NodePlacementPolicyType {
	return []NodePlacementPolicyType{NodePlacementPolicyTypeRegional, NodePlacementPolicyTypeZonal}
}

// PackageState enumerates the values for package state.
type PackageState string

const (
	// PackageStateActive The application package is ready for use.
	PackageStateActive PackageState = "Active"
	// PackageStatePending The application package has been created but has not yet been activated.
	PackageStatePending PackageState = "Pending"
)

// PossiblePackageStateValues returns an array of possible values for the PackageState const type.
func PossiblePackageStateValues() []PackageState {
	return []PackageState{PackageStateActive, PackageStatePending}
}

// PoolAllocationMode enumerates the values for pool allocation mode.
type PoolAllocationMode string

const (
	// PoolAllocationModeBatchService Pools will be allocated in subscriptions owned by the Batch service.
	PoolAllocationModeBatchService PoolAllocationMode = "BatchService"
	// PoolAllocationModeUserSubscription Pools will be allocated in a subscription owned by the user.
	PoolAllocationModeUserSubscription PoolAllocationMode = "UserSubscription"
)

// PossiblePoolAllocationModeValues returns an array of possible values for the PoolAllocationMode const type.
func PossiblePoolAllocationModeValues() []PoolAllocationMode {
	return []PoolAllocationMode{PoolAllocationModeBatchService, PoolAllocationModeUserSubscription}
}

// PoolIdentityType enumerates the values for pool identity type.
type PoolIdentityType string

const (
	// PoolIdentityTypeNone Batch pool has no identity associated with it. Setting `None` in update pool will
	// remove existing identities.
	PoolIdentityTypeNone PoolIdentityType = "None"
	// PoolIdentityTypeUserAssigned Batch pool has user assigned identities with it.
	PoolIdentityTypeUserAssigned PoolIdentityType = "UserAssigned"
)

// PossiblePoolIdentityTypeValues returns an array of possible values for the PoolIdentityType const type.
func PossiblePoolIdentityTypeValues() []PoolIdentityType {
	return []PoolIdentityType{PoolIdentityTypeNone, PoolIdentityTypeUserAssigned}
}

// PoolProvisioningState enumerates the values for pool provisioning state.
type PoolProvisioningState string

const (
	// PoolProvisioningStateDeleting The user has requested that the pool be deleted, but the delete operation
	// has not yet completed.
	PoolProvisioningStateDeleting PoolProvisioningState = "Deleting"
	// PoolProvisioningStateSucceeded The pool is available to run tasks subject to the availability of compute
	// nodes.
	PoolProvisioningStateSucceeded PoolProvisioningState = "Succeeded"
)

// PossiblePoolProvisioningStateValues returns an array of possible values for the PoolProvisioningState const type.
func PossiblePoolProvisioningStateValues() []PoolProvisioningState {
	return []PoolProvisioningState{PoolProvisioningStateDeleting, PoolProvisioningStateSucceeded}
}

// PrivateEndpointConnectionProvisioningState enumerates the values for private endpoint connection
// provisioning state.
type PrivateEndpointConnectionProvisioningState string

const (
	// PrivateEndpointConnectionProvisioningStateFailed The user requested that the connection be updated and
	// it failed. You may retry the update operation.
	PrivateEndpointConnectionProvisioningStateFailed PrivateEndpointConnectionProvisioningState = "Failed"
	// PrivateEndpointConnectionProvisioningStateSucceeded The connection status is final and is ready for use
	// if Status is Approved.
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = "Succeeded"
	// PrivateEndpointConnectionProvisioningStateUpdating The user has requested that the connection status be
	// updated, but the update operation has not yet completed. You may not reference the connection when
	// connecting the Batch account.
	PrivateEndpointConnectionProvisioningStateUpdating PrivateEndpointConnectionProvisioningState = "Updating"
)

// PossiblePrivateEndpointConnectionProvisioningStateValues returns an array of possible values for the PrivateEndpointConnectionProvisioningState const type.
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{PrivateEndpointConnectionProvisioningStateFailed, PrivateEndpointConnectionProvisioningStateSucceeded, PrivateEndpointConnectionProvisioningStateUpdating}
}

// PrivateLinkServiceConnectionStatus enumerates the values for private link service connection status.
type PrivateLinkServiceConnectionStatus string

const (
	// PrivateLinkServiceConnectionStatusApproved The private endpoint connection is approved and can be used
	// to access Batch account
	PrivateLinkServiceConnectionStatusApproved PrivateLinkServiceConnectionStatus = "Approved"
	// PrivateLinkServiceConnectionStatusDisconnected The private endpoint connection is disconnected and
	// cannot be used to access Batch account
	PrivateLinkServiceConnectionStatusDisconnected PrivateLinkServiceConnectionStatus = "Disconnected"
	// PrivateLinkServiceConnectionStatusPending The private endpoint connection is pending and cannot be used
	// to access Batch account
	PrivateLinkServiceConnectionStatusPending PrivateLinkServiceConnectionStatus = "Pending"
	// PrivateLinkServiceConnectionStatusRejected The private endpoint connection is rejected and cannot be
	// used to access Batch account
	PrivateLinkServiceConnectionStatusRejected PrivateLinkServiceConnectionStatus = "Rejected"
)

// PossiblePrivateLinkServiceConnectionStatusValues returns an array of possible values for the PrivateLinkServiceConnectionStatus const type.
func PossiblePrivateLinkServiceConnectionStatusValues() []PrivateLinkServiceConnectionStatus {
	return []PrivateLinkServiceConnectionStatus{PrivateLinkServiceConnectionStatusApproved, PrivateLinkServiceConnectionStatusDisconnected, PrivateLinkServiceConnectionStatusPending, PrivateLinkServiceConnectionStatusRejected}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// ProvisioningStateCancelled The last operation for the account is cancelled.
	ProvisioningStateCancelled ProvisioningState = "Cancelled"
	// ProvisioningStateCreating The account is being created.
	ProvisioningStateCreating ProvisioningState = "Creating"
	// ProvisioningStateDeleting The account is being deleted.
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed The last operation for the account is failed.
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateInvalid The account is in an invalid state.
	ProvisioningStateInvalid ProvisioningState = "Invalid"
	// ProvisioningStateSucceeded The account has been created and is ready for use.
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{ProvisioningStateCancelled, ProvisioningStateCreating, ProvisioningStateDeleting, ProvisioningStateFailed, ProvisioningStateInvalid, ProvisioningStateSucceeded}
}

// PublicNetworkAccessType enumerates the values for public network access type.
type PublicNetworkAccessType string

const (
	// PublicNetworkAccessTypeDisabled Disables public connectivity and enables private connectivity to Azure
	// Batch Service through private endpoint resource.
	PublicNetworkAccessTypeDisabled PublicNetworkAccessType = "Disabled"
	// PublicNetworkAccessTypeEnabled Enables connectivity to Azure Batch through public DNS.
	PublicNetworkAccessTypeEnabled PublicNetworkAccessType = "Enabled"
)

// PossiblePublicNetworkAccessTypeValues returns an array of possible values for the PublicNetworkAccessType const type.
func PossiblePublicNetworkAccessTypeValues() []PublicNetworkAccessType {
	return []PublicNetworkAccessType{PublicNetworkAccessTypeDisabled, PublicNetworkAccessTypeEnabled}
}

// ResourceIdentityType enumerates the values for resource identity type.
type ResourceIdentityType string

const (
	// ResourceIdentityTypeNone Batch account has no identity associated with it. Setting `None` in update
	// account will remove existing identities.
	ResourceIdentityTypeNone ResourceIdentityType = "None"
	// ResourceIdentityTypeSystemAssigned Batch account has a system assigned identity with it.
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = "SystemAssigned"
	// ResourceIdentityTypeUserAssigned Batch account has user assigned identities with it.
	ResourceIdentityTypeUserAssigned ResourceIdentityType = "UserAssigned"
)

// PossibleResourceIdentityTypeValues returns an array of possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{ResourceIdentityTypeNone, ResourceIdentityTypeSystemAssigned, ResourceIdentityTypeUserAssigned}
}

// StorageAccountType enumerates the values for storage account type.
type StorageAccountType string

const (
	// StorageAccountTypePremiumLRS The data disk should use premium locally redundant storage.
	StorageAccountTypePremiumLRS StorageAccountType = "Premium_LRS"
	// StorageAccountTypeStandardLRS The data disk should use standard locally redundant storage.
	StorageAccountTypeStandardLRS StorageAccountType = "Standard_LRS"
)

// PossibleStorageAccountTypeValues returns an array of possible values for the StorageAccountType const type.
func PossibleStorageAccountTypeValues() []StorageAccountType {
	return []StorageAccountType{StorageAccountTypePremiumLRS, StorageAccountTypeStandardLRS}
}
