package hybriddata

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

// JobDefinitionsClient is the client for the JobDefinitions methods of the Hybriddata service.
type JobDefinitionsClient struct {
	BaseClient
}

// NewJobDefinitionsClient creates an instance of the JobDefinitionsClient client.
func NewJobDefinitionsClient(subscriptionID string) JobDefinitionsClient {
	return NewJobDefinitionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewJobDefinitionsClientWithBaseURI creates an instance of the JobDefinitionsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewJobDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) JobDefinitionsClient {
	return JobDefinitionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a job definition.
// Parameters:
// dataServiceName - the data service type of the job definition.
// jobDefinitionName - the job definition name to be created or updated.
// jobDefinition - job Definition object to be created or updated.
// resourceGroupName - the Resource Group Name
// dataManagerName - the name of the DataManager Resource within the specified resource group. DataManager
// names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
func (client JobDefinitionsClient) CreateOrUpdate(ctx context.Context, dataServiceName string, jobDefinitionName string, jobDefinition JobDefinition, resourceGroupName string, dataManagerName string) (result JobDefinitionsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobDefinitionsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: jobDefinition,
			Constraints: []validation.Constraint{{Target: "jobDefinition.JobDefinitionProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "jobDefinition.JobDefinitionProperties.DataSourceID", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "jobDefinition.JobDefinitionProperties.DataSinkID", Name: validation.Null, Rule: true, Chain: nil},
				}}}},
		{TargetValue: dataManagerName,
			Constraints: []validation.Constraint{{Target: "dataManagerName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "dataManagerName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "dataManagerName", Name: validation.Pattern, Rule: `^[-\w\.]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hybriddata.JobDefinitionsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, dataServiceName, jobDefinitionName, jobDefinition, resourceGroupName, dataManagerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.JobDefinitionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.JobDefinitionsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client JobDefinitionsClient) CreateOrUpdatePreparer(ctx context.Context, dataServiceName string, jobDefinitionName string, jobDefinition JobDefinition, resourceGroupName string, dataManagerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dataManagerName":   autorest.Encode("path", dataManagerName),
		"dataServiceName":   autorest.Encode("path", dataServiceName),
		"jobDefinitionName": autorest.Encode("path", jobDefinitionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}", pathParameters),
		autorest.WithJSON(jobDefinition),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client JobDefinitionsClient) CreateOrUpdateSender(req *http.Request) (future JobDefinitionsCreateOrUpdateFuture, err error) {
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
func (client JobDefinitionsClient) CreateOrUpdateResponder(resp *http.Response) (result JobDefinition, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete this method deletes the given job definition.
// Parameters:
// dataServiceName - the data service type of the job definition.
// jobDefinitionName - the job definition name to be deleted.
// resourceGroupName - the Resource Group Name
// dataManagerName - the name of the DataManager Resource within the specified resource group. DataManager
// names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
func (client JobDefinitionsClient) Delete(ctx context.Context, dataServiceName string, jobDefinitionName string, resourceGroupName string, dataManagerName string) (result JobDefinitionsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobDefinitionsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: dataManagerName,
			Constraints: []validation.Constraint{{Target: "dataManagerName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "dataManagerName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "dataManagerName", Name: validation.Pattern, Rule: `^[-\w\.]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hybriddata.JobDefinitionsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, dataServiceName, jobDefinitionName, resourceGroupName, dataManagerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.JobDefinitionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.JobDefinitionsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client JobDefinitionsClient) DeletePreparer(ctx context.Context, dataServiceName string, jobDefinitionName string, resourceGroupName string, dataManagerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dataManagerName":   autorest.Encode("path", dataManagerName),
		"dataServiceName":   autorest.Encode("path", dataServiceName),
		"jobDefinitionName": autorest.Encode("path", jobDefinitionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client JobDefinitionsClient) DeleteSender(req *http.Request) (future JobDefinitionsDeleteFuture, err error) {
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
func (client JobDefinitionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get this method gets job definition object by name.
// Parameters:
// dataServiceName - the data service name of the job definition
// jobDefinitionName - the job definition name that is being queried.
// resourceGroupName - the Resource Group Name
// dataManagerName - the name of the DataManager Resource within the specified resource group. DataManager
// names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
func (client JobDefinitionsClient) Get(ctx context.Context, dataServiceName string, jobDefinitionName string, resourceGroupName string, dataManagerName string) (result JobDefinition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobDefinitionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: dataManagerName,
			Constraints: []validation.Constraint{{Target: "dataManagerName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "dataManagerName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "dataManagerName", Name: validation.Pattern, Rule: `^[-\w\.]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hybriddata.JobDefinitionsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, dataServiceName, jobDefinitionName, resourceGroupName, dataManagerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.JobDefinitionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hybriddata.JobDefinitionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.JobDefinitionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client JobDefinitionsClient) GetPreparer(ctx context.Context, dataServiceName string, jobDefinitionName string, resourceGroupName string, dataManagerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dataManagerName":   autorest.Encode("path", dataManagerName),
		"dataServiceName":   autorest.Encode("path", dataServiceName),
		"jobDefinitionName": autorest.Encode("path", jobDefinitionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client JobDefinitionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client JobDefinitionsClient) GetResponder(resp *http.Response) (result JobDefinition, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByDataManager this method gets all the job definitions of the given data manager resource.
// Parameters:
// resourceGroupName - the Resource Group Name
// dataManagerName - the name of the DataManager Resource within the specified resource group. DataManager
// names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
// filter - oData Filter options
func (client JobDefinitionsClient) ListByDataManager(ctx context.Context, resourceGroupName string, dataManagerName string, filter string) (result JobDefinitionListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobDefinitionsClient.ListByDataManager")
		defer func() {
			sc := -1
			if result.jdl.Response.Response != nil {
				sc = result.jdl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: dataManagerName,
			Constraints: []validation.Constraint{{Target: "dataManagerName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "dataManagerName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "dataManagerName", Name: validation.Pattern, Rule: `^[-\w\.]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hybriddata.JobDefinitionsClient", "ListByDataManager", err.Error())
	}

	result.fn = client.listByDataManagerNextResults
	req, err := client.ListByDataManagerPreparer(ctx, resourceGroupName, dataManagerName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.JobDefinitionsClient", "ListByDataManager", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDataManagerSender(req)
	if err != nil {
		result.jdl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hybriddata.JobDefinitionsClient", "ListByDataManager", resp, "Failure sending request")
		return
	}

	result.jdl, err = client.ListByDataManagerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.JobDefinitionsClient", "ListByDataManager", resp, "Failure responding to request")
		return
	}
	if result.jdl.hasNextLink() && result.jdl.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByDataManagerPreparer prepares the ListByDataManager request.
func (client JobDefinitionsClient) ListByDataManagerPreparer(ctx context.Context, resourceGroupName string, dataManagerName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dataManagerName":   autorest.Encode("path", dataManagerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/jobDefinitions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByDataManagerSender sends the ListByDataManager request. The method will close the
// http.Response Body if it receives an error.
func (client JobDefinitionsClient) ListByDataManagerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByDataManagerResponder handles the response to the ListByDataManager request. The method always
// closes the http.Response Body.
func (client JobDefinitionsClient) ListByDataManagerResponder(resp *http.Response) (result JobDefinitionList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByDataManagerNextResults retrieves the next set of results, if any.
func (client JobDefinitionsClient) listByDataManagerNextResults(ctx context.Context, lastResults JobDefinitionList) (result JobDefinitionList, err error) {
	req, err := lastResults.jobDefinitionListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "hybriddata.JobDefinitionsClient", "listByDataManagerNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByDataManagerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "hybriddata.JobDefinitionsClient", "listByDataManagerNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByDataManagerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.JobDefinitionsClient", "listByDataManagerNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListByDataManagerComplete enumerates all values, automatically crossing page boundaries as required.
func (client JobDefinitionsClient) ListByDataManagerComplete(ctx context.Context, resourceGroupName string, dataManagerName string, filter string) (result JobDefinitionListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobDefinitionsClient.ListByDataManager")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByDataManager(ctx, resourceGroupName, dataManagerName, filter)
	return
}

// ListByDataService this method gets all the job definitions of the given data service name.
// Parameters:
// dataServiceName - the data service type of interest.
// resourceGroupName - the Resource Group Name
// dataManagerName - the name of the DataManager Resource within the specified resource group. DataManager
// names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
// filter - oData Filter options
func (client JobDefinitionsClient) ListByDataService(ctx context.Context, dataServiceName string, resourceGroupName string, dataManagerName string, filter string) (result JobDefinitionListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobDefinitionsClient.ListByDataService")
		defer func() {
			sc := -1
			if result.jdl.Response.Response != nil {
				sc = result.jdl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: dataManagerName,
			Constraints: []validation.Constraint{{Target: "dataManagerName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "dataManagerName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "dataManagerName", Name: validation.Pattern, Rule: `^[-\w\.]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hybriddata.JobDefinitionsClient", "ListByDataService", err.Error())
	}

	result.fn = client.listByDataServiceNextResults
	req, err := client.ListByDataServicePreparer(ctx, dataServiceName, resourceGroupName, dataManagerName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.JobDefinitionsClient", "ListByDataService", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDataServiceSender(req)
	if err != nil {
		result.jdl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hybriddata.JobDefinitionsClient", "ListByDataService", resp, "Failure sending request")
		return
	}

	result.jdl, err = client.ListByDataServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.JobDefinitionsClient", "ListByDataService", resp, "Failure responding to request")
		return
	}
	if result.jdl.hasNextLink() && result.jdl.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByDataServicePreparer prepares the ListByDataService request.
func (client JobDefinitionsClient) ListByDataServicePreparer(ctx context.Context, dataServiceName string, resourceGroupName string, dataManagerName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dataManagerName":   autorest.Encode("path", dataManagerName),
		"dataServiceName":   autorest.Encode("path", dataServiceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByDataServiceSender sends the ListByDataService request. The method will close the
// http.Response Body if it receives an error.
func (client JobDefinitionsClient) ListByDataServiceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByDataServiceResponder handles the response to the ListByDataService request. The method always
// closes the http.Response Body.
func (client JobDefinitionsClient) ListByDataServiceResponder(resp *http.Response) (result JobDefinitionList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByDataServiceNextResults retrieves the next set of results, if any.
func (client JobDefinitionsClient) listByDataServiceNextResults(ctx context.Context, lastResults JobDefinitionList) (result JobDefinitionList, err error) {
	req, err := lastResults.jobDefinitionListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "hybriddata.JobDefinitionsClient", "listByDataServiceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByDataServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "hybriddata.JobDefinitionsClient", "listByDataServiceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByDataServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.JobDefinitionsClient", "listByDataServiceNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListByDataServiceComplete enumerates all values, automatically crossing page boundaries as required.
func (client JobDefinitionsClient) ListByDataServiceComplete(ctx context.Context, dataServiceName string, resourceGroupName string, dataManagerName string, filter string) (result JobDefinitionListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobDefinitionsClient.ListByDataService")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByDataService(ctx, dataServiceName, resourceGroupName, dataManagerName, filter)
	return
}

// Run this method runs a job instance of the given job definition.
// Parameters:
// dataServiceName - the data service type of the job definition.
// jobDefinitionName - name of the job definition.
// runParameters - run time parameters for the job definition.
// resourceGroupName - the Resource Group Name
// dataManagerName - the name of the DataManager Resource within the specified resource group. DataManager
// names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
func (client JobDefinitionsClient) Run(ctx context.Context, dataServiceName string, jobDefinitionName string, runParameters RunParameters, resourceGroupName string, dataManagerName string) (result JobDefinitionsRunFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobDefinitionsClient.Run")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: dataManagerName,
			Constraints: []validation.Constraint{{Target: "dataManagerName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "dataManagerName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "dataManagerName", Name: validation.Pattern, Rule: `^[-\w\.]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hybriddata.JobDefinitionsClient", "Run", err.Error())
	}

	req, err := client.RunPreparer(ctx, dataServiceName, jobDefinitionName, runParameters, resourceGroupName, dataManagerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.JobDefinitionsClient", "Run", nil, "Failure preparing request")
		return
	}

	result, err = client.RunSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.JobDefinitionsClient", "Run", result.Response(), "Failure sending request")
		return
	}

	return
}

// RunPreparer prepares the Run request.
func (client JobDefinitionsClient) RunPreparer(ctx context.Context, dataServiceName string, jobDefinitionName string, runParameters RunParameters, resourceGroupName string, dataManagerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dataManagerName":   autorest.Encode("path", dataManagerName),
		"dataServiceName":   autorest.Encode("path", dataServiceName),
		"jobDefinitionName": autorest.Encode("path", jobDefinitionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}/run", pathParameters),
		autorest.WithJSON(runParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RunSender sends the Run request. The method will close the
// http.Response Body if it receives an error.
func (client JobDefinitionsClient) RunSender(req *http.Request) (future JobDefinitionsRunFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// RunResponder handles the response to the Run request. The method always
// closes the http.Response Body.
func (client JobDefinitionsClient) RunResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
