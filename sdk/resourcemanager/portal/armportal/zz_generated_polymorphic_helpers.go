//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armportal

import "encoding/json"

func unmarshalDashboardPartMetadataClassification(rawMsg json.RawMessage) (DashboardPartMetadataClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DashboardPartMetadataClassification
	switch m["type"] {
	case "Extension/HubsExtension/PartType/MarkdownPart":
		b = &MarkdownPartMetadata{}
	default:
		b = &DashboardPartMetadata{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalDashboardPartMetadataClassificationArray(rawMsg json.RawMessage) ([]DashboardPartMetadataClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]DashboardPartMetadataClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalDashboardPartMetadataClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalDashboardPartMetadataClassificationMap(rawMsg json.RawMessage) (map[string]DashboardPartMetadataClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]DashboardPartMetadataClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalDashboardPartMetadataClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}
