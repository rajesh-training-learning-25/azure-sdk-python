//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ServerDnsAliasGet.json
func ExampleServerDNSAliasesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServerDNSAliasesClient().Get(ctx, "Default", "dns-alias-server", "dns-alias-name-1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServerDNSAlias = armsql.ServerDNSAlias{
	// 	Name: to.Ptr("dns-alias-name-1"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/dnsAliases"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/dns-alias-server/dnsAliases/dns-alias-name-1"),
	// 	Properties: &armsql.ServerDNSAliasProperties{
	// 		AzureDNSRecord: to.Ptr("dns-alias-name-1.database.windows.net"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ServerDnsAliasCreateOrUpdate.json
func ExampleServerDNSAliasesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServerDNSAliasesClient().BeginCreateOrUpdate(ctx, "Default", "dns-alias-server", "dns-alias-name-1", nil)
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
	// res.ServerDNSAlias = armsql.ServerDNSAlias{
	// 	Name: to.Ptr("dns-alias-name-1"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/dnsAliases"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/dns-alias-server/dnsAliases/dns-alias-name-1"),
	// 	Properties: &armsql.ServerDNSAliasProperties{
	// 		AzureDNSRecord: to.Ptr("dns-alias-name-1.database.windows.net"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ServerDnsAliasDelete.json
func ExampleServerDNSAliasesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServerDNSAliasesClient().BeginDelete(ctx, "Default", "dns-alias-server", "dns-alias-name-1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ServerDnsAliasList.json
func ExampleServerDNSAliasesClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServerDNSAliasesClient().NewListByServerPager("Default", "dns-alias-server", nil)
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
		// page.ServerDNSAliasListResult = armsql.ServerDNSAliasListResult{
		// 	Value: []*armsql.ServerDNSAlias{
		// 		{
		// 			Name: to.Ptr("server-dns-alias-1"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/dnsAliases"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/dns-alias-server/dnsAliases/dns-alias-1"),
		// 			Properties: &armsql.ServerDNSAliasProperties{
		// 				AzureDNSRecord: to.Ptr("dns-alias-1.database.windows.net"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("server-dns-alias-2"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/dnsAliases"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/dns-alias-server/dnsAliases/dns-alias-2"),
		// 			Properties: &armsql.ServerDNSAliasProperties{
		// 				AzureDNSRecord: to.Ptr("dns-alias-2.database.windows.net"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ServerDnsAliasAcquire.json
func ExampleServerDNSAliasesClient_BeginAcquire() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServerDNSAliasesClient().BeginAcquire(ctx, "Default", "dns-alias-new-server", "dns-alias-name-1", armsql.ServerDNSAliasAcquisition{
		OldServerDNSAliasID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/dns-alias-old-server/dnsAliases/dns-alias-name-1"),
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
	// res.ServerDNSAlias = armsql.ServerDNSAlias{
	// 	Name: to.Ptr("server-dns-alias-name-1"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/dnsAliases"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/dns-alias-new-server/dnsAliases/dns-alias-name-1"),
	// 	Properties: &armsql.ServerDNSAliasProperties{
	// 		AzureDNSRecord: to.Ptr("dns-alias-name-1.database.windows.net"),
	// 	},
	// }
}
