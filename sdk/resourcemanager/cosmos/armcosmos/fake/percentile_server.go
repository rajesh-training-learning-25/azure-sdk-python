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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
	"net/http"
	"net/url"
	"regexp"
)

// PercentileServer is a fake server for instances of the armcosmos.PercentileClient type.
type PercentileServer struct {
	// NewListMetricsPager is the fake for method PercentileClient.NewListMetricsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListMetricsPager func(resourceGroupName string, accountName string, filter string, options *armcosmos.PercentileClientListMetricsOptions) (resp azfake.PagerResponder[armcosmos.PercentileClientListMetricsResponse])
}

// NewPercentileServerTransport creates a new instance of PercentileServerTransport with the provided implementation.
// The returned PercentileServerTransport instance is connected to an instance of armcosmos.PercentileClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPercentileServerTransport(srv *PercentileServer) *PercentileServerTransport {
	return &PercentileServerTransport{
		srv:                 srv,
		newListMetricsPager: newTracker[azfake.PagerResponder[armcosmos.PercentileClientListMetricsResponse]](),
	}
}

// PercentileServerTransport connects instances of armcosmos.PercentileClient to instances of PercentileServer.
// Don't use this type directly, use NewPercentileServerTransport instead.
type PercentileServerTransport struct {
	srv                 *PercentileServer
	newListMetricsPager *tracker[azfake.PagerResponder[armcosmos.PercentileClientListMetricsResponse]]
}

// Do implements the policy.Transporter interface for PercentileServerTransport.
func (p *PercentileServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PercentileClient.NewListMetricsPager":
		resp, err = p.dispatchNewListMetricsPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PercentileServerTransport) dispatchNewListMetricsPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListMetricsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListMetricsPager not implemented")}
	}
	newListMetricsPager := p.newListMetricsPager.get(req)
	if newListMetricsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/databaseAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/percentile/metrics`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
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
		filterParam, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListMetricsPager(resourceGroupNameParam, accountNameParam, filterParam, nil)
		newListMetricsPager = &resp
		p.newListMetricsPager.add(req, newListMetricsPager)
	}
	resp, err := server.PagerResponderNext(newListMetricsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListMetricsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListMetricsPager) {
		p.newListMetricsPager.remove(req)
	}
	return resp, nil
}
