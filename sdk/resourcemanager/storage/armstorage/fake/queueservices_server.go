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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"net/http"
	"regexp"
)

// QueueServicesServer is a fake server for instances of the armstorage.QueueServicesClient type.
type QueueServicesServer struct {
	// GetServiceProperties is the fake for method QueueServicesClient.GetServiceProperties
	// HTTP status codes to indicate success: http.StatusOK
	GetServiceProperties func(ctx context.Context, resourceGroupName string, accountName string, options *armstorage.QueueServicesClientGetServicePropertiesOptions) (resp azfake.Responder[armstorage.QueueServicesClientGetServicePropertiesResponse], errResp azfake.ErrorResponder)

	// List is the fake for method QueueServicesClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, resourceGroupName string, accountName string, options *armstorage.QueueServicesClientListOptions) (resp azfake.Responder[armstorage.QueueServicesClientListResponse], errResp azfake.ErrorResponder)

	// SetServiceProperties is the fake for method QueueServicesClient.SetServiceProperties
	// HTTP status codes to indicate success: http.StatusOK
	SetServiceProperties func(ctx context.Context, resourceGroupName string, accountName string, parameters armstorage.QueueServiceProperties, options *armstorage.QueueServicesClientSetServicePropertiesOptions) (resp azfake.Responder[armstorage.QueueServicesClientSetServicePropertiesResponse], errResp azfake.ErrorResponder)
}

// NewQueueServicesServerTransport creates a new instance of QueueServicesServerTransport with the provided implementation.
// The returned QueueServicesServerTransport instance is connected to an instance of armstorage.QueueServicesClient by way of the
// undefined.Transporter field.
func NewQueueServicesServerTransport(srv *QueueServicesServer) *QueueServicesServerTransport {
	return &QueueServicesServerTransport{srv: srv}
}

// QueueServicesServerTransport connects instances of armstorage.QueueServicesClient to instances of QueueServicesServer.
// Don't use this type directly, use NewQueueServicesServerTransport instead.
type QueueServicesServerTransport struct {
	srv *QueueServicesServer
}

// Do implements the policy.Transporter interface for QueueServicesServerTransport.
func (q *QueueServicesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "QueueServicesClient.GetServiceProperties":
		resp, err = q.dispatchGetServiceProperties(req)
	case "QueueServicesClient.List":
		resp, err = q.dispatchList(req)
	case "QueueServicesClient.SetServiceProperties":
		resp, err = q.dispatchSetServiceProperties(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (q *QueueServicesServerTransport) dispatchGetServiceProperties(req *http.Request) (*http.Response, error) {
	if q.srv.GetServiceProperties == nil {
		return nil, &nonRetriableError{errors.New("method GetServiceProperties not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[a-zA-Z0-9-_]+)/queueServices/(?P<queueServiceName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := q.srv.GetServiceProperties(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("accountName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).QueueServiceProperties, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (q *QueueServicesServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if q.srv.List == nil {
		return nil, &nonRetriableError{errors.New("method List not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[a-zA-Z0-9-_]+)/queueServices"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := q.srv.List(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("accountName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ListQueueServices, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (q *QueueServicesServerTransport) dispatchSetServiceProperties(req *http.Request) (*http.Response, error) {
	if q.srv.SetServiceProperties == nil {
		return nil, &nonRetriableError{errors.New("method SetServiceProperties not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[a-zA-Z0-9-_]+)/queueServices/(?P<queueServiceName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armstorage.QueueServiceProperties](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := q.srv.SetServiceProperties(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("accountName")], body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).QueueServiceProperties, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
