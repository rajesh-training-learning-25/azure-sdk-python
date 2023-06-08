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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"net/http"
	"regexp"
)

// AccountsServer is a fake server for instances of the armstorage.AccountsClient type.
type AccountsServer struct {
	// BeginAbortHierarchicalNamespaceMigration is the fake for method AccountsClient.BeginAbortHierarchicalNamespaceMigration
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginAbortHierarchicalNamespaceMigration func(ctx context.Context, resourceGroupName string, accountName string, options *armstorage.AccountsClientBeginAbortHierarchicalNamespaceMigrationOptions) (resp azfake.PollerResponder[armstorage.AccountsClientAbortHierarchicalNamespaceMigrationResponse], errResp azfake.ErrorResponder)

	// CheckNameAvailability is the fake for method AccountsClient.CheckNameAvailability
	// HTTP status codes to indicate success: http.StatusOK
	CheckNameAvailability func(ctx context.Context, accountName armstorage.AccountCheckNameAvailabilityParameters, options *armstorage.AccountsClientCheckNameAvailabilityOptions) (resp azfake.Responder[armstorage.AccountsClientCheckNameAvailabilityResponse], errResp azfake.ErrorResponder)

	// BeginCreate is the fake for method AccountsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreate func(ctx context.Context, resourceGroupName string, accountName string, parameters armstorage.AccountCreateParameters, options *armstorage.AccountsClientBeginCreateOptions) (resp azfake.PollerResponder[armstorage.AccountsClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method AccountsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, accountName string, options *armstorage.AccountsClientDeleteOptions) (resp azfake.Responder[armstorage.AccountsClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginFailover is the fake for method AccountsClient.BeginFailover
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginFailover func(ctx context.Context, resourceGroupName string, accountName string, options *armstorage.AccountsClientBeginFailoverOptions) (resp azfake.PollerResponder[armstorage.AccountsClientFailoverResponse], errResp azfake.ErrorResponder)

	// GetProperties is the fake for method AccountsClient.GetProperties
	// HTTP status codes to indicate success: http.StatusOK
	GetProperties func(ctx context.Context, resourceGroupName string, accountName string, options *armstorage.AccountsClientGetPropertiesOptions) (resp azfake.Responder[armstorage.AccountsClientGetPropertiesResponse], errResp azfake.ErrorResponder)

	// BeginHierarchicalNamespaceMigration is the fake for method AccountsClient.BeginHierarchicalNamespaceMigration
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginHierarchicalNamespaceMigration func(ctx context.Context, resourceGroupName string, accountName string, requestType string, options *armstorage.AccountsClientBeginHierarchicalNamespaceMigrationOptions) (resp azfake.PollerResponder[armstorage.AccountsClientHierarchicalNamespaceMigrationResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method AccountsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armstorage.AccountsClientListOptions) (resp azfake.PagerResponder[armstorage.AccountsClientListResponse])

	// ListAccountSAS is the fake for method AccountsClient.ListAccountSAS
	// HTTP status codes to indicate success: http.StatusOK
	ListAccountSAS func(ctx context.Context, resourceGroupName string, accountName string, parameters armstorage.AccountSasParameters, options *armstorage.AccountsClientListAccountSASOptions) (resp azfake.Responder[armstorage.AccountsClientListAccountSASResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method AccountsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armstorage.AccountsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armstorage.AccountsClientListByResourceGroupResponse])

	// ListKeys is the fake for method AccountsClient.ListKeys
	// HTTP status codes to indicate success: http.StatusOK
	ListKeys func(ctx context.Context, resourceGroupName string, accountName string, options *armstorage.AccountsClientListKeysOptions) (resp azfake.Responder[armstorage.AccountsClientListKeysResponse], errResp azfake.ErrorResponder)

	// ListServiceSAS is the fake for method AccountsClient.ListServiceSAS
	// HTTP status codes to indicate success: http.StatusOK
	ListServiceSAS func(ctx context.Context, resourceGroupName string, accountName string, parameters armstorage.ServiceSasParameters, options *armstorage.AccountsClientListServiceSASOptions) (resp azfake.Responder[armstorage.AccountsClientListServiceSASResponse], errResp azfake.ErrorResponder)

	// RegenerateKey is the fake for method AccountsClient.RegenerateKey
	// HTTP status codes to indicate success: http.StatusOK
	RegenerateKey func(ctx context.Context, resourceGroupName string, accountName string, regenerateKey armstorage.AccountRegenerateKeyParameters, options *armstorage.AccountsClientRegenerateKeyOptions) (resp azfake.Responder[armstorage.AccountsClientRegenerateKeyResponse], errResp azfake.ErrorResponder)

	// BeginRestoreBlobRanges is the fake for method AccountsClient.BeginRestoreBlobRanges
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginRestoreBlobRanges func(ctx context.Context, resourceGroupName string, accountName string, parameters armstorage.BlobRestoreParameters, options *armstorage.AccountsClientBeginRestoreBlobRangesOptions) (resp azfake.PollerResponder[armstorage.AccountsClientRestoreBlobRangesResponse], errResp azfake.ErrorResponder)

	// RevokeUserDelegationKeys is the fake for method AccountsClient.RevokeUserDelegationKeys
	// HTTP status codes to indicate success: http.StatusOK
	RevokeUserDelegationKeys func(ctx context.Context, resourceGroupName string, accountName string, options *armstorage.AccountsClientRevokeUserDelegationKeysOptions) (resp azfake.Responder[armstorage.AccountsClientRevokeUserDelegationKeysResponse], errResp azfake.ErrorResponder)

	// Update is the fake for method AccountsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, accountName string, parameters armstorage.AccountUpdateParameters, options *armstorage.AccountsClientUpdateOptions) (resp azfake.Responder[armstorage.AccountsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewAccountsServerTransport creates a new instance of AccountsServerTransport with the provided implementation.
// The returned AccountsServerTransport instance is connected to an instance of armstorage.AccountsClient by way of the
// undefined.Transporter field.
func NewAccountsServerTransport(srv *AccountsServer) *AccountsServerTransport {
	return &AccountsServerTransport{srv: srv}
}

// AccountsServerTransport connects instances of armstorage.AccountsClient to instances of AccountsServer.
// Don't use this type directly, use NewAccountsServerTransport instead.
type AccountsServerTransport struct {
	srv                                      *AccountsServer
	beginAbortHierarchicalNamespaceMigration *azfake.PollerResponder[armstorage.AccountsClientAbortHierarchicalNamespaceMigrationResponse]
	beginCreate                              *azfake.PollerResponder[armstorage.AccountsClientCreateResponse]
	beginFailover                            *azfake.PollerResponder[armstorage.AccountsClientFailoverResponse]
	beginHierarchicalNamespaceMigration      *azfake.PollerResponder[armstorage.AccountsClientHierarchicalNamespaceMigrationResponse]
	newListPager                             *azfake.PagerResponder[armstorage.AccountsClientListResponse]
	newListByResourceGroupPager              *azfake.PagerResponder[armstorage.AccountsClientListByResourceGroupResponse]
	beginRestoreBlobRanges                   *azfake.PollerResponder[armstorage.AccountsClientRestoreBlobRangesResponse]
}

// Do implements the policy.Transporter interface for AccountsServerTransport.
func (a *AccountsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AccountsClient.BeginAbortHierarchicalNamespaceMigration":
		resp, err = a.dispatchBeginAbortHierarchicalNamespaceMigration(req)
	case "AccountsClient.CheckNameAvailability":
		resp, err = a.dispatchCheckNameAvailability(req)
	case "AccountsClient.BeginCreate":
		resp, err = a.dispatchBeginCreate(req)
	case "AccountsClient.Delete":
		resp, err = a.dispatchDelete(req)
	case "AccountsClient.BeginFailover":
		resp, err = a.dispatchBeginFailover(req)
	case "AccountsClient.GetProperties":
		resp, err = a.dispatchGetProperties(req)
	case "AccountsClient.BeginHierarchicalNamespaceMigration":
		resp, err = a.dispatchBeginHierarchicalNamespaceMigration(req)
	case "AccountsClient.NewListPager":
		resp, err = a.dispatchNewListPager(req)
	case "AccountsClient.ListAccountSAS":
		resp, err = a.dispatchListAccountSAS(req)
	case "AccountsClient.NewListByResourceGroupPager":
		resp, err = a.dispatchNewListByResourceGroupPager(req)
	case "AccountsClient.ListKeys":
		resp, err = a.dispatchListKeys(req)
	case "AccountsClient.ListServiceSAS":
		resp, err = a.dispatchListServiceSAS(req)
	case "AccountsClient.RegenerateKey":
		resp, err = a.dispatchRegenerateKey(req)
	case "AccountsClient.BeginRestoreBlobRanges":
		resp, err = a.dispatchBeginRestoreBlobRanges(req)
	case "AccountsClient.RevokeUserDelegationKeys":
		resp, err = a.dispatchRevokeUserDelegationKeys(req)
	case "AccountsClient.Update":
		resp, err = a.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AccountsServerTransport) dispatchBeginAbortHierarchicalNamespaceMigration(req *http.Request) (*http.Response, error) {
	if a.srv.BeginAbortHierarchicalNamespaceMigration == nil {
		return nil, &nonRetriableError{errors.New("method BeginAbortHierarchicalNamespaceMigration not implemented")}
	}
	if a.beginAbortHierarchicalNamespaceMigration == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourcegroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[a-zA-Z0-9-_]+)/aborthnsonmigration"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := a.srv.BeginAbortHierarchicalNamespaceMigration(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("accountName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		a.beginAbortHierarchicalNamespaceMigration = &respr
	}

	resp, err := server.PollerResponderNext(a.beginAbortHierarchicalNamespaceMigration, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(a.beginAbortHierarchicalNamespaceMigration) {
		a.beginAbortHierarchicalNamespaceMigration = nil
	}

	return resp, nil
}

func (a *AccountsServerTransport) dispatchCheckNameAvailability(req *http.Request) (*http.Response, error) {
	if a.srv.CheckNameAvailability == nil {
		return nil, &nonRetriableError{errors.New("method CheckNameAvailability not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Storage/checkNameAvailability"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armstorage.AccountCheckNameAvailabilityParameters](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.CheckNameAvailability(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CheckNameAvailabilityResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AccountsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if a.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreate not implemented")}
	}
	if a.beginCreate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armstorage.AccountCreateParameters](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginCreate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("accountName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		a.beginCreate = &respr
	}

	resp, err := server.PollerResponderNext(a.beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(a.beginCreate) {
		a.beginCreate = nil
	}

	return resp, nil
}

func (a *AccountsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if a.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("method Delete not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := a.srv.Delete(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("accountName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AccountsServerTransport) dispatchBeginFailover(req *http.Request) (*http.Response, error) {
	if a.srv.BeginFailover == nil {
		return nil, &nonRetriableError{errors.New("method BeginFailover not implemented")}
	}
	if a.beginFailover == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[a-zA-Z0-9-_]+)/failover"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		failoverTypeParam := getOptional(qp.Get("failoverType"))
		var options *armstorage.AccountsClientBeginFailoverOptions
		if failoverTypeParam != nil {
			options = &armstorage.AccountsClientBeginFailoverOptions{
				FailoverType: failoverTypeParam,
			}
		}
		respr, errRespr := a.srv.BeginFailover(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("accountName")], options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		a.beginFailover = &respr
	}

	resp, err := server.PollerResponderNext(a.beginFailover, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(a.beginFailover) {
		a.beginFailover = nil
	}

	return resp, nil
}

func (a *AccountsServerTransport) dispatchGetProperties(req *http.Request) (*http.Response, error) {
	if a.srv.GetProperties == nil {
		return nil, &nonRetriableError{errors.New("method GetProperties not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	expandParam := getOptional(armstorage.StorageAccountExpand(qp.Get("$expand")))
	var options *armstorage.AccountsClientGetPropertiesOptions
	if expandParam != nil {
		options = &armstorage.AccountsClientGetPropertiesOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := a.srv.GetProperties(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("accountName")], options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Account, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AccountsServerTransport) dispatchBeginHierarchicalNamespaceMigration(req *http.Request) (*http.Response, error) {
	if a.srv.BeginHierarchicalNamespaceMigration == nil {
		return nil, &nonRetriableError{errors.New("method BeginHierarchicalNamespaceMigration not implemented")}
	}
	if a.beginHierarchicalNamespaceMigration == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourcegroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[a-zA-Z0-9-_]+)/hnsonmigration"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		respr, errRespr := a.srv.BeginHierarchicalNamespaceMigration(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("accountName")], qp.Get("requestType"), nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		a.beginHierarchicalNamespaceMigration = &respr
	}

	resp, err := server.PollerResponderNext(a.beginHierarchicalNamespaceMigration, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(a.beginHierarchicalNamespaceMigration) {
		a.beginHierarchicalNamespaceMigration = nil
	}

	return resp, nil
}

func (a *AccountsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if a.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Storage/storageAccounts"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := a.srv.NewListPager(nil)
		a.newListPager = &resp
		server.PagerResponderInjectNextLinks(a.newListPager, req, func(page *armstorage.AccountsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(a.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(a.newListPager) {
		a.newListPager = nil
	}
	return resp, nil
}

func (a *AccountsServerTransport) dispatchListAccountSAS(req *http.Request) (*http.Response, error) {
	if a.srv.ListAccountSAS == nil {
		return nil, &nonRetriableError{errors.New("method ListAccountSAS not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[a-zA-Z0-9-_]+)/ListAccountSas"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armstorage.AccountSasParameters](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.ListAccountSAS(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("accountName")], body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ListAccountSasResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AccountsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListByResourceGroupPager not implemented")}
	}
	if a.newListByResourceGroupPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Storage/storageAccounts"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := a.srv.NewListByResourceGroupPager(matches[regex.SubexpIndex("resourceGroupName")], nil)
		a.newListByResourceGroupPager = &resp
		server.PagerResponderInjectNextLinks(a.newListByResourceGroupPager, req, func(page *armstorage.AccountsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(a.newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(a.newListByResourceGroupPager) {
		a.newListByResourceGroupPager = nil
	}
	return resp, nil
}

func (a *AccountsServerTransport) dispatchListKeys(req *http.Request) (*http.Response, error) {
	if a.srv.ListKeys == nil {
		return nil, &nonRetriableError{errors.New("method ListKeys not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[a-zA-Z0-9-_]+)/listKeys"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	expandParam := getOptional(qp.Get("$expand"))
	var options *armstorage.AccountsClientListKeysOptions
	if expandParam != nil {
		options = &armstorage.AccountsClientListKeysOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := a.srv.ListKeys(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("accountName")], options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AccountListKeysResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AccountsServerTransport) dispatchListServiceSAS(req *http.Request) (*http.Response, error) {
	if a.srv.ListServiceSAS == nil {
		return nil, &nonRetriableError{errors.New("method ListServiceSAS not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[a-zA-Z0-9-_]+)/ListServiceSas"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armstorage.ServiceSasParameters](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.ListServiceSAS(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("accountName")], body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ListServiceSasResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AccountsServerTransport) dispatchRegenerateKey(req *http.Request) (*http.Response, error) {
	if a.srv.RegenerateKey == nil {
		return nil, &nonRetriableError{errors.New("method RegenerateKey not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[a-zA-Z0-9-_]+)/regenerateKey"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armstorage.AccountRegenerateKeyParameters](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.RegenerateKey(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("accountName")], body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AccountListKeysResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AccountsServerTransport) dispatchBeginRestoreBlobRanges(req *http.Request) (*http.Response, error) {
	if a.srv.BeginRestoreBlobRanges == nil {
		return nil, &nonRetriableError{errors.New("method BeginRestoreBlobRanges not implemented")}
	}
	if a.beginRestoreBlobRanges == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[a-zA-Z0-9-_]+)/restoreBlobRanges"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armstorage.BlobRestoreParameters](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginRestoreBlobRanges(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("accountName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		a.beginRestoreBlobRanges = &respr
	}

	resp, err := server.PollerResponderNext(a.beginRestoreBlobRanges, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(a.beginRestoreBlobRanges) {
		a.beginRestoreBlobRanges = nil
	}

	return resp, nil
}

func (a *AccountsServerTransport) dispatchRevokeUserDelegationKeys(req *http.Request) (*http.Response, error) {
	if a.srv.RevokeUserDelegationKeys == nil {
		return nil, &nonRetriableError{errors.New("method RevokeUserDelegationKeys not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[a-zA-Z0-9-_]+)/revokeUserDelegationKeys"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := a.srv.RevokeUserDelegationKeys(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("accountName")], nil)
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

func (a *AccountsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("method Update not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armstorage.AccountUpdateParameters](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Update(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("accountName")], body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Account, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
