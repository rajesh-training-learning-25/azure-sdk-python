// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armdeviceregistry

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"time"
)

// Asset definition.
type Asset struct {
	// REQUIRED; The extended location.
	ExtendedLocation *ExtendedLocation

	// REQUIRED; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string

	// The resource-specific properties for this resource.
	Properties *AssetProperties

	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Asset name parameter.
	Name *string
}

// AssetEndpointProfile - Asset Endpoint Profile definition.
type AssetEndpointProfile struct {
	// REQUIRED; The extended location.
	ExtendedLocation *ExtendedLocation

	// REQUIRED; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string

	// The resource-specific properties for this resource.
	Properties *AssetEndpointProfileProperties

	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Asset Endpoint Profile name parameter.
	Name *string
}

// AssetEndpointProfileListResult - The response of a AssetEndpointProfile list operation.
type AssetEndpointProfileListResult struct {
	// REQUIRED; The AssetEndpointProfile items on this page
	Value []*AssetEndpointProfile

	// The link to the next page of items
	NextLink *string
}

// AssetEndpointProfileProperties - Defines the Asset Endpoint Profile properties.
type AssetEndpointProfileProperties struct {
	// REQUIRED; The local valid URI specifying the network address/DNS name of a southbound device. The scheme part of the targetAddress
	// URI specifies the type of the device. The additionalConfiguration field holds further connector type specific configuration.
	TargetAddress *string

	// Contains connectivity type specific further configuration (e.g. OPC UA, Modbus, ONVIF).
	AdditionalConfiguration *string

	// Provisioning state of the resource.
	ProvisioningState *ProvisioningState

	// Defines the authentication mechanism for the southbound connector connecting to the shop floor/OT device.
	TransportAuthentication *TransportAuthentication

	// Globally unique, immutable, non-reusable id.
	UUID *string

	// Defines the client authentication mechanism to the server.
	UserAuthentication *UserAuthentication
}

// AssetEndpointProfileUpdate - The type used for update operations of the AssetEndpointProfile.
type AssetEndpointProfileUpdate struct {
	Properties *AssetEndpointProfileUpdateProperties

	// Resource tags.
	Tags map[string]*string
}

// AssetEndpointProfileUpdateProperties - The updatable properties of the AssetEndpointProfile.
type AssetEndpointProfileUpdateProperties struct {
	// Contains connectivity type specific further configuration (e.g. OPC UA, Modbus, ONVIF).
	AdditionalConfiguration *string

	// The local valid URI specifying the network address/DNS name of a southbound device. The scheme part of the targetAddress
	// URI specifies the type of the device. The additionalConfiguration field holds further connector type specific configuration.
	TargetAddress *string

	// Defines the authentication mechanism for the southbound connector connecting to the shop floor/OT device.
	TransportAuthentication *TransportAuthentication

	// Defines the client authentication mechanism to the server.
	UserAuthentication *UserAuthentication
}

// AssetListResult - The response of a Asset list operation.
type AssetListResult struct {
	// REQUIRED; The Asset items on this page
	Value []*Asset

	// The link to the next page of items
	NextLink *string
}

