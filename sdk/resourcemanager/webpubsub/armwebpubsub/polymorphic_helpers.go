//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armwebpubsub

import "encoding/json"

func unmarshalEventListenerEndpointClassification(rawMsg json.RawMessage) (EventListenerEndpointClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b EventListenerEndpointClassification
	switch m["type"] {
	case string(EventListenerEndpointDiscriminatorEventHub):
		b = &EventHubEndpoint{}
	default:
		b = &EventListenerEndpoint{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalEventListenerFilterClassification(rawMsg json.RawMessage) (EventListenerFilterClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b EventListenerFilterClassification
	switch m["type"] {
	case string(EventListenerFilterDiscriminatorEventName):
		b = &EventNameFilter{}
	default:
		b = &EventListenerFilter{}
	}
	return b, json.Unmarshal(rawMsg, b)
}
