package datashare

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

// SynchronizationSettingsClient is the creates a Microsoft.DataShare management client.
type SynchronizationSettingsClient struct {
	BaseClient
}

// NewSynchronizationSettingsClient creates an instance of the SynchronizationSettingsClient client.
func NewSynchronizationSettingsClient(subscriptionID string) SynchronizationSettingsClient {
	return NewSynchronizationSettingsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSynchronizationSettingsClientWithBaseURI creates an instance of the SynchronizationSettingsClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewSynchronizationSettingsClientWithBaseURI(baseURI string, subscriptionID string) SynchronizationSettingsClient {
	return SynchronizationSettingsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create create a synchronizationSetting
// Parameters:
// resourceGroupName - the resource group name.
// accountName - the name of the share account.
// shareName - the name of the share to add the synchronization setting to.
// synchronizationSettingName - the name of the synchronizationSetting.
// synchronizationSetting - the new synchronization setting information.
func (client SynchronizationSettingsClient) Create(ctx context.Context, resourceGroupName string, accountName string, shareName string, synchronizationSettingName string, synchronizationSetting BasicSynchronizationSetting) (result SynchronizationSettingModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SynchronizationSettingsClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, resourceGroupName, accountName, shareName, synchronizationSettingName, synchronizationSetting)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.SynchronizationSettingsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datashare.SynchronizationSettingsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.SynchronizationSettingsClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client SynchronizationSettingsClient) CreatePreparer(ctx context.Context, resourceGroupName string, accountName string, shareName string, synchronizationSettingName string, synchronizationSetting BasicSynchronizationSetting) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":                autorest.Encode("path", accountName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"shareName":                  autorest.Encode("path", shareName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
		"synchronizationSettingName": autorest.Encode("path", synchronizationSettingName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/synchronizationSettings/{synchronizationSettingName}", pathParameters),
		autorest.WithJSON(synchronizationSetting),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client SynchronizationSettingsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client SynchronizationSettingsClient) CreateResponder(resp *http.Response) (result SynchronizationSettingModel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a synchronizationSetting in a share
// Parameters:
// resourceGroupName - the resource group name.
// accountName - the name of the share account.
// shareName - the name of the share.
// synchronizationSettingName - the name of the synchronizationSetting .
func (client SynchronizationSettingsClient) Delete(ctx context.Context, resourceGroupName string, accountName string, shareName string, synchronizationSettingName string) (result SynchronizationSettingsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SynchronizationSettingsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, accountName, shareName, synchronizationSettingName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.SynchronizationSettingsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.SynchronizationSettingsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SynchronizationSettingsClient) DeletePreparer(ctx context.Context, resourceGroupName string, accountName string, shareName string, synchronizationSettingName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":                autorest.Encode("path", accountName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"shareName":                  autorest.Encode("path", shareName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
		"synchronizationSettingName": autorest.Encode("path", synchronizationSettingName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/synchronizationSettings/{synchronizationSettingName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SynchronizationSettingsClient) DeleteSender(req *http.Request) (future SynchronizationSettingsDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SynchronizationSettingsClient) DeleteResponder(resp *http.Response) (result OperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get a synchronizationSetting in a share
// Parameters:
// resourceGroupName - the resource group name.
// accountName - the name of the share account.
// shareName - the name of the share.
// synchronizationSettingName - the name of the synchronizationSetting.
func (client SynchronizationSettingsClient) Get(ctx context.Context, resourceGroupName string, accountName string, shareName string, synchronizationSettingName string) (result SynchronizationSettingModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SynchronizationSettingsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, accountName, shareName, synchronizationSettingName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.SynchronizationSettingsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datashare.SynchronizationSettingsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.SynchronizationSettingsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client SynchronizationSettingsClient) GetPreparer(ctx context.Context, resourceGroupName string, accountName string, shareName string, synchronizationSettingName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":                autorest.Encode("path", accountName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"shareName":                  autorest.Encode("path", shareName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
		"synchronizationSettingName": autorest.Encode("path", synchronizationSettingName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/synchronizationSettings/{synchronizationSettingName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SynchronizationSettingsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SynchronizationSettingsClient) GetResponder(resp *http.Response) (result SynchronizationSettingModel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByShare list synchronizationSettings in a share
// Parameters:
// resourceGroupName - the resource group name.
// accountName - the name of the share account.
// shareName - the name of the share.
// skipToken - continuation token
func (client SynchronizationSettingsClient) ListByShare(ctx context.Context, resourceGroupName string, accountName string, shareName string, skipToken string) (result SynchronizationSettingListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SynchronizationSettingsClient.ListByShare")
		defer func() {
			sc := -1
			if result.ssl.Response.Response != nil {
				sc = result.ssl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByShareNextResults
	req, err := client.ListBySharePreparer(ctx, resourceGroupName, accountName, shareName, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.SynchronizationSettingsClient", "ListByShare", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByShareSender(req)
	if err != nil {
		result.ssl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datashare.SynchronizationSettingsClient", "ListByShare", resp, "Failure sending request")
		return
	}

	result.ssl, err = client.ListByShareResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.SynchronizationSettingsClient", "ListByShare", resp, "Failure responding to request")
		return
	}
	if result.ssl.hasNextLink() && result.ssl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySharePreparer prepares the ListByShare request.
func (client SynchronizationSettingsClient) ListBySharePreparer(ctx context.Context, resourceGroupName string, accountName string, shareName string, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"shareName":         autorest.Encode("path", shareName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/synchronizationSettings", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByShareSender sends the ListByShare request. The method will close the
// http.Response Body if it receives an error.
func (client SynchronizationSettingsClient) ListByShareSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByShareResponder handles the response to the ListByShare request. The method always
// closes the http.Response Body.
func (client SynchronizationSettingsClient) ListByShareResponder(resp *http.Response) (result SynchronizationSettingList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByShareNextResults retrieves the next set of results, if any.
func (client SynchronizationSettingsClient) listByShareNextResults(ctx context.Context, lastResults SynchronizationSettingList) (result SynchronizationSettingList, err error) {
	req, err := lastResults.synchronizationSettingListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "datashare.SynchronizationSettingsClient", "listByShareNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByShareSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "datashare.SynchronizationSettingsClient", "listByShareNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByShareResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.SynchronizationSettingsClient", "listByShareNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByShareComplete enumerates all values, automatically crossing page boundaries as required.
func (client SynchronizationSettingsClient) ListByShareComplete(ctx context.Context, resourceGroupName string, accountName string, shareName string, skipToken string) (result SynchronizationSettingListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SynchronizationSettingsClient.ListByShare")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByShare(ctx, resourceGroupName, accountName, shareName, skipToken)
	return
}
