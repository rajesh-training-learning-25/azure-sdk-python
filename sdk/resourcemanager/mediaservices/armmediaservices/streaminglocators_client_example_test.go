//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmediaservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-locators-list.json
func ExampleStreamingLocatorsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStreamingLocatorsClient().NewListPager("contoso", "contosomedia", &armmediaservices.StreamingLocatorsClientListOptions{Filter: nil,
		Top:     nil,
		Orderby: nil,
	})
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
		// page.StreamingLocatorCollection = armmediaservices.StreamingLocatorCollection{
		// 	Value: []*armmediaservices.StreamingLocator{
		// 		{
		// 			Name: to.Ptr("clearStreamingLocator"),
		// 			Type: to.Ptr("Microsoft.Media/mediaservices/streamingLocators"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/streamingLocators/clearStreamingLocator"),
		// 			Properties: &armmediaservices.StreamingLocatorProperties{
		// 				AssetName: to.Ptr("ClimbingMountRainier"),
		// 				ContentKeys: []*armmediaservices.StreamingLocatorContentKey{
		// 				},
		// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-08T18:29:31.934Z"); return t}()),
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "9999-12-31T23:59:59.999Z"); return t}()),
		// 				StreamingLocatorID: to.Ptr("6a116ec6-0c85-441f-9c31-89a5bc3adf0a"),
		// 				StreamingPolicyName: to.Ptr("clearStreamingPolicy"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("secureStreamingLocator"),
		// 			Type: to.Ptr("Microsoft.Media/mediaservices/streamingLocators"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/streamingLocators/secureStreamingLocator"),
		// 			Properties: &armmediaservices.StreamingLocatorProperties{
		// 				AssetName: to.Ptr("ClimbingMountRainier"),
		// 				ContentKeys: []*armmediaservices.StreamingLocatorContentKey{
		// 				},
		// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-08T18:29:31.954Z"); return t}()),
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "9999-12-31T23:59:59.999Z"); return t}()),
		// 				StreamingLocatorID: to.Ptr("7338ef90-ffc8-42de-8bff-de8f99973300"),
		// 				StreamingPolicyName: to.Ptr("secureStreamingPolicy"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-locators-get-by-name.json
func ExampleStreamingLocatorsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStreamingLocatorsClient().Get(ctx, "contoso", "contosomedia", "clearStreamingLocator", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StreamingLocator = armmediaservices.StreamingLocator{
	// 	Name: to.Ptr("clearStreamingLocator"),
	// 	Type: to.Ptr("Microsoft.Media/mediaservices/streamingLocators"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/streamingLocators/clearStreamingLocator"),
	// 	Properties: &armmediaservices.StreamingLocatorProperties{
	// 		AssetName: to.Ptr("ClimbingMountRainier"),
	// 		ContentKeys: []*armmediaservices.StreamingLocatorContentKey{
	// 		},
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-08T18:29:32.115Z"); return t}()),
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "9999-12-31T23:59:59.999Z"); return t}()),
	// 		StreamingLocatorID: to.Ptr("7684043b-f6d1-44a7-8bed-8a4aa474c5a4"),
	// 		StreamingPolicyName: to.Ptr("clearStreamingPolicy"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-locators-create-clear.json
