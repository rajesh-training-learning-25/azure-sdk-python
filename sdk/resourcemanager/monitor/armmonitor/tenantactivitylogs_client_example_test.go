//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fbfcfd2df9357735a95fc0aba82dd4577ffc1e63/specification/monitor/resource-manager/Microsoft.Insights/stable/2015-04-01/examples/GetTenantActivityLogsFiltered.json
func ExampleTenantActivityLogsClient_NewListPager_getTenantActivityLogsWithFilter() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewTenantActivityLogsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager(&armmonitor.TenantActivityLogsClientListOptions{Filter: to.Ptr("eventTimestamp ge '2015-01-21T20:00:00Z' and eventTimestamp le '2015-01-23T20:00:00Z' and resourceGroupName eq 'MSSupportGroup'"),
		Select: nil,
	})
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
		// page.EventDataCollection = armmonitor.EventDataCollection{
		// 	Value: []*armmonitor.EventData{
		// 		{
		// 			OperationID: to.Ptr("1e121103-0ba6-4300-ac9d-952bb5d0c80f"),
		// 			Description: to.Ptr(""),
		// 			Authorization: &armmonitor.SenderAuthorization{
		// 				Action: to.Ptr("microsoft.support/supporttickets/write"),
		// 				Role: to.Ptr("Subscription Admin"),
		// 				Scope: to.Ptr("/subscriptions/089bd33f-d4ec-47fe-8ba5-0753aa5c5b33/resourceGroups/MSSupportGroup/providers/microsoft.support/supporttickets/115012112305841"),
		// 			},
		// 			Caller: to.Ptr("admin@contoso.com"),
		// 			Claims: map[string]*string{
		// 				"name": to.Ptr("John Smith"),
		// 				"appid": to.Ptr("c44b4083-3bq0-49c1-b47d-974e53cbdf3c"),
		// 				"appidacr": to.Ptr("2"),
		// 				"aud": to.Ptr("https://management.core.windows.net/"),
		// 				"exp": to.Ptr("1421880271"),
		// 				"groups": to.Ptr("cacfe77c-e058-4712-83qw-f9b08849fd60,7f71d11d-4c41-4b23-99d2-d32ce7aa621c,31522864-0578-4ea0-9gdc-e66cc564d18c"),
		// 				"http://schemas.microsoft.com/claims/authnclassreference": to.Ptr("1"),
		// 				"http://schemas.microsoft.com/claims/authnmethodsreferences": to.Ptr("pwd"),
		// 				"http://schemas.microsoft.com/identity/claims/objectidentifier": to.Ptr("2468adf0-8211-44e3-95xq-85137af64708"),
		// 				"http://schemas.microsoft.com/identity/claims/scope": to.Ptr("user_impersonation"),
		// 				"http://schemas.microsoft.com/identity/claims/tenantid": to.Ptr("1e8d8218-c5e7-4578-9acc-9abbd5d23315"),
		// 				"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/givenname": to.Ptr("John"),
		// 				"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/name": to.Ptr("admin@contoso.com"),
		// 				"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/nameidentifier": to.Ptr("9vckmEGF7zDKk1YzIY8k0t1_EAPaXoeHyPRn6f413zM"),
		// 				"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/surname": to.Ptr("Smith"),
		// 				"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/upn": to.Ptr("admin@contoso.com"),
		// 				"iat": to.Ptr("1421876371"),
		// 				"iss": to.Ptr("https://sts.windows.net/72f988bf-86f1-41af-91ab-2d7cd011db47/"),
		// 				"nbf": to.Ptr("1421876371"),
		// 				"puid": to.Ptr("20030000801A118C"),
		// 				"ver": to.Ptr("1.0"),
		// 			},
		// 			CorrelationID: to.Ptr("1e121103-0ba6-4300-ac9d-952bb5d0c80f"),
		// 			EventDataID: to.Ptr("44ade6b4-3813-45e6-ae27-7420a95fa2f8"),
		// 			EventName: &armmonitor.LocalizableString{
		// 				LocalizedValue: to.Ptr("End request"),
		// 				Value: to.Ptr("EndRequest"),
		// 			},
		// 			EventTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-01-21T22:14:26.9792776Z"); return t}()),
		// 			HTTPRequest: &armmonitor.HTTPRequestInfo{
		// 				Method: to.Ptr("PUT"),
		// 				ClientIPAddress: to.Ptr("192.168.35.115"),
		// 				ClientRequestID: to.Ptr("27003b25-91d3-418f-8eb1-29e537dcb249"),
		// 			},
		// 			ID: to.Ptr("/subscriptions/089bd33f-d4ec-47fe-8ba5-0753aa5c5b33/resourceGroups/MSSupportGroup/providers/microsoft.support/supporttickets/115012112305841/events/44ade6b4-3813-45e6-ae27-7420a95fa2f8/ticks/635574752669792776"),
		// 			Level: to.Ptr(armmonitor.EventLevelInformational),
		// 			OperationName: &armmonitor.LocalizableString{
		// 				LocalizedValue: to.Ptr("microsoft.support/supporttickets/write"),
		// 				Value: to.Ptr("microsoft.support/supporttickets/write"),
		// 			},
		// 			Properties: map[string]*string{
		// 				"statusCode": to.Ptr("Created"),
		// 			},
		// 			ResourceGroupName: to.Ptr("MSSupportGroup"),
		// 			ResourceProviderName: &armmonitor.LocalizableString{
		// 				LocalizedValue: to.Ptr("microsoft.support"),
		// 				Value: to.Ptr("microsoft.support"),
		// 			},
		// 			Status: &armmonitor.LocalizableString{
		// 				LocalizedValue: to.Ptr("Succeeded"),
		// 				Value: to.Ptr("Succeeded"),
		// 			},
		// 			SubStatus: &armmonitor.LocalizableString{
		// 				LocalizedValue: to.Ptr("Created (HTTP Status Code: 201)"),
		// 				Value: to.Ptr("Created"),
		// 			},
		// 			SubmissionTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-01-21T22:14:39.9936304Z"); return t}()),
		// 			SubscriptionID: to.Ptr("089bd33f-d4ec-47fe-8ba5-0753aa5c5b33"),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fbfcfd2df9357735a95fc0aba82dd4577ffc1e63/specification/monitor/resource-manager/Microsoft.Insights/stable/2015-04-01/examples/GetTenantActivityLogsFilteredAndSelected.json
