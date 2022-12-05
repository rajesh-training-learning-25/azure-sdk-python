//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybriddatamanager

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

// JobDefinitionsClient contains the methods for the JobDefinitions group.
// Don't use this type directly, use NewJobDefinitionsClient() instead.
type JobDefinitionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewJobDefinitionsClient creates a new instance of JobDefinitionsClient with the specified values.
// subscriptionID - The Subscription Id
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewJobDefinitionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*JobDefinitionsClient, error) {
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
	client := &JobDefinitionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a job definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
// dataServiceName - The data service type of the job definition.
// jobDefinitionName - The job definition name to be created or updated.
// resourceGroupName - The Resource Group Name
// dataManagerName - The name of the DataManager Resource within the specified resource group. DataManager names must be between
// 3 and 24 characters in length and use any alphanumeric and underscore only
// jobDefinition - Job Definition object to be created or updated.
// options - JobDefinitionsClientBeginCreateOrUpdateOptions contains the optional parameters for the JobDefinitionsClient.BeginCreateOrUpdate
// method.
func (client *JobDefinitionsClient) BeginCreateOrUpdate(ctx context.Context, dataServiceName string, jobDefinitionName string, resourceGroupName string, dataManagerName string, jobDefinition JobDefinition, options *JobDefinitionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[JobDefinitionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, dataServiceName, jobDefinitionName, resourceGroupName, dataManagerName, jobDefinition, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[JobDefinitionsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[JobDefinitionsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a job definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
func (client *JobDefinitionsClient) createOrUpdate(ctx context.Context, dataServiceName string, jobDefinitionName string, resourceGroupName string, dataManagerName string, jobDefinition JobDefinition, options *JobDefinitionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, dataServiceName, jobDefinitionName, resourceGroupName, dataManagerName, jobDefinition, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *JobDefinitionsClient) createOrUpdateCreateRequest(ctx context.Context, dataServiceName string, jobDefinitionName string, resourceGroupName string, dataManagerName string, jobDefinition JobDefinition, options *JobDefinitionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}"
	if dataServiceName == "" {
		return nil, errors.New("parameter dataServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataServiceName}", url.PathEscape(dataServiceName))
	if jobDefinitionName == "" {
		return nil, errors.New("parameter jobDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobDefinitionName}", url.PathEscape(jobDefinitionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataManagerName == "" {
		return nil, errors.New("parameter dataManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataManagerName}", url.PathEscape(dataManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, jobDefinition)
}

// BeginDelete - This method deletes the given job definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
// dataServiceName - The data service type of the job definition.
// jobDefinitionName - The job definition name to be deleted.
// resourceGroupName - The Resource Group Name
// dataManagerName - The name of the DataManager Resource within the specified resource group. DataManager names must be between
// 3 and 24 characters in length and use any alphanumeric and underscore only
// options - JobDefinitionsClientBeginDeleteOptions contains the optional parameters for the JobDefinitionsClient.BeginDelete
// method.
func (client *JobDefinitionsClient) BeginDelete(ctx context.Context, dataServiceName string, jobDefinitionName string, resourceGroupName string, dataManagerName string, options *JobDefinitionsClientBeginDeleteOptions) (*runtime.Poller[JobDefinitionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, dataServiceName, jobDefinitionName, resourceGroupName, dataManagerName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[JobDefinitionsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[JobDefinitionsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - This method deletes the given job definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
func (client *JobDefinitionsClient) deleteOperation(ctx context.Context, dataServiceName string, jobDefinitionName string, resourceGroupName string, dataManagerName string, options *JobDefinitionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, dataServiceName, jobDefinitionName, resourceGroupName, dataManagerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *JobDefinitionsClient) deleteCreateRequest(ctx context.Context, dataServiceName string, jobDefinitionName string, resourceGroupName string, dataManagerName string, options *JobDefinitionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}"
	if dataServiceName == "" {
		return nil, errors.New("parameter dataServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataServiceName}", url.PathEscape(dataServiceName))
	if jobDefinitionName == "" {
		return nil, errors.New("parameter jobDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobDefinitionName}", url.PathEscape(jobDefinitionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataManagerName == "" {
		return nil, errors.New("parameter dataManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataManagerName}", url.PathEscape(dataManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - This method gets job definition object by name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
// dataServiceName - The data service name of the job definition
// jobDefinitionName - The job definition name that is being queried.
// resourceGroupName - The Resource Group Name
// dataManagerName - The name of the DataManager Resource within the specified resource group. DataManager names must be between
// 3 and 24 characters in length and use any alphanumeric and underscore only
// options - JobDefinitionsClientGetOptions contains the optional parameters for the JobDefinitionsClient.Get method.
func (client *JobDefinitionsClient) Get(ctx context.Context, dataServiceName string, jobDefinitionName string, resourceGroupName string, dataManagerName string, options *JobDefinitionsClientGetOptions) (JobDefinitionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, dataServiceName, jobDefinitionName, resourceGroupName, dataManagerName, options)
	if err != nil {
		return JobDefinitionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobDefinitionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobDefinitionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *JobDefinitionsClient) getCreateRequest(ctx context.Context, dataServiceName string, jobDefinitionName string, resourceGroupName string, dataManagerName string, options *JobDefinitionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}"
	if dataServiceName == "" {
		return nil, errors.New("parameter dataServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataServiceName}", url.PathEscape(dataServiceName))
	if jobDefinitionName == "" {
		return nil, errors.New("parameter jobDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobDefinitionName}", url.PathEscape(jobDefinitionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataManagerName == "" {
		return nil, errors.New("parameter dataManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataManagerName}", url.PathEscape(dataManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *JobDefinitionsClient) getHandleResponse(resp *http.Response) (JobDefinitionsClientGetResponse, error) {
	result := JobDefinitionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobDefinition); err != nil {
		return JobDefinitionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDataManagerPager - This method gets all the job definitions of the given data manager resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
// resourceGroupName - The Resource Group Name
// dataManagerName - The name of the DataManager Resource within the specified resource group. DataManager names must be between
// 3 and 24 characters in length and use any alphanumeric and underscore only
// options - JobDefinitionsClientListByDataManagerOptions contains the optional parameters for the JobDefinitionsClient.ListByDataManager
// method.
func (client *JobDefinitionsClient) NewListByDataManagerPager(resourceGroupName string, dataManagerName string, options *JobDefinitionsClientListByDataManagerOptions) *runtime.Pager[JobDefinitionsClientListByDataManagerResponse] {
	return runtime.NewPager(runtime.PagingHandler[JobDefinitionsClientListByDataManagerResponse]{
		More: func(page JobDefinitionsClientListByDataManagerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *JobDefinitionsClientListByDataManagerResponse) (JobDefinitionsClientListByDataManagerResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByDataManagerCreateRequest(ctx, resourceGroupName, dataManagerName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return JobDefinitionsClientListByDataManagerResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return JobDefinitionsClientListByDataManagerResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return JobDefinitionsClientListByDataManagerResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDataManagerHandleResponse(resp)
		},
	})
}

// listByDataManagerCreateRequest creates the ListByDataManager request.
func (client *JobDefinitionsClient) listByDataManagerCreateRequest(ctx context.Context, resourceGroupName string, dataManagerName string, options *JobDefinitionsClientListByDataManagerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/jobDefinitions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataManagerName == "" {
		return nil, errors.New("parameter dataManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataManagerName}", url.PathEscape(dataManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDataManagerHandleResponse handles the ListByDataManager response.
func (client *JobDefinitionsClient) listByDataManagerHandleResponse(resp *http.Response) (JobDefinitionsClientListByDataManagerResponse, error) {
	result := JobDefinitionsClientListByDataManagerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobDefinitionList); err != nil {
		return JobDefinitionsClientListByDataManagerResponse{}, err
	}
	return result, nil
}

// NewListByDataServicePager - This method gets all the job definitions of the given data service name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
// dataServiceName - The data service type of interest.
// resourceGroupName - The Resource Group Name
// dataManagerName - The name of the DataManager Resource within the specified resource group. DataManager names must be between
// 3 and 24 characters in length and use any alphanumeric and underscore only
// options - JobDefinitionsClientListByDataServiceOptions contains the optional parameters for the JobDefinitionsClient.ListByDataService
// method.
func (client *JobDefinitionsClient) NewListByDataServicePager(dataServiceName string, resourceGroupName string, dataManagerName string, options *JobDefinitionsClientListByDataServiceOptions) *runtime.Pager[JobDefinitionsClientListByDataServiceResponse] {
	return runtime.NewPager(runtime.PagingHandler[JobDefinitionsClientListByDataServiceResponse]{
		More: func(page JobDefinitionsClientListByDataServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *JobDefinitionsClientListByDataServiceResponse) (JobDefinitionsClientListByDataServiceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByDataServiceCreateRequest(ctx, dataServiceName, resourceGroupName, dataManagerName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return JobDefinitionsClientListByDataServiceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return JobDefinitionsClientListByDataServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return JobDefinitionsClientListByDataServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDataServiceHandleResponse(resp)
		},
	})
}

// listByDataServiceCreateRequest creates the ListByDataService request.
func (client *JobDefinitionsClient) listByDataServiceCreateRequest(ctx context.Context, dataServiceName string, resourceGroupName string, dataManagerName string, options *JobDefinitionsClientListByDataServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions"
	if dataServiceName == "" {
		return nil, errors.New("parameter dataServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataServiceName}", url.PathEscape(dataServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataManagerName == "" {
		return nil, errors.New("parameter dataManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataManagerName}", url.PathEscape(dataManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDataServiceHandleResponse handles the ListByDataService response.
func (client *JobDefinitionsClient) listByDataServiceHandleResponse(resp *http.Response) (JobDefinitionsClientListByDataServiceResponse, error) {
	result := JobDefinitionsClientListByDataServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobDefinitionList); err != nil {
		return JobDefinitionsClientListByDataServiceResponse{}, err
	}
	return result, nil
}

// BeginRun - This method runs a job instance of the given job definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
// dataServiceName - The data service type of the job definition.
// jobDefinitionName - Name of the job definition.
// resourceGroupName - The Resource Group Name
// dataManagerName - The name of the DataManager Resource within the specified resource group. DataManager names must be between
// 3 and 24 characters in length and use any alphanumeric and underscore only
// runParameters - Run time parameters for the job definition.
// options - JobDefinitionsClientBeginRunOptions contains the optional parameters for the JobDefinitionsClient.BeginRun method.
func (client *JobDefinitionsClient) BeginRun(ctx context.Context, dataServiceName string, jobDefinitionName string, resourceGroupName string, dataManagerName string, runParameters RunParameters, options *JobDefinitionsClientBeginRunOptions) (*runtime.Poller[JobDefinitionsClientRunResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.run(ctx, dataServiceName, jobDefinitionName, resourceGroupName, dataManagerName, runParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[JobDefinitionsClientRunResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[JobDefinitionsClientRunResponse](options.ResumeToken, client.pl, nil)
	}
}

// Run - This method runs a job instance of the given job definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
func (client *JobDefinitionsClient) run(ctx context.Context, dataServiceName string, jobDefinitionName string, resourceGroupName string, dataManagerName string, runParameters RunParameters, options *JobDefinitionsClientBeginRunOptions) (*http.Response, error) {
	req, err := client.runCreateRequest(ctx, dataServiceName, jobDefinitionName, resourceGroupName, dataManagerName, runParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// runCreateRequest creates the Run request.
func (client *JobDefinitionsClient) runCreateRequest(ctx context.Context, dataServiceName string, jobDefinitionName string, resourceGroupName string, dataManagerName string, runParameters RunParameters, options *JobDefinitionsClientBeginRunOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}/run"
	if dataServiceName == "" {
		return nil, errors.New("parameter dataServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataServiceName}", url.PathEscape(dataServiceName))
	if jobDefinitionName == "" {
		return nil, errors.New("parameter jobDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobDefinitionName}", url.PathEscape(jobDefinitionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataManagerName == "" {
		return nil, errors.New("parameter dataManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataManagerName}", url.PathEscape(dataManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, runtime.MarshalAsJSON(req, runParameters)
}