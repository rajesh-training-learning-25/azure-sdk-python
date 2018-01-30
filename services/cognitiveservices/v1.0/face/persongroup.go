package face

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
	"net/http"
)

// PersonGroupClient is the an API for face detection, verification, and identification.
type PersonGroupClient struct {
	BaseClient
}

// NewPersonGroupClient creates an instance of the PersonGroupClient client.
func NewPersonGroupClient(azureRegion AzureRegions) PersonGroupClient {
	return PersonGroupClient{New(azureRegion)}
}

// Create create a new person group with specified personGroupId, name and user-provided userData.
//
// personGroupID is user-provided personGroupId as a string. body is request body for creating new person group.
func (client PersonGroupClient) Create(ctx context.Context, personGroupID string, body CreatePersonGroupRequest) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: personGroupID,
			Constraints: []validation.Constraint{{Target: "personGroupID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "personGroupID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}},
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.Name", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "body.Name", Name: validation.MaxLength, Rule: 128, Chain: nil}}},
				{Target: "body.UserData", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "body.UserData", Name: validation.MaxLength, Rule: 16384, Chain: nil}}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "face.PersonGroupClient", "Create")
	}

	req, err := client.CreatePreparer(ctx, personGroupID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client PersonGroupClient) CreatePreparer(ctx context.Context, personGroupID string, body CreatePersonGroupRequest) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	pathParameters := map[string]interface{}{
		"personGroupId": autorest.Encode("path", personGroupID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPathParameters("/persongroups/{personGroupId}", pathParameters),
		autorest.WithJSON(body))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client PersonGroupClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client PersonGroupClient) CreateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete delete an existing person group. Persisted face images of all people in the person group will also be
// deleted.
//
// personGroupID is the personGroupId of the person group to be deleted.
func (client PersonGroupClient) Delete(ctx context.Context, personGroupID string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: personGroupID,
			Constraints: []validation.Constraint{{Target: "personGroupID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "personGroupID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "face.PersonGroupClient", "Delete")
	}

	req, err := client.DeletePreparer(ctx, personGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client PersonGroupClient) DeletePreparer(ctx context.Context, personGroupID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	pathParameters := map[string]interface{}{
		"personGroupId": autorest.Encode("path", personGroupID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPathParameters("/persongroups/{personGroupId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PersonGroupClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client PersonGroupClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieve the information of a person group, including its name and userData.
//
// personGroupID is personGroupId of the target person group.
func (client PersonGroupClient) Get(ctx context.Context, personGroupID string) (result PersonGroupResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: personGroupID,
			Constraints: []validation.Constraint{{Target: "personGroupID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "personGroupID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "face.PersonGroupClient", "Get")
	}

	req, err := client.GetPreparer(ctx, personGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client PersonGroupClient) GetPreparer(ctx context.Context, personGroupID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	pathParameters := map[string]interface{}{
		"personGroupId": autorest.Encode("path", personGroupID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPathParameters("/persongroups/{personGroupId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PersonGroupClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PersonGroupClient) GetResponder(resp *http.Response) (result PersonGroupResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetTrainingStatus retrieve the training status of a person group (completed or ongoing).
//
// personGroupID is personGroupId of target person group.
func (client PersonGroupClient) GetTrainingStatus(ctx context.Context, personGroupID string) (result TrainingStatus, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: personGroupID,
			Constraints: []validation.Constraint{{Target: "personGroupID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "personGroupID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "face.PersonGroupClient", "GetTrainingStatus")
	}

	req, err := client.GetTrainingStatusPreparer(ctx, personGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "GetTrainingStatus", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetTrainingStatusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "GetTrainingStatus", resp, "Failure sending request")
		return
	}

	result, err = client.GetTrainingStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "GetTrainingStatus", resp, "Failure responding to request")
	}

	return
}

// GetTrainingStatusPreparer prepares the GetTrainingStatus request.
func (client PersonGroupClient) GetTrainingStatusPreparer(ctx context.Context, personGroupID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	pathParameters := map[string]interface{}{
		"personGroupId": autorest.Encode("path", personGroupID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPathParameters("/persongroups/{personGroupId}/training", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetTrainingStatusSender sends the GetTrainingStatus request. The method will close the
// http.Response Body if it receives an error.
func (client PersonGroupClient) GetTrainingStatusSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetTrainingStatusResponder handles the response to the GetTrainingStatus request. The method always
// closes the http.Response Body.
func (client PersonGroupClient) GetTrainingStatusResponder(resp *http.Response) (result TrainingStatus, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list person groups and their information.
//
// start is list person groups from the least personGroupId greater than the "start". top is the number of person
// groups to list.
func (client PersonGroupClient) List(ctx context.Context, start string, top *int32) (result ListPersonGroupResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: start,
			Constraints: []validation.Constraint{{Target: "start", Name: validation.MaxLength, Rule: 64, Chain: nil}}},
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: 1000, Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "face.PersonGroupClient", "List")
	}

	req, err := client.ListPreparer(ctx, start, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client PersonGroupClient) ListPreparer(ctx context.Context, start string, top *int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	queryParameters := map[string]interface{}{}
	if len(start) > 0 {
		queryParameters["start"] = autorest.Encode("query", start)
	}
	if top != nil {
		queryParameters["top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPath("/persongroups"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client PersonGroupClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client PersonGroupClient) ListResponder(resp *http.Response) (result ListPersonGroupResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Train queue a person group training task, the training task may not be started immediately.
//
// personGroupID is target person group to be trained.
func (client PersonGroupClient) Train(ctx context.Context, personGroupID string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: personGroupID,
			Constraints: []validation.Constraint{{Target: "personGroupID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "personGroupID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "face.PersonGroupClient", "Train")
	}

	req, err := client.TrainPreparer(ctx, personGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "Train", nil, "Failure preparing request")
		return
	}

	resp, err := client.TrainSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "Train", resp, "Failure sending request")
		return
	}

	result, err = client.TrainResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "Train", resp, "Failure responding to request")
	}

	return
}

// TrainPreparer prepares the Train request.
func (client PersonGroupClient) TrainPreparer(ctx context.Context, personGroupID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	pathParameters := map[string]interface{}{
		"personGroupId": autorest.Encode("path", personGroupID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPathParameters("/persongroups/{personGroupId}/train", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// TrainSender sends the Train request. The method will close the
// http.Response Body if it receives an error.
func (client PersonGroupClient) TrainSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// TrainResponder handles the response to the Train request. The method always
// closes the http.Response Body.
func (client PersonGroupClient) TrainResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update update an existing person group's display name and userData. The properties which does not appear in request
// body will not be updated.
//
// personGroupID is personGroupId of the person group to be updated. body is request body for updating person group.
func (client PersonGroupClient) Update(ctx context.Context, personGroupID string, body CreatePersonGroupRequest) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: personGroupID,
			Constraints: []validation.Constraint{{Target: "personGroupID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "personGroupID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "face.PersonGroupClient", "Update")
	}

	req, err := client.UpdatePreparer(ctx, personGroupID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client PersonGroupClient) UpdatePreparer(ctx context.Context, personGroupID string, body CreatePersonGroupRequest) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	pathParameters := map[string]interface{}{
		"personGroupId": autorest.Encode("path", personGroupID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPathParameters("/persongroups/{personGroupId}", pathParameters),
		autorest.WithJSON(body))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client PersonGroupClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client PersonGroupClient) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
