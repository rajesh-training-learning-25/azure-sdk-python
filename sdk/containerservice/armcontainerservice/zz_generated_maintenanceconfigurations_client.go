//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerservice

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

// MaintenanceConfigurationsClient contains the methods for the MaintenanceConfigurations group.
// Don't use this type directly, use NewMaintenanceConfigurationsClient() instead.
type MaintenanceConfigurationsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewMaintenanceConfigurationsClient creates a new instance of MaintenanceConfigurationsClient with the specified values.
func NewMaintenanceConfigurationsClient(con *arm.Connection, subscriptionID string) *MaintenanceConfigurationsClient {
	return &MaintenanceConfigurationsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CreateOrUpdate - Creates or updates a maintenance configuration in the specified managed cluster.
// If the operation fails it returns the *CloudError error type.
func (client *MaintenanceConfigurationsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, configName string, parameters MaintenanceConfiguration, options *MaintenanceConfigurationsCreateOrUpdateOptions) (MaintenanceConfigurationsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, configName, parameters, options)
	if err != nil {
		return MaintenanceConfigurationsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MaintenanceConfigurationsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MaintenanceConfigurationsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *MaintenanceConfigurationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, configName string, parameters MaintenanceConfiguration, options *MaintenanceConfigurationsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/maintenanceConfigurations/{configName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if configName == "" {
		return nil, errors.New("parameter configName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configName}", url.PathEscape(configName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *MaintenanceConfigurationsClient) createOrUpdateHandleResponse(resp *http.Response) (MaintenanceConfigurationsCreateOrUpdateResponse, error) {
	result := MaintenanceConfigurationsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MaintenanceConfiguration); err != nil {
		return MaintenanceConfigurationsCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *MaintenanceConfigurationsClient) createOrUpdateHandleError(resp *http.Response) error {
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

// Delete - Deletes a maintenance configuration.
// If the operation fails it returns the *CloudError error type.
func (client *MaintenanceConfigurationsClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, configName string, options *MaintenanceConfigurationsDeleteOptions) (MaintenanceConfigurationsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, configName, options)
	if err != nil {
		return MaintenanceConfigurationsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MaintenanceConfigurationsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return MaintenanceConfigurationsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return MaintenanceConfigurationsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *MaintenanceConfigurationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, configName string, options *MaintenanceConfigurationsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/maintenanceConfigurations/{configName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if configName == "" {
		return nil, errors.New("parameter configName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configName}", url.PathEscape(configName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *MaintenanceConfigurationsClient) deleteHandleError(resp *http.Response) error {
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

// Get - Gets the specified maintenance configuration of a managed cluster.
// If the operation fails it returns the *CloudError error type.
func (client *MaintenanceConfigurationsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, configName string, options *MaintenanceConfigurationsGetOptions) (MaintenanceConfigurationsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, configName, options)
	if err != nil {
		return MaintenanceConfigurationsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MaintenanceConfigurationsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MaintenanceConfigurationsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *MaintenanceConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, configName string, options *MaintenanceConfigurationsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/maintenanceConfigurations/{configName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if configName == "" {
		return nil, errors.New("parameter configName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configName}", url.PathEscape(configName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MaintenanceConfigurationsClient) getHandleResponse(resp *http.Response) (MaintenanceConfigurationsGetResponse, error) {
	result := MaintenanceConfigurationsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MaintenanceConfiguration); err != nil {
		return MaintenanceConfigurationsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *MaintenanceConfigurationsClient) getHandleError(resp *http.Response) error {
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

// ListByManagedCluster - Gets a list of maintenance configurations in the specified managed cluster.
// If the operation fails it returns the *CloudError error type.
func (client *MaintenanceConfigurationsClient) ListByManagedCluster(resourceGroupName string, resourceName string, options *MaintenanceConfigurationsListByManagedClusterOptions) *MaintenanceConfigurationsListByManagedClusterPager {
	return &MaintenanceConfigurationsListByManagedClusterPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByManagedClusterCreateRequest(ctx, resourceGroupName, resourceName, options)
		},
		advancer: func(ctx context.Context, resp MaintenanceConfigurationsListByManagedClusterResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.MaintenanceConfigurationListResult.NextLink)
		},
	}
}

// listByManagedClusterCreateRequest creates the ListByManagedCluster request.
func (client *MaintenanceConfigurationsClient) listByManagedClusterCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *MaintenanceConfigurationsListByManagedClusterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/maintenanceConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByManagedClusterHandleResponse handles the ListByManagedCluster response.
func (client *MaintenanceConfigurationsClient) listByManagedClusterHandleResponse(resp *http.Response) (MaintenanceConfigurationsListByManagedClusterResponse, error) {
	result := MaintenanceConfigurationsListByManagedClusterResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MaintenanceConfigurationListResult); err != nil {
		return MaintenanceConfigurationsListByManagedClusterResponse{}, err
	}
	return result, nil
}

// listByManagedClusterHandleError handles the ListByManagedCluster error response.
func (client *MaintenanceConfigurationsClient) listByManagedClusterHandleError(resp *http.Response) error {
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
