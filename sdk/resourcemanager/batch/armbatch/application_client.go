//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armbatch

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ApplicationClient contains the methods for the Application group.
// Don't use this type directly, use NewApplicationClient() instead.
type ApplicationClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewApplicationClient creates a new instance of ApplicationClient with the specified values.
// subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewApplicationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ApplicationClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ApplicationClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Create - Adds an application to the specified Batch account.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01
// resourceGroupName - The name of the resource group that contains the Batch account.
// accountName - The name of the Batch account.
// applicationName - The name of the application. This must be unique within the account.
// options - ApplicationClientCreateOptions contains the optional parameters for the ApplicationClient.Create method.
func (client *ApplicationClient) Create(ctx context.Context, resourceGroupName string, accountName string, applicationName string, options *ApplicationClientCreateOptions) (ApplicationClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, applicationName, options)
	if err != nil {
		return ApplicationClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApplicationClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *ApplicationClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, applicationName string, options *ApplicationClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Parameters != nil {
		return req, runtime.MarshalAsJSON(req, *options.Parameters)
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *ApplicationClient) createHandleResponse(resp *http.Response) (ApplicationClientCreateResponse, error) {
	result := ApplicationClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Application); err != nil {
		return ApplicationClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an application.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01
// resourceGroupName - The name of the resource group that contains the Batch account.
// accountName - The name of the Batch account.
// applicationName - The name of the application. This must be unique within the account.
// options - ApplicationClientDeleteOptions contains the optional parameters for the ApplicationClient.Delete method.
func (client *ApplicationClient) Delete(ctx context.Context, resourceGroupName string, accountName string, applicationName string, options *ApplicationClientDeleteOptions) (ApplicationClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, applicationName, options)
	if err != nil {
		return ApplicationClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ApplicationClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ApplicationClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ApplicationClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, applicationName string, options *ApplicationClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets information about the specified application.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01
// resourceGroupName - The name of the resource group that contains the Batch account.
// accountName - The name of the Batch account.
// applicationName - The name of the application. This must be unique within the account.
// options - ApplicationClientGetOptions contains the optional parameters for the ApplicationClient.Get method.
func (client *ApplicationClient) Get(ctx context.Context, resourceGroupName string, accountName string, applicationName string, options *ApplicationClientGetOptions) (ApplicationClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, applicationName, options)
	if err != nil {
		return ApplicationClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApplicationClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ApplicationClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, applicationName string, options *ApplicationClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ApplicationClient) getHandleResponse(resp *http.Response) (ApplicationClientGetResponse, error) {
	result := ApplicationClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Application); err != nil {
		return ApplicationClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all of the applications in the specified account.
// Generated from API version 2022-10-01
// resourceGroupName - The name of the resource group that contains the Batch account.
// accountName - The name of the Batch account.
// options - ApplicationClientListOptions contains the optional parameters for the ApplicationClient.List method.
func (client *ApplicationClient) NewListPager(resourceGroupName string, accountName string, options *ApplicationClientListOptions) *runtime.Pager[ApplicationClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ApplicationClientListResponse]{
		More: func(page ApplicationClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ApplicationClientListResponse) (ApplicationClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, accountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ApplicationClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ApplicationClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ApplicationClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ApplicationClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *ApplicationClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Maxresults != nil {
		reqQP.Set("maxresults", strconv.FormatInt(int64(*options.Maxresults), 10))
	}
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ApplicationClient) listHandleResponse(resp *http.Response) (ApplicationClientListResponse, error) {
	result := ApplicationClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListApplicationsResult); err != nil {
		return ApplicationClientListResponse{}, err
	}
	return result, nil
}

// Update - Updates settings for the specified application.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01
// resourceGroupName - The name of the resource group that contains the Batch account.
// accountName - The name of the Batch account.
// applicationName - The name of the application. This must be unique within the account.
// parameters - The parameters for the request.
// options - ApplicationClientUpdateOptions contains the optional parameters for the ApplicationClient.Update method.
func (client *ApplicationClient) Update(ctx context.Context, resourceGroupName string, accountName string, applicationName string, parameters Application, options *ApplicationClientUpdateOptions) (ApplicationClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, applicationName, parameters, options)
	if err != nil {
		return ApplicationClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApplicationClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ApplicationClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, applicationName string, parameters Application, options *ApplicationClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *ApplicationClient) updateHandleResponse(resp *http.Response) (ApplicationClientUpdateResponse, error) {
	result := ApplicationClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Application); err != nil {
		return ApplicationClientUpdateResponse{}, err
	}
	return result, nil
}