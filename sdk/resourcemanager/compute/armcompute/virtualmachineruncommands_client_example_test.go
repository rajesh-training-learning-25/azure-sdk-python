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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/runCommandExamples/RunCommand_List.json
func ExampleVirtualMachineRunCommandsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachineRunCommandsClient().NewListPager("SoutheastAsia", nil)
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
		// page.RunCommandListResult = armcompute.RunCommandListResult{
		// 	Value: []*armcompute.RunCommandDocumentBase{
		// 		{
		// 			Description: to.Ptr("Configure the machine to enable remote PowerShell."),
		// 			Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
		// 			ID: to.Ptr("EnableRemotePS"),
		// 			Label: to.Ptr("Enable remote PowerShell"),
		// 			OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 		},
		// 		{
		// 			Description: to.Ptr("Shows detailed information for the IP address, subnet mask and default gateway for each adapter bound to TCP/IP."),
		// 			Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
		// 			ID: to.Ptr("IPConfig"),
		// 			Label: to.Ptr("List IP configuration"),
		// 			OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 		},
		// 		{
		// 			Description: to.Ptr("Custom multiline PowerShell script should be defined in script property. Optional parameters can be set in parameters property."),
		// 			Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
		// 			ID: to.Ptr("RunPowerShellScript"),
		// 			Label: to.Ptr("Executes a PowerShell script"),
		// 			OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 		},
		// 		{
		// 			Description: to.Ptr("Custom multiline shell script should be defined in script property. Optional parameters can be set in parameters property."),
		// 			Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
		// 			ID: to.Ptr("RunShellScript"),
		// 			Label: to.Ptr("Executes a Linux shell script"),
		// 			OSType: to.Ptr(armcompute.OperatingSystemTypesLinux),
		// 		},
		// 		{
		// 			Description: to.Ptr("Get the configuration of all network interfaces."),
		// 			Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
		// 			ID: to.Ptr("ifconfig"),
		// 			Label: to.Ptr("List network configuration"),
		// 			OSType: to.Ptr(armcompute.OperatingSystemTypesLinux),
		// 		},
		// 		{
		// 			Description: to.Ptr("Checks if the local Administrator account is disabled, and if so enables it."),
		// 			Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
		// 			ID: to.Ptr("EnableAdminAccount"),
		// 			Label: to.Ptr("Enable administrator account"),
		// 			OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 		},
		// 		{
		// 			Description: to.Ptr("Reset built-in Administrator account password."),
		// 			Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
		// 			ID: to.Ptr("ResetAccountPassword"),
		// 			Label: to.Ptr("Reset built-in Administrator account password"),
		// 			OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 		},
		// 		{
		// 			Description: to.Ptr("Checks registry settings and domain policy settings. Suggests policy actions if machine is part of a domain or modifies the settings to default values."),
		// 			Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
		// 			ID: to.Ptr("RDPSettings"),
		// 			Label: to.Ptr("Verify RDP Listener Settings"),
		// 			OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 		},
		// 		{
		// 			Description: to.Ptr("Sets the default or user specified port number for Remote Desktop connections. Enables firewall rule for inbound access to the port."),
		// 			Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
		// 			ID: to.Ptr("SetRDPPort"),
		// 			Label: to.Ptr("Set Remote Desktop port"),
		// 			OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 		},
		// 		{
		// 			Description: to.Ptr("Removes the SSL certificate tied to the RDP listener and restores the RDP listerner security to default. Use this script if you see any issues with the certificate."),
		// 			Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
		// 			ID: to.Ptr("ResetRDPCert"),
		// 			Label: to.Ptr("Restore RDP Authentication mode to defaults"),
		// 			OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/runCommandExamples/RunCommand_Get.json
func ExampleVirtualMachineRunCommandsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineRunCommandsClient().Get(ctx, "SoutheastAsia", "RunPowerShellScript", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RunCommandDocument = armcompute.RunCommandDocument{
	// 	Description: to.Ptr("Custom multiline PowerShell script should be defined in script property. Optional parameters can be set in parameters property."),
	// 	Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
	// 	ID: to.Ptr("RunPowerShellScript"),
	// 	Label: to.Ptr("Executes a PowerShell script"),
	// 	OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 	Parameters: []*armcompute.RunCommandParameterDefinition{
	// 		{
	// 			Name: to.Ptr("arg1"),
	// 			Type: to.Ptr("string"),
	// 			DefaultValue: to.Ptr("value1"),
	// 		},
	// 		{
	// 			Name: to.Ptr("arg2"),
	// 			Type: to.Ptr("string"),
	// 			DefaultValue: to.Ptr("value2"),
	// 	}},
	// 	Script: []*string{
	// 		to.Ptr("param("),
	// 		to.Ptr("    [string]$arg1,"),
	// 		to.Ptr("    [string]$arg2"),
	// 		to.Ptr(")"),
	// 		to.Ptr("Write-Host This is a sample script with parameters $arg1 $arg2")},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/runCommandExamples/VirtualMachineRunCommand_CreateOrUpdate.json
func ExampleVirtualMachineRunCommandsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineRunCommandsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myVM", "myRunCommand", armcompute.VirtualMachineRunCommand{
		Location: to.Ptr("West US"),
		Properties: &armcompute.VirtualMachineRunCommandProperties{
			AsyncExecution: to.Ptr(false),
			ErrorBlobURI:   to.Ptr("https://mystorageaccount.blob.core.windows.net/scriptcontainer/scriptURI"),
			OutputBlobManagedIdentity: &armcompute.RunCommandManagedIdentity{
				ClientID: to.Ptr("22d35efb-0c99-4041-8c5b-6d24db33a69a"),
			},
			OutputBlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/myscriptoutputcontainer/MyScriptoutput.txt"),
			Parameters: []*armcompute.RunCommandInputParameter{
				{
					Name:  to.Ptr("param1"),
					Value: to.Ptr("value1"),
				},
				{
					Name:  to.Ptr("param2"),
					Value: to.Ptr("value2"),
				}},
			RunAsPassword: to.Ptr("<runAsPassword>"),
			RunAsUser:     to.Ptr("user1"),
			Source: &armcompute.VirtualMachineRunCommandScriptSource{
				ScriptURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/scriptcontainer/scriptURI"),
			},
			TimeoutInSeconds:                to.Ptr[int32](3600),
			TreatFailureAsDeploymentFailure: to.Ptr(false),
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
	// res.VirtualMachineRunCommand = armcompute.VirtualMachineRunCommand{
	// 	Name: to.Ptr("myRunCommand"),
	// 	Type: to.Ptr("Microsoft.Compute/virtualMachines/runCommands"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM/runCommands/myRunCommand"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armcompute.VirtualMachineRunCommandProperties{
	// 		AsyncExecution: to.Ptr(false),
	// 		ErrorBlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/mycontainer/MyScriptError.txt"),
	// 		OutputBlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/myscriptoutputcontainer/MyScriptoutput.txt"),
	// 		Parameters: []*armcompute.RunCommandInputParameter{
	// 			{
	// 				Name: to.Ptr("param1"),
	// 				Value: to.Ptr("value1"),
	// 			},
	// 			{
	// 				Name: to.Ptr("param2"),
	// 				Value: to.Ptr("value2"),
	// 		}},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		RunAsUser: to.Ptr("user1"),
	// 		Source: &armcompute.VirtualMachineRunCommandScriptSource{
	// 			ScriptURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/scriptcontainer/MyScript.ps1"),
	// 		},
	// 		TimeoutInSeconds: to.Ptr[int32](3600),
	// 		TreatFailureAsDeploymentFailure: to.Ptr(false),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/runCommandExamples/VirtualMachineRunCommand_Update.json
func ExampleVirtualMachineRunCommandsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineRunCommandsClient().BeginUpdate(ctx, "myResourceGroup", "myVM", "myRunCommand", armcompute.VirtualMachineRunCommandUpdate{
		Properties: &armcompute.VirtualMachineRunCommandProperties{
			AsyncExecution: to.Ptr(false),
			ErrorBlobManagedIdentity: &armcompute.RunCommandManagedIdentity{
				ObjectID: to.Ptr("4231e4d2-33e4-4e23-96b2-17888afa6072"),
			},
			ErrorBlobURI:  to.Ptr("https://mystorageaccount.blob.core.windows.net/mycontainer/MyScriptError.txt"),
			OutputBlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/myscriptoutputcontainer/outputUri"),
			Parameters: []*armcompute.RunCommandInputParameter{
				{
					Name:  to.Ptr("param1"),
					Value: to.Ptr("value1"),
				},
				{
					Name:  to.Ptr("param2"),
					Value: to.Ptr("value2"),
				}},
			RunAsPassword: to.Ptr("<runAsPassword>"),
			RunAsUser:     to.Ptr("user1"),
			Source: &armcompute.VirtualMachineRunCommandScriptSource{
				Script: to.Ptr("Write-Host Hello World! ; Remove-Item C:	est	estFile.txt"),
			},
			TimeoutInSeconds: to.Ptr[int32](3600),
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
	// res.VirtualMachineRunCommand = armcompute.VirtualMachineRunCommand{
	// 	Name: to.Ptr("myRunCommand"),
	// 	Type: to.Ptr("Microsoft.Compute/virtualMachines/runCommands"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM/runCommands/myRunCommand"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armcompute.VirtualMachineRunCommandProperties{
	// 		AsyncExecution: to.Ptr(false),
	// 		ErrorBlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/mycontainer/MyScriptError.txt"),
	// 		OutputBlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/myscriptoutputcontainer/MyScriptoutput.txt"),
	// 		Parameters: []*armcompute.RunCommandInputParameter{
	// 			{
	// 				Name: to.Ptr("param1"),
	// 				Value: to.Ptr("value1"),
	// 			},
	// 			{
	// 				Name: to.Ptr("param2"),
	// 				Value: to.Ptr("value2"),
	// 		}},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		RunAsUser: to.Ptr("user1"),
	// 		Source: &armcompute.VirtualMachineRunCommandScriptSource{
	// 			Script: to.Ptr("Write-Host Hello World! ; Remove-Item C:	est	estFile.txt"),
	// 		},
	// 		TimeoutInSeconds: to.Ptr[int32](3600),
	// 		TreatFailureAsDeploymentFailure: to.Ptr(false),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/runCommandExamples/VirtualMachineRunCommand_Delete.json
func ExampleVirtualMachineRunCommandsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineRunCommandsClient().BeginDelete(ctx, "myResourceGroup", "myVM", "myRunCommand", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/runCommandExamples/VirtualMachineRunCommand_Get.json
func ExampleVirtualMachineRunCommandsClient_GetByVirtualMachine() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineRunCommandsClient().GetByVirtualMachine(ctx, "myResourceGroup", "myVM", "myRunCommand", &armcompute.VirtualMachineRunCommandsClientGetByVirtualMachineOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineRunCommand = armcompute.VirtualMachineRunCommand{
	// 	Name: to.Ptr("myRunCommand"),
	// 	Type: to.Ptr("Microsoft.Compute/virtualMachines/runCommands"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM/runCommands/myRunCommand"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armcompute.VirtualMachineRunCommandProperties{
	// 		AsyncExecution: to.Ptr(false),
	// 		ErrorBlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/mycontainer/MyScriptError.txt"),
	// 		OutputBlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/myscriptoutputcontainer/MyScriptoutput.txt"),
	// 		Parameters: []*armcompute.RunCommandInputParameter{
	// 			{
	// 				Name: to.Ptr("param1"),
	// 				Value: to.Ptr("value1"),
	// 			},
	// 			{
	// 				Name: to.Ptr("param2"),
	// 				Value: to.Ptr("value2"),
	// 		}},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		RunAsUser: to.Ptr("user1"),
	// 		Source: &armcompute.VirtualMachineRunCommandScriptSource{
	// 			Script: to.Ptr("Write-Host Hello World! ; Remove-Item C:	est	estFile.txt"),
	// 		},
	// 		TimeoutInSeconds: to.Ptr[int32](3600),
	// 		TreatFailureAsDeploymentFailure: to.Ptr(false),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/runCommandExamples/VirtualMachineRunCommand_List.json
func ExampleVirtualMachineRunCommandsClient_NewListByVirtualMachinePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachineRunCommandsClient().NewListByVirtualMachinePager("myResourceGroup", "myVM", &armcompute.VirtualMachineRunCommandsClientListByVirtualMachineOptions{Expand: nil})
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
		// page.VirtualMachineRunCommandsListResult = armcompute.VirtualMachineRunCommandsListResult{
		// 	Value: []*armcompute.VirtualMachineRunCommand{
		// 		{
		// 			Name: to.Ptr("myRunCommand"),
		// 			Type: to.Ptr("Microsoft.Compute/virtualMachines/runCommands"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM/runCommands/myRunCommand"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"tag1": to.Ptr("value1"),
		// 				"tag2": to.Ptr("value2"),
		// 			},
		// 			Properties: &armcompute.VirtualMachineRunCommandProperties{
		// 				AsyncExecution: to.Ptr(false),
		// 				ErrorBlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/mycontainer/MyScriptError.txt"),
		// 				OutputBlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/myscriptoutputcontainer/MyScriptoutput.txt"),
		// 				Parameters: []*armcompute.RunCommandInputParameter{
		// 					{
		// 						Name: to.Ptr("param1"),
		// 						Value: to.Ptr("value1"),
		// 					},
		// 					{
		// 						Name: to.Ptr("param2"),
		// 						Value: to.Ptr("value2"),
		// 				}},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				RunAsUser: to.Ptr("user1"),
		// 				Source: &armcompute.VirtualMachineRunCommandScriptSource{
		// 					Script: to.Ptr("Write-Host Hello World!"),
		// 				},
		// 				TimeoutInSeconds: to.Ptr[int32](0),
		// 				TreatFailureAsDeploymentFailure: to.Ptr(false),
		// 			},
		// 	}},
		// }
	}
}
