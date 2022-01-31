//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapplicationinsights

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// FavoritesClient contains the methods for the Favorites group.
// Don't use this type directly, use NewFavoritesClient() instead.
type FavoritesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewFavoritesClient creates a new instance of FavoritesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewFavoritesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *FavoritesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &FavoritesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Add - Adds a new favorites to an Application Insights component.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// favoriteID - The Id of a specific favorite defined in the Application Insights component
// favoriteProperties - Properties that need to be specified to create a new favorite and add it to an Application Insights
// component.
// options - FavoritesClientAddOptions contains the optional parameters for the FavoritesClient.Add method.
func (client *FavoritesClient) Add(ctx context.Context, resourceGroupName string, resourceName string, favoriteID string, favoriteProperties ComponentFavorite, options *FavoritesClientAddOptions) (FavoritesClientAddResponse, error) {
	req, err := client.addCreateRequest(ctx, resourceGroupName, resourceName, favoriteID, favoriteProperties, options)
	if err != nil {
		return FavoritesClientAddResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FavoritesClientAddResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FavoritesClientAddResponse{}, runtime.NewResponseError(resp)
	}
	return client.addHandleResponse(resp)
}

// addCreateRequest creates the Add request.
func (client *FavoritesClient) addCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, favoriteID string, favoriteProperties ComponentFavorite, options *FavoritesClientAddOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/favorites/{favoriteId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if favoriteID == "" {
		return nil, errors.New("parameter favoriteID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{favoriteId}", url.PathEscape(favoriteID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, favoriteProperties)
}

// addHandleResponse handles the Add response.
func (client *FavoritesClient) addHandleResponse(resp *http.Response) (FavoritesClientAddResponse, error) {
	result := FavoritesClientAddResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentFavorite); err != nil {
		return FavoritesClientAddResponse{}, err
	}
	return result, nil
}

// Delete - Remove a favorite that is associated to an Application Insights component.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// favoriteID - The Id of a specific favorite defined in the Application Insights component
// options - FavoritesClientDeleteOptions contains the optional parameters for the FavoritesClient.Delete method.
func (client *FavoritesClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, favoriteID string, options *FavoritesClientDeleteOptions) (FavoritesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, favoriteID, options)
	if err != nil {
		return FavoritesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FavoritesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FavoritesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return FavoritesClientDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *FavoritesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, favoriteID string, options *FavoritesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/favorites/{favoriteId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if favoriteID == "" {
		return nil, errors.New("parameter favoriteID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{favoriteId}", url.PathEscape(favoriteID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Get a single favorite by its FavoriteId, defined within an Application Insights component.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// favoriteID - The Id of a specific favorite defined in the Application Insights component
// options - FavoritesClientGetOptions contains the optional parameters for the FavoritesClient.Get method.
func (client *FavoritesClient) Get(ctx context.Context, resourceGroupName string, resourceName string, favoriteID string, options *FavoritesClientGetOptions) (FavoritesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, favoriteID, options)
	if err != nil {
		return FavoritesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FavoritesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FavoritesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *FavoritesClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, favoriteID string, options *FavoritesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/favorites/{favoriteId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if favoriteID == "" {
		return nil, errors.New("parameter favoriteID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{favoriteId}", url.PathEscape(favoriteID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FavoritesClient) getHandleResponse(resp *http.Response) (FavoritesClientGetResponse, error) {
	result := FavoritesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentFavorite); err != nil {
		return FavoritesClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets a list of favorites defined within an Application Insights component.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// options - FavoritesClientListOptions contains the optional parameters for the FavoritesClient.List method.
func (client *FavoritesClient) List(ctx context.Context, resourceGroupName string, resourceName string, options *FavoritesClientListOptions) (FavoritesClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return FavoritesClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FavoritesClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FavoritesClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *FavoritesClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *FavoritesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/favorites"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	if options != nil && options.FavoriteType != nil {
		reqQP.Set("favoriteType", string(*options.FavoriteType))
	}
	if options != nil && options.SourceType != nil {
		reqQP.Set("sourceType", string(*options.SourceType))
	}
	if options != nil && options.CanFetchContent != nil {
		reqQP.Set("canFetchContent", strconv.FormatBool(*options.CanFetchContent))
	}
	if options != nil && options.Tags != nil {
		reqQP.Set("tags", strings.Join(options.Tags, ","))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *FavoritesClient) listHandleResponse(resp *http.Response) (FavoritesClientListResponse, error) {
	result := FavoritesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentFavoriteArray); err != nil {
		return FavoritesClientListResponse{}, err
	}
	return result, nil
}

// Update - Updates a favorite that has already been added to an Application Insights component.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// favoriteID - The Id of a specific favorite defined in the Application Insights component
// favoriteProperties - Properties that need to be specified to update the existing favorite.
// options - FavoritesClientUpdateOptions contains the optional parameters for the FavoritesClient.Update method.
func (client *FavoritesClient) Update(ctx context.Context, resourceGroupName string, resourceName string, favoriteID string, favoriteProperties ComponentFavorite, options *FavoritesClientUpdateOptions) (FavoritesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, favoriteID, favoriteProperties, options)
	if err != nil {
		return FavoritesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FavoritesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FavoritesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *FavoritesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, favoriteID string, favoriteProperties ComponentFavorite, options *FavoritesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/favorites/{favoriteId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if favoriteID == "" {
		return nil, errors.New("parameter favoriteID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{favoriteId}", url.PathEscape(favoriteID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, favoriteProperties)
}

// updateHandleResponse handles the Update response.
func (client *FavoritesClient) updateHandleResponse(resp *http.Response) (FavoritesClientUpdateResponse, error) {
	result := FavoritesClientUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentFavorite); err != nil {
		return FavoritesClientUpdateResponse{}, err
	}
	return result, nil
}