func ExampleTenantActivityLogsClient_NewListPager_getTenantActivityLogsWithFilterAndSelect() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewTenantActivityLogsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager(&armmonitor.TenantActivityLogsClientListOptions{Filter: to.Ptr("eventTimestamp ge '2015-01-21T20:00:00Z' and eventTimestamp le '2015-01-23T20:00:00Z' and resourceGroupName eq 'MSSupportGroup'"),
		Select: to.Ptr("eventName,id,resourceGroupName,resourceProviderName,operationName,status,eventTimestamp,correlationId,submissionTimestamp,level"),
	})
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
		// page.EventDataCollection = armmonitor.EventDataCollection{
		// 	Value: []*armmonitor.EventData{
		// 		{
		// 			CorrelationID: to.Ptr("1e121103-0ba6-4300-ac9d-952bb5d0c80f"),
		// 			EventName: &armmonitor.LocalizableString{
		// 				LocalizedValue: to.Ptr("End request"),
		// 				Value: to.Ptr("EndRequest"),
		// 			},
		// 			EventTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-01-21T22:14:26.9792776Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/089bd33f-d4ec-47fe-8ba5-0753aa5c5b33/resourceGroups/MSSupportGroup/providers/microsoft.support/supporttickets/115012112305841/events/44ade6b4-3813-45e6-ae27-7420a95fa2f8/ticks/635574752669792776"),
		// 			Level: to.Ptr(armmonitor.EventLevelInformational),
		// 			OperationName: &armmonitor.LocalizableString{
		// 				LocalizedValue: to.Ptr("microsoft.support/supporttickets/write"),
		// 				Value: to.Ptr("microsoft.support/supporttickets/write"),
		// 			},
		// 			ResourceGroupName: to.Ptr("MSSupportGroup"),
		// 			ResourceProviderName: &armmonitor.LocalizableString{
		// 				LocalizedValue: to.Ptr("microsoft.support"),
		// 				Value: to.Ptr("microsoft.support"),
		// 			},
		// 			Status: &armmonitor.LocalizableString{
		// 				LocalizedValue: to.Ptr("Succeeded"),
		// 				Value: to.Ptr("Succeeded"),
		// 			},
		// 			SubmissionTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-01-21T22:14:39.9936304Z"); return t}()),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fbfcfd2df9357735a95fc0aba82dd4577ffc1e63/specification/monitor/resource-manager/Microsoft.Insights/stable/2015-04-01/examples/GetTenantActivityLogsSelected.json
