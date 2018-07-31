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
	"github.com/satori/go.uuid"
	"io"
	"net/http"
)

// PersonGroupPersonClient is the an API for face detection, verification, and identification.
type PersonGroupPersonClient struct {
	BaseClient
}

// NewPersonGroupPersonClient creates an instance of the PersonGroupPersonClient client.
func NewPersonGroupPersonClient(endpoint string) PersonGroupPersonClient {
	return PersonGroupPersonClient{New(endpoint)}
}

// AddPersonFaceFromStream add a representative face to a person for identification. The input face is specified as an
// image with a targetFace rectangle.
// Parameters:
// personGroupID - id referencing a particular person group.
// personID - id referencing a particular person.
// imageParameter - an image stream.
// userData - user-specified data about the face for any purpose. The maximum length is 1KB.
// targetFace - a face rectangle to specify the target face to be added to a person in the format of
// "targetFace=left,top,width,height". E.g. "targetFace=10,10,100,100". If there is more than one face in the
// image, targetFace is required to specify which face to add. No targetFace means there is only one face
// detected in the entire image.
func (client PersonGroupPersonClient) AddPersonFaceFromStream(ctx context.Context, personGroupID string, personID uuid.UUID, imageParameter io.ReadCloser, userData string, targetFace []int32) (result PersistedFace, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: personGroupID,
			Constraints: []validation.Constraint{{Target: "personGroupID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "personGroupID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}},
		{TargetValue: userData,
			Constraints: []validation.Constraint{{Target: "userData", Name: validation.MaxLength, Rule: 1024, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.PersonGroupPersonClient", "AddPersonFaceFromStream", err.Error())
	}

	req, err := client.AddPersonFaceFromStreamPreparer(ctx, personGroupID, personID, imageParameter, userData, targetFace)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupPersonClient", "AddPersonFaceFromStream", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddPersonFaceFromStreamSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.PersonGroupPersonClient", "AddPersonFaceFromStream", resp, "Failure sending request")
		return
	}

	result, err = client.AddPersonFaceFromStreamResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupPersonClient", "AddPersonFaceFromStream", resp, "Failure responding to request")
	}

	return
}

