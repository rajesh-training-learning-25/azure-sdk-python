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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
	"net/http"
)

// ProviderOperationsServer is a fake server for instances of the armdevtestlabs.ProviderOperationsClient type.
type ProviderOperationsServer struct {
	// NewListPager is the fake for method ProviderOperationsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armdevtestlabs.ProviderOperationsClientListOptions) (resp azfake.PagerResponder[armdevtestlabs.ProviderOperationsClientListResponse])
}

// NewProviderOperationsServerTransport creates a new instance of ProviderOperationsServerTransport with the provided implementation.
// The returned ProviderOperationsServerTransport instance is connected to an instance of armdevtestlabs.ProviderOperationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewProviderOperationsServerTransport(srv *ProviderOperationsServer) *ProviderOperationsServerTransport {
	return &ProviderOperationsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armdevtestlabs.ProviderOperationsClientListResponse]](),
	}
}

// ProviderOperationsServerTransport connects instances of armdevtestlabs.ProviderOperationsClient to instances of ProviderOperationsServer.
// Don't use this type directly, use NewProviderOperationsServerTransport instead.
type ProviderOperationsServerTransport struct {
	srv          *ProviderOperationsServer
	newListPager *tracker[azfake.PagerResponder[armdevtestlabs.ProviderOperationsClientListResponse]]
}

// Do implements the policy.Transporter interface for ProviderOperationsServerTransport.
func (p *ProviderOperationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ProviderOperationsClient.NewListPager":
		resp, err = p.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *ProviderOperationsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := p.newListPager.get(req)
	if newListPager == nil {
		resp := p.srv.NewListPager(nil)
		newListPager = &resp
		p.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armdevtestlabs.ProviderOperationsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		p.newListPager.remove(req)
	}
	return resp, nil
}
