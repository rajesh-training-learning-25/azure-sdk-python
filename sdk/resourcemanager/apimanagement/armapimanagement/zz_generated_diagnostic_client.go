//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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

// DiagnosticClient contains the methods for the Diagnostic group.
// Don't use this type directly, use NewDiagnosticClient() instead.
type DiagnosticClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDiagnosticClient creates a new instance of DiagnosticClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDiagnosticClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DiagnosticClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &DiagnosticClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates a new Diagnostic or updates an existing one.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// diagnosticID - Diagnostic identifier. Must be unique in the current API Management service instance.
// parameters - Create parameters.
// options - DiagnosticClientCreateOrUpdateOptions contains the optional parameters for the DiagnosticClient.CreateOrUpdate
// method.
func (client *DiagnosticClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, diagnosticID string, parameters DiagnosticContract, options *DiagnosticClientCreateOrUpdateOptions) (DiagnosticClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, diagnosticID, parameters, options)
	if err != nil {
		return DiagnosticClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DiagnosticClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return DiagnosticClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DiagnosticClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, diagnosticID string, parameters DiagnosticContract, options *DiagnosticClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/diagnostics/{diagnosticId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if diagnosticID == "" {
		return nil, errors.New("parameter diagnosticID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diagnosticId}", url.PathEscape(diagnosticID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DiagnosticClient) createOrUpdateHandleResponse(resp *http.Response) (DiagnosticClientCreateOrUpdateResponse, error) {
	result := DiagnosticClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiagnosticContract); err != nil {
		return DiagnosticClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the specified Diagnostic.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// diagnosticID - Diagnostic identifier. Must be unique in the current API Management service instance.
// ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
// it should be * for unconditional update.
// options - DiagnosticClientDeleteOptions contains the optional parameters for the DiagnosticClient.Delete method.
func (client *DiagnosticClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, diagnosticID string, ifMatch string, options *DiagnosticClientDeleteOptions) (DiagnosticClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, diagnosticID, ifMatch, options)
	if err != nil {
		return DiagnosticClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DiagnosticClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return DiagnosticClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return DiagnosticClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DiagnosticClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, diagnosticID string, ifMatch string, options *DiagnosticClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/diagnostics/{diagnosticId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if diagnosticID == "" {
		return nil, errors.New("parameter diagnosticID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diagnosticId}", url.PathEscape(diagnosticID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("If-Match", ifMatch)
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the details of the Diagnostic specified by its identifier.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// diagnosticID - Diagnostic identifier. Must be unique in the current API Management service instance.
// options - DiagnosticClientGetOptions contains the optional parameters for the DiagnosticClient.Get method.
func (client *DiagnosticClient) Get(ctx context.Context, resourceGroupName string, serviceName string, diagnosticID string, options *DiagnosticClientGetOptions) (DiagnosticClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, diagnosticID, options)
	if err != nil {
		return DiagnosticClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DiagnosticClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DiagnosticClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DiagnosticClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, diagnosticID string, options *DiagnosticClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/diagnostics/{diagnosticId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if diagnosticID == "" {
		return nil, errors.New("parameter diagnosticID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diagnosticId}", url.PathEscape(diagnosticID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DiagnosticClient) getHandleResponse(resp *http.Response) (DiagnosticClientGetResponse, error) {
	result := DiagnosticClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiagnosticContract); err != nil {
		return DiagnosticClientGetResponse{}, err
	}
	return result, nil
}

// GetEntityTag - Gets the entity state (Etag) version of the Diagnostic specified by its identifier.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// diagnosticID - Diagnostic identifier. Must be unique in the current API Management service instance.
// options - DiagnosticClientGetEntityTagOptions contains the optional parameters for the DiagnosticClient.GetEntityTag method.
func (client *DiagnosticClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, diagnosticID string, options *DiagnosticClientGetEntityTagOptions) (DiagnosticClientGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, diagnosticID, options)
	if err != nil {
		return DiagnosticClientGetEntityTagResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DiagnosticClientGetEntityTagResponse{}, err
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *DiagnosticClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, diagnosticID string, options *DiagnosticClientGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/diagnostics/{diagnosticId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if diagnosticID == "" {
		return nil, errors.New("parameter diagnosticID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diagnosticId}", url.PathEscape(diagnosticID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *DiagnosticClient) getEntityTagHandleResponse(resp *http.Response) (DiagnosticClientGetEntityTagResponse, error) {
	result := DiagnosticClientGetEntityTagResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// ListByService - Lists all diagnostics of the API Management service instance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// options - DiagnosticClientListByServiceOptions contains the optional parameters for the DiagnosticClient.ListByService
// method.
func (client *DiagnosticClient) ListByService(resourceGroupName string, serviceName string, options *DiagnosticClientListByServiceOptions) *runtime.Pager[DiagnosticClientListByServiceResponse] {
	return runtime.NewPager(runtime.PageProcessor[DiagnosticClientListByServiceResponse]{
		More: func(page DiagnosticClientListByServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DiagnosticClientListByServiceResponse) (DiagnosticClientListByServiceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DiagnosticClientListByServiceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DiagnosticClientListByServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DiagnosticClientListByServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServiceHandleResponse(resp)
		},
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *DiagnosticClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *DiagnosticClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/diagnostics"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *DiagnosticClient) listByServiceHandleResponse(resp *http.Response) (DiagnosticClientListByServiceResponse, error) {
	result := DiagnosticClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiagnosticCollection); err != nil {
		return DiagnosticClientListByServiceResponse{}, err
	}
	return result, nil
}

// Update - Updates the details of the Diagnostic specified by its identifier.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// diagnosticID - Diagnostic identifier. Must be unique in the current API Management service instance.
// ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
// it should be * for unconditional update.
// parameters - Diagnostic Update parameters.
// options - DiagnosticClientUpdateOptions contains the optional parameters for the DiagnosticClient.Update method.
func (client *DiagnosticClient) Update(ctx context.Context, resourceGroupName string, serviceName string, diagnosticID string, ifMatch string, parameters DiagnosticContract, options *DiagnosticClientUpdateOptions) (DiagnosticClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, diagnosticID, ifMatch, parameters, options)
	if err != nil {
		return DiagnosticClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DiagnosticClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DiagnosticClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *DiagnosticClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, diagnosticID string, ifMatch string, parameters DiagnosticContract, options *DiagnosticClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/diagnostics/{diagnosticId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if diagnosticID == "" {
		return nil, errors.New("parameter diagnosticID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diagnosticId}", url.PathEscape(diagnosticID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("If-Match", ifMatch)
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *DiagnosticClient) updateHandleResponse(resp *http.Response) (DiagnosticClientUpdateResponse, error) {
	result := DiagnosticClientUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiagnosticContract); err != nil {
		return DiagnosticClientUpdateResponse{}, err
	}
	return result, nil
}
