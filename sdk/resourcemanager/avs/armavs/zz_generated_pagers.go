//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armavs

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// AddonsListPager provides operations for iterating over paged responses.
type AddonsListPager struct {
	client    *AddonsClient
	current   AddonsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AddonsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AddonsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AddonsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AddonList.NextLink == nil || len(*p.current.AddonList.NextLink) == 0 {
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
		p.err = p.client.listHandleError(resp)
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

// PageResponse returns the current AddonsListResponse page.
func (p *AddonsListPager) PageResponse() AddonsListResponse {
	return p.current
}

// AuthorizationsListPager provides operations for iterating over paged responses.
type AuthorizationsListPager struct {
	client    *AuthorizationsClient
	current   AuthorizationsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AuthorizationsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AuthorizationsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AuthorizationsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ExpressRouteAuthorizationList.NextLink == nil || len(*p.current.ExpressRouteAuthorizationList.NextLink) == 0 {
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
		p.err = p.client.listHandleError(resp)
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

// PageResponse returns the current AuthorizationsListResponse page.
func (p *AuthorizationsListPager) PageResponse() AuthorizationsListResponse {
	return p.current
}

// CloudLinksListPager provides operations for iterating over paged responses.
type CloudLinksListPager struct {
	client    *CloudLinksClient
	current   CloudLinksListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, CloudLinksListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *CloudLinksListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *CloudLinksListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.CloudLinkList.NextLink == nil || len(*p.current.CloudLinkList.NextLink) == 0 {
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
		p.err = p.client.listHandleError(resp)
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

// PageResponse returns the current CloudLinksListResponse page.
func (p *CloudLinksListPager) PageResponse() CloudLinksListResponse {
	return p.current
}

// ClustersListPager provides operations for iterating over paged responses.
type ClustersListPager struct {
	client    *ClustersClient
	current   ClustersListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClustersListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ClustersListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ClustersListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ClusterList.NextLink == nil || len(*p.current.ClusterList.NextLink) == 0 {
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
		p.err = p.client.listHandleError(resp)
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

// PageResponse returns the current ClustersListResponse page.
func (p *ClustersListPager) PageResponse() ClustersListResponse {
	return p.current
}

// DatastoresListPager provides operations for iterating over paged responses.
type DatastoresListPager struct {
	client    *DatastoresClient
	current   DatastoresListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DatastoresListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DatastoresListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DatastoresListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DatastoreList.NextLink == nil || len(*p.current.DatastoreList.NextLink) == 0 {
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
		p.err = p.client.listHandleError(resp)
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

// PageResponse returns the current DatastoresListResponse page.
func (p *DatastoresListPager) PageResponse() DatastoresListResponse {
	return p.current
}

// GlobalReachConnectionsListPager provides operations for iterating over paged responses.
type GlobalReachConnectionsListPager struct {
	client    *GlobalReachConnectionsClient
	current   GlobalReachConnectionsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, GlobalReachConnectionsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *GlobalReachConnectionsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *GlobalReachConnectionsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.GlobalReachConnectionList.NextLink == nil || len(*p.current.GlobalReachConnectionList.NextLink) == 0 {
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
		p.err = p.client.listHandleError(resp)
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

// PageResponse returns the current GlobalReachConnectionsListResponse page.
func (p *GlobalReachConnectionsListPager) PageResponse() GlobalReachConnectionsListResponse {
	return p.current
}

// HcxEnterpriseSitesListPager provides operations for iterating over paged responses.
type HcxEnterpriseSitesListPager struct {
	client    *HcxEnterpriseSitesClient
	current   HcxEnterpriseSitesListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, HcxEnterpriseSitesListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *HcxEnterpriseSitesListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *HcxEnterpriseSitesListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.HcxEnterpriseSiteList.NextLink == nil || len(*p.current.HcxEnterpriseSiteList.NextLink) == 0 {
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
		p.err = p.client.listHandleError(resp)
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

// PageResponse returns the current HcxEnterpriseSitesListResponse page.
func (p *HcxEnterpriseSitesListPager) PageResponse() HcxEnterpriseSitesListResponse {
	return p.current
}

// OperationsListPager provides operations for iterating over paged responses.
type OperationsListPager struct {
	client    *OperationsClient
	current   OperationsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, OperationsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *OperationsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *OperationsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OperationList.NextLink == nil || len(*p.current.OperationList.NextLink) == 0 {
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
		p.err = p.client.listHandleError(resp)
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

// PageResponse returns the current OperationsListResponse page.
func (p *OperationsListPager) PageResponse() OperationsListResponse {
	return p.current
}

// PlacementPoliciesListPager provides operations for iterating over paged responses.
type PlacementPoliciesListPager struct {
	client    *PlacementPoliciesClient
	current   PlacementPoliciesListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PlacementPoliciesListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PlacementPoliciesListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PlacementPoliciesListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PlacementPoliciesList.NextLink == nil || len(*p.current.PlacementPoliciesList.NextLink) == 0 {
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
		p.err = p.client.listHandleError(resp)
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

// PageResponse returns the current PlacementPoliciesListResponse page.
func (p *PlacementPoliciesListPager) PageResponse() PlacementPoliciesListResponse {
	return p.current
}

// PrivateCloudsListInSubscriptionPager provides operations for iterating over paged responses.
type PrivateCloudsListInSubscriptionPager struct {
	client    *PrivateCloudsClient
	current   PrivateCloudsListInSubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PrivateCloudsListInSubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PrivateCloudsListInSubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PrivateCloudsListInSubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PrivateCloudList.NextLink == nil || len(*p.current.PrivateCloudList.NextLink) == 0 {
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
		p.err = p.client.listInSubscriptionHandleError(resp)
		return false
	}
	result, err := p.client.listInSubscriptionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PrivateCloudsListInSubscriptionResponse page.
func (p *PrivateCloudsListInSubscriptionPager) PageResponse() PrivateCloudsListInSubscriptionResponse {
	return p.current
}

// PrivateCloudsListPager provides operations for iterating over paged responses.
type PrivateCloudsListPager struct {
	client    *PrivateCloudsClient
	current   PrivateCloudsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PrivateCloudsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PrivateCloudsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PrivateCloudsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PrivateCloudList.NextLink == nil || len(*p.current.PrivateCloudList.NextLink) == 0 {
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
		p.err = p.client.listHandleError(resp)
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

// PageResponse returns the current PrivateCloudsListResponse page.
func (p *PrivateCloudsListPager) PageResponse() PrivateCloudsListResponse {
	return p.current
}

// ScriptCmdletsListPager provides operations for iterating over paged responses.
type ScriptCmdletsListPager struct {
	client    *ScriptCmdletsClient
	current   ScriptCmdletsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ScriptCmdletsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ScriptCmdletsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ScriptCmdletsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ScriptCmdletsList.NextLink == nil || len(*p.current.ScriptCmdletsList.NextLink) == 0 {
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
		p.err = p.client.listHandleError(resp)
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

// PageResponse returns the current ScriptCmdletsListResponse page.
func (p *ScriptCmdletsListPager) PageResponse() ScriptCmdletsListResponse {
	return p.current
}

// ScriptExecutionsListPager provides operations for iterating over paged responses.
type ScriptExecutionsListPager struct {
	client    *ScriptExecutionsClient
	current   ScriptExecutionsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ScriptExecutionsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ScriptExecutionsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ScriptExecutionsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ScriptExecutionsList.NextLink == nil || len(*p.current.ScriptExecutionsList.NextLink) == 0 {
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
		p.err = p.client.listHandleError(resp)
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

// PageResponse returns the current ScriptExecutionsListResponse page.
func (p *ScriptExecutionsListPager) PageResponse() ScriptExecutionsListResponse {
	return p.current
}

// ScriptPackagesListPager provides operations for iterating over paged responses.
type ScriptPackagesListPager struct {
	client    *ScriptPackagesClient
	current   ScriptPackagesListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ScriptPackagesListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ScriptPackagesListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ScriptPackagesListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ScriptPackagesList.NextLink == nil || len(*p.current.ScriptPackagesList.NextLink) == 0 {
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
		p.err = p.client.listHandleError(resp)
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

// PageResponse returns the current ScriptPackagesListResponse page.
func (p *ScriptPackagesListPager) PageResponse() ScriptPackagesListResponse {
	return p.current
}

// VirtualMachinesListPager provides operations for iterating over paged responses.
type VirtualMachinesListPager struct {
	client    *VirtualMachinesClient
	current   VirtualMachinesListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, VirtualMachinesListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *VirtualMachinesListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *VirtualMachinesListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.VirtualMachinesList.NextLink == nil || len(*p.current.VirtualMachinesList.NextLink) == 0 {
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
		p.err = p.client.listHandleError(resp)
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

// PageResponse returns the current VirtualMachinesListResponse page.
func (p *VirtualMachinesListPager) PageResponse() VirtualMachinesListResponse {
	return p.current
}

// WorkloadNetworksListDNSServicesPager provides operations for iterating over paged responses.
type WorkloadNetworksListDNSServicesPager struct {
	client    *WorkloadNetworksClient
	current   WorkloadNetworksListDNSServicesResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WorkloadNetworksListDNSServicesResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WorkloadNetworksListDNSServicesPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WorkloadNetworksListDNSServicesPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WorkloadNetworkDNSServicesList.NextLink == nil || len(*p.current.WorkloadNetworkDNSServicesList.NextLink) == 0 {
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
		p.err = p.client.listDNSServicesHandleError(resp)
		return false
	}
	result, err := p.client.listDNSServicesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current WorkloadNetworksListDNSServicesResponse page.
func (p *WorkloadNetworksListDNSServicesPager) PageResponse() WorkloadNetworksListDNSServicesResponse {
	return p.current
}

// WorkloadNetworksListDNSZonesPager provides operations for iterating over paged responses.
type WorkloadNetworksListDNSZonesPager struct {
	client    *WorkloadNetworksClient
	current   WorkloadNetworksListDNSZonesResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WorkloadNetworksListDNSZonesResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WorkloadNetworksListDNSZonesPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WorkloadNetworksListDNSZonesPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WorkloadNetworkDNSZonesList.NextLink == nil || len(*p.current.WorkloadNetworkDNSZonesList.NextLink) == 0 {
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
		p.err = p.client.listDNSZonesHandleError(resp)
		return false
	}
	result, err := p.client.listDNSZonesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current WorkloadNetworksListDNSZonesResponse page.
func (p *WorkloadNetworksListDNSZonesPager) PageResponse() WorkloadNetworksListDNSZonesResponse {
	return p.current
}

// WorkloadNetworksListDhcpPager provides operations for iterating over paged responses.
type WorkloadNetworksListDhcpPager struct {
	client    *WorkloadNetworksClient
	current   WorkloadNetworksListDhcpResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WorkloadNetworksListDhcpResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WorkloadNetworksListDhcpPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WorkloadNetworksListDhcpPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WorkloadNetworkDhcpList.NextLink == nil || len(*p.current.WorkloadNetworkDhcpList.NextLink) == 0 {
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
		p.err = p.client.listDhcpHandleError(resp)
		return false
	}
	result, err := p.client.listDhcpHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current WorkloadNetworksListDhcpResponse page.
func (p *WorkloadNetworksListDhcpPager) PageResponse() WorkloadNetworksListDhcpResponse {
	return p.current
}

// WorkloadNetworksListGatewaysPager provides operations for iterating over paged responses.
type WorkloadNetworksListGatewaysPager struct {
	client    *WorkloadNetworksClient
	current   WorkloadNetworksListGatewaysResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WorkloadNetworksListGatewaysResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WorkloadNetworksListGatewaysPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WorkloadNetworksListGatewaysPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WorkloadNetworkGatewayList.NextLink == nil || len(*p.current.WorkloadNetworkGatewayList.NextLink) == 0 {
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
		p.err = p.client.listGatewaysHandleError(resp)
		return false
	}
	result, err := p.client.listGatewaysHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current WorkloadNetworksListGatewaysResponse page.
func (p *WorkloadNetworksListGatewaysPager) PageResponse() WorkloadNetworksListGatewaysResponse {
	return p.current
}

// WorkloadNetworksListPortMirroringPager provides operations for iterating over paged responses.
type WorkloadNetworksListPortMirroringPager struct {
	client    *WorkloadNetworksClient
	current   WorkloadNetworksListPortMirroringResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WorkloadNetworksListPortMirroringResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WorkloadNetworksListPortMirroringPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WorkloadNetworksListPortMirroringPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WorkloadNetworkPortMirroringList.NextLink == nil || len(*p.current.WorkloadNetworkPortMirroringList.NextLink) == 0 {
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
		p.err = p.client.listPortMirroringHandleError(resp)
		return false
	}
	result, err := p.client.listPortMirroringHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current WorkloadNetworksListPortMirroringResponse page.
func (p *WorkloadNetworksListPortMirroringPager) PageResponse() WorkloadNetworksListPortMirroringResponse {
	return p.current
}

// WorkloadNetworksListPublicIPsPager provides operations for iterating over paged responses.
type WorkloadNetworksListPublicIPsPager struct {
	client    *WorkloadNetworksClient
	current   WorkloadNetworksListPublicIPsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WorkloadNetworksListPublicIPsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WorkloadNetworksListPublicIPsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WorkloadNetworksListPublicIPsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WorkloadNetworkPublicIPsList.NextLink == nil || len(*p.current.WorkloadNetworkPublicIPsList.NextLink) == 0 {
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
		p.err = p.client.listPublicIPsHandleError(resp)
		return false
	}
	result, err := p.client.listPublicIPsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current WorkloadNetworksListPublicIPsResponse page.
func (p *WorkloadNetworksListPublicIPsPager) PageResponse() WorkloadNetworksListPublicIPsResponse {
	return p.current
}

// WorkloadNetworksListSegmentsPager provides operations for iterating over paged responses.
type WorkloadNetworksListSegmentsPager struct {
	client    *WorkloadNetworksClient
	current   WorkloadNetworksListSegmentsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WorkloadNetworksListSegmentsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WorkloadNetworksListSegmentsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WorkloadNetworksListSegmentsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WorkloadNetworkSegmentsList.NextLink == nil || len(*p.current.WorkloadNetworkSegmentsList.NextLink) == 0 {
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
		p.err = p.client.listSegmentsHandleError(resp)
		return false
	}
	result, err := p.client.listSegmentsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current WorkloadNetworksListSegmentsResponse page.
func (p *WorkloadNetworksListSegmentsPager) PageResponse() WorkloadNetworksListSegmentsResponse {
	return p.current
}

// WorkloadNetworksListVMGroupsPager provides operations for iterating over paged responses.
type WorkloadNetworksListVMGroupsPager struct {
	client    *WorkloadNetworksClient
	current   WorkloadNetworksListVMGroupsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WorkloadNetworksListVMGroupsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WorkloadNetworksListVMGroupsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WorkloadNetworksListVMGroupsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WorkloadNetworkVMGroupsList.NextLink == nil || len(*p.current.WorkloadNetworkVMGroupsList.NextLink) == 0 {
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
		p.err = p.client.listVMGroupsHandleError(resp)
		return false
	}
	result, err := p.client.listVMGroupsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current WorkloadNetworksListVMGroupsResponse page.
func (p *WorkloadNetworksListVMGroupsPager) PageResponse() WorkloadNetworksListVMGroupsResponse {
	return p.current
}

// WorkloadNetworksListVirtualMachinesPager provides operations for iterating over paged responses.
type WorkloadNetworksListVirtualMachinesPager struct {
	client    *WorkloadNetworksClient
	current   WorkloadNetworksListVirtualMachinesResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WorkloadNetworksListVirtualMachinesResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WorkloadNetworksListVirtualMachinesPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WorkloadNetworksListVirtualMachinesPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WorkloadNetworkVirtualMachinesList.NextLink == nil || len(*p.current.WorkloadNetworkVirtualMachinesList.NextLink) == 0 {
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
		p.err = p.client.listVirtualMachinesHandleError(resp)
		return false
	}
	result, err := p.client.listVirtualMachinesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current WorkloadNetworksListVirtualMachinesResponse page.
func (p *WorkloadNetworksListVirtualMachinesPager) PageResponse() WorkloadNetworksListVirtualMachinesResponse {
	return p.current
}
