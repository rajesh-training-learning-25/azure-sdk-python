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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
	"net/http"
	"net/url"
	"regexp"
)

// SourceControlSyncJobStreamsServer is a fake server for instances of the armautomation.SourceControlSyncJobStreamsClient type.
type SourceControlSyncJobStreamsServer struct {
	// Get is the fake for method SourceControlSyncJobStreamsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, automationAccountName string, sourceControlName string, sourceControlSyncJobID string, streamID string, options *armautomation.SourceControlSyncJobStreamsClientGetOptions) (resp azfake.Responder[armautomation.SourceControlSyncJobStreamsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListBySyncJobPager is the fake for method SourceControlSyncJobStreamsClient.NewListBySyncJobPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySyncJobPager func(resourceGroupName string, automationAccountName string, sourceControlName string, sourceControlSyncJobID string, options *armautomation.SourceControlSyncJobStreamsClientListBySyncJobOptions) (resp azfake.PagerResponder[armautomation.SourceControlSyncJobStreamsClientListBySyncJobResponse])
}

// NewSourceControlSyncJobStreamsServerTransport creates a new instance of SourceControlSyncJobStreamsServerTransport with the provided implementation.
// The returned SourceControlSyncJobStreamsServerTransport instance is connected to an instance of armautomation.SourceControlSyncJobStreamsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSourceControlSyncJobStreamsServerTransport(srv *SourceControlSyncJobStreamsServer) *SourceControlSyncJobStreamsServerTransport {
	return &SourceControlSyncJobStreamsServerTransport{
		srv:                   srv,
		newListBySyncJobPager: newTracker[azfake.PagerResponder[armautomation.SourceControlSyncJobStreamsClientListBySyncJobResponse]](),
	}
}

// SourceControlSyncJobStreamsServerTransport connects instances of armautomation.SourceControlSyncJobStreamsClient to instances of SourceControlSyncJobStreamsServer.
// Don't use this type directly, use NewSourceControlSyncJobStreamsServerTransport instead.
type SourceControlSyncJobStreamsServerTransport struct {
	srv                   *SourceControlSyncJobStreamsServer
	newListBySyncJobPager *tracker[azfake.PagerResponder[armautomation.SourceControlSyncJobStreamsClientListBySyncJobResponse]]
}

// Do implements the policy.Transporter interface for SourceControlSyncJobStreamsServerTransport.
func (s *SourceControlSyncJobStreamsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SourceControlSyncJobStreamsClient.Get":
		resp, err = s.dispatchGet(req)
	case "SourceControlSyncJobStreamsClient.NewListBySyncJobPager":
		resp, err = s.dispatchNewListBySyncJobPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SourceControlSyncJobStreamsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automation/automationAccounts/(?P<automationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sourceControls/(?P<sourceControlName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sourceControlSyncJobs/(?P<sourceControlSyncJobId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/streams/(?P<streamId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	automationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("automationAccountName")])
	if err != nil {
		return nil, err
	}
	sourceControlNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sourceControlName")])
	if err != nil {
		return nil, err
	}
	sourceControlSyncJobIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("sourceControlSyncJobId")])
	if err != nil {
		return nil, err
	}
	streamIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("streamId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, automationAccountNameParam, sourceControlNameParam, sourceControlSyncJobIDParam, streamIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SourceControlSyncJobStreamByID, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SourceControlSyncJobStreamsServerTransport) dispatchNewListBySyncJobPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListBySyncJobPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySyncJobPager not implemented")}
	}
	newListBySyncJobPager := s.newListBySyncJobPager.get(req)
	if newListBySyncJobPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automation/automationAccounts/(?P<automationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sourceControls/(?P<sourceControlName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sourceControlSyncJobs/(?P<sourceControlSyncJobId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/streams`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		automationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("automationAccountName")])
		if err != nil {
			return nil, err
		}
		sourceControlNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sourceControlName")])
		if err != nil {
			return nil, err
		}
		sourceControlSyncJobIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("sourceControlSyncJobId")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armautomation.SourceControlSyncJobStreamsClientListBySyncJobOptions
		if filterParam != nil {
			options = &armautomation.SourceControlSyncJobStreamsClientListBySyncJobOptions{
				Filter: filterParam,
			}
		}
		resp := s.srv.NewListBySyncJobPager(resourceGroupNameParam, automationAccountNameParam, sourceControlNameParam, sourceControlSyncJobIDParam, options)
		newListBySyncJobPager = &resp
		s.newListBySyncJobPager.add(req, newListBySyncJobPager)
		server.PagerResponderInjectNextLinks(newListBySyncJobPager, req, func(page *armautomation.SourceControlSyncJobStreamsClientListBySyncJobResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySyncJobPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListBySyncJobPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySyncJobPager) {
		s.newListBySyncJobPager.remove(req)
	}
	return resp, nil
}
