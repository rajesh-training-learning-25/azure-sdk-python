//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armstorage

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

// EncryptionScopesClient contains the methods for the EncryptionScopes group.
// Don't use this type directly, use NewEncryptionScopesClient() instead.
type EncryptionScopesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewEncryptionScopesClient creates a new instance of EncryptionScopesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewEncryptionScopesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*EncryptionScopesClient, error) {
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
	client := &EncryptionScopesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Returns the properties for the specified encryption scope.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// accountName - The name of the storage account within the specified resource group. Storage account names must be between
// 3 and 24 characters in length and use numbers and lower-case letters only.
// encryptionScopeName - The name of the encryption scope within the specified storage account. Encryption scope names must
// be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every
// dash (-) character must be immediately preceded and followed by a letter or number.
// options - EncryptionScopesClientGetOptions contains the optional parameters for the EncryptionScopesClient.Get method.
func (client *EncryptionScopesClient) Get(ctx context.Context, resourceGroupName string, accountName string, encryptionScopeName string, options *EncryptionScopesClientGetOptions) (EncryptionScopesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, encryptionScopeName, options)
	if err != nil {
		return EncryptionScopesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EncryptionScopesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EncryptionScopesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *EncryptionScopesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, encryptionScopeName string, options *EncryptionScopesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/encryptionScopes/{encryptionScopeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if encryptionScopeName == "" {
		return nil, errors.New("parameter encryptionScopeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{encryptionScopeName}", url.PathEscape(encryptionScopeName))
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

// getHandleResponse handles the Get response.
func (client *EncryptionScopesClient) getHandleResponse(resp *http.Response) (EncryptionScopesClientGetResponse, error) {
	result := EncryptionScopesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EncryptionScope); err != nil {
		return EncryptionScopesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all the encryption scopes available under the specified storage account.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// accountName - The name of the storage account within the specified resource group. Storage account names must be between
// 3 and 24 characters in length and use numbers and lower-case letters only.
// options - EncryptionScopesClientListOptions contains the optional parameters for the EncryptionScopesClient.List method.
func (client *EncryptionScopesClient) NewListPager(resourceGroupName string, accountName string, options *EncryptionScopesClientListOptions) *runtime.Pager[EncryptionScopesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[EncryptionScopesClientListResponse]{
		More: func(page EncryptionScopesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EncryptionScopesClientListResponse) (EncryptionScopesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, accountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return EncryptionScopesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return EncryptionScopesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return EncryptionScopesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *EncryptionScopesClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *EncryptionScopesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/encryptionScopes"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
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
func (client *EncryptionScopesClient) listHandleResponse(resp *http.Response) (EncryptionScopesClientListResponse, error) {
	result := EncryptionScopesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EncryptionScopeListResult); err != nil {
		return EncryptionScopesClientListResponse{}, err
	}
	return result, nil
}

// Patch - Update encryption scope properties as specified in the request body. Update fails if the specified encryption scope
// does not already exist.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// accountName - The name of the storage account within the specified resource group. Storage account names must be between
// 3 and 24 characters in length and use numbers and lower-case letters only.
// encryptionScopeName - The name of the encryption scope within the specified storage account. Encryption scope names must
// be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every
// dash (-) character must be immediately preceded and followed by a letter or number.
// encryptionScope - Encryption scope properties to be used for the update.
// options - EncryptionScopesClientPatchOptions contains the optional parameters for the EncryptionScopesClient.Patch method.
func (client *EncryptionScopesClient) Patch(ctx context.Context, resourceGroupName string, accountName string, encryptionScopeName string, encryptionScope EncryptionScope, options *EncryptionScopesClientPatchOptions) (EncryptionScopesClientPatchResponse, error) {
	req, err := client.patchCreateRequest(ctx, resourceGroupName, accountName, encryptionScopeName, encryptionScope, options)
	if err != nil {
		return EncryptionScopesClientPatchResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EncryptionScopesClientPatchResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EncryptionScopesClientPatchResponse{}, runtime.NewResponseError(resp)
	}
	return client.patchHandleResponse(resp)
}

// patchCreateRequest creates the Patch request.
func (client *EncryptionScopesClient) patchCreateRequest(ctx context.Context, resourceGroupName string, accountName string, encryptionScopeName string, encryptionScope EncryptionScope, options *EncryptionScopesClientPatchOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/encryptionScopes/{encryptionScopeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if encryptionScopeName == "" {
		return nil, errors.New("parameter encryptionScopeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{encryptionScopeName}", url.PathEscape(encryptionScopeName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, encryptionScope)
}

// patchHandleResponse handles the Patch response.
func (client *EncryptionScopesClient) patchHandleResponse(resp *http.Response) (EncryptionScopesClientPatchResponse, error) {
	result := EncryptionScopesClientPatchResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EncryptionScope); err != nil {
		return EncryptionScopesClientPatchResponse{}, err
	}
	return result, nil
}

// Put - Synchronously creates or updates an encryption scope under the specified storage account. If an encryption scope
// is already created and a subsequent request is issued with different properties, the
// encryption scope properties will be updated per the specified request.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// accountName - The name of the storage account within the specified resource group. Storage account names must be between
// 3 and 24 characters in length and use numbers and lower-case letters only.
// encryptionScopeName - The name of the encryption scope within the specified storage account. Encryption scope names must
// be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every
// dash (-) character must be immediately preceded and followed by a letter or number.
// encryptionScope - Encryption scope properties to be used for the create or update.
// options - EncryptionScopesClientPutOptions contains the optional parameters for the EncryptionScopesClient.Put method.
func (client *EncryptionScopesClient) Put(ctx context.Context, resourceGroupName string, accountName string, encryptionScopeName string, encryptionScope EncryptionScope, options *EncryptionScopesClientPutOptions) (EncryptionScopesClientPutResponse, error) {
	req, err := client.putCreateRequest(ctx, resourceGroupName, accountName, encryptionScopeName, encryptionScope, options)
	if err != nil {
		return EncryptionScopesClientPutResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EncryptionScopesClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return EncryptionScopesClientPutResponse{}, runtime.NewResponseError(resp)
	}
	return client.putHandleResponse(resp)
}

// putCreateRequest creates the Put request.
func (client *EncryptionScopesClient) putCreateRequest(ctx context.Context, resourceGroupName string, accountName string, encryptionScopeName string, encryptionScope EncryptionScope, options *EncryptionScopesClientPutOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/encryptionScopes/{encryptionScopeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if encryptionScopeName == "" {
		return nil, errors.New("parameter encryptionScopeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{encryptionScopeName}", url.PathEscape(encryptionScopeName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, encryptionScope)
}

// putHandleResponse handles the Put response.
func (client *EncryptionScopesClient) putHandleResponse(resp *http.Response) (EncryptionScopesClientPutResponse, error) {
	result := EncryptionScopesClientPutResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EncryptionScope); err != nil {
		return EncryptionScopesClientPutResponse{}, err
	}
	return result, nil
}