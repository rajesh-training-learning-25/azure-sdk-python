//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogic

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

// IntegrationAccountAssembliesClient contains the methods for the IntegrationAccountAssemblies group.
// Don't use this type directly, use NewIntegrationAccountAssembliesClient() instead.
type IntegrationAccountAssembliesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewIntegrationAccountAssembliesClient creates a new instance of IntegrationAccountAssembliesClient with the specified values.
func NewIntegrationAccountAssembliesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *IntegrationAccountAssembliesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &IntegrationAccountAssembliesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CreateOrUpdate - Create or update an assembly for an integration account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationAccountAssembliesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, integrationAccountName string, assemblyArtifactName string, assemblyArtifact AssemblyDefinition, options *IntegrationAccountAssembliesCreateOrUpdateOptions) (IntegrationAccountAssembliesCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, integrationAccountName, assemblyArtifactName, assemblyArtifact, options)
	if err != nil {
		return IntegrationAccountAssembliesCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountAssembliesCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return IntegrationAccountAssembliesCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IntegrationAccountAssembliesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, assemblyArtifactName string, assemblyArtifact AssemblyDefinition, options *IntegrationAccountAssembliesCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/assemblies/{assemblyArtifactName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if assemblyArtifactName == "" {
		return nil, errors.New("parameter assemblyArtifactName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assemblyArtifactName}", url.PathEscape(assemblyArtifactName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, assemblyArtifact)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *IntegrationAccountAssembliesClient) createOrUpdateHandleResponse(resp *http.Response) (IntegrationAccountAssembliesCreateOrUpdateResponse, error) {
	result := IntegrationAccountAssembliesCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssemblyDefinition); err != nil {
		return IntegrationAccountAssembliesCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *IntegrationAccountAssembliesClient) createOrUpdateHandleError(resp *http.Response) error {
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

// Delete - Delete an assembly for an integration account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationAccountAssembliesClient) Delete(ctx context.Context, resourceGroupName string, integrationAccountName string, assemblyArtifactName string, options *IntegrationAccountAssembliesDeleteOptions) (IntegrationAccountAssembliesDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, integrationAccountName, assemblyArtifactName, options)
	if err != nil {
		return IntegrationAccountAssembliesDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountAssembliesDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return IntegrationAccountAssembliesDeleteResponse{}, client.deleteHandleError(resp)
	}
	return IntegrationAccountAssembliesDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IntegrationAccountAssembliesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, assemblyArtifactName string, options *IntegrationAccountAssembliesDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/assemblies/{assemblyArtifactName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if assemblyArtifactName == "" {
		return nil, errors.New("parameter assemblyArtifactName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assemblyArtifactName}", url.PathEscape(assemblyArtifactName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *IntegrationAccountAssembliesClient) deleteHandleError(resp *http.Response) error {
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

// Get - Get an assembly for an integration account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationAccountAssembliesClient) Get(ctx context.Context, resourceGroupName string, integrationAccountName string, assemblyArtifactName string, options *IntegrationAccountAssembliesGetOptions) (IntegrationAccountAssembliesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, integrationAccountName, assemblyArtifactName, options)
	if err != nil {
		return IntegrationAccountAssembliesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountAssembliesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationAccountAssembliesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IntegrationAccountAssembliesClient) getCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, assemblyArtifactName string, options *IntegrationAccountAssembliesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/assemblies/{assemblyArtifactName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if assemblyArtifactName == "" {
		return nil, errors.New("parameter assemblyArtifactName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assemblyArtifactName}", url.PathEscape(assemblyArtifactName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IntegrationAccountAssembliesClient) getHandleResponse(resp *http.Response) (IntegrationAccountAssembliesGetResponse, error) {
	result := IntegrationAccountAssembliesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssemblyDefinition); err != nil {
		return IntegrationAccountAssembliesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *IntegrationAccountAssembliesClient) getHandleError(resp *http.Response) error {
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

// List - List the assemblies for an integration account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationAccountAssembliesClient) List(ctx context.Context, resourceGroupName string, integrationAccountName string, options *IntegrationAccountAssembliesListOptions) (IntegrationAccountAssembliesListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, integrationAccountName, options)
	if err != nil {
		return IntegrationAccountAssembliesListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountAssembliesListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationAccountAssembliesListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *IntegrationAccountAssembliesClient) listCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, options *IntegrationAccountAssembliesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/assemblies"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IntegrationAccountAssembliesClient) listHandleResponse(resp *http.Response) (IntegrationAccountAssembliesListResponse, error) {
	result := IntegrationAccountAssembliesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssemblyCollection); err != nil {
		return IntegrationAccountAssembliesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *IntegrationAccountAssembliesClient) listHandleError(resp *http.Response) error {
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

// ListContentCallbackURL - Get the content callback url for an integration account assembly.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationAccountAssembliesClient) ListContentCallbackURL(ctx context.Context, resourceGroupName string, integrationAccountName string, assemblyArtifactName string, options *IntegrationAccountAssembliesListContentCallbackURLOptions) (IntegrationAccountAssembliesListContentCallbackURLResponse, error) {
	req, err := client.listContentCallbackURLCreateRequest(ctx, resourceGroupName, integrationAccountName, assemblyArtifactName, options)
	if err != nil {
		return IntegrationAccountAssembliesListContentCallbackURLResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountAssembliesListContentCallbackURLResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationAccountAssembliesListContentCallbackURLResponse{}, client.listContentCallbackURLHandleError(resp)
	}
	return client.listContentCallbackURLHandleResponse(resp)
}

// listContentCallbackURLCreateRequest creates the ListContentCallbackURL request.
func (client *IntegrationAccountAssembliesClient) listContentCallbackURLCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, assemblyArtifactName string, options *IntegrationAccountAssembliesListContentCallbackURLOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/assemblies/{assemblyArtifactName}/listContentCallbackUrl"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if assemblyArtifactName == "" {
		return nil, errors.New("parameter assemblyArtifactName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assemblyArtifactName}", url.PathEscape(assemblyArtifactName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listContentCallbackURLHandleResponse handles the ListContentCallbackURL response.
func (client *IntegrationAccountAssembliesClient) listContentCallbackURLHandleResponse(resp *http.Response) (IntegrationAccountAssembliesListContentCallbackURLResponse, error) {
	result := IntegrationAccountAssembliesListContentCallbackURLResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowTriggerCallbackURL); err != nil {
		return IntegrationAccountAssembliesListContentCallbackURLResponse{}, err
	}
	return result, nil
}

// listContentCallbackURLHandleError handles the ListContentCallbackURL error response.
func (client *IntegrationAccountAssembliesClient) listContentCallbackURLHandleError(resp *http.Response) error {
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
