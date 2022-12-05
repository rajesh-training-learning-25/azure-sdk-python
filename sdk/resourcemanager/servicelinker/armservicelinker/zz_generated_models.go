//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicelinker

import "time"

// AuthInfoBaseClassification provides polymorphic access to related types.
// Call the interface's GetAuthInfoBase() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AuthInfoBase, *SecretAuthInfo, *ServicePrincipalCertificateAuthInfo, *ServicePrincipalSecretAuthInfo, *SystemAssignedIdentityAuthInfo,
// - *UserAssignedIdentityAuthInfo
type AuthInfoBaseClassification interface {
	// GetAuthInfoBase returns the AuthInfoBase content of the underlying type.
	GetAuthInfoBase() *AuthInfoBase
}

// AuthInfoBase - The authentication info
type AuthInfoBase struct {
	// REQUIRED; The authentication type.
	AuthType *AuthType `json:"authType,omitempty"`
}

// GetAuthInfoBase implements the AuthInfoBaseClassification interface for type AuthInfoBase.
func (a *AuthInfoBase) GetAuthInfoBase() *AuthInfoBase { return a }

// AzureKeyVaultProperties - The resource properties when type is Azure Key Vault
type AzureKeyVaultProperties struct {
	// REQUIRED; The azure resource type.
	Type *AzureResourceType `json:"type,omitempty"`

	// True if connect via Kubernetes CSI Driver.
	ConnectAsKubernetesCsiDriver *bool `json:"connectAsKubernetesCsiDriver,omitempty"`
}

// GetAzureResourcePropertiesBase implements the AzureResourcePropertiesBaseClassification interface for type AzureKeyVaultProperties.
func (a *AzureKeyVaultProperties) GetAzureResourcePropertiesBase() *AzureResourcePropertiesBase {
	return &AzureResourcePropertiesBase{
		Type: a.Type,
	}
}

// AzureResource - The azure resource info when target service type is AzureResource
type AzureResource struct {
	// REQUIRED; The target service type.
	Type *TargetServiceType `json:"type,omitempty"`

	// The Id of azure resource.
	ID *string `json:"id,omitempty"`

	// The azure resource connection related properties.
	ResourceProperties AzureResourcePropertiesBaseClassification `json:"resourceProperties,omitempty"`
}

// GetTargetServiceBase implements the TargetServiceBaseClassification interface for type AzureResource.
func (a *AzureResource) GetTargetServiceBase() *TargetServiceBase {
	return &TargetServiceBase{
		Type: a.Type,
	}
}

// AzureResourcePropertiesBaseClassification provides polymorphic access to related types.
// Call the interface's GetAzureResourcePropertiesBase() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AzureKeyVaultProperties, *AzureResourcePropertiesBase
type AzureResourcePropertiesBaseClassification interface {
	// GetAzureResourcePropertiesBase returns the AzureResourcePropertiesBase content of the underlying type.
	GetAzureResourcePropertiesBase() *AzureResourcePropertiesBase
}

// AzureResourcePropertiesBase - The azure resource properties
type AzureResourcePropertiesBase struct {
	// REQUIRED; The azure resource type.
	Type *AzureResourceType `json:"type,omitempty"`
}

// GetAzureResourcePropertiesBase implements the AzureResourcePropertiesBaseClassification interface for type AzureResourcePropertiesBase.
func (a *AzureResourcePropertiesBase) GetAzureResourcePropertiesBase() *AzureResourcePropertiesBase {
	return a
}

// ConfluentBootstrapServer - The service properties when target service type is ConfluentBootstrapServer
type ConfluentBootstrapServer struct {
	// REQUIRED; The target service type.
	Type *TargetServiceType `json:"type,omitempty"`

	// The endpoint of service.
	Endpoint *string `json:"endpoint,omitempty"`
}

// GetTargetServiceBase implements the TargetServiceBaseClassification interface for type ConfluentBootstrapServer.
func (c *ConfluentBootstrapServer) GetTargetServiceBase() *TargetServiceBase {
	return &TargetServiceBase{
		Type: c.Type,
	}
}

// ConfluentSchemaRegistry - The service properties when target service type is ConfluentSchemaRegistry
type ConfluentSchemaRegistry struct {
	// REQUIRED; The target service type.
	Type *TargetServiceType `json:"type,omitempty"`

	// The endpoint of service.
	Endpoint *string `json:"endpoint,omitempty"`
}

