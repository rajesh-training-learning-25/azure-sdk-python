package storsimple

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

// AlertsClient is the client for the Alerts methods of the Storsimple service.
type AlertsClient struct {
	BaseClient
}

// NewAlertsClient creates an instance of the AlertsClient client.
func NewAlertsClient(subscriptionID string) AlertsClient {
	return NewAlertsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAlertsClientWithBaseURI creates an instance of the AlertsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewAlertsClientWithBaseURI(baseURI string, subscriptionID string) AlertsClient {
	return AlertsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Clear clear the alerts.
// Parameters:
// request - the clear alert request.
// resourceGroupName - the resource group name
// managerName - the manager name
func (client AlertsClient) Clear(ctx context.Context, request ClearAlertRequest, resourceGroupName string, managerName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AlertsClient.Clear")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: request,
			Constraints: []validation.Constraint{{Target: "request.Alerts", Name: validation.Null, Rule: true, Chain: nil}}},
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storsimple.AlertsClient", "Clear", err.Error())
	}

	req, err := client.ClearPreparer(ctx, request, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.AlertsClient", "Clear", nil, "Failure preparing request")
		return
	}

	resp, err := client.ClearSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "storsimple.AlertsClient", "Clear", resp, "Failure sending request")
		return
	}

	result, err = client.ClearResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.AlertsClient", "Clear", resp, "Failure responding to request")
		return
	}

	return
}

// ClearPreparer prepares the Clear request.
func (client AlertsClient) ClearPreparer(ctx context.Context, request ClearAlertRequest, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managerName":       autorest.Encode("path", managerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/clearAlerts", pathParameters),
		autorest.WithJSON(request),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ClearSender sends the Clear request. The method will close the
// http.Response Body if it receives an error.
func (client AlertsClient) ClearSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ClearResponder handles the response to the Clear request. The method always
// closes the http.Response Body.
func (client AlertsClient) ClearResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ListByManager retrieves all the alerts in a manager.
// Parameters:
// resourceGroupName - the resource group name
// managerName - the manager name
// filter - oData Filter options
func (client AlertsClient) ListByManager(ctx context.Context, resourceGroupName string, managerName string, filter string) (result AlertListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AlertsClient.ListByManager")
		defer func() {
			sc := -1
			if result.al.Response.Response != nil {
				sc = result.al.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storsimple.AlertsClient", "ListByManager", err.Error())
	}

	result.fn = client.listByManagerNextResults
	req, err := client.ListByManagerPreparer(ctx, resourceGroupName, managerName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.AlertsClient", "ListByManager", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByManagerSender(req)
	if err != nil {
		result.al.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storsimple.AlertsClient", "ListByManager", resp, "Failure sending request")
		return
	}

	result.al, err = client.ListByManagerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.AlertsClient", "ListByManager", resp, "Failure responding to request")
		return
	}
	if result.al.hasNextLink() && result.al.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByManagerPreparer prepares the ListByManager request.
func (client AlertsClient) ListByManagerPreparer(ctx context.Context, resourceGroupName string, managerName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managerName":       autorest.Encode("path", managerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/alerts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByManagerSender sends the ListByManager request. The method will close the
// http.Response Body if it receives an error.
func (client AlertsClient) ListByManagerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByManagerResponder handles the response to the ListByManager request. The method always
// closes the http.Response Body.
func (client AlertsClient) ListByManagerResponder(resp *http.Response) (result AlertList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByManagerNextResults retrieves the next set of results, if any.
func (client AlertsClient) listByManagerNextResults(ctx context.Context, lastResults AlertList) (result AlertList, err error) {
	req, err := lastResults.alertListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "storsimple.AlertsClient", "listByManagerNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByManagerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "storsimple.AlertsClient", "listByManagerNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByManagerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.AlertsClient", "listByManagerNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListByManagerComplete enumerates all values, automatically crossing page boundaries as required.
func (client AlertsClient) ListByManagerComplete(ctx context.Context, resourceGroupName string, managerName string, filter string) (result AlertListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AlertsClient.ListByManager")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByManager(ctx, resourceGroupName, managerName, filter)
	return
}

// SendTestEmail sends a test alert email.
// Parameters:
// deviceName - the device name.
// request - the send test alert email request.
// resourceGroupName - the resource group name
// managerName - the manager name
func (client AlertsClient) SendTestEmail(ctx context.Context, deviceName string, request SendTestAlertEmailRequest, resourceGroupName string, managerName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AlertsClient.SendTestEmail")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: request,
			Constraints: []validation.Constraint{{Target: "request.EmailList", Name: validation.Null, Rule: true, Chain: nil}}},
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storsimple.AlertsClient", "SendTestEmail", err.Error())
	}

	req, err := client.SendTestEmailPreparer(ctx, deviceName, request, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.AlertsClient", "SendTestEmail", nil, "Failure preparing request")
		return
	}

	resp, err := client.SendTestEmailSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "storsimple.AlertsClient", "SendTestEmail", resp, "Failure sending request")
		return
	}

	result, err = client.SendTestEmailResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.AlertsClient", "SendTestEmail", resp, "Failure responding to request")
		return
	}

	return
}

// SendTestEmailPreparer prepares the SendTestEmail request.
func (client AlertsClient) SendTestEmailPreparer(ctx context.Context, deviceName string, request SendTestAlertEmailRequest, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        autorest.Encode("path", deviceName),
		"managerName":       autorest.Encode("path", managerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/sendTestAlertEmail", pathParameters),
		autorest.WithJSON(request),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SendTestEmailSender sends the SendTestEmail request. The method will close the
// http.Response Body if it receives an error.
func (client AlertsClient) SendTestEmailSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// SendTestEmailResponder handles the response to the SendTestEmail request. The method always
// closes the http.Response Body.
func (client AlertsClient) SendTestEmailResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
