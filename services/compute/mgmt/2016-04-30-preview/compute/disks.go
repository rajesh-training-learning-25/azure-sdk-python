package compute

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
	"bytes"
	"context"
	"encoding/json"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"io/ioutil"
	"net/http"
)

// DisksClient is the compute Client
type DisksClient struct {
	ManagementClient
}

// NewDisksClient creates an instance of the DisksClient client.
func NewDisksClient(p pipeline.Pipeline) DisksClient {
	return DisksClient{NewManagementClient(p)}
}

// CreateOrUpdate creates or updates a disk. This method may poll for completion. Polling can be canceled by passing
// the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. diskName is the name of the managed disk that is being created.
// The name can't be changed after the disk is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The
// maximum name length is 80 characters. disk is disk object supplied in the body of the Put disk operation.
func (client DisksClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, diskName string, disk Disk) (*Disk, error) {
	if err := validate([]validation{
		{targetValue: disk,
			constraints: []constraint{{target: "disk.DiskProperties", name: null, rule: false,
				chain: []constraint{{target: "disk.DiskProperties.CreationData", name: null, rule: true,
					chain: []constraint{{target: "disk.DiskProperties.CreationData.ImageReference", name: null, rule: false,
						chain: []constraint{{target: "disk.DiskProperties.CreationData.ImageReference.ID", name: null, rule: true, chain: nil}}},
					}},
					{target: "disk.DiskProperties.EncryptionSettings", name: null, rule: false,
						chain: []constraint{{target: "disk.DiskProperties.EncryptionSettings.DiskEncryptionKey", name: null, rule: false,
							chain: []constraint{{target: "disk.DiskProperties.EncryptionSettings.DiskEncryptionKey.SourceVault", name: null, rule: true, chain: nil},
								{target: "disk.DiskProperties.EncryptionSettings.DiskEncryptionKey.SecretURL", name: null, rule: true, chain: nil},
							}},
							{target: "disk.DiskProperties.EncryptionSettings.KeyEncryptionKey", name: null, rule: false,
								chain: []constraint{{target: "disk.DiskProperties.EncryptionSettings.KeyEncryptionKey.SourceVault", name: null, rule: true, chain: nil},
									{target: "disk.DiskProperties.EncryptionSettings.KeyEncryptionKey.KeyURL", name: null, rule: true, chain: nil},
								}},
						}},
				}}}}}); err != nil {
		return nil, err
	}
	req, err := client.createOrUpdatePreparer(resourceGroupName, diskName, disk)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createOrUpdateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*Disk), err
}

// createOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DisksClient) createOrUpdatePreparer(resourceGroupName string, diskName string, disk Disk) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}"
	req, err := pipeline.NewRequest("PUT", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(disk)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// createOrUpdateResponder handles the response to the CreateOrUpdate request.
func (client DisksClient) createOrUpdateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	result := &Disk{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// Delete deletes a disk. This method may poll for completion. Polling can be canceled by passing the cancel channel
// argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. diskName is the name of the managed disk that is being created.
// The name can't be changed after the disk is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The
// maximum name length is 80 characters.
func (client DisksClient) Delete(ctx context.Context, resourceGroupName string, diskName string) (*OperationStatusResponse, error) {
	req, err := client.deletePreparer(resourceGroupName, diskName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.deleteResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*OperationStatusResponse), err
}

// deletePreparer prepares the Delete request.
func (client DisksClient) deletePreparer(resourceGroupName string, diskName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}"
	req, err := pipeline.NewRequest("DELETE", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// deleteResponder handles the response to the Delete request.
func (client DisksClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent)
	if resp == nil {
		return nil, err
	}
	result := &OperationStatusResponse{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// Get gets information about a disk.
//
// resourceGroupName is the name of the resource group. diskName is the name of the managed disk that is being created.
// The name can't be changed after the disk is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The
// maximum name length is 80 characters.
func (client DisksClient) Get(ctx context.Context, resourceGroupName string, diskName string) (*Disk, error) {
	req, err := client.getPreparer(resourceGroupName, diskName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*Disk), err
}

// getPreparer prepares the Get request.
func (client DisksClient) getPreparer(resourceGroupName string, diskName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getResponder handles the response to the Get request.
func (client DisksClient) getResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &Disk{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// GrantAccess grants access to a disk. This method may poll for completion. Polling can be canceled by passing the
// cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. diskName is the name of the managed disk that is being created.
// The name can't be changed after the disk is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The
// maximum name length is 80 characters. grantAccessData is access data object supplied in the body of the get disk
// access operation.
func (client DisksClient) GrantAccess(ctx context.Context, resourceGroupName string, diskName string, grantAccessData GrantAccessData) (*AccessURI, error) {
	if err := validate([]validation{
		{targetValue: grantAccessData,
			constraints: []constraint{{target: "grantAccessData.DurationInSeconds", name: null, rule: true, chain: nil}}}}); err != nil {
		return nil, err
	}
	req, err := client.grantAccessPreparer(resourceGroupName, diskName, grantAccessData)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.grantAccessResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*AccessURI), err
}

// grantAccessPreparer prepares the GrantAccess request.
func (client DisksClient) grantAccessPreparer(resourceGroupName string, diskName string, grantAccessData GrantAccessData) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}/beginGetAccess"
	req, err := pipeline.NewRequest("POST", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(grantAccessData)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// grantAccessResponder handles the response to the GrantAccess request.
func (client DisksClient) grantAccessResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	result := &AccessURI{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// List lists all the disks under a subscription.
func (client DisksClient) List(ctx context.Context) (*DiskList, error) {
	req, err := client.listPreparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*DiskList), err
}

// listPreparer prepares the List request.
func (client DisksClient) listPreparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/disks"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listResponder handles the response to the List request.
func (client DisksClient) listResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &DiskList{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// ListByResourceGroup lists all the disks under a resource group.
//
// resourceGroupName is the name of the resource group.
func (client DisksClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (*DiskList, error) {
	req, err := client.listByResourceGroupPreparer(resourceGroupName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listByResourceGroupResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*DiskList), err
}

// listByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client DisksClient) listByResourceGroupPreparer(resourceGroupName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listByResourceGroupResponder handles the response to the ListByResourceGroup request.
func (client DisksClient) listByResourceGroupResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &DiskList{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// RevokeAccess revokes access to a disk. This method may poll for completion. Polling can be canceled by passing the
// cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. diskName is the name of the managed disk that is being created.
// The name can't be changed after the disk is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The
// maximum name length is 80 characters.
func (client DisksClient) RevokeAccess(ctx context.Context, resourceGroupName string, diskName string) (*OperationStatusResponse, error) {
	req, err := client.revokeAccessPreparer(resourceGroupName, diskName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.revokeAccessResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*OperationStatusResponse), err
}

// revokeAccessPreparer prepares the RevokeAccess request.
func (client DisksClient) revokeAccessPreparer(resourceGroupName string, diskName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}/endGetAccess"
	req, err := pipeline.NewRequest("POST", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// revokeAccessResponder handles the response to the RevokeAccess request.
func (client DisksClient) revokeAccessResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	result := &OperationStatusResponse{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// Update updates (patches) a disk. This method may poll for completion. Polling can be canceled by passing the cancel
// channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. diskName is the name of the managed disk that is being created.
// The name can't be changed after the disk is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The
// maximum name length is 80 characters. disk is disk object supplied in the body of the Patch disk operation.
func (client DisksClient) Update(ctx context.Context, resourceGroupName string, diskName string, disk DiskUpdate) (*Disk, error) {
	req, err := client.updatePreparer(resourceGroupName, diskName, disk)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.updateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*Disk), err
}

// updatePreparer prepares the Update request.
func (client DisksClient) updatePreparer(resourceGroupName string, diskName string, disk DiskUpdate) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}"
	req, err := pipeline.NewRequest("PATCH", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(disk)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// updateResponder handles the response to the Update request.
func (client DisksClient) updateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	result := &Disk{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}
