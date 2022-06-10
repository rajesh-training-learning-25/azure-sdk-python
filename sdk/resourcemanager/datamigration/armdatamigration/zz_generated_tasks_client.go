//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatamigration

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

// TasksClient contains the methods for the Tasks group.
// Don't use this type directly, use NewTasksClient() instead.
type TasksClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewTasksClient creates a new instance of TasksClient with the specified values.
// subscriptionID - Subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewTasksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TasksClient, error) {
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
	client := &TasksClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Cancel - The tasks resource is a nested, proxy-only resource representing work performed by a DMS instance. This method
// cancels a task if it's currently queued or running.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-30-preview
// groupName - Name of the resource group
// serviceName - Name of the service
// projectName - Name of the project
// taskName - Name of the Task
// options - TasksClientCancelOptions contains the optional parameters for the TasksClient.Cancel method.
func (client *TasksClient) Cancel(ctx context.Context, groupName string, serviceName string, projectName string, taskName string, options *TasksClientCancelOptions) (TasksClientCancelResponse, error) {
	req, err := client.cancelCreateRequest(ctx, groupName, serviceName, projectName, taskName, options)
	if err != nil {
		return TasksClientCancelResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TasksClientCancelResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TasksClientCancelResponse{}, runtime.NewResponseError(resp)
	}
	return client.cancelHandleResponse(resp)
}

