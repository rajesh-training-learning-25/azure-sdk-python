//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmaintenance

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ConfigurationAssignmentsForSubscriptionsClient contains the methods for the ConfigurationAssignmentsForSubscriptions group.
// Don't use this type directly, use NewConfigurationAssignmentsForSubscriptionsClient() instead.
type ConfigurationAssignmentsForSubscriptionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewConfigurationAssignmentsForSubscriptionsClient creates a new instance of ConfigurationAssignmentsForSubscriptionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewConfigurationAssignmentsForSubscriptionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConfigurationAssignmentsForSubscriptionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ConfigurationAssignmentsForSubscriptionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Register configuration for resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - configurationAssignmentName - Configuration assignment name
//   - configurationAssignment - The configurationAssignment
//   - options - ConfigurationAssignmentsForSubscriptionsClientCreateOrUpdateOptions contains the optional parameters for the
//     ConfigurationAssignmentsForSubscriptionsClient.CreateOrUpdate method.
func (client *ConfigurationAssignmentsForSubscriptionsClient) CreateOrUpdate(ctx context.Context, configurationAssignmentName string, configurationAssignment ConfigurationAssignment, options *ConfigurationAssignmentsForSubscriptionsClientCreateOrUpdateOptions) (ConfigurationAssignmentsForSubscriptionsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "ConfigurationAssignmentsForSubscriptionsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, configurationAssignmentName, configurationAssignment, options)
	if err != nil {
		return ConfigurationAssignmentsForSubscriptionsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConfigurationAssignmentsForSubscriptionsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ConfigurationAssignmentsForSubscriptionsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ConfigurationAssignmentsForSubscriptionsClient) createOrUpdateCreateRequest(ctx context.Context, configurationAssignmentName string, configurationAssignment ConfigurationAssignment, options *ConfigurationAssignmentsForSubscriptionsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Maintenance/configurationAssignments/{configurationAssignmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if configurationAssignmentName == "" {
		return nil, errors.New("parameter configurationAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationAssignmentName}", url.PathEscape(configurationAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, configurationAssignment); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ConfigurationAssignmentsForSubscriptionsClient) createOrUpdateHandleResponse(resp *http.Response) (ConfigurationAssignmentsForSubscriptionsClientCreateOrUpdateResponse, error) {
	result := ConfigurationAssignmentsForSubscriptionsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationAssignment); err != nil {
		return ConfigurationAssignmentsForSubscriptionsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Unregister configuration for resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - configurationAssignmentName - Unique configuration assignment name
//   - options - ConfigurationAssignmentsForSubscriptionsClientDeleteOptions contains the optional parameters for the ConfigurationAssignmentsForSubscriptionsClient.Delete
//     method.
func (client *ConfigurationAssignmentsForSubscriptionsClient) Delete(ctx context.Context, configurationAssignmentName string, options *ConfigurationAssignmentsForSubscriptionsClientDeleteOptions) (ConfigurationAssignmentsForSubscriptionsClientDeleteResponse, error) {
	var err error
	const operationName = "ConfigurationAssignmentsForSubscriptionsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, configurationAssignmentName, options)
	if err != nil {
		return ConfigurationAssignmentsForSubscriptionsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConfigurationAssignmentsForSubscriptionsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ConfigurationAssignmentsForSubscriptionsClientDeleteResponse{}, err
	}
	resp, err := client.deleteHandleResponse(httpResp)
	return resp, err
}

// deleteCreateRequest creates the Delete request.
func (client *ConfigurationAssignmentsForSubscriptionsClient) deleteCreateRequest(ctx context.Context, configurationAssignmentName string, options *ConfigurationAssignmentsForSubscriptionsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Maintenance/configurationAssignments/{configurationAssignmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if configurationAssignmentName == "" {
		return nil, errors.New("parameter configurationAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationAssignmentName}", url.PathEscape(configurationAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *ConfigurationAssignmentsForSubscriptionsClient) deleteHandleResponse(resp *http.Response) (ConfigurationAssignmentsForSubscriptionsClientDeleteResponse, error) {
	result := ConfigurationAssignmentsForSubscriptionsClientDeleteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationAssignment); err != nil {
		return ConfigurationAssignmentsForSubscriptionsClientDeleteResponse{}, err
	}
	return result, nil
}

// Get - Get configuration assignment for resource..
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - configurationAssignmentName - Configuration assignment name
//   - options - ConfigurationAssignmentsForSubscriptionsClientGetOptions contains the optional parameters for the ConfigurationAssignmentsForSubscriptionsClient.Get
//     method.
func (client *ConfigurationAssignmentsForSubscriptionsClient) Get(ctx context.Context, configurationAssignmentName string, options *ConfigurationAssignmentsForSubscriptionsClientGetOptions) (ConfigurationAssignmentsForSubscriptionsClientGetResponse, error) {
	var err error
	const operationName = "ConfigurationAssignmentsForSubscriptionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, configurationAssignmentName, options)
	if err != nil {
		return ConfigurationAssignmentsForSubscriptionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConfigurationAssignmentsForSubscriptionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConfigurationAssignmentsForSubscriptionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ConfigurationAssignmentsForSubscriptionsClient) getCreateRequest(ctx context.Context, configurationAssignmentName string, options *ConfigurationAssignmentsForSubscriptionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Maintenance/configurationAssignments/{configurationAssignmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if configurationAssignmentName == "" {
		return nil, errors.New("parameter configurationAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationAssignmentName}", url.PathEscape(configurationAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConfigurationAssignmentsForSubscriptionsClient) getHandleResponse(resp *http.Response) (ConfigurationAssignmentsForSubscriptionsClientGetResponse, error) {
	result := ConfigurationAssignmentsForSubscriptionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationAssignment); err != nil {
		return ConfigurationAssignmentsForSubscriptionsClientGetResponse{}, err
	}
	return result, nil
}

// Update - Register configuration for resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - configurationAssignmentName - Configuration assignment name
//   - configurationAssignment - The configurationAssignment
//   - options - ConfigurationAssignmentsForSubscriptionsClientUpdateOptions contains the optional parameters for the ConfigurationAssignmentsForSubscriptionsClient.Update
//     method.
func (client *ConfigurationAssignmentsForSubscriptionsClient) Update(ctx context.Context, configurationAssignmentName string, configurationAssignment ConfigurationAssignment, options *ConfigurationAssignmentsForSubscriptionsClientUpdateOptions) (ConfigurationAssignmentsForSubscriptionsClientUpdateResponse, error) {
	var err error
	const operationName = "ConfigurationAssignmentsForSubscriptionsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, configurationAssignmentName, configurationAssignment, options)
	if err != nil {
		return ConfigurationAssignmentsForSubscriptionsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConfigurationAssignmentsForSubscriptionsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConfigurationAssignmentsForSubscriptionsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *ConfigurationAssignmentsForSubscriptionsClient) updateCreateRequest(ctx context.Context, configurationAssignmentName string, configurationAssignment ConfigurationAssignment, options *ConfigurationAssignmentsForSubscriptionsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Maintenance/configurationAssignments/{configurationAssignmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if configurationAssignmentName == "" {
		return nil, errors.New("parameter configurationAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationAssignmentName}", url.PathEscape(configurationAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, configurationAssignment); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ConfigurationAssignmentsForSubscriptionsClient) updateHandleResponse(resp *http.Response) (ConfigurationAssignmentsForSubscriptionsClientUpdateResponse, error) {
	result := ConfigurationAssignmentsForSubscriptionsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationAssignment); err != nil {
		return ConfigurationAssignmentsForSubscriptionsClientUpdateResponse{}, err
	}
	return result, nil
}
