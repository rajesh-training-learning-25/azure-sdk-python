//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armredhatopenshift

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

// SyncSetsClient contains the methods for the SyncSets group.
// Don't use this type directly, use NewSyncSetsClient() instead.
type SyncSetsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSyncSetsClient creates a new instance of SyncSetsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSyncSetsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SyncSetsClient, error) {
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
	client := &SyncSetsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - The operation returns properties of a SyncSet.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-04
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the OpenShift cluster resource.
// childResourceName - The name of the SyncSet resource.
// parameters - The SyncSet resource.
// options - SyncSetsClientCreateOrUpdateOptions contains the optional parameters for the SyncSetsClient.CreateOrUpdate method.
func (client *SyncSetsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, parameters SyncSet, options *SyncSetsClientCreateOrUpdateOptions) (SyncSetsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, childResourceName, parameters, options)
	if err != nil {
		return SyncSetsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SyncSetsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return SyncSetsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SyncSetsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, parameters SyncSet, options *SyncSetsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/openshiftclusters/{resourceName}/syncSet/{childResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if childResourceName == "" {
		return nil, errors.New("parameter childResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{childResourceName}", url.PathEscape(childResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SyncSetsClient) createOrUpdateHandleResponse(resp *http.Response) (SyncSetsClientCreateOrUpdateResponse, error) {
	result := SyncSetsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SyncSet); err != nil {
		return SyncSetsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - The operation returns nothing.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-04
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the OpenShift cluster resource.
// childResourceName - The name of the SyncSet resource.
// options - SyncSetsClientDeleteOptions contains the optional parameters for the SyncSetsClient.Delete method.
func (client *SyncSetsClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, options *SyncSetsClientDeleteOptions) (SyncSetsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, childResourceName, options)
	if err != nil {
		return SyncSetsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SyncSetsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return SyncSetsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return SyncSetsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SyncSetsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, options *SyncSetsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/openshiftclusters/{resourceName}/syncSet/{childResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if childResourceName == "" {
		return nil, errors.New("parameter childResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{childResourceName}", url.PathEscape(childResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - The operation returns properties of a SyncSet.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-04
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the OpenShift cluster resource.
// childResourceName - The name of the SyncSet resource.
// options - SyncSetsClientGetOptions contains the optional parameters for the SyncSetsClient.Get method.
func (client *SyncSetsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, options *SyncSetsClientGetOptions) (SyncSetsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, childResourceName, options)
	if err != nil {
		return SyncSetsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SyncSetsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SyncSetsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SyncSetsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, options *SyncSetsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/openshiftclusters/{resourceName}/syncSet/{childResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if childResourceName == "" {
		return nil, errors.New("parameter childResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{childResourceName}", url.PathEscape(childResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SyncSetsClient) getHandleResponse(resp *http.Response) (SyncSetsClientGetResponse, error) {
	result := SyncSetsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SyncSet); err != nil {
		return SyncSetsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - The operation returns properties of each SyncSet.
// Generated from API version 2022-09-04
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the OpenShift cluster resource.
// options - SyncSetsClientListOptions contains the optional parameters for the SyncSetsClient.List method.
func (client *SyncSetsClient) NewListPager(resourceGroupName string, resourceName string, options *SyncSetsClientListOptions) *runtime.Pager[SyncSetsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SyncSetsClientListResponse]{
		More: func(page SyncSetsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SyncSetsClientListResponse) (SyncSetsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, resourceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SyncSetsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SyncSetsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SyncSetsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SyncSetsClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *SyncSetsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/openShiftCluster/{resourceName}/syncSets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SyncSetsClient) listHandleResponse(resp *http.Response) (SyncSetsClientListResponse, error) {
	result := SyncSetsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SyncSetList); err != nil {
		return SyncSetsClientListResponse{}, err
	}
	return result, nil
}

// Update - The operation returns properties of a SyncSet.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-04
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the OpenShift cluster resource.
// childResourceName - The name of the SyncSet resource.
// parameters - The SyncSet resource.
// options - SyncSetsClientUpdateOptions contains the optional parameters for the SyncSetsClient.Update method.
func (client *SyncSetsClient) Update(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, parameters SyncSetUpdate, options *SyncSetsClientUpdateOptions) (SyncSetsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, childResourceName, parameters, options)
	if err != nil {
		return SyncSetsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SyncSetsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SyncSetsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *SyncSetsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, parameters SyncSetUpdate, options *SyncSetsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/openshiftclusters/{resourceName}/syncSet/{childResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if childResourceName == "" {
		return nil, errors.New("parameter childResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{childResourceName}", url.PathEscape(childResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *SyncSetsClient) updateHandleResponse(resp *http.Response) (SyncSetsClientUpdateResponse, error) {
	result := SyncSetsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SyncSet); err != nil {
		return SyncSetsClientUpdateResponse{}, err
	}
	return result, nil
}
