//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsqlvirtualmachine

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

// SQLVirtualMachinesClient contains the methods for the SQLVirtualMachines group.
// Don't use this type directly, use NewSQLVirtualMachinesClient() instead.
type SQLVirtualMachinesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSQLVirtualMachinesClient creates a new instance of SQLVirtualMachinesClient with the specified values.
// subscriptionID - Subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSQLVirtualMachinesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SQLVirtualMachinesClient, error) {
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
	client := &SQLVirtualMachinesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a SQL virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01-preview
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// sqlVirtualMachineName - Name of the SQL virtual machine.
// parameters - The SQL virtual machine.
// options - SQLVirtualMachinesClientBeginCreateOrUpdateOptions contains the optional parameters for the SQLVirtualMachinesClient.BeginCreateOrUpdate
// method.
func (client *SQLVirtualMachinesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, parameters SQLVirtualMachine, options *SQLVirtualMachinesClientBeginCreateOrUpdateOptions) (*runtime.Poller[SQLVirtualMachinesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, sqlVirtualMachineName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[SQLVirtualMachinesClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[SQLVirtualMachinesClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a SQL virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01-preview
func (client *SQLVirtualMachinesClient) createOrUpdate(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, parameters SQLVirtualMachine, options *SQLVirtualMachinesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, sqlVirtualMachineName, parameters, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SQLVirtualMachinesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, parameters SQLVirtualMachine, options *SQLVirtualMachinesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/{sqlVirtualMachineName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlVirtualMachineName == "" {
		return nil, errors.New("parameter sqlVirtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlVirtualMachineName}", url.PathEscape(sqlVirtualMachineName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a SQL virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01-preview
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// sqlVirtualMachineName - Name of the SQL virtual machine.
// options - SQLVirtualMachinesClientBeginDeleteOptions contains the optional parameters for the SQLVirtualMachinesClient.BeginDelete
// method.
func (client *SQLVirtualMachinesClient) BeginDelete(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, options *SQLVirtualMachinesClientBeginDeleteOptions) (*runtime.Poller[SQLVirtualMachinesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, sqlVirtualMachineName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[SQLVirtualMachinesClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[SQLVirtualMachinesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a SQL virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01-preview
func (client *SQLVirtualMachinesClient) deleteOperation(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, options *SQLVirtualMachinesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, sqlVirtualMachineName, options)
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
func (client *SQLVirtualMachinesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, options *SQLVirtualMachinesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/{sqlVirtualMachineName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlVirtualMachineName == "" {
		return nil, errors.New("parameter sqlVirtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlVirtualMachineName}", url.PathEscape(sqlVirtualMachineName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a SQL virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01-preview
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// sqlVirtualMachineName - Name of the SQL virtual machine.
// options - SQLVirtualMachinesClientGetOptions contains the optional parameters for the SQLVirtualMachinesClient.Get method.
func (client *SQLVirtualMachinesClient) Get(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, options *SQLVirtualMachinesClientGetOptions) (SQLVirtualMachinesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, sqlVirtualMachineName, options)
	if err != nil {
		return SQLVirtualMachinesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLVirtualMachinesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SQLVirtualMachinesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SQLVirtualMachinesClient) getCreateRequest(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, options *SQLVirtualMachinesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/{sqlVirtualMachineName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlVirtualMachineName == "" {
		return nil, errors.New("parameter sqlVirtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlVirtualMachineName}", url.PathEscape(sqlVirtualMachineName))
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
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SQLVirtualMachinesClient) getHandleResponse(resp *http.Response) (SQLVirtualMachinesClientGetResponse, error) {
	result := SQLVirtualMachinesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLVirtualMachine); err != nil {
		return SQLVirtualMachinesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all SQL virtual machines in a subscription.
// Generated from API version 2022-07-01-preview
// options - SQLVirtualMachinesClientListOptions contains the optional parameters for the SQLVirtualMachinesClient.List method.
func (client *SQLVirtualMachinesClient) NewListPager(options *SQLVirtualMachinesClientListOptions) *runtime.Pager[SQLVirtualMachinesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SQLVirtualMachinesClientListResponse]{
		More: func(page SQLVirtualMachinesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SQLVirtualMachinesClientListResponse) (SQLVirtualMachinesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SQLVirtualMachinesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SQLVirtualMachinesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SQLVirtualMachinesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SQLVirtualMachinesClient) listCreateRequest(ctx context.Context, options *SQLVirtualMachinesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SQLVirtualMachinesClient) listHandleResponse(resp *http.Response) (SQLVirtualMachinesClientListResponse, error) {
	result := SQLVirtualMachinesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return SQLVirtualMachinesClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets all SQL virtual machines in a resource group.
// Generated from API version 2022-07-01-preview
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// options - SQLVirtualMachinesClientListByResourceGroupOptions contains the optional parameters for the SQLVirtualMachinesClient.ListByResourceGroup
// method.
func (client *SQLVirtualMachinesClient) NewListByResourceGroupPager(resourceGroupName string, options *SQLVirtualMachinesClientListByResourceGroupOptions) *runtime.Pager[SQLVirtualMachinesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[SQLVirtualMachinesClientListByResourceGroupResponse]{
		More: func(page SQLVirtualMachinesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SQLVirtualMachinesClientListByResourceGroupResponse) (SQLVirtualMachinesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SQLVirtualMachinesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SQLVirtualMachinesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SQLVirtualMachinesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SQLVirtualMachinesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *SQLVirtualMachinesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *SQLVirtualMachinesClient) listByResourceGroupHandleResponse(resp *http.Response) (SQLVirtualMachinesClientListByResourceGroupResponse, error) {
	result := SQLVirtualMachinesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return SQLVirtualMachinesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySQLVMGroupPager - Gets the list of sql virtual machines in a SQL virtual machine group.
// Generated from API version 2022-07-01-preview
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// sqlVirtualMachineGroupName - Name of the SQL virtual machine group.
// options - SQLVirtualMachinesClientListBySQLVMGroupOptions contains the optional parameters for the SQLVirtualMachinesClient.ListBySQLVMGroup
// method.
func (client *SQLVirtualMachinesClient) NewListBySQLVMGroupPager(resourceGroupName string, sqlVirtualMachineGroupName string, options *SQLVirtualMachinesClientListBySQLVMGroupOptions) *runtime.Pager[SQLVirtualMachinesClientListBySQLVMGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[SQLVirtualMachinesClientListBySQLVMGroupResponse]{
		More: func(page SQLVirtualMachinesClientListBySQLVMGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SQLVirtualMachinesClientListBySQLVMGroupResponse) (SQLVirtualMachinesClientListBySQLVMGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySQLVMGroupCreateRequest(ctx, resourceGroupName, sqlVirtualMachineGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SQLVirtualMachinesClientListBySQLVMGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SQLVirtualMachinesClientListBySQLVMGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SQLVirtualMachinesClientListBySQLVMGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySQLVMGroupHandleResponse(resp)
		},
	})
}

// listBySQLVMGroupCreateRequest creates the ListBySQLVMGroup request.
func (client *SQLVirtualMachinesClient) listBySQLVMGroupCreateRequest(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, options *SQLVirtualMachinesClientListBySQLVMGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/{sqlVirtualMachineGroupName}/sqlVirtualMachines"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlVirtualMachineGroupName == "" {
		return nil, errors.New("parameter sqlVirtualMachineGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlVirtualMachineGroupName}", url.PathEscape(sqlVirtualMachineGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySQLVMGroupHandleResponse handles the ListBySQLVMGroup response.
func (client *SQLVirtualMachinesClient) listBySQLVMGroupHandleResponse(resp *http.Response) (SQLVirtualMachinesClientListBySQLVMGroupResponse, error) {
	result := SQLVirtualMachinesClientListBySQLVMGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return SQLVirtualMachinesClientListBySQLVMGroupResponse{}, err
	}
	return result, nil
}

// BeginRedeploy - Uninstalls and reinstalls the SQL Iaas Extension.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01-preview
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// sqlVirtualMachineName - Name of the SQL virtual machine.
// options - SQLVirtualMachinesClientBeginRedeployOptions contains the optional parameters for the SQLVirtualMachinesClient.BeginRedeploy
// method.
func (client *SQLVirtualMachinesClient) BeginRedeploy(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, options *SQLVirtualMachinesClientBeginRedeployOptions) (*runtime.Poller[SQLVirtualMachinesClientRedeployResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.redeploy(ctx, resourceGroupName, sqlVirtualMachineName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[SQLVirtualMachinesClientRedeployResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[SQLVirtualMachinesClientRedeployResponse](options.ResumeToken, client.pl, nil)
	}
}

// Redeploy - Uninstalls and reinstalls the SQL Iaas Extension.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01-preview
func (client *SQLVirtualMachinesClient) redeploy(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, options *SQLVirtualMachinesClientBeginRedeployOptions) (*http.Response, error) {
	req, err := client.redeployCreateRequest(ctx, resourceGroupName, sqlVirtualMachineName, options)
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

// redeployCreateRequest creates the Redeploy request.
func (client *SQLVirtualMachinesClient) redeployCreateRequest(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, options *SQLVirtualMachinesClientBeginRedeployOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/{sqlVirtualMachineName}/redeploy"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlVirtualMachineName == "" {
		return nil, errors.New("parameter sqlVirtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlVirtualMachineName}", url.PathEscape(sqlVirtualMachineName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// BeginStartAssessment - Starts Assessment on SQL virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01-preview
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// sqlVirtualMachineName - Name of the SQL virtual machine.
// options - SQLVirtualMachinesClientBeginStartAssessmentOptions contains the optional parameters for the SQLVirtualMachinesClient.BeginStartAssessment
// method.
func (client *SQLVirtualMachinesClient) BeginStartAssessment(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, options *SQLVirtualMachinesClientBeginStartAssessmentOptions) (*runtime.Poller[SQLVirtualMachinesClientStartAssessmentResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.startAssessment(ctx, resourceGroupName, sqlVirtualMachineName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[SQLVirtualMachinesClientStartAssessmentResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[SQLVirtualMachinesClientStartAssessmentResponse](options.ResumeToken, client.pl, nil)
	}
}

// StartAssessment - Starts Assessment on SQL virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01-preview
func (client *SQLVirtualMachinesClient) startAssessment(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, options *SQLVirtualMachinesClientBeginStartAssessmentOptions) (*http.Response, error) {
	req, err := client.startAssessmentCreateRequest(ctx, resourceGroupName, sqlVirtualMachineName, options)
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

// startAssessmentCreateRequest creates the StartAssessment request.
func (client *SQLVirtualMachinesClient) startAssessmentCreateRequest(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, options *SQLVirtualMachinesClientBeginStartAssessmentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/{sqlVirtualMachineName}/startAssessment"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlVirtualMachineName == "" {
		return nil, errors.New("parameter sqlVirtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlVirtualMachineName}", url.PathEscape(sqlVirtualMachineName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// BeginUpdate - Updates a SQL virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01-preview
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// sqlVirtualMachineName - Name of the SQL virtual machine.
// parameters - The SQL virtual machine.
// options - SQLVirtualMachinesClientBeginUpdateOptions contains the optional parameters for the SQLVirtualMachinesClient.BeginUpdate
// method.
func (client *SQLVirtualMachinesClient) BeginUpdate(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, parameters Update, options *SQLVirtualMachinesClientBeginUpdateOptions) (*runtime.Poller[SQLVirtualMachinesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, sqlVirtualMachineName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[SQLVirtualMachinesClientUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[SQLVirtualMachinesClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Updates a SQL virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01-preview
func (client *SQLVirtualMachinesClient) update(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, parameters Update, options *SQLVirtualMachinesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, sqlVirtualMachineName, parameters, options)
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
func (client *SQLVirtualMachinesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, parameters Update, options *SQLVirtualMachinesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/{sqlVirtualMachineName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlVirtualMachineName == "" {
		return nil, errors.New("parameter sqlVirtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlVirtualMachineName}", url.PathEscape(sqlVirtualMachineName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}