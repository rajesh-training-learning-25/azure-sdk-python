//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmobilenetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/eba34b9c764e877193788a87a81cebfa915eb858/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-06-01/examples/SimDelete.json
func ExampleSimsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmobilenetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSimsClient().BeginDelete(ctx, "testResourceGroupName", "testSimGroup", "testSim", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/eba34b9c764e877193788a87a81cebfa915eb858/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-06-01/examples/SimGet.json
func ExampleSimsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmobilenetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSimsClient().Get(ctx, "testResourceGroupName", "testSimGroup", "testSimName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Sim = armmobilenetwork.Sim{
	// 	Name: to.Ptr("testSim"),
	// 	Type: to.Ptr("Microsoft.MobileNetwork/simGroups/sims"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/simGroups/testSimGroup/sims/testSim"),
	// 	SystemData: &armmobilenetwork.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
	// 	},
	// 	Properties: &armmobilenetwork.SimPropertiesFormat{
	// 		DeviceType: to.Ptr("Video camera"),
	// 		IntegratedCircuitCardIdentifier: to.Ptr("8900000000000000000"),
	// 		InternationalMobileSubscriberIdentity: to.Ptr("00000"),
	// 		ProvisioningState: to.Ptr(armmobilenetwork.ProvisioningStateSucceeded),
	// 		SimPolicy: &armmobilenetwork.SimPolicyResourceID{
	// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/simPolicies/MySimPolicy"),
	// 		},
	// 		SimState: to.Ptr(armmobilenetwork.SimStateEnabled),
	// 		SiteProvisioningState: map[string]*armmobilenetwork.SiteProvisioningState{
	// 			"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/sites/testSite": to.Ptr(armmobilenetwork.SiteProvisioningStateProvisioned),
	// 			"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/sites/testSite2": to.Ptr(armmobilenetwork.SiteProvisioningStateProvisioned),
	// 		},
	// 		StaticIPConfiguration: []*armmobilenetwork.SimStaticIPProperties{
	// 			{
	// 				AttachedDataNetwork: &armmobilenetwork.AttachedDataNetworkResourceID{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP/attachedDataNetworks/TestAttachedDataNetwork"),
	// 				},
	// 				Slice: &armmobilenetwork.SliceResourceID{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice"),
	// 				},
	// 				StaticIP: &armmobilenetwork.SimStaticIPPropertiesStaticIP{
	// 					IPv4Address: to.Ptr("2.4.0.1"),
	// 				},
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/eba34b9c764e877193788a87a81cebfa915eb858/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-06-01/examples/SimCreate.json
func ExampleSimsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmobilenetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSimsClient().BeginCreateOrUpdate(ctx, "rg1", "testSimGroup", "testSim", armmobilenetwork.Sim{
		Properties: &armmobilenetwork.SimPropertiesFormat{
			DeviceType:                            to.Ptr("Video camera"),
			IntegratedCircuitCardIdentifier:       to.Ptr("8900000000000000000"),
			InternationalMobileSubscriberIdentity: to.Ptr("00000"),
			SimPolicy: &armmobilenetwork.SimPolicyResourceID{
				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/simPolicies/MySimPolicy"),
			},
			StaticIPConfiguration: []*armmobilenetwork.SimStaticIPProperties{
				{
					AttachedDataNetwork: &armmobilenetwork.AttachedDataNetworkResourceID{
						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP/attachedDataNetworks/TestAttachedDataNetwork"),
					},
					Slice: &armmobilenetwork.SliceResourceID{
						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice"),
					},
					StaticIP: &armmobilenetwork.SimStaticIPPropertiesStaticIP{
						IPv4Address: to.Ptr("2.4.0.1"),
					},
				}},
			AuthenticationKey: to.Ptr("00000000000000000000000000000000"),
			OperatorKeyCode:   to.Ptr("00000000000000000000000000000000"),
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
	// res.Sim = armmobilenetwork.Sim{
	// 	Name: to.Ptr("testSim"),
	// 	Type: to.Ptr("Microsoft.MobileNetwork/simGroups/sims"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/simGroups/testSimGroup/sims/testSim"),
	// 	SystemData: &armmobilenetwork.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
	// 	},
	// 	Properties: &armmobilenetwork.SimPropertiesFormat{
	// 		DeviceType: to.Ptr("Video camera"),
	// 		IntegratedCircuitCardIdentifier: to.Ptr("8900000000000000000"),
	// 		InternationalMobileSubscriberIdentity: to.Ptr("00000"),
	// 		ProvisioningState: to.Ptr(armmobilenetwork.ProvisioningStateSucceeded),
	// 		SimPolicy: &armmobilenetwork.SimPolicyResourceID{
	// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/simPolicies/MySimPolicy"),
	// 		},
	// 		SimState: to.Ptr(armmobilenetwork.SimStateEnabled),
	// 		SiteProvisioningState: map[string]*armmobilenetwork.SiteProvisioningState{
	// 			"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/sites/testSite": to.Ptr(armmobilenetwork.SiteProvisioningStateProvisioned),
	// 			"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/sites/testSite2": to.Ptr(armmobilenetwork.SiteProvisioningStateProvisioned),
	// 		},
	// 		StaticIPConfiguration: []*armmobilenetwork.SimStaticIPProperties{
	// 			{
	// 				AttachedDataNetwork: &armmobilenetwork.AttachedDataNetworkResourceID{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP/attachedDataNetworks/TestAttachedDataNetwork"),
	// 				},
	// 				Slice: &armmobilenetwork.SliceResourceID{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice"),
	// 				},
	// 				StaticIP: &armmobilenetwork.SimStaticIPPropertiesStaticIP{
	// 					IPv4Address: to.Ptr("2.4.0.1"),
	// 				},
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/eba34b9c764e877193788a87a81cebfa915eb858/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-06-01/examples/SimListBySimGroup.json
func ExampleSimsClient_NewListByGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmobilenetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSimsClient().NewListByGroupPager("rg1", "testSimGroup", nil)
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
		// page.SimListResult = armmobilenetwork.SimListResult{
		// 	Value: []*armmobilenetwork.Sim{
		// 		{
		// 			Name: to.Ptr("testSim"),
		// 			Type: to.Ptr("Microsoft.MobileNetwork/simGroups/sims"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/simGroups/testSimGroup/sims/testSim"),
		// 			SystemData: &armmobilenetwork.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
		// 			},
		// 			Properties: &armmobilenetwork.SimPropertiesFormat{
		// 				DeviceType: to.Ptr("Video camera"),
		// 				IntegratedCircuitCardIdentifier: to.Ptr("8900000000000000000"),
		// 				InternationalMobileSubscriberIdentity: to.Ptr("00000"),
		// 				ProvisioningState: to.Ptr(armmobilenetwork.ProvisioningStateSucceeded),
		// 				SimPolicy: &armmobilenetwork.SimPolicyResourceID{
		// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/simPolicies/MySimPolicy"),
		// 				},
		// 				SimState: to.Ptr(armmobilenetwork.SimStateEnabled),
		// 				SiteProvisioningState: map[string]*armmobilenetwork.SiteProvisioningState{
		// 					"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/sites/testSite": to.Ptr(armmobilenetwork.SiteProvisioningStateProvisioned),
		// 					"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/sites/testSite2": to.Ptr(armmobilenetwork.SiteProvisioningStateProvisioned),
		// 				},
		// 				StaticIPConfiguration: []*armmobilenetwork.SimStaticIPProperties{
		// 					{
		// 						AttachedDataNetwork: &armmobilenetwork.AttachedDataNetworkResourceID{
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP/attachedDataNetworks/TestAttachedDataNetwork"),
		// 						},
		// 						Slice: &armmobilenetwork.SliceResourceID{
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice"),
		// 						},
		// 						StaticIP: &armmobilenetwork.SimStaticIPPropertiesStaticIP{
		// 							IPv4Address: to.Ptr("2.4.0.1"),
		// 						},
		// 				}},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/eba34b9c764e877193788a87a81cebfa915eb858/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-06-01/examples/SimBulkUpload.json
func ExampleSimsClient_BeginBulkUpload() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmobilenetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSimsClient().BeginBulkUpload(ctx, "rg1", "testSimGroup", armmobilenetwork.SimUploadList{
		Sims: []*armmobilenetwork.SimNameAndProperties{
			{
				Name: to.Ptr("testSim"),
				Properties: &armmobilenetwork.SimPropertiesFormat{
					DeviceType:                            to.Ptr("Video camera"),
					IntegratedCircuitCardIdentifier:       to.Ptr("8900000000000000000"),
					InternationalMobileSubscriberIdentity: to.Ptr("00000"),
					SimPolicy: &armmobilenetwork.SimPolicyResourceID{
						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/simPolicies/MySimPolicy"),
					},
					StaticIPConfiguration: []*armmobilenetwork.SimStaticIPProperties{
						{
							AttachedDataNetwork: &armmobilenetwork.AttachedDataNetworkResourceID{
								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP/attachedDataNetworks/TestAttachedDataNetwork"),
							},
							Slice: &armmobilenetwork.SliceResourceID{
								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice"),
							},
							StaticIP: &armmobilenetwork.SimStaticIPPropertiesStaticIP{
								IPv4Address: to.Ptr("2.4.0.1"),
							},
						}},
					AuthenticationKey: to.Ptr("00000000000000000000000000000000"),
					OperatorKeyCode:   to.Ptr("00000000000000000000000000000000"),
				},
			},
			{
				Name: to.Ptr("testSim2"),
				Properties: &armmobilenetwork.SimPropertiesFormat{
					DeviceType:                            to.Ptr("Video camera"),
					IntegratedCircuitCardIdentifier:       to.Ptr("8900000000000000001"),
					InternationalMobileSubscriberIdentity: to.Ptr("00000"),
					SimPolicy: &armmobilenetwork.SimPolicyResourceID{
						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/simPolicies/MySimPolicy"),
					},
					StaticIPConfiguration: []*armmobilenetwork.SimStaticIPProperties{
						{
							AttachedDataNetwork: &armmobilenetwork.AttachedDataNetworkResourceID{
								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP/attachedDataNetworks/TestAttachedDataNetwork"),
							},
							Slice: &armmobilenetwork.SliceResourceID{
								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice"),
							},
							StaticIP: &armmobilenetwork.SimStaticIPPropertiesStaticIP{
								IPv4Address: to.Ptr("2.4.0.2"),
							},
						}},
					AuthenticationKey: to.Ptr("00000000000000000000000000000000"),
					OperatorKeyCode:   to.Ptr("00000000000000000000000000000000"),
				},
			}},
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
	// res.AsyncOperationStatus = armmobilenetwork.AsyncOperationStatus{
	// 	Name: to.Ptr("testOperation"),
	// 	EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T03:38:07.000Z"); return t}()),
	// 	ID: to.Ptr("/providers/Microsoft.MobileNetwork/locations/testLocation/operationStatuses/testOperation"),
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T03:36:07.000Z"); return t}()),
	// 	Status: to.Ptr("Succeeded"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/eba34b9c764e877193788a87a81cebfa915eb858/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-06-01/examples/SimBulkDelete.json
func ExampleSimsClient_BeginBulkDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmobilenetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSimsClient().BeginBulkDelete(ctx, "testResourceGroupName", "testSimGroup", armmobilenetwork.SimDeleteList{
		Sims: []*string{
			to.Ptr("testSim"),
			to.Ptr("testSim2")},
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
	// res.AsyncOperationStatus = armmobilenetwork.AsyncOperationStatus{
	// 	Name: to.Ptr("testOperation"),
	// 	EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T03:38:07.000Z"); return t}()),
	// 	ID: to.Ptr("/providers/Microsoft.MobileNetwork/locations/testLocation/operationStatuses/testOperation"),
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T03:36:07.000Z"); return t}()),
	// 	Status: to.Ptr("Succeeded"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/eba34b9c764e877193788a87a81cebfa915eb858/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-06-01/examples/SimBulkUploadEncrypted.json
func ExampleSimsClient_BeginBulkUploadEncrypted() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmobilenetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSimsClient().BeginBulkUploadEncrypted(ctx, "rg1", "testSimGroup", armmobilenetwork.EncryptedSimUploadList{
		AzureKeyIdentifier:    to.Ptr[int32](1),
		EncryptedTransportKey: to.Ptr("ABC123"),
		SignedTransportKey:    to.Ptr("ABC123"),
		Sims: []*armmobilenetwork.SimNameAndEncryptedProperties{
			{
				Name: to.Ptr("testSim"),
				Properties: &armmobilenetwork.EncryptedSimPropertiesFormat{
					DeviceType:                            to.Ptr("Video camera"),
					IntegratedCircuitCardIdentifier:       to.Ptr("8900000000000000000"),
					InternationalMobileSubscriberIdentity: to.Ptr("00000"),
					SimPolicy: &armmobilenetwork.SimPolicyResourceID{
						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/simPolicies/MySimPolicy"),
					},
					StaticIPConfiguration: []*armmobilenetwork.SimStaticIPProperties{
						{
							AttachedDataNetwork: &armmobilenetwork.AttachedDataNetworkResourceID{
								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP/attachedDataNetworks/TestAttachedDataNetwork"),
							},
							Slice: &armmobilenetwork.SliceResourceID{
								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice"),
							},
							StaticIP: &armmobilenetwork.SimStaticIPPropertiesStaticIP{
								IPv4Address: to.Ptr("2.4.0.1"),
							},
						}},
					EncryptedCredentials: to.Ptr("ABC123"),
				},
			},
			{
				Name: to.Ptr("testSim2"),
				Properties: &armmobilenetwork.EncryptedSimPropertiesFormat{
					DeviceType:                            to.Ptr("Video camera"),
					IntegratedCircuitCardIdentifier:       to.Ptr("8900000000000000001"),
					InternationalMobileSubscriberIdentity: to.Ptr("00000"),
					SimPolicy: &armmobilenetwork.SimPolicyResourceID{
						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/simPolicies/MySimPolicy"),
					},
					StaticIPConfiguration: []*armmobilenetwork.SimStaticIPProperties{
						{
							AttachedDataNetwork: &armmobilenetwork.AttachedDataNetworkResourceID{
								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP/attachedDataNetworks/TestAttachedDataNetwork"),
							},
							Slice: &armmobilenetwork.SliceResourceID{
								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice"),
							},
							StaticIP: &armmobilenetwork.SimStaticIPPropertiesStaticIP{
								IPv4Address: to.Ptr("2.4.0.2"),
							},
						}},
					EncryptedCredentials: to.Ptr("ABC123"),
				},
			}},
		VendorKeyFingerprint: to.Ptr("ABC123"),
		Version:              to.Ptr[int32](1),
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
	// res.AsyncOperationStatus = armmobilenetwork.AsyncOperationStatus{
	// 	Name: to.Ptr("testOperation"),
	// 	EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T03:38:07.000Z"); return t}()),
	// 	ID: to.Ptr("/providers/Microsoft.MobileNetwork/locations/testLocation/operationStatuses/testOperation"),
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T03:36:07.000Z"); return t}()),
	// 	Status: to.Ptr("Succeeded"),
	// }
}
