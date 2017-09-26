package redis

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
// Code generated by Microsoft (R) AutoRest Code Generator 2.2.22.0
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// LinkedServerClient is the REST API for Azure Redis Cache Service.
type LinkedServerClient struct {
	ManagementClient
}

// NewLinkedServerClient creates an instance of the LinkedServerClient client.
func NewLinkedServerClient(subscriptionID string) LinkedServerClient {
	return NewLinkedServerClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLinkedServerClientWithBaseURI creates an instance of the LinkedServerClient client.
func NewLinkedServerClientWithBaseURI(baseURI string, subscriptionID string) LinkedServerClient {
	return LinkedServerClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create adds a linked server to the Redis cache (requires Premium SKU). This method may poll for completion. Polling
// can be canceled by passing the cancel channel argument. The channel will be used to cancel polling and any
// outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. name is the name of the Redis cache. linkedServerName is the
// name of the linked server that is being added to the Redis cache. parameters is parameters supplied to the Create
// Linked server operation.
func (client LinkedServerClient) Create(resourceGroupName string, name string, linkedServerName string, parameters LinkedServerCreateParameters, cancel <-chan struct{}) (<-chan LinkedServerWithProperties, <-chan error) {
	resultChan := make(chan LinkedServerWithProperties, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.LinkedServerCreateProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.LinkedServerCreateProperties.LinkedRedisCacheID", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.LinkedServerCreateProperties.LinkedRedisCacheLocation", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "redis.LinkedServerClient", "Create")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result LinkedServerWithProperties
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreatePreparer(resourceGroupName, name, linkedServerName, parameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "redis.LinkedServerClient", "Create", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "redis.LinkedServerClient", "Create", resp, "Failure sending request")
			return
		}

		result, err = client.CreateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "redis.LinkedServerClient", "Create", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreatePreparer prepares the Create request.
func (client LinkedServerClient) CreatePreparer(resourceGroupName string, name string, linkedServerName string, parameters LinkedServerCreateParameters, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"linkedServerName":  autorest.Encode("path", linkedServerName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/linkedServers/{linkedServerName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client LinkedServerClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client LinkedServerClient) CreateResponder(resp *http.Response) (result LinkedServerWithProperties, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the linked server from a redis cache (requires Premium SKU).
//
// resourceGroupName is the name of the resource group. name is the name of the redis cache. linkedServerName is the
// name of the linked server that is being added to the Redis cache.
func (client LinkedServerClient) Delete(resourceGroupName string, name string, linkedServerName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(resourceGroupName, name, linkedServerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.LinkedServerClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "redis.LinkedServerClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.LinkedServerClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client LinkedServerClient) DeletePreparer(resourceGroupName string, name string, linkedServerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"linkedServerName":  autorest.Encode("path", linkedServerName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/linkedServers/{linkedServerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client LinkedServerClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client LinkedServerClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the detailed information about a linked server of a redis cache (requires Premium SKU).
//
// resourceGroupName is the name of the resource group. name is the name of the redis cache. linkedServerName is the
// name of the linked server.
func (client LinkedServerClient) Get(resourceGroupName string, name string, linkedServerName string) (result LinkedServerWithProperties, err error) {
	req, err := client.GetPreparer(resourceGroupName, name, linkedServerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.LinkedServerClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "redis.LinkedServerClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.LinkedServerClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client LinkedServerClient) GetPreparer(resourceGroupName string, name string, linkedServerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"linkedServerName":  autorest.Encode("path", linkedServerName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/linkedServers/{linkedServerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client LinkedServerClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client LinkedServerClient) GetResponder(resp *http.Response) (result LinkedServerWithProperties, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets the list of linked servers associated with this redis cache (requires Premium SKU).
//
// resourceGroupName is the name of the resource group. name is the name of the redis cache.
func (client LinkedServerClient) List(resourceGroupName string, name string) (result LinkedServerWithPropertiesList, err error) {
	req, err := client.ListPreparer(resourceGroupName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.LinkedServerClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "redis.LinkedServerClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.LinkedServerClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client LinkedServerClient) ListPreparer(resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/linkedServers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client LinkedServerClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client LinkedServerClient) ListResponder(resp *http.Response) (result LinkedServerWithPropertiesList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
