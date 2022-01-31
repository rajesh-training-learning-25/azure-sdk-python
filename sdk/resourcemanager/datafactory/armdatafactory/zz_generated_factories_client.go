//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatafactory

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
	"strings"
)

// FactoriesClient contains the methods for the Factories group.
// Don't use this type directly, use NewFactoriesClient() instead.
type FactoriesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewFactoriesClient creates a new instance of FactoriesClient with the specified values.
// subscriptionID - The subscription identifier.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewFactoriesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *FactoriesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &FactoriesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// ConfigureFactoryRepo - Updates a factory's repo information.
// If the operation fails it returns an *azcore.ResponseError type.
// locationID - The location identifier.
// factoryRepoUpdate - Update factory repo request definition.
// options - FactoriesClientConfigureFactoryRepoOptions contains the optional parameters for the FactoriesClient.ConfigureFactoryRepo
// method.
func (client *FactoriesClient) ConfigureFactoryRepo(ctx context.Context, locationID string, factoryRepoUpdate FactoryRepoUpdate, options *FactoriesClientConfigureFactoryRepoOptions) (FactoriesClientConfigureFactoryRepoResponse, error) {
	req, err := client.configureFactoryRepoCreateRequest(ctx, locationID, factoryRepoUpdate, options)
	if err != nil {
		return FactoriesClientConfigureFactoryRepoResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FactoriesClientConfigureFactoryRepoResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FactoriesClientConfigureFactoryRepoResponse{}, runtime.NewResponseError(resp)
	}
	return client.configureFactoryRepoHandleResponse(resp)
}

// configureFactoryRepoCreateRequest creates the ConfigureFactoryRepo request.
func (client *FactoriesClient) configureFactoryRepoCreateRequest(ctx context.Context, locationID string, factoryRepoUpdate FactoryRepoUpdate, options *FactoriesClientConfigureFactoryRepoOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DataFactory/locations/{locationId}/configureFactoryRepo"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationID == "" {
		return nil, errors.New("parameter locationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationId}", url.PathEscape(locationID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, factoryRepoUpdate)
}

// configureFactoryRepoHandleResponse handles the ConfigureFactoryRepo response.
func (client *FactoriesClient) configureFactoryRepoHandleResponse(resp *http.Response) (FactoriesClientConfigureFactoryRepoResponse, error) {
	result := FactoriesClientConfigureFactoryRepoResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Factory); err != nil {
		return FactoriesClientConfigureFactoryRepoResponse{}, err
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a factory.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// factoryName - The factory name.
// factory - Factory resource definition.
// options - FactoriesClientCreateOrUpdateOptions contains the optional parameters for the FactoriesClient.CreateOrUpdate
// method.
func (client *FactoriesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, factoryName string, factory Factory, options *FactoriesClientCreateOrUpdateOptions) (FactoriesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, factoryName, factory, options)
	if err != nil {
		return FactoriesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FactoriesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FactoriesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *FactoriesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, factory Factory, options *FactoriesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, factory)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *FactoriesClient) createOrUpdateHandleResponse(resp *http.Response) (FactoriesClientCreateOrUpdateResponse, error) {
	result := FactoriesClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Factory); err != nil {
		return FactoriesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a factory.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// factoryName - The factory name.
// options - FactoriesClientDeleteOptions contains the optional parameters for the FactoriesClient.Delete method.
func (client *FactoriesClient) Delete(ctx context.Context, resourceGroupName string, factoryName string, options *FactoriesClientDeleteOptions) (FactoriesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, factoryName, options)
	if err != nil {
		return FactoriesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FactoriesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return FactoriesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return FactoriesClientDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *FactoriesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, options *FactoriesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets a factory.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// factoryName - The factory name.
// options - FactoriesClientGetOptions contains the optional parameters for the FactoriesClient.Get method.
func (client *FactoriesClient) Get(ctx context.Context, resourceGroupName string, factoryName string, options *FactoriesClientGetOptions) (FactoriesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, factoryName, options)
	if err != nil {
		return FactoriesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FactoriesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotModified) {
		return FactoriesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *FactoriesClient) getCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, options *FactoriesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FactoriesClient) getHandleResponse(resp *http.Response) (FactoriesClientGetResponse, error) {
	result := FactoriesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Factory); err != nil {
		return FactoriesClientGetResponse{}, err
	}
	return result, nil
}

// GetDataPlaneAccess - Get Data Plane access.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// factoryName - The factory name.
// policy - Data Plane user access policy definition.
// options - FactoriesClientGetDataPlaneAccessOptions contains the optional parameters for the FactoriesClient.GetDataPlaneAccess
// method.
func (client *FactoriesClient) GetDataPlaneAccess(ctx context.Context, resourceGroupName string, factoryName string, policy UserAccessPolicy, options *FactoriesClientGetDataPlaneAccessOptions) (FactoriesClientGetDataPlaneAccessResponse, error) {
	req, err := client.getDataPlaneAccessCreateRequest(ctx, resourceGroupName, factoryName, policy, options)
	if err != nil {
		return FactoriesClientGetDataPlaneAccessResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FactoriesClientGetDataPlaneAccessResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FactoriesClientGetDataPlaneAccessResponse{}, runtime.NewResponseError(resp)
	}
	return client.getDataPlaneAccessHandleResponse(resp)
}

// getDataPlaneAccessCreateRequest creates the GetDataPlaneAccess request.
func (client *FactoriesClient) getDataPlaneAccessCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, policy UserAccessPolicy, options *FactoriesClientGetDataPlaneAccessOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/getDataPlaneAccess"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, policy)
}

