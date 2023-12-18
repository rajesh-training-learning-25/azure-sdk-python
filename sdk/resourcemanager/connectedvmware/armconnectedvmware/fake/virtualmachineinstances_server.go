//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// Code generated by @autorest/go. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
)

// VirtualMachineInstancesServer is a fake server for instances of the armconnectedvmware.VirtualMachineInstancesClient type.
type VirtualMachineInstancesServer struct {
	// BeginCreateOrUpdate is the fake for method VirtualMachineInstancesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceURI string, body armconnectedvmware.VirtualMachineInstance, options *armconnectedvmware.VirtualMachineInstancesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armconnectedvmware.VirtualMachineInstancesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method VirtualMachineInstancesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceURI string, options *armconnectedvmware.VirtualMachineInstancesClientBeginDeleteOptions) (resp azfake.PollerResponder[armconnectedvmware.VirtualMachineInstancesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VirtualMachineInstancesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceURI string, options *armconnectedvmware.VirtualMachineInstancesClientGetOptions) (resp azfake.Responder[armconnectedvmware.VirtualMachineInstancesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method VirtualMachineInstancesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceURI string, options *armconnectedvmware.VirtualMachineInstancesClientListOptions) (resp azfake.PagerResponder[armconnectedvmware.VirtualMachineInstancesClientListResponse])

	// BeginRestart is the fake for method VirtualMachineInstancesClient.BeginRestart
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginRestart func(ctx context.Context, resourceURI string, options *armconnectedvmware.VirtualMachineInstancesClientBeginRestartOptions) (resp azfake.PollerResponder[armconnectedvmware.VirtualMachineInstancesClientRestartResponse], errResp azfake.ErrorResponder)

	// BeginStart is the fake for method VirtualMachineInstancesClient.BeginStart
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginStart func(ctx context.Context, resourceURI string, options *armconnectedvmware.VirtualMachineInstancesClientBeginStartOptions) (resp azfake.PollerResponder[armconnectedvmware.VirtualMachineInstancesClientStartResponse], errResp azfake.ErrorResponder)

	// BeginStop is the fake for method VirtualMachineInstancesClient.BeginStop
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginStop func(ctx context.Context, resourceURI string, options *armconnectedvmware.VirtualMachineInstancesClientBeginStopOptions) (resp azfake.PollerResponder[armconnectedvmware.VirtualMachineInstancesClientStopResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method VirtualMachineInstancesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceURI string, body armconnectedvmware.VirtualMachineInstanceUpdate, options *armconnectedvmware.VirtualMachineInstancesClientBeginUpdateOptions) (resp azfake.PollerResponder[armconnectedvmware.VirtualMachineInstancesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewVirtualMachineInstancesServerTransport creates a new instance of VirtualMachineInstancesServerTransport with the provided implementation.
// The returned VirtualMachineInstancesServerTransport instance is connected to an instance of armconnectedvmware.VirtualMachineInstancesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVirtualMachineInstancesServerTransport(srv *VirtualMachineInstancesServer) *VirtualMachineInstancesServerTransport {
	return &VirtualMachineInstancesServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armconnectedvmware.VirtualMachineInstancesClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armconnectedvmware.VirtualMachineInstancesClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armconnectedvmware.VirtualMachineInstancesClientListResponse]](),
		beginRestart:        newTracker[azfake.PollerResponder[armconnectedvmware.VirtualMachineInstancesClientRestartResponse]](),
		beginStart:          newTracker[azfake.PollerResponder[armconnectedvmware.VirtualMachineInstancesClientStartResponse]](),
		beginStop:           newTracker[azfake.PollerResponder[armconnectedvmware.VirtualMachineInstancesClientStopResponse]](),
		beginUpdate:         newTracker[azfake.PollerResponder[armconnectedvmware.VirtualMachineInstancesClientUpdateResponse]](),
	}
}

