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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"net/http"
	"net/url"
	"regexp"
)

// ElasticPoolActivitiesServer is a fake server for instances of the armsql.ElasticPoolActivitiesClient type.
type ElasticPoolActivitiesServer struct {
	// NewListByElasticPoolPager is the fake for method ElasticPoolActivitiesClient.NewListByElasticPoolPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByElasticPoolPager func(resourceGroupName string, serverName string, elasticPoolName string, options *armsql.ElasticPoolActivitiesClientListByElasticPoolOptions) (resp azfake.PagerResponder[armsql.ElasticPoolActivitiesClientListByElasticPoolResponse])
}

// NewElasticPoolActivitiesServerTransport creates a new instance of ElasticPoolActivitiesServerTransport with the provided implementation.
// The returned ElasticPoolActivitiesServerTransport instance is connected to an instance of armsql.ElasticPoolActivitiesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewElasticPoolActivitiesServerTransport(srv *ElasticPoolActivitiesServer) *ElasticPoolActivitiesServerTransport {
	return &ElasticPoolActivitiesServerTransport{
		srv:                       srv,
		newListByElasticPoolPager: newTracker[azfake.PagerResponder[armsql.ElasticPoolActivitiesClientListByElasticPoolResponse]](),
	}
}

// ElasticPoolActivitiesServerTransport connects instances of armsql.ElasticPoolActivitiesClient to instances of ElasticPoolActivitiesServer.
// Don't use this type directly, use NewElasticPoolActivitiesServerTransport instead.
type ElasticPoolActivitiesServerTransport struct {
	srv                       *ElasticPoolActivitiesServer
	newListByElasticPoolPager *tracker[azfake.PagerResponder[armsql.ElasticPoolActivitiesClientListByElasticPoolResponse]]
}

// Do implements the policy.Transporter interface for ElasticPoolActivitiesServerTransport.
func (e *ElasticPoolActivitiesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ElasticPoolActivitiesClient.NewListByElasticPoolPager":
		resp, err = e.dispatchNewListByElasticPoolPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *ElasticPoolActivitiesServerTransport) dispatchNewListByElasticPoolPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListByElasticPoolPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByElasticPoolPager not implemented")}
	}
	newListByElasticPoolPager := e.newListByElasticPoolPager.get(req)
	if newListByElasticPoolPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/elasticPools/(?P<elasticPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/elasticPoolActivity`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		elasticPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("elasticPoolName")])
		if err != nil {
			return nil, err
		}
		resp := e.srv.NewListByElasticPoolPager(resourceGroupNameParam, serverNameParam, elasticPoolNameParam, nil)
		newListByElasticPoolPager = &resp
		e.newListByElasticPoolPager.add(req, newListByElasticPoolPager)
	}
	resp, err := server.PagerResponderNext(newListByElasticPoolPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newListByElasticPoolPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByElasticPoolPager) {
		e.newListByElasticPoolPager.remove(req)
	}
	return resp, nil
}
