//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DatabaseClient contains the methods for the Database group.
// Don't use this type directly, use NewDatabaseClient() instead.
type DatabaseClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewDatabaseClient creates a new instance of DatabaseClient with the specified values.
func NewDatabaseClient(con *arm.Connection, subscriptionID string) *DatabaseClient {
	return &DatabaseClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// ListMetricDefinitions - Retrieves metric definitions for the given database.
// If the operation fails it returns a generic error.
func (client *DatabaseClient) ListMetricDefinitions(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, options *DatabaseListMetricDefinitionsOptions) (DatabaseListMetricDefinitionsResponse, error) {
	req, err := client.listMetricDefinitionsCreateRequest(ctx, resourceGroupName, accountName, databaseRid, options)
	if err != nil {
		return DatabaseListMetricDefinitionsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DatabaseListMetricDefinitionsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DatabaseListMetricDefinitionsResponse{}, client.listMetricDefinitionsHandleError(resp)
	}
	return client.listMetricDefinitionsHandleResponse(resp)
}

// listMetricDefinitionsCreateRequest creates the ListMetricDefinitions request.
func (client *DatabaseClient) listMetricDefinitionsCreateRequest(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, options *DatabaseListMetricDefinitionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/metricDefinitions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if databaseRid == "" {
		return nil, errors.New("parameter databaseRid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseRid}", url.PathEscape(databaseRid))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listMetricDefinitionsHandleResponse handles the ListMetricDefinitions response.
func (client *DatabaseClient) listMetricDefinitionsHandleResponse(resp *http.Response) (DatabaseListMetricDefinitionsResponse, error) {
	result := DatabaseListMetricDefinitionsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricDefinitionsListResult); err != nil {
		return DatabaseListMetricDefinitionsResponse{}, err
	}
	return result, nil
}

// listMetricDefinitionsHandleError handles the ListMetricDefinitions error response.
func (client *DatabaseClient) listMetricDefinitionsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListMetrics - Retrieves the metrics determined by the given filter for the given database account and database.
// If the operation fails it returns a generic error.
func (client *DatabaseClient) ListMetrics(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, filter string, options *DatabaseListMetricsOptions) (DatabaseListMetricsResponse, error) {
	req, err := client.listMetricsCreateRequest(ctx, resourceGroupName, accountName, databaseRid, filter, options)
	if err != nil {
		return DatabaseListMetricsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DatabaseListMetricsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DatabaseListMetricsResponse{}, client.listMetricsHandleError(resp)
	}
	return client.listMetricsHandleResponse(resp)
}

// listMetricsCreateRequest creates the ListMetrics request.
func (client *DatabaseClient) listMetricsCreateRequest(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, filter string, options *DatabaseListMetricsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/metrics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if databaseRid == "" {
		return nil, errors.New("parameter databaseRid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseRid}", url.PathEscape(databaseRid))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-15")
	reqQP.Set("$filter", filter)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listMetricsHandleResponse handles the ListMetrics response.
func (client *DatabaseClient) listMetricsHandleResponse(resp *http.Response) (DatabaseListMetricsResponse, error) {
	result := DatabaseListMetricsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricListResult); err != nil {
		return DatabaseListMetricsResponse{}, err
	}
	return result, nil
}

// listMetricsHandleError handles the ListMetrics error response.
func (client *DatabaseClient) listMetricsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListUsages - Retrieves the usages (most recent data) for the given database.
// If the operation fails it returns a generic error.
func (client *DatabaseClient) ListUsages(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, options *DatabaseListUsagesOptions) (DatabaseListUsagesResponse, error) {
	req, err := client.listUsagesCreateRequest(ctx, resourceGroupName, accountName, databaseRid, options)
	if err != nil {
		return DatabaseListUsagesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DatabaseListUsagesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DatabaseListUsagesResponse{}, client.listUsagesHandleError(resp)
	}
	return client.listUsagesHandleResponse(resp)
}

// listUsagesCreateRequest creates the ListUsages request.
func (client *DatabaseClient) listUsagesCreateRequest(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, options *DatabaseListUsagesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/usages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if databaseRid == "" {
		return nil, errors.New("parameter databaseRid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseRid}", url.PathEscape(databaseRid))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-15")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listUsagesHandleResponse handles the ListUsages response.
func (client *DatabaseClient) listUsagesHandleResponse(resp *http.Response) (DatabaseListUsagesResponse, error) {
	result := DatabaseListUsagesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.UsagesResult); err != nil {
		return DatabaseListUsagesResponse{}, err
	}
	return result, nil
}

// listUsagesHandleError handles the ListUsages error response.
func (client *DatabaseClient) listUsagesHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