// getDataPlaneAccessHandleResponse handles the GetDataPlaneAccess response.
func (client *FactoriesClient) getDataPlaneAccessHandleResponse(resp *http.Response) (FactoriesClientGetDataPlaneAccessResponse, error) {
	result := FactoriesClientGetDataPlaneAccessResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessPolicyResponse); err != nil {
		return FactoriesClientGetDataPlaneAccessResponse{}, err
	}
	return result, nil
}

// GetGitHubAccessToken - Get GitHub Access Token.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// factoryName - The factory name.
// gitHubAccessTokenRequest - Get GitHub access token request definition.
// options - FactoriesClientGetGitHubAccessTokenOptions contains the optional parameters for the FactoriesClient.GetGitHubAccessToken
// method.
func (client *FactoriesClient) GetGitHubAccessToken(ctx context.Context, resourceGroupName string, factoryName string, gitHubAccessTokenRequest GitHubAccessTokenRequest, options *FactoriesClientGetGitHubAccessTokenOptions) (FactoriesClientGetGitHubAccessTokenResponse, error) {
	req, err := client.getGitHubAccessTokenCreateRequest(ctx, resourceGroupName, factoryName, gitHubAccessTokenRequest, options)
	if err != nil {
		return FactoriesClientGetGitHubAccessTokenResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FactoriesClientGetGitHubAccessTokenResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FactoriesClientGetGitHubAccessTokenResponse{}, runtime.NewResponseError(resp)
	}
	return client.getGitHubAccessTokenHandleResponse(resp)
}

// getGitHubAccessTokenCreateRequest creates the GetGitHubAccessToken request.
func (client *FactoriesClient) getGitHubAccessTokenCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, gitHubAccessTokenRequest GitHubAccessTokenRequest, options *FactoriesClientGetGitHubAccessTokenOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/getGitHubAccessToken"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, gitHubAccessTokenRequest)
}

// getGitHubAccessTokenHandleResponse handles the GetGitHubAccessToken response.
func (client *FactoriesClient) getGitHubAccessTokenHandleResponse(resp *http.Response) (FactoriesClientGetGitHubAccessTokenResponse, error) {
	result := FactoriesClientGetGitHubAccessTokenResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.GitHubAccessTokenResponse); err != nil {
		return FactoriesClientGetGitHubAccessTokenResponse{}, err
	}
	return result, nil
}

// List - Lists factories under the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - FactoriesClientListOptions contains the optional parameters for the FactoriesClient.List method.
func (client *FactoriesClient) List(options *FactoriesClientListOptions) *FactoriesClientListPager {
	return &FactoriesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp FactoriesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.FactoryListResponse.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *FactoriesClient) listCreateRequest(ctx context.Context, options *FactoriesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DataFactory/factories"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *FactoriesClient) listHandleResponse(resp *http.Response) (FactoriesClientListResponse, error) {
	result := FactoriesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.FactoryListResponse); err != nil {
		return FactoriesClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Lists factories.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// options - FactoriesClientListByResourceGroupOptions contains the optional parameters for the FactoriesClient.ListByResourceGroup
// method.
func (client *FactoriesClient) ListByResourceGroup(resourceGroupName string, options *FactoriesClientListByResourceGroupOptions) *FactoriesClientListByResourceGroupPager {
	return &FactoriesClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp FactoriesClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.FactoryListResponse.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *FactoriesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *FactoriesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories"
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
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *FactoriesClient) listByResourceGroupHandleResponse(resp *http.Response) (FactoriesClientListByResourceGroupResponse, error) {
	result := FactoriesClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.FactoryListResponse); err != nil {
		return FactoriesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// Update - Updates a factory.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// factoryName - The factory name.
// factoryUpdateParameters - The parameters for updating a factory.
// options - FactoriesClientUpdateOptions contains the optional parameters for the FactoriesClient.Update method.
func (client *FactoriesClient) Update(ctx context.Context, resourceGroupName string, factoryName string, factoryUpdateParameters FactoryUpdateParameters, options *FactoriesClientUpdateOptions) (FactoriesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, factoryName, factoryUpdateParameters, options)
	if err != nil {
		return FactoriesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FactoriesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FactoriesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *FactoriesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, factoryUpdateParameters FactoryUpdateParameters, options *FactoriesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, factoryUpdateParameters)
}

// updateHandleResponse handles the Update response.
func (client *FactoriesClient) updateHandleResponse(resp *http.Response) (FactoriesClientUpdateResponse, error) {
	result := FactoriesClientUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Factory); err != nil {
		return FactoriesClientUpdateResponse{}, err
	}
	return result, nil
}
