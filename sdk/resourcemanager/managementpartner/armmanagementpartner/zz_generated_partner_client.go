//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagementpartner

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

// PartnerClient contains the methods for the Partner group.
// Don't use this type directly, use NewPartnerClient() instead.
type PartnerClient struct {
	host string
	pl   runtime.Pipeline
}

// NewPartnerClient creates a new instance of PartnerClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPartnerClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*PartnerClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &PartnerClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// Create - Create a management partner for the objectId and tenantId.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-02-01
// partnerID - Id of the Partner
// options - PartnerClientCreateOptions contains the optional parameters for the PartnerClient.Create method.
func (client *PartnerClient) Create(ctx context.Context, partnerID string, options *PartnerClientCreateOptions) (PartnerClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, partnerID, options)
	if err != nil {
		return PartnerClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PartnerClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PartnerClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *PartnerClient) createCreateRequest(ctx context.Context, partnerID string, options *PartnerClientCreateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.ManagementPartner/partners/{partnerId}"
	if partnerID == "" {
		return nil, errors.New("parameter partnerID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerId}", url.PathEscape(partnerID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *PartnerClient) createHandleResponse(resp *http.Response) (PartnerClientCreateResponse, error) {
	result := PartnerClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PartnerResponse); err != nil {
		return PartnerClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Delete the management partner for the objectId and tenantId.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-02-01
// partnerID - Id of the Partner
// options - PartnerClientDeleteOptions contains the optional parameters for the PartnerClient.Delete method.
func (client *PartnerClient) Delete(ctx context.Context, partnerID string, options *PartnerClientDeleteOptions) (PartnerClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, partnerID, options)
	if err != nil {
		return PartnerClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PartnerClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PartnerClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return PartnerClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PartnerClient) deleteCreateRequest(ctx context.Context, partnerID string, options *PartnerClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.ManagementPartner/partners/{partnerId}"
	if partnerID == "" {
		return nil, errors.New("parameter partnerID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerId}", url.PathEscape(partnerID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the management partner using the partnerId, objectId and tenantId.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-02-01
// partnerID - Id of the Partner
// options - PartnerClientGetOptions contains the optional parameters for the PartnerClient.Get method.
func (client *PartnerClient) Get(ctx context.Context, partnerID string, options *PartnerClientGetOptions) (PartnerClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, partnerID, options)
	if err != nil {
		return PartnerClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PartnerClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PartnerClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PartnerClient) getCreateRequest(ctx context.Context, partnerID string, options *PartnerClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.ManagementPartner/partners/{partnerId}"
	if partnerID == "" {
		return nil, errors.New("parameter partnerID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerId}", url.PathEscape(partnerID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PartnerClient) getHandleResponse(resp *http.Response) (PartnerClientGetResponse, error) {
	result := PartnerClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PartnerResponse); err != nil {
		return PartnerClientGetResponse{}, err
	}
	return result, nil
}

// Update - Update the management partner for the objectId and tenantId.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-02-01
// partnerID - Id of the Partner
// options - PartnerClientUpdateOptions contains the optional parameters for the PartnerClient.Update method.
func (client *PartnerClient) Update(ctx context.Context, partnerID string, options *PartnerClientUpdateOptions) (PartnerClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, partnerID, options)
	if err != nil {
		return PartnerClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PartnerClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PartnerClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *PartnerClient) updateCreateRequest(ctx context.Context, partnerID string, options *PartnerClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.ManagementPartner/partners/{partnerId}"
	if partnerID == "" {
		return nil, errors.New("parameter partnerID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerId}", url.PathEscape(partnerID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *PartnerClient) updateHandleResponse(resp *http.Response) (PartnerClientUpdateResponse, error) {
	result := PartnerClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PartnerResponse); err != nil {
		return PartnerClientUpdateResponse{}, err
	}
	return result, nil
}