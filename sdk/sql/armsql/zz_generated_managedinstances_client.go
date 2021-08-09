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
	"strconv"
	"strings"
	"time"
)

// ManagedInstancesClient contains the methods for the ManagedInstances group.
// Don't use this type directly, use NewManagedInstancesClient() instead.
type ManagedInstancesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewManagedInstancesClient creates a new instance of ManagedInstancesClient with the specified values.
func NewManagedInstancesClient(con *armcore.Connection, subscriptionID string) *ManagedInstancesClient {
	return &ManagedInstancesClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a managed instance.
// If the operation fails it returns a generic error.
func (client *ManagedInstancesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters ManagedInstance, options *ManagedInstancesBeginCreateOrUpdateOptions) (ManagedInstancesCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, managedInstanceName, parameters, options)
	if err != nil {
		return ManagedInstancesCreateOrUpdatePollerResponse{}, err
	}
	result := ManagedInstancesCreateOrUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ManagedInstancesClient.CreateOrUpdate", "", resp, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return ManagedInstancesCreateOrUpdatePollerResponse{}, err
	}
	poller := &managedInstancesCreateOrUpdatePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ManagedInstancesCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new ManagedInstancesCreateOrUpdatePoller from the specified resume token.
// token - The value must come from a previous call to ManagedInstancesCreateOrUpdatePoller.ResumeToken().
func (client *ManagedInstancesClient) ResumeCreateOrUpdate(ctx context.Context, token string) (ManagedInstancesCreateOrUpdatePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ManagedInstancesClient.CreateOrUpdate", token, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return ManagedInstancesCreateOrUpdatePollerResponse{}, err
	}
	poller := &managedInstancesCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ManagedInstancesCreateOrUpdatePollerResponse{}, err
	}
	result := ManagedInstancesCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ManagedInstancesCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a managed instance.
// If the operation fails it returns a generic error.
func (client *ManagedInstancesClient) createOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters ManagedInstance, options *ManagedInstancesBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, managedInstanceName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ManagedInstancesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters ManagedInstance, options *ManagedInstancesBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ManagedInstancesClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginDelete - Deletes a managed instance.
// If the operation fails it returns a generic error.
func (client *ManagedInstancesClient) BeginDelete(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstancesBeginDeleteOptions) (ManagedInstancesDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, managedInstanceName, options)
	if err != nil {
		return ManagedInstancesDeletePollerResponse{}, err
	}
	result := ManagedInstancesDeletePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ManagedInstancesClient.Delete", "", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return ManagedInstancesDeletePollerResponse{}, err
	}
	poller := &managedInstancesDeletePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ManagedInstancesDeleteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new ManagedInstancesDeletePoller from the specified resume token.
// token - The value must come from a previous call to ManagedInstancesDeletePoller.ResumeToken().
func (client *ManagedInstancesClient) ResumeDelete(ctx context.Context, token string) (ManagedInstancesDeletePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ManagedInstancesClient.Delete", token, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return ManagedInstancesDeletePollerResponse{}, err
	}
	poller := &managedInstancesDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ManagedInstancesDeletePollerResponse{}, err
	}
	result := ManagedInstancesDeletePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ManagedInstancesDeleteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Deletes a managed instance.
// If the operation fails it returns a generic error.
func (client *ManagedInstancesClient) deleteOperation(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstancesBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ManagedInstancesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstancesBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ManagedInstancesClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginFailover - Failovers a managed instance.
// If the operation fails it returns a generic error.
func (client *ManagedInstancesClient) BeginFailover(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstancesBeginFailoverOptions) (ManagedInstancesFailoverPollerResponse, error) {
	resp, err := client.failover(ctx, resourceGroupName, managedInstanceName, options)
	if err != nil {
		return ManagedInstancesFailoverPollerResponse{}, err
	}
	result := ManagedInstancesFailoverPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ManagedInstancesClient.Failover", "", resp, client.con.Pipeline(), client.failoverHandleError)
	if err != nil {
		return ManagedInstancesFailoverPollerResponse{}, err
	}
	poller := &managedInstancesFailoverPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ManagedInstancesFailoverResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeFailover creates a new ManagedInstancesFailoverPoller from the specified resume token.
// token - The value must come from a previous call to ManagedInstancesFailoverPoller.ResumeToken().
func (client *ManagedInstancesClient) ResumeFailover(ctx context.Context, token string) (ManagedInstancesFailoverPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ManagedInstancesClient.Failover", token, client.con.Pipeline(), client.failoverHandleError)
	if err != nil {
		return ManagedInstancesFailoverPollerResponse{}, err
	}
	poller := &managedInstancesFailoverPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ManagedInstancesFailoverPollerResponse{}, err
	}
	result := ManagedInstancesFailoverPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ManagedInstancesFailoverResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Failover - Failovers a managed instance.
// If the operation fails it returns a generic error.
func (client *ManagedInstancesClient) failover(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstancesBeginFailoverOptions) (*azcore.Response, error) {
	req, err := client.failoverCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.failoverHandleError(resp)
	}
	return resp, nil
}

// failoverCreateRequest creates the Failover request.
func (client *ManagedInstancesClient) failoverCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstancesBeginFailoverOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/failover"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.ReplicaType != nil {
		reqQP.Set("replicaType", string(*options.ReplicaType))
	}
	reqQP.Set("api-version", "2020-11-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	return req, nil
}

