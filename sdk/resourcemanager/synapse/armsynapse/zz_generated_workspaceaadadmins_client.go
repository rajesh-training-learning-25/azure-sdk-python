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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// WorkspaceAADAdminsClient contains the methods for the WorkspaceAADAdmins group.
// Don't use this type directly, use NewWorkspaceAADAdminsClient() instead.
type WorkspaceAADAdminsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWorkspaceAADAdminsClient creates a new instance of WorkspaceAADAdminsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWorkspaceAADAdminsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *WorkspaceAADAdminsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &WorkspaceAADAdminsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a workspace active directory admin
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// aadAdminInfo - Workspace active directory administrator properties
// options - WorkspaceAADAdminsClientBeginCreateOrUpdateOptions contains the optional parameters for the WorkspaceAADAdminsClient.BeginCreateOrUpdate
// method.
func (client *WorkspaceAADAdminsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, aadAdminInfo WorkspaceAADAdminInfo, options *WorkspaceAADAdminsClientBeginCreateOrUpdateOptions) (WorkspaceAADAdminsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, aadAdminInfo, options)
	if err != nil {
		return WorkspaceAADAdminsClientCreateOrUpdatePollerResponse{}, err
	}
	result := WorkspaceAADAdminsClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("WorkspaceAADAdminsClient.CreateOrUpdate", "location", resp, client.pl)
	if err != nil {
		return WorkspaceAADAdminsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &WorkspaceAADAdminsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a workspace active directory admin
// If the operation fails it returns an *azcore.ResponseError type.
func (client *WorkspaceAADAdminsClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, aadAdminInfo WorkspaceAADAdminInfo, options *WorkspaceAADAdminsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, aadAdminInfo, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WorkspaceAADAdminsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, aadAdminInfo WorkspaceAADAdminInfo, options *WorkspaceAADAdminsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/administrators/activeDirectory"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, aadAdminInfo)
}

// BeginDelete - Deletes a workspace active directory admin
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// options - WorkspaceAADAdminsClientBeginDeleteOptions contains the optional parameters for the WorkspaceAADAdminsClient.BeginDelete
// method.
func (client *WorkspaceAADAdminsClient) BeginDelete(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspaceAADAdminsClientBeginDeleteOptions) (WorkspaceAADAdminsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, workspaceName, options)
	if err != nil {
		return WorkspaceAADAdminsClientDeletePollerResponse{}, err
	}
	result := WorkspaceAADAdminsClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("WorkspaceAADAdminsClient.Delete", "location", resp, client.pl)
	if err != nil {
		return WorkspaceAADAdminsClientDeletePollerResponse{}, err
	}
	result.Poller = &WorkspaceAADAdminsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a workspace active directory admin
// If the operation fails it returns an *azcore.ResponseError type.
func (client *WorkspaceAADAdminsClient) deleteOperation(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspaceAADAdminsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WorkspaceAADAdminsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspaceAADAdminsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/administrators/activeDirectory"
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets a workspace active directory admin
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// options - WorkspaceAADAdminsClientGetOptions contains the optional parameters for the WorkspaceAADAdminsClient.Get method.
func (client *WorkspaceAADAdminsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspaceAADAdminsClientGetOptions) (WorkspaceAADAdminsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, options)
	if err != nil {
		return WorkspaceAADAdminsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceAADAdminsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceAADAdminsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkspaceAADAdminsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspaceAADAdminsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/administrators/activeDirectory"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
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
func (client *WorkspaceAADAdminsClient) getHandleResponse(resp *http.Response) (WorkspaceAADAdminsClientGetResponse, error) {
	result := WorkspaceAADAdminsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceAADAdminInfo); err != nil {
		return WorkspaceAADAdminsClientGetResponse{}, err
	}
	return result, nil
}
