//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// MetricsClient contains the methods for the Metrics group.
// Don't use this type directly, use NewMetricsClient() instead.
type MetricsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewMetricsClient creates a new instance of MetricsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMetricsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MetricsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &MetricsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// List - Lists the metric values for a resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-05-01
//   - resourceURI - The identifier of the resource.
//   - options - MetricsClientListOptions contains the optional parameters for the MetricsClient.List method.
func (client *MetricsClient) List(ctx context.Context, resourceURI string, options *MetricsClientListOptions) (MetricsClientListResponse, error) {
	var err error
	const operationName = "MetricsClient.List"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listCreateRequest(ctx, resourceURI, options)
	if err != nil {
		return MetricsClientListResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MetricsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MetricsClientListResponse{}, err
	}
	resp, err := client.listHandleResponse(httpResp)
	return resp, err
}

// listCreateRequest creates the List request.
func (client *MetricsClient) listCreateRequest(ctx context.Context, resourceURI string, options *MetricsClientListOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/metrics"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Timespan != nil {
		reqQP.Set("timespan", *options.Timespan)
	}
	if options != nil && options.Interval != nil {
		reqQP.Set("interval", *options.Interval)
	}
	if options != nil && options.Metricnames != nil {
		reqQP.Set("metricnames", *options.Metricnames)
	}
	if options != nil && options.Aggregation != nil {
		reqQP.Set("aggregation", *options.Aggregation)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("orderby", *options.Orderby)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.ResultType != nil {
		reqQP.Set("resultType", string(*options.ResultType))
	}
	reqQP.Set("api-version", "2021-05-01")
	if options != nil && options.Metricnamespace != nil {
		reqQP.Set("metricnamespace", *options.Metricnamespace)
	}
	if options != nil && options.AutoAdjustTimegrain != nil {
		reqQP.Set("AutoAdjustTimegrain", strconv.FormatBool(*options.AutoAdjustTimegrain))
	}
	if options != nil && options.ValidateDimensions != nil {
		reqQP.Set("ValidateDimensions", strconv.FormatBool(*options.ValidateDimensions))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MetricsClient) listHandleResponse(resp *http.Response) (MetricsClientListResponse, error) {
	result := MetricsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Response); err != nil {
		return MetricsClientListResponse{}, err
	}
	return result, nil
}

// ListAtSubscriptionScope - Lists the metric data for a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-05-01
//   - region - The region where the metrics you want reside.
//   - options - MetricsClientListAtSubscriptionScopeOptions contains the optional parameters for the MetricsClient.ListAtSubscriptionScope
//     method.
func (client *MetricsClient) ListAtSubscriptionScope(ctx context.Context, region string, options *MetricsClientListAtSubscriptionScopeOptions) (MetricsClientListAtSubscriptionScopeResponse, error) {
	var err error
	const operationName = "MetricsClient.ListAtSubscriptionScope"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listAtSubscriptionScopeCreateRequest(ctx, region, options)
	if err != nil {
		return MetricsClientListAtSubscriptionScopeResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MetricsClientListAtSubscriptionScopeResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MetricsClientListAtSubscriptionScopeResponse{}, err
	}
	resp, err := client.listAtSubscriptionScopeHandleResponse(httpResp)
	return resp, err
}

