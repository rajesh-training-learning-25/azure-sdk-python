package hybridnetwork

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

// RoleInstancesClient is the client for the RoleInstances methods of the Hybridnetwork service.
type RoleInstancesClient struct {
	BaseClient
}

// NewRoleInstancesClient creates an instance of the RoleInstancesClient client.
func NewRoleInstancesClient(subscriptionID string) RoleInstancesClient {
	return NewRoleInstancesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRoleInstancesClientWithBaseURI creates an instance of the RoleInstancesClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewRoleInstancesClientWithBaseURI(baseURI string, subscriptionID string) RoleInstancesClient {
	return RoleInstancesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets the information of role instance of vendor network function.
// Parameters:
// locationName - the Azure region where the network function resource was created by customer.
// vendorName - the name of the vendor.
// serviceKey - the GUID for the vendor network function.
// roleInstanceName - the name of the role instance of the vendor network function.
func (client RoleInstancesClient) Get(ctx context.Context, locationName string, vendorName string, serviceKey string, roleInstanceName string) (result RoleInstance, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleInstancesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hybridnetwork.RoleInstancesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, locationName, vendorName, serviceKey, roleInstanceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridnetwork.RoleInstancesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hybridnetwork.RoleInstancesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridnetwork.RoleInstancesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client RoleInstancesClient) GetPreparer(ctx context.Context, locationName string, vendorName string, serviceKey string, roleInstanceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"locationName":     autorest.Encode("path", locationName),
		"roleInstanceName": autorest.Encode("path", roleInstanceName),
		"serviceKey":       autorest.Encode("path", serviceKey),
		"subscriptionId":   autorest.Encode("path", client.SubscriptionID),
		"vendorName":       autorest.Encode("path", vendorName),
	}

	const APIVersion = "2020-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/locations/{locationName}/vendors/{vendorName}/networkFunctions/{serviceKey}/roleInstances/{roleInstanceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RoleInstancesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RoleInstancesClient) GetResponder(resp *http.Response) (result RoleInstance, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the information of role instances of vendor network function.
// Parameters:
// locationName - the Azure region where the network function resource was created by customer.
// vendorName - the name of the vendor.
// serviceKey - the GUID for the vendor network function.
func (client RoleInstancesClient) List(ctx context.Context, locationName string, vendorName string, serviceKey string) (result NetworkFunctionRoleInstanceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleInstancesClient.List")
		defer func() {
			sc := -1
			if result.nfrilr.Response.Response != nil {
				sc = result.nfrilr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hybridnetwork.RoleInstancesClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, locationName, vendorName, serviceKey)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridnetwork.RoleInstancesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.nfrilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hybridnetwork.RoleInstancesClient", "List", resp, "Failure sending request")
		return
	}

	result.nfrilr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridnetwork.RoleInstancesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.nfrilr.hasNextLink() && result.nfrilr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client RoleInstancesClient) ListPreparer(ctx context.Context, locationName string, vendorName string, serviceKey string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"locationName":   autorest.Encode("path", locationName),
		"serviceKey":     autorest.Encode("path", serviceKey),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
		"vendorName":     autorest.Encode("path", vendorName),
	}

	const APIVersion = "2020-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/locations/{locationName}/vendors/{vendorName}/networkFunctions/{serviceKey}/roleInstances", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RoleInstancesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RoleInstancesClient) ListResponder(resp *http.Response) (result NetworkFunctionRoleInstanceListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client RoleInstancesClient) listNextResults(ctx context.Context, lastResults NetworkFunctionRoleInstanceListResult) (result NetworkFunctionRoleInstanceListResult, err error) {
	req, err := lastResults.networkFunctionRoleInstanceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "hybridnetwork.RoleInstancesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "hybridnetwork.RoleInstancesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridnetwork.RoleInstancesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client RoleInstancesClient) ListComplete(ctx context.Context, locationName string, vendorName string, serviceKey string) (result NetworkFunctionRoleInstanceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleInstancesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, locationName, vendorName, serviceKey)
	return
}

// Restart restarts a role instance of a vendor network function.
// Parameters:
// locationName - the Azure region where the network function resource was created by customer.
// vendorName - the name of the vendor.
// serviceKey - the GUID for the vendor network function.
// roleInstanceName - the name of the role instance of the vendor network function.
func (client RoleInstancesClient) Restart(ctx context.Context, locationName string, vendorName string, serviceKey string, roleInstanceName string) (result RoleInstancesRestartFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleInstancesClient.Restart")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hybridnetwork.RoleInstancesClient", "Restart", err.Error())
	}

	req, err := client.RestartPreparer(ctx, locationName, vendorName, serviceKey, roleInstanceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridnetwork.RoleInstancesClient", "Restart", nil, "Failure preparing request")
		return
	}

	result, err = client.RestartSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridnetwork.RoleInstancesClient", "Restart", nil, "Failure sending request")
		return
	}

	return
}

