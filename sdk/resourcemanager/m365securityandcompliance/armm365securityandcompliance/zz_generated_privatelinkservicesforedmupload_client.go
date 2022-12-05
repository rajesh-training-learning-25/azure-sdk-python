//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armm365securityandcompliance

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

// PrivateLinkServicesForEDMUploadClient contains the methods for the PrivateLinkServicesForEDMUpload group.
// Don't use this type directly, use NewPrivateLinkServicesForEDMUploadClient() instead.
type PrivateLinkServicesForEDMUploadClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPrivateLinkServicesForEDMUploadClient creates a new instance of PrivateLinkServicesForEDMUploadClient with the specified values.
// subscriptionID - The subscription identifier.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPrivateLinkServicesForEDMUploadClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateLinkServicesForEDMUploadClient, error) {
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
	client := &PrivateLinkServicesForEDMUploadClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update the metadata of a privateLinkServicesForEDMUpload instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-03-25-preview
// resourceGroupName - The name of the resource group that contains the service instance.
// resourceName - The name of the service instance.
// privateLinkServicesForEDMUploadDescription - The service instance metadata.
// options - PrivateLinkServicesForEDMUploadClientBeginCreateOrUpdateOptions contains the optional parameters for the PrivateLinkServicesForEDMUploadClient.BeginCreateOrUpdate
// method.
func (client *PrivateLinkServicesForEDMUploadClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, privateLinkServicesForEDMUploadDescription PrivateLinkServicesForEDMUploadDescription, options *PrivateLinkServicesForEDMUploadClientBeginCreateOrUpdateOptions) (*runtime.Poller[PrivateLinkServicesForEDMUploadClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, resourceName, privateLinkServicesForEDMUploadDescription, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[PrivateLinkServicesForEDMUploadClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[PrivateLinkServicesForEDMUploadClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update the metadata of a privateLinkServicesForEDMUpload instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-03-25-preview
func (client *PrivateLinkServicesForEDMUploadClient) createOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, privateLinkServicesForEDMUploadDescription PrivateLinkServicesForEDMUploadDescription, options *PrivateLinkServicesForEDMUploadClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, privateLinkServicesForEDMUploadDescription, options)
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
func (client *PrivateLinkServicesForEDMUploadClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, privateLinkServicesForEDMUploadDescription PrivateLinkServicesForEDMUploadDescription, options *PrivateLinkServicesForEDMUploadClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForEDMUpload/{resourceName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, privateLinkServicesForEDMUploadDescription)
}

// Get - Get the metadata of a privateLinkServicesForEDMUpload resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-03-25-preview
// resourceGroupName - The name of the resource group that contains the service instance.
// resourceName - The name of the service instance.
// options - PrivateLinkServicesForEDMUploadClientGetOptions contains the optional parameters for the PrivateLinkServicesForEDMUploadClient.Get
// method.
func (client *PrivateLinkServicesForEDMUploadClient) Get(ctx context.Context, resourceGroupName string, resourceName string, options *PrivateLinkServicesForEDMUploadClientGetOptions) (PrivateLinkServicesForEDMUploadClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return PrivateLinkServicesForEDMUploadClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateLinkServicesForEDMUploadClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateLinkServicesForEDMUploadClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateLinkServicesForEDMUploadClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *PrivateLinkServicesForEDMUploadClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForEDMUpload/{resourceName}"
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
	reqQP.Set("api-version", "2021-03-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateLinkServicesForEDMUploadClient) getHandleResponse(resp *http.Response) (PrivateLinkServicesForEDMUploadClientGetResponse, error) {
	result := PrivateLinkServicesForEDMUploadClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkServicesForEDMUploadDescription); err != nil {
		return PrivateLinkServicesForEDMUploadClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get all the privateLinkServicesForEDMUpload instances in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-03-25-preview
// options - PrivateLinkServicesForEDMUploadClientListOptions contains the optional parameters for the PrivateLinkServicesForEDMUploadClient.List
// method.
func (client *PrivateLinkServicesForEDMUploadClient) NewListPager(options *PrivateLinkServicesForEDMUploadClientListOptions) *runtime.Pager[PrivateLinkServicesForEDMUploadClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateLinkServicesForEDMUploadClientListResponse]{
		More: func(page PrivateLinkServicesForEDMUploadClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateLinkServicesForEDMUploadClientListResponse) (PrivateLinkServicesForEDMUploadClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PrivateLinkServicesForEDMUploadClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PrivateLinkServicesForEDMUploadClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PrivateLinkServicesForEDMUploadClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *PrivateLinkServicesForEDMUploadClient) listCreateRequest(ctx context.Context, options *PrivateLinkServicesForEDMUploadClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForEDMUpload"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PrivateLinkServicesForEDMUploadClient) listHandleResponse(resp *http.Response) (PrivateLinkServicesForEDMUploadClientListResponse, error) {
	result := PrivateLinkServicesForEDMUploadClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkServicesForEDMUploadDescriptionListResult); err != nil {
		return PrivateLinkServicesForEDMUploadClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Get all the service instances in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-03-25-preview
// resourceGroupName - The name of the resource group that contains the service instance.
// options - PrivateLinkServicesForEDMUploadClientListByResourceGroupOptions contains the optional parameters for the PrivateLinkServicesForEDMUploadClient.ListByResourceGroup
// method.
func (client *PrivateLinkServicesForEDMUploadClient) NewListByResourceGroupPager(resourceGroupName string, options *PrivateLinkServicesForEDMUploadClientListByResourceGroupOptions) *runtime.Pager[PrivateLinkServicesForEDMUploadClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateLinkServicesForEDMUploadClientListByResourceGroupResponse]{
		More: func(page PrivateLinkServicesForEDMUploadClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateLinkServicesForEDMUploadClientListByResourceGroupResponse) (PrivateLinkServicesForEDMUploadClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PrivateLinkServicesForEDMUploadClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PrivateLinkServicesForEDMUploadClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PrivateLinkServicesForEDMUploadClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *PrivateLinkServicesForEDMUploadClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *PrivateLinkServicesForEDMUploadClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForEDMUpload"
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
	reqQP.Set("api-version", "2021-03-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *PrivateLinkServicesForEDMUploadClient) listByResourceGroupHandleResponse(resp *http.Response) (PrivateLinkServicesForEDMUploadClientListByResourceGroupResponse, error) {
	result := PrivateLinkServicesForEDMUploadClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkServicesForEDMUploadDescriptionListResult); err != nil {
		return PrivateLinkServicesForEDMUploadClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update the metadata of a privateLinkServicesForEDMUpload instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-03-25-preview
// resourceGroupName - The name of the resource group that contains the service instance.
// resourceName - The name of the service instance.
// servicePatchDescription - The service instance metadata and security metadata.
// options - PrivateLinkServicesForEDMUploadClientBeginUpdateOptions contains the optional parameters for the PrivateLinkServicesForEDMUploadClient.BeginUpdate
// method.
func (client *PrivateLinkServicesForEDMUploadClient) BeginUpdate(ctx context.Context, resourceGroupName string, resourceName string, servicePatchDescription ServicesPatchDescription, options *PrivateLinkServicesForEDMUploadClientBeginUpdateOptions) (*runtime.Poller[PrivateLinkServicesForEDMUploadClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, resourceName, servicePatchDescription, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[PrivateLinkServicesForEDMUploadClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[PrivateLinkServicesForEDMUploadClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Update the metadata of a privateLinkServicesForEDMUpload instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-03-25-preview
func (client *PrivateLinkServicesForEDMUploadClient) update(ctx context.Context, resourceGroupName string, resourceName string, servicePatchDescription ServicesPatchDescription, options *PrivateLinkServicesForEDMUploadClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, servicePatchDescription, options)
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
func (client *PrivateLinkServicesForEDMUploadClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, servicePatchDescription ServicesPatchDescription, options *PrivateLinkServicesForEDMUploadClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForEDMUpload/{resourceName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, servicePatchDescription)
}