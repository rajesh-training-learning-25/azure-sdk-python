package speechservices

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
	"github.com/Azure/go-autorest/tracing"
	"github.com/satori/go.uuid"
	"io"
	"net/http"
)

// VoiceDatasetClient is the speech Services API v2.0.
type VoiceDatasetClient struct {
	BaseClient
}

// NewVoiceDatasetClient creates an instance of the VoiceDatasetClient client.
func NewVoiceDatasetClient() VoiceDatasetClient {
	return NewVoiceDatasetClientWithBaseURI(DefaultBaseURI)
}

// NewVoiceDatasetClientWithBaseURI creates an instance of the VoiceDatasetClient client.
func NewVoiceDatasetClientWithBaseURI(baseURI string) VoiceDatasetClient {
	return VoiceDatasetClient{NewWithBaseURI(baseURI)}
}

// Delete sends the delete request.
// Parameters:
// ID - the identifier of the voice dataset
func (client VoiceDatasetClient) Delete(ctx context.Context, ID uuid.UUID) (result ErrorContent, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VoiceDatasetClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, ID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "speechservices.VoiceDatasetClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "speechservices.VoiceDatasetClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "speechservices.VoiceDatasetClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client VoiceDatasetClient) DeletePreparer(ctx context.Context, ID uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"id": autorest.Encode("path", ID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/api/texttospeech/v2.0/datasets/{id}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VoiceDatasetClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client VoiceDatasetClient) DeleteResponder(resp *http.Response) (result ErrorContent, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusUnauthorized, http.StatusForbidden, http.StatusMethodNotAllowed, http.StatusTooManyRequests),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update sends the update request.
// Parameters:
// ID - the identifier of the voice dataset.
// datasetUpdate - the updated values for the voice dataset.
func (client VoiceDatasetClient) Update(ctx context.Context, ID uuid.UUID, datasetUpdate DatasetUpdate) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VoiceDatasetClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, ID, datasetUpdate)
	if err != nil {
		err = autorest.NewErrorWithError(err, "speechservices.VoiceDatasetClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "speechservices.VoiceDatasetClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "speechservices.VoiceDatasetClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client VoiceDatasetClient) UpdatePreparer(ctx context.Context, ID uuid.UUID, datasetUpdate DatasetUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"id": autorest.Encode("path", ID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/api/texttospeech/v2.0/datasets/{id}", pathParameters),
		autorest.WithJSON(datasetUpdate))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client VoiceDatasetClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client VoiceDatasetClient) UpdateResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusConflict, http.StatusUnsupportedMediaType, http.StatusTooManyRequests),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Upload sends the upload request.
// Parameters:
// name - the name of this data import (always add this string for any import).
// description - optional description of this data import.
// locale - the locale of this data import (always add this string for any import).
// dataImportKind - the kind of the data import (always add this string for any import).
// properties - optional properties of this data import (json serialized object with key/values, where all
// values must be strings)
// audiodata - a zip file containing the audio data.
// transcriptions - the transcriptions text file of the audio data.
func (client VoiceDatasetClient) Upload(ctx context.Context, name string, description string, locale string, dataImportKind string, properties string, audiodata io.ReadCloser, transcriptions io.ReadCloser) (result VoiceDatasetUploadFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VoiceDatasetClient.Upload")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UploadPreparer(ctx, name, description, locale, dataImportKind, properties, audiodata, transcriptions)
	if err != nil {
		err = autorest.NewErrorWithError(err, "speechservices.VoiceDatasetClient", "Upload", nil, "Failure preparing request")
		return
	}

	result, err = client.UploadSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "speechservices.VoiceDatasetClient", "Upload", result.Response(), "Failure sending request")
		return
	}

	return
}

// UploadPreparer prepares the Upload request.
func (client VoiceDatasetClient) UploadPreparer(ctx context.Context, name string, description string, locale string, dataImportKind string, properties string, audiodata io.ReadCloser, transcriptions io.ReadCloser) (*http.Request, error) {
	formDataParameters := map[string]interface{}{}
	if len(name) > 0 {
		formDataParameters["name"] = name
	}
	if len(description) > 0 {
		formDataParameters["description"] = description
	}
	if len(locale) > 0 {
		formDataParameters["locale"] = locale
	}
	if len(string(dataImportKind)) > 0 {
		formDataParameters["dataImportKind"] = dataImportKind
	}
	if len(properties) > 0 {
		formDataParameters["properties"] = properties
	}
	if audiodata != nil {
		formDataParameters["audiodata"] = audiodata
	}
	if transcriptions != nil {
		formDataParameters["transcriptions"] = transcriptions
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/api/texttospeech/v2.0/datasets/upload"),
		autorest.WithMultiPartFormData(formDataParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UploadSender sends the Upload request. The method will close the
// http.Response Body if it receives an error.
func (client VoiceDatasetClient) UploadSender(req *http.Request) (future VoiceDatasetUploadFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// UploadResponder handles the response to the Upload request. The method always
// closes the http.Response Body.
func (client VoiceDatasetClient) UploadResponder(resp *http.Response) (result ErrorContent, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusUnsupportedMediaType, http.StatusTooManyRequests),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
