//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsynapse

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// WorkspaceManagedSQLServerBlobAuditingPoliciesClient contains the methods for the WorkspaceManagedSQLServerBlobAuditingPolicies group.
// Don't use this type directly, use NewWorkspaceManagedSQLServerBlobAuditingPoliciesClient() instead.
type WorkspaceManagedSQLServerBlobAuditingPoliciesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWorkspaceManagedSQLServerBlobAuditingPoliciesClient creates a new instance of WorkspaceManagedSQLServerBlobAuditingPoliciesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWorkspaceManagedSQLServerBlobAuditingPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkspaceManagedSQLServerBlobAuditingPoliciesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &WorkspaceManagedSQLServerBlobAuditingPoliciesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or Update a workspace managed sql server's blob auditing policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - blobAuditingPolicyName - The name of the blob auditing policy.
//   - parameters - Properties of extended blob auditing policy.
//   - options - WorkspaceManagedSQLServerBlobAuditingPoliciesClientBeginCreateOrUpdateOptions contains the optional parameters
//     for the WorkspaceManagedSQLServerBlobAuditingPoliciesClient.BeginCreateOrUpdate method.
func (client *WorkspaceManagedSQLServerBlobAuditingPoliciesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, blobAuditingPolicyName BlobAuditingPolicyName, parameters ServerBlobAuditingPolicy, options *WorkspaceManagedSQLServerBlobAuditingPoliciesClientBeginCreateOrUpdateOptions) (*runtime.Poller[WorkspaceManagedSQLServerBlobAuditingPoliciesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, blobAuditingPolicyName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[WorkspaceManagedSQLServerBlobAuditingPoliciesClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[WorkspaceManagedSQLServerBlobAuditingPoliciesClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or Update a workspace managed sql server's blob auditing policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
func (client *WorkspaceManagedSQLServerBlobAuditingPoliciesClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, blobAuditingPolicyName BlobAuditingPolicyName, parameters ServerBlobAuditingPolicy, options *WorkspaceManagedSQLServerBlobAuditingPoliciesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, blobAuditingPolicyName, parameters, options)
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
func (client *WorkspaceManagedSQLServerBlobAuditingPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, blobAuditingPolicyName BlobAuditingPolicyName, parameters ServerBlobAuditingPolicy, options *WorkspaceManagedSQLServerBlobAuditingPoliciesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/auditingSettings/{blobAuditingPolicyName}"
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
	if blobAuditingPolicyName == "" {
		return nil, errors.New("parameter blobAuditingPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blobAuditingPolicyName}", url.PathEscape(string(blobAuditingPolicyName)))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// Get - Get a workspace managed sql server's blob auditing policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - blobAuditingPolicyName - The name of the blob auditing policy.
//   - options - WorkspaceManagedSQLServerBlobAuditingPoliciesClientGetOptions contains the optional parameters for the WorkspaceManagedSQLServerBlobAuditingPoliciesClient.Get
//     method.
func (client *WorkspaceManagedSQLServerBlobAuditingPoliciesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, blobAuditingPolicyName BlobAuditingPolicyName, options *WorkspaceManagedSQLServerBlobAuditingPoliciesClientGetOptions) (WorkspaceManagedSQLServerBlobAuditingPoliciesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, blobAuditingPolicyName, options)
	if err != nil {
		return WorkspaceManagedSQLServerBlobAuditingPoliciesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceManagedSQLServerBlobAuditingPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceManagedSQLServerBlobAuditingPoliciesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkspaceManagedSQLServerBlobAuditingPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, blobAuditingPolicyName BlobAuditingPolicyName, options *WorkspaceManagedSQLServerBlobAuditingPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/auditingSettings/{blobAuditingPolicyName}"
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
	if blobAuditingPolicyName == "" {
		return nil, errors.New("parameter blobAuditingPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blobAuditingPolicyName}", url.PathEscape(string(blobAuditingPolicyName)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkspaceManagedSQLServerBlobAuditingPoliciesClient) getHandleResponse(resp *http.Response) (WorkspaceManagedSQLServerBlobAuditingPoliciesClientGetResponse, error) {
	result := WorkspaceManagedSQLServerBlobAuditingPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerBlobAuditingPolicy); err != nil {
		return WorkspaceManagedSQLServerBlobAuditingPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByWorkspacePager - List workspace managed sql server's blob auditing policies.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - WorkspaceManagedSQLServerBlobAuditingPoliciesClientListByWorkspaceOptions contains the optional parameters for
//     the WorkspaceManagedSQLServerBlobAuditingPoliciesClient.NewListByWorkspacePager method.
func (client *WorkspaceManagedSQLServerBlobAuditingPoliciesClient) NewListByWorkspacePager(resourceGroupName string, workspaceName string, options *WorkspaceManagedSQLServerBlobAuditingPoliciesClientListByWorkspaceOptions) *runtime.Pager[WorkspaceManagedSQLServerBlobAuditingPoliciesClientListByWorkspaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkspaceManagedSQLServerBlobAuditingPoliciesClientListByWorkspaceResponse]{
		More: func(page WorkspaceManagedSQLServerBlobAuditingPoliciesClientListByWorkspaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkspaceManagedSQLServerBlobAuditingPoliciesClientListByWorkspaceResponse) (WorkspaceManagedSQLServerBlobAuditingPoliciesClientListByWorkspaceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return WorkspaceManagedSQLServerBlobAuditingPoliciesClientListByWorkspaceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return WorkspaceManagedSQLServerBlobAuditingPoliciesClientListByWorkspaceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WorkspaceManagedSQLServerBlobAuditingPoliciesClientListByWorkspaceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByWorkspaceHandleResponse(resp)
		},
	})
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *WorkspaceManagedSQLServerBlobAuditingPoliciesClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspaceManagedSQLServerBlobAuditingPoliciesClientListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/auditingSettings"
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
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *WorkspaceManagedSQLServerBlobAuditingPoliciesClient) listByWorkspaceHandleResponse(resp *http.Response) (WorkspaceManagedSQLServerBlobAuditingPoliciesClientListByWorkspaceResponse, error) {
	result := WorkspaceManagedSQLServerBlobAuditingPoliciesClientListByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerBlobAuditingPolicyListResult); err != nil {
		return WorkspaceManagedSQLServerBlobAuditingPoliciesClientListByWorkspaceResponse{}, err
	}
	return result, nil
}
