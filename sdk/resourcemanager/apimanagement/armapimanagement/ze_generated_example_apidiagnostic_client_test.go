//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListApiDiagnostics.json
func ExampleAPIDiagnosticClient_ListByService() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armapimanagement.NewAPIDiagnosticClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.ListByService("<resource-group-name>",
		"<service-name>",
		"<api-id>",
		&armapimanagement.APIDiagnosticClientListByServiceOptions{Filter: nil,
			Top:  nil,
			Skip: nil,
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadApiDiagnostic.json
func ExampleAPIDiagnosticClient_GetEntityTag() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armapimanagement.NewAPIDiagnosticClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.GetEntityTag(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<api-id>",
		"<diagnostic-id>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetApiDiagnostic.json
func ExampleAPIDiagnosticClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armapimanagement.NewAPIDiagnosticClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<api-id>",
		"<diagnostic-id>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateApiDiagnostic.json
func ExampleAPIDiagnosticClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armapimanagement.NewAPIDiagnosticClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<api-id>",
		"<diagnostic-id>",
		armapimanagement.DiagnosticContract{
			Properties: &armapimanagement.DiagnosticContractProperties{
				AlwaysLog: to.Ptr(armapimanagement.AlwaysLogAllErrors),
				Backend: &armapimanagement.PipelineDiagnosticSettings{
					Response: &armapimanagement.HTTPMessageDiagnostic{
						Body: &armapimanagement.BodyDiagnosticSettings{
							Bytes: to.Ptr[int32](512),
						},
						Headers: []*string{
							to.Ptr("Content-type")},
					},
					Request: &armapimanagement.HTTPMessageDiagnostic{
						Body: &armapimanagement.BodyDiagnosticSettings{
							Bytes: to.Ptr[int32](512),
						},
						Headers: []*string{
							to.Ptr("Content-type")},
					},
				},
				Frontend: &armapimanagement.PipelineDiagnosticSettings{
					Response: &armapimanagement.HTTPMessageDiagnostic{
						Body: &armapimanagement.BodyDiagnosticSettings{
							Bytes: to.Ptr[int32](512),
						},
						Headers: []*string{
							to.Ptr("Content-type")},
					},
					Request: &armapimanagement.HTTPMessageDiagnostic{
						Body: &armapimanagement.BodyDiagnosticSettings{
							Bytes: to.Ptr[int32](512),
						},
						Headers: []*string{
							to.Ptr("Content-type")},
					},
				},
				LoggerID: to.Ptr("<logger-id>"),
				Sampling: &armapimanagement.SamplingSettings{
					Percentage:   to.Ptr[float64](50),
					SamplingType: to.Ptr(armapimanagement.SamplingTypeFixed),
				},
			},
		},
		&armapimanagement.APIDiagnosticClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdateApiDiagnostic.json
func ExampleAPIDiagnosticClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armapimanagement.NewAPIDiagnosticClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<api-id>",
		"<diagnostic-id>",
		"<if-match>",
		armapimanagement.DiagnosticContract{
			Properties: &armapimanagement.DiagnosticContractProperties{
				AlwaysLog: to.Ptr(armapimanagement.AlwaysLogAllErrors),
				Backend: &armapimanagement.PipelineDiagnosticSettings{
					Response: &armapimanagement.HTTPMessageDiagnostic{
						Body: &armapimanagement.BodyDiagnosticSettings{
							Bytes: to.Ptr[int32](512),
						},
						Headers: []*string{
							to.Ptr("Content-type")},
					},
					Request: &armapimanagement.HTTPMessageDiagnostic{
						Body: &armapimanagement.BodyDiagnosticSettings{
							Bytes: to.Ptr[int32](512),
						},
						Headers: []*string{
							to.Ptr("Content-type")},
					},
				},
				Frontend: &armapimanagement.PipelineDiagnosticSettings{
					Response: &armapimanagement.HTTPMessageDiagnostic{
						Body: &armapimanagement.BodyDiagnosticSettings{
							Bytes: to.Ptr[int32](512),
						},
						Headers: []*string{
							to.Ptr("Content-type")},
					},
					Request: &armapimanagement.HTTPMessageDiagnostic{
						Body: &armapimanagement.BodyDiagnosticSettings{
							Bytes: to.Ptr[int32](512),
						},
						Headers: []*string{
							to.Ptr("Content-type")},
					},
				},
				LoggerID: to.Ptr("<logger-id>"),
				Sampling: &armapimanagement.SamplingSettings{
					Percentage:   to.Ptr[float64](50),
					SamplingType: to.Ptr(armapimanagement.SamplingTypeFixed),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeleteApiDiagnostic.json
func ExampleAPIDiagnosticClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armapimanagement.NewAPIDiagnosticClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<api-id>",
		"<diagnostic-id>",
		"<if-match>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}
