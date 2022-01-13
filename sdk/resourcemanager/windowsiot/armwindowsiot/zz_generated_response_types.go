//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armwindowsiot

import "net/http"

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsClientListResult contains the result from method OperationsClient.List.
type OperationsClientListResult struct {
	OperationListResult
}

// ServicesClientCheckDeviceServiceNameAvailabilityResponse contains the response from method ServicesClient.CheckDeviceServiceNameAvailability.
type ServicesClientCheckDeviceServiceNameAvailabilityResponse struct {
	ServicesClientCheckDeviceServiceNameAvailabilityResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServicesClientCheckDeviceServiceNameAvailabilityResult contains the result from method ServicesClient.CheckDeviceServiceNameAvailability.
type ServicesClientCheckDeviceServiceNameAvailabilityResult struct {
	DeviceServiceNameAvailabilityInfo
}

// ServicesClientCreateOrUpdateResponse contains the response from method ServicesClient.CreateOrUpdate.
type ServicesClientCreateOrUpdateResponse struct {
	ServicesClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServicesClientCreateOrUpdateResult contains the result from method ServicesClient.CreateOrUpdate.
type ServicesClientCreateOrUpdateResult struct {
	DeviceService
}

// ServicesClientDeleteResponse contains the response from method ServicesClient.Delete.
type ServicesClientDeleteResponse struct {
	ServicesClientDeleteResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServicesClientDeleteResult contains the result from method ServicesClient.Delete.
type ServicesClientDeleteResult struct {
	DeviceService
}

// ServicesClientGetResponse contains the response from method ServicesClient.Get.
type ServicesClientGetResponse struct {
	ServicesClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServicesClientGetResult contains the result from method ServicesClient.Get.
type ServicesClientGetResult struct {
	DeviceService
}

// ServicesClientListByResourceGroupResponse contains the response from method ServicesClient.ListByResourceGroup.
type ServicesClientListByResourceGroupResponse struct {
	ServicesClientListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServicesClientListByResourceGroupResult contains the result from method ServicesClient.ListByResourceGroup.
type ServicesClientListByResourceGroupResult struct {
	DeviceServiceDescriptionListResult
}

// ServicesClientListResponse contains the response from method ServicesClient.List.
type ServicesClientListResponse struct {
	ServicesClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServicesClientListResult contains the result from method ServicesClient.List.
type ServicesClientListResult struct {
	DeviceServiceDescriptionListResult
}

// ServicesClientUpdateResponse contains the response from method ServicesClient.Update.
type ServicesClientUpdateResponse struct {
	ServicesClientUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServicesClientUpdateResult contains the result from method ServicesClient.Update.
type ServicesClientUpdateResult struct {
	DeviceService
}
