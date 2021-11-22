//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdomainservices

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
	"time"
)

// CloudError - An error response from the Domain Services.
// Implements the error and azcore.HTTPResponse interfaces.
type CloudError struct {
	raw string
	// An error response from the Domain Services.
	InnerError *CloudErrorBody `json:"error,omitempty"`
}

// Error implements the error interface for type CloudError.
// The contents of the error text are not contractual and subject to change.
func (e CloudError) Error() string {
	return e.raw
}

// CloudErrorBody - An error response from the Domain Services.
type CloudErrorBody struct {
	// An identifier for the error. Codes are invariant and are intended to be consumed programmatically.
	Code *string `json:"code,omitempty"`

	// A list of additional details about the error.
	Details []*CloudErrorBody `json:"details,omitempty"`

	// A message describing the error, intended to be suitable for display in a user interface.
	Message *string `json:"message,omitempty"`

	// The target of the particular error. For example, the name of the property in error.
	Target *string `json:"target,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type CloudErrorBody.
func (c CloudErrorBody) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", c.Code)
	populate(objectMap, "details", c.Details)
	populate(objectMap, "message", c.Message)
	populate(objectMap, "target", c.Target)
	return json.Marshal(objectMap)
}

// ConfigDiagnostics - Configuration Diagnostics
type ConfigDiagnostics struct {
	// Last domain configuration diagnostics DateTime
	LastExecuted *time.Time `json:"lastExecuted,omitempty"`

	// List of Configuration Diagnostics validator results.
	ValidatorResults []*ConfigDiagnosticsValidatorResult `json:"validatorResults,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ConfigDiagnostics.
func (c ConfigDiagnostics) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC1123(objectMap, "lastExecuted", c.LastExecuted)
	populate(objectMap, "validatorResults", c.ValidatorResults)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ConfigDiagnostics.
func (c *ConfigDiagnostics) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "lastExecuted":
			err = unpopulateTimeRFC1123(val, &c.LastExecuted)
			delete(rawMsg, key)
		case "validatorResults":
			err = unpopulate(val, &c.ValidatorResults)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// ConfigDiagnosticsValidatorResult - Config Diagnostics validator result data
type ConfigDiagnosticsValidatorResult struct {
	// List of resource config validation issues.
	Issues []*ConfigDiagnosticsValidatorResultIssue `json:"issues,omitempty"`

	// Replica set location and subnet name
	ReplicaSetSubnetDisplayName *string `json:"replicaSetSubnetDisplayName,omitempty"`

	// Status for individual validator after running diagnostics.
	Status *Status `json:"status,omitempty"`

	// Validator identifier
	ValidatorID *string `json:"validatorId,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ConfigDiagnosticsValidatorResult.
func (c ConfigDiagnosticsValidatorResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "issues", c.Issues)
	populate(objectMap, "replicaSetSubnetDisplayName", c.ReplicaSetSubnetDisplayName)
	populate(objectMap, "status", c.Status)
	populate(objectMap, "validatorId", c.ValidatorID)
	return json.Marshal(objectMap)
}

// ConfigDiagnosticsValidatorResultIssue - Specific issue for a particular config diagnostics validator
type ConfigDiagnosticsValidatorResultIssue struct {
	// List of domain resource property name or values used to compose a rich description.
	DescriptionParams []*string `json:"descriptionParams,omitempty"`

	// Validation issue identifier.
	ID *string `json:"id,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ConfigDiagnosticsValidatorResultIssue.
func (c ConfigDiagnosticsValidatorResultIssue) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "descriptionParams", c.DescriptionParams)
	populate(objectMap, "id", c.ID)
	return json.Marshal(objectMap)
}

