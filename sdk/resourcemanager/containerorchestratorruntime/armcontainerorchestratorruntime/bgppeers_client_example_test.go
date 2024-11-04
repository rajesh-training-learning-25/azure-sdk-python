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

// Generated from example definition: 2024-03-01/BgpPeers_CreateOrUpdate.json
func ExampleBgpPeersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerorchestratorruntime.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBgpPeersClient().BeginCreateOrUpdate(ctx, "subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1", "testpeer", armcontainerorchestratorruntime.BgpPeer{
		Properties: &armcontainerorchestratorruntime.BgpPeerProperties{
			MyAsn:       to.Ptr[int32](64500),
			PeerAsn:     to.Ptr[int32](64501),
			PeerAddress: to.Ptr("10.0.0.1"),
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
	// res = armcontainerorchestratorruntime.BgpPeersClientCreateOrUpdateResponse{
	// 	BgpPeer: &armcontainerorchestratorruntime.BgpPeer{
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1/providers/Microsoft.KubernetesRuntime/bgpPeers/testpeer"),
	// 		Name: to.Ptr("testpeer"),
	// 		Type: to.Ptr("Microsoft.KubernetesRuntime/BgpPeers"),
	// 		Properties: &armcontainerorchestratorruntime.BgpPeerProperties{
	// 			ProvisioningState: to.Ptr(armcontainerorchestratorruntime.ProvisioningStateSucceeded),
	// 			MyAsn: to.Ptr[int32](64500),
	// 			PeerAsn: to.Ptr[int32](64501),
	// 			PeerAddress: to.Ptr("10.0.0.1"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-03-01/BgpPeers_Delete.json
func ExampleBgpPeersClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerorchestratorruntime.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBgpPeersClient().Delete(ctx, "subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1", "testpeer", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerorchestratorruntime.BgpPeersClientDeleteResponse{
	// }
}

// Generated from example definition: 2024-03-01/BgpPeers_Get.json
func ExampleBgpPeersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerorchestratorruntime.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBgpPeersClient().Get(ctx, "subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1", "testpeer", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerorchestratorruntime.BgpPeersClientGetResponse{
	// 	BgpPeer: &armcontainerorchestratorruntime.BgpPeer{
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1/providers/Microsoft.KubernetesRuntime/bgpPeers/testpeer"),
	// 		Name: to.Ptr("testpeer"),
	// 		Type: to.Ptr("Microsoft.KubernetesRuntime/BgpPeers"),
	// 		Properties: &armcontainerorchestratorruntime.BgpPeerProperties{
	// 			ProvisioningState: to.Ptr(armcontainerorchestratorruntime.ProvisioningStateSucceeded),
	// 			MyAsn: to.Ptr[int32](64500),
	// 			PeerAsn: to.Ptr[int32](64501),
	// 			PeerAddress: to.Ptr("10.0.0.1"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-03-01/BgpPeers_List.json
func ExampleBgpPeersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerorchestratorruntime.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBgpPeersClient().NewListPager("subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1", nil)
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
		// page = armcontainerorchestratorruntime.BgpPeersClientListResponse{
		// 	BgpPeerListResult: armcontainerorchestratorruntime.BgpPeerListResult{
		// 		Value: []*armcontainerorchestratorruntime.BgpPeer{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1/providers/Microsoft.KubernetesRuntime/bgpPeers/testpeer"),
		// 				Name: to.Ptr("testpeer"),
		// 				Type: to.Ptr("Microsoft.KubernetesRuntime/BgpPeers"),
		// 				Properties: &armcontainerorchestratorruntime.BgpPeerProperties{
		// 					ProvisioningState: to.Ptr(armcontainerorchestratorruntime.ProvisioningStateSucceeded),
		// 					MyAsn: to.Ptr[int32](64500),
		// 					PeerAsn: to.Ptr[int32](64501),
		// 					PeerAddress: to.Ptr("10.0.0.1"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
