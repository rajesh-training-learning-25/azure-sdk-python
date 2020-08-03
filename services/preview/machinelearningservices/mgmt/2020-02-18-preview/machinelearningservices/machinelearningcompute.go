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

// MachineLearningComputeClient is the these APIs allow end users to operate on Azure Machine Learning Workspace
// resources.
type MachineLearningComputeClient struct {
	BaseClient
}

// NewMachineLearningComputeClient creates an instance of the MachineLearningComputeClient client.
func NewMachineLearningComputeClient(subscriptionID string) MachineLearningComputeClient {
	return NewMachineLearningComputeClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewMachineLearningComputeClientWithBaseURI creates an instance of the MachineLearningComputeClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewMachineLearningComputeClientWithBaseURI(baseURI string, subscriptionID string) MachineLearningComputeClient {
	return MachineLearningComputeClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates compute. This call will overwrite a compute if it exists. This is a nonrecoverable
// operation. If your intent is to create a new compute, do a GET first to verify that it does not exist yet.
// Parameters:
// resourceGroupName - name of the resource group in which workspace is located.
// workspaceName - name of Azure Machine Learning workspace.
// computeName - name of the Azure Machine Learning compute.
// parameters - payload with Machine Learning compute definition.
func (client MachineLearningComputeClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, parameters ComputeResource) (result MachineLearningComputeCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MachineLearningComputeClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, workspaceName, computeName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinelearningservices.MachineLearningComputeClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinelearningservices.MachineLearningComputeClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client MachineLearningComputeClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, parameters ComputeResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"computeName":       autorest.Encode("path", computeName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-02-18-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client MachineLearningComputeClient) CreateOrUpdateSender(req *http.Request) (future MachineLearningComputeCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client MachineLearningComputeClient) CreateOrUpdateResponder(resp *http.Response) (result ComputeResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes specified Machine Learning compute.
// Parameters:
// resourceGroupName - name of the resource group in which workspace is located.
// workspaceName - name of Azure Machine Learning workspace.
// computeName - name of the Azure Machine Learning compute.
// underlyingResourceAction - delete the underlying compute if 'Delete', or detach the underlying compute from
// workspace if 'Detach'.
func (client MachineLearningComputeClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, underlyingResourceAction UnderlyingResourceAction) (result MachineLearningComputeDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MachineLearningComputeClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, workspaceName, computeName, underlyingResourceAction)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinelearningservices.MachineLearningComputeClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinelearningservices.MachineLearningComputeClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client MachineLearningComputeClient) DeletePreparer(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, underlyingResourceAction UnderlyingResourceAction) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"computeName":       autorest.Encode("path", computeName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-02-18-preview"
	queryParameters := map[string]interface{}{
		"api-version":              APIVersion,
		"underlyingResourceAction": autorest.Encode("query", underlyingResourceAction),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client MachineLearningComputeClient) DeleteSender(req *http.Request) (future MachineLearningComputeDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client MachineLearningComputeClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets compute definition by its name. Any secrets (storage keys, service credentials, etc) are not returned - use
// 'keys' nested resource to get them.
// Parameters:
// resourceGroupName - name of the resource group in which workspace is located.
// workspaceName - name of Azure Machine Learning workspace.
// computeName - name of the Azure Machine Learning compute.
func (client MachineLearningComputeClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, computeName string) (result ComputeResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MachineLearningComputeClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName, computeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinelearningservices.MachineLearningComputeClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "machinelearningservices.MachineLearningComputeClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinelearningservices.MachineLearningComputeClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client MachineLearningComputeClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string, computeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"computeName":       autorest.Encode("path", computeName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-02-18-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client MachineLearningComputeClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client MachineLearningComputeClient) GetResponder(resp *http.Response) (result ComputeResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByWorkspace gets computes in specified workspace.
// Parameters:
// resourceGroupName - name of the resource group in which workspace is located.
// workspaceName - name of Azure Machine Learning workspace.
// skiptoken - continuation token for pagination.
func (client MachineLearningComputeClient) ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string, skiptoken string) (result PaginatedComputeResourcesListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MachineLearningComputeClient.ListByWorkspace")
		defer func() {
			sc := -1
			if result.pcrl.Response.Response != nil {
				sc = result.pcrl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByWorkspaceNextResults
	req, err := client.ListByWorkspacePreparer(ctx, resourceGroupName, workspaceName, skiptoken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinelearningservices.MachineLearningComputeClient", "ListByWorkspace", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByWorkspaceSender(req)
	if err != nil {
		result.pcrl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "machinelearningservices.MachineLearningComputeClient", "ListByWorkspace", resp, "Failure sending request")
		return
	}

	result.pcrl, err = client.ListByWorkspaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinelearningservices.MachineLearningComputeClient", "ListByWorkspace", resp, "Failure responding to request")
	}
	if result.pcrl.hasNextLink() && result.pcrl.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByWorkspacePreparer prepares the ListByWorkspace request.
func (client MachineLearningComputeClient) ListByWorkspacePreparer(ctx context.Context, resourceGroupName string, workspaceName string, skiptoken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-02-18-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByWorkspaceSender sends the ListByWorkspace request. The method will close the
// http.Response Body if it receives an error.
func (client MachineLearningComputeClient) ListByWorkspaceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByWorkspaceResponder handles the response to the ListByWorkspace request. The method always
// closes the http.Response Body.
func (client MachineLearningComputeClient) ListByWorkspaceResponder(resp *http.Response) (result PaginatedComputeResourcesList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByWorkspaceNextResults retrieves the next set of results, if any.
func (client MachineLearningComputeClient) listByWorkspaceNextResults(ctx context.Context, lastResults PaginatedComputeResourcesList) (result PaginatedComputeResourcesList, err error) {
	req, err := lastResults.paginatedComputeResourcesListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "machinelearningservices.MachineLearningComputeClient", "listByWorkspaceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByWorkspaceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "machinelearningservices.MachineLearningComputeClient", "listByWorkspaceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByWorkspaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinelearningservices.MachineLearningComputeClient", "listByWorkspaceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByWorkspaceComplete enumerates all values, automatically crossing page boundaries as required.
func (client MachineLearningComputeClient) ListByWorkspaceComplete(ctx context.Context, resourceGroupName string, workspaceName string, skiptoken string) (result PaginatedComputeResourcesListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MachineLearningComputeClient.ListByWorkspace")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByWorkspace(ctx, resourceGroupName, workspaceName, skiptoken)
	return
}

// ListKeys gets secrets related to Machine Learning compute (storage keys, service credentials, etc).
// Parameters:
// resourceGroupName - name of the resource group in which workspace is located.
// workspaceName - name of Azure Machine Learning workspace.
// computeName - name of the Azure Machine Learning compute.
func (client MachineLearningComputeClient) ListKeys(ctx context.Context, resourceGroupName string, workspaceName string, computeName string) (result ComputeSecretsModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MachineLearningComputeClient.ListKeys")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListKeysPreparer(ctx, resourceGroupName, workspaceName, computeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinelearningservices.MachineLearningComputeClient", "ListKeys", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListKeysSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "machinelearningservices.MachineLearningComputeClient", "ListKeys", resp, "Failure sending request")
		return
	}

	result, err = client.ListKeysResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinelearningservices.MachineLearningComputeClient", "ListKeys", resp, "Failure responding to request")
	}

	return
}

// ListKeysPreparer prepares the ListKeys request.
func (client MachineLearningComputeClient) ListKeysPreparer(ctx context.Context, resourceGroupName string, workspaceName string, computeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"computeName":       autorest.Encode("path", computeName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-02-18-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}/listKeys", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListKeysSender sends the ListKeys request. The method will close the
// http.Response Body if it receives an error.
func (client MachineLearningComputeClient) ListKeysSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListKeysResponder handles the response to the ListKeys request. The method always
// closes the http.Response Body.
func (client MachineLearningComputeClient) ListKeysResponder(resp *http.Response) (result ComputeSecretsModel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNodes get the details (e.g IP address, port etc) of all the compute nodes in the compute.
// Parameters:
// resourceGroupName - name of the resource group in which workspace is located.
// workspaceName - name of Azure Machine Learning workspace.
// computeName - name of the Azure Machine Learning compute.
func (client MachineLearningComputeClient) ListNodes(ctx context.Context, resourceGroupName string, workspaceName string, computeName string) (result AmlComputeNodesInformation, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MachineLearningComputeClient.ListNodes")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListNodesPreparer(ctx, resourceGroupName, workspaceName, computeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinelearningservices.MachineLearningComputeClient", "ListNodes", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListNodesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "machinelearningservices.MachineLearningComputeClient", "ListNodes", resp, "Failure sending request")
		return
	}

	result, err = client.ListNodesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinelearningservices.MachineLearningComputeClient", "ListNodes", resp, "Failure responding to request")
	}

	return
}

// ListNodesPreparer prepares the ListNodes request.
func (client MachineLearningComputeClient) ListNodesPreparer(ctx context.Context, resourceGroupName string, workspaceName string, computeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"computeName":       autorest.Encode("path", computeName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-02-18-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}/listNodes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListNodesSender sends the ListNodes request. The method will close the
// http.Response Body if it receives an error.
func (client MachineLearningComputeClient) ListNodesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListNodesResponder handles the response to the ListNodes request. The method always
// closes the http.Response Body.
func (client MachineLearningComputeClient) ListNodesResponder(resp *http.Response) (result AmlComputeNodesInformation, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Restart posts a restart action to a compute instance
// Parameters:
// resourceGroupName - name of the resource group in which workspace is located.
// workspaceName - name of Azure Machine Learning workspace.
// computeName - name of the Azure Machine Learning compute.
func (client MachineLearningComputeClient) Restart(ctx context.Context, resourceGroupName string, workspaceName string, computeName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MachineLearningComputeClient.Restart")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RestartPreparer(ctx, resourceGroupName, workspaceName, computeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinelearningservices.MachineLearningComputeClient", "Restart", nil, "Failure preparing request")
		return
	}

	resp, err := client.RestartSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "machinelearningservices.MachineLearningComputeClient", "Restart", resp, "Failure sending request")
		return
	}

	result, err = client.RestartResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinelearningservices.MachineLearningComputeClient", "Restart", resp, "Failure responding to request")
	}

	return
}

// RestartPreparer prepares the Restart request.
func (client MachineLearningComputeClient) RestartPreparer(ctx context.Context, resourceGroupName string, workspaceName string, computeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"computeName":       autorest.Encode("path", computeName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-02-18-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}/restart", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RestartSender sends the Restart request. The method will close the
// http.Response Body if it receives an error.
func (client MachineLearningComputeClient) RestartSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// RestartResponder handles the response to the Restart request. The method always
// closes the http.Response Body.
func (client MachineLearningComputeClient) RestartResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Start posts a start action to a compute instance
// Parameters:
// resourceGroupName - name of the resource group in which workspace is located.
// workspaceName - name of Azure Machine Learning workspace.
// computeName - name of the Azure Machine Learning compute.
func (client MachineLearningComputeClient) Start(ctx context.Context, resourceGroupName string, workspaceName string, computeName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MachineLearningComputeClient.Start")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StartPreparer(ctx, resourceGroupName, workspaceName, computeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinelearningservices.MachineLearningComputeClient", "Start", nil, "Failure preparing request")
		return
	}

	resp, err := client.StartSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "machinelearningservices.MachineLearningComputeClient", "Start", resp, "Failure sending request")
		return
	}

	result, err = client.StartResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinelearningservices.MachineLearningComputeClient", "Start", resp, "Failure responding to request")
	}

	return
}

