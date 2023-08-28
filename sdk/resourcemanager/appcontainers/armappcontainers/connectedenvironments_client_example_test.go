//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/58be094c6b365f8d4d73a91e293dfb4818e57cf6/specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/ConnectedEnvironments_ListBySubscription.json
func ExampleConnectedEnvironmentsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConnectedEnvironmentsClient().NewListBySubscriptionPager(nil)
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
		// page.ConnectedEnvironmentCollection = armappcontainers.ConnectedEnvironmentCollection{
		// 	Value: []*armappcontainers.ConnectedEnvironment{
		// 		{
		// 			Name: to.Ptr("sample1"),
		// 			Type: to.Ptr("Microsoft.App/connectedEnvironments"),
		// 			ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/connectedEnvironments/sample1"),
		// 			Location: to.Ptr("North Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armappcontainers.ConnectedEnvironmentProperties{
		// 				CustomDomainConfiguration: &armappcontainers.CustomDomainConfiguration{
		// 					CustomDomainVerificationID: to.Ptr("custom domain verification id"),
		// 					DNSSuffix: to.Ptr("www.my-name.com"),
		// 					ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-06T04:00:00Z"); return t}()),
		// 					SubjectName: to.Ptr("CN=www.my-name.com"),
		// 					Thumbprint: to.Ptr("CERTIFICATE_THUMBPRINT"),
		// 				},
		// 				DefaultDomain: to.Ptr("sample1.k4apps.io"),
		// 				ProvisioningState: to.Ptr(armappcontainers.ConnectedEnvironmentProvisioningStateSucceeded),
		// 				StaticIP: to.Ptr("20.42.33.145"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sample2"),
		// 			Type: to.Ptr("Microsoft.App/connectedEnvironments"),
		// 			ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/DemoRG/providers/Microsoft.App/connectedEnvironments/sample2"),
		// 			Location: to.Ptr("North Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armappcontainers.ConnectedEnvironmentProperties{
		// 				CustomDomainConfiguration: &armappcontainers.CustomDomainConfiguration{
		// 					CustomDomainVerificationID: to.Ptr("custom domain verification id"),
		// 					DNSSuffix: to.Ptr("www.my-name2.com"),
		// 					ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-06T04:00:00Z"); return t}()),
		// 					SubjectName: to.Ptr("CN=www.my-name2.com"),
		// 					Thumbprint: to.Ptr("CERTIFICATE_THUMBPRINT"),
		// 				},
		// 				DefaultDomain: to.Ptr("sample2.k4apps.io"),
		// 				ProvisioningState: to.Ptr(armappcontainers.ConnectedEnvironmentProvisioningStateSucceeded),
		// 				StaticIP: to.Ptr("52.142.21.61"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/58be094c6b365f8d4d73a91e293dfb4818e57cf6/specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/ConnectedEnvironments_ListByResourceGroup.json
func ExampleConnectedEnvironmentsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConnectedEnvironmentsClient().NewListByResourceGroupPager("examplerg", nil)
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
		// page.ConnectedEnvironmentCollection = armappcontainers.ConnectedEnvironmentCollection{
		// 	Value: []*armappcontainers.ConnectedEnvironment{
		// 		{
		// 			Name: to.Ptr("sample1"),
		// 			Type: to.Ptr("Microsoft.App/connectedEnvironments"),
		// 			ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/connectedEnvironments/sample1"),
		// 			Location: to.Ptr("North Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			ExtendedLocation: &armappcontainers.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation"),
		// 				Type: to.Ptr(armappcontainers.ExtendedLocationTypesCustomLocation),
		// 			},
		// 			Properties: &armappcontainers.ConnectedEnvironmentProperties{
		// 				CustomDomainConfiguration: &armappcontainers.CustomDomainConfiguration{
		// 					CustomDomainVerificationID: to.Ptr("custom domain verification id"),
		// 					DNSSuffix: to.Ptr("www.my-name.com"),
		// 					ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-06T04:00:00Z"); return t}()),
		// 					SubjectName: to.Ptr("CN=www.my-name.com"),
		// 					Thumbprint: to.Ptr("CERTIFICATE_THUMBPRINT"),
		// 				},
		// 				DefaultDomain: to.Ptr("sample1.k4apps.io"),
		// 				ProvisioningState: to.Ptr(armappcontainers.ConnectedEnvironmentProvisioningStateSucceeded),
		// 				StaticIP: to.Ptr("20.42.33.145"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sample2"),
		// 			Type: to.Ptr("Microsoft.App/connectedEnvironments"),
		// 			ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/connectedEnvironments/sample2"),
		// 			Location: to.Ptr("North Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			ExtendedLocation: &armappcontainers.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation"),
		// 				Type: to.Ptr(armappcontainers.ExtendedLocationTypesCustomLocation),
		// 			},
		// 			Properties: &armappcontainers.ConnectedEnvironmentProperties{
		// 				CustomDomainConfiguration: &armappcontainers.CustomDomainConfiguration{
		// 					CustomDomainVerificationID: to.Ptr("custom domain verification id"),
		// 					DNSSuffix: to.Ptr("www.my-name2.com"),
		// 					ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-06T04:00:00Z"); return t}()),
		// 					SubjectName: to.Ptr("CN=www.my-name2.com"),
		// 					Thumbprint: to.Ptr("CERTIFICATE_THUMBPRINT"),
		// 				},
		// 				DefaultDomain: to.Ptr("sample2.k4apps.io"),
		// 				ProvisioningState: to.Ptr(armappcontainers.ConnectedEnvironmentProvisioningStateSucceeded),
		// 				StaticIP: to.Ptr("52.142.21.61"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/58be094c6b365f8d4d73a91e293dfb4818e57cf6/specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/ConnectedEnvironments_Get.json
func ExampleConnectedEnvironmentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectedEnvironmentsClient().Get(ctx, "examplerg", "examplekenv", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConnectedEnvironment = armappcontainers.ConnectedEnvironment{
	// 	Name: to.Ptr("examplekenv"),
	// 	Type: to.Ptr("Microsoft.App/kubeEnvironments"),
	// 	ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/kubeEnvironments/examplekenv"),
	// 	Location: to.Ptr("North Central US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	ExtendedLocation: &armappcontainers.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation"),
	// 		Type: to.Ptr(armappcontainers.ExtendedLocationTypesCustomLocation),
	// 	},
	// 	Properties: &armappcontainers.ConnectedEnvironmentProperties{
	// 		CustomDomainConfiguration: &armappcontainers.CustomDomainConfiguration{
	// 			CustomDomainVerificationID: to.Ptr("custom domain verification id"),
	// 			DNSSuffix: to.Ptr("www.my-name.com"),
	// 			ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-06T04:00:00Z"); return t}()),
	// 			SubjectName: to.Ptr("CN=www.my-name.com"),
	// 			Thumbprint: to.Ptr("CERTIFICATE_THUMBPRINT"),
	// 		},
	// 		DefaultDomain: to.Ptr("examplekenv.k4apps.io"),
	// 		ProvisioningState: to.Ptr(armappcontainers.ConnectedEnvironmentProvisioningStateSucceeded),
	// 		StaticIP: to.Ptr("20.42.33.145"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/58be094c6b365f8d4d73a91e293dfb4818e57cf6/specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/ConnectedEnvironments_CreateOrUpdate.json
func ExampleConnectedEnvironmentsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConnectedEnvironmentsClient().BeginCreateOrUpdate(ctx, "examplerg", "testenv", armappcontainers.ConnectedEnvironment{
		Location: to.Ptr("East US"),
		Properties: &armappcontainers.ConnectedEnvironmentProperties{
			CustomDomainConfiguration: &armappcontainers.CustomDomainConfiguration{
				CertificatePassword: to.Ptr("private key password"),
				CertificateValue:    []byte("Y2VydA=="),
				DNSSuffix:           to.Ptr("www.my-name.com"),
			},
			DaprAIConnectionString: to.Ptr("InstrumentationKey=00000000-0000-0000-0000-000000000000;IngestionEndpoint=https://northcentralus-0.in.applicationinsights.azure.com/"),
			StaticIP:               to.Ptr("1.2.3.4"),
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
	// res.ConnectedEnvironment = armappcontainers.ConnectedEnvironment{
	// 	Name: to.Ptr("testenv"),
	// 	Type: to.Ptr("Microsoft.App/connectedEnvironments"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/connectedEnvironments/testenv"),
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	ExtendedLocation: &armappcontainers.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation"),
	// 		Type: to.Ptr(armappcontainers.ExtendedLocationTypesCustomLocation),
	// 	},
	// 	Properties: &armappcontainers.ConnectedEnvironmentProperties{
	// 		CustomDomainConfiguration: &armappcontainers.CustomDomainConfiguration{
	// 			CustomDomainVerificationID: to.Ptr("custom domain verification id"),
	// 			DNSSuffix: to.Ptr("www.my-name.com"),
	// 			ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-06T04:00:00Z"); return t}()),
	// 			SubjectName: to.Ptr("CN=www.my-name.com"),
	// 			Thumbprint: to.Ptr("CERTIFICATE_THUMBPRINT"),
	// 		},
	// 		DefaultDomain: to.Ptr("testenv.k4apps.io"),
	// 		ProvisioningState: to.Ptr(armappcontainers.ConnectedEnvironmentProvisioningStateSucceeded),
	// 		StaticIP: to.Ptr("1.2.3.4"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/58be094c6b365f8d4d73a91e293dfb4818e57cf6/specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/ConnectedEnvironments_Delete.json
func ExampleConnectedEnvironmentsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConnectedEnvironmentsClient().BeginDelete(ctx, "examplerg", "examplekenv", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/58be094c6b365f8d4d73a91e293dfb4818e57cf6/specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/ConnectedEnvironments_Patch.json
func ExampleConnectedEnvironmentsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectedEnvironmentsClient().Update(ctx, "examplerg", "testenv", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConnectedEnvironment = armappcontainers.ConnectedEnvironment{
	// 	Name: to.Ptr("testenv"),
	// 	Type: to.Ptr("Microsoft.App/connectedEnvironments"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/connectedEnvironments/testenv"),
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armappcontainers.ConnectedEnvironmentProperties{
	// 		CustomDomainConfiguration: &armappcontainers.CustomDomainConfiguration{
	// 			CustomDomainVerificationID: to.Ptr("custom domain verification id"),
	// 			DNSSuffix: to.Ptr("www.my-name.com"),
	// 			ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-06T04:00:00Z"); return t}()),
	// 			SubjectName: to.Ptr("CN=www.my-name.com"),
	// 			Thumbprint: to.Ptr("CERTIFICATE_THUMBPRINT"),
	// 		},
	// 		DefaultDomain: to.Ptr("testenv.k4apps.io"),
	// 		ProvisioningState: to.Ptr(armappcontainers.ConnectedEnvironmentProvisioningStateSucceeded),
	// 		StaticIP: to.Ptr("20.42.33.145"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/58be094c6b365f8d4d73a91e293dfb4818e57cf6/specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/ConnectedEnvironmentsCertificates_CheckNameAvailability.json
func ExampleConnectedEnvironmentsClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectedEnvironmentsClient().CheckNameAvailability(ctx, "examplerg", "testcontainerenv", armappcontainers.CheckNameAvailabilityRequest{
		Name: to.Ptr("testcertificatename"),
		Type: to.Ptr("Microsoft.App/connectedEnvironments/certificates"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckNameAvailabilityResponse = armappcontainers.CheckNameAvailabilityResponse{
	// 	Message: to.Ptr(""),
	// 	NameAvailable: to.Ptr(true),
	// 	Reason: to.Ptr(armappcontainers.CheckNameAvailabilityReason("None")),
	// }
}
