//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcognitiveservices

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// CognitiveServicesManagementClient contains the methods for the CognitiveServicesManagementClient group.
// Don't use this type directly, use NewCognitiveServicesManagementClient() instead.
type CognitiveServicesManagementClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewCognitiveServicesManagementClient creates a new instance of CognitiveServicesManagementClient with the specified values.
func NewCognitiveServicesManagementClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *CognitiveServicesManagementClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &CognitiveServicesManagementClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CheckDomainAvailability - Check whether a domain is available.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CognitiveServicesManagementClient) CheckDomainAvailability(ctx context.Context, parameters CheckDomainAvailabilityParameter, options *CognitiveServicesManagementClientCheckDomainAvailabilityOptions) (CognitiveServicesManagementClientCheckDomainAvailabilityResponse, error) {
	req, err := client.checkDomainAvailabilityCreateRequest(ctx, parameters, options)
	if err != nil {
		return CognitiveServicesManagementClientCheckDomainAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CognitiveServicesManagementClientCheckDomainAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CognitiveServicesManagementClientCheckDomainAvailabilityResponse{}, client.checkDomainAvailabilityHandleError(resp)
	}
	return client.checkDomainAvailabilityHandleResponse(resp)
}

// checkDomainAvailabilityCreateRequest creates the CheckDomainAvailability request.
func (client *CognitiveServicesManagementClient) checkDomainAvailabilityCreateRequest(ctx context.Context, parameters CheckDomainAvailabilityParameter, options *CognitiveServicesManagementClientCheckDomainAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.CognitiveServices/checkDomainAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// checkDomainAvailabilityHandleResponse handles the CheckDomainAvailability response.
func (client *CognitiveServicesManagementClient) checkDomainAvailabilityHandleResponse(resp *http.Response) (CognitiveServicesManagementClientCheckDomainAvailabilityResponse, error) {
	result := CognitiveServicesManagementClientCheckDomainAvailabilityResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DomainAvailability); err != nil {
		return CognitiveServicesManagementClientCheckDomainAvailabilityResponse{}, err
	}
	return result, nil
}

// checkDomainAvailabilityHandleError handles the CheckDomainAvailability error response.
func (client *CognitiveServicesManagementClient) checkDomainAvailabilityHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// CheckSKUAvailability - Check available SKUs.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CognitiveServicesManagementClient) CheckSKUAvailability(ctx context.Context, location string, parameters CheckSKUAvailabilityParameter, options *CognitiveServicesManagementClientCheckSKUAvailabilityOptions) (CognitiveServicesManagementClientCheckSKUAvailabilityResponse, error) {
	req, err := client.checkSKUAvailabilityCreateRequest(ctx, location, parameters, options)
	if err != nil {
		return CognitiveServicesManagementClientCheckSKUAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CognitiveServicesManagementClientCheckSKUAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CognitiveServicesManagementClientCheckSKUAvailabilityResponse{}, client.checkSKUAvailabilityHandleError(resp)
	}
	return client.checkSKUAvailabilityHandleResponse(resp)
}

// checkSKUAvailabilityCreateRequest creates the CheckSKUAvailability request.
func (client *CognitiveServicesManagementClient) checkSKUAvailabilityCreateRequest(ctx context.Context, location string, parameters CheckSKUAvailabilityParameter, options *CognitiveServicesManagementClientCheckSKUAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.CognitiveServices/locations/{location}/checkSkuAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// checkSKUAvailabilityHandleResponse handles the CheckSKUAvailability response.
func (client *CognitiveServicesManagementClient) checkSKUAvailabilityHandleResponse(resp *http.Response) (CognitiveServicesManagementClientCheckSKUAvailabilityResponse, error) {
	result := CognitiveServicesManagementClientCheckSKUAvailabilityResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SKUAvailabilityListResult); err != nil {
		return CognitiveServicesManagementClientCheckSKUAvailabilityResponse{}, err
	}
	return result, nil
}

// checkSKUAvailabilityHandleError handles the CheckSKUAvailability error response.
func (client *CognitiveServicesManagementClient) checkSKUAvailabilityHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
