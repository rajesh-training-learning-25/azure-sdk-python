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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v3"
	"net/http"
	"net/url"
	"regexp"
)

// AsyncOperationStatusServer is a fake server for instances of the armredis.AsyncOperationStatusClient type.
type AsyncOperationStatusServer struct {
	// Get is the fake for method AsyncOperationStatusClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, location string, operationID string, options *armredis.AsyncOperationStatusClientGetOptions) (resp azfake.Responder[armredis.AsyncOperationStatusClientGetResponse], errResp azfake.ErrorResponder)
}

// NewAsyncOperationStatusServerTransport creates a new instance of AsyncOperationStatusServerTransport with the provided implementation.
// The returned AsyncOperationStatusServerTransport instance is connected to an instance of armredis.AsyncOperationStatusClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAsyncOperationStatusServerTransport(srv *AsyncOperationStatusServer) *AsyncOperationStatusServerTransport {
	return &AsyncOperationStatusServerTransport{srv: srv}
}

// AsyncOperationStatusServerTransport connects instances of armredis.AsyncOperationStatusClient to instances of AsyncOperationStatusServer.
// Don't use this type directly, use NewAsyncOperationStatusServerTransport instead.
type AsyncOperationStatusServerTransport struct {
	srv *AsyncOperationStatusServer
}

// Do implements the policy.Transporter interface for AsyncOperationStatusServerTransport.
func (a *AsyncOperationStatusServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AsyncOperationStatusClient.Get":
		resp, err = a.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AsyncOperationStatusServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cache/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/asyncOperations/(?P<operationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	operationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), locationParam, operationIDParam, nil)
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
