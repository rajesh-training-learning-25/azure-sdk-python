//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresources

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type Deployment.
func (d Deployment) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "location", d.Location)
	populate(objectMap, "properties", d.Properties)
	populate(objectMap, "tags", d.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DeploymentOperationProperties.
func (d *DeploymentOperationProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "duration":
			err = unpopulate(val, "Duration", &d.Duration)
			delete(rawMsg, key)
		case "provisioningOperation":
			err = unpopulate(val, "ProvisioningOperation", &d.ProvisioningOperation)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &d.ProvisioningState)
			delete(rawMsg, key)
		case "request":
			err = unpopulate(val, "Request", &d.Request)
			delete(rawMsg, key)
		case "response":
			err = unpopulate(val, "Response", &d.Response)
			delete(rawMsg, key)
		case "serviceRequestId":
			err = unpopulate(val, "ServiceRequestID", &d.ServiceRequestID)
			delete(rawMsg, key)
		case "statusCode":
			err = unpopulate(val, "StatusCode", &d.StatusCode)
			delete(rawMsg, key)
		case "statusMessage":
			err = unpopulate(val, "StatusMessage", &d.StatusMessage)
			delete(rawMsg, key)
		case "targetResource":
			err = unpopulate(val, "TargetResource", &d.TargetResource)
			delete(rawMsg, key)
		case "timestamp":
			err = unpopulateTimeRFC3339(val, "Timestamp", &d.Timestamp)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DeploymentPropertiesExtended.
func (d *DeploymentPropertiesExtended) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "correlationId":
			err = unpopulate(val, "CorrelationID", &d.CorrelationID)
			delete(rawMsg, key)
		case "debugSetting":
			err = unpopulate(val, "DebugSetting", &d.DebugSetting)
			delete(rawMsg, key)
		case "dependencies":
			err = unpopulate(val, "Dependencies", &d.Dependencies)
			delete(rawMsg, key)
		case "duration":
			err = unpopulate(val, "Duration", &d.Duration)
			delete(rawMsg, key)
		case "error":
			err = unpopulate(val, "Error", &d.Error)
			delete(rawMsg, key)
		case "mode":
			err = unpopulate(val, "Mode", &d.Mode)
			delete(rawMsg, key)
		case "onErrorDeployment":
			err = unpopulate(val, "OnErrorDeployment", &d.OnErrorDeployment)
			delete(rawMsg, key)
		case "outputResources":
			err = unpopulate(val, "OutputResources", &d.OutputResources)
			delete(rawMsg, key)
		case "outputs":
			err = unpopulate(val, "Outputs", &d.Outputs)
			delete(rawMsg, key)
		case "parameters":
			err = unpopulate(val, "Parameters", &d.Parameters)
			delete(rawMsg, key)
		case "parametersLink":
			err = unpopulate(val, "ParametersLink", &d.ParametersLink)
			delete(rawMsg, key)
		case "providers":
			err = unpopulate(val, "Providers", &d.Providers)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &d.ProvisioningState)
			delete(rawMsg, key)
		case "templateHash":
			err = unpopulate(val, "TemplateHash", &d.TemplateHash)
			delete(rawMsg, key)
		case "templateLink":
			err = unpopulate(val, "TemplateLink", &d.TemplateLink)
			delete(rawMsg, key)
		case "timestamp":
			err = unpopulateTimeRFC3339(val, "Timestamp", &d.Timestamp)
			delete(rawMsg, key)
		case "validatedResources":
			err = unpopulate(val, "ValidatedResources", &d.ValidatedResources)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ExportTemplateRequest.
func (e ExportTemplateRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "options", e.Options)
	populate(objectMap, "resources", e.Resources)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type GenericResource.
