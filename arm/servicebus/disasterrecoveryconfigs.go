package servicebus

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

// DisasterRecoveryConfigsClient is the azure Service Bus client
type DisasterRecoveryConfigsClient struct {
	ManagementClient
}

// NewDisasterRecoveryConfigsClient creates an instance of the DisasterRecoveryConfigsClient client.
func NewDisasterRecoveryConfigsClient(subscriptionID string) DisasterRecoveryConfigsClient {
	return NewDisasterRecoveryConfigsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDisasterRecoveryConfigsClientWithBaseURI creates an instance of the DisasterRecoveryConfigsClient client.
func NewDisasterRecoveryConfigsClientWithBaseURI(baseURI string, subscriptionID string) DisasterRecoveryConfigsClient {
	return DisasterRecoveryConfigsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// BreakPairing this operation disables the Disaster Recovery and stops replicating changes from primary to secondary
// namespaces
//
// resourceGroupName is name of the Resource group within the Azure subscription. namespaceName is the namespace name
// alias is the Disaster Recovery configuration name
func (client DisasterRecoveryConfigsClient) BreakPairing(resourceGroupName string, namespaceName string, alias string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}},
		{TargetValue: alias,
			Constraints: []validation.Constraint{{Target: "alias", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "alias", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "servicebus.DisasterRecoveryConfigsClient", "BreakPairing")
	}

	req, err := client.BreakPairingPreparer(resourceGroupName, namespaceName, alias)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.DisasterRecoveryConfigsClient", "BreakPairing", nil, "Failure preparing request")
		return
	}

	resp, err := client.BreakPairingSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "servicebus.DisasterRecoveryConfigsClient", "BreakPairing", resp, "Failure sending request")
		return
	}

	result, err = client.BreakPairingResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.DisasterRecoveryConfigsClient", "BreakPairing", resp, "Failure responding to request")
	}

	return
}

// BreakPairingPreparer prepares the BreakPairing request.
func (client DisasterRecoveryConfigsClient) BreakPairingPreparer(resourceGroupName string, namespaceName string, alias string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"alias":             autorest.Encode("path", alias),
		"namespaceName":     autorest.Encode("path", namespaceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}/breakPairing", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// BreakPairingSender sends the BreakPairing request. The method will close the
// http.Response Body if it receives an error.
func (client DisasterRecoveryConfigsClient) BreakPairingSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// BreakPairingResponder handles the response to the BreakPairing request. The method always
// closes the http.Response Body.
func (client DisasterRecoveryConfigsClient) BreakPairingResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CreateOrUpdate creates or updates a new Alias(Disaster Recovery configuration)
//
// resourceGroupName is name of the Resource group within the Azure subscription. namespaceName is the namespace name
// alias is the Disaster Recovery configuration name parameters is parameters required to create an Alias(Disaster
// Recovery configuration)
func (client DisasterRecoveryConfigsClient) CreateOrUpdate(resourceGroupName string, namespaceName string, alias string, parameters ArmDisasterRecovery) (result ArmDisasterRecovery, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}},
		{TargetValue: alias,
			Constraints: []validation.Constraint{{Target: "alias", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "alias", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "servicebus.DisasterRecoveryConfigsClient", "CreateOrUpdate")
	}

	req, err := client.CreateOrUpdatePreparer(resourceGroupName, namespaceName, alias, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.DisasterRecoveryConfigsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicebus.DisasterRecoveryConfigsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.DisasterRecoveryConfigsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DisasterRecoveryConfigsClient) CreateOrUpdatePreparer(resourceGroupName string, namespaceName string, alias string, parameters ArmDisasterRecovery) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"alias":             autorest.Encode("path", alias),
		"namespaceName":     autorest.Encode("path", namespaceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DisasterRecoveryConfigsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DisasterRecoveryConfigsClient) CreateOrUpdateResponder(resp *http.Response) (result ArmDisasterRecovery, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an Alias(Disaster Recovery configuration)
//
// resourceGroupName is name of the Resource group within the Azure subscription. namespaceName is the namespace name
// alias is the Disaster Recovery configuration name
func (client DisasterRecoveryConfigsClient) Delete(resourceGroupName string, namespaceName string, alias string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}},
		{TargetValue: alias,
			Constraints: []validation.Constraint{{Target: "alias", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "alias", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "servicebus.DisasterRecoveryConfigsClient", "Delete")
	}

	req, err := client.DeletePreparer(resourceGroupName, namespaceName, alias)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.DisasterRecoveryConfigsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "servicebus.DisasterRecoveryConfigsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.DisasterRecoveryConfigsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DisasterRecoveryConfigsClient) DeletePreparer(resourceGroupName string, namespaceName string, alias string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"alias":             autorest.Encode("path", alias),
		"namespaceName":     autorest.Encode("path", namespaceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DisasterRecoveryConfigsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DisasterRecoveryConfigsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// FailOver envokes GEO DR failover and reconfigure the alias to point to the secondary namespace
//
// resourceGroupName is name of the Resource group within the Azure subscription. namespaceName is the namespace name
// alias is the Disaster Recovery configuration name
func (client DisasterRecoveryConfigsClient) FailOver(resourceGroupName string, namespaceName string, alias string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}},
		{TargetValue: alias,
			Constraints: []validation.Constraint{{Target: "alias", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "alias", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "servicebus.DisasterRecoveryConfigsClient", "FailOver")
	}

	req, err := client.FailOverPreparer(resourceGroupName, namespaceName, alias)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.DisasterRecoveryConfigsClient", "FailOver", nil, "Failure preparing request")
		return
	}

	resp, err := client.FailOverSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "servicebus.DisasterRecoveryConfigsClient", "FailOver", resp, "Failure sending request")
		return
	}

	result, err = client.FailOverResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.DisasterRecoveryConfigsClient", "FailOver", resp, "Failure responding to request")
	}

	return
}

