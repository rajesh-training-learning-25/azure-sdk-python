//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armwindowsesu

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
	"strings"
)

// MultipleActivationKeysClient contains the methods for the MultipleActivationKeys group.
// Don't use this type directly, use NewMultipleActivationKeysClient() instead.
type MultipleActivationKeysClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewMultipleActivationKeysClient creates a new instance of MultipleActivationKeysClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewMultipleActivationKeysClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MultipleActivationKeysClient, error) {
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
	client := &MultipleActivationKeysClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Create a MAK key.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-09-16-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// multipleActivationKeyName - The name of the MAK key.
// multipleActivationKey - Details of the MAK key.
// options - MultipleActivationKeysClientBeginCreateOptions contains the optional parameters for the MultipleActivationKeysClient.BeginCreate
// method.
func (client *MultipleActivationKeysClient) BeginCreate(ctx context.Context, resourceGroupName string, multipleActivationKeyName string, multipleActivationKey MultipleActivationKey, options *MultipleActivationKeysClientBeginCreateOptions) (*runtime.Poller[MultipleActivationKeysClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, multipleActivationKeyName, multipleActivationKey, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[MultipleActivationKeysClientCreateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[MultipleActivationKeysClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Create a MAK key.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-09-16-preview
func (client *MultipleActivationKeysClient) create(ctx context.Context, resourceGroupName string, multipleActivationKeyName string, multipleActivationKey MultipleActivationKey, options *MultipleActivationKeysClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, multipleActivationKeyName, multipleActivationKey, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *MultipleActivationKeysClient) createCreateRequest(ctx context.Context, resourceGroupName string, multipleActivationKeyName string, multipleActivationKey MultipleActivationKey, options *MultipleActivationKeysClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.WindowsESU/multipleActivationKeys/{multipleActivationKeyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if multipleActivationKeyName == "" {
		return nil, errors.New("parameter multipleActivationKeyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{multipleActivationKeyName}", url.PathEscape(multipleActivationKeyName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-09-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, multipleActivationKey)
}

// Delete - Delete a MAK key.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-09-16-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// multipleActivationKeyName - The name of the MAK key.
// options - MultipleActivationKeysClientDeleteOptions contains the optional parameters for the MultipleActivationKeysClient.Delete
// method.
func (client *MultipleActivationKeysClient) Delete(ctx context.Context, resourceGroupName string, multipleActivationKeyName string, options *MultipleActivationKeysClientDeleteOptions) (MultipleActivationKeysClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, multipleActivationKeyName, options)
	if err != nil {
		return MultipleActivationKeysClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MultipleActivationKeysClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return MultipleActivationKeysClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return MultipleActivationKeysClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *MultipleActivationKeysClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, multipleActivationKeyName string, options *MultipleActivationKeysClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.WindowsESU/multipleActivationKeys/{multipleActivationKeyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if multipleActivationKeyName == "" {
		return nil, errors.New("parameter multipleActivationKeyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{multipleActivationKeyName}", url.PathEscape(multipleActivationKeyName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-09-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a MAK key.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-09-16-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// multipleActivationKeyName - The name of the MAK key.
// options - MultipleActivationKeysClientGetOptions contains the optional parameters for the MultipleActivationKeysClient.Get
// method.
func (client *MultipleActivationKeysClient) Get(ctx context.Context, resourceGroupName string, multipleActivationKeyName string, options *MultipleActivationKeysClientGetOptions) (MultipleActivationKeysClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, multipleActivationKeyName, options)
	if err != nil {
		return MultipleActivationKeysClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MultipleActivationKeysClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MultipleActivationKeysClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *MultipleActivationKeysClient) getCreateRequest(ctx context.Context, resourceGroupName string, multipleActivationKeyName string, options *MultipleActivationKeysClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.WindowsESU/multipleActivationKeys/{multipleActivationKeyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if multipleActivationKeyName == "" {
		return nil, errors.New("parameter multipleActivationKeyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{multipleActivationKeyName}", url.PathEscape(multipleActivationKeyName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-09-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MultipleActivationKeysClient) getHandleResponse(resp *http.Response) (MultipleActivationKeysClientGetResponse, error) {
	result := MultipleActivationKeysClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MultipleActivationKey); err != nil {
		return MultipleActivationKeysClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List all Multiple Activation Keys (MAK) created for a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-09-16-preview
// options - MultipleActivationKeysClientListOptions contains the optional parameters for the MultipleActivationKeysClient.List
// method.
func (client *MultipleActivationKeysClient) NewListPager(options *MultipleActivationKeysClientListOptions) *runtime.Pager[MultipleActivationKeysClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[MultipleActivationKeysClientListResponse]{
		More: func(page MultipleActivationKeysClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MultipleActivationKeysClientListResponse) (MultipleActivationKeysClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return MultipleActivationKeysClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return MultipleActivationKeysClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return MultipleActivationKeysClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *MultipleActivationKeysClient) listCreateRequest(ctx context.Context, options *MultipleActivationKeysClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.WindowsESU/multipleActivationKeys"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-09-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MultipleActivationKeysClient) listHandleResponse(resp *http.Response) (MultipleActivationKeysClientListResponse, error) {
	result := MultipleActivationKeysClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MultipleActivationKeyList); err != nil {
		return MultipleActivationKeysClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List all Multiple Activation Keys (MAK) in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-09-16-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - MultipleActivationKeysClientListByResourceGroupOptions contains the optional parameters for the MultipleActivationKeysClient.ListByResourceGroup
// method.
func (client *MultipleActivationKeysClient) NewListByResourceGroupPager(resourceGroupName string, options *MultipleActivationKeysClientListByResourceGroupOptions) *runtime.Pager[MultipleActivationKeysClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[MultipleActivationKeysClientListByResourceGroupResponse]{
		More: func(page MultipleActivationKeysClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MultipleActivationKeysClientListByResourceGroupResponse) (MultipleActivationKeysClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return MultipleActivationKeysClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return MultipleActivationKeysClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return MultipleActivationKeysClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *MultipleActivationKeysClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *MultipleActivationKeysClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.WindowsESU/multipleActivationKeys"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-09-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *MultipleActivationKeysClient) listByResourceGroupHandleResponse(resp *http.Response) (MultipleActivationKeysClientListByResourceGroupResponse, error) {
	result := MultipleActivationKeysClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MultipleActivationKeyList); err != nil {
		return MultipleActivationKeysClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// Update - Update a MAK key.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-09-16-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// multipleActivationKeyName - The name of the MAK key.
// multipleActivationKey - Details of the MAK key.
// options - MultipleActivationKeysClientUpdateOptions contains the optional parameters for the MultipleActivationKeysClient.Update
// method.
func (client *MultipleActivationKeysClient) Update(ctx context.Context, resourceGroupName string, multipleActivationKeyName string, multipleActivationKey MultipleActivationKeyUpdate, options *MultipleActivationKeysClientUpdateOptions) (MultipleActivationKeysClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, multipleActivationKeyName, multipleActivationKey, options)
	if err != nil {
		return MultipleActivationKeysClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MultipleActivationKeysClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MultipleActivationKeysClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *MultipleActivationKeysClient) updateCreateRequest(ctx context.Context, resourceGroupName string, multipleActivationKeyName string, multipleActivationKey MultipleActivationKeyUpdate, options *MultipleActivationKeysClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.WindowsESU/multipleActivationKeys/{multipleActivationKeyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if multipleActivationKeyName == "" {
		return nil, errors.New("parameter multipleActivationKeyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{multipleActivationKeyName}", url.PathEscape(multipleActivationKeyName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-09-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, multipleActivationKey)
}

// updateHandleResponse handles the Update response.
func (client *MultipleActivationKeysClient) updateHandleResponse(resp *http.Response) (MultipleActivationKeysClientUpdateResponse, error) {
	result := MultipleActivationKeysClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MultipleActivationKey); err != nil {
		return MultipleActivationKeysClientUpdateResponse{}, err
	}
	return result, nil
}