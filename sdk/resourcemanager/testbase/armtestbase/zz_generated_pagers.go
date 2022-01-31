//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtestbase

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// AccountsClientListByResourceGroupPager provides operations for iterating over paged responses.
type AccountsClientListByResourceGroupPager struct {
	client    *AccountsClient
	current   AccountsClientListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AccountsClientListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AccountsClientListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AccountsClientListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AccountListResult.NextLink == nil || len(*p.current.AccountListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current AccountsClientListByResourceGroupResponse page.
func (p *AccountsClientListByResourceGroupPager) PageResponse() AccountsClientListByResourceGroupResponse {
	return p.current
}

// AccountsClientListBySubscriptionPager provides operations for iterating over paged responses.
type AccountsClientListBySubscriptionPager struct {
	client    *AccountsClient
	current   AccountsClientListBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AccountsClientListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AccountsClientListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AccountsClientListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AccountListResult.NextLink == nil || len(*p.current.AccountListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listBySubscriptionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current AccountsClientListBySubscriptionResponse page.
func (p *AccountsClientListBySubscriptionPager) PageResponse() AccountsClientListBySubscriptionResponse {
	return p.current
}

// AvailableOSClientListPager provides operations for iterating over paged responses.
type AvailableOSClientListPager struct {
	client    *AvailableOSClient
	current   AvailableOSClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AvailableOSClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AvailableOSClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AvailableOSClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AvailableOSListResult.NextLink == nil || len(*p.current.AvailableOSListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current AvailableOSClientListResponse page.
func (p *AvailableOSClientListPager) PageResponse() AvailableOSClientListResponse {
	return p.current
}

// CustomerEventsClientListByTestBaseAccountPager provides operations for iterating over paged responses.
type CustomerEventsClientListByTestBaseAccountPager struct {
	client    *CustomerEventsClient
	current   CustomerEventsClientListByTestBaseAccountResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, CustomerEventsClientListByTestBaseAccountResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *CustomerEventsClientListByTestBaseAccountPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *CustomerEventsClientListByTestBaseAccountPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.CustomerEventListResult.NextLink == nil || len(*p.current.CustomerEventListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByTestBaseAccountHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current CustomerEventsClientListByTestBaseAccountResponse page.
func (p *CustomerEventsClientListByTestBaseAccountPager) PageResponse() CustomerEventsClientListByTestBaseAccountResponse {
	return p.current
}

// EmailEventsClientListPager provides operations for iterating over paged responses.
type EmailEventsClientListPager struct {
	client    *EmailEventsClient
	current   EmailEventsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, EmailEventsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *EmailEventsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *EmailEventsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.EmailEventListResult.NextLink == nil || len(*p.current.EmailEventListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current EmailEventsClientListResponse page.
func (p *EmailEventsClientListPager) PageResponse() EmailEventsClientListResponse {
	return p.current
}

// FavoriteProcessesClientListPager provides operations for iterating over paged responses.
type FavoriteProcessesClientListPager struct {
	client    *FavoriteProcessesClient
	current   FavoriteProcessesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, FavoriteProcessesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *FavoriteProcessesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *FavoriteProcessesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.FavoriteProcessListResult.NextLink == nil || len(*p.current.FavoriteProcessListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current FavoriteProcessesClientListResponse page.
func (p *FavoriteProcessesClientListPager) PageResponse() FavoriteProcessesClientListResponse {
	return p.current
}

// FlightingRingsClientListPager provides operations for iterating over paged responses.
type FlightingRingsClientListPager struct {
	client    *FlightingRingsClient
	current   FlightingRingsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, FlightingRingsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *FlightingRingsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *FlightingRingsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.FlightingRingListResult.NextLink == nil || len(*p.current.FlightingRingListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current FlightingRingsClientListResponse page.
func (p *FlightingRingsClientListPager) PageResponse() FlightingRingsClientListResponse {
	return p.current
}

// OSUpdatesClientListPager provides operations for iterating over paged responses.
type OSUpdatesClientListPager struct {
	client    *OSUpdatesClient
	current   OSUpdatesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, OSUpdatesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *OSUpdatesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *OSUpdatesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OSUpdateListResult.NextLink == nil || len(*p.current.OSUpdateListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current OSUpdatesClientListResponse page.
func (p *OSUpdatesClientListPager) PageResponse() OSUpdatesClientListResponse {
	return p.current
}

// OperationsClientListPager provides operations for iterating over paged responses.
type OperationsClientListPager struct {
	client    *OperationsClient
	current   OperationsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, OperationsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *OperationsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *OperationsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OperationListResult.NextLink == nil || len(*p.current.OperationListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current OperationsClientListResponse page.
func (p *OperationsClientListPager) PageResponse() OperationsClientListResponse {
	return p.current
}

// PackagesClientListByTestBaseAccountPager provides operations for iterating over paged responses.
type PackagesClientListByTestBaseAccountPager struct {
	client    *PackagesClient
	current   PackagesClientListByTestBaseAccountResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PackagesClientListByTestBaseAccountResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PackagesClientListByTestBaseAccountPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PackagesClientListByTestBaseAccountPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PackageListResult.NextLink == nil || len(*p.current.PackageListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByTestBaseAccountHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PackagesClientListByTestBaseAccountResponse page.
func (p *PackagesClientListByTestBaseAccountPager) PageResponse() PackagesClientListByTestBaseAccountResponse {
	return p.current
}

// SKUsClientListPager provides operations for iterating over paged responses.
type SKUsClientListPager struct {
	client    *SKUsClient
	current   SKUsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SKUsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SKUsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SKUsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AccountSKUListResult.NextLink == nil || len(*p.current.AccountSKUListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current SKUsClientListResponse page.
func (p *SKUsClientListPager) PageResponse() SKUsClientListResponse {
	return p.current
}

// TestResultsClientListPager provides operations for iterating over paged responses.
type TestResultsClientListPager struct {
	client    *TestResultsClient
	current   TestResultsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, TestResultsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *TestResultsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *TestResultsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.TestResultListResult.NextLink == nil || len(*p.current.TestResultListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current TestResultsClientListResponse page.
func (p *TestResultsClientListPager) PageResponse() TestResultsClientListResponse {
	return p.current
}

// TestSummariesClientListPager provides operations for iterating over paged responses.
type TestSummariesClientListPager struct {
	client    *TestSummariesClient
	current   TestSummariesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, TestSummariesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *TestSummariesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *TestSummariesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.TestSummaryListResult.NextLink == nil || len(*p.current.TestSummaryListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current TestSummariesClientListResponse page.
func (p *TestSummariesClientListPager) PageResponse() TestSummariesClientListResponse {
	return p.current
}

// TestTypesClientListPager provides operations for iterating over paged responses.
type TestTypesClientListPager struct {
	client    *TestTypesClient
	current   TestTypesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, TestTypesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *TestTypesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *TestTypesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.TestTypeListResult.NextLink == nil || len(*p.current.TestTypeListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current TestTypesClientListResponse page.
func (p *TestTypesClientListPager) PageResponse() TestTypesClientListResponse {
	return p.current
}

// UsageClientListPager provides operations for iterating over paged responses.
type UsageClientListPager struct {
	client    *UsageClient
	current   UsageClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, UsageClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *UsageClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *UsageClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AccountUsageDataList.NextLink == nil || len(*p.current.AccountUsageDataList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current UsageClientListResponse page.
func (p *UsageClientListPager) PageResponse() UsageClientListResponse {
	return p.current
}
