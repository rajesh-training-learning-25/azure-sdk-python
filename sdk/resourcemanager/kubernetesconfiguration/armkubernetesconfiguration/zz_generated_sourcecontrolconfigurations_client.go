//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkubernetesconfiguration

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

// SourceControlConfigurationsClient contains the methods for the SourceControlConfigurations group.
// Don't use this type directly, use NewSourceControlConfigurationsClient() instead.
type SourceControlConfigurationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSourceControlConfigurationsClient creates a new instance of SourceControlConfigurationsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSourceControlConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SourceControlConfigurationsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &SourceControlConfigurationsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Create a new Kubernetes Source Control Configuration.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// clusterRp - The Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes, Microsoft.HybridContainerService.
// clusterResourceName - The Kubernetes cluster resource name - i.e. managedClusters, connectedClusters, provisionedClusters.
// clusterName - The name of the kubernetes cluster.
// sourceControlConfigurationName - Name of the Source Control Configuration.
// sourceControlConfiguration - Properties necessary to Create KubernetesConfiguration.
// options - SourceControlConfigurationsClientCreateOrUpdateOptions contains the optional parameters for the SourceControlConfigurationsClient.CreateOrUpdate
// method.
func (client *SourceControlConfigurationsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, sourceControlConfigurationName string, sourceControlConfiguration SourceControlConfiguration, options *SourceControlConfigurationsClientCreateOrUpdateOptions) (SourceControlConfigurationsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, sourceControlConfigurationName, sourceControlConfiguration, options)
	if err != nil {
		return SourceControlConfigurationsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SourceControlConfigurationsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return SourceControlConfigurationsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SourceControlConfigurationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, sourceControlConfigurationName string, sourceControlConfiguration SourceControlConfiguration, options *SourceControlConfigurationsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/sourceControlConfigurations/{sourceControlConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterRp == "" {
		return nil, errors.New("parameter clusterRp cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterRp}", url.PathEscape(clusterRp))
	if clusterResourceName == "" {
		return nil, errors.New("parameter clusterResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterResourceName}", url.PathEscape(clusterResourceName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if sourceControlConfigurationName == "" {
		return nil, errors.New("parameter sourceControlConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sourceControlConfigurationName}", url.PathEscape(sourceControlConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, sourceControlConfiguration)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SourceControlConfigurationsClient) createOrUpdateHandleResponse(resp *http.Response) (SourceControlConfigurationsClientCreateOrUpdateResponse, error) {
	result := SourceControlConfigurationsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SourceControlConfiguration); err != nil {
		return SourceControlConfigurationsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - This will delete the YAML file used to set up the Source control configuration, thus stopping future sync
// from the source repo.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// clusterRp - The Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes, Microsoft.HybridContainerService.
// clusterResourceName - The Kubernetes cluster resource name - i.e. managedClusters, connectedClusters, provisionedClusters.
// clusterName - The name of the kubernetes cluster.
// sourceControlConfigurationName - Name of the Source Control Configuration.
// options - SourceControlConfigurationsClientBeginDeleteOptions contains the optional parameters for the SourceControlConfigurationsClient.BeginDelete
// method.
func (client *SourceControlConfigurationsClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, sourceControlConfigurationName string, options *SourceControlConfigurationsClientBeginDeleteOptions) (*armruntime.Poller[SourceControlConfigurationsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, sourceControlConfigurationName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[SourceControlConfigurationsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[SourceControlConfigurationsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - This will delete the YAML file used to set up the Source control configuration, thus stopping future sync from
// the source repo.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SourceControlConfigurationsClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, sourceControlConfigurationName string, options *SourceControlConfigurationsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, sourceControlConfigurationName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SourceControlConfigurationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, sourceControlConfigurationName string, options *SourceControlConfigurationsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/sourceControlConfigurations/{sourceControlConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterRp == "" {
		return nil, errors.New("parameter clusterRp cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterRp}", url.PathEscape(clusterRp))
	if clusterResourceName == "" {
		return nil, errors.New("parameter clusterResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterResourceName}", url.PathEscape(clusterResourceName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if sourceControlConfigurationName == "" {
		return nil, errors.New("parameter sourceControlConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sourceControlConfigurationName}", url.PathEscape(sourceControlConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets details of the Source Control Configuration.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// clusterRp - The Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes, Microsoft.HybridContainerService.
// clusterResourceName - The Kubernetes cluster resource name - i.e. managedClusters, connectedClusters, provisionedClusters.
// clusterName - The name of the kubernetes cluster.
// sourceControlConfigurationName - Name of the Source Control Configuration.
// options - SourceControlConfigurationsClientGetOptions contains the optional parameters for the SourceControlConfigurationsClient.Get
// method.
func (client *SourceControlConfigurationsClient) Get(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, sourceControlConfigurationName string, options *SourceControlConfigurationsClientGetOptions) (SourceControlConfigurationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, sourceControlConfigurationName, options)
	if err != nil {
		return SourceControlConfigurationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SourceControlConfigurationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SourceControlConfigurationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SourceControlConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, sourceControlConfigurationName string, options *SourceControlConfigurationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/sourceControlConfigurations/{sourceControlConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterRp == "" {
		return nil, errors.New("parameter clusterRp cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterRp}", url.PathEscape(clusterRp))
	if clusterResourceName == "" {
		return nil, errors.New("parameter clusterResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterResourceName}", url.PathEscape(clusterResourceName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if sourceControlConfigurationName == "" {
		return nil, errors.New("parameter sourceControlConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sourceControlConfigurationName}", url.PathEscape(sourceControlConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SourceControlConfigurationsClient) getHandleResponse(resp *http.Response) (SourceControlConfigurationsClientGetResponse, error) {
	result := SourceControlConfigurationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SourceControlConfiguration); err != nil {
		return SourceControlConfigurationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List all Source Control Configurations.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// clusterRp - The Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes, Microsoft.HybridContainerService.
// clusterResourceName - The Kubernetes cluster resource name - i.e. managedClusters, connectedClusters, provisionedClusters.
// clusterName - The name of the kubernetes cluster.
// options - SourceControlConfigurationsClientListOptions contains the optional parameters for the SourceControlConfigurationsClient.List
// method.
func (client *SourceControlConfigurationsClient) NewListPager(resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, options *SourceControlConfigurationsClientListOptions) *runtime.Pager[SourceControlConfigurationsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[SourceControlConfigurationsClientListResponse]{
		More: func(page SourceControlConfigurationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SourceControlConfigurationsClientListResponse) (SourceControlConfigurationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SourceControlConfigurationsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SourceControlConfigurationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SourceControlConfigurationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SourceControlConfigurationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, options *SourceControlConfigurationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/sourceControlConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterRp == "" {
		return nil, errors.New("parameter clusterRp cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterRp}", url.PathEscape(clusterRp))
	if clusterResourceName == "" {
		return nil, errors.New("parameter clusterResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterResourceName}", url.PathEscape(clusterResourceName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SourceControlConfigurationsClient) listHandleResponse(resp *http.Response) (SourceControlConfigurationsClientListResponse, error) {
	result := SourceControlConfigurationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SourceControlConfigurationList); err != nil {
		return SourceControlConfigurationsClientListResponse{}, err
	}
	return result, nil
}
