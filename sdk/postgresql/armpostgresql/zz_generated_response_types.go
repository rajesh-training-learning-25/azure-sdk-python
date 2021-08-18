// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresql

import (
	"context"
	"net/http"
	"time"
)

// CheckNameAvailabilityExecuteResponse contains the response from method CheckNameAvailability.Execute.
type CheckNameAvailabilityExecuteResponse struct {
	CheckNameAvailabilityExecuteResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CheckNameAvailabilityExecuteResult contains the result from method CheckNameAvailability.Execute.
type CheckNameAvailabilityExecuteResult struct {
	NameAvailability
}

// ConfigurationsCreateOrUpdatePollerResponse contains the response from method Configurations.CreateOrUpdate.
type ConfigurationsCreateOrUpdatePollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (ConfigurationsCreateOrUpdateResponse, error)

	// Poller contains an initialized poller.
	Poller ConfigurationsCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConfigurationsCreateOrUpdateResponse contains the response from method Configurations.CreateOrUpdate.
type ConfigurationsCreateOrUpdateResponse struct {
	ConfigurationsCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConfigurationsCreateOrUpdateResult contains the result from method Configurations.CreateOrUpdate.
type ConfigurationsCreateOrUpdateResult struct {
	Configuration
}

// ConfigurationsGetResponse contains the response from method Configurations.Get.
type ConfigurationsGetResponse struct {
	ConfigurationsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConfigurationsGetResult contains the result from method Configurations.Get.
type ConfigurationsGetResult struct {
	Configuration
}

// ConfigurationsListByServerResponse contains the response from method Configurations.ListByServer.
type ConfigurationsListByServerResponse struct {
	ConfigurationsListByServerResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConfigurationsListByServerResult contains the result from method Configurations.ListByServer.
type ConfigurationsListByServerResult struct {
	ConfigurationListResult
}

// DatabasesCreateOrUpdatePollerResponse contains the response from method Databases.CreateOrUpdate.
type DatabasesCreateOrUpdatePollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (DatabasesCreateOrUpdateResponse, error)

	// Poller contains an initialized poller.
	Poller DatabasesCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DatabasesCreateOrUpdateResponse contains the response from method Databases.CreateOrUpdate.
type DatabasesCreateOrUpdateResponse struct {
	DatabasesCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DatabasesCreateOrUpdateResult contains the result from method Databases.CreateOrUpdate.
type DatabasesCreateOrUpdateResult struct {
	Database
}

// DatabasesDeletePollerResponse contains the response from method Databases.Delete.
type DatabasesDeletePollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (DatabasesDeleteResponse, error)

	// Poller contains an initialized poller.
	Poller DatabasesDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DatabasesDeleteResponse contains the response from method Databases.Delete.
type DatabasesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DatabasesGetResponse contains the response from method Databases.Get.
type DatabasesGetResponse struct {
	DatabasesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DatabasesGetResult contains the result from method Databases.Get.
type DatabasesGetResult struct {
	Database
}

// DatabasesListByServerResponse contains the response from method Databases.ListByServer.
type DatabasesListByServerResponse struct {
	DatabasesListByServerResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DatabasesListByServerResult contains the result from method Databases.ListByServer.
type DatabasesListByServerResult struct {
	DatabaseListResult
}

// FirewallRulesCreateOrUpdatePollerResponse contains the response from method FirewallRules.CreateOrUpdate.
type FirewallRulesCreateOrUpdatePollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (FirewallRulesCreateOrUpdateResponse, error)

	// Poller contains an initialized poller.
	Poller FirewallRulesCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FirewallRulesCreateOrUpdateResponse contains the response from method FirewallRules.CreateOrUpdate.
type FirewallRulesCreateOrUpdateResponse struct {
	FirewallRulesCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FirewallRulesCreateOrUpdateResult contains the result from method FirewallRules.CreateOrUpdate.
type FirewallRulesCreateOrUpdateResult struct {
	FirewallRule
}

// FirewallRulesDeletePollerResponse contains the response from method FirewallRules.Delete.
type FirewallRulesDeletePollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (FirewallRulesDeleteResponse, error)

	// Poller contains an initialized poller.
	Poller FirewallRulesDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FirewallRulesDeleteResponse contains the response from method FirewallRules.Delete.
type FirewallRulesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FirewallRulesGetResponse contains the response from method FirewallRules.Get.
type FirewallRulesGetResponse struct {
	FirewallRulesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FirewallRulesGetResult contains the result from method FirewallRules.Get.
type FirewallRulesGetResult struct {
	FirewallRule
}

// FirewallRulesListByServerResponse contains the response from method FirewallRules.ListByServer.
type FirewallRulesListByServerResponse struct {
	FirewallRulesListByServerResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FirewallRulesListByServerResult contains the result from method FirewallRules.ListByServer.
type FirewallRulesListByServerResult struct {
	FirewallRuleListResult
}

// LocationBasedPerformanceTierListResponse contains the response from method LocationBasedPerformanceTier.List.
type LocationBasedPerformanceTierListResponse struct {
	LocationBasedPerformanceTierListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// LocationBasedPerformanceTierListResult contains the result from method LocationBasedPerformanceTier.List.
type LocationBasedPerformanceTierListResult struct {
	PerformanceTierListResult
}

// LogFilesListByServerResponse contains the response from method LogFiles.ListByServer.
type LogFilesListByServerResponse struct {
	LogFilesListByServerResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// LogFilesListByServerResult contains the result from method LogFiles.ListByServer.
type LogFilesListByServerResult struct {
	LogFileListResult
}

// OperationsListResponse contains the response from method Operations.List.
type OperationsListResponse struct {
	OperationsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsListResult contains the result from method Operations.List.
type OperationsListResult struct {
	OperationListResult
}

// PrivateEndpointConnectionsCreateOrUpdatePollerResponse contains the response from method PrivateEndpointConnections.CreateOrUpdate.
type PrivateEndpointConnectionsCreateOrUpdatePollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (PrivateEndpointConnectionsCreateOrUpdateResponse, error)

	// Poller contains an initialized poller.
	Poller PrivateEndpointConnectionsCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsCreateOrUpdateResponse contains the response from method PrivateEndpointConnections.CreateOrUpdate.
type PrivateEndpointConnectionsCreateOrUpdateResponse struct {
	PrivateEndpointConnectionsCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsCreateOrUpdateResult contains the result from method PrivateEndpointConnections.CreateOrUpdate.
type PrivateEndpointConnectionsCreateOrUpdateResult struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsDeletePollerResponse contains the response from method PrivateEndpointConnections.Delete.
type PrivateEndpointConnectionsDeletePollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (PrivateEndpointConnectionsDeleteResponse, error)

	// Poller contains an initialized poller.
	Poller PrivateEndpointConnectionsDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsDeleteResponse contains the response from method PrivateEndpointConnections.Delete.
type PrivateEndpointConnectionsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsGetResponse contains the response from method PrivateEndpointConnections.Get.
type PrivateEndpointConnectionsGetResponse struct {
	PrivateEndpointConnectionsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsGetResult contains the result from method PrivateEndpointConnections.Get.
type PrivateEndpointConnectionsGetResult struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsListByServerResponse contains the response from method PrivateEndpointConnections.ListByServer.
type PrivateEndpointConnectionsListByServerResponse struct {
	PrivateEndpointConnectionsListByServerResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsListByServerResult contains the result from method PrivateEndpointConnections.ListByServer.
type PrivateEndpointConnectionsListByServerResult struct {
	PrivateEndpointConnectionListResult
}

// PrivateEndpointConnectionsUpdateTagsPollerResponse contains the response from method PrivateEndpointConnections.UpdateTags.
type PrivateEndpointConnectionsUpdateTagsPollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (PrivateEndpointConnectionsUpdateTagsResponse, error)

	// Poller contains an initialized poller.
	Poller PrivateEndpointConnectionsUpdateTagsPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsUpdateTagsResponse contains the response from method PrivateEndpointConnections.UpdateTags.
type PrivateEndpointConnectionsUpdateTagsResponse struct {
	PrivateEndpointConnectionsUpdateTagsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsUpdateTagsResult contains the result from method PrivateEndpointConnections.UpdateTags.
type PrivateEndpointConnectionsUpdateTagsResult struct {
	PrivateEndpointConnection
}

// PrivateLinkResourcesGetResponse contains the response from method PrivateLinkResources.Get.
type PrivateLinkResourcesGetResponse struct {
	PrivateLinkResourcesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateLinkResourcesGetResult contains the result from method PrivateLinkResources.Get.
type PrivateLinkResourcesGetResult struct {
	PrivateLinkResource
}

// PrivateLinkResourcesListByServerResponse contains the response from method PrivateLinkResources.ListByServer.
type PrivateLinkResourcesListByServerResponse struct {
	PrivateLinkResourcesListByServerResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateLinkResourcesListByServerResult contains the result from method PrivateLinkResources.ListByServer.
type PrivateLinkResourcesListByServerResult struct {
	PrivateLinkResourceListResult
}

// RecoverableServersGetResponse contains the response from method RecoverableServers.Get.
type RecoverableServersGetResponse struct {
	RecoverableServersGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RecoverableServersGetResult contains the result from method RecoverableServers.Get.
type RecoverableServersGetResult struct {
	RecoverableServerResource
}

// ReplicasListByServerResponse contains the response from method Replicas.ListByServer.
type ReplicasListByServerResponse struct {
	ReplicasListByServerResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ReplicasListByServerResult contains the result from method Replicas.ListByServer.
type ReplicasListByServerResult struct {
	ServerListResult
}

// ServerAdministratorsCreateOrUpdatePollerResponse contains the response from method ServerAdministrators.CreateOrUpdate.
type ServerAdministratorsCreateOrUpdatePollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (ServerAdministratorsCreateOrUpdateResponse, error)

	// Poller contains an initialized poller.
	Poller ServerAdministratorsCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServerAdministratorsCreateOrUpdateResponse contains the response from method ServerAdministrators.CreateOrUpdate.
type ServerAdministratorsCreateOrUpdateResponse struct {
	ServerAdministratorsCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServerAdministratorsCreateOrUpdateResult contains the result from method ServerAdministrators.CreateOrUpdate.
type ServerAdministratorsCreateOrUpdateResult struct {
	ServerAdministratorResource
}

// ServerAdministratorsDeletePollerResponse contains the response from method ServerAdministrators.Delete.
type ServerAdministratorsDeletePollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (ServerAdministratorsDeleteResponse, error)

	// Poller contains an initialized poller.
	Poller ServerAdministratorsDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServerAdministratorsDeleteResponse contains the response from method ServerAdministrators.Delete.
type ServerAdministratorsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServerAdministratorsGetResponse contains the response from method ServerAdministrators.Get.
type ServerAdministratorsGetResponse struct {
	ServerAdministratorsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServerAdministratorsGetResult contains the result from method ServerAdministrators.Get.
type ServerAdministratorsGetResult struct {
	ServerAdministratorResource
}

// ServerAdministratorsListResponse contains the response from method ServerAdministrators.List.
type ServerAdministratorsListResponse struct {
	ServerAdministratorsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServerAdministratorsListResult contains the result from method ServerAdministrators.List.
type ServerAdministratorsListResult struct {
	ServerAdministratorResourceListResult
}

// ServerBasedPerformanceTierListResponse contains the response from method ServerBasedPerformanceTier.List.
type ServerBasedPerformanceTierListResponse struct {
	ServerBasedPerformanceTierListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServerBasedPerformanceTierListResult contains the result from method ServerBasedPerformanceTier.List.
type ServerBasedPerformanceTierListResult struct {
	PerformanceTierListResult
}

// ServerKeysCreateOrUpdatePollerResponse contains the response from method ServerKeys.CreateOrUpdate.
type ServerKeysCreateOrUpdatePollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (ServerKeysCreateOrUpdateResponse, error)

	// Poller contains an initialized poller.
	Poller ServerKeysCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServerKeysCreateOrUpdateResponse contains the response from method ServerKeys.CreateOrUpdate.
type ServerKeysCreateOrUpdateResponse struct {
	ServerKeysCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServerKeysCreateOrUpdateResult contains the result from method ServerKeys.CreateOrUpdate.
type ServerKeysCreateOrUpdateResult struct {
	ServerKey
}

// ServerKeysDeletePollerResponse contains the response from method ServerKeys.Delete.
type ServerKeysDeletePollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (ServerKeysDeleteResponse, error)

	// Poller contains an initialized poller.
	Poller ServerKeysDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServerKeysDeleteResponse contains the response from method ServerKeys.Delete.
type ServerKeysDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServerKeysGetResponse contains the response from method ServerKeys.Get.
type ServerKeysGetResponse struct {
	ServerKeysGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServerKeysGetResult contains the result from method ServerKeys.Get.
type ServerKeysGetResult struct {
	ServerKey
}

// ServerKeysListResponse contains the response from method ServerKeys.List.
type ServerKeysListResponse struct {
	ServerKeysListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServerKeysListResult contains the result from method ServerKeys.List.
type ServerKeysListResult struct {
	ServerKeyListResult
}

// ServerParametersListUpdateConfigurationsPollerResponse contains the response from method ServerParameters.ListUpdateConfigurations.
type ServerParametersListUpdateConfigurationsPollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (ServerParametersListUpdateConfigurationsResponse, error)

	// Poller contains an initialized poller.
	Poller ServerParametersListUpdateConfigurationsPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServerParametersListUpdateConfigurationsResponse contains the response from method ServerParameters.ListUpdateConfigurations.
type ServerParametersListUpdateConfigurationsResponse struct {
	ServerParametersListUpdateConfigurationsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServerParametersListUpdateConfigurationsResult contains the result from method ServerParameters.ListUpdateConfigurations.
type ServerParametersListUpdateConfigurationsResult struct {
	ConfigurationListResult
}

// ServerSecurityAlertPoliciesCreateOrUpdatePollerResponse contains the response from method ServerSecurityAlertPolicies.CreateOrUpdate.
type ServerSecurityAlertPoliciesCreateOrUpdatePollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (ServerSecurityAlertPoliciesCreateOrUpdateResponse, error)

	// Poller contains an initialized poller.
	Poller ServerSecurityAlertPoliciesCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServerSecurityAlertPoliciesCreateOrUpdateResponse contains the response from method ServerSecurityAlertPolicies.CreateOrUpdate.
type ServerSecurityAlertPoliciesCreateOrUpdateResponse struct {
	ServerSecurityAlertPoliciesCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServerSecurityAlertPoliciesCreateOrUpdateResult contains the result from method ServerSecurityAlertPolicies.CreateOrUpdate.
type ServerSecurityAlertPoliciesCreateOrUpdateResult struct {
	ServerSecurityAlertPolicy
}

// ServerSecurityAlertPoliciesGetResponse contains the response from method ServerSecurityAlertPolicies.Get.
type ServerSecurityAlertPoliciesGetResponse struct {
	ServerSecurityAlertPoliciesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServerSecurityAlertPoliciesGetResult contains the result from method ServerSecurityAlertPolicies.Get.
type ServerSecurityAlertPoliciesGetResult struct {
	ServerSecurityAlertPolicy
}

// ServerSecurityAlertPoliciesListByServerResponse contains the response from method ServerSecurityAlertPolicies.ListByServer.
type ServerSecurityAlertPoliciesListByServerResponse struct {
	ServerSecurityAlertPoliciesListByServerResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServerSecurityAlertPoliciesListByServerResult contains the result from method ServerSecurityAlertPolicies.ListByServer.
type ServerSecurityAlertPoliciesListByServerResult struct {
	ServerSecurityAlertPolicyListResult
}

// ServersCreatePollerResponse contains the response from method Servers.Create.
type ServersCreatePollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (ServersCreateResponse, error)

	// Poller contains an initialized poller.
	Poller ServersCreatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersCreateResponse contains the response from method Servers.Create.
type ServersCreateResponse struct {
	ServersCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersCreateResult contains the result from method Servers.Create.
type ServersCreateResult struct {
	Server
}

// ServersDeletePollerResponse contains the response from method Servers.Delete.
type ServersDeletePollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (ServersDeleteResponse, error)

	// Poller contains an initialized poller.
	Poller ServersDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersDeleteResponse contains the response from method Servers.Delete.
type ServersDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersGetResponse contains the response from method Servers.Get.
type ServersGetResponse struct {
	ServersGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersGetResult contains the result from method Servers.Get.
type ServersGetResult struct {
	Server
}

// ServersListByResourceGroupResponse contains the response from method Servers.ListByResourceGroup.
type ServersListByResourceGroupResponse struct {
	ServersListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersListByResourceGroupResult contains the result from method Servers.ListByResourceGroup.
type ServersListByResourceGroupResult struct {
	ServerListResult
}

// ServersListResponse contains the response from method Servers.List.
type ServersListResponse struct {
	ServersListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersListResult contains the result from method Servers.List.
type ServersListResult struct {
	ServerListResult
}

// ServersRestartPollerResponse contains the response from method Servers.Restart.
type ServersRestartPollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (ServersRestartResponse, error)

	// Poller contains an initialized poller.
	Poller ServersRestartPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersRestartResponse contains the response from method Servers.Restart.
type ServersRestartResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersUpdatePollerResponse contains the response from method Servers.Update.
type ServersUpdatePollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (ServersUpdateResponse, error)

	// Poller contains an initialized poller.
	Poller ServersUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersUpdateResponse contains the response from method Servers.Update.
type ServersUpdateResponse struct {
	ServersUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersUpdateResult contains the result from method Servers.Update.
type ServersUpdateResult struct {
	Server
}

// VirtualNetworkRulesCreateOrUpdatePollerResponse contains the response from method VirtualNetworkRules.CreateOrUpdate.
type VirtualNetworkRulesCreateOrUpdatePollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (VirtualNetworkRulesCreateOrUpdateResponse, error)

	// Poller contains an initialized poller.
	Poller VirtualNetworkRulesCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VirtualNetworkRulesCreateOrUpdateResponse contains the response from method VirtualNetworkRules.CreateOrUpdate.
type VirtualNetworkRulesCreateOrUpdateResponse struct {
	VirtualNetworkRulesCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VirtualNetworkRulesCreateOrUpdateResult contains the result from method VirtualNetworkRules.CreateOrUpdate.
type VirtualNetworkRulesCreateOrUpdateResult struct {
	VirtualNetworkRule
}

// VirtualNetworkRulesDeletePollerResponse contains the response from method VirtualNetworkRules.Delete.
type VirtualNetworkRulesDeletePollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (VirtualNetworkRulesDeleteResponse, error)

	// Poller contains an initialized poller.
	Poller VirtualNetworkRulesDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VirtualNetworkRulesDeleteResponse contains the response from method VirtualNetworkRules.Delete.
type VirtualNetworkRulesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VirtualNetworkRulesGetResponse contains the response from method VirtualNetworkRules.Get.
type VirtualNetworkRulesGetResponse struct {
	VirtualNetworkRulesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VirtualNetworkRulesGetResult contains the result from method VirtualNetworkRules.Get.
type VirtualNetworkRulesGetResult struct {
	VirtualNetworkRule
}

// VirtualNetworkRulesListByServerResponse contains the response from method VirtualNetworkRules.ListByServer.
type VirtualNetworkRulesListByServerResponse struct {
	VirtualNetworkRulesListByServerResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VirtualNetworkRulesListByServerResult contains the result from method VirtualNetworkRules.ListByServer.
type VirtualNetworkRulesListByServerResult struct {
	VirtualNetworkRuleListResult
}
