//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// TenantAccessClient contains the methods for the TenantAccess group.
// Don't use this type directly, use NewTenantAccessClient() instead.
type TenantAccessClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewTenantAccessClient creates a new instance of TenantAccessClient with the specified values.
func NewTenantAccessClient(con *arm.Connection, subscriptionID string) *TenantAccessClient {
	return &TenantAccessClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Create - Update tenant access information details.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TenantAccessClient) Create(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, ifMatch string, parameters AccessInformationCreateParameters, options *TenantAccessCreateOptions) (TenantAccessCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, serviceName, accessName, ifMatch, parameters, options)
	if err != nil {
		return TenantAccessCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TenantAccessCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TenantAccessCreateResponse{}, client.createHandleError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *TenantAccessClient) createCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, ifMatch string, parameters AccessInformationCreateParameters, options *TenantAccessCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if accessName == "" {
		return nil, errors.New("parameter accessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessName}", url.PathEscape(string(accessName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("If-Match", ifMatch)
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createHandleResponse handles the Create response.
func (client *TenantAccessClient) createHandleResponse(resp *http.Response) (TenantAccessCreateResponse, error) {
	result := TenantAccessCreateResponse{RawResponse: resp}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessInformationContract); err != nil {
		return TenantAccessCreateResponse{}, err
	}
	return result, nil
}

// createHandleError handles the Create error response.
func (client *TenantAccessClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Get tenant access information details without secrets.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TenantAccessClient) Get(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, options *TenantAccessGetOptions) (TenantAccessGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, accessName, options)
	if err != nil {
		return TenantAccessGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TenantAccessGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TenantAccessGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TenantAccessClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, options *TenantAccessGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}"
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
	if accessName == "" {
		return nil, errors.New("parameter accessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessName}", url.PathEscape(string(accessName)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TenantAccessClient) getHandleResponse(resp *http.Response) (TenantAccessGetResponse, error) {
	result := TenantAccessGetResponse{RawResponse: resp}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessInformationContract); err != nil {
		return TenantAccessGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *TenantAccessClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetEntityTag - Tenant access metadata
// If the operation fails it returns the *ErrorResponse error type.
func (client *TenantAccessClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, options *TenantAccessGetEntityTagOptions) (TenantAccessGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, accessName, options)
	if err != nil {
		return TenantAccessGetEntityTagResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TenantAccessGetEntityTagResponse{}, err
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *TenantAccessClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, options *TenantAccessGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}"
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
	if accessName == "" {
		return nil, errors.New("parameter accessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessName}", url.PathEscape(string(accessName)))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *TenantAccessClient) getEntityTagHandleResponse(resp *http.Response) (TenantAccessGetEntityTagResponse, error) {
	result := TenantAccessGetEntityTagResponse{RawResponse: resp}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// ListByService - Returns list of access infos - for Git and Management endpoints.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TenantAccessClient) ListByService(resourceGroupName string, serviceName string, options *TenantAccessListByServiceOptions) *TenantAccessListByServicePager {
	return &TenantAccessListByServicePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, options)
		},
		advancer: func(ctx context.Context, resp TenantAccessListByServiceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AccessInformationCollection.NextLink)
		},
	}
}

// listByServiceCreateRequest creates the ListByService request.
func (client *TenantAccessClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *TenantAccessListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2021-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *TenantAccessClient) listByServiceHandleResponse(resp *http.Response) (TenantAccessListByServiceResponse, error) {
	result := TenantAccessListByServiceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessInformationCollection); err != nil {
		return TenantAccessListByServiceResponse{}, err
	}
	return result, nil
}

// listByServiceHandleError handles the ListByService error response.
func (client *TenantAccessClient) listByServiceHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListSecrets - Get tenant access information details.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TenantAccessClient) ListSecrets(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, options *TenantAccessListSecretsOptions) (TenantAccessListSecretsResponse, error) {
	req, err := client.listSecretsCreateRequest(ctx, resourceGroupName, serviceName, accessName, options)
	if err != nil {
		return TenantAccessListSecretsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TenantAccessListSecretsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TenantAccessListSecretsResponse{}, client.listSecretsHandleError(resp)
	}
	return client.listSecretsHandleResponse(resp)
}

