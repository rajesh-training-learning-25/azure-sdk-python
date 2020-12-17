package synapse

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

// WorkspaceManagedSQLServerRecoverableSqlpoolsClient is the azure Synapse Analytics Management Client
type WorkspaceManagedSQLServerRecoverableSqlpoolsClient struct {
	BaseClient
}

// NewWorkspaceManagedSQLServerRecoverableSqlpoolsClient creates an instance of the
// WorkspaceManagedSQLServerRecoverableSqlpoolsClient client.
func NewWorkspaceManagedSQLServerRecoverableSqlpoolsClient(subscriptionID string) WorkspaceManagedSQLServerRecoverableSqlpoolsClient {
	return NewWorkspaceManagedSQLServerRecoverableSqlpoolsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkspaceManagedSQLServerRecoverableSqlpoolsClientWithBaseURI creates an instance of the
// WorkspaceManagedSQLServerRecoverableSqlpoolsClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWorkspaceManagedSQLServerRecoverableSqlpoolsClientWithBaseURI(baseURI string, subscriptionID string) WorkspaceManagedSQLServerRecoverableSqlpoolsClient {
	return WorkspaceManagedSQLServerRecoverableSqlpoolsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get recoverable sql pools for workspace managed sql server.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// SQLComputeName - the name of the sql compute
func (client WorkspaceManagedSQLServerRecoverableSqlpoolsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, SQLComputeName string) (result RecoverableSQLPool, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceManagedSQLServerRecoverableSqlpoolsClient.Get")
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
		return result, validation.NewError("synapse.WorkspaceManagedSQLServerRecoverableSqlpoolsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName, SQLComputeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerRecoverableSqlpoolsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerRecoverableSqlpoolsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerRecoverableSqlpoolsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client WorkspaceManagedSQLServerRecoverableSqlpoolsClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string, SQLComputeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"sqlComputeName":    autorest.Encode("path", SQLComputeName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/recoverableSqlPools/{sqlComputeName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceManagedSQLServerRecoverableSqlpoolsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client WorkspaceManagedSQLServerRecoverableSqlpoolsClient) GetResponder(resp *http.Response) (result RecoverableSQLPool, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get list of recoverable sql pools for workspace managed sql server.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
func (client WorkspaceManagedSQLServerRecoverableSqlpoolsClient) List(ctx context.Context, resourceGroupName string, workspaceName string) (result RecoverableSQLPoolListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceManagedSQLServerRecoverableSqlpoolsClient.List")
		defer func() {
			sc := -1
			if result.rsplr.Response.Response != nil {
				sc = result.rsplr.Response.Response.StatusCode
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
		return result, validation.NewError("synapse.WorkspaceManagedSQLServerRecoverableSqlpoolsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, workspaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerRecoverableSqlpoolsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rsplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerRecoverableSqlpoolsClient", "List", resp, "Failure sending request")
		return
	}

	result.rsplr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerRecoverableSqlpoolsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.rsplr.hasNextLink() && result.rsplr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client WorkspaceManagedSQLServerRecoverableSqlpoolsClient) ListPreparer(ctx context.Context, resourceGroupName string, workspaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/recoverableSqlpools", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceManagedSQLServerRecoverableSqlpoolsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client WorkspaceManagedSQLServerRecoverableSqlpoolsClient) ListResponder(resp *http.Response) (result RecoverableSQLPoolListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client WorkspaceManagedSQLServerRecoverableSqlpoolsClient) listNextResults(ctx context.Context, lastResults RecoverableSQLPoolListResult) (result RecoverableSQLPoolListResult, err error) {
	req, err := lastResults.recoverableSQLPoolListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerRecoverableSqlpoolsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerRecoverableSqlpoolsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerRecoverableSqlpoolsClient", "listNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client WorkspaceManagedSQLServerRecoverableSqlpoolsClient) ListComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result RecoverableSQLPoolListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceManagedSQLServerRecoverableSqlpoolsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, workspaceName)
	return
}
