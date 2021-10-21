//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatalakestore

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// LocationsClient contains the methods for the Locations group.
// Don't use this type directly, use NewLocationsClient() instead.
type LocationsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewLocationsClient creates a new instance of LocationsClient with the specified values.
func NewLocationsClient(con *arm.Connection, subscriptionID string) *LocationsClient {
	return &LocationsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// GetCapability - Gets subscription-level properties and limits for Data Lake Store specified by resource location.
// If the operation fails it returns a generic error.
func (client *LocationsClient) GetCapability(ctx context.Context, location string, options *LocationsGetCapabilityOptions) (LocationsGetCapabilityResponse, error) {
	req, err := client.getCapabilityCreateRequest(ctx, location, options)
	if err != nil {
		return LocationsGetCapabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LocationsGetCapabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotFound) {
		return LocationsGetCapabilityResponse{}, client.getCapabilityHandleError(resp)
	}
	return client.getCapabilityHandleResponse(resp)
}

// getCapabilityCreateRequest creates the GetCapability request.
func (client *LocationsClient) getCapabilityCreateRequest(ctx context.Context, location string, options *LocationsGetCapabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DataLakeStore/locations/{location}/capability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getCapabilityHandleResponse handles the GetCapability response.
func (client *LocationsClient) getCapabilityHandleResponse(resp *http.Response) (LocationsGetCapabilityResponse, error) {
	result := LocationsGetCapabilityResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapabilityInformation); err != nil {
		return LocationsGetCapabilityResponse{}, err
	}
	return result, nil
}

// getCapabilityHandleError handles the GetCapability error response.
func (client *LocationsClient) getCapabilityHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// GetUsage - Gets the current usage count and the limit for the resources of the location under the subscription.
// If the operation fails it returns a generic error.
func (client *LocationsClient) GetUsage(ctx context.Context, location string, options *LocationsGetUsageOptions) (LocationsGetUsageResponse, error) {
	req, err := client.getUsageCreateRequest(ctx, location, options)
	if err != nil {
		return LocationsGetUsageResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LocationsGetUsageResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LocationsGetUsageResponse{}, client.getUsageHandleError(resp)
	}
	return client.getUsageHandleResponse(resp)
}

// getUsageCreateRequest creates the GetUsage request.
func (client *LocationsClient) getUsageCreateRequest(ctx context.Context, location string, options *LocationsGetUsageOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DataLakeStore/locations/{location}/usages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getUsageHandleResponse handles the GetUsage response.
func (client *LocationsClient) getUsageHandleResponse(resp *http.Response) (LocationsGetUsageResponse, error) {
	result := LocationsGetUsageResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.UsageListResult); err != nil {
		return LocationsGetUsageResponse{}, err
	}
	return result, nil
}

// getUsageHandleError handles the GetUsage error response.
func (client *LocationsClient) getUsageHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
