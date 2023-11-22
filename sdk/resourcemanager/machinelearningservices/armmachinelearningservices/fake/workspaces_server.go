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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearningservices/armmachinelearningservices"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// WorkspacesServer is a fake server for instances of the armmachinelearningservices.WorkspacesClient type.
type WorkspacesServer struct {
	// BeginCreateOrUpdate is the fake for method WorkspacesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, workspaceName string, parameters armmachinelearningservices.Workspace, options *armmachinelearningservices.WorkspacesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armmachinelearningservices.WorkspacesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method WorkspacesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, workspaceName string, options *armmachinelearningservices.WorkspacesClientBeginDeleteOptions) (resp azfake.PollerResponder[armmachinelearningservices.WorkspacesClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginDiagnose is the fake for method WorkspacesClient.BeginDiagnose
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginDiagnose func(ctx context.Context, resourceGroupName string, workspaceName string, options *armmachinelearningservices.WorkspacesClientBeginDiagnoseOptions) (resp azfake.PollerResponder[armmachinelearningservices.WorkspacesClientDiagnoseResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method WorkspacesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, workspaceName string, options *armmachinelearningservices.WorkspacesClientGetOptions) (resp azfake.Responder[armmachinelearningservices.WorkspacesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method WorkspacesClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armmachinelearningservices.WorkspacesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armmachinelearningservices.WorkspacesClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method WorkspacesClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armmachinelearningservices.WorkspacesClientListBySubscriptionOptions) (resp azfake.PagerResponder[armmachinelearningservices.WorkspacesClientListBySubscriptionResponse])

	// ListKeys is the fake for method WorkspacesClient.ListKeys
	// HTTP status codes to indicate success: http.StatusOK
	ListKeys func(ctx context.Context, resourceGroupName string, workspaceName string, options *armmachinelearningservices.WorkspacesClientListKeysOptions) (resp azfake.Responder[armmachinelearningservices.WorkspacesClientListKeysResponse], errResp azfake.ErrorResponder)

	// ListNotebookAccessToken is the fake for method WorkspacesClient.ListNotebookAccessToken
	// HTTP status codes to indicate success: http.StatusOK
	ListNotebookAccessToken func(ctx context.Context, resourceGroupName string, workspaceName string, options *armmachinelearningservices.WorkspacesClientListNotebookAccessTokenOptions) (resp azfake.Responder[armmachinelearningservices.WorkspacesClientListNotebookAccessTokenResponse], errResp azfake.ErrorResponder)

	// ListNotebookKeys is the fake for method WorkspacesClient.ListNotebookKeys
	// HTTP status codes to indicate success: http.StatusOK
	ListNotebookKeys func(ctx context.Context, resourceGroupName string, workspaceName string, options *armmachinelearningservices.WorkspacesClientListNotebookKeysOptions) (resp azfake.Responder[armmachinelearningservices.WorkspacesClientListNotebookKeysResponse], errResp azfake.ErrorResponder)

	// ListOutboundNetworkDependenciesEndpoints is the fake for method WorkspacesClient.ListOutboundNetworkDependenciesEndpoints
	// HTTP status codes to indicate success: http.StatusOK
	ListOutboundNetworkDependenciesEndpoints func(ctx context.Context, resourceGroupName string, workspaceName string, options *armmachinelearningservices.WorkspacesClientListOutboundNetworkDependenciesEndpointsOptions) (resp azfake.Responder[armmachinelearningservices.WorkspacesClientListOutboundNetworkDependenciesEndpointsResponse], errResp azfake.ErrorResponder)

	// ListStorageAccountKeys is the fake for method WorkspacesClient.ListStorageAccountKeys
	// HTTP status codes to indicate success: http.StatusOK
	ListStorageAccountKeys func(ctx context.Context, resourceGroupName string, workspaceName string, options *armmachinelearningservices.WorkspacesClientListStorageAccountKeysOptions) (resp azfake.Responder[armmachinelearningservices.WorkspacesClientListStorageAccountKeysResponse], errResp azfake.ErrorResponder)

	// BeginPrepareNotebook is the fake for method WorkspacesClient.BeginPrepareNotebook
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginPrepareNotebook func(ctx context.Context, resourceGroupName string, workspaceName string, options *armmachinelearningservices.WorkspacesClientBeginPrepareNotebookOptions) (resp azfake.PollerResponder[armmachinelearningservices.WorkspacesClientPrepareNotebookResponse], errResp azfake.ErrorResponder)

	// BeginResyncKeys is the fake for method WorkspacesClient.BeginResyncKeys
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginResyncKeys func(ctx context.Context, resourceGroupName string, workspaceName string, options *armmachinelearningservices.WorkspacesClientBeginResyncKeysOptions) (resp azfake.PollerResponder[armmachinelearningservices.WorkspacesClientResyncKeysResponse], errResp azfake.ErrorResponder)

	// Update is the fake for method WorkspacesClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, workspaceName string, parameters armmachinelearningservices.WorkspaceUpdateParameters, options *armmachinelearningservices.WorkspacesClientUpdateOptions) (resp azfake.Responder[armmachinelearningservices.WorkspacesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewWorkspacesServerTransport creates a new instance of WorkspacesServerTransport with the provided implementation.
// The returned WorkspacesServerTransport instance is connected to an instance of armmachinelearningservices.WorkspacesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWorkspacesServerTransport(srv *WorkspacesServer) *WorkspacesServerTransport {
	return &WorkspacesServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armmachinelearningservices.WorkspacesClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armmachinelearningservices.WorkspacesClientDeleteResponse]](),
		beginDiagnose:               newTracker[azfake.PollerResponder[armmachinelearningservices.WorkspacesClientDiagnoseResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armmachinelearningservices.WorkspacesClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armmachinelearningservices.WorkspacesClientListBySubscriptionResponse]](),
		beginPrepareNotebook:        newTracker[azfake.PollerResponder[armmachinelearningservices.WorkspacesClientPrepareNotebookResponse]](),
		beginResyncKeys:             newTracker[azfake.PollerResponder[armmachinelearningservices.WorkspacesClientResyncKeysResponse]](),
	}
}

