// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armcontainerorchestratorruntime_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerorchestratorruntime/armcontainerorchestratorruntime"
	"log"
)

// Generated from example definition: 2024-03-01/StorageClass_CreateOrUpdate.json
func ExampleStorageClassClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerorchestratorruntime.NewClientFactory(, cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewStorageClassClient().BeginCreateOrUpdate(ctx, "subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1", "testrwx", armcontainerorchestratorruntime.StorageClassResource{
		Properties: &armcontainerorchestratorruntime.StorageClassProperties{
			TypeProperties: &armcontainerorchestratorruntime.RwxStorageClassTypeProperties{
				Type: to.Ptr("RWX"),
				BackingStorageClassName: to.Ptr("default"),
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
	// res = armcontainerorchestratorruntime.StorageClassClientCreateOrUpdateResponse{
	// 	StorageClassResource: &armcontainerorchestratorruntime.StorageClassResource{
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1/providers/Microsoft.KubernetesRuntime/storageclasses/testrwx"),
	// 		Name: to.Ptr("testrwx"),
	// 		Type: to.Ptr("microsoft.kubernetesruntime/storageclass"),
	// 		Properties: &armcontainerorchestratorruntime.StorageClassProperties{
	// 			Performance: to.Ptr(armcontainerorchestratorruntime.PerformanceTierBasic),
	// 			TypeProperties: &armcontainerorchestratorruntime.RwxStorageClassTypeProperties{
	// 				Type: to.Ptr("RWX"),
	// 				BackingStorageClassName: to.Ptr("default"),
	// 			},
	// 			AccessModes: []*armcontainerorchestratorruntime.AccessMode{
	// 				to.Ptr(armcontainerorchestratorruntime.AccessModeReadWriteOnce),
	// 				to.Ptr(armcontainerorchestratorruntime.AccessModeReadWriteMany),
	// 			},
	// 			AllowVolumeExpansion: to.Ptr(armcontainerorchestratorruntime.VolumeExpansionAllow),
	// 			Provisioner: to.Ptr("rwx.csi.microsoft.com"),
	// 			VolumeBindingMode: to.Ptr(armcontainerorchestratorruntime.VolumeBindingModeImmediate),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-03-01/StorageClass_Delete.json
func ExampleStorageClassClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerorchestratorruntime.NewClientFactory(, cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewStorageClassClient().BeginDelete(ctx, "subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1", "testrwx", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: 2024-03-01/StorageClass_Get.json
func ExampleStorageClassClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerorchestratorruntime.NewClientFactory(, cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStorageClassClient().Get(ctx, "subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1", "testrwx", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerorchestratorruntime.StorageClassClientGetResponse{
	// 	StorageClassResource: &armcontainerorchestratorruntime.StorageClassResource{
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1/providers/Microsoft.KubernetesRuntime/storageclasses/testrwx"),
	// 		Name: to.Ptr("testrwx"),
	// 		Type: to.Ptr("microsoft.kubernetesruntime/storageclass"),
	// 		Properties: &armcontainerorchestratorruntime.StorageClassProperties{
	// 			Performance: to.Ptr(armcontainerorchestratorruntime.PerformanceTierBasic),
	// 			TypeProperties: &armcontainerorchestratorruntime.RwxStorageClassTypeProperties{
	// 				Type: to.Ptr("RWX"),
	// 				BackingStorageClassName: to.Ptr("default"),
	// 			},
	// 			AccessModes: []*armcontainerorchestratorruntime.AccessMode{
	// 				to.Ptr(armcontainerorchestratorruntime.AccessModeReadWriteOnce),
	// 				to.Ptr(armcontainerorchestratorruntime.AccessModeReadWriteMany),
	// 			},
	// 			AllowVolumeExpansion: to.Ptr(armcontainerorchestratorruntime.VolumeExpansionAllow),
	// 			Provisioner: to.Ptr("rwx.csi.microsoft.com"),
	// 			VolumeBindingMode: to.Ptr(armcontainerorchestratorruntime.VolumeBindingModeImmediate),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-03-01/StorageClass_List.json
func ExampleStorageClassClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerorchestratorruntime.NewClientFactory(, cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStorageClassClient().NewListPager("subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1", nil)
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
		// page = armcontainerorchestratorruntime.StorageClassClientListResponse{
		// 	StorageClassResourceListResult: armcontainerorchestratorruntime.StorageClassResourceListResult{
		// 		Value: []*armcontainerorchestratorruntime.StorageClassResource{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1/providers/Microsoft.KubernetesRuntime/storageclasses/testrwx"),
		// 				Name: to.Ptr("testrwx"),
		// 				Type: to.Ptr("microsoft.kubernetesruntime/storageclass"),
		// 				Properties: &armcontainerorchestratorruntime.StorageClassProperties{
		// 					ProvisioningState: to.Ptr(armcontainerorchestratorruntime.ProvisioningStateSucceeded),
		// 					Performance: to.Ptr(armcontainerorchestratorruntime.PerformanceTierBasic),
		// 					TypeProperties: &armcontainerorchestratorruntime.RwxStorageClassTypeProperties{
		// 						Type: to.Ptr("RWX"),
		// 						BackingStorageClassName: to.Ptr("default"),
		// 					},
		// 					AccessModes: []*armcontainerorchestratorruntime.AccessMode{
		// 						to.Ptr(armcontainerorchestratorruntime.AccessModeReadWriteOnce),
		// 						to.Ptr(armcontainerorchestratorruntime.AccessModeReadWriteMany),
		// 					},
		// 					AllowVolumeExpansion: to.Ptr(armcontainerorchestratorruntime.VolumeExpansionAllow),
		// 					Provisioner: to.Ptr("rwx.csi.microsoft.com"),
		// 					VolumeBindingMode: to.Ptr(armcontainerorchestratorruntime.VolumeBindingModeImmediate),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

// Generated from example definition: 2024-03-01/StorageClass_Update.json
func ExampleStorageClassClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerorchestratorruntime.NewClientFactory(, cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewStorageClassClient().BeginUpdate(ctx, "subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1", "testrwx", armcontainerorchestratorruntime.StorageClassResourceUpdate{
		Properties: &armcontainerorchestratorruntime.StorageClassPropertiesUpdate{
			TypeProperties: &armcontainerorchestratorruntime.StorageClassTypePropertiesUpdate{
				BackingStorageClassName: to.Ptr("default"),
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
	// res = armcontainerorchestratorruntime.StorageClassClientUpdateResponse{
	// 	StorageClassResource: &armcontainerorchestratorruntime.StorageClassResource{
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1/providers/Microsoft.KubernetesRuntime/storageclasses/testrwx"),
	// 		Name: to.Ptr("testrwx"),
	// 		Type: to.Ptr("microsoft.kubernetesruntime/storageclass"),
	// 		Properties: &armcontainerorchestratorruntime.StorageClassProperties{
	// 			Performance: to.Ptr(armcontainerorchestratorruntime.PerformanceTierBasic),
	// 			TypeProperties: &armcontainerorchestratorruntime.RwxStorageClassTypeProperties{
	// 				Type: to.Ptr("RWX"),
	// 				BackingStorageClassName: to.Ptr("default"),
	// 			},
	// 			AccessModes: []*armcontainerorchestratorruntime.AccessMode{
	// 				to.Ptr(armcontainerorchestratorruntime.AccessModeReadWriteOnce),
	// 				to.Ptr(armcontainerorchestratorruntime.AccessModeReadWriteMany),
	// 			},
	// 			AllowVolumeExpansion: to.Ptr(armcontainerorchestratorruntime.VolumeExpansionAllow),
	// 			Provisioner: to.Ptr("rwx.csi.microsoft.com"),
	// 			VolumeBindingMode: to.Ptr(armcontainerorchestratorruntime.VolumeBindingModeImmediate),
	// 		},
	// 	},
	// }
}

