//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

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

// InformationProtectionPoliciesClient contains the methods for the InformationProtectionPolicies group.
// Don't use this type directly, use NewInformationProtectionPoliciesClient() instead.
type InformationProtectionPoliciesClient struct {
	ep string
	pl runtime.Pipeline
}

// NewInformationProtectionPoliciesClient creates a new instance of InformationProtectionPoliciesClient with the specified values.
func NewInformationProtectionPoliciesClient(con *arm.Connection) *InformationProtectionPoliciesClient {
	return &InformationProtectionPoliciesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version)}
}

// CreateOrUpdate - Details of the information protection policy.
// If the operation fails it returns the *CloudError error type.
func (client *InformationProtectionPoliciesClient) CreateOrUpdate(ctx context.Context, scope string, informationProtectionPolicyName Enum18, informationProtectionPolicy InformationProtectionPolicy, options *InformationProtectionPoliciesCreateOrUpdateOptions) (InformationProtectionPoliciesCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, scope, informationProtectionPolicyName, informationProtectionPolicy, options)
	if err != nil {
		return InformationProtectionPoliciesCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return InformationProtectionPoliciesCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return InformationProtectionPoliciesCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *InformationProtectionPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, scope string, informationProtectionPolicyName Enum18, informationProtectionPolicy InformationProtectionPolicy, options *InformationProtectionPoliciesCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Security/informationProtectionPolicies/{informationProtectionPolicyName}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if informationProtectionPolicyName == "" {
		return nil, errors.New("parameter informationProtectionPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{informationProtectionPolicyName}", url.PathEscape(string(informationProtectionPolicyName)))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, informationProtectionPolicy)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *InformationProtectionPoliciesClient) createOrUpdateHandleResponse(resp *http.Response) (InformationProtectionPoliciesCreateOrUpdateResponse, error) {
	result := InformationProtectionPoliciesCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.InformationProtectionPolicy); err != nil {
		return InformationProtectionPoliciesCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *InformationProtectionPoliciesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Details of the information protection policy.
// If the operation fails it returns the *CloudError error type.
func (client *InformationProtectionPoliciesClient) Get(ctx context.Context, scope string, informationProtectionPolicyName Enum18, options *InformationProtectionPoliciesGetOptions) (InformationProtectionPoliciesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, informationProtectionPolicyName, options)
	if err != nil {
		return InformationProtectionPoliciesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return InformationProtectionPoliciesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return InformationProtectionPoliciesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *InformationProtectionPoliciesClient) getCreateRequest(ctx context.Context, scope string, informationProtectionPolicyName Enum18, options *InformationProtectionPoliciesGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Security/informationProtectionPolicies/{informationProtectionPolicyName}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if informationProtectionPolicyName == "" {
		return nil, errors.New("parameter informationProtectionPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{informationProtectionPolicyName}", url.PathEscape(string(informationProtectionPolicyName)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *InformationProtectionPoliciesClient) getHandleResponse(resp *http.Response) (InformationProtectionPoliciesGetResponse, error) {
	result := InformationProtectionPoliciesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.InformationProtectionPolicy); err != nil {
		return InformationProtectionPoliciesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *InformationProtectionPoliciesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Information protection policies of a specific management group.
// If the operation fails it returns the *CloudError error type.
func (client *InformationProtectionPoliciesClient) List(scope string, options *InformationProtectionPoliciesListOptions) *InformationProtectionPoliciesListPager {
	return &InformationProtectionPoliciesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, scope, options)
		},
		advancer: func(ctx context.Context, resp InformationProtectionPoliciesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.InformationProtectionPolicyList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *InformationProtectionPoliciesClient) listCreateRequest(ctx context.Context, scope string, options *InformationProtectionPoliciesListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Security/informationProtectionPolicies"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *InformationProtectionPoliciesClient) listHandleResponse(resp *http.Response) (InformationProtectionPoliciesListResponse, error) {
	result := InformationProtectionPoliciesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.InformationProtectionPolicyList); err != nil {
		return InformationProtectionPoliciesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *InformationProtectionPoliciesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