// ContainerAccount - Container Account Description
type ContainerAccount struct {
	// The account name
	AccountName *string `json:"accountName,omitempty"`

	// The account password
	Password *string `json:"password,omitempty"`

	// The account spn
	Spn *string `json:"spn,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ContainerAccount.
func (c ContainerAccount) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "accountName", c.AccountName)
	populate(objectMap, "password", c.Password)
	populate(objectMap, "spn", c.Spn)
	return json.Marshal(objectMap)
}

// DomainSecuritySettings - Domain Security Settings
type DomainSecuritySettings struct {
	// A flag to determine whether or not KerberosArmoring is enabled or disabled.
	KerberosArmoring *KerberosArmoring `json:"kerberosArmoring,omitempty"`

	// A flag to determine whether or not KerberosRc4Encryption is enabled or disabled.
	KerberosRc4Encryption *KerberosRc4Encryption `json:"kerberosRc4Encryption,omitempty"`

	// A flag to determine whether or not NtlmV1 is enabled or disabled.
	NtlmV1 *NtlmV1 `json:"ntlmV1,omitempty"`

	// A flag to determine whether or not SyncKerberosPasswords is enabled or disabled.
	SyncKerberosPasswords *SyncKerberosPasswords `json:"syncKerberosPasswords,omitempty"`

	// A flag to determine whether or not SyncNtlmPasswords is enabled or disabled.
	SyncNtlmPasswords *SyncNtlmPasswords `json:"syncNtlmPasswords,omitempty"`

	// A flag to determine whether or not SyncOnPremPasswords is enabled or disabled.
	SyncOnPremPasswords *SyncOnPremPasswords `json:"syncOnPremPasswords,omitempty"`

	// A flag to determine whether or not TlsV1 is enabled or disabled.
	TLSV1 *TLSV1 `json:"tlsV1,omitempty"`
}

// DomainService - Domain service.
type DomainService struct {
	Resource
	// Domain service properties
	Properties *DomainServiceProperties `json:"properties,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type DomainService.
func (d DomainService) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	d.Resource.marshalInternal(objectMap)
	populate(objectMap, "properties", d.Properties)
	return json.Marshal(objectMap)
}