// StartPreparer prepares the Start request.
func (client MachineLearningComputeClient) StartPreparer(ctx context.Context, resourceGroupName string, workspaceName string, computeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"computeName":       autorest.Encode("path", computeName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-02-18-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}/start", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StartSender sends the Start request. The method will close the
// http.Response Body if it receives an error.
func (client MachineLearningComputeClient) StartSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// StartResponder handles the response to the Start request. The method always
// closes the http.Response Body.
func (client MachineLearningComputeClient) StartResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Stop posts a stop action to a compute instance
// Parameters:
// resourceGroupName - name of the resource group in which workspace is located.
// workspaceName - name of Azure Machine Learning workspace.
// computeName - name of the Azure Machine Learning compute.
func (client MachineLearningComputeClient) Stop(ctx context.Context, resourceGroupName string, workspaceName string, computeName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MachineLearningComputeClient.Stop")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StopPreparer(ctx, resourceGroupName, workspaceName, computeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinelearningservices.MachineLearningComputeClient", "Stop", nil, "Failure preparing request")
		return
	}

	resp, err := client.StopSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "machinelearningservices.MachineLearningComputeClient", "Stop", resp, "Failure sending request")
		return
	}

	result, err = client.StopResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinelearningservices.MachineLearningComputeClient", "Stop", resp, "Failure responding to request")
	}

	return
}

