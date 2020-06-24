// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresources

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// The type of the paths for alias.
type AliasPathType struct {
	// The API versions.
	APIVersions *[]string `json:"apiVersions,omitempty"`

	// The path of an alias.
	Path *string `json:"path,omitempty"`
}

// The alias type.
type AliasType struct {
	// The alias name.
	Name *string `json:"name,omitempty"`

	// The paths for an alias.
	Paths *[]AliasPathType `json:"paths,omitempty"`
}

// Deployment dependency information.
type BasicDependency struct {
	// The ID of the dependency.
	ID *string `json:"id,omitempty"`

	// The dependency resource name.
	ResourceName *string `json:"resourceName,omitempty"`

	// The dependency resource type.
	ResourceType *string `json:"resourceType,omitempty"`
}

// An error response for a resource management request.
type CloudError struct {
	// The resource management error response.
	InnerError *ErrorResponse `json:"error,omitempty"`
}

func (e CloudError) Error() string {
	msg := ""
	if e.InnerError != nil {
		msg += fmt.Sprintf("InnerError: %v\n", *e.InnerError)
	}
	if msg == "" {
		msg = "missing error info"
	}
	return msg
}

type ComponentsSgqdofSchemasIDentityPropertiesUserassignedidentitiesAdditionalproperties struct {
	// The client id of user assigned identity.
	ClientID *string `json:"clientId,omitempty"`

	// The principal id of user assigned identity.
	PrincipalID *string `json:"principalId,omitempty"`
}

// The debug setting.
type DebugSetting struct {
	// Specifies the type of information to log for debugging. The permitted values are none, requestContent, responseContent,
	// or both requestContent and responseContent separated by a comma. The default is none. When setting this value, carefully
	// consider the type of information you are passing in during deployment. By logging information about the request or response,
	// you could potentially expose sensitive data that is retrieved through the deployment operations.
	DetailLevel *string `json:"detailLevel,omitempty"`
}

// Deployment dependency information.
type Dependency struct {
	// The list of dependencies.
	DependsOn *[]BasicDependency `json:"dependsOn,omitempty"`

	// The ID of the dependency.
	ID *string `json:"id,omitempty"`

	// The dependency resource name.
	ResourceName *string `json:"resourceName,omitempty"`

	// The dependency resource type.
	ResourceType *string `json:"resourceType,omitempty"`
}

// Deployment operation parameters.
type Deployment struct {
	// The location to store the deployment data.
	Location *string `json:"location,omitempty"`

	// The deployment properties.
	Properties *DeploymentProperties `json:"properties,omitempty"`
}

// The deployment export result.
type DeploymentExportResult struct {
	// The template content.
	Template *interface{} `json:"template,omitempty"`
}

