//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package appservice_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/ListPublishingCredentialsPolicies.json
func ExampleWebAppsClient_ListBasicPublishingCredentialsPolicies() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	pager := client.ListBasicPublishingCredentialsPolicies("<resource-group-name>",
		"<name>",
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

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetPublishingCredentialsPolicy.json
func ExampleWebAppsClient_GetFtpAllowed() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	res, err := client.GetFtpAllowed(ctx,
		"<resource-group-name>",
		"<name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebAppsClientGetFtpAllowedResult)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/UpdatePublishingCredentialsPolicy.json
func ExampleWebAppsClient_UpdateFtpAllowed() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	res, err := client.UpdateFtpAllowed(ctx,
		"<resource-group-name>",
		"<name>",
		appservice.CsmPublishingCredentialsPoliciesEntity{
			Properties: &appservice.CsmPublishingCredentialsPoliciesEntityProperties{
				Allow: to.BoolPtr(true),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebAppsClientUpdateFtpAllowedResult)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetPublishingCredentialsPolicy.json
func ExampleWebAppsClient_GetScmAllowed() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	res, err := client.GetScmAllowed(ctx,
		"<resource-group-name>",
		"<name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebAppsClientGetScmAllowedResult)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/UpdatePublishingCredentialsPolicy.json
func ExampleWebAppsClient_UpdateScmAllowed() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	res, err := client.UpdateScmAllowed(ctx,
		"<resource-group-name>",
		"<name>",
		appservice.CsmPublishingCredentialsPoliciesEntity{
			Properties: &appservice.CsmPublishingCredentialsPoliciesEntityProperties{
				Allow: to.BoolPtr(true),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebAppsClientUpdateScmAllowedResult)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetKeyVaultReferencesForAppSettings.json
func ExampleWebAppsClient_GetAppSettingsKeyVaultReferences() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	pager := client.GetAppSettingsKeyVaultReferences("<resource-group-name>",
		"<name>",
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

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetKeyVaultReferencesForAppSetting.json
func ExampleWebAppsClient_GetAppSettingKeyVaultReference() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	res, err := client.GetAppSettingKeyVaultReference(ctx,
		"<resource-group-name>",
		"<name>",
		"<app-setting-key>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebAppsClientGetAppSettingKeyVaultReferenceResult)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetSiteInstanceInfo.json
func ExampleWebAppsClient_GetInstanceInfo() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	res, err := client.GetInstanceInfo(ctx,
		"<resource-group-name>",
		"<name>",
		"<instance-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebAppsClientGetInstanceInfoResult)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/ListSlotBackups.json
func ExampleWebAppsClient_ListSiteBackups() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	pager := client.ListSiteBackups("<resource-group-name>",
		"<name>",
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

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetWebSiteNetworkTraceOperation.json
func ExampleWebAppsClient_GetNetworkTraceOperation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	res, err := client.GetNetworkTraceOperation(ctx,
		"<resource-group-name>",
		"<name>",
		"<operation-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebAppsClientGetNetworkTraceOperationResult)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/StartWebSiteNetworkTraceOperation.json
func ExampleWebAppsClient_BeginStartWebSiteNetworkTraceOperation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginStartWebSiteNetworkTraceOperation(ctx,
		"<resource-group-name>",
		"<name>",
		&appservice.WebAppsClientBeginStartWebSiteNetworkTraceOperationOptions{DurationInSeconds: to.Int32Ptr(60),
			MaxFrameLength: nil,
			SasURL:         nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebAppsClientStartWebSiteNetworkTraceOperationResult)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/StopWebSiteNetworkTrace.json
func ExampleWebAppsClient_StopWebSiteNetworkTrace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	_, err = client.StopWebSiteNetworkTrace(ctx,
		"<resource-group-name>",
		"<name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetWebSiteNetworkTraces.json
func ExampleWebAppsClient_GetNetworkTraces() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	res, err := client.GetNetworkTraces(ctx,
		"<resource-group-name>",
		"<name>",
		"<operation-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebAppsClientGetNetworkTracesResult)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetWebSiteNetworkTraceOperation.json
func ExampleWebAppsClient_GetNetworkTraceOperationV2() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	res, err := client.GetNetworkTraceOperationV2(ctx,
		"<resource-group-name>",
		"<name>",
		"<operation-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebAppsClientGetNetworkTraceOperationV2Result)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetWebSiteNetworkTraces.json
func ExampleWebAppsClient_GetNetworkTracesV2() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	res, err := client.GetNetworkTracesV2(ctx,
		"<resource-group-name>",
		"<name>",
		"<operation-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebAppsClientGetNetworkTracesV2Result)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetSitePrivateEndpointConnection.json
func ExampleWebAppsClient_GetPrivateEndpointConnection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	res, err := client.GetPrivateEndpointConnection(ctx,
		"<resource-group-name>",
		"<name>",
		"<private-endpoint-connection-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebAppsClientGetPrivateEndpointConnectionResult)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/ApproveRejectSitePrivateEndpointConnection.json
func ExampleWebAppsClient_BeginApproveOrRejectPrivateEndpointConnection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginApproveOrRejectPrivateEndpointConnection(ctx,
		"<resource-group-name>",
		"<name>",
		"<private-endpoint-connection-name>",
		appservice.PrivateLinkConnectionApprovalRequestResource{
			Properties: &appservice.PrivateLinkConnectionApprovalRequest{
				PrivateLinkServiceConnectionState: &appservice.PrivateLinkConnectionState{
					Description:     to.StringPtr("<description>"),
					ActionsRequired: to.StringPtr("<actions-required>"),
					Status:          to.StringPtr("<status>"),
				},
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
	log.Printf("Response result: %#v\n", res.WebAppsClientApproveOrRejectPrivateEndpointConnectionResult)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/DeleteSitePrivateEndpointConnection.json
func ExampleWebAppsClient_BeginDeletePrivateEndpointConnection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDeletePrivateEndpointConnection(ctx,
		"<resource-group-name>",
		"<name>",
		"<private-endpoint-connection-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebAppsClientDeletePrivateEndpointConnectionResult)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetSitePrivateLinkResources.json
func ExampleWebAppsClient_GetPrivateLinkResources() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	res, err := client.GetPrivateLinkResources(ctx,
		"<resource-group-name>",
		"<name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebAppsClientGetPrivateLinkResourcesResult)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/ListPublishingCredentialsPoliciesSlot.json
func ExampleWebAppsClient_ListBasicPublishingCredentialsPoliciesSlot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	pager := client.ListBasicPublishingCredentialsPoliciesSlot("<resource-group-name>",
		"<name>",
		"<slot>",
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

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetPublishingCredentialsPolicySlot.json
func ExampleWebAppsClient_GetFtpAllowedSlot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	res, err := client.GetFtpAllowedSlot(ctx,
		"<resource-group-name>",
		"<name>",
		"<slot>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebAppsClientGetFtpAllowedSlotResult)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/UpdatePublishingCredentialsPolicySlot.json
func ExampleWebAppsClient_UpdateFtpAllowedSlot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	res, err := client.UpdateFtpAllowedSlot(ctx,
		"<resource-group-name>",
		"<name>",
		"<slot>",
		appservice.CsmPublishingCredentialsPoliciesEntity{
			Properties: &appservice.CsmPublishingCredentialsPoliciesEntityProperties{
				Allow: to.BoolPtr(true),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebAppsClientUpdateFtpAllowedSlotResult)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetPublishingCredentialsPolicySlot.json
func ExampleWebAppsClient_GetScmAllowedSlot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	res, err := client.GetScmAllowedSlot(ctx,
		"<resource-group-name>",
		"<name>",
		"<slot>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebAppsClientGetScmAllowedSlotResult)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/UpdatePublishingCredentialsPolicySlot.json
func ExampleWebAppsClient_UpdateScmAllowedSlot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	res, err := client.UpdateScmAllowedSlot(ctx,
		"<resource-group-name>",
		"<name>",
		"<slot>",
		appservice.CsmPublishingCredentialsPoliciesEntity{
			Properties: &appservice.CsmPublishingCredentialsPoliciesEntityProperties{
				Allow: to.BoolPtr(true),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebAppsClientUpdateScmAllowedSlotResult)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetKeyVaultReferencesForAppSettingsSlot.json
func ExampleWebAppsClient_GetAppSettingsKeyVaultReferencesSlot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	pager := client.GetAppSettingsKeyVaultReferencesSlot("<resource-group-name>",
		"<name>",
		"<slot>",
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

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetKeyVaultReferencesForAppSettingSlot.json
func ExampleWebAppsClient_GetAppSettingKeyVaultReferenceSlot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	res, err := client.GetAppSettingKeyVaultReferenceSlot(ctx,
		"<resource-group-name>",
		"<name>",
		"<app-setting-key>",
		"<slot>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebAppsClientGetAppSettingKeyVaultReferenceSlotResult)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetSiteInstanceInfo.json
func ExampleWebAppsClient_GetInstanceInfoSlot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	res, err := client.GetInstanceInfoSlot(ctx,
		"<resource-group-name>",
		"<name>",
		"<instance-id>",
		"<slot>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebAppsClientGetInstanceInfoSlotResult)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/ListSlotBackups.json
func ExampleWebAppsClient_ListSiteBackupsSlot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	pager := client.ListSiteBackupsSlot("<resource-group-name>",
		"<name>",
		"<slot>",
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

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetWebSiteNetworkTraceOperation.json
func ExampleWebAppsClient_GetNetworkTraceOperationSlot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	res, err := client.GetNetworkTraceOperationSlot(ctx,
		"<resource-group-name>",
		"<name>",
		"<operation-id>",
		"<slot>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebAppsClientGetNetworkTraceOperationSlotResult)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/StartWebSiteNetworkTraceOperation.json
func ExampleWebAppsClient_BeginStartWebSiteNetworkTraceOperationSlot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginStartWebSiteNetworkTraceOperationSlot(ctx,
		"<resource-group-name>",
		"<name>",
		"<slot>",
		&appservice.WebAppsClientBeginStartWebSiteNetworkTraceOperationSlotOptions{DurationInSeconds: to.Int32Ptr(60),
			MaxFrameLength: nil,
			SasURL:         nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebAppsClientStartWebSiteNetworkTraceOperationSlotResult)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/StopWebSiteNetworkTrace.json
func ExampleWebAppsClient_StopWebSiteNetworkTraceSlot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	_, err = client.StopWebSiteNetworkTraceSlot(ctx,
		"<resource-group-name>",
		"<name>",
		"<slot>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetWebSiteNetworkTraces.json
func ExampleWebAppsClient_GetNetworkTracesSlot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	res, err := client.GetNetworkTracesSlot(ctx,
		"<resource-group-name>",
		"<name>",
		"<operation-id>",
		"<slot>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebAppsClientGetNetworkTracesSlotResult)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetWebSiteNetworkTraceOperation.json
func ExampleWebAppsClient_GetNetworkTraceOperationSlotV2() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	res, err := client.GetNetworkTraceOperationSlotV2(ctx,
		"<resource-group-name>",
		"<name>",
		"<operation-id>",
		"<slot>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebAppsClientGetNetworkTraceOperationSlotV2Result)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetWebSiteNetworkTraces.json
func ExampleWebAppsClient_GetNetworkTracesSlotV2() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	res, err := client.GetNetworkTracesSlotV2(ctx,
		"<resource-group-name>",
		"<name>",
		"<operation-id>",
		"<slot>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebAppsClientGetNetworkTracesSlotV2Result)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetSitePrivateEndpointConnectionSlot.json
func ExampleWebAppsClient_GetPrivateEndpointConnectionSlot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	res, err := client.GetPrivateEndpointConnectionSlot(ctx,
		"<resource-group-name>",
		"<name>",
		"<private-endpoint-connection-name>",
		"<slot>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebAppsClientGetPrivateEndpointConnectionSlotResult)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/ApproveRejectSitePrivateEndpointConnectionSlot.json
func ExampleWebAppsClient_BeginApproveOrRejectPrivateEndpointConnectionSlot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginApproveOrRejectPrivateEndpointConnectionSlot(ctx,
		"<resource-group-name>",
		"<name>",
		"<private-endpoint-connection-name>",
		"<slot>",
		appservice.PrivateLinkConnectionApprovalRequestResource{
			Properties: &appservice.PrivateLinkConnectionApprovalRequest{
				PrivateLinkServiceConnectionState: &appservice.PrivateLinkConnectionState{
					Description:     to.StringPtr("<description>"),
					ActionsRequired: to.StringPtr("<actions-required>"),
					Status:          to.StringPtr("<status>"),
				},
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
	log.Printf("Response result: %#v\n", res.WebAppsClientApproveOrRejectPrivateEndpointConnectionSlotResult)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/DeleteSitePrivateEndpointConnectionSlot.json
func ExampleWebAppsClient_BeginDeletePrivateEndpointConnectionSlot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDeletePrivateEndpointConnectionSlot(ctx,
		"<resource-group-name>",
		"<name>",
		"<private-endpoint-connection-name>",
		"<slot>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebAppsClientDeletePrivateEndpointConnectionSlotResult)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetSitePrivateLinkResourcesSlot.json
func ExampleWebAppsClient_GetPrivateLinkResourcesSlot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	res, err := client.GetPrivateLinkResourcesSlot(ctx,
		"<resource-group-name>",
		"<name>",
		"<slot>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebAppsClientGetPrivateLinkResourcesSlotResult)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/StartWebSiteNetworkTraceOperation.json
func ExampleWebAppsClient_BeginStartNetworkTraceSlot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginStartNetworkTraceSlot(ctx,
		"<resource-group-name>",
		"<name>",
		"<slot>",
		&appservice.WebAppsClientBeginStartNetworkTraceSlotOptions{DurationInSeconds: to.Int32Ptr(60),
			MaxFrameLength: nil,
			SasURL:         nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebAppsClientStartNetworkTraceSlotResult)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/StopWebSiteNetworkTrace.json
func ExampleWebAppsClient_StopNetworkTraceSlot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	_, err = client.StopNetworkTraceSlot(ctx,
		"<resource-group-name>",
		"<name>",
		"<slot>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/StartWebSiteNetworkTraceOperation.json
func ExampleWebAppsClient_BeginStartNetworkTrace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginStartNetworkTrace(ctx,
		"<resource-group-name>",
		"<name>",
		&appservice.WebAppsClientBeginStartNetworkTraceOptions{DurationInSeconds: to.Int32Ptr(60),
			MaxFrameLength: nil,
			SasURL:         nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebAppsClientStartNetworkTraceResult)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/StopWebSiteNetworkTrace.json
func ExampleWebAppsClient_StopNetworkTrace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := appservice.NewWebAppsClient("<subscription-id>", cred, nil)
	_, err = client.StopNetworkTrace(ctx,
		"<resource-group-name>",
		"<name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
