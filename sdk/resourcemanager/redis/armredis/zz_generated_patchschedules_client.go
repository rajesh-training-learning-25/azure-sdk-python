//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armredis

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// PatchSchedulesClient contains the methods for the PatchSchedules group.
// Don't use this type directly, use NewPatchSchedulesClient() instead.
type PatchSchedulesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewPatchSchedulesClient creates a new instance of PatchSchedulesClient with the specified values.
func NewPatchSchedulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *PatchSchedulesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &PatchSchedulesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CreateOrUpdate - Create or replace the patching schedule for Redis cache.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PatchSchedulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, name string, defaultParam DefaultName, parameters RedisPatchSchedule, options *PatchSchedulesCreateOrUpdateOptions) (PatchSchedulesCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, name, defaultParam, parameters, options)
	if err != nil {
		return PatchSchedulesCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PatchSchedulesCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return PatchSchedulesCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PatchSchedulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, name string, defaultParam DefaultName, parameters RedisPatchSchedule, options *PatchSchedulesCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/patchSchedules/{default}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if defaultParam == "" {
		return nil, errors.New("parameter defaultParam cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{default}", url.PathEscape(string(defaultParam)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *PatchSchedulesClient) createOrUpdateHandleResponse(resp *http.Response) (PatchSchedulesCreateOrUpdateResponse, error) {
	result := PatchSchedulesCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RedisPatchSchedule); err != nil {
		return PatchSchedulesCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *PatchSchedulesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Deletes the patching schedule of a redis cache.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PatchSchedulesClient) Delete(ctx context.Context, resourceGroupName string, name string, defaultParam DefaultName, options *PatchSchedulesDeleteOptions) (PatchSchedulesDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, name, defaultParam, options)
	if err != nil {
		return PatchSchedulesDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PatchSchedulesDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return PatchSchedulesDeleteResponse{}, client.deleteHandleError(resp)
	}
	return PatchSchedulesDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PatchSchedulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, name string, defaultParam DefaultName, options *PatchSchedulesDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/patchSchedules/{default}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if defaultParam == "" {
		return nil, errors.New("parameter defaultParam cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{default}", url.PathEscape(string(defaultParam)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *PatchSchedulesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets the patching schedule of a redis cache.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PatchSchedulesClient) Get(ctx context.Context, resourceGroupName string, name string, defaultParam DefaultName, options *PatchSchedulesGetOptions) (PatchSchedulesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, name, defaultParam, options)
	if err != nil {
		return PatchSchedulesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PatchSchedulesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PatchSchedulesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PatchSchedulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, name string, defaultParam DefaultName, options *PatchSchedulesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/patchSchedules/{default}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if defaultParam == "" {
		return nil, errors.New("parameter defaultParam cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{default}", url.PathEscape(string(defaultParam)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PatchSchedulesClient) getHandleResponse(resp *http.Response) (PatchSchedulesGetResponse, error) {
	result := PatchSchedulesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RedisPatchSchedule); err != nil {
		return PatchSchedulesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *PatchSchedulesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByRedisResource - Gets all patch schedules in the specified redis cache (there is only one).
// If the operation fails it returns the *ErrorResponse error type.
func (client *PatchSchedulesClient) ListByRedisResource(resourceGroupName string, cacheName string, options *PatchSchedulesListByRedisResourceOptions) *PatchSchedulesListByRedisResourcePager {
	return &PatchSchedulesListByRedisResourcePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByRedisResourceCreateRequest(ctx, resourceGroupName, cacheName, options)
		},
		advancer: func(ctx context.Context, resp PatchSchedulesListByRedisResourceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RedisPatchScheduleListResult.NextLink)
		},
	}
}

// listByRedisResourceCreateRequest creates the ListByRedisResource request.
func (client *PatchSchedulesClient) listByRedisResourceCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, options *PatchSchedulesListByRedisResourceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{cacheName}/patchSchedules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByRedisResourceHandleResponse handles the ListByRedisResource response.
func (client *PatchSchedulesClient) listByRedisResourceHandleResponse(resp *http.Response) (PatchSchedulesListByRedisResourceResponse, error) {
	result := PatchSchedulesListByRedisResourceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RedisPatchScheduleListResult); err != nil {
		return PatchSchedulesListByRedisResourceResponse{}, err
	}
	return result, nil
}

// listByRedisResourceHandleError handles the ListByRedisResource error response.
func (client *PatchSchedulesClient) listByRedisResourceHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
