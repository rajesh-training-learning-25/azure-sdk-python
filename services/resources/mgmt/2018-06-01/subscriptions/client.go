// Deprecated: Please note, this package has been deprecated. A replacement package is available github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions. We strongly encourage you to upgrade to continue receiving updates. See Migration Guide for guidance on upgrading. Refer to our deprecation policy for more details.
//
// Package subscriptions implements the Azure ARM Subscriptions service API version 2018-06-01.
//
// All resource groups and resources exist within subscriptions. These operation enable you get information about your
// subscriptions and tenants. A tenant is a dedicated instance of Azure Active Directory (Azure AD) for your
// organization.
package subscriptions

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
	// DefaultBaseURI is the default URI used for the service Subscriptions
	DefaultBaseURI = "https://management.azure.com"
)

// BaseClient is the base client for Subscriptions.
type BaseClient struct {
	autorest.Client
	BaseURI string
}

// New creates an instance of the BaseClient client.
func New() BaseClient {
	return NewWithBaseURI(DefaultBaseURI)
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string) BaseClient {
	return BaseClient{
		Client:  autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI: baseURI,
	}
}

// CheckResourceName a resource name is valid if it is not a reserved word, does not contains a reserved word and does
// not start with a reserved word
// Parameters:
// resourceNameDefinition - resource object with values for resource name and resource type
func (client BaseClient) CheckResourceName(ctx context.Context, resourceNameDefinition *ResourceName) (result CheckResourceNameResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.CheckResourceName")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceNameDefinition,
			Constraints: []validation.Constraint{{Target: "resourceNameDefinition", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "resourceNameDefinition.Name", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "resourceNameDefinition.Type", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("subscriptions.BaseClient", "CheckResourceName", err.Error())
	}

	req, err := client.CheckResourceNamePreparer(ctx, resourceNameDefinition)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscriptions.BaseClient", "CheckResourceName", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckResourceNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "subscriptions.BaseClient", "CheckResourceName", resp, "Failure sending request")
		return
	}

	result, err = client.CheckResourceNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscriptions.BaseClient", "CheckResourceName", resp, "Failure responding to request")
		return
	}

	return
}

// CheckResourceNamePreparer prepares the CheckResourceName request.
func (client BaseClient) CheckResourceNamePreparer(ctx context.Context, resourceNameDefinition *ResourceName) (*http.Request, error) {
	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Resources/checkResourceName"),
		autorest.WithQueryParameters(queryParameters))
	if resourceNameDefinition != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(resourceNameDefinition))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckResourceNameSender sends the CheckResourceName request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) CheckResourceNameSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CheckResourceNameResponder handles the response to the CheckResourceName request. The method always
// closes the http.Response Body.
func (client BaseClient) CheckResourceNameResponder(resp *http.Response) (result CheckResourceNameResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
