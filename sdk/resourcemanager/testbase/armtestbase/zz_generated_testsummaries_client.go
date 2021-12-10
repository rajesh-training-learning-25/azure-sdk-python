//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtestbase

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

// TestSummariesClient contains the methods for the TestSummaries group.
// Don't use this type directly, use NewTestSummariesClient() instead.
type TestSummariesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewTestSummariesClient creates a new instance of TestSummariesClient with the specified values.
func NewTestSummariesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *TestSummariesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &TestSummariesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Gets a Test Summary with specific name from all the Test Summaries of all the packages under a Test Base Account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TestSummariesClient) Get(ctx context.Context, resourceGroupName string, testBaseAccountName string, testSummaryName string, options *TestSummariesGetOptions) (TestSummariesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, testBaseAccountName, testSummaryName, options)
	if err != nil {
		return TestSummariesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TestSummariesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TestSummariesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TestSummariesClient) getCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, testSummaryName string, options *TestSummariesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/testSummaries/{testSummaryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	if testSummaryName == "" {
		return nil, errors.New("parameter testSummaryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testSummaryName}", url.PathEscape(testSummaryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TestSummariesClient) getHandleResponse(resp *http.Response) (TestSummariesGetResponse, error) {
	result := TestSummariesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TestSummaryResource); err != nil {
		return TestSummariesGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *TestSummariesClient) getHandleError(resp *http.Response) error {
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

// List - Lists the Test Summaries of all the packages under a Test Base Account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TestSummariesClient) List(resourceGroupName string, testBaseAccountName string, options *TestSummariesListOptions) *TestSummariesListPager {
	return &TestSummariesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, testBaseAccountName, options)
		},
		advancer: func(ctx context.Context, resp TestSummariesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.TestSummaryListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *TestSummariesClient) listCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, options *TestSummariesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/testSummaries"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TestSummariesClient) listHandleResponse(resp *http.Response) (TestSummariesListResponse, error) {
	result := TestSummariesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TestSummaryListResult); err != nil {
		return TestSummariesListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *TestSummariesClient) listHandleError(resp *http.Response) error {
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
