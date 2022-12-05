//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicebus

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

// MigrationConfigsClient contains the methods for the MigrationConfigs group.
// Don't use this type directly, use NewMigrationConfigsClient() instead.
type MigrationConfigsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewMigrationConfigsClient creates a new instance of MigrationConfigsClient with the specified values.
// subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewMigrationConfigsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MigrationConfigsClient, error) {
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
	client := &MigrationConfigsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CompleteMigration - This operation Completes Migration of entities by pointing the connection strings to Premium namespace
// and any entities created after the operation will be under Premium Namespace. CompleteMigration
// operation will fail when entity migration is in-progress.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// resourceGroupName - Name of the Resource group within the Azure subscription.
// namespaceName - The namespace name
// configName - The configuration name. Should always be "$default".
// options - MigrationConfigsClientCompleteMigrationOptions contains the optional parameters for the MigrationConfigsClient.CompleteMigration
// method.
func (client *MigrationConfigsClient) CompleteMigration(ctx context.Context, resourceGroupName string, namespaceName string, configName MigrationConfigurationName, options *MigrationConfigsClientCompleteMigrationOptions) (MigrationConfigsClientCompleteMigrationResponse, error) {
	req, err := client.completeMigrationCreateRequest(ctx, resourceGroupName, namespaceName, configName, options)
	if err != nil {
		return MigrationConfigsClientCompleteMigrationResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MigrationConfigsClientCompleteMigrationResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MigrationConfigsClientCompleteMigrationResponse{}, runtime.NewResponseError(resp)
	}
	return MigrationConfigsClientCompleteMigrationResponse{}, nil
}

