package costmanagement

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

// QueryClient is the client for the Query methods of the Costmanagement service.
type QueryClient struct {
	BaseClient
}

// NewQueryClient creates an instance of the QueryClient client.
func NewQueryClient(subscriptionID string) QueryClient {
	return NewQueryClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewQueryClientWithBaseURI creates an instance of the QueryClient client.
func NewQueryClientWithBaseURI(baseURI string, subscriptionID string) QueryClient {
	return QueryClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Usage query the usage data for scope defined.
// Parameters:
// scope - the scope associated with query and export operations. This includes
// '/subscriptions/{subscriptionId}/' for subscription scope,
// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope and
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department
// scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}'
// for EnrollmentAccount scope, '/providers/Microsoft.Management/managementGroups/{managementGroupId} for
// Management Group scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for
// billingProfile scope,
// 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}'
// for invoiceSection scope,
// 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for
// partners, 'providers/Microsoft.CostManagement/ExternalSubscriptions/{externalSubscriptionId}' for linked
// account and 'providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountId}' for
// consolidated account
// parameters - parameters supplied to the CreateOrUpdate Query Config operation.
func (client QueryClient) Usage(ctx context.Context, scope string, parameters QueryDefinition) (result QueryResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QueryClient.Usage")
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
			Constraints: []validation.Constraint{{Target: "parameters.Type", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.TimePeriod", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.TimePeriod.From", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.TimePeriod.To", Name: validation.Null, Rule: true, Chain: nil},
					}},
				{Target: "parameters.Dataset", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.Dataset.Grouping", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.Dataset.Grouping", Name: validation.MaxItems, Rule: 2, Chain: nil}}},
						{Target: "parameters.Dataset.Filter", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.And", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.And", Name: validation.MinItems, Rule: 2, Chain: nil}}},
								{Target: "parameters.Dataset.Filter.Or", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Or", Name: validation.MinItems, Rule: 2, Chain: nil}}},
								{Target: "parameters.Dataset.Filter.Not", Name: validation.Null, Rule: false, Chain: nil},
								{Target: "parameters.Dataset.Filter.Dimension", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Dimension.Name", Name: validation.Null, Rule: true, Chain: nil},
										{Target: "parameters.Dataset.Filter.Dimension.Operator", Name: validation.Null, Rule: true, Chain: nil},
										{Target: "parameters.Dataset.Filter.Dimension.Values", Name: validation.Null, Rule: true,
											Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Dimension.Values", Name: validation.MinItems, Rule: 1, Chain: nil}}},
									}},
								{Target: "parameters.Dataset.Filter.Tag", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Tag.Name", Name: validation.Null, Rule: true, Chain: nil},
										{Target: "parameters.Dataset.Filter.Tag.Operator", Name: validation.Null, Rule: true, Chain: nil},
										{Target: "parameters.Dataset.Filter.Tag.Values", Name: validation.Null, Rule: true,
											Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Tag.Values", Name: validation.MinItems, Rule: 1, Chain: nil}}},
									}},
							}},
					}}}}}); err != nil {
		return result, validation.NewError("costmanagement.QueryClient", "Usage", err.Error())
	}

	req, err := client.UsagePreparer(ctx, scope, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.QueryClient", "Usage", nil, "Failure preparing request")
		return
	}

	resp, err := client.UsageSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.QueryClient", "Usage", resp, "Failure sending request")
		return
	}

	result, err = client.UsageResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.QueryClient", "Usage", resp, "Failure responding to request")
	}

	return
}

// UsagePreparer prepares the Usage request.
func (client QueryClient) UsagePreparer(ctx context.Context, scope string, parameters QueryDefinition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scope": scope,
	}

	const APIVersion = "2019-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.CostManagement/query", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UsageSender sends the Usage request. The method will close the
// http.Response Body if it receives an error.
func (client QueryClient) UsageSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// UsageResponder handles the response to the Usage request. The method always
// closes the http.Response Body.
func (client QueryClient) UsageResponder(resp *http.Response) (result QueryResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
