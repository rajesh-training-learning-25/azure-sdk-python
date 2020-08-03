package enterpriseknowledgegraphservice

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

// EnterpriseKnowledgeGraphClient is the azure Enterprise Knowledge Graph Service is a platform for creating knowledge
// graphs at scale.
type EnterpriseKnowledgeGraphClient struct {
	BaseClient
}

// NewEnterpriseKnowledgeGraphClient creates an instance of the EnterpriseKnowledgeGraphClient client.
func NewEnterpriseKnowledgeGraphClient(subscriptionID string) EnterpriseKnowledgeGraphClient {
	return NewEnterpriseKnowledgeGraphClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewEnterpriseKnowledgeGraphClientWithBaseURI creates an instance of the EnterpriseKnowledgeGraphClient client using
// a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewEnterpriseKnowledgeGraphClientWithBaseURI(baseURI string, subscriptionID string) EnterpriseKnowledgeGraphClient {
	return EnterpriseKnowledgeGraphClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates a EnterpriseKnowledgeGraph Service. EnterpriseKnowledgeGraph Service is a resource group wide
// resource type.
// Parameters:
// resourceGroupName - the name of the EnterpriseKnowledgeGraph resource group in the user subscription.
// resourceName - the name of the EnterpriseKnowledgeGraph resource.
// parameters - the parameters to provide for the created EnterpriseKnowledgeGraph.
func (client EnterpriseKnowledgeGraphClient) Create(ctx context.Context, resourceGroupName string, resourceName string, parameters EnterpriseKnowledgeGraph) (result EnterpriseKnowledgeGraph, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnterpriseKnowledgeGraphClient.Create")
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
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("enterpriseknowledgegraphservice.EnterpriseKnowledgeGraphClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, resourceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "enterpriseknowledgegraphservice.EnterpriseKnowledgeGraphClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "enterpriseknowledgegraphservice.EnterpriseKnowledgeGraphClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "enterpriseknowledgegraphservice.EnterpriseKnowledgeGraphClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client EnterpriseKnowledgeGraphClient) CreatePreparer(ctx context.Context, resourceGroupName string, resourceName string, parameters EnterpriseKnowledgeGraph) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-12-03"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EnterpriseKnowledgeGraph/services/{resourceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client EnterpriseKnowledgeGraphClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client EnterpriseKnowledgeGraphClient) CreateResponder(resp *http.Response) (result EnterpriseKnowledgeGraph, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a EnterpriseKnowledgeGraph Service from the resource group.
// Parameters:
// resourceGroupName - the name of the EnterpriseKnowledgeGraph resource group in the user subscription.
// resourceName - the name of the EnterpriseKnowledgeGraph resource.
func (client EnterpriseKnowledgeGraphClient) Delete(ctx context.Context, resourceGroupName string, resourceName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnterpriseKnowledgeGraphClient.Delete")
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
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("enterpriseknowledgegraphservice.EnterpriseKnowledgeGraphClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "enterpriseknowledgegraphservice.EnterpriseKnowledgeGraphClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "enterpriseknowledgegraphservice.EnterpriseKnowledgeGraphClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "enterpriseknowledgegraphservice.EnterpriseKnowledgeGraphClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client EnterpriseKnowledgeGraphClient) DeletePreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-12-03"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EnterpriseKnowledgeGraph/services/{resourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client EnterpriseKnowledgeGraphClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client EnterpriseKnowledgeGraphClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get returns a EnterpriseKnowledgeGraph service specified by the parameters.
// Parameters:
// resourceGroupName - the name of the EnterpriseKnowledgeGraph resource group in the user subscription.
// resourceName - the name of the EnterpriseKnowledgeGraph resource.
func (client EnterpriseKnowledgeGraphClient) Get(ctx context.Context, resourceGroupName string, resourceName string) (result EnterpriseKnowledgeGraph, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnterpriseKnowledgeGraphClient.Get")
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
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("enterpriseknowledgegraphservice.EnterpriseKnowledgeGraphClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "enterpriseknowledgegraphservice.EnterpriseKnowledgeGraphClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "enterpriseknowledgegraphservice.EnterpriseKnowledgeGraphClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "enterpriseknowledgegraphservice.EnterpriseKnowledgeGraphClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client EnterpriseKnowledgeGraphClient) GetPreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-12-03"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EnterpriseKnowledgeGraph/services/{resourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client EnterpriseKnowledgeGraphClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client EnterpriseKnowledgeGraphClient) GetResponder(resp *http.Response) (result EnterpriseKnowledgeGraph, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List returns all the resources of a particular type belonging to a subscription.
func (client EnterpriseKnowledgeGraphClient) List(ctx context.Context) (result EnterpriseKnowledgeGraphResponseListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnterpriseKnowledgeGraphClient.List")
		defer func() {
			sc := -1
			if result.ekgrl.Response.Response != nil {
				sc = result.ekgrl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "enterpriseknowledgegraphservice.EnterpriseKnowledgeGraphClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.ekgrl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "enterpriseknowledgegraphservice.EnterpriseKnowledgeGraphClient", "List", resp, "Failure sending request")
		return
	}

	result.ekgrl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "enterpriseknowledgegraphservice.EnterpriseKnowledgeGraphClient", "List", resp, "Failure responding to request")
	}
	if result.ekgrl.hasNextLink() && result.ekgrl.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client EnterpriseKnowledgeGraphClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-12-03"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.EnterpriseKnowledgeGraph/services", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client EnterpriseKnowledgeGraphClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client EnterpriseKnowledgeGraphClient) ListResponder(resp *http.Response) (result EnterpriseKnowledgeGraphResponseList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client EnterpriseKnowledgeGraphClient) listNextResults(ctx context.Context, lastResults EnterpriseKnowledgeGraphResponseList) (result EnterpriseKnowledgeGraphResponseList, err error) {
	req, err := lastResults.enterpriseKnowledgeGraphResponseListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "enterpriseknowledgegraphservice.EnterpriseKnowledgeGraphClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "enterpriseknowledgegraphservice.EnterpriseKnowledgeGraphClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "enterpriseknowledgegraphservice.EnterpriseKnowledgeGraphClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client EnterpriseKnowledgeGraphClient) ListComplete(ctx context.Context) (result EnterpriseKnowledgeGraphResponseListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnterpriseKnowledgeGraphClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// ListByResourceGroup returns all the resources of a particular type belonging to a resource group
// Parameters:
// resourceGroupName - the name of the EnterpriseKnowledgeGraph resource group in the user subscription.
func (client EnterpriseKnowledgeGraphClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result EnterpriseKnowledgeGraphResponseListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnterpriseKnowledgeGraphClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.ekgrl.Response.Response != nil {
				sc = result.ekgrl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("enterpriseknowledgegraphservice.EnterpriseKnowledgeGraphClient", "ListByResourceGroup", err.Error())
	}

	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "enterpriseknowledgegraphservice.EnterpriseKnowledgeGraphClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.ekgrl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "enterpriseknowledgegraphservice.EnterpriseKnowledgeGraphClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.ekgrl, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "enterpriseknowledgegraphservice.EnterpriseKnowledgeGraphClient", "ListByResourceGroup", resp, "Failure responding to request")
	}
	if result.ekgrl.hasNextLink() && result.ekgrl.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client EnterpriseKnowledgeGraphClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-12-03"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EnterpriseKnowledgeGraph/services", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client EnterpriseKnowledgeGraphClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client EnterpriseKnowledgeGraphClient) ListByResourceGroupResponder(resp *http.Response) (result EnterpriseKnowledgeGraphResponseList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client EnterpriseKnowledgeGraphClient) listByResourceGroupNextResults(ctx context.Context, lastResults EnterpriseKnowledgeGraphResponseList) (result EnterpriseKnowledgeGraphResponseList, err error) {
	req, err := lastResults.enterpriseKnowledgeGraphResponseListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "enterpriseknowledgegraphservice.EnterpriseKnowledgeGraphClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "enterpriseknowledgegraphservice.EnterpriseKnowledgeGraphClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "enterpriseknowledgegraphservice.EnterpriseKnowledgeGraphClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client EnterpriseKnowledgeGraphClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result EnterpriseKnowledgeGraphResponseListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnterpriseKnowledgeGraphClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// Update updates a EnterpriseKnowledgeGraph Service
// Parameters:
// resourceGroupName - the name of the EnterpriseKnowledgeGraph resource group in the user subscription.
// resourceName - the name of the EnterpriseKnowledgeGraph resource.
// parameters - the parameters to provide for the created EnterpriseKnowledgeGraph.
func (client EnterpriseKnowledgeGraphClient) Update(ctx context.Context, resourceGroupName string, resourceName string, parameters EnterpriseKnowledgeGraph) (result EnterpriseKnowledgeGraph, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnterpriseKnowledgeGraphClient.Update")
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
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("enterpriseknowledgegraphservice.EnterpriseKnowledgeGraphClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, resourceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "enterpriseknowledgegraphservice.EnterpriseKnowledgeGraphClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "enterpriseknowledgegraphservice.EnterpriseKnowledgeGraphClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "enterpriseknowledgegraphservice.EnterpriseKnowledgeGraphClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client EnterpriseKnowledgeGraphClient) UpdatePreparer(ctx context.Context, resourceGroupName string, resourceName string, parameters EnterpriseKnowledgeGraph) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-12-03"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EnterpriseKnowledgeGraph/services/{resourceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client EnterpriseKnowledgeGraphClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client EnterpriseKnowledgeGraphClient) UpdateResponder(resp *http.Response) (result EnterpriseKnowledgeGraph, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
