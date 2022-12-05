//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armserialconsole

// CloudError - An error response from the service.
type CloudError struct {
	// Cloud error body.
	Error *CloudErrorBody `json:"error,omitempty"`
}

// CloudErrorBody - An error response from the Batch service.
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

// DisableSerialConsoleResult - Returns whether or not Serial Console is disabled.
type DisableSerialConsoleResult struct {
	// Whether or not Serial Console is disabled.
	Disabled *bool `json:"disabled,omitempty"`
}

// EnableSerialConsoleResult - Returns whether or not Serial Console is disabled (enabled).
type EnableSerialConsoleResult struct {
	// Whether or not Serial Console is disabled (enabled).
	Disabled *bool `json:"disabled,omitempty"`
}

// GetSerialConsoleSubscriptionNotFound - Error saying that the provided subscription could not be found
type GetSerialConsoleSubscriptionNotFound struct {
	// Error code
	Code *string `json:"code,omitempty"`

	// Subscription not found message
	Message *string `json:"message,omitempty"`
}

// MicrosoftSerialConsoleClientDisableConsoleOptions contains the optional parameters for the MicrosoftSerialConsoleClient.DisableConsole
// method.
type MicrosoftSerialConsoleClientDisableConsoleOptions struct {
	// placeholder for future optional parameters
}

// MicrosoftSerialConsoleClientEnableConsoleOptions contains the optional parameters for the MicrosoftSerialConsoleClient.EnableConsole
// method.
type MicrosoftSerialConsoleClientEnableConsoleOptions struct {
	// placeholder for future optional parameters
}

// MicrosoftSerialConsoleClientGetConsoleStatusOptions contains the optional parameters for the MicrosoftSerialConsoleClient.GetConsoleStatus
// method.
type MicrosoftSerialConsoleClientGetConsoleStatusOptions struct {
	// placeholder for future optional parameters
}

// MicrosoftSerialConsoleClientListOperationsOptions contains the optional parameters for the MicrosoftSerialConsoleClient.ListOperations
// method.
type MicrosoftSerialConsoleClientListOperationsOptions struct {
	// placeholder for future optional parameters
}

// Operations - Serial Console operations
type Operations struct {
	// A list of Serial Console operations
	Value []*OperationsValueItem `json:"value,omitempty"`
}

type OperationsValueItem struct {
	Display      *OperationsValueItemDisplay `json:"display,omitempty"`
	IsDataAction *string                     `json:"isDataAction,omitempty"`
	Name         *string                     `json:"name,omitempty"`
}

type OperationsValueItemDisplay struct {
	Description *string `json:"description,omitempty"`
	Operation   *string `json:"operation,omitempty"`
	Provider    *string `json:"provider,omitempty"`
	Resource    *string `json:"resource,omitempty"`
}

// ProxyResource - The resource model definition for a ARM proxy resource. It will have everything other than required location
// and tags
type ProxyResource struct {
	// READ-ONLY; Resource Id
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Resource name
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Resource type
	Type *string `json:"type,omitempty" azure:"ro"`
}

// Resource - The Resource model definition.
type Resource struct {
	// READ-ONLY; Resource Id
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Resource name
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Resource type
	Type *string `json:"type,omitempty" azure:"ro"`
}

// SerialPort - Represents the serial port of the parent resource.
type SerialPort struct {
	// The properties of the serial port.
	Properties *SerialPortProperties `json:"properties,omitempty"`

	// READ-ONLY; Resource Id
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Resource name
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Resource type
	Type *string `json:"type,omitempty" azure:"ro"`
}

// SerialPortConnectResult - Returns a connection string to the serial port of the resource.
type SerialPortConnectResult struct {
	// Connection string to the serial port of the resource.
	ConnectionString *string `json:"connectionString,omitempty"`
}

// SerialPortListResult - The list serial ports operation response.
type SerialPortListResult struct {
	// The list of serial ports.
	Value []*SerialPort `json:"value,omitempty"`
}

// SerialPortProperties - The properties of the serial port.
type SerialPortProperties struct {
	// Specifies whether the port is enabled for a serial console connection.
	State *SerialPortState `json:"state,omitempty"`
}

// SerialPortsClientConnectOptions contains the optional parameters for the SerialPortsClient.Connect method.
type SerialPortsClientConnectOptions struct {
	// placeholder for future optional parameters
}

// SerialPortsClientCreateOptions contains the optional parameters for the SerialPortsClient.Create method.
type SerialPortsClientCreateOptions struct {
	// placeholder for future optional parameters
}

// SerialPortsClientDeleteOptions contains the optional parameters for the SerialPortsClient.Delete method.
type SerialPortsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// SerialPortsClientGetOptions contains the optional parameters for the SerialPortsClient.Get method.
type SerialPortsClientGetOptions struct {
	// placeholder for future optional parameters
}

// SerialPortsClientListBySubscriptionsOptions contains the optional parameters for the SerialPortsClient.ListBySubscriptions
// method.
type SerialPortsClientListBySubscriptionsOptions struct {
	// placeholder for future optional parameters
}

// SerialPortsClientListOptions contains the optional parameters for the SerialPortsClient.List method.
type SerialPortsClientListOptions struct {
	// placeholder for future optional parameters
}

// Status - Returns whether or not Serial Console is disabled.
type Status struct {
	// Whether or not Serial Console is disabled.
	Disabled *bool `json:"disabled,omitempty"`
}