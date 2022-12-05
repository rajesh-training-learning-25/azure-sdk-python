//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurityinsights

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

// ProductSettingsClient contains the methods for the ProductSettings group.
// Don't use this type directly, use NewProductSettingsClient() instead.
type ProductSettingsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewProductSettingsClient creates a new instance of ProductSettingsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewProductSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ProductSettingsClient, error) {
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
	client := &ProductSettingsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Delete - Delete setting of the product.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// settingsName - The setting name. Supports - Anomalies, EyesOn, EntityAnalytics, Ueba
// options - ProductSettingsClientDeleteOptions contains the optional parameters for the ProductSettingsClient.Delete method.
func (client *ProductSettingsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, settingsName string, options *ProductSettingsClientDeleteOptions) (ProductSettingsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, settingsName, options)
	if err != nil {
		return ProductSettingsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProductSettingsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ProductSettingsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ProductSettingsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ProductSettingsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, settingsName string, options *ProductSettingsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/settings/{settingsName}"
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
	if settingsName == "" {
		return nil, errors.New("parameter settingsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{settingsName}", url.PathEscape(settingsName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a setting.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// settingsName - The setting name. Supports - Anomalies, EyesOn, EntityAnalytics, Ueba
// options - ProductSettingsClientGetOptions contains the optional parameters for the ProductSettingsClient.Get method.
func (client *ProductSettingsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, settingsName string, options *ProductSettingsClientGetOptions) (ProductSettingsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, settingsName, options)
	if err != nil {
		return ProductSettingsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProductSettingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProductSettingsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ProductSettingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, settingsName string, options *ProductSettingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/settings/{settingsName}"
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
	if settingsName == "" {
		return nil, errors.New("parameter settingsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{settingsName}", url.PathEscape(settingsName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProductSettingsClient) getHandleResponse(resp *http.Response) (ProductSettingsClientGetResponse, error) {
	result := ProductSettingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return ProductSettingsClientGetResponse{}, err
	}
	return result, nil
}

// List - List of all the settings
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// options - ProductSettingsClientListOptions contains the optional parameters for the ProductSettingsClient.List method.
func (client *ProductSettingsClient) List(ctx context.Context, resourceGroupName string, workspaceName string, options *ProductSettingsClientListOptions) (ProductSettingsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
	if err != nil {
		return ProductSettingsClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProductSettingsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProductSettingsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ProductSettingsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *ProductSettingsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/settings"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ProductSettingsClient) listHandleResponse(resp *http.Response) (ProductSettingsClientListResponse, error) {
	result := ProductSettingsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SettingList); err != nil {
		return ProductSettingsClientListResponse{}, err
	}
	return result, nil
}

// Update - Updates setting.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// settingsName - The setting name. Supports - Anomalies, EyesOn, EntityAnalytics, Ueba
// settings - The setting
// options - ProductSettingsClientUpdateOptions contains the optional parameters for the ProductSettingsClient.Update method.
func (client *ProductSettingsClient) Update(ctx context.Context, resourceGroupName string, workspaceName string, settingsName string, settings SettingsClassification, options *ProductSettingsClientUpdateOptions) (ProductSettingsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, workspaceName, settingsName, settings, options)
	if err != nil {
		return ProductSettingsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProductSettingsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProductSettingsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ProductSettingsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, settingsName string, settings SettingsClassification, options *ProductSettingsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/settings/{settingsName}"
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
	if settingsName == "" {
		return nil, errors.New("parameter settingsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{settingsName}", url.PathEscape(settingsName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, settings)
}

// updateHandleResponse handles the Update response.
func (client *ProductSettingsClient) updateHandleResponse(resp *http.Response) (ProductSettingsClientUpdateResponse, error) {
	result := ProductSettingsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return ProductSettingsClientUpdateResponse{}, err
	}
	return result, nil
}