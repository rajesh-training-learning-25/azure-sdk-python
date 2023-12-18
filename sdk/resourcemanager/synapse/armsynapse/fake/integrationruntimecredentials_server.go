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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
	"net/http"
	"net/url"
	"regexp"
)

// IntegrationRuntimeCredentialsServer is a fake server for instances of the armsynapse.IntegrationRuntimeCredentialsClient type.
type IntegrationRuntimeCredentialsServer struct {
	// Sync is the fake for method IntegrationRuntimeCredentialsClient.Sync
	// HTTP status codes to indicate success: http.StatusOK
	Sync func(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, options *armsynapse.IntegrationRuntimeCredentialsClientSyncOptions) (resp azfake.Responder[armsynapse.IntegrationRuntimeCredentialsClientSyncResponse], errResp azfake.ErrorResponder)
}

// NewIntegrationRuntimeCredentialsServerTransport creates a new instance of IntegrationRuntimeCredentialsServerTransport with the provided implementation.
// The returned IntegrationRuntimeCredentialsServerTransport instance is connected to an instance of armsynapse.IntegrationRuntimeCredentialsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewIntegrationRuntimeCredentialsServerTransport(srv *IntegrationRuntimeCredentialsServer) *IntegrationRuntimeCredentialsServerTransport {
	return &IntegrationRuntimeCredentialsServerTransport{srv: srv}
}

// IntegrationRuntimeCredentialsServerTransport connects instances of armsynapse.IntegrationRuntimeCredentialsClient to instances of IntegrationRuntimeCredentialsServer.
// Don't use this type directly, use NewIntegrationRuntimeCredentialsServerTransport instead.
type IntegrationRuntimeCredentialsServerTransport struct {
	srv *IntegrationRuntimeCredentialsServer
}

// Do implements the policy.Transporter interface for IntegrationRuntimeCredentialsServerTransport.
func (i *IntegrationRuntimeCredentialsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "IntegrationRuntimeCredentialsClient.Sync":
		resp, err = i.dispatchSync(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *IntegrationRuntimeCredentialsServerTransport) dispatchSync(req *http.Request) (*http.Response, error) {
	if i.srv.Sync == nil {
		return nil, &nonRetriableError{errors.New("fake for method Sync not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Synapse/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/integrationRuntimes/(?P<integrationRuntimeName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/syncCredentials`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
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
	integrationRuntimeNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("integrationRuntimeName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Sync(req.Context(), resourceGroupNameParam, workspaceNameParam, integrationRuntimeNameParam, nil)
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
