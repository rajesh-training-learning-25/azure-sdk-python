//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurearcdata

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type ActiveDirectoryConnectorDNSDetails.
func (a ActiveDirectoryConnectorDNSDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "domainName", a.DomainName)
	populate(objectMap, "nameserverIPAddresses", a.NameserverIPAddresses)
	populate(objectMap, "preferK8sDnsForPtrLookups", a.PreferK8SDNSForPtrLookups)
	populate(objectMap, "replicas", a.Replicas)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ActiveDirectoryConnectorStatus.
func (a ActiveDirectoryConnectorStatus) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "lastUpdateTime", a.LastUpdateTime)
	populate(objectMap, "observedGeneration", a.ObservedGeneration)
	populate(objectMap, "state", a.State)
	if a.AdditionalProperties != nil {
		for key, val := range a.AdditionalProperties {
			objectMap[key] = val
		}
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ActiveDirectoryConnectorStatus.
func (a *ActiveDirectoryConnectorStatus) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "lastUpdateTime":
			err = unpopulate(val, "LastUpdateTime", &a.LastUpdateTime)
			delete(rawMsg, key)
		case "observedGeneration":
			err = unpopulate(val, "ObservedGeneration", &a.ObservedGeneration)
			delete(rawMsg, key)
		case "state":
			err = unpopulate(val, "State", &a.State)
			delete(rawMsg, key)
		default:
			if a.AdditionalProperties == nil {
				a.AdditionalProperties = map[string]interface{}{}
			}
			if val != nil {
				var aux interface{}
				err = json.Unmarshal(val, &aux)
				a.AdditionalProperties[key] = aux
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ActiveDirectoryDomainControllers.
func (a ActiveDirectoryDomainControllers) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "primaryDomainController", a.PrimaryDomainController)
	populate(objectMap, "secondaryDomainControllers", a.SecondaryDomainControllers)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DataControllerProperties.
func (d DataControllerProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "basicLoginInformation", d.BasicLoginInformation)
	populate(objectMap, "clusterId", d.ClusterID)
	populate(objectMap, "extensionId", d.ExtensionID)
	populate(objectMap, "infrastructure", d.Infrastructure)
	populate(objectMap, "k8sRaw", &d.K8SRaw)
	populateTimeRFC3339(objectMap, "lastUploadedDate", d.LastUploadedDate)
	populate(objectMap, "logAnalyticsWorkspaceConfig", d.LogAnalyticsWorkspaceConfig)
	populate(objectMap, "logsDashboardCredential", d.LogsDashboardCredential)
	populate(objectMap, "metricsDashboardCredential", d.MetricsDashboardCredential)
	populate(objectMap, "onPremiseProperty", d.OnPremiseProperty)
	populate(objectMap, "provisioningState", d.ProvisioningState)
	populate(objectMap, "uploadServicePrincipal", d.UploadServicePrincipal)
	populate(objectMap, "uploadWatermark", d.UploadWatermark)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DataControllerProperties.
func (d *DataControllerProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "basicLoginInformation":
			err = unpopulate(val, "BasicLoginInformation", &d.BasicLoginInformation)
			delete(rawMsg, key)
		case "clusterId":
			err = unpopulate(val, "ClusterID", &d.ClusterID)
			delete(rawMsg, key)
		case "extensionId":
			err = unpopulate(val, "ExtensionID", &d.ExtensionID)
			delete(rawMsg, key)
		case "infrastructure":
			err = unpopulate(val, "Infrastructure", &d.Infrastructure)
			delete(rawMsg, key)
		case "k8sRaw":
			err = unpopulate(val, "K8SRaw", &d.K8SRaw)
			delete(rawMsg, key)
		case "lastUploadedDate":
			err = unpopulateTimeRFC3339(val, "LastUploadedDate", &d.LastUploadedDate)
			delete(rawMsg, key)
		case "logAnalyticsWorkspaceConfig":
			err = unpopulate(val, "LogAnalyticsWorkspaceConfig", &d.LogAnalyticsWorkspaceConfig)
			delete(rawMsg, key)
		case "logsDashboardCredential":
			err = unpopulate(val, "LogsDashboardCredential", &d.LogsDashboardCredential)
			delete(rawMsg, key)
		case "metricsDashboardCredential":
			err = unpopulate(val, "MetricsDashboardCredential", &d.MetricsDashboardCredential)
			delete(rawMsg, key)
		case "onPremiseProperty":
			err = unpopulate(val, "OnPremiseProperty", &d.OnPremiseProperty)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &d.ProvisioningState)
			delete(rawMsg, key)
		case "uploadServicePrincipal":
			err = unpopulate(val, "UploadServicePrincipal", &d.UploadServicePrincipal)
			delete(rawMsg, key)
		case "uploadWatermark":
			err = unpopulate(val, "UploadWatermark", &d.UploadWatermark)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DataControllerResource.
func (d DataControllerResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "extendedLocation", d.ExtendedLocation)
	populate(objectMap, "id", d.ID)
	populate(objectMap, "location", d.Location)
	populate(objectMap, "name", d.Name)
	populate(objectMap, "properties", d.Properties)
	populate(objectMap, "systemData", d.SystemData)
	populate(objectMap, "tags", d.Tags)
	populate(objectMap, "type", d.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DataControllerUpdate.
func (d DataControllerUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", d.Properties)
	populate(objectMap, "tags", d.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type K8SResourceRequirements.
func (k K8SResourceRequirements) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "limits", k.Limits)
	populate(objectMap, "requests", k.Requests)
	if k.AdditionalProperties != nil {
		for key, val := range k.AdditionalProperties {
			objectMap[key] = val
		}
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type K8SResourceRequirements.
func (k *K8SResourceRequirements) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", k, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "limits":
			err = unpopulate(val, "Limits", &k.Limits)
			delete(rawMsg, key)
		case "requests":
			err = unpopulate(val, "Requests", &k.Requests)
			delete(rawMsg, key)
		default:
			if k.AdditionalProperties == nil {
				k.AdditionalProperties = map[string]interface{}{}
			}
			if val != nil {
				var aux interface{}
				err = json.Unmarshal(val, &aux)
				k.AdditionalProperties[key] = aux
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", k, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type K8SScheduling.
func (k K8SScheduling) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "default", k.Default)
	if k.AdditionalProperties != nil {
		for key, val := range k.AdditionalProperties {
			objectMap[key] = val
		}
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type K8SScheduling.
func (k *K8SScheduling) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", k, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "default":
			err = unpopulate(val, "Default", &k.Default)
			delete(rawMsg, key)
		default:
			if k.AdditionalProperties == nil {
				k.AdditionalProperties = map[string]interface{}{}
			}
			if val != nil {
				var aux interface{}
				err = json.Unmarshal(val, &aux)
				k.AdditionalProperties[key] = aux
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", k, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type K8SSchedulingOptions.
func (k K8SSchedulingOptions) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "resources", k.Resources)
	if k.AdditionalProperties != nil {
		for key, val := range k.AdditionalProperties {
			objectMap[key] = val
		}
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type K8SSchedulingOptions.
func (k *K8SSchedulingOptions) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", k, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "resources":
			err = unpopulate(val, "Resources", &k.Resources)
			delete(rawMsg, key)
		default:
			if k.AdditionalProperties == nil {
				k.AdditionalProperties = map[string]interface{}{}
			}
			if val != nil {
				var aux interface{}
				err = json.Unmarshal(val, &aux)
				k.AdditionalProperties[key] = aux
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", k, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PostgresInstance.
func (p PostgresInstance) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "extendedLocation", p.ExtendedLocation)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "location", p.Location)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "properties", p.Properties)
	populate(objectMap, "sku", p.SKU)
	populate(objectMap, "systemData", p.SystemData)
	populate(objectMap, "tags", p.Tags)
	populate(objectMap, "type", p.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PostgresInstanceProperties.
func (p PostgresInstanceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "admin", p.Admin)
	populate(objectMap, "basicLoginInformation", p.BasicLoginInformation)
	populate(objectMap, "dataControllerId", p.DataControllerID)
	populate(objectMap, "k8sRaw", &p.K8SRaw)
	populateTimeRFC3339(objectMap, "lastUploadedDate", p.LastUploadedDate)
	populate(objectMap, "provisioningState", p.ProvisioningState)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PostgresInstanceProperties.
func (p *PostgresInstanceProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "admin":
			err = unpopulate(val, "Admin", &p.Admin)
			delete(rawMsg, key)
		case "basicLoginInformation":
			err = unpopulate(val, "BasicLoginInformation", &p.BasicLoginInformation)
			delete(rawMsg, key)
		case "dataControllerId":
			err = unpopulate(val, "DataControllerID", &p.DataControllerID)
			delete(rawMsg, key)
		case "k8sRaw":
			err = unpopulate(val, "K8SRaw", &p.K8SRaw)
			delete(rawMsg, key)
		case "lastUploadedDate":
			err = unpopulateTimeRFC3339(val, "LastUploadedDate", &p.LastUploadedDate)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &p.ProvisioningState)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PostgresInstanceUpdate.
func (p PostgresInstanceUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", p.Properties)
	populate(objectMap, "tags", p.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SQLManagedInstance.
func (s SQLManagedInstance) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "extendedLocation", s.ExtendedLocation)
	populate(objectMap, "id", s.ID)
	populate(objectMap, "location", s.Location)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "sku", s.SKU)
	populate(objectMap, "systemData", s.SystemData)
	populate(objectMap, "tags", s.Tags)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SQLManagedInstanceK8SRaw.
func (s SQLManagedInstanceK8SRaw) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "spec", s.Spec)
	if s.AdditionalProperties != nil {
		for key, val := range s.AdditionalProperties {
			objectMap[key] = val
		}
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SQLManagedInstanceK8SRaw.
func (s *SQLManagedInstanceK8SRaw) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "spec":
			err = unpopulate(val, "Spec", &s.Spec)
			delete(rawMsg, key)
		default:
			if s.AdditionalProperties == nil {
				s.AdditionalProperties = map[string]interface{}{}
			}
			if val != nil {
				var aux interface{}
				err = json.Unmarshal(val, &aux)
				s.AdditionalProperties[key] = aux
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SQLManagedInstanceK8SSpec.
func (s SQLManagedInstanceK8SSpec) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "replicas", s.Replicas)
	populate(objectMap, "scheduling", s.Scheduling)
	if s.AdditionalProperties != nil {
		for key, val := range s.AdditionalProperties {
			objectMap[key] = val
		}
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SQLManagedInstanceK8SSpec.
func (s *SQLManagedInstanceK8SSpec) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "replicas":
			err = unpopulate(val, "Replicas", &s.Replicas)
			delete(rawMsg, key)
		case "scheduling":
			err = unpopulate(val, "Scheduling", &s.Scheduling)
			delete(rawMsg, key)
		default:
			if s.AdditionalProperties == nil {
				s.AdditionalProperties = map[string]interface{}{}
			}
			if val != nil {
				var aux interface{}
				err = json.Unmarshal(val, &aux)
				s.AdditionalProperties[key] = aux
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SQLManagedInstanceProperties.
func (s SQLManagedInstanceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "activeDirectoryInformation", s.ActiveDirectoryInformation)
	populate(objectMap, "admin", s.Admin)
	populate(objectMap, "basicLoginInformation", s.BasicLoginInformation)
	populate(objectMap, "clusterId", s.ClusterID)
	populate(objectMap, "dataControllerId", s.DataControllerID)
	populate(objectMap, "endTime", s.EndTime)
	populate(objectMap, "extensionId", s.ExtensionID)
	populate(objectMap, "k8sRaw", s.K8SRaw)
	populateTimeRFC3339(objectMap, "lastUploadedDate", s.LastUploadedDate)
	populate(objectMap, "licenseType", s.LicenseType)
	populate(objectMap, "provisioningState", s.ProvisioningState)
	populate(objectMap, "startTime", s.StartTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SQLManagedInstanceProperties.
func (s *SQLManagedInstanceProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "activeDirectoryInformation":
			err = unpopulate(val, "ActiveDirectoryInformation", &s.ActiveDirectoryInformation)
			delete(rawMsg, key)
		case "admin":
			err = unpopulate(val, "Admin", &s.Admin)
			delete(rawMsg, key)
		case "basicLoginInformation":
			err = unpopulate(val, "BasicLoginInformation", &s.BasicLoginInformation)
			delete(rawMsg, key)
		case "clusterId":
			err = unpopulate(val, "ClusterID", &s.ClusterID)
			delete(rawMsg, key)
		case "dataControllerId":
			err = unpopulate(val, "DataControllerID", &s.DataControllerID)
			delete(rawMsg, key)
		case "endTime":
			err = unpopulate(val, "EndTime", &s.EndTime)
			delete(rawMsg, key)
		case "extensionId":
			err = unpopulate(val, "ExtensionID", &s.ExtensionID)
			delete(rawMsg, key)
		case "k8sRaw":
			err = unpopulate(val, "K8SRaw", &s.K8SRaw)
			delete(rawMsg, key)
		case "lastUploadedDate":
			err = unpopulateTimeRFC3339(val, "LastUploadedDate", &s.LastUploadedDate)
			delete(rawMsg, key)
		case "licenseType":
			err = unpopulate(val, "LicenseType", &s.LicenseType)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &s.ProvisioningState)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulate(val, "StartTime", &s.StartTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SQLManagedInstanceUpdate.
func (s SQLManagedInstanceUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", s.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SQLServerInstance.
func (s SQLServerInstance) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", s.ID)
	populate(objectMap, "location", s.Location)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "systemData", s.SystemData)
	populate(objectMap, "tags", s.Tags)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SQLServerInstanceProperties.
func (s SQLServerInstanceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "azureDefenderStatus", s.AzureDefenderStatus)
	populateTimeRFC3339(objectMap, "azureDefenderStatusLastUpdated", s.AzureDefenderStatusLastUpdated)
	populate(objectMap, "collation", s.Collation)
	populate(objectMap, "containerResourceId", s.ContainerResourceID)
	populate(objectMap, "createTime", s.CreateTime)
	populate(objectMap, "currentVersion", s.CurrentVersion)
	populate(objectMap, "edition", s.Edition)
	populate(objectMap, "hostType", s.HostType)
	populate(objectMap, "instanceName", s.InstanceName)
	populate(objectMap, "licenseType", s.LicenseType)
	populate(objectMap, "patchLevel", s.PatchLevel)
	populate(objectMap, "productId", s.ProductID)
	populate(objectMap, "provisioningState", s.ProvisioningState)
	populate(objectMap, "status", s.Status)
	populate(objectMap, "tcpDynamicPorts", s.TCPDynamicPorts)
	populate(objectMap, "tcpStaticPorts", s.TCPStaticPorts)
	populate(objectMap, "vCore", s.VCore)
	populate(objectMap, "version", s.Version)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SQLServerInstanceProperties.
func (s *SQLServerInstanceProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "azureDefenderStatus":
			err = unpopulate(val, "AzureDefenderStatus", &s.AzureDefenderStatus)
			delete(rawMsg, key)
		case "azureDefenderStatusLastUpdated":
			err = unpopulateTimeRFC3339(val, "AzureDefenderStatusLastUpdated", &s.AzureDefenderStatusLastUpdated)
			delete(rawMsg, key)
		case "collation":
			err = unpopulate(val, "Collation", &s.Collation)
			delete(rawMsg, key)
		case "containerResourceId":
			err = unpopulate(val, "ContainerResourceID", &s.ContainerResourceID)
			delete(rawMsg, key)
		case "createTime":
			err = unpopulate(val, "CreateTime", &s.CreateTime)
			delete(rawMsg, key)
		case "currentVersion":
			err = unpopulate(val, "CurrentVersion", &s.CurrentVersion)
			delete(rawMsg, key)
		case "edition":
			err = unpopulate(val, "Edition", &s.Edition)
			delete(rawMsg, key)
		case "hostType":
			err = unpopulate(val, "HostType", &s.HostType)
			delete(rawMsg, key)
		case "instanceName":
			err = unpopulate(val, "InstanceName", &s.InstanceName)
			delete(rawMsg, key)
		case "licenseType":
			err = unpopulate(val, "LicenseType", &s.LicenseType)
			delete(rawMsg, key)
		case "patchLevel":
			err = unpopulate(val, "PatchLevel", &s.PatchLevel)
			delete(rawMsg, key)
		case "productId":
			err = unpopulate(val, "ProductID", &s.ProductID)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &s.ProvisioningState)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &s.Status)
			delete(rawMsg, key)
		case "tcpDynamicPorts":
			err = unpopulate(val, "TCPDynamicPorts", &s.TCPDynamicPorts)
			delete(rawMsg, key)
		case "tcpStaticPorts":
			err = unpopulate(val, "TCPStaticPorts", &s.TCPStaticPorts)
			delete(rawMsg, key)
		case "vCore":
			err = unpopulate(val, "VCore", &s.VCore)
			delete(rawMsg, key)
		case "version":
			err = unpopulate(val, "Version", &s.Version)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SQLServerInstanceUpdate.
func (s SQLServerInstanceUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", s.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, "CreatedAt", &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, "CreatedBy", &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, "CreatedByType", &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, "LastModifiedAt", &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, "LastModifiedBy", &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, "LastModifiedByType", &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
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
	populate(objectMap, "systemData", t.SystemData)
	populate(objectMap, "tags", t.Tags)
	populate(objectMap, "type", t.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type UploadWatermark.
func (u UploadWatermark) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "logs", u.Logs)
	populateTimeRFC3339(objectMap, "metrics", u.Metrics)
	populateTimeRFC3339(objectMap, "usages", u.Usages)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UploadWatermark.
func (u *UploadWatermark) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", u, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "logs":
			err = unpopulateTimeRFC3339(val, "Logs", &u.Logs)
			delete(rawMsg, key)
		case "metrics":
			err = unpopulateTimeRFC3339(val, "Metrics", &u.Metrics)
			delete(rawMsg, key)
		case "usages":
			err = unpopulateTimeRFC3339(val, "Usages", &u.Usages)
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