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

// ManagedInstanceTdeCertificatesClient contains the methods for the ManagedInstanceTdeCertificates group.
// Don't use this type directly, use NewManagedInstanceTdeCertificatesClient() instead.
type ManagedInstanceTdeCertificatesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewManagedInstanceTdeCertificatesClient creates a new instance of ManagedInstanceTdeCertificatesClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewManagedInstanceTdeCertificatesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedInstanceTdeCertificatesClient, error) {
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
	client := &ManagedInstanceTdeCertificatesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Creates a TDE certificate for a given server.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// parameters - The requested TDE certificate to be created or updated.
// options - ManagedInstanceTdeCertificatesClientBeginCreateOptions contains the optional parameters for the ManagedInstanceTdeCertificatesClient.BeginCreate
// method.
func (client *ManagedInstanceTdeCertificatesClient) BeginCreate(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters TdeCertificate, options *ManagedInstanceTdeCertificatesClientBeginCreateOptions) (*runtime.Poller[ManagedInstanceTdeCertificatesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, managedInstanceName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ManagedInstanceTdeCertificatesClientCreateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ManagedInstanceTdeCertificatesClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Creates a TDE certificate for a given server.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
func (client *ManagedInstanceTdeCertificatesClient) create(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters TdeCertificate, options *ManagedInstanceTdeCertificatesClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, managedInstanceName, parameters, options)
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

// createCreateRequest creates the Create request.
func (client *ManagedInstanceTdeCertificatesClient) createCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters TdeCertificate, options *ManagedInstanceTdeCertificatesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/tdeCertificates"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, runtime.MarshalAsJSON(req, parameters)
}