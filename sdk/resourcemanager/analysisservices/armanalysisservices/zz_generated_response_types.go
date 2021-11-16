//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armanalysisservices

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// OperationsListResponse contains the response from method Operations.List.
type OperationsListResponse struct {
	OperationsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsListResult contains the result from method Operations.List.
type OperationsListResult struct {
	OperationListResult
}

// ServersCheckNameAvailabilityResponse contains the response from method Servers.CheckNameAvailability.
type ServersCheckNameAvailabilityResponse struct {
	ServersCheckNameAvailabilityResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersCheckNameAvailabilityResult contains the result from method Servers.CheckNameAvailability.
type ServersCheckNameAvailabilityResult struct {
	CheckServerNameAvailabilityResult
}

// ServersCreatePollerResponse contains the response from method Servers.Create.
type ServersCreatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ServersCreatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ServersCreatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ServersCreateResponse, error) {
	respType := ServersCreateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.AnalysisServicesServer)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ServersCreatePollerResponse from the provided client and resume token.
func (l *ServersCreatePollerResponse) Resume(ctx context.Context, client *ServersClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ServersClient.Create", token, client.pl, client.createHandleError)
	if err != nil {
		return err
	}
	poller := &ServersCreatePoller{
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

// ServersCreateResponse contains the response from method Servers.Create.
type ServersCreateResponse struct {
	ServersCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersCreateResult contains the result from method Servers.Create.
type ServersCreateResult struct {
	AnalysisServicesServer
}

// ServersDeletePollerResponse contains the response from method Servers.Delete.
type ServersDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ServersDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ServersDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ServersDeleteResponse, error) {
	respType := ServersDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ServersDeletePollerResponse from the provided client and resume token.
func (l *ServersDeletePollerResponse) Resume(ctx context.Context, client *ServersClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ServersClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &ServersDeletePoller{
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

// ServersDeleteResponse contains the response from method Servers.Delete.
type ServersDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersDissociateGatewayResponse contains the response from method Servers.DissociateGateway.
type ServersDissociateGatewayResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersGetDetailsResponse contains the response from method Servers.GetDetails.
type ServersGetDetailsResponse struct {
	ServersGetDetailsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersGetDetailsResult contains the result from method Servers.GetDetails.
type ServersGetDetailsResult struct {
	AnalysisServicesServer
}

// ServersListByResourceGroupResponse contains the response from method Servers.ListByResourceGroup.
type ServersListByResourceGroupResponse struct {
	ServersListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersListByResourceGroupResult contains the result from method Servers.ListByResourceGroup.
type ServersListByResourceGroupResult struct {
	AnalysisServicesServers
}

// ServersListGatewayStatusResponse contains the response from method Servers.ListGatewayStatus.
type ServersListGatewayStatusResponse struct {
	ServersListGatewayStatusResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersListGatewayStatusResult contains the result from method Servers.ListGatewayStatus.
type ServersListGatewayStatusResult struct {
	GatewayListStatusLive
}

// ServersListOperationResultsResponse contains the response from method Servers.ListOperationResults.
type ServersListOperationResultsResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersListOperationStatusesResponse contains the response from method Servers.ListOperationStatuses.
type ServersListOperationStatusesResponse struct {
	ServersListOperationStatusesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersListOperationStatusesResult contains the result from method Servers.ListOperationStatuses.
type ServersListOperationStatusesResult struct {
	OperationStatus
}

// ServersListResponse contains the response from method Servers.List.
type ServersListResponse struct {
	ServersListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersListResult contains the result from method Servers.List.
type ServersListResult struct {
	AnalysisServicesServers
}

// ServersListSKUsForExistingResponse contains the response from method Servers.ListSKUsForExisting.
type ServersListSKUsForExistingResponse struct {
	ServersListSKUsForExistingResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersListSKUsForExistingResult contains the result from method Servers.ListSKUsForExisting.
type ServersListSKUsForExistingResult struct {
	SKUEnumerationForExistingResourceResult
}

// ServersListSKUsForNewResponse contains the response from method Servers.ListSKUsForNew.
type ServersListSKUsForNewResponse struct {
	ServersListSKUsForNewResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersListSKUsForNewResult contains the result from method Servers.ListSKUsForNew.
type ServersListSKUsForNewResult struct {
	SKUEnumerationForNewResourceResult
}

// ServersResumePollerResponse contains the response from method Servers.Resume.
type ServersResumePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ServersResumePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ServersResumePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ServersResumeResponse, error) {
	respType := ServersResumeResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ServersResumePollerResponse from the provided client and resume token.
func (l *ServersResumePollerResponse) Resume(ctx context.Context, client *ServersClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ServersClient.Resume", token, client.pl, client.resumeHandleError)
	if err != nil {
		return err
	}
	poller := &ServersResumePoller{
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

// ServersResumeResponse contains the response from method Servers.Resume.
type ServersResumeResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersSuspendPollerResponse contains the response from method Servers.Suspend.
type ServersSuspendPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ServersSuspendPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ServersSuspendPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ServersSuspendResponse, error) {
	respType := ServersSuspendResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ServersSuspendPollerResponse from the provided client and resume token.
func (l *ServersSuspendPollerResponse) Resume(ctx context.Context, client *ServersClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ServersClient.Suspend", token, client.pl, client.suspendHandleError)
	if err != nil {
		return err
	}
	poller := &ServersSuspendPoller{
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

// ServersSuspendResponse contains the response from method Servers.Suspend.
type ServersSuspendResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersUpdatePollerResponse contains the response from method Servers.Update.
type ServersUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ServersUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ServersUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ServersUpdateResponse, error) {
	respType := ServersUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.AnalysisServicesServer)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ServersUpdatePollerResponse from the provided client and resume token.
func (l *ServersUpdatePollerResponse) Resume(ctx context.Context, client *ServersClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ServersClient.Update", token, client.pl, client.updateHandleError)
	if err != nil {
		return err
	}
	poller := &ServersUpdatePoller{
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

// ServersUpdateResponse contains the response from method Servers.Update.
type ServersUpdateResponse struct {
	ServersUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersUpdateResult contains the result from method Servers.Update.
type ServersUpdateResult struct {
	AnalysisServicesServer
}