// RestartPreparer prepares the Restart request.
func (client RoleInstancesClient) RestartPreparer(ctx context.Context, locationName string, vendorName string, serviceKey string, roleInstanceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"locationName":     autorest.Encode("path", locationName),
		"roleInstanceName": autorest.Encode("path", roleInstanceName),
		"serviceKey":       autorest.Encode("path", serviceKey),
		"subscriptionId":   autorest.Encode("path", client.SubscriptionID),
		"vendorName":       autorest.Encode("path", vendorName),
	}

	const APIVersion = "2020-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/locations/{locationName}/vendors/{vendorName}/networkFunctions/{serviceKey}/roleInstances/{roleInstanceName}/restart", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RestartSender sends the Restart request. The method will close the
// http.Response Body if it receives an error.
func (client RoleInstancesClient) RestartSender(req *http.Request) (future RoleInstancesRestartFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client RoleInstancesClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "hybridnetwork.RoleInstancesRestartFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("hybridnetwork.RoleInstancesRestartFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// RestartResponder handles the response to the Restart request. The method always
// closes the http.Response Body.
func (client RoleInstancesClient) RestartResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Start starts a role instance of a vendor network function.
// Parameters:
// locationName - the Azure region where the network function resource was created by customer.
// vendorName - the name of the vendor.
// serviceKey - the GUID for the vendor network function.
// roleInstanceName - the name of the role instance of the vendor network function.
func (client RoleInstancesClient) Start(ctx context.Context, locationName string, vendorName string, serviceKey string, roleInstanceName string) (result RoleInstancesStartFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleInstancesClient.Start")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hybridnetwork.RoleInstancesClient", "Start", err.Error())
	}

	req, err := client.StartPreparer(ctx, locationName, vendorName, serviceKey, roleInstanceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridnetwork.RoleInstancesClient", "Start", nil, "Failure preparing request")
		return
	}

	result, err = client.StartSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridnetwork.RoleInstancesClient", "Start", nil, "Failure sending request")
		return
	}

	return
}

// StartPreparer prepares the Start request.
func (client RoleInstancesClient) StartPreparer(ctx context.Context, locationName string, vendorName string, serviceKey string, roleInstanceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"locationName":     autorest.Encode("path", locationName),
		"roleInstanceName": autorest.Encode("path", roleInstanceName),
		"serviceKey":       autorest.Encode("path", serviceKey),
		"subscriptionId":   autorest.Encode("path", client.SubscriptionID),
		"vendorName":       autorest.Encode("path", vendorName),
	}

	const APIVersion = "2020-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/locations/{locationName}/vendors/{vendorName}/networkFunctions/{serviceKey}/roleInstances/{roleInstanceName}/start", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StartSender sends the Start request. The method will close the
// http.Response Body if it receives an error.
func (client RoleInstancesClient) StartSender(req *http.Request) (future RoleInstancesStartFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client RoleInstancesClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "hybridnetwork.RoleInstancesStartFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("hybridnetwork.RoleInstancesStartFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// StartResponder handles the response to the Start request. The method always
// closes the http.Response Body.
func (client RoleInstancesClient) StartResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Stop powers off (stop) a role instance of a vendor network function.
// Parameters:
// locationName - the Azure region where the network function resource was created by customer.
// vendorName - the name of the vendor.
// serviceKey - the GUID for the vendor network function.
// roleInstanceName - the name of the role instance of the vendor network function.
func (client RoleInstancesClient) Stop(ctx context.Context, locationName string, vendorName string, serviceKey string, roleInstanceName string) (result RoleInstancesStopFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleInstancesClient.Stop")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hybridnetwork.RoleInstancesClient", "Stop", err.Error())
	}

	req, err := client.StopPreparer(ctx, locationName, vendorName, serviceKey, roleInstanceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridnetwork.RoleInstancesClient", "Stop", nil, "Failure preparing request")
		return
	}

	result, err = client.StopSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridnetwork.RoleInstancesClient", "Stop", nil, "Failure sending request")
		return
	}

	return
}

// StopPreparer prepares the Stop request.
func (client RoleInstancesClient) StopPreparer(ctx context.Context, locationName string, vendorName string, serviceKey string, roleInstanceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"locationName":     autorest.Encode("path", locationName),
		"roleInstanceName": autorest.Encode("path", roleInstanceName),
		"serviceKey":       autorest.Encode("path", serviceKey),
		"subscriptionId":   autorest.Encode("path", client.SubscriptionID),
		"vendorName":       autorest.Encode("path", vendorName),
	}

	const APIVersion = "2020-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/locations/{locationName}/vendors/{vendorName}/networkFunctions/{serviceKey}/roleInstances/{roleInstanceName}/stop", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StopSender sends the Stop request. The method will close the
// http.Response Body if it receives an error.
func (client RoleInstancesClient) StopSender(req *http.Request) (future RoleInstancesStopFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client RoleInstancesClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "hybridnetwork.RoleInstancesStopFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("hybridnetwork.RoleInstancesStopFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// StopResponder handles the response to the Stop request. The method always
// closes the http.Response Body.
func (client RoleInstancesClient) StopResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
