//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsubscription_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/subscription/resource-manager/Microsoft.Subscription/stable/2021-10-01/examples/changeTenantPolicy.json
func ExamplePolicyClient_AddUpdatePolicyForTenant() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsubscription.NewPolicyClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.AddUpdatePolicyForTenant(ctx,
		armsubscription.PutTenantPolicyRequestProperties{
			BlockSubscriptionsIntoTenant:    to.Ptr(true),
			BlockSubscriptionsLeavingTenant: to.Ptr(true),
			ExemptedPrincipals: []*string{
				to.Ptr("e879cf0f-2b4d-5431-109a-f72fc9868693"),
				to.Ptr("9792da87-c97b-410d-a97d-27021ba09ce6")},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/subscription/resource-manager/Microsoft.Subscription/stable/2021-10-01/examples/getTenantPolicy.json
func ExamplePolicyClient_GetPolicyForTenant() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsubscription.NewPolicyClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetPolicyForTenant(ctx,
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/subscription/resource-manager/Microsoft.Subscription/stable/2021-10-01/examples/getTenantPolicyList.json
func ExamplePolicyClient_NewListPolicyForTenantPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsubscription.NewPolicyClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPolicyForTenantPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}