// completeMigrationCreateRequest creates the CompleteMigration request.
func (client *MigrationConfigsClient) completeMigrationCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, configName MigrationConfigurationName, options *MigrationConfigsClientCompleteMigrationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/migrationConfigurations/{configName}/upgrade"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if configName == "" {
		return nil, errors.New("parameter configName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configName}", url.PathEscape(string(configName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginCreateAndStartMigration - Creates Migration configuration and starts migration of entities from Standard to Premium
// namespace
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// resourceGroupName - Name of the Resource group within the Azure subscription.
// namespaceName - The namespace name
// configName - The configuration name. Should always be "$default".
// parameters - Parameters required to create Migration Configuration
// options - MigrationConfigsClientBeginCreateAndStartMigrationOptions contains the optional parameters for the MigrationConfigsClient.BeginCreateAndStartMigration
// method.
func (client *MigrationConfigsClient) BeginCreateAndStartMigration(ctx context.Context, resourceGroupName string, namespaceName string, configName MigrationConfigurationName, parameters MigrationConfigProperties, options *MigrationConfigsClientBeginCreateAndStartMigrationOptions) (*runtime.Poller[MigrationConfigsClientCreateAndStartMigrationResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createAndStartMigration(ctx, resourceGroupName, namespaceName, configName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[MigrationConfigsClientCreateAndStartMigrationResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[MigrationConfigsClientCreateAndStartMigrationResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateAndStartMigration - Creates Migration configuration and starts migration of entities from Standard to Premium namespace
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
func (client *MigrationConfigsClient) createAndStartMigration(ctx context.Context, resourceGroupName string, namespaceName string, configName MigrationConfigurationName, parameters MigrationConfigProperties, options *MigrationConfigsClientBeginCreateAndStartMigrationOptions) (*http.Response, error) {
	req, err := client.createAndStartMigrationCreateRequest(ctx, resourceGroupName, namespaceName, configName, parameters, options)
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

// createAndStartMigrationCreateRequest creates the CreateAndStartMigration request.
func (client *MigrationConfigsClient) createAndStartMigrationCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, configName MigrationConfigurationName, parameters MigrationConfigProperties, options *MigrationConfigsClientBeginCreateAndStartMigrationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/migrationConfigurations/{configName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if configName == "" {
		return nil, errors.New("parameter configName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configName}", url.PathEscape(string(configName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// Delete - Deletes a MigrationConfiguration
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// resourceGroupName - Name of the Resource group within the Azure subscription.
// namespaceName - The namespace name
// configName - The configuration name. Should always be "$default".
// options - MigrationConfigsClientDeleteOptions contains the optional parameters for the MigrationConfigsClient.Delete method.
func (client *MigrationConfigsClient) Delete(ctx context.Context, resourceGroupName string, namespaceName string, configName MigrationConfigurationName, options *MigrationConfigsClientDeleteOptions) (MigrationConfigsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, namespaceName, configName, options)
	if err != nil {
		return MigrationConfigsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MigrationConfigsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return MigrationConfigsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return MigrationConfigsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *MigrationConfigsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, configName MigrationConfigurationName, options *MigrationConfigsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/migrationConfigurations/{configName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if configName == "" {
		return nil, errors.New("parameter configName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configName}", url.PathEscape(string(configName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves Migration Config
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// resourceGroupName - Name of the Resource group within the Azure subscription.
// namespaceName - The namespace name
// configName - The configuration name. Should always be "$default".
// options - MigrationConfigsClientGetOptions contains the optional parameters for the MigrationConfigsClient.Get method.
func (client *MigrationConfigsClient) Get(ctx context.Context, resourceGroupName string, namespaceName string, configName MigrationConfigurationName, options *MigrationConfigsClientGetOptions) (MigrationConfigsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, namespaceName, configName, options)
	if err != nil {
		return MigrationConfigsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MigrationConfigsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MigrationConfigsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *MigrationConfigsClient) getCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, configName MigrationConfigurationName, options *MigrationConfigsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/migrationConfigurations/{configName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if configName == "" {
		return nil, errors.New("parameter configName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configName}", url.PathEscape(string(configName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MigrationConfigsClient) getHandleResponse(resp *http.Response) (MigrationConfigsClientGetResponse, error) {
	result := MigrationConfigsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MigrationConfigProperties); err != nil {
		return MigrationConfigsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all migrationConfigurations
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// resourceGroupName - Name of the Resource group within the Azure subscription.
// namespaceName - The namespace name
// options - MigrationConfigsClientListOptions contains the optional parameters for the MigrationConfigsClient.List method.
func (client *MigrationConfigsClient) NewListPager(resourceGroupName string, namespaceName string, options *MigrationConfigsClientListOptions) *runtime.Pager[MigrationConfigsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[MigrationConfigsClientListResponse]{
		More: func(page MigrationConfigsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MigrationConfigsClientListResponse) (MigrationConfigsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, namespaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return MigrationConfigsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return MigrationConfigsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return MigrationConfigsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *MigrationConfigsClient) listCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, options *MigrationConfigsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/migrationConfigurations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MigrationConfigsClient) listHandleResponse(resp *http.Response) (MigrationConfigsClientListResponse, error) {
	result := MigrationConfigsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MigrationConfigListResult); err != nil {
		return MigrationConfigsClientListResponse{}, err
	}
	return result, nil
}

// Revert - This operation reverts Migration
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// resourceGroupName - Name of the Resource group within the Azure subscription.
// namespaceName - The namespace name
// configName - The configuration name. Should always be "$default".
// options - MigrationConfigsClientRevertOptions contains the optional parameters for the MigrationConfigsClient.Revert method.
func (client *MigrationConfigsClient) Revert(ctx context.Context, resourceGroupName string, namespaceName string, configName MigrationConfigurationName, options *MigrationConfigsClientRevertOptions) (MigrationConfigsClientRevertResponse, error) {
	req, err := client.revertCreateRequest(ctx, resourceGroupName, namespaceName, configName, options)
	if err != nil {
		return MigrationConfigsClientRevertResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MigrationConfigsClientRevertResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MigrationConfigsClientRevertResponse{}, runtime.NewResponseError(resp)
	}
	return MigrationConfigsClientRevertResponse{}, nil
}

// revertCreateRequest creates the Revert request.
func (client *MigrationConfigsClient) revertCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, configName MigrationConfigurationName, options *MigrationConfigsClientRevertOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/migrationConfigurations/{configName}/revert"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if configName == "" {
		return nil, errors.New("parameter configName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configName}", url.PathEscape(string(configName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}