func (g GenericResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "extendedLocation", g.ExtendedLocation)
	populate(objectMap, "id", g.ID)
	populate(objectMap, "identity", g.Identity)
	populate(objectMap, "kind", g.Kind)
	populate(objectMap, "location", g.Location)
	populate(objectMap, "managedBy", g.ManagedBy)
	populate(objectMap, "name", g.Name)
	populate(objectMap, "plan", g.Plan)
	populate(objectMap, "properties", &g.Properties)
	populate(objectMap, "sku", g.SKU)
	populate(objectMap, "tags", g.Tags)
	populate(objectMap, "type", g.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type GenericResourceExpanded.
func (g GenericResourceExpanded) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "changedTime", g.ChangedTime)
	populateTimeRFC3339(objectMap, "createdTime", g.CreatedTime)
	populate(objectMap, "extendedLocation", g.ExtendedLocation)
	populate(objectMap, "id", g.ID)
	populate(objectMap, "identity", g.Identity)
	populate(objectMap, "kind", g.Kind)
	populate(objectMap, "location", g.Location)
	populate(objectMap, "managedBy", g.ManagedBy)
	populate(objectMap, "name", g.Name)
	populate(objectMap, "plan", g.Plan)
	populate(objectMap, "properties", &g.Properties)
	populate(objectMap, "provisioningState", g.ProvisioningState)
	populate(objectMap, "sku", g.SKU)
	populate(objectMap, "tags", g.Tags)
	populate(objectMap, "type", g.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GenericResourceExpanded.
func (g *GenericResourceExpanded) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", g, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "changedTime":
			err = unpopulateTimeRFC3339(val, "ChangedTime", &g.ChangedTime)
			delete(rawMsg, key)
		case "createdTime":
			err = unpopulateTimeRFC3339(val, "CreatedTime", &g.CreatedTime)
			delete(rawMsg, key)
		case "extendedLocation":
			err = unpopulate(val, "ExtendedLocation", &g.ExtendedLocation)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &g.ID)
			delete(rawMsg, key)
		case "identity":
			err = unpopulate(val, "Identity", &g.Identity)
			delete(rawMsg, key)
		case "kind":
			err = unpopulate(val, "Kind", &g.Kind)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, "Location", &g.Location)
			delete(rawMsg, key)
		case "managedBy":
			err = unpopulate(val, "ManagedBy", &g.ManagedBy)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &g.Name)
			delete(rawMsg, key)
		case "plan":
			err = unpopulate(val, "Plan", &g.Plan)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &g.Properties)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &g.ProvisioningState)
			delete(rawMsg, key)
		case "sku":
			err = unpopulate(val, "SKU", &g.SKU)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &g.Tags)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &g.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", g, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Identity.
func (i Identity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "principalId", i.PrincipalID)
	populate(objectMap, "tenantId", i.TenantID)
	populate(objectMap, "type", i.Type)
	populate(objectMap, "userAssignedIdentities", i.UserAssignedIdentities)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MoveInfo.
func (m MoveInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "resources", m.Resources)
	populate(objectMap, "targetResourceGroup", m.TargetResourceGroup)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "extendedLocation", r.ExtendedLocation)
	populate(objectMap, "id", r.ID)
	populate(objectMap, "location", r.Location)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "tags", r.Tags)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceGroup.
func (r ResourceGroup) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", r.ID)
	populate(objectMap, "location", r.Location)
	populate(objectMap, "managedBy", r.ManagedBy)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "properties", r.Properties)
	populate(objectMap, "tags", r.Tags)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceGroupPatchable.
func (r ResourceGroupPatchable) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "managedBy", r.ManagedBy)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "properties", r.Properties)
	populate(objectMap, "tags", r.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ScopedDeployment.
func (s ScopedDeployment) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "location", s.Location)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "tags", s.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Tags.
func (t Tags) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", t.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type TagsPatchResource.
func (t TagsPatchResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "operation", t.Operation)
	populate(objectMap, "properties", t.Properties)
	return json.Marshal(objectMap)
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v interface{}) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}