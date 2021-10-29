//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SecureScoreControlDefinitionsClient contains the methods for the SecureScoreControlDefinitions group.
// Don't use this type directly, use NewSecureScoreControlDefinitionsClient() instead.
type SecureScoreControlDefinitionsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewSecureScoreControlDefinitionsClient creates a new instance of SecureScoreControlDefinitionsClient with the specified values.
func NewSecureScoreControlDefinitionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SecureScoreControlDefinitionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &SecureScoreControlDefinitionsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// List - List the available security controls, their assessments, and the max score
// If the operation fails it returns the *CloudError error type.
func (client *SecureScoreControlDefinitionsClient) List(options *SecureScoreControlDefinitionsListOptions) *SecureScoreControlDefinitionsListPager {
	return &SecureScoreControlDefinitionsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp SecureScoreControlDefinitionsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SecureScoreControlDefinitionList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *SecureScoreControlDefinitionsClient) listCreateRequest(ctx context.Context, options *SecureScoreControlDefinitionsListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Security/secureScoreControlDefinitions"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SecureScoreControlDefinitionsClient) listHandleResponse(resp *http.Response) (SecureScoreControlDefinitionsListResponse, error) {
	result := SecureScoreControlDefinitionsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecureScoreControlDefinitionList); err != nil {
		return SecureScoreControlDefinitionsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *SecureScoreControlDefinitionsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListBySubscription - For a specified subscription, list the available security controls, their assessments, and the max score
// If the operation fails it returns the *CloudError error type.
func (client *SecureScoreControlDefinitionsClient) ListBySubscription(options *SecureScoreControlDefinitionsListBySubscriptionOptions) *SecureScoreControlDefinitionsListBySubscriptionPager {
	return &SecureScoreControlDefinitionsListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp SecureScoreControlDefinitionsListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SecureScoreControlDefinitionList.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *SecureScoreControlDefinitionsClient) listBySubscriptionCreateRequest(ctx context.Context, options *SecureScoreControlDefinitionsListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/secureScoreControlDefinitions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *SecureScoreControlDefinitionsClient) listBySubscriptionHandleResponse(resp *http.Response) (SecureScoreControlDefinitionsListBySubscriptionResponse, error) {
	result := SecureScoreControlDefinitionsListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecureScoreControlDefinitionList); err != nil {
		return SecureScoreControlDefinitionsListBySubscriptionResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *SecureScoreControlDefinitionsClient) listBySubscriptionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
