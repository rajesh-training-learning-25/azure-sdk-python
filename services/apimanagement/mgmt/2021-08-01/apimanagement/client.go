// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
//
// Package apimanagement implements the Azure ARM Apimanagement service API version 2021-08-01.
//
// ApiManagement Client
package apimanagement

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
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

const (
	// DefaultBaseURI is the default URI used for the service Apimanagement
	DefaultBaseURI = "https://management.azure.com"
)

// BaseClient is the base client for Apimanagement.
type BaseClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the BaseClient client.
func New(subscriptionID string) BaseClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return BaseClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}

// PerformConnectivityCheckAsync performs a connectivity check between the API Management service and a given
// destination, and returns metrics for the connection, as well as errors encountered while trying to establish it.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// connectivityCheckRequestParams - connectivity Check request parameters.
func (client BaseClient) PerformConnectivityCheckAsync(ctx context.Context, resourceGroupName string, serviceName string, connectivityCheckRequestParams ConnectivityCheckRequest) (result PerformConnectivityCheckAsyncFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.PerformConnectivityCheckAsync")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: connectivityCheckRequestParams,
			Constraints: []validation.Constraint{{Target: "connectivityCheckRequestParams.Source", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "connectivityCheckRequestParams.Source.Region", Name: validation.Null, Rule: true, Chain: nil}}},
				{Target: "connectivityCheckRequestParams.Destination", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "connectivityCheckRequestParams.Destination.Address", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "connectivityCheckRequestParams.Destination.Port", Name: validation.Null, Rule: true, Chain: nil},
					}}}}}); err != nil {
		return result, validation.NewError("apimanagement.BaseClient", "PerformConnectivityCheckAsync", err.Error())
	}

	req, err := client.PerformConnectivityCheckAsyncPreparer(ctx, resourceGroupName, serviceName, connectivityCheckRequestParams)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.BaseClient", "PerformConnectivityCheckAsync", nil, "Failure preparing request")
		return
	}

	result, err = client.PerformConnectivityCheckAsyncSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.BaseClient", "PerformConnectivityCheckAsync", result.Response(), "Failure sending request")
		return
	}

	return
}

// PerformConnectivityCheckAsyncPreparer prepares the PerformConnectivityCheckAsync request.
func (client BaseClient) PerformConnectivityCheckAsyncPreparer(ctx context.Context, resourceGroupName string, serviceName string, connectivityCheckRequestParams ConnectivityCheckRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/connectivityCheck", pathParameters),
		autorest.WithJSON(connectivityCheckRequestParams),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PerformConnectivityCheckAsyncSender sends the PerformConnectivityCheckAsync request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) PerformConnectivityCheckAsyncSender(req *http.Request) (future PerformConnectivityCheckAsyncFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// PerformConnectivityCheckAsyncResponder handles the response to the PerformConnectivityCheckAsync request. The method always
// closes the http.Response Body.
func (client BaseClient) PerformConnectivityCheckAsyncResponder(resp *http.Response) (result ConnectivityCheckResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
