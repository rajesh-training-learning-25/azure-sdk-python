// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// RestorableMongodbResourcesClient contains the methods for the RestorableMongodbResources group.
// Don't use this type directly, use NewRestorableMongodbResourcesClient() instead.
type RestorableMongodbResourcesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewRestorableMongodbResourcesClient creates a new instance of RestorableMongodbResourcesClient with the specified values.
func NewRestorableMongodbResourcesClient(con *armcore.Connection, subscriptionID string) *RestorableMongodbResourcesClient {
	return &RestorableMongodbResourcesClient{con: con, subscriptionID: subscriptionID}
}

// List - Return a list of database and collection combo that exist on the account at the given timestamp and location. This helps in scenarios to validate
// what resources exist at given timestamp and location.
// This API requires 'Microsoft.DocumentDB/locations/restorableDatabaseAccounts/…/read' permission.
// If the operation fails it returns the *CloudError error type.
func (client *RestorableMongodbResourcesClient) List(ctx context.Context, location string, instanceID string, options *RestorableMongodbResourcesListOptions) (RestorableMongodbResourcesListResponse, error) {
	req, err := client.listCreateRequest(ctx, location, instanceID, options)
	if err != nil {
		return RestorableMongodbResourcesListResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RestorableMongodbResourcesListResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return RestorableMongodbResourcesListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *RestorableMongodbResourcesClient) listCreateRequest(ctx context.Context, location string, instanceID string, options *RestorableMongodbResourcesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{instanceId}/restorableMongodbResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-15")
	if options != nil && options.RestoreLocation != nil {
		reqQP.Set("restoreLocation", *options.RestoreLocation)
	}
	if options != nil && options.RestoreTimestampInUTC != nil {
		reqQP.Set("restoreTimestampInUtc", *options.RestoreTimestampInUTC)
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RestorableMongodbResourcesClient) listHandleResponse(resp *azcore.Response) (RestorableMongodbResourcesListResponse, error) {
	result := RestorableMongodbResourcesListResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.RestorableMongodbResourcesListResult); err != nil {
		return RestorableMongodbResourcesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *RestorableMongodbResourcesClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
