//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armwebpubsub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/webpubsub/armwebpubsub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/Usages_List.json
func ExampleUsagesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armwebpubsub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUsagesClient().NewListPager("eastus", nil)
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
		// page.SignalRServiceUsageList = armwebpubsub.SignalRServiceUsageList{
		// 	Value: []*armwebpubsub.SignalRServiceUsage{
		// 		{
		// 			Name: &armwebpubsub.SignalRServiceUsageName{
		// 				LocalizedValue: to.Ptr("Usage1"),
		// 				Value: to.Ptr("Usage1"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.SignalRService/locations/eastus/usages/Usage1"),
		// 			Limit: to.Ptr[int64](100),
		// 			Unit: to.Ptr("Count"),
		// 		},
		// 		{
		// 			Name: &armwebpubsub.SignalRServiceUsageName{
		// 				LocalizedValue: to.Ptr("Usage2"),
		// 				Value: to.Ptr("Usage2"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.SignalRService/locations/eastus/usages/Usage2"),
		// 			Limit: to.Ptr[int64](100),
		// 			Unit: to.Ptr("Count"),
		// 	}},
		// }
	}
}
