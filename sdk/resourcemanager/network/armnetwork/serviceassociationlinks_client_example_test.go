//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/network/resource-manager/Microsoft.Network/stable/2022-09-01/examples/VirtualNetworkGetServiceAssociationLinks.json
func ExampleServiceAssociationLinksClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewServiceAssociationLinksClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.List(ctx, "rg1", "vnet", "subnet", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServiceAssociationLinksListResult = armnetwork.ServiceAssociationLinksListResult{
	// 	Value: []*armnetwork.ServiceAssociationLink{
	// 		{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet/subnets/subnet/serviceAssociationLinks/acisal"),
	// 			Name: to.Ptr("acisal"),
	// 			Type: to.Ptr("Microsoft.Network/virtualNetworks/subnets/serviceAssociationLinks"),
	// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
	// 			Properties: &armnetwork.ServiceAssociationLinkPropertiesFormat{
	// 				AllowDelete: to.Ptr(true),
	// 				LinkedResourceType: to.Ptr("Microsoft.ContainerInstance/containerGroups"),
	// 				Locations: []*string{
	// 					to.Ptr("westus")},
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 				},
	// 		}},
	// 	}
}
