//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapplicationinsights

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type APIKeyRequest.
func (a APIKeyRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "linkedReadProperties", a.LinkedReadProperties)
	populate(objectMap, "linkedWriteProperties", a.LinkedWriteProperties)
	populate(objectMap, "name", a.Name)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Annotation.
func (a Annotation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "AnnotationName", a.AnnotationName)
	populate(objectMap, "Category", a.Category)
	populateTimeRFC3339(objectMap, "EventTime", a.EventTime)
	populate(objectMap, "Id", a.ID)
	populate(objectMap, "Properties", a.Properties)
	populate(objectMap, "RelatedAnnotation", a.RelatedAnnotation)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Annotation.
func (a *Annotation) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "AnnotationName":
			err = unpopulate(val, "AnnotationName", &a.AnnotationName)
			delete(rawMsg, key)
		case "Category":
			err = unpopulate(val, "Category", &a.Category)
			delete(rawMsg, key)
		case "EventTime":
			err = unpopulateTimeRFC3339(val, "EventTime", &a.EventTime)
			delete(rawMsg, key)
		case "Id":
			err = unpopulate(val, "ID", &a.ID)
			delete(rawMsg, key)
		case "Properties":
			err = unpopulate(val, "Properties", &a.Properties)
			delete(rawMsg, key)
		case "RelatedAnnotation":
			err = unpopulate(val, "RelatedAnnotation", &a.RelatedAnnotation)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Component.
func (c Component) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", c.Etag)
	populate(objectMap, "id", c.ID)
	populate(objectMap, "kind", c.Kind)
	populate(objectMap, "location", c.Location)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "tags", c.Tags)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ComponentBillingFeatures.
func (c ComponentBillingFeatures) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "CurrentBillingFeatures", c.CurrentBillingFeatures)
	populate(objectMap, "DataVolumeCap", c.DataVolumeCap)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ComponentFavorite.
func (c ComponentFavorite) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "Category", c.Category)
	populate(objectMap, "Config", c.Config)
	populate(objectMap, "FavoriteId", c.FavoriteID)
	populate(objectMap, "FavoriteType", c.FavoriteType)
	populate(objectMap, "IsGeneratedFromTemplate", c.IsGeneratedFromTemplate)
	populate(objectMap, "Name", c.Name)
	populate(objectMap, "SourceType", c.SourceType)
	populate(objectMap, "Tags", c.Tags)
	populate(objectMap, "TimeModified", c.TimeModified)
	populate(objectMap, "UserId", c.UserID)
	populate(objectMap, "Version", c.Version)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ComponentLinkedStorageAccountsPatch.
func (c ComponentLinkedStorageAccountsPatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", c.Properties)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ComponentProactiveDetectionConfiguration.
func (c ComponentProactiveDetectionConfiguration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "CustomEmails", c.CustomEmails)
	populate(objectMap, "Enabled", c.Enabled)
	populate(objectMap, "LastUpdatedTime", c.LastUpdatedTime)
	populate(objectMap, "Name", c.Name)
	populate(objectMap, "RuleDefinitions", c.RuleDefinitions)
	populate(objectMap, "SendEmailsToSubscriptionOwners", c.SendEmailsToSubscriptionOwners)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ComponentProperties.
func (c ComponentProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "AppId", c.AppID)
	populate(objectMap, "ApplicationId", c.ApplicationID)
	populate(objectMap, "Application_Type", c.ApplicationType)
	populate(objectMap, "ConnectionString", c.ConnectionString)
	populateTimeRFC3339(objectMap, "CreationDate", c.CreationDate)
	populate(objectMap, "DisableIpMasking", c.DisableIPMasking)
	populate(objectMap, "DisableLocalAuth", c.DisableLocalAuth)
	populate(objectMap, "Flow_Type", c.FlowType)
	populate(objectMap, "ForceCustomerStorageForProfiler", c.ForceCustomerStorageForProfiler)
	populate(objectMap, "HockeyAppId", c.HockeyAppID)
	populate(objectMap, "HockeyAppToken", c.HockeyAppToken)
	populate(objectMap, "ImmediatePurgeDataOn30Days", c.ImmediatePurgeDataOn30Days)
	populate(objectMap, "IngestionMode", c.IngestionMode)
	populate(objectMap, "InstrumentationKey", c.InstrumentationKey)
	populateTimeRFC3339(objectMap, "LaMigrationDate", c.LaMigrationDate)
	populate(objectMap, "Name", c.Name)
	populate(objectMap, "PrivateLinkScopedResources", c.PrivateLinkScopedResources)
	populate(objectMap, "provisioningState", c.ProvisioningState)
	populate(objectMap, "publicNetworkAccessForIngestion", c.PublicNetworkAccessForIngestion)
	populate(objectMap, "publicNetworkAccessForQuery", c.PublicNetworkAccessForQuery)
	populate(objectMap, "Request_Source", c.RequestSource)
	populate(objectMap, "RetentionInDays", c.RetentionInDays)
	populate(objectMap, "SamplingPercentage", c.SamplingPercentage)
	populate(objectMap, "TenantId", c.TenantID)
	populate(objectMap, "WorkspaceResourceId", c.WorkspaceResourceID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ComponentProperties.
func (c *ComponentProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "AppId":
			err = unpopulate(val, "AppID", &c.AppID)
			delete(rawMsg, key)
		case "ApplicationId":
			err = unpopulate(val, "ApplicationID", &c.ApplicationID)
			delete(rawMsg, key)
		case "Application_Type":
			err = unpopulate(val, "ApplicationType", &c.ApplicationType)
			delete(rawMsg, key)
		case "ConnectionString":
			err = unpopulate(val, "ConnectionString", &c.ConnectionString)
			delete(rawMsg, key)
		case "CreationDate":
			err = unpopulateTimeRFC3339(val, "CreationDate", &c.CreationDate)
			delete(rawMsg, key)
		case "DisableIpMasking":
			err = unpopulate(val, "DisableIPMasking", &c.DisableIPMasking)
			delete(rawMsg, key)
		case "DisableLocalAuth":
			err = unpopulate(val, "DisableLocalAuth", &c.DisableLocalAuth)
			delete(rawMsg, key)
		case "Flow_Type":
			err = unpopulate(val, "FlowType", &c.FlowType)
			delete(rawMsg, key)
		case "ForceCustomerStorageForProfiler":
			err = unpopulate(val, "ForceCustomerStorageForProfiler", &c.ForceCustomerStorageForProfiler)
			delete(rawMsg, key)
		case "HockeyAppId":
			err = unpopulate(val, "HockeyAppID", &c.HockeyAppID)
			delete(rawMsg, key)
		case "HockeyAppToken":
			err = unpopulate(val, "HockeyAppToken", &c.HockeyAppToken)
			delete(rawMsg, key)
		case "ImmediatePurgeDataOn30Days":
			err = unpopulate(val, "ImmediatePurgeDataOn30Days", &c.ImmediatePurgeDataOn30Days)
			delete(rawMsg, key)
		case "IngestionMode":
			err = unpopulate(val, "IngestionMode", &c.IngestionMode)
			delete(rawMsg, key)
		case "InstrumentationKey":
			err = unpopulate(val, "InstrumentationKey", &c.InstrumentationKey)
			delete(rawMsg, key)
		case "LaMigrationDate":
			err = unpopulateTimeRFC3339(val, "LaMigrationDate", &c.LaMigrationDate)
			delete(rawMsg, key)
		case "Name":
			err = unpopulate(val, "Name", &c.Name)
			delete(rawMsg, key)
		case "PrivateLinkScopedResources":
			err = unpopulate(val, "PrivateLinkScopedResources", &c.PrivateLinkScopedResources)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &c.ProvisioningState)
			delete(rawMsg, key)
		case "publicNetworkAccessForIngestion":
			err = unpopulate(val, "PublicNetworkAccessForIngestion", &c.PublicNetworkAccessForIngestion)
			delete(rawMsg, key)
		case "publicNetworkAccessForQuery":
			err = unpopulate(val, "PublicNetworkAccessForQuery", &c.PublicNetworkAccessForQuery)
			delete(rawMsg, key)
		case "Request_Source":
			err = unpopulate(val, "RequestSource", &c.RequestSource)
			delete(rawMsg, key)
		case "RetentionInDays":
			err = unpopulate(val, "RetentionInDays", &c.RetentionInDays)
			delete(rawMsg, key)
		case "SamplingPercentage":
			err = unpopulate(val, "SamplingPercentage", &c.SamplingPercentage)
			delete(rawMsg, key)
		case "TenantId":
			err = unpopulate(val, "TenantID", &c.TenantID)
			delete(rawMsg, key)
		case "WorkspaceResourceId":
			err = unpopulate(val, "WorkspaceResourceID", &c.WorkspaceResourceID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ComponentPurgeBody.
func (c ComponentPurgeBody) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "filters", c.Filters)
	populate(objectMap, "table", c.Table)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ComponentsResource.
func (c ComponentsResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", c.ID)
	populate(objectMap, "location", c.Location)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "tags", c.Tags)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type InnerError.
func (i *InnerError) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "diagnosticcontext":
			err = unpopulate(val, "Diagnosticcontext", &i.Diagnosticcontext)
			delete(rawMsg, key)
		case "time":
			err = unpopulateTimeRFC3339(val, "Time", &i.Time)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ManagedServiceIdentity.
func (m ManagedServiceIdentity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "principalId", m.PrincipalID)
	populate(objectMap, "tenantId", m.TenantID)
	populate(objectMap, "type", m.Type)
	populate(objectMap, "userAssignedIdentities", m.UserAssignedIdentities)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MyWorkbook.
func (m MyWorkbook) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", m.Etag)
	populate(objectMap, "id", m.ID)
	populate(objectMap, "identity", m.Identity)
	populate(objectMap, "kind", m.Kind)
	populate(objectMap, "location", m.Location)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "properties", m.Properties)
	populate(objectMap, "systemData", m.SystemData)
	populate(objectMap, "tags", m.Tags)
	populate(objectMap, "type", m.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MyWorkbookProperties.
func (m MyWorkbookProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "category", m.Category)
	populate(objectMap, "displayName", m.DisplayName)
	populate(objectMap, "serializedData", m.SerializedData)
	populate(objectMap, "sourceId", m.SourceID)
	populate(objectMap, "storageUri", m.StorageURI)
	populate(objectMap, "tags", m.Tags)
	populate(objectMap, "timeModified", m.TimeModified)
	populate(objectMap, "userId", m.UserID)
	populate(objectMap, "version", m.Version)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MyWorkbookResource.
func (m MyWorkbookResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", m.Etag)
	populate(objectMap, "id", m.ID)
	populate(objectMap, "identity", m.Identity)
	populate(objectMap, "location", m.Location)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "tags", m.Tags)
	populate(objectMap, "type", m.Type)
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

// MarshalJSON implements the json.Marshaller interface for type TagsResource.
func (t TagsResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", t.Tags)
	return json.Marshal(objectMap)
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

// MarshalJSON implements the json.Marshaller interface for type WebTest.
func (w WebTest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", w.ID)
	populate(objectMap, "kind", w.Kind)
	populate(objectMap, "location", w.Location)
	populate(objectMap, "name", w.Name)
	populate(objectMap, "properties", w.Properties)
	populate(objectMap, "tags", w.Tags)
	populate(objectMap, "type", w.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type WebTestProperties.
func (w WebTestProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "Configuration", w.Configuration)
	populate(objectMap, "Description", w.Description)
	populate(objectMap, "Enabled", w.Enabled)
	populate(objectMap, "Frequency", w.Frequency)
	populate(objectMap, "Locations", w.Locations)
	populate(objectMap, "provisioningState", w.ProvisioningState)
	populate(objectMap, "RetryEnabled", w.RetryEnabled)
	populate(objectMap, "SyntheticMonitorId", w.SyntheticMonitorID)
	populate(objectMap, "Timeout", w.Timeout)
	populate(objectMap, "Kind", w.WebTestKind)
	populate(objectMap, "Name", w.WebTestName)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type WebtestsResource.
func (w WebtestsResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", w.ID)
	populate(objectMap, "location", w.Location)
	populate(objectMap, "name", w.Name)
	populate(objectMap, "tags", w.Tags)
	populate(objectMap, "type", w.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type WorkItemCreateConfiguration.
func (w WorkItemCreateConfiguration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "ConnectorDataConfiguration", w.ConnectorDataConfiguration)
	populate(objectMap, "ConnectorId", w.ConnectorID)
	populate(objectMap, "ValidateOnly", w.ValidateOnly)
	populate(objectMap, "WorkItemProperties", w.WorkItemProperties)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Workbook.
func (w Workbook) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", w.Etag)
	populate(objectMap, "id", w.ID)
	populate(objectMap, "identity", w.Identity)
	populate(objectMap, "kind", w.Kind)
	populate(objectMap, "location", w.Location)
	populate(objectMap, "name", w.Name)
	populate(objectMap, "properties", w.Properties)
	populate(objectMap, "systemData", w.SystemData)
	populate(objectMap, "tags", w.Tags)
	populate(objectMap, "type", w.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type WorkbookProperties.
func (w WorkbookProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "category", w.Category)
	populate(objectMap, "description", w.Description)
	populate(objectMap, "displayName", w.DisplayName)
	populate(objectMap, "revision", w.Revision)
	populate(objectMap, "serializedData", w.SerializedData)
	populate(objectMap, "sourceId", w.SourceID)
	populate(objectMap, "storageUri", w.StorageURI)
	populate(objectMap, "tags", w.Tags)
	populateTimeRFC3339(objectMap, "timeModified", w.TimeModified)
	populate(objectMap, "userId", w.UserID)
	populate(objectMap, "version", w.Version)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type WorkbookProperties.
func (w *WorkbookProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", w, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "category":
			err = unpopulate(val, "Category", &w.Category)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &w.Description)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, "DisplayName", &w.DisplayName)
			delete(rawMsg, key)
		case "revision":
			err = unpopulate(val, "Revision", &w.Revision)
			delete(rawMsg, key)
		case "serializedData":
			err = unpopulate(val, "SerializedData", &w.SerializedData)
			delete(rawMsg, key)
		case "sourceId":
			err = unpopulate(val, "SourceID", &w.SourceID)
			delete(rawMsg, key)
		case "storageUri":
			err = unpopulate(val, "StorageURI", &w.StorageURI)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &w.Tags)
			delete(rawMsg, key)
		case "timeModified":
			err = unpopulateTimeRFC3339(val, "TimeModified", &w.TimeModified)
			delete(rawMsg, key)
		case "userId":
			err = unpopulate(val, "UserID", &w.UserID)
			delete(rawMsg, key)
		case "version":
			err = unpopulate(val, "Version", &w.Version)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", w, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type WorkbookPropertiesUpdateParameters.
func (w WorkbookPropertiesUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "category", w.Category)
	populate(objectMap, "description", w.Description)
	populate(objectMap, "displayName", w.DisplayName)
	populate(objectMap, "revision", w.Revision)
	populate(objectMap, "serializedData", w.SerializedData)
	populate(objectMap, "tags", w.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type WorkbookResource.
func (w WorkbookResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", w.Etag)
	populate(objectMap, "id", w.ID)
	populate(objectMap, "identity", w.Identity)
	populate(objectMap, "kind", w.Kind)
	populate(objectMap, "location", w.Location)
	populate(objectMap, "name", w.Name)
	populate(objectMap, "tags", w.Tags)
	populate(objectMap, "type", w.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type WorkbookResourceIdentity.
func (w WorkbookResourceIdentity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "principalId", w.PrincipalID)
	populate(objectMap, "tenantId", w.TenantID)
	populate(objectMap, "type", w.Type)
	populate(objectMap, "userAssignedIdentities", w.UserAssignedIdentities)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type WorkbookTemplate.
func (w WorkbookTemplate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", w.ID)
	populate(objectMap, "location", w.Location)
	populate(objectMap, "name", w.Name)
	populate(objectMap, "properties", w.Properties)
	populate(objectMap, "tags", w.Tags)
	populate(objectMap, "type", w.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type WorkbookTemplateLocalizedGallery.
func (w WorkbookTemplateLocalizedGallery) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "galleries", w.Galleries)
	populate(objectMap, "templateData", &w.TemplateData)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type WorkbookTemplateProperties.
func (w WorkbookTemplateProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "author", w.Author)
	populate(objectMap, "galleries", w.Galleries)
	populate(objectMap, "localized", w.Localized)
	populate(objectMap, "priority", w.Priority)
	populate(objectMap, "templateData", &w.TemplateData)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type WorkbookTemplateResource.
func (w WorkbookTemplateResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", w.ID)
	populate(objectMap, "location", w.Location)
	populate(objectMap, "name", w.Name)
	populate(objectMap, "tags", w.Tags)
	populate(objectMap, "type", w.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type WorkbookTemplateUpdateParameters.
func (w WorkbookTemplateUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", w.Properties)
	populate(objectMap, "tags", w.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type WorkbookUpdateParameters.
func (w WorkbookUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "kind", w.Kind)
	populate(objectMap, "properties", w.Properties)
	populate(objectMap, "tags", w.Tags)
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