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

// WorkloadClassifiersClient contains the methods for the WorkloadClassifiers group.
// Don't use this type directly, use NewWorkloadClassifiersClient() instead.
type WorkloadClassifiersClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewWorkloadClassifiersClient creates a new instance of WorkloadClassifiersClient with the specified values.
func NewWorkloadClassifiersClient(con *arm.Connection, subscriptionID string) *WorkloadClassifiersClient {
	return &WorkloadClassifiersClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a workload classifier.
// If the operation fails it returns a generic error.
func (client *WorkloadClassifiersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, workloadGroupName string, workloadClassifierName string, parameters WorkloadClassifier, options *WorkloadClassifiersBeginCreateOrUpdateOptions) (WorkloadClassifiersCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, databaseName, workloadGroupName, workloadClassifierName, parameters, options)
	if err != nil {
		return WorkloadClassifiersCreateOrUpdatePollerResponse{}, err
	}
	result := WorkloadClassifiersCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("WorkloadClassifiersClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return WorkloadClassifiersCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &WorkloadClassifiersCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a workload classifier.
// If the operation fails it returns a generic error.
func (client *WorkloadClassifiersClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, workloadGroupName string, workloadClassifierName string, parameters WorkloadClassifier, options *WorkloadClassifiersBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, databaseName, workloadGroupName, workloadClassifierName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WorkloadClassifiersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, workloadGroupName string, workloadClassifierName string, parameters WorkloadClassifier, options *WorkloadClassifiersBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/workloadGroups/{workloadGroupName}/workloadClassifiers/{workloadClassifierName}"
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
	if workloadGroupName == "" {
		return nil, errors.New("parameter workloadGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workloadGroupName}", url.PathEscape(workloadGroupName))
	if workloadClassifierName == "" {
		return nil, errors.New("parameter workloadClassifierName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workloadClassifierName}", url.PathEscape(workloadClassifierName))
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
func (client *WorkloadClassifiersClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginDelete - Deletes a workload classifier.
// If the operation fails it returns a generic error.
func (client *WorkloadClassifiersClient) BeginDelete(ctx context.Context, resourceGroupName string, serverName string, databaseName string, workloadGroupName string, workloadClassifierName string, options *WorkloadClassifiersBeginDeleteOptions) (WorkloadClassifiersDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, serverName, databaseName, workloadGroupName, workloadClassifierName, options)
	if err != nil {
		return WorkloadClassifiersDeletePollerResponse{}, err
	}
	result := WorkloadClassifiersDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("WorkloadClassifiersClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return WorkloadClassifiersDeletePollerResponse{}, err
	}
	result.Poller = &WorkloadClassifiersDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a workload classifier.
// If the operation fails it returns a generic error.
func (client *WorkloadClassifiersClient) deleteOperation(ctx context.Context, resourceGroupName string, serverName string, databaseName string, workloadGroupName string, workloadClassifierName string, options *WorkloadClassifiersBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, databaseName, workloadGroupName, workloadClassifierName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WorkloadClassifiersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, workloadGroupName string, workloadClassifierName string, options *WorkloadClassifiersBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/workloadGroups/{workloadGroupName}/workloadClassifiers/{workloadClassifierName}"
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
	if workloadGroupName == "" {
		return nil, errors.New("parameter workloadGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workloadGroupName}", url.PathEscape(workloadGroupName))
	if workloadClassifierName == "" {
		return nil, errors.New("parameter workloadClassifierName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workloadClassifierName}", url.PathEscape(workloadClassifierName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *WorkloadClassifiersClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Gets a workload classifier
// If the operation fails it returns a generic error.
func (client *WorkloadClassifiersClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, workloadGroupName string, workloadClassifierName string, options *WorkloadClassifiersGetOptions) (WorkloadClassifiersGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, databaseName, workloadGroupName, workloadClassifierName, options)
	if err != nil {
		return WorkloadClassifiersGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkloadClassifiersGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkloadClassifiersGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkloadClassifiersClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, workloadGroupName string, workloadClassifierName string, options *WorkloadClassifiersGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/workloadGroups/{workloadGroupName}/workloadClassifiers/{workloadClassifierName}"
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
	if workloadGroupName == "" {
		return nil, errors.New("parameter workloadGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workloadGroupName}", url.PathEscape(workloadGroupName))
	if workloadClassifierName == "" {
		return nil, errors.New("parameter workloadClassifierName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workloadClassifierName}", url.PathEscape(workloadClassifierName))
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
func (client *WorkloadClassifiersClient) getHandleResponse(resp *http.Response) (WorkloadClassifiersGetResponse, error) {
	result := WorkloadClassifiersGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkloadClassifier); err != nil {
		return WorkloadClassifiersGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *WorkloadClassifiersClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByWorkloadGroup - Gets the list of workload classifiers for a workload group
// If the operation fails it returns a generic error.
func (client *WorkloadClassifiersClient) ListByWorkloadGroup(resourceGroupName string, serverName string, databaseName string, workloadGroupName string, options *WorkloadClassifiersListByWorkloadGroupOptions) *WorkloadClassifiersListByWorkloadGroupPager {
	return &WorkloadClassifiersListByWorkloadGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByWorkloadGroupCreateRequest(ctx, resourceGroupName, serverName, databaseName, workloadGroupName, options)
		},
		advancer: func(ctx context.Context, resp WorkloadClassifiersListByWorkloadGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.WorkloadClassifierListResult.NextLink)
		},
	}
}

// listByWorkloadGroupCreateRequest creates the ListByWorkloadGroup request.
func (client *WorkloadClassifiersClient) listByWorkloadGroupCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, workloadGroupName string, options *WorkloadClassifiersListByWorkloadGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/workloadGroups/{workloadGroupName}/workloadClassifiers"
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
	if workloadGroupName == "" {
		return nil, errors.New("parameter workloadGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workloadGroupName}", url.PathEscape(workloadGroupName))
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

// listByWorkloadGroupHandleResponse handles the ListByWorkloadGroup response.
func (client *WorkloadClassifiersClient) listByWorkloadGroupHandleResponse(resp *http.Response) (WorkloadClassifiersListByWorkloadGroupResponse, error) {
	result := WorkloadClassifiersListByWorkloadGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkloadClassifierListResult); err != nil {
		return WorkloadClassifiersListByWorkloadGroupResponse{}, err
	}
	return result, nil
}

// listByWorkloadGroupHandleError handles the ListByWorkloadGroup error response.
func (client *WorkloadClassifiersClient) listByWorkloadGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
