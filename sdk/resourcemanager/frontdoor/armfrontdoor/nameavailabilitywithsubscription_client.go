//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfrontdoor

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

// NameAvailabilityWithSubscriptionClient contains the methods for the FrontDoorNameAvailabilityWithSubscription group.
// Don't use this type directly, use NewNameAvailabilityWithSubscriptionClient() instead.
type NameAvailabilityWithSubscriptionClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNameAvailabilityWithSubscriptionClient creates a new instance of NameAvailabilityWithSubscriptionClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNameAvailabilityWithSubscriptionClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NameAvailabilityWithSubscriptionClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NameAvailabilityWithSubscriptionClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Check - Check the availability of a Front Door subdomain.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - checkFrontDoorNameAvailabilityInput - Input to check.
//   - options - NameAvailabilityWithSubscriptionClientCheckOptions contains the optional parameters for the NameAvailabilityWithSubscriptionClient.Check
//     method.
func (client *NameAvailabilityWithSubscriptionClient) Check(ctx context.Context, checkFrontDoorNameAvailabilityInput CheckNameAvailabilityInput, options *NameAvailabilityWithSubscriptionClientCheckOptions) (NameAvailabilityWithSubscriptionClientCheckResponse, error) {
	var err error
	const operationName = "NameAvailabilityWithSubscriptionClient.Check"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.checkCreateRequest(ctx, checkFrontDoorNameAvailabilityInput, options)
	if err != nil {
		return NameAvailabilityWithSubscriptionClientCheckResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NameAvailabilityWithSubscriptionClientCheckResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NameAvailabilityWithSubscriptionClientCheckResponse{}, err
	}
	resp, err := client.checkHandleResponse(httpResp)
	return resp, err
}

// checkCreateRequest creates the Check request.
func (client *NameAvailabilityWithSubscriptionClient) checkCreateRequest(ctx context.Context, checkFrontDoorNameAvailabilityInput CheckNameAvailabilityInput, options *NameAvailabilityWithSubscriptionClientCheckOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/checkFrontDoorNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, checkFrontDoorNameAvailabilityInput); err != nil {
		return nil, err
	}
	return req, nil
}

// checkHandleResponse handles the Check response.
func (client *NameAvailabilityWithSubscriptionClient) checkHandleResponse(resp *http.Response) (NameAvailabilityWithSubscriptionClientCheckResponse, error) {
	result := NameAvailabilityWithSubscriptionClientCheckResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityOutput); err != nil {
		return NameAvailabilityWithSubscriptionClientCheckResponse{}, err
	}
	return result, nil
}
