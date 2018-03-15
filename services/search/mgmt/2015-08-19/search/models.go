package search

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
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// AdminKeyKind enumerates the values for admin key kind.
type AdminKeyKind string

const (
	// Primary ...
	Primary AdminKeyKind = "primary"
	// Secondary ...
	Secondary AdminKeyKind = "secondary"
)

// PossibleAdminKeyKindValues returns an array of possible values for the AdminKeyKind const type.
func PossibleAdminKeyKindValues() [2]AdminKeyKind {
	return [2]AdminKeyKind{Primary, Secondary}
}

// HostingMode enumerates the values for hosting mode.
type HostingMode string

const (
	// Default ...
	Default HostingMode = "default"
	// HighDensity ...
	HighDensity HostingMode = "highDensity"
)

// PossibleHostingModeValues returns an array of possible values for the HostingMode const type.
func PossibleHostingModeValues() [2]HostingMode {
	return [2]HostingMode{Default, HighDensity}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Failed ...
	Failed ProvisioningState = "failed"
	// Provisioning ...
	Provisioning ProvisioningState = "provisioning"
	// Succeeded ...
	Succeeded ProvisioningState = "succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() [3]ProvisioningState {
	return [3]ProvisioningState{Failed, Provisioning, Succeeded}
}

// ServiceStatus enumerates the values for service status.
type ServiceStatus string

const (
	// ServiceStatusDegraded ...
	ServiceStatusDegraded ServiceStatus = "degraded"
	// ServiceStatusDeleting ...
	ServiceStatusDeleting ServiceStatus = "deleting"
	// ServiceStatusDisabled ...
	ServiceStatusDisabled ServiceStatus = "disabled"
	// ServiceStatusError ...
	ServiceStatusError ServiceStatus = "error"
	// ServiceStatusProvisioning ...
	ServiceStatusProvisioning ServiceStatus = "provisioning"
	// ServiceStatusRunning ...
	ServiceStatusRunning ServiceStatus = "running"
)

// PossibleServiceStatusValues returns an array of possible values for the ServiceStatus const type.
func PossibleServiceStatusValues() [6]ServiceStatus {
	return [6]ServiceStatus{ServiceStatusDegraded, ServiceStatusDeleting, ServiceStatusDisabled, ServiceStatusError, ServiceStatusProvisioning, ServiceStatusRunning}
}

// SkuName enumerates the values for sku name.
type SkuName string

const (
	// Basic ...
	Basic SkuName = "basic"
	// Free ...
	Free SkuName = "free"
	// Standard ...
	Standard SkuName = "standard"
	// Standard2 ...
	Standard2 SkuName = "standard2"
	// Standard3 ...
	Standard3 SkuName = "standard3"
)

// PossibleSkuNameValues returns an array of possible values for the SkuName const type.
func PossibleSkuNameValues() [5]SkuName {
	return [5]SkuName{Basic, Free, Standard, Standard2, Standard3}
}

// UnavailableNameReason enumerates the values for unavailable name reason.
type UnavailableNameReason string

const (
	// AlreadyExists ...
	AlreadyExists UnavailableNameReason = "AlreadyExists"
	// Invalid ...
	Invalid UnavailableNameReason = "Invalid"
)

// PossibleUnavailableNameReasonValues returns an array of possible values for the UnavailableNameReason const type.
func PossibleUnavailableNameReasonValues() [2]UnavailableNameReason {
	return [2]UnavailableNameReason{AlreadyExists, Invalid}
}

// AdminKeyResult response containing the primary and secondary admin API keys for a given Azure Search service.
type AdminKeyResult struct {
	autorest.Response `json:"-"`
	// PrimaryKey - The primary admin API key of the Search service.
	PrimaryKey *string `json:"primaryKey,omitempty"`
	// SecondaryKey - The secondary admin API key of the Search service.
	SecondaryKey *string `json:"secondaryKey,omitempty"`
}

// CheckNameAvailabilityInput input of check name availability API.
type CheckNameAvailabilityInput struct {
	// Name - The Search service name to validate. Search service names must only contain lowercase letters, digits or dashes, cannot use dash as the first two or last one characters, cannot contain consecutive dashes, and must be between 2 and 60 characters in length.
	Name *string `json:"name,omitempty"`
	// Type - The type of the resource whose name is to be validated. This value must always be 'searchServices'.
	Type *string `json:"type,omitempty"`
}

// CheckNameAvailabilityOutput output of check name availability API.
type CheckNameAvailabilityOutput struct {
	autorest.Response `json:"-"`
	// IsNameAvailable - A value indicating whether the name is available.
	IsNameAvailable *bool `json:"nameAvailable,omitempty"`
	// Reason - The reason why the name is not available. 'Invalid' indicates the name provided does not match the naming requirements (incorrect length, unsupported characters, etc.). 'AlreadyExists' indicates that the name is already in use and is therefore unavailable. Possible values include: 'Invalid', 'AlreadyExists'
	Reason UnavailableNameReason `json:"reason,omitempty"`
	// Message - A message that explains why the name is invalid and provides resource naming requirements. Available only if 'Invalid' is returned in the 'reason' property.
	Message *string `json:"message,omitempty"`
}

