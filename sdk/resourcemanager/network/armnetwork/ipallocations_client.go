//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

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

// IPAllocationsClient contains the methods for the IPAllocations group.
// Don't use this type directly, use NewIPAllocationsClient() instead.
type IPAllocationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewIPAllocationsClient creates a new instance of IPAllocationsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewIPAllocationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IPAllocationsClient, error) {
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
	client := &IPAllocationsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates an IpAllocation in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group.
// ipAllocationName - The name of the IpAllocation.
// parameters - Parameters supplied to the create or update virtual network operation.
// options - IPAllocationsClientBeginCreateOrUpdateOptions contains the optional parameters for the IPAllocationsClient.BeginCreateOrUpdate
// method.
func (client *IPAllocationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, ipAllocationName string, parameters IPAllocation, options *IPAllocationsClientBeginCreateOrUpdateOptions) (*runtime.Poller[IPAllocationsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, ipAllocationName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[IPAllocationsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[IPAllocationsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates an IpAllocation in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
func (client *IPAllocationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, ipAllocationName string, parameters IPAllocation, options *IPAllocationsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, ipAllocationName, parameters, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IPAllocationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, ipAllocationName string, parameters IPAllocation, options *IPAllocationsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/IpAllocations/{ipAllocationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ipAllocationName == "" {
		return nil, errors.New("parameter ipAllocationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipAllocationName}", url.PathEscape(ipAllocationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified IpAllocation.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group.
// ipAllocationName - The name of the IpAllocation.
// options - IPAllocationsClientBeginDeleteOptions contains the optional parameters for the IPAllocationsClient.BeginDelete
// method.
func (client *IPAllocationsClient) BeginDelete(ctx context.Context, resourceGroupName string, ipAllocationName string, options *IPAllocationsClientBeginDeleteOptions) (*runtime.Poller[IPAllocationsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, ipAllocationName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[IPAllocationsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[IPAllocationsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified IpAllocation.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
func (client *IPAllocationsClient) deleteOperation(ctx context.Context, resourceGroupName string, ipAllocationName string, options *IPAllocationsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, ipAllocationName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IPAllocationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, ipAllocationName string, options *IPAllocationsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/IpAllocations/{ipAllocationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ipAllocationName == "" {
		return nil, errors.New("parameter ipAllocationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipAllocationName}", url.PathEscape(ipAllocationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified IpAllocation by resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group.
// ipAllocationName - The name of the IpAllocation.
// options - IPAllocationsClientGetOptions contains the optional parameters for the IPAllocationsClient.Get method.
func (client *IPAllocationsClient) Get(ctx context.Context, resourceGroupName string, ipAllocationName string, options *IPAllocationsClientGetOptions) (IPAllocationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, ipAllocationName, options)
	if err != nil {
		return IPAllocationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IPAllocationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IPAllocationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IPAllocationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, ipAllocationName string, options *IPAllocationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/IpAllocations/{ipAllocationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ipAllocationName == "" {
		return nil, errors.New("parameter ipAllocationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipAllocationName}", url.PathEscape(ipAllocationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IPAllocationsClient) getHandleResponse(resp *http.Response) (IPAllocationsClientGetResponse, error) {
	result := IPAllocationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPAllocation); err != nil {
		return IPAllocationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all IpAllocations in a subscription.
// Generated from API version 2022-05-01
// options - IPAllocationsClientListOptions contains the optional parameters for the IPAllocationsClient.List method.
func (client *IPAllocationsClient) NewListPager(options *IPAllocationsClientListOptions) *runtime.Pager[IPAllocationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[IPAllocationsClientListResponse]{
		More: func(page IPAllocationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IPAllocationsClientListResponse) (IPAllocationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return IPAllocationsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return IPAllocationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return IPAllocationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *IPAllocationsClient) listCreateRequest(ctx context.Context, options *IPAllocationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/IpAllocations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IPAllocationsClient) listHandleResponse(resp *http.Response) (IPAllocationsClientListResponse, error) {
	result := IPAllocationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPAllocationListResult); err != nil {
		return IPAllocationsClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets all IpAllocations in a resource group.
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group.
// options - IPAllocationsClientListByResourceGroupOptions contains the optional parameters for the IPAllocationsClient.ListByResourceGroup
// method.
func (client *IPAllocationsClient) NewListByResourceGroupPager(resourceGroupName string, options *IPAllocationsClientListByResourceGroupOptions) *runtime.Pager[IPAllocationsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[IPAllocationsClientListByResourceGroupResponse]{
		More: func(page IPAllocationsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IPAllocationsClientListByResourceGroupResponse) (IPAllocationsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return IPAllocationsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return IPAllocationsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return IPAllocationsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *IPAllocationsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *IPAllocationsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/IpAllocations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *IPAllocationsClient) listByResourceGroupHandleResponse(resp *http.Response) (IPAllocationsClientListByResourceGroupResponse, error) {
	result := IPAllocationsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPAllocationListResult); err != nil {
		return IPAllocationsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates a IpAllocation tags.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group.
// ipAllocationName - The name of the IpAllocation.
// parameters - Parameters supplied to update IpAllocation tags.
// options - IPAllocationsClientUpdateTagsOptions contains the optional parameters for the IPAllocationsClient.UpdateTags
// method.
func (client *IPAllocationsClient) UpdateTags(ctx context.Context, resourceGroupName string, ipAllocationName string, parameters TagsObject, options *IPAllocationsClientUpdateTagsOptions) (IPAllocationsClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, ipAllocationName, parameters, options)
	if err != nil {
		return IPAllocationsClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IPAllocationsClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IPAllocationsClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *IPAllocationsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, ipAllocationName string, parameters TagsObject, options *IPAllocationsClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/IpAllocations/{ipAllocationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ipAllocationName == "" {
		return nil, errors.New("parameter ipAllocationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipAllocationName}", url.PathEscape(ipAllocationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *IPAllocationsClient) updateTagsHandleResponse(resp *http.Response) (IPAllocationsClientUpdateTagsResponse, error) {
	result := IPAllocationsClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPAllocation); err != nil {
		return IPAllocationsClientUpdateTagsResponse{}, err
	}
	return result, nil
}