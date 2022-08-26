//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrecommendationsservice

import "time"

// AccountPatchResource - Account resource patch details.
type AccountPatchResource struct {
	// The identity used for the resource.
	Identity *ManagedServiceIdentity `json:"identity,omitempty"`

	// Account resource patch properties.
	Properties *AccountPatchResourceProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`
}

// AccountPatchResourceProperties - Account resource patch properties.
type AccountPatchResourceProperties struct {
	// The list of CORS details.
	Cors []*CorsRule `json:"cors,omitempty"`

	// The list of service endpoints authentication details.
	EndpointAuthentications []*EndpointAuthentication `json:"endpointAuthentications,omitempty"`

	// Connection string to write Accounts reports to.
	ReportsConnectionString *string `json:"reportsConnectionString,omitempty"`
}

// AccountResource - Account resource details.
type AccountResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// The identity used for the resource.
	Identity *ManagedServiceIdentity `json:"identity,omitempty"`

	// Account resource properties.
	Properties *AccountResourceProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// AccountResourceList - The list of RecommendationsService Account resources.
type AccountResourceList struct {
	// The link used to get the next page of RecommendationsService Account resources list.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of RecommendationsService Account resources.
	Value []*AccountResource `json:"value,omitempty"`
}

// AccountResourceProperties - Account resource properties.
type AccountResourceProperties struct {
	// Account configuration. This can only be set at RecommendationsService Account creation.
	Configuration *AccountConfiguration `json:"configuration,omitempty"`

	// The list of CORS details.
	Cors []*CorsRule `json:"cors,omitempty"`

	// The list of service endpoints authentication details.
	EndpointAuthentications []*EndpointAuthentication `json:"endpointAuthentications,omitempty"`

	// Connection string to write Accounts reports to.
	ReportsConnectionString *string `json:"reportsConnectionString,omitempty"`

	// READ-ONLY; The resource provisioning state.
	ProvisioningState *string `json:"provisioningState,omitempty" azure:"ro"`
}

// AccountStatus - Account status.
type AccountStatus struct {
	// The list of scopes statuses.
	ScopesStatuses []*ScopeStatuses `json:"scopesStatuses,omitempty"`
}

// AccountsClientBeginCreateOrUpdateOptions contains the optional parameters for the AccountsClient.BeginCreateOrUpdate method.
type AccountsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AccountsClientBeginDeleteOptions contains the optional parameters for the AccountsClient.BeginDelete method.
type AccountsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AccountsClientBeginUpdateOptions contains the optional parameters for the AccountsClient.BeginUpdate method.
type AccountsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AccountsClientCheckNameAvailabilityOptions contains the optional parameters for the AccountsClient.CheckNameAvailability
// method.
type AccountsClientCheckNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientGetOptions contains the optional parameters for the AccountsClient.Get method.
type AccountsClientGetOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientGetStatusOptions contains the optional parameters for the AccountsClient.GetStatus method.
type AccountsClientGetStatusOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientListByResourceGroupOptions contains the optional parameters for the AccountsClient.ListByResourceGroup method.
type AccountsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientListBySubscriptionOptions contains the optional parameters for the AccountsClient.ListBySubscription method.
type AccountsClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// CheckNameAvailabilityRequest - The check availability request body.
type CheckNameAvailabilityRequest struct {
	// The name of the resource for which availability needs to be checked.
	Name *string `json:"name,omitempty"`

	// The resource type.
	Type *string `json:"type,omitempty"`
}

// CheckNameAvailabilityResponse - The check availability result.
type CheckNameAvailabilityResponse struct {
	// Detailed reason why the given name is available.
	Message *string `json:"message,omitempty"`

	// Indicates if the resource name is available.
	NameAvailable *bool `json:"nameAvailable,omitempty"`

	// The reason why the given name is not available.
	Reason *CheckNameAvailabilityReason `json:"reason,omitempty"`
}

// CorsRule - CORS details.
type CorsRule struct {
	// REQUIRED; The origin domains that are permitted to make a request against the service via CORS.
	AllowedOrigins []*string `json:"allowedOrigins,omitempty"`

	// The request headers that the origin domain may specify on the CORS request.
	AllowedHeaders []*string `json:"allowedHeaders,omitempty"`

	// The methods (HTTP request verbs) that the origin domain may use for a CORS request.
	AllowedMethods []*string `json:"allowedMethods,omitempty"`

	// The response headers to expose to CORS clients.
	ExposedHeaders []*string `json:"exposedHeaders,omitempty"`

	// The number of seconds that the client/browser should cache a preflight response.
	MaxAgeInSeconds *int32 `json:"maxAgeInSeconds,omitempty"`
}

// EndpointAuthentication - Service endpoints authentication details.
type EndpointAuthentication struct {
	// AAD tenant ID.
	AADTenantID *string `json:"aadTenantID,omitempty"`

	// AAD principal ID.
	PrincipalID *string `json:"principalID,omitempty"`

	// AAD principal type.
	PrincipalType *PrincipalType `json:"principalType,omitempty"`
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info interface{} `json:"info,omitempty" azure:"ro"`

	// READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo `json:"additionalInfo,omitempty" azure:"ro"`

	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; The error details.
	Details []*ErrorDetail `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; The error message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The error target.
	Target *string `json:"target,omitempty" azure:"ro"`
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail `json:"error,omitempty"`
}

