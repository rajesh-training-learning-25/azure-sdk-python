package operationalinsights

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

// WorkspaceClient is the operational Insights Client
type WorkspaceClient struct {
	BaseClient
}

// NewWorkspaceClient creates an instance of the WorkspaceClient client.
func NewWorkspaceClient(subscriptionID string, purgeID string) WorkspaceClient {
	return NewWorkspaceClientWithBaseURI(DefaultBaseURI, subscriptionID, purgeID)
}

// NewWorkspaceClientWithBaseURI creates an instance of the WorkspaceClient client.
func NewWorkspaceClientWithBaseURI(baseURI string, subscriptionID string, purgeID string) WorkspaceClient {
	return WorkspaceClient{NewWithBaseURI(baseURI, subscriptionID, purgeID)}
}

// Purge purges data in an Log Analytics workspace by a set of user-defined filters.
//
// resourceGroupName is the name of the resource group to get. The name is case insensitive. workspaceName is log
// Analytics workspace name body is describes the body of a request to purge data in a single table of an Log
// Analytics Workspace
func (client WorkspaceClient) Purge(ctx context.Context, resourceGroupName string, workspaceName string, body WorkspacePurgeBody) (result WorkspacePurgeFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.Table", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "body.Filters", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("operationalinsights.WorkspaceClient", "Purge", err.Error())
	}

	req, err := client.PurgePreparer(ctx, resourceGroupName, workspaceName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.WorkspaceClient", "Purge", nil, "Failure preparing request")
		return
	}

	result, err = client.PurgeSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.WorkspaceClient", "Purge", result.Response(), "Failure sending request")
		return
	}

	return
}

// PurgePreparer prepares the Purge request.
func (client WorkspaceClient) PurgePreparer(ctx context.Context, resourceGroupName string, workspaceName string, body WorkspacePurgeBody) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2015-03-20"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/purge", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PurgeSender sends the Purge request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceClient) PurgeSender(req *http.Request) (future WorkspacePurgeFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// PurgeResponder handles the response to the Purge request. The method always
// closes the http.Response Body.
func (client WorkspaceClient) PurgeResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
