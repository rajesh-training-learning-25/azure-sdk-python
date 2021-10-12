//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// CassandraClustersClient contains the methods for the CassandraClusters group.
// Don't use this type directly, use NewCassandraClustersClient() instead.
type CassandraClustersClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewCassandraClustersClient creates a new instance of CassandraClustersClient with the specified values.
func NewCassandraClustersClient(con *arm.Connection, subscriptionID string) *CassandraClustersClient {
	return &CassandraClustersClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateUpdate - Create or update a managed Cassandra cluster. When updating, you must specify all writable properties. To update only some properties,
// use PATCH.
// If the operation fails it returns the *CloudError error type.
func (client *CassandraClustersClient) BeginCreateUpdate(ctx context.Context, resourceGroupName string, clusterName string, body ClusterResource, options *CassandraClustersBeginCreateUpdateOptions) (CassandraClustersCreateUpdatePollerResponse, error) {
	resp, err := client.createUpdate(ctx, resourceGroupName, clusterName, body, options)
	if err != nil {
		return CassandraClustersCreateUpdatePollerResponse{}, err
	}
	result := CassandraClustersCreateUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("CassandraClustersClient.CreateUpdate", "", resp, client.pl, client.createUpdateHandleError)
	if err != nil {
		return CassandraClustersCreateUpdatePollerResponse{}, err
	}
	result.Poller = &CassandraClustersCreateUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateUpdate - Create or update a managed Cassandra cluster. When updating, you must specify all writable properties. To update only some properties,
// use PATCH.
// If the operation fails it returns the *CloudError error type.
func (client *CassandraClustersClient) createUpdate(ctx context.Context, resourceGroupName string, clusterName string, body ClusterResource, options *CassandraClustersBeginCreateUpdateOptions) (*http.Response, error) {
	req, err := client.createUpdateCreateRequest(ctx, resourceGroupName, clusterName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createUpdateHandleError(resp)
	}
	return resp, nil
}

// createUpdateCreateRequest creates the CreateUpdate request.
func (client *CassandraClustersClient) createUpdateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, body ClusterResource, options *CassandraClustersBeginCreateUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/cassandraClusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// createUpdateHandleError handles the CreateUpdate error response.
func (client *CassandraClustersClient) createUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes a managed Cassandra cluster.
// If the operation fails it returns the *CloudError error type.
func (client *CassandraClustersClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, options *CassandraClustersBeginDeleteOptions) (CassandraClustersDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return CassandraClustersDeletePollerResponse{}, err
	}
	result := CassandraClustersDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("CassandraClustersClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return CassandraClustersDeletePollerResponse{}, err
	}
	result.Poller = &CassandraClustersDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a managed Cassandra cluster.
// If the operation fails it returns the *CloudError error type.
func (client *CassandraClustersClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, options *CassandraClustersBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CassandraClustersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *CassandraClustersBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/cassandraClusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *CassandraClustersClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginFetchNodeStatus - Request the status of all nodes in the cluster (as returned by 'nodetool status').
// If the operation fails it returns the *CloudError error type.
func (client *CassandraClustersClient) BeginFetchNodeStatus(ctx context.Context, resourceGroupName string, clusterName string, options *CassandraClustersBeginFetchNodeStatusOptions) (CassandraClustersFetchNodeStatusPollerResponse, error) {
	resp, err := client.fetchNodeStatus(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return CassandraClustersFetchNodeStatusPollerResponse{}, err
	}
	result := CassandraClustersFetchNodeStatusPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("CassandraClustersClient.FetchNodeStatus", "", resp, client.pl, client.fetchNodeStatusHandleError)
	if err != nil {
		return CassandraClustersFetchNodeStatusPollerResponse{}, err
	}
	result.Poller = &CassandraClustersFetchNodeStatusPoller{
		pt: pt,
	}
	return result, nil
}

// FetchNodeStatus - Request the status of all nodes in the cluster (as returned by 'nodetool status').
// If the operation fails it returns the *CloudError error type.
func (client *CassandraClustersClient) fetchNodeStatus(ctx context.Context, resourceGroupName string, clusterName string, options *CassandraClustersBeginFetchNodeStatusOptions) (*http.Response, error) {
	req, err := client.fetchNodeStatusCreateRequest(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.fetchNodeStatusHandleError(resp)
	}
	return resp, nil
}

// fetchNodeStatusCreateRequest creates the FetchNodeStatus request.
func (client *CassandraClustersClient) fetchNodeStatusCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *CassandraClustersBeginFetchNodeStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/cassandraClusters/{clusterName}/fetchNodeStatus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// fetchNodeStatusHandleError handles the FetchNodeStatus error response.
func (client *CassandraClustersClient) fetchNodeStatusHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Get the properties of a managed Cassandra cluster.
// If the operation fails it returns the *CloudError error type.
func (client *CassandraClustersClient) Get(ctx context.Context, resourceGroupName string, clusterName string, options *CassandraClustersGetOptions) (CassandraClustersGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return CassandraClustersGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CassandraClustersGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CassandraClustersGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CassandraClustersClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *CassandraClustersGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/cassandraClusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CassandraClustersClient) getHandleResponse(resp *http.Response) (CassandraClustersGetResponse, error) {
	result := CassandraClustersGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterResource); err != nil {
		return CassandraClustersGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *CassandraClustersClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetBackup - Get the properties of an individual backup of this cluster that is available to restore.
// If the operation fails it returns the *CloudError error type.
func (client *CassandraClustersClient) GetBackup(ctx context.Context, resourceGroupName string, clusterName string, backupID string, options *CassandraClustersGetBackupOptions) (CassandraClustersGetBackupResponse, error) {
	req, err := client.getBackupCreateRequest(ctx, resourceGroupName, clusterName, backupID, options)
	if err != nil {
		return CassandraClustersGetBackupResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CassandraClustersGetBackupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CassandraClustersGetBackupResponse{}, client.getBackupHandleError(resp)
	}
	return client.getBackupHandleResponse(resp)
}

// getBackupCreateRequest creates the GetBackup request.
func (client *CassandraClustersClient) getBackupCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, backupID string, options *CassandraClustersGetBackupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/cassandraClusters/{clusterName}/backups/{backupId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if backupID == "" {
		return nil, errors.New("parameter backupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupId}", url.PathEscape(backupID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getBackupHandleResponse handles the GetBackup response.
func (client *CassandraClustersClient) getBackupHandleResponse(resp *http.Response) (CassandraClustersGetBackupResponse, error) {
	result := CassandraClustersGetBackupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupResource); err != nil {
		return CassandraClustersGetBackupResponse{}, err
	}
	return result, nil
}

// getBackupHandleError handles the GetBackup error response.
func (client *CassandraClustersClient) getBackupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListBackups - List the backups of this cluster that are available to restore.
// If the operation fails it returns the *CloudError error type.
func (client *CassandraClustersClient) ListBackups(ctx context.Context, resourceGroupName string, clusterName string, options *CassandraClustersListBackupsOptions) (CassandraClustersListBackupsResponse, error) {
	req, err := client.listBackupsCreateRequest(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return CassandraClustersListBackupsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CassandraClustersListBackupsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CassandraClustersListBackupsResponse{}, client.listBackupsHandleError(resp)
	}
	return client.listBackupsHandleResponse(resp)
}

// listBackupsCreateRequest creates the ListBackups request.
func (client *CassandraClustersClient) listBackupsCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *CassandraClustersListBackupsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/cassandraClusters/{clusterName}/backups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBackupsHandleResponse handles the ListBackups response.
func (client *CassandraClustersClient) listBackupsHandleResponse(resp *http.Response) (CassandraClustersListBackupsResponse, error) {
	result := CassandraClustersListBackupsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListBackups); err != nil {
		return CassandraClustersListBackupsResponse{}, err
	}
	return result, nil
}

// listBackupsHandleError handles the ListBackups error response.
func (client *CassandraClustersClient) listBackupsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - List all managed Cassandra clusters in this resource group.
// If the operation fails it returns the *CloudError error type.
func (client *CassandraClustersClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, options *CassandraClustersListByResourceGroupOptions) (CassandraClustersListByResourceGroupResponse, error) {
	req, err := client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
	if err != nil {
		return CassandraClustersListByResourceGroupResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CassandraClustersListByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CassandraClustersListByResourceGroupResponse{}, client.listByResourceGroupHandleError(resp)
	}
	return client.listByResourceGroupHandleResponse(resp)
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *CassandraClustersClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *CassandraClustersListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/cassandraClusters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *CassandraClustersClient) listByResourceGroupHandleResponse(resp *http.Response) (CassandraClustersListByResourceGroupResponse, error) {
	result := CassandraClustersListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListClusters); err != nil {
		return CassandraClustersListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *CassandraClustersClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListBySubscription - List all managed Cassandra clusters in this subscription.
// If the operation fails it returns the *CloudError error type.
func (client *CassandraClustersClient) ListBySubscription(ctx context.Context, options *CassandraClustersListBySubscriptionOptions) (CassandraClustersListBySubscriptionResponse, error) {
	req, err := client.listBySubscriptionCreateRequest(ctx, options)
	if err != nil {
		return CassandraClustersListBySubscriptionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CassandraClustersListBySubscriptionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CassandraClustersListBySubscriptionResponse{}, client.listBySubscriptionHandleError(resp)
	}
	return client.listBySubscriptionHandleResponse(resp)
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *CassandraClustersClient) listBySubscriptionCreateRequest(ctx context.Context, options *CassandraClustersListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/cassandraClusters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *CassandraClustersClient) listBySubscriptionHandleResponse(resp *http.Response) (CassandraClustersListBySubscriptionResponse, error) {
	result := CassandraClustersListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListClusters); err != nil {
		return CassandraClustersListBySubscriptionResponse{}, err
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *CassandraClustersClient) listBySubscriptionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginRequestRepair - Request that repair begin on this cluster as soon as possible.
// If the operation fails it returns the *CloudError error type.
func (client *CassandraClustersClient) BeginRequestRepair(ctx context.Context, resourceGroupName string, clusterName string, body RepairPostBody, options *CassandraClustersBeginRequestRepairOptions) (CassandraClustersRequestRepairPollerResponse, error) {
	resp, err := client.requestRepair(ctx, resourceGroupName, clusterName, body, options)
	if err != nil {
		return CassandraClustersRequestRepairPollerResponse{}, err
	}
	result := CassandraClustersRequestRepairPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("CassandraClustersClient.RequestRepair", "", resp, client.pl, client.requestRepairHandleError)
	if err != nil {
		return CassandraClustersRequestRepairPollerResponse{}, err
	}
	result.Poller = &CassandraClustersRequestRepairPoller{
		pt: pt,
	}
	return result, nil
}

// RequestRepair - Request that repair begin on this cluster as soon as possible.
// If the operation fails it returns the *CloudError error type.
func (client *CassandraClustersClient) requestRepair(ctx context.Context, resourceGroupName string, clusterName string, body RepairPostBody, options *CassandraClustersBeginRequestRepairOptions) (*http.Response, error) {
	req, err := client.requestRepairCreateRequest(ctx, resourceGroupName, clusterName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, client.requestRepairHandleError(resp)
	}
	return resp, nil
}

// requestRepairCreateRequest creates the RequestRepair request.
func (client *CassandraClustersClient) requestRepairCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, body RepairPostBody, options *CassandraClustersBeginRequestRepairOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/cassandraClusters/{clusterName}/repair"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// requestRepairHandleError handles the RequestRepair error response.
func (client *CassandraClustersClient) requestRepairHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginUpdate - Updates some of the properties of a managed Cassandra cluster.
// If the operation fails it returns the *CloudError error type.
func (client *CassandraClustersClient) BeginUpdate(ctx context.Context, resourceGroupName string, clusterName string, body ClusterResource, options *CassandraClustersBeginUpdateOptions) (CassandraClustersUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, clusterName, body, options)
	if err != nil {
		return CassandraClustersUpdatePollerResponse{}, err
	}
	result := CassandraClustersUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("CassandraClustersClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return CassandraClustersUpdatePollerResponse{}, err
	}
	result.Poller = &CassandraClustersUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates some of the properties of a managed Cassandra cluster.
// If the operation fails it returns the *CloudError error type.
func (client *CassandraClustersClient) update(ctx context.Context, resourceGroupName string, clusterName string, body ClusterResource, options *CassandraClustersBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, clusterName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *CassandraClustersClient) updateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, body ClusterResource, options *CassandraClustersBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/cassandraClusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// updateHandleError handles the Update error response.
func (client *CassandraClustersClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
