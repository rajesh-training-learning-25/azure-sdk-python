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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservicesdatareplication/armrecoveryservicesdatareplication"
	"net/http"
	"net/url"
	"regexp"
)

// ReplicationExtensionOperationStatusServer is a fake server for instances of the armrecoveryservicesdatareplication.ReplicationExtensionOperationStatusClient type.
type ReplicationExtensionOperationStatusServer struct {
	// Get is the fake for method ReplicationExtensionOperationStatusClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, vaultName string, replicationExtensionName string, operationID string, options *armrecoveryservicesdatareplication.ReplicationExtensionOperationStatusClientGetOptions) (resp azfake.Responder[armrecoveryservicesdatareplication.ReplicationExtensionOperationStatusClientGetResponse], errResp azfake.ErrorResponder)
}

// NewReplicationExtensionOperationStatusServerTransport creates a new instance of ReplicationExtensionOperationStatusServerTransport with the provided implementation.
// The returned ReplicationExtensionOperationStatusServerTransport instance is connected to an instance of armrecoveryservicesdatareplication.ReplicationExtensionOperationStatusClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewReplicationExtensionOperationStatusServerTransport(srv *ReplicationExtensionOperationStatusServer) *ReplicationExtensionOperationStatusServerTransport {
	return &ReplicationExtensionOperationStatusServerTransport{srv: srv}
}

// ReplicationExtensionOperationStatusServerTransport connects instances of armrecoveryservicesdatareplication.ReplicationExtensionOperationStatusClient to instances of ReplicationExtensionOperationStatusServer.
// Don't use this type directly, use NewReplicationExtensionOperationStatusServerTransport instead.
type ReplicationExtensionOperationStatusServerTransport struct {
	srv *ReplicationExtensionOperationStatusServer
}

// Do implements the policy.Transporter interface for ReplicationExtensionOperationStatusServerTransport.
func (r *ReplicationExtensionOperationStatusServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ReplicationExtensionOperationStatusClient.Get":
		resp, err = r.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *ReplicationExtensionOperationStatusServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataReplication/replicationVaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationExtensions/(?P<replicationExtensionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/operations/(?P<operationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	replicationExtensionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("replicationExtensionName")])
	if err != nil {
		return nil, err
	}
	operationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceGroupNameParam, vaultNameParam, replicationExtensionNameParam, operationIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).OperationStatus, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
