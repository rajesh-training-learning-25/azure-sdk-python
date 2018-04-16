package servicefabric

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// ApplicationTypesClient is the service Fabric Management Client
type ApplicationTypesClient struct {
	BaseClient
}

// NewApplicationTypesClient creates an instance of the ApplicationTypesClient client.
func NewApplicationTypesClient() ApplicationTypesClient {
	return NewApplicationTypesClientWithBaseURI(DefaultBaseURI)
}

// NewApplicationTypesClientWithBaseURI creates an instance of the ApplicationTypesClient client.
func NewApplicationTypesClientWithBaseURI(baseURI string) ApplicationTypesClient {
	return ApplicationTypesClient{NewWithBaseURI(baseURI)}
}

// Create create or update a Service Fabric application type name resource with the specified name.
//
// subscriptionID is the customer subscription identifier. resourceGroupName is the name of the resource group.
// clusterName is the name of the cluster resource. applicationTypeName is the name of the application type name
// resource. parameters is the application type name resource.
func (client ApplicationTypesClient) Create(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string, applicationTypeName string, parameters ApplicationTypeResource) (result ApplicationTypeResource, err error) {
	req, err := client.CreatePreparer(ctx, subscriptionID, resourceGroupName, clusterName, applicationTypeName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationTypesClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationTypesClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationTypesClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ApplicationTypesClient) CreatePreparer(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string, applicationTypeName string, parameters ApplicationTypeResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationTypeName": autorest.Encode("path", applicationTypeName),
		"clusterName":         autorest.Encode("path", clusterName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2017-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationTypesClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ApplicationTypesClient) CreateResponder(resp *http.Response) (result ApplicationTypeResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a Service Fabric application type name resource with the specified name.
//
// subscriptionID is the customer subscription identifier. resourceGroupName is the name of the resource group.
// clusterName is the name of the cluster resource. applicationTypeName is the name of the application type name
// resource.
func (client ApplicationTypesClient) Delete(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string, applicationTypeName string) (result ApplicationTypesDeleteFuture, err error) {
	req, err := client.DeletePreparer(ctx, subscriptionID, resourceGroupName, clusterName, applicationTypeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationTypesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationTypesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ApplicationTypesClient) DeletePreparer(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string, applicationTypeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationTypeName": autorest.Encode("path", applicationTypeName),
		"clusterName":         autorest.Encode("path", clusterName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2017-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationTypesClient) DeleteSender(req *http.Request) (future ApplicationTypesDeleteFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ApplicationTypesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get a Service Fabric application type name resource created or in the process of being created in the Service
// Fabric cluster resource.
//
// subscriptionID is the customer subscription identifier. resourceGroupName is the name of the resource group.
// clusterName is the name of the cluster resource. applicationTypeName is the name of the application type name
// resource.
func (client ApplicationTypesClient) Get(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string, applicationTypeName string) (result ApplicationTypeResource, err error) {
	req, err := client.GetPreparer(ctx, subscriptionID, resourceGroupName, clusterName, applicationTypeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationTypesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationTypesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationTypesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ApplicationTypesClient) GetPreparer(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string, applicationTypeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationTypeName": autorest.Encode("path", applicationTypeName),
		"clusterName":         autorest.Encode("path", clusterName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2017-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationTypesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ApplicationTypesClient) GetResponder(resp *http.Response) (result ApplicationTypeResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all application type name resources created or in the process of being created in the Service Fabric
// cluster resource.
//
// subscriptionID is the customer subscription identifier. resourceGroupName is the name of the resource group.
// clusterName is the name of the cluster resource.
func (client ApplicationTypesClient) List(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string) (result ApplicationTypeResourceList, err error) {
	req, err := client.ListPreparer(ctx, subscriptionID, resourceGroupName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationTypesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationTypesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationTypesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ApplicationTypesClient) ListPreparer(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2017-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationTypesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ApplicationTypesClient) ListResponder(resp *http.Response) (result ApplicationTypeResourceList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
