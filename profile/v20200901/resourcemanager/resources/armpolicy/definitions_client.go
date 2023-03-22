//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armpolicy

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/profile/v20200901/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DefinitionsClient contains the methods for the PolicyDefinitions group.
// Don't use this type directly, use NewDefinitionsClient() instead.
type DefinitionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDefinitionsClient creates a new instance of DefinitionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDefinitionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DefinitionsClient, error) {
	cl, err := arm.NewClient(internal.ModuleName+"/armpolicy.DefinitionsClient", internal.ModuleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DefinitionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a policy definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-12-01
//   - policyDefinitionName - The name of the policy definition to create.
//   - parameters - The policy definition properties.
//   - options - DefinitionsClientCreateOrUpdateOptions contains the optional parameters for the DefinitionsClient.CreateOrUpdate
//     method.
func (client *DefinitionsClient) CreateOrUpdate(ctx context.Context, policyDefinitionName string, parameters Definition, options *DefinitionsClientCreateOrUpdateOptions) (DefinitionsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, policyDefinitionName, parameters, options)
	if err != nil {
		return DefinitionsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DefinitionsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return DefinitionsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DefinitionsClient) createOrUpdateCreateRequest(ctx context.Context, policyDefinitionName string, parameters Definition, options *DefinitionsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}"
	if policyDefinitionName == "" {
		return nil, errors.New("parameter policyDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyDefinitionName}", url.PathEscape(policyDefinitionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json, text/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DefinitionsClient) createOrUpdateHandleResponse(resp *http.Response) (DefinitionsClientCreateOrUpdateResponse, error) {
	result := DefinitionsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Definition); err != nil {
		return DefinitionsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// CreateOrUpdateAtManagementGroup - Creates or updates a policy definition at management group level.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-12-01
//   - policyDefinitionName - The name of the policy definition to create.
//   - managementGroupID - The ID of the management group.
//   - parameters - The policy definition properties.
//   - options - DefinitionsClientCreateOrUpdateAtManagementGroupOptions contains the optional parameters for the DefinitionsClient.CreateOrUpdateAtManagementGroup
//     method.
func (client *DefinitionsClient) CreateOrUpdateAtManagementGroup(ctx context.Context, policyDefinitionName string, managementGroupID string, parameters Definition, options *DefinitionsClientCreateOrUpdateAtManagementGroupOptions) (DefinitionsClientCreateOrUpdateAtManagementGroupResponse, error) {
	req, err := client.createOrUpdateAtManagementGroupCreateRequest(ctx, policyDefinitionName, managementGroupID, parameters, options)
	if err != nil {
		return DefinitionsClientCreateOrUpdateAtManagementGroupResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DefinitionsClientCreateOrUpdateAtManagementGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return DefinitionsClientCreateOrUpdateAtManagementGroupResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateAtManagementGroupHandleResponse(resp)
}

// createOrUpdateAtManagementGroupCreateRequest creates the CreateOrUpdateAtManagementGroup request.
func (client *DefinitionsClient) createOrUpdateAtManagementGroupCreateRequest(ctx context.Context, policyDefinitionName string, managementGroupID string, parameters Definition, options *DefinitionsClientCreateOrUpdateAtManagementGroupOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}"
	if policyDefinitionName == "" {
		return nil, errors.New("parameter policyDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyDefinitionName}", url.PathEscape(policyDefinitionName))
	if managementGroupID == "" {
		return nil, errors.New("parameter managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(managementGroupID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json, text/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateAtManagementGroupHandleResponse handles the CreateOrUpdateAtManagementGroup response.
func (client *DefinitionsClient) createOrUpdateAtManagementGroupHandleResponse(resp *http.Response) (DefinitionsClientCreateOrUpdateAtManagementGroupResponse, error) {
	result := DefinitionsClientCreateOrUpdateAtManagementGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Definition); err != nil {
		return DefinitionsClientCreateOrUpdateAtManagementGroupResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a policy definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-12-01
//   - policyDefinitionName - The name of the policy definition to delete.
//   - options - DefinitionsClientDeleteOptions contains the optional parameters for the DefinitionsClient.Delete method.
func (client *DefinitionsClient) Delete(ctx context.Context, policyDefinitionName string, options *DefinitionsClientDeleteOptions) (DefinitionsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, policyDefinitionName, options)
	if err != nil {
		return DefinitionsClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DefinitionsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return DefinitionsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return DefinitionsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DefinitionsClient) deleteCreateRequest(ctx context.Context, policyDefinitionName string, options *DefinitionsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}"
	if policyDefinitionName == "" {
		return nil, errors.New("parameter policyDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyDefinitionName}", url.PathEscape(policyDefinitionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// DeleteAtManagementGroup - Deletes a policy definition at management group level.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-12-01
//   - policyDefinitionName - The name of the policy definition to delete.
//   - managementGroupID - The ID of the management group.
//   - options - DefinitionsClientDeleteAtManagementGroupOptions contains the optional parameters for the DefinitionsClient.DeleteAtManagementGroup
//     method.
func (client *DefinitionsClient) DeleteAtManagementGroup(ctx context.Context, policyDefinitionName string, managementGroupID string, options *DefinitionsClientDeleteAtManagementGroupOptions) (DefinitionsClientDeleteAtManagementGroupResponse, error) {
	req, err := client.deleteAtManagementGroupCreateRequest(ctx, policyDefinitionName, managementGroupID, options)
	if err != nil {
		return DefinitionsClientDeleteAtManagementGroupResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DefinitionsClientDeleteAtManagementGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return DefinitionsClientDeleteAtManagementGroupResponse{}, runtime.NewResponseError(resp)
	}
	return DefinitionsClientDeleteAtManagementGroupResponse{}, nil
}

// deleteAtManagementGroupCreateRequest creates the DeleteAtManagementGroup request.
func (client *DefinitionsClient) deleteAtManagementGroupCreateRequest(ctx context.Context, policyDefinitionName string, managementGroupID string, options *DefinitionsClientDeleteAtManagementGroupOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}"
	if policyDefinitionName == "" {
		return nil, errors.New("parameter policyDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyDefinitionName}", url.PathEscape(policyDefinitionName))
	if managementGroupID == "" {
		return nil, errors.New("parameter managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(managementGroupID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets the policy definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-12-01
//   - policyDefinitionName - The name of the policy definition to get.
//   - options - DefinitionsClientGetOptions contains the optional parameters for the DefinitionsClient.Get method.
func (client *DefinitionsClient) Get(ctx context.Context, policyDefinitionName string, options *DefinitionsClientGetOptions) (DefinitionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, policyDefinitionName, options)
	if err != nil {
		return DefinitionsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DefinitionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DefinitionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DefinitionsClient) getCreateRequest(ctx context.Context, policyDefinitionName string, options *DefinitionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}"
	if policyDefinitionName == "" {
		return nil, errors.New("parameter policyDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyDefinitionName}", url.PathEscape(policyDefinitionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json, text/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DefinitionsClient) getHandleResponse(resp *http.Response) (DefinitionsClientGetResponse, error) {
	result := DefinitionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Definition); err != nil {
		return DefinitionsClientGetResponse{}, err
	}
	return result, nil
}

// GetAtManagementGroup - Gets the policy definition at management group level.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-12-01
//   - policyDefinitionName - The name of the policy definition to get.
//   - managementGroupID - The ID of the management group.
//   - options - DefinitionsClientGetAtManagementGroupOptions contains the optional parameters for the DefinitionsClient.GetAtManagementGroup
//     method.
func (client *DefinitionsClient) GetAtManagementGroup(ctx context.Context, policyDefinitionName string, managementGroupID string, options *DefinitionsClientGetAtManagementGroupOptions) (DefinitionsClientGetAtManagementGroupResponse, error) {
	req, err := client.getAtManagementGroupCreateRequest(ctx, policyDefinitionName, managementGroupID, options)
	if err != nil {
		return DefinitionsClientGetAtManagementGroupResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DefinitionsClientGetAtManagementGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DefinitionsClientGetAtManagementGroupResponse{}, runtime.NewResponseError(resp)
	}
	return client.getAtManagementGroupHandleResponse(resp)
}

// getAtManagementGroupCreateRequest creates the GetAtManagementGroup request.
func (client *DefinitionsClient) getAtManagementGroupCreateRequest(ctx context.Context, policyDefinitionName string, managementGroupID string, options *DefinitionsClientGetAtManagementGroupOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}"
	if policyDefinitionName == "" {
		return nil, errors.New("parameter policyDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyDefinitionName}", url.PathEscape(policyDefinitionName))
	if managementGroupID == "" {
		return nil, errors.New("parameter managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(managementGroupID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json, text/json"}
	return req, nil
}

// getAtManagementGroupHandleResponse handles the GetAtManagementGroup response.
func (client *DefinitionsClient) getAtManagementGroupHandleResponse(resp *http.Response) (DefinitionsClientGetAtManagementGroupResponse, error) {
	result := DefinitionsClientGetAtManagementGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Definition); err != nil {
		return DefinitionsClientGetAtManagementGroupResponse{}, err
	}
	return result, nil
}

// GetBuiltIn - Gets the built in policy definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-12-01
//   - policyDefinitionName - The name of the built in policy definition to get.
//   - options - DefinitionsClientGetBuiltInOptions contains the optional parameters for the DefinitionsClient.GetBuiltIn method.
func (client *DefinitionsClient) GetBuiltIn(ctx context.Context, policyDefinitionName string, options *DefinitionsClientGetBuiltInOptions) (DefinitionsClientGetBuiltInResponse, error) {
	req, err := client.getBuiltInCreateRequest(ctx, policyDefinitionName, options)
	if err != nil {
		return DefinitionsClientGetBuiltInResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DefinitionsClientGetBuiltInResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DefinitionsClientGetBuiltInResponse{}, runtime.NewResponseError(resp)
	}
	return client.getBuiltInHandleResponse(resp)
}

// getBuiltInCreateRequest creates the GetBuiltIn request.
func (client *DefinitionsClient) getBuiltInCreateRequest(ctx context.Context, policyDefinitionName string, options *DefinitionsClientGetBuiltInOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}"
	if policyDefinitionName == "" {
		return nil, errors.New("parameter policyDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyDefinitionName}", url.PathEscape(policyDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json, text/json"}
	return req, nil
}

// getBuiltInHandleResponse handles the GetBuiltIn response.
func (client *DefinitionsClient) getBuiltInHandleResponse(resp *http.Response) (DefinitionsClientGetBuiltInResponse, error) {
	result := DefinitionsClientGetBuiltInResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Definition); err != nil {
		return DefinitionsClientGetBuiltInResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all the policy definitions for a subscription.
//
// Generated from API version 2016-12-01
//   - options - DefinitionsClientListOptions contains the optional parameters for the DefinitionsClient.NewListPager method.
func (client *DefinitionsClient) NewListPager(options *DefinitionsClientListOptions) *runtime.Pager[DefinitionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DefinitionsClientListResponse]{
		More: func(page DefinitionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DefinitionsClientListResponse) (DefinitionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DefinitionsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DefinitionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DefinitionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *DefinitionsClient) listCreateRequest(ctx context.Context, options *DefinitionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json, text/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DefinitionsClient) listHandleResponse(resp *http.Response) (DefinitionsClientListResponse, error) {
	result := DefinitionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DefinitionListResult); err != nil {
		return DefinitionsClientListResponse{}, err
	}
	return result, nil
}

// NewListBuiltInPager - Gets all the built in policy definitions.
//
// Generated from API version 2016-12-01
//   - options - DefinitionsClientListBuiltInOptions contains the optional parameters for the DefinitionsClient.NewListBuiltInPager
//     method.
func (client *DefinitionsClient) NewListBuiltInPager(options *DefinitionsClientListBuiltInOptions) *runtime.Pager[DefinitionsClientListBuiltInResponse] {
	return runtime.NewPager(runtime.PagingHandler[DefinitionsClientListBuiltInResponse]{
		More: func(page DefinitionsClientListBuiltInResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DefinitionsClientListBuiltInResponse) (DefinitionsClientListBuiltInResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBuiltInCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DefinitionsClientListBuiltInResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DefinitionsClientListBuiltInResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DefinitionsClientListBuiltInResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBuiltInHandleResponse(resp)
		},
	})
}

// listBuiltInCreateRequest creates the ListBuiltIn request.
func (client *DefinitionsClient) listBuiltInCreateRequest(ctx context.Context, options *DefinitionsClientListBuiltInOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Authorization/policyDefinitions"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json, text/json"}
	return req, nil
}

// listBuiltInHandleResponse handles the ListBuiltIn response.
func (client *DefinitionsClient) listBuiltInHandleResponse(resp *http.Response) (DefinitionsClientListBuiltInResponse, error) {
	result := DefinitionsClientListBuiltInResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DefinitionListResult); err != nil {
		return DefinitionsClientListBuiltInResponse{}, err
	}
	return result, nil
}

// NewListByManagementGroupPager - Gets all the policy definitions for a subscription at management group level.
//
// Generated from API version 2016-12-01
//   - managementGroupID - The ID of the management group.
//   - options - DefinitionsClientListByManagementGroupOptions contains the optional parameters for the DefinitionsClient.NewListByManagementGroupPager
//     method.
func (client *DefinitionsClient) NewListByManagementGroupPager(managementGroupID string, options *DefinitionsClientListByManagementGroupOptions) *runtime.Pager[DefinitionsClientListByManagementGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[DefinitionsClientListByManagementGroupResponse]{
		More: func(page DefinitionsClientListByManagementGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DefinitionsClientListByManagementGroupResponse) (DefinitionsClientListByManagementGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByManagementGroupCreateRequest(ctx, managementGroupID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DefinitionsClientListByManagementGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DefinitionsClientListByManagementGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DefinitionsClientListByManagementGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByManagementGroupHandleResponse(resp)
		},
	})
}

// listByManagementGroupCreateRequest creates the ListByManagementGroup request.
func (client *DefinitionsClient) listByManagementGroupCreateRequest(ctx context.Context, managementGroupID string, options *DefinitionsClientListByManagementGroupOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions"
	if managementGroupID == "" {
		return nil, errors.New("parameter managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(managementGroupID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json, text/json"}
	return req, nil
}

// listByManagementGroupHandleResponse handles the ListByManagementGroup response.
func (client *DefinitionsClient) listByManagementGroupHandleResponse(resp *http.Response) (DefinitionsClientListByManagementGroupResponse, error) {
	result := DefinitionsClientListByManagementGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DefinitionListResult); err != nil {
		return DefinitionsClientListByManagementGroupResponse{}, err
	}
	return result, nil
}
