//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSetVMs_Reimage_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetVMsClient_BeginReimage() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetVMsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginReimage(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		"<instance-id>",
		&armcompute.VirtualMachineScaleSetVMsClientBeginReimageOptions{VMScaleSetVMReimageInput: &armcompute.VirtualMachineScaleSetVMReimageParameters{
			TempDisk: to.Ptr(true),
		},
			ResumeToken: "",
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSetVMs_ReimageAll_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetVMsClient_BeginReimageAll() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetVMsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginReimageAll(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		"<instance-id>",
		&armcompute.VirtualMachineScaleSetVMsClientBeginReimageAllOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSetVMs_Deallocate_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetVMsClient_BeginDeallocate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetVMsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginDeallocate(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		"<instance-id>",
		&armcompute.VirtualMachineScaleSetVMsClientBeginDeallocateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSetVMs_Update_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetVMsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetVMsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		"<instance-id>",
		armcompute.VirtualMachineScaleSetVM{
			Location: to.Ptr("<location>"),
			Tags:     map[string]*string{},
			Plan: &armcompute.Plan{
				Name:          to.Ptr("<name>"),
				Product:       to.Ptr("<product>"),
				PromotionCode: to.Ptr("<promotion-code>"),
				Publisher:     to.Ptr("<publisher>"),
			},
			Properties: &armcompute.VirtualMachineScaleSetVMProperties{
				AdditionalCapabilities: &armcompute.AdditionalCapabilities{
					HibernationEnabled: to.Ptr(true),
					UltraSSDEnabled:    to.Ptr(true),
				},
				AvailabilitySet: &armcompute.SubResource{
					ID: to.Ptr("<id>"),
				},
				DiagnosticsProfile: &armcompute.DiagnosticsProfile{
					BootDiagnostics: &armcompute.BootDiagnostics{
						Enabled:    to.Ptr(true),
						StorageURI: to.Ptr("<storage-uri>"),
					},
				},
				HardwareProfile: &armcompute.HardwareProfile{
					VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesBasicA0),
					VMSizeProperties: &armcompute.VMSizeProperties{
						VCPUsAvailable: to.Ptr[int32](9),
						VCPUsPerCore:   to.Ptr[int32](12),
					},
				},
				InstanceView: &armcompute.VirtualMachineScaleSetVMInstanceView{
					BootDiagnostics: &armcompute.BootDiagnosticsInstanceView{
						Status: &armcompute.InstanceViewStatus{
							Code:          to.Ptr("<code>"),
							DisplayStatus: to.Ptr("<display-status>"),
							Level:         to.Ptr(armcompute.StatusLevelTypesInfo),
							Message:       to.Ptr("<message>"),
							Time:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
						},
					},
					Disks: []*armcompute.DiskInstanceView{
						{
							Name: to.Ptr("<name>"),
							EncryptionSettings: []*armcompute.DiskEncryptionSettings{
								{
									DiskEncryptionKey: &armcompute.KeyVaultSecretReference{
										SecretURL: to.Ptr("<secret-url>"),
										SourceVault: &armcompute.SubResource{
											ID: to.Ptr("<id>"),
										},
									},
									Enabled: to.Ptr(true),
									KeyEncryptionKey: &armcompute.KeyVaultKeyReference{
										KeyURL: to.Ptr("<key-url>"),
										SourceVault: &armcompute.SubResource{
											ID: to.Ptr("<id>"),
										},
									},
								}},
							Statuses: []*armcompute.InstanceViewStatus{
								{
									Code:          to.Ptr("<code>"),
									DisplayStatus: to.Ptr("<display-status>"),
									Level:         to.Ptr(armcompute.StatusLevelTypesInfo),
									Message:       to.Ptr("<message>"),
									Time:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
								}},
						}},
					MaintenanceRedeployStatus: &armcompute.MaintenanceRedeployStatus{
						IsCustomerInitiatedMaintenanceAllowed: to.Ptr(true),
						LastOperationMessage:                  to.Ptr("<last-operation-message>"),
						LastOperationResultCode:               to.Ptr(armcompute.MaintenanceOperationResultCodeTypesNone),
						MaintenanceWindowEndTime:              to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.531Z"); return t }()),
						MaintenanceWindowStartTime:            to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.531Z"); return t }()),
						PreMaintenanceWindowEndTime:           to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.531Z"); return t }()),
						PreMaintenanceWindowStartTime:         to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.531Z"); return t }()),
					},
					PlacementGroupID:     to.Ptr("<placement-group-id>"),
					PlatformFaultDomain:  to.Ptr[int32](14),
					PlatformUpdateDomain: to.Ptr[int32](23),
					RdpThumbPrint:        to.Ptr("<rdp-thumb-print>"),
					Statuses: []*armcompute.InstanceViewStatus{
						{
							Code:          to.Ptr("<code>"),
							DisplayStatus: to.Ptr("<display-status>"),
							Level:         to.Ptr(armcompute.StatusLevelTypesInfo),
							Message:       to.Ptr("<message>"),
							Time:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
						}},
					VMAgent: &armcompute.VirtualMachineAgentInstanceView{
						ExtensionHandlers: []*armcompute.VirtualMachineExtensionHandlerInstanceView{
							{
								Type: to.Ptr("<type>"),
								Status: &armcompute.InstanceViewStatus{
									Code:          to.Ptr("<code>"),
									DisplayStatus: to.Ptr("<display-status>"),
									Level:         to.Ptr(armcompute.StatusLevelTypesInfo),
									Message:       to.Ptr("<message>"),
									Time:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
								},
								TypeHandlerVersion: to.Ptr("<type-handler-version>"),
							}},
						Statuses: []*armcompute.InstanceViewStatus{
							{
								Code:          to.Ptr("<code>"),
								DisplayStatus: to.Ptr("<display-status>"),
								Level:         to.Ptr(armcompute.StatusLevelTypesInfo),
								Message:       to.Ptr("<message>"),
								Time:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
							}},
						VMAgentVersion: to.Ptr("<vmagent-version>"),
					},
					VMHealth: &armcompute.VirtualMachineHealthStatus{
						Status: &armcompute.InstanceViewStatus{
							Code:          to.Ptr("<code>"),
							DisplayStatus: to.Ptr("<display-status>"),
							Level:         to.Ptr(armcompute.StatusLevelTypesInfo),
							Message:       to.Ptr("<message>"),
							Time:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
						},
					},
					Extensions: []*armcompute.VirtualMachineExtensionInstanceView{
						{
							Name: to.Ptr("<name>"),
							Type: to.Ptr("<type>"),
							Statuses: []*armcompute.InstanceViewStatus{
								{
									Code:          to.Ptr("<code>"),
									DisplayStatus: to.Ptr("<display-status>"),
									Level:         to.Ptr(armcompute.StatusLevelTypesInfo),
									Message:       to.Ptr("<message>"),
									Time:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
								}},
							Substatuses: []*armcompute.InstanceViewStatus{
								{
									Code:          to.Ptr("<code>"),
									DisplayStatus: to.Ptr("<display-status>"),
									Level:         to.Ptr(armcompute.StatusLevelTypesInfo),
									Message:       to.Ptr("<message>"),
									Time:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
								}},
							TypeHandlerVersion: to.Ptr("<type-handler-version>"),
						}},
				},
				LicenseType: to.Ptr("<license-type>"),
				NetworkProfile: &armcompute.NetworkProfile{
					NetworkAPIVersion: to.Ptr(armcompute.NetworkAPIVersionTwoThousandTwenty1101),
					NetworkInterfaceConfigurations: []*armcompute.VirtualMachineNetworkInterfaceConfiguration{
						{
							Name: to.Ptr("<name>"),
							Properties: &armcompute.VirtualMachineNetworkInterfaceConfigurationProperties{
								DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
								DNSSettings: &armcompute.VirtualMachineNetworkInterfaceDNSSettingsConfiguration{
									DNSServers: []*string{
										to.Ptr("aaaaaa")},
								},
								DscpConfiguration: &armcompute.SubResource{
									ID: to.Ptr("<id>"),
								},
								EnableAcceleratedNetworking: to.Ptr(true),
								EnableFpga:                  to.Ptr(true),
								EnableIPForwarding:          to.Ptr(true),
								IPConfigurations: []*armcompute.VirtualMachineNetworkInterfaceIPConfiguration{
									{
										Name: to.Ptr("<name>"),
										Properties: &armcompute.VirtualMachineNetworkInterfaceIPConfigurationProperties{
											ApplicationGatewayBackendAddressPools: []*armcompute.SubResource{
												{
													ID: to.Ptr("<id>"),
												}},
											ApplicationSecurityGroups: []*armcompute.SubResource{
												{
													ID: to.Ptr("<id>"),
												}},
											LoadBalancerBackendAddressPools: []*armcompute.SubResource{
												{
													ID: to.Ptr("<id>"),
												}},
											Primary:                 to.Ptr(true),
											PrivateIPAddressVersion: to.Ptr(armcompute.IPVersionsIPv4),
											PublicIPAddressConfiguration: &armcompute.VirtualMachinePublicIPAddressConfiguration{
												Name: to.Ptr("<name>"),
												Properties: &armcompute.VirtualMachinePublicIPAddressConfigurationProperties{
													DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
													DNSSettings: &armcompute.VirtualMachinePublicIPAddressDNSSettingsConfiguration{
														DomainNameLabel: to.Ptr("<domain-name-label>"),
													},
													IdleTimeoutInMinutes: to.Ptr[int32](2),
													IPTags: []*armcompute.VirtualMachineIPTag{
														{
															IPTagType: to.Ptr("<iptag-type>"),
															Tag:       to.Ptr("<tag>"),
														}},
													PublicIPAddressVersion:   to.Ptr(armcompute.IPVersionsIPv4),
													PublicIPAllocationMethod: to.Ptr(armcompute.PublicIPAllocationMethodDynamic),
													PublicIPPrefix: &armcompute.SubResource{
														ID: to.Ptr("<id>"),
													},
												},
												SKU: &armcompute.PublicIPAddressSKU{
													Name: to.Ptr(armcompute.PublicIPAddressSKUNameBasic),
													Tier: to.Ptr(armcompute.PublicIPAddressSKUTierRegional),
												},
											},
											Subnet: &armcompute.SubResource{
												ID: to.Ptr("<id>"),
											},
										},
									}},
								NetworkSecurityGroup: &armcompute.SubResource{
									ID: to.Ptr("<id>"),
								},
								Primary: to.Ptr(true),
							},
						}},
					NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
						{
							ID: to.Ptr("<id>"),
							Properties: &armcompute.NetworkInterfaceReferenceProperties{
								DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
								Primary:      to.Ptr(true),
							},
						}},
				},
				NetworkProfileConfiguration: &armcompute.VirtualMachineScaleSetVMNetworkProfileConfiguration{
					NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
						{
							ID:   to.Ptr("<id>"),
							Name: to.Ptr("<name>"),
							Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
								DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
								DNSSettings: &armcompute.VirtualMachineScaleSetNetworkConfigurationDNSSettings{
									DNSServers: []*string{},
								},
								EnableAcceleratedNetworking: to.Ptr(true),
								EnableFpga:                  to.Ptr(true),
								EnableIPForwarding:          to.Ptr(true),
								IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
									{
										ID:   to.Ptr("<id>"),
										Name: to.Ptr("<name>"),
										Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
											ApplicationGatewayBackendAddressPools: []*armcompute.SubResource{
												{
													ID: to.Ptr("<id>"),
												}},
											ApplicationSecurityGroups: []*armcompute.SubResource{
												{
													ID: to.Ptr("<id>"),
												}},
											LoadBalancerBackendAddressPools: []*armcompute.SubResource{
												{
													ID: to.Ptr("<id>"),
												}},
											LoadBalancerInboundNatPools: []*armcompute.SubResource{
												{
													ID: to.Ptr("<id>"),
												}},
											Primary:                 to.Ptr(true),
											PrivateIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
											PublicIPAddressConfiguration: &armcompute.VirtualMachineScaleSetPublicIPAddressConfiguration{
												Name: to.Ptr("<name>"),
												Properties: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationProperties{
													DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
													DNSSettings: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationDNSSettings{
														DomainNameLabel: to.Ptr("<domain-name-label>"),
													},
													IdleTimeoutInMinutes: to.Ptr[int32](18),
													IPTags: []*armcompute.VirtualMachineScaleSetIPTag{
														{
															IPTagType: to.Ptr("<iptag-type>"),
															Tag:       to.Ptr("<tag>"),
														}},
													PublicIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
													PublicIPPrefix: &armcompute.SubResource{
														ID: to.Ptr("<id>"),
													},
												},
												SKU: &armcompute.PublicIPAddressSKU{
													Name: to.Ptr(armcompute.PublicIPAddressSKUNameBasic),
													Tier: to.Ptr(armcompute.PublicIPAddressSKUTierRegional),
												},
											},
											Subnet: &armcompute.APIEntityReference{
												ID: to.Ptr("<id>"),
											},
										},
									}},
								NetworkSecurityGroup: &armcompute.SubResource{
									ID: to.Ptr("<id>"),
								},
								Primary: to.Ptr(true),
							},
						}},
				},
				OSProfile: &armcompute.OSProfile{
					AdminPassword:            to.Ptr("<admin-password>"),
					AdminUsername:            to.Ptr("<admin-username>"),
					AllowExtensionOperations: to.Ptr(true),
					ComputerName:             to.Ptr("<computer-name>"),
					CustomData:               to.Ptr("<custom-data>"),
					LinuxConfiguration: &armcompute.LinuxConfiguration{
						DisablePasswordAuthentication: to.Ptr(true),
						PatchSettings: &armcompute.LinuxPatchSettings{
							AssessmentMode: to.Ptr(armcompute.LinuxPatchAssessmentModeImageDefault),
							PatchMode:      to.Ptr(armcompute.LinuxVMGuestPatchModeImageDefault),
						},
						ProvisionVMAgent: to.Ptr(true),
						SSH: &armcompute.SSHConfiguration{
							PublicKeys: []*armcompute.SSHPublicKey{
								{
									Path:    to.Ptr("<path>"),
									KeyData: to.Ptr("<key-data>"),
								}},
						},
					},
					RequireGuestProvisionSignal: to.Ptr(true),
					Secrets:                     []*armcompute.VaultSecretGroup{},
					WindowsConfiguration: &armcompute.WindowsConfiguration{
						AdditionalUnattendContent: []*armcompute.AdditionalUnattendContent{
							{
								ComponentName: to.Ptr("<component-name>"),
								Content:       to.Ptr("<content>"),
								PassName:      to.Ptr("<pass-name>"),
								SettingName:   to.Ptr(armcompute.SettingNamesAutoLogon),
							}},
						EnableAutomaticUpdates: to.Ptr(true),
						PatchSettings: &armcompute.PatchSettings{
							AssessmentMode:    to.Ptr(armcompute.WindowsPatchAssessmentModeImageDefault),
							EnableHotpatching: to.Ptr(true),
							PatchMode:         to.Ptr(armcompute.WindowsVMGuestPatchModeManual),
						},
						ProvisionVMAgent: to.Ptr(true),
						TimeZone:         to.Ptr("<time-zone>"),
						WinRM: &armcompute.WinRMConfiguration{
							Listeners: []*armcompute.WinRMListener{
								{
									CertificateURL: to.Ptr("<certificate-url>"),
									Protocol:       to.Ptr(armcompute.ProtocolTypesHTTP),
								}},
						},
					},
				},
				ProtectionPolicy: &armcompute.VirtualMachineScaleSetVMProtectionPolicy{
					ProtectFromScaleIn:         to.Ptr(true),
					ProtectFromScaleSetActions: to.Ptr(true),
				},
				SecurityProfile: &armcompute.SecurityProfile{
					EncryptionAtHost: to.Ptr(true),
					SecurityType:     to.Ptr(armcompute.SecurityTypesTrustedLaunch),
					UefiSettings: &armcompute.UefiSettings{
						SecureBootEnabled: to.Ptr(true),
						VTpmEnabled:       to.Ptr(true),
					},
				},
				StorageProfile: &armcompute.StorageProfile{
					DataDisks: []*armcompute.DataDisk{
						{
							Name:         to.Ptr("<name>"),
							Caching:      to.Ptr(armcompute.CachingTypesNone),
							CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesEmpty),
							DeleteOption: to.Ptr(armcompute.DiskDeleteOptionTypesDelete),
							DetachOption: to.Ptr(armcompute.DiskDetachOptionTypesForceDetach),
							DiskSizeGB:   to.Ptr[int32](128),
							Image: &armcompute.VirtualHardDisk{
								URI: to.Ptr("<uri>"),
							},
							Lun: to.Ptr[int32](1),
							ManagedDisk: &armcompute.ManagedDiskParameters{
								ID: to.Ptr("<id>"),
								DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
									ID: to.Ptr("<id>"),
								},
								StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
							},
							ToBeDetached: to.Ptr(true),
							Vhd: &armcompute.VirtualHardDisk{
								URI: to.Ptr("<uri>"),
							},
							WriteAcceleratorEnabled: to.Ptr(true),
						}},
					ImageReference: &armcompute.ImageReference{
						ID:                   to.Ptr("<id>"),
						Offer:                to.Ptr("<offer>"),
						Publisher:            to.Ptr("<publisher>"),
						SharedGalleryImageID: to.Ptr("<shared-gallery-image-id>"),
						SKU:                  to.Ptr("<sku>"),
						Version:              to.Ptr("<version>"),
					},
					OSDisk: &armcompute.OSDisk{
						Name:         to.Ptr("<name>"),
						Caching:      to.Ptr(armcompute.CachingTypesNone),
						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
						DeleteOption: to.Ptr(armcompute.DiskDeleteOptionTypesDelete),
						DiffDiskSettings: &armcompute.DiffDiskSettings{
							Option:    to.Ptr(armcompute.DiffDiskOptionsLocal),
							Placement: to.Ptr(armcompute.DiffDiskPlacementCacheDisk),
						},
						DiskSizeGB: to.Ptr[int32](127),
						EncryptionSettings: &armcompute.DiskEncryptionSettings{
							DiskEncryptionKey: &armcompute.KeyVaultSecretReference{
								SecretURL: to.Ptr("<secret-url>"),
								SourceVault: &armcompute.SubResource{
									ID: to.Ptr("<id>"),
								},
							},
							Enabled: to.Ptr(true),
							KeyEncryptionKey: &armcompute.KeyVaultKeyReference{
								KeyURL: to.Ptr("<key-url>"),
								SourceVault: &armcompute.SubResource{
									ID: to.Ptr("<id>"),
								},
							},
						},
						Image: &armcompute.VirtualHardDisk{
							URI: to.Ptr("<uri>"),
						},
						ManagedDisk: &armcompute.ManagedDiskParameters{
							ID: to.Ptr("<id>"),
							DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
								ID: to.Ptr("<id>"),
							},
							StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
						},
						OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
						Vhd: &armcompute.VirtualHardDisk{
							URI: to.Ptr("<uri>"),
						},
						WriteAcceleratorEnabled: to.Ptr(true),
					},
				},
				UserData: to.Ptr("<user-data>"),
			},
			SKU: &armcompute.SKU{
				Name:     to.Ptr("<name>"),
				Capacity: to.Ptr[int64](29),
				Tier:     to.Ptr("<tier>"),
			},
		},
		&armcompute.VirtualMachineScaleSetVMsClientBeginUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/ForceDeleteVirtualMachineScaleSetVM.json
func ExampleVirtualMachineScaleSetVMsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetVMsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		"<instance-id>",
		&armcompute.VirtualMachineScaleSetVMsClientBeginDeleteOptions{ForceDeletion: to.Ptr(true),
			ResumeToken: "",
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/GetVirtualMachineScaleSetVMWithUserData.json
func ExampleVirtualMachineScaleSetVMsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetVMsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		"<instance-id>",
		&armcompute.VirtualMachineScaleSetVMsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/GetVirtualMachineScaleSetVMInstanceViewAutoPlacedOnDedicatedHostGroup.json
func ExampleVirtualMachineScaleSetVMsClient_GetInstanceView() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetVMsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.GetInstanceView(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		"<instance-id>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSetVMs_List_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetVMsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetVMsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.NewListPager("<resource-group-name>",
		"<virtual-machine-scale-set-name>",
		&armcompute.VirtualMachineScaleSetVMsClientListOptions{Filter: to.Ptr("<filter>"),
			Select: to.Ptr("<select>"),
			Expand: to.Ptr("<expand>"),
		})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSetVMs_PowerOff_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetVMsClient_BeginPowerOff() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetVMsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginPowerOff(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		"<instance-id>",
		&armcompute.VirtualMachineScaleSetVMsClientBeginPowerOffOptions{SkipShutdown: to.Ptr(true),
			ResumeToken: "",
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSetVMs_Restart_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetVMsClient_BeginRestart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetVMsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginRestart(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		"<instance-id>",
		&armcompute.VirtualMachineScaleSetVMsClientBeginRestartOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSetVMs_Start_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetVMsClient_BeginStart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetVMsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginStart(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		"<instance-id>",
		&armcompute.VirtualMachineScaleSetVMsClientBeginStartOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSetVMs_Redeploy_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetVMsClient_BeginRedeploy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetVMsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginRedeploy(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		"<instance-id>",
		&armcompute.VirtualMachineScaleSetVMsClientBeginRedeployOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/RetrieveBootDiagnosticsDataVMScaleSetVM.json
func ExampleVirtualMachineScaleSetVMsClient_RetrieveBootDiagnosticsData() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetVMsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.RetrieveBootDiagnosticsData(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		"<instance-id>",
		&armcompute.VirtualMachineScaleSetVMsClientRetrieveBootDiagnosticsDataOptions{SasURIExpirationTimeInMinutes: to.Ptr[int32](60)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSetVMs_PerformMaintenance_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetVMsClient_BeginPerformMaintenance() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetVMsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginPerformMaintenance(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		"<instance-id>",
		&armcompute.VirtualMachineScaleSetVMsClientBeginPerformMaintenanceOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/SimulateEvictionOfVmssVM.json
func ExampleVirtualMachineScaleSetVMsClient_SimulateEviction() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetVMsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.SimulateEviction(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		"<instance-id>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/runCommands/VMScaleSetRunCommand.json
func ExampleVirtualMachineScaleSetVMsClient_BeginRunCommand() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetVMsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginRunCommand(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		"<instance-id>",
		armcompute.RunCommandInput{
			CommandID: to.Ptr("<command-id>"),
			Script: []*string{
				to.Ptr("Write-Host Hello World!")},
		},
		&armcompute.VirtualMachineScaleSetVMsClientBeginRunCommandOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
