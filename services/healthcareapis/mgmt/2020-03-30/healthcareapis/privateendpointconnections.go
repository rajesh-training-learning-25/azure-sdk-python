package healthcareapis

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

// PrivateEndpointConnectionsClient is the azure Healthcare APIs Client
type PrivateEndpointConnectionsClient struct {
	BaseClient
}

// NewPrivateEndpointConnectionsClient creates an instance of the PrivateEndpointConnectionsClient client.
func NewPrivateEndpointConnectionsClient(subscriptionID string) PrivateEndpointConnectionsClient {
	return NewPrivateEndpointConnectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPrivateEndpointConnectionsClientWithBaseURI creates an instance of the PrivateEndpointConnectionsClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewPrivateEndpointConnectionsClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionsClient {
	return PrivateEndpointConnectionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate update the state of the specified private endpoint connection associated with the service.
// Parameters:
// resourceGroupName - the name of the resource group that contains the service instance.
// resourceName - the name of the service instance.
// privateEndpointConnectionName - the name of the private endpoint connection associated with the Azure
// resource
// properties - the private endpoint connection properties.
func (client PrivateEndpointConnectionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, properties PrivateEndpointConnection) (result PrivateEndpointConnectionsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointConnectionsClient.CreateOrUpdate")
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
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: properties,
			Constraints: []validation.Constraint{{Target: "properties.PrivateEndpointConnectionProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "properties.PrivateEndpointConnectionProperties.PrivateLinkServiceConnectionState", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("healthcareapis.PrivateEndpointConnectionsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, resourceName, privateEndpointConnectionName, properties)
	if err != nil {
		err = autorest.NewErrorWithError(err, "healthcareapis.PrivateEndpointConnectionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "healthcareapis.PrivateEndpointConnectionsClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client PrivateEndpointConnectionsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, properties PrivateEndpointConnection) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateEndpointConnectionName": autorest.Encode("path", privateEndpointConnectionName),
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"resourceName":                  autorest.Encode("path", resourceName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/services/{resourceName}/privateEndpointConnections/{privateEndpointConnectionName}", pathParameters),
		autorest.WithJSON(properties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateEndpointConnectionsClient) CreateOrUpdateSender(req *http.Request) (future PrivateEndpointConnectionsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client PrivateEndpointConnectionsClient) (pec PrivateEndpointConnection, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "healthcareapis.PrivateEndpointConnectionsCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("healthcareapis.PrivateEndpointConnectionsCreateOrUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		if pec.Response.Response, err = future.GetResult(sender); err == nil && pec.Response.Response.StatusCode != http.StatusNoContent {
			pec, err = client.CreateOrUpdateResponder(pec.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "healthcareapis.PrivateEndpointConnectionsCreateOrUpdateFuture", "Result", pec.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client PrivateEndpointConnectionsClient) CreateOrUpdateResponder(resp *http.Response) (result PrivateEndpointConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a private endpoint connection.
// Parameters:
// resourceGroupName - the name of the resource group that contains the service instance.
// resourceName - the name of the service instance.
// privateEndpointConnectionName - the name of the private endpoint connection associated with the Azure
// resource
func (client PrivateEndpointConnectionsClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string) (result PrivateEndpointConnectionsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointConnectionsClient.Delete")
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
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("healthcareapis.PrivateEndpointConnectionsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, resourceName, privateEndpointConnectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "healthcareapis.PrivateEndpointConnectionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "healthcareapis.PrivateEndpointConnectionsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client PrivateEndpointConnectionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateEndpointConnectionName": autorest.Encode("path", privateEndpointConnectionName),
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"resourceName":                  autorest.Encode("path", resourceName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/services/{resourceName}/privateEndpointConnections/{privateEndpointConnectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateEndpointConnectionsClient) DeleteSender(req *http.Request) (future PrivateEndpointConnectionsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client PrivateEndpointConnectionsClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "healthcareapis.PrivateEndpointConnectionsDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("healthcareapis.PrivateEndpointConnectionsDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client PrivateEndpointConnectionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified private endpoint connection associated with the service.
// Parameters:
// resourceGroupName - the name of the resource group that contains the service instance.
// resourceName - the name of the service instance.
// privateEndpointConnectionName - the name of the private endpoint connection associated with the Azure
// resource
func (client PrivateEndpointConnectionsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string) (result PrivateEndpointConnection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointConnectionsClient.Get")
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
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("healthcareapis.PrivateEndpointConnectionsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, resourceName, privateEndpointConnectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "healthcareapis.PrivateEndpointConnectionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "healthcareapis.PrivateEndpointConnectionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "healthcareapis.PrivateEndpointConnectionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client PrivateEndpointConnectionsClient) GetPreparer(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateEndpointConnectionName": autorest.Encode("path", privateEndpointConnectionName),
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"resourceName":                  autorest.Encode("path", resourceName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/services/{resourceName}/privateEndpointConnections/{privateEndpointConnectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateEndpointConnectionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PrivateEndpointConnectionsClient) GetResponder(resp *http.Response) (result PrivateEndpointConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByService lists all private endpoint connections for a service.
// Parameters:
// resourceGroupName - the name of the resource group that contains the service instance.
// resourceName - the name of the service instance.
func (client PrivateEndpointConnectionsClient) ListByService(ctx context.Context, resourceGroupName string, resourceName string) (result PrivateEndpointConnectionListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointConnectionsClient.ListByService")
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
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("healthcareapis.PrivateEndpointConnectionsClient", "ListByService", err.Error())
	}

	req, err := client.ListByServicePreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "healthcareapis.PrivateEndpointConnectionsClient", "ListByService", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "healthcareapis.PrivateEndpointConnectionsClient", "ListByService", resp, "Failure sending request")
		return
	}

	result, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "healthcareapis.PrivateEndpointConnectionsClient", "ListByService", resp, "Failure responding to request")
		return
	}

	return
}

// ListByServicePreparer prepares the ListByService request.
func (client PrivateEndpointConnectionsClient) ListByServicePreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/services/{resourceName}/privateEndpointConnections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByServiceSender sends the ListByService request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateEndpointConnectionsClient) ListByServiceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByServiceResponder handles the response to the ListByService request. The method always
// closes the http.Response Body.
func (client PrivateEndpointConnectionsClient) ListByServiceResponder(resp *http.Response) (result PrivateEndpointConnectionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
