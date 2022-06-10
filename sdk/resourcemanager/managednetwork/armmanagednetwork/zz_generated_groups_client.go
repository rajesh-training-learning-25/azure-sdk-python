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

// GroupsClient contains the methods for the ManagedNetworkGroups group.
// Don't use this type directly, use NewGroupsClient() instead.
type GroupsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewGroupsClient creates a new instance of GroupsClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GroupsClient, error) {
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
	client := &GroupsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - The Put ManagedNetworkGroups operation creates or updates a Managed Network Group resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
// resourceGroupName - The name of the resource group.
// managedNetworkName - The name of the Managed Network.
// managedNetworkGroupName - The name of the Managed Network Group.
// managedNetworkGroup - Parameters supplied to the create/update a Managed Network Group resource
// options - GroupsClientBeginCreateOrUpdateOptions contains the optional parameters for the GroupsClient.BeginCreateOrUpdate
// method.
func (client *GroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkGroupName string, managedNetworkGroup Group, options *GroupsClientBeginCreateOrUpdateOptions) (*runtime.Poller[GroupsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, managedNetworkName, managedNetworkGroupName, managedNetworkGroup, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[GroupsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[GroupsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - The Put ManagedNetworkGroups operation creates or updates a Managed Network Group resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
func (client *GroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkGroupName string, managedNetworkGroup Group, options *GroupsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, managedNetworkName, managedNetworkGroupName, managedNetworkGroup, options)
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
func (client *GroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkGroupName string, managedNetworkGroup Group, options *GroupsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}/managedNetworkGroups/{managedNetworkGroupName}"
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
	if managedNetworkGroupName == "" {
		return nil, errors.New("parameter managedNetworkGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedNetworkGroupName}", url.PathEscape(managedNetworkGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, managedNetworkGroup)
}

// BeginDelete - The Delete ManagedNetworkGroups operation deletes a Managed Network Group specified by the resource group,
// Managed Network name, and group name
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
// resourceGroupName - The name of the resource group.
// managedNetworkName - The name of the Managed Network.
// managedNetworkGroupName - The name of the Managed Network Group.
// options - GroupsClientBeginDeleteOptions contains the optional parameters for the GroupsClient.BeginDelete method.
func (client *GroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkGroupName string, options *GroupsClientBeginDeleteOptions) (*runtime.Poller[GroupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, managedNetworkName, managedNetworkGroupName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[GroupsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[GroupsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - The Delete ManagedNetworkGroups operation deletes a Managed Network Group specified by the resource group, Managed
// Network name, and group name
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
func (client *GroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkGroupName string, options *GroupsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, managedNetworkName, managedNetworkGroupName, options)
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
func (client *GroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkGroupName string, options *GroupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}/managedNetworkGroups/{managedNetworkGroupName}"
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
	if managedNetworkGroupName == "" {
		return nil, errors.New("parameter managedNetworkGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedNetworkGroupName}", url.PathEscape(managedNetworkGroupName))
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

// Get - The Get ManagedNetworkGroups operation gets a Managed Network Group specified by the resource group, Managed Network
// name, and group name
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
// resourceGroupName - The name of the resource group.
// managedNetworkName - The name of the Managed Network.
// managedNetworkGroupName - The name of the Managed Network Group.
// options - GroupsClientGetOptions contains the optional parameters for the GroupsClient.Get method.
func (client *GroupsClient) Get(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkGroupName string, options *GroupsClientGetOptions) (GroupsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedNetworkName, managedNetworkGroupName, options)
	if err != nil {
		return GroupsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GroupsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *GroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkGroupName string, options *GroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}/managedNetworkGroups/{managedNetworkGroupName}"
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
	if managedNetworkGroupName == "" {
		return nil, errors.New("parameter managedNetworkGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedNetworkGroupName}", url.PathEscape(managedNetworkGroupName))
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
func (client *GroupsClient) getHandleResponse(resp *http.Response) (GroupsClientGetResponse, error) {
	result := GroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Group); err != nil {
		return GroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByManagedNetworkPager - The ListByManagedNetwork ManagedNetworkGroup operation retrieves all the Managed Network
// Groups in a specified Managed Networks in a paginated format.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
// resourceGroupName - The name of the resource group.
// managedNetworkName - The name of the Managed Network.
// options - GroupsClientListByManagedNetworkOptions contains the optional parameters for the GroupsClient.ListByManagedNetwork
// method.
func (client *GroupsClient) NewListByManagedNetworkPager(resourceGroupName string, managedNetworkName string, options *GroupsClientListByManagedNetworkOptions) *runtime.Pager[GroupsClientListByManagedNetworkResponse] {
	return runtime.NewPager(runtime.PagingHandler[GroupsClientListByManagedNetworkResponse]{
		More: func(page GroupsClientListByManagedNetworkResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GroupsClientListByManagedNetworkResponse) (GroupsClientListByManagedNetworkResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByManagedNetworkCreateRequest(ctx, resourceGroupName, managedNetworkName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return GroupsClientListByManagedNetworkResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return GroupsClientListByManagedNetworkResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return GroupsClientListByManagedNetworkResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByManagedNetworkHandleResponse(resp)
		},
	})
}

// listByManagedNetworkCreateRequest creates the ListByManagedNetwork request.
func (client *GroupsClient) listByManagedNetworkCreateRequest(ctx context.Context, resourceGroupName string, managedNetworkName string, options *GroupsClientListByManagedNetworkOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}/managedNetworkGroups"
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

// listByManagedNetworkHandleResponse handles the ListByManagedNetwork response.
func (client *GroupsClient) listByManagedNetworkHandleResponse(resp *http.Response) (GroupsClientListByManagedNetworkResponse, error) {
	result := GroupsClientListByManagedNetworkResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GroupListResult); err != nil {
		return GroupsClientListByManagedNetworkResponse{}, err
	}
	return result, nil
}
