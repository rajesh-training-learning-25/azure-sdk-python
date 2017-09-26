package powerbiembedded

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
	"net/http"
)

// WorkspaceCollectionsClient is the client to manage your Power BI Embedded workspace collections and retrieve
// workspaces.
type WorkspaceCollectionsClient struct {
	ManagementClient
}

// NewWorkspaceCollectionsClient creates an instance of the WorkspaceCollectionsClient client.
func NewWorkspaceCollectionsClient(subscriptionID string) WorkspaceCollectionsClient {
	return NewWorkspaceCollectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkspaceCollectionsClientWithBaseURI creates an instance of the WorkspaceCollectionsClient client.
func NewWorkspaceCollectionsClientWithBaseURI(baseURI string, subscriptionID string) WorkspaceCollectionsClient {
	return WorkspaceCollectionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CheckNameAvailability verify the specified Power BI Workspace Collection name is valid and not already in use.
//
// location is azure location body is check name availability request
func (client WorkspaceCollectionsClient) CheckNameAvailability(location string, body CheckNameRequest) (result CheckNameResponse, err error) {
	req, err := client.CheckNameAvailabilityPreparer(location, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "CheckNameAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckNameAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "CheckNameAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckNameAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "CheckNameAvailability", resp, "Failure responding to request")
	}

	return
}

// CheckNameAvailabilityPreparer prepares the CheckNameAvailability request.
func (client WorkspaceCollectionsClient) CheckNameAvailabilityPreparer(location string, body CheckNameRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-01-29"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.PowerBI/locations/{location}/checkNameAvailability", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CheckNameAvailabilitySender sends the CheckNameAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceCollectionsClient) CheckNameAvailabilitySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CheckNameAvailabilityResponder handles the response to the CheckNameAvailability request. The method always
// closes the http.Response Body.
func (client WorkspaceCollectionsClient) CheckNameAvailabilityResponder(resp *http.Response) (result CheckNameResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Create creates a new Power BI Workspace Collection with the specified properties. A Power BI Workspace Collection
// contains one or more workspaces, and can be used to provision keys that provide API access to those workspaces.
//
// resourceGroupName is azure resource group workspaceCollectionName is power BI Embedded Workspace Collection name
// body is create workspace collection request
func (client WorkspaceCollectionsClient) Create(resourceGroupName string, workspaceCollectionName string, body CreateWorkspaceCollectionRequest) (result WorkspaceCollection, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.Sku", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "body.Sku.Name", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "body.Sku.Tier", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "powerbiembedded.WorkspaceCollectionsClient", "Create")
	}

	req, err := client.CreatePreparer(resourceGroupName, workspaceCollectionName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client WorkspaceCollectionsClient) CreatePreparer(resourceGroupName string, workspaceCollectionName string, body CreateWorkspaceCollectionRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
		"workspaceCollectionName": autorest.Encode("path", workspaceCollectionName),
	}

	const APIVersion = "2016-01-29"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceCollectionsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client WorkspaceCollectionsClient) CreateResponder(resp *http.Response) (result WorkspaceCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a Power BI Workspace Collection. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is azure resource group workspaceCollectionName is power BI Embedded Workspace Collection name
func (client WorkspaceCollectionsClient) Delete(resourceGroupName string, workspaceCollectionName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
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
		req, err := client.DeletePreparer(resourceGroupName, workspaceCollectionName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "Delete", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeleteSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "Delete", resp, "Failure sending request")
			return
		}

		result, err = client.DeleteResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "Delete", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeletePreparer prepares the Delete request.
func (client WorkspaceCollectionsClient) DeletePreparer(resourceGroupName string, workspaceCollectionName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
		"workspaceCollectionName": autorest.Encode("path", workspaceCollectionName),
	}

	const APIVersion = "2016-01-29"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceCollectionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client WorkspaceCollectionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetAccessKeys retrieves the primary and secondary access keys for the specified Power BI Workspace Collection.
//
// resourceGroupName is azure resource group workspaceCollectionName is power BI Embedded Workspace Collection name
func (client WorkspaceCollectionsClient) GetAccessKeys(resourceGroupName string, workspaceCollectionName string) (result WorkspaceCollectionAccessKeys, err error) {
	req, err := client.GetAccessKeysPreparer(resourceGroupName, workspaceCollectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "GetAccessKeys", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAccessKeysSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "GetAccessKeys", resp, "Failure sending request")
		return
	}

	result, err = client.GetAccessKeysResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "GetAccessKeys", resp, "Failure responding to request")
	}

	return
}

// GetAccessKeysPreparer prepares the GetAccessKeys request.
func (client WorkspaceCollectionsClient) GetAccessKeysPreparer(resourceGroupName string, workspaceCollectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
		"workspaceCollectionName": autorest.Encode("path", workspaceCollectionName),
	}

	const APIVersion = "2016-01-29"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}/listKeys", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetAccessKeysSender sends the GetAccessKeys request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceCollectionsClient) GetAccessKeysSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetAccessKeysResponder handles the response to the GetAccessKeys request. The method always
// closes the http.Response Body.
func (client WorkspaceCollectionsClient) GetAccessKeysResponder(resp *http.Response) (result WorkspaceCollectionAccessKeys, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetByName retrieves an existing Power BI Workspace Collection.
//
// resourceGroupName is azure resource group workspaceCollectionName is power BI Embedded Workspace Collection name
func (client WorkspaceCollectionsClient) GetByName(resourceGroupName string, workspaceCollectionName string) (result WorkspaceCollection, err error) {
	req, err := client.GetByNamePreparer(resourceGroupName, workspaceCollectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "GetByName", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "GetByName", resp, "Failure sending request")
		return
	}

	result, err = client.GetByNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "GetByName", resp, "Failure responding to request")
	}

	return
}

