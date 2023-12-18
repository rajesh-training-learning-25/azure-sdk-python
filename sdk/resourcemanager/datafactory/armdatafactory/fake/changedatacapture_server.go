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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
	"net/http"
	"net/url"
	"regexp"
)

// ChangeDataCaptureServer is a fake server for instances of the armdatafactory.ChangeDataCaptureClient type.
type ChangeDataCaptureServer struct {
	// CreateOrUpdate is the fake for method ChangeDataCaptureClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, factoryName string, changeDataCaptureName string, changeDataCapture armdatafactory.ChangeDataCaptureResource, options *armdatafactory.ChangeDataCaptureClientCreateOrUpdateOptions) (resp azfake.Responder[armdatafactory.ChangeDataCaptureClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ChangeDataCaptureClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, factoryName string, changeDataCaptureName string, options *armdatafactory.ChangeDataCaptureClientDeleteOptions) (resp azfake.Responder[armdatafactory.ChangeDataCaptureClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ChangeDataCaptureClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, factoryName string, changeDataCaptureName string, options *armdatafactory.ChangeDataCaptureClientGetOptions) (resp azfake.Responder[armdatafactory.ChangeDataCaptureClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByFactoryPager is the fake for method ChangeDataCaptureClient.NewListByFactoryPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByFactoryPager func(resourceGroupName string, factoryName string, options *armdatafactory.ChangeDataCaptureClientListByFactoryOptions) (resp azfake.PagerResponder[armdatafactory.ChangeDataCaptureClientListByFactoryResponse])

	// Start is the fake for method ChangeDataCaptureClient.Start
	// HTTP status codes to indicate success: http.StatusOK
	Start func(ctx context.Context, resourceGroupName string, factoryName string, changeDataCaptureName string, options *armdatafactory.ChangeDataCaptureClientStartOptions) (resp azfake.Responder[armdatafactory.ChangeDataCaptureClientStartResponse], errResp azfake.ErrorResponder)

	// Status is the fake for method ChangeDataCaptureClient.Status
	// HTTP status codes to indicate success: http.StatusOK
	Status func(ctx context.Context, resourceGroupName string, factoryName string, changeDataCaptureName string, options *armdatafactory.ChangeDataCaptureClientStatusOptions) (resp azfake.Responder[armdatafactory.ChangeDataCaptureClientStatusResponse], errResp azfake.ErrorResponder)

	// Stop is the fake for method ChangeDataCaptureClient.Stop
	// HTTP status codes to indicate success: http.StatusOK
	Stop func(ctx context.Context, resourceGroupName string, factoryName string, changeDataCaptureName string, options *armdatafactory.ChangeDataCaptureClientStopOptions) (resp azfake.Responder[armdatafactory.ChangeDataCaptureClientStopResponse], errResp azfake.ErrorResponder)
}

// NewChangeDataCaptureServerTransport creates a new instance of ChangeDataCaptureServerTransport with the provided implementation.
// The returned ChangeDataCaptureServerTransport instance is connected to an instance of armdatafactory.ChangeDataCaptureClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewChangeDataCaptureServerTransport(srv *ChangeDataCaptureServer) *ChangeDataCaptureServerTransport {
	return &ChangeDataCaptureServerTransport{
		srv:                   srv,
		newListByFactoryPager: newTracker[azfake.PagerResponder[armdatafactory.ChangeDataCaptureClientListByFactoryResponse]](),
	}
}

// ChangeDataCaptureServerTransport connects instances of armdatafactory.ChangeDataCaptureClient to instances of ChangeDataCaptureServer.
// Don't use this type directly, use NewChangeDataCaptureServerTransport instead.
type ChangeDataCaptureServerTransport struct {
	srv                   *ChangeDataCaptureServer
	newListByFactoryPager *tracker[azfake.PagerResponder[armdatafactory.ChangeDataCaptureClientListByFactoryResponse]]
}

