package recoveryservicessiterecovery

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

// ReplicationStorageClassificationsClient is the client for the ReplicationStorageClassifications methods of the
// Recoveryservicessiterecovery service.
type ReplicationStorageClassificationsClient struct {
	ManagementClient
}

// NewReplicationStorageClassificationsClient creates an instance of the ReplicationStorageClassificationsClient
// client.
func NewReplicationStorageClassificationsClient(subscriptionID string, resourceGroupName string, resourceName string) ReplicationStorageClassificationsClient {
	return NewReplicationStorageClassificationsClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName, resourceName)
}

// NewReplicationStorageClassificationsClientWithBaseURI creates an instance of the
// ReplicationStorageClassificationsClient client.
func NewReplicationStorageClassificationsClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string, resourceName string) ReplicationStorageClassificationsClient {
	return ReplicationStorageClassificationsClient{NewWithBaseURI(baseURI, subscriptionID, resourceGroupName, resourceName)}
}

// Get gets the details of the specified storage classification.
//
// fabricName is fabric name. storageClassificationName is storage classification name.
func (client ReplicationStorageClassificationsClient) Get(fabricName string, storageClassificationName string) (result StorageClassification, err error) {
	req, err := client.GetPreparer(fabricName, storageClassificationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationStorageClassificationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationStorageClassificationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationStorageClassificationsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ReplicationStorageClassificationsClient) GetPreparer(fabricName string, storageClassificationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":                autorest.Encode("path", fabricName),
		"resourceGroupName":         autorest.Encode("path", client.ResourceGroupName),
		"resourceName":              autorest.Encode("path", client.ResourceName),
		"storageClassificationName": autorest.Encode("path", storageClassificationName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications/{storageClassificationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationStorageClassificationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ReplicationStorageClassificationsClient) GetResponder(resp *http.Response) (result StorageClassification, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the storage classifications in the vault.
func (client ReplicationStorageClassificationsClient) List() (result StorageClassificationCollection, err error) {
	req, err := client.ListPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationStorageClassificationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationStorageClassificationsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationStorageClassificationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ReplicationStorageClassificationsClient) ListPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationStorageClassifications", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationStorageClassificationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ReplicationStorageClassificationsClient) ListResponder(resp *http.Response) (result StorageClassificationCollection, err error) {
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
func (client ReplicationStorageClassificationsClient) ListNextResults(lastResults StorageClassificationCollection) (result StorageClassificationCollection, err error) {
	req, err := lastResults.StorageClassificationCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationStorageClassificationsClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationStorageClassificationsClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationStorageClassificationsClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client ReplicationStorageClassificationsClient) ListComplete(cancel <-chan struct{}) (<-chan StorageClassification, <-chan error) {
	resultChan := make(chan StorageClassification)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.List()
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

// ListByReplicationFabrics lists the storage classifications available in the specified fabric.
//
// fabricName is site name of interest.
func (client ReplicationStorageClassificationsClient) ListByReplicationFabrics(fabricName string) (result StorageClassificationCollection, err error) {
	req, err := client.ListByReplicationFabricsPreparer(fabricName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationStorageClassificationsClient", "ListByReplicationFabrics", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByReplicationFabricsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationStorageClassificationsClient", "ListByReplicationFabrics", resp, "Failure sending request")
		return
	}

	result, err = client.ListByReplicationFabricsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationStorageClassificationsClient", "ListByReplicationFabrics", resp, "Failure responding to request")
	}

	return
}

// ListByReplicationFabricsPreparer prepares the ListByReplicationFabrics request.
func (client ReplicationStorageClassificationsClient) ListByReplicationFabricsPreparer(fabricName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":        autorest.Encode("path", fabricName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByReplicationFabricsSender sends the ListByReplicationFabrics request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationStorageClassificationsClient) ListByReplicationFabricsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByReplicationFabricsResponder handles the response to the ListByReplicationFabrics request. The method always
// closes the http.Response Body.
func (client ReplicationStorageClassificationsClient) ListByReplicationFabricsResponder(resp *http.Response) (result StorageClassificationCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByReplicationFabricsNextResults retrieves the next set of results, if any.
func (client ReplicationStorageClassificationsClient) ListByReplicationFabricsNextResults(lastResults StorageClassificationCollection) (result StorageClassificationCollection, err error) {
	req, err := lastResults.StorageClassificationCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationStorageClassificationsClient", "ListByReplicationFabrics", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByReplicationFabricsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationStorageClassificationsClient", "ListByReplicationFabrics", resp, "Failure sending next results request")
	}

	result, err = client.ListByReplicationFabricsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationStorageClassificationsClient", "ListByReplicationFabrics", resp, "Failure responding to next results request")
	}

	return
}

// ListByReplicationFabricsComplete gets all elements from the list without paging.
func (client ReplicationStorageClassificationsClient) ListByReplicationFabricsComplete(fabricName string, cancel <-chan struct{}) (<-chan StorageClassification, <-chan error) {
	resultChan := make(chan StorageClassification)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListByReplicationFabrics(fabricName)
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
			list, err = client.ListByReplicationFabricsNextResults(list)
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
