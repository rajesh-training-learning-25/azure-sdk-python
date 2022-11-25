//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerservice

// AgentPoolsClientAbortLatestOperationResponse contains the response from method AgentPoolsClient.AbortLatestOperation.
type AgentPoolsClientAbortLatestOperationResponse struct {
	// placeholder for future response values
}

// AgentPoolsClientCreateOrUpdateResponse contains the response from method AgentPoolsClient.CreateOrUpdate.
type AgentPoolsClientCreateOrUpdateResponse struct {
	AgentPool
}

// AgentPoolsClientDeleteResponse contains the response from method AgentPoolsClient.Delete.
type AgentPoolsClientDeleteResponse struct {
	// placeholder for future response values
}

// AgentPoolsClientGetAvailableAgentPoolVersionsResponse contains the response from method AgentPoolsClient.GetAvailableAgentPoolVersions.
type AgentPoolsClientGetAvailableAgentPoolVersionsResponse struct {
	AgentPoolAvailableVersions
}

// AgentPoolsClientGetResponse contains the response from method AgentPoolsClient.Get.
type AgentPoolsClientGetResponse struct {
	AgentPool
}

// AgentPoolsClientGetUpgradeProfileResponse contains the response from method AgentPoolsClient.GetUpgradeProfile.
type AgentPoolsClientGetUpgradeProfileResponse struct {
	AgentPoolUpgradeProfile
}

// AgentPoolsClientListResponse contains the response from method AgentPoolsClient.List.
type AgentPoolsClientListResponse struct {
	AgentPoolListResult
}

// AgentPoolsClientUpgradeNodeImageVersionResponse contains the response from method AgentPoolsClient.UpgradeNodeImageVersion.
type AgentPoolsClientUpgradeNodeImageVersionResponse struct {
	AgentPool
}

// FleetMembersClientCreateOrUpdateResponse contains the response from method FleetMembersClient.CreateOrUpdate.
type FleetMembersClientCreateOrUpdateResponse struct {
	FleetMember
}

// FleetMembersClientDeleteResponse contains the response from method FleetMembersClient.Delete.
type FleetMembersClientDeleteResponse struct {
	// placeholder for future response values
}

// FleetMembersClientGetResponse contains the response from method FleetMembersClient.Get.
type FleetMembersClientGetResponse struct {
	FleetMember
}

// FleetMembersClientListByFleetResponse contains the response from method FleetMembersClient.ListByFleet.
type FleetMembersClientListByFleetResponse struct {
	FleetMembersListResult
}

// FleetsClientCreateOrUpdateResponse contains the response from method FleetsClient.CreateOrUpdate.
type FleetsClientCreateOrUpdateResponse struct {
	Fleet
}

// FleetsClientDeleteResponse contains the response from method FleetsClient.Delete.
type FleetsClientDeleteResponse struct {
	// placeholder for future response values
}

// FleetsClientGetResponse contains the response from method FleetsClient.Get.
type FleetsClientGetResponse struct {
	Fleet
}

// FleetsClientListByResourceGroupResponse contains the response from method FleetsClient.ListByResourceGroup.
type FleetsClientListByResourceGroupResponse struct {
	FleetListResult
}

// FleetsClientListCredentialsResponse contains the response from method FleetsClient.ListCredentials.
type FleetsClientListCredentialsResponse struct {
	FleetCredentialResults
}

// FleetsClientListResponse contains the response from method FleetsClient.List.
type FleetsClientListResponse struct {
	FleetListResult
}

// FleetsClientUpdateResponse contains the response from method FleetsClient.Update.
type FleetsClientUpdateResponse struct {
	Fleet
}

// MaintenanceConfigurationsClientCreateOrUpdateResponse contains the response from method MaintenanceConfigurationsClient.CreateOrUpdate.
type MaintenanceConfigurationsClientCreateOrUpdateResponse struct {
	MaintenanceConfiguration
}

// MaintenanceConfigurationsClientDeleteResponse contains the response from method MaintenanceConfigurationsClient.Delete.
type MaintenanceConfigurationsClientDeleteResponse struct {
	// placeholder for future response values
}

// MaintenanceConfigurationsClientGetResponse contains the response from method MaintenanceConfigurationsClient.Get.
type MaintenanceConfigurationsClientGetResponse struct {
	MaintenanceConfiguration
}

// MaintenanceConfigurationsClientListByManagedClusterResponse contains the response from method MaintenanceConfigurationsClient.ListByManagedCluster.
type MaintenanceConfigurationsClientListByManagedClusterResponse struct {
	MaintenanceConfigurationListResult
}

