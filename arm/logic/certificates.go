package logic

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

// CertificatesClient is the REST API for Azure Logic Apps.
type CertificatesClient struct {
	ManagementClient
}

// NewCertificatesClient creates an instance of the CertificatesClient client.
func NewCertificatesClient(subscriptionID string) CertificatesClient {
	return NewCertificatesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCertificatesClientWithBaseURI creates an instance of the CertificatesClient client.
func NewCertificatesClientWithBaseURI(baseURI string, subscriptionID string) CertificatesClient {
	return CertificatesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates an integration account certificate.
//
// resourceGroupName is the resource group name. integrationAccountName is the integration account name.
// certificateName is the integration account certificate name. certificate is the integration account certificate.
func (client CertificatesClient) CreateOrUpdate(resourceGroupName string, integrationAccountName string, certificateName string, certificate IntegrationAccountCertificate) (result IntegrationAccountCertificate, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: certificate,
			Constraints: []validation.Constraint{{Target: "certificate.IntegrationAccountCertificateProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "certificate.IntegrationAccountCertificateProperties.Key", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "certificate.IntegrationAccountCertificateProperties.Key.KeyVault", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "certificate.IntegrationAccountCertificateProperties.Key.KeyName", Name: validation.Null, Rule: true, Chain: nil},
					}},
				}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "logic.CertificatesClient", "CreateOrUpdate")
	}

	req, err := client.CreateOrUpdatePreparer(resourceGroupName, integrationAccountName, certificateName, certificate)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.CertificatesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.CertificatesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.CertificatesClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client CertificatesClient) CreateOrUpdatePreparer(resourceGroupName string, integrationAccountName string, certificateName string, certificate IntegrationAccountCertificate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"certificateName":        autorest.Encode("path", certificateName),
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/certificates/{certificateName}", pathParameters),
		autorest.WithJSON(certificate),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client CertificatesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client CertificatesClient) CreateOrUpdateResponder(resp *http.Response) (result IntegrationAccountCertificate, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an integration account certificate.
//
// resourceGroupName is the resource group name. integrationAccountName is the integration account name.
// certificateName is the integration account certificate name.
func (client CertificatesClient) Delete(resourceGroupName string, integrationAccountName string, certificateName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(resourceGroupName, integrationAccountName, certificateName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.CertificatesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "logic.CertificatesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.CertificatesClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client CertificatesClient) DeletePreparer(resourceGroupName string, integrationAccountName string, certificateName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"certificateName":        autorest.Encode("path", certificateName),
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/certificates/{certificateName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client CertificatesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client CertificatesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets an integration account certificate.
//
// resourceGroupName is the resource group name. integrationAccountName is the integration account name.
// certificateName is the integration account certificate name.
func (client CertificatesClient) Get(resourceGroupName string, integrationAccountName string, certificateName string) (result IntegrationAccountCertificate, err error) {
	req, err := client.GetPreparer(resourceGroupName, integrationAccountName, certificateName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.CertificatesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.CertificatesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.CertificatesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client CertificatesClient) GetPreparer(resourceGroupName string, integrationAccountName string, certificateName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"certificateName":        autorest.Encode("path", certificateName),
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/certificates/{certificateName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client CertificatesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client CertificatesClient) GetResponder(resp *http.Response) (result IntegrationAccountCertificate, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByIntegrationAccounts gets a list of integration account certificates.
//
// resourceGroupName is the resource group name. integrationAccountName is the integration account name. top is the
// number of items to be included in the result.
func (client CertificatesClient) ListByIntegrationAccounts(resourceGroupName string, integrationAccountName string, top *int32) (result IntegrationAccountCertificateListResult, err error) {
	req, err := client.ListByIntegrationAccountsPreparer(resourceGroupName, integrationAccountName, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.CertificatesClient", "ListByIntegrationAccounts", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByIntegrationAccountsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.CertificatesClient", "ListByIntegrationAccounts", resp, "Failure sending request")
		return
	}

	result, err = client.ListByIntegrationAccountsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.CertificatesClient", "ListByIntegrationAccounts", resp, "Failure responding to request")
	}

	return
}

// ListByIntegrationAccountsPreparer prepares the ListByIntegrationAccounts request.
func (client CertificatesClient) ListByIntegrationAccountsPreparer(resourceGroupName string, integrationAccountName string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/certificates", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByIntegrationAccountsSender sends the ListByIntegrationAccounts request. The method will close the
// http.Response Body if it receives an error.
func (client CertificatesClient) ListByIntegrationAccountsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByIntegrationAccountsResponder handles the response to the ListByIntegrationAccounts request. The method always
// closes the http.Response Body.
func (client CertificatesClient) ListByIntegrationAccountsResponder(resp *http.Response) (result IntegrationAccountCertificateListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByIntegrationAccountsNextResults retrieves the next set of results, if any.
func (client CertificatesClient) ListByIntegrationAccountsNextResults(lastResults IntegrationAccountCertificateListResult) (result IntegrationAccountCertificateListResult, err error) {
	req, err := lastResults.IntegrationAccountCertificateListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.CertificatesClient", "ListByIntegrationAccounts", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByIntegrationAccountsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic.CertificatesClient", "ListByIntegrationAccounts", resp, "Failure sending next results request")
	}

	result, err = client.ListByIntegrationAccountsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.CertificatesClient", "ListByIntegrationAccounts", resp, "Failure responding to next results request")
	}

	return
}

// ListByIntegrationAccountsComplete gets all elements from the list without paging.
func (client CertificatesClient) ListByIntegrationAccountsComplete(resourceGroupName string, integrationAccountName string, top *int32, cancel <-chan struct{}) (<-chan IntegrationAccountCertificate, <-chan error) {
	resultChan := make(chan IntegrationAccountCertificate)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListByIntegrationAccounts(resourceGroupName, integrationAccountName, top)
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
			list, err = client.ListByIntegrationAccountsNextResults(list)
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
