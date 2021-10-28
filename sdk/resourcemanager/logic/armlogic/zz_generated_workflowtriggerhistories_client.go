//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// WorkflowTriggerHistoriesClient contains the methods for the WorkflowTriggerHistories group.
// Don't use this type directly, use NewWorkflowTriggerHistoriesClient() instead.
type WorkflowTriggerHistoriesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewWorkflowTriggerHistoriesClient creates a new instance of WorkflowTriggerHistoriesClient with the specified values.
func NewWorkflowTriggerHistoriesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *WorkflowTriggerHistoriesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &WorkflowTriggerHistoriesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Gets a workflow trigger history.
// If the operation fails it returns the *ErrorResponse error type.
func (client *WorkflowTriggerHistoriesClient) Get(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, historyName string, options *WorkflowTriggerHistoriesGetOptions) (WorkflowTriggerHistoriesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workflowName, triggerName, historyName, options)
	if err != nil {
		return WorkflowTriggerHistoriesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkflowTriggerHistoriesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkflowTriggerHistoriesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkflowTriggerHistoriesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, historyName string, options *WorkflowTriggerHistoriesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}/histories/{historyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	if historyName == "" {
		return nil, errors.New("parameter historyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{historyName}", url.PathEscape(historyName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkflowTriggerHistoriesClient) getHandleResponse(resp *http.Response) (WorkflowTriggerHistoriesGetResponse, error) {
	result := WorkflowTriggerHistoriesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowTriggerHistory); err != nil {
		return WorkflowTriggerHistoriesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *WorkflowTriggerHistoriesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Gets a list of workflow trigger histories.
// If the operation fails it returns the *ErrorResponse error type.
func (client *WorkflowTriggerHistoriesClient) List(resourceGroupName string, workflowName string, triggerName string, options *WorkflowTriggerHistoriesListOptions) *WorkflowTriggerHistoriesListPager {
	return &WorkflowTriggerHistoriesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, workflowName, triggerName, options)
		},
		advancer: func(ctx context.Context, resp WorkflowTriggerHistoriesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.WorkflowTriggerHistoryListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *WorkflowTriggerHistoriesClient) listCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, options *WorkflowTriggerHistoriesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}/histories"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WorkflowTriggerHistoriesClient) listHandleResponse(resp *http.Response) (WorkflowTriggerHistoriesListResponse, error) {
	result := WorkflowTriggerHistoriesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowTriggerHistoryListResult); err != nil {
		return WorkflowTriggerHistoriesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *WorkflowTriggerHistoriesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Resubmit - Resubmits a workflow run based on the trigger history.
// If the operation fails it returns the *ErrorResponse error type.
func (client *WorkflowTriggerHistoriesClient) Resubmit(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, historyName string, options *WorkflowTriggerHistoriesResubmitOptions) (WorkflowTriggerHistoriesResubmitResponse, error) {
	req, err := client.resubmitCreateRequest(ctx, resourceGroupName, workflowName, triggerName, historyName, options)
	if err != nil {
		return WorkflowTriggerHistoriesResubmitResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkflowTriggerHistoriesResubmitResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return WorkflowTriggerHistoriesResubmitResponse{}, client.resubmitHandleError(resp)
	}
	return WorkflowTriggerHistoriesResubmitResponse{RawResponse: resp}, nil
}

// resubmitCreateRequest creates the Resubmit request.
func (client *WorkflowTriggerHistoriesClient) resubmitCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, historyName string, options *WorkflowTriggerHistoriesResubmitOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}/histories/{historyName}/resubmit"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	if historyName == "" {
		return nil, errors.New("parameter historyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{historyName}", url.PathEscape(historyName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// resubmitHandleError handles the Resubmit error response.
func (client *WorkflowTriggerHistoriesClient) resubmitHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
