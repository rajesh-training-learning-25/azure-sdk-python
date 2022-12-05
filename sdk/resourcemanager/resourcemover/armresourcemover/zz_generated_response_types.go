//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcemover

// MoveCollectionsClientBulkRemoveResponse contains the response from method MoveCollectionsClient.BulkRemove.
type MoveCollectionsClientBulkRemoveResponse struct {
	OperationStatus
}

// MoveCollectionsClientCommitResponse contains the response from method MoveCollectionsClient.Commit.
type MoveCollectionsClientCommitResponse struct {
	OperationStatus
}

// MoveCollectionsClientCreateResponse contains the response from method MoveCollectionsClient.Create.
type MoveCollectionsClientCreateResponse struct {
	MoveCollection
}

// MoveCollectionsClientDeleteResponse contains the response from method MoveCollectionsClient.Delete.
type MoveCollectionsClientDeleteResponse struct {
	OperationStatus
}

// MoveCollectionsClientDiscardResponse contains the response from method MoveCollectionsClient.Discard.
type MoveCollectionsClientDiscardResponse struct {
	OperationStatus
}

// MoveCollectionsClientGetResponse contains the response from method MoveCollectionsClient.Get.
type MoveCollectionsClientGetResponse struct {
	MoveCollection
}

// MoveCollectionsClientInitiateMoveResponse contains the response from method MoveCollectionsClient.InitiateMove.
type MoveCollectionsClientInitiateMoveResponse struct {
	OperationStatus
}

// MoveCollectionsClientListMoveCollectionsByResourceGroupResponse contains the response from method MoveCollectionsClient.ListMoveCollectionsByResourceGroup.
type MoveCollectionsClientListMoveCollectionsByResourceGroupResponse struct {
	MoveCollectionResultList
}

// MoveCollectionsClientListMoveCollectionsBySubscriptionResponse contains the response from method MoveCollectionsClient.ListMoveCollectionsBySubscription.
type MoveCollectionsClientListMoveCollectionsBySubscriptionResponse struct {
	MoveCollectionResultList
}

// MoveCollectionsClientListRequiredForResponse contains the response from method MoveCollectionsClient.ListRequiredFor.
type MoveCollectionsClientListRequiredForResponse struct {
	RequiredForResourcesCollection
}

// MoveCollectionsClientPrepareResponse contains the response from method MoveCollectionsClient.Prepare.
type MoveCollectionsClientPrepareResponse struct {
	OperationStatus
}

// MoveCollectionsClientResolveDependenciesResponse contains the response from method MoveCollectionsClient.ResolveDependencies.
type MoveCollectionsClientResolveDependenciesResponse struct {
	OperationStatus
}

// MoveCollectionsClientUpdateResponse contains the response from method MoveCollectionsClient.Update.
type MoveCollectionsClientUpdateResponse struct {
	MoveCollection
}

// MoveResourcesClientCreateResponse contains the response from method MoveResourcesClient.Create.
type MoveResourcesClientCreateResponse struct {
	MoveResource
}

// MoveResourcesClientDeleteResponse contains the response from method MoveResourcesClient.Delete.
type MoveResourcesClientDeleteResponse struct {
	OperationStatus
}

// MoveResourcesClientGetResponse contains the response from method MoveResourcesClient.Get.
type MoveResourcesClientGetResponse struct {
	MoveResource
}

// MoveResourcesClientListResponse contains the response from method MoveResourcesClient.List.
type MoveResourcesClientListResponse struct {
	MoveResourceCollection
}

// OperationsDiscoveryClientGetResponse contains the response from method OperationsDiscoveryClient.Get.
type OperationsDiscoveryClientGetResponse struct {
	OperationsDiscoveryCollection
}

// UnresolvedDependenciesClientGetResponse contains the response from method UnresolvedDependenciesClient.Get.
type UnresolvedDependenciesClientGetResponse struct {
	UnresolvedDependencyCollection
}