// ManagedServiceIdentity - Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity struct {
	// REQUIRED; Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
	Type *ManagedServiceIdentityType `json:"type,omitempty"`

	// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM
	// resource ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}.
	// The dictionary values can be empty objects ({}) in
	// requests.
	UserAssignedIdentities map[string]*UserAssignedIdentity `json:"userAssignedIdentities,omitempty"`

	// READ-ONLY; The service principal ID of the system assigned identity. This property will only be provided for a system assigned
	// identity.
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`

	// READ-ONLY; The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}

// ModelingClientBeginCreateOrUpdateOptions contains the optional parameters for the ModelingClient.BeginCreateOrUpdate method.
type ModelingClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ModelingClientBeginDeleteOptions contains the optional parameters for the ModelingClient.BeginDelete method.
type ModelingClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ModelingClientBeginUpdateOptions contains the optional parameters for the ModelingClient.BeginUpdate method.
type ModelingClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ModelingClientGetOptions contains the optional parameters for the ModelingClient.Get method.
type ModelingClientGetOptions struct {
	// placeholder for future optional parameters
}

// ModelingClientListByAccountResourceOptions contains the optional parameters for the ModelingClient.ListByAccountResource
// method.
type ModelingClientListByAccountResourceOptions struct {
	// placeholder for future optional parameters
}

// ModelingInputData - The configuration to raw CDM data to be used as Modeling resource input.
type ModelingInputData struct {
	// Connection string to raw input data.
	ConnectionString *string `json:"connectionString,omitempty"`
}

// ModelingPatchResource - Modeling resource patch details.
type ModelingPatchResource struct {
	// Modeling resource properties to update.
	Properties *ModelingPatchResourceProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`
}

// ModelingPatchResourceProperties - Modeling resource properties to update.
type ModelingPatchResourceProperties struct {
	// The configuration to raw CDM data to be used as Modeling resource input.
	InputData *ModelingInputData `json:"inputData,omitempty"`
}

// ModelingResource - Modeling resource details.
type ModelingResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Modeling resource properties.
	Properties *ModelingResourceProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ModelingResourceList - The list of Modeling resources.
type ModelingResourceList struct {
	// The link used to get the next page of Modeling resources list.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of Modeling resources.
	Value []*ModelingResource `json:"value,omitempty"`
}

// ModelingResourceProperties - Modeling resource properties.
type ModelingResourceProperties struct {
	// Modeling features controls the set of supported scenarios\models being computed. This can only be set at Modeling creation.
	Features *ModelingFeatures `json:"features,omitempty"`

	// Modeling frequency controls the modeling compute frequency.
	Frequency *ModelingFrequency `json:"frequency,omitempty"`

	// The configuration to raw CDM data to be used as Modeling resource input.
	InputData *ModelingInputData `json:"inputData,omitempty"`

	// Modeling size controls the maximum supported input data size.
	Size *ModelingSize `json:"size,omitempty"`

	// READ-ONLY; The resource provisioning state.
	ProvisioningState *string `json:"provisioningState,omitempty" azure:"ro"`
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType `json:"actionType,omitempty" azure:"ro"`

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool `json:"isDataAction,omitempty" azure:"ro"`

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin `json:"origin,omitempty" azure:"ro"`
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation `json:"value,omitempty" azure:"ro"`
}

