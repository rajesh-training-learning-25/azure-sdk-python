// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventgrid

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

// SystemTopicsClient contains the methods for the SystemTopics group.
// Don't use this type directly, use NewSystemTopicsClient() instead.
type SystemTopicsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewSystemTopicsClient creates a new instance of SystemTopicsClient with the specified values.
func NewSystemTopicsClient(con *armcore.Connection, subscriptionID string) *SystemTopicsClient {
	return &SystemTopicsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Asynchronously creates a new system topic with the specified parameters.
// If the operation fails it returns a generic error.
func (client *SystemTopicsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, systemTopicName string, systemTopicInfo SystemTopic, options *SystemTopicsBeginCreateOrUpdateOptions) (SystemTopicsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, systemTopicName, systemTopicInfo, options)
	if err != nil {
		return SystemTopicsCreateOrUpdatePollerResponse{}, err
	}
	result := SystemTopicsCreateOrUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("SystemTopicsClient.CreateOrUpdate", "", resp, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return SystemTopicsCreateOrUpdatePollerResponse{}, err
	}
	poller := &systemTopicsCreateOrUpdatePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SystemTopicsCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new SystemTopicsCreateOrUpdatePoller from the specified resume token.
// token - The value must come from a previous call to SystemTopicsCreateOrUpdatePoller.ResumeToken().
func (client *SystemTopicsClient) ResumeCreateOrUpdate(ctx context.Context, token string) (SystemTopicsCreateOrUpdatePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("SystemTopicsClient.CreateOrUpdate", token, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return SystemTopicsCreateOrUpdatePollerResponse{}, err
	}
	poller := &systemTopicsCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return SystemTopicsCreateOrUpdatePollerResponse{}, err
	}
	result := SystemTopicsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SystemTopicsCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Asynchronously creates a new system topic with the specified parameters.
// If the operation fails it returns a generic error.
func (client *SystemTopicsClient) createOrUpdate(ctx context.Context, resourceGroupName string, systemTopicName string, systemTopicInfo SystemTopic, options *SystemTopicsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, systemTopicName, systemTopicInfo, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SystemTopicsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, systemTopicName string, systemTopicInfo SystemTopic, options *SystemTopicsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if systemTopicName == "" {
		return nil, errors.New("parameter systemTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemTopicName}", url.PathEscape(systemTopicName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(systemTopicInfo)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *SystemTopicsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginDelete - Delete existing system topic.
// If the operation fails it returns a generic error.
func (client *SystemTopicsClient) BeginDelete(ctx context.Context, resourceGroupName string, systemTopicName string, options *SystemTopicsBeginDeleteOptions) (SystemTopicsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, systemTopicName, options)
	if err != nil {
		return SystemTopicsDeletePollerResponse{}, err
	}
	result := SystemTopicsDeletePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("SystemTopicsClient.Delete", "", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return SystemTopicsDeletePollerResponse{}, err
	}
	poller := &systemTopicsDeletePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SystemTopicsDeleteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new SystemTopicsDeletePoller from the specified resume token.
// token - The value must come from a previous call to SystemTopicsDeletePoller.ResumeToken().
func (client *SystemTopicsClient) ResumeDelete(ctx context.Context, token string) (SystemTopicsDeletePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("SystemTopicsClient.Delete", token, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return SystemTopicsDeletePollerResponse{}, err
	}
	poller := &systemTopicsDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return SystemTopicsDeletePollerResponse{}, err
	}
	result := SystemTopicsDeletePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SystemTopicsDeleteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Delete existing system topic.
// If the operation fails it returns a generic error.
func (client *SystemTopicsClient) deleteOperation(ctx context.Context, resourceGroupName string, systemTopicName string, options *SystemTopicsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, systemTopicName, options)
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
func (client *SystemTopicsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, systemTopicName string, options *SystemTopicsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if systemTopicName == "" {
		return nil, errors.New("parameter systemTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemTopicName}", url.PathEscape(systemTopicName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *SystemTopicsClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Get - Get properties of a system topic.
// If the operation fails it returns a generic error.
func (client *SystemTopicsClient) Get(ctx context.Context, resourceGroupName string, systemTopicName string, options *SystemTopicsGetOptions) (SystemTopicsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, systemTopicName, options)
	if err != nil {
		return SystemTopicsGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SystemTopicsGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SystemTopicsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SystemTopicsClient) getCreateRequest(ctx context.Context, resourceGroupName string, systemTopicName string, options *SystemTopicsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if systemTopicName == "" {
		return nil, errors.New("parameter systemTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemTopicName}", url.PathEscape(systemTopicName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SystemTopicsClient) getHandleResponse(resp *azcore.Response) (SystemTopicsGetResponse, error) {
	result := SystemTopicsGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.SystemTopic); err != nil {
		return SystemTopicsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *SystemTopicsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListByResourceGroup - List all the system topics under a resource group.
// If the operation fails it returns a generic error.
func (client *SystemTopicsClient) ListByResourceGroup(resourceGroupName string, options *SystemTopicsListByResourceGroupOptions) SystemTopicsListByResourceGroupPager {
	return &systemTopicsListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp SystemTopicsListByResourceGroupResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.SystemTopicsListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SystemTopicsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *SystemTopicsListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *SystemTopicsClient) listByResourceGroupHandleResponse(resp *azcore.Response) (SystemTopicsListByResourceGroupResponse, error) {
	result := SystemTopicsListByResourceGroupResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.SystemTopicsListResult); err != nil {
		return SystemTopicsListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *SystemTopicsClient) listByResourceGroupHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListBySubscription - List all the system topics under an Azure subscription.
// If the operation fails it returns a generic error.
func (client *SystemTopicsClient) ListBySubscription(options *SystemTopicsListBySubscriptionOptions) SystemTopicsListBySubscriptionPager {
	return &systemTopicsListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp SystemTopicsListBySubscriptionResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.SystemTopicsListResult.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *SystemTopicsClient) listBySubscriptionCreateRequest(ctx context.Context, options *SystemTopicsListBySubscriptionOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.EventGrid/systemTopics"
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
	reqQP.Set("api-version", "2021-06-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *SystemTopicsClient) listBySubscriptionHandleResponse(resp *azcore.Response) (SystemTopicsListBySubscriptionResponse, error) {
	result := SystemTopicsListBySubscriptionResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.SystemTopicsListResult); err != nil {
		return SystemTopicsListBySubscriptionResponse{}, err
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *SystemTopicsClient) listBySubscriptionHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginUpdate - Asynchronously updates a system topic with the specified parameters.
// If the operation fails it returns a generic error.
func (client *SystemTopicsClient) BeginUpdate(ctx context.Context, resourceGroupName string, systemTopicName string, systemTopicUpdateParameters SystemTopicUpdateParameters, options *SystemTopicsBeginUpdateOptions) (SystemTopicsUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, systemTopicName, systemTopicUpdateParameters, options)
	if err != nil {
		return SystemTopicsUpdatePollerResponse{}, err
	}
	result := SystemTopicsUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("SystemTopicsClient.Update", "", resp, client.con.Pipeline(), client.updateHandleError)
	if err != nil {
		return SystemTopicsUpdatePollerResponse{}, err
	}
	poller := &systemTopicsUpdatePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SystemTopicsUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeUpdate creates a new SystemTopicsUpdatePoller from the specified resume token.
// token - The value must come from a previous call to SystemTopicsUpdatePoller.ResumeToken().
func (client *SystemTopicsClient) ResumeUpdate(ctx context.Context, token string) (SystemTopicsUpdatePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("SystemTopicsClient.Update", token, client.con.Pipeline(), client.updateHandleError)
	if err != nil {
		return SystemTopicsUpdatePollerResponse{}, err
	}
	poller := &systemTopicsUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return SystemTopicsUpdatePollerResponse{}, err
	}
	result := SystemTopicsUpdatePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SystemTopicsUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Update - Asynchronously updates a system topic with the specified parameters.
// If the operation fails it returns a generic error.
func (client *SystemTopicsClient) update(ctx context.Context, resourceGroupName string, systemTopicName string, systemTopicUpdateParameters SystemTopicUpdateParameters, options *SystemTopicsBeginUpdateOptions) (*azcore.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, systemTopicName, systemTopicUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *SystemTopicsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, systemTopicName string, systemTopicUpdateParameters SystemTopicUpdateParameters, options *SystemTopicsBeginUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if systemTopicName == "" {
		return nil, errors.New("parameter systemTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemTopicName}", url.PathEscape(systemTopicName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(systemTopicUpdateParameters)
}

// updateHandleError handles the Update error response.
func (client *SystemTopicsClient) updateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
