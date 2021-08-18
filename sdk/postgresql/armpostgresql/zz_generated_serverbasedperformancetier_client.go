// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresql

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

// ServerBasedPerformanceTierClient contains the methods for the ServerBasedPerformanceTier group.
// Don't use this type directly, use NewServerBasedPerformanceTierClient() instead.
type ServerBasedPerformanceTierClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewServerBasedPerformanceTierClient creates a new instance of ServerBasedPerformanceTierClient with the specified values.
func NewServerBasedPerformanceTierClient(con *armcore.Connection, subscriptionID string) *ServerBasedPerformanceTierClient {
	return &ServerBasedPerformanceTierClient{con: con, subscriptionID: subscriptionID}
}

// List - List all the performance tiers for a PostgreSQL server.
// If the operation fails it returns the *CloudError error type.
func (client *ServerBasedPerformanceTierClient) List(ctx context.Context, resourceGroupName string, serverName string, options *ServerBasedPerformanceTierListOptions) (ServerBasedPerformanceTierListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return ServerBasedPerformanceTierListResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ServerBasedPerformanceTierListResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ServerBasedPerformanceTierListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ServerBasedPerformanceTierClient) listCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerBasedPerformanceTierListOptions) (*azcore.Request, error) {
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
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServerBasedPerformanceTierClient) listHandleResponse(resp *azcore.Response) (ServerBasedPerformanceTierListResponse, error) {
	result := ServerBasedPerformanceTierListResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.PerformanceTierListResult); err != nil {
		return ServerBasedPerformanceTierListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ServerBasedPerformanceTierClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
