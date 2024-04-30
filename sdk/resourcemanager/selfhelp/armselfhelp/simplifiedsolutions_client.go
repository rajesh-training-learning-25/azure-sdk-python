//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armselfhelp

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

// SimplifiedSolutionsClient contains the methods for the SimplifiedSolutions group.
// Don't use this type directly, use NewSimplifiedSolutionsClient() instead.
type SimplifiedSolutionsClient struct {
	internal *arm.Client
}

// NewSimplifiedSolutionsClient creates a new instance of SimplifiedSolutionsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSimplifiedSolutionsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*SimplifiedSolutionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SimplifiedSolutionsClient{
		internal: cl,
	}
	return client, nil
}

// BeginCreate - Creates Simplified Solutions for an Azure subscription using 'solutionId' from Discovery Solutions as the
// input.
// Simplified Solutions API makes the consumption of solutions APIs easier while still providing access to the same powerful
// solutions rendered in Solutions API. With Simplified Solutions, users don't
// have to worry about stitching together the article using replacement maps and can use the content in the API response to
// directly render as HTML content.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01-preview
//   - scope - scope = resourceUri of affected resource.
//     For example: /subscriptions/0d0fcd2e-c4fd-4349-8497-200edb3923c6/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-non-read
//   - simplifiedSolutionsResourceName - Simplified Solutions Resource Name.
//   - simplifiedSolutionsRequestBody - The required request body for simplified Solutions resource creation.
//   - options - SimplifiedSolutionsClientBeginCreateOptions contains the optional parameters for the SimplifiedSolutionsClient.BeginCreate
//     method.
func (client *SimplifiedSolutionsClient) BeginCreate(ctx context.Context, scope string, simplifiedSolutionsResourceName string, simplifiedSolutionsRequestBody SimplifiedSolutionsResource, options *SimplifiedSolutionsClientBeginCreateOptions) (*runtime.Poller[SimplifiedSolutionsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, scope, simplifiedSolutionsResourceName, simplifiedSolutionsRequestBody, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SimplifiedSolutionsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SimplifiedSolutionsClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Creates Simplified Solutions for an Azure subscription using 'solutionId' from Discovery Solutions as the input.
// Simplified Solutions API makes the consumption of solutions APIs easier while still providing access to the same powerful
// solutions rendered in Solutions API. With Simplified Solutions, users don't
// have to worry about stitching together the article using replacement maps and can use the content in the API response to
// directly render as HTML content.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01-preview
func (client *SimplifiedSolutionsClient) create(ctx context.Context, scope string, simplifiedSolutionsResourceName string, simplifiedSolutionsRequestBody SimplifiedSolutionsResource, options *SimplifiedSolutionsClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "SimplifiedSolutionsClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, scope, simplifiedSolutionsResourceName, simplifiedSolutionsRequestBody, options)
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

// createCreateRequest creates the Create request.
func (client *SimplifiedSolutionsClient) createCreateRequest(ctx context.Context, scope string, simplifiedSolutionsResourceName string, simplifiedSolutionsRequestBody SimplifiedSolutionsResource, options *SimplifiedSolutionsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Help/simplifiedSolutions/{simplifiedSolutionsResourceName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if simplifiedSolutionsResourceName == "" {
		return nil, errors.New("parameter simplifiedSolutionsResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{simplifiedSolutionsResourceName}", url.PathEscape(simplifiedSolutionsResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, simplifiedSolutionsRequestBody); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Get the simplified Solutions using the applicable solutionResourceName while creating the simplified Solutions.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01-preview
//   - scope - scope = resourceUri of affected resource.
//     For example: /subscriptions/0d0fcd2e-c4fd-4349-8497-200edb3923c6/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-non-read
//   - simplifiedSolutionsResourceName - Simplified Solutions Resource Name.
//   - options - SimplifiedSolutionsClientGetOptions contains the optional parameters for the SimplifiedSolutionsClient.Get method.
func (client *SimplifiedSolutionsClient) Get(ctx context.Context, scope string, simplifiedSolutionsResourceName string, options *SimplifiedSolutionsClientGetOptions) (SimplifiedSolutionsClientGetResponse, error) {
	var err error
	const operationName = "SimplifiedSolutionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, scope, simplifiedSolutionsResourceName, options)
	if err != nil {
		return SimplifiedSolutionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SimplifiedSolutionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SimplifiedSolutionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SimplifiedSolutionsClient) getCreateRequest(ctx context.Context, scope string, simplifiedSolutionsResourceName string, options *SimplifiedSolutionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Help/simplifiedSolutions/{simplifiedSolutionsResourceName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if simplifiedSolutionsResourceName == "" {
		return nil, errors.New("parameter simplifiedSolutionsResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{simplifiedSolutionsResourceName}", url.PathEscape(simplifiedSolutionsResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SimplifiedSolutionsClient) getHandleResponse(resp *http.Response) (SimplifiedSolutionsClientGetResponse, error) {
	result := SimplifiedSolutionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SimplifiedSolutionsResource); err != nil {
		return SimplifiedSolutionsClientGetResponse{}, err
	}
	return result, nil
}
