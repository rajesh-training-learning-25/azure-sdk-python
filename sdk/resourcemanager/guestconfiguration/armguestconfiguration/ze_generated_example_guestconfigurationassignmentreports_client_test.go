//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armguestconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2022-01-25/examples/listAllGuestConfigurationAssignmentReports.json
func ExampleAssignmentReportsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armguestconfiguration.NewAssignmentReportsClient("mySubscriptionid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.List(ctx,
		"myResourceGroupName",
		"AuditSecureProtocol",
		"myVMName",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2022-01-25/examples/getGuestConfigurationAssignmentReportById.json
func ExampleAssignmentReportsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armguestconfiguration.NewAssignmentReportsClient("mySubscriptionid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"myResourceGroupName",
		"AuditSecureProtocol",
		"7367cbb8-ae99-47d0-a33b-a283564d2cb1",
		"myvm",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}