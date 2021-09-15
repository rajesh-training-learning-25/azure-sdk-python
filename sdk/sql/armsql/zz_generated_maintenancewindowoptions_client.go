//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// MaintenanceWindowOptionsClient contains the methods for the MaintenanceWindowOptions group.
// Don't use this type directly, use NewMaintenanceWindowOptionsClient() instead.
type MaintenanceWindowOptionsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewMaintenanceWindowOptionsClient creates a new instance of MaintenanceWindowOptionsClient with the specified values.
func NewMaintenanceWindowOptionsClient(con *arm.Connection, subscriptionID string) *MaintenanceWindowOptionsClient {
	return &MaintenanceWindowOptionsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Get - Gets a list of available maintenance windows.
// If the operation fails it returns a generic error.
func (client *MaintenanceWindowOptionsClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, maintenanceWindowOptionsName string, options *MaintenanceWindowOptionsGetOptions) (MaintenanceWindowOptionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, databaseName, maintenanceWindowOptionsName, options)
	if err != nil {
		return MaintenanceWindowOptionsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MaintenanceWindowOptionsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MaintenanceWindowOptionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *MaintenanceWindowOptionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, maintenanceWindowOptionsName string, options *MaintenanceWindowOptionsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/maintenanceWindowOptions/current"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("maintenanceWindowOptionsName", maintenanceWindowOptionsName)
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MaintenanceWindowOptionsClient) getHandleResponse(resp *http.Response) (MaintenanceWindowOptionsGetResponse, error) {
	result := MaintenanceWindowOptionsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MaintenanceWindowOptions); err != nil {
		return MaintenanceWindowOptionsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *MaintenanceWindowOptionsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
