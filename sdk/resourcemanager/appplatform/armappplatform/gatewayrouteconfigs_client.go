//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappplatform

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

// GatewayRouteConfigsClient contains the methods for the GatewayRouteConfigs group.
// Don't use this type directly, use NewGatewayRouteConfigsClient() instead.
type GatewayRouteConfigsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewGatewayRouteConfigsClient creates a new instance of GatewayRouteConfigsClient with the specified values.
// subscriptionID - Gets subscription ID which uniquely identify the Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewGatewayRouteConfigsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GatewayRouteConfigsClient, error) {
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
	client := &GatewayRouteConfigsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create the default Spring Cloud Gateway route configs or update the existing Spring Cloud Gateway
// route configs.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// gatewayName - The name of Spring Cloud Gateway.
// routeConfigName - The name of the Spring Cloud Gateway route config.
// gatewayRouteConfigResource - The Spring Cloud Gateway route config for the create or update operation
// options - GatewayRouteConfigsClientBeginCreateOrUpdateOptions contains the optional parameters for the GatewayRouteConfigsClient.BeginCreateOrUpdate
// method.
func (client *GatewayRouteConfigsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, routeConfigName string, gatewayRouteConfigResource GatewayRouteConfigResource, options *GatewayRouteConfigsClientBeginCreateOrUpdateOptions) (*runtime.Poller[GatewayRouteConfigsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serviceName, gatewayName, routeConfigName, gatewayRouteConfigResource, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[GatewayRouteConfigsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[GatewayRouteConfigsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create the default Spring Cloud Gateway route configs or update the existing Spring Cloud Gateway route
// configs.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01-preview
func (client *GatewayRouteConfigsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, routeConfigName string, gatewayRouteConfigResource GatewayRouteConfigResource, options *GatewayRouteConfigsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, gatewayName, routeConfigName, gatewayRouteConfigResource, options)
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
func (client *GatewayRouteConfigsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, routeConfigName string, gatewayRouteConfigResource GatewayRouteConfigResource, options *GatewayRouteConfigsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/gateways/{gatewayName}/routeConfigs/{routeConfigName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if routeConfigName == "" {
		return nil, errors.New("parameter routeConfigName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeConfigName}", url.PathEscape(routeConfigName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, gatewayRouteConfigResource)
}

// BeginDelete - Delete the Spring Cloud Gateway route config.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// gatewayName - The name of Spring Cloud Gateway.
// routeConfigName - The name of the Spring Cloud Gateway route config.
// options - GatewayRouteConfigsClientBeginDeleteOptions contains the optional parameters for the GatewayRouteConfigsClient.BeginDelete
// method.
func (client *GatewayRouteConfigsClient) BeginDelete(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, routeConfigName string, options *GatewayRouteConfigsClientBeginDeleteOptions) (*runtime.Poller[GatewayRouteConfigsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serviceName, gatewayName, routeConfigName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[GatewayRouteConfigsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[GatewayRouteConfigsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete the Spring Cloud Gateway route config.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01-preview
func (client *GatewayRouteConfigsClient) deleteOperation(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, routeConfigName string, options *GatewayRouteConfigsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, gatewayName, routeConfigName, options)
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
func (client *GatewayRouteConfigsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, routeConfigName string, options *GatewayRouteConfigsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/gateways/{gatewayName}/routeConfigs/{routeConfigName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if routeConfigName == "" {
		return nil, errors.New("parameter routeConfigName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeConfigName}", url.PathEscape(routeConfigName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the Spring Cloud Gateway route configs.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// gatewayName - The name of Spring Cloud Gateway.
// routeConfigName - The name of the Spring Cloud Gateway route config.
// options - GatewayRouteConfigsClientGetOptions contains the optional parameters for the GatewayRouteConfigsClient.Get method.
func (client *GatewayRouteConfigsClient) Get(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, routeConfigName string, options *GatewayRouteConfigsClientGetOptions) (GatewayRouteConfigsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, gatewayName, routeConfigName, options)
	if err != nil {
		return GatewayRouteConfigsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GatewayRouteConfigsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GatewayRouteConfigsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *GatewayRouteConfigsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, routeConfigName string, options *GatewayRouteConfigsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/gateways/{gatewayName}/routeConfigs/{routeConfigName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if routeConfigName == "" {
		return nil, errors.New("parameter routeConfigName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeConfigName}", url.PathEscape(routeConfigName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GatewayRouteConfigsClient) getHandleResponse(resp *http.Response) (GatewayRouteConfigsClientGetResponse, error) {
	result := GatewayRouteConfigsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GatewayRouteConfigResource); err != nil {
		return GatewayRouteConfigsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Handle requests to list all Spring Cloud Gateway route configs.
// Generated from API version 2022-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// gatewayName - The name of Spring Cloud Gateway.
// options - GatewayRouteConfigsClientListOptions contains the optional parameters for the GatewayRouteConfigsClient.List
// method.
func (client *GatewayRouteConfigsClient) NewListPager(resourceGroupName string, serviceName string, gatewayName string, options *GatewayRouteConfigsClientListOptions) *runtime.Pager[GatewayRouteConfigsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[GatewayRouteConfigsClientListResponse]{
		More: func(page GatewayRouteConfigsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GatewayRouteConfigsClientListResponse) (GatewayRouteConfigsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, serviceName, gatewayName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return GatewayRouteConfigsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return GatewayRouteConfigsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return GatewayRouteConfigsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *GatewayRouteConfigsClient) listCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, options *GatewayRouteConfigsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/gateways/{gatewayName}/routeConfigs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *GatewayRouteConfigsClient) listHandleResponse(resp *http.Response) (GatewayRouteConfigsClientListResponse, error) {
	result := GatewayRouteConfigsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GatewayRouteConfigResourceCollection); err != nil {
		return GatewayRouteConfigsClientListResponse{}, err
	}
	return result, nil
}