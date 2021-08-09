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
	"time"
)

// ServerTrustGroupsClient contains the methods for the ServerTrustGroups group.
// Don't use this type directly, use NewServerTrustGroupsClient() instead.
type ServerTrustGroupsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewServerTrustGroupsClient creates a new instance of ServerTrustGroupsClient with the specified values.
func NewServerTrustGroupsClient(con *armcore.Connection, subscriptionID string) *ServerTrustGroupsClient {
	return &ServerTrustGroupsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a server trust group.
// If the operation fails it returns a generic error.
func (client *ServerTrustGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, locationName string, serverTrustGroupName string, parameters ServerTrustGroup, options *ServerTrustGroupsBeginCreateOrUpdateOptions) (ServerTrustGroupsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, locationName, serverTrustGroupName, parameters, options)
	if err != nil {
		return ServerTrustGroupsCreateOrUpdatePollerResponse{}, err
	}
	result := ServerTrustGroupsCreateOrUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ServerTrustGroupsClient.CreateOrUpdate", "", resp, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return ServerTrustGroupsCreateOrUpdatePollerResponse{}, err
	}
	poller := &serverTrustGroupsCreateOrUpdatePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ServerTrustGroupsCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new ServerTrustGroupsCreateOrUpdatePoller from the specified resume token.
// token - The value must come from a previous call to ServerTrustGroupsCreateOrUpdatePoller.ResumeToken().
func (client *ServerTrustGroupsClient) ResumeCreateOrUpdate(ctx context.Context, token string) (ServerTrustGroupsCreateOrUpdatePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ServerTrustGroupsClient.CreateOrUpdate", token, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return ServerTrustGroupsCreateOrUpdatePollerResponse{}, err
	}
	poller := &serverTrustGroupsCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ServerTrustGroupsCreateOrUpdatePollerResponse{}, err
	}
	result := ServerTrustGroupsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ServerTrustGroupsCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a server trust group.
