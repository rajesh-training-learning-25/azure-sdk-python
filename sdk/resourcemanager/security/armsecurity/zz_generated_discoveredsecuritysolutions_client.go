//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

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
	"strings"
)

// DiscoveredSecuritySolutionsClient contains the methods for the DiscoveredSecuritySolutions group.
// Don't use this type directly, use NewDiscoveredSecuritySolutionsClient() instead.
type DiscoveredSecuritySolutionsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
	ascLocation    string
}

// NewDiscoveredSecuritySolutionsClient creates a new instance of DiscoveredSecuritySolutionsClient with the specified values.
func NewDiscoveredSecuritySolutionsClient(subscriptionID string, ascLocation string, credential azcore.TokenCredential, options *arm.ClientOptions) *DiscoveredSecuritySolutionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &DiscoveredSecuritySolutionsClient{subscriptionID: subscriptionID, ascLocation: ascLocation, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Gets a specific discovered Security Solution.
// If the operation fails it returns the *CloudError error type.
func (client *DiscoveredSecuritySolutionsClient) Get(ctx context.Context, resourceGroupName string, discoveredSecuritySolutionName string, options *DiscoveredSecuritySolutionsGetOptions) (DiscoveredSecuritySolutionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, discoveredSecuritySolutionName, options)
	if err != nil {
		return DiscoveredSecuritySolutionsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DiscoveredSecuritySolutionsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DiscoveredSecuritySolutionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DiscoveredSecuritySolutionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, discoveredSecuritySolutionName string, options *DiscoveredSecuritySolutionsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/locations/{ascLocation}/discoveredSecuritySolutions/{discoveredSecuritySolutionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.ascLocation == "" {
		return nil, errors.New("parameter client.ascLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ascLocation}", url.PathEscape(client.ascLocation))
	if discoveredSecuritySolutionName == "" {
		return nil, errors.New("parameter discoveredSecuritySolutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{discoveredSecuritySolutionName}", url.PathEscape(discoveredSecuritySolutionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DiscoveredSecuritySolutionsClient) getHandleResponse(resp *http.Response) (DiscoveredSecuritySolutionsGetResponse, error) {
	result := DiscoveredSecuritySolutionsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiscoveredSecuritySolution); err != nil {
		return DiscoveredSecuritySolutionsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DiscoveredSecuritySolutionsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Gets a list of discovered Security Solutions for the subscription.
// If the operation fails it returns the *CloudError error type.
func (client *DiscoveredSecuritySolutionsClient) List(options *DiscoveredSecuritySolutionsListOptions) *DiscoveredSecuritySolutionsListPager {
	return &DiscoveredSecuritySolutionsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp DiscoveredSecuritySolutionsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DiscoveredSecuritySolutionList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *DiscoveredSecuritySolutionsClient) listCreateRequest(ctx context.Context, options *DiscoveredSecuritySolutionsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/discoveredSecuritySolutions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DiscoveredSecuritySolutionsClient) listHandleResponse(resp *http.Response) (DiscoveredSecuritySolutionsListResponse, error) {
	result := DiscoveredSecuritySolutionsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiscoveredSecuritySolutionList); err != nil {
		return DiscoveredSecuritySolutionsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *DiscoveredSecuritySolutionsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByHomeRegion - Gets a list of discovered Security Solutions for the subscription and location.
// If the operation fails it returns the *CloudError error type.
func (client *DiscoveredSecuritySolutionsClient) ListByHomeRegion(options *DiscoveredSecuritySolutionsListByHomeRegionOptions) *DiscoveredSecuritySolutionsListByHomeRegionPager {
	return &DiscoveredSecuritySolutionsListByHomeRegionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByHomeRegionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp DiscoveredSecuritySolutionsListByHomeRegionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DiscoveredSecuritySolutionList.NextLink)
		},
	}
}

// listByHomeRegionCreateRequest creates the ListByHomeRegion request.
func (client *DiscoveredSecuritySolutionsClient) listByHomeRegionCreateRequest(ctx context.Context, options *DiscoveredSecuritySolutionsListByHomeRegionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/discoveredSecuritySolutions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if client.ascLocation == "" {
		return nil, errors.New("parameter client.ascLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ascLocation}", url.PathEscape(client.ascLocation))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByHomeRegionHandleResponse handles the ListByHomeRegion response.
func (client *DiscoveredSecuritySolutionsClient) listByHomeRegionHandleResponse(resp *http.Response) (DiscoveredSecuritySolutionsListByHomeRegionResponse, error) {
	result := DiscoveredSecuritySolutionsListByHomeRegionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiscoveredSecuritySolutionList); err != nil {
		return DiscoveredSecuritySolutionsListByHomeRegionResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByHomeRegionHandleError handles the ListByHomeRegion error response.
func (client *DiscoveredSecuritySolutionsClient) listByHomeRegionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
