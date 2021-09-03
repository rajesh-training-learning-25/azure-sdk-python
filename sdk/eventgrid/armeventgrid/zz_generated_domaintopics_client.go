//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventgrid

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// DomainTopicsClient contains the methods for the DomainTopics group.
// Don't use this type directly, use NewDomainTopicsClient() instead.
type DomainTopicsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewDomainTopicsClient creates a new instance of DomainTopicsClient with the specified values.
func NewDomainTopicsClient(con *arm.Connection, subscriptionID string) *DomainTopicsClient {
	return &DomainTopicsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Asynchronously creates or updates a new domain topic with the specified parameters.
// If the operation fails it returns a generic error.
func (client *DomainTopicsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, domainName string, domainTopicName string, options *DomainTopicsBeginCreateOrUpdateOptions) (DomainTopicsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, domainName, domainTopicName, options)
	if err != nil {
		return DomainTopicsCreateOrUpdatePollerResponse{}, err
	}
	result := DomainTopicsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DomainTopicsClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return DomainTopicsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &DomainTopicsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Asynchronously creates or updates a new domain topic with the specified parameters.
// If the operation fails it returns a generic error.
func (client *DomainTopicsClient) createOrUpdate(ctx context.Context, resourceGroupName string, domainName string, domainTopicName string, options *DomainTopicsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, domainName, domainTopicName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DomainTopicsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, domainName string, domainTopicName string, options *DomainTopicsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/topics/{domainTopicName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	if domainTopicName == "" {
		return nil, errors.New("parameter domainTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainTopicName}", url.PathEscape(domainTopicName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *DomainTopicsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginDelete - Delete existing domain topic.
// If the operation fails it returns a generic error.
func (client *DomainTopicsClient) BeginDelete(ctx context.Context, resourceGroupName string, domainName string, domainTopicName string, options *DomainTopicsBeginDeleteOptions) (DomainTopicsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, domainName, domainTopicName, options)
	if err != nil {
		return DomainTopicsDeletePollerResponse{}, err
	}
	result := DomainTopicsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DomainTopicsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return DomainTopicsDeletePollerResponse{}, err
	}
	result.Poller = &DomainTopicsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete existing domain topic.
// If the operation fails it returns a generic error.
func (client *DomainTopicsClient) deleteOperation(ctx context.Context, resourceGroupName string, domainName string, domainTopicName string, options *DomainTopicsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, domainName, domainTopicName, options)
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
func (client *DomainTopicsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, domainName string, domainTopicName string, options *DomainTopicsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/topics/{domainTopicName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	if domainTopicName == "" {
		return nil, errors.New("parameter domainTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainTopicName}", url.PathEscape(domainTopicName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *DomainTopicsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Get properties of a domain topic.
// If the operation fails it returns a generic error.
func (client *DomainTopicsClient) Get(ctx context.Context, resourceGroupName string, domainName string, domainTopicName string, options *DomainTopicsGetOptions) (DomainTopicsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, domainName, domainTopicName, options)
	if err != nil {
		return DomainTopicsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DomainTopicsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DomainTopicsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DomainTopicsClient) getCreateRequest(ctx context.Context, resourceGroupName string, domainName string, domainTopicName string, options *DomainTopicsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/topics/{domainTopicName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	if domainTopicName == "" {
		return nil, errors.New("parameter domainTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainTopicName}", url.PathEscape(domainTopicName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DomainTopicsClient) getHandleResponse(resp *http.Response) (DomainTopicsGetResponse, error) {
	result := DomainTopicsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DomainTopic); err != nil {
		return DomainTopicsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DomainTopicsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByDomain - List all the topics in a domain.
// If the operation fails it returns a generic error.
func (client *DomainTopicsClient) ListByDomain(resourceGroupName string, domainName string, options *DomainTopicsListByDomainOptions) *DomainTopicsListByDomainPager {
	return &DomainTopicsListByDomainPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByDomainCreateRequest(ctx, resourceGroupName, domainName, options)
		},
		advancer: func(ctx context.Context, resp DomainTopicsListByDomainResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DomainTopicsListResult.NextLink)
		},
	}
}

// listByDomainCreateRequest creates the ListByDomain request.
func (client *DomainTopicsClient) listByDomainCreateRequest(ctx context.Context, resourceGroupName string, domainName string, options *DomainTopicsListByDomainOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/topics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByDomainHandleResponse handles the ListByDomain response.
func (client *DomainTopicsClient) listByDomainHandleResponse(resp *http.Response) (DomainTopicsListByDomainResponse, error) {
	result := DomainTopicsListByDomainResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DomainTopicsListResult); err != nil {
		return DomainTopicsListByDomainResponse{}, err
	}
	return result, nil
}

// listByDomainHandleError handles the ListByDomain error response.
func (client *DomainTopicsClient) listByDomainHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
