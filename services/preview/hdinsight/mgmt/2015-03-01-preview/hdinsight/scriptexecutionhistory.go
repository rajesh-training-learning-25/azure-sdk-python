package hdinsight

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
	"net/http"
)

// ScriptExecutionHistoryClient is the hDInsight Management Client
type ScriptExecutionHistoryClient struct {
	BaseClient
}

// NewScriptExecutionHistoryClient creates an instance of the ScriptExecutionHistoryClient client.
func NewScriptExecutionHistoryClient(subscriptionID string, resourceGroupName string, clusterName string, applicationName string, location string, configurationName string, extensionName string, scriptName string, scriptExecutionID string) ScriptExecutionHistoryClient {
	return NewScriptExecutionHistoryClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName, clusterName, applicationName, location, configurationName, extensionName, scriptName, scriptExecutionID)
}

// NewScriptExecutionHistoryClientWithBaseURI creates an instance of the ScriptExecutionHistoryClient client.
func NewScriptExecutionHistoryClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string, clusterName string, applicationName string, location string, configurationName string, extensionName string, scriptName string, scriptExecutionID string) ScriptExecutionHistoryClient {
	return ScriptExecutionHistoryClient{NewWithBaseURI(baseURI, subscriptionID, resourceGroupName, clusterName, applicationName, location, configurationName, extensionName, scriptName, scriptExecutionID)}
}

// List lists all scripts' execution history for the specified cluster.
func (client ScriptExecutionHistoryClient) List(ctx context.Context) (result ScriptActionExecutionHistoryListPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ScriptExecutionHistoryClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.saehl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hdinsight.ScriptExecutionHistoryClient", "List", resp, "Failure sending request")
		return
	}

	result.saehl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ScriptExecutionHistoryClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ScriptExecutionHistoryClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", client.ClusterName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/scriptExecutionHistory", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ScriptExecutionHistoryClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ScriptExecutionHistoryClient) ListResponder(resp *http.Response) (result ScriptActionExecutionHistoryList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ScriptExecutionHistoryClient) listNextResults(lastResults ScriptActionExecutionHistoryList) (result ScriptActionExecutionHistoryList, err error) {
	req, err := lastResults.scriptActionExecutionHistoryListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "hdinsight.ScriptExecutionHistoryClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "hdinsight.ScriptExecutionHistoryClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ScriptExecutionHistoryClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ScriptExecutionHistoryClient) ListComplete(ctx context.Context) (result ScriptActionExecutionHistoryListIterator, err error) {
	result.page, err = client.List(ctx)
	return
}

// Promote promotes the specified ad-hoc script execution to a persisted script.
func (client ScriptExecutionHistoryClient) Promote(ctx context.Context) (result autorest.Response, err error) {
	req, err := client.PromotePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ScriptExecutionHistoryClient", "Promote", nil, "Failure preparing request")
		return
	}

	resp, err := client.PromoteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "hdinsight.ScriptExecutionHistoryClient", "Promote", resp, "Failure sending request")
		return
	}

	result, err = client.PromoteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ScriptExecutionHistoryClient", "Promote", resp, "Failure responding to request")
	}

	return
}

// PromotePreparer prepares the Promote request.
func (client ScriptExecutionHistoryClient) PromotePreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", client.ClusterName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"scriptExecutionId": autorest.Encode("path", client.ScriptExecutionID),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/scriptExecutionHistory/{scriptExecutionId}/promote", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PromoteSender sends the Promote request. The method will close the
// http.Response Body if it receives an error.
func (client ScriptExecutionHistoryClient) PromoteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// PromoteResponder handles the response to the Promote request. The method always
// closes the http.Response Body.
func (client ScriptExecutionHistoryClient) PromoteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
