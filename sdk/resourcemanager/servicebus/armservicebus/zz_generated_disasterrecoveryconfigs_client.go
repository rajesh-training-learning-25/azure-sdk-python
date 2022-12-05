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

// DisasterRecoveryConfigsClient contains the methods for the DisasterRecoveryConfigs group.
// Don't use this type directly, use NewDisasterRecoveryConfigsClient() instead.
type DisasterRecoveryConfigsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDisasterRecoveryConfigsClient creates a new instance of DisasterRecoveryConfigsClient with the specified values.
// subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDisasterRecoveryConfigsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DisasterRecoveryConfigsClient, error) {
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
	client := &DisasterRecoveryConfigsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BreakPairing - This operation disables the Disaster Recovery and stops replicating changes from primary to secondary namespaces
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// resourceGroupName - Name of the Resource group within the Azure subscription.
// namespaceName - The namespace name
// alias - The Disaster Recovery configuration name
// options - DisasterRecoveryConfigsClientBreakPairingOptions contains the optional parameters for the DisasterRecoveryConfigsClient.BreakPairing
// method.
func (client *DisasterRecoveryConfigsClient) BreakPairing(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsClientBreakPairingOptions) (DisasterRecoveryConfigsClientBreakPairingResponse, error) {
	req, err := client.breakPairingCreateRequest(ctx, resourceGroupName, namespaceName, alias, options)
	if err != nil {
		return DisasterRecoveryConfigsClientBreakPairingResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DisasterRecoveryConfigsClientBreakPairingResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DisasterRecoveryConfigsClientBreakPairingResponse{}, runtime.NewResponseError(resp)
	}
	return DisasterRecoveryConfigsClientBreakPairingResponse{}, nil
}

// breakPairingCreateRequest creates the BreakPairing request.
func (client *DisasterRecoveryConfigsClient) breakPairingCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsClientBreakPairingOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}/breakPairing"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if alias == "" {
		return nil, errors.New("parameter alias cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alias}", url.PathEscape(alias))
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