// failoverHandleError handles the Failover error response.
func (client *ManagedInstancesClient) failoverHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Get - Gets a managed instance.
// If the operation fails it returns a generic error.
func (client *ManagedInstancesClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstancesGetOptions) (ManagedInstancesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
	if err != nil {
		return ManagedInstancesGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ManagedInstancesGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ManagedInstancesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ManagedInstancesClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstancesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
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
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2020-11-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedInstancesClient) getHandleResponse(resp *azcore.Response) (ManagedInstancesGetResponse, error) {
	result := ManagedInstancesGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ManagedInstance); err != nil {
		return ManagedInstancesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ManagedInstancesClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// List - Gets a list of all managed instances in the subscription.
// If the operation fails it returns a generic error.
func (client *ManagedInstancesClient) List(options *ManagedInstancesListOptions) ManagedInstancesListPager {
	return &managedInstancesListPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ManagedInstancesListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ManagedInstanceListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ManagedInstancesClient) listCreateRequest(ctx context.Context, options *ManagedInstancesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/managedInstances"
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
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2020-11-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ManagedInstancesClient) listHandleResponse(resp *azcore.Response) (ManagedInstancesListResponse, error) {
	result := ManagedInstancesListResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ManagedInstanceListResult); err != nil {
		return ManagedInstancesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ManagedInstancesClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListByInstancePool - Gets a list of all managed instances in an instance pool.
// If the operation fails it returns a generic error.
func (client *ManagedInstancesClient) ListByInstancePool(resourceGroupName string, instancePoolName string, options *ManagedInstancesListByInstancePoolOptions) ManagedInstancesListByInstancePoolPager {
	return &managedInstancesListByInstancePoolPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByInstancePoolCreateRequest(ctx, resourceGroupName, instancePoolName, options)
		},
		advancer: func(ctx context.Context, resp ManagedInstancesListByInstancePoolResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ManagedInstanceListResult.NextLink)
		},
	}
}

// listByInstancePoolCreateRequest creates the ListByInstancePool request.
func (client *ManagedInstancesClient) listByInstancePoolCreateRequest(ctx context.Context, resourceGroupName string, instancePoolName string, options *ManagedInstancesListByInstancePoolOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/instancePools/{instancePoolName}/managedInstances"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if instancePoolName == "" {
		return nil, errors.New("parameter instancePoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instancePoolName}", url.PathEscape(instancePoolName))
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
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2020-11-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByInstancePoolHandleResponse handles the ListByInstancePool response.
func (client *ManagedInstancesClient) listByInstancePoolHandleResponse(resp *azcore.Response) (ManagedInstancesListByInstancePoolResponse, error) {
	result := ManagedInstancesListByInstancePoolResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ManagedInstanceListResult); err != nil {
		return ManagedInstancesListByInstancePoolResponse{}, err
	}
	return result, nil
}

// listByInstancePoolHandleError handles the ListByInstancePool error response.
func (client *ManagedInstancesClient) listByInstancePoolHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListByManagedInstance - Get top resource consuming queries of a managed instance.
// If the operation fails it returns a generic error.
func (client *ManagedInstancesClient) ListByManagedInstance(resourceGroupName string, managedInstanceName string, options *ManagedInstancesListByManagedInstanceOptions) ManagedInstancesListByManagedInstancePager {
	return &managedInstancesListByManagedInstancePager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByManagedInstanceCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
		},
		advancer: func(ctx context.Context, resp ManagedInstancesListByManagedInstanceResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.TopQueriesListResult.NextLink)
		},
	}
}

