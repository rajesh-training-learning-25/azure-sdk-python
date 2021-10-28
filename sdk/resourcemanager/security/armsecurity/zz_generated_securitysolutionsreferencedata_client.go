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

// SecuritySolutionsReferenceDataClient contains the methods for the SecuritySolutionsReferenceData group.
// Don't use this type directly, use NewSecuritySolutionsReferenceDataClient() instead.
type SecuritySolutionsReferenceDataClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
	ascLocation    string
}

// NewSecuritySolutionsReferenceDataClient creates a new instance of SecuritySolutionsReferenceDataClient with the specified values.
func NewSecuritySolutionsReferenceDataClient(subscriptionID string, ascLocation string, credential azcore.TokenCredential, options *arm.ClientOptions) *SecuritySolutionsReferenceDataClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &SecuritySolutionsReferenceDataClient{subscriptionID: subscriptionID, ascLocation: ascLocation, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// List - Gets a list of all supported Security Solutions for the subscription.
// If the operation fails it returns the *CloudError error type.
func (client *SecuritySolutionsReferenceDataClient) List(ctx context.Context, options *SecuritySolutionsReferenceDataListOptions) (SecuritySolutionsReferenceDataListResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return SecuritySolutionsReferenceDataListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SecuritySolutionsReferenceDataListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SecuritySolutionsReferenceDataListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *SecuritySolutionsReferenceDataClient) listCreateRequest(ctx context.Context, options *SecuritySolutionsReferenceDataListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/securitySolutionsReferenceData"
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
func (client *SecuritySolutionsReferenceDataClient) listHandleResponse(resp *http.Response) (SecuritySolutionsReferenceDataListResponse, error) {
	result := SecuritySolutionsReferenceDataListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecuritySolutionsReferenceDataList); err != nil {
		return SecuritySolutionsReferenceDataListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *SecuritySolutionsReferenceDataClient) listHandleError(resp *http.Response) error {
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

// ListByHomeRegion - Gets list of all supported Security Solutions for subscription and location.
// If the operation fails it returns the *CloudError error type.
func (client *SecuritySolutionsReferenceDataClient) ListByHomeRegion(ctx context.Context, options *SecuritySolutionsReferenceDataListByHomeRegionOptions) (SecuritySolutionsReferenceDataListByHomeRegionResponse, error) {
	req, err := client.listByHomeRegionCreateRequest(ctx, options)
	if err != nil {
		return SecuritySolutionsReferenceDataListByHomeRegionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SecuritySolutionsReferenceDataListByHomeRegionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SecuritySolutionsReferenceDataListByHomeRegionResponse{}, client.listByHomeRegionHandleError(resp)
	}
	return client.listByHomeRegionHandleResponse(resp)
}

// listByHomeRegionCreateRequest creates the ListByHomeRegion request.
func (client *SecuritySolutionsReferenceDataClient) listByHomeRegionCreateRequest(ctx context.Context, options *SecuritySolutionsReferenceDataListByHomeRegionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/securitySolutionsReferenceData"
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
func (client *SecuritySolutionsReferenceDataClient) listByHomeRegionHandleResponse(resp *http.Response) (SecuritySolutionsReferenceDataListByHomeRegionResponse, error) {
	result := SecuritySolutionsReferenceDataListByHomeRegionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecuritySolutionsReferenceDataList); err != nil {
		return SecuritySolutionsReferenceDataListByHomeRegionResponse{}, err
	}
	return result, nil
}

// listByHomeRegionHandleError handles the ListByHomeRegion error response.
func (client *SecuritySolutionsReferenceDataClient) listByHomeRegionHandleError(resp *http.Response) error {
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
