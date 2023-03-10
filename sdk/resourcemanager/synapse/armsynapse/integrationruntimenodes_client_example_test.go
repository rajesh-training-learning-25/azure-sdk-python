//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimeNodes_Get.json
func ExampleIntegrationRuntimeNodesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewIntegrationRuntimeNodesClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "exampleResourceGroup", "exampleWorkspace", "exampleIntegrationRuntime", "Node_1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SelfHostedIntegrationRuntimeNode = armsynapse.SelfHostedIntegrationRuntimeNode{
	// 	Capabilities: map[string]*string{
	// 		"connectedToResourceManager": to.Ptr("True"),
	// 		"credentialInSync": to.Ptr("True"),
	// 		"httpsPortEnabled": to.Ptr("True"),
	// 		"nodeEnabled": to.Ptr("True"),
	// 		"serviceBusConnected": to.Ptr("True"),
	// 	},
	// 	HostServiceURI: to.Ptr("https://yanzhang-dt.fareast.corp.microsoft.com:8050/HostServiceRemote.svc/"),
	// 	IsActiveDispatcher: to.Ptr(true),
	// 	LastConnectTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-17T06:30:46.6262976Z"); return t}()),
	// 	LastStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-17T03:45:30.8499851Z"); return t}()),
	// 	LastUpdateResult: to.Ptr(armsynapse.IntegrationRuntimeUpdateResultNone),
	// 	MachineName: to.Ptr("YANZHANG-DT"),
	// 	MaxConcurrentJobs: to.Ptr[int32](20),
	// 	NodeName: to.Ptr("Node_1"),
	// 	RegisterTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-17T03:44:55.8012825Z"); return t}()),
	// 	Status: to.Ptr(armsynapse.SelfHostedIntegrationRuntimeNodeStatusOnline),
	// 	Version: to.Ptr("3.8.6743.6"),
	// 	VersionStatus: to.Ptr("UpToDate"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimeNodes_Update.json
func ExampleIntegrationRuntimeNodesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewIntegrationRuntimeNodesClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "exampleResourceGroup", "exampleWorkspace", "exampleIntegrationRuntime", "Node_1", armsynapse.UpdateIntegrationRuntimeNodeRequest{
		ConcurrentJobsLimit: to.Ptr[int32](2),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SelfHostedIntegrationRuntimeNode = armsynapse.SelfHostedIntegrationRuntimeNode{
	// 	Capabilities: map[string]*string{
	// 		"connectedToResourceManager": to.Ptr("True"),
	// 		"credentialInSync": to.Ptr("True"),
	// 		"httpsPortEnabled": to.Ptr("True"),
	// 		"nodeEnabled": to.Ptr("True"),
	// 		"serviceBusConnected": to.Ptr("True"),
	// 	},
	// 	ConcurrentJobsLimit: to.Ptr[int32](2),
	// 	HostServiceURI: to.Ptr("https://yanzhang-dt.fareast.corp.microsoft.com:8050/HostServiceRemote.svc/"),
	// 	IsActiveDispatcher: to.Ptr(true),
	// 	LastConnectTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-14T15:00:35.7544322Z"); return t}()),
	// 	LastStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-14T14:52:59.8933313Z"); return t}()),
	// 	LastUpdateResult: to.Ptr(armsynapse.IntegrationRuntimeUpdateResultNone),
	// 	MachineName: to.Ptr("YANZHANG-DT"),
	// 	MaxConcurrentJobs: to.Ptr[int32](56),
	// 	NodeName: to.Ptr("Node_1"),
	// 	RegisterTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-14T14:51:44.9237069Z"); return t}()),
	// 	Status: to.Ptr(armsynapse.SelfHostedIntegrationRuntimeNodeStatusOnline),
	// 	Version: to.Ptr("3.8.6730.2"),
	// 	VersionStatus: to.Ptr("UpToDate"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimeNodes_Delete.json
func ExampleIntegrationRuntimeNodesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewIntegrationRuntimeNodesClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx, "exampleResourceGroup", "exampleWorkspace", "exampleIntegrationRuntime", "Node_1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
