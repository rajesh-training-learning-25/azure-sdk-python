//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevtestlabs

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
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
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewArtifactsClient creates a new instance of ArtifactsClient with the specified values.
func NewArtifactsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ArtifactsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ArtifactsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// GenerateArmTemplate - Generates an ARM template for the given artifact, uploads the required files to a storage account, and validates the generated
// artifact.
// If the operation fails it returns the *CloudError error type.
func (client *ArtifactsClient) GenerateArmTemplate(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, name string, generateArmTemplateRequest GenerateArmTemplateRequest, options *ArtifactsGenerateArmTemplateOptions) (ArtifactsGenerateArmTemplateResponse, error) {
	req, err := client.generateArmTemplateCreateRequest(ctx, resourceGroupName, labName, artifactSourceName, name, generateArmTemplateRequest, options)
	if err != nil {
		return ArtifactsGenerateArmTemplateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ArtifactsGenerateArmTemplateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ArtifactsGenerateArmTemplateResponse{}, client.generateArmTemplateHandleError(resp)
	}
	return client.generateArmTemplateHandleResponse(resp)
}

// generateArmTemplateCreateRequest creates the GenerateArmTemplate request.
func (client *ArtifactsClient) generateArmTemplateCreateRequest(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, name string, generateArmTemplateRequest GenerateArmTemplateRequest, options *ArtifactsGenerateArmTemplateOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, generateArmTemplateRequest)
}

// generateArmTemplateHandleResponse handles the GenerateArmTemplate response.
func (client *ArtifactsClient) generateArmTemplateHandleResponse(resp *http.Response) (ArtifactsGenerateArmTemplateResponse, error) {
	result := ArtifactsGenerateArmTemplateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArmTemplateInfo); err != nil {
		return ArtifactsGenerateArmTemplateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// generateArmTemplateHandleError handles the GenerateArmTemplate error response.
func (client *ArtifactsClient) generateArmTemplateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Get artifact.
// If the operation fails it returns the *CloudError error type.
func (client *ArtifactsClient) Get(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, name string, options *ArtifactsGetOptions) (ArtifactsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, labName, artifactSourceName, name, options)
	if err != nil {
		return ArtifactsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ArtifactsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ArtifactsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ArtifactsClient) getCreateRequest(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, name string, options *ArtifactsGetOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ArtifactsClient) getHandleResponse(resp *http.Response) (ArtifactsGetResponse, error) {
	result := ArtifactsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Artifact); err != nil {
		return ArtifactsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ArtifactsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - List artifacts in a given artifact source.
// If the operation fails it returns the *CloudError error type.
func (client *ArtifactsClient) List(resourceGroupName string, labName string, artifactSourceName string, options *ArtifactsListOptions) *ArtifactsListPager {
	return &ArtifactsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, labName, artifactSourceName, options)
		},
		advancer: func(ctx context.Context, resp ArtifactsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ArtifactList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ArtifactsClient) listCreateRequest(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, options *ArtifactsListOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
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
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ArtifactsClient) listHandleResponse(resp *http.Response) (ArtifactsListResponse, error) {
	result := ArtifactsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArtifactList); err != nil {
		return ArtifactsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ArtifactsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
