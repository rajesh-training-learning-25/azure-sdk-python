//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurestackhci

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

// LogicalNetworksClient contains the methods for the LogicalNetworks group.
// Don't use this type directly, use NewLogicalNetworksClient() instead.
type LogicalNetworksClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewLogicalNetworksClient creates a new instance of LogicalNetworksClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLogicalNetworksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LogicalNetworksClient, error) {
	cl, err := arm.NewClient(moduleName+".LogicalNetworksClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &LogicalNetworksClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - The operation to create or update a logical network. Please note some properties can be set only
// during logical network creation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - logicalNetworkName - Name of the logical network
//   - options - LogicalNetworksClientBeginCreateOrUpdateOptions contains the optional parameters for the LogicalNetworksClient.BeginCreateOrUpdate
//     method.
func (client *LogicalNetworksClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, logicalNetworkName string, logicalNetworks LogicalNetworks, options *LogicalNetworksClientBeginCreateOrUpdateOptions) (*runtime.Poller[LogicalNetworksClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, logicalNetworkName, logicalNetworks, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LogicalNetworksClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[LogicalNetworksClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - The operation to create or update a logical network. Please note some properties can be set only during
// logical network creation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
func (client *LogicalNetworksClient) createOrUpdate(ctx context.Context, resourceGroupName string, logicalNetworkName string, logicalNetworks LogicalNetworks, options *LogicalNetworksClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, logicalNetworkName, logicalNetworks, options)
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
func (client *LogicalNetworksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, logicalNetworkName string, logicalNetworks LogicalNetworks, options *LogicalNetworksClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/logicalNetworks/{logicalNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if logicalNetworkName == "" {
		return nil, errors.New("parameter logicalNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{logicalNetworkName}", url.PathEscape(logicalNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, logicalNetworks); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - The operation to delete a logical network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - logicalNetworkName - Name of the logical network
//   - options - LogicalNetworksClientBeginDeleteOptions contains the optional parameters for the LogicalNetworksClient.BeginDelete
//     method.
func (client *LogicalNetworksClient) BeginDelete(ctx context.Context, resourceGroupName string, logicalNetworkName string, options *LogicalNetworksClientBeginDeleteOptions) (*runtime.Poller[LogicalNetworksClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, logicalNetworkName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LogicalNetworksClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[LogicalNetworksClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - The operation to delete a logical network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
func (client *LogicalNetworksClient) deleteOperation(ctx context.Context, resourceGroupName string, logicalNetworkName string, options *LogicalNetworksClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, logicalNetworkName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LogicalNetworksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, logicalNetworkName string, options *LogicalNetworksClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/logicalNetworks/{logicalNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if logicalNetworkName == "" {
		return nil, errors.New("parameter logicalNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{logicalNetworkName}", url.PathEscape(logicalNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get -
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - logicalNetworkName - Name of the logical network
//   - options - LogicalNetworksClientGetOptions contains the optional parameters for the LogicalNetworksClient.Get method.
func (client *LogicalNetworksClient) Get(ctx context.Context, resourceGroupName string, logicalNetworkName string, options *LogicalNetworksClientGetOptions) (LogicalNetworksClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, logicalNetworkName, options)
	if err != nil {
		return LogicalNetworksClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LogicalNetworksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LogicalNetworksClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *LogicalNetworksClient) getCreateRequest(ctx context.Context, resourceGroupName string, logicalNetworkName string, options *LogicalNetworksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/logicalNetworks/{logicalNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if logicalNetworkName == "" {
		return nil, errors.New("parameter logicalNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{logicalNetworkName}", url.PathEscape(logicalNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LogicalNetworksClient) getHandleResponse(resp *http.Response) (LogicalNetworksClientGetResponse, error) {
	result := LogicalNetworksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LogicalNetworks); err != nil {
		return LogicalNetworksClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all of the logical networks in the specified resource group. Use the nextLink property in the response
// to get the next page of logical networks.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - LogicalNetworksClientListOptions contains the optional parameters for the LogicalNetworksClient.NewListPager
//     method.
func (client *LogicalNetworksClient) NewListPager(resourceGroupName string, options *LogicalNetworksClientListOptions) *runtime.Pager[LogicalNetworksClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[LogicalNetworksClientListResponse]{
		More: func(page LogicalNetworksClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LogicalNetworksClientListResponse) (LogicalNetworksClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return LogicalNetworksClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return LogicalNetworksClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LogicalNetworksClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *LogicalNetworksClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *LogicalNetworksClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/logicalNetworks"
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
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LogicalNetworksClient) listHandleResponse(resp *http.Response) (LogicalNetworksClientListResponse, error) {
	result := LogicalNetworksClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LogicalNetworksListResult); err != nil {
		return LogicalNetworksClientListResponse{}, err
	}
	return result, nil
}

// NewListAllPager - Lists all of the logical networks in the specified subscription. Use the nextLink property in the response
// to get the next page of logical networks.
//
// Generated from API version 2023-09-01-preview
//   - options - LogicalNetworksClientListAllOptions contains the optional parameters for the LogicalNetworksClient.NewListAllPager
//     method.
func (client *LogicalNetworksClient) NewListAllPager(options *LogicalNetworksClientListAllOptions) *runtime.Pager[LogicalNetworksClientListAllResponse] {
	return runtime.NewPager(runtime.PagingHandler[LogicalNetworksClientListAllResponse]{
		More: func(page LogicalNetworksClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LogicalNetworksClientListAllResponse) (LogicalNetworksClientListAllResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAllCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return LogicalNetworksClientListAllResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return LogicalNetworksClientListAllResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LogicalNetworksClientListAllResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAllHandleResponse(resp)
		},
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *LogicalNetworksClient) listAllCreateRequest(ctx context.Context, options *LogicalNetworksClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AzureStackHCI/logicalNetworks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *LogicalNetworksClient) listAllHandleResponse(resp *http.Response) (LogicalNetworksClientListAllResponse, error) {
	result := LogicalNetworksClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LogicalNetworksListResult); err != nil {
		return LogicalNetworksClientListAllResponse{}, err
	}
	return result, nil
}

// BeginUpdate - The operation to update a logical network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - logicalNetworkName - Name of the logical network
//   - options - LogicalNetworksClientBeginUpdateOptions contains the optional parameters for the LogicalNetworksClient.BeginUpdate
//     method.
func (client *LogicalNetworksClient) BeginUpdate(ctx context.Context, resourceGroupName string, logicalNetworkName string, logicalNetworks LogicalNetworksUpdateRequest, options *LogicalNetworksClientBeginUpdateOptions) (*runtime.Poller[LogicalNetworksClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, logicalNetworkName, logicalNetworks, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LogicalNetworksClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[LogicalNetworksClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - The operation to update a logical network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
func (client *LogicalNetworksClient) update(ctx context.Context, resourceGroupName string, logicalNetworkName string, logicalNetworks LogicalNetworksUpdateRequest, options *LogicalNetworksClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, logicalNetworkName, logicalNetworks, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *LogicalNetworksClient) updateCreateRequest(ctx context.Context, resourceGroupName string, logicalNetworkName string, logicalNetworks LogicalNetworksUpdateRequest, options *LogicalNetworksClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/logicalNetworks/{logicalNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if logicalNetworkName == "" {
		return nil, errors.New("parameter logicalNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{logicalNetworkName}", url.PathEscape(logicalNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, logicalNetworks); err != nil {
		return nil, err
	}
	return req, nil
}
