//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsubscriptions

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/profile/v20200901/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// SubscriptionClient contains the methods for the SubscriptionClient group.
// Don't use this type directly, use NewSubscriptionClient() instead.
type SubscriptionClient struct {
	host string
	pl   runtime.Pipeline
}

// NewSubscriptionClient creates a new instance of SubscriptionClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSubscriptionClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*SubscriptionClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(internal.ModuleName, internal.ModuleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &SubscriptionClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// CheckResourceName - A resource name is valid if it is not a reserved word, does not contains a reserved word and does not
// start with a reserved word
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2016-06-01
// options - SubscriptionClientCheckResourceNameOptions contains the optional parameters for the SubscriptionClient.CheckResourceName
// method.
func (client *SubscriptionClient) CheckResourceName(ctx context.Context, options *SubscriptionClientCheckResourceNameOptions) (SubscriptionClientCheckResourceNameResponse, error) {
	req, err := client.checkResourceNameCreateRequest(ctx, options)
	if err != nil {
		return SubscriptionClientCheckResourceNameResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SubscriptionClientCheckResourceNameResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubscriptionClientCheckResourceNameResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkResourceNameHandleResponse(resp)
}

// checkResourceNameCreateRequest creates the CheckResourceName request.
func (client *SubscriptionClient) checkResourceNameCreateRequest(ctx context.Context, options *SubscriptionClientCheckResourceNameOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Resources/checkResourceName"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.ResourceNameDefinition != nil {
		return req, runtime.MarshalAsJSON(req, *options.ResourceNameDefinition)
	}
	return req, nil
}

// checkResourceNameHandleResponse handles the CheckResourceName response.
func (client *SubscriptionClient) checkResourceNameHandleResponse(resp *http.Response) (SubscriptionClientCheckResourceNameResponse, error) {
	result := SubscriptionClientCheckResourceNameResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckResourceNameResult); err != nil {
		return SubscriptionClientCheckResourceNameResponse{}, err
	}
	return result, nil
}
