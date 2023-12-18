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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservicesdatareplication/armrecoveryservicesdatareplication"
	"net/http"
	"net/url"
	"regexp"
)

// RecoveryPointsServer is a fake server for instances of the armrecoveryservicesdatareplication.RecoveryPointsClient type.
type RecoveryPointsServer struct {
	// Get is the fake for method RecoveryPointsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, vaultName string, protectedItemName string, recoveryPointName string, options *armrecoveryservicesdatareplication.RecoveryPointsClientGetOptions) (resp azfake.Responder[armrecoveryservicesdatareplication.RecoveryPointsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method RecoveryPointsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, vaultName string, protectedItemName string, options *armrecoveryservicesdatareplication.RecoveryPointsClientListOptions) (resp azfake.PagerResponder[armrecoveryservicesdatareplication.RecoveryPointsClientListResponse])
}

// NewRecoveryPointsServerTransport creates a new instance of RecoveryPointsServerTransport with the provided implementation.
// The returned RecoveryPointsServerTransport instance is connected to an instance of armrecoveryservicesdatareplication.RecoveryPointsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRecoveryPointsServerTransport(srv *RecoveryPointsServer) *RecoveryPointsServerTransport {
	return &RecoveryPointsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armrecoveryservicesdatareplication.RecoveryPointsClientListResponse]](),
	}
}

// RecoveryPointsServerTransport connects instances of armrecoveryservicesdatareplication.RecoveryPointsClient to instances of RecoveryPointsServer.
// Don't use this type directly, use NewRecoveryPointsServerTransport instead.
type RecoveryPointsServerTransport struct {
	srv          *RecoveryPointsServer
	newListPager *tracker[azfake.PagerResponder[armrecoveryservicesdatareplication.RecoveryPointsClientListResponse]]
}

// Do implements the policy.Transporter interface for RecoveryPointsServerTransport.
func (r *RecoveryPointsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "RecoveryPointsClient.Get":
		resp, err = r.dispatchGet(req)
	case "RecoveryPointsClient.NewListPager":
		resp, err = r.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RecoveryPointsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataReplication/replicationVaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/protectedItems/(?P<protectedItemName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/recoveryPoints/(?P<recoveryPointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
	if err != nil {
		return nil, err
	}
	protectedItemNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("protectedItemName")])
	if err != nil {
		return nil, err
	}
	recoveryPointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("recoveryPointName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceGroupNameParam, vaultNameParam, protectedItemNameParam, recoveryPointNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RecoveryPointModel, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RecoveryPointsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := r.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataReplication/replicationVaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/protectedItems/(?P<protectedItemName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/recoveryPoints`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
		if err != nil {
			return nil, err
		}
		protectedItemNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("protectedItemName")])
		if err != nil {
			return nil, err
		}
		resp := r.srv.NewListPager(resourceGroupNameParam, vaultNameParam, protectedItemNameParam, nil)
		newListPager = &resp
		r.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armrecoveryservicesdatareplication.RecoveryPointsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		r.newListPager.remove(req)
	}
	return resp, nil
}
