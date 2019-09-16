package vmwarecloudsimple

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

// DedicatedCloudNodesClient is the description of the new service
type DedicatedCloudNodesClient struct {
	BaseClient
}

// NewDedicatedCloudNodesClient creates an instance of the DedicatedCloudNodesClient client.
func NewDedicatedCloudNodesClient(referer string, subscriptionID string) DedicatedCloudNodesClient {
	return NewDedicatedCloudNodesClientWithBaseURI(DefaultBaseURI, referer, subscriptionID)
}

// NewDedicatedCloudNodesClientWithBaseURI creates an instance of the DedicatedCloudNodesClient client.
func NewDedicatedCloudNodesClientWithBaseURI(baseURI string, referer string, subscriptionID string) DedicatedCloudNodesClient {
	return DedicatedCloudNodesClient{NewWithBaseURI(baseURI, referer, subscriptionID)}
}

// CreateOrUpdate returns dedicated cloud node by its name
// Parameters:
// resourceGroupName - the name of the resource group
// dedicatedCloudNodeName - dedicated cloud node name
// dedicatedCloudNodeRequest - create Dedicated Cloud Node request
func (client DedicatedCloudNodesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, dedicatedCloudNodeName string, dedicatedCloudNodeRequest CreateDedicatedCloudNodeRequest) (result DedicatedCloudNodesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedCloudNodesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: dedicatedCloudNodeName,
			Constraints: []validation.Constraint{{Target: "dedicatedCloudNodeName", Name: validation.Pattern, Rule: `^[-a-zA-Z0-9]+$`, Chain: nil}}},
		{TargetValue: dedicatedCloudNodeRequest,
			Constraints: []validation.Constraint{{Target: "dedicatedCloudNodeRequest.CreateDedicatedCloudNodeProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "dedicatedCloudNodeRequest.CreateDedicatedCloudNodeProperties.NodesCount", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("vmwarecloudsimple.DedicatedCloudNodesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, dedicatedCloudNodeName, dedicatedCloudNodeRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudNodesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudNodesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DedicatedCloudNodesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, dedicatedCloudNodeName string, dedicatedCloudNodeRequest CreateDedicatedCloudNodeRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dedicatedCloudNodeName": autorest.Encode("path", dedicatedCloudNodeName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes/{dedicatedCloudNodeName}", pathParameters),
		autorest.WithJSON(dedicatedCloudNodeRequest),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("Referer", client.Referer))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DedicatedCloudNodesClient) CreateOrUpdateSender(req *http.Request) (future DedicatedCloudNodesCreateOrUpdateFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DedicatedCloudNodesClient) CreateOrUpdateResponder(resp *http.Response) (result DedicatedCloudNode, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete dedicated cloud node
// Parameters:
// resourceGroupName - the name of the resource group
// dedicatedCloudNodeName - dedicated cloud node name
func (client DedicatedCloudNodesClient) Delete(ctx context.Context, resourceGroupName string, dedicatedCloudNodeName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedCloudNodesClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: dedicatedCloudNodeName,
			Constraints: []validation.Constraint{{Target: "dedicatedCloudNodeName", Name: validation.Pattern, Rule: `^[-a-zA-Z0-9]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("vmwarecloudsimple.DedicatedCloudNodesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, dedicatedCloudNodeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudNodesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudNodesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudNodesClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DedicatedCloudNodesClient) DeletePreparer(ctx context.Context, resourceGroupName string, dedicatedCloudNodeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dedicatedCloudNodeName": autorest.Encode("path", dedicatedCloudNodeName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes/{dedicatedCloudNodeName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DedicatedCloudNodesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DedicatedCloudNodesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get returns dedicated cloud node
// Parameters:
// resourceGroupName - the name of the resource group
// dedicatedCloudNodeName - dedicated cloud node name
func (client DedicatedCloudNodesClient) Get(ctx context.Context, resourceGroupName string, dedicatedCloudNodeName string) (result DedicatedCloudNode, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedCloudNodesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: dedicatedCloudNodeName,
			Constraints: []validation.Constraint{{Target: "dedicatedCloudNodeName", Name: validation.Pattern, Rule: `^[-a-zA-Z0-9]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("vmwarecloudsimple.DedicatedCloudNodesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, dedicatedCloudNodeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudNodesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudNodesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudNodesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client DedicatedCloudNodesClient) GetPreparer(ctx context.Context, resourceGroupName string, dedicatedCloudNodeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dedicatedCloudNodeName": autorest.Encode("path", dedicatedCloudNodeName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes/{dedicatedCloudNodeName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DedicatedCloudNodesClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DedicatedCloudNodesClient) GetResponder(resp *http.Response) (result DedicatedCloudNode, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup returns list of dedicate cloud nodes within resource group
// Parameters:
// resourceGroupName - the name of the resource group
// filter - the filter to apply on the list operation
// top - the maximum number of record sets to return
// skipToken - to be used by nextLink implementation
func (client DedicatedCloudNodesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, top *int32, skipToken string) (result DedicatedCloudNodeListResponsePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedCloudNodesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.dcnlr.Response.Response != nil {
				sc = result.dcnlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, filter, top, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudNodesClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.dcnlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudNodesClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.dcnlr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudNodesClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client DedicatedCloudNodesClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, filter string, top *int32, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client DedicatedCloudNodesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client DedicatedCloudNodesClient) ListByResourceGroupResponder(resp *http.Response) (result DedicatedCloudNodeListResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client DedicatedCloudNodesClient) listByResourceGroupNextResults(ctx context.Context, lastResults DedicatedCloudNodeListResponse) (result DedicatedCloudNodeListResponse, err error) {
	req, err := lastResults.dedicatedCloudNodeListResponsePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudNodesClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudNodesClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudNodesClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client DedicatedCloudNodesClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string, top *int32, skipToken string) (result DedicatedCloudNodeListResponseIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedCloudNodesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, filter, top, skipToken)
	return
}

// ListBySubscription returns list of dedicate cloud nodes within subscription
// Parameters:
// filter - the filter to apply on the list operation
// top - the maximum number of record sets to return
// skipToken - to be used by nextLink implementation
func (client DedicatedCloudNodesClient) ListBySubscription(ctx context.Context, filter string, top *int32, skipToken string) (result DedicatedCloudNodeListResponsePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedCloudNodesClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.dcnlr.Response.Response != nil {
				sc = result.dcnlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx, filter, top, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudNodesClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.dcnlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudNodesClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.dcnlr, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudNodesClient", "ListBySubscription", resp, "Failure responding to request")
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client DedicatedCloudNodesClient) ListBySubscriptionPreparer(ctx context.Context, filter string, top *int32, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client DedicatedCloudNodesClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client DedicatedCloudNodesClient) ListBySubscriptionResponder(resp *http.Response) (result DedicatedCloudNodeListResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client DedicatedCloudNodesClient) listBySubscriptionNextResults(ctx context.Context, lastResults DedicatedCloudNodeListResponse) (result DedicatedCloudNodeListResponse, err error) {
	req, err := lastResults.dedicatedCloudNodeListResponsePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudNodesClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudNodesClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudNodesClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client DedicatedCloudNodesClient) ListBySubscriptionComplete(ctx context.Context, filter string, top *int32, skipToken string) (result DedicatedCloudNodeListResponseIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedCloudNodesClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscription(ctx, filter, top, skipToken)
	return
}

// Update patches dedicated node properties
// Parameters:
// resourceGroupName - the name of the resource group
// dedicatedCloudNodeName - dedicated cloud node name
// dedicatedCloudNodeRequest - patch Dedicated Cloud Node request
func (client DedicatedCloudNodesClient) Update(ctx context.Context, resourceGroupName string, dedicatedCloudNodeName string, dedicatedCloudNodeRequest PatchPayload) (result DedicatedCloudNode, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedCloudNodesClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: dedicatedCloudNodeName,
			Constraints: []validation.Constraint{{Target: "dedicatedCloudNodeName", Name: validation.Pattern, Rule: `^[-a-zA-Z0-9]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("vmwarecloudsimple.DedicatedCloudNodesClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, dedicatedCloudNodeName, dedicatedCloudNodeRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudNodesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudNodesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.DedicatedCloudNodesClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client DedicatedCloudNodesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, dedicatedCloudNodeName string, dedicatedCloudNodeRequest PatchPayload) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dedicatedCloudNodeName": autorest.Encode("path", dedicatedCloudNodeName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes/{dedicatedCloudNodeName}", pathParameters),
		autorest.WithJSON(dedicatedCloudNodeRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client DedicatedCloudNodesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client DedicatedCloudNodesClient) UpdateResponder(resp *http.Response) (result DedicatedCloudNode, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
