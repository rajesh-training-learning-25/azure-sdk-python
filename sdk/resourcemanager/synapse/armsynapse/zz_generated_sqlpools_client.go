//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

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

// SQLPoolsClient contains the methods for the SQLPools group.
// Don't use this type directly, use NewSQLPoolsClient() instead.
type SQLPoolsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewSQLPoolsClient creates a new instance of SQLPoolsClient with the specified values.
func NewSQLPoolsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SQLPoolsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &SQLPoolsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreate - Create a SQL pool
// If the operation fails it returns the *ErrorResponse error type.
func (client *SQLPoolsClient) BeginCreate(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, sqlPoolInfo SQLPool, options *SQLPoolsBeginCreateOptions) (SQLPoolsCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, workspaceName, sqlPoolName, sqlPoolInfo, options)
	if err != nil {
		return SQLPoolsCreatePollerResponse{}, err
	}
	result := SQLPoolsCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SQLPoolsClient.Create", "location", resp, client.pl, client.createHandleError)
	if err != nil {
		return SQLPoolsCreatePollerResponse{}, err
	}
	result.Poller = &SQLPoolsCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Create a SQL pool
// If the operation fails it returns the *ErrorResponse error type.
func (client *SQLPoolsClient) create(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, sqlPoolInfo SQLPool, options *SQLPoolsBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, sqlPoolInfo, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNotFound) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *SQLPoolsClient) createCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, sqlPoolInfo SQLPool, options *SQLPoolsBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, sqlPoolInfo)
}

// createHandleError handles the Create error response.
func (client *SQLPoolsClient) createHandleError(resp *http.Response) error {
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

// BeginDelete - Delete a SQL pool
// If the operation fails it returns the *ErrorResponse error type.
func (client *SQLPoolsClient) BeginDelete(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, options *SQLPoolsBeginDeleteOptions) (SQLPoolsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, workspaceName, sqlPoolName, options)
	if err != nil {
		return SQLPoolsDeletePollerResponse{}, err
	}
	result := SQLPoolsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SQLPoolsClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return SQLPoolsDeletePollerResponse{}, err
	}
	result.Poller = &SQLPoolsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete a SQL pool
// If the operation fails it returns the *ErrorResponse error type.
func (client *SQLPoolsClient) deleteOperation(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, options *SQLPoolsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SQLPoolsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, options *SQLPoolsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *SQLPoolsClient) deleteHandleError(resp *http.Response) error {
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

// Get - Get SQL pool properties
// If the operation fails it returns the *ErrorResponse error type.
func (client *SQLPoolsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, options *SQLPoolsGetOptions) (SQLPoolsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, options)
	if err != nil {
		return SQLPoolsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLPoolsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SQLPoolsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SQLPoolsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, options *SQLPoolsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SQLPoolsClient) getHandleResponse(resp *http.Response) (SQLPoolsGetResponse, error) {
	result := SQLPoolsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLPool); err != nil {
		return SQLPoolsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *SQLPoolsClient) getHandleError(resp *http.Response) error {
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

// ListByWorkspace - List all SQL pools
// If the operation fails it returns the *ErrorResponse error type.
func (client *SQLPoolsClient) ListByWorkspace(resourceGroupName string, workspaceName string, options *SQLPoolsListByWorkspaceOptions) *SQLPoolsListByWorkspacePager {
	return &SQLPoolsListByWorkspacePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
		},
		advancer: func(ctx context.Context, resp SQLPoolsListByWorkspaceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SQLPoolInfoListResult.NextLink)
		},
	}
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *SQLPoolsClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *SQLPoolsListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *SQLPoolsClient) listByWorkspaceHandleResponse(resp *http.Response) (SQLPoolsListByWorkspaceResponse, error) {
	result := SQLPoolsListByWorkspaceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLPoolInfoListResult); err != nil {
		return SQLPoolsListByWorkspaceResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByWorkspaceHandleError handles the ListByWorkspace error response.
func (client *SQLPoolsClient) listByWorkspaceHandleError(resp *http.Response) error {
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

// BeginPause - Pause a SQL pool
// If the operation fails it returns the *ErrorResponse error type.
func (client *SQLPoolsClient) BeginPause(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, options *SQLPoolsBeginPauseOptions) (SQLPoolsPausePollerResponse, error) {
	resp, err := client.pause(ctx, resourceGroupName, workspaceName, sqlPoolName, options)
	if err != nil {
		return SQLPoolsPausePollerResponse{}, err
	}
	result := SQLPoolsPausePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SQLPoolsClient.Pause", "location", resp, client.pl, client.pauseHandleError)
	if err != nil {
		return SQLPoolsPausePollerResponse{}, err
	}
	result.Poller = &SQLPoolsPausePoller{
		pt: pt,
	}
	return result, nil
}

// Pause - Pause a SQL pool
// If the operation fails it returns the *ErrorResponse error type.
func (client *SQLPoolsClient) pause(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, options *SQLPoolsBeginPauseOptions) (*http.Response, error) {
	req, err := client.pauseCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.pauseHandleError(resp)
	}
	return resp, nil
}

// pauseCreateRequest creates the Pause request.
func (client *SQLPoolsClient) pauseCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, options *SQLPoolsBeginPauseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/pause"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// pauseHandleError handles the Pause error response.
func (client *SQLPoolsClient) pauseHandleError(resp *http.Response) error {
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

// Rename - Rename a SQL pool.
// If the operation fails it returns a generic error.
func (client *SQLPoolsClient) Rename(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, parameters ResourceMoveDefinition, options *SQLPoolsRenameOptions) (SQLPoolsRenameResponse, error) {
	req, err := client.renameCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, parameters, options)
	if err != nil {
		return SQLPoolsRenameResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLPoolsRenameResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SQLPoolsRenameResponse{}, client.renameHandleError(resp)
	}
	return SQLPoolsRenameResponse{RawResponse: resp}, nil
}

// renameCreateRequest creates the Rename request.
func (client *SQLPoolsClient) renameCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, parameters ResourceMoveDefinition, options *SQLPoolsRenameOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/move"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, runtime.MarshalAsJSON(req, parameters)
}

