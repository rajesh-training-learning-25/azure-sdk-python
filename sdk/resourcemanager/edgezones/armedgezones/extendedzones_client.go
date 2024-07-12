// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armedgezones

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ExtendedZonesClient - ExtendedZone operations
// Don't use this type directly, use NewExtendedZonesClient() instead.
type ExtendedZonesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewExtendedZonesClient creates a new instance of ExtendedZonesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewExtendedZonesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ExtendedZonesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ExtendedZonesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets an Azure Extended Zone for a subscription
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01-preview
//   - extendedZoneName - The name of the ExtendedZone
//   - options - ExtendedZonesClientGetOptions contains the optional parameters for the ExtendedZonesClient.Get method.
func (client *ExtendedZonesClient) Get(ctx context.Context, extendedZoneName string, options *ExtendedZonesClientGetOptions) (ExtendedZonesClientGetResponse, error) {
	var err error
	const operationName = "ExtendedZonesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, extendedZoneName, options)
	if err != nil {
		return ExtendedZonesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExtendedZonesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ExtendedZonesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ExtendedZonesClient) getCreateRequest(ctx context.Context, extendedZoneName string, _ *ExtendedZonesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.EdgeZones/extendedZones/{extendedZoneName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if extendedZoneName == "" {
		return nil, errors.New("parameter extendedZoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extendedZoneName}", url.PathEscape(extendedZoneName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExtendedZonesClient) getHandleResponse(resp *http.Response) (ExtendedZonesClientGetResponse, error) {
	result := ExtendedZonesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtendedZone); err != nil {
		return ExtendedZonesClientGetResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists the Azure Extended Zones available to a subscription
//
// Generated from API version 2024-04-01-preview
//   - options - ExtendedZonesClientListBySubscriptionOptions contains the optional parameters for the ExtendedZonesClient.NewListBySubscriptionPager
//     method.
func (client *ExtendedZonesClient) NewListBySubscriptionPager(options *ExtendedZonesClientListBySubscriptionOptions) *runtime.Pager[ExtendedZonesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExtendedZonesClientListBySubscriptionResponse]{
		More: func(page ExtendedZonesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExtendedZonesClientListBySubscriptionResponse) (ExtendedZonesClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ExtendedZonesClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ExtendedZonesClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ExtendedZonesClient) listBySubscriptionCreateRequest(ctx context.Context, _ *ExtendedZonesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.EdgeZones/extendedZones"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ExtendedZonesClient) listBySubscriptionHandleResponse(resp *http.Response) (ExtendedZonesClientListBySubscriptionResponse, error) {
	result := ExtendedZonesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtendedZoneListResult); err != nil {
		return ExtendedZonesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Register - Registers a subscription for an Extended Zone
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01-preview
//   - extendedZoneName - The name of the ExtendedZone
//   - options - ExtendedZonesClientRegisterOptions contains the optional parameters for the ExtendedZonesClient.Register method.
func (client *ExtendedZonesClient) Register(ctx context.Context, extendedZoneName string, options *ExtendedZonesClientRegisterOptions) (ExtendedZonesClientRegisterResponse, error) {
	var err error
	const operationName = "ExtendedZonesClient.Register"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.registerCreateRequest(ctx, extendedZoneName, options)
	if err != nil {
		return ExtendedZonesClientRegisterResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExtendedZonesClientRegisterResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ExtendedZonesClientRegisterResponse{}, err
	}
	resp, err := client.registerHandleResponse(httpResp)
	return resp, err
}

// registerCreateRequest creates the Register request.
func (client *ExtendedZonesClient) registerCreateRequest(ctx context.Context, extendedZoneName string, _ *ExtendedZonesClientRegisterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.EdgeZones/extendedZones/{extendedZoneName}/register"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if extendedZoneName == "" {
		return nil, errors.New("parameter extendedZoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extendedZoneName}", url.PathEscape(extendedZoneName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// registerHandleResponse handles the Register response.
func (client *ExtendedZonesClient) registerHandleResponse(resp *http.Response) (ExtendedZonesClientRegisterResponse, error) {
	result := ExtendedZonesClientRegisterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtendedZone); err != nil {
		return ExtendedZonesClientRegisterResponse{}, err
	}
	return result, nil
}

// Unregister - Unregisters a subscription for an Extended Zone
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01-preview
//   - extendedZoneName - The name of the ExtendedZone
//   - options - ExtendedZonesClientUnregisterOptions contains the optional parameters for the ExtendedZonesClient.Unregister
//     method.
func (client *ExtendedZonesClient) Unregister(ctx context.Context, extendedZoneName string, options *ExtendedZonesClientUnregisterOptions) (ExtendedZonesClientUnregisterResponse, error) {
	var err error
	const operationName = "ExtendedZonesClient.Unregister"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.unregisterCreateRequest(ctx, extendedZoneName, options)
	if err != nil {
		return ExtendedZonesClientUnregisterResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExtendedZonesClientUnregisterResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ExtendedZonesClientUnregisterResponse{}, err
	}
	resp, err := client.unregisterHandleResponse(httpResp)
	return resp, err
}

// unregisterCreateRequest creates the Unregister request.
func (client *ExtendedZonesClient) unregisterCreateRequest(ctx context.Context, extendedZoneName string, _ *ExtendedZonesClientUnregisterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.EdgeZones/extendedZones/{extendedZoneName}/unregister"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if extendedZoneName == "" {
		return nil, errors.New("parameter extendedZoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extendedZoneName}", url.PathEscape(extendedZoneName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// unregisterHandleResponse handles the Unregister response.
func (client *ExtendedZonesClient) unregisterHandleResponse(resp *http.Response) (ExtendedZonesClientUnregisterResponse, error) {
	result := ExtendedZonesClientUnregisterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtendedZone); err != nil {
		return ExtendedZonesClientUnregisterResponse{}, err
	}
	return result, nil
}
