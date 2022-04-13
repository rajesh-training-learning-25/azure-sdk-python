//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragesync

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ServerEndpointsClient contains the methods for the ServerEndpoints group.
// Don't use this type directly, use NewServerEndpointsClient() instead.
type ServerEndpointsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewServerEndpointsClient creates a new instance of ServerEndpointsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewServerEndpointsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServerEndpointsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ServerEndpointsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Create a new ServerEndpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageSyncServiceName - Name of Storage Sync Service resource.
// syncGroupName - Name of Sync Group resource.
// serverEndpointName - Name of Server Endpoint object.
// parameters - Body of Server Endpoint object.
// options - ServerEndpointsClientBeginCreateOptions contains the optional parameters for the ServerEndpointsClient.BeginCreate
// method.
func (client *ServerEndpointsClient) BeginCreate(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, serverEndpointName string, parameters ServerEndpointCreateParameters, options *ServerEndpointsClientBeginCreateOptions) (*armruntime.Poller[ServerEndpointsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, serverEndpointName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ServerEndpointsClientCreateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ServerEndpointsClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Create a new ServerEndpoint.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ServerEndpointsClient) create(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, serverEndpointName string, parameters ServerEndpointCreateParameters, options *ServerEndpointsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, serverEndpointName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *ServerEndpointsClient) createCreateRequest(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, serverEndpointName string, parameters ServerEndpointCreateParameters, options *ServerEndpointsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/serverEndpoints/{serverEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageSyncServiceName == "" {
		return nil, errors.New("parameter storageSyncServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageSyncServiceName}", url.PathEscape(storageSyncServiceName))
	if syncGroupName == "" {
		return nil, errors.New("parameter syncGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncGroupName}", url.PathEscape(syncGroupName))
	if serverEndpointName == "" {
		return nil, errors.New("parameter serverEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverEndpointName}", url.PathEscape(serverEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Delete a given ServerEndpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageSyncServiceName - Name of Storage Sync Service resource.
// syncGroupName - Name of Sync Group resource.
// serverEndpointName - Name of Server Endpoint object.
// options - ServerEndpointsClientBeginDeleteOptions contains the optional parameters for the ServerEndpointsClient.BeginDelete
// method.
func (client *ServerEndpointsClient) BeginDelete(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, serverEndpointName string, options *ServerEndpointsClientBeginDeleteOptions) (*armruntime.Poller[ServerEndpointsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, serverEndpointName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ServerEndpointsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ServerEndpointsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete a given ServerEndpoint.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ServerEndpointsClient) deleteOperation(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, serverEndpointName string, options *ServerEndpointsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, serverEndpointName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ServerEndpointsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, serverEndpointName string, options *ServerEndpointsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/serverEndpoints/{serverEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageSyncServiceName == "" {
		return nil, errors.New("parameter storageSyncServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageSyncServiceName}", url.PathEscape(storageSyncServiceName))
	if syncGroupName == "" {
		return nil, errors.New("parameter syncGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncGroupName}", url.PathEscape(syncGroupName))
	if serverEndpointName == "" {
		return nil, errors.New("parameter serverEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverEndpointName}", url.PathEscape(serverEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get a ServerEndpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageSyncServiceName - Name of Storage Sync Service resource.
// syncGroupName - Name of Sync Group resource.
// serverEndpointName - Name of Server Endpoint object.
// options - ServerEndpointsClientGetOptions contains the optional parameters for the ServerEndpointsClient.Get method.
func (client *ServerEndpointsClient) Get(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, serverEndpointName string, options *ServerEndpointsClientGetOptions) (ServerEndpointsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, serverEndpointName, options)
	if err != nil {
		return ServerEndpointsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServerEndpointsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServerEndpointsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServerEndpointsClient) getCreateRequest(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, serverEndpointName string, options *ServerEndpointsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/serverEndpoints/{serverEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageSyncServiceName == "" {
		return nil, errors.New("parameter storageSyncServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageSyncServiceName}", url.PathEscape(storageSyncServiceName))
	if syncGroupName == "" {
		return nil, errors.New("parameter syncGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncGroupName}", url.PathEscape(syncGroupName))
	if serverEndpointName == "" {
		return nil, errors.New("parameter serverEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverEndpointName}", url.PathEscape(serverEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServerEndpointsClient) getHandleResponse(resp *http.Response) (ServerEndpointsClientGetResponse, error) {
	result := ServerEndpointsClientGetResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if val := resp.Header.Get("x-ms-correlation-request-id"); val != "" {
		result.XMSCorrelationRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerEndpoint); err != nil {
		return ServerEndpointsClientGetResponse{}, err
	}
	return result, nil
}

// ListBySyncGroup - Get a ServerEndpoint list.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageSyncServiceName - Name of Storage Sync Service resource.
// syncGroupName - Name of Sync Group resource.
// options - ServerEndpointsClientListBySyncGroupOptions contains the optional parameters for the ServerEndpointsClient.ListBySyncGroup
// method.
func (client *ServerEndpointsClient) ListBySyncGroup(resourceGroupName string, storageSyncServiceName string, syncGroupName string, options *ServerEndpointsClientListBySyncGroupOptions) *runtime.Pager[ServerEndpointsClientListBySyncGroupResponse] {
	return runtime.NewPager(runtime.PageProcessor[ServerEndpointsClientListBySyncGroupResponse]{
		More: func(page ServerEndpointsClientListBySyncGroupResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ServerEndpointsClientListBySyncGroupResponse) (ServerEndpointsClientListBySyncGroupResponse, error) {
			req, err := client.listBySyncGroupCreateRequest(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, options)
			if err != nil {
				return ServerEndpointsClientListBySyncGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ServerEndpointsClientListBySyncGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServerEndpointsClientListBySyncGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySyncGroupHandleResponse(resp)
		},
	})
}

// listBySyncGroupCreateRequest creates the ListBySyncGroup request.
func (client *ServerEndpointsClient) listBySyncGroupCreateRequest(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, options *ServerEndpointsClientListBySyncGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/serverEndpoints"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageSyncServiceName == "" {
		return nil, errors.New("parameter storageSyncServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageSyncServiceName}", url.PathEscape(storageSyncServiceName))
	if syncGroupName == "" {
		return nil, errors.New("parameter syncGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncGroupName}", url.PathEscape(syncGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySyncGroupHandleResponse handles the ListBySyncGroup response.
func (client *ServerEndpointsClient) listBySyncGroupHandleResponse(resp *http.Response) (ServerEndpointsClientListBySyncGroupResponse, error) {
	result := ServerEndpointsClientListBySyncGroupResponse{}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if val := resp.Header.Get("x-ms-correlation-request-id"); val != "" {
		result.XMSCorrelationRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerEndpointArray); err != nil {
		return ServerEndpointsClientListBySyncGroupResponse{}, err
	}
	return result, nil
}

// BeginRecallAction - Recall a server endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageSyncServiceName - Name of Storage Sync Service resource.
// syncGroupName - Name of Sync Group resource.
// serverEndpointName - Name of Server Endpoint object.
// parameters - Body of Recall Action object.
// options - ServerEndpointsClientBeginRecallActionOptions contains the optional parameters for the ServerEndpointsClient.BeginRecallAction
// method.
func (client *ServerEndpointsClient) BeginRecallAction(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, serverEndpointName string, parameters RecallActionParameters, options *ServerEndpointsClientBeginRecallActionOptions) (*armruntime.Poller[ServerEndpointsClientRecallActionResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.recallAction(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, serverEndpointName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ServerEndpointsClientRecallActionResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ServerEndpointsClientRecallActionResponse](options.ResumeToken, client.pl, nil)
	}
}

// RecallAction - Recall a server endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ServerEndpointsClient) recallAction(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, serverEndpointName string, parameters RecallActionParameters, options *ServerEndpointsClientBeginRecallActionOptions) (*http.Response, error) {
	req, err := client.recallActionCreateRequest(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, serverEndpointName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// recallActionCreateRequest creates the RecallAction request.
func (client *ServerEndpointsClient) recallActionCreateRequest(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, serverEndpointName string, parameters RecallActionParameters, options *ServerEndpointsClientBeginRecallActionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/serverEndpoints/{serverEndpointName}/recallAction"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageSyncServiceName == "" {
		return nil, errors.New("parameter storageSyncServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageSyncServiceName}", url.PathEscape(storageSyncServiceName))
	if syncGroupName == "" {
		return nil, errors.New("parameter syncGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncGroupName}", url.PathEscape(syncGroupName))
	if serverEndpointName == "" {
		return nil, errors.New("parameter serverEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverEndpointName}", url.PathEscape(serverEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginUpdate - Patch a given ServerEndpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageSyncServiceName - Name of Storage Sync Service resource.
// syncGroupName - Name of Sync Group resource.
// serverEndpointName - Name of Server Endpoint object.
// options - ServerEndpointsClientBeginUpdateOptions contains the optional parameters for the ServerEndpointsClient.BeginUpdate
// method.
func (client *ServerEndpointsClient) BeginUpdate(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, serverEndpointName string, options *ServerEndpointsClientBeginUpdateOptions) (*armruntime.Poller[ServerEndpointsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, serverEndpointName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ServerEndpointsClientUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ServerEndpointsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Patch a given ServerEndpoint.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ServerEndpointsClient) update(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, serverEndpointName string, options *ServerEndpointsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, serverEndpointName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ServerEndpointsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, serverEndpointName string, options *ServerEndpointsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/serverEndpoints/{serverEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageSyncServiceName == "" {
		return nil, errors.New("parameter storageSyncServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageSyncServiceName}", url.PathEscape(storageSyncServiceName))
	if syncGroupName == "" {
		return nil, errors.New("parameter syncGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncGroupName}", url.PathEscape(syncGroupName))
	if serverEndpointName == "" {
		return nil, errors.New("parameter serverEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverEndpointName}", url.PathEscape(serverEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Parameters != nil {
		return req, runtime.MarshalAsJSON(req, *options.Parameters)
	}
	return req, nil
}