// renameHandleError handles the Rename error response.
func (client *SQLPoolsClient) renameHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginResume - Resume a SQL pool
// If the operation fails it returns the *ErrorResponse error type.
func (client *SQLPoolsClient) BeginResume(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, options *SQLPoolsBeginResumeOptions) (SQLPoolsResumePollerResponse, error) {
	resp, err := client.resume(ctx, resourceGroupName, workspaceName, sqlPoolName, options)
	if err != nil {
		return SQLPoolsResumePollerResponse{}, err
	}
	result := SQLPoolsResumePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SQLPoolsClient.Resume", "location", resp, client.pl, client.resumeHandleError)
	if err != nil {
		return SQLPoolsResumePollerResponse{}, err
	}
	result.Poller = &SQLPoolsResumePoller{
		pt: pt,
	}
	return result, nil
}

// Resume - Resume a SQL pool
// If the operation fails it returns the *ErrorResponse error type.
func (client *SQLPoolsClient) resume(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, options *SQLPoolsBeginResumeOptions) (*http.Response, error) {
	req, err := client.resumeCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.resumeHandleError(resp)
	}
	return resp, nil
}

// resumeCreateRequest creates the Resume request.
func (client *SQLPoolsClient) resumeCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, options *SQLPoolsBeginResumeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/resume"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// resumeHandleError handles the Resume error response.
func (client *SQLPoolsClient) resumeHandleError(resp *http.Response) error {
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

// Update - Apply a partial update to a SQL pool
// If the operation fails it returns the *ErrorResponse error type.
func (client *SQLPoolsClient) Update(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, sqlPoolInfo SQLPoolPatchInfo, options *SQLPoolsUpdateOptions) (SQLPoolsUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, sqlPoolInfo, options)
	if err != nil {
		return SQLPoolsUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLPoolsUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return SQLPoolsUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *SQLPoolsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, sqlPoolInfo SQLPoolPatchInfo, options *SQLPoolsUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, sqlPoolInfo)
}

// updateHandleResponse handles the Update response.
func (client *SQLPoolsClient) updateHandleResponse(resp *http.Response) (SQLPoolsUpdateResponse, error) {
	result := SQLPoolsUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLPool); err != nil {
		return SQLPoolsUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *SQLPoolsClient) updateHandleError(resp *http.Response) error {
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
