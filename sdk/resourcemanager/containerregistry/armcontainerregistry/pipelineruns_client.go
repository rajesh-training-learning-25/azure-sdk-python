//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerregistry

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

// PipelineRunsClient contains the methods for the PipelineRuns group.
// Don't use this type directly, use NewPipelineRunsClient() instead.
type PipelineRunsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPipelineRunsClient creates a new instance of PipelineRunsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPipelineRunsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PipelineRunsClient, error) {
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
	client := &PipelineRunsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Creates a pipeline run for a container registry with the specified parameters
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - pipelineRunName - The name of the pipeline run.
//   - pipelineRunCreateParameters - The parameters for creating a pipeline run.
//   - options - PipelineRunsClientBeginCreateOptions contains the optional parameters for the PipelineRunsClient.BeginCreate
//     method.
func (client *PipelineRunsClient) BeginCreate(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, pipelineRunCreateParameters PipelineRun, options *PipelineRunsClientBeginCreateOptions) (*runtime.Poller[PipelineRunsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, registryName, pipelineRunName, pipelineRunCreateParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[PipelineRunsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[PipelineRunsClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Creates a pipeline run for a container registry with the specified parameters
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
func (client *PipelineRunsClient) create(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, pipelineRunCreateParameters PipelineRun, options *PipelineRunsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, registryName, pipelineRunName, pipelineRunCreateParameters, options)
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

// createCreateRequest creates the Create request.
func (client *PipelineRunsClient) createCreateRequest(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, pipelineRunCreateParameters PipelineRun, options *PipelineRunsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/pipelineRuns/{pipelineRunName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if pipelineRunName == "" {
		return nil, errors.New("parameter pipelineRunName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineRunName}", url.PathEscape(pipelineRunName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, pipelineRunCreateParameters)
}

// BeginDelete - Deletes a pipeline run from a container registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - pipelineRunName - The name of the pipeline run.
//   - options - PipelineRunsClientBeginDeleteOptions contains the optional parameters for the PipelineRunsClient.BeginDelete
//     method.
func (client *PipelineRunsClient) BeginDelete(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, options *PipelineRunsClientBeginDeleteOptions) (*runtime.Poller[PipelineRunsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, registryName, pipelineRunName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[PipelineRunsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[PipelineRunsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a pipeline run from a container registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
func (client *PipelineRunsClient) deleteOperation(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, options *PipelineRunsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, registryName, pipelineRunName, options)
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
func (client *PipelineRunsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, options *PipelineRunsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/pipelineRuns/{pipelineRunName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if pipelineRunName == "" {
		return nil, errors.New("parameter pipelineRunName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineRunName}", url.PathEscape(pipelineRunName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the detailed information for a given pipeline run.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - pipelineRunName - The name of the pipeline run.
//   - options - PipelineRunsClientGetOptions contains the optional parameters for the PipelineRunsClient.Get method.
func (client *PipelineRunsClient) Get(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, options *PipelineRunsClientGetOptions) (PipelineRunsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, pipelineRunName, options)
	if err != nil {
		return PipelineRunsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PipelineRunsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PipelineRunsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PipelineRunsClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, options *PipelineRunsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/pipelineRuns/{pipelineRunName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if pipelineRunName == "" {
		return nil, errors.New("parameter pipelineRunName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineRunName}", url.PathEscape(pipelineRunName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PipelineRunsClient) getHandleResponse(resp *http.Response) (PipelineRunsClientGetResponse, error) {
	result := PipelineRunsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineRun); err != nil {
		return PipelineRunsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all the pipeline runs for the specified container registry.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - options - PipelineRunsClientListOptions contains the optional parameters for the PipelineRunsClient.NewListPager method.
func (client *PipelineRunsClient) NewListPager(resourceGroupName string, registryName string, options *PipelineRunsClientListOptions) *runtime.Pager[PipelineRunsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PipelineRunsClientListResponse]{
		More: func(page PipelineRunsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PipelineRunsClientListResponse) (PipelineRunsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, registryName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PipelineRunsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PipelineRunsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PipelineRunsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *PipelineRunsClient) listCreateRequest(ctx context.Context, resourceGroupName string, registryName string, options *PipelineRunsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/pipelineRuns"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PipelineRunsClient) listHandleResponse(resp *http.Response) (PipelineRunsClientListResponse, error) {
	result := PipelineRunsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineRunListResult); err != nil {
		return PipelineRunsClientListResponse{}, err
	}
	return result, nil
}
