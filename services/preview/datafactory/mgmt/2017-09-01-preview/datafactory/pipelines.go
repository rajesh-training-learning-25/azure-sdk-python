package datafactory

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

// PipelinesClient is the the Azure Data Factory V2 management API provides a RESTful set of web services that interact
// with Azure Data Factory V2 services.
type PipelinesClient struct {
	BaseClient
}

// NewPipelinesClient creates an instance of the PipelinesClient client.
func NewPipelinesClient(subscriptionID string) PipelinesClient {
	return NewPipelinesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPipelinesClientWithBaseURI creates an instance of the PipelinesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewPipelinesClientWithBaseURI(baseURI string, subscriptionID string) PipelinesClient {
	return PipelinesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a pipeline.
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
// pipelineName - the pipeline name.
// pipeline - pipeline resource definition.
// ifMatch - eTag of the pipeline entity.  Should only be specified for update, for which it should match
// existing entity or can be * for unconditional update.
func (client PipelinesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, pipeline PipelineResource, ifMatch string) (result PipelineResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PipelinesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: pipelineName,
			Constraints: []validation.Constraint{{Target: "pipelineName", Name: validation.MaxLength, Rule: 260, Chain: nil},
				{Target: "pipelineName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "pipelineName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil}}},
		{TargetValue: pipeline,
			Constraints: []validation.Constraint{{Target: "pipeline.Pipeline", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "pipeline.Pipeline.Concurrency", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "pipeline.Pipeline.Concurrency", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil}}},
				}}}}}); err != nil {
		return result, validation.NewError("datafactory.PipelinesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, factoryName, pipelineName, pipeline, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.PipelinesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datafactory.PipelinesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.PipelinesClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client PipelinesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, pipeline PipelineResource, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":       autorest.Encode("path", factoryName),
		"pipelineName":      autorest.Encode("path", pipelineName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelines/{pipelineName}", pathParameters),
		autorest.WithJSON(pipeline),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client PipelinesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client PipelinesClient) CreateOrUpdateResponder(resp *http.Response) (result PipelineResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateRun creates a run of a pipeline.
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
// pipelineName - the pipeline name.
// parameters - parameters of the pipeline run.
func (client PipelinesClient) CreateRun(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, parameters map[string]interface{}) (result CreateRunResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PipelinesClient.CreateRun")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: pipelineName,
			Constraints: []validation.Constraint{{Target: "pipelineName", Name: validation.MaxLength, Rule: 260, Chain: nil},
				{Target: "pipelineName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "pipelineName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.PipelinesClient", "CreateRun", err.Error())
	}

	req, err := client.CreateRunPreparer(ctx, resourceGroupName, factoryName, pipelineName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.PipelinesClient", "CreateRun", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateRunSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datafactory.PipelinesClient", "CreateRun", resp, "Failure sending request")
		return
	}

	result, err = client.CreateRunResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.PipelinesClient", "CreateRun", resp, "Failure responding to request")
		return
	}

	return
}

// CreateRunPreparer prepares the CreateRun request.
func (client PipelinesClient) CreateRunPreparer(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, parameters map[string]interface{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":       autorest.Encode("path", factoryName),
		"pipelineName":      autorest.Encode("path", pipelineName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelines/{pipelineName}/createRun", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if parameters != nil && len(parameters) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(parameters))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateRunSender sends the CreateRun request. The method will close the
// http.Response Body if it receives an error.
func (client PipelinesClient) CreateRunSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateRunResponder handles the response to the CreateRun request. The method always
// closes the http.Response Body.
func (client PipelinesClient) CreateRunResponder(resp *http.Response) (result CreateRunResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a pipeline.
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
// pipelineName - the pipeline name.
func (client PipelinesClient) Delete(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PipelinesClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: pipelineName,
			Constraints: []validation.Constraint{{Target: "pipelineName", Name: validation.MaxLength, Rule: 260, Chain: nil},
				{Target: "pipelineName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "pipelineName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.PipelinesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, factoryName, pipelineName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.PipelinesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "datafactory.PipelinesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.PipelinesClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client PipelinesClient) DeletePreparer(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":       autorest.Encode("path", factoryName),
		"pipelineName":      autorest.Encode("path", pipelineName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelines/{pipelineName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PipelinesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client PipelinesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a pipeline.
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
// pipelineName - the pipeline name.
func (client PipelinesClient) Get(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string) (result PipelineResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PipelinesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: pipelineName,
			Constraints: []validation.Constraint{{Target: "pipelineName", Name: validation.MaxLength, Rule: 260, Chain: nil},
				{Target: "pipelineName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "pipelineName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.PipelinesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, factoryName, pipelineName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.PipelinesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datafactory.PipelinesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.PipelinesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client PipelinesClient) GetPreparer(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":       autorest.Encode("path", factoryName),
		"pipelineName":      autorest.Encode("path", pipelineName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelines/{pipelineName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PipelinesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PipelinesClient) GetResponder(resp *http.Response) (result PipelineResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByFactory lists pipelines.
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
func (client PipelinesClient) ListByFactory(ctx context.Context, resourceGroupName string, factoryName string) (result PipelineListResponsePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PipelinesClient.ListByFactory")
		defer func() {
			sc := -1
			if result.plr.Response.Response != nil {
				sc = result.plr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.PipelinesClient", "ListByFactory", err.Error())
	}

	result.fn = client.listByFactoryNextResults
	req, err := client.ListByFactoryPreparer(ctx, resourceGroupName, factoryName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.PipelinesClient", "ListByFactory", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByFactorySender(req)
	if err != nil {
		result.plr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datafactory.PipelinesClient", "ListByFactory", resp, "Failure sending request")
		return
	}

	result.plr, err = client.ListByFactoryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.PipelinesClient", "ListByFactory", resp, "Failure responding to request")
		return
	}
	if result.plr.hasNextLink() && result.plr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByFactoryPreparer prepares the ListByFactory request.
func (client PipelinesClient) ListByFactoryPreparer(ctx context.Context, resourceGroupName string, factoryName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":       autorest.Encode("path", factoryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelines", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByFactorySender sends the ListByFactory request. The method will close the
// http.Response Body if it receives an error.
func (client PipelinesClient) ListByFactorySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByFactoryResponder handles the response to the ListByFactory request. The method always
// closes the http.Response Body.
func (client PipelinesClient) ListByFactoryResponder(resp *http.Response) (result PipelineListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByFactoryNextResults retrieves the next set of results, if any.
func (client PipelinesClient) listByFactoryNextResults(ctx context.Context, lastResults PipelineListResponse) (result PipelineListResponse, err error) {
	req, err := lastResults.pipelineListResponsePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "datafactory.PipelinesClient", "listByFactoryNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByFactorySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "datafactory.PipelinesClient", "listByFactoryNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByFactoryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.PipelinesClient", "listByFactoryNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListByFactoryComplete enumerates all values, automatically crossing page boundaries as required.
func (client PipelinesClient) ListByFactoryComplete(ctx context.Context, resourceGroupName string, factoryName string) (result PipelineListResponseIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PipelinesClient.ListByFactory")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByFactory(ctx, resourceGroupName, factoryName)
	return
}
