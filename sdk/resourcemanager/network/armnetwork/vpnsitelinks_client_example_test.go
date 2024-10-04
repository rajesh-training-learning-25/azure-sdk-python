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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4883fa5dbf6f2c9093fac8ce334547e9dfac68fa/specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/VpnSiteLinkGet.json
func ExampleVPNSiteLinksClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVPNSiteLinksClient().Get(ctx, "rg1", "vpnSite1", "vpnSiteLink1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VPNSiteLink = armnetwork.VPNSiteLink{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/vpnSite1/vpnSiteLinks/vpnSiteLink1"),
	// 	Name: to.Ptr("vpnSiteLink1"),
	// 	Type: to.Ptr("Microsoft.Network/vpnSites/vpnSiteLinks"),
	// 	Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
	// 	Properties: &armnetwork.VPNSiteLinkProperties{
	// 		BgpProperties: &armnetwork.VPNLinkBgpSettings{
	// 			Asn: to.Ptr[int64](1234),
	// 			BgpPeeringAddress: to.Ptr("192.168.0.0"),
	// 		},
	// 		IPAddress: to.Ptr("50.50.50.56"),
	// 		LinkProperties: &armnetwork.VPNLinkProviderProperties{
	// 			LinkSpeedInMbps: to.Ptr[int32](0),
	// 		},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4883fa5dbf6f2c9093fac8ce334547e9dfac68fa/specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/VpnSiteLinkListByVpnSite.json
func ExampleVPNSiteLinksClient_NewListByVPNSitePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVPNSiteLinksClient().NewListByVPNSitePager("rg1", "vpnSite1", nil)
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
		// page.ListVPNSiteLinksResult = armnetwork.ListVPNSiteLinksResult{
		// 	Value: []*armnetwork.VPNSiteLink{
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/vpnSite1/vpnSiteLinks/vpnSiteLink1"),
		// 			Name: to.Ptr("vpnSiteLink1"),
		// 			Type: to.Ptr("Microsoft.Network/vpnSites/vpnSiteLinks"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
		// 			Properties: &armnetwork.VPNSiteLinkProperties{
		// 				BgpProperties: &armnetwork.VPNLinkBgpSettings{
		// 					Asn: to.Ptr[int64](1234),
		// 					BgpPeeringAddress: to.Ptr("192.168.0.0"),
		// 				},
		// 				IPAddress: to.Ptr("50.50.50.56"),
		// 				LinkProperties: &armnetwork.VPNLinkProviderProperties{
		// 					LinkSpeedInMbps: to.Ptr[int32](200),
		// 				},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/vpnSite1/vpnSiteLinks/vpnSiteLink2"),
		// 			Name: to.Ptr("vpnSiteLink2"),
		// 			Type: to.Ptr("Microsoft.Network/vpnSites/vpnSiteLinks"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
		// 			Properties: &armnetwork.VPNSiteLinkProperties{
		// 				BgpProperties: &armnetwork.VPNLinkBgpSettings{
		// 					Asn: to.Ptr[int64](1234),
		// 					BgpPeeringAddress: to.Ptr("192.168.0.1"),
		// 				},
		// 				IPAddress: to.Ptr("40.40.40.46"),
		// 				LinkProperties: &armnetwork.VPNLinkProviderProperties{
		// 					LinkSpeedInMbps: to.Ptr[int32](200),
		// 				},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
