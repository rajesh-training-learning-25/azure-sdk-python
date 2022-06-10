// Deprecated: Please note, this package has been deprecated. A replacement package is available github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices. We strongly encourage you to upgrade to continue receiving updates. See Migration Guide for guidance on upgrading. Refer to our deprecation policy for more details.
//
// Package recoveryservices implements the Azure ARM Recoveryservices service API version 2021-08-01.
//
// Recovery Services Client
package recoveryservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

const (
	// DefaultBaseURI is the default URI used for the service Recoveryservices
	DefaultBaseURI = "https://management.azure.com"
)

// BaseClient is the base client for Recoveryservices.
type BaseClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the BaseClient client.
func New(subscriptionID string) BaseClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return BaseClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}

// GetOperationResult gets the operation result for a resource.
// Parameters:
// resourceGroupName - the name of the resource group where the recovery services vault is present.
// vaultName - the name of the recovery services vault.
func (client BaseClient) GetOperationResult(ctx context.Context, resourceGroupName string, vaultName string, operationID string) (result Vault, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GetOperationResult")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetOperationResultPreparer(ctx, resourceGroupName, vaultName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservices.BaseClient", "GetOperationResult", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetOperationResultSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "recoveryservices.BaseClient", "GetOperationResult", resp, "Failure sending request")
		return
	}

	result, err = client.GetOperationResultResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservices.BaseClient", "GetOperationResult", resp, "Failure responding to request")
		return
	}

	return
}

// GetOperationResultPreparer prepares the GetOperationResult request.
func (client BaseClient) GetOperationResultPreparer(ctx context.Context, resourceGroupName string, vaultName string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationId":       autorest.Encode("path", operationID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2021-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/operationResults/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetOperationResultSender sends the GetOperationResult request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GetOperationResultSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetOperationResultResponder handles the response to the GetOperationResult request. The method always
// closes the http.Response Body.
func (client BaseClient) GetOperationResultResponder(resp *http.Response) (result Vault, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetOperationStatus gets the operation status for a resource.
// Parameters:
// resourceGroupName - the name of the resource group where the recovery services vault is present.
// vaultName - the name of the recovery services vault.
func (client BaseClient) GetOperationStatus(ctx context.Context, resourceGroupName string, vaultName string, operationID string) (result OperationResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GetOperationStatus")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetOperationStatusPreparer(ctx, resourceGroupName, vaultName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservices.BaseClient", "GetOperationStatus", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetOperationStatusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "recoveryservices.BaseClient", "GetOperationStatus", resp, "Failure sending request")
		return
	}

	result, err = client.GetOperationStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservices.BaseClient", "GetOperationStatus", resp, "Failure responding to request")
		return
	}

	return
}

// GetOperationStatusPreparer prepares the GetOperationStatus request.
func (client BaseClient) GetOperationStatusPreparer(ctx context.Context, resourceGroupName string, vaultName string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationId":       autorest.Encode("path", operationID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2021-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/operationStatus/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetOperationStatusSender sends the GetOperationStatus request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GetOperationStatusSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetOperationStatusResponder handles the response to the GetOperationStatus request. The method always
// closes the http.Response Body.
func (client BaseClient) GetOperationStatusResponder(resp *http.Response) (result OperationResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
