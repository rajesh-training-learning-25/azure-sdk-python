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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AccessReviewInstancesAssignedForMyApprovalClient contains the methods for the AccessReviewInstancesAssignedForMyApproval group.
// Don't use this type directly, use NewAccessReviewInstancesAssignedForMyApprovalClient() instead.
type AccessReviewInstancesAssignedForMyApprovalClient struct {
	ep string
	pl runtime.Pipeline
}

// NewAccessReviewInstancesAssignedForMyApprovalClient creates a new instance of AccessReviewInstancesAssignedForMyApprovalClient with the specified values.
func NewAccessReviewInstancesAssignedForMyApprovalClient(credential azcore.TokenCredential, options *arm.ClientOptions) *AccessReviewInstancesAssignedForMyApprovalClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &AccessReviewInstancesAssignedForMyApprovalClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// GetByID - Get single access review instance assigned for my approval.
// If the operation fails it returns the *ErrorDefinition error type.
func (client *AccessReviewInstancesAssignedForMyApprovalClient) GetByID(ctx context.Context, scheduleDefinitionID string, id string, options *AccessReviewInstancesAssignedForMyApprovalGetByIDOptions) (AccessReviewInstancesAssignedForMyApprovalGetByIDResponse, error) {
	req, err := client.getByIDCreateRequest(ctx, scheduleDefinitionID, id, options)
	if err != nil {
		return AccessReviewInstancesAssignedForMyApprovalGetByIDResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccessReviewInstancesAssignedForMyApprovalGetByIDResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccessReviewInstancesAssignedForMyApprovalGetByIDResponse{}, client.getByIDHandleError(resp)
	}
	return client.getByIDHandleResponse(resp)
}

// getByIDCreateRequest creates the GetByID request.
func (client *AccessReviewInstancesAssignedForMyApprovalClient) getByIDCreateRequest(ctx context.Context, scheduleDefinitionID string, id string, options *AccessReviewInstancesAssignedForMyApprovalGetByIDOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances/{id}"
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getByIDHandleResponse handles the GetByID response.
func (client *AccessReviewInstancesAssignedForMyApprovalClient) getByIDHandleResponse(resp *http.Response) (AccessReviewInstancesAssignedForMyApprovalGetByIDResponse, error) {
	result := AccessReviewInstancesAssignedForMyApprovalGetByIDResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewInstance); err != nil {
		return AccessReviewInstancesAssignedForMyApprovalGetByIDResponse{}, err
	}
	return result, nil
}

// getByIDHandleError handles the GetByID error response.
func (client *AccessReviewInstancesAssignedForMyApprovalClient) getByIDHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorDefinition{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Get access review instances assigned for my approval.
// If the operation fails it returns the *ErrorDefinition error type.
func (client *AccessReviewInstancesAssignedForMyApprovalClient) List(scheduleDefinitionID string, options *AccessReviewInstancesAssignedForMyApprovalListOptions) *AccessReviewInstancesAssignedForMyApprovalListPager {
	return &AccessReviewInstancesAssignedForMyApprovalListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, scheduleDefinitionID, options)
		},
		advancer: func(ctx context.Context, resp AccessReviewInstancesAssignedForMyApprovalListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AccessReviewInstanceListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *AccessReviewInstancesAssignedForMyApprovalClient) listCreateRequest(ctx context.Context, scheduleDefinitionID string, options *AccessReviewInstancesAssignedForMyApprovalListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances"
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AccessReviewInstancesAssignedForMyApprovalClient) listHandleResponse(resp *http.Response) (AccessReviewInstancesAssignedForMyApprovalListResponse, error) {
	result := AccessReviewInstancesAssignedForMyApprovalListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewInstanceListResult); err != nil {
		return AccessReviewInstancesAssignedForMyApprovalListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *AccessReviewInstancesAssignedForMyApprovalClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorDefinition{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
