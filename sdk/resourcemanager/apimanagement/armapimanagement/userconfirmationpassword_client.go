//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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

// UserConfirmationPasswordClient contains the methods for the UserConfirmationPassword group.
// Don't use this type directly, use NewUserConfirmationPasswordClient() instead.
type UserConfirmationPasswordClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewUserConfirmationPasswordClient creates a new instance of UserConfirmationPasswordClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewUserConfirmationPasswordClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*UserConfirmationPasswordClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &UserConfirmationPasswordClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Send - Sends confirmation
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - userID - User identifier. Must be unique in the current API Management service instance.
//   - options - UserConfirmationPasswordClientSendOptions contains the optional parameters for the UserConfirmationPasswordClient.Send
//     method.
func (client *UserConfirmationPasswordClient) Send(ctx context.Context, resourceGroupName string, serviceName string, userID string, options *UserConfirmationPasswordClientSendOptions) (UserConfirmationPasswordClientSendResponse, error) {
	var err error
	const operationName = "UserConfirmationPasswordClient.Send"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.sendCreateRequest(ctx, resourceGroupName, serviceName, userID, options)
	if err != nil {
		return UserConfirmationPasswordClientSendResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return UserConfirmationPasswordClientSendResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return UserConfirmationPasswordClientSendResponse{}, err
	}
	return UserConfirmationPasswordClientSendResponse{}, nil
}

// sendCreateRequest creates the Send request.
func (client *UserConfirmationPasswordClient) sendCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, userID string, options *UserConfirmationPasswordClientSendOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}/confirmations/password/send"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if userID == "" {
		return nil, errors.New("parameter userID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userId}", url.PathEscape(userID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	if options != nil && options.AppType != nil {
		reqQP.Set("appType", string(*options.AppType))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
