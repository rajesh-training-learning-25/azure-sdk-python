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

// ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient contains the methods for the ManagedRestorableDroppedDatabaseBackupShortTermRetentionPolicies group.
// Don't use this type directly, use NewManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient() instead.
type ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient creates a new instance of ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient with the specified values.
func NewManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient(con *armcore.Connection, subscriptionID string) *ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient {
	return &ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Sets a database's short term retention policy.
// If the operation fails it returns a generic error.
func (client *ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, restorableDroppedDatabaseID string, policyName ManagedShortTermRetentionPolicyName, parameters ManagedBackupShortTermRetentionPolicy, options *ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesBeginCreateOrUpdateOptions) (ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, managedInstanceName, restorableDroppedDatabaseID, policyName, parameters, options)
	if err != nil {
		return ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesCreateOrUpdatePollerResponse{}, err
	}
	result := ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesCreateOrUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient.CreateOrUpdate", "", resp, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesCreateOrUpdatePollerResponse{}, err
	}
	poller := &managedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesCreateOrUpdatePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesCreateOrUpdatePoller from the specified resume token.
// token - The value must come from a previous call to ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesCreateOrUpdatePoller.ResumeToken().
func (client *ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) ResumeCreateOrUpdate(ctx context.Context, token string) (ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesCreateOrUpdatePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient.CreateOrUpdate", token, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesCreateOrUpdatePollerResponse{}, err
	}
	poller := &managedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesCreateOrUpdatePollerResponse{}, err
	}
	result := ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Sets a database's short term retention policy.
