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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/labservices/armlabservices"
	"net/http"
	"net/url"
	"regexp"
)

// OperationResultsServer is a fake server for instances of the armlabservices.OperationResultsClient type.
type OperationResultsServer struct {
	// Get is the fake for method OperationResultsClient.Get
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Get func(ctx context.Context, operationResultID string, options *armlabservices.OperationResultsClientGetOptions) (resp azfake.Responder[armlabservices.OperationResultsClientGetResponse], errResp azfake.ErrorResponder)
}

// NewOperationResultsServerTransport creates a new instance of OperationResultsServerTransport with the provided implementation.
// The returned OperationResultsServerTransport instance is connected to an instance of armlabservices.OperationResultsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewOperationResultsServerTransport(srv *OperationResultsServer) *OperationResultsServerTransport {
	return &OperationResultsServerTransport{srv: srv}
}

// OperationResultsServerTransport connects instances of armlabservices.OperationResultsClient to instances of OperationResultsServer.
// Don't use this type directly, use NewOperationResultsServerTransport instead.
type OperationResultsServerTransport struct {
	srv *OperationResultsServer
}

// Do implements the policy.Transporter interface for OperationResultsServerTransport.
func (o *OperationResultsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "OperationResultsClient.Get":
		resp, err = o.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (o *OperationResultsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if o.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.LabServices/operationResults/(?P<operationResultId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	operationResultIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationResultId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.Get(req.Context(), operationResultIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).OperationResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
