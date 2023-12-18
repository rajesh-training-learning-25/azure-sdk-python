//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armstorsimple8000series_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/BackupSchedulesListByBackupPolicy.json
func ExampleBackupSchedulesClient_NewListByBackupPolicyPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBackupSchedulesClient().NewListByBackupPolicyPager("Device05ForSDKTest", "BkUpPolicy01ForSDKTest", "ResourceGroupForSDKTest", "ManagerForSDKTest1", nil)
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
		// page.BackupScheduleList = armstorsimple8000series.BackupScheduleList{
		// 	Value: []*armstorsimple8000series.BackupSchedule{
		// 		{
		// 			Name: to.Ptr("schedule1"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/backupPolicies/schedules"),
		// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/backupPolicies/BkUpPolicy01ForSDKTest/schedules/schedule1"),
		// 			Kind: to.Ptr("Series8000"),
		// 			Properties: &armstorsimple8000series.BackupScheduleProperties{
		// 				BackupType: to.Ptr(armstorsimple8000series.BackupTypeCloudSnapshot),
		// 				RetentionCount: to.Ptr[int64](1),
		// 				ScheduleRecurrence: &armstorsimple8000series.ScheduleRecurrence{
		// 					RecurrenceType: to.Ptr(armstorsimple8000series.RecurrenceTypeDaily),
		// 					RecurrenceValue: to.Ptr[int32](1),
		// 				},
		// 				ScheduleStatus: to.Ptr(armstorsimple8000series.ScheduleStatusEnabled),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-24T00:00:00.000Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("schedule2"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/backupPolicies/schedules"),
		// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/backupPolicies/BkUpPolicy01ForSDKTest/schedules/schedule2"),
		// 			Kind: to.Ptr("Series8000"),
		// 			Properties: &armstorsimple8000series.BackupScheduleProperties{
		// 				BackupType: to.Ptr(armstorsimple8000series.BackupTypeCloudSnapshot),
		// 				RetentionCount: to.Ptr[int64](1),
		// 				ScheduleRecurrence: &armstorsimple8000series.ScheduleRecurrence{
		// 					RecurrenceType: to.Ptr(armstorsimple8000series.RecurrenceTypeWeekly),
		// 					RecurrenceValue: to.Ptr[int32](1),
		// 					WeeklyDaysList: []*armstorsimple8000series.DayOfWeek{
		// 						to.Ptr(armstorsimple8000series.DayOfWeekMonday),
		// 						to.Ptr(armstorsimple8000series.DayOfWeekThursday),
		// 						to.Ptr(armstorsimple8000series.DayOfWeekFriday)},
		// 					},
		// 					ScheduleStatus: to.Ptr(armstorsimple8000series.ScheduleStatusEnabled),
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-24T01:00:00.000Z"); return t}()),
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/BackupSchedulesGet.json
func ExampleBackupSchedulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBackupSchedulesClient().Get(ctx, "Device05ForSDKTest", "BkUpPolicy01ForSDKTest", "schedule2", "ResourceGroupForSDKTest", "ManagerForSDKTest1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BackupSchedule = armstorsimple8000series.BackupSchedule{
	// 	Name: to.Ptr("schedule2"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/devices/backupPolicies/schedules"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/backupPolicies/BkUpPolicy01ForSDKTest/schedules/schedule2"),
	// 	Kind: to.Ptr("Series8000"),
	// 	Properties: &armstorsimple8000series.BackupScheduleProperties{
	// 		BackupType: to.Ptr(armstorsimple8000series.BackupTypeCloudSnapshot),
	// 		RetentionCount: to.Ptr[int64](1),
	// 		ScheduleRecurrence: &armstorsimple8000series.ScheduleRecurrence{
	// 			RecurrenceType: to.Ptr(armstorsimple8000series.RecurrenceTypeWeekly),
	// 			RecurrenceValue: to.Ptr[int32](1),
	// 			WeeklyDaysList: []*armstorsimple8000series.DayOfWeek{
	// 				to.Ptr(armstorsimple8000series.DayOfWeekMonday),
	// 				to.Ptr(armstorsimple8000series.DayOfWeekThursday),
	// 				to.Ptr(armstorsimple8000series.DayOfWeekFriday)},
	// 			},
	// 			ScheduleStatus: to.Ptr(armstorsimple8000series.ScheduleStatusEnabled),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-24T01:00:00.000Z"); return t}()),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/BackupSchedulesCreateOrUpdate.json
func ExampleBackupSchedulesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBackupSchedulesClient().BeginCreateOrUpdate(ctx, "Device05ForSDKTest", "BkUpPolicy01ForSDKTest", "schedule2", "ResourceGroupForSDKTest", "ManagerForSDKTest1", armstorsimple8000series.BackupSchedule{
		Kind: to.Ptr("Series8000"),
		Properties: &armstorsimple8000series.BackupScheduleProperties{
			BackupType:     to.Ptr(armstorsimple8000series.BackupTypeCloudSnapshot),
			RetentionCount: to.Ptr[int64](1),
			ScheduleRecurrence: &armstorsimple8000series.ScheduleRecurrence{
				RecurrenceType:  to.Ptr(armstorsimple8000series.RecurrenceTypeWeekly),
				RecurrenceValue: to.Ptr[int32](1),
				WeeklyDaysList: []*armstorsimple8000series.DayOfWeek{
					to.Ptr(armstorsimple8000series.DayOfWeekFriday),
					to.Ptr(armstorsimple8000series.DayOfWeekThursday),
					to.Ptr(armstorsimple8000series.DayOfWeekMonday)},
			},
			ScheduleStatus: to.Ptr(armstorsimple8000series.ScheduleStatusEnabled),
			StartTime:      to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-24T01:00:00.000Z"); return t }()),
		},
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
	// res.BackupSchedule = armstorsimple8000series.BackupSchedule{
	// 	Name: to.Ptr("schedule2"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/devices/backupPolicies/schedules"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/backupPolicies/BkUpPolicy01ForSDKTest/schedules/schedule2"),
	// 	Kind: to.Ptr("Series8000"),
	// 	Properties: &armstorsimple8000series.BackupScheduleProperties{
	// 		BackupType: to.Ptr(armstorsimple8000series.BackupTypeCloudSnapshot),
	// 		RetentionCount: to.Ptr[int64](1),
	// 		ScheduleRecurrence: &armstorsimple8000series.ScheduleRecurrence{
	// 			RecurrenceType: to.Ptr(armstorsimple8000series.RecurrenceTypeWeekly),
	// 			RecurrenceValue: to.Ptr[int32](1),
	// 			WeeklyDaysList: []*armstorsimple8000series.DayOfWeek{
	// 				to.Ptr(armstorsimple8000series.DayOfWeekMonday),
	// 				to.Ptr(armstorsimple8000series.DayOfWeekThursday),
	// 				to.Ptr(armstorsimple8000series.DayOfWeekFriday)},
	// 			},
	// 			ScheduleStatus: to.Ptr(armstorsimple8000series.ScheduleStatusEnabled),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-24T01:00:00.000Z"); return t}()),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/BackupSchedulesDelete.json
func ExampleBackupSchedulesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBackupSchedulesClient().BeginDelete(ctx, "Device05ForSDKTest", "BkUpPolicy01ForSDKTest", "schedule1", "ResourceGroupForSDKTest", "ManagerForSDKTest1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
