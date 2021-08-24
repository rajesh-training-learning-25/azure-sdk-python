//go:build go1.13
// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// GalleryApplicationsClient contains the methods for the GalleryApplications group.
// Don't use this type directly, use NewGalleryApplicationsClient() instead.
type GalleryApplicationsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewGalleryApplicationsClient creates a new instance of GalleryApplicationsClient with the specified values.
func NewGalleryApplicationsClient(con *armcore.Connection, subscriptionID string) *GalleryApplicationsClient {
	return &GalleryApplicationsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Create or update a gallery Application Definition.
// If the operation fails it returns the *CloudError error type.
func (client *GalleryApplicationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplication GalleryApplication, options *GalleryApplicationsBeginCreateOrUpdateOptions) (GalleryApplicationsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplication, options)
	if err != nil {
		return GalleryApplicationsCreateOrUpdatePollerResponse{}, err
	}
	result := GalleryApplicationsCreateOrUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("GalleryApplicationsClient.CreateOrUpdate", "", resp, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return GalleryApplicationsCreateOrUpdatePollerResponse{}, err
	}
	poller := &galleryApplicationsCreateOrUpdatePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (GalleryApplicationsCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new GalleryApplicationsCreateOrUpdatePoller from the specified resume token.
// token - The value must come from a previous call to GalleryApplicationsCreateOrUpdatePoller.ResumeToken().
func (client *GalleryApplicationsClient) ResumeCreateOrUpdate(ctx context.Context, token string) (GalleryApplicationsCreateOrUpdatePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("GalleryApplicationsClient.CreateOrUpdate", token, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return GalleryApplicationsCreateOrUpdatePollerResponse{}, err
	}
	poller := &galleryApplicationsCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return GalleryApplicationsCreateOrUpdatePollerResponse{}, err
	}
	result := GalleryApplicationsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (GalleryApplicationsCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Create or update a gallery Application Definition.
// If the operation fails it returns the *CloudError error type.
func (client *GalleryApplicationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplication GalleryApplication, options *GalleryApplicationsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplication, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *GalleryApplicationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplication GalleryApplication, options *GalleryApplicationsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if galleryApplicationName == "" {
		return nil, errors.New("parameter galleryApplicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationName}", url.PathEscape(galleryApplicationName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-09-30")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(galleryApplication)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *GalleryApplicationsClient) createOrUpdateHandleError(resp *azcore.Response) error {
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

// BeginDelete - Delete a gallery Application.
// If the operation fails it returns the *CloudError error type.
func (client *GalleryApplicationsClient) BeginDelete(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, options *GalleryApplicationsBeginDeleteOptions) (GalleryApplicationsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, galleryName, galleryApplicationName, options)
	if err != nil {
		return GalleryApplicationsDeletePollerResponse{}, err
	}
	result := GalleryApplicationsDeletePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("GalleryApplicationsClient.Delete", "", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return GalleryApplicationsDeletePollerResponse{}, err
	}
	poller := &galleryApplicationsDeletePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (GalleryApplicationsDeleteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new GalleryApplicationsDeletePoller from the specified resume token.
// token - The value must come from a previous call to GalleryApplicationsDeletePoller.ResumeToken().
func (client *GalleryApplicationsClient) ResumeDelete(ctx context.Context, token string) (GalleryApplicationsDeletePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("GalleryApplicationsClient.Delete", token, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return GalleryApplicationsDeletePollerResponse{}, err
	}
	poller := &galleryApplicationsDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return GalleryApplicationsDeletePollerResponse{}, err
	}
	result := GalleryApplicationsDeletePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (GalleryApplicationsDeleteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Delete a gallery Application.
// If the operation fails it returns the *CloudError error type.
func (client *GalleryApplicationsClient) deleteOperation(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, options *GalleryApplicationsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, galleryName, galleryApplicationName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *GalleryApplicationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, options *GalleryApplicationsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if galleryApplicationName == "" {
		return nil, errors.New("parameter galleryApplicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationName}", url.PathEscape(galleryApplicationName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-09-30")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *GalleryApplicationsClient) deleteHandleError(resp *azcore.Response) error {
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

// Get - Retrieves information about a gallery Application Definition.
// If the operation fails it returns the *CloudError error type.
func (client *GalleryApplicationsClient) Get(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, options *GalleryApplicationsGetOptions) (GalleryApplicationsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, galleryName, galleryApplicationName, options)
	if err != nil {
		return GalleryApplicationsGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return GalleryApplicationsGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return GalleryApplicationsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *GalleryApplicationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, options *GalleryApplicationsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if galleryApplicationName == "" {
		return nil, errors.New("parameter galleryApplicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationName}", url.PathEscape(galleryApplicationName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-09-30")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GalleryApplicationsClient) getHandleResponse(resp *azcore.Response) (GalleryApplicationsGetResponse, error) {
	result := GalleryApplicationsGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.GalleryApplication); err != nil {
		return GalleryApplicationsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *GalleryApplicationsClient) getHandleError(resp *azcore.Response) error {
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

// ListByGallery - List gallery Application Definitions in a gallery.
// If the operation fails it returns the *CloudError error type.
func (client *GalleryApplicationsClient) ListByGallery(resourceGroupName string, galleryName string, options *GalleryApplicationsListByGalleryOptions) GalleryApplicationsListByGalleryPager {
	return &galleryApplicationsListByGalleryPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByGalleryCreateRequest(ctx, resourceGroupName, galleryName, options)
		},
		advancer: func(ctx context.Context, resp GalleryApplicationsListByGalleryResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.GalleryApplicationList.NextLink)
		},
	}
}

// listByGalleryCreateRequest creates the ListByGallery request.
func (client *GalleryApplicationsClient) listByGalleryCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, options *GalleryApplicationsListByGalleryOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-09-30")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByGalleryHandleResponse handles the ListByGallery response.
func (client *GalleryApplicationsClient) listByGalleryHandleResponse(resp *azcore.Response) (GalleryApplicationsListByGalleryResponse, error) {
	result := GalleryApplicationsListByGalleryResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.GalleryApplicationList); err != nil {
		return GalleryApplicationsListByGalleryResponse{}, err
	}
	return result, nil
}

// listByGalleryHandleError handles the ListByGallery error response.
func (client *GalleryApplicationsClient) listByGalleryHandleError(resp *azcore.Response) error {
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

// BeginUpdate - Update a gallery Application Definition.
// If the operation fails it returns the *CloudError error type.
func (client *GalleryApplicationsClient) BeginUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplication GalleryApplicationUpdate, options *GalleryApplicationsBeginUpdateOptions) (GalleryApplicationsUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplication, options)
	if err != nil {
		return GalleryApplicationsUpdatePollerResponse{}, err
	}
	result := GalleryApplicationsUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("GalleryApplicationsClient.Update", "", resp, client.con.Pipeline(), client.updateHandleError)
	if err != nil {
		return GalleryApplicationsUpdatePollerResponse{}, err
	}
	poller := &galleryApplicationsUpdatePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (GalleryApplicationsUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeUpdate creates a new GalleryApplicationsUpdatePoller from the specified resume token.
// token - The value must come from a previous call to GalleryApplicationsUpdatePoller.ResumeToken().
func (client *GalleryApplicationsClient) ResumeUpdate(ctx context.Context, token string) (GalleryApplicationsUpdatePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("GalleryApplicationsClient.Update", token, client.con.Pipeline(), client.updateHandleError)
	if err != nil {
		return GalleryApplicationsUpdatePollerResponse{}, err
	}
	poller := &galleryApplicationsUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return GalleryApplicationsUpdatePollerResponse{}, err
	}
	result := GalleryApplicationsUpdatePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (GalleryApplicationsUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Update - Update a gallery Application Definition.
// If the operation fails it returns the *CloudError error type.
func (client *GalleryApplicationsClient) update(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplication GalleryApplicationUpdate, options *GalleryApplicationsBeginUpdateOptions) (*azcore.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplication, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *GalleryApplicationsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplication GalleryApplicationUpdate, options *GalleryApplicationsBeginUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if galleryApplicationName == "" {
		return nil, errors.New("parameter galleryApplicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationName}", url.PathEscape(galleryApplicationName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-09-30")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(galleryApplication)
}

// updateHandleError handles the Update error response.
func (client *GalleryApplicationsClient) updateHandleError(resp *azcore.Response) error {
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
