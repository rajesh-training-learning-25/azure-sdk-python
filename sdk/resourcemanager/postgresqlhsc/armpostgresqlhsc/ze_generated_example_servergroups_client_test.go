//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresqlhsc_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2020-10-05-privatepreview/examples/ServerGroupList.json
func ExampleServerGroupsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpostgresqlhsc.NewServerGroupsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2020-10-05-privatepreview/examples/ServerGroupListByResourceGroup.json
func ExampleServerGroupsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpostgresqlhsc.NewServerGroupsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByResourceGroupPager("<resource-group-name>",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2020-10-05-privatepreview/examples/ServerGroupCreate.json
func ExampleServerGroupsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpostgresqlhsc.NewServerGroupsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<server-group-name>",
		armpostgresqlhsc.ServerGroup{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"ElasticServer": to.Ptr("1"),
			},
			Properties: &armpostgresqlhsc.ServerGroupProperties{
				AdministratorLogin:         to.Ptr("<administrator-login>"),
				AdministratorLoginPassword: to.Ptr("<administrator-login-password>"),
				AvailabilityZone:           to.Ptr("<availability-zone>"),
				BackupRetentionDays:        to.Ptr[int32](35),
				CitusVersion:               to.Ptr(armpostgresqlhsc.CitusVersionNine5),
				DelegatedSubnetArguments: &armpostgresqlhsc.ServerGroupPropertiesDelegatedSubnetArguments{
					SubnetArmResourceID: to.Ptr("<subnet-arm-resource-id>"),
				},
				EnableMx:          to.Ptr(true),
				EnableZfs:         to.Ptr(false),
				PostgresqlVersion: to.Ptr(armpostgresqlhsc.PostgreSQLVersionTwelve),
				PrivateDNSZoneArguments: &armpostgresqlhsc.ServerGroupPropertiesPrivateDNSZoneArguments{
					PrivateDNSZoneArmResourceID: to.Ptr("<private-dnszone-arm-resource-id>"),
				},
				ServerRoleGroups: []*armpostgresqlhsc.ServerRoleGroup{
					{
						EnableHa:         to.Ptr(true),
						ServerEdition:    to.Ptr(armpostgresqlhsc.ServerEditionGeneralPurpose),
						StorageQuotaInMb: to.Ptr[int64](524288),
						VCores:           to.Ptr[int64](4),
						Name:             to.Ptr("<name>"),
						Role:             to.Ptr(armpostgresqlhsc.ServerRoleCoordinator),
						ServerCount:      to.Ptr[int32](1),
					},
					{
						EnableHa:         to.Ptr(false),
						ServerEdition:    to.Ptr(armpostgresqlhsc.ServerEditionMemoryOptimized),
						StorageQuotaInMb: to.Ptr[int64](524288),
						VCores:           to.Ptr[int64](4),
						Name:             to.Ptr("<name>"),
						Role:             to.Ptr(armpostgresqlhsc.ServerRoleWorker),
						ServerCount:      to.Ptr[int32](3),
					}},
				StandbyAvailabilityZone: to.Ptr("<standby-availability-zone>"),
			},
		},
		&armpostgresqlhsc.ServerGroupsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2020-10-05-privatepreview/examples/ServerGroupGet.json
func ExampleServerGroupsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpostgresqlhsc.NewServerGroupsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<server-group-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2020-10-05-privatepreview/examples/ServerGroupDelete.json
func ExampleServerGroupsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpostgresqlhsc.NewServerGroupsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<server-group-name>",
		&armpostgresqlhsc.ServerGroupsClientBeginDeleteOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2020-10-05-privatepreview/examples/ServerGroupAddNode.json
func ExampleServerGroupsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpostgresqlhsc.NewServerGroupsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<server-group-name>",
		armpostgresqlhsc.ServerGroupForUpdate{
			Location: to.Ptr("<location>"),
			Properties: &armpostgresqlhsc.ServerGroupPropertiesForUpdate{
				ServerRoleGroups: []*armpostgresqlhsc.ServerRoleGroup{
					{
						Name:        to.Ptr("<name>"),
						Role:        to.Ptr(armpostgresqlhsc.ServerRoleWorker),
						ServerCount: to.Ptr[int32](10),
					}},
			},
		},
		&armpostgresqlhsc.ServerGroupsClientBeginUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2020-10-05-privatepreview/examples/ServerGroupRestart.json
func ExampleServerGroupsClient_BeginRestart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpostgresqlhsc.NewServerGroupsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginRestart(ctx,
		"<resource-group-name>",
		"<server-group-name>",
		&armpostgresqlhsc.ServerGroupsClientBeginRestartOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2020-10-05-privatepreview/examples/ServerGroupStart.json
func ExampleServerGroupsClient_BeginStart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpostgresqlhsc.NewServerGroupsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginStart(ctx,
		"<resource-group-name>",
		"<server-group-name>",
		&armpostgresqlhsc.ServerGroupsClientBeginStartOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2020-10-05-privatepreview/examples/ServerGroupStop.json
func ExampleServerGroupsClient_BeginStop() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpostgresqlhsc.NewServerGroupsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginStop(ctx,
		"<resource-group-name>",
		"<server-group-name>",
		&armpostgresqlhsc.ServerGroupsClientBeginStopOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2020-10-05-privatepreview/examples/CheckNameAvailability.json
func ExampleServerGroupsClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpostgresqlhsc.NewServerGroupsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CheckNameAvailability(ctx,
		armpostgresqlhsc.NameAvailabilityRequest{
			Name: to.Ptr("<name>"),
			Type: to.Ptr("<type>"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
