package costmanagement

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

// ExportsClient is the client for the Exports methods of the Costmanagement service.
type ExportsClient struct {
	BaseClient
}

// NewExportsClient creates an instance of the ExportsClient client.
func NewExportsClient(subscriptionID string) ExportsClient {
	return NewExportsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewExportsClientWithBaseURI creates an instance of the ExportsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewExportsClientWithBaseURI(baseURI string, subscriptionID string) ExportsClient {
	return ExportsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate the operation to create or update a export. Update operation requires latest eTag to be set in the
// request. You may obtain the latest eTag by performing a get operation. Create operation does not require eTag.
// Parameters:
// scope - the scope associated with query and export operations. This includes
// '/subscriptions/{subscriptionId}/' for subscription scope,
// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope and
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department
// scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}'
// for EnrollmentAccount scope, '/providers/Microsoft.Management/managementGroups/{managementGroupId} for
// Management Group scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for
// billingProfile scope,
// 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}'
// for invoiceSection scope,
// 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for
// partners, 'providers/Microsoft.CostManagement/ExternalSubscriptions/{externalSubscriptionId}' for linked
// account and 'providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountId}' for
// consolidated account
// exportName - export Name.
// parameters - parameters supplied to the CreateOrUpdate Export operation.
func (client ExportsClient) CreateOrUpdate(ctx context.Context, scope string, exportName string, parameters Export) (result Export, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExportsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ExportProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.ExportProperties.Schedule", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.ExportProperties.Schedule.RecurrencePeriod", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.ExportProperties.Schedule.RecurrencePeriod.From", Name: validation.Null, Rule: true, Chain: nil}}},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("costmanagement.ExportsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, scope, exportName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ExportsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.ExportsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ExportsClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ExportsClient) CreateOrUpdatePreparer(ctx context.Context, scope string, exportName string, parameters Export) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"exportName": autorest.Encode("path", exportName),
		"scope":      scope,
	}

	const APIVersion = "2019-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.CostManagement/exports/{exportName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ExportsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ExportsClient) CreateOrUpdateResponder(resp *http.Response) (result Export, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the operation to delete a export.
// Parameters:
// scope - the scope associated with query and export operations. This includes
// '/subscriptions/{subscriptionId}/' for subscription scope,
// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope and
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department
// scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}'
// for EnrollmentAccount scope, '/providers/Microsoft.Management/managementGroups/{managementGroupId} for
// Management Group scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for
// billingProfile scope,
// 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}'
// for invoiceSection scope,
// 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for
// partners, 'providers/Microsoft.CostManagement/ExternalSubscriptions/{externalSubscriptionId}' for linked
// account and 'providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountId}' for
// consolidated account
// exportName - export Name.
func (client ExportsClient) Delete(ctx context.Context, scope string, exportName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExportsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, scope, exportName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ExportsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "costmanagement.ExportsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ExportsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ExportsClient) DeletePreparer(ctx context.Context, scope string, exportName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"exportName": autorest.Encode("path", exportName),
		"scope":      scope,
	}

	const APIVersion = "2019-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.CostManagement/exports/{exportName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ExportsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ExportsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Execute the operation to execute a export.
// Parameters:
// scope - the scope associated with query and export operations. This includes
// '/subscriptions/{subscriptionId}/' for subscription scope,
// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope and
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department
// scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}'
// for EnrollmentAccount scope, '/providers/Microsoft.Management/managementGroups/{managementGroupId} for
// Management Group scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for
// billingProfile scope,
// 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}'
// for invoiceSection scope,
// 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for
// partners, 'providers/Microsoft.CostManagement/ExternalSubscriptions/{externalSubscriptionId}' for linked
// account and 'providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountId}' for
// consolidated account
// exportName - export Name.
func (client ExportsClient) Execute(ctx context.Context, scope string, exportName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExportsClient.Execute")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ExecutePreparer(ctx, scope, exportName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ExportsClient", "Execute", nil, "Failure preparing request")
		return
	}

	resp, err := client.ExecuteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "costmanagement.ExportsClient", "Execute", resp, "Failure sending request")
		return
	}

	result, err = client.ExecuteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ExportsClient", "Execute", resp, "Failure responding to request")
		return
	}

	return
}

// ExecutePreparer prepares the Execute request.
func (client ExportsClient) ExecutePreparer(ctx context.Context, scope string, exportName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"exportName": autorest.Encode("path", exportName),
		"scope":      scope,
	}

	const APIVersion = "2019-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.CostManagement/exports/{exportName}/run", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ExecuteSender sends the Execute request. The method will close the
