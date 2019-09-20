package machinelearningservices

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// WorkspaceFeaturesClient is the these APIs allow end users to operate on Azure Machine Learning Workspace resources.
type WorkspaceFeaturesClient struct {
	BaseClient
}

// NewWorkspaceFeaturesClient creates an instance of the WorkspaceFeaturesClient client.
func NewWorkspaceFeaturesClient(subscriptionID string) WorkspaceFeaturesClient {
	return NewWorkspaceFeaturesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkspaceFeaturesClientWithBaseURI creates an instance of the WorkspaceFeaturesClient client.
func NewWorkspaceFeaturesClientWithBaseURI(baseURI string, subscriptionID string) WorkspaceFeaturesClient {
	return WorkspaceFeaturesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists all enabled features for a workspace
// Parameters:
// resourceGroupName - name of the resource group in which workspace is located.
// workspaceName - name of Azure Machine Learning workspace.
func (client WorkspaceFeaturesClient) List(ctx context.Context, resourceGroupName string, workspaceName string) (result ListAmlUserFeatureResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceFeaturesClient.List")
		defer func() {
			sc := -1
			if result.laufr.Response.Response != nil {
				sc = result.laufr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, workspaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinelearningservices.WorkspaceFeaturesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.laufr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "machinelearningservices.WorkspaceFeaturesClient", "List", resp, "Failure sending request")
		return
	}

	result.laufr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinelearningservices.WorkspaceFeaturesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client WorkspaceFeaturesClient) ListPreparer(ctx context.Context, resourceGroupName string, workspaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/features", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceFeaturesClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client WorkspaceFeaturesClient) ListResponder(resp *http.Response) (result ListAmlUserFeatureResult, err error) {
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
func (client WorkspaceFeaturesClient) listNextResults(ctx context.Context, lastResults ListAmlUserFeatureResult) (result ListAmlUserFeatureResult, err error) {
	req, err := lastResults.listAmlUserFeatureResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "machinelearningservices.WorkspaceFeaturesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "machinelearningservices.WorkspaceFeaturesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinelearningservices.WorkspaceFeaturesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client WorkspaceFeaturesClient) ListComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result ListAmlUserFeatureResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceFeaturesClient.List")
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
