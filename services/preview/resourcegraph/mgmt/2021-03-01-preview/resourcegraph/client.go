// Deprecated: Please note, this package has been deprecated. A replacement package is available github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph. We strongly encourage you to upgrade to continue receiving updates. See Migration Guide for guidance on upgrading. Refer to our deprecation policy for more details.
//
// Package resourcegraph implements the Azure ARM Resourcegraph service API version .
//
// Azure Resource Graph API Reference
package resourcegraph

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
	// DefaultBaseURI is the default URI used for the service Resourcegraph
	DefaultBaseURI = "https://management.azure.com"
)

// BaseClient is the base client for Resourcegraph.
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

// ResourceChangeDetails get resource change details.
// Parameters:
// parameters - the parameters for this request for resource change details.
func (client BaseClient) ResourceChangeDetails(ctx context.Context, parameters ResourceChangeDetailsRequestParameters) (result ListResourceChangeData, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.ResourceChangeDetails")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ResourceIds", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.ChangeIds", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("resourcegraph.BaseClient", "ResourceChangeDetails", err.Error())
	}

	req, err := client.ResourceChangeDetailsPreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcegraph.BaseClient", "ResourceChangeDetails", nil, "Failure preparing request")
		return
	}

	resp, err := client.ResourceChangeDetailsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resourcegraph.BaseClient", "ResourceChangeDetails", resp, "Failure sending request")
		return
	}

	result, err = client.ResourceChangeDetailsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcegraph.BaseClient", "ResourceChangeDetails", resp, "Failure responding to request")
		return
	}

	return
}

// ResourceChangeDetailsPreparer prepares the ResourceChangeDetails request.
func (client BaseClient) ResourceChangeDetailsPreparer(ctx context.Context, parameters ResourceChangeDetailsRequestParameters) (*http.Request, error) {
	const APIVersion = "2020-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.ResourceGraph/resourceChangeDetails"),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ResourceChangeDetailsSender sends the ResourceChangeDetails request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) ResourceChangeDetailsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ResourceChangeDetailsResponder handles the response to the ResourceChangeDetails request. The method always
// closes the http.Response Body.
func (client BaseClient) ResourceChangeDetailsResponder(resp *http.Response) (result ListResourceChangeData, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ResourceChanges list changes to a resource for a given time interval.
// Parameters:
// parameters - the parameters for this request for changes.
func (client BaseClient) ResourceChanges(ctx context.Context, parameters ResourceChangesRequestParameters) (result ResourceChangeList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.ResourceChanges")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Interval", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.Top", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.Top", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
						{Target: "parameters.Top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
					}}}}}); err != nil {
		return result, validation.NewError("resourcegraph.BaseClient", "ResourceChanges", err.Error())
	}

	req, err := client.ResourceChangesPreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcegraph.BaseClient", "ResourceChanges", nil, "Failure preparing request")
		return
	}

	resp, err := client.ResourceChangesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resourcegraph.BaseClient", "ResourceChanges", resp, "Failure sending request")
		return
	}

	result, err = client.ResourceChangesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcegraph.BaseClient", "ResourceChanges", resp, "Failure responding to request")
		return
	}

	return
}

// ResourceChangesPreparer prepares the ResourceChanges request.
func (client BaseClient) ResourceChangesPreparer(ctx context.Context, parameters ResourceChangesRequestParameters) (*http.Request, error) {
	const APIVersion = "2020-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.ResourceGraph/resourceChanges"),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ResourceChangesSender sends the ResourceChanges request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) ResourceChangesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ResourceChangesResponder handles the response to the ResourceChanges request. The method always
// closes the http.Response Body.
func (client BaseClient) ResourceChangesResponder(resp *http.Response) (result ResourceChangeList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Resources queries the resources managed by Azure Resource Manager for scopes specified in the request.
// Parameters:
// query - request specifying query and its options.
func (client BaseClient) Resources(ctx context.Context, query QueryRequest) (result QueryResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.Resources")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: query,
			Constraints: []validation.Constraint{{Target: "query.Query", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "query.Options", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "query.Options.Top", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "query.Options.Top", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
							{Target: "query.Options.Top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
						}},
						{Target: "query.Options.Skip", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "query.Options.Skip", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil}}},
					}}}}}); err != nil {
		return result, validation.NewError("resourcegraph.BaseClient", "Resources", err.Error())
	}

	req, err := client.ResourcesPreparer(ctx, query)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcegraph.BaseClient", "Resources", nil, "Failure preparing request")
		return
	}

	resp, err := client.ResourcesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resourcegraph.BaseClient", "Resources", resp, "Failure sending request")
		return
	}

	result, err = client.ResourcesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcegraph.BaseClient", "Resources", resp, "Failure responding to request")
		return
	}

	return
}

// ResourcesPreparer prepares the Resources request.
func (client BaseClient) ResourcesPreparer(ctx context.Context, query QueryRequest) (*http.Request, error) {
	const APIVersion = "2021-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.ResourceGraph/resources"),
		autorest.WithJSON(query),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ResourcesSender sends the Resources request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) ResourcesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ResourcesResponder handles the response to the Resources request. The method always
// closes the http.Response Body.
func (client BaseClient) ResourcesResponder(resp *http.Response) (result QueryResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ResourcesHistory list all snapshots of a resource for a given time interval.
func (client BaseClient) ResourcesHistory(ctx context.Context, request ResourcesHistoryRequest) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.ResourcesHistory")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: request,
			Constraints: []validation.Constraint{{Target: "request.Options", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "request.Options.Interval", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "request.Options.Interval.Start", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "request.Options.Interval.End", Name: validation.Null, Rule: true, Chain: nil},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("resourcegraph.BaseClient", "ResourcesHistory", err.Error())
	}

	req, err := client.ResourcesHistoryPreparer(ctx, request)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcegraph.BaseClient", "ResourcesHistory", nil, "Failure preparing request")
		return
	}

	resp, err := client.ResourcesHistorySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resourcegraph.BaseClient", "ResourcesHistory", resp, "Failure sending request")
		return
	}

	result, err = client.ResourcesHistoryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcegraph.BaseClient", "ResourcesHistory", resp, "Failure responding to request")
		return
	}

	return
}

// ResourcesHistoryPreparer prepares the ResourcesHistory request.
func (client BaseClient) ResourcesHistoryPreparer(ctx context.Context, request ResourcesHistoryRequest) (*http.Request, error) {
	const APIVersion = "2020-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.ResourceGraph/resourcesHistory"),
		autorest.WithJSON(request),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ResourcesHistorySender sends the ResourcesHistory request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) ResourcesHistorySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ResourcesHistoryResponder handles the response to the ResourcesHistory request. The method always
// closes the http.Response Body.
func (client BaseClient) ResourcesHistoryResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
