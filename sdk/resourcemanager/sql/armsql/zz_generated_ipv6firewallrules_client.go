//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// IPv6FirewallRulesClient contains the methods for the IPv6FirewallRules group.
// Don't use this type directly, use NewIPv6FirewallRulesClient() instead.
type IPv6FirewallRulesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewIPv6FirewallRulesClient creates a new instance of IPv6FirewallRulesClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewIPv6FirewallRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IPv6FirewallRulesClient, error) {
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
	client := &IPv6FirewallRulesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates an IPv6 firewall rule.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// firewallRuleName - The name of the firewall rule.
// parameters - The required parameters for creating or updating an IPv6 firewall rule.
// options - IPv6FirewallRulesClientCreateOrUpdateOptions contains the optional parameters for the IPv6FirewallRulesClient.CreateOrUpdate
// method.
func (client *IPv6FirewallRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, parameters IPv6FirewallRule, options *IPv6FirewallRulesClientCreateOrUpdateOptions) (IPv6FirewallRulesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, firewallRuleName, parameters, options)
	if err != nil {
		return IPv6FirewallRulesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IPv6FirewallRulesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return IPv6FirewallRulesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IPv6FirewallRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, parameters IPv6FirewallRule, options *IPv6FirewallRulesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/ipv6FirewallRules/{firewallRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if firewallRuleName == "" {
		return nil, errors.New("parameter firewallRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallRuleName}", url.PathEscape(firewallRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *IPv6FirewallRulesClient) createOrUpdateHandleResponse(resp *http.Response) (IPv6FirewallRulesClientCreateOrUpdateResponse, error) {
	result := IPv6FirewallRulesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPv6FirewallRule); err != nil {
		return IPv6FirewallRulesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an IPv6 firewall rule.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// firewallRuleName - The name of the firewall rule.
// options - IPv6FirewallRulesClientDeleteOptions contains the optional parameters for the IPv6FirewallRulesClient.Delete
// method.
func (client *IPv6FirewallRulesClient) Delete(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, options *IPv6FirewallRulesClientDeleteOptions) (IPv6FirewallRulesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, firewallRuleName, options)
	if err != nil {
		return IPv6FirewallRulesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IPv6FirewallRulesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return IPv6FirewallRulesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return IPv6FirewallRulesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IPv6FirewallRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, options *IPv6FirewallRulesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/ipv6FirewallRules/{firewallRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if firewallRuleName == "" {
		return nil, errors.New("parameter firewallRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallRuleName}", url.PathEscape(firewallRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets an IPv6 firewall rule.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// firewallRuleName - The name of the firewall rule.
// options - IPv6FirewallRulesClientGetOptions contains the optional parameters for the IPv6FirewallRulesClient.Get method.
func (client *IPv6FirewallRulesClient) Get(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, options *IPv6FirewallRulesClientGetOptions) (IPv6FirewallRulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, firewallRuleName, options)
	if err != nil {
		return IPv6FirewallRulesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IPv6FirewallRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IPv6FirewallRulesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IPv6FirewallRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, options *IPv6FirewallRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/ipv6FirewallRules/{firewallRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if firewallRuleName == "" {
		return nil, errors.New("parameter firewallRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallRuleName}", url.PathEscape(firewallRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IPv6FirewallRulesClient) getHandleResponse(resp *http.Response) (IPv6FirewallRulesClientGetResponse, error) {
	result := IPv6FirewallRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPv6FirewallRule); err != nil {
		return IPv6FirewallRulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServerPager - Gets a list of IPv6 firewall rules.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// options - IPv6FirewallRulesClientListByServerOptions contains the optional parameters for the IPv6FirewallRulesClient.ListByServer
// method.
func (client *IPv6FirewallRulesClient) NewListByServerPager(resourceGroupName string, serverName string, options *IPv6FirewallRulesClientListByServerOptions) *runtime.Pager[IPv6FirewallRulesClientListByServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[IPv6FirewallRulesClientListByServerResponse]{
		More: func(page IPv6FirewallRulesClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IPv6FirewallRulesClientListByServerResponse) (IPv6FirewallRulesClientListByServerResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return IPv6FirewallRulesClientListByServerResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return IPv6FirewallRulesClientListByServerResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return IPv6FirewallRulesClientListByServerResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServerHandleResponse(resp)
		},
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *IPv6FirewallRulesClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *IPv6FirewallRulesClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/ipv6FirewallRules"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *IPv6FirewallRulesClient) listByServerHandleResponse(resp *http.Response) (IPv6FirewallRulesClientListByServerResponse, error) {
	result := IPv6FirewallRulesClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPv6FirewallRuleListResult); err != nil {
		return IPv6FirewallRulesClientListByServerResponse{}, err
	}
	return result, nil
}