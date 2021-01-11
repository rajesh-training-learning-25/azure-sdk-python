package apimanagement

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

// TenantConfigurationClient is the client for the TenantConfiguration methods of the Apimanagement service.
type TenantConfigurationClient struct {
	BaseClient
}

// NewTenantConfigurationClient creates an instance of the TenantConfigurationClient client.
func NewTenantConfigurationClient() TenantConfigurationClient {
	return TenantConfigurationClient{New()}
}

// Deploy this operation applies changes from the specified Git branch to the configuration database. This is a long
// running operation and could take several minutes to complete.
// Parameters:
// apimBaseURL - the management endpoint of the API Management service, for example
// https://myapimservice.management.azure-api.net.
// parameters - deploy Configuration parameters.
func (client TenantConfigurationClient) Deploy(ctx context.Context, apimBaseURL string, parameters DeployConfigurationParameters) (result TenantConfigurationDeployFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TenantConfigurationClient.Deploy")
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
			Constraints: []validation.Constraint{{Target: "parameters.Branch", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.TenantConfigurationClient", "Deploy", err.Error())
	}

	req, err := client.DeployPreparer(ctx, apimBaseURL, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationClient", "Deploy", nil, "Failure preparing request")
		return
	}

	result, err = client.DeploySender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationClient", "Deploy", nil, "Failure sending request")
		return
	}

	return
}

