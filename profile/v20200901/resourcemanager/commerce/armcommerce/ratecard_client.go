//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcommerce

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/profile/v20200901/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RateCardClient contains the methods for the RateCard group.
// Don't use this type directly, use NewRateCardClient() instead.
type RateCardClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRateCardClient creates a new instance of RateCardClient with the specified values.
//   - subscriptionID - It uniquely identifies Microsoft Azure subscription. The subscription ID forms part of the URI for every
//     service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRateCardClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RateCardClient, error) {
	cl, err := arm.NewClient(internal.ModuleName+"/armcommerce.RateCardClient", internal.ModuleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RateCardClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Enables you to query for the resource/meter metadata and related prices used in a given subscription by Offer ID,
// Currency, Locale and Region. The metadata associated with the billing meters,
// including but not limited to service names, types, resources, units of measure, and regions, is subject to change at any
// time and without notice. If you intend to use this billing data in an automated
// fashion, please use the billing meter GUID to uniquely identify each billable item. If the billing meter GUID is scheduled
// to change due to a new billing model, you will be notified in advance of the
// change.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-06-01-preview
//   - filter - The filter to apply on the operation. It ONLY supports the 'eq' and 'and' logical operators at this time. All
//     the 4 query parameters 'OfferDurableId', 'Currency', 'Locale', 'Region' are required to be
//     a part of the $filter.
//   - options - RateCardClientGetOptions contains the optional parameters for the RateCardClient.Get method.
func (client *RateCardClient) Get(ctx context.Context, filter string, options *RateCardClientGetOptions) (RateCardClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, filter, options)
	if err != nil {
		return RateCardClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RateCardClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RateCardClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RateCardClient) getCreateRequest(ctx context.Context, filter string, options *RateCardClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Commerce/RateCard"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("$filter", filter)
	reqQP.Set("api-version", "2015-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json, text/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RateCardClient) getHandleResponse(resp *http.Response) (RateCardClientGetResponse, error) {
	result := RateCardClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceRateCardInfo); err != nil {
		return RateCardClientGetResponse{}, err
	}
	return result, nil
}
