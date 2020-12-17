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

// StreamingEndpointsClient is the client for the StreamingEndpoints methods of the Media service.
type StreamingEndpointsClient struct {
	BaseClient
}

// NewStreamingEndpointsClient creates an instance of the StreamingEndpointsClient client.
func NewStreamingEndpointsClient(subscriptionID string) StreamingEndpointsClient {
	return NewStreamingEndpointsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewStreamingEndpointsClientWithBaseURI creates an instance of the StreamingEndpointsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewStreamingEndpointsClientWithBaseURI(baseURI string, subscriptionID string) StreamingEndpointsClient {
	return StreamingEndpointsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates a StreamingEndpoint.
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// streamingEndpointName - the name of the StreamingEndpoint.
// parameters - streamingEndpoint properties needed for creation.
// autoStart - the flag indicates if the resource should be automatically started on creation.
func (client StreamingEndpointsClient) Create(ctx context.Context, resourceGroupName string, accountName string, streamingEndpointName string, parameters StreamingEndpoint, autoStart *bool) (result StreamingEndpointsCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StreamingEndpointsClient.Create")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: streamingEndpointName,
			Constraints: []validation.Constraint{{Target: "streamingEndpointName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "streamingEndpointName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "streamingEndpointName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-*[a-zA-Z0-9])*$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.StreamingEndpointProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.StreamingEndpointProperties.ScaleUnits", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("media.StreamingEndpointsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, accountName, streamingEndpointName, parameters, autoStart)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingEndpointsClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingEndpointsClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client StreamingEndpointsClient) CreatePreparer(ctx context.Context, resourceGroupName string, accountName string, streamingEndpointName string, parameters StreamingEndpoint, autoStart *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":           autorest.Encode("path", accountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"streamingEndpointName": autorest.Encode("path", streamingEndpointName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01-preview"
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/streamingEndpoints/{streamingEndpointName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client StreamingEndpointsClient) CreateSender(req *http.Request) (future StreamingEndpointsCreateFuture, err error) {
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
func (client StreamingEndpointsClient) CreateResponder(resp *http.Response) (result StreamingEndpoint, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a StreamingEndpoint.
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// streamingEndpointName - the name of the StreamingEndpoint.
func (client StreamingEndpointsClient) Delete(ctx context.Context, resourceGroupName string, accountName string, streamingEndpointName string) (result StreamingEndpointsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StreamingEndpointsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: streamingEndpointName,
			Constraints: []validation.Constraint{{Target: "streamingEndpointName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "streamingEndpointName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "streamingEndpointName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-*[a-zA-Z0-9])*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("media.StreamingEndpointsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, accountName, streamingEndpointName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingEndpointsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingEndpointsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client StreamingEndpointsClient) DeletePreparer(ctx context.Context, resourceGroupName string, accountName string, streamingEndpointName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":           autorest.Encode("path", accountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"streamingEndpointName": autorest.Encode("path", streamingEndpointName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/streamingEndpoints/{streamingEndpointName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client StreamingEndpointsClient) DeleteSender(req *http.Request) (future StreamingEndpointsDeleteFuture, err error) {
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
func (client StreamingEndpointsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a StreamingEndpoint.
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// streamingEndpointName - the name of the StreamingEndpoint.
func (client StreamingEndpointsClient) Get(ctx context.Context, resourceGroupName string, accountName string, streamingEndpointName string) (result StreamingEndpoint, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StreamingEndpointsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: streamingEndpointName,
			Constraints: []validation.Constraint{{Target: "streamingEndpointName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "streamingEndpointName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "streamingEndpointName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-*[a-zA-Z0-9])*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("media.StreamingEndpointsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, accountName, streamingEndpointName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingEndpointsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "media.StreamingEndpointsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingEndpointsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client StreamingEndpointsClient) GetPreparer(ctx context.Context, resourceGroupName string, accountName string, streamingEndpointName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":           autorest.Encode("path", accountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"streamingEndpointName": autorest.Encode("path", streamingEndpointName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/streamingEndpoints/{streamingEndpointName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client StreamingEndpointsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client StreamingEndpointsClient) GetResponder(resp *http.Response) (result StreamingEndpoint, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the StreamingEndpoints in the account.
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
func (client StreamingEndpointsClient) List(ctx context.Context, resourceGroupName string, accountName string) (result StreamingEndpointListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StreamingEndpointsClient.List")
		defer func() {
			sc := -1
			if result.selr.Response.Response != nil {
				sc = result.selr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, accountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingEndpointsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.selr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "media.StreamingEndpointsClient", "List", resp, "Failure sending request")
		return
	}

	result.selr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingEndpointsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.selr.hasNextLink() && result.selr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client StreamingEndpointsClient) ListPreparer(ctx context.Context, resourceGroupName string, accountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/streamingEndpoints", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client StreamingEndpointsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client StreamingEndpointsClient) ListResponder(resp *http.Response) (result StreamingEndpointListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client StreamingEndpointsClient) listNextResults(ctx context.Context, lastResults StreamingEndpointListResult) (result StreamingEndpointListResult, err error) {
	req, err := lastResults.streamingEndpointListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "media.StreamingEndpointsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "media.StreamingEndpointsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingEndpointsClient", "listNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client StreamingEndpointsClient) ListComplete(ctx context.Context, resourceGroupName string, accountName string) (result StreamingEndpointListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StreamingEndpointsClient.List")
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

// Scale scales an existing StreamingEndpoint.
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// streamingEndpointName - the name of the StreamingEndpoint.
// parameters - streamingEndpoint scale parameters
func (client StreamingEndpointsClient) Scale(ctx context.Context, resourceGroupName string, accountName string, streamingEndpointName string, parameters StreamingEntityScaleUnit) (result StreamingEndpointsScaleFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StreamingEndpointsClient.Scale")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: streamingEndpointName,
			Constraints: []validation.Constraint{{Target: "streamingEndpointName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "streamingEndpointName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "streamingEndpointName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-*[a-zA-Z0-9])*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("media.StreamingEndpointsClient", "Scale", err.Error())
	}

	req, err := client.ScalePreparer(ctx, resourceGroupName, accountName, streamingEndpointName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingEndpointsClient", "Scale", nil, "Failure preparing request")
		return
	}

	result, err = client.ScaleSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingEndpointsClient", "Scale", result.Response(), "Failure sending request")
		return
	}

	return
}

// ScalePreparer prepares the Scale request.
func (client StreamingEndpointsClient) ScalePreparer(ctx context.Context, resourceGroupName string, accountName string, streamingEndpointName string, parameters StreamingEntityScaleUnit) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":           autorest.Encode("path", accountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"streamingEndpointName": autorest.Encode("path", streamingEndpointName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/streamingEndpoints/{streamingEndpointName}/scale", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ScaleSender sends the Scale request. The method will close the
// http.Response Body if it receives an error.
func (client StreamingEndpointsClient) ScaleSender(req *http.Request) (future StreamingEndpointsScaleFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// ScaleResponder handles the response to the Scale request. The method always
// closes the http.Response Body.
func (client StreamingEndpointsClient) ScaleResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Start starts an existing StreamingEndpoint.
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// streamingEndpointName - the name of the StreamingEndpoint.
func (client StreamingEndpointsClient) Start(ctx context.Context, resourceGroupName string, accountName string, streamingEndpointName string) (result StreamingEndpointsStartFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StreamingEndpointsClient.Start")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: streamingEndpointName,
			Constraints: []validation.Constraint{{Target: "streamingEndpointName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "streamingEndpointName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "streamingEndpointName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-*[a-zA-Z0-9])*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("media.StreamingEndpointsClient", "Start", err.Error())
	}

	req, err := client.StartPreparer(ctx, resourceGroupName, accountName, streamingEndpointName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingEndpointsClient", "Start", nil, "Failure preparing request")
		return
	}

	result, err = client.StartSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingEndpointsClient", "Start", result.Response(), "Failure sending request")
		return
	}

	return
}

// StartPreparer prepares the Start request.
func (client StreamingEndpointsClient) StartPreparer(ctx context.Context, resourceGroupName string, accountName string, streamingEndpointName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":           autorest.Encode("path", accountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"streamingEndpointName": autorest.Encode("path", streamingEndpointName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/streamingEndpoints/{streamingEndpointName}/start", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StartSender sends the Start request. The method will close the
// http.Response Body if it receives an error.
func (client StreamingEndpointsClient) StartSender(req *http.Request) (future StreamingEndpointsStartFuture, err error) {
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
func (client StreamingEndpointsClient) StartResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Stop stops an existing StreamingEndpoint.
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// streamingEndpointName - the name of the StreamingEndpoint.
func (client StreamingEndpointsClient) Stop(ctx context.Context, resourceGroupName string, accountName string, streamingEndpointName string) (result StreamingEndpointsStopFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StreamingEndpointsClient.Stop")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: streamingEndpointName,
			Constraints: []validation.Constraint{{Target: "streamingEndpointName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "streamingEndpointName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "streamingEndpointName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-*[a-zA-Z0-9])*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("media.StreamingEndpointsClient", "Stop", err.Error())
	}

	req, err := client.StopPreparer(ctx, resourceGroupName, accountName, streamingEndpointName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingEndpointsClient", "Stop", nil, "Failure preparing request")
		return
	}

	result, err = client.StopSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingEndpointsClient", "Stop", result.Response(), "Failure sending request")
		return
	}

	return
}

// StopPreparer prepares the Stop request.
func (client StreamingEndpointsClient) StopPreparer(ctx context.Context, resourceGroupName string, accountName string, streamingEndpointName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":           autorest.Encode("path", accountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"streamingEndpointName": autorest.Encode("path", streamingEndpointName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/streamingEndpoints/{streamingEndpointName}/stop", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StopSender sends the Stop request. The method will close the
// http.Response Body if it receives an error.
func (client StreamingEndpointsClient) StopSender(req *http.Request) (future StreamingEndpointsStopFuture, err error) {
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
func (client StreamingEndpointsClient) StopResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update updates a existing StreamingEndpoint.
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// streamingEndpointName - the name of the StreamingEndpoint.
// parameters - streamingEndpoint properties needed for creation.
func (client StreamingEndpointsClient) Update(ctx context.Context, resourceGroupName string, accountName string, streamingEndpointName string, parameters StreamingEndpoint) (result StreamingEndpointsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StreamingEndpointsClient.Update")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: streamingEndpointName,
			Constraints: []validation.Constraint{{Target: "streamingEndpointName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "streamingEndpointName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "streamingEndpointName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-*[a-zA-Z0-9])*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("media.StreamingEndpointsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, accountName, streamingEndpointName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingEndpointsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingEndpointsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client StreamingEndpointsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, accountName string, streamingEndpointName string, parameters StreamingEndpoint) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":           autorest.Encode("path", accountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"streamingEndpointName": autorest.Encode("path", streamingEndpointName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/streamingEndpoints/{streamingEndpointName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client StreamingEndpointsClient) UpdateSender(req *http.Request) (future StreamingEndpointsUpdateFuture, err error) {
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
func (client StreamingEndpointsClient) UpdateResponder(resp *http.Response) (result StreamingEndpoint, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
