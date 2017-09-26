package automation

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
// Code generated by Microsoft (R) AutoRest Code Generator 2.2.22.0
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"io"
	"net/http"
)

// RunbookDraftClient is the automation Client
type RunbookDraftClient struct {
	ManagementClient
}

// NewRunbookDraftClient creates an instance of the RunbookDraftClient client.
func NewRunbookDraftClient(subscriptionID string) RunbookDraftClient {
	return NewRunbookDraftClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRunbookDraftClientWithBaseURI creates an instance of the RunbookDraftClient client.
func NewRunbookDraftClientWithBaseURI(baseURI string, subscriptionID string) RunbookDraftClient {
	return RunbookDraftClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate updates the runbook draft with runbookStream as its content. This method may poll for completion.
// Polling can be canceled by passing the cancel channel argument. The channel will be used to cancel polling and any
// outstanding HTTP requests.
//
// resourceGroupName is the resource group name. automationAccountName is the automation account name. runbookName is
// the runbook name. runbookContent is the runbook draft content. runbookContent will be closed upon successful return.
// Callers should ensure closure when receiving an error.
func (client RunbookDraftClient) CreateOrUpdate(resourceGroupName string, automationAccountName string, runbookName string, runbookContent io.ReadCloser, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "automation.RunbookDraftClient", "CreateOrUpdate")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result autorest.Response
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreateOrUpdatePreparer(resourceGroupName, automationAccountName, runbookName, runbookContent, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "CreateOrUpdate", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateOrUpdateSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "CreateOrUpdate", resp, "Failure sending request")
			return
		}

		result, err = client.CreateOrUpdateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "CreateOrUpdate", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client RunbookDraftClient) CreateOrUpdatePreparer(resourceGroupName string, automationAccountName string, runbookName string, runbookContent io.ReadCloser, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"runbookName":           autorest.Encode("path", runbookName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/content", pathParameters),
		autorest.WithFile(runbookContent),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client RunbookDraftClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client RunbookDraftClient) CreateOrUpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusAccepted, http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieve the runbook draft identified by runbook name.
//
// resourceGroupName is the resource group name. automationAccountName is the automation account name. runbookName is
// the runbook name.
func (client RunbookDraftClient) Get(resourceGroupName string, automationAccountName string, runbookName string) (result RunbookDraft, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.RunbookDraftClient", "Get")
	}

	req, err := client.GetPreparer(resourceGroupName, automationAccountName, runbookName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RunbookDraftClient) GetPreparer(resourceGroupName string, automationAccountName string, runbookName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"runbookName":           autorest.Encode("path", runbookName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RunbookDraftClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RunbookDraftClient) GetResponder(resp *http.Response) (result RunbookDraft, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetContent retrieve the content of runbook draft identified by runbook name.
//
// resourceGroupName is the resource group name. automationAccountName is the automation account name. runbookName is
// the runbook name.
func (client RunbookDraftClient) GetContent(resourceGroupName string, automationAccountName string, runbookName string) (result ReadCloser, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.RunbookDraftClient", "GetContent")
	}

	req, err := client.GetContentPreparer(resourceGroupName, automationAccountName, runbookName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "GetContent", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetContentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "GetContent", resp, "Failure sending request")
		return
	}

	result, err = client.GetContentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "GetContent", resp, "Failure responding to request")
	}

	return
}

// GetContentPreparer prepares the GetContent request.
func (client RunbookDraftClient) GetContentPreparer(resourceGroupName string, automationAccountName string, runbookName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"runbookName":           autorest.Encode("path", runbookName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/content", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetContentSender sends the GetContent request. The method will close the
// http.Response Body if it receives an error.
func (client RunbookDraftClient) GetContentSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetContentResponder handles the response to the GetContent request. The method always
// closes the http.Response Body.
func (client RunbookDraftClient) GetContentResponder(resp *http.Response) (result ReadCloser, err error) {
	result.Value = &resp.Body
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK))
	result.Response = autorest.Response{Response: resp}
	return
}

// Publish publish runbook draft. This method may poll for completion. Polling can be canceled by passing the cancel
// channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the resource group name. automationAccountName is the automation account name. runbookName is
// the parameters supplied to the publish runbook operation.
func (client RunbookDraftClient) Publish(resourceGroupName string, automationAccountName string, runbookName string, cancel <-chan struct{}) (<-chan Runbook, <-chan error) {
	resultChan := make(chan Runbook, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "automation.RunbookDraftClient", "Publish")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result Runbook
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.PublishPreparer(resourceGroupName, automationAccountName, runbookName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "Publish", nil, "Failure preparing request")
			return
		}

		resp, err := client.PublishSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "Publish", resp, "Failure sending request")
			return
		}

		result, err = client.PublishResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "Publish", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// PublishPreparer prepares the Publish request.
func (client RunbookDraftClient) PublishPreparer(resourceGroupName string, automationAccountName string, runbookName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"runbookName":           autorest.Encode("path", runbookName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/publish", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// PublishSender sends the Publish request. The method will close the
// http.Response Body if it receives an error.
func (client RunbookDraftClient) PublishSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// PublishResponder handles the response to the Publish request. The method always
// closes the http.Response Body.
func (client RunbookDraftClient) PublishResponder(resp *http.Response) (result Runbook, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusAccepted, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UndoEdit retrieve the runbook identified by runbook name.
//
// resourceGroupName is the resource group name. automationAccountName is the automation account name. runbookName is
// the runbook name.
func (client RunbookDraftClient) UndoEdit(resourceGroupName string, automationAccountName string, runbookName string) (result RunbookDraftUndoEditResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.RunbookDraftClient", "UndoEdit")
	}

	req, err := client.UndoEditPreparer(resourceGroupName, automationAccountName, runbookName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "UndoEdit", nil, "Failure preparing request")
		return
	}

	resp, err := client.UndoEditSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "UndoEdit", resp, "Failure sending request")
		return
	}

	result, err = client.UndoEditResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "UndoEdit", resp, "Failure responding to request")
	}

	return
}

// UndoEditPreparer prepares the UndoEdit request.
func (client RunbookDraftClient) UndoEditPreparer(resourceGroupName string, automationAccountName string, runbookName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"runbookName":           autorest.Encode("path", runbookName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/undoEdit", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// UndoEditSender sends the UndoEdit request. The method will close the
// http.Response Body if it receives an error.
func (client RunbookDraftClient) UndoEditSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UndoEditResponder handles the response to the UndoEdit request. The method always
// closes the http.Response Body.
func (client RunbookDraftClient) UndoEditResponder(resp *http.Response) (result RunbookDraftUndoEditResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
