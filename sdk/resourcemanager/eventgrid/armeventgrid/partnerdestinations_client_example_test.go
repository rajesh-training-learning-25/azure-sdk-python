//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armeventgrid_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8691fbfca8fcdc5a241a0b501c32fd4a76bb0cd/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/PartnerDestinations_Get.json
func ExamplePartnerDestinationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPartnerDestinationsClient().Get(ctx, "examplerg", "examplePartnerDestinationName1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PartnerDestination = armeventgrid.PartnerDestination{
	// 	Name: to.Ptr("examplePartnerDestinationName1"),
	// 	Type: to.Ptr("Microsoft.EventGrid/partnerDestinations"),
	// 	ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/partnerDestinations/examplePartnerDestinationName1"),
	// 	Location: to.Ptr("centraluseuap"),
	// 	Properties: &armeventgrid.PartnerDestinationProperties{
	// 		ActivationState: to.Ptr(armeventgrid.PartnerDestinationActivationStateNeverActivated),
	// 		EndpointBaseURL: to.Ptr("https://somepartnerhostname"),
	// 		EndpointServiceContext: to.Ptr("ContosoCorp.Accounts.User1"),
	// 		ExpirationTimeIfNotActivatedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-21T22:50:25.410Z"); return t}()),
	// 		MessageForActivation: to.Ptr("Some message to the approver"),
	// 		ProvisioningState: to.Ptr(armeventgrid.PartnerDestinationProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8691fbfca8fcdc5a241a0b501c32fd4a76bb0cd/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/PartnerDestinations_CreateOrUpdate.json
func ExamplePartnerDestinationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPartnerDestinationsClient().BeginCreateOrUpdate(ctx, "examplerg", "examplePartnerDestinationName1", armeventgrid.PartnerDestination{
		Location: to.Ptr("westus2"),
		Properties: &armeventgrid.PartnerDestinationProperties{
			EndpointBaseURL:                 to.Ptr("https://www.example/endpoint"),
			EndpointServiceContext:          to.Ptr("This is an example"),
			ExpirationTimeIfNotActivatedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-14T19:33:43.430Z"); return t }()),
			MessageForActivation:            to.Ptr("Sample Activation message"),
			PartnerRegistrationImmutableID:  to.Ptr("0bd70ee2-7d95-447e-ab1f-c4f320019404"),
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
	// res.PartnerDestination = armeventgrid.PartnerDestination{
	// 	Name: to.Ptr("examplePartnerDestinationName1"),
	// 	Type: to.Ptr("Microsoft.EventGrid/partnerDestinations"),
	// 	ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/partnerDestinations/examplePartnerDestinationName1"),
	// 	Location: to.Ptr("westus2"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armeventgrid.PartnerDestinationProperties{
	// 		ActivationState: to.Ptr(armeventgrid.PartnerDestinationActivationStateNeverActivated),
	// 		EndpointBaseURL: to.Ptr("https://www.example/endpoint"),
	// 		EndpointServiceContext: to.Ptr("string"),
	// 		ExpirationTimeIfNotActivatedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-14T19:33:43.430Z"); return t}()),
	// 		MessageForActivation: to.Ptr("Sample Activation message"),
	// 		PartnerRegistrationImmutableID: to.Ptr("0bd70ee2-7d95-447e-ab1f-c4f320019404"),
	// 		ProvisioningState: to.Ptr(armeventgrid.PartnerDestinationProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8691fbfca8fcdc5a241a0b501c32fd4a76bb0cd/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/PartnerDestinations_Delete.json
func ExamplePartnerDestinationsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPartnerDestinationsClient().BeginDelete(ctx, "examplerg", "examplePartnerDestinationName1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8691fbfca8fcdc5a241a0b501c32fd4a76bb0cd/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/PartnerDestinations_Update.json
func ExamplePartnerDestinationsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPartnerDestinationsClient().BeginUpdate(ctx, "examplerg", "examplePartnerDestinationName1", armeventgrid.PartnerDestinationUpdateParameters{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
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
	// res.PartnerDestination = armeventgrid.PartnerDestination{
	// 	Name: to.Ptr("examplePartnerDestinationName1"),
	// 	Type: to.Ptr("Microsoft.EventGrid/partnerDestinations"),
	// 	ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/partnerDestinations/examplePartnerDestinationName1"),
	// 	Location: to.Ptr("centraluseuap"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armeventgrid.PartnerDestinationProperties{
	// 		ActivationState: to.Ptr(armeventgrid.PartnerDestinationActivationStateNeverActivated),
	// 		EndpointBaseURL: to.Ptr("https://somepartnerhostname"),
	// 		EndpointServiceContext: to.Ptr("ContosoCorp.Accounts.User1"),
	// 		ExpirationTimeIfNotActivatedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-21T22:50:25.410Z"); return t}()),
	// 		MessageForActivation: to.Ptr("Some message to the approver"),
	// 		ProvisioningState: to.Ptr(armeventgrid.PartnerDestinationProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8691fbfca8fcdc5a241a0b501c32fd4a76bb0cd/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/PartnerDestinations_ListBySubscription.json
func ExamplePartnerDestinationsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPartnerDestinationsClient().NewListBySubscriptionPager(&armeventgrid.PartnerDestinationsClientListBySubscriptionOptions{Filter: nil,
		Top: nil,
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
		// page.PartnerDestinationsListResult = armeventgrid.PartnerDestinationsListResult{
		// 	Value: []*armeventgrid.PartnerDestination{
		// 		{
		// 			Name: to.Ptr("examplePartnerDestinationName1"),
		// 			Type: to.Ptr("Microsoft.EventGrid/partnerDestinations"),
		// 			ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/partnerDestinations/examplePartnerDestinationName1"),
		// 			Location: to.Ptr("centraluseuap"),
		// 			Properties: &armeventgrid.PartnerDestinationProperties{
		// 				ActivationState: to.Ptr(armeventgrid.PartnerDestinationActivationStateNeverActivated),
		// 				EndpointBaseURL: to.Ptr("https://somepartnerhostname"),
		// 				EndpointServiceContext: to.Ptr("ContosoCorp.Accounts.User1"),
		// 				ExpirationTimeIfNotActivatedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-21T22:50:25.410Z"); return t}()),
		// 				MessageForActivation: to.Ptr("Some message to the approver"),
		// 				ProvisioningState: to.Ptr(armeventgrid.PartnerDestinationProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8691fbfca8fcdc5a241a0b501c32fd4a76bb0cd/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/PartnerDestinations_ListByResourceGroup.json
func ExamplePartnerDestinationsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPartnerDestinationsClient().NewListByResourceGroupPager("examplerg", &armeventgrid.PartnerDestinationsClientListByResourceGroupOptions{Filter: nil,
		Top: nil,
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
		// page.PartnerDestinationsListResult = armeventgrid.PartnerDestinationsListResult{
		// 	Value: []*armeventgrid.PartnerDestination{
		// 		{
		// 			Name: to.Ptr("examplePartnerDestinationName1"),
		// 			Type: to.Ptr("Microsoft.EventGrid/partnerDestinations"),
		// 			ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/partnerDestinations/examplePartnerDestinationName1"),
		// 			Location: to.Ptr("centraluseuap"),
		// 			Properties: &armeventgrid.PartnerDestinationProperties{
		// 				ActivationState: to.Ptr(armeventgrid.PartnerDestinationActivationStateNeverActivated),
		// 				EndpointBaseURL: to.Ptr("https://somepartnerhostname"),
		// 				EndpointServiceContext: to.Ptr("ContosoCorp.Accounts.User1"),
		// 				ExpirationTimeIfNotActivatedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-21T22:50:25.410Z"); return t}()),
		// 				MessageForActivation: to.Ptr("Some message to the approver"),
		// 				ProvisioningState: to.Ptr(armeventgrid.PartnerDestinationProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8691fbfca8fcdc5a241a0b501c32fd4a76bb0cd/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/PartnerDestinations_Activate.json
func ExamplePartnerDestinationsClient_Activate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPartnerDestinationsClient().Activate(ctx, "examplerg", "examplePartnerDestination1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PartnerDestination = armeventgrid.PartnerDestination{
	// 	Name: to.Ptr("examplePartnerDestinationName1"),
	// 	Type: to.Ptr("Microsoft.EventGrid/partnerDestinations"),
	// 	ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/partnerDestinations/examplePartnerDestinationName1"),
	// 	Location: to.Ptr("centraluseuap"),
	// 	Properties: &armeventgrid.PartnerDestinationProperties{
	// 		ActivationState: to.Ptr(armeventgrid.PartnerDestinationActivationStateActivated),
	// 		EndpointBaseURL: to.Ptr("https://somepartnerhostname"),
	// 		EndpointServiceContext: to.Ptr("ContosoCorp.Accounts.User1"),
	// 		MessageForActivation: to.Ptr("Some message to the approver"),
	// 		ProvisioningState: to.Ptr(armeventgrid.PartnerDestinationProvisioningStateSucceeded),
	// 	},
	// }
}
