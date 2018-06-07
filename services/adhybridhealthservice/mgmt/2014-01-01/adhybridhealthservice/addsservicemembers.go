package adhybridhealthservice

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
	"github.com/satori/go.uuid"
	"net/http"
)

// AddsServiceMembersClient is the REST APIs for Azure Active Drectory Connect Health
type AddsServiceMembersClient struct {
	BaseClient
}

// NewAddsServiceMembersClient creates an instance of the AddsServiceMembersClient client.
func NewAddsServiceMembersClient() AddsServiceMembersClient {
	return NewAddsServiceMembersClientWithBaseURI(DefaultBaseURI)
}

// NewAddsServiceMembersClientWithBaseURI creates an instance of the AddsServiceMembersClient client.
func NewAddsServiceMembersClientWithBaseURI(baseURI string) AddsServiceMembersClient {
	return AddsServiceMembersClient{NewWithBaseURI(baseURI)}
}

// Add onboards  a server, for a given Active Directory Domain Controller service, to Azure Active Directory Connect
// Health Service.
// Parameters:
// serviceName - the name of the service under which the server is to be onboarded.
// serviceMember - the server object.
func (client AddsServiceMembersClient) Add(ctx context.Context, serviceName string, serviceMember ServiceMember) (result ServiceMember, err error) {
	req, err := client.AddPreparer(ctx, serviceName, serviceMember)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "Add", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "Add", resp, "Failure sending request")
		return
	}

	result, err = client.AddResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "Add", resp, "Failure responding to request")
	}

	return
}

// AddPreparer prepares the Add request.
func (client AddsServiceMembersClient) AddPreparer(ctx context.Context, serviceName string, serviceMember ServiceMember) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"serviceName": autorest.Encode("path", serviceName),
	}

	const APIVersion = "2014-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/servicemembers", pathParameters),
		autorest.WithJSON(serviceMember),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddSender sends the Add request. The method will close the
// http.Response Body if it receives an error.
func (client AddsServiceMembersClient) AddSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AddResponder handles the response to the Add request. The method always
// closes the http.Response Body.
func (client AddsServiceMembersClient) AddResponder(resp *http.Response) (result ServiceMember, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a Active Directory Domain Controller server that has been onboarded to Azure Active Directory Connect
// Health Service.
// Parameters:
// serviceName - the name of the service.
// serviceMemberID - the server Id.
// confirm - indicates if the server will be permanently deleted or disabled. True indicates that the server
// will be permanently deleted and False indicates that the server will be marked disabled and then deleted
// after 30 days, if it is not re-registered.
func (client AddsServiceMembersClient) Delete(ctx context.Context, serviceName string, serviceMemberID uuid.UUID, confirm *bool) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(ctx, serviceName, serviceMemberID, confirm)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client AddsServiceMembersClient) DeletePreparer(ctx context.Context, serviceName string, serviceMemberID uuid.UUID, confirm *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"serviceMemberId": autorest.Encode("path", serviceMemberID),
		"serviceName":     autorest.Encode("path", serviceName),
	}

	const APIVersion = "2014-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if confirm != nil {
		queryParameters["confirm"] = autorest.Encode("query", *confirm)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/servicemembers/{serviceMemberId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client AddsServiceMembersClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client AddsServiceMembersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the details of a server, for a given Active Directory Domain Controller service, that are onboarded to
// Azure Active Directory Connect Health Service.
// Parameters:
// serviceName - the name of the service.
// serviceMemberID - the server Id.
func (client AddsServiceMembersClient) Get(ctx context.Context, serviceName string, serviceMemberID uuid.UUID) (result ServiceMember, err error) {
	req, err := client.GetPreparer(ctx, serviceName, serviceMemberID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client AddsServiceMembersClient) GetPreparer(ctx context.Context, serviceName string, serviceMemberID uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"serviceMemberId": autorest.Encode("path", serviceMemberID),
		"serviceName":     autorest.Encode("path", serviceName),
	}

	const APIVersion = "2014-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/servicemembers/{serviceMemberId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AddsServiceMembersClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AddsServiceMembersClient) GetResponder(resp *http.Response) (result ServiceMember, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets the details of the Active Directory Domain servers, for a given Active Directory Domain Service, that are
// onboarded to Azure Active Directory Connect Health.
// Parameters:
// serviceName - the name of the service.
// filter - the server property filter to apply.
func (client AddsServiceMembersClient) List(ctx context.Context, serviceName string, filter string) (result AddsServiceMembersPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, serviceName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.asm.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "List", resp, "Failure sending request")
		return
	}

	result.asm, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client AddsServiceMembersClient) ListPreparer(ctx context.Context, serviceName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"serviceName": autorest.Encode("path", serviceName),
	}

	const APIVersion = "2014-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/addsservicemembers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AddsServiceMembersClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AddsServiceMembersClient) ListResponder(resp *http.Response) (result AddsServiceMembers, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client AddsServiceMembersClient) listNextResults(lastResults AddsServiceMembers) (result AddsServiceMembers, err error) {
	req, err := lastResults.addsServiceMembersPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client AddsServiceMembersClient) ListComplete(ctx context.Context, serviceName string, filter string) (result AddsServiceMembersIterator, err error) {
	result.page, err = client.List(ctx, serviceName, filter)
	return
}

