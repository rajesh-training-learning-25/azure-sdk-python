//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapp

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ManagedEnvironmentsStoragesClient contains the methods for the ManagedEnvironmentsStorages group.
// Don't use this type directly, use NewManagedEnvironmentsStoragesClient() instead.
type ManagedEnvironmentsStoragesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewManagedEnvironmentsStoragesClient creates a new instance of ManagedEnvironmentsStoragesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewManagedEnvironmentsStoragesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ManagedEnvironmentsStoragesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ManagedEnvironmentsStoragesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Create or update storage for a managedEnvironment.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// envName - Name of the Environment.
// name - Name of the storage.
// storageEnvelope - Configuration details of storage.
// options - ManagedEnvironmentsStoragesClientCreateOrUpdateOptions contains the optional parameters for the ManagedEnvironmentsStoragesClient.CreateOrUpdate
// method.
func (client *ManagedEnvironmentsStoragesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, envName string, name string, storageEnvelope ManagedEnvironmentStorage, options *ManagedEnvironmentsStoragesClientCreateOrUpdateOptions) (ManagedEnvironmentsStoragesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, envName, name, storageEnvelope, options)
	if err != nil {
		return ManagedEnvironmentsStoragesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedEnvironmentsStoragesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedEnvironmentsStoragesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ManagedEnvironmentsStoragesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, envName string, name string, storageEnvelope ManagedEnvironmentStorage, options *ManagedEnvironmentsStoragesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{envName}/storages/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if envName == "" {
		return nil, errors.New("parameter envName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{envName}", url.PathEscape(envName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, storageEnvelope)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ManagedEnvironmentsStoragesClient) createOrUpdateHandleResponse(resp *http.Response) (ManagedEnvironmentsStoragesClientCreateOrUpdateResponse, error) {
	result := ManagedEnvironmentsStoragesClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedEnvironmentStorage); err != nil {
		return ManagedEnvironmentsStoragesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete storage for a managedEnvironment.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// envName - Name of the Environment.
// name - Name of the storage.
// options - ManagedEnvironmentsStoragesClientDeleteOptions contains the optional parameters for the ManagedEnvironmentsStoragesClient.Delete
// method.
func (client *ManagedEnvironmentsStoragesClient) Delete(ctx context.Context, resourceGroupName string, envName string, name string, options *ManagedEnvironmentsStoragesClientDeleteOptions) (ManagedEnvironmentsStoragesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, envName, name, options)
	if err != nil {
		return ManagedEnvironmentsStoragesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedEnvironmentsStoragesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ManagedEnvironmentsStoragesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ManagedEnvironmentsStoragesClientDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ManagedEnvironmentsStoragesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, envName string, name string, options *ManagedEnvironmentsStoragesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{envName}/storages/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if envName == "" {
		return nil, errors.New("parameter envName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{envName}", url.PathEscape(envName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get storage for a managedEnvironment.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// envName - Name of the Environment.
// name - Name of the storage.
// options - ManagedEnvironmentsStoragesClientGetOptions contains the optional parameters for the ManagedEnvironmentsStoragesClient.Get
// method.
func (client *ManagedEnvironmentsStoragesClient) Get(ctx context.Context, resourceGroupName string, envName string, name string, options *ManagedEnvironmentsStoragesClientGetOptions) (ManagedEnvironmentsStoragesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, envName, name, options)
	if err != nil {
		return ManagedEnvironmentsStoragesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedEnvironmentsStoragesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedEnvironmentsStoragesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ManagedEnvironmentsStoragesClient) getCreateRequest(ctx context.Context, resourceGroupName string, envName string, name string, options *ManagedEnvironmentsStoragesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{envName}/storages/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if envName == "" {
		return nil, errors.New("parameter envName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{envName}", url.PathEscape(envName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedEnvironmentsStoragesClient) getHandleResponse(resp *http.Response) (ManagedEnvironmentsStoragesClientGetResponse, error) {
	result := ManagedEnvironmentsStoragesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedEnvironmentStorage); err != nil {
		return ManagedEnvironmentsStoragesClientGetResponse{}, err
	}
	return result, nil
}

// List - Get all storages for a managedEnvironment.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// envName - Name of the Environment.
// options - ManagedEnvironmentsStoragesClientListOptions contains the optional parameters for the ManagedEnvironmentsStoragesClient.List
// method.
func (client *ManagedEnvironmentsStoragesClient) List(ctx context.Context, resourceGroupName string, envName string, options *ManagedEnvironmentsStoragesClientListOptions) (ManagedEnvironmentsStoragesClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, envName, options)
	if err != nil {
		return ManagedEnvironmentsStoragesClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedEnvironmentsStoragesClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedEnvironmentsStoragesClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ManagedEnvironmentsStoragesClient) listCreateRequest(ctx context.Context, resourceGroupName string, envName string, options *ManagedEnvironmentsStoragesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{envName}/storages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if envName == "" {
		return nil, errors.New("parameter envName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{envName}", url.PathEscape(envName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ManagedEnvironmentsStoragesClient) listHandleResponse(resp *http.Response) (ManagedEnvironmentsStoragesClientListResponse, error) {
	result := ManagedEnvironmentsStoragesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedEnvironmentStoragesCollection); err != nil {
		return ManagedEnvironmentsStoragesClientListResponse{}, err
	}
	return result, nil
}
