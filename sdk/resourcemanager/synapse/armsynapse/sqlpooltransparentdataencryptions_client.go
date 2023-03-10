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

// SQLPoolTransparentDataEncryptionsClient contains the methods for the SQLPoolTransparentDataEncryptions group.
// Don't use this type directly, use NewSQLPoolTransparentDataEncryptionsClient() instead.
type SQLPoolTransparentDataEncryptionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSQLPoolTransparentDataEncryptionsClient creates a new instance of SQLPoolTransparentDataEncryptionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSQLPoolTransparentDataEncryptionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SQLPoolTransparentDataEncryptionsClient, error) {
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
	client := &SQLPoolTransparentDataEncryptionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a Sql pool's transparent data encryption configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - sqlPoolName - SQL pool name
//   - transparentDataEncryptionName - The name of the transparent data encryption configuration.
//   - parameters - The required parameters for creating or updating transparent data encryption.
//   - options - SQLPoolTransparentDataEncryptionsClientCreateOrUpdateOptions contains the optional parameters for the SQLPoolTransparentDataEncryptionsClient.CreateOrUpdate
//     method.
func (client *SQLPoolTransparentDataEncryptionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, transparentDataEncryptionName TransparentDataEncryptionName, parameters TransparentDataEncryption, options *SQLPoolTransparentDataEncryptionsClientCreateOrUpdateOptions) (SQLPoolTransparentDataEncryptionsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, transparentDataEncryptionName, parameters, options)
	if err != nil {
		return SQLPoolTransparentDataEncryptionsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLPoolTransparentDataEncryptionsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return SQLPoolTransparentDataEncryptionsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SQLPoolTransparentDataEncryptionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, transparentDataEncryptionName TransparentDataEncryptionName, parameters TransparentDataEncryption, options *SQLPoolTransparentDataEncryptionsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/transparentDataEncryption/{transparentDataEncryptionName}"
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
	if transparentDataEncryptionName == "" {
		return nil, errors.New("parameter transparentDataEncryptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{transparentDataEncryptionName}", url.PathEscape(string(transparentDataEncryptionName)))
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

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SQLPoolTransparentDataEncryptionsClient) createOrUpdateHandleResponse(resp *http.Response) (SQLPoolTransparentDataEncryptionsClientCreateOrUpdateResponse, error) {
	result := SQLPoolTransparentDataEncryptionsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TransparentDataEncryption); err != nil {
		return SQLPoolTransparentDataEncryptionsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Get - Get a SQL pool's transparent data encryption configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - sqlPoolName - SQL pool name
//   - transparentDataEncryptionName - The name of the transparent data encryption configuration.
//   - options - SQLPoolTransparentDataEncryptionsClientGetOptions contains the optional parameters for the SQLPoolTransparentDataEncryptionsClient.Get
//     method.
func (client *SQLPoolTransparentDataEncryptionsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, transparentDataEncryptionName TransparentDataEncryptionName, options *SQLPoolTransparentDataEncryptionsClientGetOptions) (SQLPoolTransparentDataEncryptionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, transparentDataEncryptionName, options)
	if err != nil {
		return SQLPoolTransparentDataEncryptionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLPoolTransparentDataEncryptionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SQLPoolTransparentDataEncryptionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SQLPoolTransparentDataEncryptionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, transparentDataEncryptionName TransparentDataEncryptionName, options *SQLPoolTransparentDataEncryptionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/transparentDataEncryption/{transparentDataEncryptionName}"
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
	if transparentDataEncryptionName == "" {
		return nil, errors.New("parameter transparentDataEncryptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{transparentDataEncryptionName}", url.PathEscape(string(transparentDataEncryptionName)))
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
func (client *SQLPoolTransparentDataEncryptionsClient) getHandleResponse(resp *http.Response) (SQLPoolTransparentDataEncryptionsClientGetResponse, error) {
	result := SQLPoolTransparentDataEncryptionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TransparentDataEncryption); err != nil {
		return SQLPoolTransparentDataEncryptionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get list of SQL pool's transparent data encryption configurations.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - sqlPoolName - SQL pool name
//   - options - SQLPoolTransparentDataEncryptionsClientListOptions contains the optional parameters for the SQLPoolTransparentDataEncryptionsClient.NewListPager
//     method.
func (client *SQLPoolTransparentDataEncryptionsClient) NewListPager(resourceGroupName string, workspaceName string, sqlPoolName string, options *SQLPoolTransparentDataEncryptionsClientListOptions) *runtime.Pager[SQLPoolTransparentDataEncryptionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SQLPoolTransparentDataEncryptionsClientListResponse]{
		More: func(page SQLPoolTransparentDataEncryptionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SQLPoolTransparentDataEncryptionsClientListResponse) (SQLPoolTransparentDataEncryptionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SQLPoolTransparentDataEncryptionsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SQLPoolTransparentDataEncryptionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SQLPoolTransparentDataEncryptionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SQLPoolTransparentDataEncryptionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, options *SQLPoolTransparentDataEncryptionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/transparentDataEncryption"
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

// listHandleResponse handles the List response.
func (client *SQLPoolTransparentDataEncryptionsClient) listHandleResponse(resp *http.Response) (SQLPoolTransparentDataEncryptionsClientListResponse, error) {
	result := SQLPoolTransparentDataEncryptionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TransparentDataEncryptionListResult); err != nil {
		return SQLPoolTransparentDataEncryptionsClientListResponse{}, err
	}
	return result, nil
}
