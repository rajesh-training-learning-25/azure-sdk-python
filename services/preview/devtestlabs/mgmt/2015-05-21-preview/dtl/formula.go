package dtl

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// FormulaClient is the azure DevTest Labs REST API version 2015-05-21-preview.
type FormulaClient struct {
	BaseClient
}

// NewFormulaClient creates an instance of the FormulaClient client.
func NewFormulaClient(subscriptionID string) FormulaClient {
	return NewFormulaClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewFormulaClientWithBaseURI creates an instance of the FormulaClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewFormulaClientWithBaseURI(baseURI string, subscriptionID string) FormulaClient {
	return FormulaClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdateResource create or replace an existing Formula. This operation can take a while to complete.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// name - the name of the formula.
func (client FormulaClient) CreateOrUpdateResource(ctx context.Context, resourceGroupName string, labName string, name string, formula Formula) (result FormulaCreateOrUpdateResourceFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FormulaClient.CreateOrUpdateResource")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdateResourcePreparer(ctx, resourceGroupName, labName, name, formula)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.FormulaClient", "CreateOrUpdateResource", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateResourceSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.FormulaClient", "CreateOrUpdateResource", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdateResourcePreparer prepares the CreateOrUpdateResource request.
func (client FormulaClient) CreateOrUpdateResourcePreparer(ctx context.Context, resourceGroupName string, labName string, name string, formula Formula) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-21-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/formulas/{name}", pathParameters),
		autorest.WithJSON(formula),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateResourceSender sends the CreateOrUpdateResource request. The method will close the
// http.Response Body if it receives an error.
func (client FormulaClient) CreateOrUpdateResourceSender(req *http.Request) (future FormulaCreateOrUpdateResourceFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResourceResponder handles the response to the CreateOrUpdateResource request. The method always
// closes the http.Response Body.
func (client FormulaClient) CreateOrUpdateResourceResponder(resp *http.Response) (result Formula, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteResource delete formula.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// name - the name of the formula.
func (client FormulaClient) DeleteResource(ctx context.Context, resourceGroupName string, labName string, name string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FormulaClient.DeleteResource")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteResourcePreparer(ctx, resourceGroupName, labName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.FormulaClient", "DeleteResource", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteResourceSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "dtl.FormulaClient", "DeleteResource", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.FormulaClient", "DeleteResource", resp, "Failure responding to request")
	}

	return
}

// DeleteResourcePreparer prepares the DeleteResource request.
func (client FormulaClient) DeleteResourcePreparer(ctx context.Context, resourceGroupName string, labName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-21-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/formulas/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteResourceSender sends the DeleteResource request. The method will close the
// http.Response Body if it receives an error.
func (client FormulaClient) DeleteResourceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResourceResponder handles the response to the DeleteResource request. The method always
// closes the http.Response Body.
func (client FormulaClient) DeleteResourceResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetResource get formula.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// name - the name of the formula.
func (client FormulaClient) GetResource(ctx context.Context, resourceGroupName string, labName string, name string) (result Formula, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FormulaClient.GetResource")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetResourcePreparer(ctx, resourceGroupName, labName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.FormulaClient", "GetResource", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.FormulaClient", "GetResource", resp, "Failure sending request")
		return
	}

	result, err = client.GetResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.FormulaClient", "GetResource", resp, "Failure responding to request")
	}

	return
}

// GetResourcePreparer prepares the GetResource request.
func (client FormulaClient) GetResourcePreparer(ctx context.Context, resourceGroupName string, labName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-21-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/formulas/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetResourceSender sends the GetResource request. The method will close the
// http.Response Body if it receives an error.
func (client FormulaClient) GetResourceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResourceResponder handles the response to the GetResource request. The method always
// closes the http.Response Body.
func (client FormulaClient) GetResourceResponder(resp *http.Response) (result Formula, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list formulas.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// filter - the filter to apply on the operation.
func (client FormulaClient) List(ctx context.Context, resourceGroupName string, labName string, filter string, top *int32, orderBy string) (result ResponseWithContinuationFormulaPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FormulaClient.List")
		defer func() {
			sc := -1
			if result.rwcf.Response.Response != nil {
				sc = result.rwcf.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, labName, filter, top, orderBy)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.FormulaClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rwcf.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.FormulaClient", "List", resp, "Failure sending request")
		return
	}

	result.rwcf, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.FormulaClient", "List", resp, "Failure responding to request")
	}
	if result.rwcf.hasNextLink() && result.rwcf.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client FormulaClient) ListPreparer(ctx context.Context, resourceGroupName string, labName string, filter string, top *int32, orderBy string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-21-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderBy) > 0 {
		queryParameters["$orderBy"] = autorest.Encode("query", orderBy)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/formulas", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client FormulaClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client FormulaClient) ListResponder(resp *http.Response) (result ResponseWithContinuationFormula, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client FormulaClient) listNextResults(ctx context.Context, lastResults ResponseWithContinuationFormula) (result ResponseWithContinuationFormula, err error) {
	req, err := lastResults.responseWithContinuationFormulaPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dtl.FormulaClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dtl.FormulaClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.FormulaClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client FormulaClient) ListComplete(ctx context.Context, resourceGroupName string, labName string, filter string, top *int32, orderBy string) (result ResponseWithContinuationFormulaIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FormulaClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, labName, filter, top, orderBy)
	return
}
