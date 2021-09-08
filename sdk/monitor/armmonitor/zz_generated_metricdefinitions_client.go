//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// MetricDefinitionsClient contains the methods for the MetricDefinitions group.
// Don't use this type directly, use NewMetricDefinitionsClient() instead.
type MetricDefinitionsClient struct {
	ep string
	pl runtime.Pipeline
}

// NewMetricDefinitionsClient creates a new instance of MetricDefinitionsClient with the specified values.
func NewMetricDefinitionsClient(con *arm.Connection) *MetricDefinitionsClient {
	return &MetricDefinitionsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version)}
}

// List - Lists the metric definitions for the resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MetricDefinitionsClient) List(ctx context.Context, resourceURI string, options *MetricDefinitionsListOptions) (MetricDefinitionsListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceURI, options)
	if err != nil {
		return MetricDefinitionsListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MetricDefinitionsListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MetricDefinitionsListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *MetricDefinitionsClient) listCreateRequest(ctx context.Context, resourceURI string, options *MetricDefinitionsListOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/metricDefinitions"
	if resourceURI == "" {
		return nil, errors.New("parameter resourceURI cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
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
func (client *MetricDefinitionsClient) listHandleResponse(resp *http.Response) (MetricDefinitionsListResponse, error) {
	result := MetricDefinitionsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricDefinitionCollection); err != nil {
		return MetricDefinitionsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *MetricDefinitionsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
