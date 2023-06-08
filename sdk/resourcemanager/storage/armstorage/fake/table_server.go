//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"net/http"
	"reflect"
	"regexp"
)

// TableServer is a fake server for instances of the armstorage.TableClient type.
type TableServer struct {
	// Create is the fake for method TableClient.Create
	// HTTP status codes to indicate success: http.StatusOK
	Create func(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *armstorage.TableClientCreateOptions) (resp azfake.Responder[armstorage.TableClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method TableClient.Delete
	// HTTP status codes to indicate success: http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *armstorage.TableClientDeleteOptions) (resp azfake.Responder[armstorage.TableClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method TableClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *armstorage.TableClientGetOptions) (resp azfake.Responder[armstorage.TableClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method TableClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, accountName string, options *armstorage.TableClientListOptions) (resp azfake.PagerResponder[armstorage.TableClientListResponse])

	// Update is the fake for method TableClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *armstorage.TableClientUpdateOptions) (resp azfake.Responder[armstorage.TableClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewTableServerTransport creates a new instance of TableServerTransport with the provided implementation.
// The returned TableServerTransport instance is connected to an instance of armstorage.TableClient by way of the
// undefined.Transporter field.
func NewTableServerTransport(srv *TableServer) *TableServerTransport {
	return &TableServerTransport{srv: srv}
}

// TableServerTransport connects instances of armstorage.TableClient to instances of TableServer.
// Don't use this type directly, use NewTableServerTransport instead.
type TableServerTransport struct {
	srv          *TableServer
	newListPager *azfake.PagerResponder[armstorage.TableClientListResponse]
}

// Do implements the policy.Transporter interface for TableServerTransport.
func (t *TableServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "TableClient.Create":
		resp, err = t.dispatchCreate(req)
	case "TableClient.Delete":
		resp, err = t.dispatchDelete(req)
	case "TableClient.Get":
		resp, err = t.dispatchGet(req)
	case "TableClient.NewListPager":
		resp, err = t.dispatchNewListPager(req)
	case "TableClient.Update":
		resp, err = t.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *TableServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if t.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("method Create not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[a-zA-Z0-9-_]+)/tableServices/default/tables/(?P<tableName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armstorage.Table](req)
	if err != nil {
		return nil, err
	}
	var options *armstorage.TableClientCreateOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armstorage.TableClientCreateOptions{
			Parameters: &body,
		}
	}
	respr, errRespr := t.srv.Create(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("accountName")], matches[regex.SubexpIndex("tableName")], options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Table, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TableServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if t.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("method Delete not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[a-zA-Z0-9-_]+)/tableServices/default/tables/(?P<tableName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := t.srv.Delete(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("accountName")], matches[regex.SubexpIndex("tableName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TableServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if t.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[a-zA-Z0-9-_]+)/tableServices/default/tables/(?P<tableName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := t.srv.Get(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("accountName")], matches[regex.SubexpIndex("tableName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Table, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TableServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if t.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if t.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[a-zA-Z0-9-_]+)/tableServices/default/tables"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := t.srv.NewListPager(matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("accountName")], nil)
		t.newListPager = &resp
		server.PagerResponderInjectNextLinks(t.newListPager, req, func(page *armstorage.TableClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(t.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(t.newListPager) {
		t.newListPager = nil
	}
	return resp, nil
}

func (t *TableServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if t.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("method Update not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[a-zA-Z0-9-_]+)/tableServices/default/tables/(?P<tableName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armstorage.Table](req)
	if err != nil {
		return nil, err
	}
	var options *armstorage.TableClientUpdateOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armstorage.TableClientUpdateOptions{
			Parameters: &body,
		}
	}
	respr, errRespr := t.srv.Update(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("accountName")], matches[regex.SubexpIndex("tableName")], options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Table, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
