//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsaas

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// ApplicationsClientListResponse contains the response from method ApplicationsClient.List.
type ApplicationsClientListResponse struct {
	ApplicationsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ApplicationsClientListResult contains the result from method ApplicationsClient.List.
type ApplicationsClientListResult struct {
	AppResponseWithContinuation
}

// ClientCreateResourcePollerResponse contains the response from method Client.CreateResource.
type ClientCreateResourcePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ClientCreateResourcePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ClientCreateResourcePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ClientCreateResourceResponse, error) {
	respType := ClientCreateResourceResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Resource)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ClientCreateResourcePollerResponse from the provided client and resume token.
func (l *ClientCreateResourcePollerResponse) Resume(ctx context.Context, client *Client, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("Client.CreateResource", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ClientCreateResourcePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ClientCreateResourceResponse contains the response from method Client.CreateResource.
type ClientCreateResourceResponse struct {
	ClientCreateResourceResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClientCreateResourceResult contains the result from method Client.CreateResource.
type ClientCreateResourceResult struct {
	Resource
}

// ClientDeletePollerResponse contains the response from method Client.Delete.
type ClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ClientDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ClientDeleteResponse, error) {
	respType := ClientDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ClientDeletePollerResponse from the provided client and resume token.
func (l *ClientDeletePollerResponse) Resume(ctx context.Context, client *Client, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("Client.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ClientDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ClientDeleteResponse contains the response from method Client.Delete.
type ClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClientGetResourceResponse contains the response from method Client.GetResource.
type ClientGetResourceResponse struct {
	ClientGetResourceResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClientGetResourceResult contains the result from method Client.GetResource.
type ClientGetResourceResult struct {
	Resource
}

// ClientUpdateResourcePollerResponse contains the response from method Client.UpdateResource.
type ClientUpdateResourcePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ClientUpdateResourcePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ClientUpdateResourcePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ClientUpdateResourceResponse, error) {
	respType := ClientUpdateResourceResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Resource)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ClientUpdateResourcePollerResponse from the provided client and resume token.
func (l *ClientUpdateResourcePollerResponse) Resume(ctx context.Context, client *Client, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("Client.UpdateResource", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ClientUpdateResourcePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ClientUpdateResourceResponse contains the response from method Client.UpdateResource.
type ClientUpdateResourceResponse struct {
	ClientUpdateResourceResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClientUpdateResourceResult contains the result from method Client.UpdateResource.
type ClientUpdateResourceResult struct {
	Resource
}

// OperationClientGetPollerResponse contains the response from method OperationClient.Get.
type OperationClientGetPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *OperationClientGetPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l OperationClientGetPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (OperationClientGetResponse, error) {
	respType := OperationClientGetResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Resource)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a OperationClientGetPollerResponse from the provided client and resume token.
func (l *OperationClientGetPollerResponse) Resume(ctx context.Context, client *OperationClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("OperationClient.Get", token, client.pl)
	if err != nil {
		return err
	}
	poller := &OperationClientGetPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// OperationClientGetResponse contains the response from method OperationClient.Get.
type OperationClientGetResponse struct {
	OperationClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationClientGetResult contains the result from method OperationClient.Get.
type OperationClientGetResult struct {
	Resource
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsClientListResult contains the result from method OperationsClient.List.
type OperationsClientListResult struct {
	AppOperationsResponseWithContinuation
}

// ResourcesClientListAccessTokenResponse contains the response from method ResourcesClient.ListAccessToken.
type ResourcesClientListAccessTokenResponse struct {
	ResourcesClientListAccessTokenResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ResourcesClientListAccessTokenResult contains the result from method ResourcesClient.ListAccessToken.
type ResourcesClientListAccessTokenResult struct {
	AccessTokenResult
}

// ResourcesClientListResponse contains the response from method ResourcesClient.List.
type ResourcesClientListResponse struct {
	ResourcesClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ResourcesClientListResult contains the result from method ResourcesClient.List.
type ResourcesClientListResult struct {
	ResourceResponseWithContinuation
}

// SubscriptionLevelClientCreateOrUpdatePollerResponse contains the response from method SubscriptionLevelClient.CreateOrUpdate.
type SubscriptionLevelClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *SubscriptionLevelClientCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l SubscriptionLevelClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (SubscriptionLevelClientCreateOrUpdateResponse, error) {
	respType := SubscriptionLevelClientCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Resource)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a SubscriptionLevelClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *SubscriptionLevelClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *SubscriptionLevelClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("SubscriptionLevelClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &SubscriptionLevelClientCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// SubscriptionLevelClientCreateOrUpdateResponse contains the response from method SubscriptionLevelClient.CreateOrUpdate.
type SubscriptionLevelClientCreateOrUpdateResponse struct {
	SubscriptionLevelClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionLevelClientCreateOrUpdateResult contains the result from method SubscriptionLevelClient.CreateOrUpdate.
type SubscriptionLevelClientCreateOrUpdateResult struct {
	Resource
}

// SubscriptionLevelClientDeletePollerResponse contains the response from method SubscriptionLevelClient.Delete.
type SubscriptionLevelClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *SubscriptionLevelClientDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l SubscriptionLevelClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (SubscriptionLevelClientDeleteResponse, error) {
	respType := SubscriptionLevelClientDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a SubscriptionLevelClientDeletePollerResponse from the provided client and resume token.
func (l *SubscriptionLevelClientDeletePollerResponse) Resume(ctx context.Context, client *SubscriptionLevelClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("SubscriptionLevelClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &SubscriptionLevelClientDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// SubscriptionLevelClientDeleteResponse contains the response from method SubscriptionLevelClient.Delete.
type SubscriptionLevelClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionLevelClientGetResponse contains the response from method SubscriptionLevelClient.Get.
type SubscriptionLevelClientGetResponse struct {
	SubscriptionLevelClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionLevelClientGetResult contains the result from method SubscriptionLevelClient.Get.
type SubscriptionLevelClientGetResult struct {
	Resource
}

// SubscriptionLevelClientListAccessTokenResponse contains the response from method SubscriptionLevelClient.ListAccessToken.
type SubscriptionLevelClientListAccessTokenResponse struct {
	SubscriptionLevelClientListAccessTokenResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionLevelClientListAccessTokenResult contains the result from method SubscriptionLevelClient.ListAccessToken.
type SubscriptionLevelClientListAccessTokenResult struct {
	AccessTokenResult
}

// SubscriptionLevelClientListByAzureSubscriptionResponse contains the response from method SubscriptionLevelClient.ListByAzureSubscription.
type SubscriptionLevelClientListByAzureSubscriptionResponse struct {
	SubscriptionLevelClientListByAzureSubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionLevelClientListByAzureSubscriptionResult contains the result from method SubscriptionLevelClient.ListByAzureSubscription.
type SubscriptionLevelClientListByAzureSubscriptionResult struct {
	ResourceResponseWithContinuation
}

// SubscriptionLevelClientListByResourceGroupResponse contains the response from method SubscriptionLevelClient.ListByResourceGroup.
type SubscriptionLevelClientListByResourceGroupResponse struct {
	SubscriptionLevelClientListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionLevelClientListByResourceGroupResult contains the result from method SubscriptionLevelClient.ListByResourceGroup.
type SubscriptionLevelClientListByResourceGroupResult struct {
	ResourceResponseWithContinuation
}

// SubscriptionLevelClientMoveResourcesPollerResponse contains the response from method SubscriptionLevelClient.MoveResources.
type SubscriptionLevelClientMoveResourcesPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *SubscriptionLevelClientMoveResourcesPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l SubscriptionLevelClientMoveResourcesPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (SubscriptionLevelClientMoveResourcesResponse, error) {
	respType := SubscriptionLevelClientMoveResourcesResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a SubscriptionLevelClientMoveResourcesPollerResponse from the provided client and resume token.
func (l *SubscriptionLevelClientMoveResourcesPollerResponse) Resume(ctx context.Context, client *SubscriptionLevelClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("SubscriptionLevelClient.MoveResources", token, client.pl)
	if err != nil {
		return err
	}
	poller := &SubscriptionLevelClientMoveResourcesPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// SubscriptionLevelClientMoveResourcesResponse contains the response from method SubscriptionLevelClient.MoveResources.
type SubscriptionLevelClientMoveResourcesResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionLevelClientUpdatePollerResponse contains the response from method SubscriptionLevelClient.Update.
type SubscriptionLevelClientUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *SubscriptionLevelClientUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l SubscriptionLevelClientUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (SubscriptionLevelClientUpdateResponse, error) {
	respType := SubscriptionLevelClientUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Resource)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a SubscriptionLevelClientUpdatePollerResponse from the provided client and resume token.
func (l *SubscriptionLevelClientUpdatePollerResponse) Resume(ctx context.Context, client *SubscriptionLevelClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("SubscriptionLevelClient.Update", token, client.pl)
	if err != nil {
		return err
	}
	poller := &SubscriptionLevelClientUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// SubscriptionLevelClientUpdateResponse contains the response from method SubscriptionLevelClient.Update.
type SubscriptionLevelClientUpdateResponse struct {
	SubscriptionLevelClientUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionLevelClientUpdateResult contains the result from method SubscriptionLevelClient.Update.
type SubscriptionLevelClientUpdateResult struct {
	Resource
}

// SubscriptionLevelClientUpdateToUnsubscribedPollerResponse contains the response from method SubscriptionLevelClient.UpdateToUnsubscribed.
type SubscriptionLevelClientUpdateToUnsubscribedPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *SubscriptionLevelClientUpdateToUnsubscribedPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l SubscriptionLevelClientUpdateToUnsubscribedPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (SubscriptionLevelClientUpdateToUnsubscribedResponse, error) {
	respType := SubscriptionLevelClientUpdateToUnsubscribedResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a SubscriptionLevelClientUpdateToUnsubscribedPollerResponse from the provided client and resume token.
func (l *SubscriptionLevelClientUpdateToUnsubscribedPollerResponse) Resume(ctx context.Context, client *SubscriptionLevelClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("SubscriptionLevelClient.UpdateToUnsubscribed", token, client.pl)
	if err != nil {
		return err
	}
	poller := &SubscriptionLevelClientUpdateToUnsubscribedPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// SubscriptionLevelClientUpdateToUnsubscribedResponse contains the response from method SubscriptionLevelClient.UpdateToUnsubscribed.
type SubscriptionLevelClientUpdateToUnsubscribedResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionLevelClientValidateMoveResourcesResponse contains the response from method SubscriptionLevelClient.ValidateMoveResources.
type SubscriptionLevelClientValidateMoveResourcesResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}
