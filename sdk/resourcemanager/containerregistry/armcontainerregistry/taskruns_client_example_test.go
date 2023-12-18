//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/TaskRunsGet.json
func ExampleTaskRunsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTaskRunsClient().Get(ctx, "myResourceGroup", "myRegistry", "myRun", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TaskRun = armcontainerregistry.TaskRun{
	// 	Name: to.Ptr("myRun"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/TaskRuns"),
	// 	ID: to.Ptr("/subscriptions/3647315e-0c5b-4ce4-8739-b071e144b2c9/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/taskRuns/myRun"),
	// 	Properties: &armcontainerregistry.TaskRunProperties{
	// 		ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		RunRequest: &armcontainerregistry.EncodedTaskRunRequest{
	// 			Type: to.Ptr("EncodedTaskRunRequest"),
	// 			IsArchiveEnabled: to.Ptr(true),
	// 			Credentials: &armcontainerregistry.Credentials{
	// 			},
	// 			EncodedTaskContent: to.Ptr("c3RlcHM6IAogIC0gY21kOiB7eyAuVmFsdWVzLmNvbW1hbmQgfX0K"),
	// 			EncodedValuesContent: to.Ptr("Y29tbWFuZDogYmFzaCBlY2hvIHt7LlJ1bi5SZWdpc3RyeX19Cg=="),
	// 			Platform: &armcontainerregistry.PlatformProperties{
	// 				Architecture: to.Ptr(armcontainerregistry.ArchitectureAmd64),
	// 				OS: to.Ptr(armcontainerregistry.OSLinux),
	// 			},
	// 			Values: []*armcontainerregistry.SetValue{
	// 			},
	// 		},
	// 		RunResult: &armcontainerregistry.Run{
	// 			Name: to.Ptr("yd5"),
	// 			Type: to.Ptr("Microsoft.ContainerRegistry/registries/runs"),
	// 			ID: to.Ptr("/subscriptions/3647315e-0c5b-4ce4-8739-b071e144b2c9/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/runs/yd5"),
	// 			Properties: &armcontainerregistry.RunProperties{
	// 				AgentConfiguration: &armcontainerregistry.AgentProperties{
	// 					CPU: to.Ptr[int32](2),
	// 				},
	// 				CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-06T17:23:21.926Z"); return t}()),
	// 				FinishTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-06T17:23:29.879Z"); return t}()),
	// 				IsArchiveEnabled: to.Ptr(true),
	// 				LastUpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-06T17:23:29.000Z"); return t}()),
	// 				Platform: &armcontainerregistry.PlatformProperties{
	// 					Architecture: to.Ptr(armcontainerregistry.ArchitectureAmd64),
	// 					OS: to.Ptr(armcontainerregistry.OSLinux),
	// 				},
	// 				ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 				RunID: to.Ptr("yd5"),
	// 				RunType: to.Ptr(armcontainerregistry.RunTypeQuickRun),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-06T17:23:22.134Z"); return t}()),
	// 				Status: to.Ptr(armcontainerregistry.RunStatusSucceeded),
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/TaskRunsCreate.json
func ExampleTaskRunsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTaskRunsClient().BeginCreate(ctx, "myResourceGroup", "myRegistry", "myRun", armcontainerregistry.TaskRun{
		Properties: &armcontainerregistry.TaskRunProperties{
			ForceUpdateTag: to.Ptr("test"),
			RunRequest: &armcontainerregistry.EncodedTaskRunRequest{
				Type:                 to.Ptr("EncodedTaskRunRequest"),
				Credentials:          &armcontainerregistry.Credentials{},
				EncodedTaskContent:   to.Ptr("c3RlcHM6IAogIC0gY21kOiB7eyAuVmFsdWVzLmNvbW1hbmQgfX0K"),
				EncodedValuesContent: to.Ptr("Y29tbWFuZDogYmFzaCBlY2hvIHt7LlJ1bi5SZWdpc3RyeX19Cg=="),
				Platform: &armcontainerregistry.PlatformProperties{
					Architecture: to.Ptr(armcontainerregistry.ArchitectureAmd64),
					OS:           to.Ptr(armcontainerregistry.OSLinux),
				},
				Values: []*armcontainerregistry.SetValue{},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TaskRun = armcontainerregistry.TaskRun{
	// 	Name: to.Ptr("myrun"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/TaskRuns"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/taskRuns/myRun"),
	// 	Properties: &armcontainerregistry.TaskRunProperties{
	// 		ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		RunRequest: &armcontainerregistry.EncodedTaskRunRequest{
	// 			Type: to.Ptr("EncodedTaskRunRequest"),
	// 			IsArchiveEnabled: to.Ptr(true),
	// 			Credentials: &armcontainerregistry.Credentials{
	// 			},
	// 			EncodedTaskContent: to.Ptr("c3RlcHM6IAogIC0gY21kOiB7eyAuVmFsdWVzLmNvbW1hbmQgfX0K"),
	// 			EncodedValuesContent: to.Ptr("Y29tbWFuZDogYmFzaCBlY2hvIHt7LlJ1bi5SZWdpc3RyeX19Cg=="),
	// 			Platform: &armcontainerregistry.PlatformProperties{
	// 				Architecture: to.Ptr(armcontainerregistry.ArchitectureAmd64),
	// 				OS: to.Ptr(armcontainerregistry.OSLinux),
	// 			},
	// 			Values: []*armcontainerregistry.SetValue{
	// 			},
	// 		},
	// 		RunResult: &armcontainerregistry.Run{
	// 			Name: to.Ptr("yd5"),
	// 			Type: to.Ptr("Microsoft.ContainerRegistry/registries/runs"),
	// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/runs/yd5"),
	// 			Properties: &armcontainerregistry.RunProperties{
	// 				AgentConfiguration: &armcontainerregistry.AgentProperties{
	// 					CPU: to.Ptr[int32](2),
	// 				},
	// 				CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-06T17:23:21.926Z"); return t}()),
	// 				IsArchiveEnabled: to.Ptr(true),
	// 				LastUpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-06T17:23:21.000Z"); return t}()),
	// 				Platform: &armcontainerregistry.PlatformProperties{
	// 					Architecture: to.Ptr(armcontainerregistry.ArchitectureAmd64),
	// 					OS: to.Ptr(armcontainerregistry.OSLinux),
	// 				},
	// 				ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 				RunID: to.Ptr("yd5"),
	// 				RunType: to.Ptr(armcontainerregistry.RunTypeQuickRun),
	// 				Status: to.Ptr(armcontainerregistry.RunStatusQueued),
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/TaskRunsDelete.json
func ExampleTaskRunsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTaskRunsClient().BeginDelete(ctx, "myResourceGroup", "myRegistry", "myRun", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/TaskRunsUpdate.json
func ExampleTaskRunsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTaskRunsClient().BeginUpdate(ctx, "myResourceGroup", "myRegistry", "myRun", armcontainerregistry.TaskRunUpdateParameters{
		Properties: &armcontainerregistry.TaskRunPropertiesUpdateParameters{
			ForceUpdateTag: to.Ptr("test"),
			RunRequest: &armcontainerregistry.EncodedTaskRunRequest{
				Type:                 to.Ptr("EncodedTaskRunRequest"),
				IsArchiveEnabled:     to.Ptr(true),
				Credentials:          &armcontainerregistry.Credentials{},
				EncodedTaskContent:   to.Ptr("c3RlcHM6IAogIC0gY21kOiB7eyAuVmFsdWVzLmNvbW1hbmQgfX0K"),
				EncodedValuesContent: to.Ptr("Y29tbWFuZDogYmFzaCBlY2hvIHt7LlJ1bi5SZWdpc3RyeX19Cg=="),
				Platform: &armcontainerregistry.PlatformProperties{
					Architecture: to.Ptr(armcontainerregistry.ArchitectureAmd64),
					OS:           to.Ptr(armcontainerregistry.OSLinux),
				},
				Values: []*armcontainerregistry.SetValue{},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TaskRun = armcontainerregistry.TaskRun{
	// 	Name: to.Ptr("mytestrun"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/TaskRuns"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/taskRuns/myRun"),
	// 	Properties: &armcontainerregistry.TaskRunProperties{
	// 		ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		RunRequest: &armcontainerregistry.EncodedTaskRunRequest{
	// 			Type: to.Ptr("EncodedTaskRunRequest"),
	// 			IsArchiveEnabled: to.Ptr(true),
	// 			Credentials: &armcontainerregistry.Credentials{
	// 			},
	// 			EncodedTaskContent: to.Ptr("c3RlcHM6IAogIC0gY21kOiB7eyAuVmFsdWVzLmNvbW1hbmQgfX0K"),
	// 			EncodedValuesContent: to.Ptr("Y29tbWFuZDogYmFzaCBlY2hvIHt7LlJ1bi5SZWdpc3RyeX19Cg=="),
	// 			Platform: &armcontainerregistry.PlatformProperties{
	// 				Architecture: to.Ptr(armcontainerregistry.ArchitectureAmd64),
	// 				OS: to.Ptr(armcontainerregistry.OSLinux),
	// 			},
	// 			Values: []*armcontainerregistry.SetValue{
	// 			},
	// 		},
	// 		RunResult: &armcontainerregistry.Run{
	// 			Name: to.Ptr("yd6"),
	// 			Type: to.Ptr("Microsoft.ContainerRegistry/registries/runs"),
	// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/runs/yd6"),
	// 			Properties: &armcontainerregistry.RunProperties{
	// 				AgentConfiguration: &armcontainerregistry.AgentProperties{
	// 					CPU: to.Ptr[int32](2),
	// 				},
	// 				CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-06T17:31:34.118Z"); return t}()),
	// 				IsArchiveEnabled: to.Ptr(true),
	// 				LastUpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-06T17:31:34.000Z"); return t}()),
	// 				Platform: &armcontainerregistry.PlatformProperties{
	// 					Architecture: to.Ptr(armcontainerregistry.ArchitectureAmd64),
	// 					OS: to.Ptr(armcontainerregistry.OSLinux),
	// 				},
	// 				ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 				RunID: to.Ptr("yd6"),
	// 				RunType: to.Ptr(armcontainerregistry.RunTypeQuickRun),
	// 				Status: to.Ptr(armcontainerregistry.RunStatusQueued),
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/TaskRunsGetDetails.json
func ExampleTaskRunsClient_GetDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTaskRunsClient().GetDetails(ctx, "myResourceGroup", "myRegistry", "myRun", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TaskRun = armcontainerregistry.TaskRun{
	// 	Name: to.Ptr("myRun"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/TaskRuns"),
	// 	ID: to.Ptr("/subscriptions/3647315e-0c5b-4ce4-8739-b071e144b2c9/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/taskRuns/myRun"),
	// 	Properties: &armcontainerregistry.TaskRunProperties{
	// 		ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		RunRequest: &armcontainerregistry.EncodedTaskRunRequest{
	// 			Type: to.Ptr("EncodedTaskRunRequest"),
	// 			IsArchiveEnabled: to.Ptr(true),
	// 			Credentials: &armcontainerregistry.Credentials{
	// 			},
	// 			EncodedTaskContent: to.Ptr("c3RlcHM6IAogIC0gY21kOiB7eyAuVmFsdWVzLmNvbW1hbmQgfX0K"),
	// 			EncodedValuesContent: to.Ptr("Y29tbWFuZDogYmFzaCBlY2hvIHt7LlJ1bi5SZWdpc3RyeX19Cg=="),
	// 			Platform: &armcontainerregistry.PlatformProperties{
	// 				Architecture: to.Ptr(armcontainerregistry.ArchitectureAmd64),
	// 				OS: to.Ptr(armcontainerregistry.OSLinux),
	// 			},
	// 			Values: []*armcontainerregistry.SetValue{
	// 			},
	// 		},
	// 		RunResult: &armcontainerregistry.Run{
	// 			Name: to.Ptr("yd5"),
	// 			Type: to.Ptr("Microsoft.ContainerRegistry/registries/runs"),
	// 			ID: to.Ptr("/subscriptions/3647315e-0c5b-4ce4-8739-b071e144b2c9/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/runs/yd5"),
	// 			Properties: &armcontainerregistry.RunProperties{
	// 				AgentConfiguration: &armcontainerregistry.AgentProperties{
	// 					CPU: to.Ptr[int32](2),
	// 				},
	// 				CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-06T17:23:21.926Z"); return t}()),
	// 				FinishTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-06T17:23:29.879Z"); return t}()),
	// 				IsArchiveEnabled: to.Ptr(true),
	// 				LastUpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-06T17:23:29.000Z"); return t}()),
	// 				Platform: &armcontainerregistry.PlatformProperties{
	// 					Architecture: to.Ptr(armcontainerregistry.ArchitectureAmd64),
	// 					OS: to.Ptr(armcontainerregistry.OSLinux),
	// 				},
	// 				ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 				RunID: to.Ptr("yd5"),
	// 				RunType: to.Ptr(armcontainerregistry.RunTypeQuickRun),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-06T17:23:22.134Z"); return t}()),
	// 				Status: to.Ptr(armcontainerregistry.RunStatusSucceeded),
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/TaskRunsList.json
func ExampleTaskRunsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTaskRunsClient().NewListPager("myResourceGroup", "myRegistry", nil)
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
		// page.TaskRunListResult = armcontainerregistry.TaskRunListResult{
		// 	Value: []*armcontainerregistry.TaskRun{
		// 		{
		// 			Name: to.Ptr("mytestrun"),
		// 			Type: to.Ptr("Microsoft.ContainerRegistry/registries/TaskRuns"),
		// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/taskRuns/myRun"),
		// 			Properties: &armcontainerregistry.TaskRunProperties{
		// 				ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
		// 				RunRequest: &armcontainerregistry.EncodedTaskRunRequest{
		// 					Type: to.Ptr("EncodedTaskRunRequest"),
		// 					IsArchiveEnabled: to.Ptr(true),
		// 					Credentials: &armcontainerregistry.Credentials{
		// 					},
		// 					EncodedTaskContent: to.Ptr("c3RlcHM6IAogIC0gY21kOiB7eyAuVmFsdWVzLmNvbW1hbmQgfX0K"),
		// 					EncodedValuesContent: to.Ptr("Y29tbWFuZDogYmFzaCBlY2hvIHt7LlJ1bi5SZWdpc3RyeX19Cg=="),
		// 					Platform: &armcontainerregistry.PlatformProperties{
		// 						Architecture: to.Ptr(armcontainerregistry.ArchitectureAmd64),
		// 						OS: to.Ptr(armcontainerregistry.OSLinux),
		// 					},
		// 					Values: []*armcontainerregistry.SetValue{
		// 					},
		// 				},
		// 				RunResult: &armcontainerregistry.Run{
		// 					Name: to.Ptr("yd4"),
		// 					Type: to.Ptr("Microsoft.ContainerRegistry/registries/runs"),
		// 					ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/runs/yd4"),
		// 					Properties: &armcontainerregistry.RunProperties{
		// 						AgentConfiguration: &armcontainerregistry.AgentProperties{
		// 							CPU: to.Ptr[int32](2),
		// 						},
		// 						CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-04T17:15:29.227Z"); return t}()),
		// 						FinishTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-04T17:15:37.034Z"); return t}()),
		// 						IsArchiveEnabled: to.Ptr(true),
		// 						LastUpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-04T17:15:37.000Z"); return t}()),
		// 						Platform: &armcontainerregistry.PlatformProperties{
		// 							Architecture: to.Ptr(armcontainerregistry.ArchitectureAmd64),
		// 							OS: to.Ptr(armcontainerregistry.OSLinux),
		// 						},
		// 						ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
		// 						RunID: to.Ptr("yd4"),
		// 						RunType: to.Ptr(armcontainerregistry.RunTypeQuickRun),
		// 						StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-04T17:15:29.458Z"); return t}()),
		// 						Status: to.Ptr(armcontainerregistry.RunStatusSucceeded),
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
