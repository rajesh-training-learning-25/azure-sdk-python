package consumption

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

// PriceSheetClient is the consumption management client provides access to consumption resources for Azure Enterprise
// Subscriptions.
type PriceSheetClient struct {
	BaseClient
}

// NewPriceSheetClient creates an instance of the PriceSheetClient client.
func NewPriceSheetClient(subscriptionID string) PriceSheetClient {
	return NewPriceSheetClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPriceSheetClientWithBaseURI creates an instance of the PriceSheetClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewPriceSheetClientWithBaseURI(baseURI string, subscriptionID string) PriceSheetClient {
	return PriceSheetClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets the price sheet for a scope by subscriptionId. Price sheet is available via this API only for May 1, 2014
// or later.
// Parameters:
// expand - may be used to expand the properties/meterDetails within a price sheet. By default, these fields
// are not included when returning price sheet.
// skiptoken - skiptoken is only used if a previous operation returned a partial result. If a previous response
// contains a nextLink element, the value of the nextLink element will include a skiptoken parameter that
// specifies a starting point to use for subsequent calls.
// top - may be used to limit the number of results to the top N results.
func (client PriceSheetClient) Get(ctx context.Context, expand string, skiptoken string, top *int32) (result PriceSheetResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PriceSheetClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
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
		return
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
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PriceSheetClient) GetResponder(resp *http.Response) (result PriceSheetResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetByBillingPeriod get the price sheet for a scope by subscriptionId and billing period. Price sheet is available
// via this API only for May 1, 2014 or later.
// Parameters:
// billingPeriodName - billing Period Name.
// expand - may be used to expand the properties/meterDetails within a price sheet. By default, these fields
// are not included when returning price sheet.
// skiptoken - skiptoken is only used if a previous operation returned a partial result. If a previous response
// contains a nextLink element, the value of the nextLink element will include a skiptoken parameter that
// specifies a starting point to use for subsequent calls.
// top - may be used to limit the number of results to the top N results.
func (client PriceSheetClient) GetByBillingPeriod(ctx context.Context, billingPeriodName string, expand string, skiptoken string, top *int32) (result PriceSheetResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PriceSheetClient.GetByBillingPeriod")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
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
		return
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
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetByBillingPeriodResponder handles the response to the GetByBillingPeriod request. The method always
// closes the http.Response Body.
func (client PriceSheetClient) GetByBillingPeriodResponder(resp *http.Response) (result PriceSheetResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
