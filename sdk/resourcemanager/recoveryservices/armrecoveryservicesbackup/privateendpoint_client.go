//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesbackup

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

// PrivateEndpointClient contains the methods for the PrivateEndpoint group.
// Don't use this type directly, use NewPrivateEndpointClient() instead.
type PrivateEndpointClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPrivateEndpointClient creates a new instance of PrivateEndpointClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPrivateEndpointClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateEndpointClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PrivateEndpointClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// GetOperationStatus - Gets the operation status for a private endpoint connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - privateEndpointConnectionName - The name of the private endpoint connection.
//   - operationID - Operation id
//   - options - PrivateEndpointClientGetOperationStatusOptions contains the optional parameters for the PrivateEndpointClient.GetOperationStatus
//     method.
func (client *PrivateEndpointClient) GetOperationStatus(ctx context.Context, vaultName string, resourceGroupName string, privateEndpointConnectionName string, operationID string, options *PrivateEndpointClientGetOperationStatusOptions) (PrivateEndpointClientGetOperationStatusResponse, error) {
	var err error
	const operationName = "PrivateEndpointClient.GetOperationStatus"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getOperationStatusCreateRequest(ctx, vaultName, resourceGroupName, privateEndpointConnectionName, operationID, options)
	if err != nil {
		return PrivateEndpointClientGetOperationStatusResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateEndpointClientGetOperationStatusResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PrivateEndpointClientGetOperationStatusResponse{}, err
	}
	resp, err := client.getOperationStatusHandleResponse(httpResp)
	return resp, err
}

// getOperationStatusCreateRequest creates the GetOperationStatus request.
func (client *PrivateEndpointClient) getOperationStatusCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, privateEndpointConnectionName string, operationID string, options *PrivateEndpointClientGetOperationStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/privateEndpointConnections/{privateEndpointConnectionName}/operationsStatus/{operationId}"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
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

// getOperationStatusHandleResponse handles the GetOperationStatus response.
func (client *PrivateEndpointClient) getOperationStatusHandleResponse(resp *http.Response) (PrivateEndpointClientGetOperationStatusResponse, error) {
	result := PrivateEndpointClientGetOperationStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationStatus); err != nil {
		return PrivateEndpointClientGetOperationStatusResponse{}, err
	}
	return result, nil
}
