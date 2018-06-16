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

// ExtensionClient is the hDInsight Management Client
type ExtensionClient struct {
	BaseClient
}

// NewExtensionClient creates an instance of the ExtensionClient client.
func NewExtensionClient(subscriptionID string, resourceGroupName string, clusterName string, applicationName string, location string, configurationName string, extensionName string, scriptName string, scriptExecutionID string) ExtensionClient {
	return NewExtensionClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName, clusterName, applicationName, location, configurationName, extensionName, scriptName, scriptExecutionID)
}

// NewExtensionClientWithBaseURI creates an instance of the ExtensionClient client.
func NewExtensionClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string, clusterName string, applicationName string, location string, configurationName string, extensionName string, scriptName string, scriptExecutionID string) ExtensionClient {
	return ExtensionClient{NewWithBaseURI(baseURI, subscriptionID, resourceGroupName, clusterName, applicationName, location, configurationName, extensionName, scriptName, scriptExecutionID)}
}

// Create creates an HDInsight cluster extension.
// Parameters:
// parameters - the cluster extensions create request.
func (client ExtensionClient) Create(ctx context.Context, parameters Extension) (result autorest.Response, err error) {
	req, err := client.CreatePreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ExtensionClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "hdinsight.ExtensionClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ExtensionClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ExtensionClient) CreatePreparer(ctx context.Context, parameters Extension) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", client.ClusterName),
		"extensionName":     autorest.Encode("path", client.ExtensionName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/{extensionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ExtensionClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ExtensionClient) CreateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete deletes the specified extension for HDInsight cluster.
func (client ExtensionClient) Delete(ctx context.Context) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ExtensionClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "hdinsight.ExtensionClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ExtensionClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ExtensionClient) DeletePreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", client.ClusterName),
		"extensionName":     autorest.Encode("path", client.ExtensionName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/{extensionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ExtensionClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ExtensionClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DisableMonitoring disables the Operations Management Suite (OMS) on the HDInsight cluster.
func (client ExtensionClient) DisableMonitoring(ctx context.Context) (result ExtensionDisableMonitoringFuture, err error) {
	req, err := client.DisableMonitoringPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ExtensionClient", "DisableMonitoring", nil, "Failure preparing request")
		return
	}

	result, err = client.DisableMonitoringSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ExtensionClient", "DisableMonitoring", result.Response(), "Failure sending request")
		return
	}

	return
}

// DisableMonitoringPreparer prepares the DisableMonitoring request.
func (client ExtensionClient) DisableMonitoringPreparer(ctx context.Context) (*http.Request, error) {
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
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/clustermonitoring", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DisableMonitoringSender sends the DisableMonitoring request. The method will close the
// http.Response Body if it receives an error.
func (client ExtensionClient) DisableMonitoringSender(req *http.Request) (future ExtensionDisableMonitoringFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	err = autorest.Respond(resp, azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DisableMonitoringResponder handles the response to the DisableMonitoring request. The method always
// closes the http.Response Body.
func (client ExtensionClient) DisableMonitoringResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// EnableMonitoring enables the Operations Management Suite (OMS) on the HDInsight cluster.
// Parameters:
// parameters - the Operations Management Suite (OMS) workspace parameters.
func (client ExtensionClient) EnableMonitoring(ctx context.Context, parameters ClusterMonitoringRequest) (result ExtensionEnableMonitoringFuture, err error) {
	req, err := client.EnableMonitoringPreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ExtensionClient", "EnableMonitoring", nil, "Failure preparing request")
		return
	}

	result, err = client.EnableMonitoringSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ExtensionClient", "EnableMonitoring", result.Response(), "Failure sending request")
		return
	}

	return
}

// EnableMonitoringPreparer prepares the EnableMonitoring request.
func (client ExtensionClient) EnableMonitoringPreparer(ctx context.Context, parameters ClusterMonitoringRequest) (*http.Request, error) {
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
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/clustermonitoring", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// EnableMonitoringSender sends the EnableMonitoring request. The method will close the
// http.Response Body if it receives an error.
func (client ExtensionClient) EnableMonitoringSender(req *http.Request) (future ExtensionEnableMonitoringFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	err = autorest.Respond(resp, azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// EnableMonitoringResponder handles the response to the EnableMonitoring request. The method always
// closes the http.Response Body.
func (client ExtensionClient) EnableMonitoringResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the extension properties for the specified HDInsight cluster extension.
func (client ExtensionClient) Get(ctx context.Context) (result Extension, err error) {
	req, err := client.GetPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ExtensionClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hdinsight.ExtensionClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ExtensionClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ExtensionClient) GetPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", client.ClusterName),
		"extensionName":     autorest.Encode("path", client.ExtensionName),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/{extensionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ExtensionClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ExtensionClient) GetResponder(resp *http.Response) (result Extension, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetMonitoringStatus gets the status of Operations Management Suite (OMS) on the HDInsight cluster.
func (client ExtensionClient) GetMonitoringStatus(ctx context.Context) (result ClusterMonitoringResponse, err error) {
	req, err := client.GetMonitoringStatusPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ExtensionClient", "GetMonitoringStatus", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetMonitoringStatusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hdinsight.ExtensionClient", "GetMonitoringStatus", resp, "Failure sending request")
		return
	}

	result, err = client.GetMonitoringStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ExtensionClient", "GetMonitoringStatus", resp, "Failure responding to request")
	}

	return
}

// GetMonitoringStatusPreparer prepares the GetMonitoringStatus request.
func (client ExtensionClient) GetMonitoringStatusPreparer(ctx context.Context) (*http.Request, error) {
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/clustermonitoring", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetMonitoringStatusSender sends the GetMonitoringStatus request. The method will close the
// http.Response Body if it receives an error.
func (client ExtensionClient) GetMonitoringStatusSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetMonitoringStatusResponder handles the response to the GetMonitoringStatus request. The method always
// closes the http.Response Body.
func (client ExtensionClient) GetMonitoringStatusResponder(resp *http.Response) (result ClusterMonitoringResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
