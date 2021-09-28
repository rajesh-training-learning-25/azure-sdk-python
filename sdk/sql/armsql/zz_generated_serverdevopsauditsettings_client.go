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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ServerDevOpsAuditSettingsClient contains the methods for the ServerDevOpsAuditSettings group.
// Don't use this type directly, use NewServerDevOpsAuditSettingsClient() instead.
type ServerDevOpsAuditSettingsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewServerDevOpsAuditSettingsClient creates a new instance of ServerDevOpsAuditSettingsClient with the specified values.
func NewServerDevOpsAuditSettingsClient(con *arm.Connection, subscriptionID string) *ServerDevOpsAuditSettingsClient {
	return &ServerDevOpsAuditSettingsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a server's DevOps audit settings.
// If the operation fails it returns a generic error.
func (client *ServerDevOpsAuditSettingsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, devOpsAuditingSettingsName string, parameters ServerDevOpsAuditingSettings, options *ServerDevOpsAuditSettingsBeginCreateOrUpdateOptions) (ServerDevOpsAuditSettingsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, devOpsAuditingSettingsName, parameters, options)
	if err != nil {
		return ServerDevOpsAuditSettingsCreateOrUpdatePollerResponse{}, err
	}
	result := ServerDevOpsAuditSettingsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ServerDevOpsAuditSettingsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return ServerDevOpsAuditSettingsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ServerDevOpsAuditSettingsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a server's DevOps audit settings.
// If the operation fails it returns a generic error.
func (client *ServerDevOpsAuditSettingsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, devOpsAuditingSettingsName string, parameters ServerDevOpsAuditingSettings, options *ServerDevOpsAuditSettingsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, devOpsAuditingSettingsName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServerDevOpsAuditSettingsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, devOpsAuditingSettingsName string, parameters ServerDevOpsAuditingSettings, options *ServerDevOpsAuditSettingsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/devOpsAuditingSettings/{devOpsAuditingSettingsName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if devOpsAuditingSettingsName == "" {
		return nil, errors.New("parameter devOpsAuditingSettingsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devOpsAuditingSettingsName}", url.PathEscape(devOpsAuditingSettingsName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ServerDevOpsAuditSettingsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Gets a server's DevOps audit settings.
// If the operation fails it returns a generic error.
func (client *ServerDevOpsAuditSettingsClient) Get(ctx context.Context, resourceGroupName string, serverName string, devOpsAuditingSettingsName string, options *ServerDevOpsAuditSettingsGetOptions) (ServerDevOpsAuditSettingsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, devOpsAuditingSettingsName, options)
	if err != nil {
		return ServerDevOpsAuditSettingsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServerDevOpsAuditSettingsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServerDevOpsAuditSettingsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServerDevOpsAuditSettingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, devOpsAuditingSettingsName string, options *ServerDevOpsAuditSettingsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/devOpsAuditingSettings/{devOpsAuditingSettingsName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if devOpsAuditingSettingsName == "" {
		return nil, errors.New("parameter devOpsAuditingSettingsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devOpsAuditingSettingsName}", url.PathEscape(devOpsAuditingSettingsName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
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
func (client *ServerDevOpsAuditSettingsClient) getHandleResponse(resp *http.Response) (ServerDevOpsAuditSettingsGetResponse, error) {
	result := ServerDevOpsAuditSettingsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerDevOpsAuditingSettings); err != nil {
		return ServerDevOpsAuditSettingsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ServerDevOpsAuditSettingsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByServer - Lists DevOps audit settings of a server.
// If the operation fails it returns a generic error.
func (client *ServerDevOpsAuditSettingsClient) ListByServer(resourceGroupName string, serverName string, options *ServerDevOpsAuditSettingsListByServerOptions) *ServerDevOpsAuditSettingsListByServerPager {
	return &ServerDevOpsAuditSettingsListByServerPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
		},
		advancer: func(ctx context.Context, resp ServerDevOpsAuditSettingsListByServerResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ServerDevOpsAuditSettingsListResult.NextLink)
		},
	}
}

// listByServerCreateRequest creates the ListByServer request.
func (client *ServerDevOpsAuditSettingsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerDevOpsAuditSettingsListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/devOpsAuditingSettings"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *ServerDevOpsAuditSettingsClient) listByServerHandleResponse(resp *http.Response) (ServerDevOpsAuditSettingsListByServerResponse, error) {
	result := ServerDevOpsAuditSettingsListByServerResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerDevOpsAuditSettingsListResult); err != nil {
		return ServerDevOpsAuditSettingsListByServerResponse{}, err
	}
	return result, nil
}

// listByServerHandleError handles the ListByServer error response.
func (client *ServerDevOpsAuditSettingsClient) listByServerHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
