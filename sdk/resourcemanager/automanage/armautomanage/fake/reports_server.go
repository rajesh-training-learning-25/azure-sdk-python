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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automanage/armautomanage"
	"net/http"
	"net/url"
	"regexp"
)

// ReportsServer is a fake server for instances of the armautomanage.ReportsClient type.
type ReportsServer struct {
	// Get is the fake for method ReportsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, configurationProfileAssignmentName string, reportName string, vmName string, options *armautomanage.ReportsClientGetOptions) (resp azfake.Responder[armautomanage.ReportsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByConfigurationProfileAssignmentsPager is the fake for method ReportsClient.NewListByConfigurationProfileAssignmentsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByConfigurationProfileAssignmentsPager func(resourceGroupName string, configurationProfileAssignmentName string, vmName string, options *armautomanage.ReportsClientListByConfigurationProfileAssignmentsOptions) (resp azfake.PagerResponder[armautomanage.ReportsClientListByConfigurationProfileAssignmentsResponse])
}

// NewReportsServerTransport creates a new instance of ReportsServerTransport with the provided implementation.
// The returned ReportsServerTransport instance is connected to an instance of armautomanage.ReportsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewReportsServerTransport(srv *ReportsServer) *ReportsServerTransport {
	return &ReportsServerTransport{
		srv: srv,
		newListByConfigurationProfileAssignmentsPager: newTracker[azfake.PagerResponder[armautomanage.ReportsClientListByConfigurationProfileAssignmentsResponse]](),
	}
}

// ReportsServerTransport connects instances of armautomanage.ReportsClient to instances of ReportsServer.
// Don't use this type directly, use NewReportsServerTransport instead.
type ReportsServerTransport struct {
	srv                                           *ReportsServer
	newListByConfigurationProfileAssignmentsPager *tracker[azfake.PagerResponder[armautomanage.ReportsClientListByConfigurationProfileAssignmentsResponse]]
}

// Do implements the policy.Transporter interface for ReportsServerTransport.
func (r *ReportsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ReportsClient.Get":
		resp, err = r.dispatchGet(req)
	case "ReportsClient.NewListByConfigurationProfileAssignmentsPager":
		resp, err = r.dispatchNewListByConfigurationProfileAssignmentsPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *ReportsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/virtualMachines/(?P<vmName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automanage/configurationProfileAssignments/(?P<configurationProfileAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reports/(?P<reportName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	configurationProfileAssignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configurationProfileAssignmentName")])
	if err != nil {
		return nil, err
	}
	reportNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("reportName")])
	if err != nil {
		return nil, err
	}
	vmNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vmName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceGroupNameParam, configurationProfileAssignmentNameParam, reportNameParam, vmNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Report, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *ReportsServerTransport) dispatchNewListByConfigurationProfileAssignmentsPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListByConfigurationProfileAssignmentsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByConfigurationProfileAssignmentsPager not implemented")}
	}
	newListByConfigurationProfileAssignmentsPager := r.newListByConfigurationProfileAssignmentsPager.get(req)
	if newListByConfigurationProfileAssignmentsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/virtualMachines/(?P<vmName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automanage/configurationProfileAssignments/(?P<configurationProfileAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reports`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		configurationProfileAssignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configurationProfileAssignmentName")])
		if err != nil {
			return nil, err
		}
		vmNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vmName")])
		if err != nil {
			return nil, err
		}
		resp := r.srv.NewListByConfigurationProfileAssignmentsPager(resourceGroupNameParam, configurationProfileAssignmentNameParam, vmNameParam, nil)
		newListByConfigurationProfileAssignmentsPager = &resp
		r.newListByConfigurationProfileAssignmentsPager.add(req, newListByConfigurationProfileAssignmentsPager)
	}
	resp, err := server.PagerResponderNext(newListByConfigurationProfileAssignmentsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListByConfigurationProfileAssignmentsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByConfigurationProfileAssignmentsPager) {
		r.newListByConfigurationProfileAssignmentsPager.remove(req)
	}
	return resp, nil
}
