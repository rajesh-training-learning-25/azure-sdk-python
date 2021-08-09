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

// DataWarehouseUserActivitiesClient contains the methods for the DataWarehouseUserActivities group.
// Don't use this type directly, use NewDataWarehouseUserActivitiesClient() instead.
type DataWarehouseUserActivitiesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewDataWarehouseUserActivitiesClient creates a new instance of DataWarehouseUserActivitiesClient with the specified values.
func NewDataWarehouseUserActivitiesClient(con *armcore.Connection, subscriptionID string) *DataWarehouseUserActivitiesClient {
	return &DataWarehouseUserActivitiesClient{con: con, subscriptionID: subscriptionID}
}

// Get - Gets the user activities of a data warehouse which includes running and suspended queries
// If the operation fails it returns a generic error.
func (client *DataWarehouseUserActivitiesClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, dataWarehouseUserActivityName DataWarehouseUserActivityName, options *DataWarehouseUserActivitiesGetOptions) (DataWarehouseUserActivitiesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, databaseName, dataWarehouseUserActivityName, options)
	if err != nil {
		return DataWarehouseUserActivitiesGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DataWarehouseUserActivitiesGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DataWarehouseUserActivitiesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DataWarehouseUserActivitiesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, dataWarehouseUserActivityName DataWarehouseUserActivityName, options *DataWarehouseUserActivitiesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/dataWarehouseUserActivities/{dataWarehouseUserActivityName}"
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
	if dataWarehouseUserActivityName == "" {
		return nil, errors.New("parameter dataWarehouseUserActivityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataWarehouseUserActivityName}", url.PathEscape(string(dataWarehouseUserActivityName)))
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
func (client *DataWarehouseUserActivitiesClient) getHandleResponse(resp *azcore.Response) (DataWarehouseUserActivitiesGetResponse, error) {
	result := DataWarehouseUserActivitiesGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.DataWarehouseUserActivities); err != nil {
		return DataWarehouseUserActivitiesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DataWarehouseUserActivitiesClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListByDatabase - List the user activities of a data warehouse which includes running and suspended queries
// If the operation fails it returns a generic error.
func (client *DataWarehouseUserActivitiesClient) ListByDatabase(resourceGroupName string, serverName string, databaseName string, options *DataWarehouseUserActivitiesListByDatabaseOptions) DataWarehouseUserActivitiesListByDatabasePager {
	return &dataWarehouseUserActivitiesListByDatabasePager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByDatabaseCreateRequest(ctx, resourceGroupName, serverName, databaseName, options)
		},
		advancer: func(ctx context.Context, resp DataWarehouseUserActivitiesListByDatabaseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.DataWarehouseUserActivitiesListResult.NextLink)
		},
	}
}

// listByDatabaseCreateRequest creates the ListByDatabase request.
func (client *DataWarehouseUserActivitiesClient) listByDatabaseCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, options *DataWarehouseUserActivitiesListByDatabaseOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/dataWarehouseUserActivities"
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

// listByDatabaseHandleResponse handles the ListByDatabase response.
func (client *DataWarehouseUserActivitiesClient) listByDatabaseHandleResponse(resp *azcore.Response) (DataWarehouseUserActivitiesListByDatabaseResponse, error) {
	result := DataWarehouseUserActivitiesListByDatabaseResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.DataWarehouseUserActivitiesListResult); err != nil {
		return DataWarehouseUserActivitiesListByDatabaseResponse{}, err
	}
	return result, nil
}

// listByDatabaseHandleError handles the ListByDatabase error response.
func (client *DataWarehouseUserActivitiesClient) listByDatabaseHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
