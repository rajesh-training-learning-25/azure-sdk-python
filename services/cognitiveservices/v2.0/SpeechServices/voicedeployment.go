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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"github.com/satori/go.uuid"
	"net/http"
)

// VoiceDeploymentClient is the speech Services API v2.0.
type VoiceDeploymentClient struct {
	BaseClient
}

// NewVoiceDeploymentClient creates an instance of the VoiceDeploymentClient client.
func NewVoiceDeploymentClient() VoiceDeploymentClient {
	return NewVoiceDeploymentClientWithBaseURI(DefaultBaseURI)
}

// NewVoiceDeploymentClientWithBaseURI creates an instance of the VoiceDeploymentClient client.
func NewVoiceDeploymentClientWithBaseURI(baseURI string) VoiceDeploymentClient {
	return VoiceDeploymentClient{NewWithBaseURI(baseURI)}
}

// Create sends the create request.
func (client VoiceDeploymentClient) Create(ctx context.Context, endpoint EndpointDefinition) (result VoiceDeploymentCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VoiceDeploymentClient.Create")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: endpoint,
			Constraints: []validation.Constraint{{Target: "endpoint.ConcurrentRecognitions", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "endpoint.ConcurrentRecognitions", Name: validation.InclusiveMaximum, Rule: int64(20), Chain: nil},
					{Target: "endpoint.ConcurrentRecognitions", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
				}},
				{Target: "endpoint.ModelsProperty", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "endpoint.Name", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("speechservices.VoiceDeploymentClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, endpoint)
	if err != nil {
		err = autorest.NewErrorWithError(err, "speechservices.VoiceDeploymentClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "speechservices.VoiceDeploymentClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client VoiceDeploymentClient) CreatePreparer(ctx context.Context, endpoint EndpointDefinition) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/api/texttospeech/v2.0/endpoints"),
		autorest.WithJSON(endpoint))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client VoiceDeploymentClient) CreateSender(req *http.Request) (future VoiceDeploymentCreateFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client VoiceDeploymentClient) CreateResponder(resp *http.Response) (result ErrorContent, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusTooManyRequests),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get sends the get request.
func (client VoiceDeploymentClient) Get(ctx context.Context, ID uuid.UUID) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VoiceDeploymentClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, ID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "speechservices.VoiceDeploymentClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "speechservices.VoiceDeploymentClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "speechservices.VoiceDeploymentClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client VoiceDeploymentClient) GetPreparer(ctx context.Context, ID uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"id": autorest.Encode("path", ID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/api/texttospeech/v2.0/endpoints/{id}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VoiceDeploymentClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VoiceDeploymentClient) GetResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusUnauthorized, http.StatusForbidden, http.StatusTooManyRequests),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