// VirtualMachineInstancesServerTransport connects instances of armconnectedvmware.VirtualMachineInstancesClient to instances of VirtualMachineInstancesServer.
// Don't use this type directly, use NewVirtualMachineInstancesServerTransport instead.
type VirtualMachineInstancesServerTransport struct {
	srv                 *VirtualMachineInstancesServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armconnectedvmware.VirtualMachineInstancesClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armconnectedvmware.VirtualMachineInstancesClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armconnectedvmware.VirtualMachineInstancesClientListResponse]]
	beginRestart        *tracker[azfake.PollerResponder[armconnectedvmware.VirtualMachineInstancesClientRestartResponse]]
	beginStart          *tracker[azfake.PollerResponder[armconnectedvmware.VirtualMachineInstancesClientStartResponse]]
	beginStop           *tracker[azfake.PollerResponder[armconnectedvmware.VirtualMachineInstancesClientStopResponse]]
	beginUpdate         *tracker[azfake.PollerResponder[armconnectedvmware.VirtualMachineInstancesClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for VirtualMachineInstancesServerTransport.
func (v *VirtualMachineInstancesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VirtualMachineInstancesClient.BeginCreateOrUpdate":
		resp, err = v.dispatchBeginCreateOrUpdate(req)
	case "VirtualMachineInstancesClient.BeginDelete":
		resp, err = v.dispatchBeginDelete(req)
	case "VirtualMachineInstancesClient.Get":
		resp, err = v.dispatchGet(req)
	case "VirtualMachineInstancesClient.NewListPager":
		resp, err = v.dispatchNewListPager(req)
	case "VirtualMachineInstancesClient.BeginRestart":
		resp, err = v.dispatchBeginRestart(req)
	case "VirtualMachineInstancesClient.BeginStart":
		resp, err = v.dispatchBeginStart(req)
	case "VirtualMachineInstancesClient.BeginStop":
		resp, err = v.dispatchBeginStop(req)
	case "VirtualMachineInstancesClient.BeginUpdate":
		resp, err = v.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VirtualMachineInstancesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := v.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ConnectedVMwarevSphere/virtualMachineInstances/default`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armconnectedvmware.VirtualMachineInstance](req)
		if err != nil {
			return nil, err
		}
		resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginCreateOrUpdate(req.Context(), resourceURIParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		v.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		v.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		v.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (v *VirtualMachineInstancesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := v.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ConnectedVMwarevSphere/virtualMachineInstances/default`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
		if err != nil {
			return nil, err
		}
		deleteFromHostUnescaped, err := url.QueryUnescape(qp.Get("deleteFromHost"))
		if err != nil {
			return nil, err
		}
		deleteFromHostParam, err := parseOptional(deleteFromHostUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		forceUnescaped, err := url.QueryUnescape(qp.Get("force"))
		if err != nil {
			return nil, err
		}
		forceParam, err := parseOptional(forceUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		var options *armconnectedvmware.VirtualMachineInstancesClientBeginDeleteOptions
		if deleteFromHostParam != nil || forceParam != nil {
			options = &armconnectedvmware.VirtualMachineInstancesClientBeginDeleteOptions{
				DeleteFromHost: deleteFromHostParam,
				Force:          forceParam,
			}
		}
		respr, errRespr := v.srv.BeginDelete(req.Context(), resourceURIParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		v.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		v.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		v.beginDelete.remove(req)
	}

	return resp, nil
}

func (v *VirtualMachineInstancesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ConnectedVMwarevSphere/virtualMachineInstances/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Get(req.Context(), resourceURIParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualMachineInstance, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualMachineInstancesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := v.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ConnectedVMwarevSphere/virtualMachineInstances`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
		if err != nil {
			return nil, err
		}
		resp := v.srv.NewListPager(resourceURIParam, nil)
		newListPager = &resp
		v.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armconnectedvmware.VirtualMachineInstancesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		v.newListPager.remove(req)
	}
	return resp, nil
}

func (v *VirtualMachineInstancesServerTransport) dispatchBeginRestart(req *http.Request) (*http.Response, error) {
	if v.srv.BeginRestart == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRestart not implemented")}
	}
	beginRestart := v.beginRestart.get(req)
	if beginRestart == nil {
		const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ConnectedVMwarevSphere/virtualMachineInstances/default/restart`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginRestart(req.Context(), resourceURIParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRestart = &respr
		v.beginRestart.add(req, beginRestart)
	}

	resp, err := server.PollerResponderNext(beginRestart, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		v.beginRestart.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRestart) {
		v.beginRestart.remove(req)
	}

	return resp, nil
}

func (v *VirtualMachineInstancesServerTransport) dispatchBeginStart(req *http.Request) (*http.Response, error) {
	if v.srv.BeginStart == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStart not implemented")}
	}
	beginStart := v.beginStart.get(req)
	if beginStart == nil {
		const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ConnectedVMwarevSphere/virtualMachineInstances/default/start`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginStart(req.Context(), resourceURIParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginStart = &respr
		v.beginStart.add(req, beginStart)
	}

	resp, err := server.PollerResponderNext(beginStart, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		v.beginStart.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginStart) {
		v.beginStart.remove(req)
	}

	return resp, nil
}

func (v *VirtualMachineInstancesServerTransport) dispatchBeginStop(req *http.Request) (*http.Response, error) {
	if v.srv.BeginStop == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStop not implemented")}
	}
	beginStop := v.beginStop.get(req)
	if beginStop == nil {
		const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ConnectedVMwarevSphere/virtualMachineInstances/default/stop`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armconnectedvmware.StopVirtualMachineOptions](req)
		if err != nil {
			return nil, err
		}
		resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
		if err != nil {
			return nil, err
		}
		var options *armconnectedvmware.VirtualMachineInstancesClientBeginStopOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armconnectedvmware.VirtualMachineInstancesClientBeginStopOptions{
				Body: &body,
			}
		}
		respr, errRespr := v.srv.BeginStop(req.Context(), resourceURIParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginStop = &respr
		v.beginStop.add(req, beginStop)
	}

	resp, err := server.PollerResponderNext(beginStop, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		v.beginStop.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginStop) {
		v.beginStop.remove(req)
	}

	return resp, nil
}

func (v *VirtualMachineInstancesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := v.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ConnectedVMwarevSphere/virtualMachineInstances/default`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armconnectedvmware.VirtualMachineInstanceUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginUpdate(req.Context(), resourceURIParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		v.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		v.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		v.beginUpdate.remove(req)
	}

	return resp, nil
}
