package redisenterprise

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

// Client is the REST API for managing Redis Enterprise resources in Azure.
type Client struct {
	BaseClient
}

// NewClient creates an instance of the Client client.
func NewClient(subscriptionID string) Client {
	return NewClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewClientWithBaseURI creates an instance of the Client client using a custom endpoint.  Use this when interacting
// with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewClientWithBaseURI(baseURI string, subscriptionID string) Client {
	return Client{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates or updates an existing (overwrite/recreate, with potential downtime) cache cluster
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the RedisEnterprise cluster.
// parameters - parameters supplied to the Create RedisEnterprise operation.
func (client Client) Create(ctx context.Context, resourceGroupName string, clusterName string, parameters Cluster) (result CreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Create")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Sku", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("redisenterprise.Client", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, clusterName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.Client", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.Client", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client Client) CreatePreparer(ctx context.Context, resourceGroupName string, clusterName string, parameters Cluster) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redisEnterprise/{clusterName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client Client) CreateSender(req *http.Request) (future CreateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client Client) CreateResponder(resp *http.Response) (result Cluster, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a RedisEnterprise cache cluster.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the RedisEnterprise cluster.
func (client Client) Delete(ctx context.Context, resourceGroupName string, clusterName string) (result DeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.Client", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.Client", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client Client) DeletePreparer(ctx context.Context, resourceGroupName string, clusterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redisEnterprise/{clusterName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client Client) DeleteSender(req *http.Request) (future DeleteFuture, err error) {
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
func (client Client) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetMethod gets information about a RedisEnterprise cluster
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the RedisEnterprise cluster.
func (client Client) GetMethod(ctx context.Context, resourceGroupName string, clusterName string) (result Cluster, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.GetMethod")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetMethodPreparer(ctx, resourceGroupName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.Client", "GetMethod", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetMethodSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "redisenterprise.Client", "GetMethod", resp, "Failure sending request")
		return
	}

	result, err = client.GetMethodResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.Client", "GetMethod", resp, "Failure responding to request")
		return
	}

	return
}

// GetMethodPreparer prepares the GetMethod request.
func (client Client) GetMethodPreparer(ctx context.Context, resourceGroupName string, clusterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redisEnterprise/{clusterName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetMethodSender sends the GetMethod request. The method will close the
// http.Response Body if it receives an error.
func (client Client) GetMethodSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetMethodResponder handles the response to the GetMethod request. The method always
// closes the http.Response Body.
func (client Client) GetMethodResponder(resp *http.Response) (result Cluster, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all RedisEnterprise clusters in the specified subscription.
func (client Client) List(ctx context.Context) (result ClusterListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.List")
		defer func() {
			sc := -1
			if result.cl.Response.Response != nil {
				sc = result.cl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.Client", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.cl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "redisenterprise.Client", "List", resp, "Failure sending request")
		return
	}

	result.cl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.Client", "List", resp, "Failure responding to request")
		return
	}
	if result.cl.hasNextLink() && result.cl.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client Client) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Cache/redisEnterprise", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client Client) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client Client) ListResponder(resp *http.Response) (result ClusterList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client Client) listNextResults(ctx context.Context, lastResults ClusterList) (result ClusterList, err error) {
	req, err := lastResults.clusterListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "redisenterprise.Client", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "redisenterprise.Client", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.Client", "listNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client Client) ListComplete(ctx context.Context) (result ClusterListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// ListByResourceGroup lists all RedisEnterprise clusters in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
func (client Client) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result ClusterListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.cl.Response.Response != nil {
				sc = result.cl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.Client", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.cl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "redisenterprise.Client", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.cl, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.Client", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.cl.hasNextLink() && result.cl.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client Client) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redisEnterprise", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client Client) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client Client) ListByResourceGroupResponder(resp *http.Response) (result ClusterList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client Client) listByResourceGroupNextResults(ctx context.Context, lastResults ClusterList) (result ClusterList, err error) {
	req, err := lastResults.clusterListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "redisenterprise.Client", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "redisenterprise.Client", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.Client", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client Client) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result ClusterListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// Update updates an existing RedisEnterprise cluster
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the RedisEnterprise cluster.
// parameters - parameters supplied to the Update RedisEnterprise operation.
func (client Client) Update(ctx context.Context, resourceGroupName string, clusterName string, parameters ClusterUpdate) (result UpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Update")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, clusterName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.Client", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.Client", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client Client) UpdatePreparer(ctx context.Context, resourceGroupName string, clusterName string, parameters ClusterUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redisEnterprise/{clusterName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client Client) UpdateSender(req *http.Request) (future UpdateFuture, err error) {
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
func (client Client) UpdateResponder(resp *http.Response) (result Cluster, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