// DomainServiceListResult - The response from the List Domain Services operation.
type DomainServiceListResult struct {
	// the list of domain services.
	Value []*DomainService `json:"value,omitempty"`

	// READ-ONLY; The continuation token for the next page of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type DomainServiceListResult.
func (d DomainServiceListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", d.NextLink)
	populate(objectMap, "value", d.Value)
	return json.Marshal(objectMap)
}

// DomainServiceOperationsListOptions contains the optional parameters for the DomainServiceOperations.List method.
type DomainServiceOperationsListOptions struct {
	// placeholder for future optional parameters
}

// DomainServiceProperties - Properties of the Domain Service.
type DomainServiceProperties struct {
	// Configuration diagnostics data containing latest execution from client.
	ConfigDiagnostics *ConfigDiagnostics `json:"configDiagnostics,omitempty"`

	// Domain Configuration Type
	DomainConfigurationType *string `json:"domainConfigurationType,omitempty"`

	// The name of the Azure domain that the user would like to deploy Domain Services to.
	DomainName *string `json:"domainName,omitempty"`

	// DomainSecurity Settings
	DomainSecuritySettings *DomainSecuritySettings `json:"domainSecuritySettings,omitempty"`

	// Enabled or Disabled flag to turn on Group-based filtered sync
	FilteredSync *FilteredSync `json:"filteredSync,omitempty"`

	// Secure LDAP Settings
	LdapsSettings *LdapsSettings `json:"ldapsSettings,omitempty"`

	// Notification Settings
	NotificationSettings *NotificationSettings `json:"notificationSettings,omitempty"`

	// List of ReplicaSets
	ReplicaSets []*ReplicaSet `json:"replicaSets,omitempty"`

	// Resource Forest Settings
	ResourceForestSettings *ResourceForestSettings `json:"resourceForestSettings,omitempty"`

	// Sku Type
	SKU *string `json:"sku,omitempty"`

	// READ-ONLY; Deployment Id
	DeploymentID *string `json:"deploymentId,omitempty" azure:"ro"`

	// READ-ONLY; Migration Properties
	MigrationProperties *MigrationProperties `json:"migrationProperties,omitempty" azure:"ro"`

	// READ-ONLY; the current deployment or provisioning state, which only appears in the response.
	ProvisioningState *string `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; SyncOwner ReplicaSet Id
	SyncOwner *string `json:"syncOwner,omitempty" azure:"ro"`

	// READ-ONLY; Azure Active Directory Tenant Id
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`

	// READ-ONLY; Data Model Version
	Version *int32 `json:"version,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type DomainServiceProperties.
func (d DomainServiceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "configDiagnostics", d.ConfigDiagnostics)
	populate(objectMap, "deploymentId", d.DeploymentID)
	populate(objectMap, "domainConfigurationType", d.DomainConfigurationType)
	populate(objectMap, "domainName", d.DomainName)
	populate(objectMap, "domainSecuritySettings", d.DomainSecuritySettings)
	populate(objectMap, "filteredSync", d.FilteredSync)
	populate(objectMap, "ldapsSettings", d.LdapsSettings)
	populate(objectMap, "migrationProperties", d.MigrationProperties)
	populate(objectMap, "notificationSettings", d.NotificationSettings)
	populate(objectMap, "provisioningState", d.ProvisioningState)
	populate(objectMap, "replicaSets", d.ReplicaSets)
	populate(objectMap, "resourceForestSettings", d.ResourceForestSettings)
	populate(objectMap, "sku", d.SKU)
	populate(objectMap, "syncOwner", d.SyncOwner)
	populate(objectMap, "tenantId", d.TenantID)
	populate(objectMap, "version", d.Version)
	return json.Marshal(objectMap)
}

// DomainServicesBeginCreateOrUpdateOptions contains the optional parameters for the DomainServices.BeginCreateOrUpdate method.
type DomainServicesBeginCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// DomainServicesBeginDeleteOptions contains the optional parameters for the DomainServices.BeginDelete method.
type DomainServicesBeginDeleteOptions struct {
	// placeholder for future optional parameters
}

// DomainServicesBeginUpdateOptions contains the optional parameters for the DomainServices.BeginUpdate method.
type DomainServicesBeginUpdateOptions struct {
	// placeholder for future optional parameters
}

// DomainServicesGetOptions contains the optional parameters for the DomainServices.Get method.
type DomainServicesGetOptions struct {
	// placeholder for future optional parameters
}

// DomainServicesListByResourceGroupOptions contains the optional parameters for the DomainServices.ListByResourceGroup method.
type DomainServicesListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// DomainServicesListOptions contains the optional parameters for the DomainServices.List method.
type DomainServicesListOptions struct {
	// placeholder for future optional parameters
}

// ForestTrust - Forest Trust Setting
type ForestTrust struct {
	// Friendly Name
	FriendlyName *string `json:"friendlyName,omitempty"`

	// Remote Dns ips
	RemoteDNSIPs *string `json:"remoteDnsIps,omitempty"`

	// Trust Direction
	TrustDirection *string `json:"trustDirection,omitempty"`

	// Trust Password
	TrustPassword *string `json:"trustPassword,omitempty"`

	// Trusted Domain FQDN
	TrustedDomainFqdn *string `json:"trustedDomainFqdn,omitempty"`
}

// HealthAlert - Health Alert Description
type HealthAlert struct {
	// READ-ONLY; Health Alert Id
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Health Alert Issue
	Issue *string `json:"issue,omitempty" azure:"ro"`

	// READ-ONLY; Health Alert Last Detected DateTime
	LastDetected *time.Time `json:"lastDetected,omitempty" azure:"ro"`

	// READ-ONLY; Health Alert Name
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Health Alert Raised DateTime
	Raised *time.Time `json:"raised,omitempty" azure:"ro"`

	// READ-ONLY; Health Alert TSG Link
	ResolutionURI *string `json:"resolutionUri,omitempty" azure:"ro"`

	// READ-ONLY; Health Alert Severity
	Severity *string `json:"severity,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type HealthAlert.
func (h HealthAlert) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", h.ID)
	populate(objectMap, "issue", h.Issue)
	populateTimeRFC3339(objectMap, "lastDetected", h.LastDetected)
	populate(objectMap, "name", h.Name)
	populateTimeRFC3339(objectMap, "raised", h.Raised)
	populate(objectMap, "resolutionUri", h.ResolutionURI)
	populate(objectMap, "severity", h.Severity)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type HealthAlert.
func (h *HealthAlert) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, &h.ID)
			delete(rawMsg, key)
		case "issue":
			err = unpopulate(val, &h.Issue)
			delete(rawMsg, key)
		case "lastDetected":
			err = unpopulateTimeRFC3339(val, &h.LastDetected)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, &h.Name)
			delete(rawMsg, key)
		case "raised":
			err = unpopulateTimeRFC3339(val, &h.Raised)
			delete(rawMsg, key)
		case "resolutionUri":
			err = unpopulate(val, &h.ResolutionURI)
			delete(rawMsg, key)
		case "severity":
			err = unpopulate(val, &h.Severity)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// HealthMonitor - Health Monitor Description
