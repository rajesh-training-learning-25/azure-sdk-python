//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

import "encoding/json"

func unmarshalBackupPolicyClassification(rawMsg json.RawMessage) (BackupPolicyClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b BackupPolicyClassification
	switch m["type"] {
	case string(BackupPolicyTypeContinuous):
		b = &ContinuousModeBackupPolicy{}
	case string(BackupPolicyTypePeriodic):
		b = &PeriodicModeBackupPolicy{}
	default:
		b = &BackupPolicy{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalBackupPolicyClassificationArray(rawMsg json.RawMessage) ([]BackupPolicyClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]BackupPolicyClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalBackupPolicyClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalServiceResourcePropertiesClassification(rawMsg json.RawMessage) (ServiceResourcePropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ServiceResourcePropertiesClassification
	switch m["serviceType"] {
	case string(ServiceTypeDataTransfer):
		b = &DataTransferServiceResourceProperties{}
	case string(ServiceTypeGraphAPICompute):
		b = &GraphAPIComputeServiceResourceProperties{}
	case string(ServiceTypeSQLDedicatedGateway):
		b = &SQLDedicatedGatewayServiceResourceProperties{}
	default:
		b = &ServiceResourceProperties{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalServiceResourcePropertiesClassificationArray(rawMsg json.RawMessage) ([]ServiceResourcePropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ServiceResourcePropertiesClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalServiceResourcePropertiesClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}