// If the operation fails it returns a generic error.
func (client *ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) createOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, restorableDroppedDatabaseID string, policyName ManagedShortTermRetentionPolicyName, parameters ManagedBackupShortTermRetentionPolicy, options *ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, managedInstanceName, restorableDroppedDatabaseID, policyName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, restorableDroppedDatabaseID string, policyName ManagedShortTermRetentionPolicyName, parameters ManagedBackupShortTermRetentionPolicy, options *ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/restorableDroppedDatabases/{restorableDroppedDatabaseId}/backupShortTermRetentionPolicies/{policyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if restorableDroppedDatabaseID == "" {
		return nil, errors.New("parameter restorableDroppedDatabaseID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorableDroppedDatabaseId}", url.PathEscape(restorableDroppedDatabaseID))
	if policyName == "" {
		return nil, errors.New("parameter policyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(string(policyName)))
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
func (client *ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Get - Gets a dropped database's short term retention policy.
// If the operation fails it returns a generic error.
func (client *ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, restorableDroppedDatabaseID string, policyName ManagedShortTermRetentionPolicyName, options *ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesGetOptions) (ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, restorableDroppedDatabaseID, policyName, options)
	if err != nil {
		return ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, restorableDroppedDatabaseID string, policyName ManagedShortTermRetentionPolicyName, options *ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/restorableDroppedDatabases/{restorableDroppedDatabaseId}/backupShortTermRetentionPolicies/{policyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if restorableDroppedDatabaseID == "" {
		return nil, errors.New("parameter restorableDroppedDatabaseID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorableDroppedDatabaseId}", url.PathEscape(restorableDroppedDatabaseID))
	if policyName == "" {
		return nil, errors.New("parameter policyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(string(policyName)))
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
func (client *ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) getHandleResponse(resp *azcore.Response) (ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesGetResponse, error) {
	result := ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ManagedBackupShortTermRetentionPolicy); err != nil {
		return ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListByRestorableDroppedDatabase - Gets a dropped database's short term retention policy list.
// If the operation fails it returns a generic error.
func (client *ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) ListByRestorableDroppedDatabase(resourceGroupName string, managedInstanceName string, restorableDroppedDatabaseID string, options *ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesListByRestorableDroppedDatabaseOptions) ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesListByRestorableDroppedDatabasePager {
	return &managedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesListByRestorableDroppedDatabasePager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByRestorableDroppedDatabaseCreateRequest(ctx, resourceGroupName, managedInstanceName, restorableDroppedDatabaseID, options)
		},
		advancer: func(ctx context.Context, resp ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesListByRestorableDroppedDatabaseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ManagedBackupShortTermRetentionPolicyListResult.NextLink)
		},
	}
}

// listByRestorableDroppedDatabaseCreateRequest creates the ListByRestorableDroppedDatabase request.
func (client *ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) listByRestorableDroppedDatabaseCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, restorableDroppedDatabaseID string, options *ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesListByRestorableDroppedDatabaseOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/restorableDroppedDatabases/{restorableDroppedDatabaseId}/backupShortTermRetentionPolicies"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if restorableDroppedDatabaseID == "" {
		return nil, errors.New("parameter restorableDroppedDatabaseID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorableDroppedDatabaseId}", url.PathEscape(restorableDroppedDatabaseID))
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

// listByRestorableDroppedDatabaseHandleResponse handles the ListByRestorableDroppedDatabase response.
func (client *ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) listByRestorableDroppedDatabaseHandleResponse(resp *azcore.Response) (ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesListByRestorableDroppedDatabaseResponse, error) {
	result := ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesListByRestorableDroppedDatabaseResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ManagedBackupShortTermRetentionPolicyListResult); err != nil {
		return ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesListByRestorableDroppedDatabaseResponse{}, err
	}
	return result, nil
}

// listByRestorableDroppedDatabaseHandleError handles the ListByRestorableDroppedDatabase error response.
func (client *ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) listByRestorableDroppedDatabaseHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginUpdate - Sets a database's short term retention policy.
// If the operation fails it returns a generic error.
func (client *ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) BeginUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, restorableDroppedDatabaseID string, policyName ManagedShortTermRetentionPolicyName, parameters ManagedBackupShortTermRetentionPolicy, options *ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesBeginUpdateOptions) (ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, managedInstanceName, restorableDroppedDatabaseID, policyName, parameters, options)
	if err != nil {
		return ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesUpdatePollerResponse{}, err
	}
	result := ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient.Update", "", resp, client.con.Pipeline(), client.updateHandleError)
	if err != nil {
		return ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesUpdatePollerResponse{}, err
	}
	poller := &managedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesUpdatePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeUpdate creates a new ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesUpdatePoller from the specified resume token.
// token - The value must come from a previous call to ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesUpdatePoller.ResumeToken().
func (client *ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) ResumeUpdate(ctx context.Context, token string) (ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesUpdatePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient.Update", token, client.con.Pipeline(), client.updateHandleError)
	if err != nil {
		return ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesUpdatePollerResponse{}, err
	}
	poller := &managedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesUpdatePollerResponse{}, err
	}
	result := ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesUpdatePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Update - Sets a database's short term retention policy.
// If the operation fails it returns a generic error.
func (client *ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) update(ctx context.Context, resourceGroupName string, managedInstanceName string, restorableDroppedDatabaseID string, policyName ManagedShortTermRetentionPolicyName, parameters ManagedBackupShortTermRetentionPolicy, options *ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesBeginUpdateOptions) (*azcore.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, managedInstanceName, restorableDroppedDatabaseID, policyName, parameters, options)
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
func (client *ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, restorableDroppedDatabaseID string, policyName ManagedShortTermRetentionPolicyName, parameters ManagedBackupShortTermRetentionPolicy, options *ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesBeginUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/restorableDroppedDatabases/{restorableDroppedDatabaseId}/backupShortTermRetentionPolicies/{policyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if restorableDroppedDatabaseID == "" {
		return nil, errors.New("parameter restorableDroppedDatabaseID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorableDroppedDatabaseId}", url.PathEscape(restorableDroppedDatabaseID))
	if policyName == "" {
		return nil, errors.New("parameter policyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(string(policyName)))
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
func (client *ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) updateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
