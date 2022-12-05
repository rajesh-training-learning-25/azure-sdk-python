//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// JobStepsClient contains the methods for the JobSteps group.
// Don't use this type directly, use NewJobStepsClient() instead.
type JobStepsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewJobStepsClient creates a new instance of JobStepsClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewJobStepsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*JobStepsClient, error) {
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
	client := &JobStepsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a job step. This will implicitly create a new job version.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// jobAgentName - The name of the job agent.
// jobName - The name of the job.
// stepName - The name of the job step.
// parameters - The requested state of the job step.
// options - JobStepsClientCreateOrUpdateOptions contains the optional parameters for the JobStepsClient.CreateOrUpdate method.
func (client *JobStepsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, stepName string, parameters JobStep, options *JobStepsClientCreateOrUpdateOptions) (JobStepsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, jobName, stepName, parameters, options)
	if err != nil {
		return JobStepsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobStepsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return JobStepsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *JobStepsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, stepName string, parameters JobStep, options *JobStepsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/steps/{stepName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if jobAgentName == "" {
		return nil, errors.New("parameter jobAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobAgentName}", url.PathEscape(jobAgentName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if stepName == "" {
		return nil, errors.New("parameter stepName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{stepName}", url.PathEscape(stepName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *JobStepsClient) createOrUpdateHandleResponse(resp *http.Response) (JobStepsClientCreateOrUpdateResponse, error) {
	result := JobStepsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobStep); err != nil {
		return JobStepsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a job step. This will implicitly create a new job version.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// jobAgentName - The name of the job agent.
// jobName - The name of the job.
// stepName - The name of the job step to delete.
// options - JobStepsClientDeleteOptions contains the optional parameters for the JobStepsClient.Delete method.
func (client *JobStepsClient) Delete(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, stepName string, options *JobStepsClientDeleteOptions) (JobStepsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, jobName, stepName, options)
	if err != nil {
		return JobStepsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobStepsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return JobStepsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return JobStepsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *JobStepsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, stepName string, options *JobStepsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/steps/{stepName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if jobAgentName == "" {
		return nil, errors.New("parameter jobAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobAgentName}", url.PathEscape(jobAgentName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if stepName == "" {
		return nil, errors.New("parameter stepName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{stepName}", url.PathEscape(stepName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a job step in a job's current version.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// jobAgentName - The name of the job agent.
// jobName - The name of the job.
// stepName - The name of the job step.
// options - JobStepsClientGetOptions contains the optional parameters for the JobStepsClient.Get method.
func (client *JobStepsClient) Get(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, stepName string, options *JobStepsClientGetOptions) (JobStepsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, jobName, stepName, options)
	if err != nil {
		return JobStepsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobStepsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobStepsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *JobStepsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, stepName string, options *JobStepsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/steps/{stepName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if jobAgentName == "" {
		return nil, errors.New("parameter jobAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobAgentName}", url.PathEscape(jobAgentName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if stepName == "" {
		return nil, errors.New("parameter stepName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{stepName}", url.PathEscape(stepName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *JobStepsClient) getHandleResponse(resp *http.Response) (JobStepsClientGetResponse, error) {
	result := JobStepsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobStep); err != nil {
		return JobStepsClientGetResponse{}, err
	}
	return result, nil
}

// GetByVersion - Gets the specified version of a job step.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// jobAgentName - The name of the job agent.
// jobName - The name of the job.
// jobVersion - The version of the job to get.
// stepName - The name of the job step.
// options - JobStepsClientGetByVersionOptions contains the optional parameters for the JobStepsClient.GetByVersion method.
func (client *JobStepsClient) GetByVersion(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobVersion int32, stepName string, options *JobStepsClientGetByVersionOptions) (JobStepsClientGetByVersionResponse, error) {
	req, err := client.getByVersionCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, jobName, jobVersion, stepName, options)
	if err != nil {
		return JobStepsClientGetByVersionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobStepsClientGetByVersionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobStepsClientGetByVersionResponse{}, runtime.NewResponseError(resp)
	}
	return client.getByVersionHandleResponse(resp)
}

// getByVersionCreateRequest creates the GetByVersion request.
func (client *JobStepsClient) getByVersionCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobVersion int32, stepName string, options *JobStepsClientGetByVersionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/versions/{jobVersion}/steps/{stepName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if jobAgentName == "" {
		return nil, errors.New("parameter jobAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobAgentName}", url.PathEscape(jobAgentName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	urlPath = strings.ReplaceAll(urlPath, "{jobVersion}", url.PathEscape(strconv.FormatInt(int64(jobVersion), 10)))
	if stepName == "" {
		return nil, errors.New("parameter stepName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{stepName}", url.PathEscape(stepName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByVersionHandleResponse handles the GetByVersion response.
func (client *JobStepsClient) getByVersionHandleResponse(resp *http.Response) (JobStepsClientGetByVersionResponse, error) {
	result := JobStepsClientGetByVersionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobStep); err != nil {
		return JobStepsClientGetByVersionResponse{}, err
	}
	return result, nil
}

// NewListByJobPager - Gets all job steps for a job's current version.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// jobAgentName - The name of the job agent.
// jobName - The name of the job to get.
// options - JobStepsClientListByJobOptions contains the optional parameters for the JobStepsClient.ListByJob method.
func (client *JobStepsClient) NewListByJobPager(resourceGroupName string, serverName string, jobAgentName string, jobName string, options *JobStepsClientListByJobOptions) *runtime.Pager[JobStepsClientListByJobResponse] {
	return runtime.NewPager(runtime.PagingHandler[JobStepsClientListByJobResponse]{
		More: func(page JobStepsClientListByJobResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *JobStepsClientListByJobResponse) (JobStepsClientListByJobResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByJobCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, jobName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return JobStepsClientListByJobResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return JobStepsClientListByJobResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return JobStepsClientListByJobResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByJobHandleResponse(resp)
		},
	})
}

// listByJobCreateRequest creates the ListByJob request.
func (client *JobStepsClient) listByJobCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, options *JobStepsClientListByJobOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/steps"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if jobAgentName == "" {
		return nil, errors.New("parameter jobAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobAgentName}", url.PathEscape(jobAgentName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByJobHandleResponse handles the ListByJob response.
func (client *JobStepsClient) listByJobHandleResponse(resp *http.Response) (JobStepsClientListByJobResponse, error) {
	result := JobStepsClientListByJobResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobStepListResult); err != nil {
		return JobStepsClientListByJobResponse{}, err
	}
	return result, nil
}

// NewListByVersionPager - Gets all job steps in the specified job version.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// jobAgentName - The name of the job agent.
// jobName - The name of the job to get.
// jobVersion - The version of the job to get.
// options - JobStepsClientListByVersionOptions contains the optional parameters for the JobStepsClient.ListByVersion method.
func (client *JobStepsClient) NewListByVersionPager(resourceGroupName string, serverName string, jobAgentName string, jobName string, jobVersion int32, options *JobStepsClientListByVersionOptions) *runtime.Pager[JobStepsClientListByVersionResponse] {
	return runtime.NewPager(runtime.PagingHandler[JobStepsClientListByVersionResponse]{
		More: func(page JobStepsClientListByVersionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *JobStepsClientListByVersionResponse) (JobStepsClientListByVersionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByVersionCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, jobName, jobVersion, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return JobStepsClientListByVersionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return JobStepsClientListByVersionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return JobStepsClientListByVersionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByVersionHandleResponse(resp)
		},
	})
}

// listByVersionCreateRequest creates the ListByVersion request.
func (client *JobStepsClient) listByVersionCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobVersion int32, options *JobStepsClientListByVersionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/versions/{jobVersion}/steps"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if jobAgentName == "" {
		return nil, errors.New("parameter jobAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobAgentName}", url.PathEscape(jobAgentName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	urlPath = strings.ReplaceAll(urlPath, "{jobVersion}", url.PathEscape(strconv.FormatInt(int64(jobVersion), 10)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByVersionHandleResponse handles the ListByVersion response.
func (client *JobStepsClient) listByVersionHandleResponse(resp *http.Response) (JobStepsClientListByVersionResponse, error) {
	result := JobStepsClientListByVersionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobStepListResult); err != nil {
		return JobStepsClientListByVersionResponse{}, err
	}
	return result, nil
}