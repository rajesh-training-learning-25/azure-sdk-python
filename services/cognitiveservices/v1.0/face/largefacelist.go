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

// LargeFaceListClient is the an API for face detection, verification, and identification.
type LargeFaceListClient struct {
	BaseClient
}

// NewLargeFaceListClient creates an instance of the LargeFaceListClient client.
func NewLargeFaceListClient(endpoint string) LargeFaceListClient {
	return LargeFaceListClient{New(endpoint)}
}

// AddFaceFromStream add a face to a large face list. The input face is specified as an image with a targetFace
// rectangle. It returns a persistedFaceId representing the added face, and persistedFaceId will not expire.
// Parameters:
// largeFaceListID - id referencing a particular large face list.
// imageParameter - an image stream.
// userData - user-specified data about the face for any purpose. The maximum length is 1KB.
// targetFace - a face rectangle to specify the target face to be added to a person in the format of
// "targetFace=left,top,width,height". E.g. "targetFace=10,10,100,100". If there is more than one face in the
// image, targetFace is required to specify which face to add. No targetFace means there is only one face
// detected in the entire image.
func (client LargeFaceListClient) AddFaceFromStream(ctx context.Context, largeFaceListID string, imageParameter io.ReadCloser, userData string, targetFace []int32) (result PersistedFace, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: largeFaceListID,
			Constraints: []validation.Constraint{{Target: "largeFaceListID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "largeFaceListID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}},
		{TargetValue: userData,
			Constraints: []validation.Constraint{{Target: "userData", Name: validation.MaxLength, Rule: 1024, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.LargeFaceListClient", "AddFaceFromStream", err.Error())
	}

	req, err := client.AddFaceFromStreamPreparer(ctx, largeFaceListID, imageParameter, userData, targetFace)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "AddFaceFromStream", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddFaceFromStreamSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "AddFaceFromStream", resp, "Failure sending request")
		return
	}

	result, err = client.AddFaceFromStreamResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "AddFaceFromStream", resp, "Failure responding to request")
	}

	return
}

// AddFaceFromStreamPreparer prepares the AddFaceFromStream request.
func (client LargeFaceListClient) AddFaceFromStreamPreparer(ctx context.Context, largeFaceListID string, imageParameter io.ReadCloser, userData string, targetFace []int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"largeFaceListId": autorest.Encode("path", largeFaceListID),
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
		autorest.WithPathParameters("/largefacelists/{largeFaceListId}/persistedFaces", pathParameters),
		autorest.WithFile(imageParameter),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddFaceFromStreamSender sends the AddFaceFromStream request. The method will close the
// http.Response Body if it receives an error.
func (client LargeFaceListClient) AddFaceFromStreamSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AddFaceFromStreamResponder handles the response to the AddFaceFromStream request. The method always
// closes the http.Response Body.
func (client LargeFaceListClient) AddFaceFromStreamResponder(resp *http.Response) (result PersistedFace, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// AddFaceFromURL add a face to a large face list. The input face is specified as an image with a targetFace rectangle.
// It returns a persistedFaceId representing the added face, and persistedFaceId will not expire.
// Parameters:
// largeFaceListID - id referencing a particular large face list.
// imageURL - a JSON document with a URL pointing to the image that is to be analyzed.
// userData - user-specified data about the face for any purpose. The maximum length is 1KB.
// targetFace - a face rectangle to specify the target face to be added to a person in the format of
// "targetFace=left,top,width,height". E.g. "targetFace=10,10,100,100". If there is more than one face in the
// image, targetFace is required to specify which face to add. No targetFace means there is only one face
// detected in the entire image.
func (client LargeFaceListClient) AddFaceFromURL(ctx context.Context, largeFaceListID string, imageURL ImageURL, userData string, targetFace []int32) (result PersistedFace, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: largeFaceListID,
			Constraints: []validation.Constraint{{Target: "largeFaceListID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "largeFaceListID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}},
		{TargetValue: userData,
			Constraints: []validation.Constraint{{Target: "userData", Name: validation.MaxLength, Rule: 1024, Chain: nil}}},
		{TargetValue: imageURL,
			Constraints: []validation.Constraint{{Target: "imageURL.URL", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.LargeFaceListClient", "AddFaceFromURL", err.Error())
	}

	req, err := client.AddFaceFromURLPreparer(ctx, largeFaceListID, imageURL, userData, targetFace)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "AddFaceFromURL", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddFaceFromURLSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "AddFaceFromURL", resp, "Failure sending request")
		return
	}

	result, err = client.AddFaceFromURLResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "AddFaceFromURL", resp, "Failure responding to request")
	}

	return
}

