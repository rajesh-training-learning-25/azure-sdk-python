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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
	"net/http"
	"net/url"
	"regexp"
)

// DedicatedHostsServer is a fake server for instances of the armcompute.DedicatedHostsClient type.
type DedicatedHostsServer struct {
	// BeginCreateOrUpdate is the fake for method DedicatedHostsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters armcompute.DedicatedHost, options *armcompute.DedicatedHostsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcompute.DedicatedHostsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method DedicatedHostsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, options *armcompute.DedicatedHostsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcompute.DedicatedHostsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method DedicatedHostsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, options *armcompute.DedicatedHostsClientGetOptions) (resp azfake.Responder[armcompute.DedicatedHostsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListAvailableSizesPager is the fake for method DedicatedHostsClient.NewListAvailableSizesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAvailableSizesPager func(resourceGroupName string, hostGroupName string, hostName string, options *armcompute.DedicatedHostsClientListAvailableSizesOptions) (resp azfake.PagerResponder[armcompute.DedicatedHostsClientListAvailableSizesResponse])

	// NewListByHostGroupPager is the fake for method DedicatedHostsClient.NewListByHostGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByHostGroupPager func(resourceGroupName string, hostGroupName string, options *armcompute.DedicatedHostsClientListByHostGroupOptions) (resp azfake.PagerResponder[armcompute.DedicatedHostsClientListByHostGroupResponse])

	// BeginRestart is the fake for method DedicatedHostsClient.BeginRestart
	// HTTP status codes to indicate success: http.StatusOK
	BeginRestart func(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, options *armcompute.DedicatedHostsClientBeginRestartOptions) (resp azfake.PollerResponder[armcompute.DedicatedHostsClientRestartResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method DedicatedHostsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK
	BeginUpdate func(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters armcompute.DedicatedHostUpdate, options *armcompute.DedicatedHostsClientBeginUpdateOptions) (resp azfake.PollerResponder[armcompute.DedicatedHostsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewDedicatedHostsServerTransport creates a new instance of DedicatedHostsServerTransport with the provided implementation.
// The returned DedicatedHostsServerTransport instance is connected to an instance of armcompute.DedicatedHostsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDedicatedHostsServerTransport(srv *DedicatedHostsServer) *DedicatedHostsServerTransport {
	return &DedicatedHostsServerTransport{
		srv:                        srv,
		beginCreateOrUpdate:        newTracker[azfake.PollerResponder[armcompute.DedicatedHostsClientCreateOrUpdateResponse]](),
		beginDelete:                newTracker[azfake.PollerResponder[armcompute.DedicatedHostsClientDeleteResponse]](),
		newListAvailableSizesPager: newTracker[azfake.PagerResponder[armcompute.DedicatedHostsClientListAvailableSizesResponse]](),
		newListByHostGroupPager:    newTracker[azfake.PagerResponder[armcompute.DedicatedHostsClientListByHostGroupResponse]](),
		beginRestart:               newTracker[azfake.PollerResponder[armcompute.DedicatedHostsClientRestartResponse]](),
		beginUpdate:                newTracker[azfake.PollerResponder[armcompute.DedicatedHostsClientUpdateResponse]](),
	}
}

// DedicatedHostsServerTransport connects instances of armcompute.DedicatedHostsClient to instances of DedicatedHostsServer.
// Don't use this type directly, use NewDedicatedHostsServerTransport instead.
type DedicatedHostsServerTransport struct {
	srv                        *DedicatedHostsServer
	beginCreateOrUpdate        *tracker[azfake.PollerResponder[armcompute.DedicatedHostsClientCreateOrUpdateResponse]]
	beginDelete                *tracker[azfake.PollerResponder[armcompute.DedicatedHostsClientDeleteResponse]]
	newListAvailableSizesPager *tracker[azfake.PagerResponder[armcompute.DedicatedHostsClientListAvailableSizesResponse]]
	newListByHostGroupPager    *tracker[azfake.PagerResponder[armcompute.DedicatedHostsClientListByHostGroupResponse]]
	beginRestart               *tracker[azfake.PollerResponder[armcompute.DedicatedHostsClientRestartResponse]]
	beginUpdate                *tracker[azfake.PollerResponder[armcompute.DedicatedHostsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for DedicatedHostsServerTransport.
func (d *DedicatedHostsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DedicatedHostsClient.BeginCreateOrUpdate":
		resp, err = d.dispatchBeginCreateOrUpdate(req)
	case "DedicatedHostsClient.BeginDelete":
		resp, err = d.dispatchBeginDelete(req)
	case "DedicatedHostsClient.Get":
		resp, err = d.dispatchGet(req)
	case "DedicatedHostsClient.NewListAvailableSizesPager":
		resp, err = d.dispatchNewListAvailableSizesPager(req)
	case "DedicatedHostsClient.NewListByHostGroupPager":
		resp, err = d.dispatchNewListByHostGroupPager(req)
	case "DedicatedHostsClient.BeginRestart":
		resp, err = d.dispatchBeginRestart(req)
	case "DedicatedHostsClient.BeginUpdate":
		resp, err = d.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DedicatedHostsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := d.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/hostGroups/(?P<hostGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hosts/(?P<hostName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.DedicatedHost](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		hostGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("hostGroupName")])
		if err != nil {
			return nil, err
		}
		hostNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("hostName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, hostGroupNameUnescaped, hostNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		d.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		d.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		d.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (d *DedicatedHostsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if d.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := d.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/hostGroups/(?P<hostGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hosts/(?P<hostName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		hostGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("hostGroupName")])
		if err != nil {
			return nil, err
		}
		hostNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("hostName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, hostGroupNameUnescaped, hostNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		d.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		d.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		d.beginDelete.remove(req)
	}

	return resp, nil
}

func (d *DedicatedHostsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/hostGroups/(?P<hostGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hosts/(?P<hostName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	hostGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("hostGroupName")])
	if err != nil {
		return nil, err
	}
	hostNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("hostName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(armcompute.InstanceViewTypes(expandUnescaped))
	var options *armcompute.DedicatedHostsClientGetOptions
	if expandParam != nil {
		options = &armcompute.DedicatedHostsClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := d.srv.Get(req.Context(), resourceGroupNameUnescaped, hostGroupNameUnescaped, hostNameUnescaped, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DedicatedHost, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DedicatedHostsServerTransport) dispatchNewListAvailableSizesPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListAvailableSizesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListAvailableSizesPager not implemented")}
	}
	newListAvailableSizesPager := d.newListAvailableSizesPager.get(req)
	if newListAvailableSizesPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/hostGroups/(?P<hostGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hosts/(?P<hostName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hostSizes`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		hostGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("hostGroupName")])
		if err != nil {
			return nil, err
		}
		hostNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("hostName")])
		if err != nil {
			return nil, err
		}
		resp := d.srv.NewListAvailableSizesPager(resourceGroupNameUnescaped, hostGroupNameUnescaped, hostNameUnescaped, nil)
		newListAvailableSizesPager = &resp
		d.newListAvailableSizesPager.add(req, newListAvailableSizesPager)
	}
	resp, err := server.PagerResponderNext(newListAvailableSizesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListAvailableSizesPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListAvailableSizesPager) {
		d.newListAvailableSizesPager.remove(req)
	}
	return resp, nil
}

func (d *DedicatedHostsServerTransport) dispatchNewListByHostGroupPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListByHostGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByHostGroupPager not implemented")}
	}
	newListByHostGroupPager := d.newListByHostGroupPager.get(req)
	if newListByHostGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/hostGroups/(?P<hostGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hosts`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		hostGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("hostGroupName")])
		if err != nil {
			return nil, err
		}
		resp := d.srv.NewListByHostGroupPager(resourceGroupNameUnescaped, hostGroupNameUnescaped, nil)
		newListByHostGroupPager = &resp
		d.newListByHostGroupPager.add(req, newListByHostGroupPager)
		server.PagerResponderInjectNextLinks(newListByHostGroupPager, req, func(page *armcompute.DedicatedHostsClientListByHostGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByHostGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListByHostGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByHostGroupPager) {
		d.newListByHostGroupPager.remove(req)
	}
	return resp, nil
}

func (d *DedicatedHostsServerTransport) dispatchBeginRestart(req *http.Request) (*http.Response, error) {
	if d.srv.BeginRestart == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRestart not implemented")}
	}
	beginRestart := d.beginRestart.get(req)
	if beginRestart == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/hostGroups/(?P<hostGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hosts/(?P<hostName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restart`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		hostGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("hostGroupName")])
		if err != nil {
			return nil, err
		}
		hostNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("hostName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginRestart(req.Context(), resourceGroupNameUnescaped, hostGroupNameUnescaped, hostNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRestart = &respr
		d.beginRestart.add(req, beginRestart)
	}

	resp, err := server.PollerResponderNext(beginRestart, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.beginRestart.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRestart) {
		d.beginRestart.remove(req)
	}

	return resp, nil
}

func (d *DedicatedHostsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := d.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/hostGroups/(?P<hostGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hosts/(?P<hostName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.DedicatedHostUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		hostGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("hostGroupName")])
		if err != nil {
			return nil, err
		}
		hostNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("hostName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginUpdate(req.Context(), resourceGroupNameUnescaped, hostGroupNameUnescaped, hostNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		d.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		d.beginUpdate.remove(req)
	}

	return resp, nil
}
