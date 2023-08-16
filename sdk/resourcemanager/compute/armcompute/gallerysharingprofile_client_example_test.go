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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-03-03/examples/galleryExamples/Gallery_AddToSharingProfile.json
func ExampleGallerySharingProfileClient_BeginUpdate_addSharingIdToTheSharingProfileOfAGallery() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGallerySharingProfileClient().BeginUpdate(ctx, "myResourceGroup", "myGalleryName", armcompute.SharingUpdate{
		Groups: []*armcompute.SharingProfileGroup{
			{
				Type: to.Ptr(armcompute.SharingProfileGroupTypesSubscriptions),
				IDs: []*string{
					to.Ptr("34a4ab42-0d72-47d9-bd1a-aed207386dac"),
					to.Ptr("380fd389-260b-41aa-bad9-0a83108c370b")},
			},
			{
				Type: to.Ptr(armcompute.SharingProfileGroupTypesAADTenants),
				IDs: []*string{
					to.Ptr("c24c76aa-8897-4027-9b03-8f7928b54ff6")},
			}},
		OperationType: to.Ptr(armcompute.SharingUpdateOperationTypesAdd),
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
	// res.SharingUpdate = armcompute.SharingUpdate{
	// 	Groups: []*armcompute.SharingProfileGroup{
	// 		{
	// 			Type: to.Ptr(armcompute.SharingProfileGroupTypesSubscriptions),
	// 			IDs: []*string{
	// 				to.Ptr("34a4ab42-0d72-47d9-bd1a-aed207386dac"),
	// 				to.Ptr("380fd389-260b-41aa-bad9-0a83108c370b")},
	// 			},
	// 			{
	// 				Type: to.Ptr(armcompute.SharingProfileGroupTypesAADTenants),
	// 				IDs: []*string{
	// 					to.Ptr("c24c76aa-8897-4027-9b03-8f7928b54ff6")},
	// 			}},
	// 			OperationType: to.Ptr(armcompute.SharingUpdateOperationTypesAdd),
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-03-03/examples/galleryExamples/Gallery_ResetSharingProfile.json
func ExampleGallerySharingProfileClient_BeginUpdate_resetSharingProfileOfAGallery() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGallerySharingProfileClient().BeginUpdate(ctx, "myResourceGroup", "myGalleryName", armcompute.SharingUpdate{
		OperationType: to.Ptr(armcompute.SharingUpdateOperationTypesReset),
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
	// res.SharingUpdate = armcompute.SharingUpdate{
	// 	OperationType: to.Ptr(armcompute.SharingUpdateOperationTypesReset),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-03-03/examples/galleryExamples/Gallery_EnableCommunityGallery.json
func ExampleGallerySharingProfileClient_BeginUpdate_shareAGalleryToCommunity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGallerySharingProfileClient().BeginUpdate(ctx, "myResourceGroup", "myGalleryName", armcompute.SharingUpdate{
		OperationType: to.Ptr(armcompute.SharingUpdateOperationTypesEnableCommunity),
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
	// res.SharingUpdate = armcompute.SharingUpdate{
	// 	OperationType: to.Ptr(armcompute.SharingUpdateOperationTypesEnableCommunity),
	// }
}
