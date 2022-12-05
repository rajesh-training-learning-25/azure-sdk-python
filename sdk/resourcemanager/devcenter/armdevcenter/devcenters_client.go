//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdevcenter

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
	"strconv"
	"strings"
)

// DevCentersClient contains the methods for the DevCenters group.
// Don't use this type directly, use NewDevCentersClient() instead.
type DevCentersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDevCentersClient creates a new instance of DevCentersClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDevCentersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DevCentersClient, error) {
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
	client := &DevCentersClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a devcenter resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-12-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// devCenterName - The name of the devcenter.
// body - Represents a devcenter.
// options - DevCentersClientBeginCreateOrUpdateOptions contains the optional parameters for the DevCentersClient.BeginCreateOrUpdate
// method.
func (client *DevCentersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, devCenterName string, body DevCenter, options *DevCentersClientBeginCreateOrUpdateOptions) (*runtime.Poller[DevCentersClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, devCenterName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[DevCentersClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DevCentersClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a devcenter resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-12-preview
func (client *DevCentersClient) createOrUpdate(ctx context.Context, resourceGroupName string, devCenterName string, body DevCenter, options *DevCentersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, devCenterName, body, options)
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
func (client *DevCentersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, body DevCenter, options *DevCentersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if devCenterName == "" {
		return nil, errors.New("parameter devCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devCenterName}", url.PathEscape(devCenterName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-12-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Deletes a devcenter
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-12-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// devCenterName - The name of the devcenter.
// options - DevCentersClientBeginDeleteOptions contains the optional parameters for the DevCentersClient.BeginDelete method.
func (client *DevCentersClient) BeginDelete(ctx context.Context, resourceGroupName string, devCenterName string, options *DevCentersClientBeginDeleteOptions) (*runtime.Poller[DevCentersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, devCenterName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[DevCentersClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DevCentersClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a devcenter
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-12-preview
func (client *DevCentersClient) deleteOperation(ctx context.Context, resourceGroupName string, devCenterName string, options *DevCentersClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, devCenterName, options)
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
func (client *DevCentersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, options *DevCentersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if devCenterName == "" {
		return nil, errors.New("parameter devCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devCenterName}", url.PathEscape(devCenterName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-12-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a devcenter.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-12-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// devCenterName - The name of the devcenter.
// options - DevCentersClientGetOptions contains the optional parameters for the DevCentersClient.Get method.
func (client *DevCentersClient) Get(ctx context.Context, resourceGroupName string, devCenterName string, options *DevCentersClientGetOptions) (DevCentersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, devCenterName, options)
	if err != nil {
		return DevCentersClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DevCentersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DevCentersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DevCentersClient) getCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, options *DevCentersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if devCenterName == "" {
		return nil, errors.New("parameter devCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devCenterName}", url.PathEscape(devCenterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-12-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DevCentersClient) getHandleResponse(resp *http.Response) (DevCentersClientGetResponse, error) {
	result := DevCentersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DevCenter); err != nil {
		return DevCentersClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists all devcenters in a resource group.
// Generated from API version 2022-10-12-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - DevCentersClientListByResourceGroupOptions contains the optional parameters for the DevCentersClient.ListByResourceGroup
// method.
func (client *DevCentersClient) NewListByResourceGroupPager(resourceGroupName string, options *DevCentersClientListByResourceGroupOptions) *runtime.Pager[DevCentersClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[DevCentersClientListByResourceGroupResponse]{
		More: func(page DevCentersClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DevCentersClientListByResourceGroupResponse) (DevCentersClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DevCentersClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DevCentersClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DevCentersClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DevCentersClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DevCentersClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters"
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
	reqQP.Set("api-version", "2022-10-12-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DevCentersClient) listByResourceGroupHandleResponse(resp *http.Response) (DevCentersClientListByResourceGroupResponse, error) {
	result := DevCentersClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return DevCentersClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists all devcenters in a subscription.
// Generated from API version 2022-10-12-preview
// options - DevCentersClientListBySubscriptionOptions contains the optional parameters for the DevCentersClient.ListBySubscription
// method.
func (client *DevCentersClient) NewListBySubscriptionPager(options *DevCentersClientListBySubscriptionOptions) *runtime.Pager[DevCentersClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[DevCentersClientListBySubscriptionResponse]{
		More: func(page DevCentersClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DevCentersClientListBySubscriptionResponse) (DevCentersClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DevCentersClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DevCentersClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DevCentersClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *DevCentersClient) listBySubscriptionCreateRequest(ctx context.Context, options *DevCentersClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DevCenter/devcenters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-12-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *DevCentersClient) listBySubscriptionHandleResponse(resp *http.Response) (DevCentersClientListBySubscriptionResponse, error) {
	result := DevCentersClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return DevCentersClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Partially updates a devcenter.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-12-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// devCenterName - The name of the devcenter.
// body - Updatable devcenter properties.
// options - DevCentersClientBeginUpdateOptions contains the optional parameters for the DevCentersClient.BeginUpdate method.
func (client *DevCentersClient) BeginUpdate(ctx context.Context, resourceGroupName string, devCenterName string, body Update, options *DevCentersClientBeginUpdateOptions) (*runtime.Poller[DevCentersClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, devCenterName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[DevCentersClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DevCentersClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Partially updates a devcenter.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-12-preview
func (client *DevCentersClient) update(ctx context.Context, resourceGroupName string, devCenterName string, body Update, options *DevCentersClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, devCenterName, body, options)
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
func (client *DevCentersClient) updateCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, body Update, options *DevCentersClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if devCenterName == "" {
		return nil, errors.New("parameter devCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devCenterName}", url.PathEscape(devCenterName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-12-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}