// DeployPreparer prepares the Deploy request.
func (client TenantConfigurationClient) DeployPreparer(ctx context.Context, apimBaseURL string, parameters DeployConfigurationParameters) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"apimBaseUrl": apimBaseURL,
	}

	pathParameters := map[string]interface{}{
		"configurationName": autorest.Encode("path", "configuration"),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{apimBaseUrl}", urlParameters),
		autorest.WithPathParameters("/tenant/{configurationName}/deploy", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeploySender sends the Deploy request. The method will close the
// http.Response Body if it receives an error.
func (client TenantConfigurationClient) DeploySender(req *http.Request) (future TenantConfigurationDeployFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client TenantConfigurationClient) (orc OperationResultContract, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationDeployFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("apimanagement.TenantConfigurationDeployFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		if orc.Response.Response, err = future.GetResult(sender); err == nil && orc.Response.Response.StatusCode != http.StatusNoContent {
			orc, err = client.DeployResponder(orc.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationDeployFuture", "Result", orc.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// DeployResponder handles the response to the Deploy request. The method always
// closes the http.Response Body.
func (client TenantConfigurationClient) DeployResponder(resp *http.Response) (result OperationResultContract, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetSyncState gets the status of the most recent synchronization between the configuration database and the Git
// repository.
// Parameters:
// apimBaseURL - the management endpoint of the API Management service, for example
// https://myapimservice.management.azure-api.net.
func (client TenantConfigurationClient) GetSyncState(ctx context.Context, apimBaseURL string) (result TenantConfigurationSyncStateContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TenantConfigurationClient.GetSyncState")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetSyncStatePreparer(ctx, apimBaseURL)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationClient", "GetSyncState", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSyncStateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationClient", "GetSyncState", resp, "Failure sending request")
		return
	}

	result, err = client.GetSyncStateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationClient", "GetSyncState", resp, "Failure responding to request")
		return
	}

	return
}

// GetSyncStatePreparer prepares the GetSyncState request.
func (client TenantConfigurationClient) GetSyncStatePreparer(ctx context.Context, apimBaseURL string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"apimBaseUrl": apimBaseURL,
	}

	pathParameters := map[string]interface{}{
		"configurationName": autorest.Encode("path", "configuration"),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{apimBaseUrl}", urlParameters),
		autorest.WithPathParameters("/tenant/{configurationName}/syncState", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSyncStateSender sends the GetSyncState request. The method will close the
// http.Response Body if it receives an error.
func (client TenantConfigurationClient) GetSyncStateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetSyncStateResponder handles the response to the GetSyncState request. The method always
// closes the http.Response Body.
func (client TenantConfigurationClient) GetSyncStateResponder(resp *http.Response) (result TenantConfigurationSyncStateContract, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Save this operation creates a commit with the current configuration snapshot to the specified branch in the
// repository. This is a long running operation and could take several minutes to complete.
// Parameters:
// apimBaseURL - the management endpoint of the API Management service, for example
// https://myapimservice.management.azure-api.net.
// parameters - save Configuration parameters.
func (client TenantConfigurationClient) Save(ctx context.Context, apimBaseURL string, parameters SaveConfigurationParameter) (result TenantConfigurationSaveFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TenantConfigurationClient.Save")
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
			Constraints: []validation.Constraint{{Target: "parameters.Branch", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.TenantConfigurationClient", "Save", err.Error())
	}

	req, err := client.SavePreparer(ctx, apimBaseURL, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationClient", "Save", nil, "Failure preparing request")
		return
	}

	result, err = client.SaveSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationClient", "Save", nil, "Failure sending request")
		return
	}

	return
}

// SavePreparer prepares the Save request.
func (client TenantConfigurationClient) SavePreparer(ctx context.Context, apimBaseURL string, parameters SaveConfigurationParameter) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"apimBaseUrl": apimBaseURL,
	}

	pathParameters := map[string]interface{}{
		"configurationName": autorest.Encode("path", "configuration"),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{apimBaseUrl}", urlParameters),
		autorest.WithPathParameters("/tenant/{configurationName}/save", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SaveSender sends the Save request. The method will close the
// http.Response Body if it receives an error.
func (client TenantConfigurationClient) SaveSender(req *http.Request) (future TenantConfigurationSaveFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client TenantConfigurationClient) (orc OperationResultContract, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationSaveFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("apimanagement.TenantConfigurationSaveFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		if orc.Response.Response, err = future.GetResult(sender); err == nil && orc.Response.Response.StatusCode != http.StatusNoContent {
			orc, err = client.SaveResponder(orc.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationSaveFuture", "Result", orc.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// SaveResponder handles the response to the Save request. The method always
// closes the http.Response Body.
func (client TenantConfigurationClient) SaveResponder(resp *http.Response) (result OperationResultContract, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Validate this operation validates the changes in the specified Git branch. This is a long running operation and
// could take several minutes to complete.
// Parameters:
// apimBaseURL - the management endpoint of the API Management service, for example
// https://myapimservice.management.azure-api.net.
// parameters - validate Configuration parameters.
func (client TenantConfigurationClient) Validate(ctx context.Context, apimBaseURL string, parameters DeployConfigurationParameters) (result TenantConfigurationValidateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TenantConfigurationClient.Validate")
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
			Constraints: []validation.Constraint{{Target: "parameters.Branch", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.TenantConfigurationClient", "Validate", err.Error())
	}

	req, err := client.ValidatePreparer(ctx, apimBaseURL, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationClient", "Validate", nil, "Failure preparing request")
		return
	}

	result, err = client.ValidateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationClient", "Validate", nil, "Failure sending request")
		return
	}

	return
}

// ValidatePreparer prepares the Validate request.
func (client TenantConfigurationClient) ValidatePreparer(ctx context.Context, apimBaseURL string, parameters DeployConfigurationParameters) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"apimBaseUrl": apimBaseURL,
	}

	pathParameters := map[string]interface{}{
		"configurationName": autorest.Encode("path", "configuration"),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{apimBaseUrl}", urlParameters),
		autorest.WithPathParameters("/tenant/{configurationName}/validate", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ValidateSender sends the Validate request. The method will close the
// http.Response Body if it receives an error.
func (client TenantConfigurationClient) ValidateSender(req *http.Request) (future TenantConfigurationValidateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client TenantConfigurationClient) (orc OperationResultContract, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationValidateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("apimanagement.TenantConfigurationValidateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		if orc.Response.Response, err = future.GetResult(sender); err == nil && orc.Response.Response.StatusCode != http.StatusNoContent {
			orc, err = client.ValidateResponder(orc.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationValidateFuture", "Result", orc.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// ValidateResponder handles the response to the Validate request. The method always
// closes the http.Response Body.
func (client TenantConfigurationClient) ValidateResponder(resp *http.Response) (result OperationResultContract, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
