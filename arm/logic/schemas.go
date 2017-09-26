package logic

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
// Code generated by Microsoft (R) AutoRest Code Generator 2.2.22.0
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// SchemasClient is the REST API for Azure Logic Apps.
type SchemasClient struct {
	ManagementClient
}

// NewSchemasClient creates an instance of the SchemasClient client.
func NewSchemasClient(subscriptionID string) SchemasClient {
	return NewSchemasClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSchemasClientWithBaseURI creates an instance of the SchemasClient client.
func NewSchemasClientWithBaseURI(baseURI string, subscriptionID string) SchemasClient {
	return SchemasClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates an integration account schema.
//
// resourceGroupName is the resource group name. integrationAccountName is the integration account name. schemaName is
// the integration account schema name. schema is the integration account schema.
func (client SchemasClient) CreateOrUpdate(resourceGroupName string, integrationAccountName string, schemaName string, schema IntegrationAccountSchema) (result IntegrationAccountSchema, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: schema,
			Constraints: []validation.Constraint{{Target: "schema.IntegrationAccountSchemaProperties", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "logic.SchemasClient", "CreateOrUpdate")
	}

	req, err := client.CreateOrUpdatePreparer(resourceGroupName, integrationAccountName, schemaName, schema)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.SchemasClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.SchemasClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.SchemasClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SchemasClient) CreateOrUpdatePreparer(resourceGroupName string, integrationAccountName string, schemaName string, schema IntegrationAccountSchema) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"schemaName":             autorest.Encode("path", schemaName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/schemas/{schemaName}", pathParameters),
		autorest.WithJSON(schema),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SchemasClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SchemasClient) CreateOrUpdateResponder(resp *http.Response) (result IntegrationAccountSchema, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an integration account schema.
//
// resourceGroupName is the resource group name. integrationAccountName is the integration account name. schemaName is
// the integration account schema name.
func (client SchemasClient) Delete(resourceGroupName string, integrationAccountName string, schemaName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(resourceGroupName, integrationAccountName, schemaName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.SchemasClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "logic.SchemasClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.SchemasClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SchemasClient) DeletePreparer(resourceGroupName string, integrationAccountName string, schemaName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"schemaName":             autorest.Encode("path", schemaName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/schemas/{schemaName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SchemasClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SchemasClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets an integration account schema.
//
// resourceGroupName is the resource group name. integrationAccountName is the integration account name. schemaName is
// the integration account schema name.
func (client SchemasClient) Get(resourceGroupName string, integrationAccountName string, schemaName string) (result IntegrationAccountSchema, err error) {
	req, err := client.GetPreparer(resourceGroupName, integrationAccountName, schemaName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.SchemasClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.SchemasClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.SchemasClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SchemasClient) GetPreparer(resourceGroupName string, integrationAccountName string, schemaName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"schemaName":             autorest.Encode("path", schemaName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/schemas/{schemaName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SchemasClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SchemasClient) GetResponder(resp *http.Response) (result IntegrationAccountSchema, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByIntegrationAccounts gets a list of integration account schemas.
//
// resourceGroupName is the resource group name. integrationAccountName is the integration account name. top is the
// number of items to be included in the result. filter is the filter to apply on the operation.
func (client SchemasClient) ListByIntegrationAccounts(resourceGroupName string, integrationAccountName string, top *int32, filter string) (result IntegrationAccountSchemaListResult, err error) {
	req, err := client.ListByIntegrationAccountsPreparer(resourceGroupName, integrationAccountName, top, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.SchemasClient", "ListByIntegrationAccounts", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByIntegrationAccountsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.SchemasClient", "ListByIntegrationAccounts", resp, "Failure sending request")
		return
	}

	result, err = client.ListByIntegrationAccountsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.SchemasClient", "ListByIntegrationAccounts", resp, "Failure responding to request")
	}

	return
}

// ListByIntegrationAccountsPreparer prepares the ListByIntegrationAccounts request.
func (client SchemasClient) ListByIntegrationAccountsPreparer(resourceGroupName string, integrationAccountName string, top *int32, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/schemas", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByIntegrationAccountsSender sends the ListByIntegrationAccounts request. The method will close the
// http.Response Body if it receives an error.
func (client SchemasClient) ListByIntegrationAccountsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByIntegrationAccountsResponder handles the response to the ListByIntegrationAccounts request. The method always
// closes the http.Response Body.
func (client SchemasClient) ListByIntegrationAccountsResponder(resp *http.Response) (result IntegrationAccountSchemaListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByIntegrationAccountsNextResults retrieves the next set of results, if any.
func (client SchemasClient) ListByIntegrationAccountsNextResults(lastResults IntegrationAccountSchemaListResult) (result IntegrationAccountSchemaListResult, err error) {
	req, err := lastResults.IntegrationAccountSchemaListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.SchemasClient", "ListByIntegrationAccounts", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByIntegrationAccountsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic.SchemasClient", "ListByIntegrationAccounts", resp, "Failure sending next results request")
	}

	result, err = client.ListByIntegrationAccountsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.SchemasClient", "ListByIntegrationAccounts", resp, "Failure responding to next results request")
	}

	return
}

// ListByIntegrationAccountsComplete gets all elements from the list without paging.
func (client SchemasClient) ListByIntegrationAccountsComplete(resourceGroupName string, integrationAccountName string, top *int32, filter string, cancel <-chan struct{}) (<-chan IntegrationAccountSchema, <-chan error) {
	resultChan := make(chan IntegrationAccountSchema)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListByIntegrationAccounts(resourceGroupName, integrationAccountName, top, filter)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListByIntegrationAccountsNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}
