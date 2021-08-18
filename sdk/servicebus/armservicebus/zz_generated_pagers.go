// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicebus

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"reflect"
)

type DisasterRecoveryConfigsListAuthorizationRulesPager interface {
	azcore.Pager
	// PageResponse returns the current DisasterRecoveryConfigsListAuthorizationRulesResponse.
	PageResponse() DisasterRecoveryConfigsListAuthorizationRulesResponse
}

type disasterRecoveryConfigsListAuthorizationRulesPager struct {
	client    *DisasterRecoveryConfigsClient
	current   DisasterRecoveryConfigsListAuthorizationRulesResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, DisasterRecoveryConfigsListAuthorizationRulesResponse) (*azcore.Request, error)
}

func (p *disasterRecoveryConfigsListAuthorizationRulesPager) Err() error {
	return p.err
}

func (p *disasterRecoveryConfigsListAuthorizationRulesPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SBAuthorizationRuleListResult.NextLink == nil || len(*p.current.SBAuthorizationRuleListResult.NextLink) == 0 {
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
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listAuthorizationRulesHandleError(resp)
		return false
	}
	result, err := p.client.listAuthorizationRulesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *disasterRecoveryConfigsListAuthorizationRulesPager) PageResponse() DisasterRecoveryConfigsListAuthorizationRulesResponse {
	return p.current
}

type DisasterRecoveryConfigsListPager interface {
	azcore.Pager
	// PageResponse returns the current DisasterRecoveryConfigsListResponse.
	PageResponse() DisasterRecoveryConfigsListResponse
}

type disasterRecoveryConfigsListPager struct {
	client    *DisasterRecoveryConfigsClient
	current   DisasterRecoveryConfigsListResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, DisasterRecoveryConfigsListResponse) (*azcore.Request, error)
}

func (p *disasterRecoveryConfigsListPager) Err() error {
	return p.err
}

func (p *disasterRecoveryConfigsListPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ArmDisasterRecoveryListResult.NextLink == nil || len(*p.current.ArmDisasterRecoveryListResult.NextLink) == 0 {
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
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
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

func (p *disasterRecoveryConfigsListPager) PageResponse() DisasterRecoveryConfigsListResponse {
	return p.current
}

type EventHubsListByNamespacePager interface {
	azcore.Pager
	// PageResponse returns the current EventHubsListByNamespaceResponse.
	PageResponse() EventHubsListByNamespaceResponse
}

type eventHubsListByNamespacePager struct {
	client    *EventHubsClient
	current   EventHubsListByNamespaceResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, EventHubsListByNamespaceResponse) (*azcore.Request, error)
}

func (p *eventHubsListByNamespacePager) Err() error {
	return p.err
}

func (p *eventHubsListByNamespacePager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.EventHubListResult.NextLink == nil || len(*p.current.EventHubListResult.NextLink) == 0 {
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
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listByNamespaceHandleError(resp)
		return false
	}
	result, err := p.client.listByNamespaceHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *eventHubsListByNamespacePager) PageResponse() EventHubsListByNamespaceResponse {
	return p.current
}

type MigrationConfigsListPager interface {
	azcore.Pager
	// PageResponse returns the current MigrationConfigsListResponse.
	PageResponse() MigrationConfigsListResponse
}

type migrationConfigsListPager struct {
	client    *MigrationConfigsClient
	current   MigrationConfigsListResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, MigrationConfigsListResponse) (*azcore.Request, error)
}

func (p *migrationConfigsListPager) Err() error {
	return p.err
}

func (p *migrationConfigsListPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.MigrationConfigListResult.NextLink == nil || len(*p.current.MigrationConfigListResult.NextLink) == 0 {
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
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
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

func (p *migrationConfigsListPager) PageResponse() MigrationConfigsListResponse {
	return p.current
}

type NamespacesListAuthorizationRulesPager interface {
	azcore.Pager
	// PageResponse returns the current NamespacesListAuthorizationRulesResponse.
	PageResponse() NamespacesListAuthorizationRulesResponse
}

type namespacesListAuthorizationRulesPager struct {
	client    *NamespacesClient
	current   NamespacesListAuthorizationRulesResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, NamespacesListAuthorizationRulesResponse) (*azcore.Request, error)
}

func (p *namespacesListAuthorizationRulesPager) Err() error {
	return p.err
}

func (p *namespacesListAuthorizationRulesPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SBAuthorizationRuleListResult.NextLink == nil || len(*p.current.SBAuthorizationRuleListResult.NextLink) == 0 {
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
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listAuthorizationRulesHandleError(resp)
		return false
	}
	result, err := p.client.listAuthorizationRulesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *namespacesListAuthorizationRulesPager) PageResponse() NamespacesListAuthorizationRulesResponse {
	return p.current
}

type NamespacesListByResourceGroupPager interface {
	azcore.Pager
	// PageResponse returns the current NamespacesListByResourceGroupResponse.
	PageResponse() NamespacesListByResourceGroupResponse
}

type namespacesListByResourceGroupPager struct {
	client    *NamespacesClient
	current   NamespacesListByResourceGroupResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, NamespacesListByResourceGroupResponse) (*azcore.Request, error)
}

