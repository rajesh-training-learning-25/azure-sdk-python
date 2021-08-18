// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// PrivateLinkResourcesClient contains the methods for the PrivateLinkResources group.
// Don't use this type directly, use NewPrivateLinkResourcesClient() instead.
type PrivateLinkResourcesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewPrivateLinkResourcesClient creates a new instance of PrivateLinkResourcesClient with the specified values.
func NewPrivateLinkResourcesClient(con *armcore.Connection, subscriptionID string) *PrivateLinkResourcesClient {
	return &PrivateLinkResourcesClient{con: con, subscriptionID: subscriptionID}
}

// Get - Gets the private link resources that need to be created for a Azure Monitor PrivateLinkScope.
// If the operation fails it returns a generic error.
func (client *PrivateLinkResourcesClient) Get(ctx context.Context, resourceGroupName string, scopeName string, groupName string, options *PrivateLinkResourcesGetOptions) (PrivateLinkResourceResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, scopeName, groupName, options)
	if err != nil {
		return PrivateLinkResourceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PrivateLinkResourceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return PrivateLinkResourceResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateLinkResourcesClient) getCreateRequest(ctx context.Context, resourceGroupName string, scopeName string, groupName string, options *PrivateLinkResourcesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/privateLinkScopes/{scopeName}/privateLinkResources/{groupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if scopeName == "" {
		return nil, errors.New("parameter scopeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopeName}", url.PathEscape(scopeName))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-10-17-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateLinkResourcesClient) getHandleResponse(resp *azcore.Response) (PrivateLinkResourceResponse, error) {
	var val *PrivateLinkResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return PrivateLinkResourceResponse{}, err
	}
	return PrivateLinkResourceResponse{RawResponse: resp.Response, PrivateLinkResource: val}, nil
}

// getHandleError handles the Get error response.
func (client *PrivateLinkResourcesClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListByPrivateLinkScope - Gets the private link resources that need to be created for a Azure Monitor PrivateLinkScope.
// If the operation fails it returns a generic error.
func (client *PrivateLinkResourcesClient) ListByPrivateLinkScope(resourceGroupName string, scopeName string, options *PrivateLinkResourcesListByPrivateLinkScopeOptions) PrivateLinkResourceListResultPager {
	return &privateLinkResourceListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByPrivateLinkScopeCreateRequest(ctx, resourceGroupName, scopeName, options)
		},
		responder: client.listByPrivateLinkScopeHandleResponse,
		errorer:   client.listByPrivateLinkScopeHandleError,
		advancer: func(ctx context.Context, resp PrivateLinkResourceListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.PrivateLinkResourceListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByPrivateLinkScopeCreateRequest creates the ListByPrivateLinkScope request.
func (client *PrivateLinkResourcesClient) listByPrivateLinkScopeCreateRequest(ctx context.Context, resourceGroupName string, scopeName string, options *PrivateLinkResourcesListByPrivateLinkScopeOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/privateLinkScopes/{scopeName}/privateLinkResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if scopeName == "" {
		return nil, errors.New("parameter scopeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopeName}", url.PathEscape(scopeName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-10-17-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByPrivateLinkScopeHandleResponse handles the ListByPrivateLinkScope response.
func (client *PrivateLinkResourcesClient) listByPrivateLinkScopeHandleResponse(resp *azcore.Response) (PrivateLinkResourceListResultResponse, error) {
	var val *PrivateLinkResourceListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return PrivateLinkResourceListResultResponse{}, err
	}
	return PrivateLinkResourceListResultResponse{RawResponse: resp.Response, PrivateLinkResourceListResult: val}, nil
}

// listByPrivateLinkScopeHandleError handles the ListByPrivateLinkScope error response.
func (client *PrivateLinkResourcesClient) listByPrivateLinkScopeHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
