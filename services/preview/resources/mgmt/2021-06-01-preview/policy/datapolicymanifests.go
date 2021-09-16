package policy

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

// DataPolicyManifestsClient is the client for the DataPolicyManifests methods of the Policy service.
type DataPolicyManifestsClient struct {
	BaseClient
}

// NewDataPolicyManifestsClient creates an instance of the DataPolicyManifestsClient client.
func NewDataPolicyManifestsClient(subscriptionID string) DataPolicyManifestsClient {
	return NewDataPolicyManifestsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDataPolicyManifestsClientWithBaseURI creates an instance of the DataPolicyManifestsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewDataPolicyManifestsClientWithBaseURI(baseURI string, subscriptionID string) DataPolicyManifestsClient {
	return DataPolicyManifestsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetByPolicyMode this operation retrieves the data policy manifest with the given policy mode.
// Parameters:
// policyMode - the policy mode of the data policy manifest to get.
func (client DataPolicyManifestsClient) GetByPolicyMode(ctx context.Context, policyMode string) (result DataPolicyManifest, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataPolicyManifestsClient.GetByPolicyMode")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetByPolicyModePreparer(ctx, policyMode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DataPolicyManifestsClient", "GetByPolicyMode", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByPolicyModeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.DataPolicyManifestsClient", "GetByPolicyMode", resp, "Failure sending request")
		return
	}

	result, err = client.GetByPolicyModeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DataPolicyManifestsClient", "GetByPolicyMode", resp, "Failure responding to request")
		return
	}

	return
}

// GetByPolicyModePreparer prepares the GetByPolicyMode request.
func (client DataPolicyManifestsClient) GetByPolicyModePreparer(ctx context.Context, policyMode string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyMode": autorest.Encode("path", policyMode),
	}

	const APIVersion = "2020-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Authorization/dataPolicyManifests/{policyMode}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByPolicyModeSender sends the GetByPolicyMode request. The method will close the
// http.Response Body if it receives an error.
func (client DataPolicyManifestsClient) GetByPolicyModeSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetByPolicyModeResponder handles the response to the GetByPolicyMode request. The method always
// closes the http.Response Body.
func (client DataPolicyManifestsClient) GetByPolicyModeResponder(resp *http.Response) (result DataPolicyManifest, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List this operation retrieves a list of all the data policy manifests that match the optional given $filter. Valid
// values for $filter are: "$filter=namespace eq '{0}'". If $filter is not provided, the unfiltered list includes all
// data policy manifests for data resource types. If $filter=namespace is provided, the returned list only includes all
// data policy manifests that have a namespace matching the provided value.
// Parameters:
// filter - the filter to apply on the operation. Valid values for $filter are: "namespace eq '{value}'". If
// $filter is not provided, no filtering is performed. If $filter=namespace eq '{value}' is provided, the
// returned list only includes all data policy manifests that have a namespace matching the provided value.
func (client DataPolicyManifestsClient) List(ctx context.Context, filter string) (result DataPolicyManifestListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataPolicyManifestsClient.List")
		defer func() {
			sc := -1
			if result.dpmlr.Response.Response != nil {
				sc = result.dpmlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DataPolicyManifestsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.dpmlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.DataPolicyManifestsClient", "List", resp, "Failure sending request")
		return
	}

	result.dpmlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DataPolicyManifestsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.dpmlr.hasNextLink() && result.dpmlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client DataPolicyManifestsClient) ListPreparer(ctx context.Context, filter string) (*http.Request, error) {
	const APIVersion = "2020-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = filter
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Authorization/dataPolicyManifests"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DataPolicyManifestsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DataPolicyManifestsClient) ListResponder(resp *http.Response) (result DataPolicyManifestListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client DataPolicyManifestsClient) listNextResults(ctx context.Context, lastResults DataPolicyManifestListResult) (result DataPolicyManifestListResult, err error) {
	req, err := lastResults.dataPolicyManifestListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policy.DataPolicyManifestsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policy.DataPolicyManifestsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DataPolicyManifestsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client DataPolicyManifestsClient) ListComplete(ctx context.Context, filter string) (result DataPolicyManifestListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataPolicyManifestsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, filter)
	return
}
