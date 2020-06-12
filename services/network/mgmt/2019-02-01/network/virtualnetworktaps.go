package network

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

// VirtualNetworkTapsClient is the network Client
type VirtualNetworkTapsClient struct {
	BaseClient
}

// NewVirtualNetworkTapsClient creates an instance of the VirtualNetworkTapsClient client.
func NewVirtualNetworkTapsClient(subscriptionID string) VirtualNetworkTapsClient {
	return NewVirtualNetworkTapsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualNetworkTapsClientWithBaseURI creates an instance of the VirtualNetworkTapsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewVirtualNetworkTapsClientWithBaseURI(baseURI string, subscriptionID string) VirtualNetworkTapsClient {
	return VirtualNetworkTapsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a Virtual Network Tap.
// Parameters:
// resourceGroupName - the name of the resource group.
// tapName - the name of the virtual network tap.
// parameters - parameters supplied to the create or update virtual network tap operation.
func (client VirtualNetworkTapsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, tapName string, parameters VirtualNetworkTap) (result VirtualNetworkTapsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworkTapsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.VirtualNetworkTapPropertiesFormat", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.VirtualNetworkTapPropertiesFormat.DestinationNetworkInterfaceIPConfiguration", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.VirtualNetworkTapPropertiesFormat.DestinationNetworkInterfaceIPConfiguration.InterfaceIPConfigurationPropertiesFormat", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.VirtualNetworkTapPropertiesFormat.DestinationNetworkInterfaceIPConfiguration.InterfaceIPConfigurationPropertiesFormat.PublicIPAddress", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.VirtualNetworkTapPropertiesFormat.DestinationNetworkInterfaceIPConfiguration.InterfaceIPConfigurationPropertiesFormat.PublicIPAddress.PublicIPAddressPropertiesFormat", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.VirtualNetworkTapPropertiesFormat.DestinationNetworkInterfaceIPConfiguration.InterfaceIPConfigurationPropertiesFormat.PublicIPAddress.PublicIPAddressPropertiesFormat.IPConfiguration", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.VirtualNetworkTapPropertiesFormat.DestinationNetworkInterfaceIPConfiguration.InterfaceIPConfigurationPropertiesFormat.PublicIPAddress.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat", Name: validation.Null, Rule: false,
										Chain: []validation.Constraint{{Target: "parameters.VirtualNetworkTapPropertiesFormat.DestinationNetworkInterfaceIPConfiguration.InterfaceIPConfigurationPropertiesFormat.PublicIPAddress.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat.PublicIPAddress", Name: validation.Null, Rule: false, Chain: nil}}},
									}},
								}},
							}},
						}},
					}},
					{Target: "parameters.VirtualNetworkTapPropertiesFormat.DestinationLoadBalancerFrontEndIPConfiguration", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.VirtualNetworkTapPropertiesFormat.DestinationLoadBalancerFrontEndIPConfiguration.FrontendIPConfigurationPropertiesFormat", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.VirtualNetworkTapPropertiesFormat.DestinationLoadBalancerFrontEndIPConfiguration.FrontendIPConfigurationPropertiesFormat.PublicIPAddress", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.VirtualNetworkTapPropertiesFormat.DestinationLoadBalancerFrontEndIPConfiguration.FrontendIPConfigurationPropertiesFormat.PublicIPAddress.PublicIPAddressPropertiesFormat", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.VirtualNetworkTapPropertiesFormat.DestinationLoadBalancerFrontEndIPConfiguration.FrontendIPConfigurationPropertiesFormat.PublicIPAddress.PublicIPAddressPropertiesFormat.IPConfiguration", Name: validation.Null, Rule: false,
										Chain: []validation.Constraint{{Target: "parameters.VirtualNetworkTapPropertiesFormat.DestinationLoadBalancerFrontEndIPConfiguration.FrontendIPConfigurationPropertiesFormat.PublicIPAddress.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat", Name: validation.Null, Rule: false,
											Chain: []validation.Constraint{{Target: "parameters.VirtualNetworkTapPropertiesFormat.DestinationLoadBalancerFrontEndIPConfiguration.FrontendIPConfigurationPropertiesFormat.PublicIPAddress.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat.PublicIPAddress", Name: validation.Null, Rule: false, Chain: nil}}},
										}},
									}},
								}},
							}},
						}},
				}}}}}); err != nil {
		return result, validation.NewError("network.VirtualNetworkTapsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, tapName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkTapsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkTapsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VirtualNetworkTapsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, tapName string, parameters VirtualNetworkTap) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"tapName":           autorest.Encode("path", tapName),
	}

	const APIVersion = "2019-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkTapsClient) CreateOrUpdateSender(req *http.Request) (future VirtualNetworkTapsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client VirtualNetworkTapsClient) CreateOrUpdateResponder(resp *http.Response) (result VirtualNetworkTap, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified virtual network tap.
// Parameters:
// resourceGroupName - the name of the resource group.
// tapName - the name of the virtual network tap.
func (client VirtualNetworkTapsClient) Delete(ctx context.Context, resourceGroupName string, tapName string) (result VirtualNetworkTapsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworkTapsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, tapName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkTapsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkTapsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client VirtualNetworkTapsClient) DeletePreparer(ctx context.Context, resourceGroupName string, tapName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"tapName":           autorest.Encode("path", tapName),
	}

	const APIVersion = "2019-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkTapsClient) DeleteSender(req *http.Request) (future VirtualNetworkTapsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client VirtualNetworkTapsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets information about the specified virtual network tap.
// Parameters:
// resourceGroupName - the name of the resource group.
// tapName - the name of virtual network tap.
func (client VirtualNetworkTapsClient) Get(ctx context.Context, resourceGroupName string, tapName string) (result VirtualNetworkTap, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworkTapsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, tapName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkTapsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkTapsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkTapsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client VirtualNetworkTapsClient) GetPreparer(ctx context.Context, resourceGroupName string, tapName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"tapName":           autorest.Encode("path", tapName),
	}

	const APIVersion = "2019-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkTapsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VirtualNetworkTapsClient) GetResponder(resp *http.Response) (result VirtualNetworkTap, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAll gets all the VirtualNetworkTaps in a subscription.
func (client VirtualNetworkTapsClient) ListAll(ctx context.Context) (result VirtualNetworkTapListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworkTapsClient.ListAll")
		defer func() {
			sc := -1
			if result.vntlr.Response.Response != nil {
				sc = result.vntlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listAllNextResults
	req, err := client.ListAllPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkTapsClient", "ListAll", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.vntlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkTapsClient", "ListAll", resp, "Failure sending request")
		return
	}

	result.vntlr, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkTapsClient", "ListAll", resp, "Failure responding to request")
	}

	return
}

