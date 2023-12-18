//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armscvmm

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// VmmServersClient contains the methods for the VmmServers group.
// Don't use this type directly, use NewVmmServersClient() instead.
type VmmServersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVmmServersClient creates a new instance of VmmServersClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVmmServersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VmmServersClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VmmServersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Onboards the SCVMM fabric as an Azure VmmServer resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vmmServerName - Name of the VMMServer.
//   - body - Request payload.
//   - options - VmmServersClientBeginCreateOrUpdateOptions contains the optional parameters for the VmmServersClient.BeginCreateOrUpdate
//     method.
func (client *VmmServersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vmmServerName string, body VMMServer, options *VmmServersClientBeginCreateOrUpdateOptions) (*runtime.Poller[VmmServersClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, vmmServerName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VmmServersClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VmmServersClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Onboards the SCVMM fabric as an Azure VmmServer resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
func (client *VmmServersClient) createOrUpdate(ctx context.Context, resourceGroupName string, vmmServerName string, body VMMServer, options *VmmServersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "VmmServersClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, vmmServerName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VmmServersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vmmServerName string, body VMMServer, options *VmmServersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ScVmm/vmmServers/{vmmServerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmmServerName == "" {
		return nil, errors.New("parameter vmmServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmmServerName}", url.PathEscape(vmmServerName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-07")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Removes the SCVMM fabric from Azure.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vmmServerName - Name of the VMMServer.
//   - options - VmmServersClientBeginDeleteOptions contains the optional parameters for the VmmServersClient.BeginDelete method.
func (client *VmmServersClient) BeginDelete(ctx context.Context, resourceGroupName string, vmmServerName string, options *VmmServersClientBeginDeleteOptions) (*runtime.Poller[VmmServersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, vmmServerName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VmmServersClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VmmServersClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Removes the SCVMM fabric from Azure.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
func (client *VmmServersClient) deleteOperation(ctx context.Context, resourceGroupName string, vmmServerName string, options *VmmServersClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "VmmServersClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vmmServerName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
<<<<<<< HEAD
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
=======
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
>>>>>>> 2621632e48ea508e16ce568001402f92fca4afa0
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VmmServersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vmmServerName string, options *VmmServersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ScVmm/vmmServers/{vmmServerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmmServerName == "" {
		return nil, errors.New("parameter vmmServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmmServerName}", url.PathEscape(vmmServerName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-07")
	if options != nil && options.Force != nil {
		reqQP.Set("force", string(*options.Force))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Implements VMMServer GET method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vmmServerName - Name of the VMMServer.
//   - options - VmmServersClientGetOptions contains the optional parameters for the VmmServersClient.Get method.
func (client *VmmServersClient) Get(ctx context.Context, resourceGroupName string, vmmServerName string, options *VmmServersClientGetOptions) (VmmServersClientGetResponse, error) {
	var err error
	const operationName = "VmmServersClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, vmmServerName, options)
	if err != nil {
		return VmmServersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VmmServersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VmmServersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *VmmServersClient) getCreateRequest(ctx context.Context, resourceGroupName string, vmmServerName string, options *VmmServersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ScVmm/vmmServers/{vmmServerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmmServerName == "" {
		return nil, errors.New("parameter vmmServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmmServerName}", url.PathEscape(vmmServerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-07")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VmmServersClient) getHandleResponse(resp *http.Response) (VmmServersClientGetResponse, error) {
	result := VmmServersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VMMServer); err != nil {
		return VmmServersClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List of VmmServers in a resource group.
//
// Generated from API version 2023-10-07
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - VmmServersClientListByResourceGroupOptions contains the optional parameters for the VmmServersClient.NewListByResourceGroupPager
//     method.
func (client *VmmServersClient) NewListByResourceGroupPager(resourceGroupName string, options *VmmServersClientListByResourceGroupOptions) *runtime.Pager[VmmServersClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[VmmServersClientListByResourceGroupResponse]{
		More: func(page VmmServersClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VmmServersClientListByResourceGroupResponse) (VmmServersClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "VmmServersClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return VmmServersClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *VmmServersClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *VmmServersClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ScVmm/vmmServers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-07")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *VmmServersClient) listByResourceGroupHandleResponse(resp *http.Response) (VmmServersClientListByResourceGroupResponse, error) {
	result := VmmServersClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VMMServerListResult); err != nil {
		return VmmServersClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List of VmmServers in a subscription.
//
// Generated from API version 2023-10-07
//   - options - VmmServersClientListBySubscriptionOptions contains the optional parameters for the VmmServersClient.NewListBySubscriptionPager
//     method.
func (client *VmmServersClient) NewListBySubscriptionPager(options *VmmServersClientListBySubscriptionOptions) *runtime.Pager[VmmServersClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[VmmServersClientListBySubscriptionResponse]{
		More: func(page VmmServersClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VmmServersClientListBySubscriptionResponse) (VmmServersClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "VmmServersClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return VmmServersClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *VmmServersClient) listBySubscriptionCreateRequest(ctx context.Context, options *VmmServersClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ScVmm/vmmServers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-07")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *VmmServersClient) listBySubscriptionHandleResponse(resp *http.Response) (VmmServersClientListBySubscriptionResponse, error) {
	result := VmmServersClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VMMServerListResult); err != nil {
		return VmmServersClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates the VmmServers resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vmmServerName - Name of the VMMServer.
//   - body - VmmServers patch payload.
//   - options - VmmServersClientBeginUpdateOptions contains the optional parameters for the VmmServersClient.BeginUpdate method.
func (client *VmmServersClient) BeginUpdate(ctx context.Context, resourceGroupName string, vmmServerName string, body ResourcePatch, options *VmmServersClientBeginUpdateOptions) (*runtime.Poller[VmmServersClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, vmmServerName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VmmServersClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VmmServersClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Updates the VmmServers resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
func (client *VmmServersClient) update(ctx context.Context, resourceGroupName string, vmmServerName string, body ResourcePatch, options *VmmServersClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "VmmServersClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, vmmServerName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
<<<<<<< HEAD
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
=======
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
>>>>>>> 2621632e48ea508e16ce568001402f92fca4afa0
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *VmmServersClient) updateCreateRequest(ctx context.Context, resourceGroupName string, vmmServerName string, body ResourcePatch, options *VmmServersClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ScVmm/vmmServers/{vmmServerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if vmmServerName == "" {
		return nil, errors.New("parameter vmmServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmmServerName}", url.PathEscape(vmmServerName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-07")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
