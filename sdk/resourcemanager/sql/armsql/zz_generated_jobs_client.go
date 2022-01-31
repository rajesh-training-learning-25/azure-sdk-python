//go:build go1.16
// +build go1.16

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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// JobsClient contains the methods for the Jobs group.
// Don't use this type directly, use NewJobsClient() instead.
type JobsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewJobsClient creates a new instance of JobsClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewJobsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *JobsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &JobsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// CreateOrUpdate - Creates or updates a job.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// jobAgentName - The name of the job agent.
// jobName - The name of the job to get.
// parameters - The requested job state.
// options - JobsClientCreateOrUpdateOptions contains the optional parameters for the JobsClient.CreateOrUpdate method.
func (client *JobsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, parameters Job, options *JobsClientCreateOrUpdateOptions) (JobsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, jobName, parameters, options)
	if err != nil {
		return JobsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return JobsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *JobsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, parameters Job, options *JobsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *JobsClient) createOrUpdateHandleResponse(resp *http.Response) (JobsClientCreateOrUpdateResponse, error) {
	result := JobsClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Job); err != nil {
		return JobsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a job.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// jobAgentName - The name of the job agent.
// jobName - The name of the job to delete.
// options - JobsClientDeleteOptions contains the optional parameters for the JobsClient.Delete method.
func (client *JobsClient) Delete(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, options *JobsClientDeleteOptions) (JobsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, jobName, options)
	if err != nil {
		return JobsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return JobsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return JobsClientDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *JobsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, options *JobsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a job.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// jobAgentName - The name of the job agent.
// jobName - The name of the job to get.
// options - JobsClientGetOptions contains the optional parameters for the JobsClient.Get method.
func (client *JobsClient) Get(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, options *JobsClientGetOptions) (JobsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, jobName, options)
	if err != nil {
		return JobsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *JobsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, options *JobsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}"
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
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *JobsClient) getHandleResponse(resp *http.Response) (JobsClientGetResponse, error) {
	result := JobsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Job); err != nil {
		return JobsClientGetResponse{}, err
	}
	return result, nil
}

// ListByAgent - Gets a list of jobs.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// jobAgentName - The name of the job agent.
// options - JobsClientListByAgentOptions contains the optional parameters for the JobsClient.ListByAgent method.
func (client *JobsClient) ListByAgent(resourceGroupName string, serverName string, jobAgentName string, options *JobsClientListByAgentOptions) *JobsClientListByAgentPager {
	return &JobsClientListByAgentPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByAgentCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, options)
		},
		advancer: func(ctx context.Context, resp JobsClientListByAgentResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.JobListResult.NextLink)
		},
	}
}

// listByAgentCreateRequest creates the ListByAgent request.
func (client *JobsClient) listByAgentCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, options *JobsClientListByAgentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs"
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
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByAgentHandleResponse handles the ListByAgent response.
func (client *JobsClient) listByAgentHandleResponse(resp *http.Response) (JobsClientListByAgentResponse, error) {
	result := JobsClientListByAgentResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobListResult); err != nil {
		return JobsClientListByAgentResponse{}, err
	}
	return result, nil
}
