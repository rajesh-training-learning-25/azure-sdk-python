//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azblob

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// containerClientListBlobFlatSegmentPager provides operations for iterating over paged responses.
type containerClientListBlobFlatSegmentPager struct {
	client    *containerClient
	current   containerClientListBlobFlatSegmentResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, containerClientListBlobFlatSegmentResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *containerClientListBlobFlatSegmentPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListBlobsFlatSegmentResponse.NextMarker == nil || len(*p.current.ListBlobsFlatSegmentResponse.NextMarker) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *containerClientListBlobFlatSegmentPager) NextPage(ctx context.Context) (containerClientListBlobFlatSegmentResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return containerClientListBlobFlatSegmentResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return containerClientListBlobFlatSegmentResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return containerClientListBlobFlatSegmentResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return containerClientListBlobFlatSegmentResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listBlobFlatSegmentHandleResponse(resp)
	if err != nil {
		return containerClientListBlobFlatSegmentResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// containerClientListBlobHierarchySegmentPager provides operations for iterating over paged responses.
type containerClientListBlobHierarchySegmentPager struct {
	client    *containerClient
	current   containerClientListBlobHierarchySegmentResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, containerClientListBlobHierarchySegmentResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *containerClientListBlobHierarchySegmentPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListBlobsHierarchySegmentResponse.NextMarker == nil || len(*p.current.ListBlobsHierarchySegmentResponse.NextMarker) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *containerClientListBlobHierarchySegmentPager) NextPage(ctx context.Context) (containerClientListBlobHierarchySegmentResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return containerClientListBlobHierarchySegmentResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return containerClientListBlobHierarchySegmentResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return containerClientListBlobHierarchySegmentResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return containerClientListBlobHierarchySegmentResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listBlobHierarchySegmentHandleResponse(resp)
	if err != nil {
		return containerClientListBlobHierarchySegmentResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// pageBlobClientGetPageRangesDiffPager provides operations for iterating over paged responses.
type pageBlobClientGetPageRangesDiffPager struct {
	client    *pageBlobClient
	current   pageBlobClientGetPageRangesDiffResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, pageBlobClientGetPageRangesDiffResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *pageBlobClientGetPageRangesDiffPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PageList.NextMarker == nil || len(*p.current.PageList.NextMarker) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *pageBlobClientGetPageRangesDiffPager) NextPage(ctx context.Context) (pageBlobClientGetPageRangesDiffResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return pageBlobClientGetPageRangesDiffResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return pageBlobClientGetPageRangesDiffResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return pageBlobClientGetPageRangesDiffResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return pageBlobClientGetPageRangesDiffResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.getPageRangesDiffHandleResponse(resp)
	if err != nil {
		return pageBlobClientGetPageRangesDiffResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// pageBlobClientGetPageRangesPager provides operations for iterating over paged responses.
type pageBlobClientGetPageRangesPager struct {
	client    *pageBlobClient
	current   pageBlobClientGetPageRangesResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, pageBlobClientGetPageRangesResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *pageBlobClientGetPageRangesPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PageList.NextMarker == nil || len(*p.current.PageList.NextMarker) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *pageBlobClientGetPageRangesPager) NextPage(ctx context.Context) (pageBlobClientGetPageRangesResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return pageBlobClientGetPageRangesResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return pageBlobClientGetPageRangesResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return pageBlobClientGetPageRangesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return pageBlobClientGetPageRangesResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.getPageRangesHandleResponse(resp)
	if err != nil {
		return pageBlobClientGetPageRangesResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// serviceClientListContainersSegmentPager provides operations for iterating over paged responses.
type serviceClientListContainersSegmentPager struct {
	client    *serviceClient
	current   ServiceClientListContainersSegmentResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ServiceClientListContainersSegmentResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *serviceClientListContainersSegmentPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListContainersSegmentResponse.NextMarker == nil || len(*p.current.ListContainersSegmentResponse.NextMarker) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *serviceClientListContainersSegmentPager) NextPage(ctx context.Context) (ServiceClientListContainersSegmentResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return ServiceClientListContainersSegmentResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return ServiceClientListContainersSegmentResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return ServiceClientListContainersSegmentResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return ServiceClientListContainersSegmentResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listContainersSegmentHandleResponse(resp)
	if err != nil {
		return ServiceClientListContainersSegmentResponse{}, err
	}
	p.current = result
	return p.current, nil
}
