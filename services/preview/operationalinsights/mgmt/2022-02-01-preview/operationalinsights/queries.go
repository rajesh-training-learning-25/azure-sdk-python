package operationalinsights

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

// QueriesClient is the operational Insights Client
type QueriesClient struct {
	BaseClient
}

// NewQueriesClient creates an instance of the QueriesClient client.
func NewQueriesClient(subscriptionID string) QueriesClient {
	return NewQueriesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewQueriesClientWithBaseURI creates an instance of the QueriesClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewQueriesClientWithBaseURI(baseURI string, subscriptionID string) QueriesClient {
	return QueriesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Delete deletes a specific Query defined within an Log Analytics QueryPack.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// queryPackName - the name of the Log Analytics QueryPack resource.
// ID - the id of a specific query defined in the Log Analytics QueryPack
func (client QueriesClient) Delete(ctx context.Context, resourceGroupName string, queryPackName string, ID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QueriesClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("operationalinsights.QueriesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, queryPackName, ID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.QueriesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "operationalinsights.QueriesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.QueriesClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client QueriesClient) DeletePreparer(ctx context.Context, resourceGroupName string, queryPackName string, ID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"id":                autorest.Encode("path", ID),
		"queryPackName":     autorest.Encode("path", queryPackName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/queryPacks/{queryPackName}/queries/{id}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client QueriesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client QueriesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a specific Log Analytics Query defined within a Log Analytics QueryPack.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// queryPackName - the name of the Log Analytics QueryPack resource.
// ID - the id of a specific query defined in the Log Analytics QueryPack
func (client QueriesClient) Get(ctx context.Context, resourceGroupName string, queryPackName string, ID string) (result LogAnalyticsQueryPackQuery, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QueriesClient.Get")
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
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("operationalinsights.QueriesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, queryPackName, ID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.QueriesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "operationalinsights.QueriesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.QueriesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client QueriesClient) GetPreparer(ctx context.Context, resourceGroupName string, queryPackName string, ID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"id":                autorest.Encode("path", ID),
		"queryPackName":     autorest.Encode("path", queryPackName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/queryPacks/{queryPackName}/queries/{id}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client QueriesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client QueriesClient) GetResponder(resp *http.Response) (result LogAnalyticsQueryPackQuery, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of Queries defined within a Log Analytics QueryPack.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// queryPackName - the name of the Log Analytics QueryPack resource.
// top - maximum items returned in page.
// includeBody - flag indicating whether or not to return the body of each applicable query. If false, only
// return the query information.
// skipToken - base64 encoded token used to fetch the next page of items. Default is null.
func (client QueriesClient) List(ctx context.Context, resourceGroupName string, queryPackName string, top *int64, includeBody *bool, skipToken string) (result LogAnalyticsQueryPackQueryListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QueriesClient.List")
		defer func() {
			sc := -1
			if result.laqpqlr.Response.Response != nil {
				sc = result.laqpqlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("operationalinsights.QueriesClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, queryPackName, top, includeBody, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.QueriesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.laqpqlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "operationalinsights.QueriesClient", "List", resp, "Failure sending request")
		return
	}

	result.laqpqlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.QueriesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.laqpqlr.hasNextLink() && result.laqpqlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client QueriesClient) ListPreparer(ctx context.Context, resourceGroupName string, queryPackName string, top *int64, includeBody *bool, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"queryPackName":     autorest.Encode("path", queryPackName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if includeBody != nil {
		queryParameters["includeBody"] = autorest.Encode("query", *includeBody)
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/queryPacks/{queryPackName}/queries", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client QueriesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client QueriesClient) ListResponder(resp *http.Response) (result LogAnalyticsQueryPackQueryListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client QueriesClient) listNextResults(ctx context.Context, lastResults LogAnalyticsQueryPackQueryListResult) (result LogAnalyticsQueryPackQueryListResult, err error) {
	req, err := lastResults.logAnalyticsQueryPackQueryListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "operationalinsights.QueriesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "operationalinsights.QueriesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.QueriesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client QueriesClient) ListComplete(ctx context.Context, resourceGroupName string, queryPackName string, top *int64, includeBody *bool, skipToken string) (result LogAnalyticsQueryPackQueryListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QueriesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, queryPackName, top, includeBody, skipToken)
	return
}

// Put adds or Updates a specific Query within a Log Analytics QueryPack.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// queryPackName - the name of the Log Analytics QueryPack resource.
// ID - the id of a specific query defined in the Log Analytics QueryPack
// queryPayload - properties that need to be specified to create a new query and add it to a Log Analytics
// QueryPack.
func (client QueriesClient) Put(ctx context.Context, resourceGroupName string, queryPackName string, ID string, queryPayload LogAnalyticsQueryPackQuery) (result LogAnalyticsQueryPackQuery, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QueriesClient.Put")
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
		{TargetValue: queryPayload,
			Constraints: []validation.Constraint{{Target: "queryPayload.LogAnalyticsQueryPackQueryProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "queryPayload.LogAnalyticsQueryPackQueryProperties.DisplayName", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "queryPayload.LogAnalyticsQueryPackQueryProperties.Body", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("operationalinsights.QueriesClient", "Put", err.Error())
	}

	req, err := client.PutPreparer(ctx, resourceGroupName, queryPackName, ID, queryPayload)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.QueriesClient", "Put", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "operationalinsights.QueriesClient", "Put", resp, "Failure sending request")
		return
	}

	result, err = client.PutResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.QueriesClient", "Put", resp, "Failure responding to request")
		return
	}

	return
}

// PutPreparer prepares the Put request.
func (client QueriesClient) PutPreparer(ctx context.Context, resourceGroupName string, queryPackName string, ID string, queryPayload LogAnalyticsQueryPackQuery) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"id":                autorest.Encode("path", ID),
		"queryPackName":     autorest.Encode("path", queryPackName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/queryPacks/{queryPackName}/queries/{id}", pathParameters),
		autorest.WithJSON(queryPayload),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutSender sends the Put request. The method will close the
// http.Response Body if it receives an error.
func (client QueriesClient) PutSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// PutResponder handles the response to the Put request. The method always
// closes the http.Response Body.
func (client QueriesClient) PutResponder(resp *http.Response) (result LogAnalyticsQueryPackQuery, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Search search a list of Queries defined within a Log Analytics QueryPack according to given search properties.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// queryPackName - the name of the Log Analytics QueryPack resource.
// querySearchProperties - properties by which to search queries in the given Log Analytics QueryPack.
// top - maximum items returned in page.
// includeBody - flag indicating whether or not to return the body of each applicable query. If false, only
// return the query information.
// skipToken - base64 encoded token used to fetch the next page of items. Default is null.
func (client QueriesClient) Search(ctx context.Context, resourceGroupName string, queryPackName string, querySearchProperties LogAnalyticsQueryPackQuerySearchProperties, top *int64, includeBody *bool, skipToken string) (result LogAnalyticsQueryPackQueryListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QueriesClient.Search")
		defer func() {
			sc := -1
			if result.laqpqlr.Response.Response != nil {
				sc = result.laqpqlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("operationalinsights.QueriesClient", "Search", err.Error())
	}

	result.fn = client.searchNextResults
	req, err := client.SearchPreparer(ctx, resourceGroupName, queryPackName, querySearchProperties, top, includeBody, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.QueriesClient", "Search", nil, "Failure preparing request")
		return
	}

	resp, err := client.SearchSender(req)
	if err != nil {
		result.laqpqlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "operationalinsights.QueriesClient", "Search", resp, "Failure sending request")
		return
	}

	result.laqpqlr, err = client.SearchResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.QueriesClient", "Search", resp, "Failure responding to request")
		return
	}
	if result.laqpqlr.hasNextLink() && result.laqpqlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// SearchPreparer prepares the Search request.
func (client QueriesClient) SearchPreparer(ctx context.Context, resourceGroupName string, queryPackName string, querySearchProperties LogAnalyticsQueryPackQuerySearchProperties, top *int64, includeBody *bool, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"queryPackName":     autorest.Encode("path", queryPackName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if includeBody != nil {
		queryParameters["includeBody"] = autorest.Encode("query", *includeBody)
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/queryPacks/{queryPackName}/queries/search", pathParameters),
		autorest.WithJSON(querySearchProperties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SearchSender sends the Search request. The method will close the
// http.Response Body if it receives an error.
func (client QueriesClient) SearchSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// SearchResponder handles the response to the Search request. The method always
// closes the http.Response Body.
func (client QueriesClient) SearchResponder(resp *http.Response) (result LogAnalyticsQueryPackQueryListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// searchNextResults retrieves the next set of results, if any.
func (client QueriesClient) searchNextResults(ctx context.Context, lastResults LogAnalyticsQueryPackQueryListResult) (result LogAnalyticsQueryPackQueryListResult, err error) {
	req, err := lastResults.logAnalyticsQueryPackQueryListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "operationalinsights.QueriesClient", "searchNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.SearchSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "operationalinsights.QueriesClient", "searchNextResults", resp, "Failure sending next results request")
	}
	result, err = client.SearchResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.QueriesClient", "searchNextResults", resp, "Failure responding to next results request")
	}
	return
}

// SearchComplete enumerates all values, automatically crossing page boundaries as required.
func (client QueriesClient) SearchComplete(ctx context.Context, resourceGroupName string, queryPackName string, querySearchProperties LogAnalyticsQueryPackQuerySearchProperties, top *int64, includeBody *bool, skipToken string) (result LogAnalyticsQueryPackQueryListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QueriesClient.Search")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.Search(ctx, resourceGroupName, queryPackName, querySearchProperties, top, includeBody, skipToken)
	return
}

// Update adds or Updates a specific Query within a Log Analytics QueryPack.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// queryPackName - the name of the Log Analytics QueryPack resource.
// ID - the id of a specific query defined in the Log Analytics QueryPack
// queryPayload - properties that need to be specified to create a new query and add it to a Log Analytics
// QueryPack.
func (client QueriesClient) Update(ctx context.Context, resourceGroupName string, queryPackName string, ID string, queryPayload LogAnalyticsQueryPackQuery) (result LogAnalyticsQueryPackQuery, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QueriesClient.Update")
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
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("operationalinsights.QueriesClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, queryPackName, ID, queryPayload)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.QueriesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "operationalinsights.QueriesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.QueriesClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client QueriesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, queryPackName string, ID string, queryPayload LogAnalyticsQueryPackQuery) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"id":                autorest.Encode("path", ID),
		"queryPackName":     autorest.Encode("path", queryPackName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/queryPacks/{queryPackName}/queries/{id}", pathParameters),
		autorest.WithJSON(queryPayload),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client QueriesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client QueriesClient) UpdateResponder(resp *http.Response) (result LogAnalyticsQueryPackQuery, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
