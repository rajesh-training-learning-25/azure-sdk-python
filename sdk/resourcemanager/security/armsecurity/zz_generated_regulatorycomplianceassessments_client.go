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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RegulatoryComplianceAssessmentsClient contains the methods for the RegulatoryComplianceAssessments group.
// Don't use this type directly, use NewRegulatoryComplianceAssessmentsClient() instead.
type RegulatoryComplianceAssessmentsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewRegulatoryComplianceAssessmentsClient creates a new instance of RegulatoryComplianceAssessmentsClient with the specified values.
func NewRegulatoryComplianceAssessmentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *RegulatoryComplianceAssessmentsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &RegulatoryComplianceAssessmentsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Supported regulatory compliance details and state for selected assessment
// If the operation fails it returns the *CloudError error type.
func (client *RegulatoryComplianceAssessmentsClient) Get(ctx context.Context, regulatoryComplianceStandardName string, regulatoryComplianceControlName string, regulatoryComplianceAssessmentName string, options *RegulatoryComplianceAssessmentsGetOptions) (RegulatoryComplianceAssessmentsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, regulatoryComplianceStandardName, regulatoryComplianceControlName, regulatoryComplianceAssessmentName, options)
	if err != nil {
		return RegulatoryComplianceAssessmentsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RegulatoryComplianceAssessmentsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RegulatoryComplianceAssessmentsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RegulatoryComplianceAssessmentsClient) getCreateRequest(ctx context.Context, regulatoryComplianceStandardName string, regulatoryComplianceControlName string, regulatoryComplianceAssessmentName string, options *RegulatoryComplianceAssessmentsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/regulatoryComplianceStandards/{regulatoryComplianceStandardName}/regulatoryComplianceControls/{regulatoryComplianceControlName}/regulatoryComplianceAssessments/{regulatoryComplianceAssessmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if regulatoryComplianceStandardName == "" {
		return nil, errors.New("parameter regulatoryComplianceStandardName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regulatoryComplianceStandardName}", url.PathEscape(regulatoryComplianceStandardName))
	if regulatoryComplianceControlName == "" {
		return nil, errors.New("parameter regulatoryComplianceControlName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regulatoryComplianceControlName}", url.PathEscape(regulatoryComplianceControlName))
	if regulatoryComplianceAssessmentName == "" {
		return nil, errors.New("parameter regulatoryComplianceAssessmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regulatoryComplianceAssessmentName}", url.PathEscape(regulatoryComplianceAssessmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RegulatoryComplianceAssessmentsClient) getHandleResponse(resp *http.Response) (RegulatoryComplianceAssessmentsGetResponse, error) {
	result := RegulatoryComplianceAssessmentsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegulatoryComplianceAssessment); err != nil {
		return RegulatoryComplianceAssessmentsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *RegulatoryComplianceAssessmentsClient) getHandleError(resp *http.Response) error {
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

// List - Details and state of assessments mapped to selected regulatory compliance control
// If the operation fails it returns the *CloudError error type.
func (client *RegulatoryComplianceAssessmentsClient) List(regulatoryComplianceStandardName string, regulatoryComplianceControlName string, options *RegulatoryComplianceAssessmentsListOptions) *RegulatoryComplianceAssessmentsListPager {
	return &RegulatoryComplianceAssessmentsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, regulatoryComplianceStandardName, regulatoryComplianceControlName, options)
		},
		advancer: func(ctx context.Context, resp RegulatoryComplianceAssessmentsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RegulatoryComplianceAssessmentList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *RegulatoryComplianceAssessmentsClient) listCreateRequest(ctx context.Context, regulatoryComplianceStandardName string, regulatoryComplianceControlName string, options *RegulatoryComplianceAssessmentsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/regulatoryComplianceStandards/{regulatoryComplianceStandardName}/regulatoryComplianceControls/{regulatoryComplianceControlName}/regulatoryComplianceAssessments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if regulatoryComplianceStandardName == "" {
		return nil, errors.New("parameter regulatoryComplianceStandardName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regulatoryComplianceStandardName}", url.PathEscape(regulatoryComplianceStandardName))
	if regulatoryComplianceControlName == "" {
		return nil, errors.New("parameter regulatoryComplianceControlName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regulatoryComplianceControlName}", url.PathEscape(regulatoryComplianceControlName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-01-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RegulatoryComplianceAssessmentsClient) listHandleResponse(resp *http.Response) (RegulatoryComplianceAssessmentsListResponse, error) {
	result := RegulatoryComplianceAssessmentsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegulatoryComplianceAssessmentList); err != nil {
		return RegulatoryComplianceAssessmentsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *RegulatoryComplianceAssessmentsClient) listHandleError(resp *http.Response) error {
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