// AssetProperties - Defines the asset properties.
type AssetProperties struct {
	// REQUIRED; A reference to the asset endpoint profile (connection information) used by brokers to connect to an endpoint
	// that provides data points for this asset. Must have the format <ModuleCR.metadata.namespace>/<ModuleCR.metadata.name>.
	AssetEndpointProfileURI *string

	// Resource path to asset type (model) definition.
	AssetType *string

	// A set of key-value pairs that contain custom attributes set by the customer.
	Attributes map[string]any

	// Array of data points that are part of the asset. Each data point can reference an asset type capability and have per-data
	// point configuration. See below for more details for the definition of the dataPoints element.
	DataPoints []*DataPoint

	// Protocol-specific default configuration for all data points. Each data point can have its own configuration that overrides
	// the default settings here. This assumes that each asset instance has one protocol.
	DefaultDataPointsConfiguration *string

	// Protocol-specific default configuration for all events. Each event can have its own configuration that overrides the default
	// settings here. This assumes that each asset instance has one protocol.
	DefaultEventsConfiguration *string

	// Human-readable description of the asset.
	Description *string

	// Human-readable display name.
	DisplayName *string

	// Reference to the documentation.
	DocumentationURI *string

	// Enabled/Disabled status of the asset.
	Enabled *bool

	// Array of events that are part of the asset. Each event can reference an asset type capability and have per-event configuration.
	// See below for more details about the definition of the events element.
	Events []*Event

	// Asset id provided by the customer.
	ExternalAssetID *string

	// Revision number of the hardware.
	HardwareRevision *string

	// Asset manufacturer name.
	Manufacturer *string

	// Asset manufacturer URI.
	ManufacturerURI *string

	// Asset model name.
	Model *string

	// Asset product code.
	ProductCode *string

	// Provisioning state of the resource.
	ProvisioningState *ProvisioningState

	// Asset serial number.
	SerialNumber *string

	// Revision number of the software.
	SoftwareRevision *string

	// Read only object to reflect changes that have occurred on the Edge. Similar to Kubernetes status property for custom resources.
	Status *AssetStatus

	// Globally unique, immutable, non-reusable id.
	UUID *string

	// An integer that is incremented each time the resource is modified.
	Version *int32
}

// AssetStatus - Defines the asset status properties.
type AssetStatus struct {
	// Array object to transfer and persist errors that originate from the Edge.
	Errors []*AssetStatusError

	// A read only incremental counter indicating the number of times the configuration has been modified from the perspective
	// of the current actual (Edge) state of the Asset. Edge would be the only writer of this value and would sync back up to
	// the cloud. In steady state, this should equal version.
	Version *int32
}

// AssetStatusError - Defines the asset status error properties.
type AssetStatusError struct {
	// Error code for classification of errors (ex: 400, 404, 500, etc.).
	Code *int32

	// Human readable helpful error message to provide additional context for error (ex: “capability Id 'foo' does not exist”).
	Message *string
}

// AssetUpdate - The type used for update operations of the Asset.
type AssetUpdate struct {
	Properties *AssetUpdateProperties

	// Resource tags.
	Tags map[string]*string
}

// AssetUpdateProperties - The updatable properties of the Asset.
type AssetUpdateProperties struct {
	// Resource path to asset type (model) definition.
	AssetType *string

	// A set of key-value pairs that contain custom attributes set by the customer.
	Attributes map[string]any

	// Array of data points that are part of the asset. Each data point can reference an asset type capability and have per-data
	// point configuration. See below for more details for the definition of the dataPoints element.
	DataPoints []*DataPoint

	// Protocol-specific default configuration for all data points. Each data point can have its own configuration that overrides
	// the default settings here. This assumes that each asset instance has one protocol.
	DefaultDataPointsConfiguration *string

	// Protocol-specific default configuration for all events. Each event can have its own configuration that overrides the default
	// settings here. This assumes that each asset instance has one protocol.
	DefaultEventsConfiguration *string

	// Human-readable description of the asset.
	Description *string

	// Human-readable display name.
	DisplayName *string

	// Reference to the documentation.
	DocumentationURI *string

	// Enabled/Disabled status of the asset.
	Enabled *bool

	// Array of events that are part of the asset. Each event can reference an asset type capability and have per-event configuration.
	// See below for more details about the definition of the events element.
	Events []*Event

	// Revision number of the hardware.
	HardwareRevision *string

	// Asset manufacturer name.
	Manufacturer *string

	// Asset manufacturer URI.
	ManufacturerURI *string

	// Asset model name.
	Model *string

	// Asset product code.
	ProductCode *string

	// Asset serial number.
	SerialNumber *string

	// Revision number of the software.
	SoftwareRevision *string
}

// DataPoint - Defines the data point properties.
type DataPoint struct {
	// REQUIRED; The address of the source of the data in the asset (e.g. URL) so that a client can access the data source on
	// the asset.
	DataSource *string

	// The path to the type definition of the capability (e.g. DTMI, OPC UA information model node id, etc.), for example dtmi:com:example:Robot:_contents:__prop1;1.
	CapabilityID *string

	// Protocol-specific configuration for the data point. For OPC UA, this could include configuration like, publishingInterval,
	// samplingInterval, and queueSize.
	DataPointConfiguration *string

	// The name of the data point.
	Name *string

	// An indication of how the data point should be mapped to OpenTelemetry.
	ObservabilityMode *DataPointsObservabilityMode
}