// WorkspacesServerTransport connects instances of armmachinelearningservices.WorkspacesClient to instances of WorkspacesServer.
// Don't use this type directly, use NewWorkspacesServerTransport instead.
type WorkspacesServerTransport struct {
	srv                         *WorkspacesServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armmachinelearningservices.WorkspacesClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armmachinelearningservices.WorkspacesClientDeleteResponse]]
	beginDiagnose               *tracker[azfake.PollerResponder[armmachinelearningservices.WorkspacesClientDiagnoseResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armmachinelearningservices.WorkspacesClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armmachinelearningservices.WorkspacesClientListBySubscriptionResponse]]
	beginPrepareNotebook        *tracker[azfake.PollerResponder[armmachinelearningservices.WorkspacesClientPrepareNotebookResponse]]
	beginResyncKeys             *tracker[azfake.PollerResponder[armmachinelearningservices.WorkspacesClientResyncKeysResponse]]
}

// Do implements the policy.Transporter interface for WorkspacesServerTransport.
func (w *WorkspacesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "WorkspacesClient.BeginCreateOrUpdate":
		resp, err = w.dispatchBeginCreateOrUpdate(req)
	case "WorkspacesClient.BeginDelete":
		resp, err = w.dispatchBeginDelete(req)
	case "WorkspacesClient.BeginDiagnose":
		resp, err = w.dispatchBeginDiagnose(req)
	case "WorkspacesClient.Get":
		resp, err = w.dispatchGet(req)
	case "WorkspacesClient.NewListByResourceGroupPager":
		resp, err = w.dispatchNewListByResourceGroupPager(req)
	case "WorkspacesClient.NewListBySubscriptionPager":
		resp, err = w.dispatchNewListBySubscriptionPager(req)
	case "WorkspacesClient.ListKeys":
		resp, err = w.dispatchListKeys(req)
	case "WorkspacesClient.ListNotebookAccessToken":
		resp, err = w.dispatchListNotebookAccessToken(req)
	case "WorkspacesClient.ListNotebookKeys":
		resp, err = w.dispatchListNotebookKeys(req)
	case "WorkspacesClient.ListOutboundNetworkDependenciesEndpoints":
		resp, err = w.dispatchListOutboundNetworkDependenciesEndpoints(req)
	case "WorkspacesClient.ListStorageAccountKeys":
		resp, err = w.dispatchListStorageAccountKeys(req)
	case "WorkspacesClient.BeginPrepareNotebook":
		resp, err = w.dispatchBeginPrepareNotebook(req)
	case "WorkspacesClient.BeginResyncKeys":
		resp, err = w.dispatchBeginResyncKeys(req)
	case "WorkspacesClient.Update":
		resp, err = w.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *WorkspacesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if w.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := w.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmachinelearningservices.Workspace](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := w.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, workspaceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		w.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		w.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		w.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (w *WorkspacesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if w.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := w.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
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
		respr, errRespr := w.srv.BeginDelete(req.Context(), resourceGroupNameParam, workspaceNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		w.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		w.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		w.beginDelete.remove(req)
	}

	return resp, nil
}

func (w *WorkspacesServerTransport) dispatchBeginDiagnose(req *http.Request) (*http.Response, error) {
	if w.srv.BeginDiagnose == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDiagnose not implemented")}
	}
	beginDiagnose := w.beginDiagnose.get(req)
	if beginDiagnose == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/diagnose`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmachinelearningservices.DiagnoseWorkspaceParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		var options *armmachinelearningservices.WorkspacesClientBeginDiagnoseOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armmachinelearningservices.WorkspacesClientBeginDiagnoseOptions{
				Parameters: &body,
			}
		}
		respr, errRespr := w.srv.BeginDiagnose(req.Context(), resourceGroupNameParam, workspaceNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDiagnose = &respr
		w.beginDiagnose.add(req, beginDiagnose)
	}

	resp, err := server.PollerResponderNext(beginDiagnose, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		w.beginDiagnose.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDiagnose) {
		w.beginDiagnose.remove(req)
	}

	return resp, nil
}

func (w *WorkspacesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if w.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
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
	respr, errRespr := w.srv.Get(req.Context(), resourceGroupNameParam, workspaceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Workspace, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkspacesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := w.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
		if err != nil {
			return nil, err
		}
		skipParam := getOptional(skipUnescaped)
		var options *armmachinelearningservices.WorkspacesClientListByResourceGroupOptions
		if skipParam != nil {
			options = &armmachinelearningservices.WorkspacesClientListByResourceGroupOptions{
				Skip: skipParam,
			}
		}
		resp := w.srv.NewListByResourceGroupPager(resourceGroupNameParam, options)
		newListByResourceGroupPager = &resp
		w.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armmachinelearningservices.WorkspacesClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		w.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		w.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (w *WorkspacesServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := w.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
		if err != nil {
			return nil, err
		}
		skipParam := getOptional(skipUnescaped)
		var options *armmachinelearningservices.WorkspacesClientListBySubscriptionOptions
		if skipParam != nil {
			options = &armmachinelearningservices.WorkspacesClientListBySubscriptionOptions{
				Skip: skipParam,
			}
		}
		resp := w.srv.NewListBySubscriptionPager(options)
		newListBySubscriptionPager = &resp
		w.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armmachinelearningservices.WorkspacesClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		w.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		w.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (w *WorkspacesServerTransport) dispatchListKeys(req *http.Request) (*http.Response, error) {
	if w.srv.ListKeys == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListKeys not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listKeys`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
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
	respr, errRespr := w.srv.ListKeys(req.Context(), resourceGroupNameParam, workspaceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ListWorkspaceKeysResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkspacesServerTransport) dispatchListNotebookAccessToken(req *http.Request) (*http.Response, error) {
	if w.srv.ListNotebookAccessToken == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListNotebookAccessToken not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listNotebookAccessToken`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
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
	respr, errRespr := w.srv.ListNotebookAccessToken(req.Context(), resourceGroupNameParam, workspaceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NotebookAccessTokenResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkspacesServerTransport) dispatchListNotebookKeys(req *http.Request) (*http.Response, error) {
	if w.srv.ListNotebookKeys == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListNotebookKeys not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listNotebookKeys`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
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
	respr, errRespr := w.srv.ListNotebookKeys(req.Context(), resourceGroupNameParam, workspaceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ListNotebookKeysResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkspacesServerTransport) dispatchListOutboundNetworkDependenciesEndpoints(req *http.Request) (*http.Response, error) {
	if w.srv.ListOutboundNetworkDependenciesEndpoints == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListOutboundNetworkDependenciesEndpoints not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/outboundNetworkDependenciesEndpoints`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
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
	respr, errRespr := w.srv.ListOutboundNetworkDependenciesEndpoints(req.Context(), resourceGroupNameParam, workspaceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExternalFQDNResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkspacesServerTransport) dispatchListStorageAccountKeys(req *http.Request) (*http.Response, error) {
	if w.srv.ListStorageAccountKeys == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListStorageAccountKeys not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listStorageAccountKeys`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
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
	respr, errRespr := w.srv.ListStorageAccountKeys(req.Context(), resourceGroupNameParam, workspaceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ListStorageAccountKeysResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkspacesServerTransport) dispatchBeginPrepareNotebook(req *http.Request) (*http.Response, error) {
	if w.srv.BeginPrepareNotebook == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPrepareNotebook not implemented")}
	}
	beginPrepareNotebook := w.beginPrepareNotebook.get(req)
	if beginPrepareNotebook == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/prepareNotebook`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
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
		respr, errRespr := w.srv.BeginPrepareNotebook(req.Context(), resourceGroupNameParam, workspaceNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPrepareNotebook = &respr
		w.beginPrepareNotebook.add(req, beginPrepareNotebook)
	}

	resp, err := server.PollerResponderNext(beginPrepareNotebook, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		w.beginPrepareNotebook.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPrepareNotebook) {
		w.beginPrepareNotebook.remove(req)
	}

	return resp, nil
}

func (w *WorkspacesServerTransport) dispatchBeginResyncKeys(req *http.Request) (*http.Response, error) {
	if w.srv.BeginResyncKeys == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginResyncKeys not implemented")}
	}
	beginResyncKeys := w.beginResyncKeys.get(req)
	if beginResyncKeys == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resyncKeys`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
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
		respr, errRespr := w.srv.BeginResyncKeys(req.Context(), resourceGroupNameParam, workspaceNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginResyncKeys = &respr
		w.beginResyncKeys.add(req, beginResyncKeys)
	}

	resp, err := server.PollerResponderNext(beginResyncKeys, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		w.beginResyncKeys.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginResyncKeys) {
		w.beginResyncKeys.remove(req)
	}

	return resp, nil
}

func (w *WorkspacesServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if w.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmachinelearningservices.WorkspaceUpdateParameters](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Update(req.Context(), resourceGroupNameParam, workspaceNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Workspace, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
