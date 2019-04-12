package speechservices

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
	"github.com/satori/go.uuid"
	"net/http"
)

// EndpointDataClient is the speech Services API v2.0.
type EndpointDataClient struct {
	BaseClient
}

// NewEndpointDataClient creates an instance of the EndpointDataClient client.
func NewEndpointDataClient() EndpointDataClient {
	return NewEndpointDataClientWithBaseURI(DefaultBaseURI)
}

// NewEndpointDataClientWithBaseURI creates an instance of the EndpointDataClient client.
func NewEndpointDataClientWithBaseURI(baseURI string) EndpointDataClient {
	return EndpointDataClient{NewWithBaseURI(baseURI)}
}

// Delete deletion will happen in the background and can take up to a day.
// Parameters:
// endpointID - the identifier of the endpoint.
func (client EndpointDataClient) Delete(ctx context.Context, endpointID uuid.UUID) (result ErrorContent, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EndpointDataClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, endpointID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "speechservices.EndpointDataClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "speechservices.EndpointDataClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "speechservices.EndpointDataClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client EndpointDataClient) DeletePreparer(ctx context.Context, endpointID uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointId": autorest.Encode("path", endpointID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/api/speechtotext/v2.0/endpoints/{endpointId}/data", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointDataClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client EndpointDataClient) DeleteResponder(resp *http.Response) (result ErrorContent, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusUnauthorized, http.StatusForbidden, http.StatusTooManyRequests),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
