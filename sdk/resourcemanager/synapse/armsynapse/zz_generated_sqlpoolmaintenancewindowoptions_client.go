//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

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

// SQLPoolMaintenanceWindowOptionsClient contains the methods for the SQLPoolMaintenanceWindowOptions group.
// Don't use this type directly, use NewSQLPoolMaintenanceWindowOptionsClient() instead.
type SQLPoolMaintenanceWindowOptionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSQLPoolMaintenanceWindowOptionsClient creates a new instance of SQLPoolMaintenanceWindowOptionsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSQLPoolMaintenanceWindowOptionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SQLPoolMaintenanceWindowOptionsClient, error) {
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
	client := &SQLPoolMaintenanceWindowOptionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Get list of SQL pool's available maintenance windows.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// sqlPoolName - SQL pool name
// maintenanceWindowOptionsName - Maintenance window options name.
// options - SQLPoolMaintenanceWindowOptionsClientGetOptions contains the optional parameters for the SQLPoolMaintenanceWindowOptionsClient.Get
// method.
func (client *SQLPoolMaintenanceWindowOptionsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, maintenanceWindowOptionsName string, options *SQLPoolMaintenanceWindowOptionsClientGetOptions) (SQLPoolMaintenanceWindowOptionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, maintenanceWindowOptionsName, options)
	if err != nil {
		return SQLPoolMaintenanceWindowOptionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLPoolMaintenanceWindowOptionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SQLPoolMaintenanceWindowOptionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SQLPoolMaintenanceWindowOptionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, maintenanceWindowOptionsName string, options *SQLPoolMaintenanceWindowOptionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/maintenanceWindowOptions/current"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	reqQP.Set("maintenanceWindowOptionsName", maintenanceWindowOptionsName)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SQLPoolMaintenanceWindowOptionsClient) getHandleResponse(resp *http.Response) (SQLPoolMaintenanceWindowOptionsClientGetResponse, error) {
	result := SQLPoolMaintenanceWindowOptionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MaintenanceWindowOptions); err != nil {
		return SQLPoolMaintenanceWindowOptionsClientGetResponse{}, err
	}
	return result, nil
}