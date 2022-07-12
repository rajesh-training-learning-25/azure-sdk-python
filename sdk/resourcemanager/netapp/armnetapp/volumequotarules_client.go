//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetapp

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

// VolumeQuotaRulesClient contains the methods for the VolumeQuotaRules group.
// Don't use this type directly, use NewVolumeQuotaRulesClient() instead.
type VolumeQuotaRulesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVolumeQuotaRulesClient creates a new instance of VolumeQuotaRulesClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVolumeQuotaRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VolumeQuotaRulesClient, error) {
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
	client := &VolumeQuotaRulesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Create the specified quota rule within the given volume
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// accountName - The name of the NetApp account
// poolName - The name of the capacity pool
// volumeName - The name of the volume
// volumeQuotaRuleName - The name of volume quota rule
// body - Quota rule object supplied in the body of the operation.
// options - VolumeQuotaRulesClientBeginCreateOptions contains the optional parameters for the VolumeQuotaRulesClient.BeginCreate
// method.
func (client *VolumeQuotaRulesClient) BeginCreate(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, volumeQuotaRuleName string, body VolumeQuotaRule, options *VolumeQuotaRulesClientBeginCreateOptions) (*runtime.Poller[VolumeQuotaRulesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, accountName, poolName, volumeName, volumeQuotaRuleName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[VolumeQuotaRulesClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VolumeQuotaRulesClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Create the specified quota rule within the given volume
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
func (client *VolumeQuotaRulesClient) create(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, volumeQuotaRuleName string, body VolumeQuotaRule, options *VolumeQuotaRulesClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, volumeQuotaRuleName, body, options)
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

// createCreateRequest creates the Create request.
func (client *VolumeQuotaRulesClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, volumeQuotaRuleName string, body VolumeQuotaRule, options *VolumeQuotaRulesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/volumeQuotaRules/{volumeQuotaRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	if volumeQuotaRuleName == "" {
		return nil, errors.New("parameter volumeQuotaRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeQuotaRuleName}", url.PathEscape(volumeQuotaRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Delete quota rule
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// accountName - The name of the NetApp account
// poolName - The name of the capacity pool
// volumeName - The name of the volume
// volumeQuotaRuleName - The name of volume quota rule
// options - VolumeQuotaRulesClientBeginDeleteOptions contains the optional parameters for the VolumeQuotaRulesClient.BeginDelete
// method.
func (client *VolumeQuotaRulesClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, volumeQuotaRuleName string, options *VolumeQuotaRulesClientBeginDeleteOptions) (*runtime.Poller[VolumeQuotaRulesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, poolName, volumeName, volumeQuotaRuleName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[VolumeQuotaRulesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VolumeQuotaRulesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete quota rule
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
func (client *VolumeQuotaRulesClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, volumeQuotaRuleName string, options *VolumeQuotaRulesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, volumeQuotaRuleName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VolumeQuotaRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, volumeQuotaRuleName string, options *VolumeQuotaRulesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/volumeQuotaRules/{volumeQuotaRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	if volumeQuotaRuleName == "" {
		return nil, errors.New("parameter volumeQuotaRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeQuotaRuleName}", url.PathEscape(volumeQuotaRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Get details of the specified quota rule
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// accountName - The name of the NetApp account
// poolName - The name of the capacity pool
// volumeName - The name of the volume
// volumeQuotaRuleName - The name of volume quota rule
// options - VolumeQuotaRulesClientGetOptions contains the optional parameters for the VolumeQuotaRulesClient.Get method.
func (client *VolumeQuotaRulesClient) Get(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, volumeQuotaRuleName string, options *VolumeQuotaRulesClientGetOptions) (VolumeQuotaRulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, volumeQuotaRuleName, options)
	if err != nil {
		return VolumeQuotaRulesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VolumeQuotaRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VolumeQuotaRulesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VolumeQuotaRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, volumeQuotaRuleName string, options *VolumeQuotaRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/volumeQuotaRules/{volumeQuotaRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	if volumeQuotaRuleName == "" {
		return nil, errors.New("parameter volumeQuotaRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeQuotaRuleName}", url.PathEscape(volumeQuotaRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VolumeQuotaRulesClient) getHandleResponse(resp *http.Response) (VolumeQuotaRulesClientGetResponse, error) {
	result := VolumeQuotaRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VolumeQuotaRule); err != nil {
		return VolumeQuotaRulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByVolumePager - List all quota rules associated with the volume
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// accountName - The name of the NetApp account
// poolName - The name of the capacity pool
// volumeName - The name of the volume
// options - VolumeQuotaRulesClientListByVolumeOptions contains the optional parameters for the VolumeQuotaRulesClient.ListByVolume
// method.
func (client *VolumeQuotaRulesClient) NewListByVolumePager(resourceGroupName string, accountName string, poolName string, volumeName string, options *VolumeQuotaRulesClientListByVolumeOptions) *runtime.Pager[VolumeQuotaRulesClientListByVolumeResponse] {
	return runtime.NewPager(runtime.PagingHandler[VolumeQuotaRulesClientListByVolumeResponse]{
		More: func(page VolumeQuotaRulesClientListByVolumeResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *VolumeQuotaRulesClientListByVolumeResponse) (VolumeQuotaRulesClientListByVolumeResponse, error) {
			req, err := client.listByVolumeCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, options)
			if err != nil {
				return VolumeQuotaRulesClientListByVolumeResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VolumeQuotaRulesClientListByVolumeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VolumeQuotaRulesClientListByVolumeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByVolumeHandleResponse(resp)
		},
	})
}

// listByVolumeCreateRequest creates the ListByVolume request.
func (client *VolumeQuotaRulesClient) listByVolumeCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, options *VolumeQuotaRulesClientListByVolumeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/volumeQuotaRules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByVolumeHandleResponse handles the ListByVolume response.
func (client *VolumeQuotaRulesClient) listByVolumeHandleResponse(resp *http.Response) (VolumeQuotaRulesClientListByVolumeResponse, error) {
	result := VolumeQuotaRulesClientListByVolumeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VolumeQuotaRulesList); err != nil {
		return VolumeQuotaRulesClientListByVolumeResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Patch a quota rule
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// accountName - The name of the NetApp account
// poolName - The name of the capacity pool
// volumeName - The name of the volume
// volumeQuotaRuleName - The name of volume quota rule
// body - Quota rule object supplied in the body of the operation.
// options - VolumeQuotaRulesClientBeginUpdateOptions contains the optional parameters for the VolumeQuotaRulesClient.BeginUpdate
// method.
func (client *VolumeQuotaRulesClient) BeginUpdate(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, volumeQuotaRuleName string, body VolumeQuotaRulePatch, options *VolumeQuotaRulesClientBeginUpdateOptions) (*runtime.Poller[VolumeQuotaRulesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, accountName, poolName, volumeName, volumeQuotaRuleName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[VolumeQuotaRulesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VolumeQuotaRulesClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Patch a quota rule
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
func (client *VolumeQuotaRulesClient) update(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, volumeQuotaRuleName string, body VolumeQuotaRulePatch, options *VolumeQuotaRulesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, volumeQuotaRuleName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *VolumeQuotaRulesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, volumeQuotaRuleName string, body VolumeQuotaRulePatch, options *VolumeQuotaRulesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/volumeQuotaRules/{volumeQuotaRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	if volumeQuotaRuleName == "" {
		return nil, errors.New("parameter volumeQuotaRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeQuotaRuleName}", url.PathEscape(volumeQuotaRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}
