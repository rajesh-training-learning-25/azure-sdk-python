//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappcontainers

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

// ContainerAppsAuthConfigsClient contains the methods for the ContainerAppsAuthConfigs group.
// Don't use this type directly, use NewContainerAppsAuthConfigsClient() instead.
type ContainerAppsAuthConfigsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewContainerAppsAuthConfigsClient creates a new instance of ContainerAppsAuthConfigsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewContainerAppsAuthConfigsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ContainerAppsAuthConfigsClient, error) {
	cl, err := arm.NewClient(moduleName+".ContainerAppsAuthConfigsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ContainerAppsAuthConfigsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update the AuthConfig for a Container App.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerAppName - Name of the Container App.
//   - authConfigName - Name of the Container App AuthConfig.
//   - authConfigEnvelope - Properties used to create a Container App AuthConfig
//   - options - ContainerAppsAuthConfigsClientCreateOrUpdateOptions contains the optional parameters for the ContainerAppsAuthConfigsClient.CreateOrUpdate
//     method.
func (client *ContainerAppsAuthConfigsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, containerAppName string, authConfigName string, authConfigEnvelope AuthConfig, options *ContainerAppsAuthConfigsClientCreateOrUpdateOptions) (ContainerAppsAuthConfigsClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, containerAppName, authConfigName, authConfigEnvelope, options)
	if err != nil {
		return ContainerAppsAuthConfigsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContainerAppsAuthConfigsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ContainerAppsAuthConfigsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ContainerAppsAuthConfigsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, authConfigName string, authConfigEnvelope AuthConfig, options *ContainerAppsAuthConfigsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/authConfigs/{authConfigName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	if authConfigName == "" {
		return nil, errors.New("parameter authConfigName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authConfigName}", url.PathEscape(authConfigName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, authConfigEnvelope); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ContainerAppsAuthConfigsClient) createOrUpdateHandleResponse(resp *http.Response) (ContainerAppsAuthConfigsClientCreateOrUpdateResponse, error) {
	result := ContainerAppsAuthConfigsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AuthConfig); err != nil {
		return ContainerAppsAuthConfigsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a Container App AuthConfig.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerAppName - Name of the Container App.
//   - authConfigName - Name of the Container App AuthConfig.
//   - options - ContainerAppsAuthConfigsClientDeleteOptions contains the optional parameters for the ContainerAppsAuthConfigsClient.Delete
//     method.
func (client *ContainerAppsAuthConfigsClient) Delete(ctx context.Context, resourceGroupName string, containerAppName string, authConfigName string, options *ContainerAppsAuthConfigsClientDeleteOptions) (ContainerAppsAuthConfigsClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, containerAppName, authConfigName, options)
	if err != nil {
		return ContainerAppsAuthConfigsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContainerAppsAuthConfigsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ContainerAppsAuthConfigsClientDeleteResponse{}, err
	}
	return ContainerAppsAuthConfigsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ContainerAppsAuthConfigsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, authConfigName string, options *ContainerAppsAuthConfigsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/authConfigs/{authConfigName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	if authConfigName == "" {
		return nil, errors.New("parameter authConfigName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authConfigName}", url.PathEscape(authConfigName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a AuthConfig of a Container App.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerAppName - Name of the Container App.
//   - authConfigName - Name of the Container App AuthConfig.
//   - options - ContainerAppsAuthConfigsClientGetOptions contains the optional parameters for the ContainerAppsAuthConfigsClient.Get
//     method.
func (client *ContainerAppsAuthConfigsClient) Get(ctx context.Context, resourceGroupName string, containerAppName string, authConfigName string, options *ContainerAppsAuthConfigsClientGetOptions) (ContainerAppsAuthConfigsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, containerAppName, authConfigName, options)
	if err != nil {
		return ContainerAppsAuthConfigsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContainerAppsAuthConfigsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ContainerAppsAuthConfigsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ContainerAppsAuthConfigsClient) getCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, authConfigName string, options *ContainerAppsAuthConfigsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/authConfigs/{authConfigName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	if authConfigName == "" {
		return nil, errors.New("parameter authConfigName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authConfigName}", url.PathEscape(authConfigName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ContainerAppsAuthConfigsClient) getHandleResponse(resp *http.Response) (ContainerAppsAuthConfigsClientGetResponse, error) {
	result := ContainerAppsAuthConfigsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AuthConfig); err != nil {
		return ContainerAppsAuthConfigsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByContainerAppPager - Get the Container App AuthConfigs in a given resource group.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerAppName - Name of the Container App.
//   - options - ContainerAppsAuthConfigsClientListByContainerAppOptions contains the optional parameters for the ContainerAppsAuthConfigsClient.NewListByContainerAppPager
//     method.
func (client *ContainerAppsAuthConfigsClient) NewListByContainerAppPager(resourceGroupName string, containerAppName string, options *ContainerAppsAuthConfigsClientListByContainerAppOptions) *runtime.Pager[ContainerAppsAuthConfigsClientListByContainerAppResponse] {
	return runtime.NewPager(runtime.PagingHandler[ContainerAppsAuthConfigsClientListByContainerAppResponse]{
		More: func(page ContainerAppsAuthConfigsClientListByContainerAppResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ContainerAppsAuthConfigsClientListByContainerAppResponse) (ContainerAppsAuthConfigsClientListByContainerAppResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByContainerAppCreateRequest(ctx, resourceGroupName, containerAppName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ContainerAppsAuthConfigsClientListByContainerAppResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ContainerAppsAuthConfigsClientListByContainerAppResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ContainerAppsAuthConfigsClientListByContainerAppResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByContainerAppHandleResponse(resp)
		},
	})
}

// listByContainerAppCreateRequest creates the ListByContainerApp request.
func (client *ContainerAppsAuthConfigsClient) listByContainerAppCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, options *ContainerAppsAuthConfigsClientListByContainerAppOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/authConfigs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByContainerAppHandleResponse handles the ListByContainerApp response.
func (client *ContainerAppsAuthConfigsClient) listByContainerAppHandleResponse(resp *http.Response) (ContainerAppsAuthConfigsClientListByContainerAppResponse, error) {
	result := ContainerAppsAuthConfigsClientListByContainerAppResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AuthConfigCollection); err != nil {
		return ContainerAppsAuthConfigsClientListByContainerAppResponse{}, err
	}
	return result, nil
}
