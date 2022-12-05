//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicefabricmesh

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

// ServiceReplicaClient contains the methods for the ServiceReplica group.
// Don't use this type directly, use NewServiceReplicaClient() instead.
type ServiceReplicaClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewServiceReplicaClient creates a new instance of ServiceReplicaClient with the specified values.
// subscriptionID - The customer subscription identifier
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewServiceReplicaClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServiceReplicaClient, error) {
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
	client := &ServiceReplicaClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Gets the information about the service replica with the given name. The information include the description and other
// properties of the service replica.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-09-01-preview
// resourceGroupName - Azure resource group name
// applicationResourceName - The identity of the application.
// serviceResourceName - The identity of the service.
// replicaName - Service Fabric replica name.
// options - ServiceReplicaClientGetOptions contains the optional parameters for the ServiceReplicaClient.Get method.
func (client *ServiceReplicaClient) Get(ctx context.Context, resourceGroupName string, applicationResourceName string, serviceResourceName string, replicaName string, options *ServiceReplicaClientGetOptions) (ServiceReplicaClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, applicationResourceName, serviceResourceName, replicaName, options)
	if err != nil {
		return ServiceReplicaClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServiceReplicaClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServiceReplicaClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServiceReplicaClient) getCreateRequest(ctx context.Context, resourceGroupName string, applicationResourceName string, serviceResourceName string, replicaName string, options *ServiceReplicaClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/applications/{applicationResourceName}/services/{serviceResourceName}/replicas/{replicaName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{applicationResourceName}", applicationResourceName)
	urlPath = strings.ReplaceAll(urlPath, "{serviceResourceName}", serviceResourceName)
	urlPath = strings.ReplaceAll(urlPath, "{replicaName}", replicaName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServiceReplicaClient) getHandleResponse(resp *http.Response) (ServiceReplicaClientGetResponse, error) {
	result := ServiceReplicaClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceReplicaDescription); err != nil {
		return ServiceReplicaClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets the information about all replicas of a given service of an application. The information includes the
// runtime properties of the replica instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-09-01-preview
// resourceGroupName - Azure resource group name
// applicationResourceName - The identity of the application.
// serviceResourceName - The identity of the service.
// options - ServiceReplicaClientListOptions contains the optional parameters for the ServiceReplicaClient.List method.
func (client *ServiceReplicaClient) NewListPager(resourceGroupName string, applicationResourceName string, serviceResourceName string, options *ServiceReplicaClientListOptions) *runtime.Pager[ServiceReplicaClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServiceReplicaClientListResponse]{
		More: func(page ServiceReplicaClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServiceReplicaClientListResponse) (ServiceReplicaClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, applicationResourceName, serviceResourceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ServiceReplicaClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ServiceReplicaClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServiceReplicaClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ServiceReplicaClient) listCreateRequest(ctx context.Context, resourceGroupName string, applicationResourceName string, serviceResourceName string, options *ServiceReplicaClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/applications/{applicationResourceName}/services/{serviceResourceName}/replicas"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{applicationResourceName}", applicationResourceName)
	urlPath = strings.ReplaceAll(urlPath, "{serviceResourceName}", serviceResourceName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServiceReplicaClient) listHandleResponse(resp *http.Response) (ServiceReplicaClientListResponse, error) {
	result := ServiceReplicaClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceReplicaDescriptionList); err != nil {
		return ServiceReplicaClientListResponse{}, err
	}
	return result, nil
}