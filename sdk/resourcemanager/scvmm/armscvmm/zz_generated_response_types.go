//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armscvmm

// AvailabilitySetsClientCreateOrUpdateResponse contains the response from method AvailabilitySetsClient.CreateOrUpdate.
type AvailabilitySetsClientCreateOrUpdateResponse struct {
	AvailabilitySet
}

// AvailabilitySetsClientDeleteResponse contains the response from method AvailabilitySetsClient.Delete.
type AvailabilitySetsClientDeleteResponse struct {
	// placeholder for future response values
}

// AvailabilitySetsClientGetResponse contains the response from method AvailabilitySetsClient.Get.
type AvailabilitySetsClientGetResponse struct {
	AvailabilitySet
}

// AvailabilitySetsClientListByResourceGroupResponse contains the response from method AvailabilitySetsClient.ListByResourceGroup.
type AvailabilitySetsClientListByResourceGroupResponse struct {
	AvailabilitySetListResult
}

// AvailabilitySetsClientListBySubscriptionResponse contains the response from method AvailabilitySetsClient.ListBySubscription.
type AvailabilitySetsClientListBySubscriptionResponse struct {
	AvailabilitySetListResult
}

// AvailabilitySetsClientUpdateResponse contains the response from method AvailabilitySetsClient.Update.
type AvailabilitySetsClientUpdateResponse struct {
	AvailabilitySet
}

// CloudsClientCreateOrUpdateResponse contains the response from method CloudsClient.CreateOrUpdate.
type CloudsClientCreateOrUpdateResponse struct {
	Cloud
}

// CloudsClientDeleteResponse contains the response from method CloudsClient.Delete.
type CloudsClientDeleteResponse struct {
	// placeholder for future response values
}

// CloudsClientGetResponse contains the response from method CloudsClient.Get.
type CloudsClientGetResponse struct {
	Cloud
}

// CloudsClientListByResourceGroupResponse contains the response from method CloudsClient.ListByResourceGroup.
type CloudsClientListByResourceGroupResponse struct {
	CloudListResult
}

// CloudsClientListBySubscriptionResponse contains the response from method CloudsClient.ListBySubscription.
type CloudsClientListBySubscriptionResponse struct {
	CloudListResult
}

// CloudsClientUpdateResponse contains the response from method CloudsClient.Update.
type CloudsClientUpdateResponse struct {
	Cloud
}

// InventoryItemsClientCreateResponse contains the response from method InventoryItemsClient.Create.
type InventoryItemsClientCreateResponse struct {
	InventoryItem
}

// InventoryItemsClientDeleteResponse contains the response from method InventoryItemsClient.Delete.
type InventoryItemsClientDeleteResponse struct {
	// placeholder for future response values
}

// InventoryItemsClientGetResponse contains the response from method InventoryItemsClient.Get.
type InventoryItemsClientGetResponse struct {
	InventoryItem
}

// InventoryItemsClientListByVMMServerResponse contains the response from method InventoryItemsClient.ListByVMMServer.
type InventoryItemsClientListByVMMServerResponse struct {
	InventoryItemsList
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	ResourceProviderOperationList
}

// VirtualMachineTemplatesClientCreateOrUpdateResponse contains the response from method VirtualMachineTemplatesClient.CreateOrUpdate.
type VirtualMachineTemplatesClientCreateOrUpdateResponse struct {
	VirtualMachineTemplate
}

// VirtualMachineTemplatesClientDeleteResponse contains the response from method VirtualMachineTemplatesClient.Delete.
type VirtualMachineTemplatesClientDeleteResponse struct {
	// placeholder for future response values
}

// VirtualMachineTemplatesClientGetResponse contains the response from method VirtualMachineTemplatesClient.Get.
type VirtualMachineTemplatesClientGetResponse struct {
	VirtualMachineTemplate
}

// VirtualMachineTemplatesClientListByResourceGroupResponse contains the response from method VirtualMachineTemplatesClient.ListByResourceGroup.
type VirtualMachineTemplatesClientListByResourceGroupResponse struct {
	VirtualMachineTemplateListResult
}

// VirtualMachineTemplatesClientListBySubscriptionResponse contains the response from method VirtualMachineTemplatesClient.ListBySubscription.
type VirtualMachineTemplatesClientListBySubscriptionResponse struct {
	VirtualMachineTemplateListResult
}

// VirtualMachineTemplatesClientUpdateResponse contains the response from method VirtualMachineTemplatesClient.Update.
type VirtualMachineTemplatesClientUpdateResponse struct {
	VirtualMachineTemplate
}

