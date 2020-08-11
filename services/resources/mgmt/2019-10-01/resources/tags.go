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

// TagsClient is the provides operations for working with resources and resource groups.
type TagsClient struct {
	BaseClient
}

// NewTagsClient creates an instance of the TagsClient client.
func NewTagsClient(subscriptionID string) TagsClient {
	return NewTagsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewTagsClientWithBaseURI creates an instance of the TagsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewTagsClientWithBaseURI(baseURI string, subscriptionID string) TagsClient {
	return TagsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate this operation allows adding a name to the list of predefined tag names for the given subscription. A
// tag name can have a maximum of 512 characters and is case-insensitive. Tag names cannot have the following prefixes
// which are reserved for Azure use: 'microsoft', 'azure', 'windows'.
// Parameters:
// tagName - the name of the tag to create.
func (client TagsClient) CreateOrUpdate(ctx context.Context, tagName string) (result TagDetails, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TagsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, tagName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.TagsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resources.TagsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.TagsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client TagsClient) CreateOrUpdatePreparer(ctx context.Context, tagName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
		"tagName":        autorest.Encode("path", tagName),
	}

	const APIVersion = "2019-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/tagNames/{tagName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client TagsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client TagsClient) CreateOrUpdateResponder(resp *http.Response) (result TagDetails, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdateAtScope this operation allows adding or replacing the entire set of tags on the specified resource or
// subscription. The specified entity can have a maximum of 50 tags.
// Parameters:
// scope - the resource scope.
func (client TagsClient) CreateOrUpdateAtScope(ctx context.Context, scope string, parameters TagsResource) (result TagsResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TagsClient.CreateOrUpdateAtScope")
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
			Constraints: []validation.Constraint{{Target: "parameters.Properties", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("resources.TagsClient", "CreateOrUpdateAtScope", err.Error())
	}

	req, err := client.CreateOrUpdateAtScopePreparer(ctx, scope, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.TagsClient", "CreateOrUpdateAtScope", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateAtScopeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resources.TagsClient", "CreateOrUpdateAtScope", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateAtScopeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.TagsClient", "CreateOrUpdateAtScope", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdateAtScopePreparer prepares the CreateOrUpdateAtScope request.
func (client TagsClient) CreateOrUpdateAtScopePreparer(ctx context.Context, scope string, parameters TagsResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scope": scope,
	}

	const APIVersion = "2019-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.ID = nil
	parameters.Name = nil
	parameters.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Resources/tags/default", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateAtScopeSender sends the CreateOrUpdateAtScope request. The method will close the
// http.Response Body if it receives an error.
func (client TagsClient) CreateOrUpdateAtScopeSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateOrUpdateAtScopeResponder handles the response to the CreateOrUpdateAtScope request. The method always
// closes the http.Response Body.
func (client TagsClient) CreateOrUpdateAtScopeResponder(resp *http.Response) (result TagsResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdateValue this operation allows adding a value to the list of predefined values for an existing predefined
// tag name. A tag value can have a maximum of 256 characters.
// Parameters:
// tagName - the name of the tag.
// tagValue - the value of the tag to create.
func (client TagsClient) CreateOrUpdateValue(ctx context.Context, tagName string, tagValue string) (result TagValue, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TagsClient.CreateOrUpdateValue")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdateValuePreparer(ctx, tagName, tagValue)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.TagsClient", "CreateOrUpdateValue", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateValueSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resources.TagsClient", "CreateOrUpdateValue", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateValueResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.TagsClient", "CreateOrUpdateValue", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdateValuePreparer prepares the CreateOrUpdateValue request.
func (client TagsClient) CreateOrUpdateValuePreparer(ctx context.Context, tagName string, tagValue string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
		"tagName":        autorest.Encode("path", tagName),
		"tagValue":       autorest.Encode("path", tagValue),
	}

	const APIVersion = "2019-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/tagNames/{tagName}/tagValues/{tagValue}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateValueSender sends the CreateOrUpdateValue request. The method will close the
// http.Response Body if it receives an error.
func (client TagsClient) CreateOrUpdateValueSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateValueResponder handles the response to the CreateOrUpdateValue request. The method always
// closes the http.Response Body.
func (client TagsClient) CreateOrUpdateValueResponder(resp *http.Response) (result TagValue, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete this operation allows deleting a name from the list of predefined tag names for the given subscription. The
// name being deleted must not be in use as a tag name for any resource. All predefined values for the given name must
// have already been deleted.
// Parameters:
// tagName - the name of the tag.
func (client TagsClient) Delete(ctx context.Context, tagName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TagsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, tagName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.TagsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "resources.TagsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.TagsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client TagsClient) DeletePreparer(ctx context.Context, tagName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
		"tagName":        autorest.Encode("path", tagName),
	}

	const APIVersion = "2019-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/tagNames/{tagName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client TagsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client TagsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteAtScope sends the delete at scope request.
// Parameters:
// scope - the resource scope.
func (client TagsClient) DeleteAtScope(ctx context.Context, scope string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TagsClient.DeleteAtScope")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteAtScopePreparer(ctx, scope)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.TagsClient", "DeleteAtScope", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteAtScopeSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "resources.TagsClient", "DeleteAtScope", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteAtScopeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.TagsClient", "DeleteAtScope", resp, "Failure responding to request")
	}

	return
}

// DeleteAtScopePreparer prepares the DeleteAtScope request.
func (client TagsClient) DeleteAtScopePreparer(ctx context.Context, scope string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scope": scope,
	}

	const APIVersion = "2019-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Resources/tags/default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteAtScopeSender sends the DeleteAtScope request. The method will close the
// http.Response Body if it receives an error.
func (client TagsClient) DeleteAtScopeSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteAtScopeResponder handles the response to the DeleteAtScope request. The method always
// closes the http.Response Body.
func (client TagsClient) DeleteAtScopeResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteValue this operation allows deleting a value from the list of predefined values for an existing predefined tag
// name. The value being deleted must not be in use as a tag value for the given tag name for any resource.
// Parameters:
// tagName - the name of the tag.
// tagValue - the value of the tag to delete.
func (client TagsClient) DeleteValue(ctx context.Context, tagName string, tagValue string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TagsClient.DeleteValue")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteValuePreparer(ctx, tagName, tagValue)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.TagsClient", "DeleteValue", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteValueSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "resources.TagsClient", "DeleteValue", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteValueResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.TagsClient", "DeleteValue", resp, "Failure responding to request")
	}

	return
}

// DeleteValuePreparer prepares the DeleteValue request.
func (client TagsClient) DeleteValuePreparer(ctx context.Context, tagName string, tagValue string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
		"tagName":        autorest.Encode("path", tagName),
		"tagValue":       autorest.Encode("path", tagValue),
	}

	const APIVersion = "2019-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/tagNames/{tagName}/tagValues/{tagValue}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteValueSender sends the DeleteValue request. The method will close the
// http.Response Body if it receives an error.
func (client TagsClient) DeleteValueSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteValueResponder handles the response to the DeleteValue request. The method always
// closes the http.Response Body.
func (client TagsClient) DeleteValueResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetAtScope sends the get at scope request.
// Parameters:
// scope - the resource scope.
func (client TagsClient) GetAtScope(ctx context.Context, scope string) (result TagsResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TagsClient.GetAtScope")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetAtScopePreparer(ctx, scope)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.TagsClient", "GetAtScope", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAtScopeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resources.TagsClient", "GetAtScope", resp, "Failure sending request")
		return
	}

	result, err = client.GetAtScopeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.TagsClient", "GetAtScope", resp, "Failure responding to request")
	}

	return
}

// GetAtScopePreparer prepares the GetAtScope request.
func (client TagsClient) GetAtScopePreparer(ctx context.Context, scope string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scope": scope,
	}

	const APIVersion = "2019-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Resources/tags/default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAtScopeSender sends the GetAtScope request. The method will close the
// http.Response Body if it receives an error.
func (client TagsClient) GetAtScopeSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetAtScopeResponder handles the response to the GetAtScope request. The method always
// closes the http.Response Body.
func (client TagsClient) GetAtScopeResponder(resp *http.Response) (result TagsResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List this operation performs a union of predefined tags, resource tags, resource group tags and subscription tags,
// and returns a summary of usage for each tag name and value under the given subscription. In case of a large number
// of tags, this operation may return a previously cached result.
func (client TagsClient) List(ctx context.Context) (result TagsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TagsClient.List")
		defer func() {
			sc := -1
			if result.tlr.Response.Response != nil {
				sc = result.tlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.TagsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.tlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resources.TagsClient", "List", resp, "Failure sending request")
		return
	}

	result.tlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.TagsClient", "List", resp, "Failure responding to request")
	}
	if result.tlr.hasNextLink() && result.tlr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client TagsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/tagNames", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client TagsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client TagsClient) ListResponder(resp *http.Response) (result TagsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client TagsClient) listNextResults(ctx context.Context, lastResults TagsListResult) (result TagsListResult, err error) {
	req, err := lastResults.tagsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.TagsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources.TagsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.TagsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client TagsClient) ListComplete(ctx context.Context) (result TagsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TagsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// UpdateAtScope this operation allows replacing, merging or selectively deleting tags on the specified resource or
// subscription. The specified entity can have a maximum of 50 tags at the end of the operation. The 'replace' option
// replaces the entire set of existing tags with a new set. The 'merge' option allows adding tags with new names and
// updating the values of tags with existing names. The 'delete' option allows selectively deleting tags based on given
// names or name/value pairs.
// Parameters:
// scope - the resource scope.
func (client TagsClient) UpdateAtScope(ctx context.Context, scope string, parameters TagsPatchResource) (result TagsResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TagsClient.UpdateAtScope")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdateAtScopePreparer(ctx, scope, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.TagsClient", "UpdateAtScope", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateAtScopeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resources.TagsClient", "UpdateAtScope", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateAtScopeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.TagsClient", "UpdateAtScope", resp, "Failure responding to request")
	}

	return
}

// UpdateAtScopePreparer prepares the UpdateAtScope request.
func (client TagsClient) UpdateAtScopePreparer(ctx context.Context, scope string, parameters TagsPatchResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scope": scope,
	}

	const APIVersion = "2019-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Resources/tags/default", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateAtScopeSender sends the UpdateAtScope request. The method will close the
// http.Response Body if it receives an error.
func (client TagsClient) UpdateAtScopeSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateAtScopeResponder handles the response to the UpdateAtScope request. The method always
// closes the http.Response Body.
func (client TagsClient) UpdateAtScopeResponder(resp *http.Response) (result TagsResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