type HealthMonitor struct {
	// READ-ONLY; Health Monitor Details
	Details *string `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; Health Monitor Id
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Health Monitor Name
	Name *string `json:"name,omitempty" azure:"ro"`
}

// LdapsSettings - Secure LDAP Settings
type LdapsSettings struct {
	// A flag to determine whether or not Secure LDAP access over the internet is enabled or disabled.
	ExternalAccess *ExternalAccess `json:"externalAccess,omitempty"`

	// A flag to determine whether or not Secure LDAP is enabled or disabled.
	Ldaps *Ldaps `json:"ldaps,omitempty"`

	// The certificate required to configure Secure LDAP. The parameter passed here should be a base64encoded representation of the certificate pfx file.
	PfxCertificate *string `json:"pfxCertificate,omitempty"`

	// The password to decrypt the provided Secure LDAP certificate pfx file.
	PfxCertificatePassword *string `json:"pfxCertificatePassword,omitempty"`

	// READ-ONLY; NotAfter DateTime of configure ldaps certificate.
	CertificateNotAfter *time.Time `json:"certificateNotAfter,omitempty" azure:"ro"`

	// READ-ONLY; Thumbprint of configure ldaps certificate.
	CertificateThumbprint *string `json:"certificateThumbprint,omitempty" azure:"ro"`

	// READ-ONLY; Public certificate used to configure secure ldap.
	PublicCertificate *string `json:"publicCertificate,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type LdapsSettings.
func (l LdapsSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "certificateNotAfter", l.CertificateNotAfter)
	populate(objectMap, "certificateThumbprint", l.CertificateThumbprint)
	populate(objectMap, "externalAccess", l.ExternalAccess)
	populate(objectMap, "ldaps", l.Ldaps)
	populate(objectMap, "pfxCertificate", l.PfxCertificate)
	populate(objectMap, "pfxCertificatePassword", l.PfxCertificatePassword)
	populate(objectMap, "publicCertificate", l.PublicCertificate)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LdapsSettings.
func (l *LdapsSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "certificateNotAfter":
			err = unpopulateTimeRFC3339(val, &l.CertificateNotAfter)
			delete(rawMsg, key)
		case "certificateThumbprint":
			err = unpopulate(val, &l.CertificateThumbprint)
			delete(rawMsg, key)
		case "externalAccess":
			err = unpopulate(val, &l.ExternalAccess)
			delete(rawMsg, key)
		case "ldaps":
			err = unpopulate(val, &l.Ldaps)
			delete(rawMsg, key)
		case "pfxCertificate":
			err = unpopulate(val, &l.PfxCertificate)
			delete(rawMsg, key)
		case "pfxCertificatePassword":
			err = unpopulate(val, &l.PfxCertificatePassword)
			delete(rawMsg, key)
		case "publicCertificate":
			err = unpopulate(val, &l.PublicCertificate)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MigrationProgress - Migration Progress
type MigrationProgress struct {
	// Completion Percentage
	CompletionPercentage *float64 `json:"completionPercentage,omitempty"`

	// Progress Message
	ProgressMessage *string `json:"progressMessage,omitempty"`
}

// MigrationProperties - Migration Properties
type MigrationProperties struct {
	// READ-ONLY; Migration Progress
	MigrationProgress *MigrationProgress `json:"migrationProgress,omitempty" azure:"ro"`

	// READ-ONLY; Old Subnet Id
	OldSubnetID *string `json:"oldSubnetId,omitempty" azure:"ro"`

	// READ-ONLY; Old Vnet Site Id
	OldVnetSiteID *string `json:"oldVnetSiteId,omitempty" azure:"ro"`
}

// NotificationSettings - Settings for notification
type NotificationSettings struct {
	// The list of additional recipients
	AdditionalRecipients []*string `json:"additionalRecipients,omitempty"`

	// Should domain controller admins be notified
	NotifyDcAdmins *NotifyDcAdmins `json:"notifyDcAdmins,omitempty"`

	// Should global admins be notified
	NotifyGlobalAdmins *NotifyGlobalAdmins `json:"notifyGlobalAdmins,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type NotificationSettings.
func (n NotificationSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalRecipients", n.AdditionalRecipients)
	populate(objectMap, "notifyDcAdmins", n.NotifyDcAdmins)
	populate(objectMap, "notifyGlobalAdmins", n.NotifyGlobalAdmins)
	return json.Marshal(objectMap)
}

// OperationDisplayInfo - The operation supported by Domain Services.
type OperationDisplayInfo struct {
	// The description of the operation.
	Description *string `json:"description,omitempty"`

	// The action that users can perform, based on their permission level.
	Operation *string `json:"operation,omitempty"`

	// Service provider: Domain Services.
	Provider *string `json:"provider,omitempty"`

	// Resource on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
}

// OperationEntity - The operation supported by Domain Services.
type OperationEntity struct {
	// The operation supported by Domain Services.
	Display *OperationDisplayInfo `json:"display,omitempty"`

	// Operation name: {provider}/{resource}/{operation}.
	Name *string `json:"name,omitempty"`

	// The origin of the operation.
	Origin *string `json:"origin,omitempty"`
}

// OperationEntityListResult - The list of domain service operation response.
type OperationEntityListResult struct {
	// The list of operations.
	Value []*OperationEntity `json:"value,omitempty"`

	// READ-ONLY; The continuation token for the next page of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type OperationEntityListResult.
func (o OperationEntityListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// OuContainer - Resource for OuContainer.
type OuContainer struct {
	Resource
	// OuContainer properties
	Properties *OuContainerProperties `json:"properties,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type OuContainer.
func (o OuContainer) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	o.Resource.marshalInternal(objectMap)
	populate(objectMap, "properties", o.Properties)
	return json.Marshal(objectMap)
}

// OuContainerBeginCreateOptions contains the optional parameters for the OuContainer.BeginCreate method.
type OuContainerBeginCreateOptions struct {
	// placeholder for future optional parameters
}

// OuContainerBeginDeleteOptions contains the optional parameters for the OuContainer.BeginDelete method.
type OuContainerBeginDeleteOptions struct {
	// placeholder for future optional parameters
}

// OuContainerBeginUpdateOptions contains the optional parameters for the OuContainer.BeginUpdate method.
type OuContainerBeginUpdateOptions struct {
	// placeholder for future optional parameters
}

// OuContainerGetOptions contains the optional parameters for the OuContainer.Get method.
type OuContainerGetOptions struct {
	// placeholder for future optional parameters
}

// OuContainerListOptions contains the optional parameters for the OuContainer.List method.
type OuContainerListOptions struct {
	// placeholder for future optional parameters
}

// OuContainerListResult - The response from the List OuContainer operation.
type OuContainerListResult struct {
	// The list of OuContainer.
	Value []*OuContainer `json:"value,omitempty"`

	// READ-ONLY; The continuation token for the next page of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type OuContainerListResult.
func (o OuContainerListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// OuContainerOperationsListOptions contains the optional parameters for the OuContainerOperations.List method.
type OuContainerOperationsListOptions struct {
	// placeholder for future optional parameters
}

// OuContainerProperties - Properties of the OuContainer.
type OuContainerProperties struct {
	// The list of container accounts
	Accounts []*ContainerAccount `json:"accounts,omitempty"`

	// READ-ONLY; The OuContainer name
	ContainerID *string `json:"containerId,omitempty" azure:"ro"`

	// READ-ONLY; The Deployment id
	DeploymentID *string `json:"deploymentId,omitempty" azure:"ro"`

	// READ-ONLY; Distinguished Name of OuContainer instance
	DistinguishedName *string `json:"distinguishedName,omitempty" azure:"ro"`

	// READ-ONLY; The domain name of Domain Services.
	DomainName *string `json:"domainName,omitempty" azure:"ro"`

	// READ-ONLY; The current deployment or provisioning state, which only appears in the response.
	ProvisioningState *string `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Status of OuContainer instance
	ServiceStatus *string `json:"serviceStatus,omitempty" azure:"ro"`

	// READ-ONLY; Azure Active Directory tenant id
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type OuContainerProperties.
func (o OuContainerProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "accounts", o.Accounts)
	populate(objectMap, "containerId", o.ContainerID)
	populate(objectMap, "deploymentId", o.DeploymentID)
	populate(objectMap, "distinguishedName", o.DistinguishedName)
	populate(objectMap, "domainName", o.DomainName)
	populate(objectMap, "provisioningState", o.ProvisioningState)
	populate(objectMap, "serviceStatus", o.ServiceStatus)
	populate(objectMap, "tenantId", o.TenantID)
	return json.Marshal(objectMap)
}

// ReplicaSet - Replica Set Definition
type ReplicaSet struct {
	// Virtual network location
	Location *string `json:"location,omitempty"`

	// The name of the virtual network that Domain Services will be deployed on. The id of the subnet that Domain Services will be deployed on. /virtualNetwork/vnetName/subnets/subnetName.
	SubnetID *string `json:"subnetId,omitempty"`

	// READ-ONLY; List of Domain Controller IP Address
	DomainControllerIPAddress []*string `json:"domainControllerIpAddress,omitempty" azure:"ro"`

	// READ-ONLY; External access ip address.
	ExternalAccessIPAddress *string `json:"externalAccessIpAddress,omitempty" azure:"ro"`

	// READ-ONLY; List of Domain Health Alerts
	HealthAlerts []*HealthAlert `json:"healthAlerts,omitempty" azure:"ro"`

	// READ-ONLY; Last domain evaluation run DateTime
	HealthLastEvaluated *time.Time `json:"healthLastEvaluated,omitempty" azure:"ro"`

	// READ-ONLY; List of Domain Health Monitors
	HealthMonitors []*HealthMonitor `json:"healthMonitors,omitempty" azure:"ro"`

	// READ-ONLY; ReplicaSet Id
	ReplicaSetID *string `json:"replicaSetId,omitempty" azure:"ro"`

	// READ-ONLY; Status of Domain Service instance
	ServiceStatus *string `json:"serviceStatus,omitempty" azure:"ro"`

	// READ-ONLY; Virtual network site id
	VnetSiteID *string `json:"vnetSiteId,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type ReplicaSet.
func (r ReplicaSet) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "domainControllerIpAddress", r.DomainControllerIPAddress)
	populate(objectMap, "externalAccessIpAddress", r.ExternalAccessIPAddress)
	populate(objectMap, "healthAlerts", r.HealthAlerts)
	populateTimeRFC1123(objectMap, "healthLastEvaluated", r.HealthLastEvaluated)
	populate(objectMap, "healthMonitors", r.HealthMonitors)
	populate(objectMap, "location", r.Location)
	populate(objectMap, "replicaSetId", r.ReplicaSetID)
	populate(objectMap, "serviceStatus", r.ServiceStatus)
	populate(objectMap, "subnetId", r.SubnetID)
	populate(objectMap, "vnetSiteId", r.VnetSiteID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ReplicaSet.
func (r *ReplicaSet) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "domainControllerIpAddress":
			err = unpopulate(val, &r.DomainControllerIPAddress)
			delete(rawMsg, key)
		case "externalAccessIpAddress":
			err = unpopulate(val, &r.ExternalAccessIPAddress)
			delete(rawMsg, key)
		case "healthAlerts":
			err = unpopulate(val, &r.HealthAlerts)
			delete(rawMsg, key)
		case "healthLastEvaluated":
			err = unpopulateTimeRFC1123(val, &r.HealthLastEvaluated)
			delete(rawMsg, key)
		case "healthMonitors":
			err = unpopulate(val, &r.HealthMonitors)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, &r.Location)
			delete(rawMsg, key)
		case "replicaSetId":
			err = unpopulate(val, &r.ReplicaSetID)
			delete(rawMsg, key)
		case "serviceStatus":
			err = unpopulate(val, &r.ServiceStatus)
			delete(rawMsg, key)
		case "subnetId":
			err = unpopulate(val, &r.SubnetID)
			delete(rawMsg, key)
		case "vnetSiteId":
			err = unpopulate(val, &r.VnetSiteID)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// Resource - The Resource model definition.
type Resource struct {
	// Resource etag
	Etag *string `json:"etag,omitempty"`

	// Resource location
	Location *string `json:"location,omitempty"`

	// Resource tags
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Resource Id
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Resource name
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; Resource type
	Type *string `json:"type,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	r.marshalInternal(objectMap)
	return json.Marshal(objectMap)
}

func (r Resource) marshalInternal(objectMap map[string]interface{}) {
	populate(objectMap, "etag", r.Etag)
	populate(objectMap, "id", r.ID)
	populate(objectMap, "location", r.Location)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "systemData", r.SystemData)
	populate(objectMap, "tags", r.Tags)
	populate(objectMap, "type", r.Type)
}

// ResourceForestSettings - Settings for Resource Forest
type ResourceForestSettings struct {
	// Resource Forest
	ResourceForest *string `json:"resourceForest,omitempty"`

	// List of settings for Resource Forest
	Settings []*ForestTrust `json:"settings,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ResourceForestSettings.
func (r ResourceForestSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "resourceForest", r.ResourceForest)
	populate(objectMap, "settings", r.Settings)
	return json.Marshal(objectMap)
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
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
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

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}
