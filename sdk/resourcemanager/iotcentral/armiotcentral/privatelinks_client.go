//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armiotcentral

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

// PrivateLinksClient contains the methods for the PrivateLinks group.
// Don't use this type directly, use NewPrivateLinksClient() instead.
type PrivateLinksClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPrivateLinksClient creates a new instance of PrivateLinksClient with the specified values.
// subscriptionID - The subscription identifier.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPrivateLinksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateLinksClient, error) {
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
	client := &PrivateLinksClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Get a private link resource of a IoT Central Application.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
// resourceGroupName - The name of the resource group that contains the IoT Central application.
// resourceName - The ARM resource name of the IoT Central application.
// groupID - The private link resource name.
// options - PrivateLinksClientGetOptions contains the optional parameters for the PrivateLinksClient.Get method.
func (client *PrivateLinksClient) Get(ctx context.Context, resourceGroupName string, resourceName string, groupID string, options *PrivateLinksClientGetOptions) (PrivateLinksClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, groupID, options)
	if err != nil {
		return PrivateLinksClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateLinksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateLinksClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateLinksClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, groupID string, options *PrivateLinksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTCentral/iotApps/{resourceName}/privateLinkResources/{groupId}"
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
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateLinksClient) getHandleResponse(resp *http.Response) (PrivateLinksClientGetResponse, error) {
	result := PrivateLinksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResource); err != nil {
		return PrivateLinksClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get all private link resources of a IoT Central Application.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
// resourceGroupName - The name of the resource group that contains the IoT Central application.
// resourceName - The ARM resource name of the IoT Central application.
// options - PrivateLinksClientListOptions contains the optional parameters for the PrivateLinksClient.List method.
func (client *PrivateLinksClient) NewListPager(resourceGroupName string, resourceName string, options *PrivateLinksClientListOptions) *runtime.Pager[PrivateLinksClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateLinksClientListResponse]{
		More: func(page PrivateLinksClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *PrivateLinksClientListResponse) (PrivateLinksClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceGroupName, resourceName, options)
			if err != nil {
				return PrivateLinksClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PrivateLinksClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PrivateLinksClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *PrivateLinksClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *PrivateLinksClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTCentral/iotApps/{resourceName}/privateLinkResources"
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
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PrivateLinksClient) listHandleResponse(resp *http.Response) (PrivateLinksClientListResponse, error) {
	result := PrivateLinksClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResourceListResult); err != nil {
		return PrivateLinksClientListResponse{}, err
	}
	return result, nil
}