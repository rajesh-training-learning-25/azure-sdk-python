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
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// ManagedBackupShortTermRetentionPoliciesClient contains the methods for the ManagedBackupShortTermRetentionPolicies group.
// Don't use this type directly, use NewManagedBackupShortTermRetentionPoliciesClient() instead.
type ManagedBackupShortTermRetentionPoliciesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewManagedBackupShortTermRetentionPoliciesClient creates a new instance of ManagedBackupShortTermRetentionPoliciesClient with the specified values.
func NewManagedBackupShortTermRetentionPoliciesClient(con *arm.Connection, subscriptionID string) *ManagedBackupShortTermRetentionPoliciesClient {
	return &ManagedBackupShortTermRetentionPoliciesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Updates a managed database's short term retention policy.
// If the operation fails it returns a generic error.
func (client *ManagedBackupShortTermRetentionPoliciesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, policyName ManagedShortTermRetentionPolicyName, parameters ManagedBackupShortTermRetentionPolicy, options *ManagedBackupShortTermRetentionPoliciesBeginCreateOrUpdateOptions) (ManagedBackupShortTermRetentionPoliciesCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, managedInstanceName, databaseName, policyName, parameters, options)
	if err != nil {
		return ManagedBackupShortTermRetentionPoliciesCreateOrUpdatePollerResponse{}, err
	}
	result := ManagedBackupShortTermRetentionPoliciesCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ManagedBackupShortTermRetentionPoliciesClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return ManagedBackupShortTermRetentionPoliciesCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ManagedBackupShortTermRetentionPoliciesCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Updates a managed database's short term retention policy.
// If the operation fails it returns a generic error.
func (client *ManagedBackupShortTermRetentionPoliciesClient) createOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, policyName ManagedShortTermRetentionPolicyName, parameters ManagedBackupShortTermRetentionPolicy, options *ManagedBackupShortTermRetentionPoliciesBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, policyName, parameters, options)
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
func (client *ManagedBackupShortTermRetentionPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, policyName ManagedShortTermRetentionPolicyName, parameters ManagedBackupShortTermRetentionPolicy, options *ManagedBackupShortTermRetentionPoliciesBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/backupShortTermRetentionPolicies/{policyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if policyName == "" {
		return nil, errors.New("parameter policyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(string(policyName)))
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
func (client *ManagedBackupShortTermRetentionPoliciesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Gets a managed database's short term retention policy.
// If the operation fails it returns a generic error.
func (client *ManagedBackupShortTermRetentionPoliciesClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, policyName ManagedShortTermRetentionPolicyName, options *ManagedBackupShortTermRetentionPoliciesGetOptions) (ManagedBackupShortTermRetentionPoliciesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, policyName, options)
	if err != nil {
		return ManagedBackupShortTermRetentionPoliciesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedBackupShortTermRetentionPoliciesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedBackupShortTermRetentionPoliciesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ManagedBackupShortTermRetentionPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, policyName ManagedShortTermRetentionPolicyName, options *ManagedBackupShortTermRetentionPoliciesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/backupShortTermRetentionPolicies/{policyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if policyName == "" {
		return nil, errors.New("parameter policyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(string(policyName)))
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
func (client *ManagedBackupShortTermRetentionPoliciesClient) getHandleResponse(resp *http.Response) (ManagedBackupShortTermRetentionPoliciesGetResponse, error) {
	result := ManagedBackupShortTermRetentionPoliciesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedBackupShortTermRetentionPolicy); err != nil {
		return ManagedBackupShortTermRetentionPoliciesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ManagedBackupShortTermRetentionPoliciesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByDatabase - Gets a managed database's short term retention policy list.
// If the operation fails it returns a generic error.
func (client *ManagedBackupShortTermRetentionPoliciesClient) ListByDatabase(resourceGroupName string, managedInstanceName string, databaseName string, options *ManagedBackupShortTermRetentionPoliciesListByDatabaseOptions) *ManagedBackupShortTermRetentionPoliciesListByDatabasePager {
	return &ManagedBackupShortTermRetentionPoliciesListByDatabasePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByDatabaseCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, options)
		},
		advancer: func(ctx context.Context, resp ManagedBackupShortTermRetentionPoliciesListByDatabaseResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ManagedBackupShortTermRetentionPolicyListResult.NextLink)
		},
	}
}

// listByDatabaseCreateRequest creates the ListByDatabase request.
func (client *ManagedBackupShortTermRetentionPoliciesClient) listByDatabaseCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, options *ManagedBackupShortTermRetentionPoliciesListByDatabaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/backupShortTermRetentionPolicies"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
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

// listByDatabaseHandleResponse handles the ListByDatabase response.
func (client *ManagedBackupShortTermRetentionPoliciesClient) listByDatabaseHandleResponse(resp *http.Response) (ManagedBackupShortTermRetentionPoliciesListByDatabaseResponse, error) {
	result := ManagedBackupShortTermRetentionPoliciesListByDatabaseResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedBackupShortTermRetentionPolicyListResult); err != nil {
		return ManagedBackupShortTermRetentionPoliciesListByDatabaseResponse{}, err
	}
	return result, nil
}

// listByDatabaseHandleError handles the ListByDatabase error response.
func (client *ManagedBackupShortTermRetentionPoliciesClient) listByDatabaseHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginUpdate - Updates a managed database's short term retention policy.
// If the operation fails it returns a generic error.
func (client *ManagedBackupShortTermRetentionPoliciesClient) BeginUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, policyName ManagedShortTermRetentionPolicyName, parameters ManagedBackupShortTermRetentionPolicy, options *ManagedBackupShortTermRetentionPoliciesBeginUpdateOptions) (ManagedBackupShortTermRetentionPoliciesUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, managedInstanceName, databaseName, policyName, parameters, options)
	if err != nil {
		return ManagedBackupShortTermRetentionPoliciesUpdatePollerResponse{}, err
	}
	result := ManagedBackupShortTermRetentionPoliciesUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ManagedBackupShortTermRetentionPoliciesClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return ManagedBackupShortTermRetentionPoliciesUpdatePollerResponse{}, err
	}
	result.Poller = &ManagedBackupShortTermRetentionPoliciesUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates a managed database's short term retention policy.
// If the operation fails it returns a generic error.
func (client *ManagedBackupShortTermRetentionPoliciesClient) update(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, policyName ManagedShortTermRetentionPolicyName, parameters ManagedBackupShortTermRetentionPolicy, options *ManagedBackupShortTermRetentionPoliciesBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, policyName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ManagedBackupShortTermRetentionPoliciesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, policyName ManagedShortTermRetentionPolicyName, parameters ManagedBackupShortTermRetentionPolicy, options *ManagedBackupShortTermRetentionPoliciesBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/backupShortTermRetentionPolicies/{policyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if policyName == "" {
		return nil, errors.New("parameter policyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(string(policyName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleError handles the Update error response.
func (client *ManagedBackupShortTermRetentionPoliciesClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
