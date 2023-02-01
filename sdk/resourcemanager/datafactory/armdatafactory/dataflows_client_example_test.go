//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4afa6837cfb404d8e5ffa8a604a5e09996d6f79e/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlows_Create.json
func ExampleDataFlowsClient_CreateOrUpdate_dataFlowsCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewDataFlowsClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleDataFlow", armdatafactory.DataFlowResource{
		Properties: &armdatafactory.MappingDataFlow{
			Type:        to.Ptr("MappingDataFlow"),
			Description: to.Ptr("Sample demo data flow to convert currencies showing usage of union, derive and conditional split transformation."),
			TypeProperties: &armdatafactory.MappingDataFlowTypeProperties{
				ScriptLines: []*string{
					to.Ptr("source(output("),
					to.Ptr("PreviousConversionRate as double,"),
					to.Ptr("Country as string,"),
					to.Ptr("DateTime1 as string,"),
					to.Ptr("CurrentConversionRate as double"),
					to.Ptr("),"),
					to.Ptr("allowSchemaDrift: false,"),
					to.Ptr("validateSchema: false) ~> USDCurrency"),
					to.Ptr("source(output("),
					to.Ptr("PreviousConversionRate as double,"),
					to.Ptr("Country as string,"),
					to.Ptr("DateTime1 as string,"),
					to.Ptr("CurrentConversionRate as double"),
					to.Ptr("),"),
					to.Ptr("allowSchemaDrift: true,"),
					to.Ptr("validateSchema: false) ~> CADSource"),
					to.Ptr("USDCurrency, CADSource union(byName: true)~> Union"),
					to.Ptr("Union derive(NewCurrencyRate = round(CurrentConversionRate*1.25)) ~> NewCurrencyColumn"),
					to.Ptr("NewCurrencyColumn split(Country == 'USD',"),
					to.Ptr("Country == 'CAD',disjoint: false) ~> ConditionalSplit1@(USD, CAD)"),
					to.Ptr("ConditionalSplit1@USD sink(saveMode:'overwrite' ) ~> USDSink"),
					to.Ptr("ConditionalSplit1@CAD sink(saveMode:'overwrite' ) ~> CADSink")},
				Sinks: []*armdatafactory.DataFlowSink{
					{
						Name: to.Ptr("USDSink"),
						Dataset: &armdatafactory.DatasetReference{
							Type:          to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
							ReferenceName: to.Ptr("USDOutput"),
						},
					},
					{
						Name: to.Ptr("CADSink"),
						Dataset: &armdatafactory.DatasetReference{
							Type:          to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
							ReferenceName: to.Ptr("CADOutput"),
						},
					}},
				Sources: []*armdatafactory.DataFlowSource{
					{
						Name: to.Ptr("USDCurrency"),
						Dataset: &armdatafactory.DatasetReference{
							Type:          to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
							ReferenceName: to.Ptr("CurrencyDatasetUSD"),
						},
					},
					{
						Name: to.Ptr("CADSource"),
						Dataset: &armdatafactory.DatasetReference{
							Type:          to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
							ReferenceName: to.Ptr("CurrencyDatasetCAD"),
						},
					}},
			},
		},
	}, &armdatafactory.DataFlowsClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DataFlowResource = armdatafactory.DataFlowResource{
	// 	Name: to.Ptr("exampleDataFlow"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/dataflows"),
	// 	Etag: to.Ptr("0a0066d4-0000-0000-0000-5b245bd20000"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/datasets/exampleDataset"),
	// 	Properties: &armdatafactory.MappingDataFlow{
	// 		Type: to.Ptr("MappingDataFlow"),
	// 		Description: to.Ptr("Sample demo data flow to convert currencies showing usage of union, derive and conditional split transformation."),
	// 		TypeProperties: &armdatafactory.MappingDataFlowTypeProperties{
	// 			ScriptLines: []*string{
	// 				to.Ptr("source(output("),
	// 				to.Ptr("PreviousConversionRate as double,"),
	// 				to.Ptr("Country as string,"),
	// 				to.Ptr("DateTime1 as string,"),
	// 				to.Ptr("CurrentConversionRate as double"),
	// 				to.Ptr("),"),
	// 				to.Ptr("allowSchemaDrift: false,"),
	// 				to.Ptr("validateSchema: false) ~> USDCurrency"),
	// 				to.Ptr("source(output("),
	// 				to.Ptr("PreviousConversionRate as double,"),
	// 				to.Ptr("Country as string,"),
	// 				to.Ptr("DateTime1 as string,"),
	// 				to.Ptr("CurrentConversionRate as double"),
	// 				to.Ptr("),"),
	// 				to.Ptr("allowSchemaDrift: true,"),
	// 				to.Ptr("validateSchema: false) ~> CADSource"),
	// 				to.Ptr("USDCurrency, CADSource union(byName: true)~> Union"),
	// 				to.Ptr("Union derive(NewCurrencyRate = round(CurrentConversionRate*1.25)) ~> NewCurrencyColumn"),
	// 				to.Ptr("NewCurrencyColumn split(Country == 'USD',"),
	// 				to.Ptr("Country == 'CAD',disjoint: false) ~> ConditionalSplit1@(USD, CAD)"),
	// 				to.Ptr("ConditionalSplit1@USD sink(saveMode:'overwrite' ) ~> USDSink"),
	// 				to.Ptr("ConditionalSplit1@CAD sink(saveMode:'overwrite' ) ~> CADSink")},
	// 				Sinks: []*armdatafactory.DataFlowSink{
	// 					{
	// 						Name: to.Ptr("USDSink"),
	// 						Dataset: &armdatafactory.DatasetReference{
	// 							Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
	// 							ReferenceName: to.Ptr("USDOutput"),
	// 						},
	// 					},
	// 					{
	// 						Name: to.Ptr("CADSink"),
	// 						Dataset: &armdatafactory.DatasetReference{
	// 							Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
	// 							ReferenceName: to.Ptr("CADOutput"),
	// 						},
	// 				}},
	// 				Sources: []*armdatafactory.DataFlowSource{
	// 					{
	// 						Name: to.Ptr("USDCurrency"),
	// 						Dataset: &armdatafactory.DatasetReference{
	// 							Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
	// 							ReferenceName: to.Ptr("CurrencyDatasetUSD"),
	// 						},
	// 					},
	// 					{
	// 						Name: to.Ptr("CADSource"),
	// 						Dataset: &armdatafactory.DatasetReference{
	// 							Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
	// 							ReferenceName: to.Ptr("CurrencyDatasetCAD"),
	// 						},
	// 				}},
	// 			},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4afa6837cfb404d8e5ffa8a604a5e09996d6f79e/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlows_Update.json
func ExampleDataFlowsClient_CreateOrUpdate_dataFlowsUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewDataFlowsClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleDataFlow", armdatafactory.DataFlowResource{
		Properties: &armdatafactory.MappingDataFlow{
			Type:        to.Ptr("MappingDataFlow"),
			Description: to.Ptr("Sample demo data flow to convert currencies showing usage of union, derive and conditional split transformation."),
			TypeProperties: &armdatafactory.MappingDataFlowTypeProperties{
				ScriptLines: []*string{
					to.Ptr("source(output("),
					to.Ptr("PreviousConversionRate as double,"),
					to.Ptr("Country as string,"),
					to.Ptr("DateTime1 as string,"),
					to.Ptr("CurrentConversionRate as double"),
					to.Ptr("),"),
					to.Ptr("allowSchemaDrift: false,"),
					to.Ptr("validateSchema: false) ~> USDCurrency"),
					to.Ptr("source(output("),
					to.Ptr("PreviousConversionRate as double,"),
					to.Ptr("Country as string,"),
					to.Ptr("DateTime1 as string,"),
					to.Ptr("CurrentConversionRate as double"),
					to.Ptr("),"),
					to.Ptr("allowSchemaDrift: true,"),
					to.Ptr("validateSchema: false) ~> CADSource"),
					to.Ptr("USDCurrency, CADSource union(byName: true)~> Union"),
					to.Ptr("Union derive(NewCurrencyRate = round(CurrentConversionRate*1.25)) ~> NewCurrencyColumn"),
					to.Ptr("NewCurrencyColumn split(Country == 'USD',"),
					to.Ptr("Country == 'CAD',disjoint: false) ~> ConditionalSplit1@(USD, CAD)"),
					to.Ptr("ConditionalSplit1@USD sink(saveMode:'overwrite' ) ~> USDSink"),
					to.Ptr("ConditionalSplit1@CAD sink(saveMode:'overwrite' ) ~> CADSink")},
				Sinks: []*armdatafactory.DataFlowSink{
					{
						Name: to.Ptr("USDSink"),
						Dataset: &armdatafactory.DatasetReference{
							Type:          to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
							ReferenceName: to.Ptr("USDOutput"),
						},
					},
					{
						Name: to.Ptr("CADSink"),
						Dataset: &armdatafactory.DatasetReference{
							Type:          to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
							ReferenceName: to.Ptr("CADOutput"),
						},
					}},
				Sources: []*armdatafactory.DataFlowSource{
					{
						Name: to.Ptr("USDCurrency"),
						Dataset: &armdatafactory.DatasetReference{
							Type:          to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
							ReferenceName: to.Ptr("CurrencyDatasetUSD"),
						},
					},
					{
						Name: to.Ptr("CADSource"),
						Dataset: &armdatafactory.DatasetReference{
							Type:          to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
							ReferenceName: to.Ptr("CurrencyDatasetCAD"),
						},
					}},
			},
		},
	}, &armdatafactory.DataFlowsClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DataFlowResource = armdatafactory.DataFlowResource{
	// 	Name: to.Ptr("exampleDataFlow"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/dataflows"),
	// 	Etag: to.Ptr("0a0068d4-0000-0000-0000-5b245bd30002"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/datasets/exampleDataset"),
	// 	Properties: &armdatafactory.MappingDataFlow{
	// 		Type: to.Ptr("MappingDataFlow"),
	// 		Description: to.Ptr("Sample demo data flow to convert currencies showing usage of union, derive and conditional split transformation."),
	// 		TypeProperties: &armdatafactory.MappingDataFlowTypeProperties{
	// 			ScriptLines: []*string{
	// 				to.Ptr("source(output("),
	// 				to.Ptr("PreviousConversionRate as double,"),
	// 				to.Ptr("Country as string,"),
	// 				to.Ptr("DateTime1 as string,"),
	// 				to.Ptr("CurrentConversionRate as double"),
	// 				to.Ptr("),"),
	// 				to.Ptr("allowSchemaDrift: false,"),
	// 				to.Ptr("validateSchema: false) ~> USDCurrency"),
	// 				to.Ptr("source(output("),
	// 				to.Ptr("PreviousConversionRate as double,"),
	// 				to.Ptr("Country as string,"),
	// 				to.Ptr("DateTime1 as string,"),
	// 				to.Ptr("CurrentConversionRate as double"),
	// 				to.Ptr("),"),
	// 				to.Ptr("allowSchemaDrift: true,"),
	// 				to.Ptr("validateSchema: false) ~> CADSource"),
	// 				to.Ptr("USDCurrency, CADSource union(byName: true)~> Union"),
	// 				to.Ptr("Union derive(NewCurrencyRate = round(CurrentConversionRate*1.25)) ~> NewCurrencyColumn"),
	// 				to.Ptr("NewCurrencyColumn split(Country == 'USD',"),
	// 				to.Ptr("Country == 'CAD',disjoint: false) ~> ConditionalSplit1@(USD, CAD)"),
	// 				to.Ptr("ConditionalSplit1@USD sink(saveMode:'overwrite' ) ~> USDSink"),
	// 				to.Ptr("ConditionalSplit1@CAD sink(saveMode:'overwrite' ) ~> CADSink")},
	// 				Sinks: []*armdatafactory.DataFlowSink{
	// 					{
	// 						Name: to.Ptr("USDSink"),
	// 						Dataset: &armdatafactory.DatasetReference{
	// 							Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
	// 							ReferenceName: to.Ptr("USDOutput"),
	// 						},
	// 					},
	// 					{
	// 						Name: to.Ptr("CADSink"),
	// 						Dataset: &armdatafactory.DatasetReference{
	// 							Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
	// 							ReferenceName: to.Ptr("CADOutput"),
	// 						},
	// 				}},
	// 				Sources: []*armdatafactory.DataFlowSource{
	// 					{
	// 						Name: to.Ptr("USDCurrency"),
	// 						Dataset: &armdatafactory.DatasetReference{
	// 							Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
	// 							ReferenceName: to.Ptr("CurrencyDatasetUSD"),
	// 						},
	// 					},
	// 					{
	// 						Name: to.Ptr("CADSource"),
	// 						Dataset: &armdatafactory.DatasetReference{
	// 							Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
	// 							ReferenceName: to.Ptr("CurrencyDatasetCAD"),
	// 						},
	// 				}},
	// 			},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4afa6837cfb404d8e5ffa8a604a5e09996d6f79e/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlows_Get.json
func ExampleDataFlowsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewDataFlowsClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleDataFlow", &armdatafactory.DataFlowsClientGetOptions{IfNoneMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DataFlowResource = armdatafactory.DataFlowResource{
	// 	Name: to.Ptr("exampleDataFlow"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/dataflows"),
	// 	Etag: to.Ptr("15004c4f-0000-0200-0000-5cbe090e0000"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/dataflows/exampleDataFlow"),
	// 	Properties: &armdatafactory.MappingDataFlow{
	// 		Type: to.Ptr("MappingDataFlow"),
	// 		Description: to.Ptr("Sample demo data flow to convert currencies showing usage of union, derive and conditional split transformation."),
	// 		TypeProperties: &armdatafactory.MappingDataFlowTypeProperties{
	// 			Script: to.Ptr("source(output(PreviousConversionRate as double,Country as string,DateTime1 as string,CurrentConversionRate as double),allowSchemaDrift: false,validateSchema: false) ~> USDCurrency\nsource(output(PreviousConversionRate as double,Country as string,DateTime1 as string,CurrentConversionRate as double),allowSchemaDrift: true,validateSchema: false) ~> CADSource\nUSDCurrency, CADSource union(byName: true)~> Union\nUnion derive(NewCurrencyRate = round(CurrentConversionRate*1.25)) ~> NewCurrencyColumn\nNewCurrencyColumn split(Country == 'USD',Country == 'CAD',disjoint: false) ~> ConditionalSplit1@(USD, CAD)\nConditionalSplit1@USD sink(saveMode:'overwrite' ) ~> USDSink\nConditionalSplit1@CAD sink(saveMode:'overwrite' ) ~> CADSink"),
	// 			Sinks: []*armdatafactory.DataFlowSink{
	// 				{
	// 					Name: to.Ptr("USDSink"),
	// 					Dataset: &armdatafactory.DatasetReference{
	// 						Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
	// 						ReferenceName: to.Ptr("USDOutput"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("CADSink"),
	// 					Dataset: &armdatafactory.DatasetReference{
	// 						Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
	// 						ReferenceName: to.Ptr("CADOutput"),
	// 					},
	// 			}},
	// 			Sources: []*armdatafactory.DataFlowSource{
	// 				{
	// 					Name: to.Ptr("USDCurrency"),
	// 					Dataset: &armdatafactory.DatasetReference{
	// 						Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
	// 						ReferenceName: to.Ptr("CurrencyDatasetUSD"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("CADSource"),
	// 					Dataset: &armdatafactory.DatasetReference{
	// 						Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
	// 						ReferenceName: to.Ptr("CurrencyDatasetCAD"),
	// 					},
	// 			}},
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4afa6837cfb404d8e5ffa8a604a5e09996d6f79e/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlows_Delete.json
func ExampleDataFlowsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewDataFlowsClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleDataFlow", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4afa6837cfb404d8e5ffa8a604a5e09996d6f79e/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlows_ListByFactory.json
func ExampleDataFlowsClient_NewListByFactoryPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewDataFlowsClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByFactoryPager("exampleResourceGroup", "exampleFactoryName", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.DataFlowListResponse = armdatafactory.DataFlowListResponse{
		// 	Value: []*armdatafactory.DataFlowResource{
		// 		{
		// 			Name: to.Ptr("exampleDataFlow"),
		// 			Type: to.Ptr("Microsoft.DataFactory/factories/dataflows"),
		// 			Etag: to.Ptr("0a0068d4-0000-0000-0000-5b245bd30000"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/dataflows/exampleDataFlow"),
		// 			Properties: &armdatafactory.MappingDataFlow{
		// 				Type: to.Ptr("MappingDataFlow"),
		// 				Description: to.Ptr("Sample demo data flow to convert currencies showing usage of union, derive and conditional split transformation."),
		// 				TypeProperties: &armdatafactory.MappingDataFlowTypeProperties{
		// 					Script: to.Ptr("source(output(PreviousConversionRate as double,Country as string,DateTime1 as string,CurrentConversionRate as double),allowSchemaDrift: false,validateSchema: false) ~> USDCurrency\nsource(output(PreviousConversionRate as double,Country as string,DateTime1 as string,CurrentConversionRate as double),allowSchemaDrift: true,validateSchema: false) ~> CADSource\nUSDCurrency, CADSource union(byName: true)~> Union\nUnion derive(NewCurrencyRate = round(CurrentConversionRate*1.25)) ~> NewCurrencyColumn\nNewCurrencyColumn split(Country == 'USD',Country == 'CAD',disjoint: false) ~> ConditionalSplit1@(USD, CAD)\nConditionalSplit1@USD sink(saveMode:'overwrite' ) ~> USDSink\nConditionalSplit1@CAD sink(saveMode:'overwrite' ) ~> CADSink"),
		// 					Sinks: []*armdatafactory.DataFlowSink{
		// 						{
		// 							Name: to.Ptr("USDSink"),
		// 							Dataset: &armdatafactory.DatasetReference{
		// 								Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
		// 								ReferenceName: to.Ptr("USDOutput"),
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("CADSink"),
		// 							Dataset: &armdatafactory.DatasetReference{
		// 								Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
		// 								ReferenceName: to.Ptr("CADOutput"),
		// 							},
		// 					}},
		// 					Sources: []*armdatafactory.DataFlowSource{
		// 						{
		// 							Name: to.Ptr("USDCurrency"),
		// 							Dataset: &armdatafactory.DatasetReference{
		// 								Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
		// 								ReferenceName: to.Ptr("CurrencyDatasetUSD"),
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("CADSource"),
		// 							Dataset: &armdatafactory.DatasetReference{
		// 								Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
		// 								ReferenceName: to.Ptr("CurrencyDatasetCAD"),
		// 							},
		// 					}},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