// CloudError contains information about an API error.
type CloudError struct {
	// Error - Describes a particular API error with an error code and a message.
	Error *CloudErrorBody `json:"error,omitempty"`
}

// CloudErrorBody describes a particular API error with an error code and a message.
type CloudErrorBody struct {
	// Code - An error code that describes the error condition more precisely than an HTTP status code. Can be used to programatically handle specific error cases.
	Code *string `json:"code,omitempty"`
	// Message - A message that describes the error in detail and provides debugging information.
	Message *string `json:"message,omitempty"`
	// Target - The target of the particular error (for example, the name of the property in error).
	Target *string `json:"target,omitempty"`
	// Details - Contains nested errors that are related to this error.
	Details *[]CloudErrorBody `json:"details,omitempty"`
}

// ListQueryKeysResult response containing the query API keys for a given Azure Search service.
type ListQueryKeysResult struct {
	autorest.Response `json:"-"`
	// Value - The query keys for the Azure Search service.
	Value *[]QueryKey `json:"value,omitempty"`
}

// Operation describes a REST API operation.
type Operation struct {
	// Name - The name of the operation. This name is of the form {provider}/{resource}/{operation}.
	Name *string `json:"name,omitempty"`
	// Display - The object that describes the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay the object that describes the operation.
type OperationDisplay struct {
	// Provider - The friendly name of the resource provider.
	Provider *string `json:"provider,omitempty"`
	// Operation - The operation type: read, write, delete, listKeys/action, etc.
	Operation *string `json:"operation,omitempty"`
	// Resource - The resource type on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
	// Description - The friendly name of the operation.
	Description *string `json:"description,omitempty"`
}

// OperationListResult the result of the request to list REST API operations. It contains a list of operations and
// a URL  to get the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of operations supported by the resource provider.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - The URL to get the next set of operation list results, if any.
	NextLink *string `json:"nextLink,omitempty"`
}

// QueryKey describes an API key for a given Azure Search service that has permissions for query operations only.
type QueryKey struct {
	autorest.Response `json:"-"`
	// Name - The name of the query API key; may be empty.
	Name *string `json:"name,omitempty"`
	// Key - The value of the query API key.
	Key *string `json:"key,omitempty"`
}

// Resource base type for all Azure resources.
type Resource struct {
	// ID - The ID of the resource. This can be used with the Azure Resource Manager to link resources together.
	ID *string `json:"id,omitempty"`
	// Name - The name of the resource.
	Name *string `json:"name,omitempty"`
	// Type - The resource type.
	Type *string `json:"type,omitempty"`
	// Location - The geographic location of the resource. This must be one of the supported and registered Azure Geo Regions (for example, West US, East US, Southeast Asia, and so forth). This property is required when creating a new resource.
	Location *string `json:"location,omitempty"`
	// Tags - Tags to help categorize the resource in the Azure portal.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.ID != nil {
		objectMap["id"] = r.ID
	}
	if r.Name != nil {
		objectMap["name"] = r.Name
	}
	if r.Type != nil {
		objectMap["type"] = r.Type
	}
	if r.Location != nil {
		objectMap["location"] = r.Location
	}
	if r.Tags != nil {
		objectMap["tags"] = r.Tags
	}
	return json.Marshal(objectMap)
}

// Service describes an Azure Search service and its current state.
type Service struct {
	autorest.Response `json:"-"`
	// ServiceProperties - Properties of the Search service.
	*ServiceProperties `json:"properties,omitempty"`
	// Sku - The SKU of the Search Service, which determines price tier and capacity limits. This property is required when creating a new Search Service.
	Sku *Sku `json:"sku,omitempty"`
	// ID - The ID of the resource. This can be used with the Azure Resource Manager to link resources together.
	ID *string `json:"id,omitempty"`
	// Name - The name of the resource.
	Name *string `json:"name,omitempty"`
	// Type - The resource type.
	Type *string `json:"type,omitempty"`
	// Location - The geographic location of the resource. This must be one of the supported and registered Azure Geo Regions (for example, West US, East US, Southeast Asia, and so forth). This property is required when creating a new resource.
	Location *string `json:"location,omitempty"`
	// Tags - Tags to help categorize the resource in the Azure portal.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Service.
func (s Service) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if s.ServiceProperties != nil {
		objectMap["properties"] = s.ServiceProperties
	}
	if s.Sku != nil {
		objectMap["sku"] = s.Sku
	}
	if s.ID != nil {
		objectMap["id"] = s.ID
	}
	if s.Name != nil {
		objectMap["name"] = s.Name
	}
	if s.Type != nil {
		objectMap["type"] = s.Type
	}
	if s.Location != nil {
		objectMap["location"] = s.Location
	}
	if s.Tags != nil {
		objectMap["tags"] = s.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Service struct.
func (s *Service) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var serviceProperties ServiceProperties
				err = json.Unmarshal(*v, &serviceProperties)
				if err != nil {
					return err
				}
				s.ServiceProperties = &serviceProperties
			}
		case "sku":
			if v != nil {
				var sku Sku
				err = json.Unmarshal(*v, &sku)
				if err != nil {
					return err
				}
				s.Sku = &sku
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				s.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				s.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				s.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				s.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				s.Tags = tags
			}
		}
	}

	return nil
}

