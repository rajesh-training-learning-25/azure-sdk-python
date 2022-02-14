package datamigration

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
	"github.com/gofrs/uuid"
	"net/http"
)

// DatabaseMigrationsSQLMiClient is the data Migration Client
type DatabaseMigrationsSQLMiClient struct {
	BaseClient
}

// NewDatabaseMigrationsSQLMiClient creates an instance of the DatabaseMigrationsSQLMiClient client.
func NewDatabaseMigrationsSQLMiClient(subscriptionID string) DatabaseMigrationsSQLMiClient {
	return NewDatabaseMigrationsSQLMiClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDatabaseMigrationsSQLMiClientWithBaseURI creates an instance of the DatabaseMigrationsSQLMiClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewDatabaseMigrationsSQLMiClientWithBaseURI(baseURI string, subscriptionID string) DatabaseMigrationsSQLMiClient {
	return DatabaseMigrationsSQLMiClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Cancel stop migrations in progress for the database
// Parameters:
// resourceGroupName - name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal.
// targetDbName - the name of the target database.
// parameters - required migration operation ID for which cancel will be initiated.
func (client DatabaseMigrationsSQLMiClient) Cancel(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, parameters MigrationOperationInput) (result DatabaseMigrationsSQLMiCancelFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabaseMigrationsSQLMiClient.Cancel")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CancelPreparer(ctx, resourceGroupName, managedInstanceName, targetDbName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.DatabaseMigrationsSQLMiClient", "Cancel", nil, "Failure preparing request")
		return
	}

	result, err = client.CancelSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.DatabaseMigrationsSQLMiClient", "Cancel", result.Response(), "Failure sending request")
		return
	}

	return
}

// CancelPreparer prepares the Cancel request.
func (client DatabaseMigrationsSQLMiClient) CancelPreparer(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, parameters MigrationOperationInput) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managedInstanceName": autorest.Encode("path", managedInstanceName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
		"targetDbName":        autorest.Encode("path", targetDbName),
	}

	const APIVersion = "2021-10-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/providers/Microsoft.DataMigration/databaseMigrations/{targetDbName}/cancel", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CancelSender sends the Cancel request. The method will close the
// http.Response Body if it receives an error.
func (client DatabaseMigrationsSQLMiClient) CancelSender(req *http.Request) (future DatabaseMigrationsSQLMiCancelFuture, err error) {
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

// CancelResponder handles the response to the Cancel request. The method always
// closes the http.Response Body.
func (client DatabaseMigrationsSQLMiClient) CancelResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CreateOrUpdate create or Update Database Migration resource.
// Parameters:
// resourceGroupName - name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal.
// targetDbName - the name of the target database.
// parameters - details of SqlMigrationService resource.
func (client DatabaseMigrationsSQLMiClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, parameters DatabaseMigrationSQLMi) (result DatabaseMigrationsSQLMiCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabaseMigrationsSQLMiClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, managedInstanceName, targetDbName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.DatabaseMigrationsSQLMiClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.DatabaseMigrationsSQLMiClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DatabaseMigrationsSQLMiClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, parameters DatabaseMigrationSQLMi) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managedInstanceName": autorest.Encode("path", managedInstanceName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
		"targetDbName":        autorest.Encode("path", targetDbName),
	}

	const APIVersion = "2021-10-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.SystemData = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/providers/Microsoft.DataMigration/databaseMigrations/{targetDbName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DatabaseMigrationsSQLMiClient) CreateOrUpdateSender(req *http.Request) (future DatabaseMigrationsSQLMiCreateOrUpdateFuture, err error) {
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

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DatabaseMigrationsSQLMiClient) CreateOrUpdateResponder(resp *http.Response) (result DatabaseMigrationSQLMi, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Cutover initiate cutover for online migration in progress for the database.
// Parameters:
// resourceGroupName - name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal.
// targetDbName - the name of the target database.
// parameters - required migration operation ID for which cutover will be initiated.
func (client DatabaseMigrationsSQLMiClient) Cutover(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, parameters MigrationOperationInput) (result DatabaseMigrationsSQLMiCutoverFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabaseMigrationsSQLMiClient.Cutover")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CutoverPreparer(ctx, resourceGroupName, managedInstanceName, targetDbName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.DatabaseMigrationsSQLMiClient", "Cutover", nil, "Failure preparing request")
		return
	}

	result, err = client.CutoverSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.DatabaseMigrationsSQLMiClient", "Cutover", result.Response(), "Failure sending request")
		return
	}

	return
}

// CutoverPreparer prepares the Cutover request.
func (client DatabaseMigrationsSQLMiClient) CutoverPreparer(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, parameters MigrationOperationInput) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managedInstanceName": autorest.Encode("path", managedInstanceName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
		"targetDbName":        autorest.Encode("path", targetDbName),
	}

	const APIVersion = "2021-10-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/providers/Microsoft.DataMigration/databaseMigrations/{targetDbName}/cutover", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CutoverSender sends the Cutover request. The method will close the
// http.Response Body if it receives an error.
func (client DatabaseMigrationsSQLMiClient) CutoverSender(req *http.Request) (future DatabaseMigrationsSQLMiCutoverFuture, err error) {
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

// CutoverResponder handles the response to the Cutover request. The method always
// closes the http.Response Body.
func (client DatabaseMigrationsSQLMiClient) CutoverResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieve the Database Migration resource.
// Parameters:
// resourceGroupName - name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal.
// targetDbName - the name of the target database.
// migrationOperationID - optional migration operation ID. If this is provided, then details of migration
// operation for that ID are retrieved. If not provided (default), then details related to most recent or
// current operation are retrieved.
// expand - the child resources to include in the response.
func (client DatabaseMigrationsSQLMiClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, migrationOperationID *uuid.UUID, expand string) (result DatabaseMigrationSQLMi, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabaseMigrationsSQLMiClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, managedInstanceName, targetDbName, migrationOperationID, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.DatabaseMigrationsSQLMiClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datamigration.DatabaseMigrationsSQLMiClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.DatabaseMigrationsSQLMiClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client DatabaseMigrationsSQLMiClient) GetPreparer(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, migrationOperationID *uuid.UUID, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managedInstanceName": autorest.Encode("path", managedInstanceName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
		"targetDbName":        autorest.Encode("path", targetDbName),
	}

	const APIVersion = "2021-10-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if migrationOperationID != nil {
		queryParameters["migrationOperationId"] = autorest.Encode("query", *migrationOperationID)
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/providers/Microsoft.DataMigration/databaseMigrations/{targetDbName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DatabaseMigrationsSQLMiClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DatabaseMigrationsSQLMiClient) GetResponder(resp *http.Response) (result DatabaseMigrationSQLMi, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
