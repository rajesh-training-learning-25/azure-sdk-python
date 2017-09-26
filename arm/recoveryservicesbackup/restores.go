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

// RestoresClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type RestoresClient struct {
	ManagementClient
}

// NewRestoresClient creates an instance of the RestoresClient client.
func NewRestoresClient(subscriptionID string) RestoresClient {
	return NewRestoresClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRestoresClientWithBaseURI creates an instance of the RestoresClient client.
func NewRestoresClientWithBaseURI(baseURI string, subscriptionID string) RestoresClient {
	return RestoresClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Trigger restores the specified backed up data. This is an asynchronous operation. To know the status of this API
// call, use GetProtectedItemOperationResult API.
//
// vaultName is the name of the recovery services vault. resourceGroupName is the name of the resource group where the
// recovery services vault is present. fabricName is fabric name associated with the backed up items. containerName is
// container name associated with the backed up items. protectedItemName is backed up item to be restored.
// recoveryPointID is recovery point ID which represents the backed up data to be restored. parameters is resource
// restore request
func (client RestoresClient) Trigger(vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string, parameters RestoreRequestResource) (result autorest.Response, err error) {
	req, err := client.TriggerPreparer(vaultName, resourceGroupName, fabricName, containerName, protectedItemName, recoveryPointID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.RestoresClient", "Trigger", nil, "Failure preparing request")
		return
	}

	resp, err := client.TriggerSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.RestoresClient", "Trigger", resp, "Failure sending request")
		return
	}

	result, err = client.TriggerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.RestoresClient", "Trigger", resp, "Failure responding to request")
	}

	return
}

// TriggerPreparer prepares the Trigger request.
func (client RestoresClient) TriggerPreparer(vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string, parameters RestoreRequestResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerName":     autorest.Encode("path", containerName),
		"fabricName":        autorest.Encode("path", fabricName),
		"protectedItemName": autorest.Encode("path", protectedItemName),
		"recoveryPointId":   autorest.Encode("path", recoveryPointID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}/recoveryPoints/{recoveryPointId}/restore", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// TriggerSender sends the Trigger request. The method will close the
// http.Response Body if it receives an error.
func (client RestoresClient) TriggerSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// TriggerResponder handles the response to the Trigger request. The method always
// closes the http.Response Body.
func (client RestoresClient) TriggerResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
