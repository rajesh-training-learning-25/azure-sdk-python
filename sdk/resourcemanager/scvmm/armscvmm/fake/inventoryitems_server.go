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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/scvmm/armscvmm"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// InventoryItemsServer is a fake server for instances of the armscvmm.InventoryItemsClient type.
type InventoryItemsServer struct {
	// Create is the fake for method InventoryItemsClient.Create
<<<<<<< HEAD
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	Create func(ctx context.Context, resourceGroupName string, vmmServerName string, inventoryItemResourceName string, options *armscvmm.InventoryItemsClientCreateOptions) (resp azfake.Responder[armscvmm.InventoryItemsClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method InventoryItemsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, vmmServerName string, inventoryItemResourceName string, options *armscvmm.InventoryItemsClientDeleteOptions) (resp azfake.Responder[armscvmm.InventoryItemsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method InventoryItemsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, vmmServerName string, inventoryItemResourceName string, options *armscvmm.InventoryItemsClientGetOptions) (resp azfake.Responder[armscvmm.InventoryItemsClientGetResponse], errResp azfake.ErrorResponder)
=======
	// HTTP status codes to indicate success: http.StatusOK
	Create func(ctx context.Context, resourceGroupName string, vmmServerName string, inventoryItemName string, options *armscvmm.InventoryItemsClientCreateOptions) (resp azfake.Responder[armscvmm.InventoryItemsClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method InventoryItemsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, vmmServerName string, inventoryItemName string, options *armscvmm.InventoryItemsClientDeleteOptions) (resp azfake.Responder[armscvmm.InventoryItemsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method InventoryItemsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, vmmServerName string, inventoryItemName string, options *armscvmm.InventoryItemsClientGetOptions) (resp azfake.Responder[armscvmm.InventoryItemsClientGetResponse], errResp azfake.ErrorResponder)
>>>>>>> 2621632e48ea508e16ce568001402f92fca4afa0

	// NewListByVMMServerPager is the fake for method InventoryItemsClient.NewListByVMMServerPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByVMMServerPager func(resourceGroupName string, vmmServerName string, options *armscvmm.InventoryItemsClientListByVMMServerOptions) (resp azfake.PagerResponder[armscvmm.InventoryItemsClientListByVMMServerResponse])
}

// NewInventoryItemsServerTransport creates a new instance of InventoryItemsServerTransport with the provided implementation.
// The returned InventoryItemsServerTransport instance is connected to an instance of armscvmm.InventoryItemsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewInventoryItemsServerTransport(srv *InventoryItemsServer) *InventoryItemsServerTransport {
	return &InventoryItemsServerTransport{
		srv:                     srv,
		newListByVMMServerPager: newTracker[azfake.PagerResponder[armscvmm.InventoryItemsClientListByVMMServerResponse]](),
	}
}

// InventoryItemsServerTransport connects instances of armscvmm.InventoryItemsClient to instances of InventoryItemsServer.
// Don't use this type directly, use NewInventoryItemsServerTransport instead.
type InventoryItemsServerTransport struct {
	srv                     *InventoryItemsServer
	newListByVMMServerPager *tracker[azfake.PagerResponder[armscvmm.InventoryItemsClientListByVMMServerResponse]]
}

// Do implements the policy.Transporter interface for InventoryItemsServerTransport.
func (i *InventoryItemsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "InventoryItemsClient.Create":
		resp, err = i.dispatchCreate(req)
	case "InventoryItemsClient.Delete":
		resp, err = i.dispatchDelete(req)
	case "InventoryItemsClient.Get":
		resp, err = i.dispatchGet(req)
	case "InventoryItemsClient.NewListByVMMServerPager":
		resp, err = i.dispatchNewListByVMMServerPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *InventoryItemsServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if i.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
<<<<<<< HEAD
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ScVmm/vmmServers/(?P<vmmServerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/inventoryItems/(?P<inventoryItemResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
=======
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ScVmm/vmmServers/(?P<vmmServerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/inventoryItems/(?P<inventoryItemName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
>>>>>>> 2621632e48ea508e16ce568001402f92fca4afa0
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armscvmm.InventoryItem](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	vmmServerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vmmServerName")])
	if err != nil {
		return nil, err
	}
<<<<<<< HEAD
	inventoryItemResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("inventoryItemResourceName")])
=======
	inventoryItemNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("inventoryItemName")])
>>>>>>> 2621632e48ea508e16ce568001402f92fca4afa0
	if err != nil {
		return nil, err
	}
	var options *armscvmm.InventoryItemsClientCreateOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armscvmm.InventoryItemsClientCreateOptions{
			Body: &body,
		}
	}
<<<<<<< HEAD
	respr, errRespr := i.srv.Create(req.Context(), resourceGroupNameParam, vmmServerNameParam, inventoryItemResourceNameParam, options)
=======
	respr, errRespr := i.srv.Create(req.Context(), resourceGroupNameParam, vmmServerNameParam, inventoryItemNameParam, options)
>>>>>>> 2621632e48ea508e16ce568001402f92fca4afa0
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
<<<<<<< HEAD
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
=======
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
>>>>>>> 2621632e48ea508e16ce568001402f92fca4afa0
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).InventoryItem, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *InventoryItemsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if i.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
<<<<<<< HEAD
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ScVmm/vmmServers/(?P<vmmServerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/inventoryItems/(?P<inventoryItemResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
=======
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ScVmm/vmmServers/(?P<vmmServerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/inventoryItems/(?P<inventoryItemName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
>>>>>>> 2621632e48ea508e16ce568001402f92fca4afa0
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	vmmServerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vmmServerName")])
	if err != nil {
		return nil, err
	}
<<<<<<< HEAD
	inventoryItemResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("inventoryItemResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Delete(req.Context(), resourceGroupNameParam, vmmServerNameParam, inventoryItemResourceNameParam, nil)
=======
	inventoryItemNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("inventoryItemName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Delete(req.Context(), resourceGroupNameParam, vmmServerNameParam, inventoryItemNameParam, nil)
>>>>>>> 2621632e48ea508e16ce568001402f92fca4afa0
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

func (i *InventoryItemsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
<<<<<<< HEAD
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ScVmm/vmmServers/(?P<vmmServerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/inventoryItems/(?P<inventoryItemResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
=======
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ScVmm/vmmServers/(?P<vmmServerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/inventoryItems/(?P<inventoryItemName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
>>>>>>> 2621632e48ea508e16ce568001402f92fca4afa0
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	vmmServerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vmmServerName")])
	if err != nil {
		return nil, err
	}
<<<<<<< HEAD
	inventoryItemResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("inventoryItemResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Get(req.Context(), resourceGroupNameParam, vmmServerNameParam, inventoryItemResourceNameParam, nil)
=======
	inventoryItemNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("inventoryItemName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Get(req.Context(), resourceGroupNameParam, vmmServerNameParam, inventoryItemNameParam, nil)
>>>>>>> 2621632e48ea508e16ce568001402f92fca4afa0
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).InventoryItem, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *InventoryItemsServerTransport) dispatchNewListByVMMServerPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListByVMMServerPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByVMMServerPager not implemented")}
	}
	newListByVMMServerPager := i.newListByVMMServerPager.get(req)
	if newListByVMMServerPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ScVmm/vmmServers/(?P<vmmServerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/inventoryItems`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vmmServerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vmmServerName")])
		if err != nil {
			return nil, err
		}
		resp := i.srv.NewListByVMMServerPager(resourceGroupNameParam, vmmServerNameParam, nil)
		newListByVMMServerPager = &resp
		i.newListByVMMServerPager.add(req, newListByVMMServerPager)
		server.PagerResponderInjectNextLinks(newListByVMMServerPager, req, func(page *armscvmm.InventoryItemsClientListByVMMServerResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByVMMServerPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListByVMMServerPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByVMMServerPager) {
		i.newListByVMMServerPager.remove(req)
	}
	return resp, nil
}
