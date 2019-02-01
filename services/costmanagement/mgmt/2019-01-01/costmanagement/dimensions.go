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

// DimensionsClient is the client for the Dimensions methods of the Costmanagement service.
type DimensionsClient struct {
	BaseClient
}

// NewDimensionsClient creates an instance of the DimensionsClient client.
func NewDimensionsClient(subscriptionID string) DimensionsClient {
	return NewDimensionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDimensionsClientWithBaseURI creates an instance of the DimensionsClient client.
func NewDimensionsClientWithBaseURI(baseURI string, subscriptionID string) DimensionsClient {
	return DimensionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByBillingAccount lists the dimensions by billingAccount Id.
// Parameters:
// billingAccountID - billingAccount ID
// filter - may be used to filter dimensions by properties/category, properties/usageStart,
// properties/usageEnd. Supported operators are 'eq','lt', 'gt', 'le', 'ge'.
// expand - may be used to expand the properties/data within a dimension category. By default, data is not
// included when listing dimensions.
// skiptoken - skiptoken is only used if a previous operation returned a partial result. If a previous response
// contains a nextLink element, the value of the nextLink element will include a skiptoken parameter that
// specifies a starting point to use for subsequent calls.
// top - may be used to limit the number of results to the most recent N dimension data.
func (client DimensionsClient) ListByBillingAccount(ctx context.Context, billingAccountID string, filter string, expand string, skiptoken string, top *int32) (result DimensionsListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DimensionsClient.ListByBillingAccount")
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
					{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("costmanagement.DimensionsClient", "ListByBillingAccount", err.Error())
	}

	req, err := client.ListByBillingAccountPreparer(ctx, billingAccountID, filter, expand, skiptoken, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.DimensionsClient", "ListByBillingAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBillingAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.DimensionsClient", "ListByBillingAccount", resp, "Failure sending request")
		return
	}

	result, err = client.ListByBillingAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.DimensionsClient", "ListByBillingAccount", resp, "Failure responding to request")
	}

	return
}

// ListByBillingAccountPreparer prepares the ListByBillingAccount request.
func (client DimensionsClient) ListByBillingAccountPreparer(ctx context.Context, billingAccountID string, filter string, expand string, skiptoken string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountId": autorest.Encode("path", billingAccountID),
	}

	const APIVersion = "2019-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
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
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/providers/Microsoft.CostManagement/dimensions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBillingAccountSender sends the ListByBillingAccount request. The method will close the
// http.Response Body if it receives an error.
func (client DimensionsClient) ListByBillingAccountSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByBillingAccountResponder handles the response to the ListByBillingAccount request. The method always
// closes the http.Response Body.
func (client DimensionsClient) ListByBillingAccountResponder(resp *http.Response) (result DimensionsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByDepartment lists the dimensions by Department Id.
// Parameters:
// billingAccountID - billingAccount ID
// departmentID - department ID
// filter - may be used to filter dimensions by properties/category, properties/usageStart,
// properties/usageEnd. Supported operators are 'eq','lt', 'gt', 'le', 'ge'.
// expand - may be used to expand the properties/data within a dimension category. By default, data is not
// included when listing dimensions.
// skiptoken - skiptoken is only used if a previous operation returned a partial result. If a previous response
// contains a nextLink element, the value of the nextLink element will include a skiptoken parameter that
// specifies a starting point to use for subsequent calls.
// top - may be used to limit the number of results to the most recent N dimension data.
func (client DimensionsClient) ListByDepartment(ctx context.Context, billingAccountID string, departmentID string, filter string, expand string, skiptoken string, top *int32) (result DimensionsListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DimensionsClient.ListByDepartment")
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
					{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("costmanagement.DimensionsClient", "ListByDepartment", err.Error())
	}

	req, err := client.ListByDepartmentPreparer(ctx, billingAccountID, departmentID, filter, expand, skiptoken, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.DimensionsClient", "ListByDepartment", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDepartmentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.DimensionsClient", "ListByDepartment", resp, "Failure sending request")
		return
	}

	result, err = client.ListByDepartmentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.DimensionsClient", "ListByDepartment", resp, "Failure responding to request")
	}

	return
}

