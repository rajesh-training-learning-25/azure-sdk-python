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

// FavoriteClient is the composite Swagger for Application Insights Management Client
type FavoriteClient struct {
	ManagementClient
}

// NewFavoriteClient creates an instance of the FavoriteClient client.
func NewFavoriteClient(p pipeline.Pipeline) FavoriteClient {
	return FavoriteClient{NewManagementClient(p)}
}

// Add adds a new favorites to an Application Insights component.
//
// resourceGroupName is the name of the resource group. resourceName is the name of the Application Insights component
// resource. favoriteID is the Id of a specific favorite defined in the Application Insights component
// favoriteProperties is properties that need to be specified to create a new favorite and add it to an Application
// Insights component.
func (client FavoriteClient) Add(ctx context.Context, resourceGroupName string, resourceName string, favoriteID string, favoriteProperties ApplicationInsightsComponentFavorite) (*ApplicationInsightsComponentFavorite, error) {
	req, err := client.addPreparer(resourceGroupName, resourceName, favoriteID, favoriteProperties)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.addResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ApplicationInsightsComponentFavorite), err
}

// addPreparer prepares the Add request.
func (client FavoriteClient) addPreparer(resourceGroupName string, resourceName string, favoriteID string, favoriteProperties ApplicationInsightsComponentFavorite) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/favorites/{favoriteId}"
	req, err := pipeline.NewRequest("PUT", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(favoriteProperties)
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

// addResponder handles the response to the Add request.
func (client FavoriteClient) addResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ApplicationInsightsComponentFavorite{rawResponse: resp.Response()}
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

// Delete remove a favorite that is associated to an Application Insights component.
//
// resourceGroupName is the name of the resource group. resourceName is the name of the Application Insights component
// resource. favoriteID is the Id of a specific favorite defined in the Application Insights component
func (client FavoriteClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, favoriteID string) (*http.Response, error) {
	req, err := client.deletePreparer(resourceGroupName, resourceName, favoriteID)
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
func (client FavoriteClient) deletePreparer(resourceGroupName string, resourceName string, favoriteID string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/favorites/{favoriteId}"
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
func (client FavoriteClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	return resp, err
}

// Get get a single favorite by its FavoriteId, defined within an Application Insights component.
//
// resourceGroupName is the name of the resource group. resourceName is the name of the Application Insights component
// resource. favoriteID is the Id of a specific favorite defined in the Application Insights component
func (client FavoriteClient) Get(ctx context.Context, resourceGroupName string, resourceName string, favoriteID string) (*ApplicationInsightsComponentFavorite, error) {
	req, err := client.getPreparer(resourceGroupName, resourceName, favoriteID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ApplicationInsightsComponentFavorite), err
}

// getPreparer prepares the Get request.
func (client FavoriteClient) getPreparer(resourceGroupName string, resourceName string, favoriteID string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/favorites/{favoriteId}"
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
func (client FavoriteClient) getResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ApplicationInsightsComponentFavorite{rawResponse: resp.Response()}
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

// Update updates a favorite that has already been added to an Application Insights component.
//
// resourceGroupName is the name of the resource group. resourceName is the name of the Application Insights component
// resource. favoriteID is the Id of a specific favorite defined in the Application Insights component
// favoriteProperties is properties that need to be specified to update the existing favorite.
func (client FavoriteClient) Update(ctx context.Context, resourceGroupName string, resourceName string, favoriteID string, favoriteProperties ApplicationInsightsComponentFavorite) (*ApplicationInsightsComponentFavorite, error) {
	req, err := client.updatePreparer(resourceGroupName, resourceName, favoriteID, favoriteProperties)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.updateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ApplicationInsightsComponentFavorite), err
}

// updatePreparer prepares the Update request.
func (client FavoriteClient) updatePreparer(resourceGroupName string, resourceName string, favoriteID string, favoriteProperties ApplicationInsightsComponentFavorite) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/favorites/{favoriteId}"
	req, err := pipeline.NewRequest("PATCH", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(favoriteProperties)
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
func (client FavoriteClient) updateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ApplicationInsightsComponentFavorite{rawResponse: resp.Response()}
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
