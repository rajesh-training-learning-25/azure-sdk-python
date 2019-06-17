package services

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

// QuotasClient is the these APIs allow end users to operate on Azure Machine Learning Workspace resources.
type QuotasClient struct {
	BaseClient
}

// NewQuotasClient creates an instance of the QuotasClient client.
func NewQuotasClient(subscriptionID string) QuotasClient {
	return NewQuotasClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewQuotasClientWithBaseURI creates an instance of the QuotasClient client.
func NewQuotasClientWithBaseURI(baseURI string, subscriptionID string) QuotasClient {
	return QuotasClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List gets the currently assigned Workspace Quotas based on VMFamily.
// Parameters:
// location - the location for which resource usage is queried.
func (client QuotasClient) List(ctx context.Context, location string) (result ListWorkspaceQuotasPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QuotasClient.List")
		defer func() {
			sc := -1
			if result.lwq.Response.Response != nil {
				sc = result.lwq.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: location,
			Constraints: []validation.Constraint{{Target: "location", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("services.QuotasClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "services.QuotasClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.lwq.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "services.QuotasClient", "List", resp, "Failure sending request")
		return
	}

	result.lwq, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "services.QuotasClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client QuotasClient) ListPreparer(ctx context.Context, location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearningServices/locations/{location}/Quotas", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client QuotasClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client QuotasClient) ListResponder(resp *http.Response) (result ListWorkspaceQuotas, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client QuotasClient) listNextResults(ctx context.Context, lastResults ListWorkspaceQuotas) (result ListWorkspaceQuotas, err error) {
	req, err := lastResults.listWorkspaceQuotasPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "services.QuotasClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "services.QuotasClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "services.QuotasClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client QuotasClient) ListComplete(ctx context.Context, location string) (result ListWorkspaceQuotasIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QuotasClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, location)
	return
}

// Update update quota for each VM family in workspace.
// Parameters:
// location - the location for update quota is queried.
// parameters - quota update parameters.
func (client QuotasClient) Update(ctx context.Context, location string, parameters QuotaUpdateParameters) (result QuotasUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QuotasClient.Update")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: location,
			Constraints: []validation.Constraint{{Target: "location", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("services.QuotasClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, location, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "services.QuotasClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "services.QuotasClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client QuotasClient) UpdatePreparer(ctx context.Context, location string, parameters QuotaUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.Value = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearningServices/locations/{location}/updateQuotas", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client QuotasClient) UpdateSender(req *http.Request) (future QuotasUpdateFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client QuotasClient) UpdateResponder(resp *http.Response) (result UpdateWorkspaceQuotasResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
