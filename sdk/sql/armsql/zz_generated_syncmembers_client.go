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

// SyncMembersClient contains the methods for the SyncMembers group.
// Don't use this type directly, use NewSyncMembersClient() instead.
type SyncMembersClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewSyncMembersClient creates a new instance of SyncMembersClient with the specified values.
func NewSyncMembersClient(con *armcore.Connection, subscriptionID string) *SyncMembersClient {
	return &SyncMembersClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a sync member.
// If the operation fails it returns a generic error.
func (client *SyncMembersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string, parameters SyncMember, options *SyncMembersBeginCreateOrUpdateOptions) (SyncMembersCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, databaseName, syncGroupName, syncMemberName, parameters, options)
	if err != nil {
		return SyncMembersCreateOrUpdatePollerResponse{}, err
	}
	result := SyncMembersCreateOrUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("SyncMembersClient.CreateOrUpdate", "", resp, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return SyncMembersCreateOrUpdatePollerResponse{}, err
	}
	poller := &syncMembersCreateOrUpdatePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SyncMembersCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new SyncMembersCreateOrUpdatePoller from the specified resume token.
// token - The value must come from a previous call to SyncMembersCreateOrUpdatePoller.ResumeToken().
func (client *SyncMembersClient) ResumeCreateOrUpdate(ctx context.Context, token string) (SyncMembersCreateOrUpdatePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("SyncMembersClient.CreateOrUpdate", token, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return SyncMembersCreateOrUpdatePollerResponse{}, err
	}
	poller := &syncMembersCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return SyncMembersCreateOrUpdatePollerResponse{}, err
	}
	result := SyncMembersCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SyncMembersCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a sync member.
// If the operation fails it returns a generic error.
func (client *SyncMembersClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string, parameters SyncMember, options *SyncMembersBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, databaseName, syncGroupName, syncMemberName, parameters, options)
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
func (client *SyncMembersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string, parameters SyncMember, options *SyncMembersBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/syncGroups/{syncGroupName}/syncMembers/{syncMemberName}"
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
	if syncGroupName == "" {
		return nil, errors.New("parameter syncGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncGroupName}", url.PathEscape(syncGroupName))
	if syncMemberName == "" {
		return nil, errors.New("parameter syncMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncMemberName}", url.PathEscape(syncMemberName))
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
func (client *SyncMembersClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginDelete - Deletes a sync member.
// If the operation fails it returns a generic error.
func (client *SyncMembersClient) BeginDelete(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string, options *SyncMembersBeginDeleteOptions) (SyncMembersDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, serverName, databaseName, syncGroupName, syncMemberName, options)
	if err != nil {
		return SyncMembersDeletePollerResponse{}, err
	}
	result := SyncMembersDeletePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("SyncMembersClient.Delete", "", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return SyncMembersDeletePollerResponse{}, err
	}
	poller := &syncMembersDeletePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SyncMembersDeleteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new SyncMembersDeletePoller from the specified resume token.
// token - The value must come from a previous call to SyncMembersDeletePoller.ResumeToken().
func (client *SyncMembersClient) ResumeDelete(ctx context.Context, token string) (SyncMembersDeletePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("SyncMembersClient.Delete", token, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return SyncMembersDeletePollerResponse{}, err
	}
	poller := &syncMembersDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return SyncMembersDeletePollerResponse{}, err
	}
	result := SyncMembersDeletePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SyncMembersDeleteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Deletes a sync member.
// If the operation fails it returns a generic error.
func (client *SyncMembersClient) deleteOperation(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string, options *SyncMembersBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, databaseName, syncGroupName, syncMemberName, options)
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
func (client *SyncMembersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string, options *SyncMembersBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/syncGroups/{syncGroupName}/syncMembers/{syncMemberName}"
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
	if syncGroupName == "" {
		return nil, errors.New("parameter syncGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncGroupName}", url.PathEscape(syncGroupName))
	if syncMemberName == "" {
		return nil, errors.New("parameter syncMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncMemberName}", url.PathEscape(syncMemberName))
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
func (client *SyncMembersClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Get - Gets a sync member.
// If the operation fails it returns a generic error.
func (client *SyncMembersClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string, options *SyncMembersGetOptions) (SyncMembersGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, databaseName, syncGroupName, syncMemberName, options)
	if err != nil {
		return SyncMembersGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SyncMembersGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SyncMembersGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SyncMembersClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string, options *SyncMembersGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/syncGroups/{syncGroupName}/syncMembers/{syncMemberName}"
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
	if syncGroupName == "" {
		return nil, errors.New("parameter syncGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncGroupName}", url.PathEscape(syncGroupName))
	if syncMemberName == "" {
		return nil, errors.New("parameter syncMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncMemberName}", url.PathEscape(syncMemberName))
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
func (client *SyncMembersClient) getHandleResponse(resp *azcore.Response) (SyncMembersGetResponse, error) {
	result := SyncMembersGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.SyncMember); err != nil {
		return SyncMembersGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *SyncMembersClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListBySyncGroup - Lists sync members in the given sync group.
// If the operation fails it returns a generic error.
func (client *SyncMembersClient) ListBySyncGroup(resourceGroupName string, serverName string, databaseName string, syncGroupName string, options *SyncMembersListBySyncGroupOptions) SyncMembersListBySyncGroupPager {
	return &syncMembersListBySyncGroupPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listBySyncGroupCreateRequest(ctx, resourceGroupName, serverName, databaseName, syncGroupName, options)
		},
		advancer: func(ctx context.Context, resp SyncMembersListBySyncGroupResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.SyncMemberListResult.NextLink)
		},
	}
}

// listBySyncGroupCreateRequest creates the ListBySyncGroup request.
func (client *SyncMembersClient) listBySyncGroupCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, options *SyncMembersListBySyncGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/syncGroups/{syncGroupName}/syncMembers"
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
	if syncGroupName == "" {
		return nil, errors.New("parameter syncGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncGroupName}", url.PathEscape(syncGroupName))
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

// listBySyncGroupHandleResponse handles the ListBySyncGroup response.
func (client *SyncMembersClient) listBySyncGroupHandleResponse(resp *azcore.Response) (SyncMembersListBySyncGroupResponse, error) {
	result := SyncMembersListBySyncGroupResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.SyncMemberListResult); err != nil {
		return SyncMembersListBySyncGroupResponse{}, err
	}
	return result, nil
}

// listBySyncGroupHandleError handles the ListBySyncGroup error response.
func (client *SyncMembersClient) listBySyncGroupHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListMemberSchemas - Gets a sync member database schema.
// If the operation fails it returns a generic error.
func (client *SyncMembersClient) ListMemberSchemas(resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string, options *SyncMembersListMemberSchemasOptions) SyncMembersListMemberSchemasPager {
	return &syncMembersListMemberSchemasPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listMemberSchemasCreateRequest(ctx, resourceGroupName, serverName, databaseName, syncGroupName, syncMemberName, options)
		},
		advancer: func(ctx context.Context, resp SyncMembersListMemberSchemasResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.SyncFullSchemaPropertiesListResult.NextLink)
		},
	}
}

// listMemberSchemasCreateRequest creates the ListMemberSchemas request.
func (client *SyncMembersClient) listMemberSchemasCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string, options *SyncMembersListMemberSchemasOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/syncGroups/{syncGroupName}/syncMembers/{syncMemberName}/schemas"
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
	if syncGroupName == "" {
		return nil, errors.New("parameter syncGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncGroupName}", url.PathEscape(syncGroupName))
	if syncMemberName == "" {
		return nil, errors.New("parameter syncMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncMemberName}", url.PathEscape(syncMemberName))
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

// listMemberSchemasHandleResponse handles the ListMemberSchemas response.
func (client *SyncMembersClient) listMemberSchemasHandleResponse(resp *azcore.Response) (SyncMembersListMemberSchemasResponse, error) {
	result := SyncMembersListMemberSchemasResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.SyncFullSchemaPropertiesListResult); err != nil {
		return SyncMembersListMemberSchemasResponse{}, err
	}
	return result, nil
}

// listMemberSchemasHandleError handles the ListMemberSchemas error response.
func (client *SyncMembersClient) listMemberSchemasHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginRefreshMemberSchema - Refreshes a sync member database schema.
// If the operation fails it returns a generic error.
func (client *SyncMembersClient) BeginRefreshMemberSchema(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string, options *SyncMembersBeginRefreshMemberSchemaOptions) (SyncMembersRefreshMemberSchemaPollerResponse, error) {
	resp, err := client.refreshMemberSchema(ctx, resourceGroupName, serverName, databaseName, syncGroupName, syncMemberName, options)
	if err != nil {
		return SyncMembersRefreshMemberSchemaPollerResponse{}, err
	}
	result := SyncMembersRefreshMemberSchemaPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("SyncMembersClient.RefreshMemberSchema", "", resp, client.con.Pipeline(), client.refreshMemberSchemaHandleError)
	if err != nil {
		return SyncMembersRefreshMemberSchemaPollerResponse{}, err
	}
	poller := &syncMembersRefreshMemberSchemaPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SyncMembersRefreshMemberSchemaResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeRefreshMemberSchema creates a new SyncMembersRefreshMemberSchemaPoller from the specified resume token.
// token - The value must come from a previous call to SyncMembersRefreshMemberSchemaPoller.ResumeToken().
func (client *SyncMembersClient) ResumeRefreshMemberSchema(ctx context.Context, token string) (SyncMembersRefreshMemberSchemaPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("SyncMembersClient.RefreshMemberSchema", token, client.con.Pipeline(), client.refreshMemberSchemaHandleError)
	if err != nil {
		return SyncMembersRefreshMemberSchemaPollerResponse{}, err
	}
	poller := &syncMembersRefreshMemberSchemaPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return SyncMembersRefreshMemberSchemaPollerResponse{}, err
	}
	result := SyncMembersRefreshMemberSchemaPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SyncMembersRefreshMemberSchemaResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// RefreshMemberSchema - Refreshes a sync member database schema.
// If the operation fails it returns a generic error.
func (client *SyncMembersClient) refreshMemberSchema(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string, options *SyncMembersBeginRefreshMemberSchemaOptions) (*azcore.Response, error) {
	req, err := client.refreshMemberSchemaCreateRequest(ctx, resourceGroupName, serverName, databaseName, syncGroupName, syncMemberName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.refreshMemberSchemaHandleError(resp)
	}
	return resp, nil
}

// refreshMemberSchemaCreateRequest creates the RefreshMemberSchema request.
func (client *SyncMembersClient) refreshMemberSchemaCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string, options *SyncMembersBeginRefreshMemberSchemaOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/syncGroups/{syncGroupName}/syncMembers/{syncMemberName}/refreshSchema"
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
	if syncGroupName == "" {
		return nil, errors.New("parameter syncGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncGroupName}", url.PathEscape(syncGroupName))
	if syncMemberName == "" {
		return nil, errors.New("parameter syncMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncMemberName}", url.PathEscape(syncMemberName))
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
	reqQP.Set("api-version", "2020-11-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	return req, nil
}

// refreshMemberSchemaHandleError handles the RefreshMemberSchema error response.
func (client *SyncMembersClient) refreshMemberSchemaHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginUpdate - Updates an existing sync member.
// If the operation fails it returns a generic error.
func (client *SyncMembersClient) BeginUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string, parameters SyncMember, options *SyncMembersBeginUpdateOptions) (SyncMembersUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, serverName, databaseName, syncGroupName, syncMemberName, parameters, options)
	if err != nil {
		return SyncMembersUpdatePollerResponse{}, err
	}
	result := SyncMembersUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("SyncMembersClient.Update", "", resp, client.con.Pipeline(), client.updateHandleError)
	if err != nil {
		return SyncMembersUpdatePollerResponse{}, err
	}
	poller := &syncMembersUpdatePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SyncMembersUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeUpdate creates a new SyncMembersUpdatePoller from the specified resume token.
// token - The value must come from a previous call to SyncMembersUpdatePoller.ResumeToken().
func (client *SyncMembersClient) ResumeUpdate(ctx context.Context, token string) (SyncMembersUpdatePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("SyncMembersClient.Update", token, client.con.Pipeline(), client.updateHandleError)
	if err != nil {
		return SyncMembersUpdatePollerResponse{}, err
	}
	poller := &syncMembersUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return SyncMembersUpdatePollerResponse{}, err
	}
	result := SyncMembersUpdatePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SyncMembersUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Update - Updates an existing sync member.
// If the operation fails it returns a generic error.
func (client *SyncMembersClient) update(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string, parameters SyncMember, options *SyncMembersBeginUpdateOptions) (*azcore.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serverName, databaseName, syncGroupName, syncMemberName, parameters, options)
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
func (client *SyncMembersClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string, parameters SyncMember, options *SyncMembersBeginUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/syncGroups/{syncGroupName}/syncMembers/{syncMemberName}"
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
	if syncGroupName == "" {
		return nil, errors.New("parameter syncGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncGroupName}", url.PathEscape(syncGroupName))
	if syncMemberName == "" {
		return nil, errors.New("parameter syncMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncMemberName}", url.PathEscape(syncMemberName))
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
func (client *SyncMembersClient) updateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