// cancelCreateRequest creates the Cancel request.
func (client *TasksClient) cancelCreateRequest(ctx context.Context, groupName string, serviceName string, projectName string, taskName string, options *TasksClientCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/tasks/{taskName}/cancel"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if taskName == "" {
		return nil, errors.New("parameter taskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{taskName}", url.PathEscape(taskName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// cancelHandleResponse handles the Cancel response.
func (client *TasksClient) cancelHandleResponse(resp *http.Response) (TasksClientCancelResponse, error) {
	result := TasksClientCancelResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProjectTask); err != nil {
		return TasksClientCancelResponse{}, err
	}
	return result, nil
}

// Command - The tasks resource is a nested, proxy-only resource representing work performed by a DMS instance. This method
// executes a command on a running task.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-30-preview
// groupName - Name of the resource group
// serviceName - Name of the service
// projectName - Name of the project
// taskName - Name of the Task
// parameters - Command to execute
// options - TasksClientCommandOptions contains the optional parameters for the TasksClient.Command method.
func (client *TasksClient) Command(ctx context.Context, groupName string, serviceName string, projectName string, taskName string, parameters CommandPropertiesClassification, options *TasksClientCommandOptions) (TasksClientCommandResponse, error) {
	req, err := client.commandCreateRequest(ctx, groupName, serviceName, projectName, taskName, parameters, options)
	if err != nil {
		return TasksClientCommandResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TasksClientCommandResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TasksClientCommandResponse{}, runtime.NewResponseError(resp)
	}
	return client.commandHandleResponse(resp)
}

// commandCreateRequest creates the Command request.
func (client *TasksClient) commandCreateRequest(ctx context.Context, groupName string, serviceName string, projectName string, taskName string, parameters CommandPropertiesClassification, options *TasksClientCommandOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/tasks/{taskName}/command"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if taskName == "" {
		return nil, errors.New("parameter taskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{taskName}", url.PathEscape(taskName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// commandHandleResponse handles the Command response.
func (client *TasksClient) commandHandleResponse(resp *http.Response) (TasksClientCommandResponse, error) {
	result := TasksClientCommandResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return TasksClientCommandResponse{}, err
	}
	return result, nil
}

// CreateOrUpdate - The tasks resource is a nested, proxy-only resource representing work performed by a DMS instance. The
// PUT method creates a new task or updates an existing one, although since tasks have no mutable
// custom properties, there is little reason to update an existing one.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-30-preview
// groupName - Name of the resource group
// serviceName - Name of the service
// projectName - Name of the project
// taskName - Name of the Task
// parameters - Information about the task
// options - TasksClientCreateOrUpdateOptions contains the optional parameters for the TasksClient.CreateOrUpdate method.
func (client *TasksClient) CreateOrUpdate(ctx context.Context, groupName string, serviceName string, projectName string, taskName string, parameters ProjectTask, options *TasksClientCreateOrUpdateOptions) (TasksClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, groupName, serviceName, projectName, taskName, parameters, options)
	if err != nil {
		return TasksClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TasksClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return TasksClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *TasksClient) createOrUpdateCreateRequest(ctx context.Context, groupName string, serviceName string, projectName string, taskName string, parameters ProjectTask, options *TasksClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/tasks/{taskName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if taskName == "" {
		return nil, errors.New("parameter taskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{taskName}", url.PathEscape(taskName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *TasksClient) createOrUpdateHandleResponse(resp *http.Response) (TasksClientCreateOrUpdateResponse, error) {
	result := TasksClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProjectTask); err != nil {
		return TasksClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - The tasks resource is a nested, proxy-only resource representing work performed by a DMS instance. The DELETE
// method deletes a task, canceling it first if it's running.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-30-preview
// groupName - Name of the resource group
// serviceName - Name of the service
// projectName - Name of the project
// taskName - Name of the Task
// options - TasksClientDeleteOptions contains the optional parameters for the TasksClient.Delete method.
func (client *TasksClient) Delete(ctx context.Context, groupName string, serviceName string, projectName string, taskName string, options *TasksClientDeleteOptions) (TasksClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, groupName, serviceName, projectName, taskName, options)
	if err != nil {
		return TasksClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TasksClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return TasksClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return TasksClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *TasksClient) deleteCreateRequest(ctx context.Context, groupName string, serviceName string, projectName string, taskName string, options *TasksClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/tasks/{taskName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if taskName == "" {
		return nil, errors.New("parameter taskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{taskName}", url.PathEscape(taskName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-30-preview")
	if options != nil && options.DeleteRunningTasks != nil {
		reqQP.Set("deleteRunningTasks", strconv.FormatBool(*options.DeleteRunningTasks))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - The tasks resource is a nested, proxy-only resource representing work performed by a DMS instance. The GET method
// retrieves information about a task.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-30-preview
// groupName - Name of the resource group
// serviceName - Name of the service
// projectName - Name of the project
// taskName - Name of the Task
// options - TasksClientGetOptions contains the optional parameters for the TasksClient.Get method.
func (client *TasksClient) Get(ctx context.Context, groupName string, serviceName string, projectName string, taskName string, options *TasksClientGetOptions) (TasksClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, groupName, serviceName, projectName, taskName, options)
	if err != nil {
		return TasksClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TasksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TasksClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TasksClient) getCreateRequest(ctx context.Context, groupName string, serviceName string, projectName string, taskName string, options *TasksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/tasks/{taskName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if taskName == "" {
		return nil, errors.New("parameter taskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{taskName}", url.PathEscape(taskName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-30-preview")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TasksClient) getHandleResponse(resp *http.Response) (TasksClientGetResponse, error) {
	result := TasksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProjectTask); err != nil {
		return TasksClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - The services resource is the top-level resource that represents the Database Migration Service. This method
// returns a list of tasks owned by a service resource. Some tasks may have a status of
// Unknown, which indicates that an error occurred while querying the status of that task.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-30-preview
// groupName - Name of the resource group
// serviceName - Name of the service
// projectName - Name of the project
// options - TasksClientListOptions contains the optional parameters for the TasksClient.List method.
func (client *TasksClient) NewListPager(groupName string, serviceName string, projectName string, options *TasksClientListOptions) *runtime.Pager[TasksClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[TasksClientListResponse]{
		More: func(page TasksClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TasksClientListResponse) (TasksClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, groupName, serviceName, projectName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return TasksClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return TasksClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TasksClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *TasksClient) listCreateRequest(ctx context.Context, groupName string, serviceName string, projectName string, options *TasksClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/tasks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-30-preview")
	if options != nil && options.TaskType != nil {
		reqQP.Set("taskType", *options.TaskType)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TasksClient) listHandleResponse(resp *http.Response) (TasksClientListResponse, error) {
	result := TasksClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TaskList); err != nil {
		return TasksClientListResponse{}, err
	}
	return result, nil
}

// Update - The tasks resource is a nested, proxy-only resource representing work performed by a DMS instance. The PATCH method
// updates an existing task, but since tasks have no mutable custom properties, there
// is little reason to do so.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-30-preview
// groupName - Name of the resource group
// serviceName - Name of the service
// projectName - Name of the project
// taskName - Name of the Task
// parameters - Information about the task
// options - TasksClientUpdateOptions contains the optional parameters for the TasksClient.Update method.
func (client *TasksClient) Update(ctx context.Context, groupName string, serviceName string, projectName string, taskName string, parameters ProjectTask, options *TasksClientUpdateOptions) (TasksClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, groupName, serviceName, projectName, taskName, parameters, options)
	if err != nil {
		return TasksClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TasksClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TasksClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *TasksClient) updateCreateRequest(ctx context.Context, groupName string, serviceName string, projectName string, taskName string, parameters ProjectTask, options *TasksClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/tasks/{taskName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if taskName == "" {
		return nil, errors.New("parameter taskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{taskName}", url.PathEscape(taskName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *TasksClient) updateHandleResponse(resp *http.Response) (TasksClientUpdateResponse, error) {
	result := TasksClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProjectTask); err != nil {
		return TasksClientUpdateResponse{}, err
	}
	return result, nil
}
