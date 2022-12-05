//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresql

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

// ServerBasedPerformanceTierClient contains the methods for the ServerBasedPerformanceTier group.
// Don't use this type directly, use NewServerBasedPerformanceTierClient() instead.
type ServerBasedPerformanceTierClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewServerBasedPerformanceTierClient creates a new instance of ServerBasedPerformanceTierClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewServerBasedPerformanceTierClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServerBasedPerformanceTierClient, error) {
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
	client := &ServerBasedPerformanceTierClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListPager - List all the performance tiers for a PostgreSQL server.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-12-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serverName - The name of the server.
// options - ServerBasedPerformanceTierClientListOptions contains the optional parameters for the ServerBasedPerformanceTierClient.List
// method.
func (client *ServerBasedPerformanceTierClient) NewListPager(resourceGroupName string, serverName string, options *ServerBasedPerformanceTierClientListOptions) *runtime.Pager[ServerBasedPerformanceTierClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServerBasedPerformanceTierClientListResponse]{
		More: func(page ServerBasedPerformanceTierClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ServerBasedPerformanceTierClientListResponse) (ServerBasedPerformanceTierClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceGroupName, serverName, options)
			if err != nil {
				return ServerBasedPerformanceTierClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ServerBasedPerformanceTierClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServerBasedPerformanceTierClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ServerBasedPerformanceTierClient) listCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerBasedPerformanceTierClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/performanceTiers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServerBasedPerformanceTierClient) listHandleResponse(resp *http.Response) (ServerBasedPerformanceTierClientListResponse, error) {
	result := ServerBasedPerformanceTierClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PerformanceTierListResult); err != nil {
		return ServerBasedPerformanceTierClientListResponse{}, err
	}
	return result, nil
}