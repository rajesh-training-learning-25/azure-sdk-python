//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmysqlflexibleservers

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ConfigurationsClient contains the methods for the Configurations group.
// Don't use this type directly, use NewConfigurationsClient() instead.
type ConfigurationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewConfigurationsClient creates a new instance of ConfigurationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConfigurationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ConfigurationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginBatchUpdate - Update a list of configurations in a given server.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - parameters - The parameters for updating a list of server configuration.
//   - options - ConfigurationsClientBeginBatchUpdateOptions contains the optional parameters for the ConfigurationsClient.BeginBatchUpdate
//     method.
func (client *ConfigurationsClient) BeginBatchUpdate(ctx context.Context, resourceGroupName string, serverName string, parameters ConfigurationListForBatchUpdate, options *ConfigurationsClientBeginBatchUpdateOptions) (*runtime.Poller[ConfigurationsClientBatchUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.batchUpdate(ctx, resourceGroupName, serverName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ConfigurationsClientBatchUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ConfigurationsClientBatchUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// BatchUpdate - Update a list of configurations in a given server.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *ConfigurationsClient) batchUpdate(ctx context.Context, resourceGroupName string, serverName string, parameters ConfigurationListForBatchUpdate, options *ConfigurationsClientBeginBatchUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ConfigurationsClient.BeginBatchUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.batchUpdateCreateRequest(ctx, resourceGroupName, serverName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// batchUpdateCreateRequest creates the BatchUpdate request.
func (client *ConfigurationsClient) batchUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, parameters ConfigurationListForBatchUpdate, options *ConfigurationsClientBeginBatchUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/flexibleServers/{serverName}/updateConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginCreateOrUpdate - Updates a configuration of a server.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - configurationName - The name of the server configuration.
//   - parameters - The required parameters for updating a server configuration.
//   - options - ConfigurationsClientBeginCreateOrUpdateOptions contains the optional parameters for the ConfigurationsClient.BeginCreateOrUpdate
//     method.
func (client *ConfigurationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, configurationName string, parameters Configuration, options *ConfigurationsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ConfigurationsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, configurationName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ConfigurationsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ConfigurationsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Updates a configuration of a server.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *ConfigurationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, configurationName string, parameters Configuration, options *ConfigurationsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ConfigurationsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, configurationName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ConfigurationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, configurationName string, parameters Configuration, options *ConfigurationsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/flexibleServers/{serverName}/configurations/{configurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Gets information about a configuration of server.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - configurationName - The name of the server configuration.
//   - options - ConfigurationsClientGetOptions contains the optional parameters for the ConfigurationsClient.Get method.
func (client *ConfigurationsClient) Get(ctx context.Context, resourceGroupName string, serverName string, configurationName string, options *ConfigurationsClientGetOptions) (ConfigurationsClientGetResponse, error) {
	var err error
	const operationName = "ConfigurationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, configurationName, options)
	if err != nil {
		return ConfigurationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConfigurationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConfigurationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, configurationName string, options *ConfigurationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/flexibleServers/{serverName}/configurations/{configurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConfigurationsClient) getHandleResponse(resp *http.Response) (ConfigurationsClientGetResponse, error) {
	result := ConfigurationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Configuration); err != nil {
		return ConfigurationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServerPager - List all the configurations in a given server.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - options - ConfigurationsClientListByServerOptions contains the optional parameters for the ConfigurationsClient.NewListByServerPager
//     method.
func (client *ConfigurationsClient) NewListByServerPager(resourceGroupName string, serverName string, options *ConfigurationsClientListByServerOptions) *runtime.Pager[ConfigurationsClientListByServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[ConfigurationsClientListByServerResponse]{
		More: func(page ConfigurationsClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ConfigurationsClientListByServerResponse) (ConfigurationsClientListByServerResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ConfigurationsClient.NewListByServerPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
			}, nil)
			if err != nil {
				return ConfigurationsClientListByServerResponse{}, err
			}
			return client.listByServerHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *ConfigurationsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ConfigurationsClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/flexibleServers/{serverName}/configurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	if options != nil && options.Keyword != nil {
		reqQP.Set("keyword", *options.Keyword)
	}
	if options != nil && options.Page != nil {
		reqQP.Set("page", strconv.FormatInt(int64(*options.Page), 10))
	}
	if options != nil && options.PageSize != nil {
		reqQP.Set("pageSize", strconv.FormatInt(int64(*options.PageSize), 10))
	}
	if options != nil && options.Tags != nil {
		reqQP.Set("tags", *options.Tags)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *ConfigurationsClient) listByServerHandleResponse(resp *http.Response) (ConfigurationsClientListByServerResponse, error) {
	result := ConfigurationsClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationListResult); err != nil {
		return ConfigurationsClientListByServerResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a configuration of a server.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - configurationName - The name of the server configuration.
//   - parameters - The required parameters for updating a server configuration.
//   - options - ConfigurationsClientBeginUpdateOptions contains the optional parameters for the ConfigurationsClient.BeginUpdate
//     method.
func (client *ConfigurationsClient) BeginUpdate(ctx context.Context, resourceGroupName string, serverName string, configurationName string, parameters Configuration, options *ConfigurationsClientBeginUpdateOptions) (*runtime.Poller[ConfigurationsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, serverName, configurationName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ConfigurationsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ConfigurationsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Updates a configuration of a server.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *ConfigurationsClient) update(ctx context.Context, resourceGroupName string, serverName string, configurationName string, parameters Configuration, options *ConfigurationsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ConfigurationsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serverName, configurationName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *ConfigurationsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, configurationName string, parameters Configuration, options *ConfigurationsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/flexibleServers/{serverName}/configurations/{configurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}
