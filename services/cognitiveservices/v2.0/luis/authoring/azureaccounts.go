package authoring

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
	"github.com/satori/go.uuid"
	"net/http"
)

// AzureAccountsClient is the client for the AzureAccounts methods of the Authoring service.
type AzureAccountsClient struct {
	BaseClient
}

// NewAzureAccountsClient creates an instance of the AzureAccountsClient client.
func NewAzureAccountsClient(endpoint string) AzureAccountsClient {
	return AzureAccountsClient{New(endpoint)}
}

// AssignToApp assigns an Azure account to the application.
// Parameters:
// appID - the application ID.
// azureAccountInfoObject - the Azure account information object.
func (client AzureAccountsClient) AssignToApp(ctx context.Context, appID uuid.UUID, azureAccountInfoObject *AzureAccountInfoObject) (result OperationStatus, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AzureAccountsClient.AssignToApp")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: azureAccountInfoObject,
			Constraints: []validation.Constraint{{Target: "azureAccountInfoObject", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "azureAccountInfoObject.AzureSubscriptionID", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "azureAccountInfoObject.ResourceGroup", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "azureAccountInfoObject.AccountName", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("authoring.AzureAccountsClient", "AssignToApp", err.Error())
	}

	req, err := client.AssignToAppPreparer(ctx, appID, azureAccountInfoObject)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.AzureAccountsClient", "AssignToApp", nil, "Failure preparing request")
		return
	}

	resp, err := client.AssignToAppSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authoring.AzureAccountsClient", "AssignToApp", resp, "Failure sending request")
		return
	}

	result, err = client.AssignToAppResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.AzureAccountsClient", "AssignToApp", resp, "Failure responding to request")
		return
	}

	return
}

// AssignToAppPreparer prepares the AssignToApp request.
func (client AzureAccountsClient) AssignToAppPreparer(ctx context.Context, appID uuid.UUID, azureAccountInfoObject *AzureAccountInfoObject) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"appId": autorest.Encode("path", appID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/luis/api/v2.0", urlParameters),
		autorest.WithPathParameters("/apps/{appId}/azureaccounts", pathParameters))
	if azureAccountInfoObject != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(azureAccountInfoObject))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AssignToAppSender sends the AssignToApp request. The method will close the
// http.Response Body if it receives an error.
func (client AzureAccountsClient) AssignToAppSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AssignToAppResponder handles the response to the AssignToApp request. The method always
// closes the http.Response Body.
func (client AzureAccountsClient) AssignToAppResponder(resp *http.Response) (result OperationStatus, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetAssigned gets the LUIS Azure accounts assigned to the application for the user using his ARM token.
// Parameters:
// appID - the application ID.
func (client AzureAccountsClient) GetAssigned(ctx context.Context, appID uuid.UUID) (result ListAzureAccountInfoObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AzureAccountsClient.GetAssigned")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetAssignedPreparer(ctx, appID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.AzureAccountsClient", "GetAssigned", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAssignedSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authoring.AzureAccountsClient", "GetAssigned", resp, "Failure sending request")
		return
	}

	result, err = client.GetAssignedResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.AzureAccountsClient", "GetAssigned", resp, "Failure responding to request")
		return
	}

	return
}

// GetAssignedPreparer prepares the GetAssigned request.
func (client AzureAccountsClient) GetAssignedPreparer(ctx context.Context, appID uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"appId": autorest.Encode("path", appID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/luis/api/v2.0", urlParameters),
		autorest.WithPathParameters("/apps/{appId}/azureaccounts", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAssignedSender sends the GetAssigned request. The method will close the
// http.Response Body if it receives an error.
func (client AzureAccountsClient) GetAssignedSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetAssignedResponder handles the response to the GetAssigned request. The method always
// closes the http.Response Body.
func (client AzureAccountsClient) GetAssignedResponder(resp *http.Response) (result ListAzureAccountInfoObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListUserLUISAccounts gets the LUIS Azure accounts for the user using his ARM token.
func (client AzureAccountsClient) ListUserLUISAccounts(ctx context.Context) (result ListAzureAccountInfoObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AzureAccountsClient.ListUserLUISAccounts")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListUserLUISAccountsPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.AzureAccountsClient", "ListUserLUISAccounts", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListUserLUISAccountsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authoring.AzureAccountsClient", "ListUserLUISAccounts", resp, "Failure sending request")
		return
	}

	result, err = client.ListUserLUISAccountsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.AzureAccountsClient", "ListUserLUISAccounts", resp, "Failure responding to request")
		return
	}

	return
}

// ListUserLUISAccountsPreparer prepares the ListUserLUISAccounts request.
func (client AzureAccountsClient) ListUserLUISAccountsPreparer(ctx context.Context) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/luis/api/v2.0", urlParameters),
		autorest.WithPath("/azureaccounts"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListUserLUISAccountsSender sends the ListUserLUISAccounts request. The method will close the
// http.Response Body if it receives an error.
func (client AzureAccountsClient) ListUserLUISAccountsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListUserLUISAccountsResponder handles the response to the ListUserLUISAccounts request. The method always
// closes the http.Response Body.
func (client AzureAccountsClient) ListUserLUISAccountsResponder(resp *http.Response) (result ListAzureAccountInfoObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// RemoveFromApp removes assigned Azure account from the application.
// Parameters:
// appID - the application ID.
// azureAccountInfoObject - the Azure account information object.
func (client AzureAccountsClient) RemoveFromApp(ctx context.Context, appID uuid.UUID, azureAccountInfoObject *AzureAccountInfoObject) (result OperationStatus, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AzureAccountsClient.RemoveFromApp")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: azureAccountInfoObject,
			Constraints: []validation.Constraint{{Target: "azureAccountInfoObject", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "azureAccountInfoObject.AzureSubscriptionID", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "azureAccountInfoObject.ResourceGroup", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "azureAccountInfoObject.AccountName", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("authoring.AzureAccountsClient", "RemoveFromApp", err.Error())
	}

	req, err := client.RemoveFromAppPreparer(ctx, appID, azureAccountInfoObject)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.AzureAccountsClient", "RemoveFromApp", nil, "Failure preparing request")
		return
	}

	resp, err := client.RemoveFromAppSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authoring.AzureAccountsClient", "RemoveFromApp", resp, "Failure sending request")
		return
	}

	result, err = client.RemoveFromAppResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.AzureAccountsClient", "RemoveFromApp", resp, "Failure responding to request")
		return
	}

	return
}

// RemoveFromAppPreparer prepares the RemoveFromApp request.
func (client AzureAccountsClient) RemoveFromAppPreparer(ctx context.Context, appID uuid.UUID, azureAccountInfoObject *AzureAccountInfoObject) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"appId": autorest.Encode("path", appID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("{Endpoint}/luis/api/v2.0", urlParameters),
		autorest.WithPathParameters("/apps/{appId}/azureaccounts", pathParameters))
	if azureAccountInfoObject != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(azureAccountInfoObject))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RemoveFromAppSender sends the RemoveFromApp request. The method will close the
// http.Response Body if it receives an error.
func (client AzureAccountsClient) RemoveFromAppSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// RemoveFromAppResponder handles the response to the RemoveFromApp request. The method always
// closes the http.Response Body.
func (client AzureAccountsClient) RemoveFromAppResponder(resp *http.Response) (result OperationStatus, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
