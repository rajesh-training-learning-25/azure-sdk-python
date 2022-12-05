//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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

// APIRevisionClient contains the methods for the APIRevision group.
// Don't use this type directly, use NewAPIRevisionClient() instead.
type APIRevisionClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAPIRevisionClient creates a new instance of APIRevisionClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAPIRevisionClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*APIRevisionClient, error) {
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
	client := &APIRevisionClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListByServicePager - Lists all revisions of an API.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-08-01
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// apiID - API identifier. Must be unique in the current API Management service instance.
// options - APIRevisionClientListByServiceOptions contains the optional parameters for the APIRevisionClient.ListByService
// method.
func (client *APIRevisionClient) NewListByServicePager(resourceGroupName string, serviceName string, apiID string, options *APIRevisionClientListByServiceOptions) *runtime.Pager[APIRevisionClientListByServiceResponse] {
	return runtime.NewPager(runtime.PagingHandler[APIRevisionClientListByServiceResponse]{
		More: func(page APIRevisionClientListByServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *APIRevisionClientListByServiceResponse) (APIRevisionClientListByServiceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, apiID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return APIRevisionClientListByServiceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return APIRevisionClientListByServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return APIRevisionClientListByServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServiceHandleResponse(resp)
		},
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *APIRevisionClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, options *APIRevisionClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/revisions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *APIRevisionClient) listByServiceHandleResponse(resp *http.Response) (APIRevisionClientListByServiceResponse, error) {
	result := APIRevisionClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIRevisionCollection); err != nil {
		return APIRevisionClientListByServiceResponse{}, err
	}
	return result, nil
}