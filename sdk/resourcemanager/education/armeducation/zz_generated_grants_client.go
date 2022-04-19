//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeducation

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strconv"
)

// GrantsClient contains the methods for the Grants group.
// Don't use this type directly, use NewGrantsClient() instead.
type GrantsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewGrantsClient creates a new instance of GrantsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewGrantsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*GrantsClient, error) {
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
	client := &GrantsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// NewListPager - Get a list of grants that Microsoft has provided.
// If the operation fails it returns an *azcore.ResponseError type.
// options - GrantsClientListOptions contains the optional parameters for the GrantsClient.List method.
func (client *GrantsClient) NewListPager(options *GrantsClientListOptions) *runtime.Pager[GrantsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[GrantsClientListResponse]{
		More: func(page GrantsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GrantsClientListResponse) (GrantsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return GrantsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return GrantsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return GrantsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *GrantsClient) listCreateRequest(ctx context.Context, options *GrantsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Education/grants"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.IncludeAllocatedBudget != nil {
		reqQP.Set("includeAllocatedBudget", strconv.FormatBool(*options.IncludeAllocatedBudget))
	}
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *GrantsClient) listHandleResponse(resp *http.Response) (GrantsClientListResponse, error) {
	result := GrantsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GrantListResponse); err != nil {
		return GrantsClientListResponse{}, err
	}
	return result, nil
}
