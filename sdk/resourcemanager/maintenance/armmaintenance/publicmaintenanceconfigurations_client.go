//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmaintenance

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

// PublicMaintenanceConfigurationsClient contains the methods for the PublicMaintenanceConfigurations group.
// Don't use this type directly, use NewPublicMaintenanceConfigurationsClient() instead.
type PublicMaintenanceConfigurationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPublicMaintenanceConfigurationsClient creates a new instance of PublicMaintenanceConfigurationsClient with the specified values.
//   - subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPublicMaintenanceConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PublicMaintenanceConfigurationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PublicMaintenanceConfigurationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get Public Maintenance Configuration record
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceName - Maintenance Configuration Name
//   - options - PublicMaintenanceConfigurationsClientGetOptions contains the optional parameters for the PublicMaintenanceConfigurationsClient.Get
//     method.
func (client *PublicMaintenanceConfigurationsClient) Get(ctx context.Context, resourceName string, options *PublicMaintenanceConfigurationsClientGetOptions) (PublicMaintenanceConfigurationsClientGetResponse, error) {
	var err error
	const operationName = "PublicMaintenanceConfigurationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceName, options)
	if err != nil {
		return PublicMaintenanceConfigurationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PublicMaintenanceConfigurationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PublicMaintenanceConfigurationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PublicMaintenanceConfigurationsClient) getCreateRequest(ctx context.Context, resourceName string, options *PublicMaintenanceConfigurationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PublicMaintenanceConfigurationsClient) getHandleResponse(resp *http.Response) (PublicMaintenanceConfigurationsClientGetResponse, error) {
	result := PublicMaintenanceConfigurationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Configuration); err != nil {
		return PublicMaintenanceConfigurationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get Public Maintenance Configuration records
//
// Generated from API version 2023-04-01
//   - options - PublicMaintenanceConfigurationsClientListOptions contains the optional parameters for the PublicMaintenanceConfigurationsClient.NewListPager
//     method.
func (client *PublicMaintenanceConfigurationsClient) NewListPager(options *PublicMaintenanceConfigurationsClientListOptions) *runtime.Pager[PublicMaintenanceConfigurationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PublicMaintenanceConfigurationsClientListResponse]{
		More: func(page PublicMaintenanceConfigurationsClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *PublicMaintenanceConfigurationsClientListResponse) (PublicMaintenanceConfigurationsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PublicMaintenanceConfigurationsClient.NewListPager")
			req, err := client.listCreateRequest(ctx, options)
			if err != nil {
				return PublicMaintenanceConfigurationsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PublicMaintenanceConfigurationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PublicMaintenanceConfigurationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *PublicMaintenanceConfigurationsClient) listCreateRequest(ctx context.Context, options *PublicMaintenanceConfigurationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Maintenance/publicMaintenanceConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PublicMaintenanceConfigurationsClient) listHandleResponse(resp *http.Response) (PublicMaintenanceConfigurationsClientListResponse, error) {
	result := PublicMaintenanceConfigurationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListMaintenanceConfigurationsResult); err != nil {
		return PublicMaintenanceConfigurationsClientListResponse{}, err
	}
	return result, nil
}
