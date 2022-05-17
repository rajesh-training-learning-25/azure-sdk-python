//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armproviderhub

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

// ProviderRegistrationsClient contains the methods for the ProviderRegistrations group.
// Don't use this type directly, use NewProviderRegistrationsClient() instead.
type ProviderRegistrationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewProviderRegistrationsClient creates a new instance of ProviderRegistrationsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewProviderRegistrationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ProviderRegistrationsClient, error) {
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
	client := &ProviderRegistrationsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates the provider registration.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-20
// providerNamespace - The name of the resource provider hosted within ProviderHub.
// properties - The provider registration properties supplied to the CreateOrUpdate operation.
// options - ProviderRegistrationsClientBeginCreateOrUpdateOptions contains the optional parameters for the ProviderRegistrationsClient.BeginCreateOrUpdate
// method.
func (client *ProviderRegistrationsClient) BeginCreateOrUpdate(ctx context.Context, providerNamespace string, properties ProviderRegistration, options *ProviderRegistrationsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ProviderRegistrationsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, providerNamespace, properties, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[ProviderRegistrationsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ProviderRegistrationsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates the provider registration.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-20
func (client *ProviderRegistrationsClient) createOrUpdate(ctx context.Context, providerNamespace string, properties ProviderRegistration, options *ProviderRegistrationsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, providerNamespace, properties, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ProviderRegistrationsClient) createOrUpdateCreateRequest(ctx context.Context, providerNamespace string, properties ProviderRegistration, options *ProviderRegistrationsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ProviderHub/providerRegistrations/{providerNamespace}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, properties)
}

// Delete - Deletes a provider registration.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-20
// providerNamespace - The name of the resource provider hosted within ProviderHub.
// options - ProviderRegistrationsClientDeleteOptions contains the optional parameters for the ProviderRegistrationsClient.Delete
// method.
func (client *ProviderRegistrationsClient) Delete(ctx context.Context, providerNamespace string, options *ProviderRegistrationsClientDeleteOptions) (ProviderRegistrationsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, providerNamespace, options)
	if err != nil {
		return ProviderRegistrationsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProviderRegistrationsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ProviderRegistrationsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ProviderRegistrationsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ProviderRegistrationsClient) deleteCreateRequest(ctx context.Context, providerNamespace string, options *ProviderRegistrationsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ProviderHub/providerRegistrations/{providerNamespace}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GenerateOperations - Generates the operations api for the given provider.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-20
// providerNamespace - The name of the resource provider hosted within ProviderHub.
// options - ProviderRegistrationsClientGenerateOperationsOptions contains the optional parameters for the ProviderRegistrationsClient.GenerateOperations
// method.
func (client *ProviderRegistrationsClient) GenerateOperations(ctx context.Context, providerNamespace string, options *ProviderRegistrationsClientGenerateOperationsOptions) (ProviderRegistrationsClientGenerateOperationsResponse, error) {
	req, err := client.generateOperationsCreateRequest(ctx, providerNamespace, options)
	if err != nil {
		return ProviderRegistrationsClientGenerateOperationsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProviderRegistrationsClientGenerateOperationsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProviderRegistrationsClientGenerateOperationsResponse{}, runtime.NewResponseError(resp)
	}
	return client.generateOperationsHandleResponse(resp)
}

// generateOperationsCreateRequest creates the GenerateOperations request.
func (client *ProviderRegistrationsClient) generateOperationsCreateRequest(ctx context.Context, providerNamespace string, options *ProviderRegistrationsClientGenerateOperationsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ProviderHub/providerRegistrations/{providerNamespace}/generateOperations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// generateOperationsHandleResponse handles the GenerateOperations response.
func (client *ProviderRegistrationsClient) generateOperationsHandleResponse(resp *http.Response) (ProviderRegistrationsClientGenerateOperationsResponse, error) {
	result := ProviderRegistrationsClientGenerateOperationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationsDefinitionArray); err != nil {
		return ProviderRegistrationsClientGenerateOperationsResponse{}, err
	}
	return result, nil
}

// Get - Gets the provider registration details.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-20
// providerNamespace - The name of the resource provider hosted within ProviderHub.
// options - ProviderRegistrationsClientGetOptions contains the optional parameters for the ProviderRegistrationsClient.Get
// method.
func (client *ProviderRegistrationsClient) Get(ctx context.Context, providerNamespace string, options *ProviderRegistrationsClientGetOptions) (ProviderRegistrationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, providerNamespace, options)
	if err != nil {
		return ProviderRegistrationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProviderRegistrationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProviderRegistrationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ProviderRegistrationsClient) getCreateRequest(ctx context.Context, providerNamespace string, options *ProviderRegistrationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ProviderHub/providerRegistrations/{providerNamespace}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProviderRegistrationsClient) getHandleResponse(resp *http.Response) (ProviderRegistrationsClientGetResponse, error) {
	result := ProviderRegistrationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProviderRegistration); err != nil {
		return ProviderRegistrationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets the list of the provider registrations in the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-20
// options - ProviderRegistrationsClientListOptions contains the optional parameters for the ProviderRegistrationsClient.List
// method.
func (client *ProviderRegistrationsClient) NewListPager(options *ProviderRegistrationsClientListOptions) *runtime.Pager[ProviderRegistrationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProviderRegistrationsClientListResponse]{
		More: func(page ProviderRegistrationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProviderRegistrationsClientListResponse) (ProviderRegistrationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ProviderRegistrationsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ProviderRegistrationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ProviderRegistrationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ProviderRegistrationsClient) listCreateRequest(ctx context.Context, options *ProviderRegistrationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ProviderHub/providerRegistrations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ProviderRegistrationsClient) listHandleResponse(resp *http.Response) (ProviderRegistrationsClientListResponse, error) {
	result := ProviderRegistrationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProviderRegistrationArrayResponseWithContinuation); err != nil {
		return ProviderRegistrationsClientListResponse{}, err
	}
	return result, nil
}
