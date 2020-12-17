package peering

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

// ServiceCountriesClient is the peering Client
type ServiceCountriesClient struct {
	BaseClient
}

// NewServiceCountriesClient creates an instance of the ServiceCountriesClient client.
func NewServiceCountriesClient(subscriptionID string) ServiceCountriesClient {
	return NewServiceCountriesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewServiceCountriesClientWithBaseURI creates an instance of the ServiceCountriesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewServiceCountriesClientWithBaseURI(baseURI string, subscriptionID string) ServiceCountriesClient {
	return ServiceCountriesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists all of the available countries for peering service.
func (client ServiceCountriesClient) List(ctx context.Context) (result ServiceCountryListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceCountriesClient.List")
		defer func() {
			sc := -1
			if result.sclr.Response.Response != nil {
				sc = result.sclr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.ServiceCountriesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.sclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "peering.ServiceCountriesClient", "List", resp, "Failure sending request")
		return
	}

	result.sclr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.ServiceCountriesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.sclr.hasNextLink() && result.sclr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client ServiceCountriesClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Peering/peeringServiceCountries", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceCountriesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ServiceCountriesClient) ListResponder(resp *http.Response) (result ServiceCountryListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ServiceCountriesClient) listNextResults(ctx context.Context, lastResults ServiceCountryListResult) (result ServiceCountryListResult, err error) {
	req, err := lastResults.serviceCountryListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "peering.ServiceCountriesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "peering.ServiceCountriesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.ServiceCountriesClient", "listNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ServiceCountriesClient) ListComplete(ctx context.Context) (result ServiceCountryListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceCountriesClient.List")
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
