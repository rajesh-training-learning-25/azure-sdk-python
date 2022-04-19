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
	"strconv"
	"strings"
)

// GenerateClient contains the methods for the Generate group.
// Don't use this type directly, use NewGenerateClient() instead.
type GenerateClient struct {
	host string
	pl   runtime.Pipeline
}

// NewGenerateClient creates a new instance of GenerateClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewGenerateClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*GenerateClient, error) {
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
	client := &GenerateClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// InviteCode - Generate invite code for a lab
// If the operation fails it returns an *azcore.ResponseError type.
// billingAccountName - Billing account name.
// billingProfileName - Billing profile name.
// invoiceSectionName - Invoice section name.
// parameters - Request parameters that are provided to generate invite code.
// options - GenerateClientInviteCodeOptions contains the optional parameters for the GenerateClient.InviteCode method.
func (client *GenerateClient) InviteCode(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, parameters InviteCodeGenerateRequest, options *GenerateClientInviteCodeOptions) (GenerateClientInviteCodeResponse, error) {
	req, err := client.inviteCodeCreateRequest(ctx, billingAccountName, billingProfileName, invoiceSectionName, parameters, options)
	if err != nil {
		return GenerateClientInviteCodeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GenerateClientInviteCodeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GenerateClientInviteCodeResponse{}, runtime.NewResponseError(resp)
	}
	return client.inviteCodeHandleResponse(resp)
}

// inviteCodeCreateRequest creates the InviteCode request.
func (client *GenerateClient) inviteCodeCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, parameters InviteCodeGenerateRequest, options *GenerateClientInviteCodeOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/providers/Microsoft.Education/labs/default/generateInviteCode"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.OnlyUpdateStudentCountParameter != nil {
		reqQP.Set("onlyUpdateStudentCountParameter", strconv.FormatBool(*options.OnlyUpdateStudentCountParameter))
	}
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// inviteCodeHandleResponse handles the InviteCode response.
func (client *GenerateClient) inviteCodeHandleResponse(resp *http.Response) (GenerateClientInviteCodeResponse, error) {
	result := GenerateClientInviteCodeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LabDetails); err != nil {
		return GenerateClientInviteCodeResponse{}, err
	}
	return result, nil
}
