//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armautomation

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

// SoftwareUpdateConfigurationRunsClient contains the methods for the SoftwareUpdateConfigurationRuns group.
// Don't use this type directly, use NewSoftwareUpdateConfigurationRunsClient() instead.
type SoftwareUpdateConfigurationRunsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSoftwareUpdateConfigurationRunsClient creates a new instance of SoftwareUpdateConfigurationRunsClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSoftwareUpdateConfigurationRunsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SoftwareUpdateConfigurationRunsClient, error) {
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
	client := &SoftwareUpdateConfigurationRunsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// GetByID - Get a single software update configuration Run by Id.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// softwareUpdateConfigurationRunID - The Id of the software update configuration run.
// options - SoftwareUpdateConfigurationRunsClientGetByIDOptions contains the optional parameters for the SoftwareUpdateConfigurationRunsClient.GetByID
// method.
func (client *SoftwareUpdateConfigurationRunsClient) GetByID(ctx context.Context, resourceGroupName string, automationAccountName string, softwareUpdateConfigurationRunID string, options *SoftwareUpdateConfigurationRunsClientGetByIDOptions) (SoftwareUpdateConfigurationRunsClientGetByIDResponse, error) {
	req, err := client.getByIDCreateRequest(ctx, resourceGroupName, automationAccountName, softwareUpdateConfigurationRunID, options)
	if err != nil {
		return SoftwareUpdateConfigurationRunsClientGetByIDResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SoftwareUpdateConfigurationRunsClientGetByIDResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SoftwareUpdateConfigurationRunsClientGetByIDResponse{}, runtime.NewResponseError(resp)
	}
	return client.getByIDHandleResponse(resp)
}

// getByIDCreateRequest creates the GetByID request.
func (client *SoftwareUpdateConfigurationRunsClient) getByIDCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, softwareUpdateConfigurationRunID string, options *SoftwareUpdateConfigurationRunsClientGetByIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/softwareUpdateConfigurationRuns/{softwareUpdateConfigurationRunId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	urlPath = strings.ReplaceAll(urlPath, "{softwareUpdateConfigurationRunId}", url.PathEscape(softwareUpdateConfigurationRunID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.ClientRequestID != nil {
		req.Raw().Header["clientRequestId"] = []string{*options.ClientRequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByIDHandleResponse handles the GetByID response.
func (client *SoftwareUpdateConfigurationRunsClient) getByIDHandleResponse(resp *http.Response) (SoftwareUpdateConfigurationRunsClientGetByIDResponse, error) {
	result := SoftwareUpdateConfigurationRunsClientGetByIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SoftwareUpdateConfigurationRun); err != nil {
		return SoftwareUpdateConfigurationRunsClientGetByIDResponse{}, err
	}
	return result, nil
}

// List - Return list of software update configuration runs
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// options - SoftwareUpdateConfigurationRunsClientListOptions contains the optional parameters for the SoftwareUpdateConfigurationRunsClient.List
// method.
func (client *SoftwareUpdateConfigurationRunsClient) List(ctx context.Context, resourceGroupName string, automationAccountName string, options *SoftwareUpdateConfigurationRunsClientListOptions) (SoftwareUpdateConfigurationRunsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, automationAccountName, options)
	if err != nil {
		return SoftwareUpdateConfigurationRunsClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SoftwareUpdateConfigurationRunsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SoftwareUpdateConfigurationRunsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *SoftwareUpdateConfigurationRunsClient) listCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *SoftwareUpdateConfigurationRunsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/softwareUpdateConfigurationRuns"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", *options.Top)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.ClientRequestID != nil {
		req.Raw().Header["clientRequestId"] = []string{*options.ClientRequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SoftwareUpdateConfigurationRunsClient) listHandleResponse(resp *http.Response) (SoftwareUpdateConfigurationRunsClientListResponse, error) {
	result := SoftwareUpdateConfigurationRunsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SoftwareUpdateConfigurationRunListResult); err != nil {
		return SoftwareUpdateConfigurationRunsClientListResponse{}, err
	}
	return result, nil
}