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
	"strings"
)

// JobTargetGroupsClient contains the methods for the JobTargetGroups group.
// Don't use this type directly, use NewJobTargetGroupsClient() instead.
type JobTargetGroupsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewJobTargetGroupsClient creates a new instance of JobTargetGroupsClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewJobTargetGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*JobTargetGroupsClient, error) {
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
	client := &JobTargetGroupsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a target group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// jobAgentName - The name of the job agent.
// targetGroupName - The name of the target group.
// parameters - The requested state of the target group.
// options - JobTargetGroupsClientCreateOrUpdateOptions contains the optional parameters for the JobTargetGroupsClient.CreateOrUpdate
// method.
func (client *JobTargetGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, targetGroupName string, parameters JobTargetGroup, options *JobTargetGroupsClientCreateOrUpdateOptions) (JobTargetGroupsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, targetGroupName, parameters, options)
	if err != nil {
		return JobTargetGroupsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobTargetGroupsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return JobTargetGroupsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *JobTargetGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, targetGroupName string, parameters JobTargetGroup, options *JobTargetGroupsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/targetGroups/{targetGroupName}"
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
	if targetGroupName == "" {
		return nil, errors.New("parameter targetGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{targetGroupName}", url.PathEscape(targetGroupName))
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
func (client *JobTargetGroupsClient) createOrUpdateHandleResponse(resp *http.Response) (JobTargetGroupsClientCreateOrUpdateResponse, error) {
	result := JobTargetGroupsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobTargetGroup); err != nil {
		return JobTargetGroupsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a target group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// jobAgentName - The name of the job agent.
// targetGroupName - The name of the target group.
// options - JobTargetGroupsClientDeleteOptions contains the optional parameters for the JobTargetGroupsClient.Delete method.
func (client *JobTargetGroupsClient) Delete(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, targetGroupName string, options *JobTargetGroupsClientDeleteOptions) (JobTargetGroupsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, targetGroupName, options)
	if err != nil {
		return JobTargetGroupsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobTargetGroupsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return JobTargetGroupsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return JobTargetGroupsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *JobTargetGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, targetGroupName string, options *JobTargetGroupsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/targetGroups/{targetGroupName}"
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
	if targetGroupName == "" {
		return nil, errors.New("parameter targetGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{targetGroupName}", url.PathEscape(targetGroupName))
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

// Get - Gets a target group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// jobAgentName - The name of the job agent.
// targetGroupName - The name of the target group.
// options - JobTargetGroupsClientGetOptions contains the optional parameters for the JobTargetGroupsClient.Get method.
func (client *JobTargetGroupsClient) Get(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, targetGroupName string, options *JobTargetGroupsClientGetOptions) (JobTargetGroupsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, targetGroupName, options)
	if err != nil {
		return JobTargetGroupsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobTargetGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobTargetGroupsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *JobTargetGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, targetGroupName string, options *JobTargetGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/targetGroups/{targetGroupName}"
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
	if targetGroupName == "" {
		return nil, errors.New("parameter targetGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{targetGroupName}", url.PathEscape(targetGroupName))
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
func (client *JobTargetGroupsClient) getHandleResponse(resp *http.Response) (JobTargetGroupsClientGetResponse, error) {
	result := JobTargetGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobTargetGroup); err != nil {
		return JobTargetGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByAgentPager - Gets all target groups in an agent.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// jobAgentName - The name of the job agent.
// options - JobTargetGroupsClientListByAgentOptions contains the optional parameters for the JobTargetGroupsClient.ListByAgent
// method.
func (client *JobTargetGroupsClient) NewListByAgentPager(resourceGroupName string, serverName string, jobAgentName string, options *JobTargetGroupsClientListByAgentOptions) *runtime.Pager[JobTargetGroupsClientListByAgentResponse] {
	return runtime.NewPager(runtime.PagingHandler[JobTargetGroupsClientListByAgentResponse]{
		More: func(page JobTargetGroupsClientListByAgentResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *JobTargetGroupsClientListByAgentResponse) (JobTargetGroupsClientListByAgentResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByAgentCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return JobTargetGroupsClientListByAgentResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return JobTargetGroupsClientListByAgentResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return JobTargetGroupsClientListByAgentResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByAgentHandleResponse(resp)
		},
	})
}

// listByAgentCreateRequest creates the ListByAgent request.
func (client *JobTargetGroupsClient) listByAgentCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, options *JobTargetGroupsClientListByAgentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/targetGroups"
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
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByAgentHandleResponse handles the ListByAgent response.
func (client *JobTargetGroupsClient) listByAgentHandleResponse(resp *http.Response) (JobTargetGroupsClientListByAgentResponse, error) {
	result := JobTargetGroupsClientListByAgentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobTargetGroupListResult); err != nil {
		return JobTargetGroupsClientListByAgentResponse{}, err
	}
	return result, nil
}