// If the operation fails it returns a generic error.
func (client *ServerTrustGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, locationName string, serverTrustGroupName string, parameters ServerTrustGroup, options *ServerTrustGroupsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, locationName, serverTrustGroupName, parameters, options)
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
func (client *ServerTrustGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, locationName string, serverTrustGroupName string, parameters ServerTrustGroup, options *ServerTrustGroupsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/locations/{locationName}/serverTrustGroups/{serverTrustGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if serverTrustGroupName == "" {
		return nil, errors.New("parameter serverTrustGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverTrustGroupName}", url.PathEscape(serverTrustGroupName))
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
func (client *ServerTrustGroupsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginDelete - Deletes a server trust group.
// If the operation fails it returns a generic error.
func (client *ServerTrustGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, locationName string, serverTrustGroupName string, options *ServerTrustGroupsBeginDeleteOptions) (ServerTrustGroupsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, locationName, serverTrustGroupName, options)
	if err != nil {
		return ServerTrustGroupsDeletePollerResponse{}, err
	}
	result := ServerTrustGroupsDeletePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ServerTrustGroupsClient.Delete", "", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return ServerTrustGroupsDeletePollerResponse{}, err
	}
	poller := &serverTrustGroupsDeletePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ServerTrustGroupsDeleteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new ServerTrustGroupsDeletePoller from the specified resume token.
// token - The value must come from a previous call to ServerTrustGroupsDeletePoller.ResumeToken().
func (client *ServerTrustGroupsClient) ResumeDelete(ctx context.Context, token string) (ServerTrustGroupsDeletePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ServerTrustGroupsClient.Delete", token, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return ServerTrustGroupsDeletePollerResponse{}, err
	}
	poller := &serverTrustGroupsDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ServerTrustGroupsDeletePollerResponse{}, err
	}
	result := ServerTrustGroupsDeletePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ServerTrustGroupsDeleteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Deletes a server trust group.
// If the operation fails it returns a generic error.
func (client *ServerTrustGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, locationName string, serverTrustGroupName string, options *ServerTrustGroupsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, locationName, serverTrustGroupName, options)
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
func (client *ServerTrustGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, locationName string, serverTrustGroupName string, options *ServerTrustGroupsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/locations/{locationName}/serverTrustGroups/{serverTrustGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if serverTrustGroupName == "" {
		return nil, errors.New("parameter serverTrustGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverTrustGroupName}", url.PathEscape(serverTrustGroupName))
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
func (client *ServerTrustGroupsClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Get - Gets a server trust group.
// If the operation fails it returns a generic error.
func (client *ServerTrustGroupsClient) Get(ctx context.Context, resourceGroupName string, locationName string, serverTrustGroupName string, options *ServerTrustGroupsGetOptions) (ServerTrustGroupsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, locationName, serverTrustGroupName, options)
	if err != nil {
		return ServerTrustGroupsGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ServerTrustGroupsGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ServerTrustGroupsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServerTrustGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, locationName string, serverTrustGroupName string, options *ServerTrustGroupsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/locations/{locationName}/serverTrustGroups/{serverTrustGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if serverTrustGroupName == "" {
		return nil, errors.New("parameter serverTrustGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverTrustGroupName}", url.PathEscape(serverTrustGroupName))
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
func (client *ServerTrustGroupsClient) getHandleResponse(resp *azcore.Response) (ServerTrustGroupsGetResponse, error) {
	result := ServerTrustGroupsGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ServerTrustGroup); err != nil {
		return ServerTrustGroupsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ServerTrustGroupsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListByInstance - Gets a server trust groups by instance name.
// If the operation fails it returns a generic error.
func (client *ServerTrustGroupsClient) ListByInstance(resourceGroupName string, managedInstanceName string, options *ServerTrustGroupsListByInstanceOptions) ServerTrustGroupsListByInstancePager {
	return &serverTrustGroupsListByInstancePager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByInstanceCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
		},
		advancer: func(ctx context.Context, resp ServerTrustGroupsListByInstanceResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ServerTrustGroupListResult.NextLink)
		},
	}
}

// listByInstanceCreateRequest creates the ListByInstance request.
func (client *ServerTrustGroupsClient) listByInstanceCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ServerTrustGroupsListByInstanceOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/serverTrustGroups"
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
	reqQP.Set("api-version", "2020-11-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByInstanceHandleResponse handles the ListByInstance response.
func (client *ServerTrustGroupsClient) listByInstanceHandleResponse(resp *azcore.Response) (ServerTrustGroupsListByInstanceResponse, error) {
	result := ServerTrustGroupsListByInstanceResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ServerTrustGroupListResult); err != nil {
		return ServerTrustGroupsListByInstanceResponse{}, err
	}
	return result, nil
}

// listByInstanceHandleError handles the ListByInstance error response.
func (client *ServerTrustGroupsClient) listByInstanceHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListByLocation - Lists a server trust group.
// If the operation fails it returns a generic error.
func (client *ServerTrustGroupsClient) ListByLocation(resourceGroupName string, locationName string, options *ServerTrustGroupsListByLocationOptions) ServerTrustGroupsListByLocationPager {
	return &serverTrustGroupsListByLocationPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByLocationCreateRequest(ctx, resourceGroupName, locationName, options)
		},
		advancer: func(ctx context.Context, resp ServerTrustGroupsListByLocationResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ServerTrustGroupListResult.NextLink)
		},
	}
}

// listByLocationCreateRequest creates the ListByLocation request.
func (client *ServerTrustGroupsClient) listByLocationCreateRequest(ctx context.Context, resourceGroupName string, locationName string, options *ServerTrustGroupsListByLocationOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/locations/{locationName}/serverTrustGroups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
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

// listByLocationHandleResponse handles the ListByLocation response.
func (client *ServerTrustGroupsClient) listByLocationHandleResponse(resp *azcore.Response) (ServerTrustGroupsListByLocationResponse, error) {
	result := ServerTrustGroupsListByLocationResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ServerTrustGroupListResult); err != nil {
		return ServerTrustGroupsListByLocationResponse{}, err
	}
	return result, nil
}

// listByLocationHandleError handles the ListByLocation error response.
func (client *ServerTrustGroupsClient) listByLocationHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
