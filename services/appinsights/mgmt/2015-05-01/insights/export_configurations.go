package insights

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

// ExportConfigurationsClient is the composite Swagger for Application Insights Management Client
type ExportConfigurationsClient struct {
	ManagementClient
}

// NewExportConfigurationsClient creates an instance of the ExportConfigurationsClient client.
func NewExportConfigurationsClient(p pipeline.Pipeline) ExportConfigurationsClient {
	return ExportConfigurationsClient{NewManagementClient(p)}
}

// Create create a Continuous Export configuration of an Application Insights component.
//
// resourceGroupName is the name of the resource group. resourceName is the name of the Application Insights component
// resource. exportProperties is properties that need to be specified to create a Continuous Export configuration of a
// Application Insights component.
func (client ExportConfigurationsClient) Create(ctx context.Context, resourceGroupName string, resourceName string, exportProperties ApplicationInsightsComponentExportRequest) (*CreateResponse, error) {
	req, err := client.createPreparer(resourceGroupName, resourceName, exportProperties)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*CreateResponse), err
}

// createPreparer prepares the Create request.
func (client ExportConfigurationsClient) createPreparer(resourceGroupName string, resourceName string, exportProperties ApplicationInsightsComponentExportRequest) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/exportconfiguration"
	req, err := pipeline.NewRequest("POST", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(exportProperties)
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

// createResponder handles the response to the Create request.
func (client ExportConfigurationsClient) createResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &CreateResponse{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, &result.Value)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// Delete delete a Continuous Export configuration of an Application Insights component.
//
// resourceGroupName is the name of the resource group. resourceName is the name of the Application Insights component
// resource. exportID is the Continuous Export configuration ID. This is unique within a Application Insights
// component.
func (client ExportConfigurationsClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, exportID string) (*ApplicationInsightsComponentExportConfiguration, error) {
	req, err := client.deletePreparer(resourceGroupName, resourceName, exportID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.deleteResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ApplicationInsightsComponentExportConfiguration), err
}

// deletePreparer prepares the Delete request.
func (client ExportConfigurationsClient) deletePreparer(resourceGroupName string, resourceName string, exportID string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/exportconfiguration/{exportId}"
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
func (client ExportConfigurationsClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ApplicationInsightsComponentExportConfiguration{rawResponse: resp.Response()}
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

// Get get the Continuous Export configuration for this export id.
//
// resourceGroupName is the name of the resource group. resourceName is the name of the Application Insights component
// resource. exportID is the Continuous Export configuration ID. This is unique within a Application Insights
// component.
func (client ExportConfigurationsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, exportID string) (*ApplicationInsightsComponentExportConfiguration, error) {
	req, err := client.getPreparer(resourceGroupName, resourceName, exportID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ApplicationInsightsComponentExportConfiguration), err
}

// getPreparer prepares the Get request.
func (client ExportConfigurationsClient) getPreparer(resourceGroupName string, resourceName string, exportID string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/exportconfiguration/{exportId}"
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
func (client ExportConfigurationsClient) getResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ApplicationInsightsComponentExportConfiguration{rawResponse: resp.Response()}
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

// List gets a list of Continuous Export configuration of an Application Insights component.
//
// resourceGroupName is the name of the resource group. resourceName is the name of the Application Insights component
// resource.
func (client ExportConfigurationsClient) List(ctx context.Context, resourceGroupName string, resourceName string) (*ListResponse, error) {
	req, err := client.listPreparer(resourceGroupName, resourceName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ListResponse), err
}

// listPreparer prepares the List request.
func (client ExportConfigurationsClient) listPreparer(resourceGroupName string, resourceName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/exportconfiguration"
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
func (client ExportConfigurationsClient) listResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ListResponse{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, &result.Value)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// Update update the Continuous Export configuration for this export id.
//
// resourceGroupName is the name of the resource group. resourceName is the name of the Application Insights component
// resource. exportID is the Continuous Export configuration ID. This is unique within a Application Insights
// component. exportProperties is properties that need to be specified to update the Continuous Export configuration.
func (client ExportConfigurationsClient) Update(ctx context.Context, resourceGroupName string, resourceName string, exportID string, exportProperties ApplicationInsightsComponentExportRequest) (*ApplicationInsightsComponentExportConfiguration, error) {
	req, err := client.updatePreparer(resourceGroupName, resourceName, exportID, exportProperties)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.updateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ApplicationInsightsComponentExportConfiguration), err
}

// updatePreparer prepares the Update request.
func (client ExportConfigurationsClient) updatePreparer(resourceGroupName string, resourceName string, exportID string, exportProperties ApplicationInsightsComponentExportRequest) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/exportconfiguration/{exportId}"
	req, err := pipeline.NewRequest("PUT", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(exportProperties)
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
func (client ExportConfigurationsClient) updateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ApplicationInsightsComponentExportConfiguration{rawResponse: resp.Response()}
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
