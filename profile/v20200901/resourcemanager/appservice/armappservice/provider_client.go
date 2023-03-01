//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappservice

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/profile/v20200901/internal"
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

// ProviderClient contains the methods for the Provider group.
// Don't use this type directly, use NewProviderClient() instead.
type ProviderClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewProviderClient creates a new instance of ProviderClient with the specified values.
// subscriptionID - Your Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewProviderClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ProviderClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(internal.ModuleName, internal.ModuleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ProviderClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewGetAvailableStacksPager - Get available application frameworks and their versions
// Generated from API version 2018-02-01
// options - ProviderClientGetAvailableStacksOptions contains the optional parameters for the ProviderClient.GetAvailableStacks
// method.
func (client *ProviderClient) NewGetAvailableStacksPager(options *ProviderClientGetAvailableStacksOptions) *runtime.Pager[ProviderClientGetAvailableStacksResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProviderClientGetAvailableStacksResponse]{
		More: func(page ProviderClientGetAvailableStacksResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProviderClientGetAvailableStacksResponse) (ProviderClientGetAvailableStacksResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.getAvailableStacksCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ProviderClientGetAvailableStacksResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ProviderClientGetAvailableStacksResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ProviderClientGetAvailableStacksResponse{}, runtime.NewResponseError(resp)
			}
			return client.getAvailableStacksHandleResponse(resp)
		},
	})
}

// getAvailableStacksCreateRequest creates the GetAvailableStacks request.
func (client *ProviderClient) getAvailableStacksCreateRequest(ctx context.Context, options *ProviderClientGetAvailableStacksOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Web/availableStacks"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.OSTypeSelected != nil {
		reqQP.Set("osTypeSelected", string(*options.OSTypeSelected))
	}
	reqQP.Set("api-version", "2018-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAvailableStacksHandleResponse handles the GetAvailableStacks response.
func (client *ProviderClient) getAvailableStacksHandleResponse(resp *http.Response) (ProviderClientGetAvailableStacksResponse, error) {
	result := ProviderClientGetAvailableStacksResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationStackCollection); err != nil {
		return ProviderClientGetAvailableStacksResponse{}, err
	}
	return result, nil
}

// NewGetAvailableStacksOnPremPager - Get available application frameworks and their versions
// Generated from API version 2018-02-01
// options - ProviderClientGetAvailableStacksOnPremOptions contains the optional parameters for the ProviderClient.GetAvailableStacksOnPrem
// method.
func (client *ProviderClient) NewGetAvailableStacksOnPremPager(options *ProviderClientGetAvailableStacksOnPremOptions) *runtime.Pager[ProviderClientGetAvailableStacksOnPremResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProviderClientGetAvailableStacksOnPremResponse]{
		More: func(page ProviderClientGetAvailableStacksOnPremResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProviderClientGetAvailableStacksOnPremResponse) (ProviderClientGetAvailableStacksOnPremResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.getAvailableStacksOnPremCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ProviderClientGetAvailableStacksOnPremResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ProviderClientGetAvailableStacksOnPremResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ProviderClientGetAvailableStacksOnPremResponse{}, runtime.NewResponseError(resp)
			}
			return client.getAvailableStacksOnPremHandleResponse(resp)
		},
	})
}

// getAvailableStacksOnPremCreateRequest creates the GetAvailableStacksOnPrem request.
func (client *ProviderClient) getAvailableStacksOnPremCreateRequest(ctx context.Context, options *ProviderClientGetAvailableStacksOnPremOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Web/availableStacks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.OSTypeSelected != nil {
		reqQP.Set("osTypeSelected", string(*options.OSTypeSelected))
	}
	reqQP.Set("api-version", "2018-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAvailableStacksOnPremHandleResponse handles the GetAvailableStacksOnPrem response.
func (client *ProviderClient) getAvailableStacksOnPremHandleResponse(resp *http.Response) (ProviderClientGetAvailableStacksOnPremResponse, error) {
	result := ProviderClientGetAvailableStacksOnPremResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationStackCollection); err != nil {
		return ProviderClientGetAvailableStacksOnPremResponse{}, err
	}
	return result, nil
}

// NewListOperationsPager - Gets all available operations for the Microsoft.Web resource provider. Also exposes resource metric
// definitions
// Generated from API version 2018-02-01
// options - ProviderClientListOperationsOptions contains the optional parameters for the ProviderClient.ListOperations method.
func (client *ProviderClient) NewListOperationsPager(options *ProviderClientListOperationsOptions) *runtime.Pager[ProviderClientListOperationsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProviderClientListOperationsResponse]{
		More: func(page ProviderClientListOperationsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProviderClientListOperationsResponse) (ProviderClientListOperationsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listOperationsCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ProviderClientListOperationsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ProviderClientListOperationsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ProviderClientListOperationsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listOperationsHandleResponse(resp)
		},
	})
}

// listOperationsCreateRequest creates the ListOperations request.
func (client *ProviderClient) listOperationsCreateRequest(ctx context.Context, options *ProviderClientListOperationsOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Web/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listOperationsHandleResponse handles the ListOperations response.
func (client *ProviderClient) listOperationsHandleResponse(resp *http.Response) (ProviderClientListOperationsResponse, error) {
	result := ProviderClientListOperationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CsmOperationCollection); err != nil {
		return ProviderClientListOperationsResponse{}, err
	}
	return result, nil
}
