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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/StartStopManagedInstanceScheduleList.json
func ExampleStartStopManagedInstanceSchedulesClient_NewListByInstancePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStartStopManagedInstanceSchedulesClient().NewListByInstancePager("schedulerg", "schedulemi", nil)
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
		// page.StartStopManagedInstanceScheduleListResult = armsql.StartStopManagedInstanceScheduleListResult{
		// 	Value: []*armsql.StartStopManagedInstanceSchedule{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/startStopSchedules"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/schedulerg/providers/Microsoft.Sql/managedInstances/schedulemi/startStopSchedules"),
		// 			Properties: &armsql.StartStopManagedInstanceScheduleProperties{
		// 				Description: to.Ptr("This is a schedule for our Dev/Test environment."),
		// 				NextExecutionTime: to.Ptr("2021-08-26T14:00:00"),
		// 				NextRunAction: to.Ptr("Stop"),
		// 				ScheduleList: []*armsql.ScheduleItem{
		// 					{
		// 						StartDay: to.Ptr(armsql.DayOfWeekThursday),
		// 						StartTime: to.Ptr("06:00 PM"),
		// 						StopDay: to.Ptr(armsql.DayOfWeekThursday),
		// 						StopTime: to.Ptr("05:00 PM"),
		// 					},
		// 					{
		// 						StartDay: to.Ptr(armsql.DayOfWeekThursday),
		// 						StartTime: to.Ptr("03:00 PM"),
		// 						StopDay: to.Ptr(armsql.DayOfWeekThursday),
		// 						StopTime: to.Ptr("02:00 PM"),
		// 				}},
		// 				TimeZoneID: to.Ptr("Central European Standard Time"),
		// 			},
		// 			SystemData: &armsql.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-26T04:41:33.937Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armsql.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-27T04:41:33.937Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armsql.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/StartStopManagedInstanceScheduleGet.json
func ExampleStartStopManagedInstanceSchedulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStartStopManagedInstanceSchedulesClient().Get(ctx, "schedulerg", "schedulemi", armsql.StartStopScheduleNameDefault, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StartStopManagedInstanceSchedule = armsql.StartStopManagedInstanceSchedule{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/startStopSchedules"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/schedulerg/providers/Microsoft.Sql/managedInstances/schedulemi/startStopSchedules/default"),
	// 	Properties: &armsql.StartStopManagedInstanceScheduleProperties{
	// 		Description: to.Ptr("This is a schedule for our Dev/Test environment."),
	// 		NextExecutionTime: to.Ptr("2021-08-26T14:00:00"),
	// 		NextRunAction: to.Ptr("Stop"),
	// 		ScheduleList: []*armsql.ScheduleItem{
	// 			{
	// 				StartDay: to.Ptr(armsql.DayOfWeekThursday),
	// 				StartTime: to.Ptr("06:00 PM"),
	// 				StopDay: to.Ptr(armsql.DayOfWeekThursday),
	// 				StopTime: to.Ptr("05:00 PM"),
	// 			},
	// 			{
	// 				StartDay: to.Ptr(armsql.DayOfWeekThursday),
	// 				StartTime: to.Ptr("03:00 PM"),
	// 				StopDay: to.Ptr(armsql.DayOfWeekThursday),
	// 				StopTime: to.Ptr("02:00 PM"),
	// 		}},
	// 		TimeZoneID: to.Ptr("Central European Standard Time"),
	// 	},
	// 	SystemData: &armsql.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-26T04:41:33.937Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armsql.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-27T04:41:33.937Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armsql.CreatedByTypeUser),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/StartStopManagedInstanceScheduleCreateOrUpdateMax.json
func ExampleStartStopManagedInstanceSchedulesClient_CreateOrUpdate_createsOrUpdatesTheManagedInstancesStartStopScheduleWithAllOptionalParametersSpecified() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStartStopManagedInstanceSchedulesClient().CreateOrUpdate(ctx, "schedulerg", "schedulemi", armsql.StartStopScheduleNameDefault, armsql.StartStopManagedInstanceSchedule{
		Properties: &armsql.StartStopManagedInstanceScheduleProperties{
			Description: to.Ptr("This is a schedule for our Dev/Test environment."),
			ScheduleList: []*armsql.ScheduleItem{
				{
					StartDay:  to.Ptr(armsql.DayOfWeekThursday),
					StartTime: to.Ptr("18:00"),
					StopDay:   to.Ptr(armsql.DayOfWeekThursday),
					StopTime:  to.Ptr("17:00"),
				},
				{
					StartDay:  to.Ptr(armsql.DayOfWeekThursday),
					StartTime: to.Ptr("15:00"),
					StopDay:   to.Ptr(armsql.DayOfWeekThursday),
					StopTime:  to.Ptr("14:00"),
				}},
			TimeZoneID: to.Ptr("Central European Standard Time"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StartStopManagedInstanceSchedule = armsql.StartStopManagedInstanceSchedule{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/startStopSchedules"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/schedulerg/providers/Microsoft.Sql/managedInstances/schedulemi/startStopSchedules/default"),
	// 	Properties: &armsql.StartStopManagedInstanceScheduleProperties{
	// 		Description: to.Ptr("This is a schedule for our Dev/Test environment."),
	// 		NextExecutionTime: to.Ptr("2021-08-26T14:00:00"),
	// 		NextRunAction: to.Ptr("Stop"),
	// 		ScheduleList: []*armsql.ScheduleItem{
	// 			{
	// 				StartDay: to.Ptr(armsql.DayOfWeekThursday),
	// 				StartTime: to.Ptr("06:00 PM"),
	// 				StopDay: to.Ptr(armsql.DayOfWeekThursday),
	// 				StopTime: to.Ptr("05:00 PM"),
	// 			},
	// 			{
	// 				StartDay: to.Ptr(armsql.DayOfWeekThursday),
	// 				StartTime: to.Ptr("03:00 PM"),
	// 				StopDay: to.Ptr(armsql.DayOfWeekThursday),
	// 				StopTime: to.Ptr("02:00 PM"),
	// 		}},
	// 		TimeZoneID: to.Ptr("Central European Standard Time"),
	// 	},
	// 	SystemData: &armsql.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-26T04:41:33.937Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armsql.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-27T04:41:33.937Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armsql.CreatedByTypeUser),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/StartStopManagedInstanceScheduleCreateOrUpdateMin.json
func ExampleStartStopManagedInstanceSchedulesClient_CreateOrUpdate_createsOrUpdatesTheManagedInstancesStartStopScheduleWithNoOptionalParametersSpecified() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStartStopManagedInstanceSchedulesClient().CreateOrUpdate(ctx, "schedulerg", "schedulemi", armsql.StartStopScheduleNameDefault, armsql.StartStopManagedInstanceSchedule{
		Properties: &armsql.StartStopManagedInstanceScheduleProperties{
			ScheduleList: []*armsql.ScheduleItem{
				{
					StartDay:  to.Ptr(armsql.DayOfWeekThursday),
					StartTime: to.Ptr("18:00"),
					StopDay:   to.Ptr(armsql.DayOfWeekThursday),
					StopTime:  to.Ptr("17:00"),
				},
				{
					StartDay:  to.Ptr(armsql.DayOfWeekThursday),
					StartTime: to.Ptr("15:00"),
					StopDay:   to.Ptr(armsql.DayOfWeekThursday),
					StopTime:  to.Ptr("14:00"),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StartStopManagedInstanceSchedule = armsql.StartStopManagedInstanceSchedule{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/startStopSchedules"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/schedulerg/providers/Microsoft.Sql/managedInstances/schedulemi/startStopSchedules/default"),
	// 	Properties: &armsql.StartStopManagedInstanceScheduleProperties{
	// 		Description: to.Ptr(""),
	// 		NextExecutionTime: to.Ptr("2021-08-26T14:00:00"),
	// 		NextRunAction: to.Ptr("Stop"),
	// 		ScheduleList: []*armsql.ScheduleItem{
	// 			{
	// 				StartDay: to.Ptr(armsql.DayOfWeekThursday),
	// 				StartTime: to.Ptr("06:00 PM"),
	// 				StopDay: to.Ptr(armsql.DayOfWeekThursday),
	// 				StopTime: to.Ptr("05:00 PM"),
	// 			},
	// 			{
	// 				StartDay: to.Ptr(armsql.DayOfWeekThursday),
	// 				StartTime: to.Ptr("03:00 PM"),
	// 				StopDay: to.Ptr(armsql.DayOfWeekThursday),
	// 				StopTime: to.Ptr("02:00 PM"),
	// 		}},
	// 		TimeZoneID: to.Ptr("UTC"),
	// 	},
	// 	SystemData: &armsql.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-26T04:41:33.937Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armsql.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-27T04:41:33.937Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armsql.CreatedByTypeUser),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/StartStopManagedInstanceScheduleDelete.json
func ExampleStartStopManagedInstanceSchedulesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewStartStopManagedInstanceSchedulesClient().Delete(ctx, "schedulerg", "schedulemi", armsql.StartStopScheduleNameDefault, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
