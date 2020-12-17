package securityinsight

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

// IncidentsClient is the API spec for Microsoft.SecurityInsights (Azure Security Insights) resource provider
type IncidentsClient struct {
	BaseClient
}

// NewIncidentsClient creates an instance of the IncidentsClient client.
func NewIncidentsClient(subscriptionID string) IncidentsClient {
	return NewIncidentsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewIncidentsClientWithBaseURI creates an instance of the IncidentsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewIncidentsClientWithBaseURI(baseURI string, subscriptionID string) IncidentsClient {
	return IncidentsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates the incident.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// workspaceName - the name of the workspace.
// incidentID - incident ID
// incident - the incident
func (client IncidentsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, incident Incident) (result Incident, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IncidentsClient.CreateOrUpdate")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: incident,
			Constraints: []validation.Constraint{{Target: "incident.IncidentProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "incident.IncidentProperties.Title", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("securityinsight.IncidentsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, workspaceName, incidentID, incident)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client IncidentsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, incident Incident) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"incidentId":        autorest.Encode("path", incidentID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}", pathParameters),
		autorest.WithJSON(incident),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client IncidentsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client IncidentsClient) CreateOrUpdateResponder(resp *http.Response) (result Incident, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete the incident.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// workspaceName - the name of the workspace.
// incidentID - incident ID
func (client IncidentsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IncidentsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.IncidentsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, workspaceName, incidentID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client IncidentsClient) DeletePreparer(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"incidentId":        autorest.Encode("path", incidentID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client IncidentsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client IncidentsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets an incident.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// workspaceName - the name of the workspace.
// incidentID - incident ID
func (client IncidentsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string) (result Incident, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IncidentsClient.Get")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.IncidentsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName, incidentID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client IncidentsClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"incidentId":        autorest.Encode("path", incidentID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client IncidentsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client IncidentsClient) GetResponder(resp *http.Response) (result Incident, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all incidents.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// workspaceName - the name of the workspace.
// filter - filters the results, based on a Boolean condition. Optional.
// orderby - sorts the results. Optional.
// top - returns only the first n results. Optional.
// skipToken - skiptoken is only used if a previous operation returned a partial result. If a previous response
// contains a nextLink element, the value of the nextLink element will include a skiptoken parameter that
// specifies a starting point to use for subsequent calls. Optional.
func (client IncidentsClient) List(ctx context.Context, resourceGroupName string, workspaceName string, filter string, orderby string, top *int32, skipToken string) (result IncidentListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IncidentsClient.List")
		defer func() {
			sc := -1
			if result.il.Response.Response != nil {
				sc = result.il.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.IncidentsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, workspaceName, filter, orderby, top, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.il.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "List", resp, "Failure sending request")
		return
	}

	result.il, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.il.hasNextLink() && result.il.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client IncidentsClient) ListPreparer(ctx context.Context, resourceGroupName string, workspaceName string, filter string, orderby string, top *int32, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client IncidentsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client IncidentsClient) ListResponder(resp *http.Response) (result IncidentList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client IncidentsClient) listNextResults(ctx context.Context, lastResults IncidentList) (result IncidentList, err error) {
	req, err := lastResults.incidentListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.IncidentsClient", "listNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client IncidentsClient) ListComplete(ctx context.Context, resourceGroupName string, workspaceName string, filter string, orderby string, top *int32, skipToken string) (result IncidentListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IncidentsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, workspaceName, filter, orderby, top, skipToken)
	return
}
