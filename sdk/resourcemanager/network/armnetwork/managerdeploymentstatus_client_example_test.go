//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/NetworkManagerDeploymentStatusList.json
func ExampleManagerDeploymentStatusClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagerDeploymentStatusClient().List(ctx, "resoureGroupSample", "testNetworkManager", armnetwork.ManagerDeploymentStatusParameter{
		DeploymentTypes: []*armnetwork.ConfigurationType{
			to.Ptr(armnetwork.ConfigurationTypeConnectivity),
			to.Ptr(armnetwork.ConfigurationType("AdminPolicy"))},
		Regions: []*string{
			to.Ptr("eastus"),
			to.Ptr("westus")},
		SkipToken: to.Ptr("FakeSkipTokenCode"),
	}, &armnetwork.ManagerDeploymentStatusClientListOptions{Top: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagerDeploymentStatusListResult = armnetwork.ManagerDeploymentStatusListResult{
	// 	SkipToken: to.Ptr("NextFakeSkipTokenCode"),
	// 	Value: []*armnetwork.ManagerDeploymentStatus{
	// 		{
	// 			CommitTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-26T06:58:50.883Z"); return t}()),
	// 			ConfigurationIDs: []*string{
	// 				to.Ptr("SecConfig1"),
	// 				to.Ptr("SecConfig2")},
	// 				DeploymentStatus: to.Ptr(armnetwork.DeploymentStatusDeploying),
	// 				DeploymentType: to.Ptr(armnetwork.ConfigurationType("AdminPolicy")),
	// 				ErrorMessage: to.Ptr(""),
	// 				Region: to.Ptr("eastus"),
	// 			},
	// 			{
	// 				CommitTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-26T06:58:50.883Z"); return t}()),
	// 				ConfigurationIDs: []*string{
	// 					to.Ptr("ConnConfig1"),
	// 					to.Ptr("ConnConfig2")},
	// 					DeploymentStatus: to.Ptr(armnetwork.DeploymentStatusDeployed),
	// 					DeploymentType: to.Ptr(armnetwork.ConfigurationTypeConnectivity),
	// 					ErrorMessage: to.Ptr(""),
	// 					Region: to.Ptr("eastus"),
	// 			}},
	// 		}
}
