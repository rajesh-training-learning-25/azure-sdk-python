package batch

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"github.com/gofrs/uuid"
	"net/http"
)

// CertificateClient is the a client for issuing REST requests to the Azure Batch service.
type CertificateClient struct {
	BaseClient
}

// NewCertificateClient creates an instance of the CertificateClient client.
func NewCertificateClient(batchURL string) CertificateClient {
	return CertificateClient{New(batchURL)}
}

// Add sends the add request.
// Parameters:
// certificate - the Certificate to be added.
// timeout - the maximum time that the server can spend processing the request, in seconds. The default is 30
// seconds.
// clientRequestID - the caller-generated request identity, in the form of a GUID with no decoration such as
// curly braces, e.g. 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.
// returnClientRequestID - whether the server should return the client-request-id in the response.
// ocpDate - the time the request was issued. Client libraries typically set this to the current system clock
// time; set it explicitly if you are calling the REST API directly.
func (client CertificateClient) Add(ctx context.Context, certificate CertificateAddParameter, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CertificateClient.Add")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: certificate,
			Constraints: []validation.Constraint{{Target: "certificate.Thumbprint", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "certificate.ThumbprintAlgorithm", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "certificate.Data", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batch.CertificateClient", "Add", err.Error())
	}

	req, err := client.AddPreparer(ctx, certificate, timeout, clientRequestID, returnClientRequestID, ocpDate)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "Add", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "Add", resp, "Failure sending request")
		return
	}

	result, err = client.AddResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "Add", resp, "Failure responding to request")
		return
	}

	return
}

// AddPreparer prepares the Add request.
func (client CertificateClient) AddPreparer(ctx context.Context, certificate CertificateAddParameter, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"batchUrl": client.BatchURL,
	}

	const APIVersion = "2020-09-01.12.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *timeout)
	} else {
		queryParameters["timeout"] = autorest.Encode("query", 30)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; odata=minimalmetadata; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{batchUrl}", urlParameters),
		autorest.WithPath("/certificates"),
		autorest.WithJSON(certificate),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	if returnClientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("return-client-request-id", autorest.String(returnClientRequestID)))
	} else {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("return-client-request-id", autorest.String(false)))
	}
	if ocpDate != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("ocp-date", autorest.String(ocpDate)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddSender sends the Add request. The method will close the
// http.Response Body if it receives an error.
func (client CertificateClient) AddSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AddResponder handles the response to the Add request. The method always
// closes the http.Response Body.
func (client CertificateClient) AddResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CancelDeletion if you try to delete a Certificate that is being used by a Pool or Compute Node, the status of the
// Certificate changes to deleteFailed. If you decide that you want to continue using the Certificate, you can use this
// operation to set the status of the Certificate back to active. If you intend to delete the Certificate, you do not
// need to run this operation after the deletion failed. You must make sure that the Certificate is not being used by
// any resources, and then you can try again to delete the Certificate.
// Parameters:
// thumbprintAlgorithm - the algorithm used to derive the thumbprint parameter. This must be sha1.
// thumbprint - the thumbprint of the Certificate being deleted.
// timeout - the maximum time that the server can spend processing the request, in seconds. The default is 30
// seconds.
// clientRequestID - the caller-generated request identity, in the form of a GUID with no decoration such as
// curly braces, e.g. 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.
// returnClientRequestID - whether the server should return the client-request-id in the response.
// ocpDate - the time the request was issued. Client libraries typically set this to the current system clock
// time; set it explicitly if you are calling the REST API directly.
func (client CertificateClient) CancelDeletion(ctx context.Context, thumbprintAlgorithm string, thumbprint string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CertificateClient.CancelDeletion")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CancelDeletionPreparer(ctx, thumbprintAlgorithm, thumbprint, timeout, clientRequestID, returnClientRequestID, ocpDate)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "CancelDeletion", nil, "Failure preparing request")
		return
	}

	resp, err := client.CancelDeletionSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "CancelDeletion", resp, "Failure sending request")
		return
	}

	result, err = client.CancelDeletionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "CancelDeletion", resp, "Failure responding to request")
		return
	}

	return
}