// ListAllPreparer prepares the ListAll request.
func (client VirtualNetworkTapsClient) ListAllPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualNetworkTaps", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAllSender sends the ListAll request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkTapsClient) ListAllSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListAllResponder handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (client VirtualNetworkTapsClient) ListAllResponder(resp *http.Response) (result VirtualNetworkTapListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listAllNextResults retrieves the next set of results, if any.
func (client VirtualNetworkTapsClient) listAllNextResults(ctx context.Context, lastResults VirtualNetworkTapListResult) (result VirtualNetworkTapListResult, err error) {
	req, err := lastResults.virtualNetworkTapListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworkTapsClient", "listAllNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworkTapsClient", "listAllNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkTapsClient", "listAllNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListAllComplete enumerates all values, automatically crossing page boundaries as required.
func (client VirtualNetworkTapsClient) ListAllComplete(ctx context.Context) (result VirtualNetworkTapListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworkTapsClient.ListAll")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListAll(ctx)
	return
}

// ListByResourceGroup gets all the VirtualNetworkTaps in a subscription.
// Parameters:
// resourceGroupName - the name of the resource group.
func (client VirtualNetworkTapsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result VirtualNetworkTapListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworkTapsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.vntlr.Response.Response != nil {
				sc = result.vntlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkTapsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.vntlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkTapsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.vntlr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkTapsClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client VirtualNetworkTapsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkTapsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client VirtualNetworkTapsClient) ListByResourceGroupResponder(resp *http.Response) (result VirtualNetworkTapListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client VirtualNetworkTapsClient) listByResourceGroupNextResults(ctx context.Context, lastResults VirtualNetworkTapListResult) (result VirtualNetworkTapListResult, err error) {
	req, err := lastResults.virtualNetworkTapListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworkTapsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworkTapsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkTapsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client VirtualNetworkTapsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result VirtualNetworkTapListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworkTapsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// UpdateTags updates an VirtualNetworkTap tags.
// Parameters:
// resourceGroupName - the name of the resource group.
// tapName - the name of the tap.
// tapParameters - parameters supplied to update VirtualNetworkTap tags.
func (client VirtualNetworkTapsClient) UpdateTags(ctx context.Context, resourceGroupName string, tapName string, tapParameters TagsObject) (result VirtualNetworkTapsUpdateTagsFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworkTapsClient.UpdateTags")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdateTagsPreparer(ctx, resourceGroupName, tapName, tapParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkTapsClient", "UpdateTags", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateTagsSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkTapsClient", "UpdateTags", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdateTagsPreparer prepares the UpdateTags request.
func (client VirtualNetworkTapsClient) UpdateTagsPreparer(ctx context.Context, resourceGroupName string, tapName string, tapParameters TagsObject) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"tapName":           autorest.Encode("path", tapName),
	}

	const APIVersion = "2019-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}", pathParameters),
		autorest.WithJSON(tapParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateTagsSender sends the UpdateTags request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkTapsClient) UpdateTagsSender(req *http.Request) (future VirtualNetworkTapsUpdateTagsFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// UpdateTagsResponder handles the response to the UpdateTags request. The method always
// closes the http.Response Body.
func (client VirtualNetworkTapsClient) UpdateTagsResponder(resp *http.Response) (result VirtualNetworkTap, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