// OperationStatusResult - The current status of an async operation.
type OperationStatusResult struct {
	// REQUIRED; Operation status.
	Status *string `json:"status,omitempty"`

	// The end time of the operation.
	EndTime *time.Time `json:"endTime,omitempty"`

	// If present, details of the operation error.
	Error *ErrorDetail `json:"error,omitempty"`

	// Fully qualified ID for the async operation.
	ID *string `json:"id,omitempty"`

	// Name of the async operation.
	Name *string `json:"name,omitempty"`

	// The operations list.
	Operations []*OperationStatusResult `json:"operations,omitempty"`

	// Percent of the operation that is complete.
	PercentComplete *float32 `json:"percentComplete,omitempty"`

	// The start time of the operation.
	StartTime *time.Time `json:"startTime,omitempty"`
}

// OperationStatusesClientGetOptions contains the optional parameters for the OperationStatusesClient.Get method.
type OperationStatusesClientGetOptions struct {
	// placeholder for future optional parameters
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ScopeStatuses - Scope statuses.
type ScopeStatuses struct {
	// The scope that the statuses refers to.
	Scope *string `json:"scope,omitempty"`

	// Scope stage statuses.
	Statuses []*StageStatus `json:"statuses,omitempty"`
}

// ServiceEndpointPatchResource - ServiceEndpoint resource patch details.
type ServiceEndpointPatchResource struct {
	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`
}

// ServiceEndpointResource - ServiceEndpoint resource details.
type ServiceEndpointResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// ServiceEndpoint resource properties.
	Properties *ServiceEndpointResourceProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ServiceEndpointResourceList - The list of ServiceEndpoint resources.
type ServiceEndpointResourceList struct {
	// The link used to get the next page of ServiceEndpoint resources list.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of ServiceEndpoint resources.
	Value []*ServiceEndpointResource `json:"value,omitempty"`
}

// ServiceEndpointResourceProperties - ServiceEndpoint resource properties.
type ServiceEndpointResourceProperties struct {
	// ServiceEndpoint pre-allocated capacity controls the maximum requests-per-second allowed for that endpoint. Only applicable
	// when Account configuration is Capacity.
	PreAllocatedCapacity *int32 `json:"preAllocatedCapacity,omitempty"`

	// READ-ONLY; The paired location that will be used by this ServiceEndpoint.
	PairedLocation *string `json:"pairedLocation,omitempty" azure:"ro"`

	// READ-ONLY; The resource provisioning state.
	ProvisioningState *string `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; The URL where the ServiceEndpoint API is accessible at.
	URL *string `json:"url,omitempty" azure:"ro"`
}

// ServiceEndpointsClientBeginCreateOrUpdateOptions contains the optional parameters for the ServiceEndpointsClient.BeginCreateOrUpdate
// method.
type ServiceEndpointsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServiceEndpointsClientBeginDeleteOptions contains the optional parameters for the ServiceEndpointsClient.BeginDelete method.
type ServiceEndpointsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServiceEndpointsClientBeginUpdateOptions contains the optional parameters for the ServiceEndpointsClient.BeginUpdate method.
type ServiceEndpointsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServiceEndpointsClientGetOptions contains the optional parameters for the ServiceEndpointsClient.Get method.
type ServiceEndpointsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ServiceEndpointsClientListByAccountResourceOptions contains the optional parameters for the ServiceEndpointsClient.ListByAccountResource
// method.
type ServiceEndpointsClientListByAccountResourceOptions struct {
	// placeholder for future optional parameters
}

// StageStatus - Stage status.
type StageStatus struct {
	// The stage name.
	Stage *string `json:"stage,omitempty"`

	// The status of the stage.
	Status *string `json:"status,omitempty"`

	// The time of the status.
	Time *time.Time `json:"time,omitempty"`
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags'
// and a 'location'
type TrackedResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// UserAssignedIdentity - User assigned identity properties
type UserAssignedIdentity struct {
	// READ-ONLY; The client ID of the assigned identity.
	ClientID *string `json:"clientId,omitempty" azure:"ro"`

	// READ-ONLY; The principal ID of the assigned identity.
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`
}