// listSecretsCreateRequest creates the ListSecrets request.
func (client *TenantAccessClient) listSecretsCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, options *TenantAccessListSecretsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}/listSecrets"
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
	if accessName == "" {
		return nil, errors.New("parameter accessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessName}", url.PathEscape(string(accessName)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listSecretsHandleResponse handles the ListSecrets response.
func (client *TenantAccessClient) listSecretsHandleResponse(resp *http.Response) (TenantAccessListSecretsResponse, error) {
	result := TenantAccessListSecretsResponse{RawResponse: resp}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessInformationSecretsContract); err != nil {
		return TenantAccessListSecretsResponse{}, err
	}
	return result, nil
}

// listSecretsHandleError handles the ListSecrets error response.
func (client *TenantAccessClient) listSecretsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// RegeneratePrimaryKey - Regenerate primary access key
// If the operation fails it returns the *ErrorResponse error type.
func (client *TenantAccessClient) RegeneratePrimaryKey(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, options *TenantAccessRegeneratePrimaryKeyOptions) (TenantAccessRegeneratePrimaryKeyResponse, error) {
	req, err := client.regeneratePrimaryKeyCreateRequest(ctx, resourceGroupName, serviceName, accessName, options)
	if err != nil {
		return TenantAccessRegeneratePrimaryKeyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TenantAccessRegeneratePrimaryKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return TenantAccessRegeneratePrimaryKeyResponse{}, client.regeneratePrimaryKeyHandleError(resp)
	}
	return TenantAccessRegeneratePrimaryKeyResponse{RawResponse: resp}, nil
}

// regeneratePrimaryKeyCreateRequest creates the RegeneratePrimaryKey request.
func (client *TenantAccessClient) regeneratePrimaryKeyCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, options *TenantAccessRegeneratePrimaryKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}/regeneratePrimaryKey"
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
	if accessName == "" {
		return nil, errors.New("parameter accessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessName}", url.PathEscape(string(accessName)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// regeneratePrimaryKeyHandleError handles the RegeneratePrimaryKey error response.
func (client *TenantAccessClient) regeneratePrimaryKeyHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// RegenerateSecondaryKey - Regenerate secondary access key
// If the operation fails it returns the *ErrorResponse error type.
func (client *TenantAccessClient) RegenerateSecondaryKey(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, options *TenantAccessRegenerateSecondaryKeyOptions) (TenantAccessRegenerateSecondaryKeyResponse, error) {
	req, err := client.regenerateSecondaryKeyCreateRequest(ctx, resourceGroupName, serviceName, accessName, options)
	if err != nil {
		return TenantAccessRegenerateSecondaryKeyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TenantAccessRegenerateSecondaryKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return TenantAccessRegenerateSecondaryKeyResponse{}, client.regenerateSecondaryKeyHandleError(resp)
	}
	return TenantAccessRegenerateSecondaryKeyResponse{RawResponse: resp}, nil
}

// regenerateSecondaryKeyCreateRequest creates the RegenerateSecondaryKey request.
func (client *TenantAccessClient) regenerateSecondaryKeyCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, options *TenantAccessRegenerateSecondaryKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}/regenerateSecondaryKey"
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
	if accessName == "" {
		return nil, errors.New("parameter accessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessName}", url.PathEscape(string(accessName)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// regenerateSecondaryKeyHandleError handles the RegenerateSecondaryKey error response.
func (client *TenantAccessClient) regenerateSecondaryKeyHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Update - Update tenant access information details.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TenantAccessClient) Update(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, ifMatch string, parameters AccessInformationUpdateParameters, options *TenantAccessUpdateOptions) (TenantAccessUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, accessName, ifMatch, parameters, options)
	if err != nil {
		return TenantAccessUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TenantAccessUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TenantAccessUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *TenantAccessClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, ifMatch string, parameters AccessInformationUpdateParameters, options *TenantAccessUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if accessName == "" {
		return nil, errors.New("parameter accessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessName}", url.PathEscape(string(accessName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("If-Match", ifMatch)
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *TenantAccessClient) updateHandleResponse(resp *http.Response) (TenantAccessUpdateResponse, error) {
	result := TenantAccessUpdateResponse{RawResponse: resp}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessInformationContract); err != nil {
		return TenantAccessUpdateResponse{}, err
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *TenantAccessClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