// ListByDepartmentPreparer prepares the ListByDepartment request.
func (client DimensionsClient) ListByDepartmentPreparer(ctx context.Context, billingAccountID string, departmentID string, filter string, expand string, skiptoken string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountId": autorest.Encode("path", billingAccountID),
		"departmentId":     autorest.Encode("path", departmentID),
	}

	const APIVersion = "2019-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
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
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}/providers/Microsoft.CostManagement/dimensions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByDepartmentSender sends the ListByDepartment request. The method will close the
// http.Response Body if it receives an error.
func (client DimensionsClient) ListByDepartmentSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByDepartmentResponder handles the response to the ListByDepartment request. The method always
// closes the http.Response Body.
func (client DimensionsClient) ListByDepartmentResponder(resp *http.Response) (result DimensionsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByEnrollmentAccount lists the dimensions by Enrollment Account Id.
// Parameters:
// billingAccountID - billingAccount ID
// enrollmentAccountID - enrollment Account ID
// filter - may be used to filter dimensions by properties/category, properties/usageStart,
// properties/usageEnd. Supported operators are 'eq','lt', 'gt', 'le', 'ge'.
// expand - may be used to expand the properties/data within a dimension category. By default, data is not
// included when listing dimensions.
// skiptoken - skiptoken is only used if a previous operation returned a partial result. If a previous response
// contains a nextLink element, the value of the nextLink element will include a skiptoken parameter that
// specifies a starting point to use for subsequent calls.
// top - may be used to limit the number of results to the most recent N dimension data.
func (client DimensionsClient) ListByEnrollmentAccount(ctx context.Context, billingAccountID string, enrollmentAccountID string, filter string, expand string, skiptoken string, top *int32) (result DimensionsListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DimensionsClient.ListByEnrollmentAccount")
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
					{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("costmanagement.DimensionsClient", "ListByEnrollmentAccount", err.Error())
	}

	req, err := client.ListByEnrollmentAccountPreparer(ctx, billingAccountID, enrollmentAccountID, filter, expand, skiptoken, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.DimensionsClient", "ListByEnrollmentAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByEnrollmentAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.DimensionsClient", "ListByEnrollmentAccount", resp, "Failure sending request")
		return
	}

	result, err = client.ListByEnrollmentAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.DimensionsClient", "ListByEnrollmentAccount", resp, "Failure responding to request")
	}

	return
}

