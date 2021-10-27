//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package internal

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// KeyVaultClientGetDeletedKeysPager provides operations for iterating over paged responses.
type KeyVaultClientGetDeletedKeysPager struct {
	client    *KeyVaultClient
	current   KeyVaultClientGetDeletedKeysResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, KeyVaultClientGetDeletedKeysResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *KeyVaultClientGetDeletedKeysPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *KeyVaultClientGetDeletedKeysPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeletedKeyListResult.NextLink == nil || len(*p.current.DeletedKeyListResult.NextLink) == 0 {
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
	resp, err := p.client.Con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.getDeletedKeysHandleError(resp)
		return false
	}
	result, err := p.client.getDeletedKeysHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current KeyVaultClientGetDeletedKeysResponse page.
func (p *KeyVaultClientGetDeletedKeysPager) PageResponse() KeyVaultClientGetDeletedKeysResponse {
	return p.current
}

// KeyVaultClientGetKeyVersionsPager provides operations for iterating over paged responses.
type KeyVaultClientGetKeyVersionsPager struct {
	client    *KeyVaultClient
	current   KeyVaultClientGetKeyVersionsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, KeyVaultClientGetKeyVersionsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *KeyVaultClientGetKeyVersionsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *KeyVaultClientGetKeyVersionsPager) NextPage(ctx context.Context) bool {
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
	resp, err := p.client.Con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.getKeyVersionsHandleError(resp)
		return false
	}
	result, err := p.client.getKeyVersionsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current KeyVaultClientGetKeyVersionsResponse page.
func (p *KeyVaultClientGetKeyVersionsPager) PageResponse() KeyVaultClientGetKeyVersionsResponse {
	return p.current
}

// KeyVaultClientGetKeysPager provides operations for iterating over paged responses.
type KeyVaultClientGetKeysPager struct {
	client    *KeyVaultClient
	current   KeyVaultClientGetKeysResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, KeyVaultClientGetKeysResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *KeyVaultClientGetKeysPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *KeyVaultClientGetKeysPager) NextPage(ctx context.Context) bool {
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
	resp, err := p.client.Con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.getKeysHandleError(resp)
		return false
	}
	result, err := p.client.getKeysHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current KeyVaultClientGetKeysResponse page.
func (p *KeyVaultClientGetKeysPager) PageResponse() KeyVaultClientGetKeysResponse {
	return p.current
}

// RoleAssignmentsListForScopePager provides operations for iterating over paged responses.
type RoleAssignmentsListForScopePager struct {
	client    *roleAssignmentsClient
	current   RoleAssignmentsListForScopeResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, RoleAssignmentsListForScopeResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *RoleAssignmentsListForScopePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *RoleAssignmentsListForScopePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RoleAssignmentListResult.NextLink == nil || len(*p.current.RoleAssignmentListResult.NextLink) == 0 {
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
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listForScopeHandleError(resp)
		return false
	}
	result, err := p.client.listForScopeHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current RoleAssignmentsListForScopeResponse page.
func (p *RoleAssignmentsListForScopePager) PageResponse() RoleAssignmentsListForScopeResponse {
	return p.current
}

// RoleDefinitionsListPager provides operations for iterating over paged responses.
type RoleDefinitionsListPager struct {
	client    *roleDefinitionsClient
	current   RoleDefinitionsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, RoleDefinitionsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *RoleDefinitionsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *RoleDefinitionsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RoleDefinitionListResult.NextLink == nil || len(*p.current.RoleDefinitionListResult.NextLink) == 0 {
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

// PageResponse returns the current RoleDefinitionsListResponse page.
func (p *RoleDefinitionsListPager) PageResponse() RoleDefinitionsListResponse {
	return p.current
}
