//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// DO NOT EDIT.

package armelastic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/elastic/resource-manager/Microsoft.Elastic/preview/2023-02-01-preview/examples/AllTrafficFilters_list.json
func ExampleAllTrafficFiltersClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelastic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAllTrafficFiltersClient().List(ctx, "myResourceGroup", "myMonitor", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TrafficFilterResponse = armelastic.TrafficFilterResponse{
	// 	Rulesets: []*armelastic.TrafficFilter{
	// 		{
	// 			Name: to.Ptr("IPFromApi"),
	// 			Type: to.Ptr(armelastic.TypeIP),
	// 			Description: to.Ptr("created from azure"),
	// 			ID: to.Ptr("31d91b5afb6f4c2eaaf104c97b1991dd"),
	// 			IncludeByDefault: to.Ptr(false),
	// 			Region: to.Ptr("azure-eastus"),
	// 			Rules: []*armelastic.TrafficFilterRule{
	// 				{
	// 					Description: to.Ptr("Allow inbound traffic from IP address 192.168.131.0"),
	// 					ID: to.Ptr("f0297dad72af4a5e964cddf817f35e65"),
	// 					Source: to.Ptr("192.168.131.0"),
	// 				},
	// 				{
	// 					Description: to.Ptr("Allow inbound traffic from IP address block 192.168.132.6/22"),
	// 					ID: to.Ptr("f9c00169f0e54931ae72aabde326b589"),
	// 					Source: to.Ptr("192.168.132.6/22"),
	// 			}},
	// 	}},
	// }
}
