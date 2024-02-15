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

// GetTieringCostOperationResultClient contains the methods for the GetTieringCostOperationResult group.
// Don't use this type directly, use NewGetTieringCostOperationResultClient() instead.
type GetTieringCostOperationResultClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewGetTieringCostOperationResultClient creates a new instance of GetTieringCostOperationResultClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGetTieringCostOperationResultClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GetTieringCostOperationResultClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &GetTieringCostOperationResultClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets the result of async operation for tiering cost
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - vaultName - The name of the recovery services vault.
//   - options - GetTieringCostOperationResultClientGetOptions contains the optional parameters for the GetTieringCostOperationResultClient.Get
//     method.
func (client *GetTieringCostOperationResultClient) Get(ctx context.Context, resourceGroupName string, vaultName string, operationID string, options *GetTieringCostOperationResultClientGetOptions) (GetTieringCostOperationResultClientGetResponse, error) {
	var err error
	const operationName = "GetTieringCostOperationResultClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, vaultName, operationID, options)
	if err != nil {
		return GetTieringCostOperationResultClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GetTieringCostOperationResultClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GetTieringCostOperationResultClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *GetTieringCostOperationResultClient) getCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, operationID string, options *GetTieringCostOperationResultClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupTieringCost/default/operationResults/{operationId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GetTieringCostOperationResultClient) getHandleResponse(resp *http.Response) (GetTieringCostOperationResultClientGetResponse, error) {
	result := GetTieringCostOperationResultClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return GetTieringCostOperationResultClientGetResponse{}, err
	}
	return result, nil
}