// ManagedClusterSnapshotsClientCreateOrUpdateResponse contains the response from method ManagedClusterSnapshotsClient.CreateOrUpdate.
type ManagedClusterSnapshotsClientCreateOrUpdateResponse struct {
	ManagedClusterSnapshot
}

// ManagedClusterSnapshotsClientDeleteResponse contains the response from method ManagedClusterSnapshotsClient.Delete.
type ManagedClusterSnapshotsClientDeleteResponse struct {
	// placeholder for future response values
}

// ManagedClusterSnapshotsClientGetResponse contains the response from method ManagedClusterSnapshotsClient.Get.
type ManagedClusterSnapshotsClientGetResponse struct {
	ManagedClusterSnapshot
}

// ManagedClusterSnapshotsClientListByResourceGroupResponse contains the response from method ManagedClusterSnapshotsClient.ListByResourceGroup.
type ManagedClusterSnapshotsClientListByResourceGroupResponse struct {
	ManagedClusterSnapshotListResult
}

// ManagedClusterSnapshotsClientListResponse contains the response from method ManagedClusterSnapshotsClient.List.
type ManagedClusterSnapshotsClientListResponse struct {
	ManagedClusterSnapshotListResult
}

// ManagedClusterSnapshotsClientUpdateTagsResponse contains the response from method ManagedClusterSnapshotsClient.UpdateTags.
type ManagedClusterSnapshotsClientUpdateTagsResponse struct {
	ManagedClusterSnapshot
}

// ManagedClustersClientAbortLatestOperationResponse contains the response from method ManagedClustersClient.AbortLatestOperation.
type ManagedClustersClientAbortLatestOperationResponse struct {
	// placeholder for future response values
}

// ManagedClustersClientCreateOrUpdateResponse contains the response from method ManagedClustersClient.CreateOrUpdate.
type ManagedClustersClientCreateOrUpdateResponse struct {
	ManagedCluster
}

// ManagedClustersClientDeleteResponse contains the response from method ManagedClustersClient.Delete.
type ManagedClustersClientDeleteResponse struct {
	// placeholder for future response values
}

// ManagedClustersClientGetAccessProfileResponse contains the response from method ManagedClustersClient.GetAccessProfile.
type ManagedClustersClientGetAccessProfileResponse struct {
	ManagedClusterAccessProfile
}

// ManagedClustersClientGetCommandResultResponse contains the response from method ManagedClustersClient.GetCommandResult.
type ManagedClustersClientGetCommandResultResponse struct {
	RunCommandResult
	// Location contains the information returned from the Location header response.
	Location *string
}

// ManagedClustersClientGetOSOptionsResponse contains the response from method ManagedClustersClient.GetOSOptions.
type ManagedClustersClientGetOSOptionsResponse struct {
	OSOptionProfile
}

// ManagedClustersClientGetResponse contains the response from method ManagedClustersClient.Get.
type ManagedClustersClientGetResponse struct {
	ManagedCluster
}

// ManagedClustersClientGetUpgradeProfileResponse contains the response from method ManagedClustersClient.GetUpgradeProfile.
type ManagedClustersClientGetUpgradeProfileResponse struct {
	ManagedClusterUpgradeProfile
}

// ManagedClustersClientListByResourceGroupResponse contains the response from method ManagedClustersClient.ListByResourceGroup.
type ManagedClustersClientListByResourceGroupResponse struct {
	ManagedClusterListResult
}

// ManagedClustersClientListClusterAdminCredentialsResponse contains the response from method ManagedClustersClient.ListClusterAdminCredentials.
type ManagedClustersClientListClusterAdminCredentialsResponse struct {
	CredentialResults
}

// ManagedClustersClientListClusterMonitoringUserCredentialsResponse contains the response from method ManagedClustersClient.ListClusterMonitoringUserCredentials.
type ManagedClustersClientListClusterMonitoringUserCredentialsResponse struct {
	CredentialResults
}

// ManagedClustersClientListClusterUserCredentialsResponse contains the response from method ManagedClustersClient.ListClusterUserCredentials.
type ManagedClustersClientListClusterUserCredentialsResponse struct {
	CredentialResults
}

// ManagedClustersClientListOutboundNetworkDependenciesEndpointsResponse contains the response from method ManagedClustersClient.ListOutboundNetworkDependenciesEndpoints.
type ManagedClustersClientListOutboundNetworkDependenciesEndpointsResponse struct {
	OutboundEnvironmentEndpointCollection
}

// ManagedClustersClientListResponse contains the response from method ManagedClustersClient.List.
type ManagedClustersClientListResponse struct {
	ManagedClusterListResult
}

// ManagedClustersClientResetAADProfileResponse contains the response from method ManagedClustersClient.ResetAADProfile.
type ManagedClustersClientResetAADProfileResponse struct {
	// placeholder for future response values
}