// FailOverPreparer prepares the FailOver request.
func (client DisasterRecoveryConfigsClient) FailOverPreparer(resourceGroupName string, namespaceName string, alias string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"alias":             autorest.Encode("path", alias),
		"namespaceName":     autorest.Encode("path", namespaceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}/failover", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// FailOverSender sends the FailOver request. The method will close the
// http.Response Body if it receives an error.
func (client DisasterRecoveryConfigsClient) FailOverSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// FailOverResponder handles the response to the FailOver request. The method always
// closes the http.Response Body.
func (client DisasterRecoveryConfigsClient) FailOverResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieves Alias(Disaster Recovery configuration) for primary or secondary namespace
//
// resourceGroupName is name of the Resource group within the Azure subscription. namespaceName is the namespace name
// alias is the Disaster Recovery configuration name
func (client DisasterRecoveryConfigsClient) Get(resourceGroupName string, namespaceName string, alias string) (result ArmDisasterRecovery, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}},
		{TargetValue: alias,
			Constraints: []validation.Constraint{{Target: "alias", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "alias", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "servicebus.DisasterRecoveryConfigsClient", "Get")
	}

	req, err := client.GetPreparer(resourceGroupName, namespaceName, alias)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.DisasterRecoveryConfigsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicebus.DisasterRecoveryConfigsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.DisasterRecoveryConfigsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client DisasterRecoveryConfigsClient) GetPreparer(resourceGroupName string, namespaceName string, alias string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"alias":             autorest.Encode("path", alias),
		"namespaceName":     autorest.Encode("path", namespaceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DisasterRecoveryConfigsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DisasterRecoveryConfigsClient) GetResponder(resp *http.Response) (result ArmDisasterRecovery, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all Alias(Disaster Recovery configurations)
//
// resourceGroupName is name of the Resource group within the Azure subscription. namespaceName is the namespace name
func (client DisasterRecoveryConfigsClient) List(resourceGroupName string, namespaceName string) (result ArmDisasterRecoveryListResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "servicebus.DisasterRecoveryConfigsClient", "List")
	}

	req, err := client.ListPreparer(resourceGroupName, namespaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.DisasterRecoveryConfigsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicebus.DisasterRecoveryConfigsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.DisasterRecoveryConfigsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client DisasterRecoveryConfigsClient) ListPreparer(resourceGroupName string, namespaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"namespaceName":     autorest.Encode("path", namespaceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DisasterRecoveryConfigsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DisasterRecoveryConfigsClient) ListResponder(resp *http.Response) (result ArmDisasterRecoveryListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client DisasterRecoveryConfigsClient) ListNextResults(lastResults ArmDisasterRecoveryListResult) (result ArmDisasterRecoveryListResult, err error) {
	req, err := lastResults.ArmDisasterRecoveryListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "servicebus.DisasterRecoveryConfigsClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "servicebus.DisasterRecoveryConfigsClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.DisasterRecoveryConfigsClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client DisasterRecoveryConfigsClient) ListComplete(resourceGroupName string, namespaceName string, cancel <-chan struct{}) (<-chan ArmDisasterRecovery, <-chan error) {
	resultChan := make(chan ArmDisasterRecovery)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.List(resourceGroupName, namespaceName)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}
