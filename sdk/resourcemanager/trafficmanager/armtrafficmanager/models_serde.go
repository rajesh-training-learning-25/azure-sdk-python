//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armtrafficmanager

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type CheckTrafficManagerRelativeDNSNameAvailabilityParameters.
func (c CheckTrafficManagerRelativeDNSNameAvailabilityParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "name", c.Name)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CheckTrafficManagerRelativeDNSNameAvailabilityParameters.
func (c *CheckTrafficManagerRelativeDNSNameAvailabilityParameters) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
			err = unpopulate(val, "Name", &c.Name)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &c.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DNSConfig.
func (d DNSConfig) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "fqdn", d.Fqdn)
	populate(objectMap, "relativeName", d.RelativeName)
	populate(objectMap, "ttl", d.TTL)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DNSConfig.
func (d *DNSConfig) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "fqdn":
			err = unpopulate(val, "Fqdn", &d.Fqdn)
			delete(rawMsg, key)
		case "relativeName":
			err = unpopulate(val, "RelativeName", &d.RelativeName)
			delete(rawMsg, key)
		case "ttl":
			err = unpopulate(val, "TTL", &d.TTL)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DeleteOperationResult.
func (d DeleteOperationResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "boolean", d.OperationResult)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DeleteOperationResult.
func (d *DeleteOperationResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "boolean":
			err = unpopulate(val, "OperationResult", &d.OperationResult)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Endpoint.
func (e Endpoint) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", e.ID)
	populate(objectMap, "name", e.Name)
	populate(objectMap, "properties", e.Properties)
	populate(objectMap, "type", e.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Endpoint.
func (e *Endpoint) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &e.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &e.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &e.Properties)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &e.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type EndpointProperties.
func (e EndpointProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "alwaysServe", e.AlwaysServe)
	populate(objectMap, "customHeaders", e.CustomHeaders)
	populate(objectMap, "endpointLocation", e.EndpointLocation)
	populate(objectMap, "endpointMonitorStatus", e.EndpointMonitorStatus)
	populate(objectMap, "endpointStatus", e.EndpointStatus)
	populate(objectMap, "geoMapping", e.GeoMapping)
	populate(objectMap, "minChildEndpoints", e.MinChildEndpoints)
	populate(objectMap, "minChildEndpointsIPv4", e.MinChildEndpointsIPv4)
	populate(objectMap, "minChildEndpointsIPv6", e.MinChildEndpointsIPv6)
	populate(objectMap, "priority", e.Priority)
	populate(objectMap, "subnets", e.Subnets)
	populate(objectMap, "target", e.Target)
	populate(objectMap, "targetResourceId", e.TargetResourceID)
	populate(objectMap, "weight", e.Weight)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EndpointProperties.
func (e *EndpointProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "alwaysServe":
			err = unpopulate(val, "AlwaysServe", &e.AlwaysServe)
			delete(rawMsg, key)
		case "customHeaders":
			err = unpopulate(val, "CustomHeaders", &e.CustomHeaders)
			delete(rawMsg, key)
		case "endpointLocation":
			err = unpopulate(val, "EndpointLocation", &e.EndpointLocation)
			delete(rawMsg, key)
		case "endpointMonitorStatus":
			err = unpopulate(val, "EndpointMonitorStatus", &e.EndpointMonitorStatus)
			delete(rawMsg, key)
		case "endpointStatus":
			err = unpopulate(val, "EndpointStatus", &e.EndpointStatus)
			delete(rawMsg, key)
		case "geoMapping":
			err = unpopulate(val, "GeoMapping", &e.GeoMapping)
			delete(rawMsg, key)
		case "minChildEndpoints":
			err = unpopulate(val, "MinChildEndpoints", &e.MinChildEndpoints)
			delete(rawMsg, key)
		case "minChildEndpointsIPv4":
			err = unpopulate(val, "MinChildEndpointsIPv4", &e.MinChildEndpointsIPv4)
			delete(rawMsg, key)
		case "minChildEndpointsIPv6":
			err = unpopulate(val, "MinChildEndpointsIPv6", &e.MinChildEndpointsIPv6)
			delete(rawMsg, key)
		case "priority":
			err = unpopulate(val, "Priority", &e.Priority)
			delete(rawMsg, key)
		case "subnets":
			err = unpopulate(val, "Subnets", &e.Subnets)
			delete(rawMsg, key)
		case "target":
			err = unpopulate(val, "Target", &e.Target)
			delete(rawMsg, key)
		case "targetResourceId":
			err = unpopulate(val, "TargetResourceID", &e.TargetResourceID)
			delete(rawMsg, key)
		case "weight":
			err = unpopulate(val, "Weight", &e.Weight)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type EndpointPropertiesCustomHeadersItem.
func (e EndpointPropertiesCustomHeadersItem) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "name", e.Name)
	populate(objectMap, "value", e.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EndpointPropertiesCustomHeadersItem.
func (e *EndpointPropertiesCustomHeadersItem) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
			err = unpopulate(val, "Name", &e.Name)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &e.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type EndpointPropertiesSubnetsItem.
func (e EndpointPropertiesSubnetsItem) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "first", e.First)
	populate(objectMap, "last", e.Last)
	populate(objectMap, "scope", e.Scope)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EndpointPropertiesSubnetsItem.
func (e *EndpointPropertiesSubnetsItem) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "first":
			err = unpopulate(val, "First", &e.First)
			delete(rawMsg, key)
		case "last":
			err = unpopulate(val, "Last", &e.Last)
			delete(rawMsg, key)
		case "scope":
			err = unpopulate(val, "Scope", &e.Scope)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type GeographicHierarchy.
func (g GeographicHierarchy) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", g.ID)
	populate(objectMap, "name", g.Name)
	populate(objectMap, "properties", g.Properties)
	populate(objectMap, "type", g.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GeographicHierarchy.
func (g *GeographicHierarchy) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", g, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &g.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &g.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &g.Properties)
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

// MarshalJSON implements the json.Marshaller interface for type GeographicHierarchyProperties.
func (g GeographicHierarchyProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "geographicHierarchy", g.GeographicHierarchy)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GeographicHierarchyProperties.
func (g *GeographicHierarchyProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", g, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "geographicHierarchy":
			err = unpopulate(val, "GeographicHierarchy", &g.GeographicHierarchy)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", g, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type HeatMapEndpoint.
func (h HeatMapEndpoint) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "endpointId", h.EndpointID)
	populate(objectMap, "resourceId", h.ResourceID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type HeatMapEndpoint.
func (h *HeatMapEndpoint) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", h, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "endpointId":
			err = unpopulate(val, "EndpointID", &h.EndpointID)
			delete(rawMsg, key)
		case "resourceId":
			err = unpopulate(val, "ResourceID", &h.ResourceID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", h, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type HeatMapModel.
func (h HeatMapModel) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", h.ID)
	populate(objectMap, "name", h.Name)
	populate(objectMap, "properties", h.Properties)
	populate(objectMap, "type", h.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type HeatMapModel.
func (h *HeatMapModel) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", h, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &h.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &h.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &h.Properties)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &h.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", h, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type HeatMapProperties.
func (h HeatMapProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "endTime", h.EndTime)
	populate(objectMap, "endpoints", h.Endpoints)
	populateTimeRFC3339(objectMap, "startTime", h.StartTime)
	populate(objectMap, "trafficFlows", h.TrafficFlows)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type HeatMapProperties.
func (h *HeatMapProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", h, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "endTime":
			err = unpopulateTimeRFC3339(val, "EndTime", &h.EndTime)
			delete(rawMsg, key)
		case "endpoints":
			err = unpopulate(val, "Endpoints", &h.Endpoints)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeRFC3339(val, "StartTime", &h.StartTime)
			delete(rawMsg, key)
		case "trafficFlows":
			err = unpopulate(val, "TrafficFlows", &h.TrafficFlows)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", h, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MonitorConfig.
func (m MonitorConfig) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "customHeaders", m.CustomHeaders)
	populate(objectMap, "expectedStatusCodeRanges", m.ExpectedStatusCodeRanges)
	populate(objectMap, "intervalInSeconds", m.IntervalInSeconds)
	populate(objectMap, "path", m.Path)
	populate(objectMap, "port", m.Port)
	populate(objectMap, "profileMonitorStatus", m.ProfileMonitorStatus)
	populate(objectMap, "protocol", m.Protocol)
	populate(objectMap, "timeoutInSeconds", m.TimeoutInSeconds)
	populate(objectMap, "toleratedNumberOfFailures", m.ToleratedNumberOfFailures)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MonitorConfig.
func (m *MonitorConfig) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "customHeaders":
			err = unpopulate(val, "CustomHeaders", &m.CustomHeaders)
			delete(rawMsg, key)
		case "expectedStatusCodeRanges":
			err = unpopulate(val, "ExpectedStatusCodeRanges", &m.ExpectedStatusCodeRanges)
			delete(rawMsg, key)
		case "intervalInSeconds":
			err = unpopulate(val, "IntervalInSeconds", &m.IntervalInSeconds)
			delete(rawMsg, key)
		case "path":
			err = unpopulate(val, "Path", &m.Path)
			delete(rawMsg, key)
		case "port":
			err = unpopulate(val, "Port", &m.Port)
			delete(rawMsg, key)
		case "profileMonitorStatus":
			err = unpopulate(val, "ProfileMonitorStatus", &m.ProfileMonitorStatus)
			delete(rawMsg, key)
		case "protocol":
			err = unpopulate(val, "Protocol", &m.Protocol)
			delete(rawMsg, key)
		case "timeoutInSeconds":
			err = unpopulate(val, "TimeoutInSeconds", &m.TimeoutInSeconds)
			delete(rawMsg, key)
		case "toleratedNumberOfFailures":
			err = unpopulate(val, "ToleratedNumberOfFailures", &m.ToleratedNumberOfFailures)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MonitorConfigCustomHeadersItem.
func (m MonitorConfigCustomHeadersItem) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "name", m.Name)
	populate(objectMap, "value", m.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MonitorConfigCustomHeadersItem.
func (m *MonitorConfigCustomHeadersItem) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
			err = unpopulate(val, "Name", &m.Name)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &m.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MonitorConfigExpectedStatusCodeRangesItem.
func (m MonitorConfigExpectedStatusCodeRangesItem) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "max", m.Max)
	populate(objectMap, "min", m.Min)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MonitorConfigExpectedStatusCodeRangesItem.
func (m *MonitorConfigExpectedStatusCodeRangesItem) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "max":
			err = unpopulate(val, "Max", &m.Max)
			delete(rawMsg, key)
		case "min":
			err = unpopulate(val, "Min", &m.Min)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type NameAvailability.
func (n NameAvailability) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "message", n.Message)
	populate(objectMap, "name", n.Name)
	populate(objectMap, "nameAvailable", n.NameAvailable)
	populate(objectMap, "reason", n.Reason)
	populate(objectMap, "type", n.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type NameAvailability.
func (n *NameAvailability) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", n, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "message":
			err = unpopulate(val, "Message", &n.Message)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &n.Name)
			delete(rawMsg, key)
		case "nameAvailable":
			err = unpopulate(val, "NameAvailable", &n.NameAvailable)
			delete(rawMsg, key)
		case "reason":
			err = unpopulate(val, "Reason", &n.Reason)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &n.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", n, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Profile.
func (p Profile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", p.ID)
	populate(objectMap, "location", p.Location)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "properties", p.Properties)
	populate(objectMap, "tags", p.Tags)
	populate(objectMap, "type", p.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Profile.
func (p *Profile) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &p.ID)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, "Location", &p.Location)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &p.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &p.Properties)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &p.Tags)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &p.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ProfileListResult.
func (p ProfileListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ProfileListResult.
func (p *ProfileListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
			err = unpopulate(val, "Value", &p.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ProfileProperties.
func (p ProfileProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "allowedEndpointRecordTypes", p.AllowedEndpointRecordTypes)
	populate(objectMap, "dnsConfig", p.DNSConfig)
	populate(objectMap, "endpoints", p.Endpoints)
	populate(objectMap, "maxReturn", p.MaxReturn)
	populate(objectMap, "monitorConfig", p.MonitorConfig)
	populate(objectMap, "profileStatus", p.ProfileStatus)
	populate(objectMap, "trafficRoutingMethod", p.TrafficRoutingMethod)
	populate(objectMap, "trafficViewEnrollmentStatus", p.TrafficViewEnrollmentStatus)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ProfileProperties.
func (p *ProfileProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "allowedEndpointRecordTypes":
			err = unpopulate(val, "AllowedEndpointRecordTypes", &p.AllowedEndpointRecordTypes)
			delete(rawMsg, key)
		case "dnsConfig":
			err = unpopulate(val, "DNSConfig", &p.DNSConfig)
			delete(rawMsg, key)
		case "endpoints":
			err = unpopulate(val, "Endpoints", &p.Endpoints)
			delete(rawMsg, key)
		case "maxReturn":
			err = unpopulate(val, "MaxReturn", &p.MaxReturn)
			delete(rawMsg, key)
		case "monitorConfig":
			err = unpopulate(val, "MonitorConfig", &p.MonitorConfig)
			delete(rawMsg, key)
		case "profileStatus":
			err = unpopulate(val, "ProfileStatus", &p.ProfileStatus)
			delete(rawMsg, key)
		case "trafficRoutingMethod":
			err = unpopulate(val, "TrafficRoutingMethod", &p.TrafficRoutingMethod)
			delete(rawMsg, key)
		case "trafficViewEnrollmentStatus":
			err = unpopulate(val, "TrafficViewEnrollmentStatus", &p.TrafficViewEnrollmentStatus)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ProxyResource.
func (p ProxyResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", p.ID)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "type", p.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ProxyResource.
func (p *ProxyResource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &p.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &p.Name)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &p.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type QueryExperience.
func (q QueryExperience) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "endpointId", q.EndpointID)
	populate(objectMap, "latency", q.Latency)
	populate(objectMap, "queryCount", q.QueryCount)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type QueryExperience.
func (q *QueryExperience) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", q, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "endpointId":
			err = unpopulate(val, "EndpointID", &q.EndpointID)
			delete(rawMsg, key)
		case "latency":
			err = unpopulate(val, "Latency", &q.Latency)
			delete(rawMsg, key)
		case "queryCount":
			err = unpopulate(val, "QueryCount", &q.QueryCount)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", q, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Region.
func (r Region) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", r.Code)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "regions", r.Regions)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Region.
func (r *Region) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "code":
			err = unpopulate(val, "Code", &r.Code)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &r.Name)
			delete(rawMsg, key)
		case "regions":
			err = unpopulate(val, "Regions", &r.Regions)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", r.ID)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Resource.
func (r *Resource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &r.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &r.Name)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &r.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TrackedResource.
func (t TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", t.ID)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "tags", t.Tags)
	populate(objectMap, "type", t.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TrackedResource.
func (t *TrackedResource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", t, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &t.ID)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, "Location", &t.Location)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &t.Name)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &t.Tags)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &t.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", t, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TrafficFlow.
func (t TrafficFlow) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "latitude", t.Latitude)
	populate(objectMap, "longitude", t.Longitude)
	populate(objectMap, "queryExperiences", t.QueryExperiences)
	populate(objectMap, "sourceIp", t.SourceIP)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TrafficFlow.
func (t *TrafficFlow) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", t, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "latitude":
			err = unpopulate(val, "Latitude", &t.Latitude)
			delete(rawMsg, key)
		case "longitude":
			err = unpopulate(val, "Longitude", &t.Longitude)
			delete(rawMsg, key)
		case "queryExperiences":
			err = unpopulate(val, "QueryExperiences", &t.QueryExperiences)
			delete(rawMsg, key)
		case "sourceIp":
			err = unpopulate(val, "SourceIP", &t.SourceIP)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", t, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type UserMetricsModel.
func (u UserMetricsModel) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", u.ID)
	populate(objectMap, "name", u.Name)
	populate(objectMap, "properties", u.Properties)
	populate(objectMap, "type", u.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UserMetricsModel.
func (u *UserMetricsModel) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", u, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &u.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &u.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &u.Properties)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &u.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", u, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type UserMetricsProperties.
func (u UserMetricsProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "key", u.Key)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UserMetricsProperties.
func (u *UserMetricsProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", u, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "key":
			err = unpopulate(val, "Key", &u.Key)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", u, err)
		}
	}
	return nil
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