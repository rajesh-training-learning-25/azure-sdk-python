//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerservice_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-02-01/examples/MaintenanceConfigurationsList.json
func ExampleMaintenanceConfigurationsClient_ListByManagedCluster() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewMaintenanceConfigurationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.ListByManagedCluster("<resource-group-name>",
		"<resource-name>",
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-02-01/examples/MaintenanceConfigurationsGet.json
func ExampleMaintenanceConfigurationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewMaintenanceConfigurationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<config-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-02-01/examples/MaintenanceConfigurationsCreate_Update.json
func ExampleMaintenanceConfigurationsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewMaintenanceConfigurationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<config-name>",
		armcontainerservice.MaintenanceConfiguration{
			Properties: &armcontainerservice.MaintenanceConfigurationProperties{
				NotAllowedTime: []*armcontainerservice.TimeSpan{
					{
						End:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-30T12:00:00Z"); return t }()),
						Start: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-26T03:00:00Z"); return t }()),
					}},
				TimeInWeek: []*armcontainerservice.TimeInWeek{
					{
						Day: to.Ptr(armcontainerservice.WeekDayMonday),
						HourSlots: []*int32{
							to.Ptr[int32](1),
							to.Ptr[int32](2)},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-02-01/examples/MaintenanceConfigurationsDelete.json
func ExampleMaintenanceConfigurationsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewMaintenanceConfigurationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<config-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}
