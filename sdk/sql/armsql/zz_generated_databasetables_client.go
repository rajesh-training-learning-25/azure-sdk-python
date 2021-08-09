// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// DatabaseTablesClient contains the methods for the DatabaseTables group.
// Don't use this type directly, use NewDatabaseTablesClient() instead.
type DatabaseTablesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewDatabaseTablesClient creates a new instance of DatabaseTablesClient with the specified values.
func NewDatabaseTablesClient(con *armcore.Connection, subscriptionID string) *DatabaseTablesClient {
	return &DatabaseTablesClient{con: con, subscriptionID: subscriptionID}
}

// Get - Get database table
// If the operation fails it returns a generic error.
func (client *DatabaseTablesClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, options *DatabaseTablesGetOptions) (DatabaseTablesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, databaseName, schemaName, tableName, options)
	if err != nil {
		return DatabaseTablesGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DatabaseTablesGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DatabaseTablesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DatabaseTablesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, options *DatabaseTablesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/schemas/{schemaName}/tables/{tableName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if schemaName == "" {
		return nil, errors.New("parameter schemaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaName}", url.PathEscape(schemaName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
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
	reqQP.Set("api-version", "2020-11-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DatabaseTablesClient) getHandleResponse(resp *azcore.Response) (DatabaseTablesGetResponse, error) {
	result := DatabaseTablesGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.DatabaseTable); err != nil {
		return DatabaseTablesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DatabaseTablesClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListBySchema - List database tables
// If the operation fails it returns a generic error.
func (client *DatabaseTablesClient) ListBySchema(resourceGroupName string, serverName string, databaseName string, schemaName string, options *DatabaseTablesListBySchemaOptions) DatabaseTablesListBySchemaPager {
	return &databaseTablesListBySchemaPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listBySchemaCreateRequest(ctx, resourceGroupName, serverName, databaseName, schemaName, options)
		},
		advancer: func(ctx context.Context, resp DatabaseTablesListBySchemaResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.DatabaseTableListResult.NextLink)
		},
	}
}

// listBySchemaCreateRequest creates the ListBySchema request.
func (client *DatabaseTablesClient) listBySchemaCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, options *DatabaseTablesListBySchemaOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/schemas/{schemaName}/tables"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if schemaName == "" {
		return nil, errors.New("parameter schemaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaName}", url.PathEscape(schemaName))
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
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2020-11-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listBySchemaHandleResponse handles the ListBySchema response.
func (client *DatabaseTablesClient) listBySchemaHandleResponse(resp *azcore.Response) (DatabaseTablesListBySchemaResponse, error) {
	result := DatabaseTablesListBySchemaResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.DatabaseTableListResult); err != nil {
		return DatabaseTablesListBySchemaResponse{}, err
	}
	return result, nil
}

// listBySchemaHandleError handles the ListBySchema error response.
func (client *DatabaseTablesClient) listBySchemaHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
