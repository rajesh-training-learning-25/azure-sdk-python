//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics"
)

func unmarshalFunctionRetrieveDefaultDefinitionParametersClassification(rawMsg json.RawMessage) (armstreamanalytics.FunctionRetrieveDefaultDefinitionParametersClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b armstreamanalytics.FunctionRetrieveDefaultDefinitionParametersClassification
	switch m["bindingType"] {
	case "Microsoft.MachineLearning/WebService":
		b = &armstreamanalytics.AzureMachineLearningWebServiceFunctionRetrieveDefaultDefinitionParameters{}
	case "Microsoft.StreamAnalytics/JavascriptUdf":
		b = &armstreamanalytics.JavaScriptFunctionRetrieveDefaultDefinitionParameters{}
	default:
		b = &armstreamanalytics.FunctionRetrieveDefaultDefinitionParameters{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}
