//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package internal

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ContainerRegistryBlobClient contains the methods for the ContainerRegistryBlob group.
// Don't use this type directly, use NewContainerRegistryBlobClient() instead.
type ContainerRegistryBlobClient struct {
	endpoint string
	pl       runtime.Pipeline
}

// NewContainerRegistryBlobClient creates a new instance of ContainerRegistryBlobClient with the specified values.
//   - endpoint - Registry login URL
//   - pl - the pipeline used for sending requests and handling responses.
func NewContainerRegistryBlobClient(endpoint string, pl runtime.Pipeline) *ContainerRegistryBlobClient {
	client := &ContainerRegistryBlobClient{
		endpoint: endpoint,
		pl:       pl,
	}
	return client
}

// CancelUpload - Cancel outstanding upload processes, releasing associated resources. If this is not called, the unfinished
// uploads will eventually timeout.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - nextLink - Link acquired from upload start or previous chunk. Note, do not include initial / (must do substring(1) )
//   - options - ContainerRegistryBlobClientCancelUploadOptions contains the optional parameters for the ContainerRegistryBlobClient.CancelUpload
//     method.
func (client *ContainerRegistryBlobClient) CancelUpload(ctx context.Context, nextLink string, options *ContainerRegistryBlobClientCancelUploadOptions) (ContainerRegistryBlobClientCancelUploadResponse, error) {
	req, err := client.cancelUploadCreateRequest(ctx, nextLink, options)
	if err != nil {
		return ContainerRegistryBlobClientCancelUploadResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContainerRegistryBlobClientCancelUploadResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return ContainerRegistryBlobClientCancelUploadResponse{}, runtime.NewResponseError(resp)
	}
	return ContainerRegistryBlobClientCancelUploadResponse{}, nil
}

