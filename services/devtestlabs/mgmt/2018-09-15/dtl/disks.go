package dtl

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// DisksClient is the the DevTest Labs Client.
type DisksClient struct {
	BaseClient
}

// NewDisksClient creates an instance of the DisksClient client.
func NewDisksClient(subscriptionID string) DisksClient {
	return NewDisksClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDisksClientWithBaseURI creates an instance of the DisksClient client.
func NewDisksClientWithBaseURI(baseURI string, subscriptionID string) DisksClient {
	return DisksClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Attach attach and create the lease of the disk to the virtual machine. This operation can take a while to complete.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// userName - the name of the user profile.
// name - the name of the disk.
// attachDiskProperties - properties of the disk to attach.
func (client DisksClient) Attach(ctx context.Context, resourceGroupName string, labName string, userName string, name string, attachDiskProperties AttachDiskProperties) (result DisksAttachFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DisksClient.Attach")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.AttachPreparer(ctx, resourceGroupName, labName, userName, name, attachDiskProperties)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.DisksClient", "Attach", nil, "Failure preparing request")
		return
	}

	result, err = client.AttachSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.DisksClient", "Attach", result.Response(), "Failure sending request")
		return
	}

	return
}

// AttachPreparer prepares the Attach request.
func (client DisksClient) AttachPreparer(ctx context.Context, resourceGroupName string, labName string, userName string, name string, attachDiskProperties AttachDiskProperties) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userName":          autorest.Encode("path", userName),
	}

	const APIVersion = "2018-09-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks/{name}/attach", pathParameters),
		autorest.WithJSON(attachDiskProperties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AttachSender sends the Attach request. The method will close the
// http.Response Body if it receives an error.
func (client DisksClient) AttachSender(req *http.Request) (future DisksAttachFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// AttachResponder handles the response to the Attach request. The method always
// closes the http.Response Body.
func (client DisksClient) AttachResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CreateOrUpdate create or replace an existing disk. This operation can take a while to complete.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// userName - the name of the user profile.
// name - the name of the disk.
// disk - a Disk.
func (client DisksClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, userName string, name string, disk Disk) (result DisksCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DisksClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: disk,
			Constraints: []validation.Constraint{{Target: "disk.DiskProperties", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("dtl.DisksClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, labName, userName, name, disk)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.DisksClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.DisksClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DisksClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, labName string, userName string, name string, disk Disk) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userName":          autorest.Encode("path", userName),
	}

	const APIVersion = "2018-09-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks/{name}", pathParameters),
		autorest.WithJSON(disk),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DisksClient) CreateOrUpdateSender(req *http.Request) (future DisksCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DisksClient) CreateOrUpdateResponder(resp *http.Response) (result Disk, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete disk. This operation can take a while to complete.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// userName - the name of the user profile.
// name - the name of the disk.
func (client DisksClient) Delete(ctx context.Context, resourceGroupName string, labName string, userName string, name string) (result DisksDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DisksClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, labName, userName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.DisksClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.DisksClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DisksClient) DeletePreparer(ctx context.Context, resourceGroupName string, labName string, userName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userName":          autorest.Encode("path", userName),
	}

	const APIVersion = "2018-09-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DisksClient) DeleteSender(req *http.Request) (future DisksDeleteFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DisksClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Detach detach and break the lease of the disk attached to the virtual machine. This operation can take a while to
// complete.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// userName - the name of the user profile.
// name - the name of the disk.
// detachDiskProperties - properties of the disk to detach.
func (client DisksClient) Detach(ctx context.Context, resourceGroupName string, labName string, userName string, name string, detachDiskProperties DetachDiskProperties) (result DisksDetachFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DisksClient.Detach")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DetachPreparer(ctx, resourceGroupName, labName, userName, name, detachDiskProperties)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.DisksClient", "Detach", nil, "Failure preparing request")
		return
	}

	result, err = client.DetachSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.DisksClient", "Detach", result.Response(), "Failure sending request")
		return
	}

	return
}

// DetachPreparer prepares the Detach request.
func (client DisksClient) DetachPreparer(ctx context.Context, resourceGroupName string, labName string, userName string, name string, detachDiskProperties DetachDiskProperties) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userName":          autorest.Encode("path", userName),
	}

	const APIVersion = "2018-09-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks/{name}/detach", pathParameters),
		autorest.WithJSON(detachDiskProperties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DetachSender sends the Detach request. The method will close the
// http.Response Body if it receives an error.
func (client DisksClient) DetachSender(req *http.Request) (future DisksDetachFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DetachResponder handles the response to the Detach request. The method always
// closes the http.Response Body.
func (client DisksClient) DetachResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get disk.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// userName - the name of the user profile.
// name - the name of the disk.
// expand - specify the $expand query. Example: 'properties($select=diskType)'
func (client DisksClient) Get(ctx context.Context, resourceGroupName string, labName string, userName string, name string, expand string) (result Disk, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DisksClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, labName, userName, name, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.DisksClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.DisksClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.DisksClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client DisksClient) GetPreparer(ctx context.Context, resourceGroupName string, labName string, userName string, name string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userName":          autorest.Encode("path", userName),
	}

	const APIVersion = "2018-09-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DisksClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DisksClient) GetResponder(resp *http.Response) (result Disk, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list disks in a given user profile.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// userName - the name of the user profile.
// expand - specify the $expand query. Example: 'properties($select=diskType)'
// filter - the filter to apply to the operation. Example: '$filter=contains(name,'myName')
// top - the maximum number of resources to return from the operation. Example: '$top=10'
// orderby - the ordering expression for the results, using OData notation. Example: '$orderby=name desc'
func (client DisksClient) List(ctx context.Context, resourceGroupName string, labName string, userName string, expand string, filter string, top *int32, orderby string) (result DiskListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DisksClient.List")
		defer func() {
			sc := -1
			if result.dl.Response.Response != nil {
				sc = result.dl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, labName, userName, expand, filter, top, orderby)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.DisksClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.dl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.DisksClient", "List", resp, "Failure sending request")
		return
	}

	result.dl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.DisksClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client DisksClient) ListPreparer(ctx context.Context, resourceGroupName string, labName string, userName string, expand string, filter string, top *int32, orderby string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userName":          autorest.Encode("path", userName),
	}

	const APIVersion = "2018-09-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DisksClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DisksClient) ListResponder(resp *http.Response) (result DiskList, err error) {
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
func (client DisksClient) listNextResults(ctx context.Context, lastResults DiskList) (result DiskList, err error) {
	req, err := lastResults.diskListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dtl.DisksClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dtl.DisksClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.DisksClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client DisksClient) ListComplete(ctx context.Context, resourceGroupName string, labName string, userName string, expand string, filter string, top *int32, orderby string) (result DiskListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DisksClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, labName, userName, expand, filter, top, orderby)
	return
}

// Update allows modifying tags of disks. All other properties will be ignored.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// userName - the name of the user profile.
// name - the name of the disk.
// disk - a Disk.
func (client DisksClient) Update(ctx context.Context, resourceGroupName string, labName string, userName string, name string, disk DiskFragment) (result Disk, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DisksClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, labName, userName, name, disk)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.DisksClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.DisksClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.DisksClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client DisksClient) UpdatePreparer(ctx context.Context, resourceGroupName string, labName string, userName string, name string, disk DiskFragment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userName":          autorest.Encode("path", userName),
	}

	const APIVersion = "2018-09-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks/{name}", pathParameters),
		autorest.WithJSON(disk),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client DisksClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client DisksClient) UpdateResponder(resp *http.Response) (result Disk, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
