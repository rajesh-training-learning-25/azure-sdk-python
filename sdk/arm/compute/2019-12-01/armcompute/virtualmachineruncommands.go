// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// VirtualMachineRunCommandsOperations contains the methods for the VirtualMachineRunCommands group.
type VirtualMachineRunCommandsOperations interface {
	// Get - Gets specific run command for a subscription in a location.
	Get(ctx context.Context, location string, commandId string) (*RunCommandDocumentResponse, error)
	// List - Lists all available run commands for a subscription in a location.
	List(location string) (RunCommandListResultPager, error)
}

// virtualMachineRunCommandsOperations implements the VirtualMachineRunCommandsOperations interface.
type virtualMachineRunCommandsOperations struct {
	*Client
	subscriptionID string
}

// Get - Gets specific run command for a subscription in a location.
func (client *virtualMachineRunCommandsOperations) Get(ctx context.Context, location string, commandId string) (*RunCommandDocumentResponse, error) {
	req, err := client.getCreateRequest(location, commandId)
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
func (client *virtualMachineRunCommandsOperations) getCreateRequest(location string, commandId string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/runCommands/{commandId}"
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	urlPath = strings.ReplaceAll(urlPath, "{commandId}", url.PathEscape(commandId))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-12-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *virtualMachineRunCommandsOperations) getHandleResponse(resp *azcore.Response) (*RunCommandDocumentResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := RunCommandDocumentResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.RunCommandDocument)
}

// getHandleError handles the Get error response.
func (client *virtualMachineRunCommandsOperations) getHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}

// List - Lists all available run commands for a subscription in a location.
func (client *virtualMachineRunCommandsOperations) List(location string) (RunCommandListResultPager, error) {
	req, err := client.listCreateRequest(location)
	if err != nil {
		return nil, err
	}
	return &runCommandListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *RunCommandListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.RunCommandListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.RunCommandListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *virtualMachineRunCommandsOperations) listCreateRequest(location string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/runCommands"
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-12-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listHandleResponse handles the List response.
func (client *virtualMachineRunCommandsOperations) listHandleResponse(resp *azcore.Response) (*RunCommandListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := RunCommandListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.RunCommandListResult)
}

// listHandleError handles the List error response.
func (client *virtualMachineRunCommandsOperations) listHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}
