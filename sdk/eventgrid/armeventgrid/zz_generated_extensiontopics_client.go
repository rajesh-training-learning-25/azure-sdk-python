//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventgrid

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// ExtensionTopicsClient contains the methods for the ExtensionTopics group.
// Don't use this type directly, use NewExtensionTopicsClient() instead.
type ExtensionTopicsClient struct {
	ep string
	pl runtime.Pipeline
}

// NewExtensionTopicsClient creates a new instance of ExtensionTopicsClient with the specified values.
func NewExtensionTopicsClient(con *arm.Connection) *ExtensionTopicsClient {
	return &ExtensionTopicsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version)}
}

// Get - Get the properties of an extension topic.
// If the operation fails it returns a generic error.
func (client *ExtensionTopicsClient) Get(ctx context.Context, scope string, options *ExtensionTopicsGetOptions) (ExtensionTopicsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, options)
	if err != nil {
		return ExtensionTopicsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExtensionTopicsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExtensionTopicsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExtensionTopicsClient) getCreateRequest(ctx context.Context, scope string, options *ExtensionTopicsGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.EventGrid/extensionTopics/default"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", url.PathEscape(scope))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExtensionTopicsClient) getHandleResponse(resp *http.Response) (ExtensionTopicsGetResponse, error) {
	result := ExtensionTopicsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtensionTopic); err != nil {
		return ExtensionTopicsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ExtensionTopicsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
