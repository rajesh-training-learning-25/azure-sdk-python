package timeseriesinsights

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

// ReferenceDataSetsClient is the time Series Insights client
type ReferenceDataSetsClient struct {
	BaseClient
}

// NewReferenceDataSetsClient creates an instance of the ReferenceDataSetsClient client.
func NewReferenceDataSetsClient(subscriptionID string) ReferenceDataSetsClient {
	return NewReferenceDataSetsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewReferenceDataSetsClientWithBaseURI creates an instance of the ReferenceDataSetsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewReferenceDataSetsClientWithBaseURI(baseURI string, subscriptionID string) ReferenceDataSetsClient {
	return ReferenceDataSetsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update a reference data set in the specified environment.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// environmentName - the name of the Time Series Insights environment associated with the specified resource
// group.
// referenceDataSetName - name of the reference data set.
// parameters - parameters for creating a reference data set.
func (client ReferenceDataSetsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, environmentName string, referenceDataSetName string, parameters ReferenceDataSetCreateOrUpdateParameters) (result ReferenceDataSetResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReferenceDataSetsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: referenceDataSetName,
			Constraints: []validation.Constraint{{Target: "referenceDataSetName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "referenceDataSetName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "referenceDataSetName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ReferenceDataSetCreationProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.ReferenceDataSetCreationProperties.KeyProperties", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("timeseriesinsights.ReferenceDataSetsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, environmentName, referenceDataSetName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "timeseriesinsights.ReferenceDataSetsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "timeseriesinsights.ReferenceDataSetsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "timeseriesinsights.ReferenceDataSetsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ReferenceDataSetsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, environmentName string, referenceDataSetName string, parameters ReferenceDataSetCreateOrUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"environmentName":      autorest.Encode("path", environmentName),
		"referenceDataSetName": autorest.Encode("path", referenceDataSetName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-02-28-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/referenceDataSets/{referenceDataSetName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ReferenceDataSetsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ReferenceDataSetsClient) CreateOrUpdateResponder(resp *http.Response) (result ReferenceDataSetResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the reference data set with the specified name in the specified subscription, resource group, and
// environment
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// environmentName - the name of the Time Series Insights environment associated with the specified resource
// group.
// referenceDataSetName - the name of the Time Series Insights reference data set associated with the specified
// environment.
func (client ReferenceDataSetsClient) Delete(ctx context.Context, resourceGroupName string, environmentName string, referenceDataSetName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReferenceDataSetsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, environmentName, referenceDataSetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "timeseriesinsights.ReferenceDataSetsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "timeseriesinsights.ReferenceDataSetsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "timeseriesinsights.ReferenceDataSetsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ReferenceDataSetsClient) DeletePreparer(ctx context.Context, resourceGroupName string, environmentName string, referenceDataSetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"environmentName":      autorest.Encode("path", environmentName),
		"referenceDataSetName": autorest.Encode("path", referenceDataSetName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-02-28-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/referenceDataSets/{referenceDataSetName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ReferenceDataSetsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ReferenceDataSetsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the reference data set with the specified name in the specified environment.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// environmentName - the name of the Time Series Insights environment associated with the specified resource
// group.
// referenceDataSetName - the name of the Time Series Insights reference data set associated with the specified
// environment.
func (client ReferenceDataSetsClient) Get(ctx context.Context, resourceGroupName string, environmentName string, referenceDataSetName string) (result ReferenceDataSetResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReferenceDataSetsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, environmentName, referenceDataSetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "timeseriesinsights.ReferenceDataSetsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "timeseriesinsights.ReferenceDataSetsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "timeseriesinsights.ReferenceDataSetsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ReferenceDataSetsClient) GetPreparer(ctx context.Context, resourceGroupName string, environmentName string, referenceDataSetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"environmentName":      autorest.Encode("path", environmentName),
		"referenceDataSetName": autorest.Encode("path", referenceDataSetName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-02-28-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/referenceDataSets/{referenceDataSetName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ReferenceDataSetsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ReferenceDataSetsClient) GetResponder(resp *http.Response) (result ReferenceDataSetResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByEnvironment lists all the available reference data sets associated with the subscription and within the
// specified resource group and environment.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// environmentName - the name of the Time Series Insights environment associated with the specified resource
// group.
func (client ReferenceDataSetsClient) ListByEnvironment(ctx context.Context, resourceGroupName string, environmentName string) (result ReferenceDataSetListResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReferenceDataSetsClient.ListByEnvironment")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByEnvironmentPreparer(ctx, resourceGroupName, environmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "timeseriesinsights.ReferenceDataSetsClient", "ListByEnvironment", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByEnvironmentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "timeseriesinsights.ReferenceDataSetsClient", "ListByEnvironment", resp, "Failure sending request")
		return
	}

	result, err = client.ListByEnvironmentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "timeseriesinsights.ReferenceDataSetsClient", "ListByEnvironment", resp, "Failure responding to request")
	}

	return
}

// ListByEnvironmentPreparer prepares the ListByEnvironment request.
func (client ReferenceDataSetsClient) ListByEnvironmentPreparer(ctx context.Context, resourceGroupName string, environmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"environmentName":   autorest.Encode("path", environmentName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-02-28-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/referenceDataSets", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByEnvironmentSender sends the ListByEnvironment request. The method will close the
// http.Response Body if it receives an error.
func (client ReferenceDataSetsClient) ListByEnvironmentSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByEnvironmentResponder handles the response to the ListByEnvironment request. The method always
// closes the http.Response Body.
func (client ReferenceDataSetsClient) ListByEnvironmentResponder(resp *http.Response) (result ReferenceDataSetListResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates the reference data set with the specified name in the specified subscription, resource group, and
// environment.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// environmentName - the name of the Time Series Insights environment associated with the specified resource
// group.
// referenceDataSetName - the name of the Time Series Insights reference data set associated with the specified
// environment.
// referenceDataSetUpdateParameters - request object that contains the updated information for the reference
// data set.
func (client ReferenceDataSetsClient) Update(ctx context.Context, resourceGroupName string, environmentName string, referenceDataSetName string, referenceDataSetUpdateParameters ReferenceDataSetUpdateParameters) (result ReferenceDataSetResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReferenceDataSetsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, environmentName, referenceDataSetName, referenceDataSetUpdateParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "timeseriesinsights.ReferenceDataSetsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "timeseriesinsights.ReferenceDataSetsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "timeseriesinsights.ReferenceDataSetsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ReferenceDataSetsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, environmentName string, referenceDataSetName string, referenceDataSetUpdateParameters ReferenceDataSetUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"environmentName":      autorest.Encode("path", environmentName),
		"referenceDataSetName": autorest.Encode("path", referenceDataSetName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-02-28-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/referenceDataSets/{referenceDataSetName}", pathParameters),
		autorest.WithJSON(referenceDataSetUpdateParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ReferenceDataSetsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ReferenceDataSetsClient) UpdateResponder(resp *http.Response) (result ReferenceDataSetResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
