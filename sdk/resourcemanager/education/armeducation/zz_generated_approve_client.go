//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeducation

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ApproveClient contains the methods for the Approve group.
// Don't use this type directly, use NewApproveClient() instead.
type ApproveClient struct {
	host string
	pl   runtime.Pipeline
}

// NewApproveClient creates a new instance of ApproveClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewApproveClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ApproveClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ApproveClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// Invite - Approve student joining the redeemable lab
// If the operation fails it returns an *azcore.ResponseError type.
// billingAccountName - Billing account name.
// billingProfileName - Billing profile name.
// invoiceSectionName - Invoice section name.
// joinRequestName - Join name
// options - ApproveClientInviteOptions contains the optional parameters for the ApproveClient.Invite method.
func (client *ApproveClient) Invite(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, joinRequestName string, options *ApproveClientInviteOptions) (ApproveClientInviteResponse, error) {
	req, err := client.inviteCreateRequest(ctx, billingAccountName, billingProfileName, invoiceSectionName, joinRequestName, options)
	if err != nil {
		return ApproveClientInviteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApproveClientInviteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApproveClientInviteResponse{}, runtime.NewResponseError(resp)
	}
	return ApproveClientInviteResponse{}, nil
}

// inviteCreateRequest creates the Invite request.
func (client *ApproveClient) inviteCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, joinRequestName string, options *ApproveClientInviteOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/providers/Microsoft.Education/labs/default/joinRequests/{joinRequestName}/approve"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	if invoiceSectionName == "" {
		return nil, errors.New("parameter invoiceSectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{invoiceSectionName}", url.PathEscape(invoiceSectionName))
	if joinRequestName == "" {
		return nil, errors.New("parameter joinRequestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{joinRequestName}", url.PathEscape(joinRequestName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}