// AddPersonFaceFromStreamPreparer prepares the AddPersonFaceFromStream request.
func (client PersonGroupPersonClient) AddPersonFaceFromStreamPreparer(ctx context.Context, personGroupID string, personID uuid.UUID, imageParameter io.ReadCloser, userData string, targetFace []int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"personGroupId": autorest.Encode("path", personGroupID),
		"personId":      autorest.Encode("path", personID),
	}

	queryParameters := map[string]interface{}{}
	if len(userData) > 0 {
		queryParameters["userData"] = autorest.Encode("query", userData)
	}
	if targetFace != nil && len(targetFace) > 0 {
		queryParameters["targetFace"] = autorest.Encode("query", targetFace, ",")
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/octet-stream"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPathParameters("/persongroups/{personGroupId}/persons/{personId}/persistedFaces", pathParameters),
		autorest.WithFile(imageParameter),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddPersonFaceFromStreamSender sends the AddPersonFaceFromStream request. The method will close the
// http.Response Body if it receives an error.
func (client PersonGroupPersonClient) AddPersonFaceFromStreamSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AddPersonFaceFromStreamResponder handles the response to the AddPersonFaceFromStream request. The method always
// closes the http.Response Body.
func (client PersonGroupPersonClient) AddPersonFaceFromStreamResponder(resp *http.Response) (result PersistedFace, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// AddPersonFaceFromURL add a representative face to a person for identification. The input face is specified as an
// image with a targetFace rectangle.
// Parameters:
// personGroupID - id referencing a particular person group.
// personID - id referencing a particular person.
// imageURL - a JSON document with a URL pointing to the image that is to be analyzed.
// userData - user-specified data about the face for any purpose. The maximum length is 1KB.
// targetFace - a face rectangle to specify the target face to be added to a person in the format of
// "targetFace=left,top,width,height". E.g. "targetFace=10,10,100,100". If there is more than one face in the
// image, targetFace is required to specify which face to add. No targetFace means there is only one face
// detected in the entire image.
func (client PersonGroupPersonClient) AddPersonFaceFromURL(ctx context.Context, personGroupID string, personID uuid.UUID, imageURL ImageURL, userData string, targetFace []int32) (result PersistedFace, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: personGroupID,
			Constraints: []validation.Constraint{{Target: "personGroupID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "personGroupID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}},
		{TargetValue: userData,
			Constraints: []validation.Constraint{{Target: "userData", Name: validation.MaxLength, Rule: 1024, Chain: nil}}},
		{TargetValue: imageURL,
			Constraints: []validation.Constraint{{Target: "imageURL.URL", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.PersonGroupPersonClient", "AddPersonFaceFromURL", err.Error())
	}

	req, err := client.AddPersonFaceFromURLPreparer(ctx, personGroupID, personID, imageURL, userData, targetFace)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupPersonClient", "AddPersonFaceFromURL", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddPersonFaceFromURLSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.PersonGroupPersonClient", "AddPersonFaceFromURL", resp, "Failure sending request")
		return
	}

	result, err = client.AddPersonFaceFromURLResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupPersonClient", "AddPersonFaceFromURL", resp, "Failure responding to request")
	}

	return
}

// AddPersonFaceFromURLPreparer prepares the AddPersonFaceFromURL request.
func (client PersonGroupPersonClient) AddPersonFaceFromURLPreparer(ctx context.Context, personGroupID string, personID uuid.UUID, imageURL ImageURL, userData string, targetFace []int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"personGroupId": autorest.Encode("path", personGroupID),
		"personId":      autorest.Encode("path", personID),
	}

	queryParameters := map[string]interface{}{}
	if len(userData) > 0 {
		queryParameters["userData"] = autorest.Encode("query", userData)
	}
	if targetFace != nil && len(targetFace) > 0 {
		queryParameters["targetFace"] = autorest.Encode("query", targetFace, ",")
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPathParameters("/persongroups/{personGroupId}/persons/{personId}/persistedFaces", pathParameters),
		autorest.WithJSON(imageURL),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddPersonFaceFromURLSender sends the AddPersonFaceFromURL request. The method will close the
// http.Response Body if it receives an error.
func (client PersonGroupPersonClient) AddPersonFaceFromURLSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AddPersonFaceFromURLResponder handles the response to the AddPersonFaceFromURL request. The method always
// closes the http.Response Body.
func (client PersonGroupPersonClient) AddPersonFaceFromURLResponder(resp *http.Response) (result PersistedFace, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Create create a new person in a specified person group.
// Parameters:
// personGroupID - id referencing a particular person group.
// body - request body for creating new person.
func (client PersonGroupPersonClient) Create(ctx context.Context, personGroupID string, body NameAndUserDataContract) (result Person, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: personGroupID,
			Constraints: []validation.Constraint{{Target: "personGroupID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "personGroupID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}},
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.Name", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "body.Name", Name: validation.MaxLength, Rule: 128, Chain: nil}}},
				{Target: "body.UserData", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "body.UserData", Name: validation.MaxLength, Rule: 16384, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("face.PersonGroupPersonClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, personGroupID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupPersonClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.PersonGroupPersonClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupPersonClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client PersonGroupPersonClient) CreatePreparer(ctx context.Context, personGroupID string, body NameAndUserDataContract) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"personGroupId": autorest.Encode("path", personGroupID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPathParameters("/persongroups/{personGroupId}/persons", pathParameters),
		autorest.WithJSON(body))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client PersonGroupPersonClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client PersonGroupPersonClient) CreateResponder(resp *http.Response) (result Person, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete an existing person from a person group. All stored person data, and face features in the person entry
// will be deleted.
// Parameters:
// personGroupID - id referencing a particular person group.
// personID - id referencing a particular person.
func (client PersonGroupPersonClient) Delete(ctx context.Context, personGroupID string, personID uuid.UUID) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: personGroupID,
			Constraints: []validation.Constraint{{Target: "personGroupID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "personGroupID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.PersonGroupPersonClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, personGroupID, personID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupPersonClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "face.PersonGroupPersonClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupPersonClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client PersonGroupPersonClient) DeletePreparer(ctx context.Context, personGroupID string, personID uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"personGroupId": autorest.Encode("path", personGroupID),
		"personId":      autorest.Encode("path", personID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPathParameters("/persongroups/{personGroupId}/persons/{personId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PersonGroupPersonClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client PersonGroupPersonClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteFace delete a face from a person. Relative feature for the persisted face will also be deleted.
// Parameters:
// personGroupID - id referencing a particular person group.
// personID - id referencing a particular person.
// persistedFaceID - id referencing a particular persistedFaceId of an existing face.
func (client PersonGroupPersonClient) DeleteFace(ctx context.Context, personGroupID string, personID uuid.UUID, persistedFaceID uuid.UUID) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: personGroupID,
			Constraints: []validation.Constraint{{Target: "personGroupID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "personGroupID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.PersonGroupPersonClient", "DeleteFace", err.Error())
	}

	req, err := client.DeleteFacePreparer(ctx, personGroupID, personID, persistedFaceID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupPersonClient", "DeleteFace", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteFaceSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "face.PersonGroupPersonClient", "DeleteFace", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteFaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupPersonClient", "DeleteFace", resp, "Failure responding to request")
	}

	return
}

// DeleteFacePreparer prepares the DeleteFace request.
func (client PersonGroupPersonClient) DeleteFacePreparer(ctx context.Context, personGroupID string, personID uuid.UUID, persistedFaceID uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"persistedFaceId": autorest.Encode("path", persistedFaceID),
		"personGroupId":   autorest.Encode("path", personGroupID),
		"personId":        autorest.Encode("path", personID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPathParameters("/persongroups/{personGroupId}/persons/{personId}/persistedFaces/{persistedFaceId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteFaceSender sends the DeleteFace request. The method will close the
// http.Response Body if it receives an error.
func (client PersonGroupPersonClient) DeleteFaceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteFaceResponder handles the response to the DeleteFace request. The method always
// closes the http.Response Body.
func (client PersonGroupPersonClient) DeleteFaceResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieve a person's information, including registered persisted faces, name and userData.
// Parameters:
// personGroupID - id referencing a particular person group.
// personID - id referencing a particular person.
func (client PersonGroupPersonClient) Get(ctx context.Context, personGroupID string, personID uuid.UUID) (result Person, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: personGroupID,
			Constraints: []validation.Constraint{{Target: "personGroupID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "personGroupID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.PersonGroupPersonClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, personGroupID, personID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupPersonClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.PersonGroupPersonClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupPersonClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client PersonGroupPersonClient) GetPreparer(ctx context.Context, personGroupID string, personID uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"personGroupId": autorest.Encode("path", personGroupID),
		"personId":      autorest.Encode("path", personID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPathParameters("/persongroups/{personGroupId}/persons/{personId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PersonGroupPersonClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PersonGroupPersonClient) GetResponder(resp *http.Response) (result Person, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetFace retrieve information about a persisted face (specified by persistedFaceId, personId and its belonging
// personGroupId).
// Parameters:
// personGroupID - id referencing a particular person group.
// personID - id referencing a particular person.
// persistedFaceID - id referencing a particular persistedFaceId of an existing face.
func (client PersonGroupPersonClient) GetFace(ctx context.Context, personGroupID string, personID uuid.UUID, persistedFaceID uuid.UUID) (result PersistedFace, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: personGroupID,
			Constraints: []validation.Constraint{{Target: "personGroupID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "personGroupID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.PersonGroupPersonClient", "GetFace", err.Error())
	}

	req, err := client.GetFacePreparer(ctx, personGroupID, personID, persistedFaceID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupPersonClient", "GetFace", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetFaceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.PersonGroupPersonClient", "GetFace", resp, "Failure sending request")
		return
	}

	result, err = client.GetFaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupPersonClient", "GetFace", resp, "Failure responding to request")
	}

	return
}

// GetFacePreparer prepares the GetFace request.
func (client PersonGroupPersonClient) GetFacePreparer(ctx context.Context, personGroupID string, personID uuid.UUID, persistedFaceID uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"persistedFaceId": autorest.Encode("path", persistedFaceID),
		"personGroupId":   autorest.Encode("path", personGroupID),
		"personId":        autorest.Encode("path", personID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPathParameters("/persongroups/{personGroupId}/persons/{personId}/persistedFaces/{persistedFaceId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetFaceSender sends the GetFace request. The method will close the
// http.Response Body if it receives an error.
func (client PersonGroupPersonClient) GetFaceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetFaceResponder handles the response to the GetFace request. The method always
// closes the http.Response Body.
func (client PersonGroupPersonClient) GetFaceResponder(resp *http.Response) (result PersistedFace, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list all persons in a person group, and retrieve person information (including personId, name, userData and
// persistedFaceIds of registered faces of the person).
// Parameters:
// personGroupID - id referencing a particular person group.
// start - starting person id to return (used to list a range of persons).
// top - number of persons to return starting with the person id indicated by the 'start' parameter.
func (client PersonGroupPersonClient) List(ctx context.Context, personGroupID string, start string, top *int32) (result ListPerson, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: personGroupID,
			Constraints: []validation.Constraint{{Target: "personGroupID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "personGroupID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}},
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("face.PersonGroupPersonClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx, personGroupID, start, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupPersonClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.PersonGroupPersonClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupPersonClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client PersonGroupPersonClient) ListPreparer(ctx context.Context, personGroupID string, start string, top *int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"personGroupId": autorest.Encode("path", personGroupID),
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
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPathParameters("/persongroups/{personGroupId}/persons", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client PersonGroupPersonClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client PersonGroupPersonClient) ListResponder(resp *http.Response) (result ListPerson, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update update name or userData of a person.
// Parameters:
// personGroupID - id referencing a particular person group.
// personID - id referencing a particular person.
// body - request body for person update operation.
func (client PersonGroupPersonClient) Update(ctx context.Context, personGroupID string, personID uuid.UUID, body NameAndUserDataContract) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: personGroupID,
			Constraints: []validation.Constraint{{Target: "personGroupID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "personGroupID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.PersonGroupPersonClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, personGroupID, personID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupPersonClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "face.PersonGroupPersonClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupPersonClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client PersonGroupPersonClient) UpdatePreparer(ctx context.Context, personGroupID string, personID uuid.UUID, body NameAndUserDataContract) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"personGroupId": autorest.Encode("path", personGroupID),
		"personId":      autorest.Encode("path", personID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPathParameters("/persongroups/{personGroupId}/persons/{personId}", pathParameters),
		autorest.WithJSON(body))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client PersonGroupPersonClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client PersonGroupPersonClient) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// UpdateFace update a person persisted face's userData field.
// Parameters:
// personGroupID - id referencing a particular person group.
// personID - id referencing a particular person.
// persistedFaceID - id referencing a particular persistedFaceId of an existing face.
// body - request body for updating persisted face.
func (client PersonGroupPersonClient) UpdateFace(ctx context.Context, personGroupID string, personID uuid.UUID, persistedFaceID uuid.UUID, body UpdateFaceRequest) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: personGroupID,
			Constraints: []validation.Constraint{{Target: "personGroupID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "personGroupID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.PersonGroupPersonClient", "UpdateFace", err.Error())
	}

	req, err := client.UpdateFacePreparer(ctx, personGroupID, personID, persistedFaceID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupPersonClient", "UpdateFace", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateFaceSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "face.PersonGroupPersonClient", "UpdateFace", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateFaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.PersonGroupPersonClient", "UpdateFace", resp, "Failure responding to request")
	}

	return
}

// UpdateFacePreparer prepares the UpdateFace request.
func (client PersonGroupPersonClient) UpdateFacePreparer(ctx context.Context, personGroupID string, personID uuid.UUID, persistedFaceID uuid.UUID, body UpdateFaceRequest) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"persistedFaceId": autorest.Encode("path", persistedFaceID),
		"personGroupId":   autorest.Encode("path", personGroupID),
		"personId":        autorest.Encode("path", personID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPathParameters("/persongroups/{personGroupId}/persons/{personId}/persistedFaces/{persistedFaceId}", pathParameters),
		autorest.WithJSON(body))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateFaceSender sends the UpdateFace request. The method will close the
// http.Response Body if it receives an error.
func (client PersonGroupPersonClient) UpdateFaceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateFaceResponder handles the response to the UpdateFace request. The method always
// closes the http.Response Body.
func (client PersonGroupPersonClient) UpdateFaceResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
