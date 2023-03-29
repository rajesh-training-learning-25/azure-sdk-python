//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmysql_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/TopQueryStatisticsGet.json
func ExampleTopQueryStatisticsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTopQueryStatisticsClient().Get(ctx, "testResourceGroupName", "testServerName", "66-636923268000000000-636923277000000000-avg-duration", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.QueryStatistic = armmysql.QueryStatistic{
	// 	Name: to.Ptr("66-636923268000000000-636923277000000000-avg-duration"),
	// 	Type: to.Ptr("Microsoft.DBforMySQL/servers/queryStatistics"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testResourceGroupName/providers/Microsoft.DBforMySQL/servers/testServerName/queryStatistic/66-636923268000000000-636923277000000000-avg-duration"),
	// 	Properties: &armmysql.QueryStatisticProperties{
	// 		AggregationFunction: to.Ptr("avg"),
	// 		DatabaseNames: []*string{
	// 			to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testResourceGroupName/providers/Microsoft.DBforMySQL/servers/testServerName/databases/mysql")},
	// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-01T17:15:00Z"); return t}()),
	// 			MetricDisplayName: to.Ptr("Query duration"),
	// 			MetricName: to.Ptr("duration"),
	// 			MetricValue: to.Ptr[float64](123.301446136),
	// 			MetricValueUnit: to.Ptr("milliseconds"),
	// 			QueryExecutionCount: to.Ptr[int64](1),
	// 			QueryID: to.Ptr("66"),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-01T17:00:00Z"); return t}()),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/TopQueryStatisticsListByServer.json
func ExampleTopQueryStatisticsClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTopQueryStatisticsClient().NewListByServerPager("testResourceGroupName", "testServerName", armmysql.TopQueryStatisticsInput{
		Properties: &armmysql.TopQueryStatisticsInputProperties{
			AggregationFunction:  to.Ptr("avg"),
			AggregationWindow:    to.Ptr("PT15M"),
			NumberOfTopQueries:   to.Ptr[int32](5),
			ObservationEndTime:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-07T20:00:00.000Z"); return t }()),
			ObservationStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-01T20:00:00.000Z"); return t }()),
			ObservedMetric:       to.Ptr("duration"),
		},
	}, nil)
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
		// page.TopQueryStatisticsResultList = armmysql.TopQueryStatisticsResultList{
		// 	Value: []*armmysql.QueryStatistic{
		// 		{
		// 			Name: to.Ptr("66-636923268000000000-636923277000000000-avg-duration"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/queryStatistics"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testResourceGroupName/providers/Microsoft.DBforMySQL/servers/testServerName/queryStatistic/66-636923268000000000-636923277000000000-avg-duration"),
		// 			Properties: &armmysql.QueryStatisticProperties{
		// 				AggregationFunction: to.Ptr("avg"),
		// 				DatabaseNames: []*string{
		// 					to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testResourceGroupName/providers/Microsoft.DBforMySQL/servers/testServerName/databases/mysql")},
		// 					EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-01T17:15:00Z"); return t}()),
		// 					MetricDisplayName: to.Ptr("Query duration"),
		// 					MetricName: to.Ptr("duration"),
		// 					MetricValue: to.Ptr[float64](123.301446136),
		// 					MetricValueUnit: to.Ptr("milliseconds"),
		// 					QueryExecutionCount: to.Ptr[int64](1),
		// 					QueryID: to.Ptr("66"),
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-01T17:00:00Z"); return t}()),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("66-636924483000000000-636924492000000000-avg-duration"),
		// 				Type: to.Ptr("Microsoft.DBforMySQL/servers/queryStatistics"),
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testResourceGroupName/providers/Microsoft.DBforMySQL/servers/testServerName/queryStatistic/66-636924483000000000-636924492000000000-avg-duration"),
		// 				Properties: &armmysql.QueryStatisticProperties{
		// 					AggregationFunction: to.Ptr("avg"),
		// 					DatabaseNames: []*string{
		// 						to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testResourceGroupName/providers/Microsoft.DBforMySQL/servers/testServerName/databases/mysql")},
		// 						EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-03T03:00:00Z"); return t}()),
		// 						MetricDisplayName: to.Ptr("Query duration"),
		// 						MetricName: to.Ptr("duration"),
		// 						MetricValue: to.Ptr[float64](1712.301446136),
		// 						MetricValueUnit: to.Ptr("milliseconds"),
		// 						QueryExecutionCount: to.Ptr[int64](1),
		// 						QueryID: to.Ptr("66"),
		// 						StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-03T02:45:00Z"); return t}()),
		// 					},
		// 			}},
		// 		}
	}
}
