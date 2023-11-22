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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearningservices/armmachinelearningservices"
	"net/http"
	"net/url"
	"regexp"
)

// PrivateLinkResourcesServer is a fake server for instances of the armmachinelearningservices.PrivateLinkResourcesClient type.
type PrivateLinkResourcesServer struct {
	// List is the fake for method PrivateLinkResourcesClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, resourceGroupName string, workspaceName string, options *armmachinelearningservices.PrivateLinkResourcesClientListOptions) (resp azfake.Responder[armmachinelearningservices.PrivateLinkResourcesClientListResponse], errResp azfake.ErrorResponder)
}

// NewPrivateLinkResourcesServerTransport creates a new instance of PrivateLinkResourcesServerTransport with the provided implementation.
// The returned PrivateLinkResourcesServerTransport instance is connected to an instance of armmachinelearningservices.PrivateLinkResourcesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPrivateLinkResourcesServerTransport(srv *PrivateLinkResourcesServer) *PrivateLinkResourcesServerTransport {
	return &PrivateLinkResourcesServerTransport{srv: srv}
}

// PrivateLinkResourcesServerTransport connects instances of armmachinelearningservices.PrivateLinkResourcesClient to instances of PrivateLinkResourcesServer.
// Don't use this type directly, use NewPrivateLinkResourcesServerTransport instead.
type PrivateLinkResourcesServerTransport struct {
	srv *PrivateLinkResourcesServer
}

// Do implements the policy.Transporter interface for PrivateLinkResourcesServerTransport.
func (p *PrivateLinkResourcesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PrivateLinkResourcesClient.List":
		resp, err = p.dispatchList(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PrivateLinkResourcesServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if p.srv.List == nil {
		return nil, &nonRetriableError{errors.New("fake for method List not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateLinkResources`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.List(req.Context(), resourceGroupNameParam, workspaceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateLinkResourceListResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