func (p *namespacesListByResourceGroupPager) Err() error {
	return p.err
}

func (p *namespacesListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SBNamespaceListResult.NextLink == nil || len(*p.current.SBNamespaceListResult.NextLink) == 0 {
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
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
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

func (p *namespacesListByResourceGroupPager) PageResponse() NamespacesListByResourceGroupResponse {
	return p.current
}

type NamespacesListIPFilterRulesPager interface {
	azcore.Pager
	// PageResponse returns the current NamespacesListIPFilterRulesResponse.
	PageResponse() NamespacesListIPFilterRulesResponse
}

type namespacesListIPFilterRulesPager struct {
	client    *NamespacesClient
	current   NamespacesListIPFilterRulesResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, NamespacesListIPFilterRulesResponse) (*azcore.Request, error)
}

func (p *namespacesListIPFilterRulesPager) Err() error {
	return p.err
}

func (p *namespacesListIPFilterRulesPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.IPFilterRuleListResult.NextLink == nil || len(*p.current.IPFilterRuleListResult.NextLink) == 0 {
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
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listIPFilterRulesHandleError(resp)
		return false
	}
	result, err := p.client.listIPFilterRulesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *namespacesListIPFilterRulesPager) PageResponse() NamespacesListIPFilterRulesResponse {
	return p.current
}

type NamespacesListNetworkRuleSetsPager interface {
	azcore.Pager
	// PageResponse returns the current NamespacesListNetworkRuleSetsResponse.
	PageResponse() NamespacesListNetworkRuleSetsResponse
}

type namespacesListNetworkRuleSetsPager struct {
	client    *NamespacesClient
	current   NamespacesListNetworkRuleSetsResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, NamespacesListNetworkRuleSetsResponse) (*azcore.Request, error)
}

func (p *namespacesListNetworkRuleSetsPager) Err() error {
	return p.err
}

func (p *namespacesListNetworkRuleSetsPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.NetworkRuleSetListResult.NextLink == nil || len(*p.current.NetworkRuleSetListResult.NextLink) == 0 {
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
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listNetworkRuleSetsHandleError(resp)
		return false
	}
	result, err := p.client.listNetworkRuleSetsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *namespacesListNetworkRuleSetsPager) PageResponse() NamespacesListNetworkRuleSetsResponse {
	return p.current
}

type NamespacesListPager interface {
	azcore.Pager
	// PageResponse returns the current NamespacesListResponse.
	PageResponse() NamespacesListResponse
}

type namespacesListPager struct {
	client    *NamespacesClient
	current   NamespacesListResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, NamespacesListResponse) (*azcore.Request, error)
}

func (p *namespacesListPager) Err() error {
	return p.err
}

func (p *namespacesListPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SBNamespaceListResult.NextLink == nil || len(*p.current.SBNamespaceListResult.NextLink) == 0 {
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
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
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

func (p *namespacesListPager) PageResponse() NamespacesListResponse {
	return p.current
}

type NamespacesListVirtualNetworkRulesPager interface {
	azcore.Pager
	// PageResponse returns the current NamespacesListVirtualNetworkRulesResponse.
	PageResponse() NamespacesListVirtualNetworkRulesResponse
}

type namespacesListVirtualNetworkRulesPager struct {
	client    *NamespacesClient
	current   NamespacesListVirtualNetworkRulesResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, NamespacesListVirtualNetworkRulesResponse) (*azcore.Request, error)
}

func (p *namespacesListVirtualNetworkRulesPager) Err() error {
	return p.err
}

func (p *namespacesListVirtualNetworkRulesPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.VirtualNetworkRuleListResult.NextLink == nil || len(*p.current.VirtualNetworkRuleListResult.NextLink) == 0 {
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
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listVirtualNetworkRulesHandleError(resp)
		return false
	}
	result, err := p.client.listVirtualNetworkRulesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *namespacesListVirtualNetworkRulesPager) PageResponse() NamespacesListVirtualNetworkRulesResponse {
	return p.current
}

type OperationsListPager interface {
	azcore.Pager
	// PageResponse returns the current OperationsListResponse.
	PageResponse() OperationsListResponse
}

type operationsListPager struct {
	client    *OperationsClient
	current   OperationsListResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, OperationsListResponse) (*azcore.Request, error)
}

