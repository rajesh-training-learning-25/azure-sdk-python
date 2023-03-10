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

// SQLPoolGeoBackupPoliciesClient contains the methods for the SQLPoolGeoBackupPolicies group.
// Don't use this type directly, use NewSQLPoolGeoBackupPoliciesClient() instead.
type SQLPoolGeoBackupPoliciesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSQLPoolGeoBackupPoliciesClient creates a new instance of SQLPoolGeoBackupPoliciesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSQLPoolGeoBackupPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SQLPoolGeoBackupPoliciesClient, error) {
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
	client := &SQLPoolGeoBackupPoliciesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Updates a SQL Pool geo backup policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - sqlPoolName - SQL pool name
//   - geoBackupPolicyName - The name of the geo backup policy.
//   - parameters - The required parameters for creating or updating the geo backup policy.
//   - options - SQLPoolGeoBackupPoliciesClientCreateOrUpdateOptions contains the optional parameters for the SQLPoolGeoBackupPoliciesClient.CreateOrUpdate
//     method.
func (client *SQLPoolGeoBackupPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, geoBackupPolicyName GeoBackupPolicyName, parameters GeoBackupPolicy, options *SQLPoolGeoBackupPoliciesClientCreateOrUpdateOptions) (SQLPoolGeoBackupPoliciesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, geoBackupPolicyName, parameters, options)
	if err != nil {
		return SQLPoolGeoBackupPoliciesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLPoolGeoBackupPoliciesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return SQLPoolGeoBackupPoliciesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SQLPoolGeoBackupPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, geoBackupPolicyName GeoBackupPolicyName, parameters GeoBackupPolicy, options *SQLPoolGeoBackupPoliciesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/geoBackupPolicies/{geoBackupPolicyName}"
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
	if geoBackupPolicyName == "" {
		return nil, errors.New("parameter geoBackupPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{geoBackupPolicyName}", url.PathEscape(string(geoBackupPolicyName)))
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
func (client *SQLPoolGeoBackupPoliciesClient) createOrUpdateHandleResponse(resp *http.Response) (SQLPoolGeoBackupPoliciesClientCreateOrUpdateResponse, error) {
	result := SQLPoolGeoBackupPoliciesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GeoBackupPolicy); err != nil {
		return SQLPoolGeoBackupPoliciesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Get - Get the specified SQL pool geo backup policy
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - sqlPoolName - SQL pool name
//   - geoBackupPolicyName - The name of the geo backup policy.
//   - options - SQLPoolGeoBackupPoliciesClientGetOptions contains the optional parameters for the SQLPoolGeoBackupPoliciesClient.Get
//     method.
func (client *SQLPoolGeoBackupPoliciesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, geoBackupPolicyName GeoBackupPolicyName, options *SQLPoolGeoBackupPoliciesClientGetOptions) (SQLPoolGeoBackupPoliciesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, geoBackupPolicyName, options)
	if err != nil {
		return SQLPoolGeoBackupPoliciesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLPoolGeoBackupPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SQLPoolGeoBackupPoliciesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SQLPoolGeoBackupPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, geoBackupPolicyName GeoBackupPolicyName, options *SQLPoolGeoBackupPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/geoBackupPolicies/{geoBackupPolicyName}"
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
	if geoBackupPolicyName == "" {
		return nil, errors.New("parameter geoBackupPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{geoBackupPolicyName}", url.PathEscape(string(geoBackupPolicyName)))
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
func (client *SQLPoolGeoBackupPoliciesClient) getHandleResponse(resp *http.Response) (SQLPoolGeoBackupPoliciesClientGetResponse, error) {
	result := SQLPoolGeoBackupPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GeoBackupPolicy); err != nil {
		return SQLPoolGeoBackupPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get list of SQL pool geo backup policies
//
// Generated from API version 2021-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - sqlPoolName - SQL pool name
//   - options - SQLPoolGeoBackupPoliciesClientListOptions contains the optional parameters for the SQLPoolGeoBackupPoliciesClient.NewListPager
//     method.
func (client *SQLPoolGeoBackupPoliciesClient) NewListPager(resourceGroupName string, workspaceName string, sqlPoolName string, options *SQLPoolGeoBackupPoliciesClientListOptions) *runtime.Pager[SQLPoolGeoBackupPoliciesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SQLPoolGeoBackupPoliciesClientListResponse]{
		More: func(page SQLPoolGeoBackupPoliciesClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *SQLPoolGeoBackupPoliciesClientListResponse) (SQLPoolGeoBackupPoliciesClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, options)
			if err != nil {
				return SQLPoolGeoBackupPoliciesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SQLPoolGeoBackupPoliciesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SQLPoolGeoBackupPoliciesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SQLPoolGeoBackupPoliciesClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, options *SQLPoolGeoBackupPoliciesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/geoBackupPolicies"
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
func (client *SQLPoolGeoBackupPoliciesClient) listHandleResponse(resp *http.Response) (SQLPoolGeoBackupPoliciesClientListResponse, error) {
	result := SQLPoolGeoBackupPoliciesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GeoBackupPolicyListResult); err != nil {
		return SQLPoolGeoBackupPoliciesClientListResponse{}, err
	}
	return result, nil
}
