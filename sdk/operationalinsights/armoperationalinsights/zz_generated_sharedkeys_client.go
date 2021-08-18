// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoperationalinsights

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// SharedKeysClient contains the methods for the SharedKeys group.
// Don't use this type directly, use NewSharedKeysClient() instead.
type SharedKeysClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewSharedKeysClient creates a new instance of SharedKeysClient with the specified values.
func NewSharedKeysClient(con *armcore.Connection, subscriptionID string) *SharedKeysClient {
	return &SharedKeysClient{con: con, subscriptionID: subscriptionID}
}

// GetSharedKeys - Gets the shared keys for a workspace.
// If the operation fails it returns a generic error.
func (client *SharedKeysClient) GetSharedKeys(ctx context.Context, resourceGroupName string, workspaceName string, options *SharedKeysGetSharedKeysOptions) (SharedKeysGetSharedKeysResponse, error) {
	req, err := client.getSharedKeysCreateRequest(ctx, resourceGroupName, workspaceName, options)
	if err != nil {
		return SharedKeysGetSharedKeysResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SharedKeysGetSharedKeysResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SharedKeysGetSharedKeysResponse{}, client.getSharedKeysHandleError(resp)
	}
	return client.getSharedKeysHandleResponse(resp)
}

// getSharedKeysCreateRequest creates the GetSharedKeys request.
func (client *SharedKeysClient) getSharedKeysCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *SharedKeysGetSharedKeysOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/sharedKeys"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getSharedKeysHandleResponse handles the GetSharedKeys response.
func (client *SharedKeysClient) getSharedKeysHandleResponse(resp *azcore.Response) (SharedKeysGetSharedKeysResponse, error) {
	result := SharedKeysGetSharedKeysResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.SharedKeys); err != nil {
		return SharedKeysGetSharedKeysResponse{}, err
	}
	return result, nil
}

// getSharedKeysHandleError handles the GetSharedKeys error response.
func (client *SharedKeysClient) getSharedKeysHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Regenerate - Regenerates the shared keys for a Log Analytics Workspace. These keys are used to connect Microsoft Operational Insights agents to the workspace.
// If the operation fails it returns a generic error.
func (client *SharedKeysClient) Regenerate(ctx context.Context, resourceGroupName string, workspaceName string, options *SharedKeysRegenerateOptions) (SharedKeysRegenerateResponse, error) {
	req, err := client.regenerateCreateRequest(ctx, resourceGroupName, workspaceName, options)
	if err != nil {
		return SharedKeysRegenerateResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SharedKeysRegenerateResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SharedKeysRegenerateResponse{}, client.regenerateHandleError(resp)
	}
	return client.regenerateHandleResponse(resp)
}

// regenerateCreateRequest creates the Regenerate request.
func (client *SharedKeysClient) regenerateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *SharedKeysRegenerateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/regenerateSharedKey"
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
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// regenerateHandleResponse handles the Regenerate response.
func (client *SharedKeysClient) regenerateHandleResponse(resp *azcore.Response) (SharedKeysRegenerateResponse, error) {
	result := SharedKeysRegenerateResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.SharedKeys); err != nil {
		return SharedKeysRegenerateResponse{}, err
	}
	return result, nil
}

// regenerateHandleError handles the Regenerate error response.
func (client *SharedKeysClient) regenerateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
