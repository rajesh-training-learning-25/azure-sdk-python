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
	"net/http"
)

// SupportedLocalesForLanguageGenerationEndpointsClient is the speech Services API v2.0.
type SupportedLocalesForLanguageGenerationEndpointsClient struct {
	BaseClient
}

// NewSupportedLocalesForLanguageGenerationEndpointsClient creates an instance of the
// SupportedLocalesForLanguageGenerationEndpointsClient client.
func NewSupportedLocalesForLanguageGenerationEndpointsClient() SupportedLocalesForLanguageGenerationEndpointsClient {
	return NewSupportedLocalesForLanguageGenerationEndpointsClientWithBaseURI(DefaultBaseURI)
}

// NewSupportedLocalesForLanguageGenerationEndpointsClientWithBaseURI creates an instance of the
// SupportedLocalesForLanguageGenerationEndpointsClient client.
func NewSupportedLocalesForLanguageGenerationEndpointsClientWithBaseURI(baseURI string) SupportedLocalesForLanguageGenerationEndpointsClient {
	return SupportedLocalesForLanguageGenerationEndpointsClient{NewWithBaseURI(baseURI)}
}

// Get sends the get request.
func (client SupportedLocalesForLanguageGenerationEndpointsClient) Get(ctx context.Context) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SupportedLocalesForLanguageGenerationEndpointsClient.Get")
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
		err = autorest.NewErrorWithError(err, "speechservices.SupportedLocalesForLanguageGenerationEndpointsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "speechservices.SupportedLocalesForLanguageGenerationEndpointsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "speechservices.SupportedLocalesForLanguageGenerationEndpointsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SupportedLocalesForLanguageGenerationEndpointsClient) GetPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/api/languagegeneration/v2.0/Endpoints/locales"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SupportedLocalesForLanguageGenerationEndpointsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SupportedLocalesForLanguageGenerationEndpointsClient) GetResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusUnauthorized, http.StatusForbidden, http.StatusTooManyRequests),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
