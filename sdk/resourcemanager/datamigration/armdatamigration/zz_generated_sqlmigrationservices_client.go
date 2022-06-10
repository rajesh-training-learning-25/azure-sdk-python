//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatamigration

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

// SQLMigrationServicesClient contains the methods for the SQLMigrationServices group.
// Don't use this type directly, use NewSQLMigrationServicesClient() instead.
type SQLMigrationServicesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSQLMigrationServicesClient creates a new instance of SQLMigrationServicesClient with the specified values.
// subscriptionID - Subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSQLMigrationServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SQLMigrationServicesClient, error) {
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
	client := &SQLMigrationServicesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or Update Database Migration Service.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-30-preview
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// sqlMigrationServiceName - Name of the SQL Migration Service.
// parameters - Details of SqlMigrationService resource.
// options - SQLMigrationServicesClientBeginCreateOrUpdateOptions contains the optional parameters for the SQLMigrationServicesClient.BeginCreateOrUpdate
// method.
func (client *SQLMigrationServicesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, sqlMigrationServiceName string, parameters SQLMigrationService, options *SQLMigrationServicesClientBeginCreateOrUpdateOptions) (*runtime.Poller[SQLMigrationServicesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, sqlMigrationServiceName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[SQLMigrationServicesClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[SQLMigrationServicesClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or Update Database Migration Service.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-30-preview
func (client *SQLMigrationServicesClient) createOrUpdate(ctx context.Context, resourceGroupName string, sqlMigrationServiceName string, parameters SQLMigrationService, options *SQLMigrationServicesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, sqlMigrationServiceName, parameters, options)
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
func (client *SQLMigrationServicesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, sqlMigrationServiceName string, parameters SQLMigrationService, options *SQLMigrationServicesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataMigration/sqlMigrationServices/{sqlMigrationServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlMigrationServiceName == "" {
		return nil, errors.New("parameter sqlMigrationServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlMigrationServiceName}", url.PathEscape(sqlMigrationServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Delete Database Migration Service.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-30-preview
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// sqlMigrationServiceName - Name of the SQL Migration Service.
// options - SQLMigrationServicesClientBeginDeleteOptions contains the optional parameters for the SQLMigrationServicesClient.BeginDelete
// method.
func (client *SQLMigrationServicesClient) BeginDelete(ctx context.Context, resourceGroupName string, sqlMigrationServiceName string, options *SQLMigrationServicesClientBeginDeleteOptions) (*runtime.Poller[SQLMigrationServicesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, sqlMigrationServiceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[SQLMigrationServicesClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[SQLMigrationServicesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete Database Migration Service.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-30-preview
func (client *SQLMigrationServicesClient) deleteOperation(ctx context.Context, resourceGroupName string, sqlMigrationServiceName string, options *SQLMigrationServicesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, sqlMigrationServiceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SQLMigrationServicesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, sqlMigrationServiceName string, options *SQLMigrationServicesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataMigration/sqlMigrationServices/{sqlMigrationServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlMigrationServiceName == "" {
		return nil, errors.New("parameter sqlMigrationServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlMigrationServiceName}", url.PathEscape(sqlMigrationServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// DeleteNode - Delete the integration runtime node.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-30-preview
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// sqlMigrationServiceName - Name of the SQL Migration Service.
// parameters - Details of SqlMigrationService resource.
// options - SQLMigrationServicesClientDeleteNodeOptions contains the optional parameters for the SQLMigrationServicesClient.DeleteNode
// method.
func (client *SQLMigrationServicesClient) DeleteNode(ctx context.Context, resourceGroupName string, sqlMigrationServiceName string, parameters DeleteNode, options *SQLMigrationServicesClientDeleteNodeOptions) (SQLMigrationServicesClientDeleteNodeResponse, error) {
	req, err := client.deleteNodeCreateRequest(ctx, resourceGroupName, sqlMigrationServiceName, parameters, options)
	if err != nil {
		return SQLMigrationServicesClientDeleteNodeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLMigrationServicesClientDeleteNodeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SQLMigrationServicesClientDeleteNodeResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteNodeHandleResponse(resp)
}

// deleteNodeCreateRequest creates the DeleteNode request.
func (client *SQLMigrationServicesClient) deleteNodeCreateRequest(ctx context.Context, resourceGroupName string, sqlMigrationServiceName string, parameters DeleteNode, options *SQLMigrationServicesClientDeleteNodeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataMigration/sqlMigrationServices/{sqlMigrationServiceName}/deleteNode"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlMigrationServiceName == "" {
		return nil, errors.New("parameter sqlMigrationServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlMigrationServiceName}", url.PathEscape(sqlMigrationServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// deleteNodeHandleResponse handles the DeleteNode response.
func (client *SQLMigrationServicesClient) deleteNodeHandleResponse(resp *http.Response) (SQLMigrationServicesClientDeleteNodeResponse, error) {
	result := SQLMigrationServicesClientDeleteNodeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeleteNode); err != nil {
		return SQLMigrationServicesClientDeleteNodeResponse{}, err
	}
	return result, nil
}

// Get - Retrieve the Database Migration Service
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-30-preview
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// sqlMigrationServiceName - Name of the SQL Migration Service.
// options - SQLMigrationServicesClientGetOptions contains the optional parameters for the SQLMigrationServicesClient.Get
// method.
func (client *SQLMigrationServicesClient) Get(ctx context.Context, resourceGroupName string, sqlMigrationServiceName string, options *SQLMigrationServicesClientGetOptions) (SQLMigrationServicesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, sqlMigrationServiceName, options)
	if err != nil {
		return SQLMigrationServicesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLMigrationServicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SQLMigrationServicesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SQLMigrationServicesClient) getCreateRequest(ctx context.Context, resourceGroupName string, sqlMigrationServiceName string, options *SQLMigrationServicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataMigration/sqlMigrationServices/{sqlMigrationServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlMigrationServiceName == "" {
		return nil, errors.New("parameter sqlMigrationServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlMigrationServiceName}", url.PathEscape(sqlMigrationServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SQLMigrationServicesClient) getHandleResponse(resp *http.Response) (SQLMigrationServicesClientGetResponse, error) {
	result := SQLMigrationServicesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLMigrationService); err != nil {
		return SQLMigrationServicesClientGetResponse{}, err
	}
	return result, nil
}

// ListAuthKeys - Retrieve the List of Authentication Keys for Self Hosted Integration Runtime.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-30-preview
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// sqlMigrationServiceName - Name of the SQL Migration Service.
// options - SQLMigrationServicesClientListAuthKeysOptions contains the optional parameters for the SQLMigrationServicesClient.ListAuthKeys
// method.
func (client *SQLMigrationServicesClient) ListAuthKeys(ctx context.Context, resourceGroupName string, sqlMigrationServiceName string, options *SQLMigrationServicesClientListAuthKeysOptions) (SQLMigrationServicesClientListAuthKeysResponse, error) {
	req, err := client.listAuthKeysCreateRequest(ctx, resourceGroupName, sqlMigrationServiceName, options)
	if err != nil {
		return SQLMigrationServicesClientListAuthKeysResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLMigrationServicesClientListAuthKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SQLMigrationServicesClientListAuthKeysResponse{}, runtime.NewResponseError(resp)
	}
	return client.listAuthKeysHandleResponse(resp)
}

// listAuthKeysCreateRequest creates the ListAuthKeys request.
func (client *SQLMigrationServicesClient) listAuthKeysCreateRequest(ctx context.Context, resourceGroupName string, sqlMigrationServiceName string, options *SQLMigrationServicesClientListAuthKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataMigration/sqlMigrationServices/{sqlMigrationServiceName}/listAuthKeys"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlMigrationServiceName == "" {
		return nil, errors.New("parameter sqlMigrationServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlMigrationServiceName}", url.PathEscape(sqlMigrationServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAuthKeysHandleResponse handles the ListAuthKeys response.
func (client *SQLMigrationServicesClient) listAuthKeysHandleResponse(resp *http.Response) (SQLMigrationServicesClientListAuthKeysResponse, error) {
	result := SQLMigrationServicesClientListAuthKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AuthenticationKeys); err != nil {
		return SQLMigrationServicesClientListAuthKeysResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Retrieve all SQL migration services in the resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-30-preview
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// options - SQLMigrationServicesClientListByResourceGroupOptions contains the optional parameters for the SQLMigrationServicesClient.ListByResourceGroup
// method.
func (client *SQLMigrationServicesClient) NewListByResourceGroupPager(resourceGroupName string, options *SQLMigrationServicesClientListByResourceGroupOptions) *runtime.Pager[SQLMigrationServicesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[SQLMigrationServicesClientListByResourceGroupResponse]{
		More: func(page SQLMigrationServicesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SQLMigrationServicesClientListByResourceGroupResponse) (SQLMigrationServicesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SQLMigrationServicesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SQLMigrationServicesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SQLMigrationServicesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SQLMigrationServicesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *SQLMigrationServicesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataMigration/sqlMigrationServices"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *SQLMigrationServicesClient) listByResourceGroupHandleResponse(resp *http.Response) (SQLMigrationServicesClientListByResourceGroupResponse, error) {
	result := SQLMigrationServicesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLMigrationListResult); err != nil {
		return SQLMigrationServicesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Retrieve all SQL migration services in the subscriptions.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-30-preview
// options - SQLMigrationServicesClientListBySubscriptionOptions contains the optional parameters for the SQLMigrationServicesClient.ListBySubscription
// method.
func (client *SQLMigrationServicesClient) NewListBySubscriptionPager(options *SQLMigrationServicesClientListBySubscriptionOptions) *runtime.Pager[SQLMigrationServicesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[SQLMigrationServicesClientListBySubscriptionResponse]{
		More: func(page SQLMigrationServicesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SQLMigrationServicesClientListBySubscriptionResponse) (SQLMigrationServicesClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SQLMigrationServicesClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SQLMigrationServicesClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SQLMigrationServicesClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *SQLMigrationServicesClient) listBySubscriptionCreateRequest(ctx context.Context, options *SQLMigrationServicesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DataMigration/sqlMigrationServices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *SQLMigrationServicesClient) listBySubscriptionHandleResponse(resp *http.Response) (SQLMigrationServicesClientListBySubscriptionResponse, error) {
	result := SQLMigrationServicesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLMigrationListResult); err != nil {
		return SQLMigrationServicesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// NewListMigrationsPager - Retrieve the List of database migrations attached to the service.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-30-preview
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// sqlMigrationServiceName - Name of the SQL Migration Service.
// options - SQLMigrationServicesClientListMigrationsOptions contains the optional parameters for the SQLMigrationServicesClient.ListMigrations
// method.
func (client *SQLMigrationServicesClient) NewListMigrationsPager(resourceGroupName string, sqlMigrationServiceName string, options *SQLMigrationServicesClientListMigrationsOptions) *runtime.Pager[SQLMigrationServicesClientListMigrationsResponse] {
	return runtime.NewPager(runtime.PagingHandler[SQLMigrationServicesClientListMigrationsResponse]{
		More: func(page SQLMigrationServicesClientListMigrationsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SQLMigrationServicesClientListMigrationsResponse) (SQLMigrationServicesClientListMigrationsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listMigrationsCreateRequest(ctx, resourceGroupName, sqlMigrationServiceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SQLMigrationServicesClientListMigrationsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SQLMigrationServicesClientListMigrationsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SQLMigrationServicesClientListMigrationsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listMigrationsHandleResponse(resp)
		},
	})
}

// listMigrationsCreateRequest creates the ListMigrations request.
func (client *SQLMigrationServicesClient) listMigrationsCreateRequest(ctx context.Context, resourceGroupName string, sqlMigrationServiceName string, options *SQLMigrationServicesClientListMigrationsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataMigration/sqlMigrationServices/{sqlMigrationServiceName}/listMigrations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlMigrationServiceName == "" {
		return nil, errors.New("parameter sqlMigrationServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlMigrationServiceName}", url.PathEscape(sqlMigrationServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listMigrationsHandleResponse handles the ListMigrations response.
func (client *SQLMigrationServicesClient) listMigrationsHandleResponse(resp *http.Response) (SQLMigrationServicesClientListMigrationsResponse, error) {
	result := SQLMigrationServicesClientListMigrationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseMigrationListResult); err != nil {
		return SQLMigrationServicesClientListMigrationsResponse{}, err
	}
	return result, nil
}

// ListMonitoringData - Retrieve the registered Integration Runtime nodes and their monitoring data for a given Database Migration
// Service.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-30-preview
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// sqlMigrationServiceName - Name of the SQL Migration Service.
// options - SQLMigrationServicesClientListMonitoringDataOptions contains the optional parameters for the SQLMigrationServicesClient.ListMonitoringData
// method.
func (client *SQLMigrationServicesClient) ListMonitoringData(ctx context.Context, resourceGroupName string, sqlMigrationServiceName string, options *SQLMigrationServicesClientListMonitoringDataOptions) (SQLMigrationServicesClientListMonitoringDataResponse, error) {
	req, err := client.listMonitoringDataCreateRequest(ctx, resourceGroupName, sqlMigrationServiceName, options)
	if err != nil {
		return SQLMigrationServicesClientListMonitoringDataResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLMigrationServicesClientListMonitoringDataResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SQLMigrationServicesClientListMonitoringDataResponse{}, runtime.NewResponseError(resp)
	}
	return client.listMonitoringDataHandleResponse(resp)
}

// listMonitoringDataCreateRequest creates the ListMonitoringData request.
func (client *SQLMigrationServicesClient) listMonitoringDataCreateRequest(ctx context.Context, resourceGroupName string, sqlMigrationServiceName string, options *SQLMigrationServicesClientListMonitoringDataOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataMigration/sqlMigrationServices/{sqlMigrationServiceName}/listMonitoringData"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlMigrationServiceName == "" {
		return nil, errors.New("parameter sqlMigrationServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlMigrationServiceName}", url.PathEscape(sqlMigrationServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listMonitoringDataHandleResponse handles the ListMonitoringData response.
func (client *SQLMigrationServicesClient) listMonitoringDataHandleResponse(resp *http.Response) (SQLMigrationServicesClientListMonitoringDataResponse, error) {
	result := SQLMigrationServicesClientListMonitoringDataResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationRuntimeMonitoringData); err != nil {
		return SQLMigrationServicesClientListMonitoringDataResponse{}, err
	}
	return result, nil
}

// RegenerateAuthKeys - Regenerate a new set of Authentication Keys for Self Hosted Integration Runtime.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-30-preview
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// sqlMigrationServiceName - Name of the SQL Migration Service.
// parameters - Details of SqlMigrationService resource.
// options - SQLMigrationServicesClientRegenerateAuthKeysOptions contains the optional parameters for the SQLMigrationServicesClient.RegenerateAuthKeys
// method.
func (client *SQLMigrationServicesClient) RegenerateAuthKeys(ctx context.Context, resourceGroupName string, sqlMigrationServiceName string, parameters RegenAuthKeys, options *SQLMigrationServicesClientRegenerateAuthKeysOptions) (SQLMigrationServicesClientRegenerateAuthKeysResponse, error) {
	req, err := client.regenerateAuthKeysCreateRequest(ctx, resourceGroupName, sqlMigrationServiceName, parameters, options)
	if err != nil {
		return SQLMigrationServicesClientRegenerateAuthKeysResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLMigrationServicesClientRegenerateAuthKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SQLMigrationServicesClientRegenerateAuthKeysResponse{}, runtime.NewResponseError(resp)
	}
	return client.regenerateAuthKeysHandleResponse(resp)
}

// regenerateAuthKeysCreateRequest creates the RegenerateAuthKeys request.
func (client *SQLMigrationServicesClient) regenerateAuthKeysCreateRequest(ctx context.Context, resourceGroupName string, sqlMigrationServiceName string, parameters RegenAuthKeys, options *SQLMigrationServicesClientRegenerateAuthKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataMigration/sqlMigrationServices/{sqlMigrationServiceName}/regenerateAuthKeys"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlMigrationServiceName == "" {
		return nil, errors.New("parameter sqlMigrationServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlMigrationServiceName}", url.PathEscape(sqlMigrationServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// regenerateAuthKeysHandleResponse handles the RegenerateAuthKeys response.
func (client *SQLMigrationServicesClient) regenerateAuthKeysHandleResponse(resp *http.Response) (SQLMigrationServicesClientRegenerateAuthKeysResponse, error) {
	result := SQLMigrationServicesClientRegenerateAuthKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegenAuthKeys); err != nil {
		return SQLMigrationServicesClientRegenerateAuthKeysResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update Database Migration Service.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-30-preview
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// sqlMigrationServiceName - Name of the SQL Migration Service.
// parameters - Details of SqlMigrationService resource.
// options - SQLMigrationServicesClientBeginUpdateOptions contains the optional parameters for the SQLMigrationServicesClient.BeginUpdate
// method.
func (client *SQLMigrationServicesClient) BeginUpdate(ctx context.Context, resourceGroupName string, sqlMigrationServiceName string, parameters SQLMigrationServiceUpdate, options *SQLMigrationServicesClientBeginUpdateOptions) (*runtime.Poller[SQLMigrationServicesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, sqlMigrationServiceName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[SQLMigrationServicesClientUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[SQLMigrationServicesClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Update Database Migration Service.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-30-preview
func (client *SQLMigrationServicesClient) update(ctx context.Context, resourceGroupName string, sqlMigrationServiceName string, parameters SQLMigrationServiceUpdate, options *SQLMigrationServicesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, sqlMigrationServiceName, parameters, options)
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

// updateCreateRequest creates the Update request.
func (client *SQLMigrationServicesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, sqlMigrationServiceName string, parameters SQLMigrationServiceUpdate, options *SQLMigrationServicesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataMigration/sqlMigrationServices/{sqlMigrationServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlMigrationServiceName == "" {
		return nil, errors.New("parameter sqlMigrationServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlMigrationServiceName}", url.PathEscape(sqlMigrationServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}