func (p *operationsListPager) Err() error {
	return p.err
}

func (p *operationsListPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
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
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
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

func (p *operationsListPager) PageResponse() OperationsListResponse {
	return p.current
}

type PremiumMessagingRegionsListPager interface {
	azcore.Pager
	// PageResponse returns the current PremiumMessagingRegionsListResponse.
	PageResponse() PremiumMessagingRegionsListResponse
}

type premiumMessagingRegionsListPager struct {
	client    *PremiumMessagingRegionsClient
	current   PremiumMessagingRegionsListResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, PremiumMessagingRegionsListResponse) (*azcore.Request, error)
}

func (p *premiumMessagingRegionsListPager) Err() error {
	return p.err
}

func (p *premiumMessagingRegionsListPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PremiumMessagingRegionsListResult.NextLink == nil || len(*p.current.PremiumMessagingRegionsListResult.NextLink) == 0 {
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
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
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

func (p *premiumMessagingRegionsListPager) PageResponse() PremiumMessagingRegionsListResponse {
	return p.current
}

type PrivateEndpointConnectionsListPager interface {
	azcore.Pager
	// PageResponse returns the current PrivateEndpointConnectionsListResponse.
	PageResponse() PrivateEndpointConnectionsListResponse
}

type privateEndpointConnectionsListPager struct {
	client    *PrivateEndpointConnectionsClient
	current   PrivateEndpointConnectionsListResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, PrivateEndpointConnectionsListResponse) (*azcore.Request, error)
}

func (p *privateEndpointConnectionsListPager) Err() error {
	return p.err
}

func (p *privateEndpointConnectionsListPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PrivateEndpointConnectionListResult.NextLink == nil || len(*p.current.PrivateEndpointConnectionListResult.NextLink) == 0 {
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
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
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

func (p *privateEndpointConnectionsListPager) PageResponse() PrivateEndpointConnectionsListResponse {
	return p.current
}

type QueuesListAuthorizationRulesPager interface {
	azcore.Pager
	// PageResponse returns the current QueuesListAuthorizationRulesResponse.
	PageResponse() QueuesListAuthorizationRulesResponse
}

type queuesListAuthorizationRulesPager struct {
	client    *QueuesClient
	current   QueuesListAuthorizationRulesResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, QueuesListAuthorizationRulesResponse) (*azcore.Request, error)
}

func (p *queuesListAuthorizationRulesPager) Err() error {
	return p.err
}

func (p *queuesListAuthorizationRulesPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SBAuthorizationRuleListResult.NextLink == nil || len(*p.current.SBAuthorizationRuleListResult.NextLink) == 0 {
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
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listAuthorizationRulesHandleError(resp)
		return false
	}
	result, err := p.client.listAuthorizationRulesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *queuesListAuthorizationRulesPager) PageResponse() QueuesListAuthorizationRulesResponse {
	return p.current
}

type QueuesListByNamespacePager interface {
	azcore.Pager
	// PageResponse returns the current QueuesListByNamespaceResponse.
	PageResponse() QueuesListByNamespaceResponse
}

type queuesListByNamespacePager struct {
	client    *QueuesClient
	current   QueuesListByNamespaceResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, QueuesListByNamespaceResponse) (*azcore.Request, error)
}

func (p *queuesListByNamespacePager) Err() error {
	return p.err
}

func (p *queuesListByNamespacePager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SBQueueListResult.NextLink == nil || len(*p.current.SBQueueListResult.NextLink) == 0 {
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
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listByNamespaceHandleError(resp)
		return false
	}
	result, err := p.client.listByNamespaceHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *queuesListByNamespacePager) PageResponse() QueuesListByNamespaceResponse {
	return p.current
}

type RegionsListBySKUPager interface {
	azcore.Pager
	// PageResponse returns the current RegionsListBySKUResponse.
	PageResponse() RegionsListBySKUResponse
}

type regionsListBySKUPager struct {
	client    *RegionsClient
	current   RegionsListBySKUResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, RegionsListBySKUResponse) (*azcore.Request, error)
}

