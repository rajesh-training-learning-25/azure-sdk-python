//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-05-01/examples/OnlineDeployment/list.json
func ExampleOnlineDeploymentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewOnlineDeploymentsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("test-rg",
		"my-aml-workspace",
		"testEndpointName",
		&armmachinelearning.OnlineDeploymentsClientListOptions{OrderBy: to.Ptr("string"),
			Top:  to.Ptr[int32](1),
			Skip: nil,
		})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-05-01/examples/OnlineDeployment/delete.json
func ExampleOnlineDeploymentsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewOnlineDeploymentsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		"testrg123",
		"workspace123",
		"testEndpoint",
		"testDeployment",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-05-01/examples/OnlineDeployment/KubernetesOnlineDeployment/get.json
func ExampleOnlineDeploymentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewOnlineDeploymentsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"test-rg",
		"my-aml-workspace",
		"testEndpointName",
		"testDeploymentName",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-05-01/examples/OnlineDeployment/KubernetesOnlineDeployment/update.json
func ExampleOnlineDeploymentsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewOnlineDeploymentsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"test-rg",
		"my-aml-workspace",
		"testEndpointName",
		"testDeploymentName",
		armmachinelearning.PartialMinimalTrackedResourceWithSKU{
			Tags: map[string]*string{},
			SKU: &armmachinelearning.PartialSKU{
				Name:     to.Ptr("string"),
				Capacity: to.Ptr[int32](1),
				Family:   to.Ptr("string"),
				Size:     to.Ptr("string"),
				Tier:     to.Ptr(armmachinelearning.SKUTierFree),
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-05-01/examples/OnlineDeployment/KubernetesOnlineDeployment/createOrUpdate.json
func ExampleOnlineDeploymentsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewOnlineDeploymentsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"test-rg",
		"my-aml-workspace",
		"testEndpointName",
		"testDeploymentName",
		armmachinelearning.OnlineDeployment{
			Location: to.Ptr("string"),
			Tags:     map[string]*string{},
			Identity: &armmachinelearning.ManagedServiceIdentity{
				Type: to.Ptr(armmachinelearning.ManagedServiceIdentityTypeSystemAssigned),
				UserAssignedIdentities: map[string]*armmachinelearning.UserAssignedIdentity{
					"string": {},
				},
			},
			Kind: to.Ptr("string"),
			Properties: &armmachinelearning.KubernetesOnlineDeployment{
				Description: to.Ptr("string"),
				CodeConfiguration: &armmachinelearning.CodeConfiguration{
					CodeID:        to.Ptr("string"),
					ScoringScript: to.Ptr("string"),
				},
				EnvironmentID: to.Ptr("string"),
				EnvironmentVariables: map[string]*string{
					"string": to.Ptr("string"),
				},
				Properties: map[string]*string{
					"string": to.Ptr("string"),
				},
				AppInsightsEnabled:  to.Ptr(false),
				EndpointComputeType: to.Ptr(armmachinelearning.EndpointComputeTypeKubernetes),
				InstanceType:        to.Ptr("string"),
				LivenessProbe: &armmachinelearning.ProbeSettings{
					FailureThreshold: to.Ptr[int32](1),
					InitialDelay:     to.Ptr("PT5M"),
					Period:           to.Ptr("PT5M"),
					SuccessThreshold: to.Ptr[int32](1),
					Timeout:          to.Ptr("PT5M"),
				},
				Model:          to.Ptr("string"),
				ModelMountPath: to.Ptr("string"),
				RequestSettings: &armmachinelearning.OnlineRequestSettings{
					MaxConcurrentRequestsPerInstance: to.Ptr[int32](1),
					MaxQueueWait:                     to.Ptr("PT5M"),
					RequestTimeout:                   to.Ptr("PT5M"),
				},
				ScaleSettings: &armmachinelearning.DefaultScaleSettings{
					ScaleType: to.Ptr(armmachinelearning.ScaleTypeDefault),
				},
				ContainerResourceRequirements: &armmachinelearning.ContainerResourceRequirements{
					ContainerResourceLimits: &armmachinelearning.ContainerResourceSettings{
						CPU:    to.Ptr("\"1\""),
						Gpu:    to.Ptr("\"1\""),
						Memory: to.Ptr("\"2Gi\""),
					},
					ContainerResourceRequests: &armmachinelearning.ContainerResourceSettings{
						CPU:    to.Ptr("\"1\""),
						Gpu:    to.Ptr("\"1\""),
						Memory: to.Ptr("\"2Gi\""),
					},
				},
			},
			SKU: &armmachinelearning.SKU{
				Name:     to.Ptr("string"),
				Capacity: to.Ptr[int32](1),
				Family:   to.Ptr("string"),
				Size:     to.Ptr("string"),
				Tier:     to.Ptr(armmachinelearning.SKUTierFree),
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-05-01/examples/OnlineDeployment/getLogs.json
func ExampleOnlineDeploymentsClient_GetLogs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewOnlineDeploymentsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetLogs(ctx,
		"testrg123",
		"workspace123",
		"testEndpoint",
		"testDeployment",
		armmachinelearning.DeploymentLogsRequest{
			ContainerType: to.Ptr(armmachinelearning.ContainerTypeStorageInitializer),
			Tail:          to.Ptr[int32](0),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
