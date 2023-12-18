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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
	"net/http"
	"net/url"
	"regexp"
)

// ConnectionMonitorsServer is a fake server for instances of the armnetwork.ConnectionMonitorsClient type.
type ConnectionMonitorsServer struct {
	// BeginCreateOrUpdate is the fake for method ConnectionMonitorsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, parameters armnetwork.ConnectionMonitor, options *armnetwork.ConnectionMonitorsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.ConnectionMonitorsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ConnectionMonitorsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *armnetwork.ConnectionMonitorsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.ConnectionMonitorsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ConnectionMonitorsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *armnetwork.ConnectionMonitorsClientGetOptions) (resp azfake.Responder[armnetwork.ConnectionMonitorsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ConnectionMonitorsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, networkWatcherName string, options *armnetwork.ConnectionMonitorsClientListOptions) (resp azfake.PagerResponder[armnetwork.ConnectionMonitorsClientListResponse])

	// BeginQuery is the fake for method ConnectionMonitorsClient.BeginQuery
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginQuery func(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *armnetwork.ConnectionMonitorsClientBeginQueryOptions) (resp azfake.PollerResponder[armnetwork.ConnectionMonitorsClientQueryResponse], errResp azfake.ErrorResponder)

	// BeginStart is the fake for method ConnectionMonitorsClient.BeginStart
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStart func(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *armnetwork.ConnectionMonitorsClientBeginStartOptions) (resp azfake.PollerResponder[armnetwork.ConnectionMonitorsClientStartResponse], errResp azfake.ErrorResponder)

	// BeginStop is the fake for method ConnectionMonitorsClient.BeginStop
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStop func(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *armnetwork.ConnectionMonitorsClientBeginStopOptions) (resp azfake.PollerResponder[armnetwork.ConnectionMonitorsClientStopResponse], errResp azfake.ErrorResponder)

	// UpdateTags is the fake for method ConnectionMonitorsClient.UpdateTags
	// HTTP status codes to indicate success: http.StatusOK
	UpdateTags func(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, parameters armnetwork.TagsObject, options *armnetwork.ConnectionMonitorsClientUpdateTagsOptions) (resp azfake.Responder[armnetwork.ConnectionMonitorsClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewConnectionMonitorsServerTransport creates a new instance of ConnectionMonitorsServerTransport with the provided implementation.
// The returned ConnectionMonitorsServerTransport instance is connected to an instance of armnetwork.ConnectionMonitorsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewConnectionMonitorsServerTransport(srv *ConnectionMonitorsServer) *ConnectionMonitorsServerTransport {
	return &ConnectionMonitorsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armnetwork.ConnectionMonitorsClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armnetwork.ConnectionMonitorsClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armnetwork.ConnectionMonitorsClientListResponse]](),
		beginQuery:          newTracker[azfake.PollerResponder[armnetwork.ConnectionMonitorsClientQueryResponse]](),
		beginStart:          newTracker[azfake.PollerResponder[armnetwork.ConnectionMonitorsClientStartResponse]](),
		beginStop:           newTracker[azfake.PollerResponder[armnetwork.ConnectionMonitorsClientStopResponse]](),
	}
}

// ConnectionMonitorsServerTransport connects instances of armnetwork.ConnectionMonitorsClient to instances of ConnectionMonitorsServer.
// Don't use this type directly, use NewConnectionMonitorsServerTransport instead.
type ConnectionMonitorsServerTransport struct {
	srv                 *ConnectionMonitorsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armnetwork.ConnectionMonitorsClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armnetwork.ConnectionMonitorsClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armnetwork.ConnectionMonitorsClientListResponse]]
	beginQuery          *tracker[azfake.PollerResponder[armnetwork.ConnectionMonitorsClientQueryResponse]]
	beginStart          *tracker[azfake.PollerResponder[armnetwork.ConnectionMonitorsClientStartResponse]]
	beginStop           *tracker[azfake.PollerResponder[armnetwork.ConnectionMonitorsClientStopResponse]]
}