// http.Response Body if it receives an error.
func (client ExportsClient) ExecuteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ExecuteResponder handles the response to the Execute request. The method always
// closes the http.Response Body.
func (client ExportsClient) ExecuteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get the operation to get the export for the defined scope by export name.
// Parameters:
// scope - the scope associated with query and export operations. This includes
// '/subscriptions/{subscriptionId}/' for subscription scope,
// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope and
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department
// scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}'
// for EnrollmentAccount scope, '/providers/Microsoft.Management/managementGroups/{managementGroupId} for
// Management Group scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for
// billingProfile scope,
// 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}'
// for invoiceSection scope,
// 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for
// partners, 'providers/Microsoft.CostManagement/ExternalSubscriptions/{externalSubscriptionId}' for linked
// account and 'providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountId}' for
// consolidated account
// exportName - export Name.
func (client ExportsClient) Get(ctx context.Context, scope string, exportName string) (result Export, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExportsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, scope, exportName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ExportsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.ExportsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ExportsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ExportsClient) GetPreparer(ctx context.Context, scope string, exportName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"exportName": autorest.Encode("path", exportName),
		"scope":      scope,
	}

	const APIVersion = "2019-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.CostManagement/exports/{exportName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ExportsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ExportsClient) GetResponder(resp *http.Response) (result Export, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetExecutionHistory the operation to get the execution history of an export for the defined scope by export name.
// Parameters:
// scope - the scope associated with query and export operations. This includes
// '/subscriptions/{subscriptionId}/' for subscription scope,
// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope and
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department
// scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}'
// for EnrollmentAccount scope, '/providers/Microsoft.Management/managementGroups/{managementGroupId} for
// Management Group scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for
// billingProfile scope,
// 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}'
// for invoiceSection scope,
// 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for
// partners, 'providers/Microsoft.CostManagement/ExternalSubscriptions/{externalSubscriptionId}' for linked
// account and 'providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountId}' for
// consolidated account
// exportName - export Name.
func (client ExportsClient) GetExecutionHistory(ctx context.Context, scope string, exportName string) (result ExportExecutionListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExportsClient.GetExecutionHistory")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetExecutionHistoryPreparer(ctx, scope, exportName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ExportsClient", "GetExecutionHistory", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetExecutionHistorySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.ExportsClient", "GetExecutionHistory", resp, "Failure sending request")
		return
	}

	result, err = client.GetExecutionHistoryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ExportsClient", "GetExecutionHistory", resp, "Failure responding to request")
		return
	}

	return
}

// GetExecutionHistoryPreparer prepares the GetExecutionHistory request.
func (client ExportsClient) GetExecutionHistoryPreparer(ctx context.Context, scope string, exportName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"exportName": autorest.Encode("path", exportName),
		"scope":      scope,
	}

	const APIVersion = "2019-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.CostManagement/exports/{exportName}/runHistory", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetExecutionHistorySender sends the GetExecutionHistory request. The method will close the
// http.Response Body if it receives an error.
func (client ExportsClient) GetExecutionHistorySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetExecutionHistoryResponder handles the response to the GetExecutionHistory request. The method always
// closes the http.Response Body.
func (client ExportsClient) GetExecutionHistoryResponder(resp *http.Response) (result ExportExecutionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List the operation to list all exports at the given scope.
// Parameters:
// scope - the scope associated with query and export operations. This includes
// '/subscriptions/{subscriptionId}/' for subscription scope,
// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope and
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department
// scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}'
// for EnrollmentAccount scope, '/providers/Microsoft.Management/managementGroups/{managementGroupId} for
// Management Group scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for
// billingProfile scope,
// 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}'
// for invoiceSection scope,
// 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for
// partners, 'providers/Microsoft.CostManagement/ExternalSubscriptions/{externalSubscriptionId}' for linked
// account and 'providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountId}' for
// consolidated account
func (client ExportsClient) List(ctx context.Context, scope string) (result ExportListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExportsClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, scope)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ExportsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.ExportsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ExportsClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ExportsClient) ListPreparer(ctx context.Context, scope string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scope": scope,
	}

	const APIVersion = "2019-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.CostManagement/exports", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ExportsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ExportsClient) ListResponder(resp *http.Response) (result ExportListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
