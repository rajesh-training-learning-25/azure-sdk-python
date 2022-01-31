//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkeyvault

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// KeysClientListPager provides operations for iterating over paged responses.
type KeysClientListPager struct {
	client    *KeysClient
	current   KeysClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, KeysClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *KeysClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *KeysClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.KeyListResult.NextLink == nil || len(*p.current.KeyListResult.NextLink) == 0 {
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

// PageResponse returns the current KeysClientListResponse page.
func (p *KeysClientListPager) PageResponse() KeysClientListResponse {
	return p.current
}

// KeysClientListVersionsPager provides operations for iterating over paged responses.
type KeysClientListVersionsPager struct {
	client    *KeysClient
	current   KeysClientListVersionsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, KeysClientListVersionsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *KeysClientListVersionsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *KeysClientListVersionsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.KeyListResult.NextLink == nil || len(*p.current.KeyListResult.NextLink) == 0 {
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
	result, err := p.client.listVersionsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current KeysClientListVersionsResponse page.
func (p *KeysClientListVersionsPager) PageResponse() KeysClientListVersionsResponse {
	return p.current
}

// MHSMPrivateEndpointConnectionsClientListByResourcePager provides operations for iterating over paged responses.
type MHSMPrivateEndpointConnectionsClientListByResourcePager struct {
	client    *MHSMPrivateEndpointConnectionsClient
	current   MHSMPrivateEndpointConnectionsClientListByResourceResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, MHSMPrivateEndpointConnectionsClientListByResourceResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *MHSMPrivateEndpointConnectionsClientListByResourcePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *MHSMPrivateEndpointConnectionsClientListByResourcePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.MHSMPrivateEndpointConnectionsListResult.NextLink == nil || len(*p.current.MHSMPrivateEndpointConnectionsListResult.NextLink) == 0 {
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
	result, err := p.client.listByResourceHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current MHSMPrivateEndpointConnectionsClientListByResourceResponse page.
func (p *MHSMPrivateEndpointConnectionsClientListByResourcePager) PageResponse() MHSMPrivateEndpointConnectionsClientListByResourceResponse {
	return p.current
}

// ManagedHsmsClientListByResourceGroupPager provides operations for iterating over paged responses.
type ManagedHsmsClientListByResourceGroupPager struct {
	client    *ManagedHsmsClient
	current   ManagedHsmsClientListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ManagedHsmsClientListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ManagedHsmsClientListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ManagedHsmsClientListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ManagedHsmListResult.NextLink == nil || len(*p.current.ManagedHsmListResult.NextLink) == 0 {
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

// PageResponse returns the current ManagedHsmsClientListByResourceGroupResponse page.
func (p *ManagedHsmsClientListByResourceGroupPager) PageResponse() ManagedHsmsClientListByResourceGroupResponse {
	return p.current
}

// ManagedHsmsClientListBySubscriptionPager provides operations for iterating over paged responses.
type ManagedHsmsClientListBySubscriptionPager struct {
	client    *ManagedHsmsClient
	current   ManagedHsmsClientListBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ManagedHsmsClientListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ManagedHsmsClientListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ManagedHsmsClientListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ManagedHsmListResult.NextLink == nil || len(*p.current.ManagedHsmListResult.NextLink) == 0 {
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

// PageResponse returns the current ManagedHsmsClientListBySubscriptionResponse page.
func (p *ManagedHsmsClientListBySubscriptionPager) PageResponse() ManagedHsmsClientListBySubscriptionResponse {
	return p.current
}

// ManagedHsmsClientListDeletedPager provides operations for iterating over paged responses.
type ManagedHsmsClientListDeletedPager struct {
	client    *ManagedHsmsClient
	current   ManagedHsmsClientListDeletedResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ManagedHsmsClientListDeletedResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ManagedHsmsClientListDeletedPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ManagedHsmsClientListDeletedPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeletedManagedHsmListResult.NextLink == nil || len(*p.current.DeletedManagedHsmListResult.NextLink) == 0 {
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
	result, err := p.client.listDeletedHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ManagedHsmsClientListDeletedResponse page.
func (p *ManagedHsmsClientListDeletedPager) PageResponse() ManagedHsmsClientListDeletedResponse {
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

// PrivateEndpointConnectionsClientListByResourcePager provides operations for iterating over paged responses.
type PrivateEndpointConnectionsClientListByResourcePager struct {
	client    *PrivateEndpointConnectionsClient
	current   PrivateEndpointConnectionsClientListByResourceResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PrivateEndpointConnectionsClientListByResourceResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PrivateEndpointConnectionsClientListByResourcePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PrivateEndpointConnectionsClientListByResourcePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
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
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByResourceHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PrivateEndpointConnectionsClientListByResourceResponse page.
func (p *PrivateEndpointConnectionsClientListByResourcePager) PageResponse() PrivateEndpointConnectionsClientListByResourceResponse {
	return p.current
}

// SecretsClientListPager provides operations for iterating over paged responses.
type SecretsClientListPager struct {
	client    *SecretsClient
	current   SecretsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SecretsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SecretsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SecretsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SecretListResult.NextLink == nil || len(*p.current.SecretListResult.NextLink) == 0 {
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

// PageResponse returns the current SecretsClientListResponse page.
func (p *SecretsClientListPager) PageResponse() SecretsClientListResponse {
	return p.current
}

// VaultsClientListByResourceGroupPager provides operations for iterating over paged responses.
type VaultsClientListByResourceGroupPager struct {
	client    *VaultsClient
	current   VaultsClientListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, VaultsClientListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *VaultsClientListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *VaultsClientListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.VaultListResult.NextLink == nil || len(*p.current.VaultListResult.NextLink) == 0 {
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

// PageResponse returns the current VaultsClientListByResourceGroupResponse page.
func (p *VaultsClientListByResourceGroupPager) PageResponse() VaultsClientListByResourceGroupResponse {
	return p.current
}

// VaultsClientListBySubscriptionPager provides operations for iterating over paged responses.
type VaultsClientListBySubscriptionPager struct {
	client    *VaultsClient
	current   VaultsClientListBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, VaultsClientListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *VaultsClientListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *VaultsClientListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.VaultListResult.NextLink == nil || len(*p.current.VaultListResult.NextLink) == 0 {
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

// PageResponse returns the current VaultsClientListBySubscriptionResponse page.
func (p *VaultsClientListBySubscriptionPager) PageResponse() VaultsClientListBySubscriptionResponse {
	return p.current
}

// VaultsClientListDeletedPager provides operations for iterating over paged responses.
type VaultsClientListDeletedPager struct {
	client    *VaultsClient
	current   VaultsClientListDeletedResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, VaultsClientListDeletedResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *VaultsClientListDeletedPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *VaultsClientListDeletedPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeletedVaultListResult.NextLink == nil || len(*p.current.DeletedVaultListResult.NextLink) == 0 {
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
	result, err := p.client.listDeletedHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current VaultsClientListDeletedResponse page.
func (p *VaultsClientListDeletedPager) PageResponse() VaultsClientListDeletedResponse {
	return p.current
}

// VaultsClientListPager provides operations for iterating over paged responses.
type VaultsClientListPager struct {
	client    *VaultsClient
	current   VaultsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, VaultsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *VaultsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *VaultsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ResourceListResult.NextLink == nil || len(*p.current.ResourceListResult.NextLink) == 0 {
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

// PageResponse returns the current VaultsClientListResponse page.
func (p *VaultsClientListPager) PageResponse() VaultsClientListResponse {
	return p.current
}
