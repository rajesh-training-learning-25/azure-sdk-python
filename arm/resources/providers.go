package resources

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
// Code generated by Microsoft (R) AutoRest Code Generator 0.11.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/azure/go-autorest/autorest"
	"net/http"
	"net/url"
)

// Providers Client
type ProvidersClient struct {
	ResourceManagementClient
}

func NewProvidersClient(subscriptionId string) ProvidersClient {
	return NewProvidersClientWithBaseUri(DefaultBaseUri, subscriptionId)
}

func NewProvidersClientWithBaseUri(baseUri string, subscriptionId string) ProvidersClient {
	return ProvidersClient{NewWithBaseUri(baseUri, subscriptionId)}
}

// Unregister unregisters provider from a subscription.
//
// resourceProviderNamespace is namespace of the resource provider.
func (client ProvidersClient) Unregister(resourceProviderNamespace string) (result Provider, ae autorest.Error) {
	req, err := client.NewUnregisterRequest(resourceProviderNamespace)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.ProvidersClient", "Unregister", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.ProvidersClient", "Unregister", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(
		client,
		req,
		autorest.DoErrorUnlessStatusCode(http.StatusOK))

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "resources.ProvidersClient", "Unregister", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "resources.ProvidersClient", "Unregister", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the Unregister request.
func (client ProvidersClient) NewUnregisterRequest(resourceProviderNamespace string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceProviderNamespace": url.QueryEscape(resourceProviderNamespace),
		"subscriptionId":            url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.UnregisterRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the Unregister request.
func (client ProvidersClient) UnregisterRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/{resourceProviderNamespace}/unregister"))
}

// Register registers provider to be used with a subscription.
//
// resourceProviderNamespace is namespace of the resource provider.
func (client ProvidersClient) Register(resourceProviderNamespace string) (result Provider, ae autorest.Error) {
	req, err := client.NewRegisterRequest(resourceProviderNamespace)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.ProvidersClient", "Register", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.ProvidersClient", "Register", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(
		client,
		req,
		autorest.DoErrorUnlessStatusCode(http.StatusOK))

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "resources.ProvidersClient", "Register", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "resources.ProvidersClient", "Register", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the Register request.
func (client ProvidersClient) NewRegisterRequest(resourceProviderNamespace string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceProviderNamespace": url.QueryEscape(resourceProviderNamespace),
		"subscriptionId":            url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.RegisterRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the Register request.
func (client ProvidersClient) RegisterRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/{resourceProviderNamespace}/register"))
}

// List gets a list of resource providers.
//
// top is query parameters. If null is passed returns all deployments.
func (client ProvidersClient) List(top int) (result ProviderListResult, ae autorest.Error) {
	req, err := client.NewListRequest(top)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.ProvidersClient", "List", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.ProvidersClient", "List", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(
		client,
		req,
		autorest.DoErrorUnlessStatusCode(http.StatusOK))

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "resources.ProvidersClient", "List", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "resources.ProvidersClient", "List", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the List request.
func (client ProvidersClient) NewListRequest(top int) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"$top":        top,
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.ListRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the List request.
func (client ProvidersClient) ListRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers"))
}

// Get gets a resource provider.
//
// resourceProviderNamespace is namespace of the resource provider.
func (client ProvidersClient) Get(resourceProviderNamespace string) (result Provider, ae autorest.Error) {
	req, err := client.NewGetRequest(resourceProviderNamespace)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.ProvidersClient", "Get", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.ProvidersClient", "Get", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(
		client,
		req,
		autorest.DoErrorUnlessStatusCode(http.StatusOK))

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "resources.ProvidersClient", "Get", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "resources.ProvidersClient", "Get", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the Get request.
func (client ProvidersClient) NewGetRequest(resourceProviderNamespace string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceProviderNamespace": url.QueryEscape(resourceProviderNamespace),
		"subscriptionId":            url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.GetRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the Get request.
func (client ProvidersClient) GetRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/{resourceProviderNamespace}"))
}
