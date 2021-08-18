// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armredis

import (
	"context"
	"errors"
	"fmt"
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

// ListByRedisCache - Gets the private link resources that need to be created for a redis cache.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PrivateLinkResourcesClient) ListByRedisCache(ctx context.Context, resourceGroupName string, cacheName string, options *PrivateLinkResourcesListByRedisCacheOptions) (PrivateLinkResourcesListByRedisCacheResponse, error) {
	req, err := client.listByRedisCacheCreateRequest(ctx, resourceGroupName, cacheName, options)
	if err != nil {
		return PrivateLinkResourcesListByRedisCacheResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PrivateLinkResourcesListByRedisCacheResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return PrivateLinkResourcesListByRedisCacheResponse{}, client.listByRedisCacheHandleError(resp)
	}
	return client.listByRedisCacheHandleResponse(resp)
}

// listByRedisCacheCreateRequest creates the ListByRedisCache request.
func (client *PrivateLinkResourcesClient) listByRedisCacheCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, options *PrivateLinkResourcesListByRedisCacheOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{cacheName}/privateLinkResources"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByRedisCacheHandleResponse handles the ListByRedisCache response.
func (client *PrivateLinkResourcesClient) listByRedisCacheHandleResponse(resp *azcore.Response) (PrivateLinkResourcesListByRedisCacheResponse, error) {
	result := PrivateLinkResourcesListByRedisCacheResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.PrivateLinkResourceListResult); err != nil {
		return PrivateLinkResourcesListByRedisCacheResponse{}, err
	}
	return result, nil
}

// listByRedisCacheHandleError handles the ListByRedisCache error response.
func (client *PrivateLinkResourcesClient) listByRedisCacheHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
