//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceRules/GetGovernanceRule_example.json
func ExampleGovernanceRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewGovernanceRulesClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "ad9a8e26-29d9-4829-bb30-e597a58cdbb8", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceRules/PutGovernanceRule_example.json
func ExampleGovernanceRulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewGovernanceRulesClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "ad9a8e26-29d9-4829-bb30-e597a58cdbb8", armsecurity.GovernanceRule{
		Properties: &armsecurity.GovernanceRuleProperties{
			Description: to.Ptr("A rule on critical recommendations"),
			ConditionSets: []interface{}{
				map[string]interface{}{
					"conditions": []interface{}{
						map[string]interface{}{
							"operator": "In",
							"property": "$.AssessmentKey",
							"value":    "[\"b1cd27e0-4ecc-4246-939f-49c426d9d72f\", \"fe83f80b-073d-4ccf-93d9-6797eb870201\"]",
						},
					},
				}},
			DisplayName: to.Ptr("Admin's rule"),
			GovernanceEmailNotification: &armsecurity.GovernanceRuleEmailNotification{
				DisableManagerEmailNotification: to.Ptr(false),
				DisableOwnerEmailNotification:   to.Ptr(false),
			},
			IsDisabled:    to.Ptr(false),
			IsGracePeriod: to.Ptr(true),
			OwnerSource: &armsecurity.GovernanceRuleOwnerSource{
				Type:  to.Ptr(armsecurity.GovernanceRuleOwnerSourceTypeManually),
				Value: to.Ptr("user@contoso.com"),
			},
			RemediationTimeframe: to.Ptr("7.00:00:00"),
			RulePriority:         to.Ptr[int32](200),
			RuleType:             to.Ptr(armsecurity.GovernanceRuleTypeIntegrated),
			SourceResourceType:   to.Ptr(armsecurity.GovernanceRuleSourceResourceTypeAssessments),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceRules/DeleteGovernanceRule_example.json
func ExampleGovernanceRulesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewGovernanceRulesClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx, "ad9a8e26-29d9-4829-bb30-e597a58cdbb8", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceRules/PostGovernanceRule_example.json
func ExampleGovernanceRulesClient_BeginRuleIDExecuteSingleSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewGovernanceRulesClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginRuleIDExecuteSingleSubscription(ctx, "ad9a8e26-29d9-4829-bb30-e597a58cdbb8", &armsecurity.GovernanceRulesClientBeginRuleIDExecuteSingleSubscriptionOptions{ExecuteGovernanceRuleParams: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceRules/PostSecurityConnectorGovernanceRule_example.json
func ExampleGovernanceRulesClient_BeginRuleIDExecuteSingleSecurityConnector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewGovernanceRulesClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginRuleIDExecuteSingleSecurityConnector(ctx, "gcpResourceGroup", "gcpconnector", "ad9a8e26-29d9-4829-bb30-e597a58cdbb8", &armsecurity.GovernanceRulesClientBeginRuleIDExecuteSingleSecurityConnectorOptions{ExecuteGovernanceRuleParams: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}