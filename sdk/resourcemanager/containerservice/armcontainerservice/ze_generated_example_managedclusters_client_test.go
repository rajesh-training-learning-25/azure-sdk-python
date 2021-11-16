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

// x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-09-01/examples/ContainerServiceGetOSOptions.json
func ExampleManagedClustersClient_GetOSOptions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerservice.NewManagedClustersClient("<subscription-id>", cred, nil)
	res, err := client.GetOSOptions(ctx,
		"<location>",
		&armcontainerservice.ManagedClustersGetOSOptionsOptions{ResourceType: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("OSOptionProfile.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-09-01/examples/ManagedClustersList.json
func ExampleManagedClustersClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerservice.NewManagedClustersClient("<subscription-id>", cred, nil)
	pager := client.List(nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("ManagedCluster.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-09-01/examples/ManagedClustersListByResourceGroup.json
func ExampleManagedClustersClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerservice.NewManagedClustersClient("<subscription-id>", cred, nil)
	pager := client.ListByResourceGroup("<resource-group-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("ManagedCluster.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-09-01/examples/ManagedClustersGetUpgradeProfile.json
func ExampleManagedClustersClient_GetUpgradeProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerservice.NewManagedClustersClient("<subscription-id>", cred, nil)
	res, err := client.GetUpgradeProfile(ctx,
		"<resource-group-name>",
		"<resource-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ManagedClusterUpgradeProfile.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-09-01/examples/ManagedClustersGetAccessProfile.json
func ExampleManagedClustersClient_GetAccessProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerservice.NewManagedClustersClient("<subscription-id>", cred, nil)
	res, err := client.GetAccessProfile(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<role-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ManagedClusterAccessProfile.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-09-01/examples/ManagedClustersListClusterCredentialResult.json
func ExampleManagedClustersClient_ListClusterAdminCredentials() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerservice.NewManagedClustersClient("<subscription-id>", cred, nil)
	_, err = client.ListClusterAdminCredentials(ctx,
		"<resource-group-name>",
		"<resource-name>",
		&armcontainerservice.ManagedClustersListClusterAdminCredentialsOptions{ServerFqdn: nil})
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-09-01/examples/ManagedClustersListClusterCredentialResult.json
func ExampleManagedClustersClient_ListClusterUserCredentials() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerservice.NewManagedClustersClient("<subscription-id>", cred, nil)
	_, err = client.ListClusterUserCredentials(ctx,
		"<resource-group-name>",
		"<resource-name>",
		&armcontainerservice.ManagedClustersListClusterUserCredentialsOptions{ServerFqdn: nil})
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-09-01/examples/ManagedClustersListClusterCredentialResult.json
func ExampleManagedClustersClient_ListClusterMonitoringUserCredentials() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerservice.NewManagedClustersClient("<subscription-id>", cred, nil)
	_, err = client.ListClusterMonitoringUserCredentials(ctx,
		"<resource-group-name>",
		"<resource-name>",
		&armcontainerservice.ManagedClustersListClusterMonitoringUserCredentialsOptions{ServerFqdn: nil})
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-09-01/examples/ManagedClustersGet.json
func ExampleManagedClustersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerservice.NewManagedClustersClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<resource-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ManagedCluster.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-09-01/examples/ManagedClustersCreate_Snapshot.json
func ExampleManagedClustersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerservice.NewManagedClustersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armcontainerservice.ManagedCluster{
			Resource: armcontainerservice.Resource{
				Location: to.StringPtr("<location>"),
				Tags: map[string]*string{
					"archv2": to.StringPtr(""),
					"tier":   to.StringPtr("production"),
				},
			},
			Properties: &armcontainerservice.ManagedClusterProperties{
				AddonProfiles: map[string]*armcontainerservice.ManagedClusterAddonProfile{},
				AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
					{
						ManagedClusterAgentPoolProfileProperties: armcontainerservice.ManagedClusterAgentPoolProfileProperties{
							Type:  armcontainerservice.AgentPoolTypeVirtualMachineScaleSets.ToPtr(),
							Count: to.Int32Ptr(3),
							CreationData: &armcontainerservice.CreationData{
								SourceResourceID: to.StringPtr("<source-resource-id>"),
							},
							EnableFIPS:         to.BoolPtr(true),
							EnableNodePublicIP: to.BoolPtr(true),
							Mode:               armcontainerservice.AgentPoolModeSystem.ToPtr(),
							OSType:             armcontainerservice.OSTypeLinux.ToPtr(),
							VMSize:             to.StringPtr("<vmsize>"),
						},
						Name: to.StringPtr("<name>"),
					}},
				AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
					ScaleDownDelayAfterAdd: to.StringPtr("<scale-down-delay-after-add>"),
					ScanInterval:           to.StringPtr("<scan-interval>"),
				},
				DiskEncryptionSetID:     to.StringPtr("<disk-encryption-set-id>"),
				DNSPrefix:               to.StringPtr("<dnsprefix>"),
				EnablePodSecurityPolicy: to.BoolPtr(false),
				EnableRBAC:              to.BoolPtr(true),
				KubernetesVersion:       to.StringPtr("<kubernetes-version>"),
				LinuxProfile: &armcontainerservice.ContainerServiceLinuxProfile{
					AdminUsername: to.StringPtr("<admin-username>"),
					SSH: &armcontainerservice.ContainerServiceSSHConfiguration{
						PublicKeys: []*armcontainerservice.ContainerServiceSSHPublicKey{
							{
								KeyData: to.StringPtr("<key-data>"),
							}},
					},
				},
				NetworkProfile: &armcontainerservice.ContainerServiceNetworkProfile{
					LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
						ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
							Count: to.Int32Ptr(2),
						},
					},
					LoadBalancerSKU: armcontainerservice.LoadBalancerSKUStandard.ToPtr(),
					OutboundType:    armcontainerservice.OutboundTypeLoadBalancer.ToPtr(),
				},
				ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
					ClientID: to.StringPtr("<client-id>"),
					Secret:   to.StringPtr("<secret>"),
				},
				WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
					AdminPassword: to.StringPtr("<admin-password>"),
					AdminUsername: to.StringPtr("<admin-username>"),
				},
			},
			SKU: &armcontainerservice.ManagedClusterSKU{
				Name: armcontainerservice.ManagedClusterSKUNameBasic.ToPtr(),
				Tier: armcontainerservice.ManagedClusterSKUTierFree.ToPtr(),
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
	log.Printf("ManagedCluster.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-09-01/examples/ManagedClustersUpdateTags.json
func ExampleManagedClustersClient_BeginUpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerservice.NewManagedClustersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdateTags(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armcontainerservice.TagsObject{
			Tags: map[string]*string{
				"archv3": to.StringPtr(""),
				"tier":   to.StringPtr("testing"),
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
	log.Printf("ManagedCluster.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-09-01/examples/ManagedClustersDelete.json
func ExampleManagedClustersClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerservice.NewManagedClustersClient("<subscription-id>", cred, nil)
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

// x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-09-01/examples/ManagedClustersResetServicePrincipalProfile.json
func ExampleManagedClustersClient_BeginResetServicePrincipalProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerservice.NewManagedClustersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginResetServicePrincipalProfile(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armcontainerservice.ManagedClusterServicePrincipalProfile{
			ClientID: to.StringPtr("<client-id>"),
			Secret:   to.StringPtr("<secret>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-09-01/examples/ManagedClustersResetAADProfile.json
func ExampleManagedClustersClient_BeginResetAADProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerservice.NewManagedClustersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginResetAADProfile(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armcontainerservice.ManagedClusterAADProfile{
			ClientAppID:     to.StringPtr("<client-app-id>"),
			ServerAppID:     to.StringPtr("<server-app-id>"),
			ServerAppSecret: to.StringPtr("<server-app-secret>"),
			TenantID:        to.StringPtr("<tenant-id>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-09-01/examples/ManagedClustersRotateClusterCertificates.json
func ExampleManagedClustersClient_BeginRotateClusterCertificates() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerservice.NewManagedClustersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginRotateClusterCertificates(ctx,
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

// x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-09-01/examples/ManagedClustersStop.json
func ExampleManagedClustersClient_BeginStop() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerservice.NewManagedClustersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginStop(ctx,
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

// x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-09-01/examples/ManagedClustersStart.json
func ExampleManagedClustersClient_BeginStart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerservice.NewManagedClustersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginStart(ctx,
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

// x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-09-01/examples/RunCommandRequest.json
func ExampleManagedClustersClient_BeginRunCommand() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerservice.NewManagedClustersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginRunCommand(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armcontainerservice.RunCommandRequest{
			ClusterToken: to.StringPtr("<cluster-token>"),
			Command:      to.StringPtr("<command>"),
			Context:      to.StringPtr("<context>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("RunCommandResult.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-09-01/examples/RunCommandResultFailed.json
func ExampleManagedClustersClient_GetCommandResult() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerservice.NewManagedClustersClient("<subscription-id>", cred, nil)
	res, err := client.GetCommandResult(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<command-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("RunCommandResult.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-09-01/examples/OutboundNetworkDependenciesEndpointsList.json
func ExampleManagedClustersClient_ListOutboundNetworkDependenciesEndpoints() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerservice.NewManagedClustersClient("<subscription-id>", cred, nil)
	pager := client.ListOutboundNetworkDependenciesEndpoints("<resource-group-name>",
		"<resource-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
	}
}