// cancelUploadCreateRequest creates the CancelUpload request.
func (client *ContainerRegistryBlobClient) cancelUploadCreateRequest(ctx context.Context, nextLink string, options *ContainerRegistryBlobClientCancelUploadOptions) (*policy.Request, error) {
	urlPath := "/{nextBlobUuidLink}"
	urlPath = strings.ReplaceAll(urlPath, "{nextBlobUuidLink}", nextLink)
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// CheckBlobExists - Same as GET, except only the headers are returned.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - name - Name of the image (including the namespace)
//   - digest - Digest of a BLOB
//   - options - ContainerRegistryBlobClientCheckBlobExistsOptions contains the optional parameters for the ContainerRegistryBlobClient.CheckBlobExists
//     method.
func (client *ContainerRegistryBlobClient) CheckBlobExists(ctx context.Context, name string, digest string, options *ContainerRegistryBlobClientCheckBlobExistsOptions) (ContainerRegistryBlobClientCheckBlobExistsResponse, error) {
	req, err := client.checkBlobExistsCreateRequest(ctx, name, digest, options)
	if err != nil {
		return ContainerRegistryBlobClientCheckBlobExistsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContainerRegistryBlobClientCheckBlobExistsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ContainerRegistryBlobClientCheckBlobExistsResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkBlobExistsHandleResponse(resp)
}

// checkBlobExistsCreateRequest creates the CheckBlobExists request.
func (client *ContainerRegistryBlobClient) checkBlobExistsCreateRequest(ctx context.Context, name string, digest string, options *ContainerRegistryBlobClientCheckBlobExistsOptions) (*policy.Request, error) {
	urlPath := "/v2/{name}/blobs/{digest}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if digest == "" {
		return nil, errors.New("parameter digest cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{digest}", url.PathEscape(digest))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// checkBlobExistsHandleResponse handles the CheckBlobExists response.
func (client *ContainerRegistryBlobClient) checkBlobExistsHandleResponse(resp *http.Response) (ContainerRegistryBlobClientCheckBlobExistsResponse, error) {
	result := ContainerRegistryBlobClientCheckBlobExistsResponse{}
	if val := resp.Header.Get("Content-Length"); val != "" {
		contentLength, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return ContainerRegistryBlobClientCheckBlobExistsResponse{}, err
		}
		result.ContentLength = &contentLength
	}
	if val := resp.Header.Get("Docker-Content-Digest"); val != "" {
		result.DockerContentDigest = &val
	}
	return result, nil
}

// CheckChunkExists - Same as GET, except only the headers are returned.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - name - Name of the image (including the namespace)
//   - digest - Digest of a BLOB
//   - rangeParam - Format : bytes=-, HTTP Range header specifying blob chunk.
//   - options - ContainerRegistryBlobClientCheckChunkExistsOptions contains the optional parameters for the ContainerRegistryBlobClient.CheckChunkExists
//     method.
func (client *ContainerRegistryBlobClient) CheckChunkExists(ctx context.Context, name string, digest string, rangeParam string, options *ContainerRegistryBlobClientCheckChunkExistsOptions) (ContainerRegistryBlobClientCheckChunkExistsResponse, error) {
	req, err := client.checkChunkExistsCreateRequest(ctx, name, digest, rangeParam, options)
	if err != nil {
		return ContainerRegistryBlobClientCheckChunkExistsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContainerRegistryBlobClientCheckChunkExistsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ContainerRegistryBlobClientCheckChunkExistsResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkChunkExistsHandleResponse(resp)
}

// checkChunkExistsCreateRequest creates the CheckChunkExists request.
func (client *ContainerRegistryBlobClient) checkChunkExistsCreateRequest(ctx context.Context, name string, digest string, rangeParam string, options *ContainerRegistryBlobClientCheckChunkExistsOptions) (*policy.Request, error) {
	urlPath := "/v2/{name}/blobs/{digest}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if digest == "" {
		return nil, errors.New("parameter digest cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{digest}", url.PathEscape(digest))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Range"] = []string{rangeParam}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// checkChunkExistsHandleResponse handles the CheckChunkExists response.
func (client *ContainerRegistryBlobClient) checkChunkExistsHandleResponse(resp *http.Response) (ContainerRegistryBlobClientCheckChunkExistsResponse, error) {
	result := ContainerRegistryBlobClientCheckChunkExistsResponse{}
	if val := resp.Header.Get("Content-Length"); val != "" {
		contentLength, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return ContainerRegistryBlobClientCheckChunkExistsResponse{}, err
		}
		result.ContentLength = &contentLength
	}
	if val := resp.Header.Get("Content-Range"); val != "" {
		result.ContentRange = &val
	}
	return result, nil
}

// CompleteUploadWithBinary - Complete the upload, providing all the data in the body, if necessary. A request without a body
// will just complete the upload with previously uploaded content.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - digest - Digest of a BLOB
//   - nextLink - Link acquired from upload start or previous chunk. Note, do not include initial / (must do substring(1) )
//   - options - ContainerRegistryBlobClientCompleteUploadWithBinaryOptions contains the optional parameters for the ContainerRegistryBlobClient.CompleteUploadWithBinary
//     method.
func (client *ContainerRegistryBlobClient) CompleteUploadWithBinary(ctx context.Context, digest string, nextLink string, options *ContainerRegistryBlobClientCompleteUploadWithBinaryOptions) (ContainerRegistryBlobClientCompleteUploadWithBinaryResponse, error) {
	req, err := client.completeUploadWithBinaryCreateRequest(ctx, digest, nextLink, options)
	if err != nil {
		return ContainerRegistryBlobClientCompleteUploadWithBinaryResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContainerRegistryBlobClientCompleteUploadWithBinaryResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return ContainerRegistryBlobClientCompleteUploadWithBinaryResponse{}, runtime.NewResponseError(resp)
	}
	return client.completeUploadWithBinaryHandleResponse(resp)
}

// completeUploadWithBinaryCreateRequest creates the CompleteUploadWithBinary request.
func (client *ContainerRegistryBlobClient) completeUploadWithBinaryCreateRequest(ctx context.Context, digest string, nextLink string, options *ContainerRegistryBlobClientCompleteUploadWithBinaryOptions) (*policy.Request, error) {
	urlPath := "/{nextBlobUuidLink}"
	urlPath = strings.ReplaceAll(urlPath, "{nextBlobUuidLink}", nextLink)
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("digest", digest)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Value != nil {
		return req, req.SetBody(options.Value, "application/octet-stream")
	}
	return req, nil
}

// completeUploadWithBinaryHandleResponse handles the CompleteUploadWithBinary response.
func (client *ContainerRegistryBlobClient) completeUploadWithBinaryHandleResponse(resp *http.Response) (ContainerRegistryBlobClientCompleteUploadWithBinaryResponse, error) {
	result := ContainerRegistryBlobClientCompleteUploadWithBinaryResponse{}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	if val := resp.Header.Get("Range"); val != "" {
		result.Range = &val
	}
	if val := resp.Header.Get("Docker-Content-Digest"); val != "" {
		result.DockerContentDigest = &val
	}
	return result, nil
}

// DeleteBlob - Removes an already uploaded blob.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - name - Name of the image (including the namespace)
//   - digest - Digest of a BLOB
//   - options - ContainerRegistryBlobClientDeleteBlobOptions contains the optional parameters for the ContainerRegistryBlobClient.DeleteBlob
//     method.
func (client *ContainerRegistryBlobClient) DeleteBlob(ctx context.Context, name string, digest string, options *ContainerRegistryBlobClientDeleteBlobOptions) (ContainerRegistryBlobClientDeleteBlobResponse, error) {
	req, err := client.deleteBlobCreateRequest(ctx, name, digest, options)
	if err != nil {
		return ContainerRegistryBlobClientDeleteBlobResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContainerRegistryBlobClientDeleteBlobResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return ContainerRegistryBlobClientDeleteBlobResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteBlobHandleResponse(resp)
}

// deleteBlobCreateRequest creates the DeleteBlob request.
func (client *ContainerRegistryBlobClient) deleteBlobCreateRequest(ctx context.Context, name string, digest string, options *ContainerRegistryBlobClientDeleteBlobOptions) (*policy.Request, error) {
	urlPath := "/v2/{name}/blobs/{digest}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if digest == "" {
		return nil, errors.New("parameter digest cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{digest}", url.PathEscape(digest))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	runtime.SkipBodyDownload(req)
	req.Raw().Header["Accept"] = []string{"application/octet-stream"}
	return req, nil
}

// deleteBlobHandleResponse handles the DeleteBlob response.
func (client *ContainerRegistryBlobClient) deleteBlobHandleResponse(resp *http.Response) (ContainerRegistryBlobClientDeleteBlobResponse, error) {
	result := ContainerRegistryBlobClientDeleteBlobResponse{Body: resp.Body}
	if val := resp.Header.Get("Docker-Content-Digest"); val != "" {
		result.DockerContentDigest = &val
	}
	return result, nil
}

// GetBlob - Retrieve the blob from the registry identified by digest.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - name - Name of the image (including the namespace)
//   - digest - Digest of a BLOB
//   - options - ContainerRegistryBlobClientGetBlobOptions contains the optional parameters for the ContainerRegistryBlobClient.GetBlob
//     method.
func (client *ContainerRegistryBlobClient) GetBlob(ctx context.Context, name string, digest string, options *ContainerRegistryBlobClientGetBlobOptions) (ContainerRegistryBlobClientGetBlobResponse, error) {
	req, err := client.getBlobCreateRequest(ctx, name, digest, options)
	if err != nil {
		return ContainerRegistryBlobClientGetBlobResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContainerRegistryBlobClientGetBlobResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ContainerRegistryBlobClientGetBlobResponse{}, runtime.NewResponseError(resp)
	}
	return client.getBlobHandleResponse(resp)
}

// getBlobCreateRequest creates the GetBlob request.
func (client *ContainerRegistryBlobClient) getBlobCreateRequest(ctx context.Context, name string, digest string, options *ContainerRegistryBlobClientGetBlobOptions) (*policy.Request, error) {
	urlPath := "/v2/{name}/blobs/{digest}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if digest == "" {
		return nil, errors.New("parameter digest cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{digest}", url.PathEscape(digest))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	runtime.SkipBodyDownload(req)
	req.Raw().Header["Accept"] = []string{"application/octet-stream"}
	return req, nil
}

// getBlobHandleResponse handles the GetBlob response.
func (client *ContainerRegistryBlobClient) getBlobHandleResponse(resp *http.Response) (ContainerRegistryBlobClientGetBlobResponse, error) {
	result := ContainerRegistryBlobClientGetBlobResponse{Body: resp.Body}
	if val := resp.Header.Get("Content-Length"); val != "" {
		contentLength, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return ContainerRegistryBlobClientGetBlobResponse{}, err
		}
		result.ContentLength = &contentLength
	}
	if val := resp.Header.Get("Docker-Content-Digest"); val != "" {
		result.DockerContentDigest = &val
	}
	return result, nil
}

// GetChunk - Retrieve the blob from the registry identified by digest. This endpoint may also support RFC7233 compliant range
// requests. Support can be detected by issuing a HEAD request. If the header
// Accept-Range: bytes is returned, range requests can be used to fetch partial content.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - name - Name of the image (including the namespace)
//   - digest - Digest of a BLOB
//   - rangeParam - Format : bytes=-, HTTP Range header specifying blob chunk.
//   - options - ContainerRegistryBlobClientGetChunkOptions contains the optional parameters for the ContainerRegistryBlobClient.GetChunk
//     method.
func (client *ContainerRegistryBlobClient) GetChunk(ctx context.Context, name string, digest string, rangeParam string, options *ContainerRegistryBlobClientGetChunkOptions) (ContainerRegistryBlobClientGetChunkResponse, error) {
	req, err := client.getChunkCreateRequest(ctx, name, digest, rangeParam, options)
	if err != nil {
		return ContainerRegistryBlobClientGetChunkResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContainerRegistryBlobClientGetChunkResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusPartialContent) {
		return ContainerRegistryBlobClientGetChunkResponse{}, runtime.NewResponseError(resp)
	}
	return client.getChunkHandleResponse(resp)
}

// getChunkCreateRequest creates the GetChunk request.
func (client *ContainerRegistryBlobClient) getChunkCreateRequest(ctx context.Context, name string, digest string, rangeParam string, options *ContainerRegistryBlobClientGetChunkOptions) (*policy.Request, error) {
	urlPath := "/v2/{name}/blobs/{digest}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if digest == "" {
		return nil, errors.New("parameter digest cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{digest}", url.PathEscape(digest))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	runtime.SkipBodyDownload(req)
	req.Raw().Header["Range"] = []string{rangeParam}
	req.Raw().Header["Accept"] = []string{"application/octet-stream"}
	return req, nil
}

// getChunkHandleResponse handles the GetChunk response.
func (client *ContainerRegistryBlobClient) getChunkHandleResponse(resp *http.Response) (ContainerRegistryBlobClientGetChunkResponse, error) {
	result := ContainerRegistryBlobClientGetChunkResponse{Body: resp.Body}
	if val := resp.Header.Get("Content-Length"); val != "" {
		contentLength, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return ContainerRegistryBlobClientGetChunkResponse{}, err
		}
		result.ContentLength = &contentLength
	}
	if val := resp.Header.Get("Content-Range"); val != "" {
		result.ContentRange = &val
	}
	return result, nil
}

// GetUploadStatus - Retrieve status of upload identified by uuid. The primary purpose of this endpoint is to resolve the
// current status of a resumable upload.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - nextLink - Link acquired from upload start or previous chunk. Note, do not include initial / (must do substring(1) )
//   - options - ContainerRegistryBlobClientGetUploadStatusOptions contains the optional parameters for the ContainerRegistryBlobClient.GetUploadStatus
//     method.
func (client *ContainerRegistryBlobClient) GetUploadStatus(ctx context.Context, nextLink string, options *ContainerRegistryBlobClientGetUploadStatusOptions) (ContainerRegistryBlobClientGetUploadStatusResponse, error) {
	req, err := client.getUploadStatusCreateRequest(ctx, nextLink, options)
	if err != nil {
		return ContainerRegistryBlobClientGetUploadStatusResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContainerRegistryBlobClientGetUploadStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return ContainerRegistryBlobClientGetUploadStatusResponse{}, runtime.NewResponseError(resp)
	}
	return client.getUploadStatusHandleResponse(resp)
}

// getUploadStatusCreateRequest creates the GetUploadStatus request.
func (client *ContainerRegistryBlobClient) getUploadStatusCreateRequest(ctx context.Context, nextLink string, options *ContainerRegistryBlobClientGetUploadStatusOptions) (*policy.Request, error) {
	urlPath := "/{nextBlobUuidLink}"
	urlPath = strings.ReplaceAll(urlPath, "{nextBlobUuidLink}", nextLink)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getUploadStatusHandleResponse handles the GetUploadStatus response.
func (client *ContainerRegistryBlobClient) getUploadStatusHandleResponse(resp *http.Response) (ContainerRegistryBlobClientGetUploadStatusResponse, error) {
	result := ContainerRegistryBlobClientGetUploadStatusResponse{}
	if val := resp.Header.Get("Range"); val != "" {
		result.Range = &val
	}
	if val := resp.Header.Get("Docker-Upload-UUID"); val != "" {
		result.DockerUploadUUID = &val
	}
	return result, nil
}

// MountBlob - Mount a blob identified by the mount parameter from another repository.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - name - Name of the image (including the namespace)
//   - from - Name of the source repository.
//   - mount - Digest of blob to mount from the source repository.
//   - options - ContainerRegistryBlobClientMountBlobOptions contains the optional parameters for the ContainerRegistryBlobClient.MountBlob
//     method.
func (client *ContainerRegistryBlobClient) MountBlob(ctx context.Context, name string, from string, mount string, options *ContainerRegistryBlobClientMountBlobOptions) (ContainerRegistryBlobClientMountBlobResponse, error) {
	req, err := client.mountBlobCreateRequest(ctx, name, from, mount, options)
	if err != nil {
		return ContainerRegistryBlobClientMountBlobResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContainerRegistryBlobClientMountBlobResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return ContainerRegistryBlobClientMountBlobResponse{}, runtime.NewResponseError(resp)
	}
	return client.mountBlobHandleResponse(resp)
}

// mountBlobCreateRequest creates the MountBlob request.
func (client *ContainerRegistryBlobClient) mountBlobCreateRequest(ctx context.Context, name string, from string, mount string, options *ContainerRegistryBlobClientMountBlobOptions) (*policy.Request, error) {
	urlPath := "/v2/{name}/blobs/uploads/"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("from", from)
	reqQP.Set("mount", mount)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// mountBlobHandleResponse handles the MountBlob response.
func (client *ContainerRegistryBlobClient) mountBlobHandleResponse(resp *http.Response) (ContainerRegistryBlobClientMountBlobResponse, error) {
	result := ContainerRegistryBlobClientMountBlobResponse{}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	if val := resp.Header.Get("Docker-Upload-UUID"); val != "" {
		result.DockerUploadUUID = &val
	}
	if val := resp.Header.Get("Docker-Content-Digest"); val != "" {
		result.DockerContentDigest = &val
	}
	return result, nil
}

// StartUpload - Initiate a resumable blob upload with an empty request body.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - name - Name of the image (including the namespace)
//   - options - ContainerRegistryBlobClientStartUploadOptions contains the optional parameters for the ContainerRegistryBlobClient.StartUpload
//     method.
func (client *ContainerRegistryBlobClient) StartUpload(ctx context.Context, name string, options *ContainerRegistryBlobClientStartUploadOptions) (ContainerRegistryBlobClientStartUploadResponse, error) {
	req, err := client.startUploadCreateRequest(ctx, name, options)
	if err != nil {
		return ContainerRegistryBlobClientStartUploadResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContainerRegistryBlobClientStartUploadResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return ContainerRegistryBlobClientStartUploadResponse{}, runtime.NewResponseError(resp)
	}
	return client.startUploadHandleResponse(resp)
}

// startUploadCreateRequest creates the StartUpload request.
func (client *ContainerRegistryBlobClient) startUploadCreateRequest(ctx context.Context, name string, options *ContainerRegistryBlobClientStartUploadOptions) (*policy.Request, error) {
	urlPath := "/v2/{name}/blobs/uploads/"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// startUploadHandleResponse handles the StartUpload response.
func (client *ContainerRegistryBlobClient) startUploadHandleResponse(resp *http.Response) (ContainerRegistryBlobClientStartUploadResponse, error) {
	result := ContainerRegistryBlobClientStartUploadResponse{}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	if val := resp.Header.Get("Range"); val != "" {
		result.Range = &val
	}
	if val := resp.Header.Get("Docker-Upload-UUID"); val != "" {
		result.DockerUploadUUID = &val
	}
	return result, nil
}

// UploadChunkWithBinary - Upload a stream of data without completing the upload.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - nextLink - Link acquired from upload start or previous chunk. Note, do not include initial / (must do substring(1) )
//   - value - Raw data of blob
//   - options - ContainerRegistryBlobClientUploadChunkWithBinaryOptions contains the optional parameters for the ContainerRegistryBlobClient.UploadChunkWithBinary
//     method.
func (client *ContainerRegistryBlobClient) UploadChunkWithBinary(ctx context.Context, nextLink string, value io.ReadSeekCloser, options *ContainerRegistryBlobClientUploadChunkWithBinaryOptions) (ContainerRegistryBlobClientUploadChunkWithBinaryResponse, error) {
	req, err := client.uploadChunkWithBinaryCreateRequest(ctx, nextLink, value, options)
	if err != nil {
		return ContainerRegistryBlobClientUploadChunkWithBinaryResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContainerRegistryBlobClientUploadChunkWithBinaryResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return ContainerRegistryBlobClientUploadChunkWithBinaryResponse{}, runtime.NewResponseError(resp)
	}
	return client.uploadChunkWithBinaryHandleResponse(resp)
}

// uploadChunkWithBinaryCreateRequest creates the UploadChunkWithBinary request.
func (client *ContainerRegistryBlobClient) uploadChunkWithBinaryCreateRequest(ctx context.Context, nextLink string, value io.ReadSeekCloser, options *ContainerRegistryBlobClientUploadChunkWithBinaryOptions) (*policy.Request, error) {
	urlPath := "/{nextBlobUuidLink}"
	urlPath = strings.ReplaceAll(urlPath, "{nextBlobUuidLink}", nextLink)
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, req.SetBody(value, "application/octet-stream")
}

// uploadChunkWithBinaryHandleResponse handles the UploadChunkWithBinary response.
func (client *ContainerRegistryBlobClient) uploadChunkWithBinaryHandleResponse(resp *http.Response) (ContainerRegistryBlobClientUploadChunkWithBinaryResponse, error) {
	result := ContainerRegistryBlobClientUploadChunkWithBinaryResponse{}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	if val := resp.Header.Get("Range"); val != "" {
		result.Range = &val
	}
	if val := resp.Header.Get("Docker-Upload-UUID"); val != "" {
		result.DockerUploadUUID = &val
	}
	return result, nil
}
