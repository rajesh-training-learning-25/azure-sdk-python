//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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

// ProductGroupClient contains the methods for the ProductGroup group.
// Don't use this type directly, use NewProductGroupClient() instead.
type ProductGroupClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewProductGroupClient creates a new instance of ProductGroupClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewProductGroupClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ProductGroupClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &ProductGroupClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// CheckEntityExists - Checks that Group entity specified by identifier is associated with the Product entity.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// productID - Product identifier. Must be unique in the current API Management service instance.
// groupID - Group identifier. Must be unique in the current API Management service instance.
// options - ProductGroupClientCheckEntityExistsOptions contains the optional parameters for the ProductGroupClient.CheckEntityExists
// method.
func (client *ProductGroupClient) CheckEntityExists(ctx context.Context, resourceGroupName string, serviceName string, productID string, groupID string, options *ProductGroupClientCheckEntityExistsOptions) (ProductGroupClientCheckEntityExistsResponse, error) {
	req, err := client.checkEntityExistsCreateRequest(ctx, resourceGroupName, serviceName, productID, groupID, options)
	if err != nil {
		return ProductGroupClientCheckEntityExistsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProductGroupClientCheckEntityExistsResponse{}, err
	}
	result := ProductGroupClientCheckEntityExistsResponse{RawResponse: resp}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// checkEntityExistsCreateRequest creates the CheckEntityExists request.
func (client *ProductGroupClient) checkEntityExistsCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, productID string, groupID string, options *ProductGroupClientCheckEntityExistsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/groups/{groupId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if productID == "" {
		return nil, errors.New("parameter productID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productId}", url.PathEscape(productID))
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// CreateOrUpdate - Adds the association between the specified developer group with the specified product.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// productID - Product identifier. Must be unique in the current API Management service instance.
// groupID - Group identifier. Must be unique in the current API Management service instance.
// options - ProductGroupClientCreateOrUpdateOptions contains the optional parameters for the ProductGroupClient.CreateOrUpdate
// method.
func (client *ProductGroupClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, productID string, groupID string, options *ProductGroupClientCreateOrUpdateOptions) (ProductGroupClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, productID, groupID, options)
	if err != nil {
		return ProductGroupClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProductGroupClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ProductGroupClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ProductGroupClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, productID string, groupID string, options *ProductGroupClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/groups/{groupId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if productID == "" {
		return nil, errors.New("parameter productID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productId}", url.PathEscape(productID))
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ProductGroupClient) createOrUpdateHandleResponse(resp *http.Response) (ProductGroupClientCreateOrUpdateResponse, error) {
	result := ProductGroupClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.GroupContract); err != nil {
		return ProductGroupClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the association between the specified group and product.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// productID - Product identifier. Must be unique in the current API Management service instance.
// groupID - Group identifier. Must be unique in the current API Management service instance.
// options - ProductGroupClientDeleteOptions contains the optional parameters for the ProductGroupClient.Delete method.
func (client *ProductGroupClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, productID string, groupID string, options *ProductGroupClientDeleteOptions) (ProductGroupClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, productID, groupID, options)
	if err != nil {
		return ProductGroupClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProductGroupClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ProductGroupClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ProductGroupClientDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ProductGroupClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, productID string, groupID string, options *ProductGroupClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/groups/{groupId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if productID == "" {
		return nil, errors.New("parameter productID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productId}", url.PathEscape(productID))
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// ListByProduct - Lists the collection of developer groups associated with the specified product.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// productID - Product identifier. Must be unique in the current API Management service instance.
// options - ProductGroupClientListByProductOptions contains the optional parameters for the ProductGroupClient.ListByProduct
// method.
func (client *ProductGroupClient) ListByProduct(resourceGroupName string, serviceName string, productID string, options *ProductGroupClientListByProductOptions) *ProductGroupClientListByProductPager {
	return &ProductGroupClientListByProductPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByProductCreateRequest(ctx, resourceGroupName, serviceName, productID, options)
		},
		advancer: func(ctx context.Context, resp ProductGroupClientListByProductResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.GroupCollection.NextLink)
		},
	}
}

// listByProductCreateRequest creates the ListByProduct request.
func (client *ProductGroupClient) listByProductCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, productID string, options *ProductGroupClientListByProductOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/groups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if productID == "" {
		return nil, errors.New("parameter productID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productId}", url.PathEscape(productID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByProductHandleResponse handles the ListByProduct response.
func (client *ProductGroupClient) listByProductHandleResponse(resp *http.Response) (ProductGroupClientListByProductResponse, error) {
	result := ProductGroupClientListByProductResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.GroupCollection); err != nil {
		return ProductGroupClientListByProductResponse{}, err
	}
	return result, nil
}
