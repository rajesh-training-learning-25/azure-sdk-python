//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armbotservice

// BotConnectionClientCreateResponse contains the response from method BotConnectionClient.Create.
type BotConnectionClientCreateResponse struct {
	ConnectionSetting
}

// BotConnectionClientDeleteResponse contains the response from method BotConnectionClient.Delete.
type BotConnectionClientDeleteResponse struct {
	// placeholder for future response values
}

// BotConnectionClientGetResponse contains the response from method BotConnectionClient.Get.
type BotConnectionClientGetResponse struct {
	ConnectionSetting
}

// BotConnectionClientListByBotServiceResponse contains the response from method BotConnectionClient.ListByBotService.
type BotConnectionClientListByBotServiceResponse struct {
	ConnectionSettingResponseList
}

// BotConnectionClientListServiceProvidersResponse contains the response from method BotConnectionClient.ListServiceProviders.
type BotConnectionClientListServiceProvidersResponse struct {
	ServiceProviderResponseList
}

// BotConnectionClientListWithSecretsResponse contains the response from method BotConnectionClient.ListWithSecrets.
type BotConnectionClientListWithSecretsResponse struct {
	ConnectionSetting
}

// BotConnectionClientUpdateResponse contains the response from method BotConnectionClient.Update.
type BotConnectionClientUpdateResponse struct {
	ConnectionSetting
}

// BotsClientCreateResponse contains the response from method BotsClient.Create.
type BotsClientCreateResponse struct {
	Bot
}

// BotsClientDeleteResponse contains the response from method BotsClient.Delete.
type BotsClientDeleteResponse struct {
	// placeholder for future response values
}

// BotsClientGetCheckNameAvailabilityResponse contains the response from method BotsClient.GetCheckNameAvailability.
type BotsClientGetCheckNameAvailabilityResponse struct {
	CheckNameAvailabilityResponseBody
}

// BotsClientGetResponse contains the response from method BotsClient.Get.
type BotsClientGetResponse struct {
	Bot
}

// BotsClientListByResourceGroupResponse contains the response from method BotsClient.ListByResourceGroup.
type BotsClientListByResourceGroupResponse struct {
	BotResponseList
}

// BotsClientListResponse contains the response from method BotsClient.List.
type BotsClientListResponse struct {
	BotResponseList
}

// BotsClientUpdateResponse contains the response from method BotsClient.Update.
type BotsClientUpdateResponse struct {
	Bot
}

// ChannelsClientCreateResponse contains the response from method ChannelsClient.Create.
type ChannelsClientCreateResponse struct {
	BotChannel
}

// ChannelsClientDeleteResponse contains the response from method ChannelsClient.Delete.
type ChannelsClientDeleteResponse struct {
	// placeholder for future response values
}

// ChannelsClientGetResponse contains the response from method ChannelsClient.Get.
type ChannelsClientGetResponse struct {
	BotChannel
}

// ChannelsClientListByResourceGroupResponse contains the response from method ChannelsClient.ListByResourceGroup.
type ChannelsClientListByResourceGroupResponse struct {
	ChannelResponseList
}

// ChannelsClientListWithKeysResponse contains the response from method ChannelsClient.ListWithKeys.
type ChannelsClientListWithKeysResponse struct {
	ListChannelWithKeysResponse
}

// ChannelsClientUpdateResponse contains the response from method ChannelsClient.Update.
type ChannelsClientUpdateResponse struct {
	BotChannel
}

// DirectLineClientRegenerateKeysResponse contains the response from method DirectLineClient.RegenerateKeys.
type DirectLineClientRegenerateKeysResponse struct {
	BotChannel
}

// EmailClientCreateSignInURLResponse contains the response from method EmailClient.CreateSignInURL.
type EmailClientCreateSignInURLResponse struct {
	CreateEmailSignInURLResponse
}

// HostSettingsClientGetResponse contains the response from method HostSettingsClient.Get.
type HostSettingsClientGetResponse struct {
	HostSettingsResponse
}

// OperationResultsClientGetResponse contains the response from method OperationResultsClient.Get.
type OperationResultsClientGetResponse struct {
	OperationResultsDescription
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationEntityListResult
}

// PrivateEndpointConnectionsClientCreateResponse contains the response from method PrivateEndpointConnectionsClient.Create.
type PrivateEndpointConnectionsClientCreateResponse struct {
	PrivateEndpointConnection
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

// PrivateLinkResourcesClientListByBotResourceResponse contains the response from method PrivateLinkResourcesClient.ListByBotResource.
type PrivateLinkResourcesClientListByBotResourceResponse struct {
	PrivateLinkResourceListResult
}

// QnAMakerEndpointKeysClientGetResponse contains the response from method QnAMakerEndpointKeysClient.Get.
type QnAMakerEndpointKeysClientGetResponse struct {
	QnAMakerEndpointKeysResponse
}
