//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagednetwork

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

// ManagedNetworksClient contains the methods for the ManagedNetworks group.
// Don't use this type directly, use NewManagedNetworksClient() instead.
type ManagedNetworksClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewManagedNetworksClient creates a new instance of ManagedNetworksClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewManagedNetworksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedNetworksClient, error) {
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
	client := &ManagedNetworksClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - The Put ManagedNetworks operation creates/updates a Managed Network Resource, specified by resource group
// and Managed Network name
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
// resourceGroupName - The name of the resource group.
// managedNetworkName - The name of the Managed Network.
// managedNetwork - Parameters supplied to the create/update a Managed Network Resource
// options - ManagedNetworksClientCreateOrUpdateOptions contains the optional parameters for the ManagedNetworksClient.CreateOrUpdate
// method.
func (client *ManagedNetworksClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetwork ManagedNetwork, options *ManagedNetworksClientCreateOrUpdateOptions) (ManagedNetworksClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, managedNetworkName, managedNetwork, options)
	if err != nil {
		return ManagedNetworksClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedNetworksClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ManagedNetworksClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ManagedNetworksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetwork ManagedNetwork, options *ManagedNetworksClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedNetworkName == "" {
		return nil, errors.New("parameter managedNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedNetworkName}", url.PathEscape(managedNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, managedNetwork)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ManagedNetworksClient) createOrUpdateHandleResponse(resp *http.Response) (ManagedNetworksClientCreateOrUpdateResponse, error) {
	result := ManagedNetworksClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedNetwork); err != nil {
		return ManagedNetworksClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - The Delete ManagedNetworks operation deletes a Managed Network Resource, specified by the resource group
// and Managed Network name
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
// resourceGroupName - The name of the resource group.
// managedNetworkName - The name of the Managed Network.
// options - ManagedNetworksClientBeginDeleteOptions contains the optional parameters for the ManagedNetworksClient.BeginDelete
// method.
func (client *ManagedNetworksClient) BeginDelete(ctx context.Context, resourceGroupName string, managedNetworkName string, options *ManagedNetworksClientBeginDeleteOptions) (*runtime.Poller[ManagedNetworksClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, managedNetworkName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[ManagedNetworksClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ManagedNetworksClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - The Delete ManagedNetworks operation deletes a Managed Network Resource, specified by the resource group and Managed
// Network name
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
func (client *ManagedNetworksClient) deleteOperation(ctx context.Context, resourceGroupName string, managedNetworkName string, options *ManagedNetworksClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, managedNetworkName, options)
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
func (client *ManagedNetworksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, managedNetworkName string, options *ManagedNetworksClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedNetworkName == "" {
		return nil, errors.New("parameter managedNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedNetworkName}", url.PathEscape(managedNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - The Get ManagedNetworks operation gets a Managed Network Resource, specified by the resource group and Managed Network
// name
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
// resourceGroupName - The name of the resource group.
// managedNetworkName - The name of the Managed Network.
// options - ManagedNetworksClientGetOptions contains the optional parameters for the ManagedNetworksClient.Get method.
func (client *ManagedNetworksClient) Get(ctx context.Context, resourceGroupName string, managedNetworkName string, options *ManagedNetworksClientGetOptions) (ManagedNetworksClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedNetworkName, options)
	if err != nil {
		return ManagedNetworksClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedNetworksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedNetworksClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ManagedNetworksClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedNetworkName string, options *ManagedNetworksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedNetworkName == "" {
		return nil, errors.New("parameter managedNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedNetworkName}", url.PathEscape(managedNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedNetworksClient) getHandleResponse(resp *http.Response) (ManagedNetworksClientGetResponse, error) {
	result := ManagedNetworksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedNetwork); err != nil {
		return ManagedNetworksClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - The ListByResourceGroup ManagedNetwork operation retrieves all the Managed Network resources
// in a resource group in a paginated format.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
// resourceGroupName - The name of the resource group.
// options - ManagedNetworksClientListByResourceGroupOptions contains the optional parameters for the ManagedNetworksClient.ListByResourceGroup
// method.
func (client *ManagedNetworksClient) NewListByResourceGroupPager(resourceGroupName string, options *ManagedNetworksClientListByResourceGroupOptions) *runtime.Pager[ManagedNetworksClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedNetworksClientListByResourceGroupResponse]{
		More: func(page ManagedNetworksClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedNetworksClientListByResourceGroupResponse) (ManagedNetworksClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ManagedNetworksClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ManagedNetworksClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ManagedNetworksClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ManagedNetworksClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ManagedNetworksClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks"
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
	reqQP.Set("api-version", "2019-06-01-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ManagedNetworksClient) listByResourceGroupHandleResponse(resp *http.Response) (ManagedNetworksClientListByResourceGroupResponse, error) {
	result := ManagedNetworksClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return ManagedNetworksClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - The ListBySubscription ManagedNetwork operation retrieves all the Managed Network Resources
// in the current subscription in a paginated format.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
// options - ManagedNetworksClientListBySubscriptionOptions contains the optional parameters for the ManagedNetworksClient.ListBySubscription
// method.
func (client *ManagedNetworksClient) NewListBySubscriptionPager(options *ManagedNetworksClientListBySubscriptionOptions) *runtime.Pager[ManagedNetworksClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedNetworksClientListBySubscriptionResponse]{
		More: func(page ManagedNetworksClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedNetworksClientListBySubscriptionResponse) (ManagedNetworksClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ManagedNetworksClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ManagedNetworksClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ManagedNetworksClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ManagedNetworksClient) listBySubscriptionCreateRequest(ctx context.Context, options *ManagedNetworksClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ManagedNetwork/managedNetworks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ManagedNetworksClient) listBySubscriptionHandleResponse(resp *http.Response) (ManagedNetworksClientListBySubscriptionResponse, error) {
	result := ManagedNetworksClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return ManagedNetworksClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates the specified Managed Network resource tags.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
// resourceGroupName - The name of the resource group.
// managedNetworkName - The name of the Managed Network.
// parameters - Parameters supplied to update application gateway tags and/or scope.
// options - ManagedNetworksClientBeginUpdateOptions contains the optional parameters for the ManagedNetworksClient.BeginUpdate
// method.
func (client *ManagedNetworksClient) BeginUpdate(ctx context.Context, resourceGroupName string, managedNetworkName string, parameters Update, options *ManagedNetworksClientBeginUpdateOptions) (*runtime.Poller[ManagedNetworksClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, managedNetworkName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[ManagedNetworksClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ManagedNetworksClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Updates the specified Managed Network resource tags.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
func (client *ManagedNetworksClient) update(ctx context.Context, resourceGroupName string, managedNetworkName string, parameters Update, options *ManagedNetworksClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, managedNetworkName, parameters, options)
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

// updateCreateRequest creates the Update request.
func (client *ManagedNetworksClient) updateCreateRequest(ctx context.Context, resourceGroupName string, managedNetworkName string, parameters Update, options *ManagedNetworksClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedNetworkName == "" {
		return nil, errors.New("parameter managedNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedNetworkName}", url.PathEscape(managedNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}