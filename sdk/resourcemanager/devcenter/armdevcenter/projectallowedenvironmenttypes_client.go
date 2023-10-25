//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevcenter

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ProjectAllowedEnvironmentTypesClient contains the methods for the ProjectAllowedEnvironmentTypes group.
// Don't use this type directly, use NewProjectAllowedEnvironmentTypesClient() instead.
type ProjectAllowedEnvironmentTypesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewProjectAllowedEnvironmentTypesClient creates a new instance of ProjectAllowedEnvironmentTypesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewProjectAllowedEnvironmentTypesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ProjectAllowedEnvironmentTypesClient, error) {
	cl, err := arm.NewClient(moduleName+".ProjectAllowedEnvironmentTypesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ProjectAllowedEnvironmentTypesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets an allowed environment type.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - The name of the project.
//   - environmentTypeName - The name of the environment type.
//   - options - ProjectAllowedEnvironmentTypesClientGetOptions contains the optional parameters for the ProjectAllowedEnvironmentTypesClient.Get
//     method.
func (client *ProjectAllowedEnvironmentTypesClient) Get(ctx context.Context, resourceGroupName string, projectName string, environmentTypeName string, options *ProjectAllowedEnvironmentTypesClientGetOptions) (ProjectAllowedEnvironmentTypesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, projectName, environmentTypeName, options)
	if err != nil {
		return ProjectAllowedEnvironmentTypesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProjectAllowedEnvironmentTypesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ProjectAllowedEnvironmentTypesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ProjectAllowedEnvironmentTypesClient) getCreateRequest(ctx context.Context, resourceGroupName string, projectName string, environmentTypeName string, options *ProjectAllowedEnvironmentTypesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/projects/{projectName}/allowedEnvironmentTypes/{environmentTypeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if environmentTypeName == "" {
		return nil, errors.New("parameter environmentTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentTypeName}", url.PathEscape(environmentTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProjectAllowedEnvironmentTypesClient) getHandleResponse(resp *http.Response) (ProjectAllowedEnvironmentTypesClientGetResponse, error) {
	result := ProjectAllowedEnvironmentTypesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AllowedEnvironmentType); err != nil {
		return ProjectAllowedEnvironmentTypesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists allowed environment types for a project.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - The name of the project.
//   - options - ProjectAllowedEnvironmentTypesClientListOptions contains the optional parameters for the ProjectAllowedEnvironmentTypesClient.NewListPager
//     method.
func (client *ProjectAllowedEnvironmentTypesClient) NewListPager(resourceGroupName string, projectName string, options *ProjectAllowedEnvironmentTypesClientListOptions) *runtime.Pager[ProjectAllowedEnvironmentTypesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProjectAllowedEnvironmentTypesClientListResponse]{
		More: func(page ProjectAllowedEnvironmentTypesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProjectAllowedEnvironmentTypesClientListResponse) (ProjectAllowedEnvironmentTypesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, projectName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ProjectAllowedEnvironmentTypesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ProjectAllowedEnvironmentTypesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ProjectAllowedEnvironmentTypesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ProjectAllowedEnvironmentTypesClient) listCreateRequest(ctx context.Context, resourceGroupName string, projectName string, options *ProjectAllowedEnvironmentTypesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/projects/{projectName}/allowedEnvironmentTypes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ProjectAllowedEnvironmentTypesClient) listHandleResponse(resp *http.Response) (ProjectAllowedEnvironmentTypesClientListResponse, error) {
	result := ProjectAllowedEnvironmentTypesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AllowedEnvironmentTypeListResult); err != nil {
		return ProjectAllowedEnvironmentTypesClientListResponse{}, err
	}
	return result, nil
}
