package consumption

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
	"net/http"
)

// PriceSheetClient is the consumption management client provides access to consumption resources for Azure Enterprise
// Subscriptions.
type PriceSheetClient struct {
	BaseClient
}

// NewPriceSheetClient creates an instance of the PriceSheetClient client.
func NewPriceSheetClient(subscriptionID string, grain Datagrain) PriceSheetClient {
	return NewPriceSheetClientWithBaseURI(DefaultBaseURI, subscriptionID, grain)
}

// NewPriceSheetClientWithBaseURI creates an instance of the PriceSheetClient client.
func NewPriceSheetClientWithBaseURI(baseURI string, subscriptionID string, grain Datagrain) PriceSheetClient {
	return PriceSheetClient{NewWithBaseURI(baseURI, subscriptionID, grain)}
}

// Get gets the price sheet for a scope by subscriptionId. Price sheet is available via this API only for May 1, 2014
// or later.
//
// expand is may be used to expand the properties/meterDetails within a price sheet. By default, these fields are
// not included when returning price sheet. skiptoken is skiptoken is only used if a previous operation returned a
// partial result. If a previous response contains a nextLink element, the value of the nextLink element will
// include a skiptoken parameter that specifies a starting point to use for subsequent calls. top is may be used to
// limit the number of results to the top N results.
func (client PriceSheetClient) Get(ctx context.Context, expand string, skiptoken string, top *int32) (result PriceSheetResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: 1000, Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("consumption.PriceSheetClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, expand, skiptoken, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.PriceSheetClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.PriceSheetClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.PriceSheetClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client PriceSheetClient) GetPreparer(ctx context.Context, expand string, skiptoken string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Consumption/pricesheets/default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PriceSheetClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PriceSheetClient) GetResponder(resp *http.Response) (result PriceSheetResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetByBillingPeriod get the price sheet for a scope by subscriptionId and billing period. Price sheet is available
// via this API only for May 1, 2014 or later.
//
// billingPeriodName is billing Period Name. expand is may be used to expand the properties/meterDetails within a
// price sheet. By default, these fields are not included when returning price sheet. skiptoken is skiptoken is
// only used if a previous operation returned a partial result. If a previous response contains a nextLink element,
// the value of the nextLink element will include a skiptoken parameter that specifies a starting point to use for
// subsequent calls. top is may be used to limit the number of results to the top N results.
func (client PriceSheetClient) GetByBillingPeriod(ctx context.Context, billingPeriodName string, expand string, skiptoken string, top *int32) (result PriceSheetResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: 1000, Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("consumption.PriceSheetClient", "GetByBillingPeriod", err.Error())
	}

	req, err := client.GetByBillingPeriodPreparer(ctx, billingPeriodName, expand, skiptoken, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.PriceSheetClient", "GetByBillingPeriod", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByBillingPeriodSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.PriceSheetClient", "GetByBillingPeriod", resp, "Failure sending request")
		return
	}

	result, err = client.GetByBillingPeriodResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.PriceSheetClient", "GetByBillingPeriod", resp, "Failure responding to request")
	}

	return
}

// GetByBillingPeriodPreparer prepares the GetByBillingPeriod request.
func (client PriceSheetClient) GetByBillingPeriodPreparer(ctx context.Context, billingPeriodName string, expand string, skiptoken string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingPeriodName": autorest.Encode("path", billingPeriodName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}/providers/Microsoft.Consumption/pricesheets/default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByBillingPeriodSender sends the GetByBillingPeriod request. The method will close the
// http.Response Body if it receives an error.
func (client PriceSheetClient) GetByBillingPeriodSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetByBillingPeriodResponder handles the response to the GetByBillingPeriod request. The method always
// closes the http.Response Body.
func (client PriceSheetClient) GetByBillingPeriodResponder(resp *http.Response) (result PriceSheetResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
