package recoveryservicesbackup

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
	"net/http"
)

// BackupProtectableItemsClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type BackupProtectableItemsClient struct {
	ManagementClient
}

// NewBackupProtectableItemsClient creates an instance of the BackupProtectableItemsClient client.
func NewBackupProtectableItemsClient(subscriptionID string) BackupProtectableItemsClient {
	return NewBackupProtectableItemsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewBackupProtectableItemsClientWithBaseURI creates an instance of the BackupProtectableItemsClient client.
func NewBackupProtectableItemsClientWithBaseURI(baseURI string, subscriptionID string) BackupProtectableItemsClient {
	return BackupProtectableItemsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List provides a pageable list of protectable objects within your subscription according to the query filter and the
// pagination parameters.
//
// vaultName is the name of the recovery services vault. resourceGroupName is the name of the resource group where the
// recovery services vault is present. filter is oData filter options. skipToken is skipToken Filter.
func (client BackupProtectableItemsClient) List(vaultName string, resourceGroupName string, filter string, skipToken string) (result WorkloadProtectableItemResourceList, err error) {
	req, err := client.ListPreparer(vaultName, resourceGroupName, filter, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.BackupProtectableItemsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.BackupProtectableItemsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.BackupProtectableItemsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client BackupProtectableItemsClient) ListPreparer(vaultName string, resourceGroupName string, filter string, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupProtectableItems", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client BackupProtectableItemsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client BackupProtectableItemsClient) ListResponder(resp *http.Response) (result WorkloadProtectableItemResourceList, err error) {
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
func (client BackupProtectableItemsClient) ListNextResults(lastResults WorkloadProtectableItemResourceList) (result WorkloadProtectableItemResourceList, err error) {
	req, err := lastResults.WorkloadProtectableItemResourceListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "recoveryservicesbackup.BackupProtectableItemsClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "recoveryservicesbackup.BackupProtectableItemsClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.BackupProtectableItemsClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client BackupProtectableItemsClient) ListComplete(vaultName string, resourceGroupName string, filter string, skipToken string, cancel <-chan struct{}) (<-chan WorkloadProtectableItemResource, <-chan error) {
	resultChan := make(chan WorkloadProtectableItemResource)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.List(vaultName, resourceGroupName, filter, skipToken)
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
