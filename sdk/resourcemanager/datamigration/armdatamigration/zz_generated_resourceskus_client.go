//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatamigration

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

// ResourceSKUsClient contains the methods for the ResourceSKUs group.
// Don't use this type directly, use NewResourceSKUsClient() instead.
type ResourceSKUsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewResourceSKUsClient creates a new instance of ResourceSKUsClient with the specified values.
// subscriptionID - Identifier of the subscription
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewResourceSKUsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ResourceSKUsClient, error) {
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
	client := &ResourceSKUsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListSKUsPager - The skus action returns the list of SKUs that DMS supports.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-30
// options - ResourceSKUsClientListSKUsOptions contains the optional parameters for the ResourceSKUsClient.ListSKUs method.
func (client *ResourceSKUsClient) NewListSKUsPager(options *ResourceSKUsClientListSKUsOptions) *runtime.Pager[ResourceSKUsClientListSKUsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ResourceSKUsClientListSKUsResponse]{
		More: func(page ResourceSKUsClientListSKUsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ResourceSKUsClientListSKUsResponse) (ResourceSKUsClientListSKUsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listSKUsCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ResourceSKUsClientListSKUsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ResourceSKUsClientListSKUsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ResourceSKUsClientListSKUsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listSKUsHandleResponse(resp)
		},
	})
}

// listSKUsCreateRequest creates the ListSKUs request.
func (client *ResourceSKUsClient) listSKUsCreateRequest(ctx context.Context, options *ResourceSKUsClientListSKUsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DataMigration/skus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSKUsHandleResponse handles the ListSKUs response.
func (client *ResourceSKUsClient) listSKUsHandleResponse(resp *http.Response) (ResourceSKUsClientListSKUsResponse, error) {
	result := ResourceSKUsClientListSKUsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceSKUsResult); err != nil {
		return ResourceSKUsClientListSKUsResponse{}, err
	}
	return result, nil
}