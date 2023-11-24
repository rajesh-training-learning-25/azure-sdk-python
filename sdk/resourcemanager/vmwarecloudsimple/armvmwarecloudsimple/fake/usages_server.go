//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple"
	"net/http"
	"net/url"
	"regexp"
)

// UsagesServer is a fake server for instances of the armvmwarecloudsimple.UsagesClient type.
type UsagesServer struct {
	// NewListPager is the fake for method UsagesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(regionID string, options *armvmwarecloudsimple.UsagesClientListOptions) (resp azfake.PagerResponder[armvmwarecloudsimple.UsagesClientListResponse])
}

// NewUsagesServerTransport creates a new instance of UsagesServerTransport with the provided implementation.
// The returned UsagesServerTransport instance is connected to an instance of armvmwarecloudsimple.UsagesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewUsagesServerTransport(srv *UsagesServer) *UsagesServerTransport {
	return &UsagesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armvmwarecloudsimple.UsagesClientListResponse]](),
	}
}

// UsagesServerTransport connects instances of armvmwarecloudsimple.UsagesClient to instances of UsagesServer.
// Don't use this type directly, use NewUsagesServerTransport instead.
type UsagesServerTransport struct {
	srv          *UsagesServer
	newListPager *tracker[azfake.PagerResponder[armvmwarecloudsimple.UsagesClientListResponse]]
}

// Do implements the policy.Transporter interface for UsagesServerTransport.
func (u *UsagesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "UsagesClient.NewListPager":
		resp, err = u.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (u *UsagesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if u.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := u.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.VMwareCloudSimple/locations/(?P<regionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/usages`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		regionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("regionId")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armvmwarecloudsimple.UsagesClientListOptions
		if filterParam != nil {
			options = &armvmwarecloudsimple.UsagesClientListOptions{
				Filter: filterParam,
			}
		}
		resp := u.srv.NewListPager(regionIDParam, options)
		newListPager = &resp
		u.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armvmwarecloudsimple.UsagesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		u.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		u.newListPager.remove(req)
	}
	return resp, nil
}
