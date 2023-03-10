//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

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

// SparkConfigurationClient contains the methods for the SparkConfiguration group.
// Don't use this type directly, use NewSparkConfigurationClient() instead.
type SparkConfigurationClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSparkConfigurationClient creates a new instance of SparkConfigurationClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSparkConfigurationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SparkConfigurationClient, error) {
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
	client := &SparkConfigurationClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Get SparkConfiguration by name in a workspace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - sparkConfigurationName - SparkConfiguration name
//   - workspaceName - The name of the workspace.
//   - options - SparkConfigurationClientGetOptions contains the optional parameters for the SparkConfigurationClient.Get method.
func (client *SparkConfigurationClient) Get(ctx context.Context, resourceGroupName string, sparkConfigurationName string, workspaceName string, options *SparkConfigurationClientGetOptions) (SparkConfigurationClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, sparkConfigurationName, workspaceName, options)
	if err != nil {
		return SparkConfigurationClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SparkConfigurationClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SparkConfigurationClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SparkConfigurationClient) getCreateRequest(ctx context.Context, resourceGroupName string, sparkConfigurationName string, workspaceName string, options *SparkConfigurationClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sparkconfigurations/{sparkConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sparkConfigurationName == "" {
		return nil, errors.New("parameter sparkConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sparkConfigurationName}", url.PathEscape(sparkConfigurationName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SparkConfigurationClient) getHandleResponse(resp *http.Response) (SparkConfigurationClientGetResponse, error) {
	result := SparkConfigurationClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SparkConfigurationResource); err != nil {
		return SparkConfigurationClientGetResponse{}, err
	}
	return result, nil
}
