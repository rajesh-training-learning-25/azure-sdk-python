package media

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

// LiveEventsClient is the client for the LiveEvents methods of the Media service.
type LiveEventsClient struct {
	BaseClient
}

// NewLiveEventsClient creates an instance of the LiveEventsClient client.
func NewLiveEventsClient(subscriptionID string) LiveEventsClient {
	return NewLiveEventsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLiveEventsClientWithBaseURI creates an instance of the LiveEventsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewLiveEventsClientWithBaseURI(baseURI string, subscriptionID string) LiveEventsClient {
	return LiveEventsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates a Live Event.
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// liveEventName - the name of the Live Event.
// parameters - live Event properties needed for creation.
// autoStart - the flag indicates if auto start the Live Event.
func (client LiveEventsClient) Create(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, parameters LiveEvent, autoStart *bool) (result LiveEventsCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LiveEventsClient.Create")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: liveEventName,
			Constraints: []validation.Constraint{{Target: "liveEventName", Name: validation.MaxLength, Rule: 32, Chain: nil},
				{Target: "liveEventName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "liveEventName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-*[a-zA-Z0-9])*$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.LiveEventProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.LiveEventProperties.Input", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("media.LiveEventsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, accountName, liveEventName, parameters, autoStart)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.LiveEventsClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.LiveEventsClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client LiveEventsClient) CreatePreparer(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, parameters LiveEvent, autoStart *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"liveEventName":     autorest.Encode("path", liveEventName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if autoStart != nil {
		queryParameters["autoStart"] = autorest.Encode("query", *autoStart)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client LiveEventsClient) CreateSender(req *http.Request) (future LiveEventsCreateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client LiveEventsClient) CreateResponder(resp *http.Response) (result LiveEvent, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a Live Event.
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// liveEventName - the name of the Live Event.
func (client LiveEventsClient) Delete(ctx context.Context, resourceGroupName string, accountName string, liveEventName string) (result LiveEventsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LiveEventsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: liveEventName,
			Constraints: []validation.Constraint{{Target: "liveEventName", Name: validation.MaxLength, Rule: 32, Chain: nil},
				{Target: "liveEventName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "liveEventName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-*[a-zA-Z0-9])*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("media.LiveEventsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, accountName, liveEventName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.LiveEventsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.LiveEventsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client LiveEventsClient) DeletePreparer(ctx context.Context, resourceGroupName string, accountName string, liveEventName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"liveEventName":     autorest.Encode("path", liveEventName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client LiveEventsClient) DeleteSender(req *http.Request) (future LiveEventsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client LiveEventsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a Live Event.
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// liveEventName - the name of the Live Event.
func (client LiveEventsClient) Get(ctx context.Context, resourceGroupName string, accountName string, liveEventName string) (result LiveEvent, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LiveEventsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: liveEventName,
			Constraints: []validation.Constraint{{Target: "liveEventName", Name: validation.MaxLength, Rule: 32, Chain: nil},
				{Target: "liveEventName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "liveEventName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-*[a-zA-Z0-9])*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("media.LiveEventsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, accountName, liveEventName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.LiveEventsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "media.LiveEventsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.LiveEventsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client LiveEventsClient) GetPreparer(ctx context.Context, resourceGroupName string, accountName string, liveEventName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"liveEventName":     autorest.Encode("path", liveEventName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client LiveEventsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client LiveEventsClient) GetResponder(resp *http.Response) (result LiveEvent, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the Live Events in the account.
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
func (client LiveEventsClient) List(ctx context.Context, resourceGroupName string, accountName string) (result LiveEventListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LiveEventsClient.List")
		defer func() {
			sc := -1
			if result.lelr.Response.Response != nil {
				sc = result.lelr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, accountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.LiveEventsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.lelr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "media.LiveEventsClient", "List", resp, "Failure sending request")
		return
	}

	result.lelr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.LiveEventsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.lelr.hasNextLink() && result.lelr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client LiveEventsClient) ListPreparer(ctx context.Context, resourceGroupName string, accountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client LiveEventsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client LiveEventsClient) ListResponder(resp *http.Response) (result LiveEventListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client LiveEventsClient) listNextResults(ctx context.Context, lastResults LiveEventListResult) (result LiveEventListResult, err error) {
	req, err := lastResults.liveEventListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "media.LiveEventsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "media.LiveEventsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.LiveEventsClient", "listNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client LiveEventsClient) ListComplete(ctx context.Context, resourceGroupName string, accountName string) (result LiveEventListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LiveEventsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, accountName)
	return
}

// Reset resets an existing Live Event.
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// liveEventName - the name of the Live Event.
func (client LiveEventsClient) Reset(ctx context.Context, resourceGroupName string, accountName string, liveEventName string) (result LiveEventsResetFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LiveEventsClient.Reset")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: liveEventName,
			Constraints: []validation.Constraint{{Target: "liveEventName", Name: validation.MaxLength, Rule: 32, Chain: nil},
				{Target: "liveEventName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "liveEventName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-*[a-zA-Z0-9])*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("media.LiveEventsClient", "Reset", err.Error())
	}

	req, err := client.ResetPreparer(ctx, resourceGroupName, accountName, liveEventName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.LiveEventsClient", "Reset", nil, "Failure preparing request")
		return
	}

	result, err = client.ResetSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.LiveEventsClient", "Reset", result.Response(), "Failure sending request")
		return
	}

	return
}

// ResetPreparer prepares the Reset request.
func (client LiveEventsClient) ResetPreparer(ctx context.Context, resourceGroupName string, accountName string, liveEventName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"liveEventName":     autorest.Encode("path", liveEventName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/reset", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ResetSender sends the Reset request. The method will close the
// http.Response Body if it receives an error.
func (client LiveEventsClient) ResetSender(req *http.Request) (future LiveEventsResetFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// ResetResponder handles the response to the Reset request. The method always
// closes the http.Response Body.
func (client LiveEventsClient) ResetResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Start starts an existing Live Event.
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// liveEventName - the name of the Live Event.
func (client LiveEventsClient) Start(ctx context.Context, resourceGroupName string, accountName string, liveEventName string) (result LiveEventsStartFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LiveEventsClient.Start")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: liveEventName,
			Constraints: []validation.Constraint{{Target: "liveEventName", Name: validation.MaxLength, Rule: 32, Chain: nil},
				{Target: "liveEventName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "liveEventName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-*[a-zA-Z0-9])*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("media.LiveEventsClient", "Start", err.Error())
	}

	req, err := client.StartPreparer(ctx, resourceGroupName, accountName, liveEventName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.LiveEventsClient", "Start", nil, "Failure preparing request")
		return
	}

	result, err = client.StartSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.LiveEventsClient", "Start", result.Response(), "Failure sending request")
		return
	}

	return
}

// StartPreparer prepares the Start request.
func (client LiveEventsClient) StartPreparer(ctx context.Context, resourceGroupName string, accountName string, liveEventName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"liveEventName":     autorest.Encode("path", liveEventName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/start", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StartSender sends the Start request. The method will close the
// http.Response Body if it receives an error.
func (client LiveEventsClient) StartSender(req *http.Request) (future LiveEventsStartFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// StartResponder handles the response to the Start request. The method always
// closes the http.Response Body.
func (client LiveEventsClient) StartResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Stop stops an existing Live Event.
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// liveEventName - the name of the Live Event.
// parameters - liveEvent stop parameters
func (client LiveEventsClient) Stop(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, parameters LiveEventActionInput) (result LiveEventsStopFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LiveEventsClient.Stop")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: liveEventName,
			Constraints: []validation.Constraint{{Target: "liveEventName", Name: validation.MaxLength, Rule: 32, Chain: nil},
				{Target: "liveEventName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "liveEventName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-*[a-zA-Z0-9])*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("media.LiveEventsClient", "Stop", err.Error())
	}

	req, err := client.StopPreparer(ctx, resourceGroupName, accountName, liveEventName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.LiveEventsClient", "Stop", nil, "Failure preparing request")
		return
	}

	result, err = client.StopSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.LiveEventsClient", "Stop", result.Response(), "Failure sending request")
		return
	}

	return
}

// StopPreparer prepares the Stop request.
func (client LiveEventsClient) StopPreparer(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, parameters LiveEventActionInput) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"liveEventName":     autorest.Encode("path", liveEventName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/stop", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StopSender sends the Stop request. The method will close the
// http.Response Body if it receives an error.
func (client LiveEventsClient) StopSender(req *http.Request) (future LiveEventsStopFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// StopResponder handles the response to the Stop request. The method always
// closes the http.Response Body.
func (client LiveEventsClient) StopResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update updates a existing Live Event.
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// liveEventName - the name of the Live Event.
// parameters - live Event properties needed for creation.
func (client LiveEventsClient) Update(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, parameters LiveEvent) (result LiveEventsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LiveEventsClient.Update")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: liveEventName,
			Constraints: []validation.Constraint{{Target: "liveEventName", Name: validation.MaxLength, Rule: 32, Chain: nil},
				{Target: "liveEventName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "liveEventName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-*[a-zA-Z0-9])*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("media.LiveEventsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, accountName, liveEventName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.LiveEventsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.LiveEventsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client LiveEventsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, parameters LiveEvent) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"liveEventName":     autorest.Encode("path", liveEventName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client LiveEventsClient) UpdateSender(req *http.Request) (future LiveEventsUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client LiveEventsClient) UpdateResponder(resp *http.Response) (result LiveEvent, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
