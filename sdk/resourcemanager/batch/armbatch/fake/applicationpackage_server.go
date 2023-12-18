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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v2"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
)

// ApplicationPackageServer is a fake server for instances of the armbatch.ApplicationPackageClient type.
type ApplicationPackageServer struct {
	// Activate is the fake for method ApplicationPackageClient.Activate
	// HTTP status codes to indicate success: http.StatusOK
	Activate func(ctx context.Context, resourceGroupName string, accountName string, applicationName string, versionName string, parameters armbatch.ActivateApplicationPackageParameters, options *armbatch.ApplicationPackageClientActivateOptions) (resp azfake.Responder[armbatch.ApplicationPackageClientActivateResponse], errResp azfake.ErrorResponder)

	// Create is the fake for method ApplicationPackageClient.Create
	// HTTP status codes to indicate success: http.StatusOK
	Create func(ctx context.Context, resourceGroupName string, accountName string, applicationName string, versionName string, options *armbatch.ApplicationPackageClientCreateOptions) (resp azfake.Responder[armbatch.ApplicationPackageClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ApplicationPackageClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, accountName string, applicationName string, versionName string, options *armbatch.ApplicationPackageClientDeleteOptions) (resp azfake.Responder[armbatch.ApplicationPackageClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ApplicationPackageClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, accountName string, applicationName string, versionName string, options *armbatch.ApplicationPackageClientGetOptions) (resp azfake.Responder[armbatch.ApplicationPackageClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ApplicationPackageClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, accountName string, applicationName string, options *armbatch.ApplicationPackageClientListOptions) (resp azfake.PagerResponder[armbatch.ApplicationPackageClientListResponse])
}

// NewApplicationPackageServerTransport creates a new instance of ApplicationPackageServerTransport with the provided implementation.
// The returned ApplicationPackageServerTransport instance is connected to an instance of armbatch.ApplicationPackageClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewApplicationPackageServerTransport(srv *ApplicationPackageServer) *ApplicationPackageServerTransport {
	return &ApplicationPackageServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armbatch.ApplicationPackageClientListResponse]](),
	}
}

// ApplicationPackageServerTransport connects instances of armbatch.ApplicationPackageClient to instances of ApplicationPackageServer.
// Don't use this type directly, use NewApplicationPackageServerTransport instead.
type ApplicationPackageServerTransport struct {
	srv          *ApplicationPackageServer
	newListPager *tracker[azfake.PagerResponder[armbatch.ApplicationPackageClientListResponse]]
}

// Do implements the policy.Transporter interface for ApplicationPackageServerTransport.
func (a *ApplicationPackageServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ApplicationPackageClient.Activate":
		resp, err = a.dispatchActivate(req)
	case "ApplicationPackageClient.Create":
		resp, err = a.dispatchCreate(req)
	case "ApplicationPackageClient.Delete":
		resp, err = a.dispatchDelete(req)
	case "ApplicationPackageClient.Get":
		resp, err = a.dispatchGet(req)
	case "ApplicationPackageClient.NewListPager":
		resp, err = a.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *ApplicationPackageServerTransport) dispatchActivate(req *http.Request) (*http.Response, error) {
	if a.srv.Activate == nil {
		return nil, &nonRetriableError{errors.New("fake for method Activate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Batch/batchAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications/(?P<applicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<versionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/activate`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armbatch.ActivateApplicationPackageParameters](req)
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
	applicationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationName")])
	if err != nil {
		return nil, err
	}
	versionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("versionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Activate(req.Context(), resourceGroupNameParam, accountNameParam, applicationNameParam, versionNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ApplicationPackage, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ApplicationPackageServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if a.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Batch/batchAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications/(?P<applicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<versionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armbatch.ApplicationPackage](req)
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
	applicationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationName")])
	if err != nil {
		return nil, err
	}
	versionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("versionName")])
	if err != nil {
		return nil, err
	}
	var options *armbatch.ApplicationPackageClientCreateOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armbatch.ApplicationPackageClientCreateOptions{
			Parameters: &body,
		}
	}
	respr, errRespr := a.srv.Create(req.Context(), resourceGroupNameParam, accountNameParam, applicationNameParam, versionNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ApplicationPackage, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ApplicationPackageServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if a.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Batch/batchAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications/(?P<applicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<versionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
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
	applicationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationName")])
	if err != nil {
		return nil, err
	}
	versionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("versionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Delete(req.Context(), resourceGroupNameParam, accountNameParam, applicationNameParam, versionNameParam, nil)
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

func (a *ApplicationPackageServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Batch/batchAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications/(?P<applicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<versionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
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
	applicationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationName")])
	if err != nil {
		return nil, err
	}
	versionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("versionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, accountNameParam, applicationNameParam, versionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ApplicationPackage, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ApplicationPackageServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := a.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Batch/batchAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications/(?P<applicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		applicationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationName")])
		if err != nil {
			return nil, err
		}
		maxresultsUnescaped, err := url.QueryUnescape(qp.Get("maxresults"))
		if err != nil {
			return nil, err
		}
		maxresultsParam, err := parseOptional(maxresultsUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *armbatch.ApplicationPackageClientListOptions
		if maxresultsParam != nil {
			options = &armbatch.ApplicationPackageClientListOptions{
				Maxresults: maxresultsParam,
			}
		}
		resp := a.srv.NewListPager(resourceGroupNameParam, accountNameParam, applicationNameParam, options)
		newListPager = &resp
		a.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armbatch.ApplicationPackageClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		a.newListPager.remove(req)
	}
	return resp, nil
}