func ExampleStreamingLocatorsClient_Create_createsAStreamingLocatorWithClearStreaming() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewStreamingLocatorsClient().Create(ctx, "contoso", "contosomedia", "UserCreatedClearStreamingLocator", armmediaservices.StreamingLocator{
		Properties: &armmediaservices.StreamingLocatorProperties{
			AssetName:           to.Ptr("ClimbingMountRainier"),
			StreamingPolicyName: to.Ptr("clearStreamingPolicy"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-locators-create-secure.json
func ExampleStreamingLocatorsClient_Create_createsAStreamingLocatorWithSecureStreaming() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewStreamingLocatorsClient().Create(ctx, "contoso", "contosomedia", "UserCreatedSecureStreamingLocator", armmediaservices.StreamingLocator{
		Properties: &armmediaservices.StreamingLocatorProperties{
			AssetName:           to.Ptr("ClimbingMountRainier"),
			EndTime:             to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2028-12-31T23:59:59.999Z"); return t }()),
			StartTime:           to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-01T00:00:00.000Z"); return t }()),
			StreamingPolicyName: to.Ptr("secureStreamingPolicy"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-locators-create-secure-userDefinedContentKeys.json
func ExampleStreamingLocatorsClient_Create_createsAStreamingLocatorWithUserDefinedContentKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewStreamingLocatorsClient().Create(ctx, "contoso", "contosomedia", "UserCreatedSecureStreamingLocatorWithUserDefinedContentKeys", armmediaservices.StreamingLocator{
		Properties: &armmediaservices.StreamingLocatorProperties{
			AssetName: to.Ptr("ClimbingMountRainier"),
			ContentKeys: []*armmediaservices.StreamingLocatorContentKey{
				{
					ID:                              to.Ptr("60000000-0000-0000-0000-000000000001"),
					LabelReferenceInStreamingPolicy: to.Ptr("aesDefaultKey"),
					Value:                           to.Ptr("1UqLohAfWsEGkULYxHjYZg=="),
				},
				{
					ID:                              to.Ptr("60000000-0000-0000-0000-000000000004"),
					LabelReferenceInStreamingPolicy: to.Ptr("cencDefaultKey"),
					Value:                           to.Ptr("4UqLohAfWsEGkULYxHjYZg=="),
				},
				{
					ID:                              to.Ptr("60000000-0000-0000-0000-000000000007"),
					LabelReferenceInStreamingPolicy: to.Ptr("cbcsDefaultKey"),
					Value:                           to.Ptr("7UqLohAfWsEGkULYxHjYZg=="),
				}},
			StreamingLocatorID:  to.Ptr("90000000-0000-0000-0000-00000000000A"),
			StreamingPolicyName: to.Ptr("secureStreamingPolicy"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-locators-delete.json
func ExampleStreamingLocatorsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewStreamingLocatorsClient().Delete(ctx, "contoso", "contosomedia", "clearStreamingLocator", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-locators-list-content-keys.json
func ExampleStreamingLocatorsClient_ListContentKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStreamingLocatorsClient().ListContentKeys(ctx, "contoso", "contosomedia", "secureStreamingLocator", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ListContentKeysResponse = armmediaservices.ListContentKeysResponse{
	// 	ContentKeys: []*armmediaservices.StreamingLocatorContentKey{
	// 		{
	// 			Type: to.Ptr(armmediaservices.StreamingLocatorContentKeyTypeEnvelopeEncryption),
	// 			ID: to.Ptr("9259eb06-eeee-4f77-987f-48f4ea5c649f"),
	// 			LabelReferenceInStreamingPolicy: to.Ptr("aesDefaultKey"),
	// 			PolicyName: to.Ptr("PolicyWithMultipleOptions"),
	// 			Tracks: []*armmediaservices.TrackSelection{
	// 			},
	// 			Value: to.Ptr("QpiqeQROdN5xamnfUF2Wdw=="),
	// 		},
	// 		{
	// 			Type: to.Ptr(armmediaservices.StreamingLocatorContentKeyTypeCommonEncryptionCenc),
	// 			ID: to.Ptr("06bfeff1-2bb6-4f58-af27-a2767f058bca"),
	// 			LabelReferenceInStreamingPolicy: to.Ptr("cencDefaultKey"),
	// 			PolicyName: to.Ptr("PolicyWithMultipleOptions"),
	// 			Tracks: []*armmediaservices.TrackSelection{
	// 			},
	// 			Value: to.Ptr("ZjgWhNnqnqcov/h+wrYusw=="),
	// 		},
	// 		{
	// 			Type: to.Ptr(armmediaservices.StreamingLocatorContentKeyTypeCommonEncryptionCbcs),
	// 			ID: to.Ptr("799e78a0-ed6f-4179-9222-ed4ec4223cec"),
	// 			LabelReferenceInStreamingPolicy: to.Ptr("cbcsDefaultKey"),
	// 			PolicyName: to.Ptr("PolicyWithMultipleOptions"),
	// 			Tracks: []*armmediaservices.TrackSelection{
	// 			},
	// 			Value: to.Ptr("FjZ3n3yRcVxRFftdYFbe9g=="),
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-locators-list-paths-streaming-and-download.json
func ExampleStreamingLocatorsClient_ListPaths_listPathsWhichHasStreamingPathsAndDownloadPaths() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStreamingLocatorsClient().ListPaths(ctx, "contoso", "contosomedia", "clearStreamingLocator", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ListPathsResponse = armmediaservices.ListPathsResponse{
	// 	DownloadPaths: []*string{
	// 		to.Ptr("/262a87b6-b538-4657-bac1-b6897924471d/textTrack.vtt"),
	// 		to.Ptr("/262a87b6-b538-4657-bac1-b6897924471d/video1.mp4"),
	// 		to.Ptr("/262a87b6-b538-4657-bac1-b6897924471d/video2.mp4"),
	// 		to.Ptr("/262a87b6-b538-4657-bac1-b6897924471d/video3.mp4")},
	// 		StreamingPaths: []*armmediaservices.StreamingPath{
	// 			{
	// 				EncryptionScheme: to.Ptr(armmediaservices.EncryptionSchemeNoEncryption),
	// 				Paths: []*string{
	// 					to.Ptr("/262a87b6-b538-4657-bac1-b6897924471d/videoManifest1.ism/manifest(format=m3u8-aapl)"),
	// 					to.Ptr("/262a87b6-b538-4657-bac1-b6897924471d/videoManifest2.ism/manifest(format=m3u8-aapl)"),
	// 					to.Ptr("/262a87b6-b538-4657-bac1-b6897924471d/videoManifest3.ism/manifest(format=m3u8-aapl)")},
	// 					StreamingProtocol: to.Ptr(armmediaservices.StreamingPolicyStreamingProtocolHls),
	// 				},
	// 				{
	// 					EncryptionScheme: to.Ptr(armmediaservices.EncryptionSchemeNoEncryption),
	// 					Paths: []*string{
	// 						to.Ptr("/262a87b6-b538-4657-bac1-b6897924471d/videoManifest1.ism/manifest(format=mpd-time-csf)"),
	// 						to.Ptr("/262a87b6-b538-4657-bac1-b6897924471d/videoManifest2.ism/manifest(format=mpd-time-csf)"),
	// 						to.Ptr("/262a87b6-b538-4657-bac1-b6897924471d/videoManifest3.ism/manifest(format=mpd-time-csf)")},
	// 						StreamingProtocol: to.Ptr(armmediaservices.StreamingPolicyStreamingProtocolDash),
	// 					},
	// 					{
	// 						EncryptionScheme: to.Ptr(armmediaservices.EncryptionSchemeNoEncryption),
	// 						Paths: []*string{
	// 							to.Ptr("/262a87b6-b538-4657-bac1-b6897924471d/videoManifest1.ism/manifest"),
	// 							to.Ptr("/262a87b6-b538-4657-bac1-b6897924471d/videoManifest2.ism/manifest"),
	// 							to.Ptr("/262a87b6-b538-4657-bac1-b6897924471d/videoManifest3.ism/manifest")},
	// 							StreamingProtocol: to.Ptr(armmediaservices.StreamingPolicyStreamingProtocolSmoothStreaming),
	// 					}},
	// 				}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-locators-list-paths-streaming-only.json
func ExampleStreamingLocatorsClient_ListPaths_listPathsWhichHasStreamingPathsOnly() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStreamingLocatorsClient().ListPaths(ctx, "contoso", "contosomedia", "secureStreamingLocator", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ListPathsResponse = armmediaservices.ListPathsResponse{
	// 	DownloadPaths: []*string{
	// 	},
	// 	StreamingPaths: []*armmediaservices.StreamingPath{
	// 		{
	// 			EncryptionScheme: to.Ptr(armmediaservices.EncryptionSchemeEnvelopeEncryption),
	// 			Paths: []*string{
	// 				to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest1.ism/manifest(format=m3u8-aapl,encryption=cbc)"),
	// 				to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest2.ism/manifest(format=m3u8-aapl,encryption=cbc)"),
	// 				to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest3.ism/manifest(format=m3u8-aapl,encryption=cbc)")},
	// 				StreamingProtocol: to.Ptr(armmediaservices.StreamingPolicyStreamingProtocolHls),
	// 			},
	// 			{
	// 				EncryptionScheme: to.Ptr(armmediaservices.EncryptionSchemeEnvelopeEncryption),
	// 				Paths: []*string{
	// 					to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest1.ism/manifest(format=mpd-time-csf,encryption=cbc)"),
	// 					to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest2.ism/manifest(format=mpd-time-csf,encryption=cbc)"),
	// 					to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest3.ism/manifest(format=mpd-time-csf,encryption=cbc)")},
	// 					StreamingProtocol: to.Ptr(armmediaservices.StreamingPolicyStreamingProtocolDash),
	// 				},
	// 				{
	// 					EncryptionScheme: to.Ptr(armmediaservices.EncryptionSchemeEnvelopeEncryption),
	// 					Paths: []*string{
	// 						to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest1.ism/manifest(encryption=cbc)"),
	// 						to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest2.ism/manifest(encryption=cbc)"),
	// 						to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest3.ism/manifest(encryption=cbc)")},
	// 						StreamingProtocol: to.Ptr(armmediaservices.StreamingPolicyStreamingProtocolSmoothStreaming),
	// 					},
	// 					{
	// 						EncryptionScheme: to.Ptr(armmediaservices.EncryptionSchemeCommonEncryptionCenc),
	// 						Paths: []*string{
	// 							to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest1.ism/manifest(format=mpd-time-csf,encryption=cenc)"),
	// 							to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest2.ism/manifest(format=mpd-time-csf,encryption=cenc)"),
	// 							to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest3.ism/manifest(format=mpd-time-csf,encryption=cenc)")},
	// 							StreamingProtocol: to.Ptr(armmediaservices.StreamingPolicyStreamingProtocolDash),
	// 						},
	// 						{
	// 							EncryptionScheme: to.Ptr(armmediaservices.EncryptionSchemeCommonEncryptionCenc),
	// 							Paths: []*string{
	// 								to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest1.ism/manifest(encryption=cenc)"),
	// 								to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest2.ism/manifest(encryption=cenc)"),
	// 								to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest3.ism/manifest(encryption=cenc)")},
	// 								StreamingProtocol: to.Ptr(armmediaservices.StreamingPolicyStreamingProtocolSmoothStreaming),
	// 							},
	// 							{
	// 								EncryptionScheme: to.Ptr(armmediaservices.EncryptionSchemeCommonEncryptionCbcs),
	// 								Paths: []*string{
	// 									to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest1.ism/manifest(format=m3u8-aapl,encryption=cbcs-aapl)"),
	// 									to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest2.ism/manifest(format=m3u8-aapl,encryption=cbcs-aapl)"),
	// 									to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest3.ism/manifest(format=m3u8-aapl,encryption=cbcs-aapl)")},
	// 									StreamingProtocol: to.Ptr(armmediaservices.StreamingPolicyStreamingProtocolHls),
	// 							}},
	// 						}
}
