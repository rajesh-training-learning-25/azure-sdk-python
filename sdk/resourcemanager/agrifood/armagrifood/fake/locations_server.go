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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agrifood/armagrifood"
	"net/http"
	"regexp"
)

// LocationsServer is a fake server for instances of the armagrifood.LocationsClient type.
type LocationsServer struct {
	// CheckNameAvailability is the fake for method LocationsClient.CheckNameAvailability
	// HTTP status codes to indicate success: http.StatusOK
	CheckNameAvailability func(ctx context.Context, body armagrifood.CheckNameAvailabilityRequest, options *armagrifood.LocationsClientCheckNameAvailabilityOptions) (resp azfake.Responder[armagrifood.LocationsClientCheckNameAvailabilityResponse], errResp azfake.ErrorResponder)
}

// NewLocationsServerTransport creates a new instance of LocationsServerTransport with the provided implementation.
// The returned LocationsServerTransport instance is connected to an instance of armagrifood.LocationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewLocationsServerTransport(srv *LocationsServer) *LocationsServerTransport {
	return &LocationsServerTransport{srv: srv}
}

// LocationsServerTransport connects instances of armagrifood.LocationsClient to instances of LocationsServer.
// Don't use this type directly, use NewLocationsServerTransport instead.
type LocationsServerTransport struct {
	srv *LocationsServer
}

// Do implements the policy.Transporter interface for LocationsServerTransport.
func (l *LocationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "LocationsClient.CheckNameAvailability":
		resp, err = l.dispatchCheckNameAvailability(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (l *LocationsServerTransport) dispatchCheckNameAvailability(req *http.Request) (*http.Response, error) {
	if l.srv.CheckNameAvailability == nil {
		return nil, &nonRetriableError{errors.New("fake for method CheckNameAvailability not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AgFoodPlatform/checkNameAvailability`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armagrifood.CheckNameAvailabilityRequest](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := l.srv.CheckNameAvailability(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CheckNameAvailabilityResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
