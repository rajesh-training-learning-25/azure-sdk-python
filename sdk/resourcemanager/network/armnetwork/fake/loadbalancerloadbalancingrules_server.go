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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v3"
	"net/http"
	"net/url"
	"regexp"
)

// LoadBalancerLoadBalancingRulesServer is a fake server for instances of the armnetwork.LoadBalancerLoadBalancingRulesClient type.
type LoadBalancerLoadBalancingRulesServer struct {
	// Get is the fake for method LoadBalancerLoadBalancingRulesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, loadBalancerName string, loadBalancingRuleName string, options *armnetwork.LoadBalancerLoadBalancingRulesClientGetOptions) (resp azfake.Responder[armnetwork.LoadBalancerLoadBalancingRulesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method LoadBalancerLoadBalancingRulesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, loadBalancerName string, options *armnetwork.LoadBalancerLoadBalancingRulesClientListOptions) (resp azfake.PagerResponder[armnetwork.LoadBalancerLoadBalancingRulesClientListResponse])
}

// NewLoadBalancerLoadBalancingRulesServerTransport creates a new instance of LoadBalancerLoadBalancingRulesServerTransport with the provided implementation.
// The returned LoadBalancerLoadBalancingRulesServerTransport instance is connected to an instance of armnetwork.LoadBalancerLoadBalancingRulesClient by way of the
// undefined.Transporter field.
func NewLoadBalancerLoadBalancingRulesServerTransport(srv *LoadBalancerLoadBalancingRulesServer) *LoadBalancerLoadBalancingRulesServerTransport {
	return &LoadBalancerLoadBalancingRulesServerTransport{srv: srv}
}

// LoadBalancerLoadBalancingRulesServerTransport connects instances of armnetwork.LoadBalancerLoadBalancingRulesClient to instances of LoadBalancerLoadBalancingRulesServer.
// Don't use this type directly, use NewLoadBalancerLoadBalancingRulesServerTransport instead.
type LoadBalancerLoadBalancingRulesServerTransport struct {
	srv          *LoadBalancerLoadBalancingRulesServer
	newListPager *azfake.PagerResponder[armnetwork.LoadBalancerLoadBalancingRulesClientListResponse]
}

// Do implements the policy.Transporter interface for LoadBalancerLoadBalancingRulesServerTransport.
func (l *LoadBalancerLoadBalancingRulesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "LoadBalancerLoadBalancingRulesClient.Get":
		resp, err = l.dispatchGet(req)
	case "LoadBalancerLoadBalancingRulesClient.NewListPager":
		resp, err = l.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (l *LoadBalancerLoadBalancingRulesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if l.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/loadBalancers/(?P<loadBalancerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/loadBalancingRules/(?P<loadBalancingRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	loadBalancerNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("loadBalancerName")])
	if err != nil {
		return nil, err
	}
	loadBalancingRuleNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("loadBalancingRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := l.srv.Get(req.Context(), resourceGroupNameUnescaped, loadBalancerNameUnescaped, loadBalancingRuleNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LoadBalancingRule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LoadBalancerLoadBalancingRulesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	if l.newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/loadBalancers/(?P<loadBalancerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/loadBalancingRules`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		loadBalancerNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("loadBalancerName")])
		if err != nil {
			return nil, err
		}
		resp := l.srv.NewListPager(resourceGroupNameUnescaped, loadBalancerNameUnescaped, nil)
		l.newListPager = &resp
		server.PagerResponderInjectNextLinks(l.newListPager, req, func(page *armnetwork.LoadBalancerLoadBalancingRulesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(l.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(l.newListPager) {
		l.newListPager = nil
	}
	return resp, nil
}
