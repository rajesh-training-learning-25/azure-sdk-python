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

// ServerAzureADOnlyAuthenticationsClient contains the methods for the ServerAzureADOnlyAuthentications group.
// Don't use this type directly, use NewServerAzureADOnlyAuthenticationsClient() instead.
type ServerAzureADOnlyAuthenticationsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewServerAzureADOnlyAuthenticationsClient creates a new instance of ServerAzureADOnlyAuthenticationsClient with the specified values.
func NewServerAzureADOnlyAuthenticationsClient(con *armcore.Connection, subscriptionID string) *ServerAzureADOnlyAuthenticationsClient {
	return &ServerAzureADOnlyAuthenticationsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Sets Server Active Directory only authentication property or updates an existing server Active Directory only authentication property.
// If the operation fails it returns a generic error.
func (client *ServerAzureADOnlyAuthenticationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, authenticationName AuthenticationName, parameters ServerAzureADOnlyAuthentication, options *ServerAzureADOnlyAuthenticationsBeginCreateOrUpdateOptions) (ServerAzureADOnlyAuthenticationsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, authenticationName, parameters, options)
	if err != nil {
		return ServerAzureADOnlyAuthenticationsCreateOrUpdatePollerResponse{}, err
	}
	result := ServerAzureADOnlyAuthenticationsCreateOrUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ServerAzureADOnlyAuthenticationsClient.CreateOrUpdate", "", resp, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return ServerAzureADOnlyAuthenticationsCreateOrUpdatePollerResponse{}, err
	}
	poller := &serverAzureADOnlyAuthenticationsCreateOrUpdatePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ServerAzureADOnlyAuthenticationsCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new ServerAzureADOnlyAuthenticationsCreateOrUpdatePoller from the specified resume token.
// token - The value must come from a previous call to ServerAzureADOnlyAuthenticationsCreateOrUpdatePoller.ResumeToken().
func (client *ServerAzureADOnlyAuthenticationsClient) ResumeCreateOrUpdate(ctx context.Context, token string) (ServerAzureADOnlyAuthenticationsCreateOrUpdatePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ServerAzureADOnlyAuthenticationsClient.CreateOrUpdate", token, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return ServerAzureADOnlyAuthenticationsCreateOrUpdatePollerResponse{}, err
	}
	poller := &serverAzureADOnlyAuthenticationsCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ServerAzureADOnlyAuthenticationsCreateOrUpdatePollerResponse{}, err
	}
	result := ServerAzureADOnlyAuthenticationsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ServerAzureADOnlyAuthenticationsCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Sets Server Active Directory only authentication property or updates an existing server Active Directory only authentication property.
