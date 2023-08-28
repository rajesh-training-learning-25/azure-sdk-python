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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/GetMaintenanceWindows.json
func ExampleMaintenanceWindowsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMaintenanceWindowsClient().Get(ctx, "Default-SQL-SouthEastAsia", "testsvr", "testdb", "current", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MaintenanceWindows = armsql.MaintenanceWindows{
	// 	Name: to.Ptr("current"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/maintenancewindows"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb/maintenancewindows/current"),
	// 	Properties: &armsql.MaintenanceWindowsProperties{
	// 		TimeRanges: []*armsql.MaintenanceWindowTimeRange{
	// 			{
	// 				DayOfWeek: to.Ptr(armsql.DayOfWeekSaturday),
	// 				Duration: to.Ptr("PT60M"),
	// 				StartTime: to.Ptr("00:00:00"),
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CreateOrUpdateMaintenanceWindows.json
func ExampleMaintenanceWindowsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewMaintenanceWindowsClient().CreateOrUpdate(ctx, "Default-SQL-SouthEastAsia", "testsvr", "testdwdb", "current", armsql.MaintenanceWindows{
		Properties: &armsql.MaintenanceWindowsProperties{
			TimeRanges: []*armsql.MaintenanceWindowTimeRange{
				{
					DayOfWeek: to.Ptr(armsql.DayOfWeekSaturday),
					Duration:  to.Ptr("PT60M"),
					StartTime: to.Ptr("00:00:00"),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
