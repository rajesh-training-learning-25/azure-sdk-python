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

// DataMaskingPoliciesClient contains the methods for the DataMaskingPolicies group.
// Don't use this type directly, use NewDataMaskingPoliciesClient() instead.
type DataMaskingPoliciesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDataMaskingPoliciesClient creates a new instance of DataMaskingPoliciesClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDataMaskingPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DataMaskingPoliciesClient, error) {
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
	client := &DataMaskingPoliciesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a database data masking policy
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2014-04-01
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// databaseName - The name of the database.
// parameters - Parameters for creating or updating a data masking policy.
// options - DataMaskingPoliciesClientCreateOrUpdateOptions contains the optional parameters for the DataMaskingPoliciesClient.CreateOrUpdate
// method.
func (client *DataMaskingPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters DataMaskingPolicy, options *DataMaskingPoliciesClientCreateOrUpdateOptions) (DataMaskingPoliciesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, databaseName, parameters, options)
	if err != nil {
		return DataMaskingPoliciesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataMaskingPoliciesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataMaskingPoliciesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DataMaskingPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters DataMaskingPolicy, options *DataMaskingPoliciesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/dataMaskingPolicies/{dataMaskingPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	urlPath = strings.ReplaceAll(urlPath, "{dataMaskingPolicyName}", url.PathEscape("Default"))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2014-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DataMaskingPoliciesClient) createOrUpdateHandleResponse(resp *http.Response) (DataMaskingPoliciesClientCreateOrUpdateResponse, error) {
	result := DataMaskingPoliciesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataMaskingPolicy); err != nil {
		return DataMaskingPoliciesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Get - Gets a database data masking policy.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2014-04-01
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// databaseName - The name of the database.
// options - DataMaskingPoliciesClientGetOptions contains the optional parameters for the DataMaskingPoliciesClient.Get method.
func (client *DataMaskingPoliciesClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, options *DataMaskingPoliciesClientGetOptions) (DataMaskingPoliciesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, databaseName, options)
	if err != nil {
		return DataMaskingPoliciesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataMaskingPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataMaskingPoliciesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DataMaskingPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, options *DataMaskingPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/dataMaskingPolicies/{dataMaskingPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	urlPath = strings.ReplaceAll(urlPath, "{dataMaskingPolicyName}", url.PathEscape("Default"))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2014-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DataMaskingPoliciesClient) getHandleResponse(resp *http.Response) (DataMaskingPoliciesClientGetResponse, error) {
	result := DataMaskingPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataMaskingPolicy); err != nil {
		return DataMaskingPoliciesClientGetResponse{}, err
	}
	return result, nil
}