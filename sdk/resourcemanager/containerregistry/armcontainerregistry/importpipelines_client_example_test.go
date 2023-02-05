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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/270d3cd664cca3ddc8511f92d3851a715e2c61db/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-01-01-preview/examples/ImportPipelineList.json
func ExampleImportPipelinesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewImportPipelinesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("myResourceGroup", "myRegistry", nil)
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
		// page.ImportPipelineListResult = armcontainerregistry.ImportPipelineListResult{
		// 	Value: []*armcontainerregistry.ImportPipeline{
		// 		{
		// 			Name: to.Ptr("myImportPipeline"),
		// 			Type: to.Ptr("Microsoft.ContainerRegistry/registries/importPipelines"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/importPipelines/myImportPipeline"),
		// 			Identity: &armcontainerregistry.IdentityProperties{
		// 				Type: to.Ptr(armcontainerregistry.ResourceIdentityTypeUserAssigned),
		// 				UserAssignedIdentities: map[string]*armcontainerregistry.UserIdentityProperties{
		// 					"/subscriptions/f9d7ebed-adbd-4cb4-b973-aaf82c136138/resourcegroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity2": &armcontainerregistry.UserIdentityProperties{
		// 						ClientID: to.Ptr("d3ce1bc2-f7d7-4a5b-9979-950f4e57680e"),
		// 						PrincipalID: to.Ptr("b6p9f58b-6fbf-4efd-a7e0-fvd46911a466"),
		// 					},
		// 				},
		// 			},
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armcontainerregistry.ImportPipelineProperties{
		// 				Options: []*armcontainerregistry.PipelineOptions{
		// 					to.Ptr(armcontainerregistry.PipelineOptionsOverwriteTags),
		// 					to.Ptr(armcontainerregistry.PipelineOptionsDeleteSourceBlobOnSuccess),
		// 					to.Ptr(armcontainerregistry.PipelineOptionsContinueOnErrors)},
		// 					ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
		// 					Source: &armcontainerregistry.ImportPipelineSourceProperties{
		// 						Type: to.Ptr(armcontainerregistry.PipelineSourceTypeAzureStorageBlobContainer),
		// 						KeyVaultURI: to.Ptr("https://myvault.vault.azure.net/secrets/acrimportsas"),
		// 						URI: to.Ptr("https://accountname.blob.core.windows.net/containername"),
		// 					},
		// 					Trigger: &armcontainerregistry.PipelineTriggerProperties{
		// 						SourceTrigger: &armcontainerregistry.PipelineSourceTriggerProperties{
		// 							Status: to.Ptr(armcontainerregistry.TriggerStatusEnabled),
		// 						},
		// 					},
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/270d3cd664cca3ddc8511f92d3851a715e2c61db/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-01-01-preview/examples/ImportPipelineGet.json
func ExampleImportPipelinesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewImportPipelinesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "myResourceGroup", "myRegistry", "myImportPipeline", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ImportPipeline = armcontainerregistry.ImportPipeline{
	// 	Name: to.Ptr("myImportPipeline"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/importPipelines"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/importPipelines/myImportPipeline"),
	// 	Identity: &armcontainerregistry.IdentityProperties{
	// 		Type: to.Ptr(armcontainerregistry.ResourceIdentityTypeUserAssigned),
	// 		UserAssignedIdentities: map[string]*armcontainerregistry.UserIdentityProperties{
	// 			"/subscriptions/f9d7ebed-adbd-4cb4-b973-aaf82c136138/resourcegroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity2": &armcontainerregistry.UserIdentityProperties{
	// 				ClientID: to.Ptr("d3ce1bc2-f7d7-4a5b-9979-950f4e57680e"),
	// 				PrincipalID: to.Ptr("b6p9f58b-6fbf-4efd-a7e0-fvd46911a466"),
	// 			},
	// 		},
	// 	},
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armcontainerregistry.ImportPipelineProperties{
	// 		Options: []*armcontainerregistry.PipelineOptions{
	// 			to.Ptr(armcontainerregistry.PipelineOptionsOverwriteTags),
	// 			to.Ptr(armcontainerregistry.PipelineOptionsDeleteSourceBlobOnSuccess),
	// 			to.Ptr(armcontainerregistry.PipelineOptionsContinueOnErrors)},
	// 			ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 			Source: &armcontainerregistry.ImportPipelineSourceProperties{
	// 				Type: to.Ptr(armcontainerregistry.PipelineSourceTypeAzureStorageBlobContainer),
	// 				KeyVaultURI: to.Ptr("https://myvault.vault.azure.net/secrets/acrimportsas"),
	// 				URI: to.Ptr("https://accountname.blob.core.windows.net/containername"),
	// 			},
	// 			Trigger: &armcontainerregistry.PipelineTriggerProperties{
	// 				SourceTrigger: &armcontainerregistry.PipelineSourceTriggerProperties{
	// 					Status: to.Ptr(armcontainerregistry.TriggerStatusEnabled),
	// 				},
	// 			},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/270d3cd664cca3ddc8511f92d3851a715e2c61db/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-01-01-preview/examples/ImportPipelineCreate.json
func ExampleImportPipelinesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewImportPipelinesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx, "myResourceGroup", "myRegistry", "myImportPipeline", armcontainerregistry.ImportPipeline{
		Identity: &armcontainerregistry.IdentityProperties{
			Type: to.Ptr(armcontainerregistry.ResourceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armcontainerregistry.UserIdentityProperties{
				"/subscriptions/f9d7ebed-adbd-4cb4-b973-aaf82c136138/resourcegroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity2": {},
			},
		},
		Location: to.Ptr("westus"),
		Properties: &armcontainerregistry.ImportPipelineProperties{
			Options: []*armcontainerregistry.PipelineOptions{
				to.Ptr(armcontainerregistry.PipelineOptionsOverwriteTags),
				to.Ptr(armcontainerregistry.PipelineOptionsDeleteSourceBlobOnSuccess),
				to.Ptr(armcontainerregistry.PipelineOptionsContinueOnErrors)},
			Source: &armcontainerregistry.ImportPipelineSourceProperties{
				Type:        to.Ptr(armcontainerregistry.PipelineSourceTypeAzureStorageBlobContainer),
				KeyVaultURI: to.Ptr("https://myvault.vault.azure.net/secrets/acrimportsas"),
				URI:         to.Ptr("https://accountname.blob.core.windows.net/containername"),
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
	// res.ImportPipeline = armcontainerregistry.ImportPipeline{
	// 	Name: to.Ptr("myImportPipeline"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/importPipelines"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/importPipelines/myImportPipeline"),
	// 	Identity: &armcontainerregistry.IdentityProperties{
	// 		Type: to.Ptr(armcontainerregistry.ResourceIdentityTypeUserAssigned),
	// 		UserAssignedIdentities: map[string]*armcontainerregistry.UserIdentityProperties{
	// 			"/subscriptions/f9d7ebed-adbd-4cb4-b973-aaf82c136138/resourcegroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity2": &armcontainerregistry.UserIdentityProperties{
	// 				ClientID: to.Ptr("d3ce1bc2-f7d7-4a5b-9979-950f4e57680e"),
	// 				PrincipalID: to.Ptr("b6p9f58b-6fbf-4efd-a7e0-fvd46911a466"),
	// 			},
	// 		},
	// 	},
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armcontainerregistry.ImportPipelineProperties{
	// 		Options: []*armcontainerregistry.PipelineOptions{
	// 			to.Ptr(armcontainerregistry.PipelineOptionsOverwriteTags),
	// 			to.Ptr(armcontainerregistry.PipelineOptionsDeleteSourceBlobOnSuccess),
	// 			to.Ptr(armcontainerregistry.PipelineOptionsContinueOnErrors)},
	// 			ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 			Source: &armcontainerregistry.ImportPipelineSourceProperties{
	// 				Type: to.Ptr(armcontainerregistry.PipelineSourceTypeAzureStorageBlobContainer),
	// 				KeyVaultURI: to.Ptr("https://myvault.vault.azure.net/secrets/acrimportsas"),
	// 				URI: to.Ptr("https://accountname.blob.core.windows.net/containername"),
	// 			},
	// 			Trigger: &armcontainerregistry.PipelineTriggerProperties{
	// 				SourceTrigger: &armcontainerregistry.PipelineSourceTriggerProperties{
	// 					Status: to.Ptr(armcontainerregistry.TriggerStatusEnabled),
	// 				},
	// 			},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/270d3cd664cca3ddc8511f92d3851a715e2c61db/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-01-01-preview/examples/ImportPipelineDelete.json
func ExampleImportPipelinesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewImportPipelinesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx, "myResourceGroup", "myRegistry", "myImportPipeline", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
