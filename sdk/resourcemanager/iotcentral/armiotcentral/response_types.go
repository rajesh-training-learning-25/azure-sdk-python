//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armiotcentral

// AppsClientCheckNameAvailabilityResponse contains the response from method AppsClient.CheckNameAvailability.
type AppsClientCheckNameAvailabilityResponse struct {
	AppAvailabilityInfo
}

// AppsClientCheckSubdomainAvailabilityResponse contains the response from method AppsClient.CheckSubdomainAvailability.
type AppsClientCheckSubdomainAvailabilityResponse struct {
	AppAvailabilityInfo
}

// AppsClientCreateOrUpdateResponse contains the response from method AppsClient.CreateOrUpdate.
type AppsClientCreateOrUpdateResponse struct {
	App
}

// AppsClientDeleteResponse contains the response from method AppsClient.Delete.
type AppsClientDeleteResponse struct {
	// placeholder for future response values
}

// AppsClientGetResponse contains the response from method AppsClient.Get.
type AppsClientGetResponse struct {
	App
}

// AppsClientListByResourceGroupResponse contains the response from method AppsClient.ListByResourceGroup.
type AppsClientListByResourceGroupResponse struct {
	AppListResult
}

// AppsClientListBySubscriptionResponse contains the response from method AppsClient.ListBySubscription.
type AppsClientListBySubscriptionResponse struct {
	AppListResult
}

// AppsClientListTemplatesResponse contains the response from method AppsClient.ListTemplates.
type AppsClientListTemplatesResponse struct {
	AppTemplatesResult
}

// AppsClientUpdateResponse contains the response from method AppsClient.Update.
type AppsClientUpdateResponse struct {
	// placeholder for future response values
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
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

// PrivateLinksClientGetResponse contains the response from method PrivateLinksClient.Get.
type PrivateLinksClientGetResponse struct {
	PrivateLinkResource
}

// PrivateLinksClientListResponse contains the response from method PrivateLinksClient.List.
type PrivateLinksClientListResponse struct {
	PrivateLinkResourceListResult
}