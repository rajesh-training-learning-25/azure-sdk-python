//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity

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

// SecureScoresClient contains the methods for the SecureScores group.
// Don't use this type directly, use NewSecureScoresClient() instead.
type SecureScoresClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSecureScoresClient creates a new instance of SecureScoresClient with the specified values.
// subscriptionID - Azure subscription ID
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSecureScoresClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SecureScoresClient, error) {
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
	client := &SecureScoresClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Get secure score for a specific Microsoft Defender for Cloud initiative within your current scope. For the ASC Default
// initiative, use 'ascScore'.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-01-01
// secureScoreName - The initiative name. For the ASC Default initiative, use 'ascScore' as in the sample request below.
// options - SecureScoresClientGetOptions contains the optional parameters for the SecureScoresClient.Get method.
func (client *SecureScoresClient) Get(ctx context.Context, secureScoreName string, options *SecureScoresClientGetOptions) (SecureScoresClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, secureScoreName, options)
	if err != nil {
		return SecureScoresClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SecureScoresClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SecureScoresClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SecureScoresClient) getCreateRequest(ctx context.Context, secureScoreName string, options *SecureScoresClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/secureScores/{secureScoreName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if secureScoreName == "" {
		return nil, errors.New("parameter secureScoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{secureScoreName}", url.PathEscape(secureScoreName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SecureScoresClient) getHandleResponse(resp *http.Response) (SecureScoresClientGetResponse, error) {
	result := SecureScoresClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecureScoreItem); err != nil {
		return SecureScoresClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List secure scores for all your Microsoft Defender for Cloud initiatives within your current scope.
// Generated from API version 2020-01-01
// options - SecureScoresClientListOptions contains the optional parameters for the SecureScoresClient.List method.
func (client *SecureScoresClient) NewListPager(options *SecureScoresClientListOptions) *runtime.Pager[SecureScoresClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SecureScoresClientListResponse]{
		More: func(page SecureScoresClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SecureScoresClientListResponse) (SecureScoresClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SecureScoresClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SecureScoresClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SecureScoresClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SecureScoresClient) listCreateRequest(ctx context.Context, options *SecureScoresClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/secureScores"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SecureScoresClient) listHandleResponse(resp *http.Response) (SecureScoresClientListResponse, error) {
	result := SecureScoresClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecureScoresList); err != nil {
		return SecureScoresClientListResponse{}, err
	}
	return result, nil
}