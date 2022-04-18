//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// MetricDefinitionsClient contains the methods for the MetricDefinitions group.
// Don't use this type directly, use NewMetricDefinitionsClient() instead.
type MetricDefinitionsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewMetricDefinitionsClient creates a new instance of MetricDefinitionsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewMetricDefinitionsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*MetricDefinitionsClient, error) {
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
	client := &MetricDefinitionsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// NewListPager - Lists the metric definitions for the resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceURI - The identifier of the resource.
// options - MetricDefinitionsClientListOptions contains the optional parameters for the MetricDefinitionsClient.List method.
func (client *MetricDefinitionsClient) NewListPager(resourceURI string, options *MetricDefinitionsClientListOptions) *runtime.Pager[MetricDefinitionsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[MetricDefinitionsClientListResponse]{
		More: func(page MetricDefinitionsClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *MetricDefinitionsClientListResponse) (MetricDefinitionsClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceURI, options)
			if err != nil {
				return MetricDefinitionsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return MetricDefinitionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return MetricDefinitionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *MetricDefinitionsClient) listCreateRequest(ctx context.Context, resourceURI string, options *MetricDefinitionsClientListOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/metricDefinitions"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-01-01")
	if options != nil && options.Metricnamespace != nil {
		reqQP.Set("metricnamespace", *options.Metricnamespace)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MetricDefinitionsClient) listHandleResponse(resp *http.Response) (MetricDefinitionsClientListResponse, error) {
	result := MetricDefinitionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricDefinitionCollection); err != nil {
		return MetricDefinitionsClientListResponse{}, err
	}
	return result, nil
}
