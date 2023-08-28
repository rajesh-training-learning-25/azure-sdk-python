//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ElasticPoolOperationsClient contains the methods for the ElasticPoolOperations group.
// Don't use this type directly, use NewElasticPoolOperationsClient() instead.
type ElasticPoolOperationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewElasticPoolOperationsClient creates a new instance of ElasticPoolOperationsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewElasticPoolOperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ElasticPoolOperationsClient, error) {
	cl, err := arm.NewClient(moduleName+".ElasticPoolOperationsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ElasticPoolOperationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Cancel - Cancels the asynchronous operation on the elastic pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - operationID - The operation identifier.
//   - options - ElasticPoolOperationsClientCancelOptions contains the optional parameters for the ElasticPoolOperationsClient.Cancel
//     method.
func (client *ElasticPoolOperationsClient) Cancel(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, operationID string, options *ElasticPoolOperationsClientCancelOptions) (ElasticPoolOperationsClientCancelResponse, error) {
	var err error
	req, err := client.cancelCreateRequest(ctx, resourceGroupName, serverName, elasticPoolName, operationID, options)
	if err != nil {
		return ElasticPoolOperationsClientCancelResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ElasticPoolOperationsClientCancelResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ElasticPoolOperationsClientCancelResponse{}, err
	}
	return ElasticPoolOperationsClientCancelResponse{}, nil
}

// cancelCreateRequest creates the Cancel request.
func (client *ElasticPoolOperationsClient) cancelCreateRequest(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, operationID string, options *ElasticPoolOperationsClientCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}/operations/{operationId}/cancel"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if elasticPoolName == "" {
		return nil, errors.New("parameter elasticPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{elasticPoolName}", url.PathEscape(elasticPoolName))
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// NewListByElasticPoolPager - Gets a list of operations performed on the elastic pool.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - options - ElasticPoolOperationsClientListByElasticPoolOptions contains the optional parameters for the ElasticPoolOperationsClient.NewListByElasticPoolPager
//     method.
func (client *ElasticPoolOperationsClient) NewListByElasticPoolPager(resourceGroupName string, serverName string, elasticPoolName string, options *ElasticPoolOperationsClientListByElasticPoolOptions) *runtime.Pager[ElasticPoolOperationsClientListByElasticPoolResponse] {
	return runtime.NewPager(runtime.PagingHandler[ElasticPoolOperationsClientListByElasticPoolResponse]{
		More: func(page ElasticPoolOperationsClientListByElasticPoolResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ElasticPoolOperationsClientListByElasticPoolResponse) (ElasticPoolOperationsClientListByElasticPoolResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByElasticPoolCreateRequest(ctx, resourceGroupName, serverName, elasticPoolName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ElasticPoolOperationsClientListByElasticPoolResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ElasticPoolOperationsClientListByElasticPoolResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ElasticPoolOperationsClientListByElasticPoolResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByElasticPoolHandleResponse(resp)
		},
	})
}

// listByElasticPoolCreateRequest creates the ListByElasticPool request.
func (client *ElasticPoolOperationsClient) listByElasticPoolCreateRequest(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, options *ElasticPoolOperationsClientListByElasticPoolOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}/operations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if elasticPoolName == "" {
		return nil, errors.New("parameter elasticPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{elasticPoolName}", url.PathEscape(elasticPoolName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByElasticPoolHandleResponse handles the ListByElasticPool response.
func (client *ElasticPoolOperationsClient) listByElasticPoolHandleResponse(resp *http.Response) (ElasticPoolOperationsClientListByElasticPoolResponse, error) {
	result := ElasticPoolOperationsClientListByElasticPoolResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ElasticPoolOperationListResult); err != nil {
		return ElasticPoolOperationsClientListByElasticPoolResponse{}, err
	}
	return result, nil
}
