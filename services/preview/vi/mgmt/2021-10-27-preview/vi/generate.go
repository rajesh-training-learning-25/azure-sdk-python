package vi

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

// GenerateClient is the microsoft Azure Video Analyzer for Media
type GenerateClient struct {
	BaseClient
}

// NewGenerateClient creates an instance of the GenerateClient client.
func NewGenerateClient(subscriptionID string) GenerateClient {
	return NewGenerateClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewGenerateClientWithBaseURI creates an instance of the GenerateClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewGenerateClientWithBaseURI(baseURI string, subscriptionID string) GenerateClient {
	return GenerateClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// AccessTokenMethod generate an Azure Video Analyzer for Media access token.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// accountName - the name of the Azure Video Analyzer for Media account.
// parameters - the parameters for generating access token
func (client GenerateClient) AccessTokenMethod(ctx context.Context, resourceGroupName string, accountName string, parameters *GenerateAccessTokenParameters) (result AccessToken, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GenerateClient.AccessTokenMethod")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-z0-9]+(-[a-z0-9]+)*`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("vi.GenerateClient", "AccessTokenMethod", err.Error())
	}

	req, err := client.AccessTokenMethodPreparer(ctx, resourceGroupName, accountName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vi.GenerateClient", "AccessTokenMethod", nil, "Failure preparing request")
		return
	}

	resp, err := client.AccessTokenMethodSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vi.GenerateClient", "AccessTokenMethod", resp, "Failure sending request")
		return
	}

	result, err = client.AccessTokenMethodResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vi.GenerateClient", "AccessTokenMethod", resp, "Failure responding to request")
		return
	}

	return
}

// AccessTokenMethodPreparer prepares the AccessTokenMethod request.
func (client GenerateClient) AccessTokenMethodPreparer(ctx context.Context, resourceGroupName string, accountName string, parameters *GenerateAccessTokenParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-10-27-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VideoIndexer/accounts/{accountName}/generateAccessToken", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if parameters != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(parameters))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AccessTokenMethodSender sends the AccessTokenMethod request. The method will close the
// http.Response Body if it receives an error.
func (client GenerateClient) AccessTokenMethodSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// AccessTokenMethodResponder handles the response to the AccessTokenMethod request. The method always
// closes the http.Response Body.
func (client GenerateClient) AccessTokenMethodResponder(resp *http.Response) (result AccessToken, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
