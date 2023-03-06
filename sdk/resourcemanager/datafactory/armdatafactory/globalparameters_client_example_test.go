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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4afa6837cfb404d8e5ffa8a604a5e09996d6f79e/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/GlobalParameters_ListByFactory.json
func ExampleGlobalParametersClient_NewListByFactoryPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewGlobalParametersClient("12345678-1234-1234-1234-12345678abc", cred, nil)
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
		// page.GlobalParameterListResponse = armdatafactory.GlobalParameterListResponse{
		// 	Value: []*armdatafactory.GlobalParameterResource{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.DataFactory/factories/globalParameters"),
		// 			Etag: to.Ptr("da00a1c3-0000-0400-0000-6241f3290000"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/globalParameters/default"),
		// 			Properties: map[string]*armdatafactory.GlobalParameterSpecification{
		// 				"copyPipelineTest": &armdatafactory.GlobalParameterSpecification{
		// 					Type: to.Ptr(armdatafactory.GlobalParameterTypeObject),
		// 					Value: map[string]any{
		// 						"mySinkDatasetFolderPath": "exampleOutput",
		// 						"mySourceDatasetFolderPath": "exampleInput/",
		// 						"testingEmptyContextParams": "",
		// 					},
		// 				},
		// 				"mySourceDatasetFolderPath": &armdatafactory.GlobalParameterSpecification{
		// 					Type: to.Ptr(armdatafactory.GlobalParameterTypeString),
		// 					Value: "input",
		// 				},
		// 				"url": &armdatafactory.GlobalParameterSpecification{
		// 					Type: to.Ptr(armdatafactory.GlobalParameterTypeString),
		// 					Value: "https://testuri.test",
		// 				},
		// 				"validADFOffice365Uris": &armdatafactory.GlobalParameterSpecification{
		// 					Type: to.Ptr(armdatafactory.GlobalParameterTypeArray),
		// 					Value: []any{
		// 						"https://testuri.test",
		// 						"https://testuri.test",
		// 					},
		// 				},
		// 				"waitTime": &armdatafactory.GlobalParameterSpecification{
		// 					Type: to.Ptr(armdatafactory.GlobalParameterTypeInt),
		// 					Value: float64(5),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4afa6837cfb404d8e5ffa8a604a5e09996d6f79e/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/GlobalParameters_Get.json
func ExampleGlobalParametersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewGlobalParametersClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "exampleResourceGroup", "exampleFactoryName", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GlobalParameterResource = armdatafactory.GlobalParameterResource{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/globalParameters"),
	// 	Etag: to.Ptr("72001a6a-0000-0400-0000-623d058f0000"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/globalParameters/default"),
	// 	Properties: map[string]*armdatafactory.GlobalParameterSpecification{
	// 		"waitTime": &armdatafactory.GlobalParameterSpecification{
	// 			Type: to.Ptr(armdatafactory.GlobalParameterTypeInt),
	// 			Value: float64(10),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4afa6837cfb404d8e5ffa8a604a5e09996d6f79e/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/GlobalParameters_Create.json
func ExampleGlobalParametersClient_CreateOrUpdate_globalParametersCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewGlobalParametersClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "exampleResourceGroup", "exampleFactoryName", "default", armdatafactory.GlobalParameterResource{
		Properties: map[string]*armdatafactory.GlobalParameterSpecification{
			"waitTime": {
				Type:  to.Ptr(armdatafactory.GlobalParameterTypeInt),
				Value: float64(5),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GlobalParameterResource = armdatafactory.GlobalParameterResource{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/globalParameters"),
	// 	Etag: to.Ptr("0a008ad4-0000-0000-0000-5b245c6e0000"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/globalParameters/default"),
	// 	Properties: map[string]*armdatafactory.GlobalParameterSpecification{
	// 		"waitTime": &armdatafactory.GlobalParameterSpecification{
	// 			Type: to.Ptr(armdatafactory.GlobalParameterTypeInt),
	// 			Value: float64(5),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4afa6837cfb404d8e5ffa8a604a5e09996d6f79e/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/GlobalParameters_Update.json
func ExampleGlobalParametersClient_CreateOrUpdate_globalParametersUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewGlobalParametersClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "exampleResourceGroup", "exampleFactoryName", "default", armdatafactory.GlobalParameterResource{
		Properties: map[string]*armdatafactory.GlobalParameterSpecification{
			"waitTime": {
				Type:  to.Ptr(armdatafactory.GlobalParameterTypeInt),
				Value: float64(5),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GlobalParameterResource = armdatafactory.GlobalParameterResource{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/globalParameters"),
	// 	Etag: to.Ptr("0a008ad4-0000-0000-0000-5b245c6e0000"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/globalParameters/default"),
	// 	Properties: map[string]*armdatafactory.GlobalParameterSpecification{
	// 		"waitTime": &armdatafactory.GlobalParameterSpecification{
	// 			Type: to.Ptr(armdatafactory.GlobalParameterTypeInt),
	// 			Value: float64(5),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4afa6837cfb404d8e5ffa8a604a5e09996d6f79e/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/GlobalParameters_Delete.json
func ExampleGlobalParametersClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewGlobalParametersClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx, "exampleResourceGroup", "exampleFactoryName", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
