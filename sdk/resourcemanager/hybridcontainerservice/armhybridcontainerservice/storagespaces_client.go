//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhybridcontainerservice

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

// StorageSpacesClient contains the methods for the StorageSpaces group.
// Don't use this type directly, use NewStorageSpacesClient() instead.
type StorageSpacesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewStorageSpacesClient creates a new instance of StorageSpacesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewStorageSpacesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*StorageSpacesClient, error) {
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
	client := &StorageSpacesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Puts the Hybrid AKS storage object
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageSpacesName - Parameter for the name of the storage object
// options - StorageSpacesClientBeginCreateOrUpdateOptions contains the optional parameters for the StorageSpacesClient.BeginCreateOrUpdate
// method.
func (client *StorageSpacesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, storageSpacesName string, storageSpaces StorageSpaces, options *StorageSpacesClientBeginCreateOrUpdateOptions) (*runtime.Poller[StorageSpacesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, storageSpacesName, storageSpaces, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[StorageSpacesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[StorageSpacesClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Puts the Hybrid AKS storage object
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
func (client *StorageSpacesClient) createOrUpdate(ctx context.Context, resourceGroupName string, storageSpacesName string, storageSpaces StorageSpaces, options *StorageSpacesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, storageSpacesName, storageSpaces, options)
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
func (client *StorageSpacesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, storageSpacesName string, storageSpaces StorageSpaces, options *StorageSpacesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridContainerService/storageSpaces/{storageSpacesName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageSpacesName == "" {
		return nil, errors.New("parameter storageSpacesName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageSpacesName}", url.PathEscape(storageSpacesName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, storageSpaces)
}

// Delete - Deletes the Hybrid AKS storage object
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageSpacesName - Parameter for the name of the storage object
// options - StorageSpacesClientDeleteOptions contains the optional parameters for the StorageSpacesClient.Delete method.
func (client *StorageSpacesClient) Delete(ctx context.Context, resourceGroupName string, storageSpacesName string, options *StorageSpacesClientDeleteOptions) (StorageSpacesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, storageSpacesName, options)
	if err != nil {
		return StorageSpacesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StorageSpacesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return StorageSpacesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return StorageSpacesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *StorageSpacesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, storageSpacesName string, options *StorageSpacesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridContainerService/storageSpaces/{storageSpacesName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageSpacesName == "" {
		return nil, errors.New("parameter storageSpacesName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageSpacesName}", url.PathEscape(storageSpacesName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// NewListByResourceGroupPager - List the Hybrid AKS storage object by resource group
// Generated from API version 2022-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - StorageSpacesClientListByResourceGroupOptions contains the optional parameters for the StorageSpacesClient.ListByResourceGroup
// method.
func (client *StorageSpacesClient) NewListByResourceGroupPager(resourceGroupName string, options *StorageSpacesClientListByResourceGroupOptions) *runtime.Pager[StorageSpacesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[StorageSpacesClientListByResourceGroupResponse]{
		More: func(page StorageSpacesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *StorageSpacesClientListByResourceGroupResponse) (StorageSpacesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return StorageSpacesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return StorageSpacesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return StorageSpacesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *StorageSpacesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *StorageSpacesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridContainerService/storageSpaces"
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
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *StorageSpacesClient) listByResourceGroupHandleResponse(resp *http.Response) (StorageSpacesClientListByResourceGroupResponse, error) {
	result := StorageSpacesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageSpacesListResult); err != nil {
		return StorageSpacesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List the Hybrid AKS storage object by subscription
// Generated from API version 2022-05-01-preview
// options - StorageSpacesClientListBySubscriptionOptions contains the optional parameters for the StorageSpacesClient.ListBySubscription
// method.
func (client *StorageSpacesClient) NewListBySubscriptionPager(options *StorageSpacesClientListBySubscriptionOptions) *runtime.Pager[StorageSpacesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[StorageSpacesClientListBySubscriptionResponse]{
		More: func(page StorageSpacesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *StorageSpacesClientListBySubscriptionResponse) (StorageSpacesClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return StorageSpacesClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return StorageSpacesClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return StorageSpacesClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *StorageSpacesClient) listBySubscriptionCreateRequest(ctx context.Context, options *StorageSpacesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridContainerService/storageSpaces"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *StorageSpacesClient) listBySubscriptionHandleResponse(resp *http.Response) (StorageSpacesClientListBySubscriptionResponse, error) {
	result := StorageSpacesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageSpacesListResult); err != nil {
		return StorageSpacesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Retrieve - Gets the Hybrid AKS storage space object
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageSpacesName - Parameter for the name of the storage object
// options - StorageSpacesClientRetrieveOptions contains the optional parameters for the StorageSpacesClient.Retrieve method.
func (client *StorageSpacesClient) Retrieve(ctx context.Context, resourceGroupName string, storageSpacesName string, options *StorageSpacesClientRetrieveOptions) (StorageSpacesClientRetrieveResponse, error) {
	req, err := client.retrieveCreateRequest(ctx, resourceGroupName, storageSpacesName, options)
	if err != nil {
		return StorageSpacesClientRetrieveResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StorageSpacesClientRetrieveResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StorageSpacesClientRetrieveResponse{}, runtime.NewResponseError(resp)
	}
	return client.retrieveHandleResponse(resp)
}

// retrieveCreateRequest creates the Retrieve request.
func (client *StorageSpacesClient) retrieveCreateRequest(ctx context.Context, resourceGroupName string, storageSpacesName string, options *StorageSpacesClientRetrieveOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridContainerService/storageSpaces/{storageSpacesName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageSpacesName == "" {
		return nil, errors.New("parameter storageSpacesName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageSpacesName}", url.PathEscape(storageSpacesName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// retrieveHandleResponse handles the Retrieve response.
func (client *StorageSpacesClient) retrieveHandleResponse(resp *http.Response) (StorageSpacesClientRetrieveResponse, error) {
	result := StorageSpacesClientRetrieveResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageSpaces); err != nil {
		return StorageSpacesClientRetrieveResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Patches the Hybrid AKS storage object
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageSpacesName - Parameter for the name of the storage object
// options - StorageSpacesClientBeginUpdateOptions contains the optional parameters for the StorageSpacesClient.BeginUpdate
// method.
func (client *StorageSpacesClient) BeginUpdate(ctx context.Context, resourceGroupName string, storageSpacesName string, storageSpaces StorageSpacesPatch, options *StorageSpacesClientBeginUpdateOptions) (*runtime.Poller[StorageSpacesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, storageSpacesName, storageSpaces, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[StorageSpacesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[StorageSpacesClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Patches the Hybrid AKS storage object
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
func (client *StorageSpacesClient) update(ctx context.Context, resourceGroupName string, storageSpacesName string, storageSpaces StorageSpacesPatch, options *StorageSpacesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, storageSpacesName, storageSpaces, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *StorageSpacesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, storageSpacesName string, storageSpaces StorageSpacesPatch, options *StorageSpacesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridContainerService/storageSpaces/{storageSpacesName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageSpacesName == "" {
		return nil, errors.New("parameter storageSpacesName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageSpacesName}", url.PathEscape(storageSpacesName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, storageSpaces)
}