// List1 gets the details of the servers, for a given Active Directory Domain Controller service, that are onboarded to
// Azure Active Directory Connect Health Service.
// Parameters:
// serviceName - the name of the service.
// filter - the server property filter to apply.
// dimensionType - the server specific dimension.
// dimensionSignature - the value of the dimension.
func (client AddsServiceMembersClient) List1(ctx context.Context, serviceName string, filter string, dimensionType string, dimensionSignature string) (result ServiceMembersPage, err error) {
	result.fn = client.list1NextResults
	req, err := client.List1Preparer(ctx, serviceName, filter, dimensionType, dimensionSignature)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "List1", nil, "Failure preparing request")
		return
	}

	resp, err := client.List1Sender(req)
	if err != nil {
		result.sm.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "List1", resp, "Failure sending request")
		return
	}

	result.sm, err = client.List1Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "List1", resp, "Failure responding to request")
	}

	return
}

// List1Preparer prepares the List1 request.
func (client AddsServiceMembersClient) List1Preparer(ctx context.Context, serviceName string, filter string, dimensionType string, dimensionSignature string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"serviceName": autorest.Encode("path", serviceName),
	}

	const APIVersion = "2014-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(dimensionType) > 0 {
		queryParameters["dimensionType"] = autorest.Encode("query", dimensionType)
	}
	if len(dimensionSignature) > 0 {
		queryParameters["dimensionSignature"] = autorest.Encode("query", dimensionSignature)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/servicemembers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// List1Sender sends the List1 request. The method will close the
// http.Response Body if it receives an error.
func (client AddsServiceMembersClient) List1Sender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// List1Responder handles the response to the List1 request. The method always
// closes the http.Response Body.
func (client AddsServiceMembersClient) List1Responder(resp *http.Response) (result ServiceMembers, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// list1NextResults retrieves the next set of results, if any.
func (client AddsServiceMembersClient) list1NextResults(lastResults ServiceMembers) (result ServiceMembers, err error) {
	req, err := lastResults.serviceMembersPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "list1NextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.List1Sender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "list1NextResults", resp, "Failure sending next results request")
	}
	result, err = client.List1Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "list1NextResults", resp, "Failure responding to next results request")
	}
	return
}

// List1Complete enumerates all values, automatically crossing page boundaries as required.
func (client AddsServiceMembersClient) List1Complete(ctx context.Context, serviceName string, filter string, dimensionType string, dimensionSignature string) (result ServiceMembersIterator, err error) {
	result.page, err = client.List1(ctx, serviceName, filter, dimensionType, dimensionSignature)
	return
}

// ListCredentials gets the credentials of the server which is needed by the agent to connect to Azure Active Directory
// Connect Health Service.
// Parameters:
// serviceName - the name of the service.
// serviceMemberID - the server Id.
// filter - the property filter to apply.
func (client AddsServiceMembersClient) ListCredentials(ctx context.Context, serviceName string, serviceMemberID uuid.UUID, filter string) (result Credentials, err error) {
	req, err := client.ListCredentialsPreparer(ctx, serviceName, serviceMemberID, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "ListCredentials", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListCredentialsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "ListCredentials", resp, "Failure sending request")
		return
	}

	result, err = client.ListCredentialsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "ListCredentials", resp, "Failure responding to request")
	}

	return
}

// ListCredentialsPreparer prepares the ListCredentials request.
func (client AddsServiceMembersClient) ListCredentialsPreparer(ctx context.Context, serviceName string, serviceMemberID uuid.UUID, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"serviceMemberId": autorest.Encode("path", serviceMemberID),
		"serviceName":     autorest.Encode("path", serviceName),
	}

	const APIVersion = "2014-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/servicemembers/{serviceMemberId}/credentials", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListCredentialsSender sends the ListCredentials request. The method will close the
// http.Response Body if it receives an error.
func (client AddsServiceMembersClient) ListCredentialsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListCredentialsResponder handles the response to the ListCredentials request. The method always
// closes the http.Response Body.
func (client AddsServiceMembersClient) ListCredentialsResponder(resp *http.Response) (result Credentials, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