// listAtSubscriptionScopeCreateRequest creates the ListAtSubscriptionScope request.
func (client *MetricsClient) listAtSubscriptionScopeCreateRequest(ctx context.Context, region string, options *MetricsClientListAtSubscriptionScopeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/metrics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	reqQP.Set("region", region)
	if options != nil && options.Timespan != nil {
		reqQP.Set("timespan", *options.Timespan)
	}
	if options != nil && options.Interval != nil {
		reqQP.Set("interval", *options.Interval)
	}
	if options != nil && options.Metricnames != nil {
		reqQP.Set("metricnames", *options.Metricnames)
	}
	if options != nil && options.Aggregation != nil {
		reqQP.Set("aggregation", *options.Aggregation)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("orderby", *options.Orderby)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.ResultType != nil {
		reqQP.Set("resultType", string(*options.ResultType))
	}
	if options != nil && options.Metricnamespace != nil {
		reqQP.Set("metricnamespace", *options.Metricnamespace)
	}
	if options != nil && options.AutoAdjustTimegrain != nil {
		reqQP.Set("AutoAdjustTimegrain", strconv.FormatBool(*options.AutoAdjustTimegrain))
	}
	if options != nil && options.ValidateDimensions != nil {
		reqQP.Set("ValidateDimensions", strconv.FormatBool(*options.ValidateDimensions))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAtSubscriptionScopeHandleResponse handles the ListAtSubscriptionScope response.
func (client *MetricsClient) listAtSubscriptionScopeHandleResponse(resp *http.Response) (MetricsClientListAtSubscriptionScopeResponse, error) {
	result := MetricsClientListAtSubscriptionScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionScopeMetricResponse); err != nil {
		return MetricsClientListAtSubscriptionScopeResponse{}, err
	}
	return result, nil
}

// ListAtSubscriptionScopePost - Lists the metric data for a subscription. Parameters can be specified on either query params
// or the body.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-05-01
//   - region - The region where the metrics you want reside.
//   - options - MetricsClientListAtSubscriptionScopePostOptions contains the optional parameters for the MetricsClient.ListAtSubscriptionScopePost
//     method.
func (client *MetricsClient) ListAtSubscriptionScopePost(ctx context.Context, region string, options *MetricsClientListAtSubscriptionScopePostOptions) (MetricsClientListAtSubscriptionScopePostResponse, error) {
	var err error
	const operationName = "MetricsClient.ListAtSubscriptionScopePost"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listAtSubscriptionScopePostCreateRequest(ctx, region, options)
	if err != nil {
		return MetricsClientListAtSubscriptionScopePostResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MetricsClientListAtSubscriptionScopePostResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MetricsClientListAtSubscriptionScopePostResponse{}, err
	}
	resp, err := client.listAtSubscriptionScopePostHandleResponse(httpResp)
	return resp, err
}

// listAtSubscriptionScopePostCreateRequest creates the ListAtSubscriptionScopePost request.
func (client *MetricsClient) listAtSubscriptionScopePostCreateRequest(ctx context.Context, region string, options *MetricsClientListAtSubscriptionScopePostOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/metrics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	reqQP.Set("region", region)
	if options != nil && options.Timespan != nil {
		reqQP.Set("timespan", *options.Timespan)
	}
	if options != nil && options.Interval != nil {
		reqQP.Set("interval", *options.Interval)
	}
	if options != nil && options.Metricnames != nil {
		reqQP.Set("metricnames", *options.Metricnames)
	}
	if options != nil && options.Aggregation != nil {
		reqQP.Set("aggregation", *options.Aggregation)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("orderby", *options.Orderby)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.ResultType != nil {
		reqQP.Set("resultType", string(*options.ResultType))
	}
	if options != nil && options.Metricnamespace != nil {
		reqQP.Set("metricnamespace", *options.Metricnamespace)
	}
	if options != nil && options.AutoAdjustTimegrain != nil {
		reqQP.Set("AutoAdjustTimegrain", strconv.FormatBool(*options.AutoAdjustTimegrain))
	}
	if options != nil && options.ValidateDimensions != nil {
		reqQP.Set("ValidateDimensions", strconv.FormatBool(*options.ValidateDimensions))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		if err := runtime.MarshalAsJSON(req, *options.Body); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// listAtSubscriptionScopePostHandleResponse handles the ListAtSubscriptionScopePost response.
func (client *MetricsClient) listAtSubscriptionScopePostHandleResponse(resp *http.Response) (MetricsClientListAtSubscriptionScopePostResponse, error) {
	result := MetricsClientListAtSubscriptionScopePostResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionScopeMetricResponse); err != nil {
		return MetricsClientListAtSubscriptionScopePostResponse{}, err
	}
	return result, nil
}