// GetTargetServiceBase implements the TargetServiceBaseClassification interface for type ConfluentSchemaRegistry.
func (c *ConfluentSchemaRegistry) GetTargetServiceBase() *TargetServiceBase {
	return &TargetServiceBase{
		Type: c.Type,
	}
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

// KeyVaultSecretReferenceSecretInfo - The secret info when type is keyVaultSecretReference. It's for scenario that user provides
// a secret stored in user's keyvault and source is Azure Kubernetes. The key Vault's resource id is linked to
// secretStore.keyVaultId.
type KeyVaultSecretReferenceSecretInfo struct {
	// REQUIRED; The secret type.
	SecretType *SecretType `json:"secretType,omitempty"`

	// Name of the Key Vault secret.
	Name *string `json:"name,omitempty"`

	// Version of the Key Vault secret.
	Version *string `json:"version,omitempty"`
}

// GetSecretInfoBase implements the SecretInfoBaseClassification interface for type KeyVaultSecretReferenceSecretInfo.
func (k *KeyVaultSecretReferenceSecretInfo) GetSecretInfoBase() *SecretInfoBase {
	return &SecretInfoBase{
		SecretType: k.SecretType,
	}
}

// KeyVaultSecretURISecretInfo - The secret info when type is keyVaultSecretUri. It's for scenario that user provides a secret
// stored in user's keyvault and source is Web App, Spring Cloud or Container App.
type KeyVaultSecretURISecretInfo struct {
	// REQUIRED; The secret type.
	SecretType *SecretType `json:"secretType,omitempty"`

	// URI to the keyvault secret
	Value *string `json:"value,omitempty"`
}

// GetSecretInfoBase implements the SecretInfoBaseClassification interface for type KeyVaultSecretURISecretInfo.
func (k *KeyVaultSecretURISecretInfo) GetSecretInfoBase() *SecretInfoBase {
	return &SecretInfoBase{
		SecretType: k.SecretType,
	}
}

// LinkerClientBeginCreateOrUpdateOptions contains the optional parameters for the LinkerClient.BeginCreateOrUpdate method.
type LinkerClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LinkerClientBeginDeleteOptions contains the optional parameters for the LinkerClient.BeginDelete method.
type LinkerClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LinkerClientBeginUpdateOptions contains the optional parameters for the LinkerClient.BeginUpdate method.
type LinkerClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LinkerClientBeginValidateOptions contains the optional parameters for the LinkerClient.BeginValidate method.
type LinkerClientBeginValidateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LinkerClientGetOptions contains the optional parameters for the LinkerClient.Get method.
type LinkerClientGetOptions struct {
	// placeholder for future optional parameters
}

// LinkerClientListConfigurationsOptions contains the optional parameters for the LinkerClient.ListConfigurations method.
type LinkerClientListConfigurationsOptions struct {
	// placeholder for future optional parameters
}

// LinkerClientListOptions contains the optional parameters for the LinkerClient.List method.
type LinkerClientListOptions struct {
	// placeholder for future optional parameters
}

// LinkerList - The list of Linker.
type LinkerList struct {
	// The link used to get the next page of Linker list.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of Linkers.
	Value []*LinkerResource `json:"value,omitempty"`
}

// LinkerPatch - A linker to be updated.
type LinkerPatch struct {
	// Linker properties
	Properties *LinkerProperties `json:"properties,omitempty"`
}

// LinkerProperties - The properties of the linker.
type LinkerProperties struct {
	// The authentication type.
	AuthInfo AuthInfoBaseClassification `json:"authInfo,omitempty"`

	// The application client type
	ClientType *ClientType `json:"clientType,omitempty"`

	// connection scope in source service.
	Scope *string `json:"scope,omitempty"`

	// An option to store secret value in secure place
	SecretStore *SecretStore `json:"secretStore,omitempty"`

	// The target service properties
	TargetService TargetServiceBaseClassification `json:"targetService,omitempty"`

	// The VNet solution.
	VNetSolution *VNetSolution `json:"vNetSolution,omitempty"`

	// READ-ONLY; The provisioning state.
	ProvisioningState *string `json:"provisioningState,omitempty" azure:"ro"`
}

// LinkerResource - Linker of source and target resource
type LinkerResource struct {
	// REQUIRED; The properties of the linker.
	Properties *LinkerProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The system data.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
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

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// ProxyResource - The resource model definition for a Azure Resource Manager proxy resource. It will not have tags and a
// location
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
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

// SecretAuthInfo - The authentication info when authType is secret
type SecretAuthInfo struct {
	// REQUIRED; The authentication type.
	AuthType *AuthType `json:"authType,omitempty"`

	// Username or account name for secret auth.
	Name *string `json:"name,omitempty"`

	// Password or key vault secret for secret auth.
	SecretInfo SecretInfoBaseClassification `json:"secretInfo,omitempty"`
}

// GetAuthInfoBase implements the AuthInfoBaseClassification interface for type SecretAuthInfo.
func (s *SecretAuthInfo) GetAuthInfoBase() *AuthInfoBase {
	return &AuthInfoBase{
		AuthType: s.AuthType,
	}
}

// SecretInfoBaseClassification provides polymorphic access to related types.
// Call the interface's GetSecretInfoBase() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *KeyVaultSecretReferenceSecretInfo, *KeyVaultSecretURISecretInfo, *SecretInfoBase, *ValueSecretInfo
type SecretInfoBaseClassification interface {
	// GetSecretInfoBase returns the SecretInfoBase content of the underlying type.
	GetSecretInfoBase() *SecretInfoBase
}

// SecretInfoBase - The secret info
type SecretInfoBase struct {
	// REQUIRED; The secret type.
	SecretType *SecretType `json:"secretType,omitempty"`
}

// GetSecretInfoBase implements the SecretInfoBaseClassification interface for type SecretInfoBase.
func (s *SecretInfoBase) GetSecretInfoBase() *SecretInfoBase { return s }

// SecretStore - An option to store secret value in secure place
type SecretStore struct {
	// The key vault id to store secret
	KeyVaultID *string `json:"keyVaultId,omitempty"`
}

// ServicePrincipalCertificateAuthInfo - The authentication info when authType is servicePrincipal certificate
type ServicePrincipalCertificateAuthInfo struct {
	// REQUIRED; The authentication type.
	AuthType *AuthType `json:"authType,omitempty"`

	// REQUIRED; ServicePrincipal certificate for servicePrincipal auth.
	Certificate *string `json:"certificate,omitempty"`

	// REQUIRED; Application clientId for servicePrincipal auth.
	ClientID *string `json:"clientId,omitempty"`

	// REQUIRED; Principal Id for servicePrincipal auth.
	PrincipalID *string `json:"principalId,omitempty"`
}

// GetAuthInfoBase implements the AuthInfoBaseClassification interface for type ServicePrincipalCertificateAuthInfo.
func (s *ServicePrincipalCertificateAuthInfo) GetAuthInfoBase() *AuthInfoBase {
	return &AuthInfoBase{
		AuthType: s.AuthType,
	}
}

// ServicePrincipalSecretAuthInfo - The authentication info when authType is servicePrincipal secret
type ServicePrincipalSecretAuthInfo struct {
	// REQUIRED; The authentication type.
	AuthType *AuthType `json:"authType,omitempty"`

	// REQUIRED; ServicePrincipal application clientId for servicePrincipal auth.
	ClientID *string `json:"clientId,omitempty"`

	// REQUIRED; Principal Id for servicePrincipal auth.
	PrincipalID *string `json:"principalId,omitempty"`

	// REQUIRED; Secret for servicePrincipal auth.
	Secret *string `json:"secret,omitempty"`
}

// GetAuthInfoBase implements the AuthInfoBaseClassification interface for type ServicePrincipalSecretAuthInfo.
func (s *ServicePrincipalSecretAuthInfo) GetAuthInfoBase() *AuthInfoBase {
	return &AuthInfoBase{
		AuthType: s.AuthType,
	}
}

// SourceConfiguration - A configuration item for source resource
type SourceConfiguration struct {
	// The name of setting.
	Name *string `json:"name,omitempty"`

	// The value of setting
	Value *string `json:"value,omitempty"`
}

// SourceConfigurationResult - Configurations for source resource, include appSettings, connectionString and serviceBindings
type SourceConfigurationResult struct {
	// The configuration properties for source resource.
	Configurations []*SourceConfiguration `json:"configurations,omitempty"`
}

// SystemAssignedIdentityAuthInfo - The authentication info when authType is systemAssignedIdentity
type SystemAssignedIdentityAuthInfo struct {
	// REQUIRED; The authentication type.
	AuthType *AuthType `json:"authType,omitempty"`
}

// GetAuthInfoBase implements the AuthInfoBaseClassification interface for type SystemAssignedIdentityAuthInfo.
func (s *SystemAssignedIdentityAuthInfo) GetAuthInfoBase() *AuthInfoBase {
	return &AuthInfoBase{
		AuthType: s.AuthType,
	}
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

// TargetServiceBaseClassification provides polymorphic access to related types.
// Call the interface's GetTargetServiceBase() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AzureResource, *ConfluentBootstrapServer, *ConfluentSchemaRegistry, *TargetServiceBase
type TargetServiceBaseClassification interface {
	// GetTargetServiceBase returns the TargetServiceBase content of the underlying type.
	GetTargetServiceBase() *TargetServiceBase
}

// TargetServiceBase - The target service properties
type TargetServiceBase struct {
	// REQUIRED; The target service type.
	Type *TargetServiceType `json:"type,omitempty"`
}

// GetTargetServiceBase implements the TargetServiceBaseClassification interface for type TargetServiceBase.
func (t *TargetServiceBase) GetTargetServiceBase() *TargetServiceBase { return t }

// UserAssignedIdentityAuthInfo - The authentication info when authType is userAssignedIdentity
type UserAssignedIdentityAuthInfo struct {
	// REQUIRED; The authentication type.
	AuthType *AuthType `json:"authType,omitempty"`

	// Client Id for userAssignedIdentity.
	ClientID *string `json:"clientId,omitempty"`

	// Subscription id for userAssignedIdentity.
	SubscriptionID *string `json:"subscriptionId,omitempty"`
}

// GetAuthInfoBase implements the AuthInfoBaseClassification interface for type UserAssignedIdentityAuthInfo.
func (u *UserAssignedIdentityAuthInfo) GetAuthInfoBase() *AuthInfoBase {
	return &AuthInfoBase{
		AuthType: u.AuthType,
	}
}

// VNetSolution - The VNet solution for linker
type VNetSolution struct {
	// Type of VNet solution.
	Type *VNetSolutionType `json:"type,omitempty"`
}

// ValidateOperationResult - The validation operation result for a linker.
type ValidateOperationResult struct {
	// The validation result detail.
	Properties *ValidateResult `json:"properties,omitempty"`

	// Validated linker id.
	ResourceID *string `json:"resourceId,omitempty"`

	// Validation operation status.
	Status *string `json:"status,omitempty"`
}

// ValidateResult - The validation result for a linker.
type ValidateResult struct {
	// The authentication type.
	AuthType *AuthType `json:"authType,omitempty"`

	// A boolean value indicating whether the connection is available or not
	IsConnectionAvailable *bool `json:"isConnectionAvailable,omitempty"`

	// The linker name.
	LinkerName *string `json:"linkerName,omitempty"`

	// The end time of the validation report.
	ReportEndTimeUTC *time.Time `json:"reportEndTimeUtc,omitempty"`

	// The start time of the validation report.
	ReportStartTimeUTC *time.Time `json:"reportStartTimeUtc,omitempty"`

	// The resource id of the linker source application.
	SourceID *string `json:"sourceId,omitempty"`

	// The resource Id of target service.
	TargetID *string `json:"targetId,omitempty"`

	// The detail of validation result
	ValidationDetail []*ValidationResultItem `json:"validationDetail,omitempty"`
}

// ValidationResultItem - The validation item for a linker.
type ValidationResultItem struct {
	// The display name of validation item
	Description *string `json:"description,omitempty"`

	// The error code of validation result
	ErrorCode *string `json:"errorCode,omitempty"`

	// The error message of validation result
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// The validation item name.
	Name *string `json:"name,omitempty"`

	// The result of validation
	Result *ValidationResultStatus `json:"result,omitempty"`
}

// ValueSecretInfo - The secret info when type is rawValue. It's for scenarios that user input the secret.
type ValueSecretInfo struct {
	// REQUIRED; The secret type.
	SecretType *SecretType `json:"secretType,omitempty"`

	// The actual value of the secret.
	Value *string `json:"value,omitempty"`
}

// GetSecretInfoBase implements the SecretInfoBaseClassification interface for type ValueSecretInfo.
func (v *ValueSecretInfo) GetSecretInfoBase() *SecretInfoBase {
	return &SecretInfoBase{
		SecretType: v.SecretType,
	}
}