// GetByNamePreparer prepares the GetByName request.
func (client WorkspaceCollectionsClient) GetByNamePreparer(resourceGroupName string, workspaceCollectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
		"workspaceCollectionName": autorest.Encode("path", workspaceCollectionName),
	}

	const APIVersion = "2016-01-29"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetByNameSender sends the GetByName request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceCollectionsClient) GetByNameSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetByNameResponder handles the response to the GetByName request. The method always
// closes the http.Response Body.
func (client WorkspaceCollectionsClient) GetByNameResponder(resp *http.Response) (result WorkspaceCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup retrieves all existing Power BI workspace collections in the specified resource group.
//
// resourceGroupName is azure resource group
func (client WorkspaceCollectionsClient) ListByResourceGroup(resourceGroupName string) (result WorkspaceCollectionList, err error) {
	req, err := client.ListByResourceGroupPreparer(resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client WorkspaceCollectionsClient) ListByResourceGroupPreparer(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-01-29"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceCollectionsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client WorkspaceCollectionsClient) ListByResourceGroupResponder(resp *http.Response) (result WorkspaceCollectionList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBySubscription retrieves all existing Power BI workspace collections in the specified subscription.
func (client WorkspaceCollectionsClient) ListBySubscription() (result WorkspaceCollectionList, err error) {
	req, err := client.ListBySubscriptionPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "ListBySubscription", resp, "Failure responding to request")
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client WorkspaceCollectionsClient) ListBySubscriptionPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-01-29"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.PowerBI/workspaceCollections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceCollectionsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client WorkspaceCollectionsClient) ListBySubscriptionResponder(resp *http.Response) (result WorkspaceCollectionList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Migrate migrates an existing Power BI Workspace Collection to a different resource group and/or subscription.
//
// resourceGroupName is azure resource group body is workspace migration request
func (client WorkspaceCollectionsClient) Migrate(resourceGroupName string, body MigrateWorkspaceCollectionRequest) (result autorest.Response, err error) {
	req, err := client.MigratePreparer(resourceGroupName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "Migrate", nil, "Failure preparing request")
		return
	}

	resp, err := client.MigrateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "Migrate", resp, "Failure sending request")
		return
	}

	result, err = client.MigrateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "Migrate", resp, "Failure responding to request")
	}

	return
}

// MigratePreparer prepares the Migrate request.
func (client WorkspaceCollectionsClient) MigratePreparer(resourceGroupName string, body MigrateWorkspaceCollectionRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-01-29"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/moveResources", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// MigrateSender sends the Migrate request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceCollectionsClient) MigrateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// MigrateResponder handles the response to the Migrate request. The method always
// closes the http.Response Body.
func (client WorkspaceCollectionsClient) MigrateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// RegenerateKey regenerates the primary or secondary access key for the specified Power BI Workspace Collection.
//
// resourceGroupName is azure resource group workspaceCollectionName is power BI Embedded Workspace Collection name
// body is access key to regenerate
func (client WorkspaceCollectionsClient) RegenerateKey(resourceGroupName string, workspaceCollectionName string, body WorkspaceCollectionAccessKey) (result WorkspaceCollectionAccessKeys, err error) {
	req, err := client.RegenerateKeyPreparer(resourceGroupName, workspaceCollectionName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "RegenerateKey", nil, "Failure preparing request")
		return
	}

	resp, err := client.RegenerateKeySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "RegenerateKey", resp, "Failure sending request")
		return
	}

	result, err = client.RegenerateKeyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "RegenerateKey", resp, "Failure responding to request")
	}

	return
}

// RegenerateKeyPreparer prepares the RegenerateKey request.
func (client WorkspaceCollectionsClient) RegenerateKeyPreparer(resourceGroupName string, workspaceCollectionName string, body WorkspaceCollectionAccessKey) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
		"workspaceCollectionName": autorest.Encode("path", workspaceCollectionName),
	}

	const APIVersion = "2016-01-29"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}/regenerateKey", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// RegenerateKeySender sends the RegenerateKey request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceCollectionsClient) RegenerateKeySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// RegenerateKeyResponder handles the response to the RegenerateKey request. The method always
// closes the http.Response Body.
func (client WorkspaceCollectionsClient) RegenerateKeyResponder(resp *http.Response) (result WorkspaceCollectionAccessKeys, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update update an existing Power BI Workspace Collection with the specified properties.
//
// resourceGroupName is azure resource group workspaceCollectionName is power BI Embedded Workspace Collection name
// body is update workspace collection request
func (client WorkspaceCollectionsClient) Update(resourceGroupName string, workspaceCollectionName string, body UpdateWorkspaceCollectionRequest) (result WorkspaceCollection, err error) {
	req, err := client.UpdatePreparer(resourceGroupName, workspaceCollectionName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client WorkspaceCollectionsClient) UpdatePreparer(resourceGroupName string, workspaceCollectionName string, body UpdateWorkspaceCollectionRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
		"workspaceCollectionName": autorest.Encode("path", workspaceCollectionName),
	}

	const APIVersion = "2016-01-29"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceCollectionsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client WorkspaceCollectionsClient) UpdateResponder(resp *http.Response) (result WorkspaceCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
