//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ElasticPoolsClient contains the methods for the ElasticPools group.
// Don't use this type directly, use NewElasticPoolsClient() instead.
type ElasticPoolsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewElasticPoolsClient creates a new instance of ElasticPoolsClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewElasticPoolsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ElasticPoolsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &ElasticPoolsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates an elastic pool.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// elasticPoolName - The name of the elastic pool.
// parameters - The elastic pool parameters.
// options - ElasticPoolsClientBeginCreateOrUpdateOptions contains the optional parameters for the ElasticPoolsClient.BeginCreateOrUpdate
// method.
func (client *ElasticPoolsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, parameters ElasticPool, options *ElasticPoolsClientBeginCreateOrUpdateOptions) (ElasticPoolsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, elasticPoolName, parameters, options)
	if err != nil {
		return ElasticPoolsClientCreateOrUpdatePollerResponse{}, err
	}
	result := ElasticPoolsClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ElasticPoolsClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return ElasticPoolsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ElasticPoolsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates an elastic pool.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ElasticPoolsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, parameters ElasticPool, options *ElasticPoolsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, elasticPoolName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ElasticPoolsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, parameters ElasticPool, options *ElasticPoolsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes an elastic pool.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// elasticPoolName - The name of the elastic pool.
// options - ElasticPoolsClientBeginDeleteOptions contains the optional parameters for the ElasticPoolsClient.BeginDelete
// method.
func (client *ElasticPoolsClient) BeginDelete(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, options *ElasticPoolsClientBeginDeleteOptions) (ElasticPoolsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, serverName, elasticPoolName, options)
	if err != nil {
		return ElasticPoolsClientDeletePollerResponse{}, err
	}
	result := ElasticPoolsClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ElasticPoolsClient.Delete", "", resp, client.pl)
	if err != nil {
		return ElasticPoolsClientDeletePollerResponse{}, err
	}
	result.Poller = &ElasticPoolsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes an elastic pool.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ElasticPoolsClient) deleteOperation(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, options *ElasticPoolsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, elasticPoolName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ElasticPoolsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, options *ElasticPoolsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// BeginFailover - Failovers an elastic pool.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// elasticPoolName - The name of the elastic pool to failover.
// options - ElasticPoolsClientBeginFailoverOptions contains the optional parameters for the ElasticPoolsClient.BeginFailover
// method.
func (client *ElasticPoolsClient) BeginFailover(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, options *ElasticPoolsClientBeginFailoverOptions) (ElasticPoolsClientFailoverPollerResponse, error) {
	resp, err := client.failover(ctx, resourceGroupName, serverName, elasticPoolName, options)
	if err != nil {
		return ElasticPoolsClientFailoverPollerResponse{}, err
	}
	result := ElasticPoolsClientFailoverPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ElasticPoolsClient.Failover", "", resp, client.pl)
	if err != nil {
		return ElasticPoolsClientFailoverPollerResponse{}, err
	}
	result.Poller = &ElasticPoolsClientFailoverPoller{
		pt: pt,
	}
	return result, nil
}

// Failover - Failovers an elastic pool.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ElasticPoolsClient) failover(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, options *ElasticPoolsClientBeginFailoverOptions) (*http.Response, error) {
	req, err := client.failoverCreateRequest(ctx, resourceGroupName, serverName, elasticPoolName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// failoverCreateRequest creates the Failover request.
func (client *ElasticPoolsClient) failoverCreateRequest(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, options *ElasticPoolsClientBeginFailoverOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}/failover"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets an elastic pool.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// elasticPoolName - The name of the elastic pool.
// options - ElasticPoolsClientGetOptions contains the optional parameters for the ElasticPoolsClient.Get method.
func (client *ElasticPoolsClient) Get(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, options *ElasticPoolsClientGetOptions) (ElasticPoolsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, elasticPoolName, options)
	if err != nil {
		return ElasticPoolsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ElasticPoolsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ElasticPoolsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ElasticPoolsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, options *ElasticPoolsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ElasticPoolsClient) getHandleResponse(resp *http.Response) (ElasticPoolsClientGetResponse, error) {
	result := ElasticPoolsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ElasticPool); err != nil {
		return ElasticPoolsClientGetResponse{}, err
	}
	return result, nil
}

// ListByServer - Gets all elastic pools in a server.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// options - ElasticPoolsClientListByServerOptions contains the optional parameters for the ElasticPoolsClient.ListByServer
// method.
func (client *ElasticPoolsClient) ListByServer(resourceGroupName string, serverName string, options *ElasticPoolsClientListByServerOptions) *ElasticPoolsClientListByServerPager {
	return &ElasticPoolsClientListByServerPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
		},
		advancer: func(ctx context.Context, resp ElasticPoolsClientListByServerResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ElasticPoolListResult.NextLink)
		},
	}
}

