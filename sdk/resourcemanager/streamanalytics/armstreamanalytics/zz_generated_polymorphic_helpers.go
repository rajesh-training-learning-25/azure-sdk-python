//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstreamanalytics

import "encoding/json"

func unmarshalFunctionBindingClassification(rawMsg json.RawMessage) (FunctionBindingClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b FunctionBindingClassification
	switch m["type"] {
	case "Microsoft.MachineLearning/WebService":
		b = &AzureMachineLearningWebServiceFunctionBinding{}
	case "Microsoft.StreamAnalytics/JavascriptUdf":
		b = &JavaScriptFunctionBinding{}
	default:
		b = &FunctionBinding{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalFunctionPropertiesClassification(rawMsg json.RawMessage) (FunctionPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b FunctionPropertiesClassification
	switch m["type"] {
	case "Aggregate":
		b = &AggregateFunctionProperties{}
	case "Scalar":
		b = &ScalarFunctionProperties{}
	default:
		b = &FunctionProperties{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalInputPropertiesClassification(rawMsg json.RawMessage) (InputPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b InputPropertiesClassification
	switch m["type"] {
	case "Reference":
		b = &ReferenceInputProperties{}
	case "Stream":
		b = &StreamInputProperties{}
	default:
		b = &InputProperties{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalOutputDataSourceClassification(rawMsg json.RawMessage) (OutputDataSourceClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b OutputDataSourceClassification
	switch m["type"] {
	case "Microsoft.AzureFunction":
		b = &AzureFunctionOutputDataSource{}
	case "Microsoft.DataLake/Accounts":
		b = &AzureDataLakeStoreOutputDataSource{}
	case "Microsoft.EventHub/EventHub":
		b = &EventHubV2OutputDataSource{}
	case "Microsoft.ServiceBus/EventHub":
		b = &EventHubOutputDataSource{}
	case "Microsoft.ServiceBus/Queue":
		b = &ServiceBusQueueOutputDataSource{}
	case "Microsoft.ServiceBus/Topic":
		b = &ServiceBusTopicOutputDataSource{}
	case "Microsoft.Sql/Server/DataWarehouse":
		b = &AzureSynapseOutputDataSource{}
	case "Microsoft.Sql/Server/Database":
		b = &AzureSQLDatabaseOutputDataSource{}
	case "Microsoft.Storage/Blob":
		b = &BlobOutputDataSource{}
	case "Microsoft.Storage/DocumentDB":
		b = &DocumentDbOutputDataSource{}
	case "Microsoft.Storage/Table":
		b = &AzureTableOutputDataSource{}
	case "PowerBI":
		b = &PowerBIOutputDataSource{}
	default:
		b = &OutputDataSource{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalReferenceInputDataSourceClassification(rawMsg json.RawMessage) (ReferenceInputDataSourceClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ReferenceInputDataSourceClassification
	switch m["type"] {
	case "Microsoft.Sql/Server/Database":
		b = &AzureSQLReferenceInputDataSource{}
	case "Microsoft.Storage/Blob":
		b = &BlobReferenceInputDataSource{}
	default:
		b = &ReferenceInputDataSource{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalSerializationClassification(rawMsg json.RawMessage) (SerializationClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b SerializationClassification
	switch m["type"] {
	case string(EventSerializationTypeAvro):
		b = &AvroSerialization{}
	case string(EventSerializationTypeCSV):
		b = &CSVSerialization{}
	case string(EventSerializationTypeJSON):
		b = &JSONSerialization{}
	case string(EventSerializationTypeParquet):
		b = &ParquetSerialization{}
	default:
		b = &Serialization{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalStreamInputDataSourceClassification(rawMsg json.RawMessage) (StreamInputDataSourceClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b StreamInputDataSourceClassification
	switch m["type"] {
	case "Microsoft.Devices/IotHubs":
		b = &IoTHubStreamInputDataSource{}
	case "Microsoft.EventHub/EventHub":
		b = &EventHubV2StreamInputDataSource{}
	case "Microsoft.ServiceBus/EventHub":
		b = &EventHubStreamInputDataSource{}
	case "Microsoft.Storage/Blob":
		b = &BlobStreamInputDataSource{}
	default:
		b = &StreamInputDataSource{}
	}
	return b, json.Unmarshal(rawMsg, b)
}