// Do implements the policy.Transporter interface for ChangeDataCaptureServerTransport.
func (c *ChangeDataCaptureServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ChangeDataCaptureClient.CreateOrUpdate":
		resp, err = c.dispatchCreateOrUpdate(req)
	case "ChangeDataCaptureClient.Delete":
		resp, err = c.dispatchDelete(req)
	case "ChangeDataCaptureClient.Get":
		resp, err = c.dispatchGet(req)
	case "ChangeDataCaptureClient.NewListByFactoryPager":
		resp, err = c.dispatchNewListByFactoryPager(req)
	case "ChangeDataCaptureClient.Start":
		resp, err = c.dispatchStart(req)
	case "ChangeDataCaptureClient.Status":
		resp, err = c.dispatchStatus(req)
	case "ChangeDataCaptureClient.Stop":
		resp, err = c.dispatchStop(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ChangeDataCaptureServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/adfcdcs/(?P<changeDataCaptureName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdatafactory.ChangeDataCaptureResource](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
	if err != nil {
		return nil, err
	}
	changeDataCaptureNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("changeDataCaptureName")])
	if err != nil {
		return nil, err
	}
	ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
	var options *armdatafactory.ChangeDataCaptureClientCreateOrUpdateOptions
	if ifMatchParam != nil {
		options = &armdatafactory.ChangeDataCaptureClientCreateOrUpdateOptions{
			IfMatch: ifMatchParam,
		}
	}
	respr, errRespr := c.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, factoryNameParam, changeDataCaptureNameParam, body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ChangeDataCaptureResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ChangeDataCaptureServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if c.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/adfcdcs/(?P<changeDataCaptureName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
	if err != nil {
		return nil, err
	}
	changeDataCaptureNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("changeDataCaptureName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Delete(req.Context(), resourceGroupNameParam, factoryNameParam, changeDataCaptureNameParam, nil)
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

func (c *ChangeDataCaptureServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/adfcdcs/(?P<changeDataCaptureName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
	if err != nil {
		return nil, err
	}
	changeDataCaptureNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("changeDataCaptureName")])
	if err != nil {
		return nil, err
	}
	ifNoneMatchParam := getOptional(getHeaderValue(req.Header, "If-None-Match"))
	var options *armdatafactory.ChangeDataCaptureClientGetOptions
	if ifNoneMatchParam != nil {
		options = &armdatafactory.ChangeDataCaptureClientGetOptions{
			IfNoneMatch: ifNoneMatchParam,
		}
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, factoryNameParam, changeDataCaptureNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ChangeDataCaptureResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ChangeDataCaptureServerTransport) dispatchNewListByFactoryPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListByFactoryPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByFactoryPager not implemented")}
	}
	newListByFactoryPager := c.newListByFactoryPager.get(req)
	if newListByFactoryPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/adfcdcs`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListByFactoryPager(resourceGroupNameParam, factoryNameParam, nil)
		newListByFactoryPager = &resp
		c.newListByFactoryPager.add(req, newListByFactoryPager)
		server.PagerResponderInjectNextLinks(newListByFactoryPager, req, func(page *armdatafactory.ChangeDataCaptureClientListByFactoryResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByFactoryPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListByFactoryPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByFactoryPager) {
		c.newListByFactoryPager.remove(req)
	}
	return resp, nil
}

func (c *ChangeDataCaptureServerTransport) dispatchStart(req *http.Request) (*http.Response, error) {
	if c.srv.Start == nil {
		return nil, &nonRetriableError{errors.New("fake for method Start not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/adfcdcs/(?P<changeDataCaptureName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/start`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
	if err != nil {
		return nil, err
	}
	changeDataCaptureNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("changeDataCaptureName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Start(req.Context(), resourceGroupNameParam, factoryNameParam, changeDataCaptureNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ChangeDataCaptureServerTransport) dispatchStatus(req *http.Request) (*http.Response, error) {
	if c.srv.Status == nil {
		return nil, &nonRetriableError{errors.New("fake for method Status not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/adfcdcs/(?P<changeDataCaptureName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/status`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
	if err != nil {
		return nil, err
	}
	changeDataCaptureNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("changeDataCaptureName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Status(req.Context(), resourceGroupNameParam, factoryNameParam, changeDataCaptureNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ChangeDataCaptureServerTransport) dispatchStop(req *http.Request) (*http.Response, error) {
	if c.srv.Stop == nil {
		return nil, &nonRetriableError{errors.New("fake for method Stop not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/adfcdcs/(?P<changeDataCaptureName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/stop`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
	if err != nil {
		return nil, err
	}
	changeDataCaptureNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("changeDataCaptureName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Stop(req.Context(), resourceGroupNameParam, factoryNameParam, changeDataCaptureNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
