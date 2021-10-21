package videoanalyzer

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

// PipelineJobOperationStatusesClient is the azure Video Analyzer provides a platform for you to build intelligent
// video applications that span the edge and the cloud
type PipelineJobOperationStatusesClient struct {
	BaseClient
}

// NewPipelineJobOperationStatusesClient creates an instance of the PipelineJobOperationStatusesClient client.
func NewPipelineJobOperationStatusesClient(subscriptionID string) PipelineJobOperationStatusesClient {
	return NewPipelineJobOperationStatusesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPipelineJobOperationStatusesClientWithBaseURI creates an instance of the PipelineJobOperationStatusesClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewPipelineJobOperationStatusesClientWithBaseURI(baseURI string, subscriptionID string) PipelineJobOperationStatusesClient {
	return PipelineJobOperationStatusesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get the operation status of a pipeline job with the given operationId.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// accountName - the Azure Video Analyzer account name.
// pipelineJobName - the pipeline job name.
// operationID - the operation ID.
func (client PipelineJobOperationStatusesClient) Get(ctx context.Context, resourceGroupName string, accountName string, pipelineJobName string, operationID string) (result PipelineJobOperationStatus, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PipelineJobOperationStatusesClient.Get")
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
		return result, validation.NewError("videoanalyzer.PipelineJobOperationStatusesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, accountName, pipelineJobName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "videoanalyzer.PipelineJobOperationStatusesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "videoanalyzer.PipelineJobOperationStatusesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "videoanalyzer.PipelineJobOperationStatusesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client PipelineJobOperationStatusesClient) GetPreparer(ctx context.Context, resourceGroupName string, accountName string, pipelineJobName string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"operationId":       autorest.Encode("path", operationID),
		"pipelineJobName":   autorest.Encode("path", pipelineJobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/pipelineJobs/{pipelineJobName}/operationStatuses/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PipelineJobOperationStatusesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PipelineJobOperationStatusesClient) GetResponder(resp *http.Response) (result PipelineJobOperationStatus, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
