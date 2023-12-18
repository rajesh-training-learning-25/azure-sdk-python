//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v3"
	"net/http"
	"net/url"
	"regexp"
)

// BackupProtectableItemsServer is a fake server for instances of the armrecoveryservicesbackup.BackupProtectableItemsClient type.
type BackupProtectableItemsServer struct {
	// NewListPager is the fake for method BackupProtectableItemsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(vaultName string, resourceGroupName string, options *armrecoveryservicesbackup.BackupProtectableItemsClientListOptions) (resp azfake.PagerResponder[armrecoveryservicesbackup.BackupProtectableItemsClientListResponse])
}

// NewBackupProtectableItemsServerTransport creates a new instance of BackupProtectableItemsServerTransport with the provided implementation.
// The returned BackupProtectableItemsServerTransport instance is connected to an instance of armrecoveryservicesbackup.BackupProtectableItemsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBackupProtectableItemsServerTransport(srv *BackupProtectableItemsServer) *BackupProtectableItemsServerTransport {
	return &BackupProtectableItemsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armrecoveryservicesbackup.BackupProtectableItemsClientListResponse]](),
	}
}

// BackupProtectableItemsServerTransport connects instances of armrecoveryservicesbackup.BackupProtectableItemsClient to instances of BackupProtectableItemsServer.
// Don't use this type directly, use NewBackupProtectableItemsServerTransport instead.
type BackupProtectableItemsServerTransport struct {
	srv          *BackupProtectableItemsServer
	newListPager *tracker[azfake.PagerResponder[armrecoveryservicesbackup.BackupProtectableItemsClientListResponse]]
}

// Do implements the policy.Transporter interface for BackupProtectableItemsServerTransport.
func (b *BackupProtectableItemsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "BackupProtectableItemsClient.NewListPager":
		resp, err = b.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *BackupProtectableItemsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := b.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupProtectableItems`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		var options *armrecoveryservicesbackup.BackupProtectableItemsClientListOptions
		if filterParam != nil || skipTokenParam != nil {
			options = &armrecoveryservicesbackup.BackupProtectableItemsClientListOptions{
				Filter:    filterParam,
				SkipToken: skipTokenParam,
			}
		}
		resp := b.srv.NewListPager(vaultNameParam, resourceGroupNameParam, options)
		newListPager = &resp
		b.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armrecoveryservicesbackup.BackupProtectableItemsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		b.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		b.newListPager.remove(req)
	}
	return resp, nil
}
