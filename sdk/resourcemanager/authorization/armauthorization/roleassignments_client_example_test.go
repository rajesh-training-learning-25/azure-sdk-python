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

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/53b1affe357b3bfbb53721d0a2002382a046d3b0/specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/RoleAssignments_ListForSubscription.json
func ExampleRoleAssignmentsClient_NewListForSubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRoleAssignmentsClient().NewListForSubscriptionPager(&armauthorization.RoleAssignmentsClientListForSubscriptionOptions{Filter: nil,
		TenantID: nil,
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
		// page.RoleAssignmentListResult = armauthorization.RoleAssignmentListResult{
		// 	Value: []*armauthorization.RoleAssignment{
		// 		{
		// 			Name: to.Ptr("b0f43c54-e787-4862-89b1-a653fa9cf747"),
		// 			Type: to.Ptr("Microsoft.Authorization/roleAssignments"),
		// 			ID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/providers/Microsoft.Authorization/roleAssignments/b0f43c54-e787-4862-89b1-a653fa9cf747"),
		// 			Properties: &armauthorization.RoleAssignmentProperties{
		// 				PrincipalID: to.Ptr("ce2ce14e-85d7-4629-bdbc-454d0519d987"),
		// 				PrincipalType: to.Ptr(armauthorization.PrincipalTypeUser),
		// 				RoleDefinitionID: to.Ptr("/providers/Microsoft.Authorization/roleDefinitions/0b5fe924-9a61-425c-96af-cfe6e287ca2d"),
		// 				Scope: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/53b1affe357b3bfbb53721d0a2002382a046d3b0/specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/RoleAssignments_ListForResourceGroup.json
func ExampleRoleAssignmentsClient_NewListForResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRoleAssignmentsClient().NewListForResourceGroupPager("testrg", &armauthorization.RoleAssignmentsClientListForResourceGroupOptions{Filter: nil,
		TenantID: nil,
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
		// page.RoleAssignmentListResult = armauthorization.RoleAssignmentListResult{
		// 	Value: []*armauthorization.RoleAssignment{
		// 		{
		// 			Name: to.Ptr("b0f43c54-e787-4862-89b1-a653fa9cf747"),
		// 			Type: to.Ptr("Microsoft.Authorization/roleAssignments"),
		// 			ID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/providers/Microsoft.Authorization/roleAssignments/b0f43c54-e787-4862-89b1-a653fa9cf747"),
		// 			Properties: &armauthorization.RoleAssignmentProperties{
		// 				PrincipalID: to.Ptr("ce2ce14e-85d7-4629-bdbc-454d0519d987"),
		// 				PrincipalType: to.Ptr(armauthorization.PrincipalTypeUser),
		// 				RoleDefinitionID: to.Ptr("/providers/Microsoft.Authorization/roleDefinitions/0b5fe924-9a61-425c-96af-cfe6e287ca2d"),
		// 				Scope: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("96786e4b-dede-4c2e-8736-8ab911987f08"),
		// 			Type: to.Ptr("Microsoft.Authorization/roleAssignments"),
		// 			ID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/testrg/providers/Microsoft.Authorization/roleAssignments/96786e4b-dede-4c2e-8736-8ab911987f08"),
		// 			Properties: &armauthorization.RoleAssignmentProperties{
		// 				PrincipalID: to.Ptr("ce2ce14e-85d7-4629-bdbc-454d0519d987"),
		// 				PrincipalType: to.Ptr(armauthorization.PrincipalTypeUser),
		// 				RoleDefinitionID: to.Ptr("/providers/Microsoft.Authorization/roleDefinitions/0b5fe924-9a61-425c-96af-cfe6e287ca2d"),
		// 				Scope: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/testrg"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/53b1affe357b3bfbb53721d0a2002382a046d3b0/specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/RoleAssignments_ListForResource.json
func ExampleRoleAssignmentsClient_NewListForResourcePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRoleAssignmentsClient().NewListForResourcePager("testrg", "Microsoft.DocumentDb", "databaseAccounts", "test-db-account", &armauthorization.RoleAssignmentsClientListForResourceOptions{Filter: nil,
		TenantID: nil,
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
		// page.RoleAssignmentListResult = armauthorization.RoleAssignmentListResult{
		// 	Value: []*armauthorization.RoleAssignment{
		// 		{
		// 			Name: to.Ptr("b0f43c54-e787-4862-89b1-a653fa9cf747"),
		// 			Type: to.Ptr("Microsoft.Authorization/roleAssignments"),
		// 			ID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/providers/Microsoft.Authorization/roleAssignments/b0f43c54-e787-4862-89b1-a653fa9cf747"),
		// 			Properties: &armauthorization.RoleAssignmentProperties{
		// 				PrincipalID: to.Ptr("ce2ce14e-85d7-4629-bdbc-454d0519d987"),
		// 				PrincipalType: to.Ptr(armauthorization.PrincipalTypeUser),
		// 				RoleDefinitionID: to.Ptr("/providers/Microsoft.Authorization/roleDefinitions/0b5fe924-9a61-425c-96af-cfe6e287ca2d"),
		// 				Scope: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("96786e4b-dede-4c2e-8736-8ab911987f08"),
		// 			Type: to.Ptr("Microsoft.Authorization/roleAssignments"),
		// 			ID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/testrg/providers/Microsoft.Authorization/roleAssignments/96786e4b-dede-4c2e-8736-8ab911987f08"),
		// 			Properties: &armauthorization.RoleAssignmentProperties{
		// 				PrincipalID: to.Ptr("ce2ce14e-85d7-4629-bdbc-454d0519d987"),
		// 				PrincipalType: to.Ptr(armauthorization.PrincipalTypeUser),
		// 				RoleDefinitionID: to.Ptr("/providers/Microsoft.Authorization/roleDefinitions/0b5fe924-9a61-425c-96af-cfe6e287ca2d"),
		// 				Scope: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/testrg"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("05c5a614-a7d6-4502-b150-c2fb455033ff"),
		// 			Type: to.Ptr("Microsoft.Authorization/roleAssignments"),
		// 			ID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/testrg/providers/Microsoft.DocumentDb/databaseAccounts/test-db-account/providers/Microsoft.Authorization/roleAssignments/05c5a614-a7d6-4502-b150-c2fb455033ff"),
		// 			Properties: &armauthorization.RoleAssignmentProperties{
		// 				PrincipalID: to.Ptr("ce2ce14e-85d7-4629-bdbc-454d0519d987"),
		// 				PrincipalType: to.Ptr(armauthorization.PrincipalTypeUser),
		// 				RoleDefinitionID: to.Ptr("/providers/Microsoft.Authorization/roleDefinitions/0b5fe924-9a61-425c-96af-cfe6e287ca2d"),
		// 				Scope: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/testrg/providers/Microsoft.DocumentDb/databaseAccounts/test-db-account"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/53b1affe357b3bfbb53721d0a2002382a046d3b0/specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/RoleAssignments_Get.json
func ExampleRoleAssignmentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoleAssignmentsClient().Get(ctx, "subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2", "b0f43c54-e787-4862-89b1-a653fa9cf747", &armauthorization.RoleAssignmentsClientGetOptions{TenantID: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RoleAssignment = armauthorization.RoleAssignment{
	// 	Name: to.Ptr("b0f43c54-e787-4862-89b1-a653fa9cf747"),
	// 	Type: to.Ptr("Microsoft.Authorization/roleAssignments"),
	// 	ID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/providers/Microsoft.Authorization/roleAssignments/b0f43c54-e787-4862-89b1-a653fa9cf747"),
	// 	Properties: &armauthorization.RoleAssignmentProperties{
	// 		PrincipalID: to.Ptr("ce2ce14e-85d7-4629-bdbc-454d0519d987"),
	// 		PrincipalType: to.Ptr(armauthorization.PrincipalTypeUser),
	// 		RoleDefinitionID: to.Ptr("/providers/Microsoft.Authorization/roleDefinitions/0b5fe924-9a61-425c-96af-cfe6e287ca2d"),
	// 		Scope: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/53b1affe357b3bfbb53721d0a2002382a046d3b0/specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/RoleAssignments_CreateForResource.json
func ExampleRoleAssignmentsClient_Create_createRoleAssignmentForResource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoleAssignmentsClient().Create(ctx, "subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/testrg/providers/Microsoft.DocumentDb/databaseAccounts/test-db-account", "05c5a614-a7d6-4502-b150-c2fb455033ff", armauthorization.RoleAssignmentCreateParameters{
		Properties: &armauthorization.RoleAssignmentProperties{
			PrincipalID:      to.Ptr("ce2ce14e-85d7-4629-bdbc-454d0519d987"),
			PrincipalType:    to.Ptr(armauthorization.PrincipalTypeUser),
			RoleDefinitionID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/providers/Microsoft.Authorization/roleDefinitions/0b5fe924-9a61-425c-96af-cfe6e287ca2d"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RoleAssignment = armauthorization.RoleAssignment{
	// 	Name: to.Ptr("05c5a614-a7d6-4502-b150-c2fb455033ff"),
	// 	Type: to.Ptr("Microsoft.Authorization/roleAssignments"),
	// 	ID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/testrg/providers/Microsoft.DocumentDb/databaseAccounts/test-db-account/providers/Microsoft.Authorization/roleAssignments/05c5a614-a7d6-4502-b150-c2fb455033ff"),
	// 	Properties: &armauthorization.RoleAssignmentProperties{
	// 		PrincipalID: to.Ptr("ce2ce14e-85d7-4629-bdbc-454d0519d987"),
	// 		PrincipalType: to.Ptr(armauthorization.PrincipalTypeUser),
	// 		RoleDefinitionID: to.Ptr("/providers/Microsoft.Authorization/roleDefinitions/0b5fe924-9a61-425c-96af-cfe6e287ca2d"),
	// 		Scope: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/testrg/providers/Microsoft.DocumentDb/databaseAccounts/test-db-account"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/53b1affe357b3bfbb53721d0a2002382a046d3b0/specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/RoleAssignments_CreateForResourceGroup.json
func ExampleRoleAssignmentsClient_Create_createRoleAssignmentForResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoleAssignmentsClient().Create(ctx, "subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/testrg", "05c5a614-a7d6-4502-b150-c2fb455033ff", armauthorization.RoleAssignmentCreateParameters{
		Properties: &armauthorization.RoleAssignmentProperties{
			PrincipalID:      to.Ptr("ce2ce14e-85d7-4629-bdbc-454d0519d987"),
			PrincipalType:    to.Ptr(armauthorization.PrincipalTypeUser),
			RoleDefinitionID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/providers/Microsoft.Authorization/roleDefinitions/0b5fe924-9a61-425c-96af-cfe6e287ca2d"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RoleAssignment = armauthorization.RoleAssignment{
	// 	Name: to.Ptr("05c5a614-a7d6-4502-b150-c2fb455033ff"),
	// 	Type: to.Ptr("Microsoft.Authorization/roleAssignments"),
	// 	ID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/testrg/providers/Microsoft.Authorization/roleAssignments/05c5a614-a7d6-4502-b150-c2fb455033ff"),
	// 	Properties: &armauthorization.RoleAssignmentProperties{
	// 		PrincipalID: to.Ptr("ce2ce14e-85d7-4629-bdbc-454d0519d987"),
	// 		PrincipalType: to.Ptr(armauthorization.PrincipalTypeUser),
	// 		RoleDefinitionID: to.Ptr("/providers/Microsoft.Authorization/roleDefinitions/0b5fe924-9a61-425c-96af-cfe6e287ca2d"),
	// 		Scope: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/testrg"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/53b1affe357b3bfbb53721d0a2002382a046d3b0/specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/RoleAssignments_CreateForSubscription.json
func ExampleRoleAssignmentsClient_Create_createRoleAssignmentForSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoleAssignmentsClient().Create(ctx, "subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2", "05c5a614-a7d6-4502-b150-c2fb455033ff", armauthorization.RoleAssignmentCreateParameters{
		Properties: &armauthorization.RoleAssignmentProperties{
			PrincipalID:      to.Ptr("ce2ce14e-85d7-4629-bdbc-454d0519d987"),
			PrincipalType:    to.Ptr(armauthorization.PrincipalTypeUser),
			RoleDefinitionID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/providers/Microsoft.Authorization/roleDefinitions/0b5fe924-9a61-425c-96af-cfe6e287ca2d"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RoleAssignment = armauthorization.RoleAssignment{
	// 	Name: to.Ptr("05c5a614-a7d6-4502-b150-c2fb455033ff"),
	// 	Type: to.Ptr("Microsoft.Authorization/roleAssignments"),
	// 	ID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/providers/Microsoft.Authorization/roleAssignments/05c5a614-a7d6-4502-b150-c2fb455033ff"),
	// 	Properties: &armauthorization.RoleAssignmentProperties{
	// 		PrincipalID: to.Ptr("ce2ce14e-85d7-4629-bdbc-454d0519d987"),
	// 		PrincipalType: to.Ptr(armauthorization.PrincipalTypeUser),
	// 		RoleDefinitionID: to.Ptr("/providers/Microsoft.Authorization/roleDefinitions/0b5fe924-9a61-425c-96af-cfe6e287ca2d"),
	// 		Scope: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/53b1affe357b3bfbb53721d0a2002382a046d3b0/specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/RoleAssignments_Delete.json
func ExampleRoleAssignmentsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoleAssignmentsClient().Delete(ctx, "subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2", "b0f43c54-e787-4862-89b1-a653fa9cf747", &armauthorization.RoleAssignmentsClientDeleteOptions{TenantID: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RoleAssignment = armauthorization.RoleAssignment{
	// 	Name: to.Ptr("b0f43c54-e787-4862-89b1-a653fa9cf747"),
	// 	Type: to.Ptr("Microsoft.Authorization/roleAssignments"),
	// 	ID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/providers/Microsoft.Authorization/roleAssignments/b0f43c54-e787-4862-89b1-a653fa9cf747"),
	// 	Properties: &armauthorization.RoleAssignmentProperties{
	// 		PrincipalID: to.Ptr("ce2ce14e-85d7-4629-bdbc-454d0519d987"),
	// 		PrincipalType: to.Ptr(armauthorization.PrincipalTypeUser),
	// 		RoleDefinitionID: to.Ptr("/providers/Microsoft.Authorization/roleDefinitions/0b5fe924-9a61-425c-96af-cfe6e287ca2d"),
	// 		Scope: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/53b1affe357b3bfbb53721d0a2002382a046d3b0/specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/RoleAssignments_ListForScope.json
func ExampleRoleAssignmentsClient_NewListForScopePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRoleAssignmentsClient().NewListForScopePager("subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2", &armauthorization.RoleAssignmentsClientListForScopeOptions{Filter: nil,
		TenantID:  nil,
		SkipToken: nil,
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
		// page.RoleAssignmentListResult = armauthorization.RoleAssignmentListResult{
		// 	Value: []*armauthorization.RoleAssignment{
		// 		{
		// 			Name: to.Ptr("b0f43c54-e787-4862-89b1-a653fa9cf747"),
		// 			Type: to.Ptr("Microsoft.Authorization/roleAssignments"),
		// 			ID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/providers/Microsoft.Authorization/roleAssignments/b0f43c54-e787-4862-89b1-a653fa9cf747"),
		// 			Properties: &armauthorization.RoleAssignmentProperties{
		// 				PrincipalID: to.Ptr("ce2ce14e-85d7-4629-bdbc-454d0519d987"),
		// 				PrincipalType: to.Ptr(armauthorization.PrincipalTypeUser),
		// 				RoleDefinitionID: to.Ptr("/providers/Microsoft.Authorization/roleDefinitions/0b5fe924-9a61-425c-96af-cfe6e287ca2d"),
		// 				Scope: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/53b1affe357b3bfbb53721d0a2002382a046d3b0/specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/RoleAssignments_GetById.json
func ExampleRoleAssignmentsClient_GetByID() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoleAssignmentsClient().GetByID(ctx, "subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/providers/Microsoft.Authorization/roleAssignments/b0f43c54-e787-4862-89b1-a653fa9cf747", &armauthorization.RoleAssignmentsClientGetByIDOptions{TenantID: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RoleAssignment = armauthorization.RoleAssignment{
	// 	Name: to.Ptr("b0f43c54-e787-4862-89b1-a653fa9cf747"),
	// 	Type: to.Ptr("Microsoft.Authorization/roleAssignments"),
	// 	ID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/providers/Microsoft.Authorization/roleAssignments/b0f43c54-e787-4862-89b1-a653fa9cf747"),
	// 	Properties: &armauthorization.RoleAssignmentProperties{
	// 		PrincipalID: to.Ptr("ce2ce14e-85d7-4629-bdbc-454d0519d987"),
	// 		PrincipalType: to.Ptr(armauthorization.PrincipalTypeUser),
	// 		RoleDefinitionID: to.Ptr("/providers/Microsoft.Authorization/roleDefinitions/0b5fe924-9a61-425c-96af-cfe6e287ca2d"),
	// 		Scope: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/53b1affe357b3bfbb53721d0a2002382a046d3b0/specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/RoleAssignments_CreateById.json
func ExampleRoleAssignmentsClient_CreateByID() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoleAssignmentsClient().CreateByID(ctx, "subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/providers/Microsoft.Authorization/roleAssignments/b0f43c54-e787-4862-89b1-a653fa9cf747", armauthorization.RoleAssignmentCreateParameters{
		Properties: &armauthorization.RoleAssignmentProperties{
			PrincipalID:      to.Ptr("ce2ce14e-85d7-4629-bdbc-454d0519d987"),
			PrincipalType:    to.Ptr(armauthorization.PrincipalTypeUser),
			RoleDefinitionID: to.Ptr("/providers/Microsoft.Authorization/roleDefinitions/0b5fe924-9a61-425c-96af-cfe6e287ca2d"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RoleAssignment = armauthorization.RoleAssignment{
	// 	Name: to.Ptr("b0f43c54-e787-4862-89b1-a653fa9cf747"),
	// 	Type: to.Ptr("Microsoft.Authorization/roleAssignments"),
	// 	ID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/providers/Microsoft.Authorization/roleAssignments/b0f43c54-e787-4862-89b1-a653fa9cf747"),
	// 	Properties: &armauthorization.RoleAssignmentProperties{
	// 		PrincipalID: to.Ptr("ce2ce14e-85d7-4629-bdbc-454d0519d987"),
	// 		PrincipalType: to.Ptr(armauthorization.PrincipalTypeUser),
	// 		RoleDefinitionID: to.Ptr("/providers/Microsoft.Authorization/roleDefinitions/0b5fe924-9a61-425c-96af-cfe6e287ca2d"),
	// 		Scope: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/53b1affe357b3bfbb53721d0a2002382a046d3b0/specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/RoleAssignments_DeleteById.json
func ExampleRoleAssignmentsClient_DeleteByID() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoleAssignmentsClient().DeleteByID(ctx, "subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/providers/Microsoft.Authorization/roleAssignments/b0f43c54-e787-4862-89b1-a653fa9cf747", &armauthorization.RoleAssignmentsClientDeleteByIDOptions{TenantID: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RoleAssignment = armauthorization.RoleAssignment{
	// 	Name: to.Ptr("b0f43c54-e787-4862-89b1-a653fa9cf747"),
	// 	Type: to.Ptr("Microsoft.Authorization/roleAssignments"),
	// 	ID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/providers/Microsoft.Authorization/roleAssignments/b0f43c54-e787-4862-89b1-a653fa9cf747"),
	// 	Properties: &armauthorization.RoleAssignmentProperties{
	// 		PrincipalID: to.Ptr("ce2ce14e-85d7-4629-bdbc-454d0519d987"),
	// 		PrincipalType: to.Ptr(armauthorization.PrincipalTypeUser),
	// 		RoleDefinitionID: to.Ptr("/providers/Microsoft.Authorization/roleDefinitions/0b5fe924-9a61-425c-96af-cfe6e287ca2d"),
	// 		Scope: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2"),
	// 	},
	// }
}
