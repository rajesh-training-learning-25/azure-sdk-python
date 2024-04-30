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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// GroupQuotaLimitsRequestServer is a fake server for instances of the armquota.GroupQuotaLimitsRequestClient type.
type GroupQuotaLimitsRequestServer struct {
	// BeginCreateOrUpdate is the fake for method GroupQuotaLimitsRequestClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, managementGroupID string, groupQuotaName string, resourceProviderName string, resourceName string, options *armquota.GroupQuotaLimitsRequestClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armquota.GroupQuotaLimitsRequestClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method GroupQuotaLimitsRequestClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, managementGroupID string, groupQuotaName string, requestID string, options *armquota.GroupQuotaLimitsRequestClientGetOptions) (resp azfake.Responder[armquota.GroupQuotaLimitsRequestClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method GroupQuotaLimitsRequestClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(managementGroupID string, groupQuotaName string, resourceProviderName string, filter string, options *armquota.GroupQuotaLimitsRequestClientListOptions) (resp azfake.PagerResponder[armquota.GroupQuotaLimitsRequestClientListResponse])

	// BeginUpdate is the fake for method GroupQuotaLimitsRequestClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, managementGroupID string, groupQuotaName string, resourceProviderName string, resourceName string, options *armquota.GroupQuotaLimitsRequestClientBeginUpdateOptions) (resp azfake.PollerResponder[armquota.GroupQuotaLimitsRequestClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewGroupQuotaLimitsRequestServerTransport creates a new instance of GroupQuotaLimitsRequestServerTransport with the provided implementation.
// The returned GroupQuotaLimitsRequestServerTransport instance is connected to an instance of armquota.GroupQuotaLimitsRequestClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewGroupQuotaLimitsRequestServerTransport(srv *GroupQuotaLimitsRequestServer) *GroupQuotaLimitsRequestServerTransport {
	return &GroupQuotaLimitsRequestServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armquota.GroupQuotaLimitsRequestClientCreateOrUpdateResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armquota.GroupQuotaLimitsRequestClientListResponse]](),
		beginUpdate:         newTracker[azfake.PollerResponder[armquota.GroupQuotaLimitsRequestClientUpdateResponse]](),
	}
}

// GroupQuotaLimitsRequestServerTransport connects instances of armquota.GroupQuotaLimitsRequestClient to instances of GroupQuotaLimitsRequestServer.
// Don't use this type directly, use NewGroupQuotaLimitsRequestServerTransport instead.
type GroupQuotaLimitsRequestServerTransport struct {
	srv                 *GroupQuotaLimitsRequestServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armquota.GroupQuotaLimitsRequestClientCreateOrUpdateResponse]]
	newListPager        *tracker[azfake.PagerResponder[armquota.GroupQuotaLimitsRequestClientListResponse]]
	beginUpdate         *tracker[azfake.PollerResponder[armquota.GroupQuotaLimitsRequestClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for GroupQuotaLimitsRequestServerTransport.
func (g *GroupQuotaLimitsRequestServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "GroupQuotaLimitsRequestClient.BeginCreateOrUpdate":
		resp, err = g.dispatchBeginCreateOrUpdate(req)
	case "GroupQuotaLimitsRequestClient.Get":
		resp, err = g.dispatchGet(req)
	case "GroupQuotaLimitsRequestClient.NewListPager":
		resp, err = g.dispatchNewListPager(req)
	case "GroupQuotaLimitsRequestClient.BeginUpdate":
		resp, err = g.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GroupQuotaLimitsRequestServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if g.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := g.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/providers/Microsoft\.Management/managementGroups/(?P<managementGroupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Quota/groupQuotas/(?P<groupQuotaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceProviders/(?P<resourceProviderName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/groupQuotaRequests/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armquota.SubmittedResourceRequestStatus](req)
		if err != nil {
			return nil, err
		}
		managementGroupIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupId")])
		if err != nil {
			return nil, err
		}
		groupQuotaNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("groupQuotaName")])
		if err != nil {
			return nil, err
		}
		resourceProviderNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceProviderName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		var options *armquota.GroupQuotaLimitsRequestClientBeginCreateOrUpdateOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armquota.GroupQuotaLimitsRequestClientBeginCreateOrUpdateOptions{
				GroupQuotaRequest: &body,
			}
		}
		respr, errRespr := g.srv.BeginCreateOrUpdate(req.Context(), managementGroupIDParam, groupQuotaNameParam, resourceProviderNameParam, resourceNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		g.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		g.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		g.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (g *GroupQuotaLimitsRequestServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if g.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Management/managementGroups/(?P<managementGroupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Quota/groupQuotas/(?P<groupQuotaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/groupQuotaRequests/(?P<requestId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	managementGroupIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupId")])
	if err != nil {
		return nil, err
	}
	groupQuotaNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("groupQuotaName")])
	if err != nil {
		return nil, err
	}
	requestIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("requestId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.Get(req.Context(), managementGroupIDParam, groupQuotaNameParam, requestIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SubmittedResourceRequestStatus, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GroupQuotaLimitsRequestServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if g.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := g.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/providers/Microsoft\.Management/managementGroups/(?P<managementGroupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Quota/groupQuotas/(?P<groupQuotaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceProviders/(?P<resourceProviderName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/groupQuotaRequests`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		managementGroupIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupId")])
		if err != nil {
			return nil, err
		}
		groupQuotaNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("groupQuotaName")])
		if err != nil {
			return nil, err
		}
		resourceProviderNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceProviderName")])
		if err != nil {
			return nil, err
		}
		filterParam, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		resp := g.srv.NewListPager(managementGroupIDParam, groupQuotaNameParam, resourceProviderNameParam, filterParam, nil)
		newListPager = &resp
		g.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armquota.GroupQuotaLimitsRequestClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		g.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		g.newListPager.remove(req)
	}
	return resp, nil
}

func (g *GroupQuotaLimitsRequestServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if g.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := g.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/providers/Microsoft\.Management/managementGroups/(?P<managementGroupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Quota/groupQuotas/(?P<groupQuotaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceProviders/(?P<resourceProviderName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/groupQuotaRequests/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armquota.SubmittedResourceRequestStatus](req)
		if err != nil {
			return nil, err
		}
		managementGroupIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupId")])
		if err != nil {
			return nil, err
		}
		groupQuotaNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("groupQuotaName")])
		if err != nil {
			return nil, err
		}
		resourceProviderNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceProviderName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		var options *armquota.GroupQuotaLimitsRequestClientBeginUpdateOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armquota.GroupQuotaLimitsRequestClientBeginUpdateOptions{
				GroupQuotaRequest: &body,
			}
		}
		respr, errRespr := g.srv.BeginUpdate(req.Context(), managementGroupIDParam, groupQuotaNameParam, resourceProviderNameParam, resourceNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		g.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		g.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		g.beginUpdate.remove(req)
	}

	return resp, nil
}