// If the operation fails it returns a generic error.
func (client *ServerAzureADOnlyAuthenticationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, authenticationName AuthenticationName, parameters ServerAzureADOnlyAuthentication, options *ServerAzureADOnlyAuthenticationsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, authenticationName, parameters, options)
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
func (client *ServerAzureADOnlyAuthenticationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, authenticationName AuthenticationName, parameters ServerAzureADOnlyAuthentication, options *ServerAzureADOnlyAuthenticationsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/azureADOnlyAuthentications/{authenticationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if authenticationName == "" {
		return nil, errors.New("parameter authenticationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authenticationName}", url.PathEscape(string(authenticationName)))
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
func (client *ServerAzureADOnlyAuthenticationsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginDelete - Deletes an existing server Active Directory only authentication property.
// If the operation fails it returns a generic error.
func (client *ServerAzureADOnlyAuthenticationsClient) BeginDelete(ctx context.Context, resourceGroupName string, serverName string, authenticationName AuthenticationName, options *ServerAzureADOnlyAuthenticationsBeginDeleteOptions) (ServerAzureADOnlyAuthenticationsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, serverName, authenticationName, options)
	if err != nil {
		return ServerAzureADOnlyAuthenticationsDeletePollerResponse{}, err
	}
	result := ServerAzureADOnlyAuthenticationsDeletePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ServerAzureADOnlyAuthenticationsClient.Delete", "", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return ServerAzureADOnlyAuthenticationsDeletePollerResponse{}, err
	}
	poller := &serverAzureADOnlyAuthenticationsDeletePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ServerAzureADOnlyAuthenticationsDeleteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new ServerAzureADOnlyAuthenticationsDeletePoller from the specified resume token.
// token - The value must come from a previous call to ServerAzureADOnlyAuthenticationsDeletePoller.ResumeToken().
func (client *ServerAzureADOnlyAuthenticationsClient) ResumeDelete(ctx context.Context, token string) (ServerAzureADOnlyAuthenticationsDeletePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ServerAzureADOnlyAuthenticationsClient.Delete", token, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return ServerAzureADOnlyAuthenticationsDeletePollerResponse{}, err
	}
	poller := &serverAzureADOnlyAuthenticationsDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ServerAzureADOnlyAuthenticationsDeletePollerResponse{}, err
	}
	result := ServerAzureADOnlyAuthenticationsDeletePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ServerAzureADOnlyAuthenticationsDeleteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Deletes an existing server Active Directory only authentication property.
// If the operation fails it returns a generic error.
func (client *ServerAzureADOnlyAuthenticationsClient) deleteOperation(ctx context.Context, resourceGroupName string, serverName string, authenticationName AuthenticationName, options *ServerAzureADOnlyAuthenticationsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, authenticationName, options)
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
func (client *ServerAzureADOnlyAuthenticationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, authenticationName AuthenticationName, options *ServerAzureADOnlyAuthenticationsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/azureADOnlyAuthentications/{authenticationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if authenticationName == "" {
		return nil, errors.New("parameter authenticationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authenticationName}", url.PathEscape(string(authenticationName)))
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
func (client *ServerAzureADOnlyAuthenticationsClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Get - Gets a specific Azure Active Directory only authentication property.
// If the operation fails it returns a generic error.
func (client *ServerAzureADOnlyAuthenticationsClient) Get(ctx context.Context, resourceGroupName string, serverName string, authenticationName AuthenticationName, options *ServerAzureADOnlyAuthenticationsGetOptions) (ServerAzureADOnlyAuthenticationsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, authenticationName, options)
	if err != nil {
		return ServerAzureADOnlyAuthenticationsGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ServerAzureADOnlyAuthenticationsGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ServerAzureADOnlyAuthenticationsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServerAzureADOnlyAuthenticationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, authenticationName AuthenticationName, options *ServerAzureADOnlyAuthenticationsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/azureADOnlyAuthentications/{authenticationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if authenticationName == "" {
		return nil, errors.New("parameter authenticationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authenticationName}", url.PathEscape(string(authenticationName)))
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
func (client *ServerAzureADOnlyAuthenticationsClient) getHandleResponse(resp *azcore.Response) (ServerAzureADOnlyAuthenticationsGetResponse, error) {
	result := ServerAzureADOnlyAuthenticationsGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ServerAzureADOnlyAuthentication); err != nil {
		return ServerAzureADOnlyAuthenticationsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ServerAzureADOnlyAuthenticationsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListByServer - Gets a list of server Azure Active Directory only authentications.
// If the operation fails it returns a generic error.
func (client *ServerAzureADOnlyAuthenticationsClient) ListByServer(resourceGroupName string, serverName string, options *ServerAzureADOnlyAuthenticationsListByServerOptions) ServerAzureADOnlyAuthenticationsListByServerPager {
	return &serverAzureADOnlyAuthenticationsListByServerPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
		},
		advancer: func(ctx context.Context, resp ServerAzureADOnlyAuthenticationsListByServerResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.AzureADOnlyAuthListResult.NextLink)
		},
	}
}

// listByServerCreateRequest creates the ListByServer request.
func (client *ServerAzureADOnlyAuthenticationsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerAzureADOnlyAuthenticationsListByServerOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/azureADOnlyAuthentications"
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

// listByServerHandleResponse handles the ListByServer response.
func (client *ServerAzureADOnlyAuthenticationsClient) listByServerHandleResponse(resp *azcore.Response) (ServerAzureADOnlyAuthenticationsListByServerResponse, error) {
	result := ServerAzureADOnlyAuthenticationsListByServerResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.AzureADOnlyAuthListResult); err != nil {
		return ServerAzureADOnlyAuthenticationsListByServerResponse{}, err
	}
	return result, nil
}

// listByServerHandleError handles the ListByServer error response.
func (client *ServerAzureADOnlyAuthenticationsClient) listByServerHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
