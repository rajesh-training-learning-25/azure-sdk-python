//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmediaservices

import (
	"context"
	"errors"
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

// StreamingPoliciesClient contains the methods for the StreamingPolicies group.
// Don't use this type directly, use NewStreamingPoliciesClient() instead.
type StreamingPoliciesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewStreamingPoliciesClient creates a new instance of StreamingPoliciesClient with the specified values.
// subscriptionID - The unique identifier for a Microsoft Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewStreamingPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *StreamingPoliciesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &StreamingPoliciesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Create - Create a Streaming Policy in the Media Services account
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// streamingPolicyName - The Streaming Policy name.
// parameters - The request parameters
// options - StreamingPoliciesClientCreateOptions contains the optional parameters for the StreamingPoliciesClient.Create
// method.
func (client *StreamingPoliciesClient) Create(ctx context.Context, resourceGroupName string, accountName string, streamingPolicyName string, parameters StreamingPolicy, options *StreamingPoliciesClientCreateOptions) (StreamingPoliciesClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, streamingPolicyName, parameters, options)
	if err != nil {
		return StreamingPoliciesClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StreamingPoliciesClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return StreamingPoliciesClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *StreamingPoliciesClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, streamingPolicyName string, parameters StreamingPolicy, options *StreamingPoliciesClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingPolicies/{streamingPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if streamingPolicyName == "" {
		return nil, errors.New("parameter streamingPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{streamingPolicyName}", url.PathEscape(streamingPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createHandleResponse handles the Create response.
func (client *StreamingPoliciesClient) createHandleResponse(resp *http.Response) (StreamingPoliciesClientCreateResponse, error) {
	result := StreamingPoliciesClientCreateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.StreamingPolicy); err != nil {
		return StreamingPoliciesClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a Streaming Policy in the Media Services account
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// streamingPolicyName - The Streaming Policy name.
// options - StreamingPoliciesClientDeleteOptions contains the optional parameters for the StreamingPoliciesClient.Delete
// method.
func (client *StreamingPoliciesClient) Delete(ctx context.Context, resourceGroupName string, accountName string, streamingPolicyName string, options *StreamingPoliciesClientDeleteOptions) (StreamingPoliciesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, streamingPolicyName, options)
	if err != nil {
		return StreamingPoliciesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StreamingPoliciesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return StreamingPoliciesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return StreamingPoliciesClientDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *StreamingPoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, streamingPolicyName string, options *StreamingPoliciesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingPolicies/{streamingPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if streamingPolicyName == "" {
		return nil, errors.New("parameter streamingPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{streamingPolicyName}", url.PathEscape(streamingPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get the details of a Streaming Policy in the Media Services account
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// streamingPolicyName - The Streaming Policy name.
// options - StreamingPoliciesClientGetOptions contains the optional parameters for the StreamingPoliciesClient.Get method.
func (client *StreamingPoliciesClient) Get(ctx context.Context, resourceGroupName string, accountName string, streamingPolicyName string, options *StreamingPoliciesClientGetOptions) (StreamingPoliciesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, streamingPolicyName, options)
	if err != nil {
		return StreamingPoliciesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StreamingPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StreamingPoliciesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *StreamingPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, streamingPolicyName string, options *StreamingPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingPolicies/{streamingPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if streamingPolicyName == "" {
		return nil, errors.New("parameter streamingPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{streamingPolicyName}", url.PathEscape(streamingPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *StreamingPoliciesClient) getHandleResponse(resp *http.Response) (StreamingPoliciesClientGetResponse, error) {
	result := StreamingPoliciesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.StreamingPolicy); err != nil {
		return StreamingPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists the Streaming Policies in the account
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// options - StreamingPoliciesClientListOptions contains the optional parameters for the StreamingPoliciesClient.List method.
func (client *StreamingPoliciesClient) List(resourceGroupName string, accountName string, options *StreamingPoliciesClientListOptions) *StreamingPoliciesClientListPager {
	return &StreamingPoliciesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, accountName, options)
		},
		advancer: func(ctx context.Context, resp StreamingPoliciesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.StreamingPolicyCollection.ODataNextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *StreamingPoliciesClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *StreamingPoliciesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingPolicies"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *StreamingPoliciesClient) listHandleResponse(resp *http.Response) (StreamingPoliciesClientListResponse, error) {
	result := StreamingPoliciesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.StreamingPolicyCollection); err != nil {
		return StreamingPoliciesClientListResponse{}, err
	}
	return result, nil
}
