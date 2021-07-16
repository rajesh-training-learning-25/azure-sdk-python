package costmanagement

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// AlertsClient is the client for the Alerts methods of the Costmanagement service.
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

// Dismiss dismisses the specified alert
// Parameters:
// scope - the scope associated with alerts operations. This includes '/subscriptions/{subscriptionId}/' for
// subscription scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup
// scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope and
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department
// scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}'
// for EnrollmentAccount scope, '/providers/Microsoft.Management/managementGroups/{managementGroupId} for
// Management Group scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for
// billingProfile scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}'
// for invoiceSection scope, and
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for
// partners.
// alertID - alert ID
// parameters - parameters supplied to the Dismiss Alert operation.
func (client AlertsClient) Dismiss(ctx context.Context, scope string, alertID string, parameters DismissAlertPayload) (result Alert, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AlertsClient.Dismiss")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DismissPreparer(ctx, scope, alertID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.AlertsClient", "Dismiss", nil, "Failure preparing request")
		return
	}

	resp, err := client.DismissSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.AlertsClient", "Dismiss", resp, "Failure sending request")
		return
	}

	result, err = client.DismissResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.AlertsClient", "Dismiss", resp, "Failure responding to request")
		return
	}

	return
}

// DismissPreparer prepares the Dismiss request.
func (client AlertsClient) DismissPreparer(ctx context.Context, scope string, alertID string, parameters DismissAlertPayload) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"alertId": alertID,
		"scope":   scope,
	}

	const APIVersion = "2020-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.CostManagement/alerts/{alertId}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DismissSender sends the Dismiss request. The method will close the
// http.Response Body if it receives an error.
func (client AlertsClient) DismissSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DismissResponder handles the response to the Dismiss request. The method always
// closes the http.Response Body.
func (client AlertsClient) DismissResponder(resp *http.Response) (result Alert, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets the alert for the scope by alert ID.
// Parameters:
// scope - the scope associated with alerts operations. This includes '/subscriptions/{subscriptionId}/' for
// subscription scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup
// scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope and
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department
// scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}'
// for EnrollmentAccount scope, '/providers/Microsoft.Management/managementGroups/{managementGroupId} for
// Management Group scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for
// billingProfile scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}'
// for invoiceSection scope, and
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for
// partners.
// alertID - alert ID
func (client AlertsClient) Get(ctx context.Context, scope string, alertID string) (result Alert, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AlertsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, scope, alertID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.AlertsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.AlertsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.AlertsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client AlertsClient) GetPreparer(ctx context.Context, scope string, alertID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"alertId": alertID,
		"scope":   scope,
	}

	const APIVersion = "2020-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.CostManagement/alerts/{alertId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AlertsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AlertsClient) GetResponder(resp *http.Response) (result Alert, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the alerts for scope defined.
// Parameters:
// scope - the scope associated with alerts operations. This includes '/subscriptions/{subscriptionId}/' for
// subscription scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup
// scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope and
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department
// scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}'
// for EnrollmentAccount scope, '/providers/Microsoft.Management/managementGroups/{managementGroupId} for
// Management Group scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for
// billingProfile scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}'
// for invoiceSection scope, and
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for
// partners.
func (client AlertsClient) List(ctx context.Context, scope string) (result AlertsResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AlertsClient.List")
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
		err = autorest.NewErrorWithError(err, "costmanagement.AlertsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.AlertsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.AlertsClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client AlertsClient) ListPreparer(ctx context.Context, scope string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scope": scope,
	}

	const APIVersion = "2020-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.CostManagement/alerts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AlertsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AlertsClient) ListResponder(resp *http.Response) (result AlertsResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListExternal lists the Alerts for external cloud provider type defined.
// Parameters:
// externalCloudProviderType - the external cloud provider type associated with dimension/query operations.
// This includes 'externalSubscriptions' for linked account and 'externalBillingAccounts' for consolidated
// account.
// externalCloudProviderID - this can be '{externalSubscriptionId}' for linked account or
// '{externalBillingAccountId}' for consolidated account used with dimension/query operations.
func (client AlertsClient) ListExternal(ctx context.Context, externalCloudProviderType ExternalCloudProviderType, externalCloudProviderID string) (result AlertsResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AlertsClient.ListExternal")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListExternalPreparer(ctx, externalCloudProviderType, externalCloudProviderID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.AlertsClient", "ListExternal", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListExternalSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.AlertsClient", "ListExternal", resp, "Failure sending request")
		return
	}

	result, err = client.ListExternalResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.AlertsClient", "ListExternal", resp, "Failure responding to request")
		return
	}

	return
}

// ListExternalPreparer prepares the ListExternal request.
func (client AlertsClient) ListExternalPreparer(ctx context.Context, externalCloudProviderType ExternalCloudProviderType, externalCloudProviderID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"externalCloudProviderId":   autorest.Encode("path", externalCloudProviderID),
		"externalCloudProviderType": autorest.Encode("path", externalCloudProviderType),
	}

	const APIVersion = "2020-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.CostManagement/{externalCloudProviderType}/{externalCloudProviderId}/alerts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListExternalSender sends the ListExternal request. The method will close the
// http.Response Body if it receives an error.
func (client AlertsClient) ListExternalSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListExternalResponder handles the response to the ListExternal request. The method always
// closes the http.Response Body.
func (client AlertsClient) ListExternalResponder(resp *http.Response) (result AlertsResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
