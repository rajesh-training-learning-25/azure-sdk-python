// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// NatGatewaysOperations contains the methods for the NatGateways group.
type NatGatewaysOperations interface {
	// BeginCreateOrUpdate - Creates or updates a nat gateway.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, natGatewayName string, parameters NatGateway) (*NatGatewayPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (NatGatewayPoller, error)
	// BeginDelete - Deletes the specified nat gateway.
	BeginDelete(ctx context.Context, resourceGroupName string, natGatewayName string) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the specified nat gateway in a specified resource group.
	Get(ctx context.Context, resourceGroupName string, natGatewayName string, natGatewaysGetOptions *NatGatewaysGetOptions) (*NatGatewayResponse, error)
	// List - Gets all nat gateways in a resource group.
	List(resourceGroupName string) (NatGatewayListResultPager, error)
	// ListAll - Gets all the Nat Gateways in a subscription.
	ListAll() (NatGatewayListResultPager, error)
	// UpdateTags - Updates nat gateway tags.
	UpdateTags(ctx context.Context, resourceGroupName string, natGatewayName string, parameters TagsObject) (*NatGatewayResponse, error)
}

// natGatewaysOperations implements the NatGatewaysOperations interface.
type natGatewaysOperations struct {
	*Client
	subscriptionID string
}

// CreateOrUpdate - Creates or updates a nat gateway.
func (client *natGatewaysOperations) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, natGatewayName string, parameters NatGateway) (*NatGatewayPollerResponse, error) {
	req, err := client.createOrUpdateCreateRequest(resourceGroupName, natGatewayName, parameters)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.createOrUpdateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("natGatewaysOperations.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &natGatewayPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*NatGatewayResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *natGatewaysOperations) ResumeCreateOrUpdate(token string) (NatGatewayPoller, error) {
	pt, err := resumePollingTracker("natGatewaysOperations.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &natGatewayPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *natGatewaysOperations) createOrUpdateCreateRequest(resourceGroupName string, natGatewayName string, parameters NatGateway) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{natGatewayName}", url.PathEscape(natGatewayName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *natGatewaysOperations) createOrUpdateHandleResponse(resp *azcore.Response) (*NatGatewayPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return &NatGatewayPollerResponse{RawResponse: resp.Response}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *natGatewaysOperations) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes the specified nat gateway.
func (client *natGatewaysOperations) BeginDelete(ctx context.Context, resourceGroupName string, natGatewayName string) (*HTTPPollerResponse, error) {
	req, err := client.deleteCreateRequest(resourceGroupName, natGatewayName)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.deleteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("natGatewaysOperations.Delete", "location", resp, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *natGatewaysOperations) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := resumePollingTracker("natGatewaysOperations.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *natGatewaysOperations) deleteCreateRequest(resourceGroupName string, natGatewayName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{natGatewayName}", url.PathEscape(natGatewayName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodDelete, *u)
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *natGatewaysOperations) deleteHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// deleteHandleError handles the Delete error response.
func (client *natGatewaysOperations) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Gets the specified nat gateway in a specified resource group.
func (client *natGatewaysOperations) Get(ctx context.Context, resourceGroupName string, natGatewayName string, natGatewaysGetOptions *NatGatewaysGetOptions) (*NatGatewayResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, natGatewayName, natGatewaysGetOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client *natGatewaysOperations) getCreateRequest(resourceGroupName string, natGatewayName string, natGatewaysGetOptions *NatGatewaysGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{natGatewayName}", url.PathEscape(natGatewayName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	if natGatewaysGetOptions != nil && natGatewaysGetOptions.Expand != nil {
		query.Set("$expand", *natGatewaysGetOptions.Expand)
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *natGatewaysOperations) getHandleResponse(resp *azcore.Response) (*NatGatewayResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := NatGatewayResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NatGateway)
}

// getHandleError handles the Get error response.
func (client *natGatewaysOperations) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Gets all nat gateways in a resource group.
func (client *natGatewaysOperations) List(resourceGroupName string) (NatGatewayListResultPager, error) {
	req, err := client.listCreateRequest(resourceGroupName)
	if err != nil {
		return nil, err
	}
	return &natGatewayListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *NatGatewayListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.NatGatewayListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.NatGatewayListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *natGatewaysOperations) listCreateRequest(resourceGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listHandleResponse handles the List response.
func (client *natGatewaysOperations) listHandleResponse(resp *azcore.Response) (*NatGatewayListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := NatGatewayListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NatGatewayListResult)
}

// listHandleError handles the List error response.
func (client *natGatewaysOperations) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListAll - Gets all the Nat Gateways in a subscription.
func (client *natGatewaysOperations) ListAll() (NatGatewayListResultPager, error) {
	req, err := client.listAllCreateRequest()
	if err != nil {
		return nil, err
	}
	return &natGatewayListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listAllHandleResponse,
		advancer: func(resp *NatGatewayListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.NatGatewayListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.NatGatewayListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listAllCreateRequest creates the ListAll request.
func (client *natGatewaysOperations) listAllCreateRequest() (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/natGateways"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *natGatewaysOperations) listAllHandleResponse(resp *azcore.Response) (*NatGatewayListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listAllHandleError(resp)
	}
	result := NatGatewayListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NatGatewayListResult)
}

// listAllHandleError handles the ListAll error response.
func (client *natGatewaysOperations) listAllHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// UpdateTags - Updates nat gateway tags.
func (client *natGatewaysOperations) UpdateTags(ctx context.Context, resourceGroupName string, natGatewayName string, parameters TagsObject) (*NatGatewayResponse, error) {
	req, err := client.updateTagsCreateRequest(resourceGroupName, natGatewayName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.updateTagsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *natGatewaysOperations) updateTagsCreateRequest(resourceGroupName string, natGatewayName string, parameters TagsObject) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{natGatewayName}", url.PathEscape(natGatewayName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPatch, *u)
	return req, req.MarshalAsJSON(parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *natGatewaysOperations) updateTagsHandleResponse(resp *azcore.Response) (*NatGatewayResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.updateTagsHandleError(resp)
	}
	result := NatGatewayResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NatGateway)
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *natGatewaysOperations) updateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