// Do implements the policy.Transporter interface for ConnectionMonitorsServerTransport.
func (c *ConnectionMonitorsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ConnectionMonitorsClient.BeginCreateOrUpdate":
		resp, err = c.dispatchBeginCreateOrUpdate(req)
	case "ConnectionMonitorsClient.BeginDelete":
		resp, err = c.dispatchBeginDelete(req)
	case "ConnectionMonitorsClient.Get":
		resp, err = c.dispatchGet(req)
	case "ConnectionMonitorsClient.NewListPager":
		resp, err = c.dispatchNewListPager(req)
	case "ConnectionMonitorsClient.BeginQuery":
		resp, err = c.dispatchBeginQuery(req)
	case "ConnectionMonitorsClient.BeginStart":
		resp, err = c.dispatchBeginStart(req)
	case "ConnectionMonitorsClient.BeginStop":
		resp, err = c.dispatchBeginStop(req)
	case "ConnectionMonitorsClient.UpdateTags":
		resp, err = c.dispatchUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ConnectionMonitorsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := c.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkWatchers/(?P<networkWatcherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/connectionMonitors/(?P<connectionMonitorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		body, err := server.UnmarshalRequestAsJSON[armnetwork.ConnectionMonitor](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkWatcherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkWatcherName")])
		if err != nil {
			return nil, err
		}
		connectionMonitorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectionMonitorName")])
		if err != nil {
			return nil, err
		}
		migrateUnescaped, err := url.QueryUnescape(qp.Get("migrate"))
		if err != nil {
			return nil, err
		}
		migrateParam := getOptional(migrateUnescaped)
		var options *armnetwork.ConnectionMonitorsClientBeginCreateOrUpdateOptions
		if migrateParam != nil {
			options = &armnetwork.ConnectionMonitorsClientBeginCreateOrUpdateOptions{
				Migrate: migrateParam,
			}
		}
		respr, errRespr := c.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, networkWatcherNameParam, connectionMonitorNameParam, body, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		c.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		c.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		c.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (c *ConnectionMonitorsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := c.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkWatchers/(?P<networkWatcherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/connectionMonitors/(?P<connectionMonitorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkWatcherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkWatcherName")])
		if err != nil {
			return nil, err
		}
		connectionMonitorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectionMonitorName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), resourceGroupNameParam, networkWatcherNameParam, connectionMonitorNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		c.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		c.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		c.beginDelete.remove(req)
	}

	return resp, nil
}

func (c *ConnectionMonitorsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkWatchers/(?P<networkWatcherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/connectionMonitors/(?P<connectionMonitorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkWatcherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkWatcherName")])
	if err != nil {
		return nil, err
	}
	connectionMonitorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectionMonitorName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, networkWatcherNameParam, connectionMonitorNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ConnectionMonitorResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ConnectionMonitorsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := c.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkWatchers/(?P<networkWatcherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/connectionMonitors`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkWatcherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkWatcherName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListPager(resourceGroupNameParam, networkWatcherNameParam, nil)
		newListPager = &resp
		c.newListPager.add(req, newListPager)
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		c.newListPager.remove(req)
	}
	return resp, nil
}

func (c *ConnectionMonitorsServerTransport) dispatchBeginQuery(req *http.Request) (*http.Response, error) {
	if c.srv.BeginQuery == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginQuery not implemented")}
	}
	beginQuery := c.beginQuery.get(req)
	if beginQuery == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkWatchers/(?P<networkWatcherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/connectionMonitors/(?P<connectionMonitorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/query`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkWatcherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkWatcherName")])
		if err != nil {
			return nil, err
		}
		connectionMonitorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectionMonitorName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginQuery(req.Context(), resourceGroupNameParam, networkWatcherNameParam, connectionMonitorNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginQuery = &respr
		c.beginQuery.add(req, beginQuery)
	}

	resp, err := server.PollerResponderNext(beginQuery, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginQuery.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginQuery) {
		c.beginQuery.remove(req)
	}

	return resp, nil
}

func (c *ConnectionMonitorsServerTransport) dispatchBeginStart(req *http.Request) (*http.Response, error) {
	if c.srv.BeginStart == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStart not implemented")}
	}
	beginStart := c.beginStart.get(req)
	if beginStart == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkWatchers/(?P<networkWatcherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/connectionMonitors/(?P<connectionMonitorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/start`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkWatcherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkWatcherName")])
		if err != nil {
			return nil, err
		}
		connectionMonitorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectionMonitorName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginStart(req.Context(), resourceGroupNameParam, networkWatcherNameParam, connectionMonitorNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginStart = &respr
		c.beginStart.add(req, beginStart)
	}

	resp, err := server.PollerResponderNext(beginStart, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginStart.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginStart) {
		c.beginStart.remove(req)
	}

	return resp, nil
}

func (c *ConnectionMonitorsServerTransport) dispatchBeginStop(req *http.Request) (*http.Response, error) {
	if c.srv.BeginStop == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStop not implemented")}
	}
	beginStop := c.beginStop.get(req)
	if beginStop == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkWatchers/(?P<networkWatcherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/connectionMonitors/(?P<connectionMonitorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/stop`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkWatcherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkWatcherName")])
		if err != nil {
			return nil, err
		}
		connectionMonitorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectionMonitorName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginStop(req.Context(), resourceGroupNameParam, networkWatcherNameParam, connectionMonitorNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginStop = &respr
		c.beginStop.add(req, beginStop)
	}

	resp, err := server.PollerResponderNext(beginStop, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginStop.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginStop) {
		c.beginStop.remove(req)
	}

	return resp, nil
}

func (c *ConnectionMonitorsServerTransport) dispatchUpdateTags(req *http.Request) (*http.Response, error) {
	if c.srv.UpdateTags == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateTags not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkWatchers/(?P<networkWatcherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/connectionMonitors/(?P<connectionMonitorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.TagsObject](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkWatcherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkWatcherName")])
	if err != nil {
		return nil, err
	}
	connectionMonitorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectionMonitorName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.UpdateTags(req.Context(), resourceGroupNameParam, networkWatcherNameParam, connectionMonitorNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ConnectionMonitorResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
