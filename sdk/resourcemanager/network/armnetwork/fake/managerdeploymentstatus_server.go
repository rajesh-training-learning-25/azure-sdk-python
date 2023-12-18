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
	"strconv"
)

// ManagerDeploymentStatusServer is a fake server for instances of the armnetwork.ManagerDeploymentStatusClient type.
type ManagerDeploymentStatusServer struct {
	// List is the fake for method ManagerDeploymentStatusClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, resourceGroupName string, networkManagerName string, parameters armnetwork.ManagerDeploymentStatusParameter, options *armnetwork.ManagerDeploymentStatusClientListOptions) (resp azfake.Responder[armnetwork.ManagerDeploymentStatusClientListResponse], errResp azfake.ErrorResponder)
}

// NewManagerDeploymentStatusServerTransport creates a new instance of ManagerDeploymentStatusServerTransport with the provided implementation.
// The returned ManagerDeploymentStatusServerTransport instance is connected to an instance of armnetwork.ManagerDeploymentStatusClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewManagerDeploymentStatusServerTransport(srv *ManagerDeploymentStatusServer) *ManagerDeploymentStatusServerTransport {
	return &ManagerDeploymentStatusServerTransport{srv: srv}
}

// ManagerDeploymentStatusServerTransport connects instances of armnetwork.ManagerDeploymentStatusClient to instances of ManagerDeploymentStatusServer.
// Don't use this type directly, use NewManagerDeploymentStatusServerTransport instead.
type ManagerDeploymentStatusServerTransport struct {
	srv *ManagerDeploymentStatusServer
}

// Do implements the policy.Transporter interface for ManagerDeploymentStatusServerTransport.
func (m *ManagerDeploymentStatusServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ManagerDeploymentStatusClient.List":
		resp, err = m.dispatchList(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *ManagerDeploymentStatusServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if m.srv.List == nil {
		return nil, &nonRetriableError{errors.New("fake for method List not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkManagers/(?P<networkManagerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listDeploymentStatus`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	body, err := server.UnmarshalRequestAsJSON[armnetwork.ManagerDeploymentStatusParameter](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkManagerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkManagerName")])
	if err != nil {
		return nil, err
	}
	topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
	if err != nil {
		return nil, err
	}
	topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	var options *armnetwork.ManagerDeploymentStatusClientListOptions
	if topParam != nil {
		options = &armnetwork.ManagerDeploymentStatusClientListOptions{
			Top: topParam,
		}
	}
	respr, errRespr := m.srv.List(req.Context(), resourceGroupNameParam, networkManagerNameParam, body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagerDeploymentStatusListResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
