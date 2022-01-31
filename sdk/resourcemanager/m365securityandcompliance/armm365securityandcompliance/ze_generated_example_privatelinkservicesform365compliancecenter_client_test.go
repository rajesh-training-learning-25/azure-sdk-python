//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armm365securityandcompliance_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/m365securityandcompliance/armm365securityandcompliance"
)

// x-ms-original-file: specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/ComplianceCenterServiceGet.json
func ExamplePrivateLinkServicesForM365ComplianceCenterClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armm365securityandcompliance.NewPrivateLinkServicesForM365ComplianceCenterClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<resource-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateLinkServicesForM365ComplianceCenterClientGetResult)
}

// x-ms-original-file: specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/ComplianceCenterServiceCreate.json
func ExamplePrivateLinkServicesForM365ComplianceCenterClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armm365securityandcompliance.NewPrivateLinkServicesForM365ComplianceCenterClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armm365securityandcompliance.PrivateLinkServicesForM365ComplianceCenterDescription{
			Identity: &armm365securityandcompliance.ServicesResourceIdentity{
				Type: armm365securityandcompliance.ManagedServiceIdentityType("SystemAssigned").ToPtr(),
			},
			Kind:     armm365securityandcompliance.KindFhirR4.ToPtr(),
			Location: to.StringPtr("<location>"),
			Tags:     map[string]*string{},
			Properties: &armm365securityandcompliance.ServicesProperties{
				AccessPolicies: []*armm365securityandcompliance.ServiceAccessPolicyEntry{
					{
						ObjectID: to.StringPtr("<object-id>"),
					},
					{
						ObjectID: to.StringPtr("<object-id>"),
					}},
				AuthenticationConfiguration: &armm365securityandcompliance.ServiceAuthenticationConfigurationInfo{
					Audience:          to.StringPtr("<audience>"),
					Authority:         to.StringPtr("<authority>"),
					SmartProxyEnabled: to.BoolPtr(true),
				},
				CorsConfiguration: &armm365securityandcompliance.ServiceCorsConfigurationInfo{
					AllowCredentials: to.BoolPtr(false),
					Headers: []*string{
						to.StringPtr("*")},
					MaxAge: to.Int64Ptr(1440),
					Methods: []*string{
						to.StringPtr("DELETE"),
						to.StringPtr("GET"),
						to.StringPtr("OPTIONS"),
						to.StringPtr("PATCH"),
						to.StringPtr("POST"),
						to.StringPtr("PUT")},
					Origins: []*string{
						to.StringPtr("*")},
				},
				CosmosDbConfiguration: &armm365securityandcompliance.ServiceCosmosDbConfigurationInfo{
					KeyVaultKeyURI:  to.StringPtr("<key-vault-key-uri>"),
					OfferThroughput: to.Int64Ptr(1000),
				},
				ExportConfiguration: &armm365securityandcompliance.ServiceExportConfigurationInfo{
					StorageAccountName: to.StringPtr("<storage-account-name>"),
				},
				PrivateEndpointConnections: []*armm365securityandcompliance.PrivateEndpointConnection{},
				PublicNetworkAccess:        armm365securityandcompliance.PublicNetworkAccess("Disabled").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateLinkServicesForM365ComplianceCenterClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/ComplianceCenterServicePatch.json
func ExamplePrivateLinkServicesForM365ComplianceCenterClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armm365securityandcompliance.NewPrivateLinkServicesForM365ComplianceCenterClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armm365securityandcompliance.ServicesPatchDescription{
			Tags: map[string]*string{
				"tag1": to.StringPtr("value1"),
				"tag2": to.StringPtr("value2"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateLinkServicesForM365ComplianceCenterClientUpdateResult)
}

// x-ms-original-file: specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/ComplianceCenterServiceDelete.json
func ExamplePrivateLinkServicesForM365ComplianceCenterClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armm365securityandcompliance.NewPrivateLinkServicesForM365ComplianceCenterClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<resource-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/ComplianceCenterServiceList.json
func ExamplePrivateLinkServicesForM365ComplianceCenterClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armm365securityandcompliance.NewPrivateLinkServicesForM365ComplianceCenterClient("<subscription-id>", cred, nil)
	pager := client.List(nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/ComplianceCenterServiceListByResourceGroup.json
func ExamplePrivateLinkServicesForM365ComplianceCenterClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armm365securityandcompliance.NewPrivateLinkServicesForM365ComplianceCenterClient("<subscription-id>", cred, nil)
	pager := client.ListByResourceGroup("<resource-group-name>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}
