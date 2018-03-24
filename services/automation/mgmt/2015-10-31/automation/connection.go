package automation

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
	"net/http"
)

// ConnectionClient is the automation Client
type ConnectionClient struct {
	BaseClient
}

// NewConnectionClient creates an instance of the ConnectionClient client.
func NewConnectionClient(subscriptionID string, resourceGroupName string) ConnectionClient {
	return NewConnectionClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName)
}

// NewConnectionClientWithBaseURI creates an instance of the ConnectionClient client.
func NewConnectionClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string) ConnectionClient {
	return ConnectionClient{NewWithBaseURI(baseURI, subscriptionID, resourceGroupName)}
}

// CreateOrUpdate create or update a connection.
//
// automationAccountName is the automation account name. connectionName is the parameters supplied to the create or
// update connection operation. parameters is the parameters supplied to the create or update connection operation.
func (client ConnectionClient) CreateOrUpdate(ctx context.Context, automationAccountName string, connectionName string, parameters ConnectionCreateOrUpdateParameters) (result Connection, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.ResourceGroupName,
			Constraints: []validation.Constraint{{Target: "client.ResourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "client.ResourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "client.ResourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Name", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.ConnectionCreateOrUpdateProperties", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "parameters.ConnectionCreateOrUpdateProperties.ConnectionType", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("automation.ConnectionClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, automationAccountName, connectionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.ConnectionClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.ConnectionClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.ConnectionClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ConnectionClient) CreateOrUpdatePreparer(ctx context.Context, automationAccountName string, connectionName string, parameters ConnectionCreateOrUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"connectionName":        autorest.Encode("path", connectionName),
		"resourceGroupName":     autorest.Encode("path", client.ResourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/connections/{connectionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ConnectionClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ConnectionClient) CreateOrUpdateResponder(resp *http.Response) (result Connection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete the connection.
//
// automationAccountName is the automation account name. connectionName is the name of connection.
func (client ConnectionClient) Delete(ctx context.Context, automationAccountName string, connectionName string) (result Connection, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.ResourceGroupName,
			Constraints: []validation.Constraint{{Target: "client.ResourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "client.ResourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "client.ResourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.ConnectionClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, automationAccountName, connectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.ConnectionClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.ConnectionClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.ConnectionClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ConnectionClient) DeletePreparer(ctx context.Context, automationAccountName string, connectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"connectionName":        autorest.Encode("path", connectionName),
		"resourceGroupName":     autorest.Encode("path", client.ResourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/connections/{connectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ConnectionClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ConnectionClient) DeleteResponder(resp *http.Response) (result Connection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get retrieve the connection identified by connection name.
//
// automationAccountName is the automation account name. connectionName is the name of connection.
func (client ConnectionClient) Get(ctx context.Context, automationAccountName string, connectionName string) (result Connection, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.ResourceGroupName,
			Constraints: []validation.Constraint{{Target: "client.ResourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "client.ResourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "client.ResourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.ConnectionClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, automationAccountName, connectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.ConnectionClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.ConnectionClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.ConnectionClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ConnectionClient) GetPreparer(ctx context.Context, automationAccountName string, connectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"connectionName":        autorest.Encode("path", connectionName),
		"resourceGroupName":     autorest.Encode("path", client.ResourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/connections/{connectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ConnectionClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ConnectionClient) GetResponder(resp *http.Response) (result Connection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByAutomationAccount retrieve a list of connections.
//
// automationAccountName is the automation account name.
func (client ConnectionClient) ListByAutomationAccount(ctx context.Context, automationAccountName string) (result ConnectionListResultPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.ResourceGroupName,
			Constraints: []validation.Constraint{{Target: "client.ResourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "client.ResourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "client.ResourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.ConnectionClient", "ListByAutomationAccount", err.Error())
	}

	result.fn = client.listByAutomationAccountNextResults
	req, err := client.ListByAutomationAccountPreparer(ctx, automationAccountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.ConnectionClient", "ListByAutomationAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByAutomationAccountSender(req)
	if err != nil {
		result.clr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.ConnectionClient", "ListByAutomationAccount", resp, "Failure sending request")
		return
	}

	result.clr, err = client.ListByAutomationAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.ConnectionClient", "ListByAutomationAccount", resp, "Failure responding to request")
	}

	return
}

// ListByAutomationAccountPreparer prepares the ListByAutomationAccount request.
func (client ConnectionClient) ListByAutomationAccountPreparer(ctx context.Context, automationAccountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", client.ResourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/connections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByAutomationAccountSender sends the ListByAutomationAccount request. The method will close the
// http.Response Body if it receives an error.
func (client ConnectionClient) ListByAutomationAccountSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByAutomationAccountResponder handles the response to the ListByAutomationAccount request. The method always
// closes the http.Response Body.
func (client ConnectionClient) ListByAutomationAccountResponder(resp *http.Response) (result ConnectionListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByAutomationAccountNextResults retrieves the next set of results, if any.
func (client ConnectionClient) listByAutomationAccountNextResults(lastResults ConnectionListResult) (result ConnectionListResult, err error) {
	req, err := lastResults.connectionListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "automation.ConnectionClient", "listByAutomationAccountNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByAutomationAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "automation.ConnectionClient", "listByAutomationAccountNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByAutomationAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.ConnectionClient", "listByAutomationAccountNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByAutomationAccountComplete enumerates all values, automatically crossing page boundaries as required.
func (client ConnectionClient) ListByAutomationAccountComplete(ctx context.Context, automationAccountName string) (result ConnectionListResultIterator, err error) {
	result.page, err = client.ListByAutomationAccount(ctx, automationAccountName)
	return
}

// Update update a connection.
//
// automationAccountName is the automation account name. connectionName is the parameters supplied to the update a
// connection operation. parameters is the parameters supplied to the update a connection operation.
func (client ConnectionClient) Update(ctx context.Context, automationAccountName string, connectionName string, parameters ConnectionUpdateParameters) (result Connection, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.ResourceGroupName,
			Constraints: []validation.Constraint{{Target: "client.ResourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "client.ResourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "client.ResourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.ConnectionClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, automationAccountName, connectionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.ConnectionClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.ConnectionClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.ConnectionClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ConnectionClient) UpdatePreparer(ctx context.Context, automationAccountName string, connectionName string, parameters ConnectionUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"connectionName":        autorest.Encode("path", connectionName),
		"resourceGroupName":     autorest.Encode("path", client.ResourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/connections/{connectionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ConnectionClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ConnectionClient) UpdateResponder(resp *http.Response) (result Connection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