// listByManagedInstanceCreateRequest creates the ListByManagedInstance request.
func (client *ManagedInstancesClient) listByManagedInstanceCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstancesListByManagedInstanceOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/topqueries"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
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
	if options != nil && options.NumberOfQueries != nil {
		reqQP.Set("numberOfQueries", strconv.FormatInt(int64(*options.NumberOfQueries), 10))
	}
	if options != nil && options.Databases != nil {
		reqQP.Set("databases", *options.Databases)
	}
	if options != nil && options.StartTime != nil {
		reqQP.Set("startTime", *options.StartTime)
	}
	if options != nil && options.EndTime != nil {
		reqQP.Set("endTime", *options.EndTime)
	}
	if options != nil && options.Interval != nil {
		reqQP.Set("interval", string(*options.Interval))
	}
	if options != nil && options.AggregationFunction != nil {
		reqQP.Set("aggregationFunction", string(*options.AggregationFunction))
	}
	if options != nil && options.ObservationMetric != nil {
		reqQP.Set("observationMetric", string(*options.ObservationMetric))
	}
	reqQP.Set("api-version", "2020-11-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByManagedInstanceHandleResponse handles the ListByManagedInstance response.
func (client *ManagedInstancesClient) listByManagedInstanceHandleResponse(resp *azcore.Response) (ManagedInstancesListByManagedInstanceResponse, error) {
	result := ManagedInstancesListByManagedInstanceResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.TopQueriesListResult); err != nil {
		return ManagedInstancesListByManagedInstanceResponse{}, err
	}
	return result, nil
}

// listByManagedInstanceHandleError handles the ListByManagedInstance error response.
func (client *ManagedInstancesClient) listByManagedInstanceHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListByResourceGroup - Gets a list of managed instances in a resource group.
// If the operation fails it returns a generic error.
func (client *ManagedInstancesClient) ListByResourceGroup(resourceGroupName string, options *ManagedInstancesListByResourceGroupOptions) ManagedInstancesListByResourceGroupPager {
	return &managedInstancesListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp ManagedInstancesListByResourceGroupResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ManagedInstanceListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ManagedInstancesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ManagedInstancesListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2020-11-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ManagedInstancesClient) listByResourceGroupHandleResponse(resp *azcore.Response) (ManagedInstancesListByResourceGroupResponse, error) {
	result := ManagedInstancesListByResourceGroupResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ManagedInstanceListResult); err != nil {
		return ManagedInstancesListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *ManagedInstancesClient) listByResourceGroupHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginUpdate - Updates a managed instance.
// If the operation fails it returns a generic error.
func (client *ManagedInstancesClient) BeginUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters ManagedInstanceUpdate, options *ManagedInstancesBeginUpdateOptions) (ManagedInstancesUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, managedInstanceName, parameters, options)
	if err != nil {
		return ManagedInstancesUpdatePollerResponse{}, err
	}
	result := ManagedInstancesUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ManagedInstancesClient.Update", "", resp, client.con.Pipeline(), client.updateHandleError)
	if err != nil {
		return ManagedInstancesUpdatePollerResponse{}, err
	}
	poller := &managedInstancesUpdatePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ManagedInstancesUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeUpdate creates a new ManagedInstancesUpdatePoller from the specified resume token.
// token - The value must come from a previous call to ManagedInstancesUpdatePoller.ResumeToken().
func (client *ManagedInstancesClient) ResumeUpdate(ctx context.Context, token string) (ManagedInstancesUpdatePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ManagedInstancesClient.Update", token, client.con.Pipeline(), client.updateHandleError)
	if err != nil {
		return ManagedInstancesUpdatePollerResponse{}, err
	}
	poller := &managedInstancesUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ManagedInstancesUpdatePollerResponse{}, err
	}
	result := ManagedInstancesUpdatePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ManagedInstancesUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Update - Updates a managed instance.
// If the operation fails it returns a generic error.
func (client *ManagedInstancesClient) update(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters ManagedInstanceUpdate, options *ManagedInstancesBeginUpdateOptions) (*azcore.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, managedInstanceName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ManagedInstancesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters ManagedInstanceUpdate, options *ManagedInstancesBeginUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateHandleError handles the Update error response.
func (client *ManagedInstancesClient) updateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