func ExampleTenantActivityLogsClient_NewListPager_getTenantActivityLogsWithSelect() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewTenantActivityLogsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager(&armmonitor.TenantActivityLogsClientListOptions{Filter: nil,
		Select: to.Ptr("eventName,id,resourceGroupName,resourceProviderName,operationName,status,eventTimestamp,correlationId,submissionTimestamp,level"),
	})
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
		// page.EventDataCollection = armmonitor.EventDataCollection{
		// 	Value: []*armmonitor.EventData{
		// 		{
		// 			CorrelationID: to.Ptr("1e121103-0ba6-4300-ac9d-952bb5d0c80f"),
		// 			EventName: &armmonitor.LocalizableString{
		// 				LocalizedValue: to.Ptr("End request"),
		// 				Value: to.Ptr("EndRequest"),
		// 			},
		// 			EventTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-01-21T22:14:26.9792776Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/089bd33f-d4ec-47fe-8ba5-0753aa5c5b33/resourceGroups/MSSupportGroup/providers/microsoft.support/supporttickets/115012112305841/events/44ade6b4-3813-45e6-ae27-7420a95fa2f8/ticks/635574752669792776"),
		// 			Level: to.Ptr(armmonitor.EventLevelInformational),
		// 			OperationName: &armmonitor.LocalizableString{
		// 				LocalizedValue: to.Ptr("microsoft.support/supporttickets/write"),
		// 				Value: to.Ptr("microsoft.support/supporttickets/write"),
		// 			},
		// 			ResourceGroupName: to.Ptr("MSSupportGroup"),
		// 			ResourceProviderName: &armmonitor.LocalizableString{
		// 				LocalizedValue: to.Ptr("microsoft.support"),
		// 				Value: to.Ptr("microsoft.support"),
		// 			},
		// 			Status: &armmonitor.LocalizableString{
		// 				LocalizedValue: to.Ptr("Succeeded"),
		// 				Value: to.Ptr("Succeeded"),
		// 			},
		// 			SubmissionTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-01-21T22:14:39.9936304Z"); return t}()),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fbfcfd2df9357735a95fc0aba82dd4577ffc1e63/specification/monitor/resource-manager/Microsoft.Insights/stable/2015-04-01/examples/GetTenantActivityLogsNoParams.json
