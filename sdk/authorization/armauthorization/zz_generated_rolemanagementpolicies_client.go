//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

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

// RoleManagementPoliciesClient contains the methods for the RoleManagementPolicies group.
// Don't use this type directly, use NewRoleManagementPoliciesClient() instead.
type RoleManagementPoliciesClient struct {
	ep string
	pl runtime.Pipeline
}

// NewRoleManagementPoliciesClient creates a new instance of RoleManagementPoliciesClient with the specified values.
func NewRoleManagementPoliciesClient(con *arm.Connection) *RoleManagementPoliciesClient {
	return &RoleManagementPoliciesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version)}
}

// Delete - Delete a role management policy
// If the operation fails it returns the *CloudError error type.
func (client *RoleManagementPoliciesClient) Delete(ctx context.Context, scope string, roleManagementPolicyName string, options *RoleManagementPoliciesDeleteOptions) (RoleManagementPoliciesDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, scope, roleManagementPolicyName, options)
	if err != nil {
		return RoleManagementPoliciesDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleManagementPoliciesDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return RoleManagementPoliciesDeleteResponse{}, client.deleteHandleError(resp)
	}
	return RoleManagementPoliciesDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RoleManagementPoliciesClient) deleteCreateRequest(ctx context.Context, scope string, roleManagementPolicyName string, options *RoleManagementPoliciesDeleteOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleManagementPolicies/{roleManagementPolicyName}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleManagementPolicyName == "" {
		return nil, errors.New("parameter roleManagementPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleManagementPolicyName}", url.PathEscape(roleManagementPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *RoleManagementPoliciesClient) deleteHandleError(resp *http.Response) error {
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

// Get - Get the specified role management policy for a resource scope
// If the operation fails it returns the *CloudError error type.
func (client *RoleManagementPoliciesClient) Get(ctx context.Context, scope string, roleManagementPolicyName string, options *RoleManagementPoliciesGetOptions) (RoleManagementPoliciesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, roleManagementPolicyName, options)
	if err != nil {
		return RoleManagementPoliciesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleManagementPoliciesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleManagementPoliciesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RoleManagementPoliciesClient) getCreateRequest(ctx context.Context, scope string, roleManagementPolicyName string, options *RoleManagementPoliciesGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleManagementPolicies/{roleManagementPolicyName}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleManagementPolicyName == "" {
		return nil, errors.New("parameter roleManagementPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleManagementPolicyName}", url.PathEscape(roleManagementPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RoleManagementPoliciesClient) getHandleResponse(resp *http.Response) (RoleManagementPoliciesGetResponse, error) {
	result := RoleManagementPoliciesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleManagementPolicy); err != nil {
		return RoleManagementPoliciesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *RoleManagementPoliciesClient) getHandleError(resp *http.Response) error {
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

// ListForScope - Gets role management policies for a resource scope.
// If the operation fails it returns the *CloudError error type.
func (client *RoleManagementPoliciesClient) ListForScope(scope string, options *RoleManagementPoliciesListForScopeOptions) *RoleManagementPoliciesListForScopePager {
	return &RoleManagementPoliciesListForScopePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listForScopeCreateRequest(ctx, scope, options)
		},
		advancer: func(ctx context.Context, resp RoleManagementPoliciesListForScopeResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RoleManagementPolicyListResult.NextLink)
		},
	}
}

// listForScopeCreateRequest creates the ListForScope request.
func (client *RoleManagementPoliciesClient) listForScopeCreateRequest(ctx context.Context, scope string, options *RoleManagementPoliciesListForScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleManagementPolicies"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listForScopeHandleResponse handles the ListForScope response.
func (client *RoleManagementPoliciesClient) listForScopeHandleResponse(resp *http.Response) (RoleManagementPoliciesListForScopeResponse, error) {
	result := RoleManagementPoliciesListForScopeResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleManagementPolicyListResult); err != nil {
		return RoleManagementPoliciesListForScopeResponse{}, err
	}
	return result, nil
}

// listForScopeHandleError handles the ListForScope error response.
func (client *RoleManagementPoliciesClient) listForScopeHandleError(resp *http.Response) error {
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

// Update - Update a role management policy
// If the operation fails it returns the *CloudError error type.
func (client *RoleManagementPoliciesClient) Update(ctx context.Context, scope string, roleManagementPolicyName string, parameters RoleManagementPolicy, options *RoleManagementPoliciesUpdateOptions) (RoleManagementPoliciesUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, scope, roleManagementPolicyName, parameters, options)
	if err != nil {
		return RoleManagementPoliciesUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleManagementPoliciesUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleManagementPoliciesUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *RoleManagementPoliciesClient) updateCreateRequest(ctx context.Context, scope string, roleManagementPolicyName string, parameters RoleManagementPolicy, options *RoleManagementPoliciesUpdateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleManagementPolicies/{roleManagementPolicyName}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleManagementPolicyName == "" {
		return nil, errors.New("parameter roleManagementPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleManagementPolicyName}", url.PathEscape(roleManagementPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *RoleManagementPoliciesClient) updateHandleResponse(resp *http.Response) (RoleManagementPoliciesUpdateResponse, error) {
	result := RoleManagementPoliciesUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleManagementPolicy); err != nil {
		return RoleManagementPoliciesUpdateResponse{}, err
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *RoleManagementPoliciesClient) updateHandleError(resp *http.Response) error {
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
