//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armreservations

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// CalculateExchangeClientPostPoller provides polling facilities until the operation reaches a terminal state.
type CalculateExchangeClientPostPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *CalculateExchangeClientPostPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *CalculateExchangeClientPostPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final CalculateExchangeClientPostResponse will be returned.
func (p *CalculateExchangeClientPostPoller) FinalResponse(ctx context.Context) (CalculateExchangeClientPostResponse, error) {
	respType := CalculateExchangeClientPostResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.CalculateExchangeOperationResultResponse)
	if err != nil {
		return CalculateExchangeClientPostResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *CalculateExchangeClientPostPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ExchangeClientPostPoller provides polling facilities until the operation reaches a terminal state.
type ExchangeClientPostPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ExchangeClientPostPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ExchangeClientPostPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ExchangeClientPostResponse will be returned.
func (p *ExchangeClientPostPoller) FinalResponse(ctx context.Context) (ExchangeClientPostResponse, error) {
	respType := ExchangeClientPostResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.ExchangeOperationResultResponse)
	if err != nil {
		return ExchangeClientPostResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ExchangeClientPostPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// QuotaClientCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type QuotaClientCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *QuotaClientCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *QuotaClientCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final QuotaClientCreateOrUpdateResponse will be returned.
func (p *QuotaClientCreateOrUpdatePoller) FinalResponse(ctx context.Context) (QuotaClientCreateOrUpdateResponse, error) {
	respType := QuotaClientCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.QuotaRequestOneResourceSubmitResponse)
	if err != nil {
		return QuotaClientCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *QuotaClientCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// QuotaClientUpdatePoller provides polling facilities until the operation reaches a terminal state.
type QuotaClientUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *QuotaClientUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *QuotaClientUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final QuotaClientUpdateResponse will be returned.
func (p *QuotaClientUpdatePoller) FinalResponse(ctx context.Context) (QuotaClientUpdateResponse, error) {
	respType := QuotaClientUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.QuotaRequestOneResourceSubmitResponse)
	if err != nil {
		return QuotaClientUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *QuotaClientUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ReservationClientAvailableScopesPoller provides polling facilities until the operation reaches a terminal state.
type ReservationClientAvailableScopesPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ReservationClientAvailableScopesPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ReservationClientAvailableScopesPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ReservationClientAvailableScopesResponse will be returned.
func (p *ReservationClientAvailableScopesPoller) FinalResponse(ctx context.Context) (ReservationClientAvailableScopesResponse, error) {
	respType := ReservationClientAvailableScopesResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.AvailableScopeProperties)
	if err != nil {
		return ReservationClientAvailableScopesResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ReservationClientAvailableScopesPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ReservationClientMergePoller provides polling facilities until the operation reaches a terminal state.
type ReservationClientMergePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ReservationClientMergePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ReservationClientMergePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ReservationClientMergeResponse will be returned.
func (p *ReservationClientMergePoller) FinalResponse(ctx context.Context) (ReservationClientMergeResponse, error) {
	respType := ReservationClientMergeResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.ReservationResponseArray)
	if err != nil {
		return ReservationClientMergeResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ReservationClientMergePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ReservationClientSplitPoller provides polling facilities until the operation reaches a terminal state.
type ReservationClientSplitPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ReservationClientSplitPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ReservationClientSplitPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ReservationClientSplitResponse will be returned.
func (p *ReservationClientSplitPoller) FinalResponse(ctx context.Context) (ReservationClientSplitResponse, error) {
	respType := ReservationClientSplitResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.ReservationResponseArray)
	if err != nil {
		return ReservationClientSplitResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ReservationClientSplitPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ReservationClientUpdatePoller provides polling facilities until the operation reaches a terminal state.
type ReservationClientUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ReservationClientUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ReservationClientUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ReservationClientUpdateResponse will be returned.
func (p *ReservationClientUpdatePoller) FinalResponse(ctx context.Context) (ReservationClientUpdateResponse, error) {
	respType := ReservationClientUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.ReservationResponse)
	if err != nil {
		return ReservationClientUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ReservationClientUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ReservationOrderClientPurchasePoller provides polling facilities until the operation reaches a terminal state.
type ReservationOrderClientPurchasePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ReservationOrderClientPurchasePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ReservationOrderClientPurchasePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ReservationOrderClientPurchaseResponse will be returned.
func (p *ReservationOrderClientPurchasePoller) FinalResponse(ctx context.Context) (ReservationOrderClientPurchaseResponse, error) {
	respType := ReservationOrderClientPurchaseResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.ReservationOrderResponse)
	if err != nil {
		return ReservationOrderClientPurchaseResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ReservationOrderClientPurchasePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}
