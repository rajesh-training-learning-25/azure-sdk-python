//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevtestlabs

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

// ArtifactsClient contains the methods for the Artifacts group.
// Don't use this type directly, use NewArtifactsClient() instead.
type ArtifactsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewArtifactsClient creates a new instance of ArtifactsClient with the specified values.
// subscriptionID - The subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewArtifactsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ArtifactsClient, error) {
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
	client := &ArtifactsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// GenerateArmTemplate - Generates an ARM template for the given artifact, uploads the required files to a storage account,
// and validates the generated artifact.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-09-15
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// artifactSourceName - The name of the artifact source.
// name - The name of the artifact.
// generateArmTemplateRequest - Parameters for generating an ARM template for deploying artifacts.
// options - ArtifactsClientGenerateArmTemplateOptions contains the optional parameters for the ArtifactsClient.GenerateArmTemplate
// method.
func (client *ArtifactsClient) GenerateArmTemplate(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, name string, generateArmTemplateRequest GenerateArmTemplateRequest, options *ArtifactsClientGenerateArmTemplateOptions) (ArtifactsClientGenerateArmTemplateResponse, error) {
	req, err := client.generateArmTemplateCreateRequest(ctx, resourceGroupName, labName, artifactSourceName, name, generateArmTemplateRequest, options)
	if err != nil {
		return ArtifactsClientGenerateArmTemplateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ArtifactsClientGenerateArmTemplateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ArtifactsClientGenerateArmTemplateResponse{}, runtime.NewResponseError(resp)
	}
	return client.generateArmTemplateHandleResponse(resp)
}

// generateArmTemplateCreateRequest creates the GenerateArmTemplate request.
func (client *ArtifactsClient) generateArmTemplateCreateRequest(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, name string, generateArmTemplateRequest GenerateArmTemplateRequest, options *ArtifactsClientGenerateArmTemplateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}/artifacts/{name}/generateArmTemplate"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if artifactSourceName == "" {
		return nil, errors.New("parameter artifactSourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{artifactSourceName}", url.PathEscape(artifactSourceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, generateArmTemplateRequest)
}

// generateArmTemplateHandleResponse handles the GenerateArmTemplate response.
func (client *ArtifactsClient) generateArmTemplateHandleResponse(resp *http.Response) (ArtifactsClientGenerateArmTemplateResponse, error) {
	result := ArtifactsClientGenerateArmTemplateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArmTemplateInfo); err != nil {
		return ArtifactsClientGenerateArmTemplateResponse{}, err
	}
	return result, nil
}

// Get - Get artifact.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-09-15
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// artifactSourceName - The name of the artifact source.
// name - The name of the artifact.
// options - ArtifactsClientGetOptions contains the optional parameters for the ArtifactsClient.Get method.
func (client *ArtifactsClient) Get(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, name string, options *ArtifactsClientGetOptions) (ArtifactsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, labName, artifactSourceName, name, options)
	if err != nil {
		return ArtifactsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ArtifactsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ArtifactsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ArtifactsClient) getCreateRequest(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, name string, options *ArtifactsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}/artifacts/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if artifactSourceName == "" {
		return nil, errors.New("parameter artifactSourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{artifactSourceName}", url.PathEscape(artifactSourceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ArtifactsClient) getHandleResponse(resp *http.Response) (ArtifactsClientGetResponse, error) {
	result := ArtifactsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Artifact); err != nil {
		return ArtifactsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List artifacts in a given artifact source.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-09-15
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// artifactSourceName - The name of the artifact source.
// options - ArtifactsClientListOptions contains the optional parameters for the ArtifactsClient.List method.
func (client *ArtifactsClient) NewListPager(resourceGroupName string, labName string, artifactSourceName string, options *ArtifactsClientListOptions) *runtime.Pager[ArtifactsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ArtifactsClientListResponse]{
		More: func(page ArtifactsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ArtifactsClientListResponse) (ArtifactsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, labName, artifactSourceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ArtifactsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ArtifactsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ArtifactsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ArtifactsClient) listCreateRequest(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, options *ArtifactsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}/artifacts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if artifactSourceName == "" {
		return nil, errors.New("parameter artifactSourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{artifactSourceName}", url.PathEscape(artifactSourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ArtifactsClient) listHandleResponse(resp *http.Response) (ArtifactsClientListResponse, error) {
	result := ArtifactsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArtifactList); err != nil {
		return ArtifactsClientListResponse{}, err
	}
	return result, nil
}