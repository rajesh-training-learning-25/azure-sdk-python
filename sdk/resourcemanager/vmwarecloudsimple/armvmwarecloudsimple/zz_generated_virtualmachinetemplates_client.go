//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvmwarecloudsimple

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

// VirtualMachineTemplatesClient contains the methods for the VirtualMachineTemplates group.
// Don't use this type directly, use NewVirtualMachineTemplatesClient() instead.
type VirtualMachineTemplatesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualMachineTemplatesClient creates a new instance of VirtualMachineTemplatesClient with the specified values.
// subscriptionID - The subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVirtualMachineTemplatesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualMachineTemplatesClient, error) {
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
	client := &VirtualMachineTemplatesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Returns virtual machine templates by its name
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-04-01
// regionID - The region Id (westus, eastus)
// pcName - The private cloud name
// virtualMachineTemplateName - virtual machine template id (vsphereId)
// options - VirtualMachineTemplatesClientGetOptions contains the optional parameters for the VirtualMachineTemplatesClient.Get
// method.
func (client *VirtualMachineTemplatesClient) Get(ctx context.Context, regionID string, pcName string, virtualMachineTemplateName string, options *VirtualMachineTemplatesClientGetOptions) (VirtualMachineTemplatesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, regionID, pcName, virtualMachineTemplateName, options)
	if err != nil {
		return VirtualMachineTemplatesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachineTemplatesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineTemplatesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualMachineTemplatesClient) getCreateRequest(ctx context.Context, regionID string, pcName string, virtualMachineTemplateName string, options *VirtualMachineTemplatesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/virtualMachineTemplates/{virtualMachineTemplateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if regionID == "" {
		return nil, errors.New("parameter regionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regionId}", url.PathEscape(regionID))
	if pcName == "" {
		return nil, errors.New("parameter pcName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pcName}", url.PathEscape(pcName))
	if virtualMachineTemplateName == "" {
		return nil, errors.New("parameter virtualMachineTemplateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineTemplateName}", url.PathEscape(virtualMachineTemplateName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualMachineTemplatesClient) getHandleResponse(resp *http.Response) (VirtualMachineTemplatesClientGetResponse, error) {
	result := VirtualMachineTemplatesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineTemplate); err != nil {
		return VirtualMachineTemplatesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Returns list of virtual machine templates in region for private cloud
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-04-01
// pcName - The private cloud name
// regionID - The region Id (westus, eastus)
// resourcePoolName - Resource pool used to derive vSphere cluster which contains VM templates
// options - VirtualMachineTemplatesClientListOptions contains the optional parameters for the VirtualMachineTemplatesClient.List
// method.
func (client *VirtualMachineTemplatesClient) NewListPager(pcName string, regionID string, resourcePoolName string, options *VirtualMachineTemplatesClientListOptions) *runtime.Pager[VirtualMachineTemplatesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualMachineTemplatesClientListResponse]{
		More: func(page VirtualMachineTemplatesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualMachineTemplatesClientListResponse) (VirtualMachineTemplatesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, pcName, regionID, resourcePoolName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualMachineTemplatesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VirtualMachineTemplatesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualMachineTemplatesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *VirtualMachineTemplatesClient) listCreateRequest(ctx context.Context, pcName string, regionID string, resourcePoolName string, options *VirtualMachineTemplatesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/virtualMachineTemplates"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if pcName == "" {
		return nil, errors.New("parameter pcName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pcName}", url.PathEscape(pcName))
	if regionID == "" {
		return nil, errors.New("parameter regionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regionId}", url.PathEscape(regionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-04-01")
	reqQP.Set("resourcePoolName", resourcePoolName)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualMachineTemplatesClient) listHandleResponse(resp *http.Response) (VirtualMachineTemplatesClientListResponse, error) {
	result := VirtualMachineTemplatesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineTemplateListResponse); err != nil {
		return VirtualMachineTemplatesClientListResponse{}, err
	}
	return result, nil
}