//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerregistry

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
	"strconv"
	"strings"
)

// RunsClient contains the methods for the Runs group.
// Don't use this type directly, use NewRunsClient() instead.
type RunsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewRunsClient creates a new instance of RunsClient with the specified values.
func NewRunsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *RunsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &RunsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCancel - Cancel an existing run.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RunsClient) BeginCancel(ctx context.Context, resourceGroupName string, registryName string, runID string, options *RunsBeginCancelOptions) (RunsCancelPollerResponse, error) {
	resp, err := client.cancel(ctx, resourceGroupName, registryName, runID, options)
	if err != nil {
		return RunsCancelPollerResponse{}, err
	}
	result := RunsCancelPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("RunsClient.Cancel", "", resp, client.pl, client.cancelHandleError)
	if err != nil {
		return RunsCancelPollerResponse{}, err
	}
	result.Poller = &RunsCancelPoller{
		pt: pt,
	}
	return result, nil
}

// Cancel - Cancel an existing run.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RunsClient) cancel(ctx context.Context, resourceGroupName string, registryName string, runID string, options *RunsBeginCancelOptions) (*http.Response, error) {
	req, err := client.cancelCreateRequest(ctx, resourceGroupName, registryName, runID, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.cancelHandleError(resp)
	}
	return resp, nil
}

// cancelCreateRequest creates the Cancel request.
func (client *RunsClient) cancelCreateRequest(ctx context.Context, resourceGroupName string, registryName string, runID string, options *RunsBeginCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/runs/{runId}/cancel"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// cancelHandleError handles the Cancel error response.
func (client *RunsClient) cancelHandleError(resp *http.Response) error {
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

// Get - Gets the detailed information for a given run.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RunsClient) Get(ctx context.Context, resourceGroupName string, registryName string, runID string, options *RunsGetOptions) (RunsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, runID, options)
	if err != nil {
		return RunsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RunsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RunsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RunsClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, runID string, options *RunsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/runs/{runId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RunsClient) getHandleResponse(resp *http.Response) (RunsGetResponse, error) {
	result := RunsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Run); err != nil {
		return RunsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *RunsClient) getHandleError(resp *http.Response) error {
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

// GetLogSasURL - Gets a link to download the run logs.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RunsClient) GetLogSasURL(ctx context.Context, resourceGroupName string, registryName string, runID string, options *RunsGetLogSasURLOptions) (RunsGetLogSasURLResponse, error) {
	req, err := client.getLogSasURLCreateRequest(ctx, resourceGroupName, registryName, runID, options)
	if err != nil {
		return RunsGetLogSasURLResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RunsGetLogSasURLResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RunsGetLogSasURLResponse{}, client.getLogSasURLHandleError(resp)
	}
	return client.getLogSasURLHandleResponse(resp)
}

// getLogSasURLCreateRequest creates the GetLogSasURL request.
func (client *RunsClient) getLogSasURLCreateRequest(ctx context.Context, resourceGroupName string, registryName string, runID string, options *RunsGetLogSasURLOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/runs/{runId}/listLogSasUrl"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getLogSasURLHandleResponse handles the GetLogSasURL response.
func (client *RunsClient) getLogSasURLHandleResponse(resp *http.Response) (RunsGetLogSasURLResponse, error) {
	result := RunsGetLogSasURLResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RunGetLogResult); err != nil {
		return RunsGetLogSasURLResponse{}, err
	}
	return result, nil
}

// getLogSasURLHandleError handles the GetLogSasURL error response.
func (client *RunsClient) getLogSasURLHandleError(resp *http.Response) error {
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

// List - Gets all the runs for a registry.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RunsClient) List(resourceGroupName string, registryName string, options *RunsListOptions) *RunsListPager {
	return &RunsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, registryName, options)
		},
		advancer: func(ctx context.Context, resp RunsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RunListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *RunsClient) listCreateRequest(ctx context.Context, resourceGroupName string, registryName string, options *RunsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/runs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RunsClient) listHandleResponse(resp *http.Response) (RunsListResponse, error) {
	result := RunsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RunListResult); err != nil {
		return RunsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *RunsClient) listHandleError(resp *http.Response) error {
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

// BeginUpdate - Patch the run properties.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RunsClient) BeginUpdate(ctx context.Context, resourceGroupName string, registryName string, runID string, runUpdateParameters RunUpdateParameters, options *RunsBeginUpdateOptions) (RunsUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, registryName, runID, runUpdateParameters, options)
	if err != nil {
		return RunsUpdatePollerResponse{}, err
	}
	result := RunsUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("RunsClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return RunsUpdatePollerResponse{}, err
	}
	result.Poller = &RunsUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Patch the run properties.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RunsClient) update(ctx context.Context, resourceGroupName string, registryName string, runID string, runUpdateParameters RunUpdateParameters, options *RunsBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, registryName, runID, runUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *RunsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, registryName string, runID string, runUpdateParameters RunUpdateParameters, options *RunsBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/runs/{runId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, runUpdateParameters)
}

// updateHandleError handles the Update error response.
func (client *RunsClient) updateHandleError(resp *http.Response) error {
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
