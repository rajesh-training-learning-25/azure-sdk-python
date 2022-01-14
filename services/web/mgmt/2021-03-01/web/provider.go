package web

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
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

// ProviderClient is the webSite Management Client
type ProviderClient struct {
	BaseClient
}

// NewProviderClient creates an instance of the ProviderClient client.
func NewProviderClient(subscriptionID string) ProviderClient {
	return NewProviderClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProviderClientWithBaseURI creates an instance of the ProviderClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewProviderClientWithBaseURI(baseURI string, subscriptionID string) ProviderClient {
	return ProviderClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetAvailableStacks description for Get available application frameworks and their versions
func (client ProviderClient) GetAvailableStacks(ctx context.Context, osTypeSelected string) (result ApplicationStackCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderClient.GetAvailableStacks")
		defer func() {
			sc := -1
			if result.asc.Response.Response != nil {
				sc = result.asc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getAvailableStacksNextResults
	req, err := client.GetAvailableStacksPreparer(ctx, osTypeSelected)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetAvailableStacks", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAvailableStacksSender(req)
	if err != nil {
		result.asc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetAvailableStacks", resp, "Failure sending request")
		return
	}

	result.asc, err = client.GetAvailableStacksResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetAvailableStacks", resp, "Failure responding to request")
		return
	}
	if result.asc.hasNextLink() && result.asc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// GetAvailableStacksPreparer prepares the GetAvailableStacks request.
func (client ProviderClient) GetAvailableStacksPreparer(ctx context.Context, osTypeSelected string) (*http.Request, error) {
	const APIVersion = "2021-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(osTypeSelected)) > 0 {
		queryParameters["osTypeSelected"] = autorest.Encode("query", osTypeSelected)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Web/availableStacks"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAvailableStacksSender sends the GetAvailableStacks request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderClient) GetAvailableStacksSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetAvailableStacksResponder handles the response to the GetAvailableStacks request. The method always
// closes the http.Response Body.
func (client ProviderClient) GetAvailableStacksResponder(resp *http.Response) (result ApplicationStackCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getAvailableStacksNextResults retrieves the next set of results, if any.
func (client ProviderClient) getAvailableStacksNextResults(ctx context.Context, lastResults ApplicationStackCollection) (result ApplicationStackCollection, err error) {
	req, err := lastResults.applicationStackCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ProviderClient", "getAvailableStacksNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetAvailableStacksSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ProviderClient", "getAvailableStacksNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetAvailableStacksResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "getAvailableStacksNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetAvailableStacksComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProviderClient) GetAvailableStacksComplete(ctx context.Context, osTypeSelected string) (result ApplicationStackCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderClient.GetAvailableStacks")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetAvailableStacks(ctx, osTypeSelected)
	return
}

// GetAvailableStacksOnPrem description for Get available application frameworks and their versions
func (client ProviderClient) GetAvailableStacksOnPrem(ctx context.Context, osTypeSelected string) (result ApplicationStackCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderClient.GetAvailableStacksOnPrem")
		defer func() {
			sc := -1
			if result.asc.Response.Response != nil {
				sc = result.asc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getAvailableStacksOnPremNextResults
	req, err := client.GetAvailableStacksOnPremPreparer(ctx, osTypeSelected)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetAvailableStacksOnPrem", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAvailableStacksOnPremSender(req)
	if err != nil {
		result.asc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetAvailableStacksOnPrem", resp, "Failure sending request")
		return
	}

	result.asc, err = client.GetAvailableStacksOnPremResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetAvailableStacksOnPrem", resp, "Failure responding to request")
		return
	}
	if result.asc.hasNextLink() && result.asc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// GetAvailableStacksOnPremPreparer prepares the GetAvailableStacksOnPrem request.
func (client ProviderClient) GetAvailableStacksOnPremPreparer(ctx context.Context, osTypeSelected string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(osTypeSelected)) > 0 {
		queryParameters["osTypeSelected"] = autorest.Encode("query", osTypeSelected)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Web/availableStacks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAvailableStacksOnPremSender sends the GetAvailableStacksOnPrem request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderClient) GetAvailableStacksOnPremSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetAvailableStacksOnPremResponder handles the response to the GetAvailableStacksOnPrem request. The method always
// closes the http.Response Body.
func (client ProviderClient) GetAvailableStacksOnPremResponder(resp *http.Response) (result ApplicationStackCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getAvailableStacksOnPremNextResults retrieves the next set of results, if any.
func (client ProviderClient) getAvailableStacksOnPremNextResults(ctx context.Context, lastResults ApplicationStackCollection) (result ApplicationStackCollection, err error) {
	req, err := lastResults.applicationStackCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ProviderClient", "getAvailableStacksOnPremNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetAvailableStacksOnPremSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ProviderClient", "getAvailableStacksOnPremNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetAvailableStacksOnPremResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "getAvailableStacksOnPremNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetAvailableStacksOnPremComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProviderClient) GetAvailableStacksOnPremComplete(ctx context.Context, osTypeSelected string) (result ApplicationStackCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderClient.GetAvailableStacksOnPrem")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetAvailableStacksOnPrem(ctx, osTypeSelected)
	return
}

// GetFunctionAppStacks description for Get available Function app frameworks and their versions
// Parameters:
// stackOsType - stack OS Type
func (client ProviderClient) GetFunctionAppStacks(ctx context.Context, stackOsType string) (result FunctionAppStackCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderClient.GetFunctionAppStacks")
		defer func() {
			sc := -1
			if result.fasc.Response.Response != nil {
				sc = result.fasc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getFunctionAppStacksNextResults
	req, err := client.GetFunctionAppStacksPreparer(ctx, stackOsType)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetFunctionAppStacks", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetFunctionAppStacksSender(req)
	if err != nil {
		result.fasc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetFunctionAppStacks", resp, "Failure sending request")
		return
	}

	result.fasc, err = client.GetFunctionAppStacksResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetFunctionAppStacks", resp, "Failure responding to request")
		return
	}
	if result.fasc.hasNextLink() && result.fasc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// GetFunctionAppStacksPreparer prepares the GetFunctionAppStacks request.
func (client ProviderClient) GetFunctionAppStacksPreparer(ctx context.Context, stackOsType string) (*http.Request, error) {
	const APIVersion = "2021-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(stackOsType)) > 0 {
		queryParameters["stackOsType"] = autorest.Encode("query", stackOsType)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Web/functionAppStacks"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetFunctionAppStacksSender sends the GetFunctionAppStacks request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderClient) GetFunctionAppStacksSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetFunctionAppStacksResponder handles the response to the GetFunctionAppStacks request. The method always
// closes the http.Response Body.
func (client ProviderClient) GetFunctionAppStacksResponder(resp *http.Response) (result FunctionAppStackCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getFunctionAppStacksNextResults retrieves the next set of results, if any.
func (client ProviderClient) getFunctionAppStacksNextResults(ctx context.Context, lastResults FunctionAppStackCollection) (result FunctionAppStackCollection, err error) {
	req, err := lastResults.functionAppStackCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ProviderClient", "getFunctionAppStacksNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetFunctionAppStacksSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ProviderClient", "getFunctionAppStacksNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetFunctionAppStacksResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "getFunctionAppStacksNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetFunctionAppStacksComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProviderClient) GetFunctionAppStacksComplete(ctx context.Context, stackOsType string) (result FunctionAppStackCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderClient.GetFunctionAppStacks")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetFunctionAppStacks(ctx, stackOsType)
	return
}

// GetFunctionAppStacksForLocation description for Get available Function app frameworks and their versions for
// location
// Parameters:
// location - function App stack location.
// stackOsType - stack OS Type
func (client ProviderClient) GetFunctionAppStacksForLocation(ctx context.Context, location string, stackOsType string) (result FunctionAppStackCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderClient.GetFunctionAppStacksForLocation")
		defer func() {
			sc := -1
			if result.fasc.Response.Response != nil {
				sc = result.fasc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getFunctionAppStacksForLocationNextResults
	req, err := client.GetFunctionAppStacksForLocationPreparer(ctx, location, stackOsType)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetFunctionAppStacksForLocation", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetFunctionAppStacksForLocationSender(req)
	if err != nil {
		result.fasc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetFunctionAppStacksForLocation", resp, "Failure sending request")
		return
	}

	result.fasc, err = client.GetFunctionAppStacksForLocationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetFunctionAppStacksForLocation", resp, "Failure responding to request")
		return
	}
	if result.fasc.hasNextLink() && result.fasc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// GetFunctionAppStacksForLocationPreparer prepares the GetFunctionAppStacksForLocation request.
func (client ProviderClient) GetFunctionAppStacksForLocationPreparer(ctx context.Context, location string, stackOsType string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location": autorest.Encode("path", location),
	}

	const APIVersion = "2021-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(stackOsType)) > 0 {
		queryParameters["stackOsType"] = autorest.Encode("query", stackOsType)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Web/locations/{location}/functionAppStacks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetFunctionAppStacksForLocationSender sends the GetFunctionAppStacksForLocation request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderClient) GetFunctionAppStacksForLocationSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetFunctionAppStacksForLocationResponder handles the response to the GetFunctionAppStacksForLocation request. The method always
// closes the http.Response Body.
func (client ProviderClient) GetFunctionAppStacksForLocationResponder(resp *http.Response) (result FunctionAppStackCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getFunctionAppStacksForLocationNextResults retrieves the next set of results, if any.
func (client ProviderClient) getFunctionAppStacksForLocationNextResults(ctx context.Context, lastResults FunctionAppStackCollection) (result FunctionAppStackCollection, err error) {
	req, err := lastResults.functionAppStackCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ProviderClient", "getFunctionAppStacksForLocationNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetFunctionAppStacksForLocationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ProviderClient", "getFunctionAppStacksForLocationNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetFunctionAppStacksForLocationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "getFunctionAppStacksForLocationNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetFunctionAppStacksForLocationComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProviderClient) GetFunctionAppStacksForLocationComplete(ctx context.Context, location string, stackOsType string) (result FunctionAppStackCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderClient.GetFunctionAppStacksForLocation")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetFunctionAppStacksForLocation(ctx, location, stackOsType)
	return
}

// GetWebAppStacks description for Get available Web app frameworks and their versions
// Parameters:
// stackOsType - stack OS Type
func (client ProviderClient) GetWebAppStacks(ctx context.Context, stackOsType string) (result AppStackCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderClient.GetWebAppStacks")
		defer func() {
			sc := -1
			if result.asc.Response.Response != nil {
				sc = result.asc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getWebAppStacksNextResults
	req, err := client.GetWebAppStacksPreparer(ctx, stackOsType)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetWebAppStacks", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetWebAppStacksSender(req)
	if err != nil {
		result.asc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetWebAppStacks", resp, "Failure sending request")
		return
	}

	result.asc, err = client.GetWebAppStacksResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetWebAppStacks", resp, "Failure responding to request")
		return
	}
	if result.asc.hasNextLink() && result.asc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// GetWebAppStacksPreparer prepares the GetWebAppStacks request.
func (client ProviderClient) GetWebAppStacksPreparer(ctx context.Context, stackOsType string) (*http.Request, error) {
	const APIVersion = "2021-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(stackOsType)) > 0 {
		queryParameters["stackOsType"] = autorest.Encode("query", stackOsType)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Web/webAppStacks"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetWebAppStacksSender sends the GetWebAppStacks request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderClient) GetWebAppStacksSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetWebAppStacksResponder handles the response to the GetWebAppStacks request. The method always
// closes the http.Response Body.
func (client ProviderClient) GetWebAppStacksResponder(resp *http.Response) (result AppStackCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getWebAppStacksNextResults retrieves the next set of results, if any.
func (client ProviderClient) getWebAppStacksNextResults(ctx context.Context, lastResults AppStackCollection) (result AppStackCollection, err error) {
	req, err := lastResults.appStackCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ProviderClient", "getWebAppStacksNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetWebAppStacksSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ProviderClient", "getWebAppStacksNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetWebAppStacksResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "getWebAppStacksNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetWebAppStacksComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProviderClient) GetWebAppStacksComplete(ctx context.Context, stackOsType string) (result AppStackCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderClient.GetWebAppStacks")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetWebAppStacks(ctx, stackOsType)
	return
}

// GetWebAppStacksForLocation description for Get available Web app frameworks and their versions for location
// Parameters:
// location - web App stack location.
// stackOsType - stack OS Type
func (client ProviderClient) GetWebAppStacksForLocation(ctx context.Context, location string, stackOsType string) (result AppStackCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderClient.GetWebAppStacksForLocation")
		defer func() {
			sc := -1
			if result.asc.Response.Response != nil {
				sc = result.asc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getWebAppStacksForLocationNextResults
	req, err := client.GetWebAppStacksForLocationPreparer(ctx, location, stackOsType)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetWebAppStacksForLocation", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetWebAppStacksForLocationSender(req)
	if err != nil {
		result.asc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetWebAppStacksForLocation", resp, "Failure sending request")
		return
	}

	result.asc, err = client.GetWebAppStacksForLocationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetWebAppStacksForLocation", resp, "Failure responding to request")
		return
	}
	if result.asc.hasNextLink() && result.asc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// GetWebAppStacksForLocationPreparer prepares the GetWebAppStacksForLocation request.
func (client ProviderClient) GetWebAppStacksForLocationPreparer(ctx context.Context, location string, stackOsType string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location": autorest.Encode("path", location),
	}

	const APIVersion = "2021-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(stackOsType)) > 0 {
		queryParameters["stackOsType"] = autorest.Encode("query", stackOsType)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Web/locations/{location}/webAppStacks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetWebAppStacksForLocationSender sends the GetWebAppStacksForLocation request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderClient) GetWebAppStacksForLocationSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetWebAppStacksForLocationResponder handles the response to the GetWebAppStacksForLocation request. The method always
// closes the http.Response Body.
func (client ProviderClient) GetWebAppStacksForLocationResponder(resp *http.Response) (result AppStackCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getWebAppStacksForLocationNextResults retrieves the next set of results, if any.
func (client ProviderClient) getWebAppStacksForLocationNextResults(ctx context.Context, lastResults AppStackCollection) (result AppStackCollection, err error) {
	req, err := lastResults.appStackCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ProviderClient", "getWebAppStacksForLocationNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetWebAppStacksForLocationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ProviderClient", "getWebAppStacksForLocationNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetWebAppStacksForLocationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "getWebAppStacksForLocationNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetWebAppStacksForLocationComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProviderClient) GetWebAppStacksForLocationComplete(ctx context.Context, location string, stackOsType string) (result AppStackCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderClient.GetWebAppStacksForLocation")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetWebAppStacksForLocation(ctx, location, stackOsType)
	return
}

// ListOperations description for Gets all available operations for the Microsoft.Web resource provider. Also exposes
// resource metric definitions
func (client ProviderClient) ListOperations(ctx context.Context) (result CsmOperationCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderClient.ListOperations")
		defer func() {
			sc := -1
			if result.coc.Response.Response != nil {
				sc = result.coc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listOperationsNextResults
	req, err := client.ListOperationsPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "ListOperations", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListOperationsSender(req)
	if err != nil {
		result.coc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "ListOperations", resp, "Failure sending request")
		return
	}

	result.coc, err = client.ListOperationsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "ListOperations", resp, "Failure responding to request")
		return
	}
	if result.coc.hasNextLink() && result.coc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListOperationsPreparer prepares the ListOperations request.
func (client ProviderClient) ListOperationsPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2021-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Web/operations"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListOperationsSender sends the ListOperations request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderClient) ListOperationsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListOperationsResponder handles the response to the ListOperations request. The method always
// closes the http.Response Body.
func (client ProviderClient) ListOperationsResponder(resp *http.Response) (result CsmOperationCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listOperationsNextResults retrieves the next set of results, if any.
func (client ProviderClient) listOperationsNextResults(ctx context.Context, lastResults CsmOperationCollection) (result CsmOperationCollection, err error) {
	req, err := lastResults.csmOperationCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ProviderClient", "listOperationsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListOperationsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ProviderClient", "listOperationsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListOperationsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "listOperationsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListOperationsComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProviderClient) ListOperationsComplete(ctx context.Context) (result CsmOperationCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderClient.ListOperations")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListOperations(ctx)
	return
}
