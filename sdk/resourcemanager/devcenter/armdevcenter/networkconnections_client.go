//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdevcenter

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
	"strconv"
	"strings"
)

// NetworkConnectionsClient contains the methods for the NetworkConnections group.
// Don't use this type directly, use NewNetworkConnectionsClient() instead.
type NetworkConnectionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewNetworkConnectionsClient creates a new instance of NetworkConnectionsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewNetworkConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NetworkConnectionsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &NetworkConnectionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a Network Connections resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-11-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// networkConnectionName - Name of the Network Connection that can be applied to a Pool.
// body - Represents network connection
// options - NetworkConnectionsClientBeginCreateOrUpdateOptions contains the optional parameters for the NetworkConnectionsClient.BeginCreateOrUpdate
// method.
func (client *NetworkConnectionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkConnectionName string, body NetworkConnection, options *NetworkConnectionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[NetworkConnectionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, networkConnectionName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[NetworkConnectionsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[NetworkConnectionsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a Network Connections resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-11-preview
func (client *NetworkConnectionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, networkConnectionName string, body NetworkConnection, options *NetworkConnectionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, networkConnectionName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *NetworkConnectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkConnectionName string, body NetworkConnection, options *NetworkConnectionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/networkConnections/{networkConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkConnectionName == "" {
		return nil, errors.New("parameter networkConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkConnectionName}", url.PathEscape(networkConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-11-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Deletes a Network Connections resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-11-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// networkConnectionName - Name of the Network Connection that can be applied to a Pool.
// options - NetworkConnectionsClientBeginDeleteOptions contains the optional parameters for the NetworkConnectionsClient.BeginDelete
// method.
func (client *NetworkConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, networkConnectionName string, options *NetworkConnectionsClientBeginDeleteOptions) (*runtime.Poller[NetworkConnectionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, networkConnectionName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[NetworkConnectionsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[NetworkConnectionsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a Network Connections resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-11-preview
func (client *NetworkConnectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, networkConnectionName string, options *NetworkConnectionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkConnectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *NetworkConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkConnectionName string, options *NetworkConnectionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/networkConnections/{networkConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkConnectionName == "" {
		return nil, errors.New("parameter networkConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkConnectionName}", url.PathEscape(networkConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-11-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a network connection resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-11-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// networkConnectionName - Name of the Network Connection that can be applied to a Pool.
// options - NetworkConnectionsClientGetOptions contains the optional parameters for the NetworkConnectionsClient.Get method.
func (client *NetworkConnectionsClient) Get(ctx context.Context, resourceGroupName string, networkConnectionName string, options *NetworkConnectionsClientGetOptions) (NetworkConnectionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkConnectionName, options)
	if err != nil {
		return NetworkConnectionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NetworkConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NetworkConnectionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *NetworkConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkConnectionName string, options *NetworkConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/networkConnections/{networkConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkConnectionName == "" {
		return nil, errors.New("parameter networkConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkConnectionName}", url.PathEscape(networkConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-11-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NetworkConnectionsClient) getHandleResponse(resp *http.Response) (NetworkConnectionsClientGetResponse, error) {
	result := NetworkConnectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkConnection); err != nil {
		return NetworkConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// GetHealthDetails - Gets health check status details.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-11-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// networkConnectionName - Name of the Network Connection that can be applied to a Pool.
// options - NetworkConnectionsClientGetHealthDetailsOptions contains the optional parameters for the NetworkConnectionsClient.GetHealthDetails
// method.
func (client *NetworkConnectionsClient) GetHealthDetails(ctx context.Context, resourceGroupName string, networkConnectionName string, options *NetworkConnectionsClientGetHealthDetailsOptions) (NetworkConnectionsClientGetHealthDetailsResponse, error) {
	req, err := client.getHealthDetailsCreateRequest(ctx, resourceGroupName, networkConnectionName, options)
	if err != nil {
		return NetworkConnectionsClientGetHealthDetailsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NetworkConnectionsClientGetHealthDetailsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NetworkConnectionsClientGetHealthDetailsResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHealthDetailsHandleResponse(resp)
}

// getHealthDetailsCreateRequest creates the GetHealthDetails request.
func (client *NetworkConnectionsClient) getHealthDetailsCreateRequest(ctx context.Context, resourceGroupName string, networkConnectionName string, options *NetworkConnectionsClientGetHealthDetailsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/networkConnections/{networkConnectionName}/healthChecks/latest"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkConnectionName == "" {
		return nil, errors.New("parameter networkConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkConnectionName}", url.PathEscape(networkConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-11-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHealthDetailsHandleResponse handles the GetHealthDetails response.
func (client *NetworkConnectionsClient) getHealthDetailsHandleResponse(resp *http.Response) (NetworkConnectionsClientGetHealthDetailsResponse, error) {
	result := NetworkConnectionsClientGetHealthDetailsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HealthCheckStatusDetails); err != nil {
		return NetworkConnectionsClientGetHealthDetailsResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists network connections in a resource group
// Generated from API version 2022-11-11-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - NetworkConnectionsClientListByResourceGroupOptions contains the optional parameters for the NetworkConnectionsClient.ListByResourceGroup
// method.
func (client *NetworkConnectionsClient) NewListByResourceGroupPager(resourceGroupName string, options *NetworkConnectionsClientListByResourceGroupOptions) *runtime.Pager[NetworkConnectionsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[NetworkConnectionsClientListByResourceGroupResponse]{
		More: func(page NetworkConnectionsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NetworkConnectionsClientListByResourceGroupResponse) (NetworkConnectionsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return NetworkConnectionsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return NetworkConnectionsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return NetworkConnectionsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *NetworkConnectionsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *NetworkConnectionsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/networkConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-11-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *NetworkConnectionsClient) listByResourceGroupHandleResponse(resp *http.Response) (NetworkConnectionsClientListByResourceGroupResponse, error) {
	result := NetworkConnectionsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkConnectionListResult); err != nil {
		return NetworkConnectionsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists network connections in a subscription
// Generated from API version 2022-11-11-preview
// options - NetworkConnectionsClientListBySubscriptionOptions contains the optional parameters for the NetworkConnectionsClient.ListBySubscription
// method.
func (client *NetworkConnectionsClient) NewListBySubscriptionPager(options *NetworkConnectionsClientListBySubscriptionOptions) *runtime.Pager[NetworkConnectionsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[NetworkConnectionsClientListBySubscriptionResponse]{
		More: func(page NetworkConnectionsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NetworkConnectionsClientListBySubscriptionResponse) (NetworkConnectionsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return NetworkConnectionsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return NetworkConnectionsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return NetworkConnectionsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *NetworkConnectionsClient) listBySubscriptionCreateRequest(ctx context.Context, options *NetworkConnectionsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DevCenter/networkConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-11-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *NetworkConnectionsClient) listBySubscriptionHandleResponse(resp *http.Response) (NetworkConnectionsClientListBySubscriptionResponse, error) {
	result := NetworkConnectionsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkConnectionListResult); err != nil {
		return NetworkConnectionsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// NewListHealthDetailsPager - Lists health check status details
// Generated from API version 2022-11-11-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// networkConnectionName - Name of the Network Connection that can be applied to a Pool.
// options - NetworkConnectionsClientListHealthDetailsOptions contains the optional parameters for the NetworkConnectionsClient.ListHealthDetails
// method.
func (client *NetworkConnectionsClient) NewListHealthDetailsPager(resourceGroupName string, networkConnectionName string, options *NetworkConnectionsClientListHealthDetailsOptions) *runtime.Pager[NetworkConnectionsClientListHealthDetailsResponse] {
	return runtime.NewPager(runtime.PagingHandler[NetworkConnectionsClientListHealthDetailsResponse]{
		More: func(page NetworkConnectionsClientListHealthDetailsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NetworkConnectionsClientListHealthDetailsResponse) (NetworkConnectionsClientListHealthDetailsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listHealthDetailsCreateRequest(ctx, resourceGroupName, networkConnectionName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return NetworkConnectionsClientListHealthDetailsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return NetworkConnectionsClientListHealthDetailsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return NetworkConnectionsClientListHealthDetailsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHealthDetailsHandleResponse(resp)
		},
	})
}

// listHealthDetailsCreateRequest creates the ListHealthDetails request.
func (client *NetworkConnectionsClient) listHealthDetailsCreateRequest(ctx context.Context, resourceGroupName string, networkConnectionName string, options *NetworkConnectionsClientListHealthDetailsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/networkConnections/{networkConnectionName}/healthChecks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkConnectionName == "" {
		return nil, errors.New("parameter networkConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkConnectionName}", url.PathEscape(networkConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-11-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHealthDetailsHandleResponse handles the ListHealthDetails response.
func (client *NetworkConnectionsClient) listHealthDetailsHandleResponse(resp *http.Response) (NetworkConnectionsClientListHealthDetailsResponse, error) {
	result := NetworkConnectionsClientListHealthDetailsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HealthCheckStatusDetailsListResult); err != nil {
		return NetworkConnectionsClientListHealthDetailsResponse{}, err
	}
	return result, nil
}

// BeginRunHealthChecks - Triggers a new health check run. The execution and health check result can be tracked via the network
// Connection health check details
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-11-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// networkConnectionName - Name of the Network Connection that can be applied to a Pool.
// options - NetworkConnectionsClientBeginRunHealthChecksOptions contains the optional parameters for the NetworkConnectionsClient.BeginRunHealthChecks
// method.
func (client *NetworkConnectionsClient) BeginRunHealthChecks(ctx context.Context, resourceGroupName string, networkConnectionName string, options *NetworkConnectionsClientBeginRunHealthChecksOptions) (*runtime.Poller[NetworkConnectionsClientRunHealthChecksResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.runHealthChecks(ctx, resourceGroupName, networkConnectionName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[NetworkConnectionsClientRunHealthChecksResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[NetworkConnectionsClientRunHealthChecksResponse](options.ResumeToken, client.pl, nil)
	}
}

// RunHealthChecks - Triggers a new health check run. The execution and health check result can be tracked via the network
// Connection health check details
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-11-preview
func (client *NetworkConnectionsClient) runHealthChecks(ctx context.Context, resourceGroupName string, networkConnectionName string, options *NetworkConnectionsClientBeginRunHealthChecksOptions) (*http.Response, error) {
	req, err := client.runHealthChecksCreateRequest(ctx, resourceGroupName, networkConnectionName, options)
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

// runHealthChecksCreateRequest creates the RunHealthChecks request.
func (client *NetworkConnectionsClient) runHealthChecksCreateRequest(ctx context.Context, resourceGroupName string, networkConnectionName string, options *NetworkConnectionsClientBeginRunHealthChecksOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/networkConnections/{networkConnectionName}/runHealthChecks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkConnectionName == "" {
		return nil, errors.New("parameter networkConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkConnectionName}", url.PathEscape(networkConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-11-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginUpdate - Partially updates a Network Connection
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-11-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// networkConnectionName - Name of the Network Connection that can be applied to a Pool.
// body - Represents network connection
// options - NetworkConnectionsClientBeginUpdateOptions contains the optional parameters for the NetworkConnectionsClient.BeginUpdate
// method.
func (client *NetworkConnectionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, networkConnectionName string, body NetworkConnectionUpdate, options *NetworkConnectionsClientBeginUpdateOptions) (*runtime.Poller[NetworkConnectionsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, networkConnectionName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[NetworkConnectionsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[NetworkConnectionsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Partially updates a Network Connection
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-11-preview
func (client *NetworkConnectionsClient) update(ctx context.Context, resourceGroupName string, networkConnectionName string, body NetworkConnectionUpdate, options *NetworkConnectionsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, networkConnectionName, body, options)
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
func (client *NetworkConnectionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, networkConnectionName string, body NetworkConnectionUpdate, options *NetworkConnectionsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/networkConnections/{networkConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkConnectionName == "" {
		return nil, errors.New("parameter networkConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkConnectionName}", url.PathEscape(networkConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-11-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}
