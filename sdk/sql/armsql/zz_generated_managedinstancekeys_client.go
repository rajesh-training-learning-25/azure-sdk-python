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

// ManagedInstanceKeysClient contains the methods for the ManagedInstanceKeys group.
// Don't use this type directly, use NewManagedInstanceKeysClient() instead.
type ManagedInstanceKeysClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewManagedInstanceKeysClient creates a new instance of ManagedInstanceKeysClient with the specified values.
func NewManagedInstanceKeysClient(con *armcore.Connection, subscriptionID string) *ManagedInstanceKeysClient {
	return &ManagedInstanceKeysClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a managed instance key.
// If the operation fails it returns a generic error.
func (client *ManagedInstanceKeysClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, keyName string, parameters ManagedInstanceKey, options *ManagedInstanceKeysBeginCreateOrUpdateOptions) (ManagedInstanceKeysCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, managedInstanceName, keyName, parameters, options)
	if err != nil {
		return ManagedInstanceKeysCreateOrUpdatePollerResponse{}, err
	}
	result := ManagedInstanceKeysCreateOrUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ManagedInstanceKeysClient.CreateOrUpdate", "", resp, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return ManagedInstanceKeysCreateOrUpdatePollerResponse{}, err
	}
	poller := &managedInstanceKeysCreateOrUpdatePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ManagedInstanceKeysCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new ManagedInstanceKeysCreateOrUpdatePoller from the specified resume token.
// token - The value must come from a previous call to ManagedInstanceKeysCreateOrUpdatePoller.ResumeToken().
func (client *ManagedInstanceKeysClient) ResumeCreateOrUpdate(ctx context.Context, token string) (ManagedInstanceKeysCreateOrUpdatePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ManagedInstanceKeysClient.CreateOrUpdate", token, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return ManagedInstanceKeysCreateOrUpdatePollerResponse{}, err
	}
	poller := &managedInstanceKeysCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ManagedInstanceKeysCreateOrUpdatePollerResponse{}, err
	}
	result := ManagedInstanceKeysCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ManagedInstanceKeysCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a managed instance key.
// If the operation fails it returns a generic error.
func (client *ManagedInstanceKeysClient) createOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, keyName string, parameters ManagedInstanceKey, options *ManagedInstanceKeysBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, managedInstanceName, keyName, parameters, options)
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
func (client *ManagedInstanceKeysClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, keyName string, parameters ManagedInstanceKey, options *ManagedInstanceKeysBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/keys/{keyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if keyName == "" {
		return nil, errors.New("parameter keyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{keyName}", url.PathEscape(keyName))
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
func (client *ManagedInstanceKeysClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginDelete - Deletes the managed instance key with the given name.
// If the operation fails it returns a generic error.
func (client *ManagedInstanceKeysClient) BeginDelete(ctx context.Context, resourceGroupName string, managedInstanceName string, keyName string, options *ManagedInstanceKeysBeginDeleteOptions) (ManagedInstanceKeysDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, managedInstanceName, keyName, options)
	if err != nil {
		return ManagedInstanceKeysDeletePollerResponse{}, err
	}
	result := ManagedInstanceKeysDeletePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ManagedInstanceKeysClient.Delete", "", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return ManagedInstanceKeysDeletePollerResponse{}, err
	}
	poller := &managedInstanceKeysDeletePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ManagedInstanceKeysDeleteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new ManagedInstanceKeysDeletePoller from the specified resume token.
// token - The value must come from a previous call to ManagedInstanceKeysDeletePoller.ResumeToken().
func (client *ManagedInstanceKeysClient) ResumeDelete(ctx context.Context, token string) (ManagedInstanceKeysDeletePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ManagedInstanceKeysClient.Delete", token, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return ManagedInstanceKeysDeletePollerResponse{}, err
	}
	poller := &managedInstanceKeysDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ManagedInstanceKeysDeletePollerResponse{}, err
	}
	result := ManagedInstanceKeysDeletePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ManagedInstanceKeysDeleteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Deletes the managed instance key with the given name.
// If the operation fails it returns a generic error.
func (client *ManagedInstanceKeysClient) deleteOperation(ctx context.Context, resourceGroupName string, managedInstanceName string, keyName string, options *ManagedInstanceKeysBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, managedInstanceName, keyName, options)
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
func (client *ManagedInstanceKeysClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, keyName string, options *ManagedInstanceKeysBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/keys/{keyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if keyName == "" {
		return nil, errors.New("parameter keyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{keyName}", url.PathEscape(keyName))
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
func (client *ManagedInstanceKeysClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Get - Gets a managed instance key.
// If the operation fails it returns a generic error.
func (client *ManagedInstanceKeysClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, keyName string, options *ManagedInstanceKeysGetOptions) (ManagedInstanceKeysGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, keyName, options)
	if err != nil {
		return ManagedInstanceKeysGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ManagedInstanceKeysGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ManagedInstanceKeysGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ManagedInstanceKeysClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, keyName string, options *ManagedInstanceKeysGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/keys/{keyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if keyName == "" {
		return nil, errors.New("parameter keyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{keyName}", url.PathEscape(keyName))
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
func (client *ManagedInstanceKeysClient) getHandleResponse(resp *azcore.Response) (ManagedInstanceKeysGetResponse, error) {
	result := ManagedInstanceKeysGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ManagedInstanceKey); err != nil {
		return ManagedInstanceKeysGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ManagedInstanceKeysClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListByInstance - Gets a list of managed instance keys.
// If the operation fails it returns a generic error.
func (client *ManagedInstanceKeysClient) ListByInstance(resourceGroupName string, managedInstanceName string, options *ManagedInstanceKeysListByInstanceOptions) ManagedInstanceKeysListByInstancePager {
	return &managedInstanceKeysListByInstancePager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByInstanceCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
		},
		advancer: func(ctx context.Context, resp ManagedInstanceKeysListByInstanceResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ManagedInstanceKeyListResult.NextLink)
		},
	}
}

// listByInstanceCreateRequest creates the ListByInstance request.
func (client *ManagedInstanceKeysClient) listByInstanceCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstanceKeysListByInstanceOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/keys"
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
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2020-11-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByInstanceHandleResponse handles the ListByInstance response.
func (client *ManagedInstanceKeysClient) listByInstanceHandleResponse(resp *azcore.Response) (ManagedInstanceKeysListByInstanceResponse, error) {
	result := ManagedInstanceKeysListByInstanceResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ManagedInstanceKeyListResult); err != nil {
		return ManagedInstanceKeysListByInstanceResponse{}, err
	}
	return result, nil
}

// listByInstanceHandleError handles the ListByInstance error response.
func (client *ManagedInstanceKeysClient) listByInstanceHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
