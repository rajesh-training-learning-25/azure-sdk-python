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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"net/http"
	"net/url"
	"regexp"
)

// ManagedInstanceAzureADOnlyAuthenticationsServer is a fake server for instances of the armsql.ManagedInstanceAzureADOnlyAuthenticationsClient type.
type ManagedInstanceAzureADOnlyAuthenticationsServer struct {
	// BeginCreateOrUpdate is the fake for method ManagedInstanceAzureADOnlyAuthenticationsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, managedInstanceName string, authenticationName armsql.AuthenticationName, parameters armsql.ManagedInstanceAzureADOnlyAuthentication, options *armsql.ManagedInstanceAzureADOnlyAuthenticationsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armsql.ManagedInstanceAzureADOnlyAuthenticationsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ManagedInstanceAzureADOnlyAuthenticationsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, managedInstanceName string, authenticationName armsql.AuthenticationName, options *armsql.ManagedInstanceAzureADOnlyAuthenticationsClientBeginDeleteOptions) (resp azfake.PollerResponder[armsql.ManagedInstanceAzureADOnlyAuthenticationsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ManagedInstanceAzureADOnlyAuthenticationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, managedInstanceName string, authenticationName armsql.AuthenticationName, options *armsql.ManagedInstanceAzureADOnlyAuthenticationsClientGetOptions) (resp azfake.Responder[armsql.ManagedInstanceAzureADOnlyAuthenticationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByInstancePager is the fake for method ManagedInstanceAzureADOnlyAuthenticationsClient.NewListByInstancePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByInstancePager func(resourceGroupName string, managedInstanceName string, options *armsql.ManagedInstanceAzureADOnlyAuthenticationsClientListByInstanceOptions) (resp azfake.PagerResponder[armsql.ManagedInstanceAzureADOnlyAuthenticationsClientListByInstanceResponse])
}

// NewManagedInstanceAzureADOnlyAuthenticationsServerTransport creates a new instance of ManagedInstanceAzureADOnlyAuthenticationsServerTransport with the provided implementation.
// The returned ManagedInstanceAzureADOnlyAuthenticationsServerTransport instance is connected to an instance of armsql.ManagedInstanceAzureADOnlyAuthenticationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewManagedInstanceAzureADOnlyAuthenticationsServerTransport(srv *ManagedInstanceAzureADOnlyAuthenticationsServer) *ManagedInstanceAzureADOnlyAuthenticationsServerTransport {
	return &ManagedInstanceAzureADOnlyAuthenticationsServerTransport{
		srv:                    srv,
		beginCreateOrUpdate:    newTracker[azfake.PollerResponder[armsql.ManagedInstanceAzureADOnlyAuthenticationsClientCreateOrUpdateResponse]](),
		beginDelete:            newTracker[azfake.PollerResponder[armsql.ManagedInstanceAzureADOnlyAuthenticationsClientDeleteResponse]](),
		newListByInstancePager: newTracker[azfake.PagerResponder[armsql.ManagedInstanceAzureADOnlyAuthenticationsClientListByInstanceResponse]](),
	}
}

// ManagedInstanceAzureADOnlyAuthenticationsServerTransport connects instances of armsql.ManagedInstanceAzureADOnlyAuthenticationsClient to instances of ManagedInstanceAzureADOnlyAuthenticationsServer.
// Don't use this type directly, use NewManagedInstanceAzureADOnlyAuthenticationsServerTransport instead.
type ManagedInstanceAzureADOnlyAuthenticationsServerTransport struct {
	srv                    *ManagedInstanceAzureADOnlyAuthenticationsServer
	beginCreateOrUpdate    *tracker[azfake.PollerResponder[armsql.ManagedInstanceAzureADOnlyAuthenticationsClientCreateOrUpdateResponse]]
	beginDelete            *tracker[azfake.PollerResponder[armsql.ManagedInstanceAzureADOnlyAuthenticationsClientDeleteResponse]]
	newListByInstancePager *tracker[azfake.PagerResponder[armsql.ManagedInstanceAzureADOnlyAuthenticationsClientListByInstanceResponse]]
}

// Do implements the policy.Transporter interface for ManagedInstanceAzureADOnlyAuthenticationsServerTransport.
func (m *ManagedInstanceAzureADOnlyAuthenticationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ManagedInstanceAzureADOnlyAuthenticationsClient.BeginCreateOrUpdate":
		resp, err = m.dispatchBeginCreateOrUpdate(req)
	case "ManagedInstanceAzureADOnlyAuthenticationsClient.BeginDelete":
		resp, err = m.dispatchBeginDelete(req)
	case "ManagedInstanceAzureADOnlyAuthenticationsClient.Get":
		resp, err = m.dispatchGet(req)
	case "ManagedInstanceAzureADOnlyAuthenticationsClient.NewListByInstancePager":
		resp, err = m.dispatchNewListByInstancePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *ManagedInstanceAzureADOnlyAuthenticationsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := m.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/azureADOnlyAuthentications/(?P<authenticationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsql.ManagedInstanceAzureADOnlyAuthentication](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
		if err != nil {
			return nil, err
		}
		authenticationNameParam, err := parseWithCast(matches[regex.SubexpIndex("authenticationName")], func(v string) (armsql.AuthenticationName, error) {
			p, unescapeErr := url.PathUnescape(v)
			if unescapeErr != nil {
				return "", unescapeErr
			}
			return armsql.AuthenticationName(p), nil
		})
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, managedInstanceNameParam, authenticationNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		m.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		m.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		m.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (m *ManagedInstanceAzureADOnlyAuthenticationsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if m.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := m.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/azureADOnlyAuthentications/(?P<authenticationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
		if err != nil {
			return nil, err
		}
		authenticationNameParam, err := parseWithCast(matches[regex.SubexpIndex("authenticationName")], func(v string) (armsql.AuthenticationName, error) {
			p, unescapeErr := url.PathUnescape(v)
			if unescapeErr != nil {
				return "", unescapeErr
			}
			return armsql.AuthenticationName(p), nil
		})
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginDelete(req.Context(), resourceGroupNameParam, managedInstanceNameParam, authenticationNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		m.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		m.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		m.beginDelete.remove(req)
	}

	return resp, nil
}

func (m *ManagedInstanceAzureADOnlyAuthenticationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if m.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/azureADOnlyAuthentications/(?P<authenticationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
	if err != nil {
		return nil, err
	}
	authenticationNameParam, err := parseWithCast(matches[regex.SubexpIndex("authenticationName")], func(v string) (armsql.AuthenticationName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armsql.AuthenticationName(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Get(req.Context(), resourceGroupNameParam, managedInstanceNameParam, authenticationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagedInstanceAzureADOnlyAuthentication, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedInstanceAzureADOnlyAuthenticationsServerTransport) dispatchNewListByInstancePager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListByInstancePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByInstancePager not implemented")}
	}
	newListByInstancePager := m.newListByInstancePager.get(req)
	if newListByInstancePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/azureADOnlyAuthentications`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
		if err != nil {
			return nil, err
		}
		resp := m.srv.NewListByInstancePager(resourceGroupNameParam, managedInstanceNameParam, nil)
		newListByInstancePager = &resp
		m.newListByInstancePager.add(req, newListByInstancePager)
		server.PagerResponderInjectNextLinks(newListByInstancePager, req, func(page *armsql.ManagedInstanceAzureADOnlyAuthenticationsClientListByInstanceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByInstancePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListByInstancePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByInstancePager) {
		m.newListByInstancePager.remove(req)
	}
	return resp, nil
}
