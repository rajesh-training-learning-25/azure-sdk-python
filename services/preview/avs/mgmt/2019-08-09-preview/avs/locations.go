package avs

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

// LocationsClient is the azure VMware Solution API
type LocationsClient struct {
	BaseClient
}

// NewLocationsClient creates an instance of the LocationsClient client.
func NewLocationsClient(subscriptionID string) LocationsClient {
	return NewLocationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLocationsClientWithBaseURI creates an instance of the LocationsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewLocationsClientWithBaseURI(baseURI string, subscriptionID string) LocationsClient {
	return LocationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CheckQuotaAvailability return quota for subscription by region
// Parameters:
// location - azure region
func (client LocationsClient) CheckQuotaAvailability(ctx context.Context, location string) (result Quota, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LocationsClient.CheckQuotaAvailability")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CheckQuotaAvailabilityPreparer(ctx, location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "avs.LocationsClient", "CheckQuotaAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckQuotaAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "avs.LocationsClient", "CheckQuotaAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckQuotaAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "avs.LocationsClient", "CheckQuotaAvailability", resp, "Failure responding to request")
		return
	}

	return
}

// CheckQuotaAvailabilityPreparer prepares the CheckQuotaAvailability request.
func (client LocationsClient) CheckQuotaAvailabilityPreparer(ctx context.Context, location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-09-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.AVS/locations/{location}/checkQuotaAvailability", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckQuotaAvailabilitySender sends the CheckQuotaAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client LocationsClient) CheckQuotaAvailabilitySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CheckQuotaAvailabilityResponder handles the response to the CheckQuotaAvailability request. The method always
// closes the http.Response Body.
func (client LocationsClient) CheckQuotaAvailabilityResponder(resp *http.Response) (result Quota, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CheckTrialAvailability return trial status for subscription by region
// Parameters:
// location - azure region
func (client LocationsClient) CheckTrialAvailability(ctx context.Context, location string) (result Trial, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LocationsClient.CheckTrialAvailability")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CheckTrialAvailabilityPreparer(ctx, location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "avs.LocationsClient", "CheckTrialAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckTrialAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "avs.LocationsClient", "CheckTrialAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckTrialAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "avs.LocationsClient", "CheckTrialAvailability", resp, "Failure responding to request")
		return
	}

	return
}

// CheckTrialAvailabilityPreparer prepares the CheckTrialAvailability request.
func (client LocationsClient) CheckTrialAvailabilityPreparer(ctx context.Context, location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-09-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.AVS/locations/{location}/checkTrialAvailability", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckTrialAvailabilitySender sends the CheckTrialAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client LocationsClient) CheckTrialAvailabilitySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CheckTrialAvailabilityResponder handles the response to the CheckTrialAvailability request. The method always
// closes the http.Response Body.
func (client LocationsClient) CheckTrialAvailabilityResponder(resp *http.Response) (result Trial, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
