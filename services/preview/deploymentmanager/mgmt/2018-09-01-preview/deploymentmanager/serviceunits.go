package deploymentmanager

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ServiceUnitsClient is the REST APIs for orchestrating deployments using the Azure Deployment Manager (ADM). See
// https://docs.microsoft.com/en-us/azure/azure-resource-manager/deployment-manager-overview for more information.
type ServiceUnitsClient struct {
	BaseClient
}

// NewServiceUnitsClient creates an instance of the ServiceUnitsClient client.
func NewServiceUnitsClient(subscriptionID string) ServiceUnitsClient {
	return NewServiceUnitsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewServiceUnitsClientWithBaseURI creates an instance of the ServiceUnitsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewServiceUnitsClientWithBaseURI(baseURI string, subscriptionID string) ServiceUnitsClient {
	return ServiceUnitsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate this is an asynchronous operation and can be polled to completion using the operation resource
// returned by this operation.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// serviceTopologyName - the name of the service topology .
// serviceName - the name of the service resource.
// serviceUnitName - the name of the service unit resource.
// serviceUnitInfo - the service unit resource object.
func (client ServiceUnitsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceTopologyName string, serviceName string, serviceUnitName string, serviceUnitInfo ServiceUnitResource) (result ServiceUnitsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceUnitsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: serviceUnitInfo,
			Constraints: []validation.Constraint{{Target: "serviceUnitInfo.ServiceUnitResourceProperties", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("deploymentmanager.ServiceUnitsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serviceTopologyName, serviceName, serviceUnitName, serviceUnitInfo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentmanager.ServiceUnitsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentmanager.ServiceUnitsClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ServiceUnitsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serviceTopologyName string, serviceName string, serviceUnitName string, serviceUnitInfo ServiceUnitResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"serviceName":         autorest.Encode("path", serviceName),
		"serviceTopologyName": autorest.Encode("path", serviceTopologyName),
		"serviceUnitName":     autorest.Encode("path", serviceUnitName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/serviceTopologies/{serviceTopologyName}/services/{serviceName}/serviceUnits/{serviceUnitName}", pathParameters),
		autorest.WithJSON(serviceUnitInfo),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceUnitsClient) CreateOrUpdateSender(req *http.Request) (future ServiceUnitsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ServiceUnitsClient) (sur ServiceUnitResource, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "deploymentmanager.ServiceUnitsCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("deploymentmanager.ServiceUnitsCreateOrUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		if sur.Response.Response, err = future.GetResult(sender); err == nil && sur.Response.Response.StatusCode != http.StatusNoContent {
			sur, err = client.CreateOrUpdateResponder(sur.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "deploymentmanager.ServiceUnitsCreateOrUpdateFuture", "Result", sur.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ServiceUnitsClient) CreateOrUpdateResponder(resp *http.Response) (result ServiceUnitResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete sends the delete request.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// serviceTopologyName - the name of the service topology .
// serviceName - the name of the service resource.
// serviceUnitName - the name of the service unit resource.
func (client ServiceUnitsClient) Delete(ctx context.Context, resourceGroupName string, serviceTopologyName string, serviceName string, serviceUnitName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceUnitsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("deploymentmanager.ServiceUnitsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, serviceTopologyName, serviceName, serviceUnitName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentmanager.ServiceUnitsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "deploymentmanager.ServiceUnitsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentmanager.ServiceUnitsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ServiceUnitsClient) DeletePreparer(ctx context.Context, resourceGroupName string, serviceTopologyName string, serviceName string, serviceUnitName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"serviceName":         autorest.Encode("path", serviceName),
		"serviceTopologyName": autorest.Encode("path", serviceTopologyName),
		"serviceUnitName":     autorest.Encode("path", serviceUnitName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/serviceTopologies/{serviceTopologyName}/services/{serviceName}/serviceUnits/{serviceUnitName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceUnitsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ServiceUnitsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get sends the get request.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// serviceTopologyName - the name of the service topology .
// serviceName - the name of the service resource.
// serviceUnitName - the name of the service unit resource.
func (client ServiceUnitsClient) Get(ctx context.Context, resourceGroupName string, serviceTopologyName string, serviceName string, serviceUnitName string) (result ServiceUnitResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceUnitsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("deploymentmanager.ServiceUnitsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, serviceTopologyName, serviceName, serviceUnitName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentmanager.ServiceUnitsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "deploymentmanager.ServiceUnitsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentmanager.ServiceUnitsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ServiceUnitsClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceTopologyName string, serviceName string, serviceUnitName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"serviceName":         autorest.Encode("path", serviceName),
		"serviceTopologyName": autorest.Encode("path", serviceTopologyName),
		"serviceUnitName":     autorest.Encode("path", serviceUnitName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/serviceTopologies/{serviceTopologyName}/services/{serviceName}/serviceUnits/{serviceUnitName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceUnitsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ServiceUnitsClient) GetResponder(resp *http.Response) (result ServiceUnitResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
