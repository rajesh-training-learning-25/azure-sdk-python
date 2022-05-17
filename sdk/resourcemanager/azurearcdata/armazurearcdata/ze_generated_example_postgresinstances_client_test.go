//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurearcdata_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurearcdata/armazurearcdata"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azurearcdata/resource-manager/Microsoft.AzureArcData/preview/2022-03-01-preview/examples/ListSubscriptionPostgresInstance.json
func ExamplePostgresInstancesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armazurearcdata.NewPostgresInstancesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azurearcdata/resource-manager/Microsoft.AzureArcData/preview/2022-03-01-preview/examples/ListByResourceGroupPostgresInstance.json
func ExamplePostgresInstancesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armazurearcdata.NewPostgresInstancesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByResourceGroupPager("testrg",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azurearcdata/resource-manager/Microsoft.AzureArcData/preview/2022-03-01-preview/examples/GetPostgresInstance.json
func ExamplePostgresInstancesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armazurearcdata.NewPostgresInstancesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"testrg",
		"testpostgresInstances",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azurearcdata/resource-manager/Microsoft.AzureArcData/preview/2022-03-01-preview/examples/CreateOrUpdatePostgresInstance.json
func ExamplePostgresInstancesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armazurearcdata.NewPostgresInstancesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"testrg",
		"testpostgresInstance",
		armazurearcdata.PostgresInstance{
			Location: to.Ptr("eastus"),
			ExtendedLocation: &armazurearcdata.ExtendedLocation{
				Name: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/arclocation"),
				Type: to.Ptr(armazurearcdata.ExtendedLocationTypesCustomLocation),
			},
			Properties: &armazurearcdata.PostgresInstanceProperties{
				Admin: to.Ptr("admin"),
				BasicLoginInformation: &armazurearcdata.BasicLoginInformation{
					Password: to.Ptr("********"),
					Username: to.Ptr("username"),
				},
				DataControllerID: to.Ptr("dataControllerId"),
				K8SRaw: map[string]interface{}{
					"apiVersion": "apiVersion",
					"kind":       "postgresql-12",
					"metadata": map[string]interface{}{
						"name":              "pg1",
						"creationTimestamp": "2020-08-25T14:55:10Z",
						"generation":        float64(1),
						"namespace":         "test",
						"resourceVersion":   "527780",
						"selfLink":          "/apis/arcdata.microsoft.com/v1alpha1/namespaces/test/postgresql-12s/pg1",
						"uid":               "1111aaaa-ffff-ffff-ffff-99999aaaaaaa",
					},
					"spec": map[string]interface{}{
						"backups": map[string]interface{}{
							"deltaMinutes": float64(3),
							"fullMinutes":  float64(10),
							"tiers": []interface{}{
								map[string]interface{}{
									"retention": map[string]interface{}{
										"maximums": []interface{}{
											"6",
											"512MB",
										},
										"minimums": []interface{}{
											"3",
										},
									},
									"storage": map[string]interface{}{
										"volumeSize": "1Gi",
									},
								},
							},
						},
						"engine": map[string]interface{}{
							"extensions": []interface{}{
								map[string]interface{}{
									"name": "citus",
								},
							},
						},
						"scale": map[string]interface{}{
							"shards": float64(3),
						},
						"scheduling": map[string]interface{}{
							"default": map[string]interface{}{
								"resources": map[string]interface{}{
									"requests": map[string]interface{}{
										"memory": "256Mi",
									},
								},
							},
						},
						"service": map[string]interface{}{
							"type": "NodePort",
						},
						"storage": map[string]interface{}{
							"data": map[string]interface{}{
								"className": "local-storage",
								"size":      "5Gi",
							},
							"logs": map[string]interface{}{
								"className": "local-storage",
								"size":      "5Gi",
							},
						},
					},
					"status": map[string]interface{}{
						"externalEndpoint": nil,
						"readyPods":        "4/4",
						"state":            "Ready",
					},
				},
			},
			SKU: &armazurearcdata.PostgresInstanceSKU{
				Name: to.Ptr("default"),
				Dev:  to.Ptr(true),
				Tier: to.Ptr("Hyperscale"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azurearcdata/resource-manager/Microsoft.AzureArcData/preview/2022-03-01-preview/examples/DeletePostgresInstance.json
func ExamplePostgresInstancesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armazurearcdata.NewPostgresInstancesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		"testrg",
		"testpostgresInstance",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azurearcdata/resource-manager/Microsoft.AzureArcData/preview/2022-03-01-preview/examples/UpdatePostgresInstance.json
func ExamplePostgresInstancesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armazurearcdata.NewPostgresInstancesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"testrg",
		"testpostgresInstance",
		armazurearcdata.PostgresInstanceUpdate{
			Tags: map[string]*string{
				"mytag": to.Ptr("myval"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