func ExampleTenantActivityLogsClient_NewListPager_getTenantActivityLogsWithoutFilterOrSelect() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewTenantActivityLogsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager(&armmonitor.TenantActivityLogsClientListOptions{Filter: nil,
		Select: nil,
	})
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
		// page.EventDataCollection = armmonitor.EventDataCollection{
		// 	Value: []*armmonitor.EventData{
		// 		{
		// 			OperationID: to.Ptr("1e121103-0ba6-4300-ac9d-952bb5d0c80f"),
		// 			Description: to.Ptr(""),
		// 			Authorization: &armmonitor.SenderAuthorization{
		// 				Action: to.Ptr("microsoft.support/supporttickets/write"),
		// 				Role: to.Ptr("Subscription Admin"),
		// 				Scope: to.Ptr("/subscriptions/089bd33f-d4ec-47fe-8ba5-0753aa5c5b33/resourceGroups/MSSupportGroup/providers/microsoft.support/supporttickets/115012112305841"),
		// 			},
		// 			Caller: to.Ptr("admin@contoso.com"),
		// 			Claims: map[string]*string{
		// 				"name": to.Ptr("John Smith"),
		// 				"appid": to.Ptr("c44b4083-3bq0-49c1-b47d-974e53cbdf3c"),
		// 				"appidacr": to.Ptr("2"),
		// 				"aud": to.Ptr("https://management.core.windows.net/"),
		// 				"exp": to.Ptr("1421880271"),
		// 				"groups": to.Ptr("cacfe77c-e058-4712-83qw-f9b08849fd60,7f71d11d-4c41-4b23-99d2-d32ce7aa621c,31522864-0578-4ea0-9gdc-e66cc564d18c"),
		// 				"http://schemas.microsoft.com/claims/authnclassreference": to.Ptr("1"),
		// 				"http://schemas.microsoft.com/claims/authnmethodsreferences": to.Ptr("pwd"),
		// 				"http://schemas.microsoft.com/identity/claims/objectidentifier": to.Ptr("2468adf0-8211-44e3-95xq-85137af64708"),
		// 				"http://schemas.microsoft.com/identity/claims/scope": to.Ptr("user_impersonation"),
		// 				"http://schemas.microsoft.com/identity/claims/tenantid": to.Ptr("1e8d8218-c5e7-4578-9acc-9abbd5d23315"),
		// 				"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/givenname": to.Ptr("John"),
		// 				"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/name": to.Ptr("admin@contoso.com"),
		// 				"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/nameidentifier": to.Ptr("9vckmEGF7zDKk1YzIY8k0t1_EAPaXoeHyPRn6f413zM"),
		// 				"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/surname": to.Ptr("Smith"),
		// 				"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/upn": to.Ptr("admin@contoso.com"),
		// 				"iat": to.Ptr("1421876371"),
		// 				"iss": to.Ptr("https://sts.windows.net/72f988bf-86f1-41af-91ab-2d7cd011db47/"),
		// 				"nbf": to.Ptr("1421876371"),
		// 				"puid": to.Ptr("20030000801A118C"),
		// 				"ver": to.Ptr("1.0"),
		// 			},
		// 			CorrelationID: to.Ptr("1e121103-0ba6-4300-ac9d-952bb5d0c80f"),
		// 			EventDataID: to.Ptr("44ade6b4-3813-45e6-ae27-7420a95fa2f8"),
		// 			EventName: &armmonitor.LocalizableString{
		// 				LocalizedValue: to.Ptr("End request"),
		// 				Value: to.Ptr("EndRequest"),
		// 			},
		// 			EventTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-01-21T22:14:26.9792776Z"); return t}()),
		// 			HTTPRequest: &armmonitor.HTTPRequestInfo{
		// 				Method: to.Ptr("PUT"),
		// 				ClientIPAddress: to.Ptr("192.168.35.115"),
		// 				ClientRequestID: to.Ptr("27003b25-91d3-418f-8eb1-29e537dcb249"),
		// 			},
		// 			ID: to.Ptr("/subscriptions/089bd33f-d4ec-47fe-8ba5-0753aa5c5b33/resourceGroups/MSSupportGroup/providers/microsoft.support/supporttickets/115012112305841/events/44ade6b4-3813-45e6-ae27-7420a95fa2f8/ticks/635574752669792776"),
		// 			Level: to.Ptr(armmonitor.EventLevelInformational),
		// 			OperationName: &armmonitor.LocalizableString{
		// 				LocalizedValue: to.Ptr("microsoft.support/supporttickets/write"),
		// 				Value: to.Ptr("microsoft.support/supporttickets/write"),
		// 			},
		// 			Properties: map[string]*string{
		// 				"statusCode": to.Ptr("Created"),
		// 			},
		// 			ResourceGroupName: to.Ptr("MSSupportGroup"),
		// 			ResourceProviderName: &armmonitor.LocalizableString{
		// 				LocalizedValue: to.Ptr("microsoft.support"),
		// 				Value: to.Ptr("microsoft.support"),
		// 			},
		// 			Status: &armmonitor.LocalizableString{
		// 				LocalizedValue: to.Ptr("Succeeded"),
		// 				Value: to.Ptr("Succeeded"),
		// 			},
		// 			SubStatus: &armmonitor.LocalizableString{
		// 				LocalizedValue: to.Ptr("Created (HTTP Status Code: 201)"),
		// 				Value: to.Ptr("Created"),
		// 			},
		// 			SubmissionTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-01-21T22:14:39.9936304Z"); return t}()),
		// 			SubscriptionID: to.Ptr("089bd33f-d4ec-47fe-8ba5-0753aa5c5b33"),
		// 	}},
		// }
	}
}
