//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armserialconsole

const (
	moduleName    = "armserialconsole"
	moduleVersion = "v1.0.0"
)

// SerialPortState - Specifies whether the port is enabled for a serial console connection.
type SerialPortState string

const (
	SerialPortStateEnabled  SerialPortState = "enabled"
	SerialPortStateDisabled SerialPortState = "disabled"
)

// PossibleSerialPortStateValues returns the possible values for the SerialPortState const type.
func PossibleSerialPortStateValues() []SerialPortState {
	return []SerialPortState{
		SerialPortStateEnabled,
		SerialPortStateDisabled,
	}
}