//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVMExtension_Create.json
func ExampleVirtualMachineScaleSetVMExtensionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineScaleSetVMExtensionsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myvmScaleSet", "0", "myVMExtension", armcompute.VirtualMachineScaleSetVMExtension{
		Properties: &armcompute.VirtualMachineExtensionProperties{
			Type:                    to.Ptr("extType"),
			AutoUpgradeMinorVersion: to.Ptr(true),
			Publisher:               to.Ptr("extPublisher"),
			Settings: map[string]any{
				"UserName": "xyz@microsoft.com",
			},
			TypeHandlerVersion: to.Ptr("1.2"),
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
	// res.VirtualMachineScaleSetVMExtension = armcompute.VirtualMachineScaleSetVMExtension{
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/myvmScaleSet/virtualMachines/0/extensions/myVMExtension"),
	// 	Name: to.Ptr("myVMExtension"),
	// 	Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets/virtualMachines/extensions"),
	// 	Properties: &armcompute.VirtualMachineExtensionProperties{
	// 		Type: to.Ptr("extType"),
	// 		AutoUpgradeMinorVersion: to.Ptr(true),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		Publisher: to.Ptr("extPublisher"),
	// 		Settings: map[string]any{
	// 			"UserName": "xyz@microsoft.com",
	// 		},
	// 		TypeHandlerVersion: to.Ptr("1.2"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVMExtension_Update.json
func ExampleVirtualMachineScaleSetVMExtensionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineScaleSetVMExtensionsClient().BeginUpdate(ctx, "myResourceGroup", "myvmScaleSet", "0", "myVMExtension", armcompute.VirtualMachineScaleSetVMExtensionUpdate{
		Properties: &armcompute.VirtualMachineExtensionUpdateProperties{
			Type:                    to.Ptr("extType"),
			AutoUpgradeMinorVersion: to.Ptr(true),
			Publisher:               to.Ptr("extPublisher"),
			Settings: map[string]any{
				"UserName": "xyz@microsoft.com",
			},
			TypeHandlerVersion: to.Ptr("1.2"),
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
	// res.VirtualMachineScaleSetVMExtension = armcompute.VirtualMachineScaleSetVMExtension{
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/myvmScaleSet/virtualMachines/0/extensions/myVMExtension"),
	// 	Name: to.Ptr("myVMExtension"),
	// 	Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets/virtualMachines/extensions"),
	// 	Properties: &armcompute.VirtualMachineExtensionProperties{
	// 		Type: to.Ptr("extType"),
	// 		AutoUpgradeMinorVersion: to.Ptr(true),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		Publisher: to.Ptr("extPublisher"),
	// 		Settings: map[string]any{
	// 			"UserName": "xyz@microsoft.com",
	// 		},
	// 		TypeHandlerVersion: to.Ptr("1.2"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVMExtension_Delete.json
func ExampleVirtualMachineScaleSetVMExtensionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineScaleSetVMExtensionsClient().BeginDelete(ctx, "myResourceGroup", "myvmScaleSet", "0", "myVMExtension", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVMExtension_Get.json
func ExampleVirtualMachineScaleSetVMExtensionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineScaleSetVMExtensionsClient().Get(ctx, "myResourceGroup", "myvmScaleSet", "0", "myVMExtension", &armcompute.VirtualMachineScaleSetVMExtensionsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineScaleSetVMExtension = armcompute.VirtualMachineScaleSetVMExtension{
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/myvmScaleSet/virtualMachines/0/extensions/myVMExtension"),
	// 	Name: to.Ptr("myVMExtension"),
	// 	Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets/virtualMachines/extensions"),
	// 	Properties: &armcompute.VirtualMachineExtensionProperties{
	// 		Type: to.Ptr("extType"),
	// 		AutoUpgradeMinorVersion: to.Ptr(true),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		Publisher: to.Ptr("extPublisher"),
	// 		Settings: map[string]any{
	// 			"UserName": "xyz@microsoft.com",
	// 		},
	// 		TypeHandlerVersion: to.Ptr("1.2"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVMExtension_List.json
func ExampleVirtualMachineScaleSetVMExtensionsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineScaleSetVMExtensionsClient().List(ctx, "myResourceGroup", "myvmScaleSet", "0", &armcompute.VirtualMachineScaleSetVMExtensionsClientListOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineScaleSetVMExtensionsListResult = armcompute.VirtualMachineScaleSetVMExtensionsListResult{
	// 	Value: []*armcompute.VirtualMachineScaleSetVMExtension{
	// 		{
	// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/myvmScaleSet/virtualMachines/0/extensions/myVMExtension"),
	// 			Name: to.Ptr("myVMExtension"),
	// 			Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets/virtualMachines/extensions"),
	// 			Properties: &armcompute.VirtualMachineExtensionProperties{
	// 				Type: to.Ptr("extType"),
	// 				AutoUpgradeMinorVersion: to.Ptr(true),
	// 				ProvisioningState: to.Ptr("Succeeded"),
	// 				Publisher: to.Ptr("extPublisher"),
	// 				Settings: map[string]any{
	// 					"UserName": "xyz@microsoft.com",
	// 				},
	// 				TypeHandlerVersion: to.Ptr("1.2"),
	// 			},
	// 		},
	// 		{
	// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/myvmScaleSet/virtualMachines/0/extensions/myVMExtension1"),
	// 			Name: to.Ptr("myVMExtension1"),
	// 			Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets/virtualMachines/extensions"),
	// 			Properties: &armcompute.VirtualMachineExtensionProperties{
	// 				Type: to.Ptr("extType1"),
	// 				AutoUpgradeMinorVersion: to.Ptr(true),
	// 				ProvisioningState: to.Ptr("Succeeded"),
	// 				Publisher: to.Ptr("extPublisher1"),
	// 				Settings: map[string]any{
	// 					"UserName": "xyz@microsoft.com",
	// 				},
	// 				TypeHandlerVersion: to.Ptr("1.0"),
	// 			},
	// 	}},
	// }
}
