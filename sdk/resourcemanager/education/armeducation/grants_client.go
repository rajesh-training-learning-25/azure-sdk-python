//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armeducation

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

// GrantsClient contains the methods for the Grants group.
// Don't use this type directly, use NewGrantsClient() instead.
type GrantsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewGrantsClient creates a new instance of GrantsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGrantsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*GrantsClient, error) {
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
	client := &GrantsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// Get - Get details for a specific grant linked to the provided billing account and billing profile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01-preview
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - billingProfileName - The ID that uniquely identifies a billing profile.
//   - options - GrantsClientGetOptions contains the optional parameters for the GrantsClient.Get method.
func (client *GrantsClient) Get(ctx context.Context, billingAccountName string, billingProfileName string, options *GrantsClientGetOptions) (GrantsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, billingAccountName, billingProfileName, options)
	if err != nil {
		return GrantsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GrantsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GrantsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *GrantsClient) getCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, options *GrantsClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/providers/Microsoft.Education/grants/default"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
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
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GrantsClient) getHandleResponse(resp *http.Response) (GrantsClientGetResponse, error) {
	result := GrantsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GrantDetails); err != nil {
		return GrantsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get details for a specific grant linked to the provided billing account and billing profile.
//
// Generated from API version 2021-12-01-preview
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - billingProfileName - The ID that uniquely identifies a billing profile.
//   - options - GrantsClientListOptions contains the optional parameters for the GrantsClient.NewListPager method.
func (client *GrantsClient) NewListPager(billingAccountName string, billingProfileName string, options *GrantsClientListOptions) *runtime.Pager[GrantsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[GrantsClientListResponse]{
		More: func(page GrantsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GrantsClientListResponse) (GrantsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, billingAccountName, billingProfileName, options)
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
func (client *GrantsClient) listCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, options *GrantsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/providers/Microsoft.Education/grants"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
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
	req.Raw().Header["Accept"] = []string{"application/json"}
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

// NewListAllPager - Get a list of grants that Microsoft has provided.
//
// Generated from API version 2021-12-01-preview
//   - options - GrantsClientListAllOptions contains the optional parameters for the GrantsClient.NewListAllPager method.
func (client *GrantsClient) NewListAllPager(options *GrantsClientListAllOptions) *runtime.Pager[GrantsClientListAllResponse] {
	return runtime.NewPager(runtime.PagingHandler[GrantsClientListAllResponse]{
		More: func(page GrantsClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GrantsClientListAllResponse) (GrantsClientListAllResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAllCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return GrantsClientListAllResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return GrantsClientListAllResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return GrantsClientListAllResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAllHandleResponse(resp)
		},
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *GrantsClient) listAllCreateRequest(ctx context.Context, options *GrantsClientListAllOptions) (*policy.Request, error) {
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
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *GrantsClient) listAllHandleResponse(resp *http.Response) (GrantsClientListAllResponse, error) {
	result := GrantsClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GrantListResponse); err != nil {
		return GrantsClientListAllResponse{}, err
	}
	return result, nil
}
