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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mixedreality/armmixedreality"
	"net/http"
	"net/url"
	"regexp"
)

// RemoteRenderingAccountsServer is a fake server for instances of the armmixedreality.RemoteRenderingAccountsClient type.
type RemoteRenderingAccountsServer struct {
	// Create is the fake for method RemoteRenderingAccountsClient.Create
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	Create func(ctx context.Context, resourceGroupName string, accountName string, remoteRenderingAccount armmixedreality.RemoteRenderingAccount, options *armmixedreality.RemoteRenderingAccountsClientCreateOptions) (resp azfake.Responder[armmixedreality.RemoteRenderingAccountsClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method RemoteRenderingAccountsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, accountName string, options *armmixedreality.RemoteRenderingAccountsClientDeleteOptions) (resp azfake.Responder[armmixedreality.RemoteRenderingAccountsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method RemoteRenderingAccountsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, accountName string, options *armmixedreality.RemoteRenderingAccountsClientGetOptions) (resp azfake.Responder[armmixedreality.RemoteRenderingAccountsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method RemoteRenderingAccountsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armmixedreality.RemoteRenderingAccountsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armmixedreality.RemoteRenderingAccountsClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method RemoteRenderingAccountsClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armmixedreality.RemoteRenderingAccountsClientListBySubscriptionOptions) (resp azfake.PagerResponder[armmixedreality.RemoteRenderingAccountsClientListBySubscriptionResponse])

	// ListKeys is the fake for method RemoteRenderingAccountsClient.ListKeys
	// HTTP status codes to indicate success: http.StatusOK
	ListKeys func(ctx context.Context, resourceGroupName string, accountName string, options *armmixedreality.RemoteRenderingAccountsClientListKeysOptions) (resp azfake.Responder[armmixedreality.RemoteRenderingAccountsClientListKeysResponse], errResp azfake.ErrorResponder)

	// RegenerateKeys is the fake for method RemoteRenderingAccountsClient.RegenerateKeys
	// HTTP status codes to indicate success: http.StatusOK
	RegenerateKeys func(ctx context.Context, resourceGroupName string, accountName string, regenerate armmixedreality.AccountKeyRegenerateRequest, options *armmixedreality.RemoteRenderingAccountsClientRegenerateKeysOptions) (resp azfake.Responder[armmixedreality.RemoteRenderingAccountsClientRegenerateKeysResponse], errResp azfake.ErrorResponder)

	// Update is the fake for method RemoteRenderingAccountsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, accountName string, remoteRenderingAccount armmixedreality.RemoteRenderingAccount, options *armmixedreality.RemoteRenderingAccountsClientUpdateOptions) (resp azfake.Responder[armmixedreality.RemoteRenderingAccountsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewRemoteRenderingAccountsServerTransport creates a new instance of RemoteRenderingAccountsServerTransport with the provided implementation.
// The returned RemoteRenderingAccountsServerTransport instance is connected to an instance of armmixedreality.RemoteRenderingAccountsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRemoteRenderingAccountsServerTransport(srv *RemoteRenderingAccountsServer) *RemoteRenderingAccountsServerTransport {
	return &RemoteRenderingAccountsServerTransport{
		srv:                         srv,
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armmixedreality.RemoteRenderingAccountsClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armmixedreality.RemoteRenderingAccountsClientListBySubscriptionResponse]](),
	}
}

// RemoteRenderingAccountsServerTransport connects instances of armmixedreality.RemoteRenderingAccountsClient to instances of RemoteRenderingAccountsServer.
// Don't use this type directly, use NewRemoteRenderingAccountsServerTransport instead.
type RemoteRenderingAccountsServerTransport struct {
	srv                         *RemoteRenderingAccountsServer
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armmixedreality.RemoteRenderingAccountsClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armmixedreality.RemoteRenderingAccountsClientListBySubscriptionResponse]]
}

// Do implements the policy.Transporter interface for RemoteRenderingAccountsServerTransport.
func (r *RemoteRenderingAccountsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "RemoteRenderingAccountsClient.Create":
		resp, err = r.dispatchCreate(req)
	case "RemoteRenderingAccountsClient.Delete":
		resp, err = r.dispatchDelete(req)
	case "RemoteRenderingAccountsClient.Get":
		resp, err = r.dispatchGet(req)
	case "RemoteRenderingAccountsClient.NewListByResourceGroupPager":
		resp, err = r.dispatchNewListByResourceGroupPager(req)
	case "RemoteRenderingAccountsClient.NewListBySubscriptionPager":
		resp, err = r.dispatchNewListBySubscriptionPager(req)
	case "RemoteRenderingAccountsClient.ListKeys":
		resp, err = r.dispatchListKeys(req)
	case "RemoteRenderingAccountsClient.RegenerateKeys":
		resp, err = r.dispatchRegenerateKeys(req)
	case "RemoteRenderingAccountsClient.Update":
		resp, err = r.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RemoteRenderingAccountsServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if r.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MixedReality/remoteRenderingAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmixedreality.RemoteRenderingAccount](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Create(req.Context(), resourceGroupNameParam, accountNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RemoteRenderingAccount, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RemoteRenderingAccountsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if r.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MixedReality/remoteRenderingAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Delete(req.Context(), resourceGroupNameParam, accountNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RemoteRenderingAccountsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MixedReality/remoteRenderingAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceGroupNameParam, accountNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RemoteRenderingAccount, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RemoteRenderingAccountsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := r.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MixedReality/remoteRenderingAccounts`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := r.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		r.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armmixedreality.RemoteRenderingAccountsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		r.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (r *RemoteRenderingAccountsServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := r.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MixedReality/remoteRenderingAccounts`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := r.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		r.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armmixedreality.RemoteRenderingAccountsClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		r.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (r *RemoteRenderingAccountsServerTransport) dispatchListKeys(req *http.Request) (*http.Response, error) {
	if r.srv.ListKeys == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListKeys not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MixedReality/remoteRenderingAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listKeys`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.ListKeys(req.Context(), resourceGroupNameParam, accountNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AccountKeys, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RemoteRenderingAccountsServerTransport) dispatchRegenerateKeys(req *http.Request) (*http.Response, error) {
	if r.srv.RegenerateKeys == nil {
		return nil, &nonRetriableError{errors.New("fake for method RegenerateKeys not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MixedReality/remoteRenderingAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/regenerateKeys`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmixedreality.AccountKeyRegenerateRequest](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.RegenerateKeys(req.Context(), resourceGroupNameParam, accountNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AccountKeys, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RemoteRenderingAccountsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if r.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MixedReality/remoteRenderingAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmixedreality.RemoteRenderingAccount](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Update(req.Context(), resourceGroupNameParam, accountNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RemoteRenderingAccount, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