// ManagedClustersClientResetServicePrincipalProfileResponse contains the response from method ManagedClustersClient.ResetServicePrincipalProfile.
type ManagedClustersClientResetServicePrincipalProfileResponse struct {
	// placeholder for future response values
}

// ManagedClustersClientRotateClusterCertificatesResponse contains the response from method ManagedClustersClient.RotateClusterCertificates.
type ManagedClustersClientRotateClusterCertificatesResponse struct {
	// placeholder for future response values
}

// ManagedClustersClientRotateServiceAccountSigningKeysResponse contains the response from method ManagedClustersClient.RotateServiceAccountSigningKeys.
type ManagedClustersClientRotateServiceAccountSigningKeysResponse struct {
	// placeholder for future response values
}

// ManagedClustersClientRunCommandResponse contains the response from method ManagedClustersClient.RunCommand.
type ManagedClustersClientRunCommandResponse struct {
	RunCommandResult
}

// ManagedClustersClientStartResponse contains the response from method ManagedClustersClient.Start.
type ManagedClustersClientStartResponse struct {
	// placeholder for future response values
}

// ManagedClustersClientStopResponse contains the response from method ManagedClustersClient.Stop.
type ManagedClustersClientStopResponse struct {
	// placeholder for future response values
}

// ManagedClustersClientUpdateTagsResponse contains the response from method ManagedClustersClient.UpdateTags.
type ManagedClustersClientUpdateTagsResponse struct {
	ManagedCluster
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.Delete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListResponse contains the response from method PrivateEndpointConnectionsClient.List.
type PrivateEndpointConnectionsClientListResponse struct {
	PrivateEndpointConnectionListResult
}

// PrivateEndpointConnectionsClientUpdateResponse contains the response from method PrivateEndpointConnectionsClient.Update.
type PrivateEndpointConnectionsClientUpdateResponse struct {
	PrivateEndpointConnection
}

// PrivateLinkResourcesClientListResponse contains the response from method PrivateLinkResourcesClient.List.
type PrivateLinkResourcesClientListResponse struct {
	PrivateLinkResourcesListResult
}

// ResolvePrivateLinkServiceIDClientPOSTResponse contains the response from method ResolvePrivateLinkServiceIDClient.POST.
type ResolvePrivateLinkServiceIDClientPOSTResponse struct {
	PrivateLinkResource
}

// SnapshotsClientCreateOrUpdateResponse contains the response from method SnapshotsClient.CreateOrUpdate.
type SnapshotsClientCreateOrUpdateResponse struct {
	Snapshot
}

// SnapshotsClientDeleteResponse contains the response from method SnapshotsClient.Delete.
type SnapshotsClientDeleteResponse struct {
	// placeholder for future response values
}

// SnapshotsClientGetResponse contains the response from method SnapshotsClient.Get.
type SnapshotsClientGetResponse struct {
	Snapshot
}

// SnapshotsClientListByResourceGroupResponse contains the response from method SnapshotsClient.ListByResourceGroup.
type SnapshotsClientListByResourceGroupResponse struct {
	SnapshotListResult
}

// SnapshotsClientListResponse contains the response from method SnapshotsClient.List.
type SnapshotsClientListResponse struct {
	SnapshotListResult
}

// SnapshotsClientUpdateTagsResponse contains the response from method SnapshotsClient.UpdateTags.
type SnapshotsClientUpdateTagsResponse struct {
	Snapshot
}

// TrustedAccessRoleBindingsClientCreateOrUpdateResponse contains the response from method TrustedAccessRoleBindingsClient.CreateOrUpdate.
type TrustedAccessRoleBindingsClientCreateOrUpdateResponse struct {
	TrustedAccessRoleBinding
}

// TrustedAccessRoleBindingsClientDeleteResponse contains the response from method TrustedAccessRoleBindingsClient.Delete.
type TrustedAccessRoleBindingsClientDeleteResponse struct {
	// placeholder for future response values
}

// TrustedAccessRoleBindingsClientGetResponse contains the response from method TrustedAccessRoleBindingsClient.Get.
type TrustedAccessRoleBindingsClientGetResponse struct {
	TrustedAccessRoleBinding
}

// TrustedAccessRoleBindingsClientListResponse contains the response from method TrustedAccessRoleBindingsClient.List.
type TrustedAccessRoleBindingsClientListResponse struct {
	TrustedAccessRoleBindingListResult
}

// TrustedAccessRolesClientListResponse contains the response from method TrustedAccessRolesClient.List.
type TrustedAccessRolesClientListResponse struct {
	TrustedAccessRoleListResult
}
