//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-01-preview/examples/BackupGet.json
func ExampleBackupsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBackupsClient().Get(ctx, "TestGroup", "postgresqltestserver", "daily_20210615T160516", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServerBackup = armpostgresqlflexibleservers.ServerBackup{
	// 	Name: to.Ptr("daily_20210615T160516"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/backups"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/postgresqltestserver/backups/daily_20210615T160516"),
	// 	Properties: &armpostgresqlflexibleservers.ServerBackupProperties{
	// 		BackupType: to.Ptr(armpostgresqlflexibleservers.OriginFull),
	// 		CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-15T16:05:19.902522+00:00"); return t}()),
	// 		Source: to.Ptr("Automatic"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-01-preview/examples/BackupListByServer.json
func ExampleBackupsClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBackupsClient().NewListByServerPager("TestGroup", "postgresqltestserver", nil)
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
		// page.ServerBackupListResult = armpostgresqlflexibleservers.ServerBackupListResult{
		// 	Value: []*armpostgresqlflexibleservers.ServerBackup{
		// 		{
		// 			Name: to.Ptr("daily_20210615T160516"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/backups"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/postgresqltestserver/backups/daily_20210615T160516"),
		// 			Properties: &armpostgresqlflexibleservers.ServerBackupProperties{
		// 				BackupType: to.Ptr(armpostgresqlflexibleservers.OriginFull),
		// 				CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-15T16:05:19.902522+00:00"); return t}()),
		// 				Source: to.Ptr("Automatic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("daily_20210616T160520"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/backups"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/postgresqltestserver/backups/daily_20210616T160520"),
		// 			Properties: &armpostgresqlflexibleservers.ServerBackupProperties{
		// 				BackupType: to.Ptr(armpostgresqlflexibleservers.OriginFull),
		// 				CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-16T16:05:23.9243453+00:00"); return t}()),
		// 				Source: to.Ptr("Automatic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("daily_20210617T160525"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/backups"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/postgresqltestserver/backups/daily_20210617T160525"),
		// 			Properties: &armpostgresqlflexibleservers.ServerBackupProperties{
		// 				BackupType: to.Ptr(armpostgresqlflexibleservers.OriginFull),
		// 				CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-17T16:05:28.1247488+00:00"); return t}()),
		// 				Source: to.Ptr("Automatic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("daily_20210618T160529"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/backups"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/postgresqltestserver/backups/daily_20210618T160529"),
		// 			Properties: &armpostgresqlflexibleservers.ServerBackupProperties{
		// 				BackupType: to.Ptr(armpostgresqlflexibleservers.OriginFull),
		// 				CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-18T16:05:32.2736701+00:00"); return t}()),
		// 				Source: to.Ptr("Automatic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("daily_20210619T160533"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/backups"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/postgresqltestserver/backups/daily_20210619T160533"),
		// 			Properties: &armpostgresqlflexibleservers.ServerBackupProperties{
		// 				BackupType: to.Ptr(armpostgresqlflexibleservers.OriginFull),
		// 				CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-19T16:05:36.8603354+00:00"); return t}()),
		// 				Source: to.Ptr("Automatic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("daily_20210620T160538"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/backups"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/postgresqltestserver/backups/daily_20210620T160538"),
		// 			Properties: &armpostgresqlflexibleservers.ServerBackupProperties{
		// 				BackupType: to.Ptr(armpostgresqlflexibleservers.OriginFull),
		// 				CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-20T16:05:41.9200138+00:00"); return t}()),
		// 				Source: to.Ptr("Automatic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("daily_20210621T160543"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/backups"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/postgresqltestserver/backups/daily_20210621T160543"),
		// 			Properties: &armpostgresqlflexibleservers.ServerBackupProperties{
		// 				BackupType: to.Ptr(armpostgresqlflexibleservers.OriginFull),
		// 				CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-21T16:05:48.8528447+00:00"); return t}()),
		// 				Source: to.Ptr("Automatic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("daily_20210622T160803"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/backups"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/postgresqltestserver/backups/daily_20210622T160803"),
		// 			Properties: &armpostgresqlflexibleservers.ServerBackupProperties{
		// 				BackupType: to.Ptr(armpostgresqlflexibleservers.OriginFull),
		// 				CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-22T16:08:06.3121688+00:00"); return t}()),
		// 				Source: to.Ptr("Automatic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("daily_20210622T210807"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/backups"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/postgresqltestserver/backups/daily_20210622T210807"),
		// 			Properties: &armpostgresqlflexibleservers.ServerBackupProperties{
		// 				BackupType: to.Ptr(armpostgresqlflexibleservers.OriginFull),
		// 				CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-22T21:08:10.5057354+00:00"); return t}()),
		// 				Source: to.Ptr("Automatic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("daily_20210623T212413"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/backups"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/postgresqltestserver/backups/daily_20210623T212413"),
		// 			Properties: &armpostgresqlflexibleservers.ServerBackupProperties{
		// 				BackupType: to.Ptr(armpostgresqlflexibleservers.OriginFull),
		// 				CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-23T21:24:16.9401531+00:00"); return t}()),
		// 				Source: to.Ptr("Automatic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("daily_20210624T061328"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/backups"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/postgresqltestserver/backups/daily_20210624T061328"),
		// 			Properties: &armpostgresqlflexibleservers.ServerBackupProperties{
		// 				BackupType: to.Ptr(armpostgresqlflexibleservers.OriginFull),
		// 				CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-24T06:13:31.4962137+00:00"); return t}()),
		// 				Source: to.Ptr("Automatic"),
		// 			},
		// 	}},
		// }
	}
}