// listByServerCreateRequest creates the ListByServer request.
func (client *ElasticPoolsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ElasticPoolsClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *ElasticPoolsClient) listByServerHandleResponse(resp *http.Response) (ElasticPoolsClientListByServerResponse, error) {
	result := ElasticPoolsClientListByServerResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ElasticPoolListResult); err != nil {
		return ElasticPoolsClientListByServerResponse{}, err
	}
	return result, nil
}

// ListMetricDefinitions - Returns elastic pool metric definitions.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// elasticPoolName - The name of the elastic pool.
// options - ElasticPoolsClientListMetricDefinitionsOptions contains the optional parameters for the ElasticPoolsClient.ListMetricDefinitions
// method.
func (client *ElasticPoolsClient) ListMetricDefinitions(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, options *ElasticPoolsClientListMetricDefinitionsOptions) (ElasticPoolsClientListMetricDefinitionsResponse, error) {
	req, err := client.listMetricDefinitionsCreateRequest(ctx, resourceGroupName, serverName, elasticPoolName, options)
	if err != nil {
		return ElasticPoolsClientListMetricDefinitionsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ElasticPoolsClientListMetricDefinitionsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ElasticPoolsClientListMetricDefinitionsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listMetricDefinitionsHandleResponse(resp)
}

// listMetricDefinitionsCreateRequest creates the ListMetricDefinitions request.
func (client *ElasticPoolsClient) listMetricDefinitionsCreateRequest(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, options *ElasticPoolsClientListMetricDefinitionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}/metricDefinitions"
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
	if elasticPoolName == "" {
		return nil, errors.New("parameter elasticPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{elasticPoolName}", url.PathEscape(elasticPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2014-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listMetricDefinitionsHandleResponse handles the ListMetricDefinitions response.
func (client *ElasticPoolsClient) listMetricDefinitionsHandleResponse(resp *http.Response) (ElasticPoolsClientListMetricDefinitionsResponse, error) {
	result := ElasticPoolsClientListMetricDefinitionsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricDefinitionListResult); err != nil {
		return ElasticPoolsClientListMetricDefinitionsResponse{}, err
	}
	return result, nil
}

// ListMetrics - Returns elastic pool metrics.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// elasticPoolName - The name of the elastic pool.
// filter - An OData filter expression that describes a subset of metrics to return.
// options - ElasticPoolsClientListMetricsOptions contains the optional parameters for the ElasticPoolsClient.ListMetrics
// method.
func (client *ElasticPoolsClient) ListMetrics(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, filter string, options *ElasticPoolsClientListMetricsOptions) (ElasticPoolsClientListMetricsResponse, error) {
	req, err := client.listMetricsCreateRequest(ctx, resourceGroupName, serverName, elasticPoolName, filter, options)
	if err != nil {
		return ElasticPoolsClientListMetricsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ElasticPoolsClientListMetricsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ElasticPoolsClientListMetricsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listMetricsHandleResponse(resp)
}

// listMetricsCreateRequest creates the ListMetrics request.
func (client *ElasticPoolsClient) listMetricsCreateRequest(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, filter string, options *ElasticPoolsClientListMetricsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}/metrics"
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
	if elasticPoolName == "" {
		return nil, errors.New("parameter elasticPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{elasticPoolName}", url.PathEscape(elasticPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2014-04-01")
	reqQP.Set("$filter", filter)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listMetricsHandleResponse handles the ListMetrics response.
func (client *ElasticPoolsClient) listMetricsHandleResponse(resp *http.Response) (ElasticPoolsClientListMetricsResponse, error) {
	result := ElasticPoolsClientListMetricsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricListResult); err != nil {
		return ElasticPoolsClientListMetricsResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates an elastic pool.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// elasticPoolName - The name of the elastic pool.
// parameters - The elastic pool update parameters.
// options - ElasticPoolsClientBeginUpdateOptions contains the optional parameters for the ElasticPoolsClient.BeginUpdate
// method.
func (client *ElasticPoolsClient) BeginUpdate(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, parameters ElasticPoolUpdate, options *ElasticPoolsClientBeginUpdateOptions) (ElasticPoolsClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, serverName, elasticPoolName, parameters, options)
	if err != nil {
		return ElasticPoolsClientUpdatePollerResponse{}, err
	}
	result := ElasticPoolsClientUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ElasticPoolsClient.Update", "", resp, client.pl)
	if err != nil {
		return ElasticPoolsClientUpdatePollerResponse{}, err
	}
	result.Poller = &ElasticPoolsClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates an elastic pool.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ElasticPoolsClient) update(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, parameters ElasticPoolUpdate, options *ElasticPoolsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serverName, elasticPoolName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ElasticPoolsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, parameters ElasticPoolUpdate, options *ElasticPoolsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}
