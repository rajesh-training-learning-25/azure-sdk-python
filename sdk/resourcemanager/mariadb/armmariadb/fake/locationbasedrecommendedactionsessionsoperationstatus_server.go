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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb"
	"net/http"
	"net/url"
	"regexp"
)

// LocationBasedRecommendedActionSessionsOperationStatusServer is a fake server for instances of the armmariadb.LocationBasedRecommendedActionSessionsOperationStatusClient type.
type LocationBasedRecommendedActionSessionsOperationStatusServer struct {
	// Get is the fake for method LocationBasedRecommendedActionSessionsOperationStatusClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, locationName string, operationID string, options *armmariadb.LocationBasedRecommendedActionSessionsOperationStatusClientGetOptions) (resp azfake.Responder[armmariadb.LocationBasedRecommendedActionSessionsOperationStatusClientGetResponse], errResp azfake.ErrorResponder)
}

// NewLocationBasedRecommendedActionSessionsOperationStatusServerTransport creates a new instance of LocationBasedRecommendedActionSessionsOperationStatusServerTransport with the provided implementation.
// The returned LocationBasedRecommendedActionSessionsOperationStatusServerTransport instance is connected to an instance of armmariadb.LocationBasedRecommendedActionSessionsOperationStatusClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewLocationBasedRecommendedActionSessionsOperationStatusServerTransport(srv *LocationBasedRecommendedActionSessionsOperationStatusServer) *LocationBasedRecommendedActionSessionsOperationStatusServerTransport {
	return &LocationBasedRecommendedActionSessionsOperationStatusServerTransport{srv: srv}
}

// LocationBasedRecommendedActionSessionsOperationStatusServerTransport connects instances of armmariadb.LocationBasedRecommendedActionSessionsOperationStatusClient to instances of LocationBasedRecommendedActionSessionsOperationStatusServer.
// Don't use this type directly, use NewLocationBasedRecommendedActionSessionsOperationStatusServerTransport instead.
type LocationBasedRecommendedActionSessionsOperationStatusServerTransport struct {
	srv *LocationBasedRecommendedActionSessionsOperationStatusServer
}

// Do implements the policy.Transporter interface for LocationBasedRecommendedActionSessionsOperationStatusServerTransport.
func (l *LocationBasedRecommendedActionSessionsOperationStatusServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "LocationBasedRecommendedActionSessionsOperationStatusClient.Get":
		resp, err = l.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (l *LocationBasedRecommendedActionSessionsOperationStatusServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if l.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforMariaDB/locations/(?P<locationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/recommendedActionSessionsAzureAsyncOperation/(?P<operationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("locationName")])
	if err != nil {
		return nil, err
	}
	operationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := l.srv.Get(req.Context(), locationNameParam, operationIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RecommendedActionSessionsOperationStatus, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
