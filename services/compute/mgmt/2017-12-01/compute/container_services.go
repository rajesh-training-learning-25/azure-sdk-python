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

// ContainerServicesClient is the compute Client
type ContainerServicesClient struct {
	ManagementClient
}

// NewContainerServicesClient creates an instance of the ContainerServicesClient client.
func NewContainerServicesClient(p pipeline.Pipeline) ContainerServicesClient {
	return ContainerServicesClient{NewManagementClient(p)}
}

// CreateOrUpdate creates or updates a container service with the specified configuration of orchestrator, masters, and
// agents. This method may poll for completion. Polling can be canceled by passing the cancel channel argument. The
// channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. containerServiceName is the name of the container service in
// the specified subscription and resource group. parameters is parameters supplied to the Create or Update a Container
// Service operation.
func (client ContainerServicesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, containerServiceName string, parameters ContainerService) (*ContainerService, error) {
	if err := validate([]validation{
		{targetValue: parameters,
			constraints: []constraint{{target: "parameters.ContainerServiceProperties", name: null, rule: false,
				chain: []constraint{{target: "parameters.ContainerServiceProperties.CustomProfile", name: null, rule: false,
					chain: []constraint{{target: "parameters.ContainerServiceProperties.CustomProfile.Orchestrator", name: null, rule: true, chain: nil}}},
					{target: "parameters.ContainerServiceProperties.ServicePrincipalProfile", name: null, rule: false,
						chain: []constraint{{target: "parameters.ContainerServiceProperties.ServicePrincipalProfile.ClientID", name: null, rule: true, chain: nil},
							{target: "parameters.ContainerServiceProperties.ServicePrincipalProfile.Secret", name: null, rule: true, chain: nil},
						}},
					{target: "parameters.ContainerServiceProperties.MasterProfile", name: null, rule: true,
						chain: []constraint{{target: "parameters.ContainerServiceProperties.MasterProfile.DNSPrefix", name: null, rule: true, chain: nil}}},
					{target: "parameters.ContainerServiceProperties.AgentPoolProfiles", name: null, rule: true, chain: nil},
					{target: "parameters.ContainerServiceProperties.WindowsProfile", name: null, rule: false,
						chain: []constraint{{target: "parameters.ContainerServiceProperties.WindowsProfile.AdminUsername", name: null, rule: true,
							chain: []constraint{{target: "parameters.ContainerServiceProperties.WindowsProfile.AdminUsername", name: pattern, rule: `^[a-zA-Z0-9]+([._]?[a-zA-Z0-9]+)*$`, chain: nil}}},
							{target: "parameters.ContainerServiceProperties.WindowsProfile.AdminPassword", name: null, rule: true,
								chain: []constraint{{target: "parameters.ContainerServiceProperties.WindowsProfile.AdminPassword", name: pattern, rule: `^(?=.*[a-z])(?=.*[A-Z])(?=.*[!@#$%\^&\*\(\)])[a-zA-Z\d!@#$%\^&\*\(\)]{12,123}$`, chain: nil}}},
						}},
					{target: "parameters.ContainerServiceProperties.LinuxProfile", name: null, rule: true,
						chain: []constraint{{target: "parameters.ContainerServiceProperties.LinuxProfile.AdminUsername", name: null, rule: true,
							chain: []constraint{{target: "parameters.ContainerServiceProperties.LinuxProfile.AdminUsername", name: pattern, rule: `^[a-z][a-z0-9_-]*$`, chain: nil}}},
							{target: "parameters.ContainerServiceProperties.LinuxProfile.SSH", name: null, rule: true,
								chain: []constraint{{target: "parameters.ContainerServiceProperties.LinuxProfile.SSH.PublicKeys", name: null, rule: true, chain: nil}}},
						}},
					{target: "parameters.ContainerServiceProperties.DiagnosticsProfile", name: null, rule: false,
						chain: []constraint{{target: "parameters.ContainerServiceProperties.DiagnosticsProfile.VMDiagnostics", name: null, rule: true,
							chain: []constraint{{target: "parameters.ContainerServiceProperties.DiagnosticsProfile.VMDiagnostics.Enabled", name: null, rule: true, chain: nil}}},
						}},
				}}}}}); err != nil {
		return nil, err
	}
	req, err := client.createOrUpdatePreparer(resourceGroupName, containerServiceName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createOrUpdateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ContainerService), err
}

// createOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ContainerServicesClient) createOrUpdatePreparer(resourceGroupName string, containerServiceName string, parameters ContainerService) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices/{containerServiceName}"
	req, err := pipeline.NewRequest("PUT", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-01-31")
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(parameters)
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
func (client ContainerServicesClient) createOrUpdateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	result := &ContainerService{rawResponse: resp.Response()}
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

// Delete deletes the specified container service in the specified subscription and resource group. The operation does
// not delete other resources created as part of creating a container service, including storage accounts, VMs, and
// availability sets. All the other resources created with the container service are part of the same resource group
// and can be deleted individually. This method may poll for completion. Polling can be canceled by passing the cancel
// channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. containerServiceName is the name of the container service in
// the specified subscription and resource group.
func (client ContainerServicesClient) Delete(ctx context.Context, resourceGroupName string, containerServiceName string) (*http.Response, error) {
	req, err := client.deletePreparer(resourceGroupName, containerServiceName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.deleteResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.Response(), err
}

// deletePreparer prepares the Delete request.
func (client ContainerServicesClient) deletePreparer(resourceGroupName string, containerServiceName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices/{containerServiceName}"
	req, err := pipeline.NewRequest("DELETE", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-01-31")
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// deleteResponder handles the response to the Delete request.
func (client ContainerServicesClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent)
	if resp == nil {
		return nil, err
	}
	return resp, err
}

// Get gets the properties of the specified container service in the specified subscription and resource group. The
// operation returns the properties including state, orchestrator, number of masters and agents, and FQDNs of masters
// and agents.
//
// resourceGroupName is the name of the resource group. containerServiceName is the name of the container service in
// the specified subscription and resource group.
func (client ContainerServicesClient) Get(ctx context.Context, resourceGroupName string, containerServiceName string) (*ContainerService, error) {
	req, err := client.getPreparer(resourceGroupName, containerServiceName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ContainerService), err
}

// getPreparer prepares the Get request.
func (client ContainerServicesClient) getPreparer(resourceGroupName string, containerServiceName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices/{containerServiceName}"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-01-31")
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getResponder handles the response to the Get request.
func (client ContainerServicesClient) getResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ContainerService{rawResponse: resp.Response()}
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

// List gets a list of container services in the specified subscription. The operation returns properties of each
// container service including state, orchestrator, number of masters and agents, and FQDNs of masters and agents.
func (client ContainerServicesClient) List(ctx context.Context) (*ContainerServiceListResult, error) {
	req, err := client.listPreparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ContainerServiceListResult), err
}

// listPreparer prepares the List request.
func (client ContainerServicesClient) listPreparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/containerServices"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-01-31")
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listResponder handles the response to the List request.
func (client ContainerServicesClient) listResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ContainerServiceListResult{rawResponse: resp.Response()}
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

// ListByResourceGroup gets a list of container services in the specified subscription and resource group. The
// operation returns properties of each container service including state, orchestrator, number of masters and agents,
// and FQDNs of masters and agents.
//
// resourceGroupName is the name of the resource group.
func (client ContainerServicesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (*ContainerServiceListResult, error) {
	req, err := client.listByResourceGroupPreparer(resourceGroupName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listByResourceGroupResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ContainerServiceListResult), err
}

// listByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ContainerServicesClient) listByResourceGroupPreparer(resourceGroupName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-01-31")
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listByResourceGroupResponder handles the response to the ListByResourceGroup request.
func (client ContainerServicesClient) listByResourceGroupResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ContainerServiceListResult{rawResponse: resp.Response()}
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
