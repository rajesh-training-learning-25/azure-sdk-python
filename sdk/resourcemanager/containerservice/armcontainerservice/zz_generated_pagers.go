//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerservice

import (
	"context"
	"net/http"
	"reflect"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// AgentPoolsListPager provides operations for iterating over paged responses.
type AgentPoolsListPager struct {
	client    *AgentPoolsClient
	current   AgentPoolsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AgentPoolsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AgentPoolsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AgentPoolsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AgentPoolListResult.NextLink == nil || len(*p.current.AgentPoolListResult.NextLink) == 0 {
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

// PageResponse returns the current AgentPoolsListResponse page.
func (p *AgentPoolsListPager) PageResponse() AgentPoolsListResponse {
	return p.current
}

// MaintenanceConfigurationsListByManagedClusterPager provides operations for iterating over paged responses.
type MaintenanceConfigurationsListByManagedClusterPager struct {
	client    *MaintenanceConfigurationsClient
	current   MaintenanceConfigurationsListByManagedClusterResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, MaintenanceConfigurationsListByManagedClusterResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *MaintenanceConfigurationsListByManagedClusterPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *MaintenanceConfigurationsListByManagedClusterPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.MaintenanceConfigurationListResult.NextLink == nil || len(*p.current.MaintenanceConfigurationListResult.NextLink) == 0 {
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
		p.err = p.client.listByManagedClusterHandleError(resp)
		return false
	}
	result, err := p.client.listByManagedClusterHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current MaintenanceConfigurationsListByManagedClusterResponse page.
func (p *MaintenanceConfigurationsListByManagedClusterPager) PageResponse() MaintenanceConfigurationsListByManagedClusterResponse {
	return p.current
}

// ManagedClustersListByResourceGroupPager provides operations for iterating over paged responses.
type ManagedClustersListByResourceGroupPager struct {
	client    *ManagedClustersClient
	current   ManagedClustersListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ManagedClustersListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ManagedClustersListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ManagedClustersListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ManagedClusterListResult.NextLink == nil || len(*p.current.ManagedClusterListResult.NextLink) == 0 {
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
		p.err = p.client.listByResourceGroupHandleError(resp)
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

// PageResponse returns the current ManagedClustersListByResourceGroupResponse page.
func (p *ManagedClustersListByResourceGroupPager) PageResponse() ManagedClustersListByResourceGroupResponse {
	return p.current
}

// ManagedClustersListOutboundNetworkDependenciesEndpointsPager provides operations for iterating over paged responses.
type ManagedClustersListOutboundNetworkDependenciesEndpointsPager struct {
	client    *ManagedClustersClient
	current   ManagedClustersListOutboundNetworkDependenciesEndpointsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ManagedClustersListOutboundNetworkDependenciesEndpointsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ManagedClustersListOutboundNetworkDependenciesEndpointsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ManagedClustersListOutboundNetworkDependenciesEndpointsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OutboundEnvironmentEndpointCollection.NextLink == nil || len(*p.current.OutboundEnvironmentEndpointCollection.NextLink) == 0 {
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
		p.err = p.client.listOutboundNetworkDependenciesEndpointsHandleError(resp)
		return false
	}
	result, err := p.client.listOutboundNetworkDependenciesEndpointsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ManagedClustersListOutboundNetworkDependenciesEndpointsResponse page.
func (p *ManagedClustersListOutboundNetworkDependenciesEndpointsPager) PageResponse() ManagedClustersListOutboundNetworkDependenciesEndpointsResponse {
	return p.current
}

// ManagedClustersListPager provides operations for iterating over paged responses.
type ManagedClustersListPager struct {
	client    *ManagedClustersClient
	current   ManagedClustersListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ManagedClustersListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ManagedClustersListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ManagedClustersListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ManagedClusterListResult.NextLink == nil || len(*p.current.ManagedClusterListResult.NextLink) == 0 {
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

// PageResponse returns the current ManagedClustersListResponse page.
func (p *ManagedClustersListPager) PageResponse() ManagedClustersListResponse {
	return p.current
}

// SnapshotsListByResourceGroupPager provides operations for iterating over paged responses.
type SnapshotsListByResourceGroupPager struct {
	client    *SnapshotsClient
	current   SnapshotsListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SnapshotsListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SnapshotsListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SnapshotsListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SnapshotListResult.NextLink == nil || len(*p.current.SnapshotListResult.NextLink) == 0 {
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
		p.err = p.client.listByResourceGroupHandleError(resp)
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

// PageResponse returns the current SnapshotsListByResourceGroupResponse page.
func (p *SnapshotsListByResourceGroupPager) PageResponse() SnapshotsListByResourceGroupResponse {
	return p.current
}

// SnapshotsListPager provides operations for iterating over paged responses.
type SnapshotsListPager struct {
	client    *SnapshotsClient
	current   SnapshotsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SnapshotsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SnapshotsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SnapshotsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SnapshotListResult.NextLink == nil || len(*p.current.SnapshotListResult.NextLink) == 0 {
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

// PageResponse returns the current SnapshotsListResponse page.
func (p *SnapshotsListPager) PageResponse() SnapshotsListResponse {
	return p.current
}
