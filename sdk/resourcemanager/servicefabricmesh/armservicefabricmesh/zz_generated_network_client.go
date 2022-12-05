//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicefabricmesh

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

// NetworkClient contains the methods for the Network group.
// Don't use this type directly, use NewNetworkClient() instead.
type NetworkClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewNetworkClient creates a new instance of NetworkClient with the specified values.
// subscriptionID - The customer subscription identifier
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewNetworkClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NetworkClient, error) {
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
	client := &NetworkClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Create - Creates a network resource with the specified name, description and properties. If a network resource with the
// same name exists, then it is updated with the specified description and properties.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-09-01-preview
// resourceGroupName - Azure resource group name
// networkResourceName - The identity of the network.
// networkResourceDescription - Description for creating a Network resource.
// options - NetworkClientCreateOptions contains the optional parameters for the NetworkClient.Create method.
func (client *NetworkClient) Create(ctx context.Context, resourceGroupName string, networkResourceName string, networkResourceDescription NetworkResourceDescription, options *NetworkClientCreateOptions) (NetworkClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, networkResourceName, networkResourceDescription, options)
	if err != nil {
		return NetworkClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NetworkClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return NetworkClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *NetworkClient) createCreateRequest(ctx context.Context, resourceGroupName string, networkResourceName string, networkResourceDescription NetworkResourceDescription, options *NetworkClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/networks/{networkResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkResourceName}", networkResourceName)
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, networkResourceDescription)
}

// createHandleResponse handles the Create response.
func (client *NetworkClient) createHandleResponse(resp *http.Response) (NetworkClientCreateResponse, error) {
	result := NetworkClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkResourceDescription); err != nil {
		return NetworkClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the network resource identified by the name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-09-01-preview
// resourceGroupName - Azure resource group name
// networkResourceName - The identity of the network.
// options - NetworkClientDeleteOptions contains the optional parameters for the NetworkClient.Delete method.
func (client *NetworkClient) Delete(ctx context.Context, resourceGroupName string, networkResourceName string, options *NetworkClientDeleteOptions) (NetworkClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkResourceName, options)
	if err != nil {
		return NetworkClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NetworkClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return NetworkClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return NetworkClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *NetworkClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkResourceName string, options *NetworkClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/networks/{networkResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkResourceName}", networkResourceName)
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the information about the network resource with the given name. The information include the description and
// other properties of the network.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-09-01-preview
// resourceGroupName - Azure resource group name
// networkResourceName - The identity of the network.
// options - NetworkClientGetOptions contains the optional parameters for the NetworkClient.Get method.
func (client *NetworkClient) Get(ctx context.Context, resourceGroupName string, networkResourceName string, options *NetworkClientGetOptions) (NetworkClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkResourceName, options)
	if err != nil {
		return NetworkClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NetworkClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NetworkClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *NetworkClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkResourceName string, options *NetworkClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/networks/{networkResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkResourceName}", networkResourceName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NetworkClient) getHandleResponse(resp *http.Response) (NetworkClientGetResponse, error) {
	result := NetworkClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkResourceDescription); err != nil {
		return NetworkClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets the information about all network resources in a given resource group. The information
// include the description and other properties of the Network.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-09-01-preview
// resourceGroupName - Azure resource group name
// options - NetworkClientListByResourceGroupOptions contains the optional parameters for the NetworkClient.ListByResourceGroup
// method.
func (client *NetworkClient) NewListByResourceGroupPager(resourceGroupName string, options *NetworkClientListByResourceGroupOptions) *runtime.Pager[NetworkClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[NetworkClientListByResourceGroupResponse]{
		More: func(page NetworkClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NetworkClientListByResourceGroupResponse) (NetworkClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return NetworkClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return NetworkClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return NetworkClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *NetworkClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *NetworkClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/networks"
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
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *NetworkClient) listByResourceGroupHandleResponse(resp *http.Response) (NetworkClientListByResourceGroupResponse, error) {
	result := NetworkClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkResourceDescriptionList); err != nil {
		return NetworkClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Gets the information about all network resources in a given resource group. The information
// include the description and other properties of the network.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-09-01-preview
// options - NetworkClientListBySubscriptionOptions contains the optional parameters for the NetworkClient.ListBySubscription
// method.
func (client *NetworkClient) NewListBySubscriptionPager(options *NetworkClientListBySubscriptionOptions) *runtime.Pager[NetworkClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[NetworkClientListBySubscriptionResponse]{
		More: func(page NetworkClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NetworkClientListBySubscriptionResponse) (NetworkClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return NetworkClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return NetworkClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return NetworkClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *NetworkClient) listBySubscriptionCreateRequest(ctx context.Context, options *NetworkClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabricMesh/networks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *NetworkClient) listBySubscriptionHandleResponse(resp *http.Response) (NetworkClientListBySubscriptionResponse, error) {
	result := NetworkClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkResourceDescriptionList); err != nil {
		return NetworkClientListBySubscriptionResponse{}, err
	}
	return result, nil
}