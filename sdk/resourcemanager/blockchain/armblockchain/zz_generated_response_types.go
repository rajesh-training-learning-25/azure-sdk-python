//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armblockchain

// LocationsClientCheckNameAvailabilityResponse contains the response from method LocationsClient.CheckNameAvailability.
type LocationsClientCheckNameAvailabilityResponse struct {
	NameAvailability
}

// LocationsClientListConsortiumsResponse contains the response from method LocationsClient.ListConsortiums.
type LocationsClientListConsortiumsResponse struct {
	ConsortiumCollection
}

// MemberOperationResultsClientGetResponse contains the response from method MemberOperationResultsClient.Get.
type MemberOperationResultsClientGetResponse struct {
	OperationResult
}

// MembersClientCreateResponse contains the response from method MembersClient.Create.
type MembersClientCreateResponse struct {
	Member
}

// MembersClientDeleteResponse contains the response from method MembersClient.Delete.
type MembersClientDeleteResponse struct {
	// placeholder for future response values
}

// MembersClientGetResponse contains the response from method MembersClient.Get.
type MembersClientGetResponse struct {
	Member
}

// MembersClientListAPIKeysResponse contains the response from method MembersClient.ListAPIKeys.
type MembersClientListAPIKeysResponse struct {
	APIKeyCollection
}

// MembersClientListAllResponse contains the response from method MembersClient.ListAll.
type MembersClientListAllResponse struct {
	MemberCollection
}

// MembersClientListConsortiumMembersResponse contains the response from method MembersClient.ListConsortiumMembers.
type MembersClientListConsortiumMembersResponse struct {
	ConsortiumMemberCollection
}

// MembersClientListRegenerateAPIKeysResponse contains the response from method MembersClient.ListRegenerateAPIKeys.
type MembersClientListRegenerateAPIKeysResponse struct {
	APIKeyCollection
}

// MembersClientListResponse contains the response from method MembersClient.List.
type MembersClientListResponse struct {
	MemberCollection
}

// MembersClientUpdateResponse contains the response from method MembersClient.Update.
type MembersClientUpdateResponse struct {
	Member
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	ResourceProviderOperationCollection
}

// SKUsClientListResponse contains the response from method SKUsClient.List.
type SKUsClientListResponse struct {
	ResourceTypeSKUCollection
}

// TransactionNodesClientCreateResponse contains the response from method TransactionNodesClient.Create.
type TransactionNodesClientCreateResponse struct {
	TransactionNode
}

// TransactionNodesClientDeleteResponse contains the response from method TransactionNodesClient.Delete.
type TransactionNodesClientDeleteResponse struct {
	// placeholder for future response values
}

// TransactionNodesClientGetResponse contains the response from method TransactionNodesClient.Get.
type TransactionNodesClientGetResponse struct {
	TransactionNode
}

// TransactionNodesClientListAPIKeysResponse contains the response from method TransactionNodesClient.ListAPIKeys.
type TransactionNodesClientListAPIKeysResponse struct {
	APIKeyCollection
}

// TransactionNodesClientListRegenerateAPIKeysResponse contains the response from method TransactionNodesClient.ListRegenerateAPIKeys.
type TransactionNodesClientListRegenerateAPIKeysResponse struct {
	APIKeyCollection
}

// TransactionNodesClientListResponse contains the response from method TransactionNodesClient.List.
type TransactionNodesClientListResponse struct {
	TransactionNodeCollection
}

// TransactionNodesClientUpdateResponse contains the response from method TransactionNodesClient.Update.
type TransactionNodesClientUpdateResponse struct {
	TransactionNode
}