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

// SchedulesClient contains the methods for the Schedules group.
// Don't use this type directly, use NewSchedulesClient() instead.
type SchedulesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSchedulesClient creates a new instance of SchedulesClient with the specified values.
// subscriptionID - Unique identifier of the Azure subscription. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSchedulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SchedulesClient, error) {
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
	client := &SchedulesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a Schedule.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - Name of the resource group within the Azure subscription.
// projectName - The name of the project.
// poolName - Name of the pool.
// scheduleName - The name of the schedule that uniquely identifies it.
// body - Represents a scheduled task
// options - SchedulesClientBeginCreateOrUpdateOptions contains the optional parameters for the SchedulesClient.BeginCreateOrUpdate
// method.
func (client *SchedulesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, projectName string, poolName string, scheduleName string, body Schedule, options *SchedulesClientBeginCreateOrUpdateOptions) (*runtime.Poller[SchedulesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, projectName, poolName, scheduleName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[SchedulesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[SchedulesClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a Schedule.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
func (client *SchedulesClient) createOrUpdate(ctx context.Context, resourceGroupName string, projectName string, poolName string, scheduleName string, body Schedule, options *SchedulesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, projectName, poolName, scheduleName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SchedulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, projectName string, poolName string, scheduleName string, body Schedule, options *SchedulesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/projects/{projectName}/pools/{poolName}/schedules/{scheduleName}"
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
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if scheduleName == "" {
		return nil, errors.New("parameter scheduleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleName}", url.PathEscape(scheduleName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Deletes a Scheduled.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - Name of the resource group within the Azure subscription.
// projectName - The name of the project.
// poolName - Name of the pool.
// scheduleName - The name of the schedule that uniquely identifies it.
// options - SchedulesClientBeginDeleteOptions contains the optional parameters for the SchedulesClient.BeginDelete method.
func (client *SchedulesClient) BeginDelete(ctx context.Context, resourceGroupName string, projectName string, poolName string, scheduleName string, options *SchedulesClientBeginDeleteOptions) (*runtime.Poller[SchedulesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, projectName, poolName, scheduleName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[SchedulesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[SchedulesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a Scheduled.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
func (client *SchedulesClient) deleteOperation(ctx context.Context, resourceGroupName string, projectName string, poolName string, scheduleName string, options *SchedulesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, projectName, poolName, scheduleName, options)
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
func (client *SchedulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, projectName string, poolName string, scheduleName string, options *SchedulesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/projects/{projectName}/pools/{poolName}/schedules/{scheduleName}"
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
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if scheduleName == "" {
		return nil, errors.New("parameter scheduleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleName}", url.PathEscape(scheduleName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a schedule resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - Name of the resource group within the Azure subscription.
// projectName - The name of the project.
// poolName - Name of the pool.
// scheduleName - The name of the schedule that uniquely identifies it.
// options - SchedulesClientGetOptions contains the optional parameters for the SchedulesClient.Get method.
func (client *SchedulesClient) Get(ctx context.Context, resourceGroupName string, projectName string, poolName string, scheduleName string, options *SchedulesClientGetOptions) (SchedulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, projectName, poolName, scheduleName, options)
	if err != nil {
		return SchedulesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SchedulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SchedulesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SchedulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, projectName string, poolName string, scheduleName string, options *SchedulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/projects/{projectName}/pools/{poolName}/schedules/{scheduleName}"
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
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if scheduleName == "" {
		return nil, errors.New("parameter scheduleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleName}", url.PathEscape(scheduleName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SchedulesClient) getHandleResponse(resp *http.Response) (SchedulesClientGetResponse, error) {
	result := SchedulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Schedule); err != nil {
		return SchedulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByPoolPager - Lists schedules for a pool
// Generated from API version 2022-09-01-preview
// resourceGroupName - Name of the resource group within the Azure subscription.
// projectName - The name of the project.
// poolName - Name of the pool.
// options - SchedulesClientListByPoolOptions contains the optional parameters for the SchedulesClient.ListByPool method.
func (client *SchedulesClient) NewListByPoolPager(resourceGroupName string, projectName string, poolName string, options *SchedulesClientListByPoolOptions) *runtime.Pager[SchedulesClientListByPoolResponse] {
	return runtime.NewPager(runtime.PagingHandler[SchedulesClientListByPoolResponse]{
		More: func(page SchedulesClientListByPoolResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SchedulesClientListByPoolResponse) (SchedulesClientListByPoolResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByPoolCreateRequest(ctx, resourceGroupName, projectName, poolName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SchedulesClientListByPoolResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SchedulesClientListByPoolResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SchedulesClientListByPoolResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByPoolHandleResponse(resp)
		},
	})
}

// listByPoolCreateRequest creates the ListByPool request.
func (client *SchedulesClient) listByPoolCreateRequest(ctx context.Context, resourceGroupName string, projectName string, poolName string, options *SchedulesClientListByPoolOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/projects/{projectName}/pools/{poolName}/schedules"
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
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByPoolHandleResponse handles the ListByPool response.
func (client *SchedulesClient) listByPoolHandleResponse(resp *http.Response) (SchedulesClientListByPoolResponse, error) {
	result := SchedulesClientListByPoolResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScheduleListResult); err != nil {
		return SchedulesClientListByPoolResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Partially updates a Scheduled.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - Name of the resource group within the Azure subscription.
// projectName - The name of the project.
// poolName - Name of the pool.
// scheduleName - The name of the schedule that uniquely identifies it.
// body - Represents a scheduled task.
// options - SchedulesClientBeginUpdateOptions contains the optional parameters for the SchedulesClient.BeginUpdate method.
func (client *SchedulesClient) BeginUpdate(ctx context.Context, resourceGroupName string, projectName string, poolName string, scheduleName string, body ScheduleUpdate, options *SchedulesClientBeginUpdateOptions) (*runtime.Poller[SchedulesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, projectName, poolName, scheduleName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[SchedulesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[SchedulesClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Partially updates a Scheduled.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
func (client *SchedulesClient) update(ctx context.Context, resourceGroupName string, projectName string, poolName string, scheduleName string, body ScheduleUpdate, options *SchedulesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, projectName, poolName, scheduleName, body, options)
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

// updateCreateRequest creates the Update request.
func (client *SchedulesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, projectName string, poolName string, scheduleName string, body ScheduleUpdate, options *SchedulesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/projects/{projectName}/pools/{poolName}/schedules/{scheduleName}"
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
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if scheduleName == "" {
		return nil, errors.New("parameter scheduleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleName}", url.PathEscape(scheduleName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}
