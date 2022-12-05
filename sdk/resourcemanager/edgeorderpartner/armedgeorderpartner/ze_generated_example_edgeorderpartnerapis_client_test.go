//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armedgeorderpartner_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorderpartner/armedgeorderpartner"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorderpartner/resource-manager/Microsoft.EdgeOrderPartner/preview/2020-12-01-preview/examples/ListOperationsPartner.json
func ExampleAPISClient_NewListOperationsPartnerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armedgeorderpartner.NewAPISClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListOperationsPartnerPager(nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorderpartner/resource-manager/Microsoft.EdgeOrderPartner/preview/2020-12-01-preview/examples/ManageInventoryMetadata.json
func ExampleAPISClient_BeginManageInventoryMetadata() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armedgeorderpartner.NewAPISClient("b783ea86-c85c-4175-b76d-3992656af50d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginManageInventoryMetadata(ctx,
		"AzureStackEdge",
		"westus",
		"SerialNumber1",
		armedgeorderpartner.ManageInventoryMetadataRequest{
			ConfigurationOnDevice: &armedgeorderpartner.ConfigurationOnDevice{
				ConfigurationIdentifier: to.Ptr("EdgeP_High"),
			},
			InventoryMetadata: to.Ptr("InventoryMetadata"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorderpartner/resource-manager/Microsoft.EdgeOrderPartner/preview/2020-12-01-preview/examples/ManageLink.json
func ExampleAPISClient_ManageLink() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armedgeorderpartner.NewAPISClient("b783ea86-c85c-4175-b76d-3992656af50d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.ManageLink(ctx,
		"AzureStackEdge",
		"westus",
		"SerialNumber1",
		armedgeorderpartner.ManageLinkRequest{
			ManagementResourceArmID: to.Ptr("/subscriptions/c783ea86-c85c-4175-b76d-3992656af50d/resourceGroups/EdgeTestRG/providers/Microsoft.DataBoxEdge/DataBoxEdgeDevices/TestEdgeDeviceName1"),
			Operation:               to.Ptr(armedgeorderpartner.ManageLinkOperationLink),
			TenantID:                to.Ptr("a783ea86-c85c-4175-b76d-3992656af50d"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorderpartner/resource-manager/Microsoft.EdgeOrderPartner/preview/2020-12-01-preview/examples/SearchInventories.json
func ExampleAPISClient_NewSearchInventoriesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armedgeorderpartner.NewAPISClient("b783ea86-c85c-4175-b76d-3992656af50d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewSearchInventoriesPager(armedgeorderpartner.SearchInventoriesRequest{
		FamilyIdentifier: to.Ptr("AzureStackEdge"),
		SerialNumber:     to.Ptr("SerialNumber1"),
	},
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