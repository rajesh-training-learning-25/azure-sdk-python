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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c7b98b36e4023331545051284d8500adf98f02fe/specification/sql/resource-manager/Microsoft.Sql/preview/2022-05-01-preview/examples/ManagedInstanceDtcList.json
func ExampleManagedInstanceDtcsClient_NewListByManagedInstancePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedInstanceDtcsClient().NewListByManagedInstancePager("testrg", "testinstance", nil)
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
		// page.ManagedInstanceDtcListResult = armsql.ManagedInstanceDtcListResult{
		// 	Value: []*armsql.ManagedInstanceDtc{
		// 		{
		// 			Name: to.Ptr("current"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/dtc"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testinstance/dtc/current"),
		// 			Properties: &armsql.ManagedInstanceDtcProperties{
		// 				DtcEnabled: to.Ptr(true),
		// 				DtcHostNameDNSSuffix: to.Ptr("dtcHostNameSuffixExample.com"),
		// 				ExternalDNSSuffixSearchList: []*string{
		// 					to.Ptr("dns.example1.com"),
		// 					to.Ptr("dns.example2.com")},
		// 					ProvisioningState: to.Ptr(armsql.ProvisioningStateSucceeded),
		// 					SecuritySettings: &armsql.ManagedInstanceDtcSecuritySettings{
		// 						SnaLu6Point2TransactionsEnabled: to.Ptr(false),
		// 						TransactionManagerCommunicationSettings: &armsql.ManagedInstanceDtcTransactionManagerCommunicationSettings{
		// 							AllowInboundEnabled: to.Ptr(false),
		// 							AllowOutboundEnabled: to.Ptr(true),
		// 							Authentication: to.Ptr("NoAuth"),
		// 						},
		// 						XaTransactionsDefaultTimeout: to.Ptr[int32](1000),
		// 						XaTransactionsEnabled: to.Ptr(false),
		// 						XaTransactionsMaximumTimeout: to.Ptr[int32](3000),
		// 					},
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c7b98b36e4023331545051284d8500adf98f02fe/specification/sql/resource-manager/Microsoft.Sql/preview/2022-05-01-preview/examples/ManagedInstanceDtcGet.json
func ExampleManagedInstanceDtcsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedInstanceDtcsClient().Get(ctx, "testrg", "testinstance", armsql.DtcNameCurrent, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedInstanceDtc = armsql.ManagedInstanceDtc{
	// 	Name: to.Ptr("current"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/dtc"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testinstance/dtc/current"),
	// 	Properties: &armsql.ManagedInstanceDtcProperties{
	// 		DtcEnabled: to.Ptr(true),
	// 		DtcHostNameDNSSuffix: to.Ptr("dtcHostNameSuffixExample.com"),
	// 		ExternalDNSSuffixSearchList: []*string{
	// 			to.Ptr("dns.example1.com"),
	// 			to.Ptr("dns.example2.com")},
	// 			ProvisioningState: to.Ptr(armsql.ProvisioningStateSucceeded),
	// 			SecuritySettings: &armsql.ManagedInstanceDtcSecuritySettings{
	// 				SnaLu6Point2TransactionsEnabled: to.Ptr(false),
	// 				TransactionManagerCommunicationSettings: &armsql.ManagedInstanceDtcTransactionManagerCommunicationSettings{
	// 					AllowInboundEnabled: to.Ptr(false),
	// 					AllowOutboundEnabled: to.Ptr(true),
	// 					Authentication: to.Ptr("NoAuth"),
	// 				},
	// 				XaTransactionsDefaultTimeout: to.Ptr[int32](1000),
	// 				XaTransactionsEnabled: to.Ptr(false),
	// 				XaTransactionsMaximumTimeout: to.Ptr[int32](3000),
	// 			},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c7b98b36e4023331545051284d8500adf98f02fe/specification/sql/resource-manager/Microsoft.Sql/preview/2022-05-01-preview/examples/ManagedInstanceDtcUpdateEnableDtc.json
func ExampleManagedInstanceDtcsClient_BeginCreateOrUpdate_updatesManagedInstanceDtcSettingsByEnablingDtc() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedInstanceDtcsClient().BeginCreateOrUpdate(ctx, "testrg", "testinstance", armsql.DtcNameCurrent, armsql.ManagedInstanceDtc{
		Properties: &armsql.ManagedInstanceDtcProperties{
			DtcEnabled: to.Ptr(true),
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
	// res.ManagedInstanceDtc = armsql.ManagedInstanceDtc{
	// 	Name: to.Ptr("current"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/dtc"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testinstance/dtc/current"),
	// 	Properties: &armsql.ManagedInstanceDtcProperties{
	// 		DtcEnabled: to.Ptr(true),
	// 		DtcHostNameDNSSuffix: to.Ptr("dtcHostNameSuffixExample.com"),
	// 		ExternalDNSSuffixSearchList: []*string{
	// 			to.Ptr("dns.example1.com"),
	// 			to.Ptr("dns.example2.com")},
	// 			ProvisioningState: to.Ptr(armsql.ProvisioningStateSucceeded),
	// 			SecuritySettings: &armsql.ManagedInstanceDtcSecuritySettings{
	// 				SnaLu6Point2TransactionsEnabled: to.Ptr(false),
	// 				TransactionManagerCommunicationSettings: &armsql.ManagedInstanceDtcTransactionManagerCommunicationSettings{
	// 					AllowInboundEnabled: to.Ptr(false),
	// 					AllowOutboundEnabled: to.Ptr(true),
	// 					Authentication: to.Ptr("NoAuth"),
	// 				},
	// 				XaTransactionsDefaultTimeout: to.Ptr[int32](1000),
	// 				XaTransactionsEnabled: to.Ptr(false),
	// 				XaTransactionsMaximumTimeout: to.Ptr[int32](3000),
	// 			},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c7b98b36e4023331545051284d8500adf98f02fe/specification/sql/resource-manager/Microsoft.Sql/preview/2022-05-01-preview/examples/ManagedInstanceDtcUpdateMax.json
func ExampleManagedInstanceDtcsClient_BeginCreateOrUpdate_updatesManagedInstanceDtcSettingsWithAllOptionalParametersSpecified() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedInstanceDtcsClient().BeginCreateOrUpdate(ctx, "testrg", "testinstance", armsql.DtcNameCurrent, armsql.ManagedInstanceDtc{
		Properties: &armsql.ManagedInstanceDtcProperties{
			DtcEnabled: to.Ptr(true),
			ExternalDNSSuffixSearchList: []*string{
				to.Ptr("dns.example1.com"),
				to.Ptr("dns.example2.com")},
			SecuritySettings: &armsql.ManagedInstanceDtcSecuritySettings{
				SnaLu6Point2TransactionsEnabled: to.Ptr(false),
				TransactionManagerCommunicationSettings: &armsql.ManagedInstanceDtcTransactionManagerCommunicationSettings{
					AllowInboundEnabled:  to.Ptr(false),
					AllowOutboundEnabled: to.Ptr(true),
					Authentication:       to.Ptr("NoAuth"),
				},
				XaTransactionsDefaultTimeout: to.Ptr[int32](1000),
				XaTransactionsEnabled:        to.Ptr(false),
				XaTransactionsMaximumTimeout: to.Ptr[int32](3000),
			},
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
	// res.ManagedInstanceDtc = armsql.ManagedInstanceDtc{
	// 	Name: to.Ptr("current"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/dtc"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testinstance/dtc/current"),
	// 	Properties: &armsql.ManagedInstanceDtcProperties{
	// 		DtcEnabled: to.Ptr(true),
	// 		DtcHostNameDNSSuffix: to.Ptr("dtcHostNameSuffixExample.com"),
	// 		ExternalDNSSuffixSearchList: []*string{
	// 			to.Ptr("dns.example1.com"),
	// 			to.Ptr("dns.example2.com")},
	// 			ProvisioningState: to.Ptr(armsql.ProvisioningStateSucceeded),
	// 			SecuritySettings: &armsql.ManagedInstanceDtcSecuritySettings{
	// 				SnaLu6Point2TransactionsEnabled: to.Ptr(false),
	// 				TransactionManagerCommunicationSettings: &armsql.ManagedInstanceDtcTransactionManagerCommunicationSettings{
	// 					AllowInboundEnabled: to.Ptr(false),
	// 					AllowOutboundEnabled: to.Ptr(true),
	// 					Authentication: to.Ptr("NoAuth"),
	// 				},
	// 				XaTransactionsDefaultTimeout: to.Ptr[int32](1000),
	// 				XaTransactionsEnabled: to.Ptr(false),
	// 				XaTransactionsMaximumTimeout: to.Ptr[int32](3000),
	// 			},
	// 		},
	// 	}
}
