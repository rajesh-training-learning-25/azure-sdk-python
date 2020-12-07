package resourcemover

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

// OperationsDiscoveryClient is the a first party Azure service orchestrating the move of Azure resources from one
// Azure region to another or between zones within a region.
type OperationsDiscoveryClient struct {
	BaseClient
}

// NewOperationsDiscoveryClient creates an instance of the OperationsDiscoveryClient client.
func NewOperationsDiscoveryClient(subscriptionID string) OperationsDiscoveryClient {
	return NewOperationsDiscoveryClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewOperationsDiscoveryClientWithBaseURI creates an instance of the OperationsDiscoveryClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewOperationsDiscoveryClientWithBaseURI(baseURI string, subscriptionID string) OperationsDiscoveryClient {
	return OperationsDiscoveryClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get sends the get request.
func (client OperationsDiscoveryClient) Get(ctx context.Context) (result OperationsDiscoveryCollection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationsDiscoveryClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcemover.OperationsDiscoveryClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resourcemover.OperationsDiscoveryClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcemover.OperationsDiscoveryClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client OperationsDiscoveryClient) GetPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2019-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Migrate/operations"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client OperationsDiscoveryClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client OperationsDiscoveryClient) GetResponder(resp *http.Response) (result OperationsDiscoveryCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