// AddFaceFromURLPreparer prepares the AddFaceFromURL request.
func (client LargeFaceListClient) AddFaceFromURLPreparer(ctx context.Context, largeFaceListID string, imageURL ImageURL, userData string, targetFace []int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"largeFaceListId": autorest.Encode("path", largeFaceListID),
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
		autorest.WithPathParameters("/largefacelists/{largeFaceListId}/persistedFaces", pathParameters),
		autorest.WithJSON(imageURL),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddFaceFromURLSender sends the AddFaceFromURL request. The method will close the
// http.Response Body if it receives an error.
func (client LargeFaceListClient) AddFaceFromURLSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AddFaceFromURLResponder handles the response to the AddFaceFromURL request. The method always
// closes the http.Response Body.
func (client LargeFaceListClient) AddFaceFromURLResponder(resp *http.Response) (result PersistedFace, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Create create an empty large face list. Up to 64 large face lists are allowed to exist in one subscription.
// Parameters:
// largeFaceListID - id referencing a particular large face list.
// body - request body for creating a large face list.
func (client LargeFaceListClient) Create(ctx context.Context, largeFaceListID string, body NameAndUserDataContract) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: largeFaceListID,
			Constraints: []validation.Constraint{{Target: "largeFaceListID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "largeFaceListID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}},
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.Name", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "body.Name", Name: validation.MaxLength, Rule: 128, Chain: nil}}},
				{Target: "body.UserData", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "body.UserData", Name: validation.MaxLength, Rule: 16384, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("face.LargeFaceListClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, largeFaceListID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client LargeFaceListClient) CreatePreparer(ctx context.Context, largeFaceListID string, body NameAndUserDataContract) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"largeFaceListId": autorest.Encode("path", largeFaceListID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPathParameters("/largefacelists/{largeFaceListId}", pathParameters),
		autorest.WithJSON(body))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client LargeFaceListClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client LargeFaceListClient) CreateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete delete an existing large face list according to faceListId. Persisted face images in the large face list will
// also be deleted.
// Parameters:
// largeFaceListID - id referencing a particular large face list.
func (client LargeFaceListClient) Delete(ctx context.Context, largeFaceListID string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: largeFaceListID,
			Constraints: []validation.Constraint{{Target: "largeFaceListID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "largeFaceListID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.LargeFaceListClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, largeFaceListID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client LargeFaceListClient) DeletePreparer(ctx context.Context, largeFaceListID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"largeFaceListId": autorest.Encode("path", largeFaceListID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPathParameters("/largefacelists/{largeFaceListId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client LargeFaceListClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client LargeFaceListClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteFace delete an existing face from a large face list (given by a persisitedFaceId and a largeFaceListId).
// Persisted image related to the face will also be deleted.
// Parameters:
// largeFaceListID - id referencing a particular large face list.
// persistedFaceID - id referencing a particular persistedFaceId of an existing face.
func (client LargeFaceListClient) DeleteFace(ctx context.Context, largeFaceListID string, persistedFaceID uuid.UUID) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: largeFaceListID,
			Constraints: []validation.Constraint{{Target: "largeFaceListID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "largeFaceListID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.LargeFaceListClient", "DeleteFace", err.Error())
	}

	req, err := client.DeleteFacePreparer(ctx, largeFaceListID, persistedFaceID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "DeleteFace", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteFaceSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "DeleteFace", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteFaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "DeleteFace", resp, "Failure responding to request")
	}

	return
}

// DeleteFacePreparer prepares the DeleteFace request.
func (client LargeFaceListClient) DeleteFacePreparer(ctx context.Context, largeFaceListID string, persistedFaceID uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"largeFaceListId": autorest.Encode("path", largeFaceListID),
		"persistedFaceId": autorest.Encode("path", persistedFaceID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPathParameters("/largefacelists/{largeFaceListId}/persistedFaces/{persistedFaceId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteFaceSender sends the DeleteFace request. The method will close the
// http.Response Body if it receives an error.
func (client LargeFaceListClient) DeleteFaceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteFaceResponder handles the response to the DeleteFace request. The method always
// closes the http.Response Body.
func (client LargeFaceListClient) DeleteFaceResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieve a large face list's information.
// Parameters:
// largeFaceListID - id referencing a particular large face list.
func (client LargeFaceListClient) Get(ctx context.Context, largeFaceListID string) (result LargeFaceList, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: largeFaceListID,
			Constraints: []validation.Constraint{{Target: "largeFaceListID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "largeFaceListID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.LargeFaceListClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, largeFaceListID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client LargeFaceListClient) GetPreparer(ctx context.Context, largeFaceListID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"largeFaceListId": autorest.Encode("path", largeFaceListID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPathParameters("/largefacelists/{largeFaceListId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client LargeFaceListClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client LargeFaceListClient) GetResponder(resp *http.Response) (result LargeFaceList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetFace retrieve information about a persisted face (specified by persistedFaceId and its belonging
// largeFaceListId).
// Parameters:
// largeFaceListID - id referencing a particular large face list.
// persistedFaceID - id referencing a particular persistedFaceId of an existing face.
func (client LargeFaceListClient) GetFace(ctx context.Context, largeFaceListID string, persistedFaceID uuid.UUID) (result PersistedFace, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: largeFaceListID,
			Constraints: []validation.Constraint{{Target: "largeFaceListID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "largeFaceListID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.LargeFaceListClient", "GetFace", err.Error())
	}

	req, err := client.GetFacePreparer(ctx, largeFaceListID, persistedFaceID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "GetFace", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetFaceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "GetFace", resp, "Failure sending request")
		return
	}

	result, err = client.GetFaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "GetFace", resp, "Failure responding to request")
	}

	return
}

// GetFacePreparer prepares the GetFace request.
func (client LargeFaceListClient) GetFacePreparer(ctx context.Context, largeFaceListID string, persistedFaceID uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"largeFaceListId": autorest.Encode("path", largeFaceListID),
		"persistedFaceId": autorest.Encode("path", persistedFaceID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPathParameters("/largefacelists/{largeFaceListId}/persistedFaces/{persistedFaceId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetFaceSender sends the GetFace request. The method will close the
// http.Response Body if it receives an error.
func (client LargeFaceListClient) GetFaceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetFaceResponder handles the response to the GetFace request. The method always
// closes the http.Response Body.
func (client LargeFaceListClient) GetFaceResponder(resp *http.Response) (result PersistedFace, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetTrainingStatus retrieve the training status of a large face list (completed or ongoing).
// Parameters:
// largeFaceListID - id referencing a particular large face list.
func (client LargeFaceListClient) GetTrainingStatus(ctx context.Context, largeFaceListID string) (result TrainingStatus, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: largeFaceListID,
			Constraints: []validation.Constraint{{Target: "largeFaceListID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "largeFaceListID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.LargeFaceListClient", "GetTrainingStatus", err.Error())
	}

	req, err := client.GetTrainingStatusPreparer(ctx, largeFaceListID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "GetTrainingStatus", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetTrainingStatusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "GetTrainingStatus", resp, "Failure sending request")
		return
	}

	result, err = client.GetTrainingStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "GetTrainingStatus", resp, "Failure responding to request")
	}

	return
}

// GetTrainingStatusPreparer prepares the GetTrainingStatus request.
func (client LargeFaceListClient) GetTrainingStatusPreparer(ctx context.Context, largeFaceListID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"largeFaceListId": autorest.Encode("path", largeFaceListID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPathParameters("/largefacelists/{largeFaceListId}/training", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetTrainingStatusSender sends the GetTrainingStatus request. The method will close the
// http.Response Body if it receives an error.
func (client LargeFaceListClient) GetTrainingStatusSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetTrainingStatusResponder handles the response to the GetTrainingStatus request. The method always
// closes the http.Response Body.
func (client LargeFaceListClient) GetTrainingStatusResponder(resp *http.Response) (result TrainingStatus, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List retrieve information about all existing large face lists. Only largeFaceListId, name and userData will be
// returned.
func (client LargeFaceListClient) List(ctx context.Context) (result ListLargeFaceList, err error) {
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client LargeFaceListClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPath("/largefacelists"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client LargeFaceListClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client LargeFaceListClient) ListResponder(resp *http.Response) (result ListLargeFaceList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Train queue a large face list training task, the training task may not be started immediately.
// Parameters:
// largeFaceListID - id referencing a particular large face list.
func (client LargeFaceListClient) Train(ctx context.Context, largeFaceListID string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: largeFaceListID,
			Constraints: []validation.Constraint{{Target: "largeFaceListID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "largeFaceListID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.LargeFaceListClient", "Train", err.Error())
	}

	req, err := client.TrainPreparer(ctx, largeFaceListID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "Train", nil, "Failure preparing request")
		return
	}

	resp, err := client.TrainSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "Train", resp, "Failure sending request")
		return
	}

	result, err = client.TrainResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "Train", resp, "Failure responding to request")
	}

	return
}

// TrainPreparer prepares the Train request.
func (client LargeFaceListClient) TrainPreparer(ctx context.Context, largeFaceListID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"largeFaceListId": autorest.Encode("path", largeFaceListID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPathParameters("/largefacelists/{largeFaceListId}/train", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// TrainSender sends the Train request. The method will close the
// http.Response Body if it receives an error.
func (client LargeFaceListClient) TrainSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// TrainResponder handles the response to the Train request. The method always
// closes the http.Response Body.
func (client LargeFaceListClient) TrainResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update update information of a large face list.
// Parameters:
// largeFaceListID - id referencing a particular large face list.
// body - request body for updating a large face list.
func (client LargeFaceListClient) Update(ctx context.Context, largeFaceListID string, body NameAndUserDataContract) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: largeFaceListID,
			Constraints: []validation.Constraint{{Target: "largeFaceListID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "largeFaceListID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.LargeFaceListClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, largeFaceListID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client LargeFaceListClient) UpdatePreparer(ctx context.Context, largeFaceListID string, body NameAndUserDataContract) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"largeFaceListId": autorest.Encode("path", largeFaceListID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPathParameters("/largefacelists/{largeFaceListId}", pathParameters),
		autorest.WithJSON(body))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client LargeFaceListClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client LargeFaceListClient) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// UpdateFace update a persisted face's userData field.
// Parameters:
// largeFaceListID - id referencing a particular large face list.
// persistedFaceID - id referencing a particular persistedFaceId of an existing face.
// body - request body for updating persisted face.
func (client LargeFaceListClient) UpdateFace(ctx context.Context, largeFaceListID string, persistedFaceID uuid.UUID, body UpdateFaceRequest) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: largeFaceListID,
			Constraints: []validation.Constraint{{Target: "largeFaceListID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "largeFaceListID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.LargeFaceListClient", "UpdateFace", err.Error())
	}

	req, err := client.UpdateFacePreparer(ctx, largeFaceListID, persistedFaceID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "UpdateFace", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateFaceSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "UpdateFace", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateFaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargeFaceListClient", "UpdateFace", resp, "Failure responding to request")
	}

	return
}

// UpdateFacePreparer prepares the UpdateFace request.
func (client LargeFaceListClient) UpdateFacePreparer(ctx context.Context, largeFaceListID string, persistedFaceID uuid.UUID, body UpdateFaceRequest) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"largeFaceListId": autorest.Encode("path", largeFaceListID),
		"persistedFaceId": autorest.Encode("path", persistedFaceID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPathParameters("/largefacelists/{largeFaceListId}/persistedFaces/{persistedFaceId}", pathParameters),
		autorest.WithJSON(body))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateFaceSender sends the UpdateFace request. The method will close the
// http.Response Body if it receives an error.
func (client LargeFaceListClient) UpdateFaceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateFaceResponder handles the response to the UpdateFace request. The method always
// closes the http.Response Body.
func (client LargeFaceListClient) UpdateFaceResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