// VirtualMachinesClientCreateCheckpointResponse contains the response from method VirtualMachinesClient.CreateCheckpoint.
type VirtualMachinesClientCreateCheckpointResponse struct {
	// placeholder for future response values
}

// VirtualMachinesClientCreateOrUpdateResponse contains the response from method VirtualMachinesClient.CreateOrUpdate.
type VirtualMachinesClientCreateOrUpdateResponse struct {
	VirtualMachine
}

// VirtualMachinesClientDeleteCheckpointResponse contains the response from method VirtualMachinesClient.DeleteCheckpoint.
type VirtualMachinesClientDeleteCheckpointResponse struct {
	// placeholder for future response values
}

// VirtualMachinesClientDeleteResponse contains the response from method VirtualMachinesClient.Delete.
type VirtualMachinesClientDeleteResponse struct {
	// placeholder for future response values
}

// VirtualMachinesClientGetResponse contains the response from method VirtualMachinesClient.Get.
type VirtualMachinesClientGetResponse struct {
	VirtualMachine
}

// VirtualMachinesClientListByResourceGroupResponse contains the response from method VirtualMachinesClient.ListByResourceGroup.
type VirtualMachinesClientListByResourceGroupResponse struct {
	VirtualMachineListResult
}

// VirtualMachinesClientListBySubscriptionResponse contains the response from method VirtualMachinesClient.ListBySubscription.
type VirtualMachinesClientListBySubscriptionResponse struct {
	VirtualMachineListResult
}

// VirtualMachinesClientRestartResponse contains the response from method VirtualMachinesClient.Restart.
type VirtualMachinesClientRestartResponse struct {
	// placeholder for future response values
}

// VirtualMachinesClientRestoreCheckpointResponse contains the response from method VirtualMachinesClient.RestoreCheckpoint.
type VirtualMachinesClientRestoreCheckpointResponse struct {
	// placeholder for future response values
}

// VirtualMachinesClientStartResponse contains the response from method VirtualMachinesClient.Start.
type VirtualMachinesClientStartResponse struct {
	// placeholder for future response values
}

// VirtualMachinesClientStopResponse contains the response from method VirtualMachinesClient.Stop.
type VirtualMachinesClientStopResponse struct {
	// placeholder for future response values
}

// VirtualMachinesClientUpdateResponse contains the response from method VirtualMachinesClient.Update.
type VirtualMachinesClientUpdateResponse struct {
	VirtualMachine
}

// VirtualNetworksClientCreateOrUpdateResponse contains the response from method VirtualNetworksClient.CreateOrUpdate.
type VirtualNetworksClientCreateOrUpdateResponse struct {
	VirtualNetwork
}

// VirtualNetworksClientDeleteResponse contains the response from method VirtualNetworksClient.Delete.
type VirtualNetworksClientDeleteResponse struct {
	// placeholder for future response values
}

// VirtualNetworksClientGetResponse contains the response from method VirtualNetworksClient.Get.
type VirtualNetworksClientGetResponse struct {
	VirtualNetwork
}

// VirtualNetworksClientListByResourceGroupResponse contains the response from method VirtualNetworksClient.ListByResourceGroup.
type VirtualNetworksClientListByResourceGroupResponse struct {
	VirtualNetworkListResult
}

// VirtualNetworksClientListBySubscriptionResponse contains the response from method VirtualNetworksClient.ListBySubscription.
type VirtualNetworksClientListBySubscriptionResponse struct {
	VirtualNetworkListResult
}

// VirtualNetworksClientUpdateResponse contains the response from method VirtualNetworksClient.Update.
type VirtualNetworksClientUpdateResponse struct {
	VirtualNetwork
}

// VmmServersClientCreateOrUpdateResponse contains the response from method VmmServersClient.CreateOrUpdate.
type VmmServersClientCreateOrUpdateResponse struct {
	VMMServer
}

// VmmServersClientDeleteResponse contains the response from method VmmServersClient.Delete.
type VmmServersClientDeleteResponse struct {
	// placeholder for future response values
}

// VmmServersClientGetResponse contains the response from method VmmServersClient.Get.
type VmmServersClientGetResponse struct {
	VMMServer
}

// VmmServersClientListByResourceGroupResponse contains the response from method VmmServersClient.ListByResourceGroup.
type VmmServersClientListByResourceGroupResponse struct {
	VMMServerListResult
}

// VmmServersClientListBySubscriptionResponse contains the response from method VmmServersClient.ListBySubscription.
type VmmServersClientListBySubscriptionResponse struct {
	VMMServerListResult
}

// VmmServersClientUpdateResponse contains the response from method VmmServersClient.Update.
type VmmServersClientUpdateResponse struct {
	VMMServer
}