// StopPreparer prepares the Stop request.
func (client MachineLearningComputeClient) StopPreparer(ctx context.Context, resourceGroupName string, workspaceName string, computeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"computeName":       autorest.Encode("path", computeName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-02-18-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}/stop", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StopSender sends the Stop request. The method will close the
// http.Response Body if it receives an error.
func (client MachineLearningComputeClient) StopSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// StopResponder handles the response to the Stop request. The method always
// closes the http.Response Body.
func (client MachineLearningComputeClient) StopResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update updates properties of a compute. This call will overwrite a compute if it exists. This is a nonrecoverable
// operation.
// Parameters:
// resourceGroupName - name of the resource group in which workspace is located.
// workspaceName - name of Azure Machine Learning workspace.
// computeName - name of the Azure Machine Learning compute.
// parameters - additional parameters for cluster update.
func (client MachineLearningComputeClient) Update(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, parameters ClusterUpdateParameters) (result MachineLearningComputeUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MachineLearningComputeClient.Update")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, workspaceName, computeName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinelearningservices.MachineLearningComputeClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinelearningservices.MachineLearningComputeClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client MachineLearningComputeClient) UpdatePreparer(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, parameters ClusterUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"computeName":       autorest.Encode("path", computeName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-02-18-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client MachineLearningComputeClient) UpdateSender(req *http.Request) (future MachineLearningComputeUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client MachineLearningComputeClient) UpdateResponder(resp *http.Response) (result ComputeResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