// CancelDeletionPreparer prepares the CancelDeletion request.
func (client CertificateClient) CancelDeletionPreparer(ctx context.Context, thumbprintAlgorithm string, thumbprint string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"batchUrl": client.BatchURL,
	}

	pathParameters := map[string]interface{}{
		"thumbprint":          autorest.Encode("path", thumbprint),
		"thumbprintAlgorithm": autorest.Encode("path", thumbprintAlgorithm),
	}

	const APIVersion = "2020-09-01.12.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *timeout)
	} else {
		queryParameters["timeout"] = autorest.Encode("query", 30)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{batchUrl}", urlParameters),
		autorest.WithPathParameters("/certificates(thumbprintAlgorithm={thumbprintAlgorithm},thumbprint={thumbprint})/canceldelete", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	if returnClientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("return-client-request-id", autorest.String(returnClientRequestID)))
	} else {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("return-client-request-id", autorest.String(false)))
	}
	if ocpDate != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("ocp-date", autorest.String(ocpDate)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CancelDeletionSender sends the CancelDeletion request. The method will close the
// http.Response Body if it receives an error.
func (client CertificateClient) CancelDeletionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CancelDeletionResponder handles the response to the CancelDeletion request. The method always
// closes the http.Response Body.
func (client CertificateClient) CancelDeletionResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete you cannot delete a Certificate if a resource (Pool or Compute Node) is using it. Before you can delete a
// Certificate, you must therefore make sure that the Certificate is not associated with any existing Pools, the
// Certificate is not installed on any Nodes (even if you remove a Certificate from a Pool, it is not removed from
// existing Compute Nodes in that Pool until they restart), and no running Tasks depend on the Certificate. If you try
// to delete a Certificate that is in use, the deletion fails. The Certificate status changes to deleteFailed. You can
// use Cancel Delete Certificate to set the status back to active if you decide that you want to continue using the
// Certificate.
// Parameters:
// thumbprintAlgorithm - the algorithm used to derive the thumbprint parameter. This must be sha1.
// thumbprint - the thumbprint of the Certificate to be deleted.
// timeout - the maximum time that the server can spend processing the request, in seconds. The default is 30
// seconds.
// clientRequestID - the caller-generated request identity, in the form of a GUID with no decoration such as
// curly braces, e.g. 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.
// returnClientRequestID - whether the server should return the client-request-id in the response.
// ocpDate - the time the request was issued. Client libraries typically set this to the current system clock
// time; set it explicitly if you are calling the REST API directly.
func (client CertificateClient) Delete(ctx context.Context, thumbprintAlgorithm string, thumbprint string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CertificateClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, thumbprintAlgorithm, thumbprint, timeout, clientRequestID, returnClientRequestID, ocpDate)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client CertificateClient) DeletePreparer(ctx context.Context, thumbprintAlgorithm string, thumbprint string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"batchUrl": client.BatchURL,
	}

	pathParameters := map[string]interface{}{
		"thumbprint":          autorest.Encode("path", thumbprint),
		"thumbprintAlgorithm": autorest.Encode("path", thumbprintAlgorithm),
	}

	const APIVersion = "2020-09-01.12.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *timeout)
	} else {
		queryParameters["timeout"] = autorest.Encode("query", 30)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("{batchUrl}", urlParameters),
		autorest.WithPathParameters("/certificates(thumbprintAlgorithm={thumbprintAlgorithm},thumbprint={thumbprint})", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	if returnClientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("return-client-request-id", autorest.String(returnClientRequestID)))
	} else {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("return-client-request-id", autorest.String(false)))
	}
	if ocpDate != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("ocp-date", autorest.String(ocpDate)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client CertificateClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client CertificateClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets information about the specified Certificate.
// Parameters:
// thumbprintAlgorithm - the algorithm used to derive the thumbprint parameter. This must be sha1.
// thumbprint - the thumbprint of the Certificate to get.
// selectParameter - an OData $select clause.
// timeout - the maximum time that the server can spend processing the request, in seconds. The default is 30
// seconds.
// clientRequestID - the caller-generated request identity, in the form of a GUID with no decoration such as
// curly braces, e.g. 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.
// returnClientRequestID - whether the server should return the client-request-id in the response.
// ocpDate - the time the request was issued. Client libraries typically set this to the current system clock
// time; set it explicitly if you are calling the REST API directly.
func (client CertificateClient) Get(ctx context.Context, thumbprintAlgorithm string, thumbprint string, selectParameter string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result Certificate, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CertificateClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, thumbprintAlgorithm, thumbprint, selectParameter, timeout, clientRequestID, returnClientRequestID, ocpDate)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client CertificateClient) GetPreparer(ctx context.Context, thumbprintAlgorithm string, thumbprint string, selectParameter string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"batchUrl": client.BatchURL,
	}

	pathParameters := map[string]interface{}{
		"thumbprint":          autorest.Encode("path", thumbprint),
		"thumbprintAlgorithm": autorest.Encode("path", thumbprintAlgorithm),
	}

	const APIVersion = "2020-09-01.12.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}
	if timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *timeout)
	} else {
		queryParameters["timeout"] = autorest.Encode("query", 30)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{batchUrl}", urlParameters),
		autorest.WithPathParameters("/certificates(thumbprintAlgorithm={thumbprintAlgorithm},thumbprint={thumbprint})", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	if returnClientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("return-client-request-id", autorest.String(returnClientRequestID)))
	} else {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("return-client-request-id", autorest.String(false)))
	}
	if ocpDate != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("ocp-date", autorest.String(ocpDate)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client CertificateClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client CertificateClient) GetResponder(resp *http.Response) (result Certificate, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List sends the list request.
// Parameters:
// filter - an OData $filter clause. For more information on constructing this filter, see
// https://docs.microsoft.com/en-us/rest/api/batchservice/odata-filters-in-batch#list-certificates.
// selectParameter - an OData $select clause.
// maxResults - the maximum number of items to return in the response. A maximum of 1000 Certificates can be
// returned.
// timeout - the maximum time that the server can spend processing the request, in seconds. The default is 30
// seconds.
// clientRequestID - the caller-generated request identity, in the form of a GUID with no decoration such as
// curly braces, e.g. 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.
// returnClientRequestID - whether the server should return the client-request-id in the response.
// ocpDate - the time the request was issued. Client libraries typically set this to the current system clock
// time; set it explicitly if you are calling the REST API directly.
func (client CertificateClient) List(ctx context.Context, filter string, selectParameter string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result CertificateListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CertificateClient.List")
		defer func() {
			sc := -1
			if result.clr.Response.Response != nil {
				sc = result.clr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: maxResults,
			Constraints: []validation.Constraint{{Target: "maxResults", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "maxResults", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
					{Target: "maxResults", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("batch.CertificateClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, filter, selectParameter, maxResults, timeout, clientRequestID, returnClientRequestID, ocpDate)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.clr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "List", resp, "Failure sending request")
		return
	}

	result.clr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "List", resp, "Failure responding to request")
		return
	}
	if result.clr.hasNextLink() && result.clr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client CertificateClient) ListPreparer(ctx context.Context, filter string, selectParameter string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"batchUrl": client.BatchURL,
	}

	const APIVersion = "2020-09-01.12.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}
	if maxResults != nil {
		queryParameters["maxresults"] = autorest.Encode("query", *maxResults)
	} else {
		queryParameters["maxresults"] = autorest.Encode("query", 1000)
	}
	if timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *timeout)
	} else {
		queryParameters["timeout"] = autorest.Encode("query", 30)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{batchUrl}", urlParameters),
		autorest.WithPath("/certificates"),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	if returnClientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("return-client-request-id", autorest.String(returnClientRequestID)))
	} else {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("return-client-request-id", autorest.String(false)))
	}
	if ocpDate != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("ocp-date", autorest.String(ocpDate)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client CertificateClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client CertificateClient) ListResponder(resp *http.Response) (result CertificateListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client CertificateClient) listNextResults(ctx context.Context, lastResults CertificateListResult) (result CertificateListResult, err error) {
	req, err := lastResults.certificateListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batch.CertificateClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batch.CertificateClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client CertificateClient) ListComplete(ctx context.Context, filter string, selectParameter string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result CertificateListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CertificateClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, filter, selectParameter, maxResults, timeout, clientRequestID, returnClientRequestID, ocpDate)
	return
}
