// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armcontainerorchestratorruntime

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

// BgpPeersClient contains the methods for the BgpPeers group.
// Don't use this type directly, use NewBgpPeersClient() instead.
type BgpPeersClient struct {
	internal *arm.Client
}

// NewBgpPeersClient creates a new instance of BgpPeersClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBgpPeersClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*BgpPeersClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BgpPeersClient{
	internal: cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a BgpPeer
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - bgpPeerName - The name of the BgpPeer
//   - resource - Resource create parameters.
//   - options - BgpPeersClientBeginCreateOrUpdateOptions contains the optional parameters for the BgpPeersClient.BeginCreateOrUpdate
//     method.
func (client *BgpPeersClient) BeginCreateOrUpdate(ctx context.Context, resourceURI string, bgpPeerName string, resource BgpPeer, options *BgpPeersClientBeginCreateOrUpdateOptions) (*runtime.Poller[BgpPeersClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceURI, bgpPeerName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[BgpPeersClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[BgpPeersClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a BgpPeer
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
func (client *BgpPeersClient) createOrUpdate(ctx context.Context, resourceURI string, bgpPeerName string, resource BgpPeer, options *BgpPeersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "BgpPeersClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceURI, bgpPeerName, resource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *BgpPeersClient) createOrUpdateCreateRequest(ctx context.Context, resourceURI string, bgpPeerName string, resource BgpPeer, _ *BgpPeersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.KubernetesRuntime/bgpPeers/{bgpPeerName}"
	if resourceURI == "" {
		return nil, errors.New("parameter resourceURI cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", url.PathEscape(resourceURI))
	if bgpPeerName == "" {
		return nil, errors.New("parameter bgpPeerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{bgpPeerName}", url.PathEscape(bgpPeerName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
	return nil, err
}
	return req, nil
}

// Delete - Delete a BgpPeer
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - bgpPeerName - The name of the BgpPeer
//   - options - BgpPeersClientDeleteOptions contains the optional parameters for the BgpPeersClient.Delete method.
func (client *BgpPeersClient) Delete(ctx context.Context, resourceURI string, bgpPeerName string, options *BgpPeersClientDeleteOptions) (BgpPeersClientDeleteResponse, error) {
	var err error
	const operationName = "BgpPeersClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceURI, bgpPeerName, options)
	if err != nil {
		return BgpPeersClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BgpPeersClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return BgpPeersClientDeleteResponse{}, err
	}
	return BgpPeersClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *BgpPeersClient) deleteCreateRequest(ctx context.Context, resourceURI string, bgpPeerName string, _ *BgpPeersClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.KubernetesRuntime/bgpPeers/{bgpPeerName}"
	if resourceURI == "" {
		return nil, errors.New("parameter resourceURI cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", url.PathEscape(resourceURI))
	if bgpPeerName == "" {
		return nil, errors.New("parameter bgpPeerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{bgpPeerName}", url.PathEscape(bgpPeerName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a BgpPeer
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - bgpPeerName - The name of the BgpPeer
//   - options - BgpPeersClientGetOptions contains the optional parameters for the BgpPeersClient.Get method.
func (client *BgpPeersClient) Get(ctx context.Context, resourceURI string, bgpPeerName string, options *BgpPeersClientGetOptions) (BgpPeersClientGetResponse, error) {
	var err error
	const operationName = "BgpPeersClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceURI, bgpPeerName, options)
	if err != nil {
		return BgpPeersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BgpPeersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BgpPeersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *BgpPeersClient) getCreateRequest(ctx context.Context, resourceURI string, bgpPeerName string, _ *BgpPeersClientGetOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.KubernetesRuntime/bgpPeers/{bgpPeerName}"
	if resourceURI == "" {
		return nil, errors.New("parameter resourceURI cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", url.PathEscape(resourceURI))
	if bgpPeerName == "" {
		return nil, errors.New("parameter bgpPeerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{bgpPeerName}", url.PathEscape(bgpPeerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BgpPeersClient) getHandleResponse(resp *http.Response) (BgpPeersClientGetResponse, error) {
	result := BgpPeersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BgpPeer); err != nil {
		return BgpPeersClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List BgpPeer resources by parent
//
// Generated from API version 2024-03-01
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - options - BgpPeersClientListOptions contains the optional parameters for the BgpPeersClient.NewListPager method.
func (client *BgpPeersClient) NewListPager(resourceURI string, options *BgpPeersClientListOptions) (*runtime.Pager[BgpPeersClientListResponse]) {
	return runtime.NewPager(runtime.PagingHandler[BgpPeersClientListResponse]{
		More: func(page BgpPeersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BgpPeersClientListResponse) (BgpPeersClientListResponse, error) {
		ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "BgpPeersClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceURI, options)
			}, nil)
			if err != nil {
				return BgpPeersClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
			},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *BgpPeersClient) listCreateRequest(ctx context.Context, resourceURI string, _ *BgpPeersClientListOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.KubernetesRuntime/bgpPeers"
	if resourceURI == "" {
		return nil, errors.New("parameter resourceURI cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", url.PathEscape(resourceURI))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BgpPeersClient) listHandleResponse(resp *http.Response) (BgpPeersClientListResponse, error) {
	result := BgpPeersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BgpPeerListResult); err != nil {
		return BgpPeersClientListResponse{}, err
	}
	return result, nil
}

