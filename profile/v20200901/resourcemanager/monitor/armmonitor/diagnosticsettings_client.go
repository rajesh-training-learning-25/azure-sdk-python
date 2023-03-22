//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmonitor

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/profile/v20200901/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DiagnosticSettingsClient contains the methods for the DiagnosticSettings group.
// Don't use this type directly, use NewDiagnosticSettingsClient() instead.
type DiagnosticSettingsClient struct {
	internal *arm.Client
}

// NewDiagnosticSettingsClient creates a new instance of DiagnosticSettingsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDiagnosticSettingsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*DiagnosticSettingsClient, error) {
	cl, err := arm.NewClient(internal.ModuleName+"/armmonitor.DiagnosticSettingsClient", internal.ModuleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DiagnosticSettingsClient{
		internal: cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates diagnostic settings for the specified resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-05-01-preview
//   - resourceURI - The identifier of the resource.
//   - name - The name of the diagnostic setting.
//   - parameters - Parameters supplied to the operation.
//   - options - DiagnosticSettingsClientCreateOrUpdateOptions contains the optional parameters for the DiagnosticSettingsClient.CreateOrUpdate
//     method.
func (client *DiagnosticSettingsClient) CreateOrUpdate(ctx context.Context, resourceURI string, name string, parameters DiagnosticSettingsResource, options *DiagnosticSettingsClientCreateOrUpdateOptions) (DiagnosticSettingsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceURI, name, parameters, options)
	if err != nil {
		return DiagnosticSettingsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DiagnosticSettingsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DiagnosticSettingsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DiagnosticSettingsClient) createOrUpdateCreateRequest(ctx context.Context, resourceURI string, name string, parameters DiagnosticSettingsResource, options *DiagnosticSettingsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/diagnosticSettings/{name}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DiagnosticSettingsClient) createOrUpdateHandleResponse(resp *http.Response) (DiagnosticSettingsClientCreateOrUpdateResponse, error) {
	result := DiagnosticSettingsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiagnosticSettingsResource); err != nil {
		return DiagnosticSettingsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes existing diagnostic settings for the specified resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-05-01-preview
//   - resourceURI - The identifier of the resource.
//   - name - The name of the diagnostic setting.
//   - options - DiagnosticSettingsClientDeleteOptions contains the optional parameters for the DiagnosticSettingsClient.Delete
//     method.
func (client *DiagnosticSettingsClient) Delete(ctx context.Context, resourceURI string, name string, options *DiagnosticSettingsClientDeleteOptions) (DiagnosticSettingsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceURI, name, options)
	if err != nil {
		return DiagnosticSettingsClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DiagnosticSettingsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return DiagnosticSettingsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return DiagnosticSettingsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DiagnosticSettingsClient) deleteCreateRequest(ctx context.Context, resourceURI string, name string, options *DiagnosticSettingsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/diagnosticSettings/{name}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the active diagnostic settings for the specified resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-05-01-preview
//   - resourceURI - The identifier of the resource.
//   - name - The name of the diagnostic setting.
//   - options - DiagnosticSettingsClientGetOptions contains the optional parameters for the DiagnosticSettingsClient.Get method.
func (client *DiagnosticSettingsClient) Get(ctx context.Context, resourceURI string, name string, options *DiagnosticSettingsClientGetOptions) (DiagnosticSettingsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceURI, name, options)
	if err != nil {
		return DiagnosticSettingsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DiagnosticSettingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DiagnosticSettingsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DiagnosticSettingsClient) getCreateRequest(ctx context.Context, resourceURI string, name string, options *DiagnosticSettingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/diagnosticSettings/{name}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DiagnosticSettingsClient) getHandleResponse(resp *http.Response) (DiagnosticSettingsClientGetResponse, error) {
	result := DiagnosticSettingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiagnosticSettingsResource); err != nil {
		return DiagnosticSettingsClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets the active diagnostic settings list for the specified resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-05-01-preview
//   - resourceURI - The identifier of the resource.
//   - options - DiagnosticSettingsClientListOptions contains the optional parameters for the DiagnosticSettingsClient.List method.
func (client *DiagnosticSettingsClient) List(ctx context.Context, resourceURI string, options *DiagnosticSettingsClientListOptions) (DiagnosticSettingsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceURI, options)
	if err != nil {
		return DiagnosticSettingsClientListResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DiagnosticSettingsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DiagnosticSettingsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *DiagnosticSettingsClient) listCreateRequest(ctx context.Context, resourceURI string, options *DiagnosticSettingsClientListOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/diagnosticSettings"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DiagnosticSettingsClient) listHandleResponse(resp *http.Response) (DiagnosticSettingsClientListResponse, error) {
	result := DiagnosticSettingsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiagnosticSettingsResourceCollection); err != nil {
		return DiagnosticSettingsClientListResponse{}, err
	}
	return result, nil
}
