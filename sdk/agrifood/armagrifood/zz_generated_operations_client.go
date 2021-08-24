//go:build go1.13
// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armagrifood

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// OperationsClient contains the methods for the Operations group.
// Don't use this type directly, use NewOperationsClient() instead.
type OperationsClient struct {
	con *armcore.Connection
}

// NewOperationsClient creates a new instance of OperationsClient with the specified values.
func NewOperationsClient(con *armcore.Connection) *OperationsClient {
	return &OperationsClient{con: con}
}

// List - Lists the available operations of Microsoft.AgFoodPlatform resource provider.
// If the operation fails it returns the *ErrorResponse error type.
func (client *OperationsClient) List(options *OperationsListOptions) OperationsListPager {
	return &operationsListPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp OperationsListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.OperationListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *OperationsClient) listCreateRequest(ctx context.Context, options *OperationsListOptions) (*azcore.Request, error) {
	urlPath := "/providers/Microsoft.AgFoodPlatform/operations"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-05-12-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *OperationsClient) listHandleResponse(resp *azcore.Response) (OperationsListResponse, error) {
	result := OperationsListResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.OperationListResult); err != nil {
		return OperationsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *OperationsClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
