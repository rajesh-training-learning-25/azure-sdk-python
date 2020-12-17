package workloadmonitor

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

// MonitorsClient is the aPIs for workload monitoring
type MonitorsClient struct {
	BaseClient
}

// NewMonitorsClient creates an instance of the MonitorsClient client.
func NewMonitorsClient(subscriptionID string) MonitorsClient {
	return NewMonitorsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewMonitorsClientWithBaseURI creates an instance of the MonitorsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewMonitorsClientWithBaseURI(baseURI string, subscriptionID string) MonitorsClient {
	return MonitorsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get sends the get request.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// resourceNamespace - the Namespace of the resource.
// resourceType - the type of the resource.
// resourceName - name of the resource.
// monitorID - monitor Id.
func (client MonitorsClient) Get(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, monitorID string) (result Monitor, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitorsClient.Get")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("workloadmonitor.MonitorsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, resourceNamespace, resourceType, resourceName, monitorID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadmonitor.MonitorsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "workloadmonitor.MonitorsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadmonitor.MonitorsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client MonitorsClient) GetPreparer(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, monitorID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"monitorId":         autorest.Encode("path", monitorID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"resourceNamespace": autorest.Encode("path", resourceNamespace),
		"resourceType":      autorest.Encode("path", resourceType),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-08-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceNamespace}/{resourceType}/{resourceName}/providers/Microsoft.WorkloadMonitor/monitors/{monitorId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client MonitorsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client MonitorsClient) GetResponder(resp *http.Response) (result Monitor, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResource sends the list by resource request.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// resourceNamespace - the Namespace of the resource.
// resourceType - the type of the resource.
// resourceName - name of the resource.
// filter - filter to be applied on the operation.
// skiptoken - the page-continuation token to use with a paged version of this API.
func (client MonitorsClient) ListByResource(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, filter string, skiptoken string) (result MonitorsCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitorsClient.ListByResource")
		defer func() {
			sc := -1
			if result.mc.Response.Response != nil {
				sc = result.mc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("workloadmonitor.MonitorsClient", "ListByResource", err.Error())
	}

	result.fn = client.listByResourceNextResults
	req, err := client.ListByResourcePreparer(ctx, resourceGroupName, resourceNamespace, resourceType, resourceName, filter, skiptoken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadmonitor.MonitorsClient", "ListByResource", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceSender(req)
	if err != nil {
		result.mc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "workloadmonitor.MonitorsClient", "ListByResource", resp, "Failure sending request")
		return
	}

	result.mc, err = client.ListByResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadmonitor.MonitorsClient", "ListByResource", resp, "Failure responding to request")
		return
	}
	if result.mc.hasNextLink() && result.mc.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByResourcePreparer prepares the ListByResource request.
func (client MonitorsClient) ListByResourcePreparer(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, filter string, skiptoken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"resourceNamespace": autorest.Encode("path", resourceNamespace),
		"resourceType":      autorest.Encode("path", resourceType),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-08-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceNamespace}/{resourceType}/{resourceName}/providers/Microsoft.WorkloadMonitor/monitors", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceSender sends the ListByResource request. The method will close the
// http.Response Body if it receives an error.
func (client MonitorsClient) ListByResourceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceResponder handles the response to the ListByResource request. The method always
// closes the http.Response Body.
func (client MonitorsClient) ListByResourceResponder(resp *http.Response) (result MonitorsCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceNextResults retrieves the next set of results, if any.
func (client MonitorsClient) listByResourceNextResults(ctx context.Context, lastResults MonitorsCollection) (result MonitorsCollection, err error) {
	req, err := lastResults.monitorsCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "workloadmonitor.MonitorsClient", "listByResourceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "workloadmonitor.MonitorsClient", "listByResourceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadmonitor.MonitorsClient", "listByResourceNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListByResourceComplete enumerates all values, automatically crossing page boundaries as required.
func (client MonitorsClient) ListByResourceComplete(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, filter string, skiptoken string) (result MonitorsCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitorsClient.ListByResource")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResource(ctx, resourceGroupName, resourceNamespace, resourceType, resourceName, filter, skiptoken)
	return
}

// Update sends the update request.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// resourceNamespace - the Namespace of the resource.
// resourceType - the type of the resource.
// resourceName - name of the resource.
// monitorID - monitor Id.
// body - body of the Monitor PATCH object.
func (client MonitorsClient) Update(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, monitorID string, body Monitor) (result Monitor, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitorsClient.Update")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("workloadmonitor.MonitorsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, resourceNamespace, resourceType, resourceName, monitorID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadmonitor.MonitorsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "workloadmonitor.MonitorsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadmonitor.MonitorsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client MonitorsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, monitorID string, body Monitor) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"monitorId":         autorest.Encode("path", monitorID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"resourceNamespace": autorest.Encode("path", resourceNamespace),
		"resourceType":      autorest.Encode("path", resourceType),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-08-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	body.Etag = nil
	body.MonitorProperties = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceNamespace}/{resourceType}/{resourceName}/providers/Microsoft.WorkloadMonitor/monitors/{monitorId}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client MonitorsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client MonitorsClient) UpdateResponder(resp *http.Response) (result Monitor, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
