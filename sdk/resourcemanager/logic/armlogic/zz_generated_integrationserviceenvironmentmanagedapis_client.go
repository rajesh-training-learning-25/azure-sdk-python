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
	"strings"
)

// IntegrationServiceEnvironmentManagedApisClient contains the methods for the IntegrationServiceEnvironmentManagedApis group.
// Don't use this type directly, use NewIntegrationServiceEnvironmentManagedApisClient() instead.
type IntegrationServiceEnvironmentManagedApisClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewIntegrationServiceEnvironmentManagedApisClient creates a new instance of IntegrationServiceEnvironmentManagedApisClient with the specified values.
func NewIntegrationServiceEnvironmentManagedApisClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *IntegrationServiceEnvironmentManagedApisClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &IntegrationServiceEnvironmentManagedApisClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginDelete - Deletes the integration service environment managed Api.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationServiceEnvironmentManagedApisClient) BeginDelete(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, apiName string, options *IntegrationServiceEnvironmentManagedApisBeginDeleteOptions) (IntegrationServiceEnvironmentManagedApisDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroup, integrationServiceEnvironmentName, apiName, options)
	if err != nil {
		return IntegrationServiceEnvironmentManagedApisDeletePollerResponse{}, err
	}
	result := IntegrationServiceEnvironmentManagedApisDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("IntegrationServiceEnvironmentManagedApisClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return IntegrationServiceEnvironmentManagedApisDeletePollerResponse{}, err
	}
	result.Poller = &IntegrationServiceEnvironmentManagedApisDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the integration service environment managed Api.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationServiceEnvironmentManagedApisClient) deleteOperation(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, apiName string, options *IntegrationServiceEnvironmentManagedApisBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroup, integrationServiceEnvironmentName, apiName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IntegrationServiceEnvironmentManagedApisClient) deleteCreateRequest(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, apiName string, options *IntegrationServiceEnvironmentManagedApisBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/managedApis/{apiName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if integrationServiceEnvironmentName == "" {
		return nil, errors.New("parameter integrationServiceEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationServiceEnvironmentName}", url.PathEscape(integrationServiceEnvironmentName))
	if apiName == "" {
		return nil, errors.New("parameter apiName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiName}", url.PathEscape(apiName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *IntegrationServiceEnvironmentManagedApisClient) deleteHandleError(resp *http.Response) error {
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

// Get - Gets the integration service environment managed Api.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationServiceEnvironmentManagedApisClient) Get(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, apiName string, options *IntegrationServiceEnvironmentManagedApisGetOptions) (IntegrationServiceEnvironmentManagedApisGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroup, integrationServiceEnvironmentName, apiName, options)
	if err != nil {
		return IntegrationServiceEnvironmentManagedApisGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationServiceEnvironmentManagedApisGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationServiceEnvironmentManagedApisGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IntegrationServiceEnvironmentManagedApisClient) getCreateRequest(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, apiName string, options *IntegrationServiceEnvironmentManagedApisGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/managedApis/{apiName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if integrationServiceEnvironmentName == "" {
		return nil, errors.New("parameter integrationServiceEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationServiceEnvironmentName}", url.PathEscape(integrationServiceEnvironmentName))
	if apiName == "" {
		return nil, errors.New("parameter apiName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiName}", url.PathEscape(apiName))
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
func (client *IntegrationServiceEnvironmentManagedApisClient) getHandleResponse(resp *http.Response) (IntegrationServiceEnvironmentManagedApisGetResponse, error) {
	result := IntegrationServiceEnvironmentManagedApisGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationServiceEnvironmentManagedAPI); err != nil {
		return IntegrationServiceEnvironmentManagedApisGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *IntegrationServiceEnvironmentManagedApisClient) getHandleError(resp *http.Response) error {
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

// List - Gets the integration service environment managed Apis.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationServiceEnvironmentManagedApisClient) List(resourceGroup string, integrationServiceEnvironmentName string, options *IntegrationServiceEnvironmentManagedApisListOptions) *IntegrationServiceEnvironmentManagedApisListPager {
	return &IntegrationServiceEnvironmentManagedApisListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroup, integrationServiceEnvironmentName, options)
		},
		advancer: func(ctx context.Context, resp IntegrationServiceEnvironmentManagedApisListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.IntegrationServiceEnvironmentManagedAPIListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *IntegrationServiceEnvironmentManagedApisClient) listCreateRequest(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, options *IntegrationServiceEnvironmentManagedApisListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/managedApis"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if integrationServiceEnvironmentName == "" {
		return nil, errors.New("parameter integrationServiceEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationServiceEnvironmentName}", url.PathEscape(integrationServiceEnvironmentName))
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

// listHandleResponse handles the List response.
func (client *IntegrationServiceEnvironmentManagedApisClient) listHandleResponse(resp *http.Response) (IntegrationServiceEnvironmentManagedApisListResponse, error) {
	result := IntegrationServiceEnvironmentManagedApisListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationServiceEnvironmentManagedAPIListResult); err != nil {
		return IntegrationServiceEnvironmentManagedApisListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *IntegrationServiceEnvironmentManagedApisClient) listHandleError(resp *http.Response) error {
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

// BeginPut - Puts the integration service environment managed Api.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationServiceEnvironmentManagedApisClient) BeginPut(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, apiName string, integrationServiceEnvironmentManagedAPI IntegrationServiceEnvironmentManagedAPI, options *IntegrationServiceEnvironmentManagedApisBeginPutOptions) (IntegrationServiceEnvironmentManagedApisPutPollerResponse, error) {
	resp, err := client.put(ctx, resourceGroup, integrationServiceEnvironmentName, apiName, integrationServiceEnvironmentManagedAPI, options)
	if err != nil {
		return IntegrationServiceEnvironmentManagedApisPutPollerResponse{}, err
	}
	result := IntegrationServiceEnvironmentManagedApisPutPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("IntegrationServiceEnvironmentManagedApisClient.Put", "", resp, client.pl, client.putHandleError)
	if err != nil {
		return IntegrationServiceEnvironmentManagedApisPutPollerResponse{}, err
	}
	result.Poller = &IntegrationServiceEnvironmentManagedApisPutPoller{
		pt: pt,
	}
	return result, nil
}

// Put - Puts the integration service environment managed Api.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationServiceEnvironmentManagedApisClient) put(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, apiName string, integrationServiceEnvironmentManagedAPI IntegrationServiceEnvironmentManagedAPI, options *IntegrationServiceEnvironmentManagedApisBeginPutOptions) (*http.Response, error) {
	req, err := client.putCreateRequest(ctx, resourceGroup, integrationServiceEnvironmentName, apiName, integrationServiceEnvironmentManagedAPI, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.putHandleError(resp)
	}
	return resp, nil
}

// putCreateRequest creates the Put request.
func (client *IntegrationServiceEnvironmentManagedApisClient) putCreateRequest(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, apiName string, integrationServiceEnvironmentManagedAPI IntegrationServiceEnvironmentManagedAPI, options *IntegrationServiceEnvironmentManagedApisBeginPutOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/managedApis/{apiName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if integrationServiceEnvironmentName == "" {
		return nil, errors.New("parameter integrationServiceEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationServiceEnvironmentName}", url.PathEscape(integrationServiceEnvironmentName))
	if apiName == "" {
		return nil, errors.New("parameter apiName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiName}", url.PathEscape(apiName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, integrationServiceEnvironmentManagedAPI)
}

// putHandleError handles the Put error response.
func (client *IntegrationServiceEnvironmentManagedApisClient) putHandleError(resp *http.Response) error {
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