func (p *regionsListBySKUPager) Err() error {
	return p.err
}

func (p *regionsListBySKUPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PremiumMessagingRegionsListResult.NextLink == nil || len(*p.current.PremiumMessagingRegionsListResult.NextLink) == 0 {
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
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listBySKUHandleError(resp)
		return false
	}
	result, err := p.client.listBySKUHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *regionsListBySKUPager) PageResponse() RegionsListBySKUResponse {
	return p.current
}

type RulesListBySubscriptionsPager interface {
	azcore.Pager
	// PageResponse returns the current RulesListBySubscriptionsResponse.
	PageResponse() RulesListBySubscriptionsResponse
}

type rulesListBySubscriptionsPager struct {
	client    *RulesClient
	current   RulesListBySubscriptionsResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, RulesListBySubscriptionsResponse) (*azcore.Request, error)
}

func (p *rulesListBySubscriptionsPager) Err() error {
	return p.err
}

func (p *rulesListBySubscriptionsPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RuleListResult.NextLink == nil || len(*p.current.RuleListResult.NextLink) == 0 {
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
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listBySubscriptionsHandleError(resp)
		return false
	}
	result, err := p.client.listBySubscriptionsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *rulesListBySubscriptionsPager) PageResponse() RulesListBySubscriptionsResponse {
	return p.current
}

type SubscriptionsListByTopicPager interface {
	azcore.Pager
	// PageResponse returns the current SubscriptionsListByTopicResponse.
	PageResponse() SubscriptionsListByTopicResponse
}

type subscriptionsListByTopicPager struct {
	client    *SubscriptionsClient
	current   SubscriptionsListByTopicResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, SubscriptionsListByTopicResponse) (*azcore.Request, error)
}

func (p *subscriptionsListByTopicPager) Err() error {
	return p.err
}

func (p *subscriptionsListByTopicPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SBSubscriptionListResult.NextLink == nil || len(*p.current.SBSubscriptionListResult.NextLink) == 0 {
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
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listByTopicHandleError(resp)
		return false
	}
	result, err := p.client.listByTopicHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *subscriptionsListByTopicPager) PageResponse() SubscriptionsListByTopicResponse {
	return p.current
}

type TopicsListAuthorizationRulesPager interface {
	azcore.Pager
	// PageResponse returns the current TopicsListAuthorizationRulesResponse.
	PageResponse() TopicsListAuthorizationRulesResponse
}

type topicsListAuthorizationRulesPager struct {
	client    *TopicsClient
	current   TopicsListAuthorizationRulesResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, TopicsListAuthorizationRulesResponse) (*azcore.Request, error)
}

func (p *topicsListAuthorizationRulesPager) Err() error {
	return p.err
}

func (p *topicsListAuthorizationRulesPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SBAuthorizationRuleListResult.NextLink == nil || len(*p.current.SBAuthorizationRuleListResult.NextLink) == 0 {
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
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listAuthorizationRulesHandleError(resp)
		return false
	}
	result, err := p.client.listAuthorizationRulesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *topicsListAuthorizationRulesPager) PageResponse() TopicsListAuthorizationRulesResponse {
	return p.current
}

type TopicsListByNamespacePager interface {
	azcore.Pager
	// PageResponse returns the current TopicsListByNamespaceResponse.
	PageResponse() TopicsListByNamespaceResponse
}

type topicsListByNamespacePager struct {
	client    *TopicsClient
	current   TopicsListByNamespaceResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, TopicsListByNamespaceResponse) (*azcore.Request, error)
}

func (p *topicsListByNamespacePager) Err() error {
	return p.err
}

func (p *topicsListByNamespacePager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SBTopicListResult.NextLink == nil || len(*p.current.SBTopicListResult.NextLink) == 0 {
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
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listByNamespaceHandleError(resp)
		return false
	}
	result, err := p.client.listByNamespaceHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *topicsListByNamespacePager) PageResponse() TopicsListByNamespaceResponse {
	return p.current
}
