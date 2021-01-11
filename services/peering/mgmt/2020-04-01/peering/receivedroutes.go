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

// ReceivedRoutesClient is the peering Client
type ReceivedRoutesClient struct {
	BaseClient
}

// NewReceivedRoutesClient creates an instance of the ReceivedRoutesClient client.
func NewReceivedRoutesClient(subscriptionID string) ReceivedRoutesClient {
	return NewReceivedRoutesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewReceivedRoutesClientWithBaseURI creates an instance of the ReceivedRoutesClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewReceivedRoutesClientWithBaseURI(baseURI string, subscriptionID string) ReceivedRoutesClient {
	return ReceivedRoutesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByPeering lists the prefixes received over the specified peering under the given subscription and resource
// group.
// Parameters:
// resourceGroupName - the name of the resource group.
// peeringName - the name of the peering.
// prefix - the optional prefix that can be used to filter the routes.
// asPath - the optional AS path that can be used to filter the routes.
// originAsValidationState - the optional origin AS validation state that can be used to filter the routes.
// rpkiValidationState - the optional RPKI validation state that can be used to filter the routes.
// skipToken - the optional page continuation token that is used in the event of paginated result.
func (client ReceivedRoutesClient) ListByPeering(ctx context.Context, resourceGroupName string, peeringName string, prefix string, asPath string, originAsValidationState string, rpkiValidationState string, skipToken string) (result ReceivedRouteListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReceivedRoutesClient.ListByPeering")
		defer func() {
			sc := -1
			if result.rrlr.Response.Response != nil {
				sc = result.rrlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByPeeringNextResults
	req, err := client.ListByPeeringPreparer(ctx, resourceGroupName, peeringName, prefix, asPath, originAsValidationState, rpkiValidationState, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.ReceivedRoutesClient", "ListByPeering", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByPeeringSender(req)
	if err != nil {
		result.rrlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "peering.ReceivedRoutesClient", "ListByPeering", resp, "Failure sending request")
		return
	}

	result.rrlr, err = client.ListByPeeringResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.ReceivedRoutesClient", "ListByPeering", resp, "Failure responding to request")
		return
	}
	if result.rrlr.hasNextLink() && result.rrlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByPeeringPreparer prepares the ListByPeering request.
func (client ReceivedRoutesClient) ListByPeeringPreparer(ctx context.Context, resourceGroupName string, peeringName string, prefix string, asPath string, originAsValidationState string, rpkiValidationState string, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"peeringName":       autorest.Encode("path", peeringName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(prefix) > 0 {
		queryParameters["prefix"] = autorest.Encode("query", prefix)
	}
	if len(asPath) > 0 {
		queryParameters["asPath"] = autorest.Encode("query", asPath)
	}
	if len(originAsValidationState) > 0 {
		queryParameters["originAsValidationState"] = autorest.Encode("query", originAsValidationState)
	}
	if len(rpkiValidationState) > 0 {
		queryParameters["rpkiValidationState"] = autorest.Encode("query", rpkiValidationState)
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}/receivedRoutes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByPeeringSender sends the ListByPeering request. The method will close the
// http.Response Body if it receives an error.
func (client ReceivedRoutesClient) ListByPeeringSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByPeeringResponder handles the response to the ListByPeering request. The method always
// closes the http.Response Body.
func (client ReceivedRoutesClient) ListByPeeringResponder(resp *http.Response) (result ReceivedRouteListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByPeeringNextResults retrieves the next set of results, if any.
func (client ReceivedRoutesClient) listByPeeringNextResults(ctx context.Context, lastResults ReceivedRouteListResult) (result ReceivedRouteListResult, err error) {
	req, err := lastResults.receivedRouteListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "peering.ReceivedRoutesClient", "listByPeeringNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByPeeringSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "peering.ReceivedRoutesClient", "listByPeeringNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByPeeringResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.ReceivedRoutesClient", "listByPeeringNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByPeeringComplete enumerates all values, automatically crossing page boundaries as required.
func (client ReceivedRoutesClient) ListByPeeringComplete(ctx context.Context, resourceGroupName string, peeringName string, prefix string, asPath string, originAsValidationState string, rpkiValidationState string, skipToken string) (result ReceivedRouteListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReceivedRoutesClient.ListByPeering")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByPeering(ctx, resourceGroupName, peeringName, prefix, asPath, originAsValidationState, rpkiValidationState, skipToken)
	return
}
