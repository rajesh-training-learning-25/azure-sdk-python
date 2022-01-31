//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"
)

// x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2021-06-01/examples/PrivateLinkResourcesList.json
func ExamplePrivateLinkResourceClient_ListByBatchAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbatch.NewPrivateLinkResourceClient("<subscription-id>", cred, nil)
	pager := client.ListByBatchAccount("<resource-group-name>",
		"<account-name>",
		&armbatch.PrivateLinkResourceClientListByBatchAccountOptions{Maxresults: nil})
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

// x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2021-06-01/examples/PrivateLinkResourceGet.json
func ExamplePrivateLinkResourceClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbatch.NewPrivateLinkResourceClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<private-link-resource-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateLinkResourceClientGetResult)
}
