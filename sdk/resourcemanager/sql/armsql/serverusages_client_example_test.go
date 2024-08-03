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

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c7b98b36e4023331545051284d8500adf98f02fe/specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01-legacy/examples/ServerUsageMetricsList.json
func ExampleServerUsagesClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServerUsagesClient().NewListByServerPager("sqlcrudtest-6730", "sqlcrudtest-9007", nil)
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
		// page.ServerUsageListResult = armsql.ServerUsageListResult{
		// 	Value: []*armsql.ServerUsage{
		// 		{
		// 			Name: to.Ptr("server_dtu_quota"),
		// 			CurrentValue: to.Ptr[float64](0),
		// 			DisplayName: to.Ptr("Database Throughput Unit Quota"),
		// 			Limit: to.Ptr[float64](45000),
		// 			ResourceName: to.Ptr("sqlcrudtest-9007"),
		// 			Unit: to.Ptr("DTUs"),
		// 		},
		// 		{
		// 			Name: to.Ptr("server_dtu_quota_current"),
		// 			CurrentValue: to.Ptr[float64](0),
		// 			DisplayName: to.Ptr("Database Throughput Unit Quota"),
		// 			Limit: to.Ptr[float64](45000),
		// 			ResourceName: to.Ptr("sqlcrudtest-9007"),
		// 			Unit: to.Ptr("DTUs"),
		// 	}},
		// }
	}
}
