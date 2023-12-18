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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerCreatePointInTimeRestore.json
func ExampleServersClient_BeginCreate_createADatabaseAsAPointInTimeRestore() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServersClient().BeginCreate(ctx, "TargetResourceGroup", "targetserver", armmysql.ServerForCreate{
		Location: to.Ptr("brazilsouth"),
		Properties: &armmysql.ServerPropertiesForRestore{
			CreateMode:         to.Ptr(armmysql.CreateModePointInTimeRestore),
			RestorePointInTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-14T00:00:37.467Z"); return t }()),
			SourceServerID:     to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/SourceResourceGroup/providers/Microsoft.DBforMySQL/servers/sourceserver"),
		},
		SKU: &armmysql.SKU{
			Name:     to.Ptr("GP_Gen5_2"),
			Capacity: to.Ptr[int32](2),
			Family:   to.Ptr("Gen5"),
			Tier:     to.Ptr(armmysql.SKUTierGeneralPurpose),
		},
		Tags: map[string]*string{
			"ElasticServer": to.Ptr("1"),
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
	// res.Server = armmysql.Server{
	// 	Name: to.Ptr("targetserver"),
	// 	Type: to.Ptr("Microsoft.DBforMySQL/servers"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/targetserver"),
	// 	Location: to.Ptr("brazilsouth"),
	// 	Tags: map[string]*string{
	// 		"elasticServer": to.Ptr("1"),
	// 	},
	// 	Properties: &armmysql.ServerProperties{
	// 		AdministratorLogin: to.Ptr("cloudsa"),
	// 		EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-14T18:02:41.577Z"); return t}()),
	// 		FullyQualifiedDomainName: to.Ptr("targetserver.mysql.database.azure.com"),
	// 		SSLEnforcement: to.Ptr(armmysql.SSLEnforcementEnumEnabled),
	// 		StorageProfile: &armmysql.StorageProfile{
	// 			BackupRetentionDays: to.Ptr[int32](7),
	// 			GeoRedundantBackup: to.Ptr(armmysql.GeoRedundantBackupEnabled),
	// 			StorageMB: to.Ptr[int32](128000),
	// 		},
	// 		UserVisibleState: to.Ptr(armmysql.ServerStateReady),
	// 		Version: to.Ptr(armmysql.ServerVersionFive7),
	// 	},
	// 	SKU: &armmysql.SKU{
	// 		Name: to.Ptr("GP_Gen5_2"),
	// 		Capacity: to.Ptr[int32](2),
	// 		Family: to.Ptr("Gen5"),
	// 		Tier: to.Ptr(armmysql.SKUTierGeneralPurpose),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerCreate.json
func ExampleServersClient_BeginCreate_createANewServer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServersClient().BeginCreate(ctx, "testrg", "mysqltestsvc4", armmysql.ServerForCreate{
		Location: to.Ptr("westus"),
		Properties: &armmysql.ServerPropertiesForDefaultCreate{
			CreateMode:     to.Ptr(armmysql.CreateModeDefault),
			SSLEnforcement: to.Ptr(armmysql.SSLEnforcementEnumEnabled),
			StorageProfile: &armmysql.StorageProfile{
				BackupRetentionDays: to.Ptr[int32](7),
				GeoRedundantBackup:  to.Ptr(armmysql.GeoRedundantBackupEnabled),
				StorageMB:           to.Ptr[int32](128000),
			},
			AdministratorLogin:         to.Ptr("cloudsa"),
			AdministratorLoginPassword: to.Ptr("<administratorLoginPassword>"),
		},
		SKU: &armmysql.SKU{
			Name:     to.Ptr("GP_Gen5_2"),
			Capacity: to.Ptr[int32](2),
			Family:   to.Ptr("Gen5"),
			Tier:     to.Ptr(armmysql.SKUTierGeneralPurpose),
		},
		Tags: map[string]*string{
			"ElasticServer": to.Ptr("1"),
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
	// res.Server = armmysql.Server{
	// 	Name: to.Ptr("mysqltestsvc4"),
	// 	Type: to.Ptr("Microsoft.DBforMySQL/servers"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc4"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"elasticServer": to.Ptr("1"),
	// 	},
	// 	Properties: &armmysql.ServerProperties{
	// 		AdministratorLogin: to.Ptr("cloudsa"),
	// 		EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-14T18:02:41.577Z"); return t}()),
	// 		FullyQualifiedDomainName: to.Ptr("mysqltestsvc4.mysql.database.azure.com"),
	// 		SSLEnforcement: to.Ptr(armmysql.SSLEnforcementEnumEnabled),
	// 		StorageProfile: &armmysql.StorageProfile{
	// 			BackupRetentionDays: to.Ptr[int32](7),
	// 			GeoRedundantBackup: to.Ptr(armmysql.GeoRedundantBackupEnabled),
	// 			StorageMB: to.Ptr[int32](128000),
	// 		},
	// 		UserVisibleState: to.Ptr(armmysql.ServerStateReady),
	// 		Version: to.Ptr(armmysql.ServerVersionFive7),
	// 	},
	// 	SKU: &armmysql.SKU{
	// 		Name: to.Ptr("GP_Gen5_2"),
	// 		Capacity: to.Ptr[int32](2),
	// 		Family: to.Ptr("Gen5"),
	// 		Tier: to.Ptr(armmysql.SKUTierGeneralPurpose),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerCreateReplicaMode.json
func ExampleServersClient_BeginCreate_createAReplicaServer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServersClient().BeginCreate(ctx, "TargetResourceGroup", "targetserver", armmysql.ServerForCreate{
		Location: to.Ptr("westus"),
		Properties: &armmysql.ServerPropertiesForReplica{
			CreateMode:     to.Ptr(armmysql.CreateModeReplica),
			SourceServerID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/MasterResourceGroup/providers/Microsoft.DBforMySQL/servers/masterserver"),
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
	// res.Server = armmysql.Server{
	// 	Name: to.Ptr("targetserver"),
	// 	Type: to.Ptr("Microsoft.DBforMySQL/servers"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TargetResourceGroup/providers/Microsoft.DBforMySQL/servers/targetserver"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"ElasticServer": to.Ptr("1"),
	// 	},
	// 	Properties: &armmysql.ServerProperties{
	// 		AdministratorLogin: to.Ptr("cloudsa"),
	// 		EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-14T18:02:41.577Z"); return t}()),
	// 		FullyQualifiedDomainName: to.Ptr("targetserver.mysql.database.azure.com"),
	// 		MasterServerID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/MasterResourceGroup/providers/Microsoft.DBforMySQL/servers/masterserver"),
	// 		ReplicaCapacity: to.Ptr[int32](0),
	// 		ReplicationRole: to.Ptr("Replica"),
	// 		SSLEnforcement: to.Ptr(armmysql.SSLEnforcementEnumEnabled),
	// 		StorageProfile: &armmysql.StorageProfile{
	// 			BackupRetentionDays: to.Ptr[int32](14),
	// 			GeoRedundantBackup: to.Ptr(armmysql.GeoRedundantBackupEnabled),
	// 			StorageMB: to.Ptr[int32](128000),
	// 		},
	// 		UserVisibleState: to.Ptr(armmysql.ServerStateReady),
	// 		Version: to.Ptr(armmysql.ServerVersionFive7),
	// 	},
	// 	SKU: &armmysql.SKU{
	// 		Name: to.Ptr("GP_Gen5_2"),
	// 		Capacity: to.Ptr[int32](2),
	// 		Family: to.Ptr("Gen5"),
	// 		Tier: to.Ptr(armmysql.SKUTierGeneralPurpose),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerCreateGeoRestoreMode.json
func ExampleServersClient_BeginCreate_createAServerAsAGeoRestore() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServersClient().BeginCreate(ctx, "TargetResourceGroup", "targetserver", armmysql.ServerForCreate{
		Location: to.Ptr("westus"),
		Properties: &armmysql.ServerPropertiesForGeoRestore{
			CreateMode:     to.Ptr(armmysql.CreateModeGeoRestore),
			SourceServerID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/SourceResourceGroup/providers/Microsoft.DBforMySQL/servers/sourceserver"),
		},
		SKU: &armmysql.SKU{
			Name:     to.Ptr("GP_Gen5_2"),
			Capacity: to.Ptr[int32](2),
			Family:   to.Ptr("Gen5"),
			Tier:     to.Ptr(armmysql.SKUTierGeneralPurpose),
		},
		Tags: map[string]*string{
			"ElasticServer": to.Ptr("1"),
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
	// res.Server = armmysql.Server{
	// 	Name: to.Ptr("targetserver"),
	// 	Type: to.Ptr("Microsoft.DBforMySQL/servers"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/targetserver"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"elasticServer": to.Ptr("1"),
	// 	},
	// 	Properties: &armmysql.ServerProperties{
	// 		AdministratorLogin: to.Ptr("cloudsa"),
	// 		EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-14T18:02:41.577Z"); return t}()),
	// 		FullyQualifiedDomainName: to.Ptr("targetserver.mysql.database.azure.com"),
	// 		SSLEnforcement: to.Ptr(armmysql.SSLEnforcementEnumEnabled),
	// 		StorageProfile: &armmysql.StorageProfile{
	// 			BackupRetentionDays: to.Ptr[int32](14),
	// 			GeoRedundantBackup: to.Ptr(armmysql.GeoRedundantBackupEnabled),
	// 			StorageMB: to.Ptr[int32](128000),
	// 		},
	// 		UserVisibleState: to.Ptr(armmysql.ServerStateReady),
	// 		Version: to.Ptr(armmysql.ServerVersionFive7),
	// 	},
	// 	SKU: &armmysql.SKU{
	// 		Name: to.Ptr("GP_Gen5_2"),
	// 		Capacity: to.Ptr[int32](2),
	// 		Family: to.Ptr("Gen5"),
	// 		Tier: to.Ptr(armmysql.SKUTierGeneralPurpose),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerUpdate.json
func ExampleServersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServersClient().BeginUpdate(ctx, "testrg", "mysqltestsvc4", armmysql.ServerUpdateParameters{
		Properties: &armmysql.ServerUpdateParametersProperties{
			AdministratorLoginPassword: to.Ptr("<administratorLoginPassword>"),
			SSLEnforcement:             to.Ptr(armmysql.SSLEnforcementEnumDisabled),
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
	// res.Server = armmysql.Server{
	// 	Name: to.Ptr("mysqltestsvc4"),
	// 	Type: to.Ptr("Microsoft.DBforMySQL/servers"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc4"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"ElasticServer": to.Ptr("1"),
	// 	},
	// 	Properties: &armmysql.ServerProperties{
	// 		AdministratorLogin: to.Ptr("cloudsa"),
	// 		EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-14T18:02:41.577Z"); return t}()),
	// 		FullyQualifiedDomainName: to.Ptr("mysqltestsvc4.mysql.database.azure.com"),
	// 		SSLEnforcement: to.Ptr(armmysql.SSLEnforcementEnumDisabled),
	// 		StorageProfile: &armmysql.StorageProfile{
	// 			BackupRetentionDays: to.Ptr[int32](7),
	// 			GeoRedundantBackup: to.Ptr(armmysql.GeoRedundantBackupEnabled),
	// 			StorageMB: to.Ptr[int32](128000),
	// 		},
	// 		UserVisibleState: to.Ptr(armmysql.ServerStateReady),
	// 		Version: to.Ptr(armmysql.ServerVersionFive7),
	// 	},
	// 	SKU: &armmysql.SKU{
	// 		Name: to.Ptr("GP_Gen4_2"),
	// 		Capacity: to.Ptr[int32](2),
	// 		Family: to.Ptr("Gen4"),
	// 		Tier: to.Ptr(armmysql.SKUTierGeneralPurpose),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerDelete.json
func ExampleServersClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServersClient().BeginDelete(ctx, "TestGroup", "testserver", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerGet.json
func ExampleServersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServersClient().Get(ctx, "testrg", "mysqltestsvc4", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Server = armmysql.Server{
	// 	Name: to.Ptr("mysqltestsvc4"),
	// 	Type: to.Ptr("Microsoft.DBforMySQL/servers"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc4"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"ElasticServer": to.Ptr("1"),
	// 	},
	// 	Properties: &armmysql.ServerProperties{
	// 		AdministratorLogin: to.Ptr("cloudsa"),
	// 		EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-14T18:02:41.577Z"); return t}()),
	// 		FullyQualifiedDomainName: to.Ptr("mysqltestsvc4.mysql.database.azure.com"),
	// 		MasterServerID: to.Ptr(""),
	// 		PrivateEndpointConnections: []*armmysql.ServerPrivateEndpointConnection{
	// 			{
	// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc4/privateEndpointConnections/private-endpoint-name-00000000-1111-2222-3333-444444444444"),
	// 				Properties: &armmysql.ServerPrivateEndpointConnectionProperties{
	// 					PrivateEndpoint: &armmysql.PrivateEndpointProperty{
	// 						ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
	// 					},
	// 					PrivateLinkServiceConnectionState: &armmysql.ServerPrivateLinkServiceConnectionStateProperty{
	// 						Description: to.Ptr("Auto-approved"),
	// 						ActionsRequired: to.Ptr(armmysql.PrivateLinkServiceConnectionStateActionsRequireNone),
	// 						Status: to.Ptr(armmysql.PrivateLinkServiceConnectionStateStatusApproved),
	// 					},
	// 					ProvisioningState: to.Ptr(armmysql.PrivateEndpointProvisioningState("Succeeded")),
	// 				},
	// 		}},
	// 		PublicNetworkAccess: to.Ptr(armmysql.PublicNetworkAccessEnumEnabled),
	// 		ReplicaCapacity: to.Ptr[int32](5),
	// 		ReplicationRole: to.Ptr("None"),
	// 		SSLEnforcement: to.Ptr(armmysql.SSLEnforcementEnumEnabled),
	// 		StorageProfile: &armmysql.StorageProfile{
	// 			BackupRetentionDays: to.Ptr[int32](7),
	// 			GeoRedundantBackup: to.Ptr(armmysql.GeoRedundantBackupEnabled),
	// 			StorageMB: to.Ptr[int32](128000),
	// 		},
	// 		UserVisibleState: to.Ptr(armmysql.ServerStateReady),
	// 		Version: to.Ptr(armmysql.ServerVersionFive7),
	// 	},
	// 	SKU: &armmysql.SKU{
	// 		Name: to.Ptr("GP_Gen4_2"),
	// 		Capacity: to.Ptr[int32](2),
	// 		Family: to.Ptr("Gen4"),
	// 		Tier: to.Ptr(armmysql.SKUTierGeneralPurpose),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerListByResourceGroup.json
func ExampleServersClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServersClient().NewListByResourceGroupPager("testrg", nil)
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
		// page.ServerListResult = armmysql.ServerListResult{
		// 	Value: []*armmysql.Server{
		// 		{
		// 			Name: to.Ptr("mysqltestsvc1"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armmysql.ServerProperties{
		// 				AdministratorLogin: to.Ptr("testuser"),
		// 				EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-07T18:17:35.729Z"); return t}()),
		// 				FullyQualifiedDomainName: to.Ptr("mysqltestsvc1.mysql.database.azure.com"),
		// 				PrivateEndpointConnections: []*armmysql.ServerPrivateEndpointConnection{
		// 				},
		// 				PublicNetworkAccess: to.Ptr(armmysql.PublicNetworkAccessEnumEnabled),
		// 				SSLEnforcement: to.Ptr(armmysql.SSLEnforcementEnumEnabled),
		// 				StorageProfile: &armmysql.StorageProfile{
		// 					BackupRetentionDays: to.Ptr[int32](7),
		// 					GeoRedundantBackup: to.Ptr(armmysql.GeoRedundantBackupDisabled),
		// 					StorageMB: to.Ptr[int32](5120),
		// 				},
		// 				UserVisibleState: to.Ptr(armmysql.ServerStateReady),
		// 				Version: to.Ptr(armmysql.ServerVersionFive7),
		// 			},
		// 			SKU: &armmysql.SKU{
		// 				Name: to.Ptr("B_Gen4_1"),
		// 				Capacity: to.Ptr[int32](1),
		// 				Family: to.Ptr("Gen4"),
		// 				Tier: to.Ptr(armmysql.SKUTierBasic),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("mysqltstsvc2"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltstsvc2"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armmysql.ServerProperties{
		// 				AdministratorLogin: to.Ptr("testuser"),
		// 				EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-07T18:17:35.729Z"); return t}()),
		// 				FullyQualifiedDomainName: to.Ptr("mysqltstsvc2.mysql.database.azure.com"),
		// 				PrivateEndpointConnections: []*armmysql.ServerPrivateEndpointConnection{
		// 					{
		// 						ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltstsvc2/privateEndpointConnections/private-endpoint-name-00000000-1111-2222-3333-444444444444"),
		// 						Properties: &armmysql.ServerPrivateEndpointConnectionProperties{
		// 							PrivateEndpoint: &armmysql.PrivateEndpointProperty{
		// 								ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
		// 							},
		// 							PrivateLinkServiceConnectionState: &armmysql.ServerPrivateLinkServiceConnectionStateProperty{
		// 								Description: to.Ptr("Auto-approved"),
		// 								ActionsRequired: to.Ptr(armmysql.PrivateLinkServiceConnectionStateActionsRequireNone),
		// 								Status: to.Ptr(armmysql.PrivateLinkServiceConnectionStateStatusApproved),
		// 							},
		// 							ProvisioningState: to.Ptr(armmysql.PrivateEndpointProvisioningState("Succeeded")),
		// 						},
		// 				}},
		// 				PublicNetworkAccess: to.Ptr(armmysql.PublicNetworkAccessEnumEnabled),
		// 				SSLEnforcement: to.Ptr(armmysql.SSLEnforcementEnumEnabled),
		// 				StorageProfile: &armmysql.StorageProfile{
		// 					BackupRetentionDays: to.Ptr[int32](7),
		// 					GeoRedundantBackup: to.Ptr(armmysql.GeoRedundantBackupDisabled),
		// 					StorageMB: to.Ptr[int32](5120),
		// 				},
		// 				UserVisibleState: to.Ptr(armmysql.ServerStateReady),
		// 				Version: to.Ptr(armmysql.ServerVersionFive7),
		// 			},
		// 			SKU: &armmysql.SKU{
		// 				Name: to.Ptr("GP_Gen4_2"),
		// 				Capacity: to.Ptr[int32](2),
		// 				Family: to.Ptr("Gen4"),
		// 				Tier: to.Ptr(armmysql.SKUTierGeneralPurpose),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerList.json
func ExampleServersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServersClient().NewListPager(nil)
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
		// page.ServerListResult = armmysql.ServerListResult{
		// 	Value: []*armmysql.Server{
		// 		{
		// 			Name: to.Ptr("mysqltestsvc1"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armmysql.ServerProperties{
		// 				AdministratorLogin: to.Ptr("testuser"),
		// 				EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-28T23:56:02.627Z"); return t}()),
		// 				FullyQualifiedDomainName: to.Ptr("mysqltestsvc1.mysql.database.azure.com"),
		// 				PrivateEndpointConnections: []*armmysql.ServerPrivateEndpointConnection{
		// 				},
		// 				PublicNetworkAccess: to.Ptr(armmysql.PublicNetworkAccessEnumEnabled),
		// 				SSLEnforcement: to.Ptr(armmysql.SSLEnforcementEnumEnabled),
		// 				StorageProfile: &armmysql.StorageProfile{
		// 					BackupRetentionDays: to.Ptr[int32](7),
		// 					GeoRedundantBackup: to.Ptr(armmysql.GeoRedundantBackupDisabled),
		// 					StorageMB: to.Ptr[int32](5120),
		// 				},
		// 				UserVisibleState: to.Ptr(armmysql.ServerStateReady),
		// 				Version: to.Ptr(armmysql.ServerVersionFive7),
		// 			},
		// 			SKU: &armmysql.SKU{
		// 				Name: to.Ptr("B_Gen4_2"),
		// 				Capacity: to.Ptr[int32](2),
		// 				Family: to.Ptr("Gen4"),
		// 				Tier: to.Ptr(armmysql.SKUTierBasic),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("mysqltstsvc2"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltstsvc2"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armmysql.ServerProperties{
		// 				AdministratorLogin: to.Ptr("testuser"),
		// 				EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-28T23:56:54.300Z"); return t}()),
		// 				FullyQualifiedDomainName: to.Ptr("mysqltstsvc2.mysql.database.azure.com"),
		// 				PrivateEndpointConnections: []*armmysql.ServerPrivateEndpointConnection{
		// 					{
		// 						ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltstsvc2/privateEndpointConnections/private-endpoint-name-00000000-1111-2222-3333-444444444444"),
		// 						Properties: &armmysql.ServerPrivateEndpointConnectionProperties{
		// 							PrivateEndpoint: &armmysql.PrivateEndpointProperty{
		// 								ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
		// 							},
		// 							PrivateLinkServiceConnectionState: &armmysql.ServerPrivateLinkServiceConnectionStateProperty{
		// 								Description: to.Ptr("Auto-approved"),
		// 								ActionsRequired: to.Ptr(armmysql.PrivateLinkServiceConnectionStateActionsRequireNone),
		// 								Status: to.Ptr(armmysql.PrivateLinkServiceConnectionStateStatusApproved),
		// 							},
		// 							ProvisioningState: to.Ptr(armmysql.PrivateEndpointProvisioningState("Succeeded")),
		// 						},
		// 				}},
		// 				PublicNetworkAccess: to.Ptr(armmysql.PublicNetworkAccessEnumEnabled),
		// 				SSLEnforcement: to.Ptr(armmysql.SSLEnforcementEnumEnabled),
		// 				StorageProfile: &armmysql.StorageProfile{
		// 					BackupRetentionDays: to.Ptr[int32](7),
		// 					GeoRedundantBackup: to.Ptr(armmysql.GeoRedundantBackupDisabled),
		// 					StorageMB: to.Ptr[int32](5120),
		// 				},
		// 				UserVisibleState: to.Ptr(armmysql.ServerStateReady),
		// 				Version: to.Ptr(armmysql.ServerVersionFive7),
		// 			},
		// 			SKU: &armmysql.SKU{
		// 				Name: to.Ptr("GP_Gen4_2"),
		// 				Capacity: to.Ptr[int32](2),
		// 				Family: to.Ptr("Gen4"),
		// 				Tier: to.Ptr(armmysql.SKUTierGeneralPurpose),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("mysqltestsvc3"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg1/providers/Microsoft.DBforMySQL/servers/mysqltestsvc3"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armmysql.ServerProperties{
		// 				AdministratorLogin: to.Ptr("testuser"),
		// 				EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-28T23:59:44.847Z"); return t}()),
		// 				FullyQualifiedDomainName: to.Ptr("mysqltestsvc3.mysql.database.azure.com"),
		// 				PrivateEndpointConnections: []*armmysql.ServerPrivateEndpointConnection{
		// 					{
		// 						ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc3/privateEndpointConnections/private-endpoint-name-00000000-1111-2222-3333-444444444444"),
		// 						Properties: &armmysql.ServerPrivateEndpointConnectionProperties{
		// 							PrivateEndpoint: &armmysql.PrivateEndpointProperty{
		// 								ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
		// 							},
		// 							PrivateLinkServiceConnectionState: &armmysql.ServerPrivateLinkServiceConnectionStateProperty{
		// 								Description: to.Ptr("Auto-approved"),
		// 								ActionsRequired: to.Ptr(armmysql.PrivateLinkServiceConnectionStateActionsRequireNone),
		// 								Status: to.Ptr(armmysql.PrivateLinkServiceConnectionStateStatusApproved),
		// 							},
		// 							ProvisioningState: to.Ptr(armmysql.PrivateEndpointProvisioningState("Succeeded")),
		// 						},
		// 				}},
		// 				PublicNetworkAccess: to.Ptr(armmysql.PublicNetworkAccessEnumEnabled),
		// 				SSLEnforcement: to.Ptr(armmysql.SSLEnforcementEnumEnabled),
		// 				StorageProfile: &armmysql.StorageProfile{
		// 					BackupRetentionDays: to.Ptr[int32](35),
		// 					GeoRedundantBackup: to.Ptr(armmysql.GeoRedundantBackupEnabled),
		// 					StorageMB: to.Ptr[int32](102400),
		// 				},
		// 				UserVisibleState: to.Ptr(armmysql.ServerStateReady),
		// 				Version: to.Ptr(armmysql.ServerVersionFive7),
		// 			},
		// 			SKU: &armmysql.SKU{
		// 				Name: to.Ptr("GP_Gen4_4"),
		// 				Capacity: to.Ptr[int32](4),
		// 				Family: to.Ptr("Gen4"),
		// 				Tier: to.Ptr(armmysql.SKUTierGeneralPurpose),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerRestart.json
func ExampleServersClient_BeginRestart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServersClient().BeginRestart(ctx, "TestGroup", "testserver", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2020-01-01/examples/ServerStart.json
func ExampleServersClient_BeginStart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServersClient().BeginStart(ctx, "TestGroup", "testserver", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2020-01-01/examples/ServerStop.json
func ExampleServersClient_BeginStop() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServersClient().BeginStop(ctx, "TestGroup", "testserver", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2020-01-01/examples/ServerUpgrade.json
func ExampleServersClient_BeginUpgrade() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServersClient().BeginUpgrade(ctx, "TestGroup", "testserver", armmysql.ServerUpgradeParameters{
		Properties: &armmysql.ServerUpgradeParametersProperties{
			TargetServerVersion: to.Ptr("5.7"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