// ServiceListResult response containing a list of Azure Search services.
type ServiceListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of Search services.
	Value *[]Service `json:"value,omitempty"`
}

// ServiceProperties properties of the Search service.
type ServiceProperties struct {
	// ReplicaCount - The number of replicas in the Search service. If specified, it must be a value between 1 and 12 inclusive for standard SKUs or between 1 and 3 inclusive for basic SKU.
	ReplicaCount *int32 `json:"replicaCount,omitempty"`
	// PartitionCount - The number of partitions in the Search service; if specified, it can be 1, 2, 3, 4, 6, or 12. Values greater than 1 are only valid for standard SKUs. For 'standard3' services with hostingMode set to 'highDensity', the allowed values are between 1 and 3.
	PartitionCount *int32 `json:"partitionCount,omitempty"`
	// HostingMode - Applicable only for the standard3 SKU. You can set this property to enable up to 3 high density partitions that allow up to 1000 indexes, which is much higher than the maximum indexes allowed for any other SKU. For the standard3 SKU, the value is either 'default' or 'highDensity'. For all other SKUs, this value must be 'default'. Possible values include: 'Default', 'HighDensity'
	HostingMode HostingMode `json:"hostingMode,omitempty"`
	// Status - The status of the Search service. Possible values include: 'running': The Search service is running and no provisioning operations are underway. 'provisioning': The Search service is being provisioned or scaled up or down. 'deleting': The Search service is being deleted. 'degraded': The Search service is degraded. This can occur when the underlying search units are not healthy. The Search service is most likely operational, but performance might be slow and some requests might be dropped. 'disabled': The Search service is disabled. In this state, the service will reject all API requests. 'error': The Search service is in an error state. If your service is in the degraded, disabled, or error states, it means the Azure Search team is actively investigating the underlying issue. Dedicated services in these states are still chargeable based on the number of search units provisioned. Possible values include: 'ServiceStatusRunning', 'ServiceStatusProvisioning', 'ServiceStatusDeleting', 'ServiceStatusDegraded', 'ServiceStatusDisabled', 'ServiceStatusError'
	Status ServiceStatus `json:"status,omitempty"`
	// StatusDetails - The details of the Search service status.
	StatusDetails *string `json:"statusDetails,omitempty"`
	// ProvisioningState - The state of the last provisioning operation performed on the Search service. Provisioning is an intermediate state that occurs while service capacity is being established. After capacity is set up, provisioningState changes to either 'succeeded' or 'failed'. Client applications can poll provisioning status (the recommended polling interval is from 30 seconds to one minute) by using the Get Search Service operation to see when an operation is completed. If you are using the free service, this value tends to come back as 'succeeded' directly in the call to Create Search service. This is because the free service uses capacity that is already set up. Possible values include: 'Succeeded', 'Provisioning', 'Failed'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
}

// ServicesCreateOrUpdateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type ServicesCreateOrUpdateFuture struct {
	azure.Future
	req *http.Request
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future ServicesCreateOrUpdateFuture) Result(client ServicesClient) (s Service, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.ServicesCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		return s, azure.NewAsyncOpIncompleteError("search.ServicesCreateOrUpdateFuture")
	}
	if future.PollingMethod() == azure.PollingLocation {
		s, err = client.CreateOrUpdateResponder(future.Response())
		if err != nil {
			err = autorest.NewErrorWithError(err, "search.ServicesCreateOrUpdateFuture", "Result", future.Response(), "Failure responding to request")
		}
		return
	}
	var req *http.Request
	var resp *http.Response
	if future.PollingURL() != "" {
		req, err = http.NewRequest(http.MethodGet, future.PollingURL(), nil)
		if err != nil {
			return
		}
	} else {
		req = autorest.ChangeToGet(future.req)
	}
	resp, err = autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.ServicesCreateOrUpdateFuture", "Result", resp, "Failure sending request")
		return
	}
	s, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.ServicesCreateOrUpdateFuture", "Result", resp, "Failure responding to request")
	}
	return
}

// Sku defines the SKU of an Azure Search Service, which determines price tier and capacity limits.
type Sku struct {
	// Name - The SKU of the Search service. Valid values include: 'free': Shared service. 'basic': Dedicated service with up to 3 replicas. 'standard': Dedicated service with up to 12 partitions and 12 replicas. 'standard2': Similar to standard, but with more capacity per search unit. 'standard3': Offers maximum capacity per search unit with up to 12 partitions and 12 replicas (or up to 3 partitions with more indexes if you also set the hostingMode property to 'highDensity'). Possible values include: 'Free', 'Basic', 'Standard', 'Standard2', 'Standard3'
	Name SkuName `json:"name,omitempty"`
}
