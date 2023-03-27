//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armauthorization_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/53b1affe357b3bfbb53721d0a2002382a046d3b0/specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/PutRoleEligibilityScheduleRequest.json
func ExampleRoleEligibilityScheduleRequestsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewRoleEligibilityScheduleRequestsClient().Create(ctx, "providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f", "64caffb6-55c0-4deb-a585-68e948ea1ad6", armauthorization.RoleEligibilityScheduleRequest{
		Properties: &armauthorization.RoleEligibilityScheduleRequestProperties{
			Condition:        to.Ptr("@Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName] StringEqualsIgnoreCase 'foo_storage_container'"),
			ConditionVersion: to.Ptr("1.0"),
			PrincipalID:      to.Ptr("a3bb8764-cb92-4276-9d2a-ca1e895e55ea"),
			RequestType:      to.Ptr(armauthorization.RequestTypeAdminAssign),
			RoleDefinitionID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleDefinitions/c8d4ff99-41c3-41a8-9f60-21dfdad59608"),
			ScheduleInfo: &armauthorization.RoleEligibilityScheduleRequestPropertiesScheduleInfo{
				Expiration: &armauthorization.RoleEligibilityScheduleRequestPropertiesScheduleInfoExpiration{
					Type:     to.Ptr(armauthorization.TypeAfterDuration),
					Duration: to.Ptr("P365D"),
				},
				StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-09T21:31:27.91Z"); return t }()),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/53b1affe357b3bfbb53721d0a2002382a046d3b0/specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/GetRoleEligibilityScheduleRequestByName.json
func ExampleRoleEligibilityScheduleRequestsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoleEligibilityScheduleRequestsClient().Get(ctx, "providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f", "64caffb6-55c0-4deb-a585-68e948ea1ad6", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RoleEligibilityScheduleRequest = armauthorization.RoleEligibilityScheduleRequest{
	// 	Name: to.Ptr("64caffb6-55c0-4deb-a585-68e948ea1ad6"),
	// 	Type: to.Ptr("Microsoft.Authorization/RoleEligibilityRequests"),
	// 	ID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/RoleEligibilityRequests/64caffb6-55c0-4deb-a585-68e948ea1ad6"),
	// 	Properties: &armauthorization.RoleEligibilityScheduleRequestProperties{
	// 		Condition: to.Ptr("@Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName] StringEqualsIgnoreCase 'foo_storage_container'"),
	// 		ConditionVersion: to.Ptr("1.0"),
	// 		CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-09T21:32:27.91Z"); return t}()),
	// 		ExpandedProperties: &armauthorization.ExpandedProperties{
	// 			Principal: &armauthorization.ExpandedPropertiesPrincipal{
	// 				Type: to.Ptr("User"),
	// 				DisplayName: to.Ptr("User Account"),
	// 				Email: to.Ptr("user@my-tenant.com"),
	// 				ID: to.Ptr("a3bb8764-cb92-4276-9d2a-ca1e895e55ea"),
	// 			},
	// 			RoleDefinition: &armauthorization.ExpandedPropertiesRoleDefinition{
	// 				Type: to.Ptr("BuiltInRole"),
	// 				DisplayName: to.Ptr("Contributor"),
	// 				ID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleDefinitions/c8d4ff99-41c3-41a8-9f60-21dfdad59608"),
	// 			},
	// 			Scope: &armauthorization.ExpandedPropertiesScope{
	// 				Type: to.Ptr("subscription"),
	// 				DisplayName: to.Ptr("Pay-As-You-Go"),
	// 				ID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f"),
	// 			},
	// 		},
	// 		PrincipalID: to.Ptr("a3bb8764-cb92-4276-9d2a-ca1e895e55ea"),
	// 		PrincipalType: to.Ptr(armauthorization.PrincipalTypeUser),
	// 		RequestType: to.Ptr(armauthorization.RequestTypeAdminAssign),
	// 		RequestorID: to.Ptr("a3bb8764-cb92-4276-9d2a-ca1e895e55ea"),
	// 		RoleDefinitionID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleDefinitions/c8d4ff99-41c3-41a8-9f60-21dfdad59608"),
	// 		ScheduleInfo: &armauthorization.RoleEligibilityScheduleRequestPropertiesScheduleInfo{
	// 			Expiration: &armauthorization.RoleEligibilityScheduleRequestPropertiesScheduleInfoExpiration{
	// 				Type: to.Ptr(armauthorization.TypeAfterDuration),
	// 				Duration: to.Ptr("P365D"),
	// 			},
	// 			StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-09T21:31:27.91Z"); return t}()),
	// 		},
	// 		Scope: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f"),
	// 		Status: to.Ptr(armauthorization.StatusProvisioned),
	// 		TargetRoleEligibilityScheduleID: to.Ptr("b1477448-2cc6-4ceb-93b4-54a202a89413"),
	// 		TicketInfo: &armauthorization.RoleEligibilityScheduleRequestPropertiesTicketInfo{
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/53b1affe357b3bfbb53721d0a2002382a046d3b0/specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/GetRoleEligibilityScheduleRequestByScope.json
func ExampleRoleEligibilityScheduleRequestsClient_NewListForScopePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRoleEligibilityScheduleRequestsClient().NewListForScopePager("providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f", &armauthorization.RoleEligibilityScheduleRequestsClientListForScopeOptions{Filter: to.Ptr("assignedTo('A3BB8764-CB92-4276-9D2A-CA1E895E55EA')")})
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
		// page.RoleEligibilityScheduleRequestListResult = armauthorization.RoleEligibilityScheduleRequestListResult{
		// 	Value: []*armauthorization.RoleEligibilityScheduleRequest{
		// 		{
		// 			Name: to.Ptr("64caffb6-55c0-4deb-a585-68e948ea1ad6"),
		// 			Type: to.Ptr("Microsoft.Authorization/RoleEligibilityRequests"),
		// 			ID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/RoleEligibilityRequests/64caffb6-55c0-4deb-a585-68e948ea1ad6"),
		// 			Properties: &armauthorization.RoleEligibilityScheduleRequestProperties{
		// 				Condition: to.Ptr("@Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName] StringEqualsIgnoreCase 'foo_storage_container'"),
		// 				ConditionVersion: to.Ptr("1.0"),
		// 				CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-09T21:32:27.91Z"); return t}()),
		// 				ExpandedProperties: &armauthorization.ExpandedProperties{
		// 					Principal: &armauthorization.ExpandedPropertiesPrincipal{
		// 						Type: to.Ptr("User"),
		// 						DisplayName: to.Ptr("User Account"),
		// 						Email: to.Ptr("user@my-tenant.com"),
		// 						ID: to.Ptr("a3bb8764-cb92-4276-9d2a-ca1e895e55ea"),
		// 					},
		// 					RoleDefinition: &armauthorization.ExpandedPropertiesRoleDefinition{
		// 						Type: to.Ptr("BuiltInRole"),
		// 						DisplayName: to.Ptr("Contributor"),
		// 						ID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleDefinitions/c8d4ff99-41c3-41a8-9f60-21dfdad59608"),
		// 					},
		// 					Scope: &armauthorization.ExpandedPropertiesScope{
		// 						Type: to.Ptr("subscription"),
		// 						DisplayName: to.Ptr("Pay-As-You-Go"),
		// 						ID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f"),
		// 					},
		// 				},
		// 				PrincipalID: to.Ptr("a3bb8764-cb92-4276-9d2a-ca1e895e55ea"),
		// 				PrincipalType: to.Ptr(armauthorization.PrincipalTypeUser),
		// 				RequestType: to.Ptr(armauthorization.RequestTypeAdminAssign),
		// 				RequestorID: to.Ptr("a3bb8764-cb92-4276-9d2a-ca1e895e55ea"),
		// 				RoleDefinitionID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleDefinitions/c8d4ff99-41c3-41a8-9f60-21dfdad59608"),
		// 				ScheduleInfo: &armauthorization.RoleEligibilityScheduleRequestPropertiesScheduleInfo{
		// 					Expiration: &armauthorization.RoleEligibilityScheduleRequestPropertiesScheduleInfoExpiration{
		// 						Type: to.Ptr(armauthorization.TypeAfterDuration),
		// 						Duration: to.Ptr("P365D"),
		// 					},
		// 					StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-09T21:31:27.91Z"); return t}()),
		// 				},
		// 				Scope: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f"),
		// 				Status: to.Ptr(armauthorization.StatusProvisioned),
		// 				TargetRoleEligibilityScheduleID: to.Ptr("b1477448-2cc6-4ceb-93b4-54a202a89413"),
		// 				TicketInfo: &armauthorization.RoleEligibilityScheduleRequestPropertiesTicketInfo{
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/53b1affe357b3bfbb53721d0a2002382a046d3b0/specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/CancelRoleEligibilityScheduleRequestByName.json
func ExampleRoleEligibilityScheduleRequestsClient_Cancel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewRoleEligibilityScheduleRequestsClient().Cancel(ctx, "providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f", "64caffb6-55c0-4deb-a585-68e948ea1ad6", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/53b1affe357b3bfbb53721d0a2002382a046d3b0/specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/ValidateRoleEligibilityScheduleRequestByName.json
func ExampleRoleEligibilityScheduleRequestsClient_Validate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoleEligibilityScheduleRequestsClient().Validate(ctx, "subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f", "64caffb6-55c0-4deb-a585-68e948ea1ad6", armauthorization.RoleEligibilityScheduleRequest{
		Properties: &armauthorization.RoleEligibilityScheduleRequestProperties{
			Condition:        to.Ptr("@Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName] StringEqualsIgnoreCase 'foo_storage_container'"),
			ConditionVersion: to.Ptr("1.0"),
			PrincipalID:      to.Ptr("a3bb8764-cb92-4276-9d2a-ca1e895e55ea"),
			RequestType:      to.Ptr(armauthorization.RequestTypeAdminAssign),
			RoleDefinitionID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleDefinitions/c8d4ff99-41c3-41a8-9f60-21dfdad59608"),
			ScheduleInfo: &armauthorization.RoleEligibilityScheduleRequestPropertiesScheduleInfo{
				Expiration: &armauthorization.RoleEligibilityScheduleRequestPropertiesScheduleInfoExpiration{
					Type:     to.Ptr(armauthorization.TypeAfterDuration),
					Duration: to.Ptr("P365D"),
				},
				StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-09T21:31:27.91Z"); return t }()),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RoleEligibilityScheduleRequest = armauthorization.RoleEligibilityScheduleRequest{
	// 	Name: to.Ptr("64caffb6-55c0-4deb-a585-68e948ea1ad6"),
	// 	Type: to.Ptr("Microsoft.Authorization/RoleEligibilityScheduleRequests"),
	// 	ID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/RoleEligibilityScheduleRequests/64caffb6-55c0-4deb-a585-68e948ea1ad6"),
	// 	Properties: &armauthorization.RoleEligibilityScheduleRequestProperties{
	// 		Condition: to.Ptr("@Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName] StringEqualsIgnoreCase 'foo_storage_container'"),
	// 		ConditionVersion: to.Ptr("1.0"),
	// 		CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-09T21:32:27.91Z"); return t}()),
	// 		ExpandedProperties: &armauthorization.ExpandedProperties{
	// 			Principal: &armauthorization.ExpandedPropertiesPrincipal{
	// 				Type: to.Ptr("User"),
	// 				DisplayName: to.Ptr("User Account"),
	// 				Email: to.Ptr("user@my-tenant.com"),
	// 				ID: to.Ptr("a3bb8764-cb92-4276-9d2a-ca1e895e55ea"),
	// 			},
	// 			RoleDefinition: &armauthorization.ExpandedPropertiesRoleDefinition{
	// 				Type: to.Ptr("BuiltInRole"),
	// 				DisplayName: to.Ptr("Contributor"),
	// 				ID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleDefinitions/c8d4ff99-41c3-41a8-9f60-21dfdad59608"),
	// 			},
	// 			Scope: &armauthorization.ExpandedPropertiesScope{
	// 				Type: to.Ptr("subscription"),
	// 				DisplayName: to.Ptr("Pay-As-You-Go"),
	// 				ID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f"),
	// 			},
	// 		},
	// 		PrincipalID: to.Ptr("a3bb8764-cb92-4276-9d2a-ca1e895e55ea"),
	// 		PrincipalType: to.Ptr(armauthorization.PrincipalTypeUser),
	// 		RequestType: to.Ptr(armauthorization.RequestTypeAdminAssign),
	// 		RequestorID: to.Ptr("a3bb8764-cb92-4276-9d2a-ca1e895e55ea"),
	// 		RoleDefinitionID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleDefinitions/c8d4ff99-41c3-41a8-9f60-21dfdad59608"),
	// 		ScheduleInfo: &armauthorization.RoleEligibilityScheduleRequestPropertiesScheduleInfo{
	// 			Expiration: &armauthorization.RoleEligibilityScheduleRequestPropertiesScheduleInfoExpiration{
	// 				Type: to.Ptr(armauthorization.TypeAfterDuration),
	// 				Duration: to.Ptr("P365D"),
	// 			},
	// 			StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-09T21:31:27.91Z"); return t}()),
	// 		},
	// 		Scope: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f"),
	// 		Status: to.Ptr(armauthorization.StatusProvisioned),
	// 		TargetRoleEligibilityScheduleID: to.Ptr("b1477448-2cc6-4ceb-93b4-54a202a89413"),
	// 		TicketInfo: &armauthorization.RoleEligibilityScheduleRequestPropertiesTicketInfo{
	// 		},
	// 	},
	// }
}
