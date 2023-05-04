//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e60df62e9e0d88462e6abba81a76d94eab000f0d/specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2023-05-01/examples/ContainerGroupUsage.json
func ExampleLocationClient_NewListUsagePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerinstance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLocationClient().NewListUsagePager("westcentralus", nil)
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
		// page.UsageListResult = armcontainerinstance.UsageListResult{
		// 	Value: []*armcontainerinstance.Usage{
		// 		{
		// 			Name: &armcontainerinstance.UsageName{
		// 				LocalizedValue: to.Ptr("Container Groups"),
		// 				Value: to.Ptr("ContainerGroups"),
		// 			},
		// 			CurrentValue: to.Ptr[int32](1),
		// 			Limit: to.Ptr[int32](2000),
		// 			Unit: to.Ptr("Count"),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e60df62e9e0d88462e6abba81a76d94eab000f0d/specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2023-05-01/examples/CachedImagesList.json
func ExampleLocationClient_NewListCachedImagesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerinstance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLocationClient().NewListCachedImagesPager("westcentralus", nil)
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
		// page.CachedImagesListResult = armcontainerinstance.CachedImagesListResult{
		// 	Value: []*armcontainerinstance.CachedImages{
		// 		{
		// 			Image: to.Ptr("ubuntu:16.04"),
		// 			OSType: to.Ptr("Linux"),
		// 		},
		// 		{
		// 			Image: to.Ptr("alpine:3.6"),
		// 			OSType: to.Ptr("Linux"),
		// 		},
		// 		{
		// 			Image: to.Ptr("microsoft/nanoserver:10.0.14393.2485"),
		// 			OSType: to.Ptr("Windows"),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e60df62e9e0d88462e6abba81a76d94eab000f0d/specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2023-05-01/examples/CapabilitiesList.json
func ExampleLocationClient_NewListCapabilitiesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerinstance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLocationClient().NewListCapabilitiesPager("westus", nil)
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
		// page.CapabilitiesListResult = armcontainerinstance.CapabilitiesListResult{
		// 	Value: []*armcontainerinstance.Capabilities{
		// 		{
		// 			Capabilities: &armcontainerinstance.CapabilitiesCapabilities{
		// 				MaxCPU: to.Ptr[float32](4),
		// 				MaxGpuCount: to.Ptr[float32](4),
		// 				MaxMemoryInGB: to.Ptr[float32](14),
		// 			},
		// 			Gpu: to.Ptr("K80"),
		// 			IPAddressType: to.Ptr("Public"),
		// 			Location: to.Ptr("West US"),
		// 			OSType: to.Ptr("Linux"),
		// 			ResourceType: to.Ptr("containerGroups"),
		// 		},
		// 		{
		// 			Capabilities: &armcontainerinstance.CapabilitiesCapabilities{
		// 				MaxCPU: to.Ptr[float32](4),
		// 				MaxGpuCount: to.Ptr[float32](0),
		// 				MaxMemoryInGB: to.Ptr[float32](14),
		// 			},
		// 			Gpu: to.Ptr("None"),
		// 			IPAddressType: to.Ptr("Public"),
		// 			Location: to.Ptr("West US"),
		// 			OSType: to.Ptr("Windows"),
		// 			ResourceType: to.Ptr("containerGroups"),
		// 	}},
		// }
	}
}
