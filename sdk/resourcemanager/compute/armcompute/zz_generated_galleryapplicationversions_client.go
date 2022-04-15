//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

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

// GalleryApplicationVersionsClient contains the methods for the GalleryApplicationVersions group.
// Don't use this type directly, use NewGalleryApplicationVersionsClient() instead.
type GalleryApplicationVersionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewGalleryApplicationVersionsClient creates a new instance of GalleryApplicationVersionsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewGalleryApplicationVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GalleryApplicationVersionsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &GalleryApplicationVersionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a gallery Application Version.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// galleryName - The name of the Shared Application Gallery in which the Application Definition resides.
// galleryApplicationName - The name of the gallery Application Definition in which the Application Version is to be created.
// galleryApplicationVersionName - The name of the gallery Application Version to be created. Needs to follow semantic version
// name pattern: The allowed characters are digit and period. Digits must be within the range of a 32-bit
// integer. Format: ..
// galleryApplicationVersion - Parameters supplied to the create or update gallery Application Version operation.
// options - GalleryApplicationVersionsClientBeginCreateOrUpdateOptions contains the optional parameters for the GalleryApplicationVersionsClient.BeginCreateOrUpdate
// method.
func (client *GalleryApplicationVersionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion GalleryApplicationVersion, options *GalleryApplicationVersionsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[GalleryApplicationVersionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName, galleryApplicationVersion, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[GalleryApplicationVersionsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[GalleryApplicationVersionsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update a gallery Application Version.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *GalleryApplicationVersionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion GalleryApplicationVersion, options *GalleryApplicationVersionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName, galleryApplicationVersion, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *GalleryApplicationVersionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion GalleryApplicationVersion, options *GalleryApplicationVersionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}/versions/{galleryApplicationVersionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if galleryApplicationName == "" {
		return nil, errors.New("parameter galleryApplicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationName}", url.PathEscape(galleryApplicationName))
	if galleryApplicationVersionName == "" {
		return nil, errors.New("parameter galleryApplicationVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationVersionName}", url.PathEscape(galleryApplicationVersionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, galleryApplicationVersion)
}

// BeginDelete - Delete a gallery Application Version.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// galleryName - The name of the Shared Application Gallery in which the Application Definition resides.
// galleryApplicationName - The name of the gallery Application Definition in which the Application Version resides.
// galleryApplicationVersionName - The name of the gallery Application Version to be deleted.
// options - GalleryApplicationVersionsClientBeginDeleteOptions contains the optional parameters for the GalleryApplicationVersionsClient.BeginDelete
// method.
func (client *GalleryApplicationVersionsClient) BeginDelete(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, options *GalleryApplicationVersionsClientBeginDeleteOptions) (*armruntime.Poller[GalleryApplicationVersionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[GalleryApplicationVersionsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[GalleryApplicationVersionsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete a gallery Application Version.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *GalleryApplicationVersionsClient) deleteOperation(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, options *GalleryApplicationVersionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName, options)
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
func (client *GalleryApplicationVersionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, options *GalleryApplicationVersionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}/versions/{galleryApplicationVersionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if galleryApplicationName == "" {
		return nil, errors.New("parameter galleryApplicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationName}", url.PathEscape(galleryApplicationName))
	if galleryApplicationVersionName == "" {
		return nil, errors.New("parameter galleryApplicationVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationVersionName}", url.PathEscape(galleryApplicationVersionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Retrieves information about a gallery Application Version.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// galleryName - The name of the Shared Application Gallery in which the Application Definition resides.
// galleryApplicationName - The name of the gallery Application Definition in which the Application Version resides.
// galleryApplicationVersionName - The name of the gallery Application Version to be retrieved.
// options - GalleryApplicationVersionsClientGetOptions contains the optional parameters for the GalleryApplicationVersionsClient.Get
// method.
func (client *GalleryApplicationVersionsClient) Get(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, options *GalleryApplicationVersionsClientGetOptions) (GalleryApplicationVersionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName, options)
	if err != nil {
		return GalleryApplicationVersionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GalleryApplicationVersionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GalleryApplicationVersionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *GalleryApplicationVersionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, options *GalleryApplicationVersionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}/versions/{galleryApplicationVersionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if galleryApplicationName == "" {
		return nil, errors.New("parameter galleryApplicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationName}", url.PathEscape(galleryApplicationName))
	if galleryApplicationVersionName == "" {
		return nil, errors.New("parameter galleryApplicationVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationVersionName}", url.PathEscape(galleryApplicationVersionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GalleryApplicationVersionsClient) getHandleResponse(resp *http.Response) (GalleryApplicationVersionsClientGetResponse, error) {
	result := GalleryApplicationVersionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GalleryApplicationVersion); err != nil {
		return GalleryApplicationVersionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByGalleryApplicationPager - List gallery Application Versions in a gallery Application Definition.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// galleryName - The name of the Shared Application Gallery in which the Application Definition resides.
// galleryApplicationName - The name of the Shared Application Gallery Application Definition from which the Application Versions
// are to be listed.
// options - GalleryApplicationVersionsClientListByGalleryApplicationOptions contains the optional parameters for the GalleryApplicationVersionsClient.ListByGalleryApplication
// method.
func (client *GalleryApplicationVersionsClient) NewListByGalleryApplicationPager(resourceGroupName string, galleryName string, galleryApplicationName string, options *GalleryApplicationVersionsClientListByGalleryApplicationOptions) *runtime.Pager[GalleryApplicationVersionsClientListByGalleryApplicationResponse] {
	return runtime.NewPager(runtime.PageProcessor[GalleryApplicationVersionsClientListByGalleryApplicationResponse]{
		More: func(page GalleryApplicationVersionsClientListByGalleryApplicationResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GalleryApplicationVersionsClientListByGalleryApplicationResponse) (GalleryApplicationVersionsClientListByGalleryApplicationResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByGalleryApplicationCreateRequest(ctx, resourceGroupName, galleryName, galleryApplicationName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return GalleryApplicationVersionsClientListByGalleryApplicationResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return GalleryApplicationVersionsClientListByGalleryApplicationResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return GalleryApplicationVersionsClientListByGalleryApplicationResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByGalleryApplicationHandleResponse(resp)
		},
	})
}

// listByGalleryApplicationCreateRequest creates the ListByGalleryApplication request.
func (client *GalleryApplicationVersionsClient) listByGalleryApplicationCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, options *GalleryApplicationVersionsClientListByGalleryApplicationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}/versions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if galleryApplicationName == "" {
		return nil, errors.New("parameter galleryApplicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationName}", url.PathEscape(galleryApplicationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByGalleryApplicationHandleResponse handles the ListByGalleryApplication response.
func (client *GalleryApplicationVersionsClient) listByGalleryApplicationHandleResponse(resp *http.Response) (GalleryApplicationVersionsClientListByGalleryApplicationResponse, error) {
	result := GalleryApplicationVersionsClientListByGalleryApplicationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GalleryApplicationVersionList); err != nil {
		return GalleryApplicationVersionsClientListByGalleryApplicationResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a gallery Application Version.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// galleryName - The name of the Shared Application Gallery in which the Application Definition resides.
// galleryApplicationName - The name of the gallery Application Definition in which the Application Version is to be updated.
// galleryApplicationVersionName - The name of the gallery Application Version to be updated. Needs to follow semantic version
// name pattern: The allowed characters are digit and period. Digits must be within the range of a 32-bit
// integer. Format: ..
// galleryApplicationVersion - Parameters supplied to the update gallery Application Version operation.
// options - GalleryApplicationVersionsClientBeginUpdateOptions contains the optional parameters for the GalleryApplicationVersionsClient.BeginUpdate
// method.
func (client *GalleryApplicationVersionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion GalleryApplicationVersionUpdate, options *GalleryApplicationVersionsClientBeginUpdateOptions) (*armruntime.Poller[GalleryApplicationVersionsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName, galleryApplicationVersion, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[GalleryApplicationVersionsClientUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[GalleryApplicationVersionsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Update a gallery Application Version.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *GalleryApplicationVersionsClient) update(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion GalleryApplicationVersionUpdate, options *GalleryApplicationVersionsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName, galleryApplicationVersion, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *GalleryApplicationVersionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion GalleryApplicationVersionUpdate, options *GalleryApplicationVersionsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}/versions/{galleryApplicationVersionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if galleryApplicationName == "" {
		return nil, errors.New("parameter galleryApplicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationName}", url.PathEscape(galleryApplicationName))
	if galleryApplicationVersionName == "" {
		return nil, errors.New("parameter galleryApplicationVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationVersionName}", url.PathEscape(galleryApplicationVersionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, galleryApplicationVersion)
}