// DeploymentExportResultResponse is the response envelope for operations that return a DeploymentExportResult type.
type DeploymentExportResultResponse struct {
	// The deployment export result.
	DeploymentExportResult *DeploymentExportResult

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Deployment information.
type DeploymentExtended struct {
	// The ID of the deployment.
	ID *string `json:"id,omitempty"`

	// the location of the deployment.
	Location *string `json:"location,omitempty"`

	// The name of the deployment.
	Name *string `json:"name,omitempty"`

	// Deployment properties.
	Properties *DeploymentPropertiesExtended `json:"properties,omitempty"`

	// The type of the deployment.
	Type *string `json:"type,omitempty"`
}

// Deployment filter.
type DeploymentExtendedFilter struct {
	// The provisioning state.
	ProvisioningState *string `json:"provisioningState,omitempty"`
}

// DeploymentExtendedPollerResponse is the response envelope for operations that asynchronously return a DeploymentExtended
// type.
type DeploymentExtendedPollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (*DeploymentExtendedResponse, error)

	// Poller contains an initialized poller.
	Poller DeploymentExtendedPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DeploymentExtendedResponse is the response envelope for operations that return a DeploymentExtended type.
type DeploymentExtendedResponse struct {
	// Deployment information.
	DeploymentExtended *DeploymentExtended

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// List of deployments.
type DeploymentListResult struct {
	// The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`

	// An array of deployments.
	Value *[]DeploymentExtended `json:"value,omitempty"`
}

// DeploymentListResultResponse is the response envelope for operations that return a DeploymentListResult type.
type DeploymentListResultResponse struct {
	// List of deployments.
	DeploymentListResult *DeploymentListResult

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Deployment operation information.
type DeploymentOperation struct {
	// Full deployment operation ID.
	ID *string `json:"id,omitempty"`

	// Deployment operation ID.
	OperationID *string `json:"operationId,omitempty"`

	// Deployment properties.
	Properties *DeploymentOperationProperties `json:"properties,omitempty"`
}

// Deployment operation properties.
type DeploymentOperationProperties struct {
	// The duration of the operation.
	Duration *string `json:"duration,omitempty"`

	// The state of the provisioning.
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// The HTTP request message.
	Request *HTTPMessage `json:"request,omitempty"`

	// The HTTP response message.
	Response *HTTPMessage `json:"response,omitempty"`

	// Deployment operation service request id.
	ServiceRequestID *string `json:"serviceRequestId,omitempty"`

	// Operation status code.
	StatusCode *string `json:"statusCode,omitempty"`

	// Operation status message.
	StatusMessage *interface{} `json:"statusMessage,omitempty"`

	// The target resource.
	TargetResource *TargetResource `json:"targetResource,omitempty"`

	// The date and time of the operation.
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

func (d DeploymentOperationProperties) MarshalJSON() ([]byte, error) {
	type alias DeploymentOperationProperties
	aux := &struct {
		*alias
		Timestamp *timeRFC3339 `json:"timestamp"`
	}{
		alias:     (*alias)(&d),
		Timestamp: (*timeRFC3339)(d.Timestamp),
	}
	return json.Marshal(aux)
}

func (d *DeploymentOperationProperties) UnmarshalJSON(data []byte) error {
	type alias DeploymentOperationProperties
	aux := &struct {
		*alias
		Timestamp *timeRFC3339 `json:"timestamp"`
	}{
		alias: (*alias)(d),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	d.Timestamp = (*time.Time)(aux.Timestamp)
	return nil
}

// DeploymentOperationResponse is the response envelope for operations that return a DeploymentOperation type.
type DeploymentOperationResponse struct {
	// Deployment operation information.
	DeploymentOperation *DeploymentOperation

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DeploymentOperationsListAtManagementGroupScopeOptions contains the optional parameters for the DeploymentOperations.ListAtManagementGroupScope
// method.
type DeploymentOperationsListAtManagementGroupScopeOptions struct {
	// The number of results to return.
	Top *int32
}

// DeploymentOperationsListAtSubscriptionScopeOptions contains the optional parameters for the DeploymentOperations.ListAtSubscriptionScope
// method.
type DeploymentOperationsListAtSubscriptionScopeOptions struct {
	// The number of results to return.
	Top *int32
}

// DeploymentOperationsListOptions contains the optional parameters for the DeploymentOperations.List method.
type DeploymentOperationsListOptions struct {
	// The number of results to return.
	Top *int32
}

// List of deployment operations.
type DeploymentOperationsListResult struct {
	// The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`

	// An array of deployment operations.
	Value *[]DeploymentOperation `json:"value,omitempty"`
}

// DeploymentOperationsListResultResponse is the response envelope for operations that return a DeploymentOperationsListResult
// type.
type DeploymentOperationsListResultResponse struct {
	// List of deployment operations.
	DeploymentOperationsListResult *DeploymentOperationsListResult

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Deployment properties.
type DeploymentProperties struct {
	// The debug setting of the deployment.
	DebugSetting *DebugSetting `json:"debugSetting,omitempty"`

	// The mode that is used to deploy resources. This value can be either Incremental or Complete. In Incremental mode, resources
	// are deployed without deleting existing resources that are not included in the template. In Complete mode, resources are
	// deployed and existing resources in the resource group that are not included in the template are deleted. Be careful when
	// using Complete mode as you may unintentionally delete resources.
	Mode *DeploymentMode `json:"mode,omitempty"`

	// The deployment on error behavior.
	OnErrorDeployment *OnErrorDeployment `json:"onErrorDeployment,omitempty"`

	// Name and value pairs that define the deployment parameters for the template. You use this element when you want to provide
	// the parameter values directly in the request rather than link to an existing parameter file. Use either the parametersLink
	// property or the parameters property, but not both. It can be a JObject or a well formed JSON string.
	Parameters *interface{} `json:"parameters,omitempty"`

	// The URI of parameters file. You use this element to link to an existing parameters file. Use either the parametersLink
	// property or the parameters property, but not both.
	ParametersLink *ParametersLink `json:"parametersLink,omitempty"`

	// The template content. You use this element when you want to pass the template syntax directly in the request rather than
	// link to an existing template. It can be a JObject or well-formed JSON string. Use either the templateLink property or the
	// template property, but not both.
	Template *interface{} `json:"template,omitempty"`

	// The URI of the template. Use either the templateLink property or the template property, but not both.
	TemplateLink *TemplateLink `json:"templateLink,omitempty"`
}

// Deployment properties with additional details.
type DeploymentPropertiesExtended struct {
	// The correlation ID of the deployment.
	CorrelationID *string `json:"correlationId,omitempty"`

	// The debug setting of the deployment.
	DebugSetting *DebugSetting `json:"debugSetting,omitempty"`

	// The list of deployment dependencies.
	Dependencies *[]Dependency `json:"dependencies,omitempty"`

	// The duration of the template deployment.
	Duration *string `json:"duration,omitempty"`

	// The deployment mode. Possible values are Incremental and Complete.
	Mode *DeploymentMode `json:"mode,omitempty"`

	// The deployment on error behavior.
	OnErrorDeployment *OnErrorDeploymentExtended `json:"onErrorDeployment,omitempty"`

	// Key/value pairs that represent deployment output.
	Outputs *interface{} `json:"outputs,omitempty"`

	// Deployment parameters. Use only one of Parameters or ParametersLink.
	Parameters *interface{} `json:"parameters,omitempty"`

	// The URI referencing the parameters. Use only one of Parameters or ParametersLink.
	ParametersLink *ParametersLink `json:"parametersLink,omitempty"`

	// The list of resource providers needed for the deployment.
	Providers *[]Provider `json:"providers,omitempty"`

	// The state of the provisioning.
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// The template content. Use only one of Template or TemplateLink.
	Template *interface{} `json:"template,omitempty"`

	// The URI referencing the template. Use only one of Template or TemplateLink.
	TemplateLink *TemplateLink `json:"templateLink,omitempty"`

	// The timestamp of the template deployment.
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

func (d DeploymentPropertiesExtended) MarshalJSON() ([]byte, error) {
	type alias DeploymentPropertiesExtended
	aux := &struct {
		*alias
		Timestamp *timeRFC3339 `json:"timestamp"`
	}{
		alias:     (*alias)(&d),
		Timestamp: (*timeRFC3339)(d.Timestamp),
	}
	return json.Marshal(aux)
}

func (d *DeploymentPropertiesExtended) UnmarshalJSON(data []byte) error {
	type alias DeploymentPropertiesExtended
	aux := &struct {
		*alias
		Timestamp *timeRFC3339 `json:"timestamp"`
	}{
		alias: (*alias)(d),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	d.Timestamp = (*time.Time)(aux.Timestamp)
	return nil
}

// Information from validate template deployment response.
type DeploymentValidateResult struct {
	// Validation error.
	Error *ResourceManagementErrorWithDetails `json:"error,omitempty"`

	// The template deployment properties.
	Properties *DeploymentPropertiesExtended `json:"properties,omitempty"`
}

// DeploymentValidateResultResponse is the response envelope for operations that return a DeploymentValidateResult type.
type DeploymentValidateResultResponse struct {
	// Information from validate template deployment response.
	DeploymentValidateResult *DeploymentValidateResult

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DeploymentsListAtManagementGroupScopeOptions contains the optional parameters for the Deployments.ListAtManagementGroupScope
// method.
type DeploymentsListAtManagementGroupScopeOptions struct {
	// The filter to apply on the operation. For example, you can use $filter=provisioningState eq '{state}'.
	Filter *string
	// The number of results to get. If null is passed, returns all deployments.
	Top *int32
}

// DeploymentsListAtSubscriptionScopeOptions contains the optional parameters for the Deployments.ListAtSubscriptionScope
// method.
type DeploymentsListAtSubscriptionScopeOptions struct {
	// The filter to apply on the operation. For example, you can use $filter=provisioningState eq '{state}'.
	Filter *string
	// The number of results to get. If null is passed, returns all deployments.
	Top *int32
}

// DeploymentsListByResourceGroupOptions contains the optional parameters for the Deployments.ListByResourceGroup method.
type DeploymentsListByResourceGroupOptions struct {
	// The filter to apply on the operation. For example, you can use $filter=provisioningState eq '{state}'.
	Filter *string
	// The number of results to get. If null is passed, returns all deployments.
	Top *int32
}

// The resource management error additional info.
type ErrorAdditionalInfo struct {
	// The additional info.
	Info *interface{} `json:"info,omitempty"`

	// The additional info type.
	Type *string `json:"type,omitempty"`
}

// The resource management error response.
type ErrorResponse struct {
	// The error additional info.
	AdditionalInfo *[]ErrorAdditionalInfo `json:"additionalInfo,omitempty"`

	// The error code.
	Code *string `json:"code,omitempty"`

	// The error details.
	Details *[]ErrorResponse `json:"details,omitempty"`

	// The error message.
	Message *string `json:"message,omitempty"`

	// The error target.
	Target *string `json:"target,omitempty"`
}

// Export resource group template request parameters.
type ExportTemplateRequest struct {
	// The export template options. A CSV-formatted list containing zero or more of the following: 'IncludeParameterDefaultValue',
	// 'IncludeComments', 'SkipResourceNameParameterization', 'SkipAllParameterization'
	Options *string `json:"options,omitempty"`

	// The IDs of the resources to filter the export by. To export all resources, supply an array with single entry '*'.
	Resources *[]string `json:"resources,omitempty"`
}

// Resource information.
type GenericResource struct {
	Resource
	// The identity of the resource.
	IDentity *IDentity `json:"identity,omitempty"`

	// The kind of the resource.
	Kind *string `json:"kind,omitempty"`

	// ID of the resource that manages this resource.
	ManagedBy *string `json:"managedBy,omitempty"`

	// The plan of the resource.
	Plan *Plan `json:"plan,omitempty"`

	// The resource properties.
	Properties *interface{} `json:"properties,omitempty"`

	// The SKU of the resource.
	Sku *Sku `json:"sku,omitempty"`
}

// Resource information.
type GenericResourceExpanded struct {
	GenericResource
	// The changed time of the resource. This is only present if requested via the $expand query parameter.
	ChangedTime *time.Time `json:"changedTime,omitempty"`

	// The created time of the resource. This is only present if requested via the $expand query parameter.
	CreatedTime *time.Time `json:"createdTime,omitempty"`

	// The provisioning state of the resource. This is only present if requested via the $expand query parameter.
	ProvisioningState *string `json:"provisioningState,omitempty"`
}

func (g GenericResourceExpanded) MarshalJSON() ([]byte, error) {
	type alias GenericResourceExpanded
	aux := &struct {
		*alias
		ChangedTime *timeRFC3339 `json:"changedTime"`
		CreatedTime *timeRFC3339 `json:"createdTime"`
	}{
		alias:       (*alias)(&g),
		ChangedTime: (*timeRFC3339)(g.ChangedTime),
		CreatedTime: (*timeRFC3339)(g.CreatedTime),
	}
	return json.Marshal(aux)
}

func (g *GenericResourceExpanded) UnmarshalJSON(data []byte) error {
	type alias GenericResourceExpanded
	aux := &struct {
		*alias
		ChangedTime *timeRFC3339 `json:"changedTime"`
		CreatedTime *timeRFC3339 `json:"createdTime"`
	}{
		alias: (*alias)(g),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	g.ChangedTime = (*time.Time)(aux.ChangedTime)
	g.CreatedTime = (*time.Time)(aux.CreatedTime)
	return nil
}

// Resource filter.
type GenericResourceFilter struct {
	// The resource type.
	ResourceType *string `json:"resourceType,omitempty"`

	// The tag name.
	Tagname *string `json:"tagname,omitempty"`

	// The tag value.
	Tagvalue *string `json:"tagvalue,omitempty"`
}

// GenericResourcePollerResponse is the response envelope for operations that asynchronously return a GenericResource type.
type GenericResourcePollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (*GenericResourceResponse, error)

	// Poller contains an initialized poller.
	Poller GenericResourcePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// GenericResourceResponse is the response envelope for operations that return a GenericResource type.
type GenericResourceResponse struct {
	// Resource information.
	GenericResource *GenericResource

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HTTP message.
type HTTPMessage struct {
	// HTTP message content.
	Content *interface{} `json:"content,omitempty"`
}

// HTTPPollerResponse contains the asynchronous HTTP response from the call to the service endpoint.
type HTTPPollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (*http.Response, error)

	// Poller contains an initialized poller.
	Poller HTTPPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Identity for the resource.
type IDentity struct {
	// The principal ID of resource identity.
	PrincipalID *string `json:"principalId,omitempty"`

	// The tenant ID of resource.
	TenantID *string `json:"tenantId,omitempty"`

	// The identity type.
	Type *ResourceIdentityType `json:"type,omitempty"`

	// The list of user identities associated with the resource. The user identity dictionary key references will be ARM resource
	// ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
	UserAssignedIDentities *map[string]ComponentsSgqdofSchemasIDentityPropertiesUserassignedidentitiesAdditionalproperties `json:"userAssignedIdentities,omitempty"`
}

// Deployment on error behavior.
type OnErrorDeployment struct {
	// The deployment to be used on error case.
	DeploymentName *string `json:"deploymentName,omitempty"`

	// The deployment on error behavior type. Possible values are LastSuccessful and SpecificDeployment.
	Type *OnErrorDeploymentType `json:"type,omitempty"`
}

// Deployment on error behavior with additional details.
type OnErrorDeploymentExtended struct {
	// The deployment to be used on error case.
	DeploymentName *string `json:"deploymentName,omitempty"`

	// The state of the provisioning for the on error deployment.
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// The deployment on error behavior type. Possible values are LastSuccessful and SpecificDeployment.
	Type *OnErrorDeploymentType `json:"type,omitempty"`
}

// Microsoft.Resources operation
type Operation struct {
	// The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
}

// The object that represents the operation.
type OperationDisplay struct {
	// Description of the operation.
	Description *string `json:"description,omitempty"`

	// Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`

	// Service provider: Microsoft.Resources
	Provider *string `json:"provider,omitempty"`

	// Resource on which the operation is performed: Profile, endpoint, etc.
	Resource *string `json:"resource,omitempty"`
}

// Result of the request to list Microsoft.Resources operations. It contains a list of operations and a URL link to get the
// next set of results.
type OperationListResult struct {
	// URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`

	// List of Microsoft.Resources operations.
	Value *[]Operation `json:"value,omitempty"`
}

// OperationListResultResponse is the response envelope for operations that return a OperationListResult type.
type OperationListResultResponse struct {
	// Result of the request to list Microsoft.Resources operations. It contains a list of operations and a URL link to get the
	// next set of results.
	OperationListResult *OperationListResult

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Entity representing the reference to the deployment parameters.
type ParametersLink struct {
	// If included, must match the ContentVersion in the template.
	ContentVersion *string `json:"contentVersion,omitempty"`

	// The URI of the parameters file.
	URI *string `json:"uri,omitempty"`
}

// Plan for the resource.
type Plan struct {
	// The plan ID.
	Name *string `json:"name,omitempty"`

	// The offer ID.
	Product *string `json:"product,omitempty"`

	// The promotion code.
	PromotionCode *string `json:"promotionCode,omitempty"`

	// The publisher ID.
	Publisher *string `json:"publisher,omitempty"`

	// The plan's version.
	Version *string `json:"version,omitempty"`
}

// Resource provider information.
type Provider struct {
	// The provider ID.
	ID *string `json:"id,omitempty"`

	// The namespace of the resource provider.
	Namespace *string `json:"namespace,omitempty"`

	// The registration policy of the resource provider.
	RegistrationPolicy *string `json:"registrationPolicy,omitempty"`

	// The registration state of the resource provider.
	RegistrationState *string `json:"registrationState,omitempty"`

	// The collection of provider resource types.
	ResourceTypes *[]ProviderResourceType `json:"resourceTypes,omitempty"`
}

// List of resource providers.
type ProviderListResult struct {
	// The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`

	// An array of resource providers.
	Value *[]Provider `json:"value,omitempty"`
}

// ProviderListResultResponse is the response envelope for operations that return a ProviderListResult type.
type ProviderListResultResponse struct {
	// List of resource providers.
	ProviderListResult *ProviderListResult

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Resource type managed by the resource provider.
type ProviderResourceType struct {
	// The API version.
	APIVersions *[]string `json:"apiVersions,omitempty"`

	// The aliases that are supported by this resource type.
	Aliases *[]AliasType `json:"aliases,omitempty"`

	// The additional capabilities offered by this resource type.
	Capabilities *string `json:"capabilities,omitempty"`

	// The collection of locations where this resource type can be created.
	Locations *[]string `json:"locations,omitempty"`

	// The properties.
	Properties *map[string]string `json:"properties,omitempty"`

	// The resource type.
	ResourceType *string `json:"resourceType,omitempty"`
}

// ProviderResponse is the response envelope for operations that return a Provider type.
type ProviderResponse struct {
	// Resource provider information.
	Provider *Provider

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ProvidersGetOptions contains the optional parameters for the Providers.Get method.
type ProvidersGetOptions struct {
	// The $expand query parameter. For example, to include property aliases in response, use $expand=resourceTypes/aliases.
	Expand *string
}

// ProvidersListOptions contains the optional parameters for the Providers.List method.
type ProvidersListOptions struct {
	// The properties to include in the results. For example, use &$expand=metadata in the query string to retrieve resource provider
	// metadata. To include property aliases in response, use $expand=resourceTypes/aliases.
	Expand *string
	// The number of results to return. If null is passed returns all deployments.
	Top *int32
}

// Specified resource.
type Resource struct {
	// Resource ID
	ID *string `json:"id,omitempty"`

	// Resource location
	Location *string `json:"location,omitempty"`

	// Resource name
	Name *string `json:"name,omitempty"`

	// Resource tags
	Tags *map[string]string `json:"tags,omitempty"`

	// Resource type
	Type *string `json:"type,omitempty"`
}

// Resource group information.
type ResourceGroup struct {
	// The ID of the resource group.
	ID *string `json:"id,omitempty"`

	// The location of the resource group. It cannot be changed after the resource group has been created. It must be one of the
	// supported Azure locations.
	Location *string `json:"location,omitempty"`

	// The ID of the resource that manages this resource group.
	ManagedBy *string `json:"managedBy,omitempty"`

	// The name of the resource group.
	Name *string `json:"name,omitempty"`

	// The resource group properties.
	Properties *ResourceGroupProperties `json:"properties,omitempty"`

	// The tags attached to the resource group.
	Tags *map[string]string `json:"tags,omitempty"`

	// The type of the resource group.
	Type *string `json:"type,omitempty"`
}

// Resource group export result.
type ResourceGroupExportResult struct {
	// The error.
	Error *ResourceManagementErrorWithDetails `json:"error,omitempty"`

	// The template content.
	Template *interface{} `json:"template,omitempty"`
}

// ResourceGroupExportResultResponse is the response envelope for operations that return a ResourceGroupExportResult type.
type ResourceGroupExportResultResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Resource group export result.
	ResourceGroupExportResult *ResourceGroupExportResult
}

// Resource group filter.
type ResourceGroupFilter struct {
	// The tag name.
	TagName *string `json:"tagName,omitempty"`

	// The tag value.
	TagValue *string `json:"tagValue,omitempty"`
}

// List of resource groups.
type ResourceGroupListResult struct {
	// The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`

	// An array of resource groups.
	Value *[]ResourceGroup `json:"value,omitempty"`
}

// ResourceGroupListResultResponse is the response envelope for operations that return a ResourceGroupListResult type.
type ResourceGroupListResultResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// List of resource groups.
	ResourceGroupListResult *ResourceGroupListResult
}

// Resource group information.
type ResourceGroupPatchable struct {
	// The ID of the resource that manages this resource group.
	ManagedBy *string `json:"managedBy,omitempty"`

	// The name of the resource group.
	Name *string `json:"name,omitempty"`

	// The resource group properties.
	Properties *ResourceGroupProperties `json:"properties,omitempty"`

	// The tags attached to the resource group.
	Tags *map[string]string `json:"tags,omitempty"`
}

// The resource group properties.
type ResourceGroupProperties struct {
	// The provisioning state.
	ProvisioningState *string `json:"provisioningState,omitempty"`
}

// ResourceGroupResponse is the response envelope for operations that return a ResourceGroup type.
type ResourceGroupResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Resource group information.
	ResourceGroup *ResourceGroup
}

// ResourceGroupsListOptions contains the optional parameters for the ResourceGroups.List method.
type ResourceGroupsListOptions struct {
	// The filter to apply on the operation.<br><br>You can filter by tag names and values. For example, to filter for a tag name
	// and value, use $filter=tagName eq 'tag1' and tagValue eq 'Value1'
	Filter *string
	// The number of results to return. If null is passed, returns all resource groups.
	Top *int32
}

// List of resource groups.
type ResourceListResult struct {
	// The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`

	// An array of resources.
	Value *[]GenericResourceExpanded `json:"value,omitempty"`
}

// ResourceListResultResponse is the response envelope for operations that return a ResourceListResult type.
type ResourceListResultResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// List of resource groups.
	ResourceListResult *ResourceListResult
}

// The detailed error message of resource management.
type ResourceManagementErrorWithDetails struct {
	// The error code returned when exporting the template.
	Code *string `json:"code,omitempty"`

	// Validation error.
	Details *[]ResourceManagementErrorWithDetails `json:"details,omitempty"`

	// The error message describing the export error.
	Message *string `json:"message,omitempty"`

	// The target of the error.
	Target *string `json:"target,omitempty"`
}

// Resource provider operation's display properties.
type ResourceProviderOperationDisplayProperties struct {
	// Operation description.
	Description *string `json:"description,omitempty"`

	// Resource provider operation.
	Operation *string `json:"operation,omitempty"`

	// Operation provider.
	Provider *string `json:"provider,omitempty"`

	// Operation description.
	Publisher *string `json:"publisher,omitempty"`

	// Operation resource.
	Resource *string `json:"resource,omitempty"`
}

// ResourcesListByResourceGroupOptions contains the optional parameters for the Resources.ListByResourceGroup method.
type ResourcesListByResourceGroupOptions struct {
	// Comma-separated list of additional properties to be included in the response. Valid values include `createdTime`, `changedTime`
	// and `provisioningState`. For example, `$expand=createdTime,changedTime`.
	Expand *string
	// The filter to apply on the operation.<br><br>The properties you can use for eq (equals) or ne (not equals) are: location,
	// resourceType, name, resourceGroup, identity, identity/principalId, plan, plan/publisher, plan/product, plan/name, plan/version,
	// and plan/promotionCode.<br><br>For example, to filter by a resource type, use: $filter=resourceType eq 'Microsoft.Network/virtualNetworks'<br><br>You
	// can use substringof(value, property) in the filter. The properties you can use for substring are: name and resourceGroup.<br><br>For
	// example, to get all resources with 'demo' anywhere in the name, use: $filter=substringof('demo', name)<br><br>You can link
	// more than one substringof together by adding and/or operators.<br><br>You can filter by tag names and values. For example,
	// to filter for a tag name and value, use $filter=tagName eq 'tag1' and tagValue eq 'Value1'<br><br>You can use some properties
	// together when filtering. The combinations you can use are: substringof and/or resourceType, plan and plan/publisher and
	// plan/name, identity and identity/principalId.
	Filter *string
	// The number of results to return. If null is passed, returns all resources.
	Top *int32
}

// ResourcesListOptions contains the optional parameters for the Resources.List method.
type ResourcesListOptions struct {
	// Comma-separated list of additional properties to be included in the response. Valid values include `createdTime`, `changedTime`
	// and `provisioningState`. For example, `$expand=createdTime,changedTime`.
	Expand *string
	// The filter to apply on the operation.<br><br>The properties you can use for eq (equals) or ne (not equals) are: location,
	// resourceType, name, resourceGroup, identity, identity/principalId, plan, plan/publisher, plan/product, plan/name, plan/version,
	// and plan/promotionCode.<br><br>For example, to filter by a resource type, use: $filter=resourceType eq 'Microsoft.Network/virtualNetworks'<br><br>You
	// can use substringof(value, property) in the filter. The properties you can use for substring are: name and resourceGroup.<br><br>For
	// example, to get all resources with 'demo' anywhere in the name, use: $filter=substringof('demo', name)<br><br>You can link
	// more than one substringof together by adding and/or operators.<br><br>You can filter by tag names and values. For example,
	// to filter for a tag name and value, use $filter=tagName eq 'tag1' and tagValue eq 'Value1'<br><br>You can use some properties
	// together when filtering. The combinations you can use are: substringof and/or resourceType, plan and plan/publisher and
	// plan/name, identity and identity/principalId.
	Filter *string
	// The number of results to return. If null is passed, returns all resource groups.
	Top *int32
}

// Parameters of move resources.
type ResourcesMoveInfo struct {
	// The IDs of the resources.
	Resources *[]string `json:"resources,omitempty"`

	// The target resource group.
	TargetResourceGroup *string `json:"targetResourceGroup,omitempty"`
}

// SKU for the resource.
type Sku struct {
	// The SKU capacity.
	Capacity *int32 `json:"capacity,omitempty"`

	// The SKU family.
	Family *string `json:"family,omitempty"`

	// The SKU model.
	Model *string `json:"model,omitempty"`

	// The SKU name.
	Name *string `json:"name,omitempty"`

	// The SKU size.
	Size *string `json:"size,omitempty"`

	// The SKU tier.
	Tier *string `json:"tier,omitempty"`
}

// Sub-resource.
type SubResource struct {
	// Resource ID
	ID *string `json:"id,omitempty"`
}

// Tag count.
type TagCount struct {
	// Type of count.
	Type *string `json:"type,omitempty"`

	// Value of count.
	Value *int32 `json:"value,omitempty"`
}

// Tag details.
type TagDetails struct {
	// The total number of resources that use the resource tag. When a tag is initially created and has no associated resources,
	// the value is 0.
	Count *TagCount `json:"count,omitempty"`

	// The tag ID.
	ID *string `json:"id,omitempty"`

	// The tag name.
	TagName *string `json:"tagName,omitempty"`

	// The list of tag values.
	Values *[]TagValue `json:"values,omitempty"`
}

// TagDetailsResponse is the response envelope for operations that return a TagDetails type.
type TagDetailsResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Tag details.
	TagDetails *TagDetails
}

// Tag information.
type TagValue struct {
	// The tag value count.
	Count *TagCount `json:"count,omitempty"`

	// The tag ID.
	ID *string `json:"id,omitempty"`

	// The tag value.
	TagValue *string `json:"tagValue,omitempty"`
}

// TagValueResponse is the response envelope for operations that return a TagValue type.
type TagValueResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Tag information.
	TagValue *TagValue
}

// List of subscription tags.
type TagsListResult struct {
	// The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`

	// An array of tags.
	Value *[]TagDetails `json:"value,omitempty"`
}

// TagsListResultResponse is the response envelope for operations that return a TagsListResult type.
type TagsListResultResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// List of subscription tags.
	TagsListResult *TagsListResult
}

// Target resource.
type TargetResource struct {
	// The ID of the resource.
	ID *string `json:"id,omitempty"`

	// The name of the resource.
	ResourceName *string `json:"resourceName,omitempty"`

	// The type of the resource.
	ResourceType *string `json:"resourceType,omitempty"`
}

// Result of the request to calculate template hash. It contains a string of minified template and its hash.
type TemplateHashResult struct {
	// The minified template string.
	MinifiedTemplate *string `json:"minifiedTemplate,omitempty"`

	// The template hash.
	TemplateHash *string `json:"templateHash,omitempty"`
}

// TemplateHashResultResponse is the response envelope for operations that return a TemplateHashResult type.
type TemplateHashResultResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Result of the request to calculate template hash. It contains a string of minified template and its hash.
	TemplateHashResult *TemplateHashResult
}

// Entity representing the reference to the template.
type TemplateLink struct {
	// If included, must match the ContentVersion in the template.
	ContentVersion *string `json:"contentVersion,omitempty"`

	// The URI of the template to deploy.
	URI *string `json:"uri,omitempty"`
}
