//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcdn

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

// AFDOriginsClient contains the methods for the AFDOrigins group.
// Don't use this type directly, use NewAFDOriginsClient() instead.
type AFDOriginsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAFDOriginsClient creates a new instance of AFDOriginsClient with the specified values.
//   - subscriptionID - Azure Subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAFDOriginsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AFDOriginsClient, error) {
	cl, err := arm.NewClient(moduleName+".AFDOriginsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AFDOriginsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates a new origin within the specified origin group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
//     group.
//   - originGroupName - Name of the origin group which is unique within the profile.
//   - originName - Name of the origin that is unique within the profile.
//   - origin - Origin properties
//   - options - AFDOriginsClientBeginCreateOptions contains the optional parameters for the AFDOriginsClient.BeginCreate method.
func (client *AFDOriginsClient) BeginCreate(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, originName string, origin AFDOrigin, options *AFDOriginsClientBeginCreateOptions) (*runtime.Poller[AFDOriginsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, profileName, originGroupName, originName, origin, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AFDOriginsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[AFDOriginsClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Creates a new origin within the specified origin group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
func (client *AFDOriginsClient) create(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, originName string, origin AFDOrigin, options *AFDOriginsClientBeginCreateOptions) (*http.Response, error) {
	var err error
	req, err := client.createCreateRequest(ctx, resourceGroupName, profileName, originGroupName, originName, origin, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *AFDOriginsClient) createCreateRequest(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, originName string, origin AFDOrigin, options *AFDOriginsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/originGroups/{originGroupName}/origins/{originName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if originGroupName == "" {
		return nil, errors.New("parameter originGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originGroupName}", url.PathEscape(originGroupName))
	if originName == "" {
		return nil, errors.New("parameter originName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originName}", url.PathEscape(originName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, origin); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes an existing origin within an origin group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
//     group.
//   - originGroupName - Name of the origin group which is unique within the profile.
//   - originName - Name of the origin which is unique within the profile.
//   - options - AFDOriginsClientBeginDeleteOptions contains the optional parameters for the AFDOriginsClient.BeginDelete method.
func (client *AFDOriginsClient) BeginDelete(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, originName string, options *AFDOriginsClientBeginDeleteOptions) (*runtime.Poller[AFDOriginsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, profileName, originGroupName, originName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AFDOriginsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[AFDOriginsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes an existing origin within an origin group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
func (client *AFDOriginsClient) deleteOperation(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, originName string, options *AFDOriginsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, profileName, originGroupName, originName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AFDOriginsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, originName string, options *AFDOriginsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/originGroups/{originGroupName}/origins/{originName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if originGroupName == "" {
		return nil, errors.New("parameter originGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originGroupName}", url.PathEscape(originGroupName))
	if originName == "" {
		return nil, errors.New("parameter originName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originName}", url.PathEscape(originName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// Get - Gets an existing origin within an origin group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
//     group.
//   - originGroupName - Name of the origin group which is unique within the profile.
//   - originName - Name of the origin which is unique within the profile.
//   - options - AFDOriginsClientGetOptions contains the optional parameters for the AFDOriginsClient.Get method.
func (client *AFDOriginsClient) Get(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, originName string, options *AFDOriginsClientGetOptions) (AFDOriginsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, profileName, originGroupName, originName, options)
	if err != nil {
		return AFDOriginsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AFDOriginsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AFDOriginsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AFDOriginsClient) getCreateRequest(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, originName string, options *AFDOriginsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/originGroups/{originGroupName}/origins/{originName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if originGroupName == "" {
		return nil, errors.New("parameter originGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originGroupName}", url.PathEscape(originGroupName))
	if originName == "" {
		return nil, errors.New("parameter originName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originName}", url.PathEscape(originName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
func (client *AFDOriginsClient) getHandleResponse(resp *http.Response) (AFDOriginsClientGetResponse, error) {
	result := AFDOriginsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AFDOrigin); err != nil {
		return AFDOriginsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByOriginGroupPager - Lists all of the existing origins within an origin group.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
//     group.
//   - originGroupName - Name of the origin group which is unique within the profile.
//   - options - AFDOriginsClientListByOriginGroupOptions contains the optional parameters for the AFDOriginsClient.NewListByOriginGroupPager
//     method.
func (client *AFDOriginsClient) NewListByOriginGroupPager(resourceGroupName string, profileName string, originGroupName string, options *AFDOriginsClientListByOriginGroupOptions) *runtime.Pager[AFDOriginsClientListByOriginGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[AFDOriginsClientListByOriginGroupResponse]{
		More: func(page AFDOriginsClientListByOriginGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AFDOriginsClientListByOriginGroupResponse) (AFDOriginsClientListByOriginGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByOriginGroupCreateRequest(ctx, resourceGroupName, profileName, originGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AFDOriginsClientListByOriginGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return AFDOriginsClientListByOriginGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AFDOriginsClientListByOriginGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByOriginGroupHandleResponse(resp)
		},
	})
}

// listByOriginGroupCreateRequest creates the ListByOriginGroup request.
func (client *AFDOriginsClient) listByOriginGroupCreateRequest(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, options *AFDOriginsClientListByOriginGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/originGroups/{originGroupName}/origins"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if originGroupName == "" {
		return nil, errors.New("parameter originGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originGroupName}", url.PathEscape(originGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// listByOriginGroupHandleResponse handles the ListByOriginGroup response.
func (client *AFDOriginsClient) listByOriginGroupHandleResponse(resp *http.Response) (AFDOriginsClientListByOriginGroupResponse, error) {
	result := AFDOriginsClientListByOriginGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AFDOriginListResult); err != nil {
		return AFDOriginsClientListByOriginGroupResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates an existing origin within an origin group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
//     group.
//   - originGroupName - Name of the origin group which is unique within the profile.
//   - originName - Name of the origin which is unique within the profile.
//   - originUpdateProperties - Origin properties
//   - options - AFDOriginsClientBeginUpdateOptions contains the optional parameters for the AFDOriginsClient.BeginUpdate method.
func (client *AFDOriginsClient) BeginUpdate(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, originName string, originUpdateProperties AFDOriginUpdateParameters, options *AFDOriginsClientBeginUpdateOptions) (*runtime.Poller[AFDOriginsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, profileName, originGroupName, originName, originUpdateProperties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AFDOriginsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[AFDOriginsClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Updates an existing origin within an origin group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
func (client *AFDOriginsClient) update(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, originName string, originUpdateProperties AFDOriginUpdateParameters, options *AFDOriginsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, profileName, originGroupName, originName, originUpdateProperties, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *AFDOriginsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, originName string, originUpdateProperties AFDOriginUpdateParameters, options *AFDOriginsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/originGroups/{originGroupName}/origins/{originName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if originGroupName == "" {
		return nil, errors.New("parameter originGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originGroupName}", url.PathEscape(originGroupName))
	if originName == "" {
		return nil, errors.New("parameter originName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originName}", url.PathEscape(originName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, originUpdateProperties); err != nil {
		return nil, err
	}
	return req, nil
}