// ListByEnrollmentAccountPreparer prepares the ListByEnrollmentAccount request.
func (client DimensionsClient) ListByEnrollmentAccountPreparer(ctx context.Context, billingAccountID string, enrollmentAccountID string, filter string, expand string, skiptoken string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountId":    autorest.Encode("path", billingAccountID),
		"enrollmentAccountId": autorest.Encode("path", enrollmentAccountID),
	}

	const APIVersion = "2019-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
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
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}/providers/Microsoft.CostManagement/dimensions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByEnrollmentAccountSender sends the ListByEnrollmentAccount request. The method will close the
// http.Response Body if it receives an error.
func (client DimensionsClient) ListByEnrollmentAccountSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByEnrollmentAccountResponder handles the response to the ListByEnrollmentAccount request. The method always
// closes the http.Response Body.
func (client DimensionsClient) ListByEnrollmentAccountResponder(resp *http.Response) (result DimensionsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByManagementGroup lists the dimensions by managementGroup Id.
// Parameters:
// managementGroupID - managementGroup ID
// filter - may be used to filter dimensions by properties/category, properties/usageStart,
// properties/usageEnd. Supported operators are 'eq','lt', 'gt', 'le', 'ge'.
// expand - may be used to expand the properties/data within a dimension category. By default, data is not
// included when listing dimensions.
// skiptoken - skiptoken is only used if a previous operation returned a partial result. If a previous response
// contains a nextLink element, the value of the nextLink element will include a skiptoken parameter that
// specifies a starting point to use for subsequent calls.
// top - may be used to limit the number of results to the most recent N dimension data.
func (client DimensionsClient) ListByManagementGroup(ctx context.Context, managementGroupID string, filter string, expand string, skiptoken string, top *int32) (result DimensionsListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DimensionsClient.ListByManagementGroup")
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
					{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("costmanagement.DimensionsClient", "ListByManagementGroup", err.Error())
	}

	req, err := client.ListByManagementGroupPreparer(ctx, managementGroupID, filter, expand, skiptoken, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.DimensionsClient", "ListByManagementGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByManagementGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.DimensionsClient", "ListByManagementGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListByManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.DimensionsClient", "ListByManagementGroup", resp, "Failure responding to request")
	}

	return
}

// ListByManagementGroupPreparer prepares the ListByManagementGroup request.
func (client DimensionsClient) ListByManagementGroupPreparer(ctx context.Context, managementGroupID string, filter string, expand string, skiptoken string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupId": autorest.Encode("path", managementGroupID),
	}

	const APIVersion = "2019-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
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
		autorest.WithPathParameters("/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.CostManagement/dimensions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByManagementGroupSender sends the ListByManagementGroup request. The method will close the
// http.Response Body if it receives an error.
func (client DimensionsClient) ListByManagementGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByManagementGroupResponder handles the response to the ListByManagementGroup request. The method always
// closes the http.Response Body.
func (client DimensionsClient) ListByManagementGroupResponder(resp *http.Response) (result DimensionsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup lists the dimensions by resource group Id.
// Parameters:
// resourceGroupName - azure Resource Group Name.
// filter - may be used to filter dimensions by properties/category, properties/usageStart,
// properties/usageEnd. Supported operators are 'eq','lt', 'gt', 'le', 'ge'.
// expand - may be used to expand the properties/data within a dimension category. By default, data is not
// included when listing dimensions.
// skiptoken - skiptoken is only used if a previous operation returned a partial result. If a previous response
// contains a nextLink element, the value of the nextLink element will include a skiptoken parameter that
// specifies a starting point to use for subsequent calls.
// top - may be used to limit the number of results to the most recent N dimension data.
func (client DimensionsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, expand string, skiptoken string, top *int32) (result DimensionsListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DimensionsClient.ListByResourceGroup")
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
					{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("costmanagement.DimensionsClient", "ListByResourceGroup", err.Error())
	}

	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, filter, expand, skiptoken, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.DimensionsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.DimensionsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.DimensionsClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client DimensionsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, filter string, expand string, skiptoken string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CostManagement/dimensions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client DimensionsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client DimensionsClient) ListByResourceGroupResponder(resp *http.Response) (result DimensionsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBySubscription lists the dimensions by subscription Id.
// Parameters:
// filter - may be used to filter dimensions by properties/category, properties/usageStart,
// properties/usageEnd. Supported operators are 'eq','lt', 'gt', 'le', 'ge'.
// expand - may be used to expand the properties/data within a dimension category. By default, data is not
// included when listing dimensions.
// skiptoken - skiptoken is only used if a previous operation returned a partial result. If a previous response
// contains a nextLink element, the value of the nextLink element will include a skiptoken parameter that
// specifies a starting point to use for subsequent calls.
// top - may be used to limit the number of results to the most recent N dimension data.
func (client DimensionsClient) ListBySubscription(ctx context.Context, filter string, expand string, skiptoken string, top *int32) (result DimensionsListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DimensionsClient.ListBySubscription")
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
					{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("costmanagement.DimensionsClient", "ListBySubscription", err.Error())
	}

	req, err := client.ListBySubscriptionPreparer(ctx, filter, expand, skiptoken, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.DimensionsClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.DimensionsClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.DimensionsClient", "ListBySubscription", resp, "Failure responding to request")
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client DimensionsClient) ListBySubscriptionPreparer(ctx context.Context, filter string, expand string, skiptoken string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.CostManagement/dimensions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client DimensionsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client DimensionsClient) ListBySubscriptionResponder(resp *http.Response) (result DimensionsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
