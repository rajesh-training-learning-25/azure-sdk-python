//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

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

// SoftwareUpdateConfigurationMachineRunsClient contains the methods for the SoftwareUpdateConfigurationMachineRuns group.
// Don't use this type directly, use NewSoftwareUpdateConfigurationMachineRunsClient() instead.
type SoftwareUpdateConfigurationMachineRunsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewSoftwareUpdateConfigurationMachineRunsClient creates a new instance of SoftwareUpdateConfigurationMachineRunsClient with the specified values.
func NewSoftwareUpdateConfigurationMachineRunsClient(con *arm.Connection, subscriptionID string) *SoftwareUpdateConfigurationMachineRunsClient {
	return &SoftwareUpdateConfigurationMachineRunsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// GetByID - Get a single software update configuration machine run by Id.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SoftwareUpdateConfigurationMachineRunsClient) GetByID(ctx context.Context, resourceGroupName string, automationAccountName string, softwareUpdateConfigurationMachineRunID string, options *SoftwareUpdateConfigurationMachineRunsGetByIDOptions) (SoftwareUpdateConfigurationMachineRunsGetByIDResponse, error) {
	req, err := client.getByIDCreateRequest(ctx, resourceGroupName, automationAccountName, softwareUpdateConfigurationMachineRunID, options)
	if err != nil {
		return SoftwareUpdateConfigurationMachineRunsGetByIDResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SoftwareUpdateConfigurationMachineRunsGetByIDResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SoftwareUpdateConfigurationMachineRunsGetByIDResponse{}, client.getByIDHandleError(resp)
	}
	return client.getByIDHandleResponse(resp)
}

// getByIDCreateRequest creates the GetByID request.
func (client *SoftwareUpdateConfigurationMachineRunsClient) getByIDCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, softwareUpdateConfigurationMachineRunID string, options *SoftwareUpdateConfigurationMachineRunsGetByIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/softwareUpdateConfigurationMachineRuns/{softwareUpdateConfigurationMachineRunId}"
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
	urlPath = strings.ReplaceAll(urlPath, "{softwareUpdateConfigurationMachineRunId}", url.PathEscape(softwareUpdateConfigurationMachineRunID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.ClientRequestID != nil {
		req.Raw().Header.Set("clientRequestId", *options.ClientRequestID)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getByIDHandleResponse handles the GetByID response.
func (client *SoftwareUpdateConfigurationMachineRunsClient) getByIDHandleResponse(resp *http.Response) (SoftwareUpdateConfigurationMachineRunsGetByIDResponse, error) {
	result := SoftwareUpdateConfigurationMachineRunsGetByIDResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SoftwareUpdateConfigurationMachineRun); err != nil {
		return SoftwareUpdateConfigurationMachineRunsGetByIDResponse{}, err
	}
	return result, nil
}

// getByIDHandleError handles the GetByID error response.
func (client *SoftwareUpdateConfigurationMachineRunsClient) getByIDHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Return list of software update configuration machine runs
// If the operation fails it returns the *ErrorResponse error type.
func (client *SoftwareUpdateConfigurationMachineRunsClient) List(ctx context.Context, resourceGroupName string, automationAccountName string, options *SoftwareUpdateConfigurationMachineRunsListOptions) (SoftwareUpdateConfigurationMachineRunsListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, automationAccountName, options)
	if err != nil {
		return SoftwareUpdateConfigurationMachineRunsListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SoftwareUpdateConfigurationMachineRunsListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SoftwareUpdateConfigurationMachineRunsListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *SoftwareUpdateConfigurationMachineRunsClient) listCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *SoftwareUpdateConfigurationMachineRunsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/softwareUpdateConfigurationMachineRuns"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
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
		req.Raw().Header.Set("clientRequestId", *options.ClientRequestID)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SoftwareUpdateConfigurationMachineRunsClient) listHandleResponse(resp *http.Response) (SoftwareUpdateConfigurationMachineRunsListResponse, error) {
	result := SoftwareUpdateConfigurationMachineRunsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SoftwareUpdateConfigurationMachineRunListResult); err != nil {
		return SoftwareUpdateConfigurationMachineRunsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *SoftwareUpdateConfigurationMachineRunsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
