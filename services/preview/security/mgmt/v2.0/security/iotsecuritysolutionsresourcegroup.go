package security

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

// IoTSecuritySolutionsResourceGroupClient is the API spec for Microsoft.Security (Azure Security Center) resource
// provider
type IoTSecuritySolutionsResourceGroupClient struct {
	BaseClient
}

// NewIoTSecuritySolutionsResourceGroupClient creates an instance of the IoTSecuritySolutionsResourceGroupClient
// client.
func NewIoTSecuritySolutionsResourceGroupClient(subscriptionID string) IoTSecuritySolutionsResourceGroupClient {
	return NewIoTSecuritySolutionsResourceGroupClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewIoTSecuritySolutionsResourceGroupClientWithBaseURI creates an instance of the
// IoTSecuritySolutionsResourceGroupClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewIoTSecuritySolutionsResourceGroupClientWithBaseURI(baseURI string, subscriptionID string) IoTSecuritySolutionsResourceGroupClient {
	return IoTSecuritySolutionsResourceGroupClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List list of security solutions
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// filter - filter the Security Solution with OData syntax. supporting filter by iotHubs
func (client IoTSecuritySolutionsResourceGroupClient) List(ctx context.Context, resourceGroupName string, filter string) (result IoTSecuritySolutionsListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IoTSecuritySolutionsResourceGroupClient.List")
		defer func() {
			sc := -1
			if result.itssl.Response.Response != nil {
				sc = result.itssl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.IoTSecuritySolutionsResourceGroupClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.IoTSecuritySolutionsResourceGroupClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.itssl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.IoTSecuritySolutionsResourceGroupClient", "List", resp, "Failure sending request")
		return
	}

	result.itssl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.IoTSecuritySolutionsResourceGroupClient", "List", resp, "Failure responding to request")
		return
	}
	if result.itssl.hasNextLink() && result.itssl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client IoTSecuritySolutionsResourceGroupClient) ListPreparer(ctx context.Context, resourceGroupName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client IoTSecuritySolutionsResourceGroupClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client IoTSecuritySolutionsResourceGroupClient) ListResponder(resp *http.Response) (result IoTSecuritySolutionsList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client IoTSecuritySolutionsResourceGroupClient) listNextResults(ctx context.Context, lastResults IoTSecuritySolutionsList) (result IoTSecuritySolutionsList, err error) {
	req, err := lastResults.ioTSecuritySolutionsListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "security.IoTSecuritySolutionsResourceGroupClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "security.IoTSecuritySolutionsResourceGroupClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.IoTSecuritySolutionsResourceGroupClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client IoTSecuritySolutionsResourceGroupClient) ListComplete(ctx context.Context, resourceGroupName string, filter string) (result IoTSecuritySolutionsListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IoTSecuritySolutionsResourceGroupClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, filter)
	return
}
