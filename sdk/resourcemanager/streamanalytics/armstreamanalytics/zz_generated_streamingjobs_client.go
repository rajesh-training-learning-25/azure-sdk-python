//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstreamanalytics

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

// StreamingJobsClient contains the methods for the StreamingJobs group.
// Don't use this type directly, use NewStreamingJobsClient() instead.
type StreamingJobsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewStreamingJobsClient creates a new instance of StreamingJobsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewStreamingJobsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*StreamingJobsClient, error) {
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
	client := &StreamingJobsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrReplace - Creates a streaming job or replaces an already existing streaming job.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// jobName - The name of the streaming job.
// streamingJob - The definition of the streaming job that will be used to create a new streaming job or replace the existing
// one.
// options - StreamingJobsClientBeginCreateOrReplaceOptions contains the optional parameters for the StreamingJobsClient.BeginCreateOrReplace
// method.
func (client *StreamingJobsClient) BeginCreateOrReplace(ctx context.Context, resourceGroupName string, jobName string, streamingJob StreamingJob, options *StreamingJobsClientBeginCreateOrReplaceOptions) (*runtime.Poller[StreamingJobsClientCreateOrReplaceResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrReplace(ctx, resourceGroupName, jobName, streamingJob, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[StreamingJobsClientCreateOrReplaceResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[StreamingJobsClientCreateOrReplaceResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrReplace - Creates a streaming job or replaces an already existing streaming job.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
func (client *StreamingJobsClient) createOrReplace(ctx context.Context, resourceGroupName string, jobName string, streamingJob StreamingJob, options *StreamingJobsClientBeginCreateOrReplaceOptions) (*http.Response, error) {
	req, err := client.createOrReplaceCreateRequest(ctx, resourceGroupName, jobName, streamingJob, options)
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

// createOrReplaceCreateRequest creates the CreateOrReplace request.
func (client *StreamingJobsClient) createOrReplaceCreateRequest(ctx context.Context, resourceGroupName string, jobName string, streamingJob StreamingJob, options *StreamingJobsClientBeginCreateOrReplaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, streamingJob)
}

// BeginDelete - Deletes a streaming job.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// jobName - The name of the streaming job.
// options - StreamingJobsClientBeginDeleteOptions contains the optional parameters for the StreamingJobsClient.BeginDelete
// method.
func (client *StreamingJobsClient) BeginDelete(ctx context.Context, resourceGroupName string, jobName string, options *StreamingJobsClientBeginDeleteOptions) (*runtime.Poller[StreamingJobsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, jobName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[StreamingJobsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[StreamingJobsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a streaming job.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
func (client *StreamingJobsClient) deleteOperation(ctx context.Context, resourceGroupName string, jobName string, options *StreamingJobsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, jobName, options)
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
func (client *StreamingJobsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, jobName string, options *StreamingJobsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets details about the specified streaming job.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// jobName - The name of the streaming job.
// options - StreamingJobsClientGetOptions contains the optional parameters for the StreamingJobsClient.Get method.
func (client *StreamingJobsClient) Get(ctx context.Context, resourceGroupName string, jobName string, options *StreamingJobsClientGetOptions) (StreamingJobsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, jobName, options)
	if err != nil {
		return StreamingJobsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StreamingJobsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StreamingJobsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *StreamingJobsClient) getCreateRequest(ctx context.Context, resourceGroupName string, jobName string, options *StreamingJobsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *StreamingJobsClient) getHandleResponse(resp *http.Response) (StreamingJobsClientGetResponse, error) {
	result := StreamingJobsClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.StreamingJob); err != nil {
		return StreamingJobsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all of the streaming jobs in the given subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// options - StreamingJobsClientListOptions contains the optional parameters for the StreamingJobsClient.List method.
func (client *StreamingJobsClient) NewListPager(options *StreamingJobsClientListOptions) *runtime.Pager[StreamingJobsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[StreamingJobsClientListResponse]{
		More: func(page StreamingJobsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *StreamingJobsClientListResponse) (StreamingJobsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return StreamingJobsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return StreamingJobsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return StreamingJobsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *StreamingJobsClient) listCreateRequest(ctx context.Context, options *StreamingJobsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.StreamAnalytics/streamingjobs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *StreamingJobsClient) listHandleResponse(resp *http.Response) (StreamingJobsClientListResponse, error) {
	result := StreamingJobsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StreamingJobListResult); err != nil {
		return StreamingJobsClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists all of the streaming jobs in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - StreamingJobsClientListByResourceGroupOptions contains the optional parameters for the StreamingJobsClient.ListByResourceGroup
// method.
func (client *StreamingJobsClient) NewListByResourceGroupPager(resourceGroupName string, options *StreamingJobsClientListByResourceGroupOptions) *runtime.Pager[StreamingJobsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[StreamingJobsClientListByResourceGroupResponse]{
		More: func(page StreamingJobsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *StreamingJobsClientListByResourceGroupResponse) (StreamingJobsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return StreamingJobsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return StreamingJobsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return StreamingJobsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *StreamingJobsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *StreamingJobsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs"
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
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *StreamingJobsClient) listByResourceGroupHandleResponse(resp *http.Response) (StreamingJobsClientListByResourceGroupResponse, error) {
	result := StreamingJobsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StreamingJobListResult); err != nil {
		return StreamingJobsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginScale - Scales a streaming job when the job is running.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// jobName - The name of the streaming job.
// options - StreamingJobsClientBeginScaleOptions contains the optional parameters for the StreamingJobsClient.BeginScale
// method.
func (client *StreamingJobsClient) BeginScale(ctx context.Context, resourceGroupName string, jobName string, options *StreamingJobsClientBeginScaleOptions) (*runtime.Poller[StreamingJobsClientScaleResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.scale(ctx, resourceGroupName, jobName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[StreamingJobsClientScaleResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[StreamingJobsClientScaleResponse](options.ResumeToken, client.pl, nil)
	}
}

// Scale - Scales a streaming job when the job is running.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
func (client *StreamingJobsClient) scale(ctx context.Context, resourceGroupName string, jobName string, options *StreamingJobsClientBeginScaleOptions) (*http.Response, error) {
	req, err := client.scaleCreateRequest(ctx, resourceGroupName, jobName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// scaleCreateRequest creates the Scale request.
func (client *StreamingJobsClient) scaleCreateRequest(ctx context.Context, resourceGroupName string, jobName string, options *StreamingJobsClientBeginScaleOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/scale"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.ScaleJobParameters != nil {
		return req, runtime.MarshalAsJSON(req, *options.ScaleJobParameters)
	}
	return req, nil
}

// BeginStart - Starts a streaming job. Once a job is started it will start processing input events and produce output.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// jobName - The name of the streaming job.
// options - StreamingJobsClientBeginStartOptions contains the optional parameters for the StreamingJobsClient.BeginStart
// method.
func (client *StreamingJobsClient) BeginStart(ctx context.Context, resourceGroupName string, jobName string, options *StreamingJobsClientBeginStartOptions) (*runtime.Poller[StreamingJobsClientStartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.start(ctx, resourceGroupName, jobName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[StreamingJobsClientStartResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[StreamingJobsClientStartResponse](options.ResumeToken, client.pl, nil)
	}
}

// Start - Starts a streaming job. Once a job is started it will start processing input events and produce output.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
func (client *StreamingJobsClient) start(ctx context.Context, resourceGroupName string, jobName string, options *StreamingJobsClientBeginStartOptions) (*http.Response, error) {
	req, err := client.startCreateRequest(ctx, resourceGroupName, jobName, options)
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

// startCreateRequest creates the Start request.
func (client *StreamingJobsClient) startCreateRequest(ctx context.Context, resourceGroupName string, jobName string, options *StreamingJobsClientBeginStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/start"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.StartJobParameters != nil {
		return req, runtime.MarshalAsJSON(req, *options.StartJobParameters)
	}
	return req, nil
}

// BeginStop - Stops a running streaming job. This will cause a running streaming job to stop processing input events and
// producing output.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// jobName - The name of the streaming job.
// options - StreamingJobsClientBeginStopOptions contains the optional parameters for the StreamingJobsClient.BeginStop method.
func (client *StreamingJobsClient) BeginStop(ctx context.Context, resourceGroupName string, jobName string, options *StreamingJobsClientBeginStopOptions) (*runtime.Poller[StreamingJobsClientStopResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.stop(ctx, resourceGroupName, jobName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[StreamingJobsClientStopResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[StreamingJobsClientStopResponse](options.ResumeToken, client.pl, nil)
	}
}

// Stop - Stops a running streaming job. This will cause a running streaming job to stop processing input events and producing
// output.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
func (client *StreamingJobsClient) stop(ctx context.Context, resourceGroupName string, jobName string, options *StreamingJobsClientBeginStopOptions) (*http.Response, error) {
	req, err := client.stopCreateRequest(ctx, resourceGroupName, jobName, options)
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

// stopCreateRequest creates the Stop request.
func (client *StreamingJobsClient) stopCreateRequest(ctx context.Context, resourceGroupName string, jobName string, options *StreamingJobsClientBeginStopOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/stop"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Update - Updates an existing streaming job. This can be used to partially update (ie. update one or two properties) a streaming
// job without affecting the rest the job definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// jobName - The name of the streaming job.
// streamingJob - A streaming job object. The properties specified here will overwrite the corresponding properties in the
// existing streaming job (ie. Those properties will be updated). Any properties that are set to
// null here will mean that the corresponding property in the existing input will remain the same and not change as a result
// of this PATCH operation.
// options - StreamingJobsClientUpdateOptions contains the optional parameters for the StreamingJobsClient.Update method.
func (client *StreamingJobsClient) Update(ctx context.Context, resourceGroupName string, jobName string, streamingJob StreamingJob, options *StreamingJobsClientUpdateOptions) (StreamingJobsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, jobName, streamingJob, options)
	if err != nil {
		return StreamingJobsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StreamingJobsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StreamingJobsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *StreamingJobsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, jobName string, streamingJob StreamingJob, options *StreamingJobsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, streamingJob)
}

// updateHandleResponse handles the Update response.
func (client *StreamingJobsClient) updateHandleResponse(resp *http.Response) (StreamingJobsClientUpdateResponse, error) {
	result := StreamingJobsClientUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.StreamingJob); err != nil {
		return StreamingJobsClientUpdateResponse{}, err
	}
	return result, nil
}