// CheckNameAvailability - Check the give namespace name availability.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// resourceGroupName - Name of the Resource group within the Azure subscription.
// namespaceName - The namespace name
// parameters - Parameters to check availability of the given namespace name
// options - DisasterRecoveryConfigsClientCheckNameAvailabilityOptions contains the optional parameters for the DisasterRecoveryConfigsClient.CheckNameAvailability
// method.
func (client *DisasterRecoveryConfigsClient) CheckNameAvailability(ctx context.Context, resourceGroupName string, namespaceName string, parameters CheckNameAvailability, options *DisasterRecoveryConfigsClientCheckNameAvailabilityOptions) (DisasterRecoveryConfigsClientCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, resourceGroupName, namespaceName, parameters, options)
	if err != nil {
		return DisasterRecoveryConfigsClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DisasterRecoveryConfigsClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DisasterRecoveryConfigsClientCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *DisasterRecoveryConfigsClient) checkNameAvailabilityCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, parameters CheckNameAvailability, options *DisasterRecoveryConfigsClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs/CheckNameAvailability"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *DisasterRecoveryConfigsClient) checkNameAvailabilityHandleResponse(resp *http.Response) (DisasterRecoveryConfigsClientCheckNameAvailabilityResponse, error) {
	result := DisasterRecoveryConfigsClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResult); err != nil {
		return DisasterRecoveryConfigsClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a new Alias(Disaster Recovery configuration)
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// resourceGroupName - Name of the Resource group within the Azure subscription.
// namespaceName - The namespace name
// alias - The Disaster Recovery configuration name
// parameters - Parameters required to create an Alias(Disaster Recovery configuration)
// options - DisasterRecoveryConfigsClientCreateOrUpdateOptions contains the optional parameters for the DisasterRecoveryConfigsClient.CreateOrUpdate
// method.
func (client *DisasterRecoveryConfigsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, alias string, parameters ArmDisasterRecovery, options *DisasterRecoveryConfigsClientCreateOrUpdateOptions) (DisasterRecoveryConfigsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, namespaceName, alias, parameters, options)
	if err != nil {
		return DisasterRecoveryConfigsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DisasterRecoveryConfigsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return DisasterRecoveryConfigsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DisasterRecoveryConfigsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, alias string, parameters ArmDisasterRecovery, options *DisasterRecoveryConfigsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if alias == "" {
		return nil, errors.New("parameter alias cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alias}", url.PathEscape(alias))
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

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DisasterRecoveryConfigsClient) createOrUpdateHandleResponse(resp *http.Response) (DisasterRecoveryConfigsClientCreateOrUpdateResponse, error) {
	result := DisasterRecoveryConfigsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArmDisasterRecovery); err != nil {
		return DisasterRecoveryConfigsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an Alias(Disaster Recovery configuration)
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// resourceGroupName - Name of the Resource group within the Azure subscription.
// namespaceName - The namespace name
// alias - The Disaster Recovery configuration name
// options - DisasterRecoveryConfigsClientDeleteOptions contains the optional parameters for the DisasterRecoveryConfigsClient.Delete
// method.
func (client *DisasterRecoveryConfigsClient) Delete(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsClientDeleteOptions) (DisasterRecoveryConfigsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, namespaceName, alias, options)
	if err != nil {
		return DisasterRecoveryConfigsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DisasterRecoveryConfigsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return DisasterRecoveryConfigsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return DisasterRecoveryConfigsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DisasterRecoveryConfigsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if alias == "" {
		return nil, errors.New("parameter alias cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alias}", url.PathEscape(alias))
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

// FailOver - Invokes GEO DR failover and reconfigure the alias to point to the secondary namespace
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// resourceGroupName - Name of the Resource group within the Azure subscription.
// namespaceName - The namespace name
// alias - The Disaster Recovery configuration name
// options - DisasterRecoveryConfigsClientFailOverOptions contains the optional parameters for the DisasterRecoveryConfigsClient.FailOver
// method.
func (client *DisasterRecoveryConfigsClient) FailOver(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsClientFailOverOptions) (DisasterRecoveryConfigsClientFailOverResponse, error) {
	req, err := client.failOverCreateRequest(ctx, resourceGroupName, namespaceName, alias, options)
	if err != nil {
		return DisasterRecoveryConfigsClientFailOverResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DisasterRecoveryConfigsClientFailOverResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DisasterRecoveryConfigsClientFailOverResponse{}, runtime.NewResponseError(resp)
	}
	return DisasterRecoveryConfigsClientFailOverResponse{}, nil
}

// failOverCreateRequest creates the FailOver request.
func (client *DisasterRecoveryConfigsClient) failOverCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsClientFailOverOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}/failover"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if alias == "" {
		return nil, errors.New("parameter alias cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alias}", url.PathEscape(alias))
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
	if options != nil && options.Parameters != nil {
		return req, runtime.MarshalAsJSON(req, *options.Parameters)
	}
	return req, nil
}

// Get - Retrieves Alias(Disaster Recovery configuration) for primary or secondary namespace
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// resourceGroupName - Name of the Resource group within the Azure subscription.
// namespaceName - The namespace name
// alias - The Disaster Recovery configuration name
// options - DisasterRecoveryConfigsClientGetOptions contains the optional parameters for the DisasterRecoveryConfigsClient.Get
// method.
func (client *DisasterRecoveryConfigsClient) Get(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsClientGetOptions) (DisasterRecoveryConfigsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, namespaceName, alias, options)
	if err != nil {
		return DisasterRecoveryConfigsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DisasterRecoveryConfigsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DisasterRecoveryConfigsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DisasterRecoveryConfigsClient) getCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if alias == "" {
		return nil, errors.New("parameter alias cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alias}", url.PathEscape(alias))
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
func (client *DisasterRecoveryConfigsClient) getHandleResponse(resp *http.Response) (DisasterRecoveryConfigsClientGetResponse, error) {
	result := DisasterRecoveryConfigsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArmDisasterRecovery); err != nil {
		return DisasterRecoveryConfigsClientGetResponse{}, err
	}
	return result, nil
}

// GetAuthorizationRule - Gets an authorization rule for a namespace by rule name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// resourceGroupName - Name of the Resource group within the Azure subscription.
// namespaceName - The namespace name
// alias - The Disaster Recovery configuration name
// authorizationRuleName - The authorization rule name.
// options - DisasterRecoveryConfigsClientGetAuthorizationRuleOptions contains the optional parameters for the DisasterRecoveryConfigsClient.GetAuthorizationRule
// method.
func (client *DisasterRecoveryConfigsClient) GetAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, alias string, authorizationRuleName string, options *DisasterRecoveryConfigsClientGetAuthorizationRuleOptions) (DisasterRecoveryConfigsClientGetAuthorizationRuleResponse, error) {
	req, err := client.getAuthorizationRuleCreateRequest(ctx, resourceGroupName, namespaceName, alias, authorizationRuleName, options)
	if err != nil {
		return DisasterRecoveryConfigsClientGetAuthorizationRuleResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DisasterRecoveryConfigsClientGetAuthorizationRuleResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DisasterRecoveryConfigsClientGetAuthorizationRuleResponse{}, runtime.NewResponseError(resp)
	}
	return client.getAuthorizationRuleHandleResponse(resp)
}

// getAuthorizationRuleCreateRequest creates the GetAuthorizationRule request.
func (client *DisasterRecoveryConfigsClient) getAuthorizationRuleCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, alias string, authorizationRuleName string, options *DisasterRecoveryConfigsClientGetAuthorizationRuleOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}/authorizationRules/{authorizationRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if alias == "" {
		return nil, errors.New("parameter alias cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alias}", url.PathEscape(alias))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
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

// getAuthorizationRuleHandleResponse handles the GetAuthorizationRule response.
func (client *DisasterRecoveryConfigsClient) getAuthorizationRuleHandleResponse(resp *http.Response) (DisasterRecoveryConfigsClientGetAuthorizationRuleResponse, error) {
	result := DisasterRecoveryConfigsClientGetAuthorizationRuleResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SBAuthorizationRule); err != nil {
		return DisasterRecoveryConfigsClientGetAuthorizationRuleResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all Alias(Disaster Recovery configurations)
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// resourceGroupName - Name of the Resource group within the Azure subscription.
// namespaceName - The namespace name
// options - DisasterRecoveryConfigsClientListOptions contains the optional parameters for the DisasterRecoveryConfigsClient.List
// method.
func (client *DisasterRecoveryConfigsClient) NewListPager(resourceGroupName string, namespaceName string, options *DisasterRecoveryConfigsClientListOptions) *runtime.Pager[DisasterRecoveryConfigsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DisasterRecoveryConfigsClientListResponse]{
		More: func(page DisasterRecoveryConfigsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DisasterRecoveryConfigsClientListResponse) (DisasterRecoveryConfigsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, namespaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DisasterRecoveryConfigsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DisasterRecoveryConfigsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DisasterRecoveryConfigsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *DisasterRecoveryConfigsClient) listCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, options *DisasterRecoveryConfigsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs"
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
func (client *DisasterRecoveryConfigsClient) listHandleResponse(resp *http.Response) (DisasterRecoveryConfigsClientListResponse, error) {
	result := DisasterRecoveryConfigsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArmDisasterRecoveryListResult); err != nil {
		return DisasterRecoveryConfigsClientListResponse{}, err
	}
	return result, nil
}

// NewListAuthorizationRulesPager - Gets the authorization rules for a namespace.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// resourceGroupName - Name of the Resource group within the Azure subscription.
// namespaceName - The namespace name
// alias - The Disaster Recovery configuration name
// options - DisasterRecoveryConfigsClientListAuthorizationRulesOptions contains the optional parameters for the DisasterRecoveryConfigsClient.ListAuthorizationRules
// method.
func (client *DisasterRecoveryConfigsClient) NewListAuthorizationRulesPager(resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsClientListAuthorizationRulesOptions) *runtime.Pager[DisasterRecoveryConfigsClientListAuthorizationRulesResponse] {
	return runtime.NewPager(runtime.PagingHandler[DisasterRecoveryConfigsClientListAuthorizationRulesResponse]{
		More: func(page DisasterRecoveryConfigsClientListAuthorizationRulesResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DisasterRecoveryConfigsClientListAuthorizationRulesResponse) (DisasterRecoveryConfigsClientListAuthorizationRulesResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAuthorizationRulesCreateRequest(ctx, resourceGroupName, namespaceName, alias, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DisasterRecoveryConfigsClientListAuthorizationRulesResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DisasterRecoveryConfigsClientListAuthorizationRulesResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DisasterRecoveryConfigsClientListAuthorizationRulesResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAuthorizationRulesHandleResponse(resp)
		},
	})
}

// listAuthorizationRulesCreateRequest creates the ListAuthorizationRules request.
func (client *DisasterRecoveryConfigsClient) listAuthorizationRulesCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsClientListAuthorizationRulesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}/authorizationRules"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if alias == "" {
		return nil, errors.New("parameter alias cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alias}", url.PathEscape(alias))
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

// listAuthorizationRulesHandleResponse handles the ListAuthorizationRules response.
func (client *DisasterRecoveryConfigsClient) listAuthorizationRulesHandleResponse(resp *http.Response) (DisasterRecoveryConfigsClientListAuthorizationRulesResponse, error) {
	result := DisasterRecoveryConfigsClientListAuthorizationRulesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SBAuthorizationRuleListResult); err != nil {
		return DisasterRecoveryConfigsClientListAuthorizationRulesResponse{}, err
	}
	return result, nil
}

// ListKeys - Gets the primary and secondary connection strings for the namespace.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// resourceGroupName - Name of the Resource group within the Azure subscription.
// namespaceName - The namespace name
// alias - The Disaster Recovery configuration name
// authorizationRuleName - The authorization rule name.
// options - DisasterRecoveryConfigsClientListKeysOptions contains the optional parameters for the DisasterRecoveryConfigsClient.ListKeys
// method.
func (client *DisasterRecoveryConfigsClient) ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, alias string, authorizationRuleName string, options *DisasterRecoveryConfigsClientListKeysOptions) (DisasterRecoveryConfigsClientListKeysResponse, error) {
	req, err := client.listKeysCreateRequest(ctx, resourceGroupName, namespaceName, alias, authorizationRuleName, options)
	if err != nil {
		return DisasterRecoveryConfigsClientListKeysResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DisasterRecoveryConfigsClientListKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DisasterRecoveryConfigsClientListKeysResponse{}, runtime.NewResponseError(resp)
	}
	return client.listKeysHandleResponse(resp)
}

// listKeysCreateRequest creates the ListKeys request.
func (client *DisasterRecoveryConfigsClient) listKeysCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, alias string, authorizationRuleName string, options *DisasterRecoveryConfigsClientListKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}/authorizationRules/{authorizationRuleName}/listKeys"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if alias == "" {
		return nil, errors.New("parameter alias cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alias}", url.PathEscape(alias))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
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

// listKeysHandleResponse handles the ListKeys response.
func (client *DisasterRecoveryConfigsClient) listKeysHandleResponse(resp *http.Response) (DisasterRecoveryConfigsClientListKeysResponse, error) {
	result := DisasterRecoveryConfigsClientListKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessKeys); err != nil {
		return DisasterRecoveryConfigsClientListKeysResponse{}, err
	}
	return result, nil
}