// Event - Defines the event properties.
type Event struct {
	// REQUIRED; The address of the notifier of the event in the asset (e.g. URL) so that a client can access the event on the
	// asset.
	EventNotifier *string

	// The path to the type definition of the capability (e.g. DTMI, OPC UA information model node id, etc.), for example dtmi:com:example:Robot:_contents:__prop1;1.
	CapabilityID *string

	// Protocol-specific configuration for the event. For OPC UA, this could include configuration like, publishingInterval, samplingInterval,
	// and queueSize.
	EventConfiguration *string

	// The name of the event.
	Name *string

	// An indication of how the event should be mapped to OpenTelemetry.
	ObservabilityMode *EventsObservabilityMode
}

// ExtendedLocation - The extended location.
type ExtendedLocation struct {
	// REQUIRED; The extended location name.
	Name *string

	// REQUIRED; The extended location type.
	Type *string
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// Localized display information for this particular operation.
	Display *OperationDisplay

	// Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for Azure Resource Manager/control-plane
	// operations.
	IsDataAction *bool

	// The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

	// The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default value is
	// "user,system"
	Origin *Origin
}

// OperationDisplay - Localized display information for and operation.
type OperationDisplay struct {
	// The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string

	// The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual Machine",
	// "Restart Virtual Machine".
	Operation *string

	// The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft Compute".
	Provider *string

	// The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job Schedule Collections".
	Resource *string
}

// OperationStatusResult - The current status of an async operation.
type OperationStatusResult struct {
	// REQUIRED; The operations list.
	Operations []*OperationStatusResult

	// REQUIRED; Operation status.
	Status *string

	// The end time of the operation.
	EndTime *time.Time

	// If present, details of the operation error.
	Error *azcore.ResponseError

	// Fully qualified ID for the async operation.
	ID *string

	// Name of the async operation.
	Name *string

	// Percent of the operation that is complete.
	PercentComplete *int32

	// The start time of the operation.
	StartTime *time.Time
}

// OwnCertificate - Certificate or private key that can be used by the southbound connector connecting to the shop floor/OT
// device. The accepted extensions are .der for certificates and .pfx/.pem for private keys.
type OwnCertificate struct {
	// Secret Reference Name (Pfx or Pem password).
	CertPasswordReference *string

	// Secret Reference name (cert and private key).
	CertSecretReference *string

	// Certificate thumbprint.
	CertThumbprint *string
}

// PagedOperation - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get
// the next set of results.
type PagedOperation struct {
	// REQUIRED; The Operation items on this page
	Value []*Operation

	// The link to the next page of items
	NextLink *string
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The type of identity that created the resource.
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}

// TransportAuthentication - Definition of the authentication mechanism for the southbound connector.
type TransportAuthentication struct {
	// REQUIRED; Defines a reference to a secret which contains all certificates and private keys that can be used by the southbound
	// connector connecting to the shop floor/OT device. The accepted extensions are .der for certificates and .pfx/.pem for private
	// keys.
	OwnCertificates []*OwnCertificate
}

// UserAuthentication - Definition of the client authentication mechanism to the server.
type UserAuthentication struct {
	// REQUIRED; Defines the mode to authenticate the user of the client at the server.
	Mode *UserAuthenticationMode

	// Defines the username and password references when UsernamePassword user authentication mode is selected.
	UsernamePasswordCredentials *UsernamePasswordCredentials

	// Defines the certificate reference when Certificate user authentication mode is selected.
	X509Credentials *X509Credentials
}

// UsernamePasswordCredentials - The credentials for authentication mode UsernamePassword.
type UsernamePasswordCredentials struct {
	// REQUIRED; A reference to secret containing the password.
	PasswordReference *string

	// REQUIRED; A reference to secret containing the username.
	UsernameReference *string
}

// X509Credentials - The x509 certificate for authentication mode Certificate.
type X509Credentials struct {
	// REQUIRED; A reference to secret containing the certificate and private key (e.g. stored as .der/.pem or .der/.pfx).
	CertificateReference *string
}
