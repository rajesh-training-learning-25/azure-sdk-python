//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrate

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

// PrivateLinkResourceClient contains the methods for the PrivateLinkResource group.
// Don't use this type directly, use NewPrivateLinkResourceClient() instead.
type PrivateLinkResourceClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPrivateLinkResourceClient creates a new instance of PrivateLinkResourceClient with the specified values.
// subscriptionID - Azure Subscription Id in which project was created.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPrivateLinkResourceClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateLinkResourceClient, error) {
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
	client := &PrivateLinkResourceClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Get information related to a specific private Link Resource in the project. Returns a json object of type 'privateLinkResources'
// as specified in the models section.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-10-01
// resourceGroupName - Name of the Azure Resource Group that project is part of.
// projectName - Name of the Azure Migrate project.
// privateLinkResourceName - Unique name of a private link resource within a project.
// options - PrivateLinkResourceClientGetOptions contains the optional parameters for the PrivateLinkResourceClient.Get method.
func (client *PrivateLinkResourceClient) Get(ctx context.Context, resourceGroupName string, projectName string, privateLinkResourceName string, options *PrivateLinkResourceClientGetOptions) (PrivateLinkResourceClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, projectName, privateLinkResourceName, options)
	if err != nil {
		return PrivateLinkResourceClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateLinkResourceClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateLinkResourceClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateLinkResourceClient) getCreateRequest(ctx context.Context, resourceGroupName string, projectName string, privateLinkResourceName string, options *PrivateLinkResourceClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentprojects/{projectName}/privateLinkResources/{privateLinkResourceName}"
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
	if privateLinkResourceName == "" {
		return nil, errors.New("parameter privateLinkResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateLinkResourceName}", url.PathEscape(privateLinkResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateLinkResourceClient) getHandleResponse(resp *http.Response) (PrivateLinkResourceClientGetResponse, error) {
	result := PrivateLinkResourceClientGetResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResource); err != nil {
		return PrivateLinkResourceClientGetResponse{}, err
	}
	return result, nil
}

// ListByProject - Get all private link resources created in the project. Returns a json array of objects of type 'privateLinkResources'
// as specified in the Models section.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-10-01
// resourceGroupName - Name of the Azure Resource Group that project is part of.
// projectName - Name of the Azure Migrate project.
// options - PrivateLinkResourceClientListByProjectOptions contains the optional parameters for the PrivateLinkResourceClient.ListByProject
// method.
func (client *PrivateLinkResourceClient) ListByProject(ctx context.Context, resourceGroupName string, projectName string, options *PrivateLinkResourceClientListByProjectOptions) (PrivateLinkResourceClientListByProjectResponse, error) {
	req, err := client.listByProjectCreateRequest(ctx, resourceGroupName, projectName, options)
	if err != nil {
		return PrivateLinkResourceClientListByProjectResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateLinkResourceClientListByProjectResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateLinkResourceClientListByProjectResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByProjectHandleResponse(resp)
}

// listByProjectCreateRequest creates the ListByProject request.
func (client *PrivateLinkResourceClient) listByProjectCreateRequest(ctx context.Context, resourceGroupName string, projectName string, options *PrivateLinkResourceClientListByProjectOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentprojects/{projectName}/privateLinkResources"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByProjectHandleResponse handles the ListByProject response.
func (client *PrivateLinkResourceClient) listByProjectHandleResponse(resp *http.Response) (PrivateLinkResourceClientListByProjectResponse, error) {
	result := PrivateLinkResourceClientListByProjectResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResourceCollection); err != nil {
		return PrivateLinkResourceClientListByProjectResponse{}, err
	}
	return result, nil
}